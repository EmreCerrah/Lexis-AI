# Game Systems

Lexis-AI keeps your world (events, ripples, links, hooks, canon) the same in every game. The **rules help** comes from a *system pack*: a small set of files that explains one game's mechanics. You choose the packs when you install ([Installation](installation.md)), and each world uses one of them (the `system:` field in `world.md`).

## Supported systems
| Pack (menu name) | `--system` id | Based on (the open source it restates) | The roll | Scene tag | Meta-currency | License |
|---|---|---|---|---|---|---|
| 5E SRD (fifth edition compatible) | `srd5` | System Reference Document 5.2.1 by Wizards of the Coast: the free 2024 fifth-edition core rules | d20 + modifier vs DC; Advantage/Disadvantage | scene feature | Heroic Inspiration | CC-BY-4.0 (SRD 5.2.1) |
| ORC 2E Core (three-action d20, remaster rules) | `orc-d20` | Pathfinder Player Core and GM Core (Paizo, remaster): their ORC-licensed rules | d20 + modifier vs DC, four degrees of success | circumstance | Hero Points | ORC License |
| Fate Accelerated | `fate-accelerated` | Fate Accelerated Edition (Evil Hat) | 4dF + approach vs the ladder | aspect | fate points | CC-BY 3.0 |
| Fate Core | `fate-core` | Fate Core System (Evil Hat) | 4dF + skill vs the ladder | aspect | fate points | CC-BY 3.0 |
| d100 Percentile Core (ORC) | `d100` | Basic Roleplaying: Universal Game Engine (Chaosium): the percentile engine behind classic investigative horror | d100 ≤ skill %, five success levels | situational modifier | power points as fate points (optional) | ORC License (Chaosium's Basic Roleplaying engine) |
| Forged in the Dark Core | `fitd` | Blades in the Dark SRD (One Seven Design) and Forged in the Dark games | d6 pool, highest die; position and effect | factor or clock | stress | CC-BY 3.0 (Blades in the Dark SRD) |
| Doom d20 Core (OSR, four abilities, Omens) | `doom-d20` | MÖRK BORG (compatible, under its Third Party License) | d20 + ability vs DR | DR shift or hazard | Omens | MÖRK BORG Third Party License |
| Superhero d20 Effects Core (OGL) | `hero-d20` | the effects-based superhero d20 SRD, third edition (OGL)¹ | d20 + rank vs DC, degrees of success | circumstance or condition | Victory Points | OGL 1.0a |

**Why the generic names?** Most game names are trademarks that the open licenses do not cover. The packs are named after their mechanics, and their `pack.md` files carry the exact license and attribution text. The *Based on* column names each source the way its license attribution does, so you can tell which game a pack is for. All trademarks belong to their owners; no pack is affiliated with or endorsed by them.

¹ The source's product name is Product Identity under the OGL, which forbids using it to claim compatibility, so it is not printed here. The pack follows the third-edition superhero d20 rules published as Open Game Content (ranks, effects, degrees of success, Toughness checks).

## What a pack contains, and what it never contains
**Contains:** how rolls work, difficulty scales, character math (modifiers, hit points, derived values), conditions and harm, how to set up an obstacle, the setup questions for `/world-init`, and the system sections of character and place files.

**Never contains:** any published setting (worlds, cities, gods, factions, famous characters, creatures), adventures or story, spell/power/item/monster lists, art, or rule text from games that are not open-licensed. Rules are written as short summaries in our own words. When something depends on your rulebook (a spell, a power, a monster), Lexis-AI names it and tells you to check your book.

The same rule applies to your world: Lexis-AI never brings published setting content into it on its own.

## How each system answers
The same request in different systems: *"Gabriel climbs the crypt wall in the rain."*

**5E SRD**
```
🎲 Ability check
- Check: Strength (Athletics) — hauling himself up · Dexterity (Acrobatics) — quick and light
- Suggested DC: 15 (Medium) — wet stone, but plenty of cracks
- Advantage/Disadvantage: Disadvantage for the rain unless he uses a rope
```

**ORC 2E Core**
```
🎲 Athletics: Climb
- Check: Athletics — the only real choice for climbing
- Suggested DC: 18 (level 3, hard) — slick stone
- Crit success: climbs fast and quiet · Crit failure: falls and takes damage
- Circumstance: −2 circumstance penalty for the rain
```

**Fate Accelerated**
```
🎲 Overcome
- Approach: Careful — testing each handhold · Quick — before the guards turn
- Suggested difficulty: Fair–Good (+2/+3) — "Rain-Slick Stones" works against him
```

**Fate Core**
```
🎲 Overcome
- Skill: Athletics — climbing · Physique — brute force on the drainpipe
- Suggested difficulty: Good (+3) — the wall is high and wet
```

**d100 Percentile Core**
```
🎲 Skill roll: Climb
- Skill: Climb — … · Agility roll (DEX × 5) if he has no Climb skill
- Difficulty: Average — there are handholds
- Modifiers: −20% rain-slick stone
- Special/critical: no noise at all · Fumble: falls from halfway up
```

**Forged in the Dark Core**
```
🎲 Prowl
- Action: Prowl — up the wall in the dark · Finesse — using the drainpipe
- Position: Risky — the patrol passes every few minutes · Effect: Standard
- 4-5: he makes it, but the patrol clock ticks · 1-3: he slips, harm level 1
```

**Doom d20 Core**
```
🎲 Agility test
- Ability: Agility — light and careful · Strength — pulling himself up
- Suggested DR: 14 (difficult) — rain
- Fumble: falls, d6 damage
```

**Superhero d20 Effects Core**
```
🎲 Athletics check
- Trait: Athletics — climbing · Acrobatics — wall-running
- Suggested DC: 15 (tough) — wet stone; routine for a wall-crawler
```

## Per-system notes

### 5E SRD (`srd5`)
- d20 tests: ability checks, saving throws, attack rolls; DC 5-30; Advantage/Disadvantage.
- Scene tags are scene features: light and obscurement, cover, Difficult Terrain, hazards, the 15 conditions.
- Obstacles include traps scaled by level, fear and mental stress, social checks with attitudes.
- Characters: ability scores and modifiers, Proficiency Bonus, hit points by class Hit Die, AC, spell save DC; NPCs get a one-line stat line.
- Setup asks: starting level, ability score method, advancement, character option limits, optional rules.

### ORC 2E Core (`orc-d20`)
- Checks with four degrees (critical success / success / failure / critical failure), proficiency ranks, simple and level-based DCs.
- Scene tags are circumstances: circumstance bonuses and penalties, cover, concealment, difficult terrain, valued conditions (frightened 2 …).
- Obstacles: hazards (simple and complex), victory points for multi-step challenges, social attitudes.
- Characters: attribute modifiers and boosts, proficiency, hit points, dying and wounded, Hero Points.
- The pack folder is under the ORC License, with the ORC, attribution and Reserved Material notices in its `pack.md`.

### Fate Accelerated (`fate-accelerated`) and Fate Core (`fate-core`)
- 4dF + approach (FAE) or + skill (Core) against the ladder; four actions; outcomes fail / tie / success / success with style.
- Scene tags are aspects, each with an invoke and a compel.
- Obstacles: difficulty ranges on the ladder, hazards as characters, challenges and contests (Core).
- Characters: aspects, approaches or the skill pyramid, stunts, stress and consequences.

### d100 Percentile Core (`d100`)
- d100 roll-under; critical (1/20), special (1/5), success, failure, fumble; Easy (×2) and Difficult (½); situational modifiers ±20% / ±50%; the resistance table.
- Optional Sanity with guideline losses (success/failure, e.g. 0/1D3); temporary insanity.
- Characters: characteristics, characteristic rolls (×5), hit points, major wounds, damage modifier, skill points by power level.
- The pack folder is under the ORC License, with the ORC and attribution notices in its `pack.md`.

### Forged in the Dark Core (`fitd`)
- d6 pools; 6 full success, 4-5 partial, 1-3 bad, two 6s critical; position (controlled / risky / desperate) and effect (limited / standard / great).
- Scene tags are factors (potency, scale, quality) and clocks (4/6/8 segments).
- Consequences, resistance with stress, harm levels, trauma, teamwork, flashbacks.
- Setup asks what the crew does and whether to use crew rules.

### Doom d20 Core (`doom-d20`)
- Four abilities (−3 to +3), tests against DR 6-18 (12 is normal), short lethal fights, armor tiers, Omens, reactions and morale.
- Rules are restated in our own words under the MÖRK BORG Third Party License; some details (the exact Omen uses, the broken table) are left to your book.

### Superhero d20 Effects Core (`hero-d20`)
- d20 + rank vs DC; degrees every 5 points; damage resisted with Toughness (DC 15 + rank) with the dazed / staggered / incapacitated track; extra effort; power level limits.
- The source's own names for its two point pools are Product Identity, so the pack calls them **Victory Points** and **Character Points**.
- The full Open Game License text and its Section 15 are in the pack's `pack.md`.

## Not supported
- **World of Darkness / Chronicles of Darkness:** the rules are not open-licensed (only a non-commercial, revocable fan license exists), so there is no pack.
- Any other system can be used with the closest pack plus your differences under `## Homebrew` in `system.md`. Homebrew always overrides the pack.

Want to add a system? See [Creating a system pack](creating-a-system-pack.md).
