package tui

import "fmt"

type Strings struct {
	Tabs func(n int) string
	Quit string
}

var PT = Strings{
	Tabs: func(n int) string { return fmt.Sprintf("[1-%d] abas", n) },
	Quit: "sair",
}

var EN = Strings{
	Tabs: func(n int) string { return fmt.Sprintf("[1-%d] tabs", n) },
	Quit: "quit",
}
