package components

import (
	"app/internal/components/colors"
	"app/internal/models"
	u "app/pkg/utils"
	"strings"

	css "github.com/charmbracelet/lipgloss"
)

func TranslatePage(it models.CtxMain) string {
	var langSwap = func() string {
		var styleFont = func(text string) string {
			return css.NewStyle().
				Bold(true).
				Render(text)
		}

		var styleBox = css.NewStyle().
			Padding(0, 2).
			Background(css.Color("#28282b"))

		return styleBox.Render(
			u.Ternary(it.CtxTranslate.SwitchLang, styleFont("PT > EN"), styleFont("EN > PT")),
		)
	}

	var textTranslate = css.NewStyle().
		Foreground(css.Color(colors.Blue)).
		Height(2).
		Render(it.CtxTranslate.Text)

	{ // TextInput
		it.TextInput.Prompt = prefixStyle("From ")
		it.TextInput.TextStyle = css.NewStyle().
			Foreground(css.Color(colors.White))
	}
	content := css.JoinVertical(css.Left,
		titleApp()+langSwap(),
		"",
		it.TextInput.View(),
		"",
		u.Ternary(it.Loading, loadingMessage, errorMessage(it.Err)),
		prefixStyle("To ")+textTranslate,
	)

	rows := strings.Count(content, "\n") + 2

	return box.Render(
		content,
		css.Place(it.Size.Width, it.Size.Height-rows, css.Left, css.Bottom, infoHelp),
	)
}
