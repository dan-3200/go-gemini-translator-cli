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

func SetApp(needsSetup bool) *App {
	ti := textinput.New()
	ti.Placeholder = "..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 100

	apiKeyInput := textinput.New()
	apiKeyInput.Placeholder = "Cole sua Gemini API key"
	apiKeyInput.EchoMode = textinput.EchoPassword
	apiKeyInput.EchoCharacter = '•'
	apiKeyInput.CharLimit = 512
	apiKeyInput.Width = 100
	if needsSetup {
		ti.Blur()
		apiKeyInput.Focus()
	}

	return &App{
		CtxMain: m.CtxMain{
			Size: struct{ Height, Width int }{
				Height: 0,
				Width:  0,
			},
			TextInput:   ti,
			APIKeyInput: apiKeyInput,
			Configuring: needsSetup,
			HasAPIKey:   !needsSetup,
			SwitchMode:  false,
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
