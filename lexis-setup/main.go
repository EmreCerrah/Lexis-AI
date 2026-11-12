// Command lexis-setup installs the Lexis-AI plugin for the AI tool the user picks.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

// Overridable at build time: go build -ldflags "-X main.defaultRepo=owner/repo -X main.version=v0.1.0"
var (
	defaultRepo = "EmreCerrah/Lexis-AI"
	defaultRef  = "main"
	version     = "dev"
)

// Offered in the language menu; any other language can be typed in.
var languages = []string{"English", "Türkçe (Turkish)", "Español (Spanish)", "Deutsch (German)", "Français (French)", "Português (Portuguese)"}

type options struct {
	platform   string
	scope      string
	projectDir string
	lang       string
	systems    string
	source     string
	repo       string
	ref        string
	yes        bool
	dryRun     bool
	force      bool
}

func main() {
	var o options
	flag.StringVar(&o.platform, "platform", "", "claude | codex | gemini | copilot (asked if empty)")
	flag.StringVar(&o.scope, "scope", "", "user | project (asked if empty)")
	flag.StringVar(&o.projectDir, "dir", "", "project folder (for scope=project)")
	flag.StringVar(&o.lang, "lang", "", "language the AI answers in, e.g. English, Turkish (asked if empty)")
	flag.StringVar(&o.systems, "system", "", "game system pack id, several separated by commas, or all (asked if empty)")
	flag.StringVar(&o.source, "source", "", "local source folder instead of GitHub (repo root or plugin folder)")
	flag.StringVar(&o.repo, "repo", defaultRepo, "GitHub repository to download the plugin from (owner/repo)")
	flag.StringVar(&o.ref, "ref", defaultRef, "branch or tag")
	flag.BoolVar(&o.yes, "yes", false, "install without asking for confirmation")
	flag.BoolVar(&o.dryRun, "dry-run", false, "show what would be done without writing anything")
	flag.BoolVar(&o.force, "force", false, "overwrite target folders that belong to something else")
	flag.Parse()

	if err := run(o, newPrompter(os.Stdin, os.Stdout)); err != nil {
		fmt.Fprintln(os.Stderr, "\n✖ Error:", err)
		os.Exit(1)
	}
}

func run(o options, p *prompter) error {
	fmt.Printf("Lexis-AI setup (%s)\n\n", buildVersion())

	if o.platform == "" {
		i := p.choose("Which AI tool do you use?", []string{"Claude Code", "Codex", "Gemini CLI", "GitHub Copilot"}, 0)
		o.platform = []string{"claude", "codex", "gemini", "copilot"}[i]
	}
	t, ok := targets[o.platform]
	if !ok {
		return fmt.Errorf("unknown platform %q (claude, codex, gemini, copilot)", o.platform)
	}

	if o.scope == "" {
		i := p.choose("Where do you want to install it?", []string{"For my user (available in every project)", "In a specific project folder"}, 0)
		o.scope = []string{"user", "project"}[i]
	}
	if o.scope != "user" && o.scope != "project" {
		return fmt.Errorf("unknown scope %q (user, project)", o.scope)
	}
	if o.scope == "project" && !t.supportsProject {
		fmt.Printf("⚠ %s extensions can only be installed for the user; installing for the user.\n\n", t.title)
		o.scope = "user"
	}
	if o.scope == "project" {
		if o.projectDir == "" {
			wd, _ := os.Getwd()
			o.projectDir = p.ask("Project folder", wd)
		}
		abs, err := filepath.Abs(o.projectDir)
		if err != nil {
			return err
		}
		if st, err := os.Stat(abs); err != nil || !st.IsDir() {
			return fmt.Errorf("project folder not found: %s", abs)
		}
		o.projectDir = abs
	}

	if o.lang == "" {
		opts := append(append([]string{}, languages...), "Other (type it)")
		i := p.choose("Which language should the AI answer in?", opts, 0)
		if i == len(languages) {
			o.lang = p.ask("Language", "English")
		} else {
			o.lang = languages[i]
		}
	}
	o.lang = strings.TrimSpace(o.lang)

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("home folder not found: %w", err)
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

	packs, err := listPacks(src)
	if err != nil {
		return err
	}
	chosen, err := selectPacks(o.systems, packs, p)
	if err != nil {
		return err
	}
	src, cleanStage, err := stageSource(src, chosen)
	if err != nil {
		return err
	}
	defer cleanStage()

	e := &env{opts: o, home: home, src: src, meta: meta, runner: &runner{dryRun: o.dryRun}}
	pl, err := t.plan(e)
	if err != nil {
		return err
	}

	fmt.Printf("\nInstall plan — %s %s → %s (%s, language: %s, game system: %s)\n", meta.Name, meta.Version, t.title, scopeLabel(o), o.lang, packNames(chosen))
	for _, s := range pl.steps {
		fmt.Println("  •", s.desc)
	}
	if o.dryRun {
		fmt.Println("\n(dry run: nothing was written)")
		return nil
	}
	if !o.yes && !p.confirm("\nContinue?", true) {
		fmt.Println("Cancelled.")
		return nil
	}

	fmt.Println()
	for _, s := range pl.steps {
		fmt.Println("→", s.desc)
		if err := s.run(); err != nil {
			return fmt.Errorf("%s: %w", s.desc, err)
		}
	}

	fmt.Printf("\n✅ %s installed.\n", meta.Name)
	if len(pl.next) > 0 {
		fmt.Println("\nNext steps:")
		for _, n := range pl.next {
			fmt.Println("  " + n)
		}
	}
	return nil
}

func scopeLabel(o options) string {
	if o.scope == "project" {
		return "project: " + o.projectDir
	}
	return "user"
}

// buildVersion prefers the -ldflags version, then the module version `go install` records.
func buildVersion() string {
	if version != "dev" {
		return version
	}
	if bi, ok := debug.ReadBuildInfo(); ok && bi.Main.Version != "" && bi.Main.Version != "(devel)" {
		return bi.Main.Version
	}
	return version
}

// slash keeps paths readable inside Markdown on every OS.
func slash(p string) string { return strings.ReplaceAll(p, `\`, "/") }
