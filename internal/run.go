package internal

import (
	"app/internal/agentAI"
	"app/internal/engine"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	if err := agentAI.InitGemini(); err != nil {
		fmt.Printf("Erro ao iniciar o Gemini: %v\n", err)
		return
	}
	defer func() {
		if err := agentAI.Close(); err != nil {
			fmt.Printf("Erro ao fechar o cliente Gemini: %v\n", err)
		}
	}()

	program := tea.NewProgram(engine.SetApp())
	_, err := program.Run()
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	}
}
