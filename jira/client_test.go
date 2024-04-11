package jira

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	config := NewJiraConfig("https://xxxxxxxxx/jira", "xxx", "xxx")
	c, _ := NewClient(config)
	c.Get("100")
}
