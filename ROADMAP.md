# Roadmap

Items are in priority order. Each one starts with a short design pass before any code.

## Done
- **v0.1 (MVP):** the core commands, the Fate Accelerated pack, 3 lenses, and `lexis-setup` for Claude Code, Codex, Gemini CLI and GitHub Copilot.
- **v0.2:**
  - **Game system packs:** 8 systems, mechanics only. The installer installs only the system the GM chooses. Licenses and limits: [docs/design/system-packs.md](docs/design/system-packs.md).
    - 5E SRD, ORC 2E Core, Fate Accelerated, Fate Core, d100 Percentile Core, Forged in the Dark Core, Doom d20 Core, Superhero d20 Effects Core.
    - World of Darkness was skipped: it is not open-licensed.
  - **Skills are system-neutral:** they read the world's pack by role.
  - **`/world-init` fill mode** completes an existing world:
    - a missing `system.md`
    - empty parts and deferred items
    - PCs and NPCs
    - switching the world to another system
  - **`/lexis-help`** lists the commands, the world's state and the next step.
  - **Docs:** a new README, and guides in `docs/`, including moving and backing up a world.
  - **Releases:** version tags and release notes (`vX.Y.Z` + `lexis-setup/vX.Y.Z`).
- **Table tests:** 5E SRD tested at the table.

## 1. Now: Arc (chained hooks)
Goal: follow a chain of hooks as one story arc (e.g. Empty Coffins → The Yellow Curtain → The Master's Letter → The Gate of Carcosa). Design: [docs/design/arcs.md](docs/design/arcs.md).

- An `arcs/` file lists the hooks in order, with the current link and what would move the story to the next one.
- `/event` and `/ripples` notice when an event advances, skips or threatens a link in an active arc.
- `/close-session` proposes arc progress (current hook resolved → next hook active).
- Hidden truths stay fixed; the arc only tracks which part of the story is in play.

## 2. Table tests for the other systems
- Fate Accelerated / Fate Core next, then the remaining packs as GMs try them.
- Note what reads badly at the table and fix the pack.

## 3. v1.1 commands
- `/world-status`: a one-minute pre-session briefing covering active hooks and arcs, open decisions, pending proposals and what happened last session. Overlaps with `/lexis-help`'s world block; decide whether it becomes `/lexis-help status` or its own command.
- `/check-consistency`: a between-session scan for broken Immutable Rules, contradictions between files, dangling `[[links]]` and stale `## Uncertain` lines.
- `/suggest-event`: 2-3 "what could happen now" options for a stalled scene. Nothing is recorded.

## 4. Prep mode: deep analysis report
- The 9-section analysis of an event, for use between sessions only.
- It writes nothing to canon; useful parts become proposals.

## 5. Distribution
- A no-Go install path: prebuilt `lexis-setup` binaries on GitHub Releases and/or `npx lexis-ai`.
- Optional: publish `docs/` with GitHub Pages.

## Later (lower priority)
- **Images for characters and locations.** On hold: needs a careful design first.
  - **Platforms differ:** image tools are not built into every AI tool.
  - **Prompts must not leak `🔒` content** into pictures the players will see.
  - **Fallback:** if it comes back, start with showing existing images only (an `## Images` section with `![[file.png]]` embeds). Generated images would start as proposals.
- `history/`: optional canon file(s) for the world's past events.
- More lenses: economy and trade, military and security, character relationships, geography and ecology.
