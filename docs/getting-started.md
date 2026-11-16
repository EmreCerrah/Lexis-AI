# Getting Started

This walk-through takes you from an empty folder to a closed first session. Commands are written as `/event`, `/ripples` … In your tool, call them like this:

| Tool | How to call a command | Example |
|---|---|---|
| Claude Code, user install | `/lexis-ai:<command>` | `/lexis-ai:event The players broke into the crypt` |
| Claude Code, project install | `/<command>` | `/event The players broke into the crypt` |
| Codex | `$<command>` (or pick `@lexis-ai`) | `$event The players broke into the crypt` |
| GitHub Copilot | `/<command>` | `/event The players broke into the crypt` |
| Gemini CLI | describe what you want | `event: the players broke into the crypt` |

Plain requests work everywhere too: "record what just happened", "close the session".

## 1. Make a world folder
Create an empty folder for the campaign and open your AI tool in it. Everything Lexis-AI writes goes into this folder.

```
my-campaign/
```

## 2. Set up the world: `/world-init`
Lexis-AI asks the **core questions** in one message. Answer them in one go; short answers are fine.

1. World name
2. Concept (1-3 sentences)
3. Genre
4. Language the AI answers in for this world
5. Narrative style / tone
6. Inspirations
7. Fantasy level
8. **Immutable Rules**: the facts of your fiction that must never break, e.g. "The dead never rise on their own", "Sunlight burns vampires". At least one.
9. System: if only one game system is installed, it is used automatically.

You can answer any item with **"later"** (it is noted under `## Uncertain` in `world.md`) or **"you suggest"** (the AI drafts it and you accept or change it).

Then it offers the **optional parts**: world notes, characters, locations, groups, religion and creation, story hooks. Again, "later" or "you suggest" is fine for each. Finally it asks the **system questions** of your game system (for example starting level and ability score method in 5E).

Result:
```
✅ World created: Karcsa Valley
- Files: world.md, system.md, 3 characters, 2 locations, 1 group, 2 hooks
- Immutable Rules: IR1 … IR3
- Left uncertain: religion, opening scene
- Next: report what happens at the table with /event; at the end of the session run /close-session.
```

## 3. During the session

### Something happened: `/event`
Report it in one or two sentences, the way you'd tell a friend.
```
/event Gabriel caught the gravekeeper next to an empty coffin. The keeper ran and dropped a theater ticket with a yellow sign on it.
```
```
📌 E1 recorded (session-001) · Lens: society-rumor, faith-supernatural
Ripples (hypothesis):
- R1 The keeper hides in the tavern and mutters "it wasn't a vampire"
- R2 By morning the town says "the hunter attacked the keeper"
- R3 🔒 The troupe notices the lost ticket and sends a marked corpse after the keeper
Who might know:
- [[gravekeeper]] — there
- [[order-of-light]] — may hear (member → [[gabriel]])
```
- The event is **canon** the moment you report it.
- **Ripples** are only possibilities. You decide at the end of the session which ones really happened.
- Lines with 🔒 are for you only; don't read them out.

### A rules question: `/help-skills`
```
/help-skills Gabriel wants to scare the keeper into talking
```
The answer comes in your system's terms. In 5E, for example:
```
🎲 Ability check
- Check: Charisma (Intimidation) — looming over him · Charisma (Persuasion) — promising protection
- Suggested DC: 15 (Medium) — he is terrified, but of something else
- Advantage/Disadvantage: Advantage if Gabriel shows the ticket
- Feature: —
```

### Scene material: `/create-scene`, `/create-obstacle`, `/create-aspect`, `/create-lore`
```
/create-scene old-cemetery night, fog, the gravekeeper is gone
/create-obstacle the cellar door the keeper locked behind him
/create-aspect the graveyard at midnight
/create-lore an old priest who knows about the troupe
```
Each one gives you something ready to use. `/create-scene` keeps your notes and fills in the rest: what to read aloud, who is there, three things to do with the check and difficulty for each. Scenes, obstacles and scene tags are logged; new people and places go to `proposals/` and become part of the world only if you approve them at the end.

### The scene stalled: `/suggest-event`
```
/suggest-event they have searched the whole crypt and won't leave
```
Three things that could happen now, each from the world you already wrote and each naming where it comes from. Pick one, say it at the table, then report it with `/event`. Nothing is recorded.

### More consequences: `/ripples`
When there is a pause, dig deeper into an event:
```
/ripples E1
```
You get more ripples through different lenses (power and politics, society and rumor, faith and the supernatural) and how the news could travel further.

## 4. End of session: `/close-session`
```
/close-session
```
You get one numbered list:
```
🗂 Closing session 001 — 7 decisions
B. Events
 1. E1 → [[old-cemetery]] · Description: + "a coffin was found empty" · suggested ✅
C. Ripples
 2. R1 The keeper hides in the tavern … → [[gravekeeper]] · suggested ❓
 3. R2 The town blames the hunter … → [[karcsa-village]] · suggested ❓
 …
Answer like: "all as suggested" · "3 ✅, 5 ❌, rest as suggested" · "defer 4-6"
```
Answer once, for example `all as suggested, 2 ✅`. Only then are the world files changed.

- ✅ becomes canon in the right file.
- ❓ goes to that file's `## Uncertain` section.
- ❌ is dropped.
- Deferred items stay open for the next `/close-session`.

## 5. Between sessions
- Read and edit your world files directly; they are plain Markdown. See [Your world files](world-files.md).
- Add `prep` or `detailed` to a command for a longer analysis: `/ripples E1 prep`.
- Run `/world-init` again in an existing world to fill in what you left as "later".

## Tips for the table
- Keep `/event` short and factual. Name people and places as they are named in your files.
- Put relations in the `## Links` section of each file (`- [[order-of-light]] — member`). "Who might know" only follows these links.
- Write the **🔒 Hidden Truth** of each hook. Lexis-AI will never reveal it and never let a suggestion contradict it.
- If you report something that breaks an Immutable Rule, Lexis-AI records it anyway and flags `⚠ rule conflict`. You choose: it only looks like a violation, a costly exception, or the rule changes.
