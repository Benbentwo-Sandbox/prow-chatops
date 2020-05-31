package main

import (
	actions "github.com/sethvargo/go-githubactions"
	"os"
	// "github.com/sirupsen/logrus"
	// "k8s.io/test-infra/prow/config"
	// "k8s.io/test-infra/prow/github"
	// "k8s.io/test-infra/prow/github/fakegithub"
	// "k8s.io/test-infra/prow/labels"
	// "k8s.io/test-infra/prow/plugins"
)

func main() {
	githubToken := os.Getenv("GITHUB_TOKEN")

	actions.AddMask(githubToken)
	actions.Warningf("GithubEvent: %s", actions.GetInput("githubevent"))
}
