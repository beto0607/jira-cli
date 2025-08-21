package config_cmd

import (
	"errors"
	"jira-cli/configs"
	"strings"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set <section>.<setting_name> <value>",
	Short: "Set <value> for <setting_name> property inside <section>",
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.ExactArgs(2)(cmd, args)

		if err != nil {
			return err
		}

		settingParts := strings.Split(args[0], ".")
		if len(settingParts) != 2 {
			return errors.New("Malformed path to property. Should be <section>.<setting_name>")
		}

		if len(settingParts[0]) == 0 {
			return errors.New("<section> cannot be empty")
		}

		if len(settingParts[1]) == 0 {
			return errors.New("<setting_name> cannot be empty")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		settingParts := strings.Split(args[0], ".")
		section := settingParts[0]
		settingName := strings.Join(settingParts[1:], ".")
		value := args[1]
		err := configs.UpdateConfigs(section, settingName, value, dryRun)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(setCmd)
	setCmd.Flags().Bool("dry-run", false, "Does not update the file but prints the new values.")
}
