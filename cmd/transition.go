package cmd

import (
	"errors"
	"fmt"
	"jira-cli/http"
	"jira-cli/models"
	"jira-cli/utils"
	"os"

	"github.com/spf13/cobra"
)

var transitionCmd = &cobra.Command{
	Use:   "transition [issue ID] [target ID]",
	Short: "Transition issues (aka, move tickets around)",
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var issueId string
		var targetId string
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

		search, _ := cmd.Flags().GetBool("search")
		if search {
			transitions, err := http.RequestTransitionsList(config, issueId)
			if err != nil {
				return err
			}

			fzfEnabled, _ := cmd.Flags().GetBool("fzf")
			selectedTransition, err := selectTransition(transitions, fzfEnabled || config.Fzf.Enabled)
			if err != nil {
				return err
			}
			if selectedTransition == nil {
				os.Exit(1)
			}

			targetId = selectedTransition.Id
		} else {
			if len(args) == 0 {
				return errors.New("Target transition not provided")
			}
			targetId = args[0]
		}

		_, err := http.RequestTransitionTo(config, issueId, targetId)
		if err != nil {
			return err
		}
		fmt.Println("oky doky")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(transitionCmd)

	branchUnderlined := utils.MakeUndeline(`JIRA-1234-my-awesome-branch`)
	issueIdUnderlined := utils.MakeUndeline("JIRA-1234")

	transitionCmd.Flags().BoolP("git-branch", "g", false, "Tries to get the issueId from the current git branch. i.e.: "+branchUnderlined+" it will try to get "+issueIdUnderlined)
	transitionCmd.Flags().BoolP("search", "s", false, "Makes a request to the Jira API to get all posible transitions and allows you to select one of them.")
	transitionCmd.Flags().Bool("fzf", false, "Use Fzf for search")
}

func selectTransition(transitions *models.ListTransitionsResponse, useFzf bool) (*models.Transition, error) {
	options := []string{}

	if len(transitions.Transitions) == 0 {
		return nil, errors.New("No valid transitions")
	}

	for _, transition := range transitions.Transitions {
		options = append(options, transition.Name+"(Id:"+transition.Id+")")
	}

	if !useFzf {
		selectedIndex, _ := utils.Select(options)
		return &transitions.Transitions[selectedIndex], nil
	}

	selectedIndex, _, err := utils.FzfSelect(options)

	if err != nil {
		return nil, err
	}

	if selectedIndex == -1 {
		return nil, nil
	}

	return &transitions.Transitions[selectedIndex], nil
}
