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
- **v0.3:**
  - **Arcs:** a chain of hooks followed as one story ([docs/design/arcs.md](docs/design/arcs.md)). `/event` notices a step done, skipped or threatened; `/close-session` proposes the progress; `/world-init` converts existing hook chains.
  - **`/create-scene`:** the GM names a place and adds notes; the AI fills the gaps with read-aloud text, who is there, 3 interactions with the system's check and difficulty, and a scene tag.
  - **Prebuilt binaries:** GoReleaser builds `lexis-setup` for Linux, Windows (amd64, arm64) and one universal macOS binary on every version tag, so installing needs no Go.
  - **`/suggest-event`:** the scene has stalled, so 3 things that could happen now, one per kind of pressure, unspent ripples first; records nothing ([docs/design/suggest-event.md](docs/design/suggest-event.md)).
  - **`/check-consistency`:** what in the world contradicts what, both sides named, writing nothing itself ([docs/design/check-consistency.md](docs/design/check-consistency.md)).
  - **LICENSE, NOTICE.md and CI:** MIT for the project with per-pack notices, and `gofmt`, `go vet`, `go test` and the Markdown link check on every pull request.
  - **Docs:** a new README, and guides in `docs/`, including moving and backing up a world.
  - **Releases:** version tags and release notes (`vX.Y.Z` + `lexis-setup/vX.Y.Z`).
- **Table tests:** 5E SRD tested at the table.

## 1. Now: `/world-status`
Goal: the one-minute briefing before a session, so the GM does not have to reconstruct two weeks from memory. Design: [docs/design/world-status.md](docs/design/world-status.md).

- Last session, who is in play, active hooks, each arc's step with its `Moves on when:`, what waits for a decision, loose threads.
- Decided: its own command, not `/lexis-help status` — a skill is found through its description, and one description cannot carry both jobs.
- Secrets are opt-in (`/world-status secrets`) and always last, in case players can see the screen.
- Writes nothing.

## 2. Table tests for the other systems
- Fate Accelerated / Fate Core next, then the remaining packs as GMs try them.
- Note what reads badly at the table and fix the pack.

## 3. v1.1 commands

## 4. Prep mode: deep analysis report
- The 9-section analysis of an event, for use between sessions only.
- It writes nothing to canon; useful parts become proposals.

## 5. Later, if asked for
- `npx lexis-ai` for people who have Node but not Go. The prebuilt binaries cover most of this already.
- Signing the macOS binaries so the `xattr` step is not needed (needs an Apple developer account).

## Later (lower priority)
- **Images for characters and locations.** On hold: needs a careful design first.
  - **Platforms differ:** image tools are not built into every AI tool.
  - **Prompts must not leak `🔒` content** into pictures the players will see.
  - **Fallback:** if it comes back, start with showing existing images only (an `## Images` section with `![[file.png]]` embeds). Generated images would start as proposals.
- `history/`: optional canon file(s) for the world's past events.
- More lenses: economy and trade, military and security, character relationships, geography and ecology.
