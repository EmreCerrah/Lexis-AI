# Creating a System Pack

A system pack teaches Lexis-AI one game's rules. The commands never name a system; they read the pack's files by role. Adding a system means adding one folder under `plugin/systems/` and nothing else: `lexis-setup` lists it in the game system menu automatically.

The full design is in [design/system-packs.md](design/system-packs.md).

## 1. Check the license first
A pack may only contain rules you are allowed to restate:
- **Open licenses** (CC-BY, ORC, OGL 1.0a, a publisher's third-party license that allows it): fine, with the required attribution or notice.
- **Closed rules** (no open license, or a non-commercial fan license only): no pack.

Read the license on the publisher's own page and note exactly what it requires (attribution text, notices, share-alike, trademark limits).

## 2. Content rule: mechanics only
**Allowed:** how rolls work, difficulty scales, outcome bands, character math (scores, modifiers, hit points, derived values, advancement), generic conditions and harm, generic GM procedures, setting-free examples you wrote yourself, the required license text.

**Never:**
- setting content: worlds, places, gods, factions, history, named characters or creatures, adventures, story;
- spell, power, item or monster lists copied wholesale; art; layout;
- trademarks as names. Name the pack after its mechanics (`srd5`, `d100`, `fitd`); use a brand only where its owner allows it, never as a claim of endorsement;
- rule text from games that are not open-licensed.

Write the rules **in your own short words**, even when the license would allow copying. A pack is a GM's cheat sheet, not a copy of the rulebook.

## 3. The files
```
plugin/systems/<id>/
  pack.md              # identity, license, content scope, glossary, seam, never
  resolution.md        # how a roll works + the answer shape for /help-skills
  scene-tags.md        # what a scene tag is in this system + its output block
  obstacles.md         # difficulty, opposition, hazards, costs + the obstacle output block
  entities.md          # system sections of character, location and group files; character math
  init-questions.md    # the system's questions for /world-init
  templates/system.md  # the world's system.md for this system
```
No other files are allowed in a pack folder.

### pack.md
```markdown
---
type: system-pack
id: <id, same as the folder name>
name: <menu name>
order: <position in the install menu>
version: 0.1.0
---

# System Pack: <name>
<what it is, when to use it>

## License
<attribution / notices exactly as the license requires>

## Content scope
Mechanics only … No setting content …

## Glossary
| Role | Term in this system |
|---|---|
| check | … |
| trait | … |
| difficulty | … |
| scene-tag | … |
| special-ability | … |
| harm | … |
| meta-currency | … |

## Files
## Seam (core → pack)
## Never
```
Every glossary role must have a row (write `—` if the system has nothing for it).

### resolution.md
How a roll works, the difficulty scale, what results mean, the traits (skills, abilities, approaches, actions) and a section **`## Answer shape for /help-skills`** with the steps and a fenced output example.

### scene-tags.md
What a scene tag is in this system (an aspect, a scene feature, a modifier, a clock …), the kinds with their rule effects, and an output block with **Effect**, **Use it** and **Against them** lines.

### obstacles.md
How to suggest a difficulty, passive vs active opposition, obstacle types (barrier, puzzle, hazard or trap, social, multi-step), a cost menu, and an output block that starts with a bold system label (e.g. `**5E**`).

### entities.md
Which sections replace the `{{system sections}}` line in the character, location and group templates (NPC and PC separately), a one-line NPC stat line, and the character math.

### init-questions.md and templates/system.md
5-6 optional setup questions (unanswered = standard), and the matching `system.md` with a `## Homebrew` and a `## Uncertain` section. The template's frontmatter has `system: <id>`.

## 4. Test it
```bash
cd lexis-setup
go test ./...
```
`TestSystemPacksContract` checks every pack:
- exactly the seven contract files, nothing else;
- `pack.md` has `id` (= folder name), `name`, the `## License`, `## Content scope`, `## Glossary` and `## Never` sections, and a glossary row for every role;
- `templates/system.md` names the pack;
- no file contains a term from the setting/brand denylist (outside `## License`). Add the setting names of your source system to the denylist in `lexis-setup/packs_test.go`.

Then install only your pack into a test folder and try it:
```bash
go run . --source ../plugin --platform claude --scope project --dir /tmp/test-world --lang English --system <id> --yes
```
Run `/world-init`, `/help-skills`, `/create-aspect` and `/create-obstacle` with the same scene you used for another system, and compare: the fiction part should match, the mechanics in your system's terms.

## 5. Document it
Add the pack to the table in [Game systems](game-systems.md), to the `--system` list in [Installation](installation.md) and the README, and to the pack table in [design/system-packs.md](design/system-packs.md).
