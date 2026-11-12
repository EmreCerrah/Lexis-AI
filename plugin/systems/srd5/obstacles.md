# 5E: Obstacles

## Choosing a DC (suggest, never decide)
| Situation | Suggested DC |
|---|---|
| Easy, but failure would still be interesting | 10 |
| A real challenge for a capable character | 15 |
| Hard; a specialist has a fair chance | 20 |
| Only for experts with help or Advantage | 25 |
| Nearly impossible | 30 |

Rules of thumb: if failure is not interesting, don't call for a roll, and say so. Give one DC with the reason, and name what would grant Advantage or impose Disadvantage instead of moving the DC around.

## Passive vs active
- **Passive:** a fixed DC (a wall, a lock, a storm). Hidden things can be found with Passive Perception against the DC.
- **Active:** someone opposes (a guard searching, a rival climbing): a contest of checks. Suggest the NPC's check and its modifier.

## Obstacle types
| Type | What to produce |
|---|---|
| Barrier (lock, wall, chasm) | the feature, a DC, 2-3 fitting checks with the fiction for each, what grants Advantage |
| Puzzle / riddle | the clue chain (fiction), what each check reveals (Investigation reads, Perception spots, History recalls), a DC; solvable without dice too |
| Trap | severity, trigger, effect (save + DC + damage or condition), how to detect and disarm it (see below) |
| Hazard (fire, collapse, cold, curse) | the saving throw, DC and effect each round or each exposure; how to escape it |
| Fear / mental stress | a Wisdom save; on a failure the Frightened condition (or Psychic damage) |
| Social obstacle (crowd, gatekeeper) | the NPC's attitude (Friendly / Indifferent / Hostile), what they are willing, hesitant or unwilling to do, the check and DC for "hesitant" |
| Challenge (multi-step) | 2-4 steps with different checks, each with its own DC, and what failing a step costs |

## Traps and hazards (scale with level)
A trap is a **nuisance** (unlikely to seriously harm characters of the level) or **deadly** (can grievously harm them). Detecting a trap is usually a Wisdom (Perception) or Intelligence (Investigation) check around DC 10-15; disarming it is a Dexterity (Sleight of Hand) check with thieves' tools, or a clever plan with no roll.

| Character levels | Deadly: save DC | Deadly: damage |
|---|---|---|
| 1-4 | 13 | 2d10 (≈11) |
| 5-10 | 15 | 4d10 (≈22) |
| 11-16 | 17 | 10d10 (≈55) |
| 17-20 | 19 | 18d10 (≈99) |

A nuisance trap uses about DC 10 and a condition (Restrained, Prone) or minor damage instead. On a successful save, damaging traps usually deal half damage.

## Fear and mental stress
| How terrifying | Wisdom save DC | Psychic damage (if used) |
|---|---|---|
| A harmless but frightening apparition | 10 | 1d6 |
| A personal, magical terror | 15 | 3d6 |
| A glimpse of something mind-breaking | 20 | 9d6 |

On a failure: the Frightened condition (repeat the save at the end of each turn) or the damage. Lasting effects (a phobia, Disadvantage on some checks for hours) only if the GM uses them; agree on these rules with the players first.

## Cost menu (for near misses)
A GM technique, not a rule: when a check fails by a little, offer success at a cost. 2-3 costs that fit the fiction: lose time, make noise, break a tool, take a level of Exhaustion, owe someone, give away your position. Costs may feed a hook's `## Surface` but never expose its `## 🔒 Hidden Truth`.

## Output block
```
**5E**
- Feature: "<…>" — <rule effect>
- Suggested DC: <n> (<Easy/Medium/Hard …>) — <why>
- Fitting checks: <Ability (Skill)> — <how>; <Ability (Skill)> — <how>
- Advantage if: <…> · Disadvantage if: <…>
- (If a trap/hazard) Save: <Ability> DC <n> · Effect: <damage or condition> · Detect/disarm: <check> DC <n>
- Cost options: <…>; <…>
```
