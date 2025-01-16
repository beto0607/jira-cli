package http_test

import (
	"jira-cli/http"
	"testing"
)

func TestGetBrowseUrl(t *testing.T) {
	got := http.GetBrowseUrl("test", "JIRA-1234")
	if got != "https://test.atlassian.net/browse/JIRA-1234" {
		t.Errorf("Got: %s", got)
	}
}
