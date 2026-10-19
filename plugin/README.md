# Lexis-AI

A **GM assistant** for the live table. You run the game; the agent never rolls dice and never decides outcomes. It records the events you report, shows how they ripple through the world, and updates your world files only with your approval at the end of the session.

- System: **Fate Accelerated** (MVP)
- Format: plain Markdown + `[[wiki-links]]`. Opens in Obsidian, but doesn't need it.
- Your world lives in your own folder. The plugin contains no world data.
- Language: the AI answers in the language you pick at install (each world can override it in `world.md`).

## Commands
| Command | What it does | What it writes |
|---|---|---|
| `/world-init` | Sets up a world in tiers: required core, deferrable parts, parts the agent can draft, FAE settings | all world files (only at setup) |
| `/event` | You report what happened at the table. It becomes canon at once; you get 3-5 ripples and a "who might know" list | the session log only |
| `/ripples` | Deepens a recorded event through lenses and shows how word of it spreads 2-3 steps | the session log only (hypotheses) |
| `/create-aspect` | Writes one aspect for the scene, with invoke and compel examples | the log (proposal) |
| `/create-obstacle` | Creates an obstacle, puzzle or hazard: fiction first, then FAE difficulty and approaches | the log (proposal) |
| `/create-lore` | Creates a persistent NPC, location, group or hook | `proposals/` (proposal) |
| `/help-skills` | Answers "which approach is this?" in FAE terms | nothing |
| `/close-session` | Gathers everything from the session into one numbered list for your approval and writes only what you approve | canon files |

How you call them depends on the tool: Claude Code `/lexis-ai:event`, Codex `$event` (or `@lexis-ai`), Copilot `/event`. In Gemini CLI, just describe what you want. A plain-language request works everywhere.

## Flow
```
/world-init  →  (at the table) /event … /ripples … /create-* … /help-skills  →  /close-session
```
- **Table mode** (default): short bullets, no approval questions mid-session.
- **Prep mode:** add `prep` or `detailed` to a command for deeper analysis.

## Knowledge model
| Kind | Where | How it changes |
|---|---|---|
| Canon | file body | only with `/close-session` approval |
| Hypothesis | `## Uncertain`, or `## Ripples` in the log | at session end: promoted to canon, kept as uncertain, or rejected |
| Inference | only in answers | never stored |
| Proposal | `proposals/`, or `## Generated` in the log | approved or rejected at session end |

The "who might know" list comes from the `## Links` section of each file (`- [[slug]] — relation`). The agent does not see links buried in prose, so write links in that section.

## Rules
- The **Immutable Rules** in `world.md` (IR1, IR2 …) are never broken by anything the agent produces. If an event you report breaks one, the agent does not block it. It flags it `⚠ rule conflict` and offers three options: (a) apparent violation, (b) costly exception, (c) rule change.
- A hook's **🔒 Hidden Truth** is fixed. If players try to break it, the agent suggests in-world reasons why they can't.
- Don't read lines marked `🔒` to the players.

## Install
Run the `lexis-setup` tool from the Lexis-AI repository. It asks which AI tool you use (Claude Code, Codex, Gemini CLI, GitHub Copilot), where to install, and which language the AI should answer in, then installs the plugin in that tool's format. See the README at the repo root.

If you use Codex, open your world folder in **workspace-write** mode (with file-write permission); otherwise Codex asks for permission on every write mid-session.

## Package layout
```
plugin.json                 # the only manifest (agent-plugins.org schema)
AGENTS.md                   # core rules every skill reads first
skills/<command>/SKILL.md   # 8 MVP skills (platform-neutral)
references/                 # knowledge model, folder layout, lenses (hidden references, not skills)
templates/                  # world core templates
systems/fate-accelerated/   # FAE system pack
```
The repo has a single `plugin.json`. `lexis-setup` generates each tool's own manifest (`.claude-plugin/plugin.json`, `.codex-plugin/plugin.json`, `gemini-extension.json`) from it at install time.
