---
name: create-scene
description: Build a ready-to-run scene when the players enter a place in a Lexis-AI world - the GM gives the place and a few notes, the AI fills the gaps - a read-aloud description, who is there, 3 interactions each with the world's game system check and suggested difficulty, one scene tag, and GM-only notes. Use when the GM asks for a scene, a room, a place description or what the players find there ("create-scene", "describe the place", "they enter the tavern", "sahne kur", "mekânı anlat").
---

# /create-scene

The players walk into a place. The GM gives the place and a few notes; this command fills the gaps and hands back one scene ready to run. A hybrid command: **fiction first, then a system line for each interaction.** It is a proposal: logged, not canon. You suggest checks and difficulties; the GM decides and rolls nothing through you.

```
/create-scene <place> [GM notes]
/create-scene old-cemetery night, thick fog, the gravekeeper is gone, one grave dug open
```

## Read first
`../../AGENTS.md`; `world.md` (`system:`, tone, language, `## Immutable Rules`); `system.md` (Homebrew overrides the standard); `../../systems/<system>/pack.md` (glossary), `resolution.md` (its `## Answer shape for /help-skills`), `scene-tags.md` and `obstacles.md`; the place's file in `locations/`; every file one step along its `## Links` (people, groups, hooks) and those hooks' `## 🔒 Hidden Truth`; every `arcs/*.md` with `status: active`; the open session log (where people are now, who died, what changed here); the PCs' files if named.

## Steps
1. **Place.** Match the GM's words to a file in `locations/` (slug, name or alias). No match → a new place: build it from the GM's notes alone, and add one line offering `/create-lore` to make it a lasting place. Never write the place file.
2. **The GM's notes are fixed.** Use every detail the GM gave, as given. Fill only what they left open, and never contradict the place file, the session's events, an `IR<n>`, a `## 🔒 Hidden Truth` or an arc's `## 🔒 Endgame`.
3. **Read-aloud description:** 2-4 sentences in the world's voice and language, through the senses (what they see, hear, smell), ending on something that invites action. Nothing in it may be 🔒.
4. **Who is here** (at most 3): people linked to the place (lives, works, guards, hides there) and members of linked groups, unless the files or the log put them elsewhere or dead. The GM's notes win ("the gravekeeper is gone"). For each: what they are doing now and their attitude toward the PCs, in the pack's terms if it has them (5E: Friendly / Indifferent / Hostile). A new person you add is marked `new:`.
5. **3 interactions:** varied (something to examine, someone to talk to, something to use or risk). For each:
   - **What:** one line of fiction.
   - **Can lead to:** one short line, open-ended (a clue, a person, a danger). Never the only way forward, never the next arc step as an obligation.
   - **System line:** from the pack's `resolution.md` answer shape: the kind of roll, 1-2 fitting traits (skill, approach, action …) and a suggested difficulty on the pack's scale (DC, ladder, position/effect …). If nothing is uncertain, write `no roll` and what simply happens. A hazard or a lock uses the pack's `obstacles.md` terms; for a full obstacle, point to `/create-obstacle`.
6. **One scene tag** for the whole place, from `scene-tags.md`, two-sided (helps the PCs one way, hinders them another).
7. **🔒 GM:** 1-2 lines: what is hidden here and how it ties to a hook (the place may *guard* a hidden truth, never reveal it), and a complication if the players linger.
8. **🧭** (only if it applies): which interaction could open the current step of an active arc. It makes the step *possible*, never mandatory.
9. **Log:** append to `## Generated` of the open log (create the log if there is none): `- G<n> scene (E<k>|—): [[<place>]] · "<short title>" · tag: "<scene tag>" · new: <new people or features | —>`.

## Output (table mode, in the output language)
```
🎬 G<n> · <Place> — <short title>
<read-aloud, 2-4 sentences>
Here:
- [[npc]] — what they are doing · attitude
Interactions:
1. <what> → <can lead to> · <system line, e.g. Wisdom (Perception) DC 15 | Careful, Fair (+2) | Survey, risky/standard>
2. …
3. …
✨ <scene tag> — helps: … · hinders: …
🔒 GM: …
🧭 [[arc]]: interaction 2 can open step <n>   (only if it applies)
```
Every block stays within the table-mode limit (`AGENTS.md` §9): one line per bullet.

## Prep mode
`prep` / `detailed`: 5-6 interactions, sub-areas (rooms, floors) each with a one-line description, each NPC's 🔒 motive and what they know (an inference from `## Links`, marked as such), a complication if the players linger, and the scene tag's full block from `scene-tags.md`.

## Never
- Write into a canon file. Lasting details (a feature of the place, a new person) reach world files only through `/close-session`.
- Decide what the players find on a roll, set the final difficulty, or roll.
- Put a 🔒 detail in the read-aloud text, or hint at a hidden truth or an arc endgame.
- Place a person here whom the log or the files put elsewhere, or who is dead.
- Bring in places, people or creatures from a published setting.
