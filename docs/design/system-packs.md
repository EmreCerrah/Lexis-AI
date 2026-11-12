# Design: System Packs

Status: draft for review · Roadmap item 1 · Branch `feat/system-packs`

## Goal
The GM chooses a game system in `/world-init`, and every command gives rules help in that system's own terms. The world core (events, ripples, links, canon, hooks) does not change between systems. The MVP has one pack, Fate Accelerated (FAE); this design adds a second and a third pack without forking the skills.

## What is FAE-specific today
| Place | FAE assumption | After this change |
|---|---|---|
| `/help-skills` | answers in approaches + the 4 actions + the ladder | answers in the pack's resolution terms (`resolution.md`) |
| `/create-aspect` | produces a Fate aspect with invoke/compel | produces the pack's **scene tag** (`scene-tags.md`): an aspect in Fate, an environmental feature or condition elsewhere |
| `/create-obstacle` | FAE difficulty ladder and approaches | the pack's difficulty and opposition (`obstacles.md`) |
| `/create-lore` | `## Aspects` section | the pack's NPC/entity sections (`characters.md`) |
| `/world-init` | FAE is the only choice; FAE setup questions | lists the installed packs; asks that pack's `init-questions.md` |
| `templates/character.md` | `## Aspects`, PC `## Approaches` / `## Stunts` | core keeps `## Who`, `## Wants` …; the pack adds its own sections |
| `AGENTS.md` §3, §10 | approach names as game terms | "keep the pack's game terms" from the pack glossary |

Everything else (`/event`, `/ripples`, `/close-session`, the knowledge model, lenses, hooks, Immutable Rules) is already system-neutral.

## The pack contract
Every pack is a folder `systems/<id>/` with the same files. Skills ask for a file by role and never name a system:

| File | Role | Used by |
|---|---|---|
| `pack.md` | identity: name, id, version; `## License` (source, license, attribution); `## Content scope`; the **glossary** (see below); what the pack never does | every skill that touches mechanics |
| `resolution.md` | how an action resolves: dice, the difficulty scale, outcome bands; the answer shape for a rules question | `/help-skills`, `/create-obstacle` |
| `scene-tags.md` | what a "scene item" is in this system and its output block | `/create-aspect` |
| `obstacles.md` | difficulty choice, passive vs active opposition, hazards, cost/complication options, output block | `/create-obstacle` |
| `characters.md` | sections a PC and an NPC get in their file, and a light NPC stat line | `/world-init`, `/create-lore` |
| `init-questions.md` | the system part of the world setup | `/world-init` |
| `templates/system.md` | the world's `system.md` (settings + `## Homebrew`) | `/world-init` |

### Glossary: generic roles → system terms
Skills speak in roles; the glossary in `pack.md` translates them. The agent uses the system's term in every answer.

| Role | FAE | D&D 5e (SRD) | Basic Roleplaying / Call of Cthulhu-style |
|---|---|---|---|
| `check` (the roll) | approach roll (4dF + approach) | ability check (d20 + modifier) | skill roll (d100 under skill %) |
| `trait` (what a check uses) | approach | ability + skill proficiency | skill % (characteristics for raw rolls) |
| `difficulty` | ladder, passive opposition | DC 5-30 | regular / hard (½) / extreme (⅕) |
| `scene-tag` | situation aspect | environmental feature, condition, advantage/disadvantage source | situation modifier (bonus/penalty die) |
| `special-ability` | stunt | feat, class feature | (none in the core; Homebrew) |
| `harm` | stress + consequences | hit points + conditions | hit points (+ sanity in horror games) |
| `meta-currency` | fate points | heroic inspiration | luck (optional) |

A pack may leave a role empty ("—"); skills then skip that line in their output.

### Pack content rule: mechanics only (legal scope)
A pack is **the rules engine and nothing else**. This is a hard rule for every pack, open-licensed or not.

**Allowed in a pack:**
- Resolution: dice, how a roll is read, the difficulty scale, outcome bands, advantage/bonus mechanics.
- Character math: how traits are generated and calculated (ability scores → modifiers, proficiency bonus, skill %, half/fifth values, HP and derived stats, advancement formulas).
- Generic conditions and harm (prone, frightened, wounded, stress …), generic GM procedures (setting a difficulty, opposition, cost of failure).
- Generic, setting-free examples written by us ("a locked iron door", "a guard dog").
- The license attribution the source requires.

**Never in a pack:**
- Setting lore: worlds, planes, nations, cities, gods, religions, factions, history.
- Named characters, named monsters or creatures tied to a setting, adventures, story text, plot hooks.
- Spell, item or monster lists copied wholesale; art; trade dress.
- Trademarks as names. A pack is named by its mechanics or its open license (e.g. `srd5` "5E SRD", `brp` "Basic Roleplaying"), never by a brand. A brand may appear only in a "compatible with" line, and only where the license allows it.
- Rule text from a system that is not open-licensed. For such systems, the pack would describe the mechanic in our own words and point to the GM's own book.

**Wording:** even for open-licensed sources, we write the pack in our own short words and do not paste whole SRD sections. The pack is a GM cheat sheet, not a rulebook copy.

**The same rule applies to worlds.** The agent never brings published setting content (names, places, gods, NPCs) into a world on its own. A world contains only what the GM created or explicitly wrote in. This goes into `AGENTS.md` as a core rule.

### What stays fixed in every pack
- The seam: **fiction first (core), then the pack's add-on block.**
- The agent never rolls, never decides outcomes, never spends the meta-currency.
- `system.md` Homebrew overrides the pack.
- Output in the world's language; game terms stay as the pack names them (with a gloss on first use if the language differs).

## Changes to the skills
- Skills read `systems/<system>/<role-file>` instead of `approaches.md` / `aspects.md`. FAE's files are renamed to the contract: `approaches.md` → `resolution.md`, `aspects.md` → `scene-tags.md`, plus a new `characters.md` (from the stress/stunt parts of `approaches.md` and the character template).
- `/help-skills` keeps its name; its answer shape comes from `resolution.md`.
- `/create-aspect` keeps its name for now (it is the locked MVP name); in non-Fate packs it produces the pack's scene tag. **Open decision:** add a system-neutral alias such as `/create-tag`.
- `templates/character.md` loses `## Aspects`; the pack's `characters.md` says which sections to add (FAE: `## Aspects`, PC `## Approaches`, `## Stunts`).
- Old worlds keep working: a world whose `system.md` says `fate-accelerated` gets the same output as today.

## Candidate packs
| Pack | Why | License | Effort |
|---|---|---|---|
| **D&D 5e (SRD 5.2)** | the most played system; forces the contract to cover d20, DCs, conditions, HP | SRD 5.2 is CC-BY-4.0: may quote rules text with the required attribution | medium |
| **Basic Roleplaying (BRP)**, Call of Cthulhu-style percentile horror | fits gothic/cosmic horror worlds like the sample; covers d100 and sanity-style harm | BRP core is released under the ORC License: rules may be used with attribution | medium |
| **Fate Core** | closest to FAE: skills instead of approaches; cheap check of the contract | CC-BY | low |
| World of Darkness / Chronicles of Darkness | asked for earlier; dice pools, attributes + abilities | **not open-licensed**: the pack could only explain how to apply the GM's own book, without rule text | medium, legally limited |

Recommendation: **D&D 5e SRD as the second pack** (largest audience, most different from FAE), then **BRP as the third** (horror fit, d100, open license). Fate Core can come any time as a small extra.

## Install and selection
- `lexis-setup` already copies every folder under `systems/`, so new packs ship with no installer change.
- `/world-init` lists the packs it finds in `systems/` (name from each `pack.md`) instead of offering only FAE.
- The world's `system` frontmatter field stores the pack id (`fate-accelerated`, `srd5`, `brp`).

## Testing
- Contract check (a small Go test in `lexis-setup`, it already reads the plugin): every `systems/*/` folder has all contract files, and `pack.md` has a glossary row for every role.
- Content-rule check (same Go test):
  - A pack folder may hold **only** the contract files; any extra file fails the test.
  - `pack.md` must have `## License` (source, license, attribution) and `## Content scope` (the statement "mechanics only, no setting content").
  - A denylist of setting and brand terms (setting names, gods, famous NPCs and creatures of each source system) is kept in the test; a pack file that contains one fails the test.
- Human review: every new pack is read once against the content rule before it is merged.
- Same scenario, three systems: run `/help-skills "climb the crypt wall in the rain"` and `/create-obstacle` against one world per pack; the fiction part should match, the mechanics block should be in each system's terms.
- Regression: the Karcsa Vadisi sample world (FAE) gives the same kind of output as before.

## Open decisions (for the GM)
1. Second and third system: D&D 5e SRD + BRP as recommended, or another choice?
2. `/create-aspect`: keep only that name, or add a neutral alias such as `/create-tag`?
3. One world, one system (current model), or should a world be able to switch systems later (then `/close-session` would have to convert sheets)? Recommendation: one system per world for now.
