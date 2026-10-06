package components

import (
	"app/internal/components/colors"
	"app/internal/models"
	"strings"

	css "github.com/charmbracelet/lipgloss"
)

func DictPage(it models.CtxMain) string {
	var dictFlag = css.NewStyle().
		Bold(true).
		Padding(0, 2).
		Background(css.Color("#28282b")).
		Render("Dictionary")

	{ // TextInput
		it.TextInput.Prompt = prefixStyle("Search: ")
		it.TextInput.TextStyle = css.NewStyle().Foreground(css.Color(colors.White))
	}

	content := css.JoinVertical(css.Left,
		titleApp()+dictFlag,
		"",
		it.TextInput.View(),
		"",
		loadingOrError(it.Loading, it.Err),
		fontColor(it.Dictionary.Word, colors.SteelBlue200),
		fontColor(it.Dictionary.PartOfSpeech, colors.SteelBlue400),
		fontColor(it.Dictionary.Definition, colors.Yellow),
		fontColor(`"`+it.Dictionary.Example+`"`, colors.SteelBlue200),
		fontColor(it.Dictionary.Collocations, colors.SteelBlue300),
		fontColor(it.Dictionary.Synonyms, colors.SteelBlue500),
		"",
	)

	rows := strings.Count(content, "\n") + 2

	return box.Render(
		content,
		css.Place(it.Size.Width, it.Size.Height-rows, css.Left, css.Bottom, infoHelp),
	)
}

func loadingOrError(loading bool, err error) string {
	if loading {
		return loadingMessage
	}
	return errorMessage(err)
}
