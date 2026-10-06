package internal

import (
	"app/internal/agentAI"
	"app/internal/credentials"
	"app/internal/engine"
	"errors"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	apiKey, err := apiKeyFromStoreOrEnvironment()
	if err != nil {
		fmt.Printf("Erro ao acessar a chave da API: %v\n", err)
		return
	}

	needsSetup := apiKey == ""
	if !needsSetup {
		if err := agentAI.InitGemini(apiKey); err != nil {
			fmt.Printf("Erro ao iniciar o Gemini: %v\n", err)
			return
		}
	}
	defer func() {
		if err := agentAI.Close(); err != nil {
			fmt.Printf("Erro ao fechar o cliente Gemini: %v\n", err)
		}
	}()

	program := tea.NewProgram(engine.SetApp(needsSetup))
	_, err = program.Run()
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	}
}

func apiKeyFromStoreOrEnvironment() (string, error) {
	apiKey, err := credentials.LoadAPIKey()
	if err == nil {
		return apiKey, nil
	}
	if !errors.Is(err, credentials.ErrNotFound) {
		return "", err
	}

	return strings.TrimSpace(os.Getenv("GEMINI_API_KEY")), nil
}
