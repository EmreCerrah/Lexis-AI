# World Folder Layout

A world is a plain folder of Markdown files the GM owns. The plugin never stores world data inside itself.

```
<dünya>/
  AGENTS.md            # short pointer written by /world-init (platform agents auto-load it)
  world.md             # canon: concept, tone, ## Değişmez Kurallar …
  system.md            # system pack settings: stunts, special situations, Homebrew
  karakterler/<slug>.md
  mekanlar/<slug>.md
  gruplar/<slug>.md
  hooklar/<slug>.md
  seanslar/seans-001.md   # session logs, one per session
  proposals/<slug>.md     # /create-lore output waiting for approval
```

## Files
- Templates: `templates/` (world core) and `systems/<sistem>/templates/` (system pack). Copy the template, fill it, keep every heading even when empty (write `—` under an empty heading).
- Frontmatter `tür:` is one of `dünya`, `sistem`, `karakter`, `mekan`, `grup`, `hook`, `seans`, `öneri`.
- **Slugs:** lowercase ASCII, words joined by `-`, Turkish letters transliterated (ç→c, ğ→g, ı→i, ö→o, ş→s, ü→u). Slugs are unique across the whole world because `[[slug]]` must resolve by file name (Obsidian-compatible).
- **Links:** `[[slug]]` only, no folder prefix, no `.md`. A link to a file that does not exist yet is allowed; `/close-session` lists such dangling links.

## Session logs
- `seanslar/seans-NNN.md`, three digits, numbered from 001. Template: `templates/session-log.md`.
- Frontmatter `durum:` is `açık` (open), `bekliyor` (closed but some decisions deferred) or `kapalı`.
- **The open session log** = the one with `durum: açık`. There is at most one. If a skill that appends to the log finds none, it creates the next number with `durum: açık` and `mod: masa` (or `mod: hazırlık` if the GM is preparing between sessions).
- Item IDs restart every session: events `O1…`, ripples `Y1…`, link proposals `B1…`, generated items `U1…`, open decisions `A1…` (`K` is reserved for Değişmez Kurallar: `K1`, `K2` … in `world.md`). Outside the log, refer to an item as `S003-O2`.
- Logs are append-only during the session: new lines go at the end of the right section (`## Olaylar`, `## Yankılar` …). Never rewrite an earlier line; to correct, append `- düzeltme (O2): …`. Exception: `/close-session` sets `durum:`, fills `GM:` answers in `## Açık Kararlar` and writes `## Seans Sonu Kararları`.
- A `bekliyor` log still holds deferred items; the next `/close-session` includes them. New events go into a new `açık` log, never into a `bekliyor` one.

## Canon files
Canon files are `world.md`, `system.md`, `karakterler/`, `mekanlar/`, `gruplar/`, `hooklar/`. Only `/world-init` (creating the world) and `/close-session` (approved items) write to them. See `AGENTS.md` §5.
