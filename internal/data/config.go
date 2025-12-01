package data

import (
	"encoding/json"
	"os"

	"github.com/c-heer/nuke-arsenal/internal/models"
)

// ReadConfig reads the config file
func ReadConfig() (*models.Config, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // No config yet
		}
		return nil, err
	}

	var config models.Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// WriteConfig saves the config file
func WriteConfig(config *models.Config) error {
	if err := EnsureConfigDir(); err != nil {
		return err
	}

	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// HasConfig checks if config exists
func HasConfig() bool {
	config, _ := ReadConfig()
	return config != nil && config.DataPath != ""
}
