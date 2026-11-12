# Roadmap

The MVP is done: 8 commands, the Fate Accelerated pack, 3 lenses, and `lexis-setup` for Claude Code, Codex, Gemini CLI and GitHub Copilot. Items below are in priority order. Each one starts with a short design pass before any code.

## 0. Now: real-table test → fixes → v0.2
- Run one full session at the table: `/world-init` → `/event` → `/ripples` → `/help-skills` → `/close-session`.
- Note what failed, what was too long or slow to read at the table, and where the AI broke a rule (changed canon mid-session, exposed a hidden truth, answered in the wrong language).
- Fix round, then tag `v0.2` so users can install a fixed version (`go install …@v0.2.0`).

## 1. More game systems (top priority)
Goal: a second system, then a third, with the GM choosing it in `/world-init`.

- **System pack contract.** Two real systems make the right abstraction visible. Write down what every pack must provide, based on FAE plus the second system:
  - core resolution and difficulty
  - what a "skill question" maps to
  - character sheet sections for PCs
  - which scene items exist (aspects, conditions …)
  - the setup questions for `/world-init`
  - the system's `system.md` template

  Skills keep asking for these by name (`systems/<system>/…`) instead of hard-coding FAE.
- **Make the skills system-neutral.** `/help-skills`, `/create-aspect` and `/create-obstacle` are FAE-shaped today. Each becomes "core fiction + the pack's add-on":
  - `/create-aspect` maps to a condition or tag in systems without aspects.
  - `/help-skills` answers in the pack's own terms: skills, attributes or approaches.
- **Order (GM's choice):** 5E SRD ✅ → 2nd-edition d20 remaster (ORC) ✅ → ~~World of Darkness–style~~ (skipped: not open-licensed) → Fate Core ✅ → d100 percentile (BRP engine, ORC) ✅ → Forged in the Dark ✅ → Mörk Borg–compatible ✅ → M&M 3e–style. Licenses and limits: [docs/design/system-packs.md](docs/design/system-packs.md).
- **Content rule:** packs hold mechanics only, no setting, and a Go test enforces it.
- **Homebrew stays per world** in `system.md` and overrides any pack.
- **`lexis-setup`** installs every pack; `/world-init` lists the installed packs.

## 2. Images for characters and locations
Goal: every character and location can have a picture, generated or taken from the GM's own files.

- **Show existing images:** read pictures already in the world folder (for example `images/`) and show them when the entity comes up.
- **Generate images:** create a portrait or scene picture from the entity's description, the world's tone and its aspects, save it into the world and link it.
- **File format:** each entity file gets an `## Images` section with Obsidian-compatible embeds (`![[gabriel-portrait.png]]`). An image is canon like any other file content; a generated image starts as a proposal and is kept only after approval at `/close-session`.
- **Platform reality (needs a design pass):**
  - Image generation is not built into every tool. Codex and ChatGPT have native image tools; Gemini has Imagen through extensions; Claude Code needs an MCP server or an external API.
  - Showing images also differs: a terminal can only open the file; Obsidian and IDE previews render it inline.
  - Plan a small "image provider" setting chosen at install, with "existing images only" as the fallback that works everywhere.
- **Prompt hygiene:** generated prompts must not leak `🔒` content (hidden truths, GM notes) into images the players will see.

## 3. Arc (chained hooks)
Goal: follow a chain of hooks as one story arc (e.g. Empty Coffins → The Yellow Curtain → The Master's Letter → The Gate of Carcosa).

- An `arcs/` file lists the hooks in order, with the current link and what would move the story to the next one.
- `/event` and `/ripples` notice when an event advances, skips or threatens a link in an active arc.
- `/close-session` proposes arc progress (current hook resolved → next hook active).
- Hidden truths stay fixed; the arc only tracks which part of the story is in play.

## 4. v1.1 commands
- `/world-status`: a one-minute pre-session briefing covering active hooks and arcs, open decisions, pending proposals and what happened last session.
- `/check-consistency`: a between-session scan for broken Immutable Rules, contradictions between files, dangling `[[links]]` and stale `## Uncertain` lines.
- `/suggest-event`: 2-3 "what could happen now" options for a stalled scene. Nothing is recorded.

## 5. Prep mode: deep analysis report
- The 9-section analysis of an event, for use between sessions only.
- It writes nothing to canon; useful parts become proposals.

## 6. Distribution
- README section: save, move and reopen a world (a world is just its folder).
- A no-Go install path: `npx lexis-ai` and/or prebuilt binaries on GitHub Releases.
- Version tags and release notes for every release.

## Later (lower priority)
- `history/`: optional canon file(s) for the world's past events.
- More lenses: economy and trade, military and security, character relationships, geography and ecology.
