package main

import (
	"fmt"
	"os"

	"jira-cli/commands"
	"jira-cli/configs"
)

func main() {
	argsValid := checkArgs(os.Args)
	if argsValid != 0 {
		os.Exit(argsValid)
	}
	configsValue := configs.LoadConfig()
	code := mainRun(configsValue)
	os.Exit(code)
}

type CommandFunc = func(args []string, configsValues configs.Configs) int

var commandsMap = map[string]CommandFunc{
	"assign":     commands.RunAssignCommand,
	"transition": commands.RunTransitionCommand,
	"--help":     commands.RunHelpCommand,
	"help":       commands.RunHelpCommand,
	"config":     commands.RunConfigCommand,
}

func checkArgs(args []string) int {
	if len(args) < 2 {
		return 1
	}

	return 0
}

func mainRun(configsValues configs.Configs) int {
	expandedArgs := []string{}
	if len(os.Args) > 0 {
		expandedArgs = os.Args[1:]
	}

	currentCommand := expandedArgs[0]
	if commandsMap[currentCommand] == nil {
		fmt.Fprintf(os.Stderr, `Unkown command "%s"\n`, currentCommand)
		return 1
	}

	r := commandsMap[currentCommand](expandedArgs, configsValues)
	return r
}
