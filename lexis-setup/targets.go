package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const marketplaceName = "lexis-ai"

// legacyNames are plugin names earlier installer versions used; re-running removes those installs.
var legacyNames = []string{"worldbuilding-forge"}

type target struct {
	title           string
	supportsProject bool
	plan            func(e *env) (*plan, error)
}

var targets = map[string]target{
	"claude":  {"Claude Code", true, planClaude},
	"codex":   {"Codex", true, planCodex},
	"gemini":  {"Gemini CLI", true, planGemini},
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
	out := []string{"The tool was not found or does not support plugin commands. After installing/updating it, run:"}
	for _, c := range cmds {
		out = append(out, "  "+c)
	}
	return out
}

// ---------- Claude Code ----------

func planClaude(e *env) (*plan, error) {
	if e.opts.scope == "project" {
		return planClaudeProject(e)
	}
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
	p.steps = append(p.steps, step{"Plugin files → " + root, func() error {
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
			"homepage":    e.meta.Homepage,
			"repository":  e.meta.Repository,
			"license":     e.meta.License,
			"keywords":    e.meta.Keywords,
		}); err != nil {
			return err
		}
		if err := writeJSON(filepath.Join(root, ".claude-plugin", "marketplace.json"), map[string]any{
			"name":     marketplaceName,
			"owner":    map[string]string{"name": e.meta.Author.Name},
			"metadata": map[string]string{"description": "Lexis-AI plugins"},
			"plugins": []map[string]any{{
				"name":        e.meta.Name,
				"source":      "./" + e.meta.Name,
				"description": e.meta.Description,
				"version":     e.meta.Version,
			}},
		}); err != nil {
			return err
		}
		if err := setLanguage(pdir, e.opts.lang); err != nil {
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
			registered, known := claudeMarketplacePath(e.runner, bin)
			switch {
			case !known:
				// No JSON listing: fall back to the name check.
				if out, _ := e.runner.output(bin, "plugin", "marketplace", "list"); !strings.Contains(out, marketplaceName) {
					return e.runner.run(addDir, bin, addArgs...)
				}
			case registered == "":
				return e.runner.run(addDir, bin, addArgs...)
			case !samePath(registered, root):
				// Left over from an install elsewhere (e.g. a deleted project folder): point it here.
				fmt.Println("  marketplace was registered at", registered, "— re-registering it here")
				if err := e.runner.run(workDir(e), bin, "plugin", "marketplace", "remove", marketplaceName); err != nil {
					return err
				}
				return e.runner.run(addDir, bin, addArgs...)
			}
			// Claude Code reads its own copy of the marketplace; refresh it from the files just written.
			return e.runner.run(workDir(e), bin, "plugin", "marketplace", "update", marketplaceName)
		}},
		step{"Claude Code plugin: " + selector, func() error {
			installed := claudePlugins(e.runner, bin)
			for _, pl := range installed {
				if !isLegacy(strings.TrimSuffix(pl.ID, "@"+marketplaceName)) || !strings.HasSuffix(pl.ID, "@"+marketplaceName) {
					continue
				}
				fmt.Println("  removing old name:", pl.ID, "("+pl.Scope+")")
				dir := workDir(e)
				if pl.ProjectPath != "" {
					if st, err := os.Stat(pl.ProjectPath); err == nil && st.IsDir() {
						dir = pl.ProjectPath
					}
				}
				if err := e.runner.run(dir, bin, "plugin", "uninstall", pl.ID, "--scope", pl.Scope); err != nil {
					// A stale entry must not block the new install.
					fmt.Printf("  ⚠ could not remove %s; remove it later with: claude plugin uninstall %s --scope %s\n", pl.ID, pl.ID, pl.Scope)
				}
			}
			for _, pl := range installed {
				if pl.ID == selector && pl.Scope == scope {
					fmt.Println("  (already installed; updating — an unchanged version number stays cached)")
					return e.runner.run(workDir(e), bin, "plugin", "update", selector, "--scope", scope)
				}
			}
			return e.runner.run(workDir(e), bin, instArgs...)
		}},
	)
	p.next = []string{
		"Restart Claude Code (or run /reload-plugins in a session).",
		"Commands carry the plugin name: /" + e.meta.Name + ":world-init, /" + e.meta.Name + ":event …",
	}
	return p, nil
}

// ---------- Codex ----------

func planCodex(e *env) (*plan, error) {
	if e.opts.scope == "project" {
		return planCodexProject(e)
	}
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
	if bin != "" {
		// Must run before the marketplace file is rewritten: Codex only lists plugins its marketplaces still name.
		if out, _ := e.runner.output(bin, "plugin", "list"); out != "" {
			for _, old := range legacyNames {
				if oldSel := old + "@" + mkt["name"].(string); strings.Contains(out, oldSel+" ") {
					p.steps = append(p.steps, step{"Removing old name: " + oldSel, func() error {
						return e.runner.run(workDir(e), bin, "plugin", "remove", oldSel)
					}})
				}
			}
		}
	}
	p.steps = append(p.steps, step{"Plugin files → " + pdir, func() error {
		if err := prepareDir(pdir, e.opts.force); err != nil {
			return err
		}
		if err := copyTree(e.src, pdir, nil); err != nil {
			return err
		}
		if err := setLanguage(pdir, e.opts.lang); err != nil {
			return err
		}
		// Codex's own plugins ship this overlay; generate it from the portable plugin.json.
		if err := writeJSON(filepath.Join(pdir, ".codex-plugin", "plugin.json"), map[string]any{
			"name":        e.meta.Name,
			"version":     e.meta.Version,
			"description": e.meta.Description,
			"author":      map[string]string{"name": e.meta.Author.Name},
			"homepage":    e.meta.Homepage,
			"repository":  e.meta.Repository,
			"license":     e.meta.License,
			"keywords":    e.meta.Keywords,
			"skills":      "./skills/",
			"interface":   e.meta.Extensions.OpenAI.Interface,
		}); err != nil {
			return err
		}
		if err := writeMarker(pdir, e, nil); err != nil {
			return err
		}
		for _, old := range legacyNames {
			if err := removeOwned(filepath.Join(root, "plugins", old)); err != nil {
				return err
			}
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
				fmt.Println("  (already registered)")
				return nil
			}
			return e.runner.run(workDir(e), bin, addArgs...)
		}},
		// Codex installs a copy into its cache, so this also runs on re-install to refresh it.
		step{"Codex plugin: " + selector, func() error { return e.runner.run(workDir(e), bin, instArgs...) }},
	)
	p.next = []string{
		"Quit and reopen the Codex app; skills appear in a new session.",
		"Call skills with $: $world-init, $event … (or @" + e.meta.Name + ")",
	}
	return p, nil
}

// codexMarketplace loads an existing marketplace.json (keeping its name and other entries)
// or starts a new one, then points the plugin entry at ./plugins/<name>.
func codexMarketplace(path, plugin string) (map[string]any, error) {
	m := map[string]any{}
	if b, err := os.ReadFile(path); err == nil {
		if err := json.Unmarshal(b, &m); err != nil {
			return nil, fmt.Errorf("cannot read %s: %w", path, err)
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
			if xm, ok := x.(map[string]any); ok && (xm["name"] == plugin || isLegacy(xm["name"])) {
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
	if e.opts.scope == "project" {
		return planGeminiProject(e)
	}
	stage := filepath.Join(lexisBase(e), "gemini", e.meta.Name)
	installed := filepath.Join(e.home, ".gemini", "extensions", e.meta.Name)
	if _, err := os.Stat(installed); err == nil {
		if _, err := os.Stat(filepath.Join(installed, markerFile)); err != nil && !e.opts.force {
			return nil, fmt.Errorf("%s is already installed and does not belong to Lexis-AI (use --force to overwrite)", installed)
		}
	}
	bin := e.runner.findBin([]string{"gemini"})

	p := &plan{}
	p.steps = append(p.steps, step{"Plugin files → " + stage, func() error {
		if err := prepareDir(stage, e.opts.force); err != nil {
			return err
		}
		if err := copyTree(e.src, stage, platformManifests); err != nil {
			return err
		}
		if err := setLanguage(stage, e.opts.lang); err != nil {
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
	p.steps = append(p.steps, step{"Gemini CLI extension: " + e.meta.Name, func() error {
		// Some Gemini CLI builds crash on exit on Windows after finishing the work, so the
		// result is judged by the extension folder, not by the exit code.
		exists := func() bool { _, err := os.Stat(filepath.Join(installed, "gemini-extension.json")); return err == nil }
		if exists() {
			err := e.runner.run(e.home, bin, "extensions", "uninstall", e.meta.Name)
			if exists() {
				return fmt.Errorf("could not remove the previous install: %v", err)
			}
		}
		for _, old := range legacyNames {
			oldDir := filepath.Join(e.home, ".gemini", "extensions", old)
			if _, err := os.Stat(filepath.Join(oldDir, markerFile)); err == nil {
				fmt.Println("  removing old name:", old)
				e.runner.run(e.home, bin, "extensions", "uninstall", old) // checked below by folder
				if _, err := os.Stat(oldDir); err == nil {
					return fmt.Errorf("could not remove old extension %s", old)
				}
			}
			if err := removeOwned(filepath.Join(lexisBase(e), "gemini", old)); err != nil {
				return err
			}
		}
		args := []string{"extensions", "install", stage}
		if help, _ := e.runner.output(bin, "extensions", "install", "--help"); strings.Contains(help, "--consent") {
			args = append(args, "--consent")
		}
		err := e.runner.run(e.home, bin, args...)
		if !exists() {
			if err == nil {
				err = fmt.Errorf("%s not found after install", installed)
			}
			return err
		}
		if err != nil {
			fmt.Println("  (Gemini CLI exited with an error code, but the extension is installed)")
		}
		return nil
	}})
	p.next = []string{
		"Restart Gemini CLI; check with /extensions list.",
		"Skills activate from your request: \"set up a new world with world-init\", \"event: …\".",
	}
	return p, nil
}

// ---------- GitHub Copilot ----------

func planCopilot(e *env) (*plan, error) {
	base := filepath.Join(e.home, ".copilot")
	prefix := slash(filepath.Join(base, e.meta.Name)) + "/"
	if e.opts.scope == "project" {
		base = filepath.Join(e.opts.projectDir, ".github")
		prefix = "../../" + e.meta.Name + "/" // committed with the repo, so keep it relative
	}
	p, err := planSkillsInto(e, base, prefix)
	if err != nil {
		return nil, err
	}
	p.next = []string{
		"In Copilot CLI run /skills reload, then check with /skills list.",
		"Call skills with /: /world-init, /event …",
	}
	return p, nil
}

// ---------- Project installs (all tools) ----------

// Project installs copy the files straight into the folder each tool already scans for project
// skills, so nothing has to be registered with the tool.
func planClaudeProject(e *env) (*plan, error) {
	p, err := planSkillsInto(e, filepath.Join(e.opts.projectDir, ".claude"), "../../"+e.meta.Name+"/")
	if err != nil {
		return nil, err
	}
	var cleanup []step
	if s, ok := oldProjectLayoutStep(e, "claude-code"); ok {
		cleanup = append(cleanup, s)
	}
	// An older project install registered a marketplace; drop it if it points into this
	// project or at a folder that no longer exists.
	if bin := e.runner.findBin([]string{"claude"}); bin != "" {
		if registered, _ := claudeMarketplacePath(e.runner, bin); registered != "" && (isInside(registered, e.opts.projectDir) || !exists(registered)) {
			cleanup = append(cleanup, step{"Remove old Claude Code marketplace registration (" + registered + ")", func() error {
				return e.runner.run(e.home, bin, "plugin", "marketplace", "remove", marketplaceName)
			}})
		}
	}
	settings := filepath.Join(e.opts.projectDir, ".claude", "settings.json")
	if hasOldClaudeEnable(settings, e.meta.Name) {
		cleanup = append(cleanup, step{"Remove old plugin entry from " + settings, func() error {
			return dropOldClaudeEnable(settings, e.meta.Name)
		}})
	}
	p.steps = append(cleanup, p.steps...)
	p.next = []string{
		"Open Claude Code in " + e.opts.projectDir + " (restart it if it is already open).",
		"Call skills with /: /world-init, /event …",
	}
	return p, nil
}

func planCodexProject(e *env) (*plan, error) {
	proj := e.opts.projectDir
	p, err := planSkillsInto(e, filepath.Join(proj, ".agents"), "../../"+e.meta.Name+"/")
	if err != nil {
		return nil, err
	}
	// Older project installs used a repo marketplace: <project>/plugins/<name> + .agents/plugins/marketplace.json.
	var cleanup []step
	oldPlugin := filepath.Join(proj, "plugins", e.meta.Name)
	if _, err := os.Stat(filepath.Join(oldPlugin, markerFile)); err == nil {
		mktFile := filepath.Join(proj, ".agents", "plugins", "marketplace.json")
		bin := e.runner.findBin(codexCandidates(), "plugin", "marketplace", "list")
		cleanup = append(cleanup, step{"Remove old Codex project plugin (" + oldPlugin + ")", func() error {
			mktName := marketplaceName
			if m, err := codexMarketplace(mktFile, e.meta.Name); err == nil {
				mktName = m["name"].(string)
			}
			if bin != "" {
				if out, _ := e.runner.output(bin, "plugin", "list"); strings.Contains(out, e.meta.Name+"@"+mktName+" ") {
					if err := e.runner.run(proj, bin, "plugin", "remove", e.meta.Name+"@"+mktName); err != nil {
						return err
					}
				}
			}
			if err := removeOwned(oldPlugin); err != nil {
				return err
			}
			removeEmpty(filepath.Join(proj, "plugins"))
			emptied, err := dropCodexEntries(mktFile, e.meta.Name)
			if err != nil {
				return err
			}
			if emptied && bin != "" && mktName == marketplaceName {
				return e.runner.run(proj, bin, "plugin", "marketplace", "remove", marketplaceName)
			}
			return nil
		}})
	}
	p.steps = append(cleanup, p.steps...)
	p.next = []string{
		"Open Codex in " + proj + " (restart it if it is already open).",
		"Call skills with $: $world-init, $event …",
	}
	return p, nil
}

func planGeminiProject(e *env) (*plan, error) {
	p, err := planSkillsInto(e, filepath.Join(e.opts.projectDir, ".gemini"), "../../"+e.meta.Name+"/")
	if err != nil {
		return nil, err
	}
	if s, ok := oldProjectLayoutStep(e, "gemini"); ok {
		p.steps = append([]step{s}, p.steps...)
	}
	p.next = []string{
		"Start Gemini CLI in " + e.opts.projectDir + "; check with /skills list.",
		"Skills activate from your request: \"set up a new world with world-init\", \"event: …\".",
	}
	return p, nil
}

// oldProjectLayoutStep removes <project>/.lexis-ai/<sub>, where earlier versions put project installs.
func oldProjectLayoutStep(e *env, sub string) (step, bool) {
	old := filepath.Join(e.opts.projectDir, ".lexis-ai", sub)
	if _, err := os.Stat(filepath.Join(old, markerFile)); err != nil {
		return step{}, false
	}
	return step{"Remove old project layout (" + old + ")", func() error {
		if err := removeOwned(old); err != nil {
			return err
		}
		removeEmpty(filepath.Dir(old))
		return nil
	}}, true
}

// planSkillsInto copies each skill to <base>/skills/<name> and the shared files (rules, references,
// templates, system packs) to <base>/<plugin>. prefix replaces "../../" in every SKILL.md.
func planSkillsInto(e *env, base, prefix string) (*plan, error) {
	core := filepath.Join(base, e.meta.Name)
	skillsRoot := filepath.Join(base, "skills")

	names, err := skillNames(e.src)
	if err != nil {
		return nil, err
	}
	owned := map[string]bool{}
	cores := []string{core}
	for _, old := range legacyNames {
		cores = append(cores, filepath.Join(base, old))
	}
	for _, c := range cores {
		if b, err := os.ReadFile(filepath.Join(c, markerFile)); err == nil {
			var m struct{ Skills []string }
			if json.Unmarshal(b, &m) == nil {
				for _, s := range m.Skills {
					owned[s] = true
				}
			}
		}
	}
	for _, n := range names {
		if _, err := os.Stat(filepath.Join(skillsRoot, n)); err == nil && !owned[n] && !e.opts.force {
			return nil, fmt.Errorf("another skill named %s already exists: %s (use --force to overwrite)", n, filepath.Join(skillsRoot, n))
		}
	}

	p := &plan{}
	p.steps = append(p.steps,
		step{"Shared files (rules, templates, FAE) → " + core, func() error {
			if err := prepareDir(core, e.opts.force); err != nil {
				return err
			}
			for _, old := range cores[1:] {
				if err := removeOwned(old); err != nil {
					return err
				}
			}
			if err := copyTree(e.src, core, append(platformManifests, "skills")); err != nil {
				return err
			}
			if err := setLanguage(core, e.opts.lang); err != nil {
				return err
			}
			return writeMarker(core, e, map[string]any{"skills": names})
		}},
		step{fmt.Sprintf("%d skills → %s", len(names), skillsRoot), func() error {
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
	return p, nil
}

// dropCodexEntries removes this plugin (and legacy names) from a repo marketplace file.
// It deletes the file when no plugins are left; emptied reports that.
func dropCodexEntries(path, plugin string) (emptied bool, err error) {
	b, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	m := map[string]any{}
	if err := json.Unmarshal(b, &m); err != nil {
		return false, fmt.Errorf("cannot read %s: %w", path, err)
	}
	var keep []any
	if list, ok := m["plugins"].([]any); ok {
		for _, x := range list {
			if xm, ok := x.(map[string]any); ok && (xm["name"] == plugin || isLegacy(xm["name"])) {
				continue
			}
			keep = append(keep, x)
		}
	}
	if len(keep) == 0 {
		if err := os.Remove(path); err != nil {
			return false, err
		}
		removeEmpty(filepath.Dir(path))
		return true, nil
	}
	m["plugins"] = keep
	return false, writeJSON(path, m)
}

func exists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

// isInside reports whether p is dir or below it (case-insensitive, for Windows paths).
func isInside(p, dir string) bool {
	p, dir = strings.ToLower(filepath.Clean(p)), strings.ToLower(filepath.Clean(dir))
	return p == dir || strings.HasPrefix(p, dir+string(filepath.Separator))
}

// removeEmpty deletes dir only if it is empty.
func removeEmpty(dir string) {
	if entries, err := os.ReadDir(dir); err == nil && len(entries) == 0 {
		os.Remove(dir)
	}
}

func mustRel(base, target string) string {
	r, err := filepath.Rel(base, target)
	if err != nil {
		return target
	}
	return r
}

func isLegacy(name any) bool {
	for _, old := range legacyNames {
		if name == old {
			return true
		}
	}
	return false
}

// claudeMarketplacePath returns where Claude Code has the lexis-ai marketplace registered
// ("" if it isn't). known is false when this Claude Code version has no JSON listing.
func claudeMarketplacePath(r *runner, bin string) (path string, known bool) {
	out, err := r.output(bin, "plugin", "marketplace", "list", "--json")
	var list []struct {
		Name, Path, InstallLocation string
	}
	if err != nil || json.Unmarshal([]byte(jsonPart(out)), &list) != nil {
		return "", false
	}
	for _, m := range list {
		if m.Name == marketplaceName {
			if m.Path != "" {
				return m.Path, true
			}
			return m.InstallLocation, true
		}
	}
	return "", true
}

type claudePlugin struct {
	ID          string `json:"id"`
	Scope       string `json:"scope"`
	ProjectPath string `json:"projectPath"`
}

func claudePlugins(r *runner, bin string) []claudePlugin {
	out, err := r.output(bin, "plugin", "list", "--json")
	var list []claudePlugin
	if err == nil && json.Unmarshal([]byte(jsonPart(out)), &list) == nil {
		return list
	}
	// No JSON listing: assume user scope for anything the plain listing names.
	plain, _ := r.output(bin, "plugin", "list")
	for _, name := range append([]string{marketplaceName}, legacyNames...) {
		if id := name + "@" + marketplaceName; strings.Contains(plain, id) {
			list = append(list, claudePlugin{ID: id, Scope: "user"})
		}
	}
	return list
}

// jsonPart drops any log lines a CLI prints before its JSON output.
func jsonPart(s string) string {
	if i := strings.IndexAny(s, "[{"); i >= 0 {
		return s[i:]
	}
	return s
}

func samePath(a, b string) bool {
	a, b = filepath.Clean(strings.TrimPrefix(a, `\\?\`)), filepath.Clean(strings.TrimPrefix(b, `\\?\`))
	if a == b || strings.EqualFold(a, b) {
		return true
	}
	// Windows can report the same folder as a long or a short (8.3) path.
	ra, err1 := filepath.EvalSymlinks(a)
	rb, err2 := filepath.EvalSymlinks(b)
	return err1 == nil && err2 == nil && strings.EqualFold(ra, rb)
}

// oldClaudeIDs are the plugin IDs earlier project installs enabled in <project>/.claude/settings.json.
func oldClaudeIDs(plugin string) []string {
	ids := []string{plugin + "@" + marketplaceName}
	for _, old := range legacyNames {
		ids = append(ids, old+"@"+marketplaceName)
	}
	return ids
}

func readClaudeSettings(path string) (map[string]any, bool) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}
	m := map[string]any{}
	if json.Unmarshal(b, &m) != nil {
		return nil, false
	}
	return m, true
}

func hasOldClaudeEnable(path, plugin string) bool {
	m, ok := readClaudeSettings(path)
	if !ok {
		return false
	}
	enabled, _ := m["enabledPlugins"].(map[string]any)
	for _, id := range oldClaudeIDs(plugin) {
		if _, ok := enabled[id]; ok {
			return true
		}
	}
	return false
}

// dropOldClaudeEnable removes only our enabledPlugins keys; the file is deleted only if nothing else is left.
func dropOldClaudeEnable(path, plugin string) error {
	m, ok := readClaudeSettings(path)
	if !ok {
		return nil
	}
	enabled, _ := m["enabledPlugins"].(map[string]any)
	for _, id := range oldClaudeIDs(plugin) {
		delete(enabled, id)
	}
	if len(enabled) == 0 {
		delete(m, "enabledPlugins")
	}
	if len(m) == 0 {
		return os.Remove(path)
	}
	return writeJSON(path, m)
}
