package agentAI

import (
	"app/internal/models"
	"app/pkg/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	Gemini "github.com/google/generative-ai-go/genai"

	"google.golang.org/api/option"
)

var (
	ctx    = context.Background()
	client *Gemini.Client
	model  *Gemini.GenerativeModel
)

const modelName = "gemini-3.5-flash-lite"

func InitGemini(apiKey string) error {
	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return errors.New("GEMINI_API_KEY não foi definida")
	}

	newClient, err := Gemini.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return fmt.Errorf("criar cliente Gemini: %w", err)
	}

	client = newClient
	model = client.GenerativeModel(modelName)
	return nil
}

func Close() error {
	if client == nil {
		return nil
	}
	err := client.Close()
	client = nil
	model = nil
	return err
}

func UseTranslation(text string, switchLang bool) (string, error) {
	if text == "" {
		return "...", nil
	}
	if model == nil {
		return "", errors.New("cliente Gemini não inicializado")
	}

	langs := utils.Ternary(switchLang, "PT-BR to EN", "EN to PT-BR")
	command := `Translate from %s: "%s". Return only the exact translation, no explanation.`
	prompt := Gemini.Text(
		fmt.Sprintf(command, langs, text),
	)

	response, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("gerar tradução: %w", err)
	}

	return responseText(response)
}

func UseDictionary(word string) (models.DictionaryEntry, error) {
	if len(word) <= 0 {
		return EmptyDictionaryEntry(), nil
	}
	if model == nil {
		return models.DictionaryEntry{}, errors.New("cliente Gemini não inicializado")
	}

	command := `
		Task: Return a strict JSON object for the word "%s".

		Rules:
		- Output must be only the JSON object. No text, no comments, no formatting.
		- The JSON must be valid: use double quotes and no trailing commas.
		- Use simple, learner-friendly English in all fields.
		- The example sentence must be natural, common, and end without a period.
	 	- If the word is a verb, format partOfSpeech as: "verb - [tense]".

		Schema:
		{
			"word": string,
			"partOfSpeech": string,
			"definition": string, 
			"example": string,
			"synonyms": string // comma-separated
			"collocations": string, // comma-separated; frequent word combinations 
		}
	`
	prompt := Gemini.Text(
		fmt.Sprintf(command, word),
	)

	response, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		return models.DictionaryEntry{}, fmt.Errorf("consultar dicionário: %w", err)
	}

	dataJSON, err := responseText(response)
	if err != nil {
		return models.DictionaryEntry{}, err
	}
	var data models.DictionaryEntry
	if err := json.Unmarshal([]byte(cleanJSON(dataJSON)), &data); err != nil {
		return models.DictionaryEntry{}, fmt.Errorf("interpretar resposta do dicionário: %w", err)
	}

	return data, nil
}

func EmptyDictionaryEntry() models.DictionaryEntry {
	return models.DictionaryEntry{
		Word:         "Word",
		PartOfSpeech: "Part of speech",
		Definition:   "Definition",
		Example:      "Example",
		Synonyms:     "Synonyms",
		Collocations: "Collocations",
	}
}

func responseText(response *Gemini.GenerateContentResponse) (string, error) {
	if response == nil || len(response.Candidates) == 0 || response.Candidates[0] == nil || response.Candidates[0].Content == nil {
		return "", errors.New("Gemini não retornou conteúdo")
	}

	var text strings.Builder
	for _, part := range response.Candidates[0].Content.Parts {
		if value, ok := part.(Gemini.Text); ok {
			text.WriteString(string(value))
		}
	}

	result := strings.TrimSpace(text.String())
	if result == "" {
		return "", errors.New("Gemini retornou conteúdo vazio")
	}
	return result, nil
}

func cleanJSON(value string) string {
	value = strings.TrimSpace(value)
	value = strings.TrimPrefix(value, "```json")
	value = strings.TrimPrefix(value, "```")
	value = strings.TrimSuffix(value, "```")
	return strings.TrimSpace(value)
}
