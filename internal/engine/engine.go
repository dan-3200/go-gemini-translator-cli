package engine

import (
	"app/internal/agentAI"
	comp "app/internal/components"
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
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return it, tea.Quit
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

func (it App) View() string {
	return u.Ternary(
		it.SwitchMode, comp.TranslatePage(it.CtxMain), comp.DictPage(it.CtxMain),
	)
}
