# World Folder Layout

A world is a plain folder of Markdown files the GM owns. The plugin never stores world data inside itself.

```
<world>/
  AGENTS.md              # short pointer written by /world-init (platform agents auto-load it)
  world.md               # canon: concept, tone, ## Immutable Rules …
  system.md              # system pack settings and Homebrew
  characters/<slug>.md
  locations/<slug>.md
  groups/<slug>.md
  hooks/<slug>.md
  arcs/<slug>.md         # chains of hooks (optional)
  sessions/session-001.md   # session logs, one per session
  proposals/<slug>.md       # /create-lore output waiting for approval
```

## Files
- Templates: `templates/` (world core) and `systems/<system>/templates/` (system pack). Copy the template, fill it, and keep every heading even when empty (write `—` under an empty heading).
- Headings and frontmatter keys stay in English exactly as in the templates; the content under them is in the output language (`AGENTS.md` §3).
- Frontmatter `type:` is one of `world`, `system`, `character`, `location`, `group`, `hook`, `arc`, `session`, `proposal`.
- **Slugs:** lowercase ASCII, words joined by `-`, accented letters transliterated (ç→c, ğ→g, ı→i, ö→o, ş→s, ü→u, é→e, ß→ss …). Slugs are unique across the whole world, because `[[slug]]` must resolve by file name (Obsidian-compatible).
- **Links:** `[[slug]]` only: no folder prefix, no `.md`. A link to a file that does not exist yet is allowed; `/close-session` lists such dangling links.

## Session logs
- `sessions/session-NNN.md`, three digits, numbered from 001. Template: `templates/session-log.md`.
- Frontmatter `status:` is `open`, `pending` (closed but some decisions deferred) or `closed`.
- **The open session log** = the one with `status: open`. There is at most one. If a skill that appends to the log finds none, it creates the next number with `status: open` and `mode: table` (or `mode: prep` if the GM is preparing between sessions).
- Item IDs restart every session: events `E1…`, ripples `R1…`, link proposals `L1…`, generated items `G1…`, open decisions `D1…`. `IR1`, `IR2` … are the Immutable Rules in `world.md`. Outside the log, refer to an item as `S003-E2`.
- Logs are append-only during the session: new lines go at the end of the right section (`## Events`, `## Ripples` …). Never rewrite an earlier line; to correct one, append `- correction (E2): …`. Exception: `/close-session` sets `status:`, fills the `GM:` answers in `## Open Decisions` and writes `## End-of-Session Decisions`.
- A `pending` log still holds deferred items; the next `/close-session` includes them. New events go into a new `open` log, never into a `pending` one.

## Canon files
Canon files are `world.md`, `system.md`, `characters/`, `locations/`, `groups/`, `hooks/`, `arcs/`. Only `/world-init` (creating the world) and `/close-session` (approved items) write to them. See `AGENTS.md` §6.
