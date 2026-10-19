---
name: create-lore
description: Generate one persistent world entity (NPC, location, group/faction or story hook) for a Lexis-AI world as a proposal file in proposals/, using the world's templates and links. Not canon until approved at /close-session. Use when the GM asks to create an NPC, place, faction or hook ("create-lore", "bir NPC lazım", "bir mekan üret", "yeni grup", "hook yaz").
---

# /create-lore

Produces **one** persistent entity as a **proposal**. It becomes canon only if the GM approves it at `/close-session`.

## Read first
`../../AGENTS.md`; `../../references/world-layout.md`; the matching template `../../templates/<character|location|group|hook>.md`; `../../templates/proposal.md`; `../../systems/<sistem>/aspects.md` (for the `## Aspect'ler` section); `world.md` (tone, `## Değişmez Kurallar`); the entities the GM mentions and their links; existing slugs (to avoid duplicates).

## Steps
1. **Type.** karakter / mekan / grup / hook, from the GM's words. If the GM names an existing entity, stop and say it exists (`[[slug]]`); suggest `/event` for changes.
2. **Fiction (core).** Fill the template body in the world's voice. Short: every section 1-3 sentences. Obey all `K<n>`. Never contradict an existing `## 🔒 Gizli Gerçek`; a new hook gets its own hidden truth — consistent with all others.
3. **Aspects (pack).** 1-3 aspects per `aspects.md` (character: high concept + trouble).
4. **Links.** At least 1, at most 4 `## Bağlantılar` lines to existing slugs — the relation that makes this entity matter at the table (who it serves, where it lives, which hook it feeds). These are what "who might know" will run on.
5. **Uncertain parts** → its own `## Belirsiz` (e.g. "Gerçek adı bilinmiyor").
6. **Write** `proposals/<slug>.md` = proposal frontmatter (`templates/proposal.md`) + filled body. Append `- U<n> lore: [[<slug>]] (<tür>) · proposals/<slug>.md` to `## Üretilenler` of the open log (create the log if none).

## Output (Turkish, table mode)
```
📜 U<n> · <Ad> (<tür>) → proposals/<slug>.md
- Kim/ne: …
- İstek / amaç: …
- Aspect: "…" · "…"
- Bağlantılar: [[…]] — …
- 🔒 …
```
In prep mode, show the full file instead.

## Never
Write into `karakterler/ mekanlar/ gruplar/ hooklar/` — that happens only at `/close-session`.
