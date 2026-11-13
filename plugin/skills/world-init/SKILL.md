---
name: world-init
description: Create a new Lexis-AI RPG world folder (world.md, system.md, entity files, hooks) through a short tiered interview with the GM. Use when the GM wants to start, set up or initialize a new campaign world ("world-init", "new world", "set up the world", "yeni dünya kur"), or to complete an existing world between sessions: fill deferred or empty parts, add missing characters, create a missing system.md, or switch the world to another installed game system ("complete the world", "eksikleri tamamla", "sistemi değiştir").
---

# /world-init

Sets up a world in tiers so the GM is never forced to answer everything up front. This is a prep-mode command: approvals inside this conversation are fine.

## Read first
1. `../../AGENTS.md` (core rules), `../../references/world-layout.md`, `../../references/knowledge-model.md`.
2. Templates: `../../templates/*.md`. System packs: every `../../systems/*/pack.md` (to list them); after the GM chooses, that pack's `init-questions.md` and `entities.md`.

## 0. Where
- Target folder = the current working directory unless the GM names another. It must not be inside the plugin folder.
- If the target already has `world.md`, do **not** overwrite it and do not run steps 1-6: go to **Fill mode** below.

## Fill mode (existing world): complete what is missing
Brings an existing world up to date without overwriting anything the GM wrote. Run it when the GM calls `/world-init` in a world, or asks to "complete the world", "fill in what's missing", "change the system" ("eksikleri tamamla", "sistemi değiştir").

### F1. Scan
Read `world.md`, `system.md`, the world's `AGENTS.md`, every file in `characters/ locations/ groups/ hooks/ arcs/ proposals/`, and the open or pending session log if any. Build the gap list:
1. **System**
   - `system:` missing in `world.md` → the world runs on `fate-accelerated` (`AGENTS.md` §10).
   - `system:` names a pack that is not installed (no `../../systems/<id>/pack.md`) → offer the installed packs by `name`, or re-running `lexis-setup` to add that system.
   - The GM asks for a different system → a **system switch** (F3).
2. **`system.md`** missing, or its `system:` differs from `world.md`'s.
3. **World `AGENTS.md`** missing.
4. **Folders** missing: `characters/ locations/ groups/ hooks/ sessions/ proposals/` (`arcs/` only when an arc is written).
5. **Empty parts:** an empty `characters/`, `locations/`, `groups/` or `hooks/` folder; empty sections of `world.md` (Intro Scene, Opening Scene, World Notes, Religion and Creation …).
6. **Deferred items:** `## Uncertain` lines marked `world-init, deferred`, in any file.
7. **Entity files** without the system sections the pack's `entities.md` lists for them (e.g. Fate `## Aspects` in a 5E world, or no `## Stat Line` on a fighting NPC), or without `## Links`.
8. **Hooks** without `## 🔒 Hidden Truth`, or with `{{GM writes this}}` still in it.
9. **Arcs:** chains written as `- [[next-hook]] — next link in the chain` whose hooks are in no arc yet → offer to turn each chain into an arc (one item per chain); arcs without a `## Premise`; arc steps whose hook file is missing; a hook in two arcs.
10. **Info only:** dangling `[[links]]`, open proposals in `proposals/`.

### F2. One list, one answer
Silently create the missing folders (4) and the world `AGENTS.md` (3, from `templates/world-AGENTS.md`); report them, don't ask. Show everything else as **one numbered list** in the output language, grouped (System · World · Characters · Locations · Groups · Hooks · Arcs · Entity files · Info), each item with a suggested action:
```
🧩 Missing in <world name>
System
1. system.md missing → ask the 5E setup questions and create it
World
2. Opening Scene empty → you write / I suggest / later
Characters
3. characters/ is empty → name the PCs and NPCs you want / I suggest / later
4. "Gabriel's grandson not decided yet" (deferred) → fill now / later
Entity files
5. 3 files have no system sections (groups/order-of-light, …) → add / leave
Arcs
6. Chain empty-coffins → yellow-curtain → masters-letter has no arc → make an arc / leave
Info
- proposals/: 2 open proposals, decided at /close-session
```
The GM answers once in plain words ("all", "1, 3: Gabriel, Ayşe (PC); rest later", "3 you suggest"). Unmentioned items stay as they are.

### F3. Fill
Work through the chosen items, batching questions into as few messages as possible:
- **`system.md`:** ask the pack's `init-questions.md` as one batch (unanswered = standard) and write it from the pack's `templates/system.md`.
- **System switch:** set `system:` in `world.md` to the new pack id. Rewrite `system.md` from the new pack's template and its init questions, keeping the old `## Homebrew` and `## Uncertain` text and adding `- Homebrew written for <old pack>, check it _(source: world-init, system switch)_` to `## Uncertain` when there was any. Then offer item 7 for every entity file.
- **Entity system sections (7):** draft the pack's sections from each file's fiction (who they are, what they do, where it is) and replace the old pack's sections. Show all drafts in one message; write only the accepted ones. Never touch the fiction sections, `## 🔒 GM Notes` or `## Links`.
- **Characters, locations, groups, hooks (5, 6):** as in steps 2-3: content from the GM is canon; "you suggest" → drafts shown in one message; "later" → a deferred `## Uncertain` line. The GM may mark a character **PC** or **NPC**; use the pack's PC or NPC sections from `entities.md`. Every new entity gets `## Links` to the existing files. A new hook gets its own hidden truth, consistent with the others and with every `IR<n>`.
- **Arcs (9):** for a chain, draft `arcs/<slug>.md` from `../../templates/arc.md`: a premise from the hooks' surfaces, the steps in chain order, a `Moves on when:` for each step that fits its hook. Ask whether the GM wants an endgame; write one only on a yes or "you suggest", consistent with every hidden truth. In each step's hook, replace the `next link in the chain` line with `- [[<arc-slug>]] — arc, step <n>`. Set the arc's `status:` from its hooks (any hook `active` or `resolved` → `active`; all `resolved` → `resolved`; else `waiting`). Show the draft; write only if accepted.
- **Missing `## Links` (7):** propose links from what the files already say; write only accepted ones.
- **Hidden truths (8):** draft one, or leave `{{GM writes this}}`, as the GM chooses.
- **Deferred items:** write the answer into the right canon section and delete the matching `## Uncertain` line.

Fill mode only adds or completes. It never rewrites canon fiction, never renumbers `IR<n>` and never touches `sessions/`.

### F4. Report
```
✅ World completed: <name>
- Created: system.md, characters/gabriel.md, characters/ayse.md …
- Updated: world.md (Opening Scene), groups/order-of-light.md (system sections)
- Left uncertain: <list or —>
- System: <pack name>
```

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
9. System: only the packs chosen at install are present (`../../systems/*/pack.md`). If there is exactly one, don't ask: state it in this item ("System: <name>, installed") and use it. If there are several, list them by `name`. If the GM wants a system that is not installed, say so: they can re-run `lexis-setup` and choose it, or pick an installed pack and put the differences under Homebrew.

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
- **Arcs:** after the hooks are drafted or given, ask once: "Do any of these hooks form a chain?" A yes (or "you suggest") → an arc draft as in fill mode F3, written to `arcs/<slug>.md`.
- Draft briefly, in the world's tone and language, obeying every `IR<n>`. For hooks, write `## Surface`, `## 🔒 Hidden Truth` and `## Links`; never leave the hidden truth for later. If the GM wants to write it themselves, leave `{{GM writes this}}` and add an `## Uncertain` line.
- Show all drafts in one message; the GM accepts, edits or drops each. Write only accepted drafts.
- Every entity replaces the template's `{{system sections}}` line with the sections the pack's `entities.md` lists for it.
- Every entity gets `## Links` at creation (character↔group membership, character↔location presence, hook↔its people and places, hook→next hook). "Who might know" runs on these links, so don't skip them.

## 4. System pack
Ask the questions in `../../systems/<system>/init-questions.md` as one batch. Unanswered = standard. Fill `system.md` from `../../systems/<system>/templates/system.md`.

## 5. Write
Create (see `references/world-layout.md`):
- `world.md` from the template, with `language:` and `system:` (the pack `id`) set and Immutable Rules numbered `IR1…`; `system.md`; `AGENTS.md` from `templates/world-AGENTS.md`
- `characters/`, `locations/`, `groups/`, `hooks/`, `sessions/`, `proposals/` (create the empty folders too); `arcs/` only if an arc was accepted
- one file per accepted entity, with slugs following the layout rules
Do **not** create a session log; `/event` opens one.

## 6. Report (short, in the output language)
```
✅ World created: <name>
- Files: world.md, system.md, <n> characters, <n> locations, <n> groups, <n> hooks, <n> arcs
- Immutable Rules: IR1 … IRn
- Left uncertain: <list or —>
- Next: report what happens at the table with /event; at the end of the session run /close-session.
```
