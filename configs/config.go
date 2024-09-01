package configs

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type RawConfigs map[string]map[string]string

type Configs struct {
	Auth AuthConfig
	User UserConfig
	Jira JiraConfig
}

type AuthConfig struct {
	Token string
}
type UserConfig struct {
	Email     string
	AccountId string
}
type JiraConfig struct {
	Organization string
}

const defaultPath = "/jira-cli/config.conf"

func LoadConfig() Configs {
	baseDir := getBaseDir()
	configFilePath := baseDir + defaultPath

	configs, err := loadConfigFile(configFilePath)
	if err != nil {
		log.Panic("Error while loading configs")
	}

	return *configs
}

func getBaseDir() string {
	baseDir := os.Getenv("XDG_CONFIG_HOME")
	if len(baseDir) == 0 {
		baseDir = "~/.config"
	}
	// trim trailing slash if there
	if baseDir[len(baseDir)-1] == '/' {
		baseDir = baseDir[:len(baseDir)-1]
	}
	return baseDir
}

func loadConfigFile(filePath string) (*Configs, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open file: %s", err.Error())
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	configMap, err := parseConfigFile(scanner)

	if err != nil {
		return nil, err
	}

	configs := convertMapToConfigs(configMap)

	return configs, nil
}

func convertMapToConfigs(configsMap *RawConfigs) *Configs {
	configs := Configs{
		Auth: AuthConfig{},
		User: UserConfig{},
		Jira: JiraConfig{},
	}

	configs.Auth.Token = (*configsMap)["auth"]["token"]

	configs.User.Email = (*configsMap)["user"]["email"]
	configs.User.AccountId = (*configsMap)["user"]["accountId"]

	configs.Jira.Organization = (*configsMap)["jira"]["organization"]

	return &configs
}

func parseConfigFile(scanner *bufio.Scanner) (*RawConfigs, error) {
	configMap := make(RawConfigs)

	var currentGroup string = ""

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// ignore commentes or empty lines
		if len(trimmedLine) == 0 || trimmedLine[0] == '#' {
			continue
		}
		r, err := regexp.Compile(`\[(?P<section>\w+)\]`)
		if err != nil {
			return nil, err
		}
		groups := r.FindStringSubmatch(trimmedLine)

		if len(groups) == 2 {
			currentGroup = strings.ToLower(groups[1])

			configMap[currentGroup] = map[string]string{}
			continue
		}

		values := strings.Split(trimmedLine, "=")
		if len(values) < 2 {
			continue
		}
		optionKey := strings.TrimSpace(values[0])
		optionValue := strings.TrimSpace(strings.Join(values[1:], "="))

		if configMap[currentGroup] == nil {
			return nil, errors.New("There's an error in your config file")
		}

		configMap[currentGroup][optionKey] = strings.Trim(optionValue, `"`)
	}

	return &configMap, nil
}
