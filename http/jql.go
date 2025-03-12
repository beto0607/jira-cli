package http

import "strings"

func getJQLForCurrentIssues(accountId string, query string) string {
	//"jql=assignee = 617fb6db702bd0006a310187 AND status = \"in progress\""
	jql := strings.Replace("jql=assignee = {accountId} AND {query}", "{query}", query, 1)
	return strings.Replace(jql, "{accountId}", accountId, 1)
}
