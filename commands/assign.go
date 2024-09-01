package commands

import (
	"fmt"
	"jira-cli/configs"
	"jira-cli/http"
	"jira-cli/utils"
	"os"
)

func RunAssignCommand(args []string, configsValues configs.Configs) int {
	if utils.ShouldPrintHelp(args) {
		printAssignHelp()
		return 0
	}
	if len(args) != 3 {
		printAssignHelp()
		return 1
	}

	issueId := args[1]
	if issueId == "-g" || issueId == "--git-branch" {
		issueIdFromBranch, err := utils.GetIssueIdFromBranch()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return 1
		}
		issueId = issueIdFromBranch
	}

	assigneeOption := args[2]
	if assigneeOption == "--me" {
		assigneeOption = configsValues.User.AccountId
	} else if assigneeOption == "--no-one" {
		assigneeOption = ""
	}

	_, err := http.RequestChangeAssignee(configsValues, issueId, assigneeOption)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}
	fmt.Println("oky doky")
	return 0
}

func printAssignHelp() {
	fmt.Println("Assign issues to someone")

	fmt.Println(utils.MakeBold("Usage:"))

	fmt.Println("\tjira-cli assign \t<issueId | --git-branch | -g>")
	fmt.Println("\t\t\t\t<assignee | --search-target | -s | --me | --no-one>")

	fmt.Println(utils.MakeBold("Flags:"))

	fmt.Println("\t-g, --git-branch")
	fmt.Println("\t  Tries to get the issueId from the current git branch. i.e.:")
	branchUnderlined := utils.MakeUndeline(`JIRA-1234-my-awesome-branch`)
	issueIdUnderlined := utils.MakeUndeline("JIRA-1234")
	fmt.Println("\t  " + branchUnderlined + " it will try to get " + issueIdUnderlined)
	fmt.Println("")

	// fmt.Println("\t-s, --search-target")
	// fmt.Println("\t  Makes a request to the Jira API to get all posible")
	// fmt.Println("\t  users and allows you to select one of them.")
	// fmt.Println("")
	fmt.Println("\t--me")
	fmt.Println("\t  Assigns ticket to Account ID found in config file")
	fmt.Println("")
	fmt.Println("\t--no-one")
	fmt.Println("\t  Unassigns ticket")
	fmt.Println("")
}
