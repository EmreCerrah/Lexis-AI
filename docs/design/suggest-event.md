# Design: /suggest-event (the scene has stalled)

Status: **accepted** (decisions below).

## Goal
The table goes quiet. The players have searched the room, asked the obvious questions and nobody knows what to do next. The GM needs one thing to happen now, and it has to come from the world they already wrote, not from nowhere.

`/suggest-event` gives 3 options for what could happen in the next minute. The GM picks one, says it at the table, and reports it with `/event` in their own words. Nothing is recorded.

It is the mirror of `/event`: `/event` takes what happened and finds the consequences; `/suggest-event` takes the world and finds something that could happen.

## What exists today
- `/event` records an event and lists ripples (`R<k>`), which are **hypotheses**: things that could follow but have not happened.
- `/ripples` adds more of them through lenses.
- A world has active hooks (with `## 🔒 Hidden Truth`), active arcs (with a current step), NPCs and groups with links, and `## Uncertain` lines.
- Nothing turns any of that back into "so what happens now?".

## Where the options come from
Every option names its source, so the GM can tell at a glance whether it fits. One option from each class, in this order:

| Class | Source | Example |
|---|---|---|
| **Someone acts** | an NPC or group linked to the scene, acting on a goal already in their file | the gravekeeper comes back for the thing he left behind |
| **The world answers** | a ripple from the log that has not happened yet (`R<k>`), the place itself, or the time of day | the cemetery gate is locked from outside; someone chained it while they were inside |
| **Something arrives** | news, a messenger, a sound, an object: information from one hop along `## Links` | a boy runs in: the troupe's wagon is leaving tonight |

Preferring an unspent `R<k>` matters: the GM already saw that ripple and half-expects it, so it lands as the world keeping its promises rather than as a new idea.

## Output
Table mode, three options, two lines each:
```
🎲 What could happen now
1. The gravekeeper comes back for the shovel he dropped — [[gravekeeper]], R4
   → asks: do they hide, or meet him?
2. The gate is chained from outside; the fog muffles the footsteps leaving — [[old-cemetery]]
   → asks: who locked them in, and how do they get out?
3. A boy with a theater ticket in his hat runs past the wall — [[the-yellow-curtain]] 🔒
   → asks: follow him, or stay and finish?
```
- One line of fiction, concrete and immediate: it happens on screen, now, in this place.
- The **source** in brackets: the file, plus `R<k>` when it is an unspent ripple.
- `→ asks:` what it puts to the players. Never a task they must do; always a choice they can refuse.
- `🔒` marks an option the GM should read before using, because it touches a hidden truth.

Prep mode: 5 options, each with its likely ripples, and a note when an option would make an arc step possible.

## Rules
- **Nothing is recorded.** Not the log, not canon. The option becomes real only when the GM makes it happen and reports it with `/event`. Nothing is written as `G<n>` either: an option that the GM does not use leaves no trace.
- **It never resolves what the players are doing.** No rolls, no outcomes, no "the check fails".
- Every option obeys every `IR<n>`, reveals no `## 🔒 Hidden Truth` and contradicts no arc `## 🔒 Endgame`.
- **It never makes an arc step mandatory.** An option may make the next step *possible* (someone who knows something walks in); it never drags the players there, and never blocks what they chose to do.
- **No new canon.** An option may use an unnamed person (a boy, a guard) but never invents a named NPC, place or group; for that the GM uses `/create-lore`.
- If the world has nothing to draw on (no hooks, no links, an empty log), it says so in one line and suggests `/world-init` or `/create-lore` instead of inventing a world.

## What each command keeps
- `/event`: what happened → consequences. Unchanged.
- `/ripples`: more consequences of a recorded event. Unchanged.
- `/suggest-event`: the world → what could happen. Reads the ripples; writes nothing.
- `/create-scene`: the players enter a place → the place, who is there, what to do. `/suggest-event` is for a place they are already in, when the three interactions are used up.

## Changes
- `plugin/skills/suggest-event/SKILL.md` (new).
- `AGENTS.md`: write permissions (writes nothing, like `/help-skills`).
- `plugin/skills/lexis-help/SKILL.md`, `plugin/README.md`: the command list.
- Docs: `commands.md`, `getting-started.md`, `installation.md` (the command count), `README.md`.

## Decisions
1. **Three options, one per source class**, not three variations of one idea: the GM should be choosing between different kinds of pressure.
2. **Every option names its source.** An option the GM cannot trace is an option they will not use.
3. **Unspent ripples come first.** They are already half-promised to the table.
4. **Nothing is logged, not even as a proposal.** An unused suggestion is noise at `/close-session`; a used one arrives as an `E<n>` anyway.
