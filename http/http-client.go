package http

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"jira-cli/configs"
	"log"
	"net/http"
	"strings"
)

const jiraRestV3 = "https://{organization}.atlassian.net/rest/api/3"
const transitionSuffix = "/issue/{issueId}/transitions"

// Transition options
// RFT_ID="91"
// RFD_ID="131"
// RFA_ID="71"
// PAUSED_ID="101"
// BLOCKED_ID="61"
// TO_DO_ID="81"
// IN_PROGRESS_ID="31"

func RequestTransitionTo(configsValues configs.Configs, issueId string, transitionId string) bool {
	organizationName := configsValues.Jira.Organization

	url := strings.Replace(jiraRestV3, "{organization}", organizationName, 1)
	urlSuffix := strings.Replace(transitionSuffix, "{issueId}", issueId, 1)
	fullUrl := url + urlSuffix

	data := fmt.Sprintf("{\"transition\":{\"id\":\"%s\"}}", transitionId)

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Panicln("Could not create request")
	}

	authHeader := getAuthorizationToken(configsValues)

	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Panicln("An error has occurred while requesting your data")
	}

	// Is status 2XX?
	if response.Status[0] == '2' {
		log.Println("oky doky")
	}
	return true
}

func getAuthorizationToken(configsValues configs.Configs) string {
	user := configsValues.User.Email + ":" + configsValues.Auth.Token
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(user))
}
