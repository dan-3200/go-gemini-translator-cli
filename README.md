# Gemini Translator Golang CLI

![image](https://github.com/user-attachments/assets/76b35e1d-3a01-466c-bf11-326f3fd9ef5c)

## Description

A command-line translation tool built in `Golang` using the Bubble Tea TUI library and integrated with the Gemini API. This app provides accurate and context-aware sentence `translation` and a `dictionary` feature for English words, including part of speech, definition, example, synonyms, and collocations. Developed as a practical project to learn and master the Go programming language.

## Installation & Usage

```bash
git clone https://github.com/DAN-3200/gemini-translator-cli.git
cd gemini-translator-cli
```

##### API Key Configuration

On first launch, the app asks for your **Gemini API key** and stores it in your operating system's credential manager. The key is not included in the executable.

For local development, you can alternatively create a `.env` file in the project root:

```env
GEMINI_API_KEY=your-gemini-api-key-here
```

**Important for development:**

* Replace `your-gemini-api-key-here` with your actual Gemini API key from [Google AI Studio](https://makersuite.google.com/app/apikey).
* The app uses the saved key first and falls back to `GEMINI_API_KEY` when no saved key exists.

Run directly (for development):

```bash
go run .
```

Build and execute:

```bash
go build -o app.exe
./app.exe
```

## Model

```go
type DictionaryEntry struct {
  Word         string `json:"word"`
  PartOfSpeech string `json:"partOfSpeech"`
  Definition   string `json:"definition"`
  Example      string `json:"example"`
  Synonyms     string `json:"synonyms"`
  Collocations string `json:"collocations"`
}
```

## Prompts for Gemini

prompt for `translation`
```txt
Translate from %s: "%s". Return only the exact translation, no explanation.
```

prompt for `dictionary`
```txt
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
```

## Links 

- Link Go: [https://pkg.go.dev/golang.org](https://pkg.go.dev/golang.org)
- Link Bubble Tea: [https://github.com/charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea)


## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
