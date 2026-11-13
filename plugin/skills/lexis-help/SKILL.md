---
name: lexis-help
description: Show the Lexis-AI commands, what each one does and how to call it, plus where the current world stands and what to do next. With a command name, explain that command in detail; with a question ("how do I add a character?"), point to the right command. Records nothing. Use when the GM asks "lexis-help", "Lexis-AI help", "what commands are there", "what can you do", "how do I …", "komutlar neler", "yardım".
---

# /lexis-help

A short guide to Lexis-AI, in the output language. This skill records nothing and writes nothing.

## Read first
`../../AGENTS.md` (output language, §10 system packs). If there is a world (`world.md` with `type: world` in the current folder or above): its frontmatter (`name`, `language`, `system`), whether `system.md` exists, the newest log in `sessions/` (its `status:` and how many `E`, `R`, `G`, `D` items it has), and every `arcs/*.md` with `status: active` (its steps and the status of their hooks). The installed packs: `../../systems/*/pack.md` (`name`).

## How to call a command
Show every command in the form the GM just used to call this one:
- called as `/lexis-ai:lexis-help` (Claude Code, user install) → `/lexis-ai:<command>`
- called as `/lexis-help` (Claude Code project install, GitHub Copilot) → `/<command>`
- called as `$lexis-help` (Codex) → `$<command>`
- no slash form (Gemini CLI, or plain words) → say that commands are called by describing them ("record an event: …", "close the session") and give the description next to each.
If unsure, use `/<command>` and add one line on the other forms.

## The commands
| Command | One line | Example |
|---|---|---|
| `world-init` | creates a world through a short interview; in an existing world it lists what is missing and completes it (system.md, characters, deferred items, system switch) | `/world-init` |
| `event` | records what just happened as canon, lists 3-5 ripples and who might know | `/event The keeper ran and dropped a ticket.` |
| `ripples` | more consequences of a recorded event, through lenses, and how the news spreads | `/ripples E2` |
| `help-skills` | a rules answer in the world's system: which check, difficulty, options | `/help-skills climb the wall in the rain` |
| `create-aspect` | one scene tag for the current scene, helping and hindering | `/create-aspect foggy cemetery` |
| `create-obstacle` | one obstacle, trap, puzzle or hazard: fiction first, then the numbers | `/create-obstacle locked crypt door` |
| `create-scene` | a place the players enter: read-aloud text, who is there, 3 interactions with checks and difficulties, a scene tag; your notes are kept, it fills the gaps | `/create-scene old-cemetery night, fog, a grave dug open` |
| `create-lore` | one NPC, place, group or hook as a proposal | `/create-lore a nervous gravekeeper` |
| `close-session` | one numbered list of everything to change; writes only what the GM approves | `/close-session` |
| `lexis-help` | this guide | `/lexis-help event` |

Translate the "One line" text and the example sentences into the output language; keep the command names as they are. Add: any command takes `prep` or `detailed` for a longer answer between sessions.

## Answer
**No argument:**
```
📖 Lexis-AI
<the command table, as a short list: command — one line>

🌍 <world name> · <system pack name> · <language>
- Session: <session-NNN open, E3 R7 G1 D0 | no open session>
- 🧭 <arc name>: step <n>/<total> [[current hook]]   (one line per active arc; omit if none)
- Next: <one suggestion>
```
- **Next** follows the world's state: no world → `/world-init`; no `system.md`, a `system:` that is not installed, or empty `characters/` → `/world-init` to complete the world; an open log → keep reporting with `/event`, and `/close-session` at the end; a `pending` log → `/close-session` to finish the deferred items; otherwise → `/event` when play starts.
- If the world's system is not installed, say so in the world line (re-run `lexis-setup`, or switch with `/world-init`).
- No world: replace the world block with one line: "No world here: run `/world-init` in an empty folder for your campaign."

**A command name** (`/lexis-help event`): that command's usage line, what it reads and writes, what its answer looks like (2-4 bullets), and one example. Take the details from that command's own `SKILL.md` (`../<command>/SKILL.md`).

**A question** ("how do I add a PC?", "how do I change the system?"): name the command that does it, with a one-line example. Questions about rules go to `help-skills`, not here.

## Never
- Write or change any file, or open a session log.
- Answer a rules question itself (point to `help-skills`).
- Reveal a `## 🔒 Hidden Truth`, `## 🔒 GM Notes` or an arc's `## 🔒 Endgame` while describing the world.
