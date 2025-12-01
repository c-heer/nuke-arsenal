package data

import (
	"encoding/json"
	"os"

	"github.com/c-heer/nuke-arsenal/internal/models"
)

// ReadCommands reads the commands file from configured path
func ReadCommands(dataPath string) (*models.CommandsFile, error) {
	data, err := os.ReadFile(dataPath)
	if err != nil {
		if os.IsNotExist(err) {
			return NewCommandsFile(), nil
		}
		return nil, err
	}

	var commands models.CommandsFile
	if err := json.Unmarshal(data, &commands); err != nil {
		return nil, err
	}

	return &commands, nil
}

// WriteCommands saves the commands file
func WriteCommands(dataPath string, commands *models.CommandsFile) error {
	if err := EnsureDataDir(dataPath); err != nil {
		return err
	}

	data, err := json.MarshalIndent(commands, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataPath, data, 0644)
}

// NewCommandsFile creates a new empty commands file
func NewCommandsFile() *models.CommandsFile {
	return &models.CommandsFile{
		Meta: models.Meta{
			Purpose: "Personal command knowledge base",
			Version: "1.0",
		},
		Groups: make(map[string]models.Group),
	}
}
