---
name: world-init
description: Create a new Lexis-AI RPG world folder (world.md, system.md, entity files, hooks) through a short tiered interview with the GM. Use when the GM wants to start, set up or initialize a new campaign world ("world-init", "new world", "set up the world", "yeni dünya kur"), or to fill deferred sections of an existing world between sessions.
---

# /world-init

Sets up a world in tiers so the GM is never forced to answer everything up front. This is a prep-mode command: approvals inside this conversation are fine.

## Read first
1. `../../AGENTS.md` (core rules), `../../references/world-layout.md`, `../../references/knowledge-model.md`.
2. Templates: `../../templates/*.md`. System pack: `../../systems/<system>/pack.md` and `init-questions.md`.

## 0. Where
- Target folder = the current working directory unless the GM names another. It must not be inside the plugin folder.
- If the target already has `world.md`, do **not** overwrite it. Switch to **fill mode**: list the `## Uncertain` items marked `world-init, deferred` and the empty sections, ask which to fill now, and write only those (the canon sections the GM answers, then delete the matching `## Uncertain` line). Then stop.

## 1. Core (required): one message, all questions together
Ask in the default output language (`AGENTS.md` §3), numbered, with a one-line example for each. Accept short answers. Accept "you suggest" for any item and draft it (mark it as a draft and get a yes/no in the same conversation).
1. World name
2. Concept (1-3 sentences)
3. Genre
4. Language: the language the AI answers in for this world. Default: the default output language from `AGENTS.md`.
5. Narrative style / tone
6. Inspirations
7. Fantasy level
8. Immutable Rules: the fiction's unbreakable rules (e.g. "The dead never rise on their own", "Sunlight burns vampires"). At least 1.
9. System (MVP: Fate Accelerated)

Do not continue until 1, 2, 8 and 9 have answers. Others may stay empty → `## Uncertain`. If the GM picks a different language in item 4, switch to it from the next message on.

## 2. Deferrable: one message
Offer the list and say clearly that every item is optional: "later" puts a note in `## Uncertain`, "you suggest" gets a draft.
- World notes · Characters · Locations · Groups · Religion/creation · Hooks
For each answer:
- The GM gave content → write it as canon.
- "later" / no answer → add `- <section> not decided yet _(source: world-init, deferred)_` to `## Uncertain` in `world.md`.
- "you suggest" → generate it (step 3).

## 3. Generatable: drafts
For items the GM asked you to draft (intro scene, opening scene, characters, locations, groups, hooks):
- Draft briefly, in the world's tone and language, obeying every `IR<n>`. For hooks, write `## Surface`, `## 🔒 Hidden Truth` and `## Links`; never leave the hidden truth for later. If the GM wants to write it themselves, leave `{{GM writes this}}` and add an `## Uncertain` line.
- Show all drafts in one message; the GM accepts, edits or drops each. Write only accepted drafts.
- Every entity gets `## Links` at creation (character↔group membership, character↔location presence, hook↔its people and places, hook→next hook). "Who might know" runs on these links, so don't skip them.

## 4. System pack
Ask the questions in `../../systems/<system>/init-questions.md` as one batch. Unanswered = standard. Fill `system.md` from `../../systems/<system>/templates/system.md`.

## 5. Write
Create (see `references/world-layout.md`):
- `world.md` from the template, with `language:` set and Immutable Rules numbered `IR1…`; `system.md`; `AGENTS.md` from `templates/world-AGENTS.md`
- `characters/`, `locations/`, `groups/`, `hooks/`, `sessions/`, `proposals/` (create the empty folders too)
- one file per accepted entity, with slugs following the layout rules
Do **not** create a session log; `/event` opens one.

## 6. Report (short, in the output language)
```
✅ World created: <name>
- Files: world.md, system.md, <n> characters, <n> locations, <n> groups, <n> hooks
- Immutable Rules: IR1 … IRn
- Left uncertain: <list or —>
- Next: report what happens at the table with /event; at the end of the session run /close-session.
```
