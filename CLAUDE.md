# stream-pac — CLAUDE.md

## O que é este projeto
App local para streamers. Agrega chat de múltiplas plataformas (Twitch, Kick, YouTube) em endpoints SSE,
serve overlays HTML estáticos para o OBS (Browser Source), e tem um gerador de overlays via AI (Claude API).
Origem: projeto Python `xumbra` em `../xumbra`.

## Arquitetura central

**O servidor é a fonte de verdade. O TUI é um cliente HTTP puro.**

```
Platforms (WS/polling) → Hub → EventBus → Server (HTTP + SSE)
                                              ↑
                              TUI (Bubble Tea) | curl | OBS Browser Source
```

- Qualquer interação do TUI pode ser feita via `curl localhost:7777/...`
- Não há lógica de negócio fora do servidor
- Adicionar nova plataforma = implementar `platform.Platform` + `Register()` no `init()`

## Estrutura de pastas

```
stream-pac/
├── main.go                          — wiring: config → crypto → store → bus → hub → server → tui
├── overlays/                        — HTML/CSS/JS estáticos servidos ao OBS
└── internal/
    ├── event/      — EventBus (pub/sub in-process)
    ├── platform/   — interface Platform + registry
    │   ├── twitch/
    │   ├── kick/
    │   └── youtube/
    ├── auth/       — Credentials (StaticToken, OAuthSession, BrowserSession) + Authorizer (OAuth PKCE)
    ├── crypto/     — AES-256-GCM + HKDF + OS keychain (zalando/go-keyring)
    ├── store/      — SQLite encriptado (modernc.org/sqlite, pure Go)
    ├── config/     — TOML config (BurntSushi/toml) + paths cross-platform
    ├── hub/        — fan-in de plataformas, publica eventos no bus
    ├── overlay/    — Manager (filesystem) + Watcher (hot-reload)
    ├── server/     — HTTP + SSE (stdlib net/http)
    └── tui/        — Bubble Tea root model + tabs (Dashboard, Overlays, Config, AI Gen, Logs)
        └── tabs/
```

## Dependências principais

| Pacote | Uso |
|---|---|
| `github.com/charmbracelet/bubbletea` | TUI framework (Elm architecture) |
| `github.com/charmbracelet/lipgloss` | Estilos e layout terminal |
| `github.com/BurntSushi/toml` | Config TOML |
| `modernc.org/sqlite` | SQLite pure Go (sem CGO) |
| `github.com/zalando/go-keyring` | OS keychain (Windows/macOS/Linux) |
| `golang.org/x/crypto` | HKDF para derivação de chave |

## Convenções

- **Sem implementação nos arquivos stub** — bodies retornam `panic("not implemented")`
- **Compile-time checks** — `var _ platform.Platform = (*Twitch)(nil)` em cada implementação
- **Testes** — arquivos `_test.go` com `t.Skip("not implemented")` até implementar
- **Credenciais** — duck typing, `platform.Credentials` não importa `auth`; ambos definem a mesma interface
- **Registro de plataformas** — cada pacote de plataforma chama `platform.Register()` no `init()`; `main.go` importa com `_`

## Persistência

- **Config** → TOML (`~/.config/stream-pac/config.toml`)
- **Sessões OAuth / cookies Kick** → SQLite encriptado (`~/.local/share/stream-pac/stream-pac.db`)
- **Overlays** → arquivos em `~/.local/share/stream-pac/overlays/{name}/`
- **Chave mestra** → OS keychain via `crypto.MasterKeyFromKeyring("stream-pac", "default")`

## API routes (referência rápida)

```
GET/POST /api/platforms[/{name}/connect|disconnect|status]
GET/POST/PUT/DELETE /api/overlays[/{name}[/reload]]
POST /api/overlays/generate
GET/PUT /api/config
GET /api/stats
GET /sse/chat | /sse/logs | /sse/events
GET /overlays/{name}/   ← arquivos estáticos para OBS
```

## Tema visual (Pac-Man arcade)

```go
ColorBackground = "#0A0E27"  // dark navy
ColorPrimary    = "#FFD700"  // pac-man yellow
ColorAccent     = "#1E90FF"  // maze blue
ColorText       = "#FFFFFF"
```

Font nos overlays HTML: `Press Start 2P` (Google Fonts)

## Kick — nota especial

Cloudflare bloqueia requests Go/Python diretos. O chatroom ID precisa ser obtido via browser.
A autenticação usa `auth.BrowserSession` (cookies capturados), não OAuth padrão.
Pusher key: `32cbd69e4b950bf97679`, cluster: `us2`, channel: `chatrooms.{id}.v2`

## Referências de código prontas (nos POCs em /mnt/c/Users/dougl/Dropbox/live/project)

- Twitch IRC WS, Kick Pusher WS, YouTube HTTP polling → `main.go` (~1500 linhas)
- SSE hub, file watcher → `streamhub/`
- Layout TUI 5 abas, AI Gen textarea → `streamhub-bubbletea/`
