---
name: close-session
description: End an RPG session in a Lexis-AI world - gather everything from the open session log (events, ripples, link proposals, generated items, rule conflicts) into one numbered batch for the GM to approve, then apply only the approved changes to the canon world files. Use when the GM says the session is over ("close-session", "end the session", "session's over", "update canon", "seansı kapat").
---

# /close-session

The **only** place where canon world files change after `/world-init`. One batch of decisions, one answer from the GM, then the writes. Nothing is written before the GM answers.

## Read first
`../../AGENTS.md` (§6 write permissions, §7 conflicts); `../../references/knowledge-model.md`; `../../references/world-layout.md`; the open log (`status: open`) and any log with `status: pending`; every file named in them; `proposals/`.
If no log is open or pending, say there is no open session to close, and stop.

## 1. Build the batch (no writes yet)
Number the items continuously across groups. Give each one a **suggested** answer. Keep each item to one line, plus a preview line for file edits: `→ file · section: + "text"`.

**A. Open decisions** (decide these first; they change other items)
- Each `D<m>` without a GM answer. Rule conflict → (a) apparent violation (b) costly exception (c) rule change. Hook break → (a) plan holds (b) plan changes. For (c)/(b), show the proposed new rule or hidden-truth text, or ask the GM to give it.

**B. Events → world files**
- Events are already canon. For each `E<n>`, propose the concrete edits that write its *facts* (not its ripples) into entity files: state changes, deaths, items moved, places damaged, hook `status` changes. Suggested: ✅.

**C. Ripples** (`R<k>`)
- For each: ✅ canon (it happened or was confirmed at the table) · ❓ uncertain (→ the target file's `## Uncertain`) · ❌ reject. Suggested: ❓, unless a later `E<n>` confirms it (then ✅) or contradicts it (then ❌).

**D. Link proposals** (`L<j>`) → the target file's `## Links`. ✅ / ❌.

**E. Generated items** (`G<n>`): ask "was it used at the table?"
- scene tag/obstacle: used and lasting (a place or group tag) → propose adding it to that file's system section (`entities.md` of the pack: `## Aspects` in Fate, `## Features` in 5E); used and transient → only marked as used in the log; unused → marked unused.
- lore: ✅ canon (move `proposals/<slug>.md` into its folder, drop the proposal frontmatter, `type:` = target-type) · ⏸ keep (stays in proposals/) · ❌ delete.

**F. Old uncertain items**
- Existing `## Uncertain` lines (in any file) that this session's events confirm or refute → ✅ promote to the body / ❌ remove / ❓ keep.

**G. Info** (no decision): dangling `[[links]]` to files that don't exist; unresolved names (`new: …`).

## 2. Ask once (in the output language)
```
🗂 Closing session <NNN> — <n> decisions
A. Open decisions
 1. D1 (E2 · IR2) … — (a)/(b)/(c) · suggested: a
B. Events
 2. E1 → [[old-cemetery]] · Description: + "…" · suggested ✅
C. Ripples
 3. R1 … → [[…]] · suggested ❓
 …
Answer like: "all as suggested" · "3 ✅, 5 ❌, rest as suggested" · "1 c: <new rule>" · "defer 4-6"
```
Then **stop and wait**. Do not write anything while waiting. If the GM changes an item's content, use their words.

## 3. Apply (only after the answer)
- Write exactly the approved edits. Keep every heading; append inside the right section; don't rewrite unrelated text.
- `## Uncertain` lines carry their source: `- … _(source: S<NNN>-R<k>)_`.
- Rule change (c): edit the `IR<n>` text in `world.md`, keep its number, append ` _(changed: S<NNN>)_`. Never silently delete a rule number.
- Hidden truth (b): replace the `## 🔒 Hidden Truth` text with the GM's new text only.
- Deferred items: leave them in the log and set the log to `status: pending`. Otherwise `status: closed`.
- Fill `## End-of-Session Decisions` in the log with one line per item: `<no>. <id> → ✅/❓/❌/⏸/deferred`.

## 4. Report (in the output language)
```
✅ Session <NNN> closed (<closed|pending>)
- Updated files: [[…]], [[…]]
- Added to uncertain: <n> · Rejected: <n> · Proposals kept: <n>
- Rule change: <IR2 → …> | —
```

## Never
- Write any canon file before the GM answers, or write an item the GM did not approve.
- Promote a ripple to canon just because it is likely.
- Rewrite a hidden truth or a rule unless option (b)/(c) was chosen explicitly.
