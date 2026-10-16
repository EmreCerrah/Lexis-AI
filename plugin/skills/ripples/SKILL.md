---
name: ripples
description: Deepen an already recorded event in a Worldbuilding Forge world (folder with world.md) by running it through analysis lenses (politics, rumor, faith…) and tracing how knowledge of it spreads 2-3 hops with weakening; results are hypotheses appended to the open session log. Use when the GM asks for more consequences of an event ("ripples", "yankılar", "O2'nin yankılarını derinleştir", "bu olayın sonuçları ne olur").
---

# /ripples

Takes an event that is already in a session log and asks "what else follows?" through lenses. Output is **hypothesis** only.

## Read first
1. `../../AGENTS.md`. 2. `../../references/knowledge-model.md`. 3. `../../references/lenses/index.md`, then the picked lens files.
4. World: `world.md` `## Değişmez Kurallar`, the event's files and their links up to 3 hops, linked hooks.

## Steps
1. **Find the event.** Argument may be `O2`, `S003-O2`, or a description. Default: the last `O<n>` of the open log. If the event is not in any log, tell the GM to record it with `/event` first (one line) and stop.
2. **Mode.** Table (default) or prep (the GM writes `hazırlık`/`detaylı`, or no session log is open).
3. **Pick lenses.** Table: 2-3. Prep: every lens that scores. Skip lenses the event's existing `Y<k>` lines already covered unless the GM asks again.
4. **Per lens**, answer the lens's `Sorular` against the world files. Keep only ripples that:
   - follow from a link, goal or rule in the files (or are clearly marked as a new guess),
   - obey every `K<n>` and every `## 🔒 Gizli Gerçek`,
   - do not repeat an existing `Y<k>`.
5. **Knowledge spread.** Run the algorithm to hop 2-3 using the lens's `Yayılma kanalları`. Label weakening (`söylenti olarak`, `çarpıtılmış`). Output only.
6. **Append** new ripples to `## Yankılar` of the open log (create one with `mod: hazırlık` if none is open): `- Y<k> (O<n> · <lens>) … · hedef: [[dosya]]`. New relations → `## Bağlantı Önerileri` as `B<j>`.

## Output (Turkish)
Table mode:
```
🌊 O<n> yankıları · Lens: <a>, <b>
<lens a>:
- Y4 …
- Y5 …
<lens b>:
- Y6 …
Yayılma:
- [[…]] — söylenti olarak duyabilir (… → … → …)
```
Max 5 ripples total in table mode. Prep mode: up to 4 per lens, each with one sentence "neden" and the file it would change; plus a short "Hangi hook'u yaklaştırıyor?" line (🔒).

## Never
- Edit canon files, or present a ripple as something that happened.
- Let a ripple explain or reveal a hidden truth; it may only move toward it.
