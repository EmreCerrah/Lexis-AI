# {{World name}} — Lexis-AI world

This folder is a Lexis-AI world. For any agent working in this folder:

- Rules: the `AGENTS.md` of the Lexis-AI plugin applies. If the plugin is not installed, only read the world files.
- Canon files (`world.md`, `system.md`, `characters/`, `locations/`, `groups/`, `hooks/`) change only with `/close-session` approval.
- During a session, only append to the open session log (`sessions/`, `status: open`).
- Never violate `## Immutable Rules` or a hook's `## 🔒 Hidden Truth`.
- Never read content marked `🔒` to the players.
- Output language: {{language}}.
