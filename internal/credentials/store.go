package credentials

import (
	"errors"
	"strings"

	"github.com/zalando/go-keyring"
)

const (
	serviceName = "loro-translator"
	accountName = "gemini-api-key"
)

var ErrNotFound = errors.New("chave da API não encontrada")

func LoadAPIKey() (string, error) {
	apiKey, err := keyring.Get(serviceName, accountName)
	if errors.Is(err, keyring.ErrNotFound) {
		return "", ErrNotFound
	}
	if err != nil {
		return "", err
	}

	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return "", ErrNotFound
	}
	return apiKey, nil
}

func SaveAPIKey(apiKey string) error {
	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return errors.New("informe uma chave da API")
	}
	return keyring.Set(serviceName, accountName, apiKey)
}

func DeleteAPIKey() error {
	err := keyring.Delete(serviceName, accountName)
	if errors.Is(err, keyring.ErrNotFound) {
		return nil
	}
	return err
}
