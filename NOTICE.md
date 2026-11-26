# Notices

Lexis-AI is two things under two different licenses.

## 1. The project: MIT
Everything except `plugin/systems/` is MIT-licensed (see [LICENSE](LICENSE)): the installer (`lexis-setup/`), the skills (`plugin/skills/`), the core rules (`plugin/AGENTS.md`), the references, the templates and the documentation. Use it, fork it, ship it; keep the copyright notice.

## 2. The system packs: each follows its source
A pack in `plugin/systems/<id>/` restates the mechanics of an openly licensed game **in our own words**. Each pack's full license text and notices live in its own `pack.md`, which is installed alongside it — that file, not this summary, is the operative notice.

| Pack | Based on | License |
|---|---|---|
| `srd5` | System Reference Document 5.2.1 by Wizards of the Coast LLC, https://www.dndbeyond.com/srd | CC-BY-4.0 |
| `orc-d20` | the second-edition remaster core rules Paizo released under the ORC License (Player Core, GM Core) | ORC License |
| `fate-accelerated` | Fate Accelerated Edition, a product of Evil Hat Productions, LLC | CC-BY 3.0 |
| `fate-core` | Fate Core System and Fate Accelerated Edition, products of Evil Hat Productions, LLC | CC-BY 3.0 |
| `d100` | Basic Roleplaying: Universal Game Engine, released by Chaosium under the ORC License | ORC License |
| `fitd` | Blades in the Dark, a product of One Seven Design, developed and authored by John Harper | CC-BY 3.0 |
| `doom-d20` | MÖRK BORG, copyright Ockult Örtmästare Games and Stockholm Kartell | MÖRK BORG Third Party License |
| `hero-d20` | the effects-based superhero d20 rules published as Open Game Content, third edition | OGL 1.0a |

Per-pack detail, including what each pack deliberately leaves out: [docs/game-systems.md](docs/game-systems.md).

## What the packs contain
- **Mechanics only.** No setting, no places, gods, factions, named characters, creatures or story from any published game. A world's content comes from its GM.
- **No copied rule text.** Rules are restated; where a license requires exact notices (ORC, OGL Section 15, MÖRK BORG), those notices are reproduced verbatim in the pack's `pack.md`.
- **Game names appear only to say which rules a pack follows.** All trademarks belong to their owners. Lexis-AI is not affiliated with, endorsed or sponsored by any of them. Where a license forbids naming the source to claim compatibility (OGL Product Identity, as for `hero-d20`), the name is not printed.

## Trademark note
D&D and Dungeons & Dragons are trademarks of Wizards of the Coast LLC. Pathfinder is a trademark of Paizo Inc. Fate is a trademark of Evil Hat Productions, LLC. MÖRK BORG is copyright Ockult Örtmästare Games and Stockholm Kartell. Basic Roleplaying is a trademark of Chaosium Inc. These names are used only nominatively, to identify the rules a pack restates.

## Your world files are yours
`lexis-setup` installs the plugin; it never touches a world folder's contents, and nothing in this project claims any right in the worlds you write with it.
