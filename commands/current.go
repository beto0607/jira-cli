package commands

import (
	"fmt"
	"jira-cli/configs"
	"jira-cli/http"
	"jira-cli/utils"
)

func RunCurrentWorkCommand(args []string, configsValues configs.Configs) int {
	if utils.ShouldPrintHelp(args) {
		printCurrentHelp()
		return 0
	}
	if len(args) != 2 {
		printCurrentHelp()
		return 1
	}
	url := http.GetCurrentWorkQueryUrl(configsValues.Jira.Organization, configsValues.User.AccountId, configsValues.JQL.CurrentWork)

	fmt.Println(url)
	return 0
}

func printCurrentHelp() {
	fmt.Println("Current issues someone is working on")

	fmt.Println(utils.MakeBold("Usage:"))

	fmt.Println("\tjira-cli current \t<accountId | --me>")
	// fmt.Println("\tjira-cli current \t<accountId | --search | -s | --me>")

	fmt.Println(utils.MakeBold("Flags:"))

	// fmt.Println("\t-s, --search")
	// fmt.Println("\t  Allows to query and select the desired user")
	// fmt.Println("")
	fmt.Println("\t--me")
	fmt.Println("\t  Ticket assinged to Account ID found in config file")
	fmt.Println("")
}

// func promptCurrentWorkTargetAssignee(configsValues configs.Configs) (*models.AssignableUser, error) {
// 	for {
// 		query := utils.PromptQuery("Who? (write \"none\" to unassign issue)")
//
// 		if strings.ToLower(query) == "none" {
// 			return nil, nil
// 		}
//
// 		if len(query) == 0 {
// 			return nil, errors.New("No query provided, canceling")
// 		}
//
// 		listAssignableUsers, err := http.RequestQueryAssignableUser(configsValues, issueId, query)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		options := []string{
// 			"Search again",
// 		}
//
// 		for _, assignableUser := range listAssignableUsers {
// 			options = append(options, assignableUser.DisplayName+"-"+assignableUser.EmailAddress)
// 		}
// 		var selectedIndex int
// 		if configsValues.Fzf.Enabled {
// 			selectedIndex, _, err = utils.FzfSelect(options)
// 			if err != nil {
// 				return nil, err
// 			}
// 		} else {
// 			selectedIndex, _ = utils.Select(options)
// 		}
// 		if selectedIndex <= 0 {
// 			// Search again
// 			continue
// 		}
//
// 		return &listAssignableUsers[selectedIndex-1], nil
// 	}
// }
