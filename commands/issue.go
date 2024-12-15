package commands

import (
	"fmt"
	"jira-cli/configs"
	"jira-cli/http"
	"jira-cli/utils"
	"os"
)

func RunIssueCommand(args []string, configsValues configs.Configs) int {
	if utils.ShouldPrintHelp(args) {
		printIssueHelp()
		return 0
	}
	if shouldPrintIssueHelp(args) {
		printIssueHelp()
		return 1
	}

	if utils.IsFlagInArgs(args, "-l") || utils.IsFlagInArgs(args, "--list") {
		return runIssueListCommand(args, configsValues)
	}

	arguments := utils.FilterFlags(args)

	var issueId string
	if utils.IsFlagInArgs(args, "-g") || utils.IsFlagInArgs(args, "--git-branch") {
		issueIdFromBranch, err := utils.GetIssueIdFromBranch()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return 1
		}
		issueId = issueIdFromBranch
	} else {
		issueId = arguments[0]
		arguments = arguments[1:]
	}

	_, err := http.RequestGetIssue(configsValues, issueId)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}
	fmt.Println("oky doky")
	return 0
}

func runIssueListCommand(args []string, configsValues configs.Configs) int {
	return 1

}

func printIssueHelp() {}

func shouldPrintIssueHelp(args []string) bool {

	arguments := utils.FilterFlags(args)

	if utils.IsFlagInArgs(args, "-l") || utils.IsFlagInArgs(args, "--list") {
		return len(arguments) == 1
	}

	return len(arguments) == 2
}
