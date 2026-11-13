---
name: suggest-event
description: Offer 3 things that could happen right now in a Lexis-AI world when the scene has stalled - each one drawn from an NPC's goal, an unspent ripple, the place itself or news arriving, with its source named and what it asks of the players. Records nothing. Use when the GM asks what could happen next, says the scene is stuck, the players are stalling or nobody knows what to do ("suggest-event", "what now", "the scene stalled", "ne olabilir", "sahne tıkandı").
---

# /suggest-event

The table has gone quiet. This gives the GM **3 options for what could happen in the next minute**, each from the world they already wrote. The GM picks one, says it at the table, then reports it with `/event` in their own words. This skill records nothing and writes nothing.

```
/suggest-event [what is stalling, or where they are] [prep]
/suggest-event they have searched the whole crypt and won't leave
```

## Read first
`../../AGENTS.md`; `world.md` (`## Immutable Rules`, tone, language); the open or newest session log (the recent `E<n>`, and every `R<k>` that has not happened yet); the files of the place and people in the current scene and everything one hop along their `## Links`; `hooks/*.md` with `status: active` or `waiting`, including their `## 🔒 Hidden Truth`; `arcs/*.md` with `status: active` (steps and `## 🔒 Endgame`); `../../references/lenses/index.md` if no source class is obvious.
If there is no world root: say so in one line and suggest `/world-init`. If the world has no hooks, no links and an empty log, say that in one line and suggest `/create-lore` or `/world-init` instead of inventing content.

## Steps
1. **Scene.** From the GM's words; if they are thin, from the last `E<n>` of the log (where they are, who is present, what they just did). Ask nothing.
2. **Collect candidates**, in this order of preference:
   - **unspent ripples:** `R<k>` lines in the log that no later `E<n>` has made real. The GM already saw these, so they land as the world keeping its promises.
   - **someone's goal:** an NPC or group linked to the scene, acting on what their file already says they want. Never someone the log put elsewhere or killed.
   - **the place or the time:** what this place does on its own (the tide, the shift change, the fog, the chained gate), from the place's file.
   - **news arriving:** information travelling one hop along `## Links` (a messenger, a sound, an object someone drops).
3. **Pick 3, one per class** (`## Where the options come from` in the design: someone acts · the world answers · something arrives). Three variations of one idea is a failed answer; rewrite it. If one class has nothing, take a second from the richest class and say which classes you used.
4. **Check each option:** it obeys every `IR<n>`; it reveals no `## 🔒 Hidden Truth` and contradicts no arc `## 🔒 Endgame`; it does not resolve what the players are doing; it is refusable; it names no new NPC, place or group (an unnamed boy or guard is fine).
5. **Write each option** as one line of fiction that happens here and now, the source in brackets, and a `→ asks:` line with the choice it puts to the players. Mark with `🔒` any option the GM should read first because it touches a hidden truth.

## Output (table mode, in the output language)
```
🎲 What could happen now
1. <one line of fiction> — [[source]]<, R<k> if it is an unspent ripple>
   → asks: <the choice it puts to the players>
2. …
3. …
```
Three options, two lines each, nothing else. No preamble, no "would you like me to …".

## Prep mode
`prep` / `detailed`: 5 options, each with 1-2 likely ripples and the file they would change, and a `🧭` note when an option would make an active arc's current step *possible* (never mandatory). You may use the lens files for variety.

## Never
- Write anything: not a canon file, not the session log, not a `G<n>` proposal. An option the GM does not use leaves no trace; one they use arrives as an `E<n>`.
- Decide an outcome, call for a roll, or say what the players do.
- Make an arc's next step mandatory, or block what the players chose to do.
- Invent a named NPC, place or group (that is `/create-lore`), or bring in anything from a published setting.
- Reveal a hidden truth or an arc endgame, also not by hinting at it in an option the GM will read aloud.
