package agentAI

import (
	"testing"

	Gemini "github.com/google/generative-ai-go/genai"
)

func TestCleanJSON(t *testing.T) {
	tests := map[string]string{
		"plain JSON":       `{"word":"book"}`,
		"markdown JSON":    "```json\n{\"word\":\"book\"}\n```",
		"generic markdown": "```\n{\"word\":\"book\"}\n```",
	}

	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			if got, want := cleanJSON(input), `{"word":"book"}`; got != want {
				t.Fatalf("cleanJSON() = %q, want %q", got, want)
			}
		})
	}
}

func TestResponseText(t *testing.T) {
	response := &Gemini.GenerateContentResponse{
		Candidates: []*Gemini.Candidate{{
			Content: &Gemini.Content{Parts: []Gemini.Part{Gemini.Text("hello"), Gemini.Text(" world")}},
		}},
	}

	got, err := responseText(response)
	if err != nil {
		t.Fatalf("responseText() returned an error: %v", err)
	}
	if want := "hello world"; got != want {
		t.Fatalf("responseText() = %q, want %q", got, want)
	}
}

func TestResponseTextRejectsEmptyResponse(t *testing.T) {
	if _, err := responseText(&Gemini.GenerateContentResponse{}); err == nil {
		t.Fatal("responseText() should reject an empty response")
	}
}
