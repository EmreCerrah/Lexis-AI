# Fate Core: Scene Tags = Aspects

In Fate the scene tag is the **aspect**: a short phrase that is **true in the fiction** and can be **invoked** (for a bonus) and **compelled** (for a complication).

## Kinds
| Kind | Lasts | Example |
|---|---|---|
| Character aspect (high concept, trouble, others) | until changed | *Last Hunter of the Order*, *The Blood Keeps Calling* |
| Situation / scene aspect | until the scene or situation ends | *Fog Swallows the Gravestones*, *Toppled Barrels* |
| Consequence | until healed | *Cracked Ribs* |
| Boost | one free invoke, then gone | *A Moment's Opening* |
| Location / group aspect (world files) | permanent | *Nobody Comes Here at Night* |

Write aspects in the output language; the examples above are only examples.

## A good aspect
1. **Double-edged:** it can help *and* hurt. Test: write one invoke and one compel. If you can't write both, rewrite it.
2. **Short and evocative:** 2-6 words, a picture, not a stat. Prefer concrete nouns and verbs over adjectives ("Rotten Boards Creak" beats "Dangerous Floor").
3. **True in the fiction:** it follows from what happened, it does not violate `## Immutable Rules`, and it does not reveal a `## 🔒 Hidden Truth`.
4. **In the world's voice:** it matches the tone and language of `world.md`.

## Free invokes
- Create an Advantage success → 1 free invoke (with style → 2). Tie → a boost.
- A newly discovered aspect on a hidden thing → note "whoever discovers it gets 1 free invoke", as a reminder only.

## Output block
```
**"<Aspect>"** — <kind: scene / situation / consequence / boost / permanent>
- Invoke: <who can use it, for what, one line>
- Compel: <how it complicates things, one line>
```
