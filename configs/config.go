package configs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type RawConfigs map[string]map[string]string

type Configs struct {
	auth AuthConfig
	user UserConfig
}

type AuthConfig struct {
	token string
}
type UserConfig struct {
	email     string
	accountId string
}

const defaultPath = "/jira-cli/config.conf"

func LoadConfig() *Configs {
	baseDir := getBaseDir()
	configFilePath := baseDir + defaultPath
	log.Print(configFilePath)

	configs, err := loadConfigFile(configFilePath)
	if err != nil {
		log.Panic("Error while loading configs")
	}

	return configs
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
		auth: AuthConfig{},
		user: UserConfig{},
	}

	configs.auth.token = (*configsMap)["auth"]["token"]
	configs.user.email = (*configsMap)["user"]["email"]
	configs.user.accountId = (*configsMap)["user"]["accountId"]

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
			fmt.Fprintf(os.Stderr, "Could not compile regex: %s", err.Error())
			return nil, err
		}
		groups := r.FindStringSubmatch(trimmedLine)

		if len(groups) == 2 {
			currentGroup = strings.ToLower(groups[1])

			// ignore unkown sections
			if currentGroup != "auth" && currentGroup != "user" {
				continue
			}

			fmt.Println(currentGroup)
			configMap[currentGroup] = map[string]string{}
			continue
		}

		values := strings.Split(trimmedLine, "=")
		if len(values) < 2 {
			continue
		}
		optionKey := strings.TrimSpace(values[0])
		optionValue := strings.TrimSpace(strings.Join(values[1:], "="))
		configMap[currentGroup][optionKey] = strings.Trim(optionValue, `"`)
	}

	return &configMap, nil
}
