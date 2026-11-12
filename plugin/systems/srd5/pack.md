---
type: system-pack
id: srd5
name: 5E SRD (fifth edition compatible)
order: 1
version: 0.1.0
---

# System Pack: 5E SRD

Fifth-edition rules (the 2024 core rules as published in SRD 5.2.1) on top of the world core. The core produces fiction; this pack adds d20 tests, Difficulty Classes, conditions and character math. Use this pack when the GM asks for fifth edition or 5E rules. This pack is "compatible with fifth edition"; it makes no other claim.

## License
This work includes material from the System Reference Document 5.2.1 ("SRD 5.2.1") by Wizards of the Coast LLC, available at https://www.dndbeyond.com/srd. The SRD 5.2.1 is licensed under the Creative Commons Attribution 4.0 International License, available at https://creativecommons.org/licenses/by/4.0/legalcode.

## Content scope
Mechanics only: d20 tests, abilities and skills, Difficulty Classes, Advantage and Disadvantage, conditions, hit points, rests and character math. No setting content: no worlds, planes, gods, places, named characters, named creatures, adventures or story. No spell, item or monster lists; for those the GM uses their own book. Rules are summarized in our own words.

## Glossary
| Role | Term in this system |
|---|---|
| check | D20 Test: ability check, saving throw or attack roll (d20 + modifier) |
| trait | ability (Strength, Dexterity, Constitution, Intelligence, Wisdom, Charisma) + skill proficiency |
| difficulty | Difficulty Class (DC 5-30); Armor Class for attacks |
| scene-tag | scene feature: a source of Advantage or Disadvantage, cover, obscurement, Difficult Terrain, a hazard or a condition |
| special-ability | class feature, feat, spell |
| harm | Hit Points, conditions, Exhaustion |
| meta-currency | Heroic Inspiration |

Keep these terms in English in every output language (`AGENTS.md` §3); add a short gloss the first time if the output language differs.

## Files
| File | Content |
|---|---|
| `resolution.md` | D20 Tests, abilities and skills, DCs, Advantage, social checks, Heroic Inspiration, the `/help-skills` answer shape |
| `scene-tags.md` | scene features: light, cover, terrain, conditions; the scene-feature output block |
| `obstacles.md` | choosing a DC, traps, hazards, fear, social obstacles, the obstacle output block |
| `entities.md` | the system sections of character, location and group files; character math; hit points and death |
| `init-questions.md` | the 5E part of world setup |
| `templates/system.md` | the world's `system.md` |

## Seam (core → pack)
1. The skill writes the **fiction block** first, without numbers.
2. Then it appends a **`5E` block** from this pack.
3. Homebrew in the world's `system.md` overrides anything here.

## Never
- Roll dice, decide success or failure, or tell the GM what the players must roll. Say "suggested DC" and "fitting checks".
- Award or spend Heroic Inspiration on the GM's behalf. You may remind: "this could earn Heroic Inspiration".
- Quote spell, magic item or monster entries. If a rule depends on one, name it and tell the GM to check their book.
- Use content from a published setting. The GM's world is the only setting.
