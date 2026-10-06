package components

import (
	"app/internal/components/colors"

	css "github.com/charmbracelet/lipgloss"
)

// -- Styles and Components -----------------------------------------------------

var infoHelp = css.NewStyle().
	Foreground(css.Color("#606060")).
	Render("[Ctrl+u] Clear • [Ctrl+a] Switch mode • [Ctrl+t] Switch language • [Ctrl+k] API key • [Esc] Exit")

var box = css.NewStyle().Padding(1)

var prefixStyle = func(text string) string {
	return css.NewStyle().
		Foreground(css.Color("#606060")).
		Render(text)
}

var titleApp = func() string {
	return css.NewStyle().
		Padding(0, 2).
		Bold(true).
		Background(css.Color(colors.SteelBlue500)).Foreground(css.Color(colors.White)).
		Render("Loro App")
}

var fontColor = func(text string, color string) string {
	return css.NewStyle().
		Foreground(css.Color(color)).
		Render(text)
}

var errorMessage = func(err error) string {
	if err == nil {
		return ""
	}
	return css.NewStyle().Foreground(css.Color("#ff6b6b")).Render("Error: " + err.Error())
}

var loadingMessage = css.NewStyle().
	Foreground(css.Color("#ffd166")).
	Render("Loading...")
