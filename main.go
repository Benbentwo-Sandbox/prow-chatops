package main

import (
	"encoding/json"
	actions "github.com/sethvargo/go-githubactions"
	// "github.com/sirupsen/logrus"
	// "k8s.io/test-infra/prow/config"
	"k8s.io/test-infra/prow/github"
	// "k8s.io/test-infra/prow/github/fakegithub"
	// "k8s.io/test-infra/prow/labels"
	// "k8s.io/test-infra/prow/plugins"
)

func main() {
	event := actions.GetInput("githubevent")
	actions.Debugf("Event: %s", event)
	var issueCommentEvent github.IssueCommentEvent

	err := json.Unmarshal([]byte(event), &issueCommentEvent)
	if err != nil {
		actions.Fatalf("Couldn't Parse the github event!")
	}

}
