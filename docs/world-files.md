# Your World Files

A Lexis-AI world is a plain folder of Markdown files. You own it: the plugin stores no world data, and you can read and edit every file yourself.

## Layout
```
my-campaign/
  AGENTS.md            # a short pointer for AI tools (written by /world-init)
  world.md             # concept, tone, language, system, Immutable Rules (IR1, IR2 …)
  system.md            # your game system's settings and ## Homebrew
  characters/<slug>.md
  locations/<slug>.md
  groups/<slug>.md
  hooks/<slug>.md
  sessions/session-001.md   # one log per session
  proposals/<slug>.md       # /create-lore output waiting for approval
```

- **Slugs** are lowercase file names with dashes: `old-cemetery`, `order-of-light`. Accented letters are simplified (ç→c, ş→s, ü→u …). A slug is unique in the whole world.
- **Links** are written `[[slug]]`, without folder or `.md`. This is Obsidian-compatible; open the folder in Obsidian for a graph view.
- **Headings and frontmatter keys stay in English** (`## Links`, `## Uncertain`, `type: character`), so every command can find them. The text under them is in your world's language.

## The files
### world.md
Frontmatter: `type: world`, `name`, `language`, `system` (the pack id, e.g. `srd5`), `created`. Sections: Concept, Genre and Tone, Inspirations, **Immutable Rules**, World Notes, Religion and Creation, Intro Scene, Opening Scene, Links, Uncertain.

**Immutable Rules** are numbered `IR1`, `IR2` … and never renumbered. Lexis-AI never suggests anything that breaks them. If you report an event that does, it is recorded and flagged, and you decide at `/close-session` whether the rule changes (then it is marked `_(changed: S003)_`).

### system.md
Your game system's settings from `/world-init` (starting level, optional rules …) and **`## Homebrew`**: your house rules. Homebrew always wins over the system pack.

### Characters, locations, groups
Each file has fiction sections (`## Who`, `## Wants` / `## Description` …), the **system sections** of your game (aspects in Fate, a stat line in 5E, features of a place …), `## 🔒 GM Notes`, `## Links` and `## Uncertain`.

### Hooks
A story hook has:
- `## Surface`: what the players can see.
- `## 🔒 Hidden Truth`: what is really going on. Lexis-AI never reveals it and never lets a suggestion contradict it. If players try to break it, it suggests in-world reasons why they can't.
- `## Links`: the people and places involved, and the next hook.

### Session logs
`sessions/session-NNN.md`, with `status: open`, `pending` or `closed`. During a session only the log changes:

| Section | Items |
|---|---|
| `## Events` | `E1`, `E2` … what you reported (canon) |
| `## Ripples` | `R1` … possible consequences (hypotheses) |
| `## Link Proposals` | `L1` … new relations |
| `## Generated` | `G1` … scene tags, obstacles, lore proposals |
| `## Open Decisions` | `D1` … rule conflicts and broken hidden truths |
| `## End-of-Session Decisions` | your answers at `/close-session` |

Logs are append-only: earlier lines are never rewritten. Outside the log, an item is referred to as `S003-E2`.

## Canon, hypothesis, inference, proposal
| Kind | Where it lives | How it changes |
|---|---|---|
| **Canon** | the body of a world file | only through `/close-session` (or `/world-init`) |
| **Hypothesis** | `## Uncertain`, or `## Ripples` in the log | promoted, kept or dropped at `/close-session` |
| **Inference** | only in an answer ("who might know") | never stored |
| **Proposal** | `proposals/`, or `## Generated` in the log | approved or rejected at `/close-session` |

## Who might know
"Who might know" follows only the **`## Links`** section of each file, one relation per line:
```
## Links
- [[order-of-light]] — member
- [[old-cemetery]] — keeper
```
`/event` goes one step (people present, plus those directly linked; group members count). `/ripples` goes 2-3 steps and marks how the news weakens. Relations written only in prose are not followed, so keep `## Links` up to date. It is a *possibility* of knowing; you decide what NPCs actually know.

## Editing by hand
You can edit any file at any time. Keep the headings, keep `[[slug]]` links, and don't renumber Immutable Rules. If you edit during a session, the change is canon immediately (Lexis-AI treats your files as the truth).

## Saving, moving and sharing a world
Everything is in the folder, so:
- **Back up** by copying the folder, or put it under git (`git init` in the world folder) to see every change `/close-session` made.
- **Move** it anywhere and open your AI tool there. If you used a *project install*, copy the tool's folder (`.claude/`, `.agents/`, `.gemini/` or `.github/`) together with the world, or install Lexis-AI for your user instead.
- **Close the session first** (`/close-session`) so nothing is left only in an open log. An open or pending log still works after moving; the next `/close-session` picks it up.
- **Several worlds:** each world is its own folder with its own `world.md`. Lexis-AI uses the nearest `world.md` above the folder you start the tool in.

## Older worlds
Worlds made by early versions may use Turkish headings (`Değişmez Kurallar`, `Bağlantılar`, `Belirsiz` …) and IDs (`K1`, `O1` …). They keep working; Lexis-AI reads both. A world without a `system:` field uses Fate Accelerated.
