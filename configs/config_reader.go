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

func LoadConfig() Configs {
	rawConfigs := getRawConfigs()
	configs := convertMapToConfigs(rawConfigs)
	return *configs
}

func GetRawValue(section string, settingName string) (value string, found bool) {
	rawConfigs := getRawConfigs()

	sectionMap := (*rawConfigs)[section]
	if sectionMap != nil {
		rawValue, ok := sectionMap[settingName]
		return rawValue, ok
	}
	return "", false
}

func getRawConfigs() *RawConfigs {
	baseDir := getBaseDir()
	configFilePath := baseDir + defaultPath

	rawConfigs, err := loadConfigFile(configFilePath)
	if err != nil {
		log.Panic("Error while loading configs")
	}
	return rawConfigs

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

func loadConfigFile(filePath string) (*RawConfigs, error) {
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

	return configMap, nil
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

	r, err := regexp.Compile(`\[(?P<section>\w+)\]`)
	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// ignore commentes or empty lines
		if len(trimmedLine) == 0 || trimmedLine[0] == '#' {
			continue
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
