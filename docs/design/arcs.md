# Design: Arcs (chained hooks)

Status: **accepted** (decisions below).

## Goal
A GM often plans hooks as a chain: *Empty Coffins → The Yellow Curtain → The Master's Letter → The Gate of Carcosa*. Today the chain exists only as a `- [[next-hook]] — next link in the chain` line in each hook, and nothing watches it. An **arc** makes the chain visible:
- Lexis-AI knows which step is in play.
- It notices when an event moves the story on, skips ahead or puts a later step at risk.
- `/close-session` proposes the progress.

An arc is a **plan, not a fixed future**. The players stay free; the arc only tracks where the story is and helps the GM keep it coherent. Hidden truths stay fixed as before.

## What exists today
- `hooks/<slug>.md` has `status: waiting | active | resolved`, `## Surface`, `## 🔒 Hidden Truth`, `## Links`.
- Chains are written as `- [[next-hook]] — next link in the chain` in `## Links`.
- `/event` already checks the hidden truths of linked hooks.
- `/close-session` already proposes hook `status` changes from events.

## File: `arcs/<slug>.md` (type: arc)
```markdown
---
type: arc
name: The Yellow Curtain
status: active            # waiting | active | resolved | abandoned
---

# The Yellow Curtain

## Premise
What the arc is about, in one or two sentences. GM material, but nothing secret.

## Steps
1. [[empty-coffins]]
   - Moves on when: the players learn who emptied the coffins.
2. [[the-yellow-curtain]]
   - Moves on when: the troupe's patron is exposed or the play is performed.
3. [[the-masters-letter]]
   - Moves on when: the letter is read by someone who understands it.
4. [[the-gate-of-carcosa]]

## 🔒 Endgame
Where the story ends if nobody stops it. GM only; never revealed, never contradicted.

## Links
- [[order-of-light]] — opposes it

## Uncertain
```

Rules:
- **The order lives in the arc file only.** Each step is one hook, one line, plus an optional `Moves on when:` condition.
- **Status lives in the hook files only.** The arc does not repeat the hooks' status.
- **The current step is derived, never stored.** It is the first step whose hook is not `resolved` (knowledge-model: an inference). So the arc and the hooks can never disagree.
- **One hook belongs to at most one arc** (v1). A hook not in any arc works as today.
- **Hooks point back to their arc:** `- [[the-yellow-curtain]] — arc, step 2` replaces the old `next link in the chain` line. Old lines keep working: an arc-less chain is read as before.
- `arcs/` is a canon folder: only `/world-init` and `/close-session` write it, like `hooks/`.

## What each command does

### /event
After the hidden-truth check, an **arc check** against active arcs. At most one `🧭` line per arc, only when it applies:

| The event … | Line | Logged as |
|---|---|---|
| meets the current step's `Moves on when` | `🧭 [[arc]]: step 2 may be done → step 3 [[the-masters-letter]]` | `- Arc: [[arc]] step 2 → done?` under the event |
| touches a later step before its time | `🧭 [[arc]]: E3 jumps to step 4` | `D<n>` (arc skip) |
| removes something a later step needs (a key NPC dies, a place burns) | `🧭 [[arc]]: E3 threatens step 3` + 1-2 in-world ways the step can still happen | `D<n>` (arc threat) |

- Like the hidden truth, an arc never blocks an event. The event is canon; the arc adapts.
- The GM can answer a `D<n>` now or at `/close-session`:
  - **Skip:** (a) the skipped steps happen offstage, (b) reorder the steps, (c) drop the skipped steps.
  - **Threat:** (a) the plan holds, with an in-world reason, (b) the step changes (new text), (c) drop the step.

### /ripples
In prep mode, each ripple can carry `→ moves [[arc]] toward step N`. In table mode, at most one such note in total.

### /close-session
New group **C. Arcs**, after the events:
- **Step done:** the step's hook → `resolved`, the next step's hook → `active`. Suggested ✅ only when an event in this session met the step's `Moves on when`.
- **Last step done:** the arc → `resolved`.
- **Arc decisions:** answers to the arc `D<n>` items, written into the arc file (reordered or changed steps). A changed step adds `_(changed: S004)_` to the step line.
- **Abandoned:** the GM can mark an arc `abandoned`. Its hooks stay as they are.

### /world-init
- **New world:** in the Hooks part, after the hooks: "Do any of these form a chain?" A yes writes an arc file (premise, steps, optional endgame). Drafting one on "you suggest" is fine.
- **Fill mode:** finds `next link in the chain` lines that are not yet in an arc and offers to turn each chain into an arc (one list item per chain). It also lists arcs without a `Premise`, and steps whose hook file is missing.

### /create-lore
`/create-lore an arc from <hooks>` or `/create-lore an arc about <idea>` writes `proposals/<slug>.md` with `type: arc`. New hooks are written as separate proposals, each with its own hidden truth. Becomes canon at `/close-session` like any proposal.

### /lexis-help
The world block shows active arcs and their current step: `🧭 The Yellow Curtain: step 2/4 [[the-yellow-curtain]]`.

## Never
- Reveal an arc's `## 🔒 Endgame` or a hook's hidden truth, also not by hinting at the next step to the players.
- Push the players toward the next step. Suggestions can make it *possible* (a ripple, an NPC who knows something), never mandatory.
- Change an arc's order, steps or endgame outside `/close-session` or `/world-init`.

## Changes
- `plugin/templates/arc.md` (new).
- `references/world-layout.md`: `arcs/` folder and `type: arc`.
- `references/knowledge-model.md`: the current step is an inference.
- `AGENTS.md`: `arcs/` is canon, endgame is 🔒.
- Skills: `event`, `ripples`, `close-session`, `world-init`, `create-lore`, `lexis-help`.
- Docs: `world-files.md`, `commands.md`, the getting-started example.

## Decisions
1. **An arc file**, not links only: links have no place for the premise, the `Moves on when` conditions or the endgame.
2. **`## 🔒 Endgame` is optional.** If written, the AI uses it for consistency and never reveals it.
3. **No branches in v1:** a straight chain of steps. Alternatives (`2a` / `2b`) may come later.
4. **A skip opens a decision** (`D<n>`): (a) the skipped steps happen offstage, (b) reorder, (c) drop them. The GM answers now or at `/close-session`.
