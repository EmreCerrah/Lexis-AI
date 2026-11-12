# d100: Obstacles

## Choosing difficulty (suggest, never decide)
- First the **difficulty level**: Automatic (no roll), Easy (×2), Average, Difficult (½), Impossible.
- Then **situational modifiers**: ±20% for a clear advantage or drawback, ±50% for an overwhelming one (`scene-tags.md`).
- Rule of thumb: if failure is not interesting, don't call for a roll, and say so. Suggest one difficulty level and at most one or two modifiers.

## Success levels, every time
For every obstacle, give a short result for each level that matters: critical (the best result), special (better than average), success, failure, fumble (a disaster: the opposite of the goal, a setback, harm).

## Passive vs active
- **Passive:** a fixed difficulty (a wall, a lock), or a **resistance roll** against an object's or force's rating (a door's STR, a poison's Potency, a storm's strength).
- **Active:** someone opposes: an **opposed roll** (Hide against Spot, Fast Talk against Insight).

## Obstacle types
| Type | What to produce |
|---|---|
| Barrier (lock, wall, chasm) | 2-3 fitting skills (or a characteristic roll) with the fiction for each, a difficulty level, modifiers, what a special and a fumble mean |
| Force (stuck door, current, weight) | a resistance roll: the object's rating and the PC characteristic it opposes (e.g. door STR 14 vs the PC's STR) |
| Puzzle / clue | the clue chain (fiction), what each skill reveals (Spot finds, Research digs out, Knowledge recalls, Idea roll connects), a difficulty; solvable without dice too |
| Poison / disease | its Potency against CON on the resistance table; effect on failure (damage, characteristic loss, a condition) and often a lesser effect on success |
| Hazard (fire, fall, collapse) | the avoiding roll (Dodge, Jump, Agility, Luck) and its difficulty; damage in dice (take falling and fire damage from the GM's book) |
| Horror (optional Sanity) | a SAN roll and a loss written success/failure (see below) |
| Social obstacle (crowd, gatekeeper) | the opposing skill (Insight, Persuade, Status), the NPC's stance, which Communication skill fits, what makes it Difficult |

## Sanity losses (optional rule)
The loss is written **success / failure** of the SAN roll. Guideline values:
| Source | Loss |
|---|---|
| A sudden scare, something grotesque | 0/1 |
| Finding a dead body, a scene of violent death | 0/1D3 |
| A minor unnatural creature | 0/1D3 or more |
| A mutilated body | 1/1D4 |
| Witnessing a gruesome death | 1/1D6 |
| Suffering minor torture | 1D3/1D6 |
| A freakish, horrible monster | 1D2/1D6 or more |
| A horrible transformation of oneself | 1D6/2D10 |
| A monstrous god-like entity | 1D10/1D100 |

Losing as much as the temporary insanity score (TIS) within five minutes means temporary insanity (the GM rolls its form and duration from their book). Repeated exposure to the same source should cost less.

## Cost menu (for near misses)
A GM technique: on a failure that isn't a fumble, offer progress at a cost. 2-3 costs that fit the fiction: lose time, make noise, break a tool, spend power points, take a minor wound, owe someone. Costs may feed a hook's `## Surface` but never expose its `## 🔒 Hidden Truth`.

## Output block
```
**d100**
- Tag: "<…>" — <modifier>
- Difficulty: <Easy / Average / Difficult> — <why>
- Fitting skills: <Skill> — <how>; <Skill> — <how> (or <Characteristic> × 5)
- Levels: special/critical <…> · fumble <…>
- (If force/poison) Resistance: <active rating> vs <passive characteristic> → <chance>%
- (If horror) SAN loss: <x/y>
- Cost options: <…>; <…>
```
