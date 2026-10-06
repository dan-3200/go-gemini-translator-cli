package engine

import (
	"app/internal/agentAI"
	m "app/internal/models"

	"github.com/charmbracelet/bubbles/textinput"
)

// Contexto Geral da aplicação responsável variáveis persistentes durante a renderização
type App struct {
	m.CtxMain
}

func SetApp() *App {
	var ti = textinput.New()
	ti.Placeholder = "..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 100

	return &App{
		CtxMain: m.CtxMain{
			Size: struct{ Height, Width int }{
				Height: 0,
				Width:  0,
			},
			TextInput:  ti,
			SwitchMode: false,
			CtxTranslate: m.CtxTranslate{
				Text:       "...",
				SwitchLang: false,
			},
			CtxDict: m.CtxDict{
				Dictionary: agentAI.EmptyDictionaryEntry(),
			},
		},
	}
}
