# Lexis-AI

The **Lexis-AI** plugin, a GM assistant for the live tabletop RPG table, plus `lexis-setup`, the tool that installs it into different AI tools.

```
plugin/      # the plugin (platform-neutral skills, rules, templates, FAE pack)
installer/   # the installer, written in Go
.agents/plugins/marketplace.json   # Codex: the repo can be added directly as a marketplace
```

## Install
```bash
lexis-setup
```
The tool asks three questions:
1. **Which AI tool do you use?** Claude Code, Codex, Gemini CLI or GitHub Copilot.
2. **Where do you want to install it?** For your user, or in a specific project folder.
3. **Which language should the AI answer in?** English, Türkçe, Español, Deutsch, Français, Português, or any other language you type.

It then downloads the plugin from GitHub (`EmreCerrah/Lexis-AI`, branch `main`), shows the install plan and asks for confirmation.

| Tool | What happens | User install | Project install |
|---|---|---|---|
| Claude Code | local marketplace + `claude plugin install` | `~/.lexis-ai/claude-code` | `<project>/.lexis-ai/claude-code` |
| Codex | local marketplace + `codex plugin add` | `~/.lexis-ai/codex` | `<project>/plugins/` + `<project>/.agents/plugins/marketplace.json` |
| Gemini CLI | `gemini-extension.json` + `gemini extensions install` | `~/.gemini/extensions/lexis-ai` | (Gemini supports user installs only) |
| GitHub Copilot | skills and shared files are copied | `~/.copilot/skills` + `~/.copilot/lexis-ai` | `<project>/.github/skills` + `<project>/.github/lexis-ai` |

**Language:** the chosen language is written into the installed `AGENTS.md` as the default output language. `/world-init` offers it as the world's language and stores it in `world.md`, so each world can use its own. File names, headings and frontmatter keys stay in English so every skill can find them; only the content is in the chosen language.

The repo has a single manifest, `plugin/plugin.json`. The installer generates each tool's own manifest from it.

Skills refer to the plugin root as `../../`. The installer rewrites that path for each tool:
- Claude Code: `${CLAUDE_PLUGIN_ROOT}/`
- Copilot, user install: the absolute path
- Copilot, project install: `../../lexis-ai/`
- Codex and Gemini CLI: unchanged

### Flags
| Flag | Meaning |
|---|---|
| `--platform claude\|codex\|gemini\|copilot` | skips the tool question |
| `--scope user\|project`, `--dir <folder>` | where to install |
| `--lang <language>` | skips the language question (e.g. `--lang English`, `--lang Turkish`) |
| `--source <folder>` | installs from a local copy instead of GitHub (for development) |
| `--repo owner/repo`, `--ref <branch/tag>` | another repository or version |
| `--dry-run` | shows the plan without writing anything |
| `--yes` | doesn't ask for confirmation |
| `--force` | overwrites a folder that does not belong to Lexis-AI |

Every folder the tool writes carries a `.lexis-ai.json` marker. When the tool runs again, it only replaces folders that carry this marker. The plugin used to be called `worldbuilding-forge`; an install under that name is removed and replaced by `lexis-ai` when `lexis-setup` runs again.

For a private repository, set the `GITHUB_TOKEN` environment variable.

## Build
```bash
cd installer && go test ./... && go build -o lexis-setup.exe .
```
The version label and the repository can be set at build time:
```bash
cd installer && go build -ldflags "-X main.version=v0.1.0 -X main.defaultRepo=EmreCerrah/Lexis-AI" -o lexis-setup.exe .
```
For other operating systems, build with `GOOS=linux GOARCH=amd64` or `GOOS=darwin GOARCH=arm64`.
