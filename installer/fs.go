package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// markerFile marks a directory the installer owns, so a re-run may replace it safely.
const markerFile = ".lexis-ai.json"

// Files in the plugin source that are manifests for one platform only.
var platformManifests = []string{".codex-plugin", "plugin.json"}

// copyTree copies src into dst, skipping top-level names in skip.
func copyTree(src, dst string, skip []string) error {
	return filepath.WalkDir(src, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(src, p)
		if rel == "." {
			return os.MkdirAll(dst, 0o755)
		}
		top := strings.Split(filepath.ToSlash(rel), "/")[0]
		for _, s := range skip {
			if top == s {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		if top == ".git" || top == markerFile {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		return copyFile(p, target)
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, in); err != nil {
		out.Close()
		return err
	}
	return out.Close()
}

func writeJSON(path string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}

// rewriteSkillPaths replaces the "../../" plugin-root prefix used in every SKILL.md.
func rewriteSkillPaths(skillsDir, prefix string) error {
	return filepath.WalkDir(skillsDir, func(p string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(p) != ".md" {
			return err
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		return os.WriteFile(p, []byte(strings.ReplaceAll(string(b), "../../", prefix)), 0o644)
	})
}

// prepareDir empties dir for a fresh install. A non-empty dir is replaced only when the
// installer owns it (marker present) or force is set.
func prepareDir(dir string, force bool) error {
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		return nil
	}
	if _, err := os.Stat(filepath.Join(dir, markerFile)); err != nil && !force {
		return fmt.Errorf("%s zaten var ve Lexis-AI'ye ait değil (üzerine yazmak için --force)", dir)
	}
	return os.RemoveAll(dir)
}

func writeMarker(dir string, e *env, extra map[string]any) error {
	m := map[string]any{
		"plugin":      e.meta.Name,
		"version":     e.meta.Version,
		"platform":    e.opts.platform,
		"scope":       e.opts.scope,
		"installedAt": time.Now().Format(time.RFC3339),
	}
	for k, v := range extra {
		m[k] = v
	}
	return writeJSON(filepath.Join(dir, markerFile), m)
}
