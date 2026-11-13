---
name: event
description: Record something that just happened at the RPG table as canon in the open session log of a Lexis-AI world (folder with world.md), then list 3-5 short ripples as hypotheses and who might know. Use when the GM reports an event during play ("event", "this happened", "the players did X", "şu oldu"). Never edits canon world files.
---

# /event

The GM reports what happened. It is canon **instantly**. You record it, then show what might follow. Table mode: fast, short, zero approval questions.

## Read first
1. `../../AGENTS.md`. 2. `../../references/knowledge-model.md` (who-might-know algorithm). 3. `../../references/lenses/index.md`.
4. World: `world.md` (`## Immutable Rules`), every file the event names, the targets of their `## Links` (1 hop), and the hooks linked to them (`## 🔒 Hidden Truth`). Every `arcs/*.md` with `status: active` (its `## Steps`, `## 🔒 Endgame`) and the status of its step hooks.

## Steps
1. **Open the log.** Find the file in `sessions/` with `status: open`; if there is none, create the next `session-NNN.md` from `../../templates/session-log.md`.
2. **Resolve names.** Map the people and places in the GM's words to slugs. For an unknown name, keep the plain name in the log and add a link proposal or a `new: <name>` note in the output; do not stop to ask.
3. **Record (append to `## Events`)** exactly what the GM said, in neutral words. Do not add consequences here; the event block holds only what the GM reported:
   ```
   ### E<n> — <3-6 word title>
   - What happened: <what the GM described>
   - Place: [[…]] · Who: [[…]], …
   - Hook: [[…]] | —
   - Arc: [[arc]] step <n> → done? | skip → D<m> | threat → D<m>   (only if the arc check found something)
   - Rules: consistent | ⚠ rule conflict (IR<n>) → D<m>
   ```
4. **Rule check.** Compare the event with every `IR<n>`. If it breaks one, it is still recorded; add `D<m>` to `## Open Decisions` with the options (a) apparent violation (b) costly exception (c) rule change (see `AGENTS.md` §7).
5. **Hidden-truth check.** If the event threatens a hook's `## 🔒 Hidden Truth`:
   - players *trying* to break it → give 1-3 in-world reasons it fails (output only).
   - the GM says it really broke → add `D<m>` with the options (a) plan holds, the story reconnects (b) plan changes.
6. **Arc check.** For each active arc, find the current step (the first step whose hook is not `resolved`). At most one finding per arc, only when it applies:
   - the event meets the current step's `Moves on when` → `done?` in the event block (decided at `/close-session`; no `D`).
   - the event reaches a later step before its time → **skip**: add `D<m>` with (a) the skipped steps happen offstage (b) reorder the steps (c) drop the skipped steps.
   - the event removes something a later step needs (a key NPC dies, a place is destroyed, an item is lost) → **threat**: add `D<m>` with (a) plan holds, plus 1-2 in-world ways the step can still happen (b) the step changes (c) drop the step.
   Never block the event, never reveal the endgame, never steer the players toward the next step.
7. **Pick lenses** per the index: at most 2 (3 only if clearly needed). Open only those lens files.
8. **Ripples:** 3-5 short consequences, one line each, each obeying every `IR<n>` and every hidden truth. Append them to `## Ripples` as `- R<k> (E<n> · <lens>) <ripple> · target: [[file]]` (target = the file it would change if approved). GM-only ripples start with `🔒`.
9. **Who might know:** 1 hop per the algorithm, at most 5. Output only, never written.
10. **Link proposals:** if the event creates an obvious new relation (someone now owes, hates, saw or holds something), append `- L<j> (E<n>) [[a]] → [[b]] — <relation>` to `## Link Proposals`. At most 2.

## Output (table mode, in the output language)
```
📌 E<n> recorded (session-NNN) · Lens: <a>, <b>
Ripples (hypothesis):
- R1 …
- R2 …
- R3 …
Who might know:
- [[…]] — there
- [[…]] — may hear (member → [[…]])
<only if any:> ⚠ D1: IR2 conflict — (a) apparent (b) costly (c) rule changes? Now or at session end.
<only if any:> 🔒 Why they can't: …
<only if any:> 🧭 [[arc]]: step 2 may be done → next [[hook]] · or: E3 jumps to step 4 → D2 · or: E3 threatens step 3 → D2 (a) plan holds: … (b) changes (c) drop
```
No preamble, no restating the event, no "do you approve?". If the GM writes `prep` or `detailed`, use up to 3 lenses and add one sentence of reasoning per ripple.

## Never
- Edit `world.md`, `system.md` or any file in `characters/ locations/ groups/ hooks/ arcs/`.
- Write the who-might-know list to a file.
- Decide what NPCs actually know or do; ripples are possibilities.
