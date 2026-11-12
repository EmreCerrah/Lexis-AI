# FitD: Entity Sections

The core templates hold the fiction (`## Who`, `## Wants`, `## Links` …). This file says which **system sections** to add where a template has the `{{system sections}}` line.

## Character
| Section | NPC | PC | Content |
|---|---|---|---|
| `## Threat` | ✓ | — | one line, see "NPC line" |
| `## Playbook` | — | ✓ | the character's role (the GM's own playbooks, or the GM's book), heritage/background, vice |
| `## Actions` | — | ✓ | the 12 action ratings (0-4 dots), grouped by attribute |
| `## Stress and Harm` | — | ✓ | stress (of 9), trauma conditions, harm by level, armor, load |
| `## Special Abilities` | — | ✓ | names and one-line effects |

A minor NPC needs no line at all. NPCs don't roll like PCs: their strength shows in position, effect and clocks.

### NPC line
`Scale: <a person / a small gang / a large gang> · Quality: <tier or gear> · Danger: <what they threaten, harm level> · Clock: "<name>" (<4/6/8>) to defeat or outwit them`

## Location
`## Features`: 1-3 standing factors or clocks of the place (see `scene-tags.md`), e.g. "Patrols" danger clock (6), darkness (potency for Prowl), fortified doors (quality).

## Group
`## Clocks`: 1-2 clocks for what the group is working toward or how its trouble grows (e.g. "Seize the Docks" 8). The GM ticks them between sessions. Write `—` if none.

## Character math
- **Action ratings:** 0-4 dots each. A new PC typically starts with 7 dots spread across actions, none above 2 at creation (check the GM's book or `system.md`).
- **Attribute rating:** the number of actions in that attribute with at least 1 dot (Insight, Prowess, Resolve); used for resistance.
- **Stress:** 9 boxes. When the last box is marked, the PC takes a **trauma** condition and is out of action for the scene; after the fourth trauma, the PC retires. Trauma conditions: Cold, Haunted, Obsessed, Paranoid, Reckless, Soft, Unstable, Vicious.
- **Harm:** level 1 (two slots): less effect when it applies; level 2 (two slots): −1d when it applies; level 3 (one slot): needs help to act; level 4: fatal or catastrophic. A full level pushes new harm up a level.
- **Recovery:** in downtime, a healing clock (4 segments) clears harm one level at a time.
- **Load:** light 3, normal 5, heavy 6 items for a job.
- **Advancement:** mark xp when making a desperate roll, plus end-of-session triggers; a full track gives a new action dot or special ability.

## Vice and stress relief
Between jobs, a PC indulges their vice to clear stress (a roll of the lowest attribute; overindulging has a cost). Use the world's own vices and places.
