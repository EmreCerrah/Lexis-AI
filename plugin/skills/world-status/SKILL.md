---
name: world-status
description: A one-minute briefing on where a Lexis-AI campaign stands before a session - what happened last time, who is in play, which hooks are active, each arc's current step and what moves it on, what waits for a decision, and which threads are still loose. Records nothing; hidden truths only with the "secrets" argument. Use when the GM asks where the campaign is, what happened last session, what is active or what they should remember before play ("world-status", "where were we", "catch me up", "what's active", "neredeydik", "durum ne").
---

# /world-status

The players are sitting down and the GM has to remember two weeks of campaign in a minute. This reads the world and gives that briefing. It writes nothing.

```
/world-status [secrets] [prep]
```

For what commands exist, that is `/lexis-help`. For what contradicts what, `/check-consistency`. For what is missing, `/world-init`. This command reports **state**.

## Read first
`../../AGENTS.md`; `world.md` (`name`, `system`, `language`); the newest logs in `sessions/`: the open one, any `pending` one, and the last `closed` one (their events, unanswered `D<n>`, and every `R<k>`); `hooks/*.md` (`status`, `## Surface`); `arcs/*.md` with `status: active` (`## Steps`, each step's `Moves on when:`, and the status of the step hooks); the `characters/`, `locations/` and `groups/` files the recent events name; `proposals/`; `## Uncertain` lines in `world.md` and in the files recent events touched.
Read a hook's `## 🔒 Hidden Truth` or an arc's `## 🔒 Endgame` only to count them — unless the GM passed `secrets`.
If there is no world root, say so in one line and suggest `/world-init`.

## Build the briefing
In this order, because it is the order the GM needs it. **Leave out any group that has nothing**; never print an empty one.

1. **Header:** `<world> · <system pack name> · <language> · next: session-<NNN>` (the open log, or the next number if none is open).
2. **Last session:** up to 2 lines from the last closed or pending log — the events that *changed* something, not every event. Name the log and its status.
3. **In play:** the PCs with one line each on their situation, then the NPCs and groups the recent events touched, with what they are doing now. Leave out anyone the log put elsewhere or killed.
4. **Active hooks:** one line of `## Surface` for each hook with `status: active`, plus a count of the waiting ones. Never the hidden truth.
5. **Arcs** (`🧭`): each active arc as `step <n>/<total>`, its current hook (the first step whose hook is not `resolved`), and that step's `Moves on when:` text. This is the most useful line in the report; never drop it.
6. **Waiting on you:** unanswered `D<n>` in a `pending` log (with the count), proposals in `proposals/` and how many sessions they have been open, and `## Uncertain` lines the GM said they would settle.
7. **Loose threads:** up to 3 `R<k>` from the logs that no later `E<n>` made real.
8. **Secrets count:** one line — `🔒 <n> hidden truths and <n> arc endgames are in play — /world-status secrets to see them`. Omit it if there are none.
9. **Next:** one line, the same logic as `/lexis-help`: a `pending` log → `/close-session`; an open log → keep reporting with `/event`; otherwise `/event` when play starts; something clearly missing → `/world-init`.

## Output (in the output language)
```
📋 <header>
Last session (<NNN>, <status>):
- …
In play:
- [[pc]] (PC) — … · [[pc]] (PC) — …
- [[group]] — …
Active hooks (<n> active, <n> waiting):
- [[hook]] — <surface, one line>
🧭 Arcs:
- <arc>: step <n>/<total> [[current hook]]
  moves on when: …
Waiting on you:
- …
Loose threads:
- R<k> … (not yet happened)
🔒 <n> hidden truths and <n> arc endgames are in play — /world-status secrets to see them
Next: …
```
At most 2 lines per group. `prep` / `detailed`: up to 5 per group, each with one line of why it matters now.

## `secrets`
With `secrets`, append one final block, always last so it can be scrolled past:
```
🔒 Not for the table
- [[hook]]: <its hidden truth, one line>
- <arc>: <its endgame, one line>
```
Only the active hooks' truths and the active arcs' endgames. Without the argument this block does not exist.

## Never
- Write anything: not a canon file, not a session log, not a proposal.
- Print a hidden truth or an arc endgame without `secrets`, or hint at one in the hook line.
- Answer an open decision, or decide anything: the report says what waits, `/close-session` resolves it.
- Report contradictions or missing parts, or list the commands. Those are other commands' jobs.
