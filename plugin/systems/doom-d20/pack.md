---
type: system-pack
id: doom-d20
name: Doom d20 Core (OSR, four abilities, Omens)
order: 8
version: 0.1.0
---

# System Pack: Doom d20 Core

A lethal, rules-light d20 engine on top of the world core: four abilities from −3 to +3, tests against a Difficulty Rating, short fights, Omens. The core produces fiction; this pack adds the numbers. Use this pack when the GM asks for this doom-laden OSR-style d20 game.

## License
Lexis-AI Doom d20 Core is an independent production by the Lexis-AI contributors and is not affiliated with Ockult Örtmästare Games or Stockholm Kartell. It is published under the MÖRK BORG Third Party License. MÖRK BORG is copyright Ockult Örtmästare Games and Stockholm Kartell.

Compatible with MÖRK BORG. The rules are restated in our own words; no text, art, logos or layout from the original books are used.

## Content scope
Mechanics only, restated in our own words: abilities, tests and Difficulty Ratings, combat, armor, hit points, Omens, morale, reactions, rest. No setting content: none of the original world, its prophecies, places, gods, factions, named characters, creatures, classes, scrolls or items. No text, art or layout from the original books.

## Glossary
| Role | Term in this system |
|---|---|
| check | test: d20 + ability against a Difficulty Rating (DR) |
| trait | ability: Strength, Agility, Presence, Toughness (−3 to +3) |
| difficulty | Difficulty Rating (DR 6-18; 12 is normal) |
| scene-tag | a DR shift (±2 to ±4) from the scene, or a hazard |
| special-ability | class ability, power |
| harm | hit points; broken at 0, dead below 0 |
| meta-currency | Omens |

Keep these terms in English in every output language (`AGENTS.md` §3); add a short gloss the first time if the output language differs.

## Files
| File | Content |
|---|---|
| `resolution.md` | tests, DRs, abilities, crits and fumbles, combat basics, Omens, the `/help-skills` answer shape |
| `scene-tags.md` | DR shifts and hazards as scene tags; the output block |
| `obstacles.md` | choosing a DR, traps and hazards, morale, reactions, the obstacle output block |
| `entities.md` | the system sections of character, location and group files; character math; harm and rest |
| `init-questions.md` | the setup questions for this system |
| `templates/system.md` | the world's `system.md` |

## Seam (core → pack)
1. The skill writes the **fiction block** first, without numbers.
2. Then it appends a **`d20` block** from this pack.
3. Homebrew in the world's `system.md` overrides anything here.

## Never
- Roll dice, decide outcomes, or spend Omens on the GM's or players' behalf.
- Quote or paraphrase passages of the original books; restate rules briefly in our own words only.
- Use anything from the original game's world (its end-times prophecy, places, gods, people, creatures, classes, scrolls). The GM's world is the only setting.
