---
name: check-consistency
description: Scan a Lexis-AI world between sessions for things that contradict each other - a dead character still acting, canon that breaks an Immutable Rule, arc steps resolved out of order, hypotheses a later event already settled, links only one file knows about. Reports both sides of every finding and writes nothing by itself. Use when the GM asks to check the world, look for contradictions or clean up before a session ("check-consistency", "is my world consistent", "find contradictions", "tutarlılık kontrolü", "çelişki var mı").
---

# /check-consistency

Reads the whole world and reports **what disagrees with what**. A prep command: thorough, slow, and it changes nothing on its own. For what is *missing* (no `system.md`, an empty folder, a hook without a hidden truth, a dangling link), that is `/world-init` fill mode, not this.

```
/check-consistency [area]
/check-consistency arcs
/check-consistency characters/
```

## Read first
`../../AGENTS.md`; `../../references/knowledge-model.md`; `../../references/world-layout.md`; `world.md` (every `IR<n>`, `system:`); `system.md`; every file in `characters/ locations/ groups/ hooks/ arcs/ proposals/`; every log in `sessions/` (their `status:`, events, unanswered `D<m>`, `R<k>` and what later events did to them).
Without an area argument, read everything. With one (a folder, a file, `rules`, `arcs`, `links`, `uncertain`), read what that area needs and say in the header that the check was narrowed.

## What to look for
Only report a finding you can point at in **two places**. Guessing a contradiction out of one file's prose is not a finding.

1. **Immutable Rules:** a canon statement that breaks an `IR<n>`; an `⚠ rule conflict` in any log that was never answered.
2. **State vs fiction:** a character or place whose state in one file is contradicted by what another file says it does now — dead but acting, destroyed but working, imprisoned but present, departed but keeping a place.
3. **Relations:** the same relation described differently in two files (member vs expelled, owner vs former owner); a group's member list against those members' own files; a one-way link whose relation clearly runs both ways.
4. **Arcs vs hooks:** a later step's hook `resolved` while an earlier one is not; an arc `active` with every hook resolved, or `resolved` with an unresolved hook; one hook listed in two arcs; a step whose `Moves on when` a session met but never closed.
5. **Hidden truths:** a `## 🔒 Hidden Truth` that canon already contradicts, or that events have effectively confirmed to the players; two hooks whose hidden truths cannot both be true. Mark these `🔒` and never quote the truth itself.
6. **Stale hypotheses:** an `## Uncertain` line that a later `E<n>` confirmed or refuted and that was never promoted or removed; a line whose source session is more than three sessions old.
7. **System drift:** entity files carrying the sections of a system pack the world no longer uses, or two files naming the same mechanic differently.
8. **Loose ends** (info only): a log left `pending`; a proposal open for more than three sessions; an event whose facts never reached any world file.

## Output (prep mode, in the output language)
```
🔍 Consistency check: <world> (<n> files, <n> sessions)<, area: <area>>
⚠ Contradictions (<n>)
1. <side A> vs <side B>, each with its file and section
   → (a) … (b) … (c) …
❓ Drift (<n>)
4. …
   → (a) … (b) …
ℹ Info (<n>)
- …
```
- Severity is **certainty**, not importance: `⚠` two canon statements that cannot both be true · `❓` probably stale, but the GM may have meant it · `ℹ` nothing to decide.
- One line per finding, plus one `→` line of options. Number continuously across groups.
- At most **10 findings**; end with `… and <n> more (narrow it: /check-consistency <area>)` when there are more.
- If nothing is wrong, say so in one line with what was checked. Do not pad the report.

## When the GM answers
The answers are **not** applied here. Append each one to the open session log as an open decision — create a log with `mode: prep` if none is open, as `/ripples` does — in `## Open Decisions`:
`- D<n> consistency: <the finding in one line> → <the GM's answer>`
`/close-session` applies them in group A with the usual approval. Say in one line that they are waiting there. If the GM answers nothing, write nothing at all.

## Never
- Edit a canon file, or "clean up" anything on your own.
- Quote a `## 🔒 Hidden Truth` or an arc's `## 🔒 Endgame`; name the file and say what it conflicts with.
- Report missing content, style or tone: that is `/world-init` fill mode, and tone is the GM's.
- Present a finding as settled, or decide which side is right. The GM decides.
