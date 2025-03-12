package commands

import (
	"fmt"
	"jira-cli/configs"
	"jira-cli/utils"
)

func RunHelpCommand(args []string, configsValues configs.Configs) int {
	fmt.Println(utils.MakeBold("Commands:"))
	fmt.Println("\tjira-cli assign - Assign ticket to someone")
	fmt.Println("\tjira-cli config - Access configuration for CLI tool")
	fmt.Println("\tjira-cli current - Ticket(s) currently working on")
	fmt.Println("\tjira-cli navigate - Open ticket in browser")
	fmt.Println("\tjira-cli transition - transition ticket")
	return 0
}
