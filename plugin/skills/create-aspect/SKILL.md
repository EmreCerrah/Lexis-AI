---
name: create-aspect
description: Write one ready-to-use Fate aspect (scene, situation, consequence, boost or location aspect) for the current RPG scene in a Lexis-AI world, with an invoke and a compel example. Use when the GM asks for an aspect ("create-aspect", "write an aspect", "aspect for this scene", "bir aspect yaz").
---

# /create-aspect

Produces **one** aspect for the table. It is a proposal: logged, not canon.

## Read first
`../../AGENTS.md`; `../../systems/<system>/aspects.md`; `world.md` (tone, language and `## Immutable Rules`); the open session log (recent `E<n>` for scene context); the files of the current place and people and their linked hooks.

## Steps
1. **Context.** Use the GM's words; if they are thin, use the last `E<n>` of the open log. Ask nothing unless there is no scene at all.
2. **Kind.** Infer it from the request: scene/situation (default), consequence, boost, permanent (place/group).
3. **Fiction first (core):** the phrase, short and concrete, in the world's voice and language, true to what happened, obeying every `IR<n>`, revealing no `## 🔒 Hidden Truth`.
4. **Fate add-on (pack):** one invoke line, one compel line. If you cannot write both, rewrite the phrase.
5. **Log:** append to `## Generated` of the open log (create the log if there is none): `- G<n> aspect (E<k>|—): "<aspect>" · <kind>`.

## Output (in the output language)
```
✨ G<n> · **"<Aspect>"** — <kind>
- Invoke: …
- Compel: …
```
Exactly one aspect. If the GM asks for several or for alternatives, give up to 3 lines and log each as its own `G<n>`.

## Never
Write the aspect into a canon file. Permanent aspects reach world files only through `/close-session`.
