package main

import (
	"fmt"
	"os"

	"jira-cli/commands"
	"jira-cli/configs"
)

func main() {
	configsValue := configs.LoadConfig()
	code := mainRun(configsValue)
	os.Exit(code)
}

func mainRun(configsValue *configs.Configs) int {
	expandedArgs := []string{}
	if len(os.Args) > 0 {
		expandedArgs = os.Args[1:]
	}

	isHelp := checkIfIsHelpCmd(expandedArgs)
	if isHelp {
		commands.PrintHelp()
		return 0
	}
	switch expandedArgs[0] {
	case "assign":
		break
	case "move":
		if !validateMoveCmd(expandedArgs) {
			fmt.Println("'move' expectes 2 arguments: ticket and target")
			return 2
		}
		return commands.DoMoveCommand(expandedArgs, configsValue)
	default:
		fmt.Fprintf(os.Stderr, "Unkown command\n")
		return 1
	}

	return 0
}

func checkIfIsHelpCmd(args []string) bool {
	return args[0] == "help" || args[0] == "--help"
}

func validateMoveCmd(args []string) bool {
	return len(args) == 3
}
