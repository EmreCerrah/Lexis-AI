---
type: system-pack
id: fate-accelerated
name: Fate Accelerated
order: 4
version: 0.2.0
---

# System Pack: Fate Accelerated (FAE)

Fate Accelerated mechanics on top of the world core. The core produces fiction; this pack adds approaches, aspects and the ladder. Use this pack when the GM asks for Fate Accelerated or FAE.

## License
Based on Fate Accelerated Edition (found at http://www.faterpg.com/), a product of Evil Hat Productions, LLC, developed, authored, and edited by Clark Valentine, Leonard Balsera, Fred Hicks, Mike Olson, and Amanda Valentine, and licensed for our use under the Creative Commons Attribution 3.0 Unported license (http://creativecommons.org/licenses/by/3.0/).

## Content scope
Mechanics only: dice, approaches, actions, outcomes, aspects, stress and consequences, character math. No setting content: no worlds, places, characters, creatures or story from any published book. Rules are summarized in our own words.

## Glossary
| Role | Term in this system |
|---|---|
| check | approach roll (4dF + approach) against a difficulty |
| trait | approach (Careful, Clever, Flashy, Forceful, Quick, Sneaky) |
| difficulty | the ladder (+0 Mediocre … +8 Legendary), passive or active opposition |
| scene-tag | aspect (situation aspect, boost, consequence) |
| special-ability | stunt |
| harm | stress boxes + consequences |
| meta-currency | fate points |

Keep these terms in English in every output language (`AGENTS.md` §3).

## Files
| File | Content |
|---|---|
| `resolution.md` | approaches, the four actions, outcomes, the ladder, fate points, the `/help-skills` answer shape |
| `scene-tags.md` | how to write, invoke and compel aspects; the aspect output block |
| `obstacles.md` | difficulty, passive vs active opposition, hazards as characters, the obstacle output block |
| `entities.md` | the system sections of character, location and group files; stunts, stress, consequences |
| `init-questions.md` | the FAE part of world setup |
| `templates/system.md` | the world's `system.md` |

## Seam (core → pack)
1. The skill writes the **fiction block** first, without numbers.
2. Then it appends a **`Fate` block** from this pack.
3. Homebrew in the world's `system.md` overrides anything here.

## Never
- Roll dice, decide success or failure, or tell the GM what the players must roll. Say "suggested difficulty" and "fitting approaches".
- Spend or award fate points on the GM's behalf. You may remind: "this could be a compel → +1 fate point".
- Add Fate Core skills. FAE uses approaches. A custom skill list exists only if Homebrew in `system.md` defines one.
