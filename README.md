# Lexis-AI

The **Lexis-AI** plugin, a GM assistant for the live tabletop RPG table, plus `lexis-setup`, the tool that installs it into different AI tools.

```
plugin/      # the plugin (platform-neutral skills, rules, templates, system packs)
lexis-setup/ # the installer, written in Go
.agents/plugins/marketplace.json   # Codex: the repo can be added directly as a marketplace
```

## Install
You need [Go](https://go.dev/dl/) 1.22 or newer. Install the setup tool straight from GitHub; there is nothing to clone or build:
```bash
go install github.com/EmreCerrah/Lexis-AI/lexis-setup@latest
```
Then run it:
```bash
lexis-setup
```
`go install` prints nothing when it succeeds. It puts the program in Go's bin folder: `%USERPROFILE%\go\bin` on Windows, `~/go/bin` on macOS and Linux. If your terminal says `lexis-setup` is not found, add that folder to your `PATH`, or run the program with its full path. To update the setup tool later, run the same `go install` command again.

The tool asks four questions:
1. **Which AI tool do you use?** Claude Code, Codex, Gemini CLI or GitHub Copilot.
2. **Where do you want to install it?** For your user, or in a specific project folder.
3. **Which language should the AI answer in?** English, Türkçe, Español, Deutsch, Français, Português, or any other language you type.
4. **Which game system do you play?** 5E SRD (fifth edition compatible), Fate Accelerated, or all of them. Only the systems you choose are installed; the others' files are never written. To add a system later, run `lexis-setup` again and choose it (or "All systems").

It then downloads the plugin from GitHub (`EmreCerrah/Lexis-AI`, branch `main`), shows the install plan and asks for confirmation.

| Tool | User install | Project install |
|---|---|---|
| Claude Code | plugin via a local marketplace (`~/.lexis-ai/claude-code`) + `claude plugin install` | `<project>/.claude/skills` + `<project>/.claude/lexis-ai` |
| Codex | plugin via a local marketplace (`~/.lexis-ai/codex`) + `codex plugin add` | `<project>/.agents/skills` + `<project>/.agents/lexis-ai` |
| Gemini CLI | extension `~/.gemini/extensions/lexis-ai` via `gemini extensions install` | `<project>/.gemini/skills` + `<project>/.gemini/lexis-ai` |
| GitHub Copilot | `~/.copilot/skills` + `~/.copilot/lexis-ai` | `<project>/.github/skills` + `<project>/.github/lexis-ai` |

A **project install** is a plain copy into the folder each tool already scans for project skills: the skills go to `skills/`, the shared files (rules, references, templates, system packs) go to `lexis-ai/` next to it. Nothing is registered with the tool, so the skills are there as soon as you open the tool in that folder, and you can commit them with the project. In Claude Code, project skills are called without the plugin prefix (`/event` instead of `/lexis-ai:event`). Re-running a project install removes what older versions put in `<project>/.lexis-ai/` or `<project>/plugins/`.

**Language:** the chosen language is written into the installed `AGENTS.md` as the default output language. `/world-init` offers it as the world's language and stores it in `world.md`, so each world can use its own. File names, headings and frontmatter keys stay in English so every skill can find them; only the content is in the chosen language.

The repo has a single manifest, `plugin/plugin.json`. The installer generates each tool's own manifest from it.

Skills refer to the plugin root as `../../`. The installer rewrites that path for each tool:
- Claude Code, user install: `${CLAUDE_PLUGIN_ROOT}/`
- Copilot, user install: the absolute path
- Every project install: `../../lexis-ai/`
- Codex and Gemini CLI, user install: unchanged

### Flags
| Flag | Meaning |
|---|---|
| `--platform claude\|codex\|gemini\|copilot` | skips the tool question |
| `--scope user\|project`, `--dir <folder>` | where to install |
| `--lang <language>` | skips the language question (e.g. `--lang English`, `--lang Turkish`) |
| `--system <id>` | skips the game system question: `srd5`, `fate-accelerated`, several separated by commas, or `all` |
| `--source <folder>` | installs from a local copy instead of GitHub (for development) |
| `--repo owner/repo`, `--ref <branch/tag>` | another repository or version |
| `--dry-run` | shows the plan without writing anything |
| `--yes` | doesn't ask for confirmation |
| `--force` | overwrites a folder that does not belong to Lexis-AI |

Every folder the tool writes carries a `.lexis-ai.json` marker. When the tool runs again, it only replaces folders that carry this marker. The plugin used to be called `worldbuilding-forge`; an install under that name is removed and replaced by `lexis-ai` when `lexis-setup` runs again.

For a private repository, set the `GITHUB_TOKEN` environment variable.

## Usage
Lexis-AI helps you **while you run the game**. Game systems: **Fate Accelerated** and **5E SRD** (fifth-edition rules); you pick the systems at install, and each world uses one of them (`/world-init` picks it automatically when only one is installed). System packs hold only rules, never a published setting. You stay in charge: it never rolls dice and never decides whether an action succeeds. You tell it what happened; it records that, shows how it ripples through the world, and helps with rules and scene material. Your world files change only when you approve the changes at the end of the session.

### How to call a command
Make an empty folder for your world, open your AI tool in that folder, then call a command.

| Tool | How to call it | Example |
|---|---|---|
| Claude Code | `/lexis-ai:<command>` (user install), `/<command>` (project install) | `/lexis-ai:event The players broke into the crypt` |
| Codex | `$<command>` (or pick `@lexis-ai`) | `$event The players broke into the crypt` |
| GitHub Copilot | `/<command>` | `/event The players broke into the crypt` |
| Gemini CLI | describe what you want | `event: the players broke into the crypt` |

Plain requests work in every tool too, for example "record what just happened" or "close the session". Below, commands are written as `/event`, `/ripples` and so on.

### Commands
| Command | When to use it | What you get | What it writes |
|---|---|---|---|
| `/world-init` | once, before the first session | A short interview: name, concept, genre, language, tone, inspirations, fantasy level, **Immutable Rules**, system. Characters, places, groups and hooks can be skipped ("later") or drafted by the AI ("you suggest"). | creates the world files |
| `/event <what happened>` | at the table, right after something happens | The event is recorded at once. You get 3-5 short **ripples** (possible consequences) and a **who might know** list. | session log only |
| `/ripples [E2]` | between scenes, when you want more | The event examined through lenses (power and politics, society and rumor, faith and the supernatural), and how word of it spreads 2-3 steps. | session log only |
| `/create-aspect [scene]` | you need a scene tag | One scene tag in your system: a Fate aspect with an invoke and a compel, or a 5E scene feature (light, cover, terrain, hazard, condition) with its rule effect. | session log (proposal) |
| `/create-obstacle [scene]` | you need a lock, trap, puzzle, hazard or social obstacle | The fiction first (what it is, why it is here), then your system's block: a suggested difficulty (ladder or DC), fitting approaches or checks, cost options; for traps, the save and its effect. | session log (proposal) |
| `/create-lore <what you need>` | you need an NPC, place, faction or hook | One ready entity with its system sections (aspects, a stat line …) and links. | `proposals/` (proposal) |
| `/help-skills <what the player does>` | a rules question at the table | The answer in your system's terms: in Fate the action and 2-3 approaches, in 5E the ability check and skills; a suggested difficulty; matching stunts, features or scene tags. | nothing |
| `/close-session` | at the end of the session | Everything from the session in one numbered list, each with a suggested answer. You answer once; only what you approve is written. | world files |

### A session, step by step
1. **Set up the world:** `/world-init`. Answer the core questions in one go. Say "later" to anything you don't want to decide yet; it goes to the world's `## Uncertain` section.
2. **Report what happens:** `/event Gabriel caught the gravekeeper next to an empty coffin. The keeper ran and dropped a theater ticket with a yellow sign on it.`
   ```
   📌 E1 recorded (session-001) · Lens: society-rumor, faith-supernatural
   Ripples (hypothesis):
   - R1 The keeper hides in the tavern and mutters "it wasn't a vampire"
   - R2 By morning the town says "the hunter attacked the keeper"
   - R3 🔒 The troupe notices the lost ticket and sends a marked corpse after the keeper
   Who might know:
   - [[gravekeeper]] — there
   - [[order-of-light]] — may hear (member → [[gabriel]])
   ```
3. **Ask rules questions:** `/help-skills Gabriel wants to scare the keeper into talking`. **Get scene material:** `/create-obstacle the cellar door the keeper locked behind him`.
4. **Dig deeper when there is time:** `/ripples E1`.
5. **Close the session:** `/close-session`. You get a numbered list (events to write into world files, ripples to keep or drop, new links, which generated items you used). Answer once, for example `all as suggested, 3 ❌, defer 7`.

### Table mode and prep mode
- **Table mode** is the default: at most 5 short bullets, no questions unless they are needed, and no approval questions in the middle of the game.
- **Prep mode:** add `prep` or `detailed` to a command, for example `/ripples E1 prep`, for a longer analysis between sessions.

### What it never does
- It never changes your world files during the session. Events, ripples and generated items wait in the session log (`sessions/`) until `/close-session`.
- It never breaks an **Immutable Rule** in anything it suggests. If you report an event that breaks one, it records the event anyway, flags `⚠ rule conflict` and offers three choices: (a) it only looks like a violation, (b) a costly exception, (c) the rule changes (decided at `/close-session`).
- It never reveals a hook's **🔒 Hidden Truth**. If players try to break it, it suggests in-world reasons why they can't.
- Lines marked `🔒` are for the GM only. Don't read them to the players.

### Your world folder
```
world.md        # concept, tone, Immutable Rules (IR1, IR2 …)
system.md       # system pack settings and Homebrew
characters/  locations/  groups/  hooks/   # one Markdown file per entity
sessions/       # one log per session
proposals/      # /create-lore output waiting for approval
```
These are plain Markdown files with `[[links]]`, and you can edit them yourself. Open the folder in Obsidian if you want a linked view. "Who might know" only uses the `## Links` section of each file, so put relations there, one per line: `- [[order-of-light]] — member`.

## Roadmap
Planned work, in priority order: [ROADMAP.md](ROADMAP.md).

## Build
```bash
cd lexis-setup && go test ./... && go build -o lexis-setup.exe .
```
The version label and the repository can be set at build time:
```bash
cd lexis-setup && go build -ldflags "-X main.version=v0.1.0 -X main.defaultRepo=EmreCerrah/Lexis-AI" -o lexis-setup.exe .
```
For other operating systems, build with `GOOS=linux GOARCH=amd64` or `GOOS=darwin GOARCH=arm64`.
