---
name: event
description: Record something that just happened at the RPG table as canon in the open session log of a Worldbuilding Forge world (folder with world.md), then list 3-5 short ripples as hypotheses and who might know. Use when the GM reports an event during play ("event", "olay", "şu oldu", "oyuncular X yaptı"). Never edits canon world files.
---

# /event

The GM reports what happened. It is canon **instantly**. You record it, then show what might follow. Table mode: fast, short, zero approval questions.

## Read first
1. `../../AGENTS.md`. 2. `../../references/knowledge-model.md` (who-might-know algorithm). 3. `../../references/lenses/index.md`.
4. World: `world.md` (`## Değişmez Kurallar`), every file the event names, their `## Bağlantılar` targets (1 hop), and hooks linked to them (`## 🔒 Gizli Gerçek`).

## Steps
1. **Open the log.** Find `seanslar/` file with `durum: açık`; if none, create the next `seans-NNN.md` from `../../templates/session-log.md`.
2. **Resolve names.** Map people/places in the GM's words to slugs. Unknown name → keep the plain name in the log and add a link proposal or a note `yeni: <ad>` in the output; do not stop to ask.
3. **Record (append to `## Olaylar`)** exactly what the GM said, in neutral words. Do not add consequences here — the event block holds only what the GM reported:
   ```
   ### O<n> — <3-6 kelimelik başlık>
   - Ne oldu: <GM'in anlattığı>
   - Yer: [[…]] · Kimler: [[…]], …
   - Hook: [[…]] | —
   - Kural: uyumlu | ⚠ kural çatışmalı (K<n>) → A<m>
   ```
4. **Rule check.** Compare with every `K<n>`. If broken: still recorded; add `A<m>` to `## Açık Kararlar` with options (a) görünüşte ihlal (b) bedelli istisna (c) kural değişikliği — see AGENTS.md §6.
5. **Hidden-truth check.** If the event threatens a hook's `## 🔒 Gizli Gerçek`:
   - players *trying* to break it → give 1-3 in-world reasons it fails (output only).
   - GM says it really broke → add `A<m>` with options (a) plan geçerli, yeniden bağlanır (b) plan değişir.
6. **Pick lenses** per the index: max 2 (3 only if clearly needed). Open only those lens files.
7. **Ripples:** 3-5 short consequences, each one line, each obeying all `K<n>` and hidden truths. Append to `## Yankılar` as `- Y<k> (O<n> · <lens>) <yankı> · hedef: [[dosya]]` (hedef = the file it would change if approved). GM-only ripples start with `🔒`.
8. **Who might know:** 1 hop per the algorithm, max 5. Output only — never written.
9. **Link proposals:** if the event creates an obvious new relation (someone now owes, hates, saw, holds something), append `- B<j> (O<n>) [[a]] → [[b]] — <ilişki>` to `## Bağlantı Önerileri`. Max 2.

## Output (Turkish, table mode)
```
📌 O<n> kaydedildi (seans-NNN) · Lens: <a>, <b>
Yankılar (hipotez):
- Y1 …
- Y2 …
- Y3 …
Kim bilebilir:
- [[…]] — orada
- [[…]] — duyabilir (üye → [[…]])
<only if any:> ⚠ A1: K2 çatışması — (a) görünüşte (b) bedelli (c) kural değişir? Şimdi ya da seans sonunda.
<only if any:> 🔒 Neden başaramazlar: …
```
No preamble, no restating the event, no "onaylıyor musun?". If the GM writes `hazırlık`/`detaylı`, use up to 3 lenses and add one sentence of reasoning per ripple.

## Never
- Edit `world.md`, `system.md` or any file in `karakterler/ mekanlar/ gruplar/ hooklar/`.
- Write the who-might-know list to a file.
- Decide what NPCs actually know or do — ripples are possibilities.
