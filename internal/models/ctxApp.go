package models

import "github.com/charmbracelet/bubbles/textinput"

type CtxMain struct {
	Size struct {
		Height, Width int
	}
	Err         error
	Loading     bool
	Configuring bool
	HasAPIKey   bool
	SwitchMode  bool
	TextInput   textinput.Model
	APIKeyInput textinput.Model
	CtxDict
	CtxTranslate
}

type CtxTranslate struct {
	Text       string
	SwitchLang bool
}

type CtxDict struct {
	Dictionary DictionaryEntry
}
