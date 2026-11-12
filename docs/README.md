# Lexis-AI Documentation

Lexis-AI is a GM assistant for the live tabletop RPG table. It runs inside the AI coding tool you already use (Claude Code, Codex, Gemini CLI or GitHub Copilot) and keeps your campaign world as plain Markdown files in a folder you own.

| Guide | Read it when |
|---|---|
| [Installation](installation.md) | you set Lexis-AI up for the first time, update it, or install a test version |
| [Getting started](getting-started.md) | you want to run your first session, step by step |
| [Commands](commands.md) | you want to know exactly what each command does, writes and shows |
| [Game systems](game-systems.md) | you want to know which rules systems are supported and how each one answers |
| [Your world files](world-files.md) | you want to understand, edit, back up or move a world |
| [Troubleshooting](troubleshooting.md) | something does not work as expected |
| [Creating a system pack](creating-a-system-pack.md) | you want to add a new game system |

Design notes for contributors: [design/system-packs.md](design/system-packs.md). Planned work: [ROADMAP.md](../ROADMAP.md).

## The idea in one paragraph
You run the game; Lexis-AI never rolls dice and never decides whether an action succeeds. When something happens at the table, you tell it in one sentence (`/event …`). It records the event, shows 3-5 possible consequences ("ripples") and who in the world might hear about it. When you need rules help or scene material, it answers in your game system's own terms. At the end of the session (`/close-session`) it gives you one numbered list of everything that should change in your world files, and writes only what you approve.
