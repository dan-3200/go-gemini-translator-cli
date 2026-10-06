package models

import "github.com/charmbracelet/bubbles/textinput"

type CtxMain struct {
	Size struct {
		Height, Width int
	}
	Err        error
	Loading    bool
	SwitchMode bool
	TextInput  textinput.Model
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
