# Worldbuilding Forge

Masa başında çalışan bir **GM asistanı**. Oyunu sen yönetirsin. Agent zar atmaz ve sonuca karar vermez. Bildirdiğin olayları kaydeder, dünyadaki yankılarını gösterir ve dünya dosyalarını yalnızca seans sonunda senin onayınla günceller.

- Sistem: **Fate Accelerated** (MVP)
- Format: düz Markdown + `[[wiki-link]]`. Obsidian ile açılabilir ama gerekmez.
- Dünya verin senin klasöründe durur. Eklentinin içinde dünya verisi yoktur.

## Komutlar
| Komut | Ne yapar | Ne yazar |
|---|---|---|
| `/world-init` | Dünyayı kademeli kurar: zorunlu çekirdek, ertelenebilir kısımlar, agent'ın taslak çıkarabileceği kısımlar, FAE ayarları | tüm dünya dosyaları (yalnızca kurulumda) |
| `/event` | Masada olanı bildirirsin. Olay hemen canon olur; 3-5 yankı ve "kim bilebilir" listesi gelir | yalnızca seans günlüğü |
| `/ripples` | Kayıtlı bir olayı lens'lerle derinleştirir, bilginin 2-3 adım nasıl yayıldığını gösterir | yalnızca seans günlüğü (hipotez) |
| `/create-aspect` | Sahneye tek bir aspect yazar; invoke ve compel örneğiyle | günlük (öneri) |
| `/create-obstacle` | Engel, bulmaca ya da tehlike üretir: önce kurgu, sonra FAE zorluğu ve approach'lar | günlük (öneri) |
| `/create-lore` | Kalıcı bir NPC, mekan, grup ya da hook üretir | `proposals/` (öneri) |
| `/help-skills` | "Bu hangi approach'la yapılır?" sorusunu FAE diliyle cevaplar | hiçbir şey |
| `/close-session` | Seansta birikenleri tek numaralı listede onayına sunar ve yalnızca onayladıklarını dünyaya yazar | canon dosyalar |

**Codex'te:** Codex skill'leri `/` ile çağırmaz. Aynı isimleri `$` ile yaz (`$event Gabriel mezarlık bekçisini korkuttu`). İstersen `@worldbuilding-forge` diyerek eklentiyi seçebilir ya da isteğini doğal dille yazabilirsin.

## Akış
```
/world-init  →  (masa) /event … /ripples … /create-* … /help-skills  →  /close-session
```
- **Masa modu** (varsayılan): kısa maddeler gelir, seans ortasında onay sorulmaz.
- **Hazırlık modu:** komuta `hazırlık` ya da `detaylı` eklersen daha derin analiz alırsın.

## Bilgi modeli
| Tür | Nerede | Nasıl değişir |
|---|---|---|
| Canon | dosya gövdesi | yalnızca `/close-session` onayıyla |
| Hipotez | `## Belirsiz` ya da günlükteki `## Yankılar` | seans sonunda canon'a yükselir, belirsizde kalır ya da reddedilir |
| Çıkarım | yalnızca cevapta | hiç kaydedilmez |
| Öneri | `proposals/` ya da günlükteki `## Üretilenler` | seans sonunda onaylanır ya da reddedilir |

"Kim bilebilir" listesi her dosyadaki `## Bağlantılar` bölümünden çıkarılır (`- [[slug]] — ilişki`). Düz yazının içine gömülen bağlantıları agent görmez. Bağlantıyı bu bölüme yaz.

## Kurallar
- `world.md` içindeki **Değişmez Kurallar** (K1, K2…) agent'ın ürettiği hiçbir şeyde bozulmaz. Senin bildirdiğin bir olay bir kuralı bozarsa agent bunu engellemez. Olayı `⚠ kural çatışmalı` olarak işaretler ve üç seçenek sunar: (a) görünüşte ihlal, (b) bedelli istisna, (c) kural değişikliği.
- Hook'ların **🔒 Gizli Gerçek** bölümü değişmez. Oyuncular onu bozmaya çalışırsa agent, dünyanın içinden "neden başaramazlar" gerekçeleri önerir.
- `🔒` işaretli satırları oyunculara okuma.

## Kurulum
Lexis-AI reposundaki `lexis-setup` aracını çalıştır. Araç sana hangi yapay zekâ aracını kullandığını (Claude Code, Codex, Gemini CLI, GitHub Copilot) ve kurulumu nereye istediğini sorar, sonra eklentiyi o aracın biçimine göre kurar. Ayrıntılar repo kökündeki README'de.

Codex kullanıyorsan dünya klasörünü **workspace-write** (dosya yazma izni olan) modda aç. Yoksa seans ortasında her yazma için izin sorulur.

## Paket yapısı
```
plugin.json                 # taşınabilir manifest (agent-plugins.org şeması)
.codex-plugin/plugin.json   # Codex uyumluluk manifesti
AGENTS.md                   # tüm skill'lerin önce okuduğu çekirdek kurallar
skills/<komut>/SKILL.md     # 8 MVP skill (platformdan bağımsız)
references/                 # bilgi modeli, klasör yapısı, lens'ler (skill değil, gizli referans)
templates/                  # dünya çekirdeği şablonları
systems/fate-accelerated/   # FAE sistem paketi
```
Platforma özgü olan yalnızca iki manifest dosyası. Başka bir platforma geçerken yalnızca onları değiştirmen yeterli.
