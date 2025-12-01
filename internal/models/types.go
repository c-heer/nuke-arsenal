package models

import "time"

// Command represents a single stored command
type Command struct {
	ID          int       `json:"id"`
	Cmd         string    `json:"cmd"`
	Description string    `json:"description"`
	Output      string    `json:"output,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	Note        string    `json:"note,omitempty"`
	Created     time.Time `json:"created"`
}

// Group represents a collection of related commands
type Group struct {
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	Commands    []Command `json:"commands"`
}

// CommandsFile is the root structure of commands.json
type CommandsFile struct {
	Meta   Meta             `json:"_meta"`
	Groups map[string]Group `json:"groups"`
}

// Meta contains file metadata
type Meta struct {
	Purpose string `json:"purpose"`
	Version string `json:"version"`
}

// Config represents the app configuration
type Config struct {
	DataPath string `json:"data_path"`
}
