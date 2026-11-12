# Doom d20: Entity Sections

The core templates hold the fiction (`## Who`, `## Wants`, `## Links` …). This file says which **system sections** to add where a template has the `{{system sections}}` line.

## Character
| Section | NPC | PC | Content |
|---|---|---|---|
| `## Stat Line` | ✓ | — | one line, see "NPC stat line" |
| `## Abilities` | — | ✓ | Strength, Agility, Presence, Toughness (−3 to +3) |
| `## Vitals` | — | ✓ | hit points, Omens per day, armor tier |
| `## Class and Gear` | — | ✓ | class or classless, class abilities, powers (by name), weapons, items, silver |

A minor NPC needs no stat line at all: it only needs one when it may fight or be fought.

### NPC stat line
`HP <n> · Morale <n> · Armor <tier, −dX> · Attack: <weapon> <damage> · Defend DR <n> · Special: <one line>`

## Location
`## Features`: 1-3 standing DR shifts or hazards of the place (see `scene-tags.md`), e.g. darkness (seeing DR +2), a rotten floor (Agility DR 12 or fall, d6).

## Group
No system section. Write `—` in place of the `{{system sections}}` line.

## Character math
- **Abilities:** roll 3d6 for each (classes may change this), then read the modifier: 1-4 → −3, 5-6 → −2, 7-8 → −1, 9-12 → ±0, 13-14 → +1, 15-16 → +2, 17-20 → +3.
- **Hit points:** Toughness + d8 for a classless character (classes change the die); at least 1.
- **Omens:** d2 per day for a classless character (classes change the die).
- **Armor:** none, light (−d2 damage), medium (−d4, Agility DR +2), heavy (−d6, Agility DR +4).
- **Getting better:** after surviving an adventure, the GM may let characters improve hit points and abilities (rolls from the GM's book).

## Harm and rest
- **0 hit points: broken.** The character is out of the fight and rolls on the broken table in the GM's book (knocked out, a broken limb or lost eye, bleeding out, or death).
- **Below 0: dead.**
- **Rest:** a short rest (food, water, a break) restores d4 hit points; a full night's rest restores d6 and the day's Omens. Without food or water, no healing.
