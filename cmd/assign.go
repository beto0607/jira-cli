package cmd

import (
	"errors"
	"fmt"
	"strings"

	"jira-cli/configs"
	"jira-cli/http"
	"jira-cli/models"
	"jira-cli/utils"

	"github.com/spf13/cobra"
)

var (
	assignCmd = &cobra.Command{
		Use:   "assign [issue ID] [account ID]",
		Short: "Assign issues to someone",
		Args:  cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var issueId string
			var targetAccountId string

			gitBranch, _ := cmd.Flags().GetBool("git-branch")
			if gitBranch {
				issueIdFromBranch, err := utils.GetIssueIdFromBranch()
				if err != nil {
					return err
				}
				issueId = issueIdFromBranch
			} else {
				if len(args) == 0 {
					return errors.New("Issue ID not provided")
				}
				issueId = args[0]
				args = args[1:]
			}

			assignToSelf, _ := cmd.Flags().GetBool("me")
			unassign, _ := cmd.Flags().GetBool("no-one")
			search, _ := cmd.Flags().GetBool("search")
			if unassign {
				targetAccountId = ""
			} else if search {
				selectedAssignee, err := promptAssignee(config, issueId)
				if err != nil {
					return err
				}
				targetAccountId = "" // by default, unassign issue
				if selectedAssignee != nil {
					targetAccountId = selectedAssignee.AccountId
				}
			} else if assignToSelf == true {
				targetAccountId = config.User.AccountId
			} else {
				if len(args) == 0 {
					return errors.New("AccountID not provided")
				}
				targetAccountId = args[0]
			}

			_, err := http.RequestChangeAssignee(config, issueId, targetAccountId)
			if err != nil {
				return err
			}
			fmt.Println("oky doky")
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(assignCmd)

	branchUnderlined := utils.MakeUndeline(`JIRA-1234-my-awesome-branch`)
	issueIdUnderlined := utils.MakeUndeline("JIRA-1234")
	assignCmd.Flags().BoolP("git-branch", "g", false, "Tries to get the issueId from the current git branch. i.e.: "+branchUnderlined+" it will try to get "+issueIdUnderlined)

	assignCmd.Flags().Bool("me", false, "Assigns ticket to Account ID found in config file")
	assignCmd.Flags().Bool("no-one", false, "Unassigns ticket")
	assignCmd.Flags().BoolP("search", "s", false, "Allows to query and select the desired user")
}

// move this?
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
		var selectedIndex int
		if configsValues.Fzf.Enabled {
			selectedIndex, _, err = utils.FzfSelect(options)
			if err != nil {
				return nil, err
			}
		} else {
			selectedIndex, _ = utils.Select(options)
		}
		if selectedIndex <= 0 {
			// Search again
			continue
		}

		return &listAssignableUsers[selectedIndex-1], nil
	}
}
