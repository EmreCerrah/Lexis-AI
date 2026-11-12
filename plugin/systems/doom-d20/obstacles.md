# Doom d20: Obstacles

## Choosing a DR (suggest, never decide)
- Default **DR 12**. Easier: 10, 8, 6. Harder: 14, 16, 18.
- Rule of thumb: if failure is not interesting, don't call for a test, and say so. The game is lethal: tell the GM when a failure could kill, so the players can choose another way.

## Obstacle types
| Type | What to produce |
|---|---|
| Barrier (lock, wall, chasm) | 1-2 fitting abilities with the fiction for each, a DR, what failure costs (time, noise, damage) |
| Trap | how it is noticed (Presence test or careful searching), how it is avoided or disarmed (a test and DR), the damage or effect if it goes off |
| Hazard (fire, cold, poison, fall) | the test (usually Toughness or Agility) and DR, the damage on failure (d4 small, d6-d8 serious, d10+ deadly) |
| Puzzle / riddle | the clue chain (fiction); solved by thinking, a Presence test only as a last hint |
| Social obstacle (crowd, gatekeeper) | a reaction roll if the NPC's attitude is unknown, then a Presence test and DR for anything they are reluctant to do |
| Fight | the foes' damage, hit points and armor, the DR to defend against them, their morale |

## Reactions (2d6, when an NPC's attitude is unknown)
| 2d6 | Attitude |
|---|---|
| 2-3 | kill! |
| 4-6 | angered |
| 7-8 | indifferent |
| 9-10 | almost friendly |
| 11-12 | helpful |

## Morale (2d6)
When a group loses its leader or half its number, or a lone foe is badly hurt, the GM rolls 2d6: above the foe's morale, it flees or surrenders.

## Cost menu (for near misses)
A GM technique: on a failure, offer progress at a grim cost. 2-3 costs that fit the fiction: lose time, make noise, break gear, take damage, lose an Omen, owe someone. Costs may feed a hook's `## Surface` but never expose its `## 🔒 Hidden Truth`.

## Output block
```
**d20**
- Tag: "<…>" — <DR shift or hazard>
- Test: <Ability> — <how>; <Ability> — <how>
- Suggested DR: <n> — <why>
- (If trap/hazard) Notice: Presence DR <n> · Avoid: <Ability> DR <n> · Damage: <dice>
- (If foes) HP <n> · Armor <tier> · Damage <dice> · Defend DR <n> · Morale <n>
- Cost options: <…>; <…>
```
