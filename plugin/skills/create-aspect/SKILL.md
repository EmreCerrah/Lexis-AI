---
name: create-aspect
description: Write one ready-to-use Fate aspect (scene, situation, consequence, boost or location aspect) for the current RPG scene in a Lexis-AI world, with an invoke and a compel example. Use when the GM asks for an aspect ("create-aspect", "bir aspect yaz", "sahneye aspect", "bu duruma aspect öner").
---

# /create-aspect

Produces **one** aspect for the table. It is a proposal: logged, not canon.

## Read first
`../../AGENTS.md`; `../../systems/<sistem>/aspects.md`; `world.md` (tone + `## Değişmez Kurallar`); the open session log (recent `O<n>` for scene context); files of the current place/people and linked hooks.

## Steps
1. **Context.** Use the GM's words; if thin, use the last `O<n>` of the open log. Do not ask questions unless there is no scene at all.
2. **Kind.** Infer from the request: sahne/durum (default), sonuç, boost, kalıcı (place/group).
3. **Fiction first (core):** the phrase — short, concrete, in the world's voice and language, true to what happened, obeying `K<n>`, not revealing any `## 🔒 Gizli Gerçek`.
4. **Fate add-on (pack):** one invoke line, one compel line. If you cannot write both, rewrite the phrase.
5. **Log:** append to `## Üretilenler` of the open log (create the log if none): `- U<n> aspect (O<k>|—): "<aspect>" · <tür>`.

## Output (Turkish)
```
✨ U<n> · **"<Aspect>"** — <tür>
- Invoke: …
- Compel: …
```
Exactly one aspect. If the GM writes `3 tane`/`alternatif`, give up to 3 lines but log each as its own `U<n>`.

## Never
Write the aspect into a canon file. Persistent aspects reach world files only through `/close-session`.
