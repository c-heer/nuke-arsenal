package services

import (
	"strings"
	"time"

	"github.com/c-heer/nuke-arsenal/internal/data"
	"github.com/c-heer/nuke-arsenal/internal/models"
)

// ArsenalService provides Wails bindings for the frontend
type ArsenalService struct {
	dataPath string
}

// NewArsenalService creates a new service instance
func NewArsenalService() *ArsenalService {
	return &ArsenalService{}
}

// --- Config ---

// HasConfig checks if app is configured
func (s *ArsenalService) HasConfig() bool {
	return data.HasConfig()
}

// GetConfig returns current config
func (s *ArsenalService) GetConfig() (*models.Config, error) {
	return data.ReadConfig()
}

// SetDataPath saves the data path and loads it
func (s *ArsenalService) SetDataPath(path string) error {
	config := &models.Config{DataPath: path}
	if err := data.WriteConfig(config); err != nil {
		return err
	}
	s.dataPath = path
	return nil
}

// --- Commands ---

func (s *ArsenalService) ensureDataPath() error {
	if s.dataPath == "" {
		config, err := data.ReadConfig()
		if err != nil {
			return err
		}
		if config != nil {
			s.dataPath = config.DataPath
		}
	}
	return nil
}

// GetCommands returns all commands data
func (s *ArsenalService) GetCommands() (*models.CommandsFile, error) {
	if err := s.ensureDataPath(); err != nil {
		return nil, err
	}
	return data.ReadCommands(s.dataPath)
}

// GetGroups returns just the groups map
func (s *ArsenalService) GetGroups() (map[string]models.Group, error) {
	commands, err := s.GetCommands()
	if err != nil {
		return nil, err
	}
	return commands.Groups, nil
}

// AddGroup creates a new group
func (s *ArsenalService) AddGroup(key, name, icon, description string) error {
	commands, err := s.GetCommands()
	if err != nil {
		return err
	}

	commands.Groups[key] = models.Group{
		Name:        name,
		Icon:        icon,
		Description: description,
		Commands:    []models.Command{},
	}

	return data.WriteCommands(s.dataPath, commands)
}

// DeleteGroup removes a group
func (s *ArsenalService) DeleteGroup(key string) error {
	commands, err := s.GetCommands()
	if err != nil {
		return err
	}

	delete(commands.Groups, key)
	return data.WriteCommands(s.dataPath, commands)
}

// AddCommand adds a command to a group
func (s *ArsenalService) AddCommand(groupKey string, cmd, description, output, note string, tags []string) error {
	commands, err := s.GetCommands()
	if err != nil {
		return err
	}

	group, exists := commands.Groups[groupKey]
	if !exists {
		return nil // Group doesn't exist
	}

	// Find next ID
	maxID := 0
	for _, c := range group.Commands {
		if c.ID > maxID {
			maxID = c.ID
		}
	}

	newCmd := models.Command{
		ID:          maxID + 1,
		Cmd:         cmd,
		Description: description,
		Output:      output,
		Note:        note,
		Tags:        tags,
		Created:     time.Now(),
	}

	group.Commands = append(group.Commands, newCmd)
	commands.Groups[groupKey] = group

	return data.WriteCommands(s.dataPath, commands)
}

// UpdateCommand updates an existing command
func (s *ArsenalService) UpdateCommand(groupKey string, id int, cmd, description, output, note string, tags []string) error {
	commands, err := s.GetCommands()
	if err != nil {
		return err
	}

	group, exists := commands.Groups[groupKey]
	if !exists {
		return nil
	}

	for i, c := range group.Commands {
		if c.ID == id {
			group.Commands[i].Cmd = cmd
			group.Commands[i].Description = description
			group.Commands[i].Output = output
			group.Commands[i].Note = note
			group.Commands[i].Tags = tags
			break
		}
	}

	commands.Groups[groupKey] = group
	return data.WriteCommands(s.dataPath, commands)
}

// DeleteCommand removes a command from a group
func (s *ArsenalService) DeleteCommand(groupKey string, id int) error {
	commands, err := s.GetCommands()
	if err != nil {
		return err
	}

	group, exists := commands.Groups[groupKey]
	if !exists {
		return nil
	}

	for i, c := range group.Commands {
		if c.ID == id {
			group.Commands = append(group.Commands[:i], group.Commands[i+1:]...)
			break
		}
	}

	commands.Groups[groupKey] = group
	return data.WriteCommands(s.dataPath, commands)
}

// Search searches across all groups
func (s *ArsenalService) Search(query string) ([]SearchResult, error) {
	commands, err := s.GetCommands()
	if err != nil {
		return nil, err
	}

	var results []SearchResult
	query = strings.ToLower(query)

	for groupKey, group := range commands.Groups {
		for _, cmd := range group.Commands {
			if strings.Contains(strings.ToLower(cmd.Cmd), query) ||
				strings.Contains(strings.ToLower(cmd.Description), query) ||
				containsTag(cmd.Tags, query) {
				results = append(results, SearchResult{
					GroupKey:  groupKey,
					GroupName: group.Name,
					Command:   cmd,
				})
			}
		}
	}

	return results, nil
}

// SearchResult represents a search match
type SearchResult struct {
	GroupKey  string         `json:"groupKey"`
	GroupName string         `json:"groupName"`
	Command   models.Command `json:"command"`
}

func containsTag(tags []string, query string) bool {
	for _, tag := range tags {
		if strings.Contains(strings.ToLower(tag), query) {
			return true
		}
	}
	return false
}
