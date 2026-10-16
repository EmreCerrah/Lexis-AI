# Lens Index

Lenses are analysis perspectives. They are **not skills**: the GM never calls them. `/event` and `/ripples` read this index, pick lenses, then open only the lens files they picked.

## How to pick
1. Read the event: what changed, where, who was involved, which hooks it touches.
2. Score each lens below by its triggers. Pick the highest.
   - `/event` (table): **max 2**, 3 only if the event clearly touches three areas.
   - `/ripples`: 2-3 in table mode; every lens that scores in prep mode.
3. If no lens fits well, use `toplum-soylenti` — every event leaves a rumor trail.
4. Name the picked lenses in the output header line, e.g. `Lens: güç-siyaset, inanç-doğaüstü`.

## Available lenses
| File | Lens | Pick when the event touches… |
|---|---|---|
| `guc-siyaset.md` | Güç ve Siyaset | rulers, factions, law, order, guards, debts of loyalty, who gains or loses authority |
| `toplum-soylenti.md` | Toplum ve Söylenti | public places, witnesses, fear, reputation, gossip channels, ordinary people's reactions |
| `inanc-dogaustu.md` | İnanç ve Doğaüstü | religion, cults, rituals, magic, monsters, curses, the uncanny, sacred places |

## Planned (not in MVP — do not invent their content)
ekonomi-ticaret, askeri-güvenlik, karakter-ilişki, coğrafya-ekoloji.

## Lens file contract
Every lens file has the same sections: **Odak**, **Tetikleyiciler**, **Sorular**, **Yayılma kanalları**, **Tipik yankılar**, **Tuzaklar**. A new lens must follow this contract and be added to the table above.
