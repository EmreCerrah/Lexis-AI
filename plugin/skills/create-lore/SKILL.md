---
name: create-lore
description: Generate one persistent world entity (NPC, location, group/faction, story hook or arc of chained hooks) for a Lexis-AI world as a proposal file in proposals/, using the world's templates and links. Not canon until approved at /close-session. Use when the GM asks to create an NPC, place, faction or hook ("create-lore", "I need an NPC", "make a location", "new faction", "write a hook", "make an arc from these hooks", "bir NPC lazım").
---

# /create-lore

Produces **one** persistent entity as a **proposal**. It becomes canon only if the GM approves it at `/close-session`.

## Read first
`../../AGENTS.md`; `../../references/world-layout.md`; the matching template `../../templates/<character|location|group|hook|arc>.md`; `../../templates/proposal.md`; `world.md` (`system:`, tone, language, `## Immutable Rules`); `../../systems/<system>/entities.md` (the system sections) and `scene-tags.md`; the entities the GM mentions and their links; existing slugs (to avoid duplicates).

## Steps
1. **Type.** character / location / group / hook / arc, from the GM's words. If the GM names an existing entity, stop, say it exists (`[[slug]]`) and suggest `/event` for changes.
2. **Fiction (core).** Fill the template body in the world's voice and language. Keep it short: 1-3 sentences per section. Obey every `IR<n>`. Never contradict an existing `## 🔒 Hidden Truth`; a new hook gets its own hidden truth, consistent with all the others.
3. **System sections (pack).** Replace the template's `{{system sections}}` line with the sections `entities.md` lists for this entity type (NPC, not PC, unless the GM says otherwise). Keep them light: what the table needs, nothing more.
4. **Links.** At least 1 and at most 4 `## Links` lines to existing slugs: the relations that make this entity matter at the table (whom it serves, where it lives, which hook it feeds). "Who might know" runs on these.
5. **Uncertain parts** → its own `## Uncertain` (e.g. "Real name unknown").
**Arc:** fill `templates/arc.md` instead of steps 2-5:
- **Premise:** 1-2 sentences.
- **Steps:** the hooks the GM names, in order, each with a `Moves on when:` that fits that hook's surface and hidden truth.
- **Endgame:** only if the GM asks for one, consistent with every hidden truth.
- A step whose hook does not exist yet becomes its own hook proposal, with its own hidden truth.
- A hook may belong to only one arc: if a named hook is already in an arc, say so.
- No system sections.

6. **Write** `proposals/<slug>.md` = the proposal frontmatter (`templates/proposal.md`) + the filled body. Append `- G<n> lore: [[<slug>]] (<type>) · proposals/<slug>.md` to `## Generated` of the open log (create the log if there is none).

## Output (table mode, in the output language)
```
📜 G<n> · <Name> (<type>) → proposals/<slug>.md
- Who/what: …
- Wants / goal: …
- System: <the key line: aspects in Fate, the stat line in 5E …>
- Links: [[…]] — …
- 🔒 …
```
In prep mode, show the full file instead.

## Never
Write into `characters/ locations/ groups/ hooks/ arcs/`; that happens only at `/close-session`.
