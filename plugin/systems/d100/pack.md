---
type: system-pack
id: d100
name: d100 Percentile Core (ORC)
order: 6
version: 0.1.0
---

# System Pack: d100 Percentile Core

The percentile universal game engine that Chaosium released under the ORC License, on top of the world core. The core produces fiction; this pack adds d100 skill rolls with five success levels, difficulty, the resistance table, wounds and optional Sanity. Use this pack when the GM asks for a percentile (d100, roll-under skill %) system, including classic percentile horror play.

## License
**ORC Notice.** This product is licensed under the ORC License located at the Library of Congress at TX 9-307-067 and available online at various locations including www.chaosium.com/orclicense, www.azoralaw.com/orclicense, www.gencon.com/orclicense and others. All warranties are disclaimed as set forth therein.

**Attribution.** This product is based on the following Licensed Material: Basic Roleplaying: Universal Game Engine ORC Content Document © 2023 Chaosium Inc., authors Jason Durall and Steve Perrin; based on the Basic Roleplaying system created by Steve Perrin, Steve Henderson, Warren James, Greg Stafford, Sandy Petersen, Ray Turney, and Lynn Willis. Basic Roleplaying is the trademark of Chaosium Inc. If you use our Licensed Material in your own published works, please credit us in your product as follows: Lexis-AI d100 Percentile Core system pack © 2026 the Lexis-AI contributors.

**Reserved Material.** Reserved Material elements in this product include all elements designated as Reserved Material under the ORC License. This pack contains no trademarks (other than the notice above), setting content or art of the licensor.

**Expressly Designated Licensed Material.** This product contains no Expressly Designated Licensed Material.

**This pack only** (the files in this folder) is released under the ORC License. The rest of Lexis-AI keeps its own license.

## Content scope
Mechanics only: characteristics, skill ratings, d100 rolls, success levels, difficulty and situational modifiers, the resistance table, opposed and cooperative rolls, wounds, healing, optional Sanity, experience. No setting content: no worlds, eras, places, named characters, named creatures, powers lists or story. Rules are summarized in our own words.

## Glossary
| Role | Term in this system |
|---|---|
| check | skill roll: d100 equal to or under the skill rating (characteristic roll or resistance roll when no skill fits) |
| trait | skill rating (%), or a characteristic roll (characteristic × 5) |
| difficulty | difficulty level (Easy ×2, Average, Difficult ½) + situational modifiers (±20%, ±50%) |
| scene-tag | situational modifier: a condition of the scene that adds or subtracts 20% or 50% |
| special-ability | power (magic, psychic ability, mutation, super power) |
| harm | hit points, major and fatal wounds; Sanity points (optional) |
| meta-currency | power points used as fate points (optional rule) |

Keep these terms in English in every output language (`AGENTS.md` §3); add a short gloss the first time if the output language differs.

## Files
| File | Content |
|---|---|
| `resolution.md` | rolls, the five success levels, difficulty, situational modifiers, the resistance table, opposed and cooperative rolls, the `/help-skills` answer shape |
| `scene-tags.md` | situational modifiers as scene tags; the output block |
| `obstacles.md` | choosing difficulty and modifiers, resistance obstacles, hazards, Sanity, social obstacles, the obstacle output block |
| `entities.md` | the system sections of character, location and group files; character math; wounds and healing |
| `init-questions.md` | the setup questions for this system (power level, optional rules) |
| `templates/system.md` | the world's `system.md` |

## Seam (core → pack)
1. The skill writes the **fiction block** first, without numbers.
2. Then it appends a **`d100` block** from this pack.
3. Homebrew in the world's `system.md` overrides anything here.

## Never
- Roll dice, decide success or failure, or tell the GM what the players must roll. Say "suggested difficulty" and "fitting skills".
- Spend power points or call for Sanity losses on the GM's behalf. You may suggest a loss, the GM decides.
- Describe powers, spells or creatures from the source book. If a rule depends on one, name it and tell the GM to check their book.
- Use content from any published setting. The GM's world is the only setting.
