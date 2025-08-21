package cmd

import (
	"jira-cli/configs"
	"os"
	"strings"

	config_cmd "jira-cli/cmd/config"

	"github.com/spf13/cobra"
)

var (
	config  configs.Configs
	rootCmd = &cobra.Command{
		Use:   "jira-cli",
		Short: "Manage tickets from your terminal",
	}
)

func Execute() {
	initCofig()

	// Check for aliased commands
	args := os.Args[1:]
	aliasedCmd, ok := config.Alias[args[0]]
	if ok {
		originalCmd := strings.Split(aliasedCmd, " ")
		os.Args = append(append([]string{""}, originalCmd...), args[1:]...)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(config_cmd.ConfigCmd)
}

func initCofig() {
	config = configs.LoadConfig()
}
