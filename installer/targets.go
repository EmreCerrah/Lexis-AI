package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const marketplaceName = "lexis-ai"

type target struct {
	title           string
	supportsProject bool
	plan            func(e *env) (*plan, error)
}

var targets = map[string]target{
	"claude":  {"Claude Code", true, planClaude},
	"codex":   {"Codex", true, planCodex},
	"gemini":  {"Gemini CLI", false, planGemini},
	"copilot": {"GitHub Copilot", true, planCopilot},
}

// lexisBase is where installer-owned copies live: ~/.lexis-ai or <project>/.lexis-ai.
func lexisBase(e *env) string {
	if e.opts.scope == "project" {
		return filepath.Join(e.opts.projectDir, ".lexis-ai")
	}
	return filepath.Join(e.home, ".lexis-ai")
}

func workDir(e *env) string {
	if e.opts.scope == "project" {
		return e.opts.projectDir
	}
	return e.home
}

func skillNames(src string) ([]string, error) {
	entries, err := os.ReadDir(filepath.Join(src, "skills"))
	if err != nil {
		return nil, err
	}
	var names []string
	for _, en := range entries {
		if _, err := os.Stat(filepath.Join(src, "skills", en.Name(), "SKILL.md")); en.IsDir() && err == nil {
			names = append(names, en.Name())
		}
	}
	return names, nil
}

func manualNext(cmds ...string) []string {
	out := []string{"Araç bulunamadı ya da eklenti komutlarını desteklemiyor. Kurduktan/güncelledikten sonra şunları çalıştır:"}
	for _, c := range cmds {
		out = append(out, "  "+c)
	}
	return out
}

// ---------- Claude Code ----------

func planClaude(e *env) (*plan, error) {
	root := filepath.Join(lexisBase(e), "claude-code")
	pdir := filepath.Join(root, e.meta.Name)
	scope := e.opts.scope
	mktSource := root
	if scope == "project" {
		mktSource = "./" + slash(mustRel(e.opts.projectDir, root)) // keeps the project's settings portable
	}
	selector := e.meta.Name + "@" + marketplaceName
	bin := e.runner.findBin([]string{"claude"})

	p := &plan{}
	p.steps = append(p.steps, step{"Eklenti dosyaları → " + root, func() error {
		if err := prepareDir(root, e.opts.force); err != nil {
			return err
		}
		if err := copyTree(e.src, pdir, platformManifests); err != nil {
			return err
		}
		if err := writeJSON(filepath.Join(pdir, ".claude-plugin", "plugin.json"), map[string]any{
			"name":        e.meta.Name,
			"version":     e.meta.Version,
			"description": e.meta.Description,
			"author":      map[string]string{"name": e.meta.Author.Name},
			"license":     e.meta.License,
			"keywords":    e.meta.Keywords,
		}); err != nil {
			return err
		}
		if err := writeJSON(filepath.Join(root, ".claude-plugin", "marketplace.json"), map[string]any{
			"name":     marketplaceName,
			"owner":    map[string]string{"name": e.meta.Author.Name},
			"metadata": map[string]string{"description": "Lexis-AI eklentileri"},
			"plugins": []map[string]any{{
				"name":        e.meta.Name,
				"source":      "./" + e.meta.Name,
				"description": e.meta.Description,
				"version":     e.meta.Version,
			}},
		}); err != nil {
			return err
		}
		// Claude Code substitutes ${CLAUDE_PLUGIN_ROOT} in skill bodies with the installed plugin path.
		if err := rewriteSkillPaths(filepath.Join(pdir, "skills"), "${CLAUDE_PLUGIN_ROOT}/"); err != nil {
			return err
		}
		return writeMarker(root, e, nil)
	}})

	addArgs := []string{"plugin", "marketplace", "add", mktSource, "--scope", scope}
	addDir := workDir(e)
	instArgs := []string{"plugin", "install", selector, "--scope", scope}
	if bin != "" {
		// Before marketplace scopes existed (e.g. 2.1.45) marketplaces were always user-level.
		if help, _ := e.runner.output(bin, "plugin", "marketplace", "add", "--help"); !strings.Contains(help, "--scope") {
			// Those versions also reject drive-letter paths, so add "./" from inside the marketplace folder.
			addArgs, addDir = []string{"plugin", "marketplace", "add", "./"}, root
		}
	}
	if bin == "" {
		p.next = manualNext(quoteCmd("claude", addArgs...), quoteCmd("claude", instArgs...))
		return p, nil
	}
	p.steps = append(p.steps,
		step{"Claude Code marketplace: " + marketplaceName, func() error {
			if out, _ := e.runner.output(bin, "plugin", "marketplace", "list"); strings.Contains(out, marketplaceName) {
				fmt.Println("  (zaten kayıtlı; dosyalar yerinde güncellendi)")
				return nil
			}
			return e.runner.run(addDir, bin, addArgs...)
		}},
		step{"Claude Code eklentisi: " + selector, func() error {
			if out, _ := e.runner.output(bin, "plugin", "list"); strings.Contains(out, selector) {
				// Claude Code keeps a per-version copy in its cache; refresh it from the marketplace.
				fmt.Println("  (zaten kurulu; güncelleniyor — aynı sürüm numarası önbellekte kalır)")
				if err := e.runner.run(workDir(e), bin, "plugin", "marketplace", "update", marketplaceName); err != nil {
					return err
				}
				return e.runner.run(workDir(e), bin, "plugin", "update", selector)
			}
			return e.runner.run(workDir(e), bin, instArgs...)
		}},
	)
	p.next = []string{
		"Claude Code'u yeniden başlat (ya da oturumda /reload-plugins).",
		"Komutlar eklenti adıyla gelir: /worldbuilding-forge:world-init, /worldbuilding-forge:event …",
	}
	return p, nil
}

// ---------- Codex ----------

func planCodex(e *env) (*plan, error) {
	var root string
	if e.opts.scope == "project" {
		root = e.opts.projectDir // repo marketplace: <repo>/.agents/plugins + <repo>/plugins/<name>
	} else {
		root = filepath.Join(lexisBase(e), "codex")
	}
	pdir := filepath.Join(root, "plugins", e.meta.Name)
	mktFile := filepath.Join(root, ".agents", "plugins", "marketplace.json")
	mkt, err := codexMarketplace(mktFile, e.meta.Name)
	if err != nil {
		return nil, err
	}
	selector := e.meta.Name + "@" + mkt["name"].(string)
	// Old builds print help for `plugin --help` (they read "plugin" as a prompt), so probe a real subcommand.
	bin := e.runner.findBin(codexCandidates(), "plugin", "marketplace", "list")

	p := &plan{}
	p.steps = append(p.steps, step{"Eklenti dosyaları → " + pdir, func() error {
		if err := prepareDir(pdir, e.opts.force); err != nil {
			return err
		}
		if err := copyTree(e.src, pdir, nil); err != nil {
			return err
		}
		if err := writeMarker(pdir, e, nil); err != nil {
			return err
		}
		return writeJSON(mktFile, mkt)
	}})

	addArgs := []string{"plugin", "marketplace", "add", root}
	instArgs := []string{"plugin", "add", selector}
	if bin == "" {
		p.next = manualNext(quoteCmd("codex", addArgs...), quoteCmd("codex", instArgs...))
		return p, nil
	}
	p.steps = append(p.steps,
		step{"Codex marketplace: " + mkt["name"].(string), func() error {
			if out, _ := e.runner.output(bin, "plugin", "marketplace", "list"); strings.Contains(out, mkt["name"].(string)) {
				fmt.Println("  (zaten kayıtlı)")
				return nil
			}
			return e.runner.run(workDir(e), bin, addArgs...)
		}},
		// Codex installs a copy into its cache, so this also runs on re-install to refresh it.
		step{"Codex eklentisi: " + selector, func() error { return e.runner.run(workDir(e), bin, instArgs...) }},
	)
	p.next = []string{
		"Codex uygulamasını tamamen kapatıp aç; skill'ler yeni oturumda görünür.",
		"Skill'leri $ ile çağır: $world-init, $event … (ya da @" + e.meta.Name + ")",
	}
	return p, nil
}

// codexMarketplace loads an existing marketplace.json (keeping its name and other entries)
// or starts a new one, then points the plugin entry at ./plugins/<name>.
func codexMarketplace(path, plugin string) (map[string]any, error) {
	m := map[string]any{}
	if b, err := os.ReadFile(path); err == nil {
		if err := json.Unmarshal(b, &m); err != nil {
			return nil, fmt.Errorf("%s okunamadı: %w", path, err)
		}
	}
	if _, ok := m["name"].(string); !ok {
		m["name"] = marketplaceName
		m["interface"] = map[string]any{"displayName": "Lexis-AI"}
	}
	entry := map[string]any{
		"name":     plugin,
		"source":   map[string]any{"source": "local", "path": "./plugins/" + plugin},
		"policy":   map[string]any{"installation": "AVAILABLE", "authentication": "ON_USE"},
		"category": "Productivity",
	}
	var plugins []any
	if existing, ok := m["plugins"].([]any); ok {
		for _, x := range existing {
			if xm, ok := x.(map[string]any); ok && xm["name"] == plugin {
				continue
			}
			plugins = append(plugins, x)
		}
	}
	m["plugins"] = append(plugins, entry)
	return m, nil
}

// ---------- Gemini CLI ----------

func planGemini(e *env) (*plan, error) {
	stage := filepath.Join(lexisBase(e), "gemini", e.meta.Name)
	installed := filepath.Join(e.home, ".gemini", "extensions", e.meta.Name)
	if _, err := os.Stat(installed); err == nil {
		if _, err := os.Stat(filepath.Join(installed, markerFile)); err != nil && !e.opts.force {
			return nil, fmt.Errorf("%s zaten kurulu ve Lexis-AI'ye ait değil (üzerine yazmak için --force)", installed)
		}
	}
	bin := e.runner.findBin([]string{"gemini"})

	p := &plan{}
	p.steps = append(p.steps, step{"Eklenti dosyaları → " + stage, func() error {
		if err := prepareDir(stage, e.opts.force); err != nil {
			return err
		}
		if err := copyTree(e.src, stage, platformManifests); err != nil {
			return err
		}
		if err := writeJSON(filepath.Join(stage, "gemini-extension.json"), map[string]any{
			"name":        e.meta.Name,
			"version":     e.meta.Version,
			"description": e.meta.Description,
		}); err != nil {
			return err
		}
		return writeMarker(stage, e, nil)
	}})

	if bin == "" {
		p.next = manualNext(quoteCmd("gemini", "extensions", "install", stage))
		return p, nil
	}
	p.steps = append(p.steps, step{"Gemini CLI eklentisi: " + e.meta.Name, func() error {
		// Some Gemini CLI builds crash on exit on Windows after finishing the work, so the
		// result is judged by the extension folder, not by the exit code.
		exists := func() bool { _, err := os.Stat(filepath.Join(installed, "gemini-extension.json")); return err == nil }
		if exists() {
			err := e.runner.run(e.home, bin, "extensions", "uninstall", e.meta.Name)
			if exists() {
				return fmt.Errorf("eski kurulum kaldırılamadı: %v", err)
			}
		}
		args := []string{"extensions", "install", stage}
		if help, _ := e.runner.output(bin, "extensions", "install", "--help"); strings.Contains(help, "--consent") {
			args = append(args, "--consent")
		}
		err := e.runner.run(e.home, bin, args...)
		if !exists() {
			if err == nil {
				err = fmt.Errorf("kurulum sonrası %s bulunamadı", installed)
			}
			return err
		}
		if err != nil {
			fmt.Println("  (Gemini CLI hata koduyla çıktı ama eklenti kuruldu)")
		}
		return nil
	}})
	p.next = []string{
		"Gemini CLI'yi yeniden başlat; /extensions list ile kontrol et.",
		"Skill'ler isteğine göre devreye girer: \"world-init ile yeni dünya kur\", \"event: …\".",
	}
	return p, nil
}

// ---------- GitHub Copilot ----------

func planCopilot(e *env) (*plan, error) {
	var base, prefix string
	if e.opts.scope == "project" {
		base = filepath.Join(e.opts.projectDir, ".github")
		prefix = "../../" + e.meta.Name + "/" // committed with the repo, so keep it relative
	} else {
		base = filepath.Join(e.home, ".copilot")
	}
	core := filepath.Join(base, e.meta.Name)
	if prefix == "" {
		prefix = slash(core) + "/"
	}
	skillsRoot := filepath.Join(base, "skills")

	names, err := skillNames(e.src)
	if err != nil {
		return nil, err
	}
	owned := map[string]bool{}
	if b, err := os.ReadFile(filepath.Join(core, markerFile)); err == nil {
		var m struct{ Skills []string }
		if json.Unmarshal(b, &m) == nil {
			for _, s := range m.Skills {
				owned[s] = true
			}
		}
	}
	for _, n := range names {
		if _, err := os.Stat(filepath.Join(skillsRoot, n)); err == nil && !owned[n] && !e.opts.force {
			return nil, fmt.Errorf("%s adında başka bir skill zaten var: %s (üzerine yazmak için --force)", n, filepath.Join(skillsRoot, n))
		}
	}

	p := &plan{}
	p.steps = append(p.steps,
		step{"Ortak dosyalar (kurallar, şablonlar, FAE) → " + core, func() error {
			if err := prepareDir(core, e.opts.force); err != nil {
				return err
			}
			if err := copyTree(e.src, core, append(platformManifests, "skills")); err != nil {
				return err
			}
			return writeMarker(core, e, map[string]any{"skills": names})
		}},
		step{fmt.Sprintf("%d skill → %s", len(names), skillsRoot), func() error {
			for _, n := range names {
				dst := filepath.Join(skillsRoot, n)
				if err := os.RemoveAll(dst); err != nil {
					return err
				}
				if err := copyTree(filepath.Join(e.src, "skills", n), dst, nil); err != nil {
					return err
				}
				if err := rewriteSkillPaths(dst, prefix); err != nil {
					return err
				}
			}
			return nil
		}},
	)
	p.next = []string{
		"Copilot CLI'de /skills reload, sonra /skills list ile kontrol et.",
		"Skill'leri / ile çağır: /world-init, /event …",
	}
	return p, nil
}

func mustRel(base, target string) string {
	r, err := filepath.Rel(base, target)
	if err != nil {
		return target
	}
	return r
}
