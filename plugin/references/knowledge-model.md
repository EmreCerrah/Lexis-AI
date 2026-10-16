# Knowledge Model

The unit of knowledge is **file + section**. There is no per-line tagging. Which category a statement belongs to follows from where it is written.

## Categories
### Canon
- What is true in the world. Lives in the **body** of a canon file (every section except `## Belirsiz`). `## Bağlantılar` is canon too.
- Sources: `/world-init` answers, GM-reported events, items approved at `/close-session`.
- A GM-reported event is canon instantly, but only inside the open session log (`## Olaylar`). World files catch up at `/close-session`.

### Hypothesis
- What might be true or might follow. Rumors, unconfirmed consequences, open questions, deferred world-init answers.
- Lives in `## Belirsiz` of the relevant file, or in `## Yankılar` of the open session log until `/close-session`.
- Line format in `## Belirsiz`: `- <claim> _(kaynak: S003-Y2)_` or `_(kaynak: world-init, ertelendi)_`.
- Lifecycle: at `/close-session` each hypothesis is **promoted** (→ canon body), **kept** (→ `## Belirsiz`), or **rejected** (deleted; the log keeps the record).
- Later events can make an old `## Belirsiz` line true or false. `/close-session` points these out and proposes promotion/removal.

### Inference
- What you derive right now: who might know, which approach fits, which lens matters.
- **Never persisted.** Re-derive every time from the current files, so it never goes stale.

### Proposal
- An item you generated with a `create-*` command.
- Persistent entities (`/create-lore`) → `proposals/<slug>.md`. Scene items (`/create-aspect`, `/create-obstacle`) → a `U<n>` line in `## Üretilenler` of the session log.
- At `/close-session`: the GM says which were used at the table. Used lore → canon file (or its uncertain parts → `## Belirsiz`); unused → deleted or kept in `proposals/` for later, the GM's choice.

## Links (`## Bağlantılar`)
- Every canon entity file has a fixed `## Bağlantılar` section. One link per line:
  `- [[slug]] — <relation>` e.g. `- [[isik-tarikati]] — üye`, `- [[eski-mezarlik]] — bekçi`.
- Links are directional in the file but read as **bidirectional** for knowledge inference: if A links B, B is reachable from A and A from B.
- New links come from `/world-init`, `/create-lore`, and link proposals (`B<n>`) raised by `/event` and `/ripples`, approved at `/close-session`.

## Who might know (inference algorithm)
Input: an event `O<n>` with its place (`Yer`) and participants (`Kimler`).
1. **Hop 0 — present:** participants and the place. Also entities whose `## Bağlantılar` link the place with a presence relation (bekçi, sakin, sahibi, çalışan, müdavim). Label: *orada / gördü*.
2. **Hop 1 — linked:** every entity linked (either direction) to a hop-0 entity. When a hop-0 character belongs to a group, the group and its other members are reached **within hop 1** (group membership does not cost an extra hop). Label: *duyabilir*.
3. **Hop 2+ (only `/ripples`):** continue from hop-1 entities. Label each step with weakening: hop 2 *söylenti olarak duyabilir*, hop 3 *çarpıtılmış söylenti*. Stop at hop 3.
4. Remove anyone the event's fiction rules out (dead, absent, cut off), and anyone whose knowing would expose a hook's `## 🔒 Gizli Gerçek` too early — instead note `🔒 bilirse gizli gerçek açığa çıkar`.
5. Output at most 5 names in table mode, strongest first: `- [[slug]] — <label> (<path, e.g. üye → [[isik-tarikati]]>)`.

The GM decides who actually knows. Never write this list to a file.

## Lifecycle summary
```
GM reports event ──► session log ## Olaylar (canon, instant)
        │
        ├─► ripples ──► session log ## Yankılar (hypothesis)
        ├─► link ideas ──► session log ## Bağlantı Önerileri (proposal)
        └─► rule/hook conflicts ──► session log ## Açık Kararlar

create-* ──► proposals/ or session log ## Üretilenler (proposal)

/close-session ──► GM approves in one batch ──► canon files / ## Belirsiz / delete
```
