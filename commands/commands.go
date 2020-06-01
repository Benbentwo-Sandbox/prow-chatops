package commands

import (
	"fmt"
	"github.com/Benbentwo/utils/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Command struct {
	Name    string   `json:"name"`
	Enabled bool     `json:"enabled,omitempty"`
	Flags   []string `json:"flags,omitempty"` // like /approve no-issue
	Args    bool     `json:"args,omitempty"`  // like /label @foo @bar ...
}

type Commands struct {
	Commands []Command `json:"commands"`
}

func SaveConfig(config *Commands, fileName string) error {
	if fileName == "" {
		return fmt.Errorf("no filename defined")
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, util.DefaultWritePermissions)
}
func LoadCommands(config *Commands, file string) error {
	if file != "" {
		exists, err := util.FileExists(file)
		if err != nil {
			return fmt.Errorf("could not check if file exists %s due to %s", file, err)
		}
		if !exists {
			return fmt.Errorf("File does " + util.ColorBold(util.ColorError("NOT")) + " exist")
		} else {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				return fmt.Errorf("failed to load file %s due to %s", file, err)
			}
			err = yaml.Unmarshal(data, config)
			if err != nil {
				return fmt.Errorf("failed to unmarshal YAML file %s due to %s", file, err)
			}
		}
	}
	return nil
}

func GetCommands() *Commands {
	var commandsFromFile *Commands
	_ = LoadCommands(commandsFromFile, "commands.yaml")
	return commandsFromFile
}
func (cmd *Commands) GetAvailableCommands() map[string]Command {
	names := make(map[string]Command, 0)
	for _, cmd := range cmd.Commands {
		names[cmd.Name] = cmd
	}
	return names
}
func getAllCommands() {

}
