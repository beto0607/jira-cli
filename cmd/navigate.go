package cmd

import (
	"errors"
	"fmt"
	"jira-cli/http"
	"jira-cli/utils"

	"github.com/spf13/cobra"
)

var (
	navigateCmd = &cobra.Command{
		Use:   "navigate [issue ID]",
		Short: "Open issue on browser",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			browser, _ := cmd.Flags().GetString("browser")
			var issueId string
			gitBranch, _ := cmd.Flags().GetBool("git-branch")
			interactive, _ := cmd.Flags().GetBool("interactive")
			if gitBranch {
				issueIdFromBranch, err := utils.GetIssueIdFromBranch()
				if err != nil {
					return err
				}
				issueId = issueIdFromBranch
			} else if interactive {
				for {
					issueId = utils.PromptQuery("Enter issue ID:")
					if len(issueId) > 0 {
						break
					}
				}
			} else {
				if len(args) == 0 {
					return errors.New("Issue ID not provided")
				}
				issueId = args[0]
				args = args[1:]
			}

			browseUrl := http.GetBrowseUrl(config.Jira.Organization, issueId)
			fmt.Printf("navigate called with %s and browser %s\n", issueId, browser)
			fmt.Printf("url: %s\n", browseUrl)
			err := utils.OpenBrowser(browseUrl, browser)

			if err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(navigateCmd)

	branchUnderlined := utils.MakeUndeline(`JIRA-1234-my-awesome-branch`)
	issueIdUnderlined := utils.MakeUndeline("JIRA-1234")
	navigateCmd.Flags().BoolP("git-branch", "g", false, "Tries to get the issueId from the current git branch. i.e.: "+branchUnderlined+" it will try to get "+issueIdUnderlined)
	navigateCmd.Flags().BoolP("interactive", "i", false, "Allows the user to write the issue ID\nDefault, if no issue or flag provided")
	navigateCmd.Flags().StringP("browser", "b", "", "Open with browser")
}
