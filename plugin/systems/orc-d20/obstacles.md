# 2E: Obstacles

## Choosing a DC (suggest, never decide)
- A task that does not depend on level: a **simple DC** (trained 15, expert 20, master 30).
- A task tied to a level (a lock built by a master, a hazard, a creature): the **level-based DC** of that level, then an adjustment (`resolution.md`).
- Rule of thumb: if failure is not interesting, don't call for a check, and say so. Give one DC with its reason; express the scene through circumstance bonuses and penalties instead of moving the DC around.

## Four degrees, every time
For every obstacle, give a short result for each degree: critical success (a bonus: faster, quieter, an extra clue), success, failure (no progress, or progress at a cost), critical failure (a real setback: damage, a fall, an alarm).

## Passive vs active
- **Passive:** a fixed DC (a wall, a lock, a storm).
- **Active:** someone opposes: the check goes against their DC (their Perception DC, Will DC or Reflex DC = 10 + their modifier).

## Obstacle types
| Type | What to produce |
|---|---|
| Barrier (lock, wall, chasm) | the DC, 2-3 fitting skills with the fiction for each, the minimum rank if any, the four degrees |
| Puzzle / riddle | the clue chain (fiction), what each Recall Knowledge or Perception check reveals, a DC; solvable without dice too |
| Hazard (trap, environmental, haunt) | level, how to notice it (Stealth DC against Perception, minimum rank if any), how to disable it (Thievery or another skill, DC, minimum rank), the trigger and the effect (save and DC), whether it resets |
| Social obstacle (crowd, gatekeeper) | the NPC's attitude, its Will DC, what Make an Impression or Request can achieve, what makes it worse |
| Challenge (multi-step) | **victory points:** the group must collect N points (3-5 per stage) through checks before something happens; success gives 1, critical success 2, critical failure may cost 1 or add a complication |

## Hazards
- **Simple hazard:** triggers once and resolves at once (a pit, a dart, a falling block).
- **Complex hazard:** acts every round once triggered, until disabled or escaped (a flooding room, a grinding wall).
- Use the level-based DC of the hazard's level for its Stealth DC, its disable DC and its save DC; make them hard (+2) for a nasty hazard.
- Noticing some hazards needs a minimum proficiency rank in Perception; disabling one may need a minimum rank in the skill.
- For damage, use the GM's book (creature and hazard damage by level); you may name a damage type and say "moderate for its level".

## Cost menu (for near misses)
A GM technique: on a failure that isn't a critical failure, offer progress at a cost. 2-3 costs that fit the fiction: lose time, make noise, break a tool, become fatigued, owe someone, give away your position. Costs may feed a hook's `## Surface` but never expose its `## 🔒 Hidden Truth`.

## Output block
```
**2E**
- Circumstance: "<…>" — <rule effect>
- Suggested DC: <n> (<simple: trained … | level n, hard …>) — <why>
- Fitting checks: <Skill> — <how>; <Skill> — <how>
- Degrees: crit success <…> · failure <…> · crit failure <…>
- (If a hazard) Notice: Perception vs Stealth DC <n> · Disable: <Skill> DC <n> (<rank>) · Effect: <save> DC <n>, <effect>
- Cost options: <…>; <…>
```
