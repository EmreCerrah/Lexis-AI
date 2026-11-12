---
type: system-pack
id: orc-d20
name: ORC 2E Core (three-action d20, remaster rules)
order: 2
version: 0.1.0
---

# System Pack: ORC 2E Core

The second-edition remaster core rules that Paizo released under the ORC License (Player Core, GM Core), on top of the world core. The core produces fiction; this pack adds checks with four degrees of success, proficiency ranks, level-based DCs, valued conditions and character math. Use this pack when the GM asks for these second-edition remaster rules or the three-action d20 system.

## License
**ORC Notice.** This product is licensed under the ORC License located at the Library of Congress at TX 9-307-067 and available online at various locations including www.paizo.com/orclicense, www.azoralaw.com/orclicense, www.gencon.com/orclicense and others. All warranties are disclaimed as set forth therein.

**Attribution.** This product is based on the following Licensed Material: Pathfinder Player Core © 2023 Paizo Inc., Designed by Logan Bonner, Jason Bulmahn, Stephen Radney-MacFarland, and Mark Seifter. Pathfinder GM Core © 2023 Paizo Inc. If you use our Licensed Material in your own published works, please credit us in your product as follows: Lexis-AI ORC 2E Core system pack © 2026 the Lexis-AI contributors.

**Reserved Material.** Reserved Material elements in this product include all elements designated as Reserved Material under the ORC License. This pack contains no trademarks, setting content or art of the licensor.

**Expressly Designated Licensed Material.** This product contains no Expressly Designated Licensed Material.

**This pack only** (the files in this folder) is released under the ORC License. The rest of Lexis-AI keeps its own license.

## Content scope
Mechanics only: checks, degrees of success, proficiency, DCs, bonuses and penalties, actions, conditions, hit points, dying and character math. No setting content: no world, gods, nations, places, organizations, named characters, named creatures or story. No ancestry, class, feat, spell, item or creature entries; for those the GM uses their own book. Rules are summarized in our own words.

## Glossary
| Role | Term in this system |
|---|---|
| check | check: d20 + modifier against a DC, with four degrees of success |
| trait | attribute modifier + skill proficiency rank (untrained, trained, expert, master, legendary) |
| difficulty | DC: simple DCs by rank, or the level-based DC with an adjustment |
| scene-tag | circumstance: a circumstance bonus or penalty, cover, concealment, difficult terrain, a hazard or a condition |
| special-ability | feat, class feature, spell |
| harm | Hit Points, dying and wounded, valued conditions |
| meta-currency | Hero Points |

Keep these terms in English in every output language (`AGENTS.md` §3); add a short gloss the first time if the output language differs.

## Files
| File | Content |
|---|---|
| `resolution.md` | checks, degrees of success, proficiency, DCs, bonus types, actions, social checks, Hero Points, the `/help-skills` answer shape |
| `scene-tags.md` | circumstances: cover, concealment, terrain, conditions; the circumstance output block |
| `obstacles.md` | choosing a DC, hazards, social obstacles, victory points, the obstacle output block |
| `entities.md` | the system sections of character, location and group files; character math; dying |
| `init-questions.md` | the setup questions for this system |
| `templates/system.md` | the world's `system.md` |

## Seam (core → pack)
1. The skill writes the **fiction block** first, without numbers.
2. Then it appends a **`2E` block** from this pack.
3. Homebrew in the world's `system.md` overrides anything here.

## Never
- Roll dice, decide success or failure, or tell the GM what the players must roll. Say "suggested DC" and "fitting checks".
- Award or spend Hero Points on the GM's behalf. You may remind: "this could earn a Hero Point".
- Quote ancestry, class, feat, spell, item or creature entries. If a rule depends on one, name it and tell the GM to check their book.
- Use any content from the publisher's setting. The GM's world is the only setting.
