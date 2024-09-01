package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"jira-cli/configs"
	"jira-cli/models"
	"net/http"
)

func RequestTransitionTo(configsValues configs.Configs, issueId string, transitionId string) (bool, error) {
	url := getTransitionUrl(configsValues.Jira.Organization, issueId)
	data := fmt.Sprintf("{\"transition\":{\"id\":\"%s\"}}", transitionId)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))

	if err != nil {
		return false, err
	}

	authorizationHeader := getAuthorizationToken(configsValues)

	prepareHeaders(authorizationHeader, req)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return false, err
	}

	// Is status 2XX?
	if response.Status[0] == '2' {
		return true, nil
	}
	return false, fmt.Errorf("Jira didn't like that. Returned: %s", response.Status)
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

func RequestChangeAssignee(configsValues configs.Configs, issueId string, accountId string) (bool, error) {
	url := getChangeAssigneeUrl(configsValues.Jira.Organization, issueId)
	data := fmt.Sprintf("{\"accountId\":\"%s\"}", accountId)
	if len(accountId) == 0 { // no accountId means unassign
		data = "{\"accountId\":null}"
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(data)))

	if err != nil {
		return false, err
	}

	authorizationHeader := getAuthorizationToken(configsValues)

	prepareHeaders(authorizationHeader, req)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return false, err
	}

	// Is status 2XX?
	if response.Status[0] == '2' {
		return true, nil
	}
	return false, fmt.Errorf("Jira didn't like that. Returned: %s", response.Status)
}

func RequestQueryAssignableUser(configsValues configs.Configs, issueId string, query string) ([]models.AssignableUser, error) {
	url := getAssignableUserUrl(configsValues.Jira.Organization, issueId, query)
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

	var assignableUserResponse []models.AssignableUser
	err = json.Unmarshal([]byte(body), &assignableUserResponse)
	if err != nil {
		return nil, err
	}

	return assignableUserResponse, nil
}
