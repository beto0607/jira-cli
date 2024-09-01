package commands

import (
	"jira-cli/configs"
	"jira-cli/http"
)

func DoMoveCommand(args []string, configsValues configs.Configs) int {
	issueId := args[1]
	targetStatus := args[2]

	if http.RequestTransitionTo(configsValues, issueId, targetStatus) {
		return 0
	}
	return 3
}
