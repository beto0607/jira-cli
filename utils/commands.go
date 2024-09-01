package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

func ShouldPrintHelp(args []string) bool {
	for _, v := range args {
		if v == "--help" || v == "help" {
			return true
		}
	}

	return false
}

func GetIssueIdFromBranch() (string, error) {
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
