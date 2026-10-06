package engine

import (
	"app/internal/agentAI"
	comp "app/internal/components"
	"app/internal/credentials"
	"app/internal/models"
	u "app/pkg/utils"

	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/textinput"
)

type translationResponseMsg struct {
	text string
	err  error
}

type dictionaryResponseMsg struct {
	entry models.DictionaryEntry
	err   error
}

type apiKeySavedMsg struct{ err error }

type apiKeyDeletedMsg struct{ err error }

func (it App) Init() tea.Cmd {
	return textinput.Blink
}

func (it App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		it.Size.Height = msg.Height
		it.Size.Width = msg.Width

	case tea.KeyMsg:
		if it.Configuring {
			switch msg.Type {
			case tea.KeyCtrlC:
				return it, tea.Quit
			case tea.KeyEsc:
				if it.HasAPIKey {
					it.Configuring = false
					it.APIKeyInput.Reset()
					it.APIKeyInput.Blur()
					it.TextInput.Focus()
					return it, nil
				}
				return it, tea.Quit
			case tea.KeyEnter:
				if !it.Loading {
					it.Err = nil
					it.Loading = true
					return it, saveAPIKeyCmd(it.APIKeyInput.Value())
				}
			case tea.KeyCtrlD:
				if !it.Loading && it.HasAPIKey {
					it.Err = nil
					it.Loading = true
					return it, deleteAPIKeyCmd()
				}
			}

			it.APIKeyInput, cmd = it.APIKeyInput.Update(msg)
			return it, cmd
		}

		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return it, tea.Quit
		case tea.KeyCtrlK:
			if !it.Loading {
				it.Err = nil
				it.Configuring = true
				it.TextInput.Blur()
				it.APIKeyInput.Focus()
				return it, nil
			}
		case tea.KeyCtrlT:
			if it.SwitchMode && !it.Loading {
				it.CtxTranslate.SwitchLang = !it.CtxTranslate.SwitchLang
			}
		case tea.KeyEnter:
			if !it.Loading {
				it.Err = nil
				it.Loading = true
				query := strings.TrimSpace(it.TextInput.Value())
				if it.SwitchMode {
					return it, translateCmd(query, it.CtxTranslate.SwitchLang)
				}
				return it, dictionaryCmd(query)
			}
		case tea.KeyCtrlU:
			if !it.Loading {
				it.TextInput.Reset()
			}
		case tea.KeyCtrlA:
			if !it.Loading {
				it.SwitchMode = !it.SwitchMode
			}
		}

	case translationResponseMsg:
		it.Loading = false
		if msg.err != nil {
			it.Err = msg.err
			return it, nil
		}
		it.CtxTranslate.Text = msg.text

	case dictionaryResponseMsg:
		it.Loading = false
		if msg.err != nil {
			it.Err = msg.err
			return it, nil
		}
		it.CtxDict.Dictionary = msg.entry

	case apiKeySavedMsg:
		it.Loading = false
		if msg.err != nil {
			it.Err = msg.err
			return it, nil
		}
		it.Err = nil
		it.HasAPIKey = true
		it.Configuring = false
		it.APIKeyInput.Reset()
		it.APIKeyInput.Blur()
		it.TextInput.Focus()

	case apiKeyDeletedMsg:
		it.Loading = false
		if msg.err != nil {
			it.Err = msg.err
			return it, nil
		}
		it.HasAPIKey = false
		it.APIKeyInput.Reset()

	case error:
		it.Err = msg
		return it, nil
	}

	it.TextInput, cmd = it.TextInput.Update(msg)
	return it, cmd
}

func translateCmd(text string, switchLang bool) tea.Cmd {
	return func() tea.Msg {
		translation, err := agentAI.UseTranslation(text, switchLang)
		return translationResponseMsg{text: translation, err: err}
	}
}

func dictionaryCmd(word string) tea.Cmd {
	return func() tea.Msg {
		entry, err := agentAI.UseDictionary(word)
		return dictionaryResponseMsg{entry: entry, err: err}
	}
}

func saveAPIKeyCmd(apiKey string) tea.Cmd {
	return func() tea.Msg {
		if err := credentials.SaveAPIKey(apiKey); err != nil {
			return apiKeySavedMsg{err: err}
		}
		if err := agentAI.Close(); err != nil {
			return apiKeySavedMsg{err: err}
		}
		return apiKeySavedMsg{err: agentAI.InitGemini(apiKey)}
	}
}

func deleteAPIKeyCmd() tea.Cmd {
	return func() tea.Msg {
		if err := credentials.DeleteAPIKey(); err != nil {
			return apiKeyDeletedMsg{err: err}
		}
		return apiKeyDeletedMsg{err: agentAI.Close()}
	}
}

func (it App) View() string {
	if it.Configuring {
		return comp.SettingsPage(it.CtxMain)
	}
	return u.Ternary(
		it.SwitchMode, comp.TranslatePage(it.CtxMain), comp.DictPage(it.CtxMain),
	)
}
