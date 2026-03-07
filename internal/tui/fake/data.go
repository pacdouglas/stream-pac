// Package fake contém dados estáticos para desenvolvimento da UI.
// Substituídos por dados reais quando integrado ao hub SSE.
package fake

import "fmt"

// ChatMsg representa uma mensagem de chat de qualquer plataforma.
type ChatMsg struct {
	Platform string // "tw", "ki", "yt", "sys"
	User     string
	Color    string // hex, ex: "#9b72ff"
	Text     string
}

// Event representa um evento de plataforma (raid, sub, gift, etc.).
type Event struct {
	Platform string
	Type     string
	User     string
	Detail   string
	Age      string
}

// Overlay representa um overlay HTML disponível.
type Overlay struct {
	Name   string
	File   string
	Width  int
	Height int
	Active bool
}

// LogEntry representa uma entrada de log de um subsistema.
type LogEntry struct {
	Timestamp string
	Source    string
	Message   string
}

// Viewers agrupa o total de viewers por plataforma.
var Viewers = struct{ Twitch, Kick, YouTube int }{80, 45, 17}

// ChatMessages são as mensagens fake exibidas no chat do Dashboard.
// Geradas via init() para simular ~1000 mensagens e testar o scroll.
var ChatMessages []ChatMsg

func init() {
	type user struct {
		platform, name, color string
	}
	users := []user{
		{"tw", "Pac", "#9b72ff"},
		{"tw", "viewer_x", "#c9b5ff"},
		{"tw", "lurker99", "#9b72ff"},
		{"tw", "dev_amigo", "#c9b5ff"},
		{"tw", "xgamer50v", "#ffd700"},
		{"tw", "sub_novo", "#58a6ff"},
		{"tw", "raider_1", "#ff7b72"},
		{"tw", "twitch_fan", "#d2a8ff"},
		{"ki", "xumbr3ga", "#53fc18"},
		{"ki", "fan_top", "#53fc18"},
		{"ki", "supporter_1", "#3fb950"},
		{"ki", "kick_lurk", "#26a641"},
		{"ki", "br_viewer", "#39d353"},
		{"yt", "sub_yt_1", "#aaaaaa"},
		{"yt", "novo_viewer", "#aaaaaa"},
		{"yt", "yt_fan", "#aaaaaa"},
	}
	texts := []string{
		"salve galera!!",
		"chegando agora, perdi o início?",
		"o7",
		"VAMO!",
		"PogChamp PogChamp",
		"que lib estão usando pra isso?",
		"esse overlay é muito bom",
		"mandou bem na última live",
		"alguém mais tá no Kick e no Twitch ao mesmo tempo?",
		"Go + SSE? isso é muito foda",
		"primeira vez aqui, gostei",
		"bora bora!!",
		"kkkkkk verdade",
		"já seguí no Kick também",
		"esse chat unificado é uma mão na roda",
		"como faz pra ter isso no meu canal?",
		"f no chat",
		"5Head essa solução",
		"muito bem feito, parabéns",
		"quando lança o open source?",
		"viewer fiel aqui o7",
		"vim pelo raid, fiquei pelo conteúdo",
		"a qualidade do stream melhorou muito",
		"GG",
		"clip isso aí",
		"vai ter tutorial disso?",
		"primeiro stream que assisto com multichat funcionando",
		"sem lag nenhum, impressionante",
		"esse Go é realmente rápido né",
		"aqui do brasil também assistindo",
	}
	sysEvents := []string{
		"Viewers: Twitch 80 · Kick 45 · YouTube 17",
		"Raid de xgamer50v com 120 viewers!",
		"sub_novo se inscreveu — Tier 1",
		"generous_1 presenteou 5 subs",
		"Viewers: Twitch 95 · Kick 52 · YouTube 20",
		"fan_top renovou — 3 meses",
		"Viewers: Twitch 110 · Kick 60 · YouTube 25",
	}

	// cabeçalho
	ChatMessages = append(ChatMessages,
		ChatMsg{Platform: "sys", Text: "Twitch conectado — canal streampac"},
		ChatMsg{Platform: "sys", Text: "Kick conectado — canal streampac"},
		ChatMsg{Platform: "sys", Text: "YouTube conectado — canal streampac"},
	)

	sysIdx := 0
	for i := 0; i < 997; i++ {
		if i%50 == 0 {
			ChatMessages = append(ChatMessages, ChatMsg{
				Platform: "sys",
				Text:     sysEvents[sysIdx%len(sysEvents)],
			})
			sysIdx++
			continue
		}
		u := users[i%len(users)]
		ChatMessages = append(ChatMessages, ChatMsg{
			Platform: u.platform,
			User:     u.name,
			Color:    u.color,
			Text:     fmt.Sprintf("[%d] %s", i+1, texts[i%len(texts)]),
		})
	}
}

// Events são os eventos fake exibidos no painel de eventos do Dashboard.
var Events = []Event{
	{Platform: "tw", Type: "Raid", User: "xgamer50v", Detail: "120 viewers", Age: "agora"},
	{Platform: "tw", Type: "Sub", User: "sub_novo", Detail: "Tier 1", Age: "2 min"},
	{Platform: "ki", Type: "Resub", User: "fan_top", Detail: "3 meses", Age: "5 min"},
	{Platform: "tw", Type: "Gift", User: "generous_1", Detail: "5 subs", Age: "12 min"},
	{Platform: "yt", Type: "Member", User: "yt_fan", Detail: "1 mês", Age: "18 min"},
}

// Overlays são os overlays HTML disponíveis.
var Overlays = []Overlay{
	{Name: "Multi Chat", File: "streampac_multichat.html", Width: 400, Height: 800, Active: true},
	{Name: "Webcam Overlay", File: "streampac_overlay_webcam.html", Width: 1920, Height: 1080, Active: true},
}

// LogEntries são as entradas de log fake exibidas na tela de Logs.
var LogEntries = []LogEntry{
	{"14:32:01.342", "server", "SSE client conectado — total: 2"},
	{"14:32:00.891", "kick", "chatroom ID encontrado: 1234567"},
	{"14:31:59.200", "twitch", "conectado ao canal streampac"},
	{"14:31:58.100", "server", "HTTP server iniciado em :7777"},
	{"14:31:57.500", "hub", "broadcast hub iniciado"},
	{"14:31:56.000", "ia", "cliente Anthropic inicializado"},
	{"14:31:55.100", "server", "carregando config de ~/.config/stream-pac"},
	{"14:33:00.001", "twitch", "mensagem recebida de Pac"},
	{"14:33:01.200", "kick", "mensagem recebida de streampac"},
	{"14:33:05.500", "hub", "broadcast para 2 clientes SSE"},
	{"14:33:10.000", "twitch", "raid de xgamer50v — 120 viewers"},
	{"14:33:10.100", "hub", "evento raid enviado para clientes"},
	{"14:33:30.000", "server", "SSE heartbeat — 2 clientes ativos"},
}
