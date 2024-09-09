package configs

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"os"
)

func UpdateConfigs(section string, settingName string, value string, dryRun bool) error {
	baseDir := getBaseDir()
	configFilePath := baseDir + defaultPath

	f, err := os.OpenFile(configFilePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	newContent := ""
	currentGroup := ""
	sectionFound := false

	r, err := regexp.Compile(`\[(?P<section>\w+)\]`)
	if err != nil {
		return err
	}

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// ignore commentes or empty lines
		if len(trimmedLine) == 0 || trimmedLine[0] == '#' {
			newContent += line + "\n"
			continue
		}
		groups := r.FindStringSubmatch(trimmedLine)

		if len(groups) == 2 {
			currentGroup = strings.ToLower(groups[1])
			newContent += line + "\n"
			continue
		}
		if currentGroup != section {
			newContent += line + "\n"
			continue
		}
		sectionFound = true
		values := strings.Split(trimmedLine, "=")
		if len(values) < 1 {
			newContent += line + "\n"
			continue
		}
		optionKey := strings.TrimSpace(values[0])
		if optionKey != settingName {
			newContent += line + "\n"
			continue
		}
		newLine := fmt.Sprintf("    %s = \"%s\"\n", settingName, value)
		newContent += newLine
	}
	// group not found
	if !sectionFound {
		newContent += fmt.Sprintf("[%s]\n", section)
		newLine := fmt.Sprintf("    %s = \"%s\"\n", settingName, value)
		newContent += newLine
	}
	if dryRun {
		fmt.Println(newContent)
		return nil
	}

	_, err = f.WriteAt([]byte(newContent), 0)
	if err != nil {
		return err
	}

	return nil
}
