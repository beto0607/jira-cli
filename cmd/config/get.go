package config_cmd

import (
	"errors"
	"fmt"
	"jira-cli/configs"
	"strings"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <section>.<setting_name>",
	Short: "Set <value> for <setting_name> property inside <section>",
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.ExactArgs(1)(cmd, args)

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
		settingParts := strings.Split(args[0], ".")
		section := settingParts[0]
		settingName := strings.Join(settingParts[1:], ".")
		rawValue, found := configs.GetRawValue(section, settingName)
		if !found {
			return fmt.Errorf("Couldn't find \"%s\"\n", args[0])
		}
		fmt.Printf("Value for \"%s\" is \"%s\"\n", args[0], rawValue)
		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(getCmd)
}
