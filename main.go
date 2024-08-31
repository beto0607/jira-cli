package main

import (
	"log"
	"os"

	"jira-cli/commands"
	"jira-cli/configs"
)

func main() {
	configs.LoadConfig()
	code := mainRun()
	os.Exit(code)
}

func mainRun() int {
	expandedArgs := []string{}
	if len(os.Args) > 0 {
		expandedArgs = os.Args[1:]
	}

	isHelp := checkIfIsHelpCmd(expandedArgs)
	if isHelp {
		commands.PrintHelp()
		return 0
	}

	log.Println(len(expandedArgs))

	log.Println(expandedArgs)

	return 0
}

func checkIfIsHelpCmd(args []string) bool {
	for _, v := range args {
		if v == "--help" || v == "help" {
			return true
		}
	}
	return false
}
