# Fate Core: Entity Sections

The core templates hold the fiction (`## Who`, `## Wants`, `## Links` …). This file says which **system sections** to add where a template has the `{{system sections}}` line.

## Character
| Section | NPC | PC | Content |
|---|---|---|---|
| `## Aspects` | ✓ | ✓ | high concept, trouble, up to 3 others (see `scene-tags.md`) |
| `## Skills` | ✓ (1-3) | ✓ | the skill pyramid |
| `## Stunts` | — | ✓ | up to 5 stunts |
| `## Stress` | — | ✓ | physical and mental stress boxes, consequence slots, refresh |

A minor NPC needs only 1-2 aspects and one or two skills: `Good (+3) Fight, Fair (+2) Notice; stress: 1`. A major NPC is built like a PC.

## Location
`## Aspects`: 1-2 permanent aspects of the place.

## Group
`## Aspects`: 1-2 permanent aspects of the group.

## Character math
- **Skill pyramid:** one at Great (+4), two at Good (+3), three at Fair (+2), four at Average (+1); every other skill is Mediocre (+0).
- **Refresh:** 3, with 3 free stunts. Up to 2 more stunts cost 1 refresh each (refresh never below 1).
- **Stress:** 2 physical and 2 mental boxes (values 1 and 2). Physique adds physical boxes and Will adds mental boxes: Average or Fair → a 3rd box; Good or Great → a 3rd and 4th box; Superb or higher → also an extra mild consequence slot for that kind of harm.
- **Consequences:** mild (absorbs 2), moderate (4), severe (6). A consequence is an aspect.
- Check `system.md` Homebrew before using these numbers.

## Stunts
Format: *"Because I [describe how I'm awesome], I get a +2 when I use [skill] to [action] when [circumstance]."*, *"Because I [...], I can use [skill] instead of [skill] when [circumstance]."* or *"Because I [...], once per session I can [do something cool]."* Look up the PC's stunts before answering; a matching stunt is the most useful thing to mention.

## Harm
- An attack's shifts are absorbed by checking **one** stress box of at least that value, and/or by taking consequences. What isn't absorbed takes the character out.
- **Taken out:** the attacker decides how the character leaves the scene (the GM rules on it).
- **Concede:** before the roll, the player gives up on their own terms and gets 1 fate point, plus 1 for each consequence taken in the conflict.
- **Recovery:** stress clears after the scene. A consequence needs a recovery roll (mild Fair +2, moderate Great +4, severe Fantastic +6, against Empathy for mental or a fitting skill like Lore for physical); then mild clears after one whole scene, moderate after one whole session, severe after one whole scenario.
