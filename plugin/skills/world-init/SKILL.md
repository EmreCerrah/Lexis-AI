---
name: world-init
description: Create a new Lexis-AI RPG world folder (world.md, system.md, entity files, hooks) through a short tiered interview with the GM. Use when the GM wants to start, set up or initialize a new campaign world ("yeni dünya kur", "world-init", "dünyayı başlat"), or to fill deferred sections of an existing world between sessions.
---

# /world-init

Sets up a world in tiers so the GM is never forced to answer everything up front. Prep-mode command: approvals inside this conversation are fine.

## Read first
1. `../../AGENTS.md` (core rules), `../../references/world-layout.md`, `../../references/knowledge-model.md`.
2. Templates: `../../templates/*.md`. System pack: `../../systems/<sistem>/pack.md` and `init-questions.md`.

## 0. Where
- Target folder = the current working directory unless the GM names another. It must not be inside the plugin folder.
- If the target already has `world.md`: do **not** overwrite. Switch to **fill mode**: list the `## Belirsiz` items marked `world-init, ertelendi` and empty sections, ask which ones to fill now, and write only those (canon sections the GM answers + delete the matching `## Belirsiz` line). Then stop.

## 1. Core (required) — one message, all questions together
Ask in Turkish, numbered, with a one-line example for each. Accept short answers; accept "sen öner" for any item and draft it (mark the draft and get a yes/no in the same conversation).
1. Dünyanın adı
2. Konsept (1-3 cümle)
3. Tür
4. Dil (varsayılan: Türkçe)
5. Anlatı tarzı / ton
6. İlham kaynakları
7. Fantazi seviyesi
8. Değişmez Kurallar — kurgunun asla bozulmayacak kuralları (ör. "Ölüler kendiliğinden dirilmez", "Güneş vampiri yakar"). En az 1.
9. Sistem (MVP: Fate Accelerated)

Do not continue until 1, 2, 8 and 9 have answers. Others may stay empty → `## Belirsiz`.

## 2. Deferrable — one message
Offer the list and say clearly: "Hepsi isteğe bağlı. 'sonra' dersen ## Belirsiz'e not düşerim; 'sen öner' dersen taslak çıkarırım."
- Dünya notları · Karakterler · Mekanlar · Gruplar · Din/yaratılış · Hook'lar
For each answer:
- GM gave content → write as canon.
- "sonra" / no answer → add `- <bölüm> henüz belirlenmedi _(kaynak: world-init, ertelendi)_` to `world.md` `## Belirsiz`.
- "sen öner" → generate (step 3).

## 3. Generatable — drafts
For items the GM asked you to draft (intro scene, opening scene, characters, locations, groups, hooks):
- Draft briefly, in the world's tone, obeying every `K<n>`. For hooks: write `## Yüzey`, `## 🔒 Gizli Gerçek`, `## Bağlantılar` — never leave the hidden truth for later; if the GM wants to write it themself, leave `{{GM yazacak}}` and add a `## Belirsiz` line.
- Show all drafts in one message; the GM accepts, edits or drops each. Only accepted drafts are written.
- Every entity gets `## Bağlantılar` links at creation (character↔group membership, character↔location presence, hook↔its people/places, hook→next hook). Links are what "who might know" runs on — do not skip them.

## 4. System pack
Ask the questions in `../../systems/<sistem>/init-questions.md` as one batch. Unanswered = standard. Fill `system.md` from `../../systems/<sistem>/templates/system.md`.

## 5. Write
Create (see `references/world-layout.md`):
- `world.md` (from template; Değişmez Kurallar numbered `K1…`), `system.md`, `AGENTS.md` (from `templates/world-AGENTS.md`)
- `karakterler/`, `mekanlar/`, `gruplar/`, `hooklar/`, `seanslar/`, `proposals/` (create empty folders too)
- one file per accepted entity, slugs per the layout rules
Do **not** create a session log; `/event` opens one.

## 6. Report (Turkish, short)
```
✅ Dünya kuruldu: <ad>
- Dosyalar: world.md, system.md, <n> karakter, <n> mekan, <n> grup, <n> hook
- Değişmez Kurallar: K1 … Kn
- Belirsiz bırakılanlar: <liste veya —>
- Sırada: masada olanları `/event` ile bildir; seans sonunda `/close-session`.
```
