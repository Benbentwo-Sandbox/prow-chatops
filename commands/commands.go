package commands

import (
	"fmt"
	"github.com/Benbentwo/utils/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Command struct {
	Name    string       `json:"name"`
	Run     func() error `json:"run,omitempty"`
	Enabled bool         `json:"enabled,omitempty"`
	Flags   []string     `json:"flags,omitempty"` // like /approve no-issue
	Args    bool         `json:"args,omitempty"`  // like /label @foo @bar ...
}

type Commands struct {
	Commands []Command `json:"commands"`
}

func LoadCommands(file string) (Commands, error) {
	var config Commands
	if file != "" {
		exists, err := util.FileExists(file)
		if err != nil {
			return config, fmt.Errorf("Could not check if file exists %s due to %s", file, err)
		}
		if !exists {
			return config, fmt.Errorf("File does " + util.ColorBold(util.ColorError("NOT")) + " exist")
		} else {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				return config, fmt.Errorf("Failed to load file %s due to %s", file, err)
			}
			err = yaml.Unmarshal(data, config)
			if err != nil {
				return config, fmt.Errorf("Failed to unmarshal YAML file %s due to %s", file, err)
			}
		}
	}
	return config, nil
}
func getAvailableCommands() {

}
func getAllCommands() {

}
