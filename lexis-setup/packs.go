package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// pack is one game system pack under plugin/systems/<id>/.
type pack struct {
	id, name string
	order    int
}

// listPacks reads every systems/*/pack.md in the plugin source, in menu order.
func listPacks(src string) ([]pack, error) {
	entries, err := os.ReadDir(filepath.Join(src, "systems"))
	if err != nil {
		return nil, fmt.Errorf("no system packs in the plugin: %w", err)
	}
	var packs []pack
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		b, err := os.ReadFile(filepath.Join(src, "systems", e.Name(), "pack.md"))
		if err != nil {
			continue
		}
		text := strings.ReplaceAll(string(b), "\r\n", "\n")
		p := pack{id: e.Name(), name: packField(text, "name"), order: 1000}
		if p.name == "" {
			p.name = p.id
		}
		if n, err := strconv.Atoi(packField(text, "order")); err == nil {
			p.order = n
		}
		packs = append(packs, p)
	}
	if len(packs) == 0 {
		return nil, fmt.Errorf("no system packs in the plugin")
	}
	sort.SliceStable(packs, func(i, j int) bool {
		if packs[i].order != packs[j].order {
			return packs[i].order < packs[j].order
		}
		return packs[i].id < packs[j].id
	})
	return packs, nil
}

// packField returns a frontmatter value of pack.md.
func packField(text, key string) string {
	if !strings.HasPrefix(text, "---\n") {
		return ""
	}
	end := strings.Index(text[4:], "\n---")
	if end < 0 {
		return ""
	}
	for _, line := range strings.Split(text[4:4+end], "\n") {
		if k, v, ok := strings.Cut(line, ":"); ok && strings.TrimSpace(k) == key {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

// selectPacks resolves the -system flag ("all", or ids separated by commas) or asks.
func selectPacks(flagValue string, packs []pack, p *prompter) ([]pack, error) {
	if flagValue == "" {
		opts := make([]string, 0, len(packs)+1)
		for _, pk := range packs {
			opts = append(opts, pk.name)
		}
		opts = append(opts, "All systems")
		i := p.choose("Which game system do you play?", opts, 0)
		if i == len(packs) {
			return packs, nil
		}
		return packs[i : i+1], nil
	}
	if strings.EqualFold(strings.TrimSpace(flagValue), "all") {
		return packs, nil
	}
	var out []pack
	seen := map[string]bool{}
	for _, id := range strings.Split(flagValue, ",") {
		id = strings.TrimSpace(id)
		if id == "" || seen[id] {
			continue
		}
		found := false
		for _, pk := range packs {
			if pk.id == id {
				out, found = append(out, pk), true
				seen[id] = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("unknown game system %q (available: %s, all)", id, packIDs(packs))
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no game system given (available: %s, all)", packIDs(packs))
	}
	return out, nil
}

func packIDs(packs []pack) string {
	ids := make([]string, len(packs))
	for i, pk := range packs {
		ids[i] = pk.id
	}
	return strings.Join(ids, ", ")
}

func packNames(packs []pack) string {
	names := make([]string, len(packs))
	for i, pk := range packs {
		names[i] = pk.name
	}
	return strings.Join(names, ", ")
}

// stageSource copies the plugin into a temporary folder with only the chosen
// system packs, so the install never writes the other systems' files.
func stageSource(src string, chosen []pack) (string, func(), error) {
	tmp, err := os.MkdirTemp("", "lexis-ai-stage-")
	if err != nil {
		return "", nil, err
	}
	cleanup := func() { os.RemoveAll(tmp) }
	if err := copyTree(src, tmp, []string{"systems"}); err != nil {
		cleanup()
		return "", nil, err
	}
	for _, pk := range chosen {
		if err := copyTree(filepath.Join(src, "systems", pk.id), filepath.Join(tmp, "systems", pk.id), nil); err != nil {
			cleanup()
			return "", nil, err
		}
	}
	return tmp, cleanup, nil
}
