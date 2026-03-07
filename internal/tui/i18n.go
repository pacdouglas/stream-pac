package tui

// Strings contém todos os textos visíveis da UI.
// Trocar idioma = trocar de Strings struct — sem magic strings espalhadas.
type Strings struct {
	// ── Tab bar ───────────────────────────────────────────────────────────────
	TabDashboard string
	TabOverlays  string
	TabConfig    string
	TabAI        string
	TabLogs      string

	// ── Dashboard ─────────────────────────────────────────────────────────────
	DashPlatforms  string
	DashViewers    string
	DashTotal      string
	DashSession    string
	DashUptime     string
	DashMessages   string
	DashClients    string
	DashChatTitle  string
	DashChatSearch string
	DashChatFilter string
	DashEvents     string

	// ── Overlays ──────────────────────────────────────────────────────────────
	OvTitle      string
	OvActive     string
	OvInactive   string
	OvAIItem     string
	OvAIDesc     string
	OvPreview    string
	OvSimulation string
	OvOBSSource  string
	OvURL        string
	OvWidth      string
	OvHeight     string
	OvCopyHint   string
	OvAIHint     string
	OvCopied     string
	OvOpening    string
	OvRemoved    string

	// ── Config ────────────────────────────────────────────────────────────────
	CfgPlatforms   string
	CfgTwitchCh    string
	CfgKickCh      string
	CfgKickRoom    string
	CfgYouTubeID   string
	CfgRoomHint    string
	CfgYouTubeHint string
	CfgServer      string
	CfgPort        string
	CfgHost        string
	CfgAI          string
	CfgAPIKey      string
	CfgModel       string
	CfgProvider    string
	CfgSave        string
	CfgDiscard     string
	CfgSaved       string

	// ── AI Gen ────────────────────────────────────────────────────────────────
	AIPromptTitle string
	AIPromptHint  string
	AIType        string
	AIChat        string
	AIWebcam      string
	AIAlert       string
	AIViewers     string
	AIHypeTrain   string
	AIOptions     string
	AITheme       string
	AIPalette     string
	AIWidth       string
	AIHeight      string
	AIGenerate    string
	AIStatus      string
	AIResult      string
	AIWaiting     string
	AIWaitingHint string
	AIGenerating  string
	AIToast       string

	// ── Logs ──────────────────────────────────────────────────────────────────
	LogTitle      string
	LogAll        string
	LogFilterHint string

	// ── Footer / global ───────────────────────────────────────────────────────
	FooterQuit   string
	FooterHelp   string
	FooterScroll string
	FooterFilter string
	FooterSearch string

	// ── Help overlay ──────────────────────────────────────────────────────────
	HelpTitle  string
	HelpGlobal string
	HelpClose  string
}

// PT é o conjunto de strings em português.
var PT = Strings{
	TabDashboard: "Painel",
	TabOverlays:  "Overlays",
	TabConfig:    "Config",
	TabAI:        "IA",
	TabLogs:      "Logs",

	DashPlatforms:  "Plataformas",
	DashViewers:    "Viewers",
	DashTotal:      "Total",
	DashSession:    "Sessão",
	DashUptime:     "Uptime",
	DashMessages:   "Msgs",
	DashClients:    "Clients",
	DashChatTitle:  "Chat ao vivo",
	DashChatSearch: "buscar",
	DashChatFilter: "filtrar",
	DashEvents:     "Eventos",

	OvTitle:      "Overlays disponíveis",
	OvActive:     "ativo",
	OvInactive:   "inativo",
	OvAIItem:     "✦ Gerar novo com IA",
	OvAIDesc:     "vai para a tela [4] IA",
	OvPreview:    "Preview",
	OvSimulation: "Simulação visual",
	OvOBSSource:  "OBS Browser Source",
	OvURL:        "URL",
	OvWidth:      "Largura",
	OvHeight:     "Altura",
	OvCopyHint:   "[enter] copiar URL  [o] abrir no browser",
	OvAIHint:     "Pressione [enter] para ir para\n  a tela [4] IA.",
	OvCopied:     "URL copiada para o clipboard!",
	OvOpening:    "Abrindo no browser...",
	OvRemoved:    "Overlay removido.",

	CfgPlatforms:   "Plataformas",
	CfgTwitchCh:    "Twitch channel:",
	CfgKickCh:      "Kick channel:",
	CfgKickRoom:    "Kick chatroom ID:",
	CfgYouTubeID:   "YouTube video ID:",
	CfgRoomHint:    "Chatroom ID vazio = auto-detect.",
	CfgYouTubeHint: "YouTube: cole o ID da live (ex: Fpfdw0iXuv8)",
	CfgServer:      "Servidor",
	CfgPort:        "Porta HTTP:",
	CfgHost:        "Host:",
	CfgAI:          "IA / Gerador",
	CfgAPIKey:      "Anthropic API Key:",
	CfgModel:       "Modelo:",
	CfgProvider:    "Provedor IA:",
	CfgSave:        " Salvar e aplicar ",
	CfgDiscard:     " Descartar ",
	CfgSaved:       "Configurações salvas!",

	AIPromptTitle: "Descreva o overlay",
	AIPromptHint:  "[ctrl+enter] gerar",
	AIType:        "Tipo",
	AIChat:        "Chat completo",
	AIWebcam:      "Webcam overlay",
	AIAlert:       "Alerta de evento",
	AIViewers:     "Contador de viewers",
	AIHypeTrain:   "Hype train",
	AIOptions:     "Opções",
	AITheme:       "Tema:",
	AIPalette:     "Paleta:",
	AIWidth:       "Largura:",
	AIHeight:      "Altura:",
	AIGenerate:    "✦ Gerar overlay",
	AIStatus:      "Status",
	AIResult:      "Resultado",
	AIWaiting:     "Aguardando prompt...",
	AIWaitingHint: "Use ctrl+enter para gerar.",
	AIGenerating:  "Gerando...",
	AIToast:       "Overlay gerado! Veja em [2] Overlays.",

	LogTitle:      "Log de eventos do sistema",
	LogAll:        "todos",
	LogFilterHint: "[f] filtrar  [c] exportar  [del] limpar",

	FooterQuit:   "sair",
	FooterHelp:   "ajuda",
	FooterScroll: "scroll",
	FooterFilter: "filtrar",
	FooterSearch: "buscar",

	HelpTitle:  " Atalhos de teclado ",
	HelpGlobal: "Global",
	HelpClose:  "fechar",
}

// EN é o conjunto de strings em inglês.
var EN = Strings{
	TabDashboard: "Dashboard",
	TabOverlays:  "Overlays",
	TabConfig:    "Config",
	TabAI:        "AI",
	TabLogs:      "Logs",

	DashPlatforms:  "Platforms",
	DashViewers:    "Viewers",
	DashTotal:      "Total",
	DashSession:    "Session",
	DashUptime:     "Uptime",
	DashMessages:   "Msgs",
	DashClients:    "Clients",
	DashChatTitle:  "Live chat",
	DashChatSearch: "search",
	DashChatFilter: "filter",
	DashEvents:     "Events",

	OvTitle:      "Available overlays",
	OvActive:     "active",
	OvInactive:   "inactive",
	OvAIItem:     "✦ Generate with AI",
	OvAIDesc:     "goes to screen [4] AI",
	OvPreview:    "Preview",
	OvSimulation: "Visual simulation",
	OvOBSSource:  "OBS Browser Source",
	OvURL:        "URL",
	OvWidth:      "Width",
	OvHeight:     "Height",
	OvCopyHint:   "[enter] copy URL  [o] open in browser",
	OvAIHint:     "Press [enter] to go to\n  screen [4] AI.",
	OvCopied:     "URL copied to clipboard!",
	OvOpening:    "Opening in browser...",
	OvRemoved:    "Overlay removed.",

	CfgPlatforms:   "Platforms",
	CfgTwitchCh:    "Twitch channel:",
	CfgKickCh:      "Kick channel:",
	CfgKickRoom:    "Kick chatroom ID:",
	CfgYouTubeID:   "YouTube video ID:",
	CfgRoomHint:    "Empty chatroom ID = auto-detect.",
	CfgYouTubeHint: "YouTube: paste the live ID (e.g. Fpfdw0iXuv8)",
	CfgServer:      "Server",
	CfgPort:        "HTTP Port:",
	CfgHost:        "Host:",
	CfgAI:          "AI / Generator",
	CfgAPIKey:      "Anthropic API Key:",
	CfgModel:       "Model:",
	CfgProvider:    "AI Provider:",
	CfgSave:        " Save & apply ",
	CfgDiscard:     " Discard ",
	CfgSaved:       "Settings saved!",

	AIPromptTitle: "Describe the overlay",
	AIPromptHint:  "[ctrl+enter] generate",
	AIType:        "Type",
	AIChat:        "Full chat",
	AIWebcam:      "Webcam overlay",
	AIAlert:       "Event alert",
	AIViewers:     "Viewer counter",
	AIHypeTrain:   "Hype train",
	AIOptions:     "Options",
	AITheme:       "Theme:",
	AIPalette:     "Palette:",
	AIWidth:       "Width:",
	AIHeight:      "Height:",
	AIGenerate:    "✦ Generate overlay",
	AIStatus:      "Status",
	AIResult:      "Result",
	AIWaiting:     "Waiting for prompt...",
	AIWaitingHint: "Use ctrl+enter to generate.",
	AIGenerating:  "Generating...",
	AIToast:       "Overlay generated! See it in [2] Overlays.",

	LogTitle:      "System event log",
	LogAll:        "all",
	LogFilterHint: "[f] filter  [c] export  [del] clear",

	FooterQuit:   "quit",
	FooterHelp:   "help",
	FooterScroll: "scroll",
	FooterFilter: "filter",
	FooterSearch: "search",

	HelpTitle:  " Keyboard shortcuts ",
	HelpGlobal: "Global",
	HelpClose:  "close",
}
