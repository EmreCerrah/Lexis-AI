# d100: Entity Sections

The core templates hold the fiction (`## Who`, `## Wants`, `## Links` …). This file says which **system sections** to add where a template has the `{{system sections}}` line.

## Character
| Section | NPC | PC | Content |
|---|---|---|---|
| `## Stat Line` | ✓ | — | one line, see "NPC stat line" |
| `## Characteristics` | — | ✓ | STR, CON, SIZ, INT, POW, DEX, CHA (EDU if used) and the characteristic rolls |
| `## Derived` | — | ✓ | hit points, major wound level, power points, damage modifier, MOV, experience bonus; Sanity and TIS if used |
| `## Profession and Skills` | — | ✓ | profession, skill ratings (%) |
| `## Powers` | — | if any | power type and the powers by name |

A minor NPC needs no stat line at all: it only needs one when it may fight, resist or be tested against.

### NPC stat line
`STR <n> CON <n> SIZ <n> INT <n> POW <n> DEX <n> · HP <n> · Damage mod <±dice> · Armor <n> · Attack: <weapon> <n>% (<damage>) · Skills: <skill n%>, <skill n%>`
For a creature from the GM's own book, write `See: <creature name>` and only the differences.

## Location
`## Features`: 1-3 standing situational modifiers of the place (see `scene-tags.md`), e.g. darkness −20% to Spot, a flooded floor making Dodge Difficult, a door of STR 16.

## Group
No system section. Write `—` in place of the `{{system sections}}` line.

## Character math
- **Characteristics:** STR, CON, POW, DEX, CHA roll 3D6; INT and SIZ roll 2D6+6 (EDU 2D6+6 if used). Up to 3 points may be moved between them; none above 21 at start. Option: 2D6+6 for all in a higher-powered game.
- **Characteristic rolls:** each characteristic × 5 (Effort, Stamina, Idea, Luck, Agility, Charm).
- **Hit points:** (CON + SIZ) ÷ 2, round up. Option: CON + SIZ for hardier heroes. **Major wound level:** half of hit points, round up.
- **Power points:** equal to POW. **Experience bonus:** INT ÷ 2, round up. **MOV:** 10 for humans.
- **Damage modifier (STR + SIZ):** 2-12 −1D6, 13-16 −1D4, 17-24 none, 25-32 +1D4, 33-40 +1D6, 41-56 +2D6, 57-72 +3D6, then +1D6 per 16 more.
- **Sanity (optional):** POW × 5. **Temporary insanity score:** half of the starting Sanity.
- **Skills:** base chance + professional points (Normal 250, Heroic 325, Epic 400, Superhuman 500) + personal points (INT × 10). Starting cap: 75% Normal, 90% Heroic, 101% Epic, none Superhuman.
- **Experience:** a skill used successfully under pressure gets an experience check (once per adventure, never on Easy rolls). After the adventure, roll d100 + experience bonus; above the current rating, add 1D6 (or a flat +3).

## Wounds and healing
- **Minor wound:** less than the major wound level in one blow. Minor wounds adding up to the major wound level call for a Luck roll or the character falls unconscious.
- **Major wound:** one blow of half the hit points or more: a lasting injury may follow (a Luck roll decides if it heals cleanly); the character can keep acting only for as many rounds as hit points remain.
- **Unconscious** at 2 or fewer hit points. **Fatal wound** at 0 or below: the character dies unless First Aid, Medicine or similar raises them to 1+ hit points in that round or the next.
- **Healing:** 1D3 hit points per week naturally; a successful First Aid restores 1D3 per wound at once.
