package http

import (
	"encoding/base64"
	"jira-cli/configs"
	"net/http"
	"strings"
)

const jiraRestV3 = "https://{organization}.atlassian.net/rest/api/3"
const transitionSuffix = "/issue/{issueId}/transitions"
const assigneeSuffix = "/issue/{issueId}/assignee"

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

func getChangeAssigneeUrl(organizationName string, issueId string) string {
	url := getBaseUrl(organizationName)
	urlSuffix := strings.Replace(assigneeSuffix, "{issueId}", issueId, 1)
	fullUrl := url + urlSuffix
	return fullUrl
}
