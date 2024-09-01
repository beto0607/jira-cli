package commands

import (
	"errors"
	"fmt"
	"jira-cli/configs"
	"jira-cli/http"
	"jira-cli/models"
	"jira-cli/utils"
	"os"
	"strings"
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
	} else if assigneeOption == "-s" || assigneeOption == "--search" {

		selectedAssignee, err := promptAssignee(configsValues, issueId)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return 1
		}
		assigneeOption = "" // by default, unassign issue
		if selectedAssignee != nil {
			assigneeOption = selectedAssignee.AccountId
		}

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
	fmt.Println("\t\t\t\t<assignee | --search | -s | --me | --no-one>")

	fmt.Println(utils.MakeBold("Flags:"))

	fmt.Println("\t-g, --git-branch")
	fmt.Println("\t  Tries to get the issueId from the current git branch. i.e.:")
	branchUnderlined := utils.MakeUndeline(`JIRA-1234-my-awesome-branch`)
	issueIdUnderlined := utils.MakeUndeline("JIRA-1234")
	fmt.Println("\t  " + branchUnderlined + " it will try to get " + issueIdUnderlined)
	fmt.Println("")

	fmt.Println("\t-s, --search")
	fmt.Println("\t  Allows to query and select the desired user")
	fmt.Println("")
	fmt.Println("\t--me")
	fmt.Println("\t  Assigns ticket to Account ID found in config file")
	fmt.Println("")
	fmt.Println("\t--no-one")
	fmt.Println("\t  Unassigns ticket")
	fmt.Println("")
}

func promptAssignee(configsValues configs.Configs, issueId string) (*models.AssignableUser, error) {
	for {
		query := utils.PromptQuery("Who's the next assignee? (write \"none\" to unassign issue)")

		if strings.ToLower(query) == "none" {
			return nil, nil
		}

		if len(query) == 0 {
			return nil, errors.New("No query provided, canceling")
		}

		listAssignableUsers, err := http.RequestQueryAssignableUser(configsValues, issueId, query)
		if err != nil {
			return nil, err
		}

		options := []string{
			"Search again",
		}

		for _, assignableUser := range listAssignableUsers {
			options = append(options, assignableUser.DisplayName+"-"+assignableUser.EmailAddress)
		}

		selectedIndex, _ := utils.Select(options)
		if selectedIndex == 0 {
			// Search again
			continue
		}

		return &listAssignableUsers[selectedIndex-1], nil
	}
}
