---
name: create-obstacle
description: Generate one ready-to-run obstacle, puzzle, hazard or challenge for the current RPG scene in a Lexis-AI world - fiction first, then Fate Accelerated mechanics (aspect, suggested difficulty, fitting approaches, costs). Use when the GM asks for an obstacle, trap, puzzle, lock, hazard or challenge ("create-obstacle", "I need a trap", "a puzzle", "engel üret").
---

# /create-obstacle

A hybrid command: **core fiction first, then the system add-on.** One obstacle, logged as a proposal. You suggest the difficulty; the GM decides and rolls nothing through you.

## Read first
`../../AGENTS.md`; `../../systems/<system>/pack.md`, `obstacles.md` and `approaches.md`; `world.md` (`## Immutable Rules`, tone, language); `system.md` (Homebrew overrides); the open log; the files of the current place and linked hooks; the PCs' files if named (stunts).

## Steps
1. **Context + type.** From the GM's words and the last `E<n>`: barrier / puzzle / hazard / social / challenge (`obstacles.md`). Default: barrier.
2. **Core block (no numbers):**
   - What: 1-2 sentences the GM can read aloud.
   - Why here: how it follows from the place, a group or a hook's `## Surface`.
   - 🔒 GM: what it protects or hides. It must stay consistent with `## 🔒 Hidden Truth`: it may guard the truth, never reveal it.
   - Puzzle: the clue chain, solvable without dice too.
3. **Fate block** in the output format of `obstacles.md`. Approaches: 2-3, with the fiction for each. Mention a PC stunt if one matches.
4. **Rule check:** every element obeys every `IR<n>`; drop anything that would need an exception.
5. **Log:** `- G<n> obstacle (E<k>|—): <short name> · <type> · <difficulty range>` in `## Generated`.

## Output (table mode, in the output language)
```
🧱 G<n> · <Obstacle name> — <type>
<What — 1-2 sentences to read aloud>
- Why here: …
- 🔒 GM: …
**Fate**
- Aspect: "…"
- Suggested difficulty: … (+n) — raises it: … / lowers it: …
- Fitting approaches: Careful — …; Clever — …
- Cost options: …; …
```

## Never
Decide the outcome, roll, or write to canon files.
