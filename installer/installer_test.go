package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// newEnv builds an env over the repo's real plugin/ folder with a fake home and no AI CLIs on PATH.
func newEnv(t *testing.T, platform, scope string) *env {
	t.Helper()
	t.Setenv("PATH", "")
	t.Setenv("LOCALAPPDATA", "")
	src, err := locatePlugin(filepath.Join("..", "plugin"))
	if err != nil {
		t.Fatal(err)
	}
	meta, err := readMeta(src)
	if err != nil {
		t.Fatal(err)
	}
	o := options{platform: platform, scope: scope}
	if scope == "project" {
		o.projectDir = t.TempDir()
	}
	return &env{opts: o, home: t.TempDir(), src: src, meta: meta, runner: &runner{}}
}

func apply(t *testing.T, e *env) *plan {
	t.Helper()
	p, err := targets[e.opts.platform].plan(e)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range p.steps {
		if err := s.run(); err != nil {
			t.Fatalf("%s: %v", s.desc, err)
		}
	}
	return p
}

func read(t *testing.T, p string) string {
	t.Helper()
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func mustExist(t *testing.T, paths ...string) {
	t.Helper()
	for _, p := range paths {
		if _, err := os.Stat(p); err != nil {
			t.Errorf("missing %s", p)
		}
	}
}

func mustNotExist(t *testing.T, paths ...string) {
	t.Helper()
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			t.Errorf("should not exist: %s", p)
		}
	}
}

func TestClaudeUser(t *testing.T) {
	e := newEnv(t, "claude", "user")
	p := apply(t, e)
	root := filepath.Join(e.home, ".lexis-ai", "claude-code")
	pdir := filepath.Join(root, "lexis-ai")
	mustExist(t, filepath.Join(root, ".claude-plugin", "marketplace.json"), filepath.Join(pdir, ".claude-plugin", "plugin.json"),
		filepath.Join(pdir, "AGENTS.md"), filepath.Join(pdir, "references", "lenses", "index.md"))
	mustNotExist(t, filepath.Join(pdir, ".codex-plugin"), filepath.Join(pdir, "plugin.json"))
	skill := read(t, filepath.Join(pdir, "skills", "event", "SKILL.md"))
	if strings.Contains(skill, "../../") || !strings.Contains(skill, "${CLAUDE_PLUGIN_ROOT}/AGENTS.md") {
		t.Error("claude skill paths not rewritten")
	}
	if !strings.Contains(strings.Join(p.next, "\n"), "claude plugin install lexis-ai@lexis-ai --scope user") {
		t.Errorf("manual next steps missing: %v", p.next)
	}
}

func TestClaudeProjectUsesRelativeMarketplace(t *testing.T) {
	e := newEnv(t, "claude", "project")
	p := apply(t, e)
	mustExist(t, filepath.Join(e.opts.projectDir, ".lexis-ai", "claude-code", "lexis-ai", "skills", "close-session", "SKILL.md"))
	if !strings.Contains(strings.Join(p.next, "\n"), "marketplace add ./.lexis-ai/claude-code --scope project") {
		t.Errorf("project marketplace should be relative: %v", p.next)
	}
}

func TestCodexUserKeepsCodexManifestAndPaths(t *testing.T) {
	e := newEnv(t, "codex", "user")
	apply(t, e)
	root := filepath.Join(e.home, ".lexis-ai", "codex")
	pdir := filepath.Join(root, "plugins", "lexis-ai")
	mustExist(t, filepath.Join(pdir, ".codex-plugin", "plugin.json"), filepath.Join(pdir, "plugin.json"))
	if !strings.Contains(read(t, filepath.Join(pdir, "skills", "event", "SKILL.md")), "../../AGENTS.md") {
		t.Error("codex skills must keep ../../ paths")
	}
	var m map[string]any
	if err := json.Unmarshal([]byte(read(t, filepath.Join(root, ".agents", "plugins", "marketplace.json"))), &m); err != nil {
		t.Fatal(err)
	}
	if m["name"] != "lexis-ai" {
		t.Errorf("marketplace name = %v", m["name"])
	}
}

func TestCodexProjectMergesExistingMarketplace(t *testing.T) {
	e := newEnv(t, "codex", "project")
	mkt := filepath.Join(e.opts.projectDir, ".agents", "plugins", "marketplace.json")
	existing := `{"name":"team","plugins":[{"name":"other","source":{"source":"local","path":"./plugins/other"}}]}`
	os.MkdirAll(filepath.Dir(mkt), 0o755)
	os.WriteFile(mkt, []byte(existing), 0o644)
	p := apply(t, e)
	var m struct {
		Name    string
		Plugins []struct{ Name string }
	}
	json.Unmarshal([]byte(read(t, mkt)), &m)
	if m.Name != "team" || len(m.Plugins) != 2 || m.Plugins[0].Name != "other" || m.Plugins[1].Name != "lexis-ai" {
		t.Errorf("merge wrong: %+v", m)
	}
	if !strings.Contains(strings.Join(p.next, "\n"), "lexis-ai@team") {
		t.Errorf("selector should use existing marketplace name: %v", p.next)
	}
}

func TestGeminiManifest(t *testing.T) {
	e := newEnv(t, "gemini", "user")
	apply(t, e)
	stage := filepath.Join(e.home, ".lexis-ai", "gemini", "lexis-ai")
	var m map[string]any
	if err := json.Unmarshal([]byte(read(t, filepath.Join(stage, "gemini-extension.json"))), &m); err != nil {
		t.Fatal(err)
	}
	if m["name"] != "lexis-ai" || m["version"] == "" {
		t.Errorf("bad gemini manifest: %v", m)
	}
	mustExist(t, filepath.Join(stage, "skills", "ripples", "SKILL.md"), filepath.Join(stage, "systems", "fate-accelerated", "pack.md"))
	mustNotExist(t, filepath.Join(stage, ".codex-plugin"))
}

func TestCopilotUserAbsoluteAndProjectRelative(t *testing.T) {
	e := newEnv(t, "copilot", "user")
	apply(t, e)
	core := filepath.Join(e.home, ".copilot", "lexis-ai")
	mustExist(t, filepath.Join(core, "AGENTS.md"), filepath.Join(core, "templates", "hook.md"))
	mustNotExist(t, filepath.Join(core, "skills"))
	skill := read(t, filepath.Join(e.home, ".copilot", "skills", "event", "SKILL.md"))
	if !strings.Contains(skill, slash(core)+"/AGENTS.md") {
		t.Error("copilot user skills should point to the absolute core path")
	}
	// The rewritten path must resolve to a real file.
	mustExist(t, filepath.FromSlash(slash(core)+"/references/knowledge-model.md"))

	ep := newEnv(t, "copilot", "project")
	apply(t, ep)
	skillDir := filepath.Join(ep.opts.projectDir, ".github", "skills", "event")
	if !strings.Contains(read(t, filepath.Join(skillDir, "SKILL.md")), "../../lexis-ai/AGENTS.md") {
		t.Error("copilot project skills should use relative core path")
	}
	mustExist(t, filepath.Join(skillDir, "..", "..", "lexis-ai", "AGENTS.md"))
}

func TestCopilotRefusesForeignSkillUnlessForce(t *testing.T) {
	e := newEnv(t, "copilot", "user")
	foreign := filepath.Join(e.home, ".copilot", "skills", "event")
	os.MkdirAll(foreign, 0o755)
	os.WriteFile(filepath.Join(foreign, "SKILL.md"), []byte("someone else's"), 0o644)
	if _, err := planCopilot(e); err == nil {
		t.Fatal("expected conflict error")
	}
	e.opts.force = true
	apply(t, e)
}

func TestReinstallReplacesOwnedDirButNotForeign(t *testing.T) {
	e := newEnv(t, "codex", "user")
	apply(t, e)
	apply(t, e) // owned: marker present → replaced without force

	e2 := newEnv(t, "gemini", "user")
	stage := filepath.Join(e2.home, ".lexis-ai", "gemini", "lexis-ai")
	os.MkdirAll(stage, 0o755)
	os.WriteFile(filepath.Join(stage, "notes.md"), []byte("mine"), 0o644)
	p, err := planGemini(e2)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.steps[0].run(); err == nil {
		t.Fatal("expected refusal to overwrite a foreign directory")
	}
}

func TestPrompterFlow(t *testing.T) {
	var out bytes.Buffer
	p := newPrompter(strings.NewReader("9\n2\n\nh\n"), &out)
	if got := p.choose("Q?", []string{"a", "b", "c", "d"}, 0); got != 1 {
		t.Errorf("choose = %d, want 1 after an invalid entry", got)
	}
	if got := p.choose("Q?", []string{"a", "b"}, 0); got != 0 {
		t.Errorf("empty input should pick the default, got %d", got)
	}
	if p.confirm("ok?", true) {
		t.Error("'h' should mean no")
	}
	if !strings.Contains(out.String(), "4) d") || !strings.Contains(out.String(), "Geçersiz seçim") {
		t.Errorf("menu output unexpected:\n%s", out.String())
	}
}

func TestExtractRejectsTraversal(t *testing.T) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)
	tw.WriteHeader(&tar.Header{Name: "../evil.txt", Mode: 0o644, Size: 1, Typeflag: tar.TypeReg})
	tw.Write([]byte("x"))
	tw.Close()
	gz.Close()
	dst := t.TempDir()
	if err := extractTarGz(&buf, dst); err == nil {
		t.Fatal("expected traversal to be rejected")
	}
	mustNotExist(t, filepath.Join(filepath.Dir(dst), "evil.txt"))
}

func TestLocatePluginInTarballLayout(t *testing.T) {
	root := t.TempDir()
	plugin := filepath.Join(root, "EmreCerrah-Lexis-AI-abc123", "plugin")
	os.MkdirAll(filepath.Join(plugin, "skills"), 0o755)
	os.WriteFile(filepath.Join(plugin, "plugin.json"), []byte(`{"name":"x","version":"1"}`), 0o644)
	got, err := locatePlugin(root)
	if err != nil || got != plugin {
		t.Errorf("locatePlugin = %q, %v", got, err)
	}
}

func TestCopilotMigratesLegacyName(t *testing.T) {
	e := newEnv(t, "copilot", "user")
	legacyCore := filepath.Join(e.home, ".copilot", "worldbuilding-forge")
	os.MkdirAll(legacyCore, 0o755)
	os.WriteFile(filepath.Join(legacyCore, markerFile), []byte(`{"skills":["event","ripples"]}`), 0o644)
	for _, s := range []string{"event", "ripples"} {
		os.MkdirAll(filepath.Join(e.home, ".copilot", "skills", s), 0o755)
	}
	apply(t, e) // legacy skills count as owned: no conflict error
	mustNotExist(t, legacyCore)
	mustExist(t, filepath.Join(e.home, ".copilot", "lexis-ai", "AGENTS.md"))
}

func TestCodexMarketplaceDropsLegacyEntry(t *testing.T) {
	path := filepath.Join(t.TempDir(), "marketplace.json")
	os.WriteFile(path, []byte(`{"name":"lexis-ai","plugins":[{"name":"worldbuilding-forge"},{"name":"other"}]}`), 0o644)
	m, err := codexMarketplace(path, "lexis-ai")
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, p := range m["plugins"].([]any) {
		names = append(names, p.(map[string]any)["name"].(string))
	}
	if strings.Join(names, ",") != "other,lexis-ai" {
		t.Errorf("plugins = %v", names)
	}
}
