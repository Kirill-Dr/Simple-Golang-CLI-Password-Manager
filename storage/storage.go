package storage

import (
	"encoding/json"
	"errors"
	"os"
)

type PasswordEntry struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const filePath = "passwords.json"

func LoadPasswords() ([]PasswordEntry, error) {
	file, err := os.ReadFile(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return []PasswordEntry{}, nil
	}
	if err != nil {
		return nil, err
	}

	var passwords []PasswordEntry
	err = json.Unmarshal(file, &passwords)
	return passwords, err
}

func SavePasswords(passwords []PasswordEntry) error {
	data, err := json.MarshalIndent(passwords, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
