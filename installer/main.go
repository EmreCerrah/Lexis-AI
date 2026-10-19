// Command lexis-setup installs the Lexis-AI plugin for the AI tool the user picks.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Overridable at build time: go build -ldflags "-X main.defaultRepo=owner/repo -X main.version=v0.1.0"
var (
	defaultRepo = "EmreCerrah/Lexis-AI"
	defaultRef  = "main"
	version     = "dev"
)

type options struct {
	platform   string
	scope      string
	projectDir string
	source     string
	repo       string
	ref        string
	yes        bool
	dryRun     bool
	force      bool
}

func main() {
	var o options
	flag.StringVar(&o.platform, "platform", "", "claude | codex | gemini | copilot (boşsa sorulur)")
	flag.StringVar(&o.scope, "scope", "", "user | project (boşsa sorulur)")
	flag.StringVar(&o.projectDir, "dir", "", "proje klasörü (scope=project için)")
	flag.StringVar(&o.source, "source", "", "GitHub yerine yerel kaynak klasörü (repo kökü ya da plugin klasörü)")
	flag.StringVar(&o.repo, "repo", defaultRepo, "eklentinin indirileceği GitHub reposu (owner/repo)")
	flag.StringVar(&o.ref, "ref", defaultRef, "dal ya da etiket")
	flag.BoolVar(&o.yes, "yes", false, "onay sormadan kur")
	flag.BoolVar(&o.dryRun, "dry-run", false, "hiçbir şey yazmadan ne yapılacağını göster")
	flag.BoolVar(&o.force, "force", false, "başka bir şeye ait hedef klasörlerin üzerine yaz")
	flag.Parse()

	if err := run(o, newPrompter(os.Stdin, os.Stdout)); err != nil {
		fmt.Fprintln(os.Stderr, "\n✖ Hata:", err)
		os.Exit(1)
	}
}

func run(o options, p *prompter) error {
	fmt.Printf("Lexis-AI kurulum (%s)\n\n", version)

	if o.platform == "" {
		i := p.choose("Hangi yapay zekâ aracını kullanıyorsun?", []string{"Claude Code", "Codex", "Gemini CLI", "GitHub Copilot"}, 0)
		o.platform = []string{"claude", "codex", "gemini", "copilot"}[i]
	}
	t, ok := targets[o.platform]
	if !ok {
		return fmt.Errorf("bilinmeyen platform %q (claude, codex, gemini, copilot)", o.platform)
	}

	if o.scope == "" {
		i := p.choose("Kurulumu nereye istiyorsun?", []string{"Kullanıcı geneli (tüm projelerde kullanılır)", "Belirli bir proje klasörü"}, 0)
		o.scope = []string{"user", "project"}[i]
	}
	if o.scope != "user" && o.scope != "project" {
		return fmt.Errorf("bilinmeyen scope %q (user, project)", o.scope)
	}
	if o.scope == "project" && !t.supportsProject {
		fmt.Printf("⚠ %s eklentileri yalnızca kullanıcı geneline kurulabiliyor; kullanıcı geneline kuruyorum.\n\n", t.title)
		o.scope = "user"
	}
	if o.scope == "project" {
		if o.projectDir == "" {
			wd, _ := os.Getwd()
			o.projectDir = p.ask("Proje klasörü", wd)
		}
		abs, err := filepath.Abs(o.projectDir)
		if err != nil {
			return err
		}
		if st, err := os.Stat(abs); err != nil || !st.IsDir() {
			return fmt.Errorf("proje klasörü bulunamadı: %s", abs)
		}
		o.projectDir = abs
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("kullanıcı klasörü bulunamadı: %w", err)
	}

	src, cleanup, err := fetchSource(o)
	if err != nil {
		return err
	}
	defer cleanup()
	meta, err := readMeta(src)
	if err != nil {
		return err
	}

	e := &env{opts: o, home: home, src: src, meta: meta, runner: &runner{dryRun: o.dryRun}}
	pl, err := t.plan(e)
	if err != nil {
		return err
	}

	fmt.Printf("\nKurulum planı — %s %s → %s (%s)\n", meta.Name, meta.Version, t.title, scopeLabel(o))
	for _, s := range pl.steps {
		fmt.Println("  •", s.desc)
	}
	if o.dryRun {
		fmt.Println("\n(dry-run: hiçbir şey yazılmadı)")
		return nil
	}
	if !o.yes && !p.confirm("\nDevam edilsin mi?", true) {
		fmt.Println("İptal edildi.")
		return nil
	}

	fmt.Println()
	for _, s := range pl.steps {
		fmt.Println("→", s.desc)
		if err := s.run(); err != nil {
			return fmt.Errorf("%s: %w", s.desc, err)
		}
	}

	fmt.Printf("\n✅ %s kuruldu.\n", meta.Name)
	if len(pl.next) > 0 {
		fmt.Println("\nSıradaki adımlar:")
		for _, n := range pl.next {
			fmt.Println("  " + n)
		}
	}
	return nil
}

func scopeLabel(o options) string {
	if o.scope == "project" {
		return "proje: " + o.projectDir
	}
	return "kullanıcı geneli"
}

// slash keeps paths readable inside Markdown on every OS.
func slash(p string) string { return strings.ReplaceAll(p, `\`, "/") }
