# Lexis-AI

Masa başı GM asistanı **Lexis-AI** eklentisi ve onu farklı yapay zekâ araçlarına kuran `lexis-setup` aracı.

```
plugin/      # eklenti (platformdan bağımsız skill'ler, kurallar, şablonlar, FAE paketi)
installer/   # Go ile yazılmış kurulum aracı
.agents/plugins/marketplace.json   # Codex: repo doğrudan marketplace olarak eklenebilir
```

## Kurulum
```bash
lexis-setup
```
Araç iki soru sorar:
1. **Hangi yapay zekâ aracını kullanıyorsun?** Claude Code, Codex, Gemini CLI ya da GitHub Copilot.
2. **Kurulumu nereye istiyorsun?** Kullanıcı geneli ya da belirli bir proje klasörü.

Ardından eklentiyi GitHub'dan (`EmreCerrah/Lexis-AI`, `main` dalı) indirir, kurulum planını gösterir ve onay ister.

| Araç | Ne yapılır | Kullanıcı geneli | Proje |
|---|---|---|---|
| Claude Code | Yerel marketplace + `claude plugin install` | `~/.lexis-ai/claude-code` | `<proje>/.lexis-ai/claude-code` |
| Codex | Yerel marketplace + `codex plugin add` | `~/.lexis-ai/codex` | `<proje>/plugins/` + `<proje>/.agents/plugins/marketplace.json` |
| Gemini CLI | `gemini-extension.json` + `gemini extensions install` | `~/.gemini/extensions/lexis-ai` | (Gemini yalnızca kullanıcı geneli) |
| GitHub Copilot | Skill'ler + ortak dosyalar kopyalanır | `~/.copilot/skills` + `~/.copilot/lexis-ai` | `<proje>/.github/skills` + `<proje>/.github/lexis-ai` |

Skill'ler eklenti köküne `../../` ile başvurur. Kurulum aracı bu yolu her araç için uygun biçime çevirir:
- Claude Code: `${CLAUDE_PLUGIN_ROOT}/`
- Copilot, kullanıcı geneli: tam yol
- Copilot, proje: `../../lexis-ai/`
- Codex ve Gemini CLI: değiştirilmez

### Bayraklar
| Bayrak | Anlamı |
|---|---|
| `--platform claude\|codex\|gemini\|copilot` | Soruyu atlar |
| `--scope user\|project`, `--dir <klasör>` | Kurulum yeri |
| `--source <klasör>` | GitHub yerine yerel kopyadan kurar (geliştirme için) |
| `--repo owner/repo`, `--ref <dal/etiket>` | Başka bir repo ya da sürüm |
| `--dry-run` | Hiçbir şey yazmadan planı gösterir |
| `--yes` | Onay sormaz |
| `--force` | Lexis-AI'ye ait olmayan bir klasörün üzerine yazar |

Aracın yazdığı her klasörde bir `.lexis-ai.json` işareti bulunur. Kurulum tekrar çalıştırıldığında yalnızca bu işareti taşıyan klasörler değiştirilir. Eklentinin eski adı `worldbuilding-forge` idi. Bu adla yapılmış bir kurulum, `lexis-setup` yeniden çalıştırıldığında kaldırılır ve yerine `lexis-ai` kurulur.

Özel repo için `GITHUB_TOKEN` ortam değişkenini ayarla.

## Derleme
```bash
cd installer && go test ./... && go build -o lexis-setup.exe .
```
Sürüm etiketi ve repo, derlemede değiştirilebilir:
```bash
cd installer && go build -ldflags "-X main.version=v0.1.0 -X main.defaultRepo=EmreCerrah/Lexis-AI" -o lexis-setup.exe .
```
Diğer işletim sistemleri için `GOOS=linux GOARCH=amd64` ya da `GOOS=darwin GOARCH=arm64` ile derle.
