# Installation

Lexis-AI is installed with `lexis-setup`, a small command-line program. It asks a few questions, downloads the plugin from GitHub and puts it where your AI tool looks for it.

## 1. What you need
- One of the supported AI tools, installed and working:
  - **Claude Code** (the `claude` command)
  - **Codex** (the CLI or the Codex desktop app)
  - **Gemini CLI** (the `gemini` command)
  - **GitHub Copilot** (Copilot CLI, or Copilot in an editor that reads skills)
- **Go 1.22 or newer** to install `lexis-setup`: https://go.dev/dl/. You don't need to know Go; it is only used to download and build the setup tool.

Check Go with:
```bash
go version
```

## 2. Install the setup tool
```bash
go install github.com/EmreCerrah/Lexis-AI/lexis-setup@latest
```
`go install` prints nothing when it succeeds. The program goes into Go's bin folder:
- Windows: `%USERPROFILE%\go\bin\lexis-setup.exe`
- macOS and Linux: `~/go/bin/lexis-setup`

If your terminal says `lexis-setup` is not found, add that folder to your `PATH`, or run the program with its full path.

## 3. Run it
```bash
lexis-setup
```
It asks four questions. Press Enter to take the default shown in brackets.

1. **Which AI tool do you use?** Claude Code, Codex, Gemini CLI or GitHub Copilot.
2. **Where do you want to install it?**
   - *For my user:* available in every folder you open the tool in.
   - *In a specific project folder:* only in that folder (for example your campaign folder). Gemini CLI extensions are always user-wide.
3. **Which language should the AI answer in?** English, Türkçe, Español, Deutsch, Français, Português, or type any other language. This becomes the default; each world can still choose its own language.
4. **Which game system do you play?** Pick one system, or "All systems". **Only the systems you choose are installed;** the files of the others are never written. See [Game systems](game-systems.md).

It then downloads the plugin, shows the install plan and asks `Continue? [Y/n]`.

```
Install plan — lexis-ai 0.2.0 → Claude Code (project: C:\Campaigns\Karcsa, language: Türkçe (Turkish), game system: 5E SRD (fifth edition compatible))
  • Shared files (rules, templates, system packs) → C:\Campaigns\Karcsa\.claude\lexis-ai
  • 10 skills → C:\Campaigns\Karcsa\.claude\skills
```

## Where the files go
| Tool | User install | Project install |
|---|---|---|
| Claude Code | a plugin through a local marketplace in `~/.lexis-ai/claude-code`, registered with `claude plugin install` | `<project>/.claude/skills` + `<project>/.claude/lexis-ai` |
| Codex | a plugin through a local marketplace in `~/.lexis-ai/codex`, registered with `codex plugin add` | `<project>/.agents/skills` + `<project>/.agents/lexis-ai` |
| Gemini CLI | an extension in `~/.gemini/extensions/lexis-ai`, via `gemini extensions install` | `<project>/.gemini/skills` + `<project>/.gemini/lexis-ai` |
| GitHub Copilot | `~/.copilot/skills` + `~/.copilot/lexis-ai` | `<project>/.github/skills` + `<project>/.github/lexis-ai` |

A **project install** is a plain copy into the folder each tool already scans: the commands go to `skills/`, the shared files (rules, templates, the chosen system packs) to `lexis-ai/` next to it. Nothing is registered, so the commands are there as soon as you open the tool in that folder, and you can commit them together with your world.

## After installing
| Tool | Next step |
|---|---|
| Claude Code | start `claude` in your world folder. User install: `/lexis-ai:world-init`. Project install: `/world-init`. |
| Codex | start Codex in your world folder (in **workspace-write** mode, so it can write files without asking every time) and call `$world-init`. |
| Gemini CLI | start `gemini` in your world folder and say "set up a new world". |
| GitHub Copilot | in Copilot CLI run `/skills reload`, check with `/skills list`, then `/world-init`. |

Continue with [Getting started](getting-started.md).

## Options (flags)
Every question can be answered with a flag, which is useful for scripts or re-installs.

| Flag | Meaning |
|---|---|
| `--platform claude\|codex\|gemini\|copilot` | the AI tool |
| `--scope user\|project` | user-wide or project install |
| `--dir <folder>` | the project folder (with `--scope project`) |
| `--lang <language>` | the answer language, e.g. `--lang English`, `--lang Turkish` |
| `--system <id>` | the game system: `srd5`, `orc-d20`, `fate-accelerated`, `fate-core`, `d100`, `fitd`, `doom-d20`, `hero-d20`; several separated by commas; or `all` |
| `--dry-run` | show the plan, write nothing |
| `--yes` | don't ask for confirmation |
| `--force` | overwrite a folder that does not belong to Lexis-AI |
| `--ref <branch or tag>` | install another version of the plugin (default `main`) |
| `--repo owner/repo` | install from another repository (a fork) |
| `--source <folder>` | install from a local copy instead of GitHub (for development) |

Example: a project install for Copilot with 5E, in English, without questions:
```bash
lexis-setup --platform copilot --scope project --dir ./my-campaign --lang English --system srd5 --yes
```

## Updating
1. Update the setup tool: run the same `go install …@latest` command again.
2. Run `lexis-setup` again with the same answers. It replaces only the folders it created before (they carry a `.lexis-ai.json` marker) and leaves everything else alone. Your world files are never touched.

To **add or change a game system**, run `lexis-setup` again and choose the system you want, or "All systems". The installed set is replaced by your new choice.

## Installing a test version
To try a branch before it is merged (for example `feat/system-packs`):
```bash
git ls-remote https://github.com/EmreCerrah/Lexis-AI refs/heads/feat/system-packs
go install github.com/EmreCerrah/Lexis-AI/lexis-setup@<commit>
lexis-setup --ref feat/system-packs
```
The first line prints the branch's latest commit id. The second builds the setup tool from that commit: Go does not accept branch names that contain `/` as a version, so use the commit id (the first 7-12 characters are enough). `--ref` makes `lexis-setup` download the plugin from the same branch.

## Uninstalling
- **Project install:** delete `lexis-ai/` and the ten command folders (`world-init`, `event`, `ripples`, `create-aspect`, `create-obstacle`, `create-scene`, `create-lore`, `help-skills`, `close-session`, `lexis-help`) from the tool's `skills/` folder in your project.
- **Claude Code (user):** `claude plugin uninstall lexis-ai@lexis-ai`, then `claude plugin marketplace remove lexis-ai`, then delete `~/.lexis-ai/claude-code`.
- **Codex (user):** remove the `lexis-ai` plugin in Codex, then delete `~/.lexis-ai/codex`.
- **Gemini CLI:** `gemini extensions uninstall lexis-ai`, then delete `~/.lexis-ai/gemini`.
- **Copilot (user):** delete the ten command folders from `~/.copilot/skills` and the folder `~/.copilot/lexis-ai`.

Your world folders are separate and stay as they are.

## Older versions
The plugin used to be called `worldbuilding-forge`. Running `lexis-setup` removes an install under that name and installs `lexis-ai` instead. Older project installs in `<project>/.lexis-ai/` or `<project>/plugins/` are cleaned up the same way.

## Private forks
For a private repository, set the `GITHUB_TOKEN` environment variable before running `lexis-setup --repo you/your-fork`.
