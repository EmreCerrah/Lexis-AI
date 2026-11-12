package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// The system pack contract (docs/design/system-packs.md): every pack has exactly these files.
var packFiles = []string{
	"pack.md",
	"resolution.md",
	"scene-tags.md",
	"obstacles.md",
	"entities.md",
	"init-questions.md",
	"templates/system.md",
}

// Glossary roles every pack.md must translate.
var packRoles = []string{"check", "trait", "difficulty", "scene-tag", "special-ability", "harm", "meta-currency"}

// Sections every pack.md must have.
var packSections = []string{"## License", "## Content scope", "## Glossary", "## Never"}

// Setting and brand terms that must never appear in a pack (outside its ## License section).
// Packs are mechanics only: no published setting content and no brand names.
var packDenylist = []string{
	// brands and publishers' game lines
	"dungeons & dragons", "d&d", "pathfinder", "starfinder", "call of cthulhu", "cthulhu",
	"world of darkness", "chronicles of darkness", "vampire: the masquerade", "werewolf: the apocalypse",
	"mörk borg", "mork borg", "mutants & masterminds", "mutants and masterminds",
	"blades in the dark",
	// settings, places and named characters or creatures
	"forgotten realms", "faerûn", "faerun", "greyhawk", "eberron", "ravenloft", "waterdeep", "baldur's gate",
	"drizzt", "elminster", "vecna", "tiamat", "lolth", "mordenkainen", "tasha", "xanathar",
	"mind flayer", "illithid", "beholder", "displacer beast",
	"golarion", "absalom", "varisia",
	"arkham", "innsmouth", "miskatonic", "dunwich", "nyarlathotep", "azathoth", "yog-sothoth",
	"camarilla", "sabbat", "anarch",
	"doskvol", "duskwall",
	"galgenbeck", "freedom city",
}

func TestSystemPacksContract(t *testing.T) {
	root := filepath.Join("..", "plugin", "systems")
	entries, err := os.ReadDir(root)
	if err != nil {
		t.Fatal(err)
	}
	var ids []string
	for _, e := range entries {
		if e.IsDir() {
			ids = append(ids, e.Name())
		}
	}
	if len(ids) < 2 {
		t.Fatalf("expected at least 2 system packs, found %v", ids)
	}
	for _, id := range ids {
		t.Run(id, func(t *testing.T) { checkPack(t, filepath.Join(root, id), id) })
	}
}

func checkPack(t *testing.T, dir, id string) {
	want := map[string]bool{}
	for _, f := range packFiles {
		want[f] = true
	}

	// Only contract files, and all of them.
	var got []string
	err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(dir, p)
		got = append(got, filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(got)
	for _, f := range got {
		if !want[f] {
			t.Errorf("extra file %s: a pack may only hold the contract files", f)
		}
		delete(want, f)
	}
	for f := range want {
		t.Errorf("missing contract file %s", f)
	}

	// pack.md: id, sections, glossary.
	pack := readPackFile(t, filepath.Join(dir, "pack.md"))
	if v := frontmatter(pack, "id"); v != id {
		t.Errorf("pack.md id = %q, want the folder name %q", v, id)
	}
	if frontmatter(pack, "name") == "" {
		t.Error("pack.md has no name")
	}
	for _, s := range packSections {
		if !regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(s) + `\s*$`).MatchString(pack) {
			t.Errorf("pack.md has no %s section", s)
		}
	}
	if strings.TrimSpace(section(pack, "## License")) == "" {
		t.Error("pack.md ## License is empty: name the source, the license and the attribution")
	}
	glossary := section(pack, "## Glossary")
	for _, r := range packRoles {
		if !regexp.MustCompile(`(?m)^\|\s*` + regexp.QuoteMeta(r) + `\s*\|\s*\S`).MatchString(glossary) {
			t.Errorf("pack.md glossary has no row for role %q", r)
		}
	}

	// templates/system.md names this pack.
	if v := frontmatter(readPackFile(t, filepath.Join(dir, "templates", "system.md")), "system"); v != id {
		t.Errorf("templates/system.md system = %q, want %q", v, id)
	}

	// Content rule: no setting or brand terms outside pack.md's ## License.
	for _, f := range got {
		text := readPackFile(t, filepath.Join(dir, filepath.FromSlash(f)))
		if f == "pack.md" {
			text = strings.Replace(text, section(text, "## License"), "", 1)
		}
		lower := strings.ToLower(text)
		for _, term := range packDenylist {
			if strings.Contains(lower, term) {
				t.Errorf("%s contains %q: packs are mechanics only, no setting content or brand names", f, term)
			}
		}
	}
}

func readPackFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return strings.ReplaceAll(string(b), "\r\n", "\n")
}

// frontmatter returns the value of key in the leading --- block.
func frontmatter(text, key string) string {
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

// section returns the body under a "## " heading, up to the next "## " heading.
func section(text, heading string) string {
	i := strings.Index(text, "\n"+heading+"\n")
	if i < 0 {
		return ""
	}
	body := text[i+len(heading)+2:]
	if j := strings.Index(body, "\n## "); j >= 0 {
		body = body[:j]
	}
	return body
}
