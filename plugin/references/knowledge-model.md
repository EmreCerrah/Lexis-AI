# Knowledge Model

The unit of knowledge is **file + section**. There is no per-line tagging. Which category a statement belongs to follows from where it is written.

## Categories
### Canon
- What is true in the world. Lives in the **body** of a canon file: every section except `## Uncertain`. `## Links` is canon too.
- Sources: `/world-init` answers, GM-reported events, items approved at `/close-session`.
- A GM-reported event is canon instantly, but only inside the open session log (`## Events`). World files catch up at `/close-session`.

### Hypothesis
- What might be true or might follow: rumors, unconfirmed consequences, open questions, deferred world-init answers.
- Lives in `## Uncertain` of the relevant file, or in `## Ripples` of the open session log until `/close-session`.
- Line format in `## Uncertain`: `- <claim> _(source: S003-R2)_` or `_(source: world-init, deferred)_`.
- Lifecycle: at `/close-session` each hypothesis is **promoted** (→ canon body), **kept** (→ `## Uncertain`) or **rejected** (deleted; the log keeps the record).
- Later events can make an old `## Uncertain` line true or false. `/close-session` points these out and proposes promotion or removal.

### Inference
- What you derive right now: who might know, which check or approach fits, which lens matters, an arc's current step (the first step whose hook is not `resolved`).
- **Never persisted.** Re-derive it every time from the current files, so it never goes stale.

### Proposal
- An item you generated with a `create-*` command.
- Persistent entities (`/create-lore`) → `proposals/<slug>.md`. Scene items (`/create-aspect`, `/create-obstacle`, `/create-scene`) → a `G<n>` line in `## Generated` of the session log.
- At `/close-session` the GM says which were used at the table. Used lore → canon file (its uncertain parts → `## Uncertain`); unused → deleted or kept in `proposals/` for later, as the GM chooses.

## Links (`## Links`)
- Every canon entity file has a fixed `## Links` section, one link per line:
  `- [[slug]] — <relation>`, for example `- [[order-of-light]] — member`, `- [[old-cemetery]] — keeper`.
- Links are directional in the file but read as **bidirectional** for knowledge inference: if A links B, B is reachable from A and A from B.
- New links come from `/world-init`, `/create-lore`, and link proposals (`L<n>`) raised by `/event` and `/ripples`, approved at `/close-session`.

## Who might know (inference algorithm)
Input: an event `E<n>` with its place (`Place`) and participants (`Who`).
1. **Hop 0, present:** the participants and the place, plus entities whose `## Links` tie them to the place with a presence relation (keeper, resident, owner, worker, regular). Label: *there / saw it*.
2. **Hop 1, linked:** every entity linked (in either direction) to a hop-0 entity. When a hop-0 character belongs to a group, the group and its other members are reached **within hop 1** (group membership does not cost an extra hop). Label: *may hear*.
3. **Hop 2+ (only `/ripples`):** continue from hop-1 entities and label each step with its weakening: hop 2 *may hear it as a rumor*, hop 3 *distorted rumor*. Stop at hop 3.
4. Remove anyone the fiction rules out (dead, absent, cut off) and anyone whose knowing would expose a hook's `## 🔒 Hidden Truth` too early; for them note `🔒 knowing would expose the hidden truth`.
5. Output at most 5 names in table mode, strongest first: `- [[slug]] — <label> (<path, e.g. member → [[order-of-light]]>)`.

The GM decides who actually knows. Never write this list to a file.

## Lifecycle summary
```
GM reports event ──► session log ## Events (canon, instant)
        │
        ├─► ripples ──► session log ## Ripples (hypothesis)
        ├─► link ideas ──► session log ## Link Proposals (proposal)
        └─► rule/hook conflicts ──► session log ## Open Decisions

create-* ──► proposals/ or session log ## Generated (proposal)

/close-session ──► GM approves in one batch ──► canon files / ## Uncertain / delete
```
