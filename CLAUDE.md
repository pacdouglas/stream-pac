# stream-pac — CLAUDE.md

## O que e este projeto
App local para streamers. Agrega chat de multiplas plataformas (Twitch, Kick, YouTube) em endpoints SSE,
serve overlays HTML estaticos para o OBS (Browser Source), e tem um gerador de overlays via AI (Claude API).
Origem: POC em `../poc`.

## Arquitetura central

**O servidor e a fonte de verdade. O TUI e um cliente HTTP puro.**

```
Platforms (WS/polling) -> Hub -> EventBus -> Server (HTTP + SSE)
                                                 ^
                                 TUI (Bubble Tea) | curl | OBS Browser Source
```

- Qualquer interacao do TUI pode ser feita via `curl localhost:7777/...`
- Nao ha logica de negocio fora do servidor
- Adicionar nova plataforma = implementar `platform.Platform` + `Register()` no `init()`

## Estrutura de pastas

```
stream-pac/
├── main.go                          — wiring: config -> crypto -> store -> bus -> hub -> server -> tui
├── overlays/                        — HTML/CSS/JS estaticos servidos ao OBS
└── internal/
    ├── event/      — EventBus (pub/sub in-process)
    ├── platform/   — interface Platform + registry
    │   ├── twitch/   — IRC over WebSocket (irc.chat.twitch.tv:6697)
    │   ├── kick/     — Pusher WebSocket (key: 32cbd69e4b950bf97679, cluster: us2)
    │   └── youtube/  — HTTP polling (youtubei/v1/live_chat, default 5s)
    ├── auth/       — Credentials (StaticToken, OAuthSession, BrowserSession) + Authorizer (OAuth PKCE)
    ├── crypto/     — AES-256-GCM + HKDF-SHA256 + OS keychain (zalando/go-keyring)
    ├── store/      — SQLite encriptado (modernc.org/sqlite, pure Go, sem CGO)
    ├── config/     — TOML config (BurntSushi/toml) + paths cross-platform
    ├── hub/        — fan-in de plataformas, publica eventos no bus
    ├── overlay/    — Manager (filesystem) + Watcher (hot-reload)
    ├── server/     — HTTP + SSE (stdlib net/http)
    └── tui/        — Bubble Tea root model + tabs
        ├── tabs/     — Dashboard, Overlays, Config, AI Gen, Logs (stubs antigos)
        └── fake/     — dados fake para desenvolvimento do TUI
```

## Dependencias principais

| Pacote | Uso |
|---|---|
| `github.com/charmbracelet/bubbletea` | TUI framework (Elm architecture) |
| `github.com/charmbracelet/bubbles` | Componentes TUI (viewport, textarea, etc) |
| `github.com/charmbracelet/lipgloss` | Estilos e layout terminal |
| `github.com/BurntSushi/toml` | Config TOML |
| `modernc.org/sqlite` | SQLite pure Go (sem CGO) |
| `github.com/zalando/go-keyring` | OS keychain (Windows/macOS/Linux) |
| `golang.org/x/crypto` | HKDF para derivacao de chave |

## Convencoes

- **Sem implementacao nos arquivos stub** — bodies retornam `panic("not implemented")`
- **Compile-time checks** — `var _ platform.Platform = (*Twitch)(nil)` em cada implementacao
- **Testes** — arquivos `_test.go` com `t.Skip("not implemented")` ate implementar
- **Credenciais** — duck typing, `platform.Credentials` nao importa `auth`; ambos definem a mesma interface
- **Registro de plataformas** — cada pacote de plataforma chama `platform.Register()` no `init()`; `main.go` importa com `_`

## Persistencia

- **Config** -> TOML (`~/.config/stream-pac/config.toml`)
- **Sessoes OAuth / cookies Kick** -> SQLite encriptado (`~/.local/share/stream-pac/stream-pac.db`)
- **Overlays** -> arquivos em `~/.local/share/stream-pac/overlays/{name}/`
- **Chave mestra** -> OS keychain via `crypto.MasterKeyFromKeyring("stream-pac", "default")`

## API routes (referencia rapida)

```
GET/POST /api/platforms[/{name}/connect|disconnect|status]
GET/POST/PUT/DELETE /api/overlays[/{name}[/reload]]
POST /api/overlays/generate
GET/PUT /api/config
GET /api/stats
GET /sse/chat | /sse/logs | /sse/events
GET /overlays/{name}/   <- arquivos estaticos para OBS
```

## TUI

- **5 abas**: [1] Dashboard, [2] Overlays, [3] Config, [4] AI Gen, [5] Logs
- **10 temas**: Dark (default arcade), Midnight, Dracula, Nord, Tokyo, Gruvbox, Catppuccin, Solarized, Monokai, Matrix
- **i18n**: PT e EN completos (100+ strings)
- **Atalhos**: [1-5] abas, [t] tema, [l] idioma, [?] help, [ctrl+q] sair
- **Tema visual Pac-Man arcade**: fundo `#0A0E27`, amarelo `#FFD700`, azul `#1E90FF`
- **Font overlays HTML**: `Press Start 2P` (Google Fonts)

## Kick — nota especial

Cloudflare bloqueia requests Go/Python diretos. O chatroom ID precisa ser obtido via browser.
A autenticacao usa `auth.BrowserSession` (cookies capturados), nao OAuth padrao.
Pusher key: `32cbd69e4b950bf97679`, cluster: `us2`, channel: `chatrooms.{id}.v2`

## Estado atual

A POC (`../poc`) tem toda a arquitetura e interfaces desenhadas. A maioria dos bodies sao stubs
(`panic("not implemented")`). O TUI ja tem visual completo com dados fake (~1000 mensagens,
5 eventos, 2 overlays, 13 log entries). Os testes existem mas com `t.Skip`.

## Sobre o usuario

Este e um projeto pessoal de aprendizado. O usuario quer ser guiado no desenvolvimento,
nao quer que o Claude programe tudo automaticamente. O objetivo e aprender.
