package main

import (
	"fmt"
	"os"
	"strings"

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
	expandedArgs := os.Args[1:]
	currentCommand := expandedArgs[0]

	val, ok := configsValues.Alias[currentCommand]
	if ok {
		originalCmd := strings.Split(val, " ")

		expandedArgs = append(originalCmd, expandedArgs[1:]...)
		currentCommand = expandedArgs[0]
		fmt.Printf("new command %s\n", strings.Join(expandedArgs, " "))
	}

	if commandsMap[currentCommand] == nil {
		fmt.Fprintf(os.Stderr, `Unkown command "%s"\n`, currentCommand)
		return 1
	}

	r := commandsMap[currentCommand](expandedArgs, configsValues)
	return r
}
