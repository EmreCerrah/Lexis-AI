# 5E: Scene Tags = Scene Features

5E has no aspects. The scene tag is a **scene feature**: something true in the fiction that changes the rules of the scene. It always has a concrete rule effect, and it should be usable by the PCs *and* able to hinder them, like a Fate aspect.

## Kinds and their rule effects
| Kind | Typical effect |
|---|---|
| Light | **Dim light** (twilight, a weak lantern, fog edge) makes an area Lightly Obscured: Disadvantage on Wisdom (Perception) checks that rely on sight. **Darkness** makes it Heavily Obscured: a creature trying to see into it is effectively Blinded. |
| Obscurement | Fog, foliage, heavy rain: Lightly Obscured (patchy) or Heavily Obscured (dense). |
| Cover | **Half cover** (a low wall, a crate): +2 to AC and Dexterity saves. **Three-quarters cover** (an arrow slit, a thick trunk): +5. **Total cover**: can't be targeted directly. |
| Terrain | **Difficult Terrain** (rubble, mud, deep snow, a crowd): every foot of movement costs 1 extra foot. Climbing and swimming also cost 1 extra foot per foot. |
| Advantage source | A position, a distraction or a prepared plan: the GM grants Advantage on a specific kind of roll while it lasts. A drawback works the other way (Disadvantage). |
| Hazard | Something that forces a saving throw: falling rocks, fire, poison gas, icy water. Give the save, the DC and the effect (see `obstacles.md`). |
| Condition | A state imposed on a creature (see below), usually by a failed save. |

## Conditions (summary)
| Condition | Main effect |
|---|---|
| Blinded | can't see; fails sight-based checks; attacks against it have Advantage, its attacks have Disadvantage |
| Charmed | can't harm the charmer; the charmer has Advantage on social checks against it |
| Deafened | can't hear; fails hearing-based checks |
| Exhaustion | cumulative levels: each level is −2 to every D20 Test and −5 ft Speed; level 6 is death; a Long Rest removes one level |
| Frightened | Disadvantage on checks and attacks while the source is in sight; can't move closer to it |
| Grappled | Speed 0; Disadvantage on attacks against anyone but the grappler; can be dragged |
| Incapacitated | no actions, Bonus Actions or Reactions; can't speak; Concentration breaks |
| Invisible | Advantage on Initiative; attacks against it have Disadvantage, its attacks have Advantage, unless it can be seen |
| Paralyzed | Incapacitated, Speed 0, fails Strength and Dexterity saves; attacks against it have Advantage; hits from within 5 ft are Critical Hits |
| Petrified | turned to stone: Incapacitated, Speed 0, fails Strength and Dexterity saves, Resistance to all damage |
| Poisoned | Disadvantage on attack rolls and ability checks |
| Prone | can only crawl or spend half its Speed to stand; Disadvantage on its attacks; melee attacks against it have Advantage, ranged ones Disadvantage |
| Restrained | Speed 0; attacks against it have Advantage, its attacks have Disadvantage; Disadvantage on Dexterity saves |
| Stunned | Incapacitated, fails Strength and Dexterity saves; attacks against it have Advantage |
| Unconscious | Incapacitated and Prone, drops what it holds, fails Strength and Dexterity saves; attacks against it have Advantage; hits from within 5 ft are Critical Hits |

A condition doesn't stack with itself (Exhaustion is the exception).

## A good scene feature
1. **Concrete and visible:** "Fallen Chandelier Chains" beats "messy room". Name it in the world's voice and language.
2. **One clear rule effect:** pick from the table above; don't invent new mechanics (unless Homebrew says so).
3. **Two-sided:** write one line on how the PCs can use it and one on how it can work against them.
4. **True in the fiction:** it follows from what happened, obeys `## Immutable Rules`, and reveals no `## 🔒 Hidden Truth`.

## Output block
```
**"<Feature name>"** — <kind: light / cover / terrain / advantage source / hazard / condition>
- Effect: <the rule effect, e.g. Lightly Obscured: Disadvantage on sight-based Perception>
- Use it: <how a PC can exploit it, one line>
- Against them: <how it can hinder the PCs, one line>
```
