package commands

import (
	"fmt"
	"jira-cli/configs"
	"jira-cli/http"
	"jira-cli/utils"
)

func RunNavigateCommand(args []string, configsValues configs.Configs) int {
	if utils.ShouldPrintHelp(args) {
		printNavigateHelp()
		return 0
	}
	if len(args) >= 3 {
		printNavigateHelp()
		return 1
	}

	arguments := utils.FilterFlags(args)

	issueId := ""
	if len(arguments) == 2 {
		issueId = arguments[1]
	} else if utils.IsFlagInArgs(args, "--git-branch") || utils.IsFlagInArgs(args, "-g") {
		issueIdFromBranch, err := utils.GetIssueIdFromBranch()
		if err != nil {
			return 1
		}
		issueId = issueIdFromBranch
	} else {
		for {
			issueId = utils.PromptQuery("Enter issue ID:")
			if len(issueId) > 0 {
				break
			}
		}
	}

	browseUrl := http.GetBrowseUrl(configsValues.Jira.Organization, issueId)
	fmt.Println("Opening: " + browseUrl)

	err := utils.OpenBrowser(browseUrl)

	if err != nil {
		return 1
	}
	return 0
}

func printNavigateHelp() {
	fmt.Println("Open issue on browser")

	fmt.Println(utils.MakeBold("Usage:"))
	fmt.Println("\tjira-cli navigate \t[<issueId | --git-branch | -g | --interactive | -i>")

	fmt.Println(utils.MakeBold("Flags:"))
	fmt.Println("\t-g, --git-branch")
	fmt.Println("\t  Tries to get the issueId from the current git branch. i.e.:")
	branchUnderlined := utils.MakeUndeline(`JIRA-1234-my-awesome-branch`)
	issueIdUnderlined := utils.MakeUndeline("JIRA-1234")
	fmt.Println("\t  " + branchUnderlined + " it will try to get " + issueIdUnderlined)
	fmt.Println("")

	fmt.Println("\t-i, --interactive")
	fmt.Println("\t  Allows the user to write the issue ID")
	fmt.Println("\t  Default, if no issue or flag provided")
	fmt.Println("")
}
