package components

import (
	"app/internal/models"

	css "github.com/charmbracelet/lipgloss"
)

func SettingsPage(it models.CtxMain) string {
	content := css.JoinVertical(css.Left,
		titleApp()+" Settings",
		"",
		"Enter your Gemini API key. It will be saved in the system credential store.",
		"",
		it.APIKeyInput.View(),
		loadingOrError(it.Loading, it.Err),
	)

	help := css.NewStyle().
		Foreground(css.Color("#606060")).
		Render("[Enter] Save • [Ctrl+d] Remove saved key • [Esc] Cancel")

	rows := 8
	return box.Render(
		content,
		css.Place(it.Size.Width, it.Size.Height-rows, css.Left, css.Bottom, help),
	)
}
