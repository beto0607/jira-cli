package commands

import (
	"errors"
	"fmt"
	"jira-cli/configs"
	"jira-cli/http"
	"jira-cli/models"
	"jira-cli/utils"
	"os"
	"os/exec"
	"regexp"
)

func RunTransitionCommand(args []string, configsValues configs.Configs) int {
	if shouldPrintHelp(args) {
		printHelp()
		return 0
	}
	if len(args) != 3 {
		printHelp()
		return 1
	}

	issueId := args[1]
	if issueId == "-g" || issueId == "--git-branch" {
		issueIdFromBranch, err := getIssueIdFromBranch()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return 2
		}
		issueId = issueIdFromBranch
	}

	targetOption := args[2]

	if targetOption == "-s" || targetOption == "--search-target" {
		transitions, err := http.RequestTransitionsList(configsValues, issueId)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return 2
		}

		selectedTransition, err := promptTransition(transitions)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return 2
		}
		targetOption = selectedTransition.Id
	}

	if http.RequestTransitionTo(configsValues, issueId, targetOption) {
		fmt.Println("oky doky")
		return 0
	}
	return 3
}

func shouldPrintHelp(args []string) bool {
	for _, v := range args {
		if v == "--help" || v == "help" {
			return true
		}
	}

	return false
}

func printHelp() {
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

func getIssueIdFromBranch() (string, error) {
	branchName, err := getBranchName()
	if err != nil {
		return "", err
	}
	fmt.Println("BranchName: " + branchName)
	issueId, err := parseBranchName(branchName)
	if err != nil {
		return "", err
	}
	fmt.Println("IssueId: " + issueId)
	return issueId, nil
}

func getBranchName() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func parseBranchName(branchName string) (string, error) {
	r, err := regexp.Compile(`(?P<section>[A-Z]{2,6}-[0-9]+)`)
	if err != nil {
		return "", err
	}
	groups := r.FindStringSubmatch(branchName)
	if len(groups) != 2 {
		return "", errors.New("Could not find a Jira ticket in branch (remember it's case sensitive)")
	}
	return groups[1], nil

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
