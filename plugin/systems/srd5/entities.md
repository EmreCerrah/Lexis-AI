# 5E: Entity Sections

The core templates hold the fiction (`## Who`, `## Wants`, `## Links` …). This file says which **system sections** to add where a template has the `{{system sections}}` line.

## Character
| Section | NPC | PC | Content |
|---|---|---|---|
| `## Stat Line` | ✓ | — | one line, see "NPC stat line" |
| `## Class and Level` | — | ✓ | class, level, subclass, species, background |
| `## Abilities` | — | ✓ | the six scores with modifiers |
| `## Proficiencies` | — | ✓ | Proficiency Bonus, saving throws, skills (mark Expertise), weapons, armor, tools |
| `## Combat` | — | ✓ | Armor Class, Hit Points, Hit Point Dice, Speed, Initiative, Passive Perception |
| `## Features` | — | ✓ | class features, feats, species traits (names and one-line effects) |
| `## Spells` | — | if a caster | spellcasting ability, spell save DC, spell attack bonus, prepared spells by name |

A minor NPC needs no stat line at all: it only needs one when it may fight, resist or be tested against.

### NPC stat line
`AC <n> · HP <n> · Speed <n> ft · Str <±n> Dex <±n> Con <±n> Int <±n> Wis <±n> Cha <±n> · Skills: <skill ±n> · Attack: <name> +<n> (<damage>) · Passive Perception <n>`
For a creature from the GM's own book, write `See: <creature name>` instead and only the differences.

## Location
`## Features`: 1-3 permanent scene features of the place (see `scene-tags.md`), e.g. darkness, Difficult Terrain, a standing hazard, a typical DC (a guarded gate: DC 15 to bluff past).

## Group
No system section. Write `—` in place of the `{{system sections}}` line.

## Character math
- **Ability scores:** standard array 15, 14, 13, 12, 10, 8; or roll 4d6 and drop the lowest die, six times; or point cost with 27 points (8 costs 0, 9: 1, 10: 2, 11: 3, 12: 4, 13: 5, 14: 7, 15: 9). The GM picks the method (`system.md`).
- **Ability modifier:** (score − 10) ÷ 2, rounded down. 8-9 → −1, 10-11 → +0, 12-13 → +1, 14-15 → +2, 16-17 → +3, 18-19 → +4, 20 → +5.
- **Proficiency Bonus by level:** 1-4: +2, 5-8: +3, 9-12: +4, 13-16: +5, 17-20: +6. **Expertise** doubles it for that skill.
- **Skill or save bonus:** ability modifier + Proficiency Bonus if proficient.
- **Hit Points:** level 1 = the maximum of the class's Hit Die + Constitution modifier. Each later level: roll the Hit Die + Con modifier (minimum 1), or take the fixed value: d12 → 7, d10 → 6, d8 → 5, d6 → 4 (+ Con modifier).
- **Hit Die by class:** Barbarian d12; Fighter, Paladin, Ranger d10; Bard, Cleric, Druid, Monk, Rogue, Warlock d8; Sorcerer, Wizard d6.
- **Armor Class:** without armor 10 + Dexterity modifier; armor or features replace this base (only one base applies).
- **Initiative:** Dexterity modifier. **Passive Perception:** 10 + Wisdom (Perception) bonus.
- **Weapon attack bonus:** Strength (melee) or Dexterity (ranged, finesse) modifier + Proficiency Bonus.
- **Spell save DC:** 8 + spellcasting ability modifier + Proficiency Bonus. **Spell attack bonus:** spellcasting ability modifier + Proficiency Bonus.
- **Advancement (XP to reach a level):** 2: 300, 3: 900, 4: 2,700, 5: 6,500, 6: 14,000, 7: 23,000, 8: 34,000, 9: 48,000, 10: 64,000; or level by milestones if `system.md` says so.

## Hit points, death and rest
- **0 Hit Points:** a monster or minor NPC dies (the GM may treat an important NPC like a PC). A PC falls Unconscious, unless the remaining damage equals or exceeds their Hit Point maximum, which kills them outright.
- **Death Saving Throws:** at the start of each turn at 0 HP, roll a d20: 10+ is a success, below 10 a failure. Three successes: Stable. Three failures: dead. A 1 counts as two failures; a 20 restores 1 Hit Point. Damage at 0 HP is a failure (two on a Critical Hit).
- **Stabilize:** a DC 10 Wisdom (Medicine) check. A Stable creature regains 1 HP after 1d4 hours.
- **Short Rest (1 hour):** spend Hit Point Dice to heal (roll each + Con modifier).
- **Long Rest (8 hours):** regain all Hit Points and all spent Hit Point Dice; remove one level of Exhaustion.
