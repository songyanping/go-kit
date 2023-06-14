package jira_client

import (
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
)

var (
	Url      string
	Username string
	Password string
)

type Client struct {
	client *jira.Client
}

func InitJiraConfig(url, username, password string) error {
	if url == "" || username == "" || password == "" {
		return errors.New("Input parameter cannot be empty")
	}
	Url = url
	Username = username
	Password = password
	return nil
}

func NewClient() (client *Client) {
	base := Url
	tp := jira.BasicAuthTransport{
		Username: Username,
		Password: Password,
	}
	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		panic(err)
	}
	return &Client{
		client: jiraClient,
	}
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
