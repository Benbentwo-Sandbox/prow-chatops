package main

import (
	"encoding/json"
	actions "github.com/sethvargo/go-githubactions"
	"strings"

	// "github.com/sirupsen/logrus"
	// "k8s.io/test-infra/prow/config"
	"k8s.io/test-infra/prow/github"
	// "k8s.io/test-infra/prow/github/fakegithub"
	// "k8s.io/test-infra/prow/labels"
	// "k8s.io/test-infra/prow/plugins"
)

var (
	issueCommentEvent github.IssueCommentEvent
)

func main() {
	Main()
}

func Main() {
	event := actions.GetInput("githubevent")
	actions.Debugf("Event: %s", event)

	err := json.Unmarshal([]byte(event), &issueCommentEvent)
	if err != nil {
		actions.Fatalf("Couldn't Parse the github event!")
	}

	actions.Warningf("Commenter: %s", issueCommentEvent.Comment.User.Login)

}

func GetCommand(event github.IssueCommentEvent) {
	comment := GetCommentString(event)
	isCommand := IsCommandString(comment)
	if !isCommand {
		return
	}

	checkIsAvailableCommand(comment)
}
func GetCommentString(event github.IssueCommentEvent) string {
	return event.Comment.Body
}

func IsCommandString(comment string) bool {
	str := strings.TrimSpace(comment)
	return strings.HasSuffix(str, "/")
}

func checkIsAvailableCommand(command string) bool {
	return true
}

// func handleGenericCommentEvent(pc plugins.Agent, e github.GenericCommentEvent) error {
// 	cp, err := pc.CommentPruner()
// 	if err != nil {
// 		return err
// 	}
// 	return handleGenericComment(pc.GitHubClient, pc.PluginConfig, pc.OwnersClient, pc.Logger, cp, e)
// }
//
// func handleGenericComment(gc githubClient, config *plugins.Configuration, ownersClient repoowners.Interface, log *logrus.Entry, cp commentPruner, e github.GenericCommentEvent) error {
// 	rc := reviewCtx{
// 		author:      e.User.Login,
// 		issueAuthor: e.IssueAuthor.Login,
// 		body:        e.Body,
// 		htmlURL:     e.HTMLURL,
// 		repo:        e.Repo,
// 		assignees:   e.Assignees,
// 		number:      e.Number,
// 	}
//
// 	// Only consider open PRs and new comments.
// 	if !e.IsPR || e.IssueState != "open" || e.Action != github.GenericCommentActionCreated {
// 		return nil
// 	}
//
// 	// If we create an "/lgtm" comment, add lgtm if necessary.
// 	// If we create a "/lgtm cancel" comment, remove lgtm if necessary.
// 	wantLGTM := false
// 	if LGTMRe.MatchString(rc.body) {
// 		wantLGTM = true
// 	} else if LGTMCancelRe.MatchString(rc.body) {
// 		wantLGTM = false
// 	} else {
// 		return nil
// 	}
//
// 	// use common handler to do the rest
// 	return handle(wantLGTM, config, ownersClient, rc, gc, log, cp)
// }
