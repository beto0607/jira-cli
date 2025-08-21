package config_cmd

import (
	"jira-cli/configs"

	"github.com/spf13/cobra"
)

var (
	config    configs.Configs
	ConfigCmd = &cobra.Command{
		Use:   "config",
		Short: "Set/Get configs",
	}
)

func init() {
	cobra.OnInitialize(initCofig)
}

func initCofig() {
	config = configs.LoadConfig()
}
