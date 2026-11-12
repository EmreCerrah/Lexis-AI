# Commands

Lexis-AI has eight commands. Call them as described in [Getting started](getting-started.md) (`/event`, `/lexis-ai:event`, `$event` … depending on your tool).

| Command | Use it | Writes |
|---|---|---|
| [`/world-init`](#world-init) | once, to create a world (or to fill in what you deferred) | the world files |
| [`/event`](#event) | right after something happens at the table | the session log |
| [`/ripples`](#ripples) | when you want more consequences of an event | the session log |
| [`/help-skills`](#help-skills) | for a rules question | nothing |
| [`/create-aspect`](#create-aspect) | for a scene tag (aspect, feature, circumstance …) | the session log |
| [`/create-obstacle`](#create-obstacle) | for a lock, trap, puzzle, hazard or social obstacle | the session log |
| [`/create-lore`](#create-lore) | for a new NPC, place, group or story hook | `proposals/` + the session log |
| [`/close-session`](#close-session) | at the end of the session | the world files (only what you approve) |

## Two modes
- **Table mode** (default): short. At most 5 bullets per block, no introduction, no approval questions in the middle of the game.
- **Prep mode:** add `prep` or `detailed` (in any language: `hazırlık`, `vorbereitung` …) to a command for more depth between sessions: `/ripples E2 prep`.

## Ground rules every command follows
- It never rolls dice, never decides whether something succeeds and never decides what an NPC really knows. It suggests; you decide.
- It never invents canon. Canon comes from you: reported events, `/world-init` answers, approvals at `/close-session`.
- Its suggestions never break an **Immutable Rule** of your world and never reveal or contradict a hook's **🔒 Hidden Truth**.
- It answers in your world's language; file names, headings and game terms stay as they are.
- It never brings content from a published setting into your world (places, gods, factions, famous characters). Your world is the only setting.

---

## /world-init
Creates a world in the current folder, in three tiers so you are never forced to decide everything up front.

1. **Core** (one message): name, concept, genre, language, tone, inspirations, fantasy level, Immutable Rules, game system. Name, concept, at least one Immutable Rule and the system are required; the rest may be "later".
2. **Optional parts**: world notes, characters, locations, groups, religion and creation, hooks. For each: give content, say "later", or "you suggest".
3. **Drafts**: everything you asked it to suggest is shown in one message; you accept, edit or drop each one.
4. **System questions**: the setup questions of your game system ([Game systems](game-systems.md)). Unanswered means standard rules.

It writes `world.md`, `system.md`, a short `AGENTS.md`, the folders `characters/ locations/ groups/ hooks/ sessions/ proposals/` and one file per accepted entity. It does not open a session log; `/event` does that.

**Fill mode:** run `/world-init` again in an existing world. It never overwrites; it lists what you deferred and fills in only what you answer now.

## /event
```
/event <what happened, in one or two sentences>
```
1. Records the event as `E<n>` in the open session log (it opens `sessions/session-NNN.md` if none is open).
2. Checks it against the Immutable Rules. A conflict is **recorded anyway** and flagged `⚠ rule conflict (IR<n>)` with three options: (a) it only looks like a violation, (b) a costly exception, (c) the rule changes. You can answer now or at `/close-session`.
3. If players try to break a hook's hidden truth, it gives 1-3 in-world reasons why they can't.
4. Lists 3-5 **ripples** (possible consequences, `R<k>`), each tied to the file it would change.
5. Lists **who might know** (one step along `## Links`; never written to a file).
6. Proposes up to 2 new links (`L<j>`) if the event obviously creates a relation.

```
📌 E2 recorded (session-001) · Lens: power-politics
Ripples (hypothesis):
- R4 The mayor orders the cemetery closed "for repairs"
- R5 🔒 The troupe's patron pays the mayor a visit
Who might know:
- [[mayor-harkin]] — may hear (owner → [[old-cemetery]])
⚠ D1: IR2 conflict — (a) apparent (b) costly (c) rule changes? Now or at session end.
```

## /ripples
```
/ripples [E2 | S003-E2 | a description] [prep]
```
Takes an event that is already recorded and looks at it through **lenses**: power and politics, society and rumor, faith and the supernatural. It adds new ripples (never repeating old ones) and traces how word could spread 2-3 steps, marking how the news weakens (`as a rumor`, `distorted`). Default event: the last one. In prep mode it uses every fitting lens and explains each ripple.

## /help-skills
```
/help-skills <what the player wants to do>
```
A rules answer in your system's terms. It records nothing.
- the kind of roll (check, action roll, test …),
- 2-3 fitting skills, abilities, approaches or actions, each with the fiction that would justify it (the player's description decides),
- a suggested difficulty with the reason,
- what the dice results would mean, where the system has degrees or bands,
- useful options: a matching feature or stunt from the character's file, a scene tag, a meta-currency use.

Examples per system: [Game systems](game-systems.md).

## /create-aspect
```
/create-aspect [scene or situation]
```
One **scene tag** for the current scene, in your system's form: a Fate aspect with an invoke and a compel; a 5E scene feature with its rule effect; a circumstance, situational modifier, DR shift, position/effect factor or clock in other systems. Every tag is two-sided: one way it helps the PCs, one way it hinders them. Logged as `G<n>`; a lasting one (a place's feature) can become canon at `/close-session`.

## /create-obstacle
```
/create-obstacle [what you need]
```
One obstacle, **fiction first**:
- what it is (to read aloud),
- why it is here (from the place, a group or a hook),
- 🔒 what it guards (never revealing a hidden truth),
- for a puzzle, the clue chain, solvable without dice too,

then your system's block: the suggested difficulty, fitting checks, what success and failure levels mean, trap or hazard numbers, and 2-3 "success at a cost" options. Logged as `G<n>`.

## /create-lore
```
/create-lore <what you need: an NPC, a place, a group, a hook>
```
One persistent entity, written to `proposals/<slug>.md` from the world's template, with the system sections your game needs (aspects in Fate, a stat line in 5E …) and 1-4 links to existing files. A new hook always gets its own hidden truth, consistent with the others. It becomes canon only if you approve it at `/close-session`.

## /close-session
```
/close-session
```
Builds **one numbered list** of everything from the session and waits for one answer. Nothing is written before you answer.

| Group | What you decide |
|---|---|
| A. Open decisions | rule conflicts and broken hidden truths: (a)/(b)/(c) |
| B. Events | the concrete changes the events make to world files (already canon; suggested ✅) |
| C. Ripples | ✅ it happened · ❓ uncertain (goes to `## Uncertain`) · ❌ drop |
| D. Link proposals | ✅ / ❌ |
| E. Generated items | was it used? lasting scene tags can join a place; lore can become canon (✅), stay a proposal (⏸) or be deleted (❌) |
| F. Old uncertain items | promote, remove or keep what this session confirmed or refuted |
| G. Info | dangling links, unknown names |

Answer in plain words: `all as suggested`, `3 ✅, 5 ❌, rest as suggested`, `1 c: <new rule text>`, `defer 4-6`. Deferred items keep the log `pending` and come back next time.

```
✅ Session 001 closed (closed)
- Updated files: [[old-cemetery]], [[gravekeeper]]
- Added to uncertain: 3 · Rejected: 2 · Proposals kept: 1
- Rule change: —
```
