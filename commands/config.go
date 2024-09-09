package commands

import (
	"fmt"
	"jira-cli/configs"
	"jira-cli/utils"
	"strings"
)

func RunConfigCommand(args []string, configsValues configs.Configs) int {
	if utils.ShouldPrintHelp(args) {
		printConfigHelp()
		return 0
	}
	arguments := utils.FilterFlags(args)
	if !isValidConfigCommand(arguments) {
		printConfigHelp()
		return 1
	}
	dryRun := utils.IsFlagInArgs(args, "--dry-run")
	operation := arguments[1]
	switch operation {
	case "set":
		return executeSet(arguments[2], arguments[3], configsValues, dryRun)
	case "get":
		return executeGet(arguments[2])
	}
	return 0
}

func executeSet(setting string, value string, configValues configs.Configs, dryRun bool) int {
	settingParts := strings.Split(setting, ".")
	section := settingParts[0]
	settingName := strings.Join(settingParts[1:], ".")
	err := configs.UpdateConfigs(section, settingName, value, dryRun)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	return 0
}

func executeGet(setting string) int {
	settingParts := strings.Split(setting, ".")
	section := settingParts[0]
	settingName := strings.Join(settingParts[1:], ".")
	rawValue, found := configs.GetRawValue(section, settingName)
	if !found {
		fmt.Printf("Couldn't find \"%s\"\n", setting)
		return 1
	}
	fmt.Printf("Value for \"%s\" is \"%s\"\n", setting, rawValue)
	return 0
}

func isValidConfigCommand(args []string) bool {
	operation := args[1]
	if operation == "get" {
		return len(args) > 2
	}
	if operation == "set" {
		return len(args) > 3
	}
	return false
}

func printConfigHelp() {
	fmt.Println("Set/Get configs")

	fmt.Println(utils.MakeBold("Usage:"))

	fmt.Println("\tjira-cli config [set <section>.<setting_name> <value>]")
	fmt.Println("\t\t\t[get <section>.<setting_name>]")

	fmt.Println(utils.MakeBold("Flags:"))
	fmt.Println("\t--dry-run")
	fmt.Println("\t Does not update the file but prints the new values.")

}
