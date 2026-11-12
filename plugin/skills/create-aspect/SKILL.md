---
name: create-aspect
description: Write one ready-to-use scene tag for the current RPG scene in a Lexis-AI world, in the world's game system - a Fate aspect with an invoke and a compel, or a 5E scene feature (light, cover, terrain, hazard, condition) with its rule effect. Use when the GM asks for an aspect, a scene tag or a scene feature ("create-aspect", "write an aspect", "aspect for this scene", "a feature for this room", "bir aspect yaz").
---

# /create-aspect

Produces **one** scene tag for the table: what the world's system calls it (glossary role `scene-tag` in `pack.md`: an aspect in Fate, a scene feature in 5E). It is a proposal: logged, not canon.

## Read first
`../../AGENTS.md`; `world.md` (`system:`, tone, language and `## Immutable Rules`); `../../systems/<system>/pack.md` and `scene-tags.md`; `system.md` (Homebrew); the open session log (recent `E<n>` for scene context); the files of the current place and people and their linked hooks.

## Steps
1. **Context.** Use the GM's words; if they are thin, use the last `E<n>` of the open log. Ask nothing unless there is no scene at all.
2. **Kind.** Infer it from the request, using the kinds in `scene-tags.md`. Default: the scene/situation kind.
3. **Fiction first (core):** the name or phrase, short and concrete, in the world's voice and language, true to what happened, obeying every `IR<n>`, revealing no `## 🔒 Hidden Truth`.
4. **System add-on (pack):** fill the output block of `scene-tags.md`. It must be two-sided: one way it helps the PCs, one way it hinders them. If you cannot write both, rewrite it.
5. **Log:** append to `## Generated` of the open log (create the log if there is none): `- G<n> <scene-tag term> (E<k>|—): "<name>" · <kind>` (e.g. `aspect`, `feature`).

## Output (in the output language)
```
✨ G<n> · <the output block of scene-tags.md>
```
Exactly one. If the GM asks for several or for alternatives, give up to 3 and log each as its own `G<n>`.

## Never
Write it into a canon file. Permanent ones (a place or group tag) reach world files only through `/close-session`.
