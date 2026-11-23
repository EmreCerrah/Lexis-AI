# Design: /world-status (the one-minute briefing)

Status: **accepted** (decisions below).

## Goal
It is two weeks since the last session. The players are sitting down, and the GM has to remember where everything was: what happened last time, which hooks are live, which arc step is next, what they promised to decide, and which secret is about to matter.

`/world-status` reads the world and gives that briefing in about fifteen lines. It writes nothing.

## Its own command, not `/lexis-help status`
The roadmap left this open. Decision: **a separate command.**

A skill is chosen by its `description`. `/lexis-help`'s description is about *which commands exist and how to call them*; folding a briefing into it would mean one description doing two unrelated jobs, and "where were we last time?" would not reliably reach it. A dedicated skill triggers on what the GM actually says at the table.

The boundary stays sharp:

| | `/lexis-help` | `/world-status` |
|---|---|---|
| Question | what can I run, and what next? | where is my campaign right now? |
| World part | 3 lines: name, system, open session, arcs | the full briefing: last session, who is in play, hooks, arcs, open decisions, loose threads |
| Also covers | the command list and per-command help | nothing else |

And against `/check-consistency`: that one reports what *contradicts* what; this one reports *state*. Neither judges the other's findings.

## What the briefing contains
In this order, because it is the order a GM needs it:

1. **Header:** world, system, language, and which session comes next.
2. **Last session in 3 lines:** what the players did and where it left them, from the last closed or pending log's events — not every event, the ones that changed something.
3. **Who is in play:** the PCs, plus the NPCs and groups recent events touched, each with one line of what they are doing now.
4. **Active hooks:** every hook with `status: active`, one line of its `## Surface` only. A count of the waiting ones.
5. **Arcs:** each active arc as `step <n>/<total>`, its current hook, and that step's `Moves on when:` — the single most useful line in the whole report.
6. **Waiting on you:** unanswered `D<n>` in any `pending` log, proposals in `proposals/`, and `## Uncertain` lines a GM said they would settle.
7. **Loose threads:** up to 3 unspent ripples (`R<k>` no later event made real) — the things the table half-expects.
8. **Next:** one line, the same logic `/lexis-help` uses (`/close-session` if a log is pending, `/event` when play starts).

## Secrets are opt-in
A GM preparing alone wants the hidden truths of the active hooks; the same GM with players reading over their shoulder does not. So the report never prints them by default. It ends with a count:

```
🔒 2 hidden truths and 1 arc endgame are in play — /world-status secrets to see them
```

`/world-status secrets` adds a final `🔒 Not for the table` block with one line per secret. That is the only way they appear, and the block is always last so it can be scrolled past.

## Output
```
📋 Vespera · 5E SRD · Türkçe · next: session-006
Last session (005, pending):
- They opened the sealed crypt and found the coffins empty
- Gabriel burned the ledger before the Order could read it
In play:
- [[gabriel-belmont]] (PC) — hunting the troupe · [[marie]] (PC) — recovering
- [[order-of-light]] — looking for the ledger they now cannot read
Active hooks (2 active, 3 waiting):
- [[the-power-of-words]] — a printer's apprentice is selling lines nobody wrote
🧭 Arcs:
- The Yellow Curtain: step 2/4 [[the-yellow-curtain]]
  moves on when: the troupe's patron is exposed or the play is performed
Waiting on you:
- session-005 is pending: 3 decisions at /close-session
- proposals/brother-gabriel.md, open 4 sessions
Loose threads:
- R7 the mayor orders the cemetery closed "for repairs" (not yet happened)
🔒 2 hidden truths and 1 arc endgame are in play — /world-status secrets to see them
Next: finish session-005 with /close-session, then /event when play starts
```

At most 2 lines per group in table mode; `prep` lifts that to 5 and adds why each item matters. A group with nothing is left out entirely rather than printed empty.

## Never
- Write anything: no log, no canon file, no proposal.
- Print a `## 🔒 Hidden Truth` or an arc `## 🔒 Endgame` without the `secrets` argument.
- Decide anything, or re-answer an open decision. It reports what waits; `/close-session` resolves it.
- Report contradictions or missing content (that is `/check-consistency` and `/world-init`), or repeat the command list (`/lexis-help`).

## Changes
- `plugin/skills/world-status/SKILL.md` (new).
- `AGENTS.md` §6: writes nothing.
- `plugin/skills/lexis-help/SKILL.md`: the command list, and its `Next` line may point here at session start.
- Docs: `commands.md`, `getting-started.md`, `installation.md` (the command count), `README.md`, `ROADMAP.md`.

## Decisions
1. **Its own command, not a mode of `/lexis-help`.** A skill is found through its description; one description cannot carry both jobs.
2. **The arc's `Moves on when:` line is the heart of the report.** It is the one thing a GM cannot reconstruct from memory.
3. **Secrets are opt-in** (`/world-status secrets`), always last, never printed by default.
4. **Empty groups are omitted.** A briefing with six "none" lines is not a briefing.
