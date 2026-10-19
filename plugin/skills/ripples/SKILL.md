---
name: ripples
description: Deepen an already recorded event in a Lexis-AI world (folder with world.md) by running it through analysis lenses (politics, rumor, faith…) and tracing how knowledge of it spreads 2-3 hops with weakening; results are hypotheses appended to the open session log. Use when the GM asks for more consequences of an event ("ripples", "deepen E2", "what else follows from this", "yankılar").
---

# /ripples

Takes an event that is already in a session log and asks "what else follows?" through lenses. The output is **hypothesis** only.

## Read first
1. `../../AGENTS.md`. 2. `../../references/knowledge-model.md`. 3. `../../references/lenses/index.md`, then the lens files you pick.
4. World: `## Immutable Rules` in `world.md`, the event's files and their links up to 3 hops, linked hooks.

## Steps
1. **Find the event.** The argument may be `E2`, `S003-E2` or a description. Default: the last `E<n>` in the open log. If the event is not in any log, tell the GM in one line to record it with `/event` first, and stop.
2. **Mode.** Table (default) or prep (the GM writes `prep`/`detailed`, or no session log is open).
3. **Pick lenses.** Table: 2-3. Prep: every lens that scores. Skip lenses the event's existing `R<k>` lines already covered, unless the GM asks again.
4. **Per lens**, answer the lens's `Questions` against the world files. Keep only ripples that:
   - follow from a link, goal or rule in the files (or are clearly marked as a new guess),
   - obey every `IR<n>` and every `## 🔒 Hidden Truth`,
   - do not repeat an existing `R<k>`.
5. **Knowledge spread.** Run the algorithm to hop 2-3 using the lens's `Spread channels`, and label the weakening (`as a rumor`, `distorted`). Output only.
6. **Append** the new ripples to `## Ripples` of the open log (create one with `mode: prep` if none is open): `- R<k> (E<n> · <lens>) … · target: [[file]]`. New relations go to `## Link Proposals` as `L<j>`.

## Output (in the output language)
Table mode:
```
🌊 Ripples of E<n> · Lens: <a>, <b>
<lens a>:
- R4 …
- R5 …
<lens b>:
- R6 …
Spread:
- [[…]] — may hear it as a rumor (… → … → …)
```
At most 5 ripples in total in table mode. Prep mode: up to 4 per lens, each with one sentence of "why" and the file it would change, plus a short "Which hook does it move closer?" line (🔒).

## Never
- Edit canon files, or present a ripple as something that happened.
- Let a ripple explain or reveal a hidden truth; it may only move toward it.
