# d20 Hero: Obstacles

## Choosing a DC (suggest, never decide)
- Use the DC table (`resolution.md`): average 10, tough 15, challenging 20, formidable 25, heroic 30.
- For something with a rank (a trap's damage, a poison, a force field): the resistance DC is **10 + rank**, damage **15 + rank**.
- Rule of thumb: if failure is not interesting, don't call for a roll, or make it routine (treat the roll as 10). Express the scene through ±2/±5 circumstances rather than moving the DC around.

## Degrees, every time
Give a short result for the degrees that matter: one degree (just made it), two or three (extra results: faster, more information, more people saved); one degree of failure (no progress), two or more (a setback, a condition, damage).

## Damage
Damage of rank R: the target makes a **Toughness** check against **15 + R**:
| Result | Effect |
|---|---|
| Success | no effect |
| 1 degree of failure | −1 to further Toughness checks against damage |
| 2 degrees | dazed until the end of its next turn, and −1 |
| 3 degrees | staggered, and −1 (staggered again → incapacitated) |
| 4 degrees | incapacitated |

The −1 penalties add up. An incapacitated target that fails again is dying; a dying target that fails again is dead. **Minions** take the worst result of any failed resistance check.

## Obstacle types
| Type | What to produce |
|---|---|
| Barrier (vault door, wall, force field) | its Toughness rank (an inch of steel is far tougher than wood), the check or effect to get through, a DC for alternatives (Technology to bypass) |
| Trap / hazard | how it is noticed (Perception DC), its effect rank and the resistance check (Dodge, Fortitude, Toughness, Will) at 10 + rank or 15 + rank for damage |
| Falling | damage rank 4 + twice the distance rank fallen (max 16); Acrobatics can lessen it |
| Puzzle / investigation | the clue chain (fiction), what Investigation, Expertise or Technology checks reveal by degree |
| Social obstacle (crowd, gatekeeper) | the opposing Insight or Will, which interaction skill fits, circumstances that help or hurt |
| Disaster in progress | a few checks against rising DCs, team checks, and what each degree of failure costs |

## Cost menu (for near misses)
A GM technique: on one degree of failure, offer success at a cost. 2-3 costs that fit the fiction: collateral damage, a bystander in danger, the villain escapes with something, the hero's complication flares up (and earns a Victory Point). Costs may feed a hook's `## Surface` but never expose its `## 🔒 Hidden Truth`.

## Output block
```
**d20 Hero**
- Tag: "<…>" — <circumstance or condition>
- Fitting traits: <Trait> — <how>; <Trait> — <how>
- Suggested DC: <n> (<name>) — <why>
- Degrees: 1 <…> · 2-3 <…> · failure <…>
- (If trap/hazard) Notice: Perception DC <n> · Resist: <defense> DC <10/15 + rank> · Effect: <condition or damage rank>
- Cost options: <…>; <…>
```
