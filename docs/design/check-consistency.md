# Design: /check-consistency (the files disagree)

Status: **accepted** (decisions below).

## Goal
A world grows for a year and starts to contradict itself. A character is dead in one file and guarding a gate in another. A session confirmed a rumor nobody ever promoted. An arc's third step resolved before its second. None of this is visible from any single file, and the GM only finds it at the table, mid-scene, in front of the players.

`/check-consistency` reads the world between sessions and reports what disagrees with what. It is a prep command: slow, thorough, and it changes nothing on its own.

## The boundary with `/world-init` fill mode
These two must never scan for the same thing. The split is **missing vs contradicting**:

| | `/world-init` (fill mode) | `/check-consistency` |
|---|---|---|
| Question | what is **not there** yet? | what is there but **cannot all be true**? |
| Examples | no `system.md`, an empty `characters/`, a hook with no hidden truth, a deferred `## Uncertain` line, an unconverted hook chain | a dead character acting, a rule broken in canon, an arc resolved out of order, a hypothesis a later event already settled |
| Writes | the files the GM answers for, now | nothing (see below) |
| When | after an install, or when something is clearly missing | between sessions, on a world with history |

Dangling `[[links]]` (a link to a file that does not exist) stay with fill mode: the target is *missing*. One-way links are a contradiction candidate and belong here.

## What it checks
Grouped by what kind of disagreement it is. Each finding names **both** sides, with file and section, so the GM can judge it in one line.

1. **Immutable Rules.** A canon statement that breaks an `IR<n>`. Also an `⚠ rule conflict` in any log that was never answered.
2. **State vs fiction.** A character's or place's state in one file against what another file says it is doing now: dead but acting, destroyed but working, imprisoned but present, gone but keeping a place.
3. **Relations.** Two files describing the same relation differently (member vs expelled, owner vs former owner). A one-way link where the relation clearly runs both ways (a reciprocal is suggested, never written). A group's member list against the members' own files.
4. **Arcs vs hooks.** A later step resolved before an earlier one; an arc `active` with every hook resolved, or `resolved` with an unresolved hook; a hook listed in two arcs; a step whose `Moves on when` was met in a session that never closed it.
5. **Hidden truths.** A `## 🔒 Hidden Truth` that canon has already contradicted or that the players have effectively confirmed; two hooks whose hidden truths cannot both be true. These are reported with `🔒` and never quoted in full.
6. **Stale hypotheses.** `## Uncertain` lines that a later event confirmed or refuted and that were never promoted or removed; lines whose source session is more than three sessions old.
7. **System drift.** Entity files carrying the sections of a system the world no longer uses (after a system switch), or two files using different terms for the same mechanic.
8. **Loose ends** (info only): a log left `pending`, proposals older than three sessions, an approved event whose facts never reached a world file.

## Output
A prep-mode report, grouped, strongest first. Severity is about certainty, not importance:

- `⚠` **contradiction:** two canon statements cannot both be true. It needs a decision.
- `❓` **drift:** probably stale, but the GM may have meant it (a one-way link, an old uncertain line).
- `ℹ` **info:** no decision needed.

```
🔍 Consistency check: Vespera (37 files, 6 sessions)
⚠ Contradictions (2)
1. [[it-recep]] died in S003-E4, but [[the-cemetery]] still lists him as keeper
   → (a) the file is stale: drop the link (b) he survived: fix the event's record
2. IR2 "the dead never rise on their own" vs [[marie]] § Background: "rose from her grave alone"
   → (a) it only looks like a violation (b) a costly exception (c) the rule changes
❓ Drift (3)
4. [[gabriel-belmont]] links [[order-of-light]], the group does not link back
   → (a) add the reciprocal link (b) one-way is correct
5. S002-R3 "the town blames the hunter" is still uncertain; S004-E2 confirmed it
   → (a) promote to canon (b) keep uncertain (c) drop
ℹ Info (2)
- session-005 is still pending: 3 decisions wait at /close-session
- proposals/brother-gabriel.md has been open for 4 sessions
```
Each finding is one line plus one `→` line of options. At most 10 findings per run; the report says how many more there are, and `/check-consistency <area>` narrows it (a folder, a file, `rules`, `arcs`, `links`, `uncertain`).

## It writes nothing by itself
Running it is always safe: a pure report. When the GM answers items ("1 a, 2 b, 4 a"), the answers do **not** become edits here. They are appended to the open session log — a `mode: prep` log is created if none is open, as `/ripples` already does — as open decisions `D<n>`, which `/close-session` applies in group A with the usual approval.

So canon still has exactly one writer. The GM can also ignore the report and fix the files by hand; nothing in the world depends on this command having run.

## Never
- Edit a canon file, "clean up" anything, or treat a finding as settled.
- Quote a `## 🔒 Hidden Truth` or an arc's `## 🔒 Endgame` in the report; name the file and say what it conflicts with.
- Report style, tone, missing content or anything fill mode covers.
- Guess at a contradiction from prose alone: a finding must point at two places in two files.

## Changes
- `plugin/skills/check-consistency/SKILL.md` (new).
- `AGENTS.md`: write permissions (the open session log only, and only `D<n>` items the GM asked for).
- `plugin/skills/close-session/SKILL.md`: group A also carries consistency decisions.
- `plugin/skills/lexis-help/SKILL.md`, `plugin/README.md`: the command list.
- Docs: `commands.md`, `getting-started.md`, `installation.md` (the command count), `README.md`, `ROADMAP.md`.

## Decisions
1. **Missing belongs to fill mode, contradicting belongs here.** One file listing is never scanned twice for the same thing.
2. **Severity is certainty, not importance:** `⚠` two statements that cannot both be true, `❓` probably stale, `ℹ` no decision.
3. **Every finding names both sides.** A finding the GM cannot check in one line is noise.
4. **It writes nothing; answers become `D<n>` for `/close-session`.** Canon keeps one writer, and running the check stays risk-free.
5. **At most 10 findings per run**, with an area argument to narrow it. A 60-item report does not get read.
