# d20 Hero: Scene Tags = Circumstances and Conditions

This system has no aspects. The scene tag is a **circumstance**: something true in the fiction that gives a ±2 (minor) or ±5 (major) modifier to a named kind of check, or a **condition** it imposes. It should be usable by the heroes *and* able to hinder them.

## Kinds and their rule effects
| Kind | Typical effect |
|---|---|
| Circumstance | a vantage point, a distraction, bad footing, smoke: ±2 (minor) or ±5 (major) to a named kind of check |
| Concealment | darkness, fog, a smoke screen: harder to target or notice (a penalty on attacks or Perception) |
| Cover | a wall, a car, a pillar: a bonus to defenses against ranged attacks |
| Environment | a collapsing building, a storm, extreme heat: resistance checks (Fortitude, Dodge) against a DC, or a condition |
| Condition | a state imposed on someone (see below) |

## Conditions (summary)
| Condition | Effect |
|---|---|
| Impaired / Disabled | −2 / −5 on checks (or on a named kind of check) |
| Dazed | only free actions and one standard action per turn |
| Stunned | no actions at all |
| Hindered / Immobile | half speed / can't move |
| Vulnerable / Defenseless | active defenses halved / active defenses 0 (attacks against it can be routine or automatic critical hits) |
| Fatigued / Exhausted | hindered / impaired and hindered; an hour of rest recovers |
| Staggered | dazed and hindered |
| Incapacitated | defenseless, stunned and unaware; usually prone |
| Prone | −5 to close attacks; enemies +5 to close attacks and −5 to ranged attacks against it |
| Restrained / Bound | hindered and vulnerable (immobile if anchored) / defenseless, immobile and impaired |
| Compelled / Controlled | one standard action chosen by another / all actions chosen by another |
| Surprised | stunned and vulnerable |
| Dying | incapacitated and near death (see `entities.md`) |

## A good scene tag
1. **Concrete and big:** "Crumbling Overpass" beats "unstable ground". Name it in the world's voice and language.
2. **One clear effect:** a ±2/±5 modifier on a named check, or a condition with its resistance check.
3. **Two-sided:** one line on how the heroes can use it, one on how it can work against them.
4. **True in the fiction:** it follows from what happened, obeys `## Immutable Rules`, and reveals no `## 🔒 Hidden Truth`.

## Output block
```
**"<Tag name>"** — <kind: circumstance / concealment / cover / environment / condition>
- Effect: <e.g. +2 to Stealth checks; or Dodge DC 15 or hindered>
- Use it: <how a hero can exploit it, one line>
- Against them: <how it can hinder the heroes, one line>
```
