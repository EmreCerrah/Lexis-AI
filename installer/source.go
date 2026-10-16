package main

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type pluginMeta struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	License     string   `json:"license"`
	Keywords    []string `json:"keywords"`
	Author      struct {
		Name string `json:"name"`
	} `json:"author"`
}

// fetchSource returns the plugin directory (the one holding plugin.json and skills/).
func fetchSource(o options) (dir string, cleanup func(), err error) {
	noop := func() {}
	if o.source != "" {
		abs, err := filepath.Abs(o.source)
		if err != nil {
			return "", noop, err
		}
		d, err := locatePlugin(abs)
		if err != nil {
			return "", noop, err
		}
		fmt.Println("Kaynak (yerel):", d)
		return d, noop, nil
	}

	tmp, err := os.MkdirTemp("", "lexis-ai-*")
	if err != nil {
		return "", noop, err
	}
	cleanup = func() { os.RemoveAll(tmp) }
	fmt.Printf("Eklenti indiriliyor: github.com/%s (%s)\n", o.repo, o.ref)
	if err := downloadTarball(o.repo, o.ref, tmp); err != nil {
		cleanup()
		return "", noop, err
	}
	d, err := locatePlugin(tmp)
	if err != nil {
		cleanup()
		return "", noop, err
	}
	return d, cleanup, nil
}

func downloadTarball(repo, ref, dst string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/tarball/%s", repo, ref)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	if tok := os.Getenv("GITHUB_TOKEN"); tok != "" {
		req.Header.Set("Authorization", "Bearer "+tok)
	}
	resp, err := (&http.Client{Timeout: 2 * time.Minute}).Do(req)
	if err != nil {
		return fmt.Errorf("indirme başarısız: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("indirme başarısız: %s → HTTP %d (repo herkese açık mı? özel repo için GITHUB_TOKEN ayarla; ya da --source kullan)", url, resp.StatusCode)
	}
	return extractTarGz(resp.Body, dst)
}

func extractTarGz(r io.Reader, dst string) error {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	tr := tar.NewReader(gz)
	root := filepath.Clean(dst) + string(os.PathSeparator)
	for {
		h, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		target := filepath.Join(dst, filepath.FromSlash(h.Name))
		if !strings.HasPrefix(target, root) {
			return fmt.Errorf("arşivde güvensiz yol: %s", h.Name)
		}
		switch h.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			f, err := os.Create(target)
			if err != nil {
				return err
			}
			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return err
			}
			f.Close()
		}
		// Links and other entry types are skipped on purpose.
	}
}

// locatePlugin finds the plugin directory under root: root itself, root/plugin, or one level deeper
// (GitHub tarballs wrap everything in "<owner>-<repo>-<sha>/").
func locatePlugin(root string) (string, error) {
	isPlugin := func(d string) bool {
		_, e1 := os.Stat(filepath.Join(d, "plugin.json"))
		st, e2 := os.Stat(filepath.Join(d, "skills"))
		return e1 == nil && e2 == nil && st.IsDir()
	}
	candidates := []string{root, filepath.Join(root, "plugin")}
	if entries, err := os.ReadDir(root); err == nil {
		for _, en := range entries {
			if en.IsDir() {
				candidates = append(candidates, filepath.Join(root, en.Name()), filepath.Join(root, en.Name(), "plugin"))
			}
		}
	}
	for _, c := range candidates {
		if isPlugin(c) {
			return c, nil
		}
	}
	return "", fmt.Errorf("eklenti bulunamadı: %s altında plugin.json + skills/ olan bir klasör yok", root)
}

func readMeta(dir string) (pluginMeta, error) {
	var m pluginMeta
	b, err := os.ReadFile(filepath.Join(dir, "plugin.json"))
	if err != nil {
		return m, err
	}
	if err := json.Unmarshal(b, &m); err != nil {
		return m, fmt.Errorf("plugin.json okunamadı: %w", err)
	}
	if m.Name == "" || m.Version == "" {
		return m, fmt.Errorf("plugin.json içinde name/version eksik")
	}
	return m, nil
}
