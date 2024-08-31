package commands

import (
	"fmt"
	"jira-cli/configs"
)

func DoMoveCommand(args []string, configsValus *configs.Configs) int {
	issueId := args[1]
	targetStatus := args[2]

	fmt.Printf("Jira ticket: %s\n", issueId)
	fmt.Printf("Transition target: %s\n", targetStatus)
	return 0
}
