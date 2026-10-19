---
name: create-obstacle
description: Generate one ready-to-run obstacle, puzzle, hazard or challenge for the current RPG scene in a Lexis-AI world - fiction first, then Fate Accelerated mechanics (aspect, suggested difficulty, fitting approaches, costs). Use when the GM asks for an obstacle, trap, puzzle, lock, hazard or challenge ("create-obstacle", "engel üret", "bir bulmaca", "tuzak lazım").
---

# /create-obstacle

Hybrid command: **core fiction first, then the system add-on.** One obstacle, logged as a proposal. You suggest difficulty; the GM decides and rolls nothing through you.

## Read first
`../../AGENTS.md`; `../../systems/<sistem>/pack.md` and `obstacles.md`, `approaches.md`; `world.md` (`## Değişmez Kurallar`, tone); `system.md` (Homebrew overrides); the open log; files of the current place and linked hooks; PCs' files if named (stunts).

## Steps
1. **Context + type.** From the GM's words and the last `O<n>`: barrier / puzzle / hazard / social / challenge (`obstacles.md`). Default: barrier.
2. **Core block (no numbers):**
   - Ne: 1-2 sentences the GM can read aloud.
   - Neden burada: how it follows from the place, a group or a hook's `## Yüzey`.
   - 🔒 GM: what it protects or hides — must stay consistent with `## 🔒 Gizli Gerçek`; it may guard the truth, never reveal it.
   - Puzzle: the clue chain, solvable without dice too.
3. **Fate block** per `obstacles.md` output format. Approaches: 2-3, with the fiction for each. Mention a PC stunt if one matches.
4. **Rule check:** every element obeys `K<n>`; drop anything that would need an exception.
5. **Log:** `- U<n> engel (O<k>|—): <kısa ad> · <tür> · <zorluk aralığı>` in `## Üretilenler`.

## Output (Turkish, table mode)
```
🧱 U<n> · <Engel adı> — <tür>
<Ne — okunabilir 1-2 cümle>
- Neden burada: …
- 🔒 GM: …
**Fate**
- Aspect: "…"
- Önerilen zorluk: … (+n) — artırır: … / düşürür: …
- Uygun approach'lar: Careful (Dikkatli) — …; Clever (Zeki) — …
- Bedel seçenekleri: …; …
```

## Never
Decide the outcome, roll, or write to canon files.
