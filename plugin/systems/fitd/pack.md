---
type: system-pack
id: fitd
name: Forged in the Dark Core
order: 7
version: 0.1.0
---

# System Pack: Forged in the Dark Core

The Forged in the Dark rules engine (from the CC-BY SRD) on top of the world core. The core produces fiction; this pack adds d6 dice pools, position and effect, consequences, stress and harm, and progress clocks. Use this pack when the GM asks for Forged in the Dark rules or a d6-pool game with position, effect and clocks.

## License
This work is based on Blades in the Dark (found at http://www.bladesinthedark.com/), product of One Seven Design, developed and authored by John Harper, and licensed for our use under the Creative Commons Attribution 3.0 Unported license (http://creativecommons.org/licenses/by/3.0/).

Blades in the Dark™ is a trademark of One Seven Design. This pack is not affiliated with or endorsed by One Seven Design.

## Content scope
Mechanics only: action ratings, dice pools, position and effect, consequences, resistance, stress, harm, teamwork, clocks, fortune and engagement rolls. No setting content: none of the original city, its districts, factions, NPCs, ghosts, history, playbook flavor, artwork or maps (the SRD license excludes them). Rules are summarized in our own words.

## Glossary
| Role | Term in this system |
|---|---|
| check | action roll: a d6 pool, read the highest die (6 full, 4-5 partial, 1-3 bad, two 6s critical) |
| trait | action rating (0-4 dots in one of 12 actions); attributes for resistance |
| difficulty | position (controlled, risky, desperate) and effect (limited, standard, great) set by the GM |
| scene-tag | a factor of position or effect (potency, scale, quality), or a progress clock |
| special-ability | special ability (playbook ability) |
| harm | stress (9 boxes) and trauma; harm levels 1-4; armor |
| meta-currency | stress spent to push, assist, resist and flash back |

Keep these terms in English in every output language (`AGENTS.md` §3); add a short gloss the first time if the output language differs.

## Files
| File | Content |
|---|---|
| `resolution.md` | dice, the 12 actions, position and effect, outcomes, pushing, teamwork, fortune roll, the `/help-skills` answer shape |
| `scene-tags.md` | position and effect factors and clocks as scene tags; the output block |
| `obstacles.md` | choosing position and effect, clocks, consequences, resistance, the obstacle output block |
| `entities.md` | the system sections of character, location and group files; harm, stress, trauma, advancement |
| `init-questions.md` | the setup questions for this system |
| `templates/system.md` | the world's `system.md` |

## Seam (core → pack)
1. The skill writes the **fiction block** first, without numbers.
2. Then it appends a **`FitD` block** from this pack.
3. Homebrew in the world's `system.md` overrides anything here.

## Never
- Roll dice, decide outcomes, set the final position or effect, or tick a clock on the GM's behalf. Say "suggested position" and "suggested effect".
- Spend stress or choose consequences for the players.
- Use anything from the original game's setting (its city, factions, ghosts, people). The GM's world is the only setting.
