package commands

import (
	"github.com/Benbentwo/utils/log"
	"github.com/magiconair/properties/assert"
	"testing"
)

var (
	// err						error
	CommandsTest Commands
)

func checkErr(err error) {
	if err != nil {
		log.Logger().Errorf("error: %s", err)
	}
}
func loadMockData() {
	cmds, err := LoadCommands("../mock-data/commands.yaml")
	CommandsTest = cmds
	checkErr(err)
}
func TestLoadCommands(t *testing.T) {
	loadMockData()
	assert.Equal(t, len(CommandsTest.Commands), 26)
}
