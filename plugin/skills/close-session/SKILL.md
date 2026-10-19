---
name: close-session
description: End an RPG session in a Lexis-AI world - gather everything from the open session log (events, ripples, link proposals, generated items, rule conflicts) into one numbered batch for the GM to approve, then apply only the approved changes to the canon world files. Use when the GM says the session is over ("close-session", "seansı kapat", "seans bitti", "canon'u güncelle").
---

# /close-session

The **only** place where canon world files change after `/world-init`. One batch of decisions, one GM answer, then writes. Nothing is written before the GM answers.

## Read first
`../../AGENTS.md` (§5 write permissions, §6 conflicts); `../../references/knowledge-model.md`; `../../references/world-layout.md`; the open log (`durum: açık`) and any log with `durum: bekliyor`; every file named in them; `proposals/`.
If no log is open or waiting: say "Kapatılacak açık seans yok." and stop.

## 1. Build the batch (no writes yet)
Number items continuously across groups. Give each a **suggested** answer (`önerilen`). Keep each item to one line plus, for file edits, one preview line `→ dosya · bölüm: + "metin"`.

**A. Açık kararlar** (decide first — they change other items)
- Each `A<m>` without a GM answer: rule conflict → (a) görünüşte ihlal (b) bedelli istisna (c) kural değişikliği; hook break → (a) plan geçerli (b) plan değişir. For (c)/(b) show the proposed new rule / hidden-truth text, or ask the GM to give it.

**B. Olaylar → dünya dosyaları**
- Events are already canon. For each `O<n>`, propose the concrete edits that write its *facts* (not ripples) into entity files: state changes, deaths, items moved, places damaged, hook `durum` changes. Önerilen: ✅.

**C. Yankılar** (`Y<k>`)
- For each: ✅ canon (it happened/was confirmed at the table) · ❓ belirsiz (→ target file `## Belirsiz`) · ❌ ret. Önerilen: ❓, unless a later `O<n>` confirms it (then ✅) or contradicts it (then ❌).

**D. Bağlantı önerileri** (`B<j>`) → target file `## Bağlantılar`. ✅ / ❌.

**E. Üretilenler** (`U<n>`) — ask "masada kullanıldı mı?"
- aspect/engel: used & lasting (a place/group aspect) → propose adding to that file's `## Aspect'ler`; used & transient → only marked used in the log; unused → marked unused.
- lore: ✅ canon (move `proposals/<slug>.md` → its folder, drop proposal frontmatter, `tür:` = hedef-tür) · ⏸ beklet (stays in proposals/) · ❌ sil.

**F. Eski belirsizler**
- Existing `## Belirsiz` lines (any file) that this session's events confirm or refute → ✅ promote to body / ❌ remove / ❓ keep.

**G. Bilgi** (no decision): dangling `[[links]]` to files that don't exist; unresolved names (`yeni: …`).

## 2. Ask once (Turkish)
```
🗂 Seans <NNN> kapanışı — <n> karar
A. Açık kararlar
 1. A1 (O2 · K2) … — (a)/(b)/(c) · önerilen: a
B. Olaylar
 2. O1 → [[eski-mezarlik]] · Tasvir: + "…" · önerilen ✅
C. Yankılar
 3. Y1 … → [[…]] · önerilen ❓
 …
Cevap örnekleri: "hepsi önerilen" · "3 ✅, 5 ❌, gerisi önerilen" · "1 c: <yeni kural>" · "ertele 4-6"
```
Then **stop and wait**. Do not write while waiting. If the GM changes the content of an item, use their words.

## 3. Apply (only after the answer)
- Write exactly the approved edits. Keep every heading; append inside the right section; don't rewrite unrelated text.
- `## Belirsiz` lines carry their source: `- … _(kaynak: S<NNN>-Y<k>)_`.
- Rule change (c): edit the `K<n>` text in `world.md`, keep its number, append ` _(değişti: S<NNN>)_`. Never delete a rule number silently.
- Hidden truth (b): replace the `## 🔒 Gizli Gerçek` text with the GM's new text only.
- Deferred items (`ertele`): leave them in the log; set the log `durum: bekliyor`. Otherwise `durum: kapalı`.
- Fill `## Seans Sonu Kararları` in the log: one line per item, `<no>. <id> → ✅/❓/❌/⏸/ertelendi`.

## 4. Report (Turkish)
```
✅ Seans <NNN> kapandı (<kapalı|bekliyor>)
- Güncellenen dosyalar: [[…]], [[…]]
- Belirsize eklenen: <n> · Reddedilen: <n> · Bekleyen öneri: <n>
- Kural değişikliği: <K2 → …> | —
```

## Never
- Write any canon file before the GM's answer, or write an item the GM did not approve.
- Promote a ripple to canon just because it is likely.
- Rewrite a hidden truth or a rule without option (b)/(c) being chosen explicitly.
