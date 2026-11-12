# 2E: Scene Tags = Circumstances

This system has no aspects. The scene tag is a **circumstance**: something true in the fiction that changes the numbers or the options of the scene. It always has a concrete rule effect, and it should be usable by the PCs *and* able to hinder them.

## Kinds and their rule effects
| Kind | Typical effect |
|---|---|
| Circumstance bonus / penalty | a position, a plan, a distraction or bad footing: +1 or +2 circumstance bonus, or −1 or −2 circumstance penalty, to a named kind of check |
| Cover | **lesser cover** (a creature in the way): +1 circumstance bonus to AC. **Standard cover** (a low wall, a crate): +2 to AC, to Reflex saves against area effects and to Stealth. **Greater cover** (an arrow slit): +4 to the same. |
| Concealment | fog, dim light, smoke: the target is **concealed**; attacks against it need a DC 5 flat check or miss. |
| Terrain | **difficult terrain** (rubble, mud, a crowd): each square costs 5 extra feet of movement. **Greater difficult terrain**: 10 extra feet. Hazardous terrain deals damage when entered. |
| Light | darkness or dim light hides creatures from those who can't see in it; they can become concealed, hidden or undetected. |
| Hazard | something that forces a save or a check: falling rocks, fire, poison gas. Give the save, the DC and the effect (see `obstacles.md`). |
| Condition | a state imposed on a creature (see below), usually by a failed save. |

## Conditions (summary)
| Condition | Main effect |
|---|---|
| Off-guard | −2 circumstance penalty to AC |
| Frightened X | status penalty X to all checks and DCs; drops by 1 at the end of each of its turns |
| Sickened X | status penalty X to all checks and DCs; can't willingly swallow anything; can retch to lower it |
| Clumsy X / Enfeebled X / Drained X / Stupefied X | status penalty X to Dexterity / Strength / Constitution / mental (Intelligence, Wisdom, Charisma) checks and DCs; drained also lowers Hit Points; stupefied can disrupt spells |
| Slowed X / Stunned X | loses X actions (stunned also ends after the lost actions) |
| Quickened | gains 1 extra (limited) action each turn |
| Blinded / Dazzled / Deafened | can't see (everything is hidden) / everything is concealed / can't hear |
| Concealed / Hidden / Undetected | DC 5 flat check to hit / DC 11 flat check and must guess the square / location unknown |
| Prone | off-guard, −2 circumstance penalty to attacks; standing up costs an action |
| Grabbed / Restrained / Immobilized | off-guard and can't move / off-guard, can't move and very limited actions / can't move |
| Fatigued | −1 status penalty to AC and saves; can't keep up exploration activities |
| Fascinated / Fleeing / Confused / Controlled | −2 to Perception and skill checks, focused on the source / must move away / acts at random / another decides its actions |
| Paralyzed / Petrified / Unconscious | can't act and is off-guard (unconscious also blinded and prone) / turned to stone |
| Persistent damage | damage at the end of each turn until a DC 15 flat check ends it |

## A good circumstance
1. **Concrete and visible:** "Collapsed Scaffolding" beats "messy area". Name it in the world's voice and language.
2. **One clear rule effect:** pick from the tables above; don't invent new mechanics (unless Homebrew says so).
3. **Two-sided:** one line on how the PCs can use it, one on how it can work against them.
4. **True in the fiction:** it follows from what happened, obeys `## Immutable Rules`, and reveals no `## 🔒 Hidden Truth`.

## Output block
```
**"<Circumstance name>"** — <kind: bonus / penalty / cover / concealment / terrain / light / hazard / condition>
- Effect: <the rule effect, e.g. standard cover: +2 to AC, Reflex vs area and Stealth>
- Use it: <how a PC can exploit it, one line>
- Against them: <how it can hinder the PCs, one line>
```
