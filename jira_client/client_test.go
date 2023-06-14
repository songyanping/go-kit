package jira_client

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	InitJiraConfig("https://xxxxxxxxx/jira", "xxx", "xxx")
	c := NewClient()
	issueId := "75070"
	c.Get(issueId)
}
