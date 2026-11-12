# 2E: Entity Sections

The core templates hold the fiction (`## Who`, `## Wants`, `## Links` …). This file says which **system sections** to add where a template has the `{{system sections}}` line.

## Character
| Section | NPC | PC | Content |
|---|---|---|---|
| `## Stat Line` | ✓ | — | one line, see "NPC stat line" |
| `## Ancestry, Background, Class` | — | ✓ | ancestry, heritage, background, class, level |
| `## Attributes` | — | ✓ | the six attribute modifiers (Str, Dex, Con, Int, Wis, Cha) |
| `## Proficiencies` | — | ✓ | ranks for Perception, saves, skills, attacks, defenses, class DC |
| `## Defenses` | — | ✓ | AC, Hit Points, Fortitude, Reflex, Will, Speed |
| `## Feats and Features` | — | ✓ | ancestry, class, skill and general feats; class features (names and one-line effects) |
| `## Spells` | — | if a caster | tradition, spell DC, spell attack, spells by name |

A minor NPC needs no stat line at all: it only needs one when it may fight, resist or be tested against.

### NPC stat line
`Level <n> · Perception +<n> · AC <n> · Fort +<n> Ref +<n> Will +<n> · HP <n> · Speed <n> ft · Skills: <skill +n> · Attack: <name> +<n> (<damage>)`
For a creature from the GM's own book, write `See: <creature name>, level <n>` instead and only the differences. A quick NPC can use the level-based DC table: its best DC ≈ the DC of its level.

## Location
`## Features`: 1-3 permanent circumstances of the place (see `scene-tags.md`), e.g. darkness, difficult terrain, a standing hazard with its level, a typical DC (a guarded gate: Deception against the guards' Perception DC 19).

## Group
No system section. Write `—` in place of the `{{system sections}}` line.

## Character math
- **Attribute modifiers** start at +0. Boosts add +1, a flaw −1: ancestry (usually two boosts, sometimes a flaw), background (two boosts), class (one boost to the key attribute), then four free boosts. Each set of boosts must go to different attributes. At 1st level no modifier goes above +4.
- **Later boosts:** four boosts at levels 5, 10, 15 and 20. A boost to a modifier that is already +4 or more is a **partial boost**; two partial boosts raise it by 1.
- **Proficiency bonus:** untrained +0; trained, expert, master, legendary = level + 2 / 4 / 6 / 8.
- **Hit Points:** ancestry HP + (class HP + Constitution modifier) × level.
- **Armor Class:** 10 + Dexterity modifier (up to the armor's cap) + proficiency bonus + the armor's item bonus.
- **Saves and Perception:** attribute modifier + proficiency bonus (Fortitude Con, Reflex Dex, Will Wis, Perception Wis).
- **Class DC:** 10 + key attribute modifier + proficiency bonus. **Spell DC:** 10 + spellcasting attribute + proficiency bonus. **Spell attack:** spellcasting attribute + proficiency bonus.
- **Strike:** attribute modifier (Str melee, Dex ranged or finesse) + proficiency bonus + item bonus.
- **Advancement:** 1,000 XP per level (or milestones if `system.md` says so).

## Hit points, dying and recovery
- **0 Hit Points:** the creature is knocked out and gains **dying 1** (dying 2 from a critical hit or a critically failed save), plus its **wounded** value. A minor NPC or monster usually just dies.
- **Recovery check** at the start of each turn while dying: a flat check against DC 10 + the dying value. Critical success: dying −2; success: dying −1; failure: dying +1; critical failure: dying +2. Taking damage while dying raises it by 1 (2 on a critical).
- **Dying 4** is death; **doomed X** lowers that threshold by X.
- Losing the dying condition gives **wounded 1**, or raises wounded by 1. Wounded ends after successful healing with Treat Wounds, or after full HP and 10 minutes of rest.
- **Rest:** a full night's rest restores Hit Points (Constitution modifier × level, minimum 1) and lowers drained and doomed by 1. **Treat Wounds** (Medicine, 10 minutes) heals out of combat.
