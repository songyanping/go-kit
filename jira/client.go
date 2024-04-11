package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func NewJiraConfig(url, username, password string) JiraConfig {
	return JiraConfig{
		Url:      url,
		Username: username,
		Password: password,
	}
}

type JiraConfig struct {
	Url      string
	Username string
	Password string
}

func NewClient(config JiraConfig) (*Client, error) {
	base := config.Url
	tp := jira.BasicAuthTransport{
		Username: config.Username,
		Password: config.Password,
	}
	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		return nil, fmt.Errorf("failed to create Jira client: %w", err)
	}
	return &Client{
		client: jiraClient,
	}, nil
}

type Client struct {
	client *jira.Client
}

func (c *Client) Get(issueId string) (summary string) {
	issue, resp, err := c.client.Issue.Get(issueId, nil)
	if err != nil {
		fmt.Sprintf("get jira issue err,issue id:%s", issueId)
		return
	}
	if resp.StatusCode == 200 {
		summary = issue.Fields.Summary
		fmt.Println(summary)
	}
	return summary
}
