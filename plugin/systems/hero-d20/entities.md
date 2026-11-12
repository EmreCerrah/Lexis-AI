# d20 Hero: Entity Sections

The core templates hold the fiction (`## Who`, `## Wants`, `## Links` …). This file says which **system sections** to add where a template has the `{{system sections}}` line.

## Character
| Section | NPC | PC | Content |
|---|---|---|---|
| `## Stat Line` | ✓ | — | one line, see "NPC stat line" |
| `## Power Level` | — | ✓ | power level, Character Points spent |
| `## Abilities` | — | ✓ | Strength, Stamina, Agility, Dexterity, Fighting, Intellect, Awareness, Presence |
| `## Defenses` | — | ✓ | Dodge, Parry, Fortitude, Toughness, Will, initiative |
| `## Skills and Advantages` | — | ✓ | skill ranks and totals; advantages by name |
| `## Powers` | — | ✓ | powers by name with their effect and rank (write-ups stay in the GM's book) |
| `## Complications` | ✓ | ✓ | motivation, weaknesses, obligations, secrets: what earns Victory Points |

A minor NPC needs no stat line at all. Henchmen and mooks are usually **minions** (any failed resistance check takes the worst result).

### NPC stat line
`PL <n> · Attack +<n> (<effect> rank <n>) · Dodge <n> Parry <n> Toughness <n> Fortitude <n> Will <n> · Key skills: <skill +n> · Powers: <name, rank> · Minion: <yes/no>`

## Location
`## Features`: 1-3 standing circumstances or hazards of the place (see `scene-tags.md`), e.g. darkness (−2 Perception), reinforced walls (Toughness rank), a security system (Technology DC 25).

## Group
No system section. Write `—` in place of the `{{system sections}}` line.

## Character math
- **Character Points:** 15 per power level (power level 10 → 150).
- **Costs:** abilities 2 per rank; defenses 1 per rank; skills 1 per 2 ranks; advantages 1 per rank; powers by effect and modifiers (GM's book).
- **Defenses** start from abilities: Dodge from Agility, Parry from Fighting, Fortitude and Toughness from Stamina, Will from Awareness. Initiative uses Agility.
- **Power level limits:** attack bonus + effect rank ≤ 2 × PL; Dodge + Toughness ≤ 2 × PL; Parry + Toughness ≤ 2 × PL; Fortitude + Will ≤ 2 × PL; a skill's total modifier ≤ PL + 10.
- **Advancement:** the GM awards Character Points between adventures (usually 1 per adventure); power level rises with the series.

## Recovery
- Damage conditions clear one per minute of rest, worst first: incapacitated, then staggered, then dazed, then each −1 Toughness penalty.
- **Dying:** make a Fortitude check (DC 15) when it happens and on later turns: success holds, two degrees of success stabilizes; failures move toward death (the GM's book has the details). Treatment and healing powers help.
- Fatigued and exhausted clear after an hour of rest.
