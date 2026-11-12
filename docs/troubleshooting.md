# Troubleshooting

## Installing
**`lexis-setup` is not found after `go install`.**
`go install` prints nothing and puts the program in `%USERPROFILE%\go\bin` (Windows) or `~/go/bin` (macOS, Linux). Add that folder to your `PATH`, or run it with the full path.

**`go install …@latest` installs an old version.**
The Go module proxy can lag behind GitHub for a while. Install a specific commit instead: `go install github.com/EmreCerrah/Lexis-AI/lexis-setup@<commit>`.

**"… already exists and does not belong to Lexis-AI (use --force to overwrite)".**
`lexis-setup` only replaces folders that carry its `.lexis-ai.json` marker. A folder with the same name made by you or another tool is left alone. Rename it, or run with `--force` if you are sure it can be overwritten.

**Download fails.**
Check your internet connection. For a private fork, set `GITHUB_TOKEN`. For development, install from a local copy: `lexis-setup --source ./Lexis-AI`.

**Gemini CLI prints an error on exit, but the extension is installed.**
Some Gemini CLI versions crash on Windows when they exit, after finishing the work. `lexis-setup` checks the extension folder instead of the exit code; if it says the step succeeded, it did.

## Commands don't show up
- **Claude Code, user install:** commands have the plugin prefix: `/lexis-ai:event`. Run `claude plugin list` to check that `lexis-ai` is installed and enabled. Restart Claude Code after installing.
- **Claude Code, project install:** commands have no prefix (`/event`) and exist only when Claude Code is started in that project folder (or below it).
- **Codex:** call commands with `$event`, or pick `@lexis-ai`. Restart Codex after installing. For a project install, start Codex inside the project folder.
- **GitHub Copilot:** run `/skills reload`, then `/skills list`.
- **Gemini CLI:** there are no slash commands; describe what you want ("record an event: …", "close the session").
- After changing the install (another system, another language), restart the tool.

## In play
**"No world here" / it wants to run `/world-init`.**
Lexis-AI looks for `world.md` (with `type: world`) in the folder you started the tool in, and above it. Start the tool in your world folder.

**It answers in the wrong language.**
The world's `language:` in `world.md` wins; otherwise the default chosen at install. Change `language:` in `world.md`, or run `lexis-setup` again with another `--lang`.

**It uses the wrong game system, or says the system is not installed.**
The world's `system:` field in `world.md` names the pack (`srd5`, `fate-core` …). Only the systems you chose at install exist. Run `lexis-setup` again and choose that system (or "All systems").

**It changed a world file during the session.**
It should not: during play only `sessions/` and `proposals/` are written; world files change only at `/close-session`. If it happens, it is a bug; please report it with the command you used.

**"Who might know" misses someone obvious.**
It only follows `## Links`. Add the relation there: `- [[order-of-light]] — member`. Lexis-AI also proposes links it notices (`L1` …); approve them at `/close-session`.

**Answers are too long at the table.**
Table mode is the default and should stay within 5 short bullets. If you wrote `prep` or `detailed`, it switches to the long form. Say "short" or leave those words out.

**It tells me to check my book.**
System packs contain rules summaries only, never spell, power, item or monster lists. For those, your own rulebook is the source.

## Reporting a problem
Open an issue on GitHub with: the tool and its version, the install type (user or project), the game system, the command you ran, and what you expected.
