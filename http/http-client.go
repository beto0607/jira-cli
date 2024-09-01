package http

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"jira-cli/configs"
	"jira-cli/models"
	"log"
	"net/http"
	"os"
	"strings"
)

const jiraRestV3 = "https://{organization}.atlassian.net/rest/api/3"
const transitionSuffix = "/issue/{issueId}/transitions"

func RequestTransitionTo(configsValues configs.Configs, issueId string, transitionId string) bool {
	url := getTransitionUrl(configsValues.Jira.Organization, issueId)
	data := fmt.Sprintf("{\"transition\":{\"id\":\"%s\"}}", transitionId)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))

	if err != nil {
		log.Panicln("Could not create request")
	}

	authorizationHeader := getAuthorizationToken(configsValues)

	prepareHeaders(authorizationHeader, req)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Panicln("An error has occurred while requesting your data")
	}

	// Is status 2XX?
	if response.Status[0] == '2' {
		return true
	}
	fmt.Fprintf(os.Stderr, "Jira didn't like that. Returned: %s", response.Status)
	return false
}

func RequestTransitionsList(configsValues configs.Configs, issueId string) (*models.ListTransitionsResponse, error) {
	url := getTransitionUrl(configsValues.Jira.Organization, issueId)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	authorizationHeader := getAuthorizationToken(configsValues)

	prepareHeaders(authorizationHeader, req)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var listTransitionsResponse models.ListTransitionsResponse
	err = json.Unmarshal([]byte(body), &listTransitionsResponse)
	if err != nil {
		return nil, err
	}

	return &listTransitionsResponse, nil
}

func prepareHeaders(authorizationHeader string, req *http.Request) {
	req.Header.Add("Authorization", authorizationHeader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func getAuthorizationToken(configsValues configs.Configs) string {
	user := configsValues.User.Email + ":" + configsValues.Auth.Token
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(user))
}

func getBaseUrl(organizationName string) string {
	return strings.Replace(jiraRestV3, "{organization}", organizationName, 1)
}

func getTransitionUrl(organizationName string, issueId string) string {
	url := getBaseUrl(organizationName)
	urlSuffix := strings.Replace(transitionSuffix, "{issueId}", issueId, 1)
	fullUrl := url + urlSuffix
	return fullUrl
}
