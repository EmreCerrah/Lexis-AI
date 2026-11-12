---
type: system-pack
id: fate-core
name: Fate Core
order: 5
version: 0.1.0
---

# System Pack: Fate Core

Fate Core mechanics on top of the world core. The core produces fiction; this pack adds skills, aspects, the ladder, stress and consequences. Use this pack when the GM asks for Fate Core (skills instead of approaches). For the approach-based version, use the Fate Accelerated pack.

## License
This work is based on Fate Core System and Fate Accelerated Edition, products of Evil Hat Productions, LLC, developed, authored, and edited by Leonard Balsera, Brian Engard, Jeremy Keller, Ryan Macklin, Mike Olson, Clark Valentine, Amanda Valentine, Fred Hicks, and Rob Donoghue, and licensed for our use under the Creative Commons Attribution 3.0 Unported license (https://creativecommons.org/licenses/by/3.0/).

Fate™ is a trademark of Evil Hat Productions, LLC.

## Content scope
Mechanics only: dice, skills, actions, outcomes, aspects, stunts, stress and consequences, character math. No setting content: no worlds, places, characters, creatures or story from any published book. Rules are summarized in our own words.

## Glossary
| Role | Term in this system |
|---|---|
| check | skill roll (4dF + skill) against opposition on the ladder |
| trait | skill (Athletics, Burglary, Contacts …) |
| difficulty | the ladder (+0 Mediocre … +8 Legendary), passive or active opposition |
| scene-tag | aspect (situation aspect, boost, consequence) |
| special-ability | stunt |
| harm | physical and mental stress tracks + consequences |
| meta-currency | fate points |

Keep these terms in English in every output language (`AGENTS.md` §3).

## Files
| File | Content |
|---|---|
| `resolution.md` | the roll, the skill list and what each skill can do, the four actions, outcomes, the ladder, teamwork, fate points, the `/help-skills` answer shape |
| `scene-tags.md` | how to write, invoke and compel aspects; the aspect output block |
| `obstacles.md` | difficulty, opposition, challenges, contests, hazards, the obstacle output block |
| `entities.md` | the system sections of character, location and group files; skill pyramid, stunts, stress, consequences |
| `init-questions.md` | the Fate Core part of world setup |
| `templates/system.md` | the world's `system.md` |

## Seam (core → pack)
1. The skill writes the **fiction block** first, without numbers.
2. Then it appends a **`Fate` block** from this pack.
3. Homebrew in the world's `system.md` overrides anything here (a custom skill list replaces the default one).

## Never
- Roll dice, decide success or failure, or tell the GM what the players must roll. Say "suggested difficulty" and "fitting skills".
- Spend or award fate points on the GM's behalf. You may remind: "this could be a compel → +1 fate point".
- Use approaches. Fate Core uses skills; approaches belong to the Fate Accelerated pack.
