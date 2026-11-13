# Lexis-AI

**A GM assistant for the live tabletop RPG table.** Lexis-AI runs inside the AI tool you already use (Claude Code, Codex, Gemini CLI or GitHub Copilot) and keeps your campaign world as plain Markdown files in your own folder.

You run the game. Lexis-AI never rolls dice and never decides whether an action succeeds. You tell it what happened; it records it, shows how it ripples through the world and who might hear about it, helps with rules in your game system's own terms, and changes your world files only with your approval at the end of the session.

```
/event Gabriel caught the gravekeeper next to an empty coffin. The keeper ran and dropped a theater ticket.

📌 E1 recorded (session-001) · Lens: society-rumor, faith-supernatural
Ripples (hypothesis):
- R1 The keeper hides in the tavern and mutters "it wasn't a vampire"
- R2 By morning the town says "the hunter attacked the keeper"
- R3 🔒 The troupe notices the lost ticket and sends someone after the keeper
Who might know:
- [[gravekeeper]] — there
- [[order-of-light]] — may hear (member → [[gabriel]])
```

## Features
- **Events become canon instantly**, consequences stay hypotheses until you decide. Nothing in your world changes behind your back.
- **Ripples through lenses:** power and politics, society and rumor, faith and the supernatural.
- **Who might know:** follows the links between your characters, places and groups.
- **Arcs:** chain hooks into one story; Lexis-AI notices when the players finish a step, jump ahead or put a later step at risk, and proposes the progress at session end.
- **Immutable Rules and hidden truths:** suggestions never break your world's core rules and never reveal a hook's secret. If *you* break a rule, it is recorded and you choose how the world deals with it.
- **Rules help in 8 game systems**, with scene material (scene tags, obstacles, traps, NPCs) in each system's terms.
- **One end-of-session list:** approve, keep as uncertain, or drop each item with a single answer.
- **Your language:** answers in English, Türkçe, Español, Deutsch, Français, Português or any language you choose; each world can have its own.
- **Plain Markdown with `[[links]]`:** readable, editable, Obsidian-compatible, easy to back up with git.

## Supported AI tools and game systems
| AI tool | Install |
|---|---|
| Claude Code | user-wide plugin or project folder |
| Codex | user-wide plugin or project folder |
| Gemini CLI | user-wide extension |
| GitHub Copilot | user-wide or project folder |

| Game system | Based on | License |
|---|---|---|
| 5E SRD | System Reference Document 5.2.1 by Wizards of the Coast: the 2024 core rules of D&D fifth edition¹ | CC-BY-4.0 |
| ORC 2E Core | Pathfinder Player Core and GM Core (Paizo, remaster) | ORC |
| Fate Accelerated | Fate Accelerated Edition (Evil Hat) | CC-BY |
| Fate Core | Fate Core System (Evil Hat) | CC-BY |
| d100 Percentile Core | Basic Roleplaying: Universal Game Engine (Chaosium), the percentile engine behind classic investigative horror | ORC |
| Forged in the Dark Core | Blades in the Dark SRD (One Seven Design) | CC-BY |
| Doom d20 Core | MÖRK BORG (compatible) | MÖRK BORG Third Party License |
| Superhero d20 Effects Core | the effects-based superhero d20 rules, third edition, from the [d20 Hero SRD](https://www.d20herosrd.com/) (OGL) | OGL 1.0a |

System packs contain **game mechanics only**: no published setting, characters, places or story, and no copied rule text. Game names appear only to say which game a pack's rules come from; all trademarks belong to their owners, and Lexis-AI is not affiliated with or endorsed by them. Details: [Game systems](docs/game-systems.md).

¹ D&D is a trademark of Wizards of the Coast LLC. Lexis-AI is not affiliated with, endorsed or sponsored by Wizards of the Coast.

## Quick start
1. Install [Go](https://go.dev/dl/) 1.22+ and the setup tool:
   ```bash
   go install github.com/EmreCerrah/Lexis-AI/lexis-setup@latest
   ```
2. Run it and answer four questions (AI tool, where to install, answer language, game system):
   ```bash
   lexis-setup
   ```
3. Make an empty folder for your campaign, open your AI tool there and create the world:
   ```
   /world-init
   ```
   (Claude Code user install: `/lexis-ai:world-init` · Codex: `$world-init` · Gemini CLI: "set up a new world")
4. At the table: `/event …` whenever something happens, `/help-skills …` for rules, `/create-scene …` when the players enter a place, `/create-obstacle …` for more scene material.
5. At the end: `/close-session`, answer once, done.

Lost? `/lexis-help` lists every command and suggests the next step.

Full walk-through: [Getting started](docs/getting-started.md).

## Commands
| Command | What it does |
|---|---|
| `/world-init` | creates a world through a short interview (anything can be "later" or "you suggest") |
| `/event` | records what happened, lists 3-5 ripples and who might know |
| `/ripples` | deepens an event through lenses and traces how the news spreads |
| `/help-skills` | answers a rules question in your system's terms |
| `/create-aspect` | one scene tag (aspect, scene feature, circumstance, clock …) |
| `/create-obstacle` | one obstacle, trap, puzzle or hazard: fiction first, then the numbers |
| `/create-scene` | a place the players enter: your notes plus read-aloud text, who is there, 3 interactions with checks and difficulties |
| `/create-lore` | one NPC, place, group or story hook, as a proposal |
| `/close-session` | one numbered list of everything to change; writes only what you approve |
| `/lexis-help` | lists the commands, shows where your world stands and what to do next |

Details and examples: [Commands](docs/commands.md).

## Documentation
| Guide | |
|---|---|
| [Installation](docs/installation.md) | install, flags, where files go, updating, test versions, uninstalling |
| [Getting started](docs/getting-started.md) | your first world and session, step by step |
| [Commands](docs/commands.md) | every command in detail |
| [Game systems](docs/game-systems.md) | the 8 system packs, how each answers, licenses |
| [Your world files](docs/world-files.md) | layout, canon vs hypothesis, links, editing, moving and backing up a world |
| [Troubleshooting](docs/troubleshooting.md) | common problems and fixes |
| [Creating a system pack](docs/creating-a-system-pack.md) | add a new game system |

## Repository layout
```
plugin/                    # the plugin: platform-neutral skills, core rules, templates, system packs
  AGENTS.md                # core rules every command reads first
  skills/<command>/        # the 10 commands
  references/              # knowledge model, world layout, lenses
  templates/               # world file templates
  systems/<id>/            # one folder per game system pack
lexis-setup/               # the installer (Go)
docs/                      # documentation
.agents/plugins/marketplace.json   # lets Codex add this repo directly as a marketplace
```

## Development
```bash
cd lexis-setup
go test ./...
go build -o lexis-setup .
go run . --source ../plugin --dry-run
```
Build with a version label: `go build -ldflags "-X main.version=v0.3.0" -o lexis-setup .`. Cross-compile with `GOOS=linux GOARCH=amd64` or `GOOS=darwin GOARCH=arm64`.

**Releasing:** bump `version` in `plugin/plugin.json`, then tag the merge commit twice: `vX.Y.Z` (the release) and `lexis-setup/vX.Y.Z` (the Go module lives in `lexis-setup/`, so `go install …@vX.Y.Z` and `@latest` read that tag).

The tests also check every system pack against the pack contract and the content rule ([Creating a system pack](docs/creating-a-system-pack.md#4-test-it)).

## Roadmap
Planned work, in priority order: [ROADMAP.md](ROADMAP.md).

## License
The plugin and the installer are MIT-licensed (`plugin/plugin.json`). Each system pack follows its source's license, stated in its `pack.md`: the ORC packs (`orc-d20`, `d100`) are released under the ORC License, the OGL pack (`hero-d20`) under the OGL 1.0a, and the others carry the attribution their CC-BY or third-party license requires.
