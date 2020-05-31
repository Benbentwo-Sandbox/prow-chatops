package main

import (
	"encoding/json"
	"github.com/Benbentwo/utils/log"
	"github.com/Benbentwo/utils/util"
	"github.com/magiconair/properties/assert"
	"io/ioutil"
	"os"
	"testing"

	// "github.com/sirupsen/logrus"
	// "k8s.io/test-infra/prow/config"
	"k8s.io/test-infra/prow/github"
	// "k8s.io/test-infra/prow/github/fakegithub"
	// "k8s.io/test-infra/prow/labels"
	// "k8s.io/test-infra/prow/plugins"
)

var (
	level                 = "debug"
	IssueCommentEventTest github.IssueCommentEvent
)

func checkErr(err error) {
	if err != nil {
		log.Logger().Errorf("error: %s", err)
	}
}
func loadMockData() {
	dat, err := ioutil.ReadFile("./mock-data/event.json")
	checkErr(err)
	err = json.Unmarshal(dat, &IssueCommentEventTest)
	checkErr(err)
}
func TestMain(m *testing.M) {

	_ = log.SetLevel(level)
	log.SetOutput(os.Stdout)
	loadMockData()
	log.Logger().Debugf("Commenter: %s", util.ColorInfo(IssueCommentEventTest.Comment.User.Login))
	os.Exit(m.Run())

}

func TestParse(t *testing.T) {
	log.Logger().Debugln("Parse Running")
	// log.Println("Parse Running")
	assert.Equal(t, IssueCommentEventTest.Comment.User.Login, "Benbentwo")
}
