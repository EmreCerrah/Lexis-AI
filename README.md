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

| Game system | For | License |
|---|---|---|
| 5E SRD | fifth-edition fantasy (2024 core rules) | CC-BY-4.0 |
| ORC 2E Core | the second-edition remaster rules Paizo released under the ORC License | ORC |
| Fate Accelerated | Fate Accelerated | CC-BY |
| Fate Core | Fate Core | CC-BY |
| d100 Percentile Core | percentile roll-under play, including investigative horror (Chaosium's Basic Roleplaying engine) | ORC |
| Forged in the Dark Core | d6 pools, position and effect, clocks | CC-BY |
| Doom d20 Core | MÖRK BORG–compatible play | MÖRK BORG Third Party License |
| Superhero d20 Effects Core | effects-based superhero d20 (third edition) | OGL 1.0a |

System packs contain **game mechanics only**: no published setting, characters, places or story, and no copied rule text. Details: [Game systems](docs/game-systems.md).

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
4. At the table: `/event …` whenever something happens, `/help-skills …` for rules, `/create-obstacle …` for scene material.
5. At the end: `/close-session`, answer once, done.

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
| `/create-lore` | one NPC, place, group or story hook, as a proposal |
| `/close-session` | one numbered list of everything to change; writes only what you approve |

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
  skills/<command>/        # the 8 commands
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
Build with a version label: `go build -ldflags "-X main.version=v0.2.0" -o lexis-setup .`. Cross-compile with `GOOS=linux GOARCH=amd64` or `GOOS=darwin GOARCH=arm64`.

The tests also check every system pack against the pack contract and the content rule ([Creating a system pack](docs/creating-a-system-pack.md#4-test-it)).

## Roadmap
Planned work, in priority order: [ROADMAP.md](ROADMAP.md).

## License
The plugin and the installer are MIT-licensed (`plugin/plugin.json`). Each system pack follows its source's license, stated in its `pack.md`: the ORC packs (`orc-d20`, `d100`) are released under the ORC License, the OGL pack (`hero-d20`) under the OGL 1.0a, and the others carry the attribution their CC-BY or third-party license requires.
