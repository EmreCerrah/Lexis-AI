# Design: System Packs

Status: contract implemented with packs `srd5`, `orc-d20`, `fate-accelerated`, `fate-core`, `d100` · Roadmap item 1 · Branch `feat/system-packs`

## Goal
The GM chooses a game system in `/world-init`, and every command gives rules help in that system's own terms. The world core (events, ripples, links, canon, hooks) does not change between systems. The MVP has one pack, Fate Accelerated (FAE); this design adds a second and a third pack without forking the skills.

## What is FAE-specific today
| Place | FAE assumption | After this change |
|---|---|---|
| `/help-skills` | answers in approaches + the 4 actions + the ladder | answers in the pack's resolution terms (`resolution.md`) |
| `/create-aspect` | produces a Fate aspect with invoke/compel | produces the pack's **scene tag** (`scene-tags.md`): an aspect in Fate, an environmental feature or condition elsewhere |
| `/create-obstacle` | FAE difficulty ladder and approaches | the pack's difficulty and opposition (`obstacles.md`) |
| `/create-lore` | `## Aspects` section | the pack's entity sections (`entities.md`) |
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
| `entities.md` | the system sections of character (PC/NPC), location and group files, character math, a light NPC stat line | `/world-init`, `/create-lore`, `/close-session` |
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
- Trademarks as names. A pack is named by its mechanics or its open license (e.g. `srd5` "5E SRD", `d100` "d100 Percentile Core"), never by a brand. A brand may appear only where its owner's license or policy allows it (e.g. SRD 5.2.1 allows "compatible with fifth edition"), never as a claim of endorsement.
- Rule text from a system that is not open-licensed. For such systems, the pack would describe the mechanic in our own words and point to the GM's own book.

**Wording:** even for open-licensed sources, we write the pack in our own short words and do not paste whole SRD sections. The pack is a GM cheat sheet, not a rulebook copy.

**The same rule applies to worlds.** The agent never brings published setting content (names, places, gods, NPCs) into a world on its own. A world contains only what the GM created or explicitly wrote in. This goes into `AGENTS.md` as a core rule.

### What stays fixed in every pack
- The seam: **fiction first (core), then the pack's add-on block.**
- The agent never rolls, never decides outcomes, never spends the meta-currency.
- `system.md` Homebrew overrides the pack.
- Output in the world's language; game terms stay as the pack names them (with a gloss on first use if the language differs).

## Changes to the skills
- Skills read `systems/<system>/<role-file>` instead of `approaches.md` / `aspects.md`. FAE's files are renamed to the contract: `approaches.md` → `resolution.md`, `aspects.md` → `scene-tags.md`, plus a new `entities.md` (from the stress/stunt parts of `approaches.md` and the `## Aspects` sections of the templates).
- `/help-skills` keeps its name; its answer shape comes from `resolution.md`.
- `/create-aspect` keeps its name for now (it is the locked MVP name); in non-Fate packs it produces the pack's scene tag. **Open decision:** add a system-neutral alias such as `/create-tag`.
- `templates/character.md`, `location.md` and `group.md` lose `## Aspects` and get a `{{system sections}}` line; the pack's `entities.md` says which sections replace it (FAE: `## Aspects`, PC `## Approaches`, `## Stunts`, `## Stress`; 5E: NPC `## Stat Line`, PC `## Abilities`, `## Proficiencies`, `## Combat` …; location `## Features`).
- Old worlds keep working: a world whose `system.md` says `fate-accelerated` gets the same output as today.

## Pack order and licenses
The GM's order. Licenses were checked against the publishers' own pages; re-check before each pack ships.

| # | System | Pack id | Source and license | What the pack may contain | Status |
|---|---|---|---|---|---|
| 1 | Fate Accelerated | `fate-accelerated` | FAE SRD, **CC-BY 3.0** | mechanics in our words + attribution | done (refactored to the contract) |
| 2 | 5E (fifth edition) | `srd5` | SRD 5.2.1, **CC-BY-4.0** | mechanics in our words + attribution; "compatible with fifth edition" only; no setting, spells, items or monsters | done |
| 3 | d20 remaster rules (2nd edition) | `orc-d20` | Player Core / GM Core, **ORC License** | mechanics; no trademarks or setting (the ORC excludes "Reserved Material"); the pack folder itself is released under the ORC License (it adapts Licensed Material) and carries the ORC, attribution and Reserved Material notices | done |
| 4 | World of Darkness–style d10 pools | — | **not open**: Paradox's Dark Pack is a non-commercial, revocable fan license | — | **skipped** (GM's decision) |
| 5 | Fate Core | `fate-core` | **CC-BY 3.0** | mechanics + attribution + the Fate™ trademark line | done |
| 6 | Percentile d100 engine | `d100` (not `brp`: "Basic Roleplaying" is a trademark) | Basic Roleplaying: Universal Game Engine ORC Content Document, **ORC License** (Chaosium) | mechanics in our words; no setting, powers lists or creatures; the pack folder is released under the ORC License with the ORC, attribution and Reserved Material notices | done |
| 7 | Forged in the Dark (d6 pools, clocks) | `fitd` | Blades in the Dark SRD, **CC-BY 3.0** | mechanics + attribution; no setting | planned |
| 8 | Mörk Borg–compatible (d20, doom) | tbd | **MÖRK BORG Third Party License** | mechanics restated in our own words only; no book text or art; "compatible" logo rules apply | planned |
| 9 | Mutants & Masterminds 3e–style (d20 effects) | tbd | **OGL 1.0a** Open Game Content (Product Identity excluded, e.g. "hero points") | mechanics; no Product Identity, no setting | planned |

## Install and selection
- `lexis-setup` asks "Which game system do you play?" (packs listed by `name`, menu order from the `order` field in `pack.md`, plus "All systems"; flag `--system`). It installs **only the chosen packs**; the other packs' files are never written. A re-run with another choice replaces the installed set. New packs need no installer change.
- `/world-init` uses the installed pack directly when there is one, and lists them when there are several.
- The world's `system` frontmatter field stores the pack id (`fate-accelerated`, `srd5` …).

## Testing
- Contract check (a small Go test in `lexis-setup`, it already reads the plugin): every `systems/*/` folder has all contract files, and `pack.md` has a glossary row for every role.
- Content-rule check (same Go test):
  - A pack folder may hold **only** the contract files; any extra file fails the test.
  - `pack.md` must have `## License` (source, license, attribution) and `## Content scope` (the statement "mechanics only, no setting content").
  - A denylist of setting and brand terms (setting names, gods, famous NPCs and creatures of each source system) is kept in the test; a pack file that contains one fails the test.
- Human review: every new pack is read once against the content rule before it is merged.
- Same scenario, three systems: run `/help-skills "climb the crypt wall in the rain"` and `/create-obstacle` against one world per pack; the fiction part should match, the mechanics block should be in each system's terms.
- Regression: the Karcsa Vadisi sample world (FAE) gives the same kind of output as before.

## Decisions
1. Systems and order: the table above (GM's choice).
2. `/create-aspect` keeps its name; it makes the pack's scene tag in every system. A neutral alias can come later.
3. One world, one system. A world without a `system:` field uses `fate-accelerated`.

4. World of Darkness–style: skipped, its rules are not open-licensed.
