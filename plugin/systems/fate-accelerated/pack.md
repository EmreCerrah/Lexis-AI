# System Pack: Fate Accelerated (FAE)

This pack adds Fate Accelerated mechanics on top of the world core. The core produces fiction; this pack adds numbers, approaches and aspect mechanics. No plugin API: a skill that needs mechanics reads this folder directly.

## Files
| File | Used by | Content |
|---|---|---|
| `approaches.md` | `/help-skills`, `/create-obstacle` | the six approaches, four actions, outcomes, the ladder |
| `aspects.md` | `/create-aspect`, `/create-lore`, `/world-init` | how to write, invoke and compel aspects |
| `obstacles.md` | `/create-obstacle` | difficulty, passive vs active opposition, hazards as characters |
| `init-questions.md` | `/world-init` | the FAE part of world setup |
| `templates/system.md` | `/world-init` | the world's `system.md` |

## Seam (core → pack)
1. The skill writes the **fiction block** first, without numbers: what it is, how it looks, why it fits the world, which `## Değişmez Kurallar` and hooks it respects.
2. Then it appends a **`Fate` block** from this pack: aspects, suggested difficulty, fitting approaches, stress/consequences if relevant, invoke/compel ideas.
3. The world's `system.md` Homebrew overrides anything here.

## Never
- Roll dice, decide success/failure, or tell the GM what the players must roll. Say "önerilen zorluk", "uygun approach'lar".
- Spend or award fate points on the GM's behalf. You may remind: "bu bir compel olabilir → +1 fate point".
- Add Fate Core skills. FAE uses approaches. A custom skill list exists only if `system.md` Homebrew defines one.
