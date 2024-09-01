package commands

import (
	"errors"
	"fmt"
	"jira-cli/configs"
	"jira-cli/http"
	"jira-cli/models"
	"jira-cli/utils"
	"os"
)

func RunTransitionCommand(args []string, configsValues configs.Configs) int {
	if utils.ShouldPrintHelp(args) {
		printTransitionsHelp()
		return 0
	}
	if len(args) != 3 {
		printTransitionsHelp()
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

	targetOption := args[2]

	if targetOption == "-s" || targetOption == "--search-target" {
		transitions, err := http.RequestTransitionsList(configsValues, issueId)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return 1
		}

		selectedTransition, err := promptTransition(transitions)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return 1
		}
		targetOption = selectedTransition.Id
	}

	_, err := http.RequestTransitionTo(configsValues, issueId, targetOption)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}
	fmt.Println("oky doky")
	return 0
}

func printTransitionsHelp() {
	fmt.Println("Transition issues (aka, move tickets around)")

	fmt.Println(utils.MakeBold("Usage:"))

	fmt.Println("\tjira-cli transition \t<issueId | --git-branch | -g>")
	fmt.Println("\t\t\t\t<targetTransitionId | --search-target | -s>")

	fmt.Println(utils.MakeBold("Flags:"))

	fmt.Println("\t-g, --git-branch")
	fmt.Println("\t  Tries to get the issueId from the current git branch. i.e.:")
	branchUnderlined := utils.MakeUndeline(`JIRA-1234-my-awesome-branch`)
	issueIdUnderlined := utils.MakeUndeline("JIRA-1234")
	fmt.Println("\t  " + branchUnderlined + " it will try to get " + issueIdUnderlined)
	fmt.Println("")

	fmt.Println("\t-s, --search-target")
	fmt.Println("\t  Makes a request to the Jira API to get all posible")
	fmt.Println("\t  transitions and allows you to select one of them.")
	fmt.Println("")
}

func promptTransition(transitions *models.ListTransitionsResponse) (*models.Transition, error) {
	options := []string{}

	if len(transitions.Transitions) == 0 {
		return nil, errors.New("No valid transitions")
	}

	for _, transition := range transitions.Transitions {
		options = append(options, transition.Name)
	}

	selectedIndex, _ := utils.Prompt(options)

	return &transitions.Transitions[selectedIndex], nil
}
