---
name: create-lore
description: Generate one persistent world entity (NPC, location, group/faction or story hook) for a Lexis-AI world as a proposal file in proposals/, using the world's templates and links. Not canon until approved at /close-session. Use when the GM asks to create an NPC, place, faction or hook ("create-lore", "I need an NPC", "make a location", "new faction", "write a hook", "bir NPC lazım").
---

# /create-lore

Produces **one** persistent entity as a **proposal**. It becomes canon only if the GM approves it at `/close-session`.

## Read first
`../../AGENTS.md`; `../../references/world-layout.md`; the matching template `../../templates/<character|location|group|hook>.md`; `../../templates/proposal.md`; `../../systems/<system>/aspects.md` (for the `## Aspects` section); `world.md` (tone, language, `## Immutable Rules`); the entities the GM mentions and their links; existing slugs (to avoid duplicates).

## Steps
1. **Type.** character / location / group / hook, from the GM's words. If the GM names an existing entity, stop, say it exists (`[[slug]]`) and suggest `/event` for changes.
2. **Fiction (core).** Fill the template body in the world's voice and language. Keep it short: 1-3 sentences per section. Obey every `IR<n>`. Never contradict an existing `## 🔒 Hidden Truth`; a new hook gets its own hidden truth, consistent with all the others.
3. **Aspects (pack).** 1-3 aspects per `aspects.md` (for a character: high concept + trouble).
4. **Links.** At least 1 and at most 4 `## Links` lines to existing slugs: the relations that make this entity matter at the table (whom it serves, where it lives, which hook it feeds). "Who might know" runs on these.
5. **Uncertain parts** → its own `## Uncertain` (e.g. "Real name unknown").
6. **Write** `proposals/<slug>.md` = the proposal frontmatter (`templates/proposal.md`) + the filled body. Append `- G<n> lore: [[<slug>]] (<type>) · proposals/<slug>.md` to `## Generated` of the open log (create the log if there is none).

## Output (table mode, in the output language)
```
📜 G<n> · <Name> (<type>) → proposals/<slug>.md
- Who/what: …
- Wants / goal: …
- Aspects: "…" · "…"
- Links: [[…]] — …
- 🔒 …
```
In prep mode, show the full file instead.

## Never
Write into `characters/ locations/ groups/ hooks/`; that happens only at `/close-session`.
