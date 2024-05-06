package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"time"
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

func (c *Client) Create(summary, description string) (err error) {
	var comps []*jira.Component
	comp := jira.Component{
		Name: "运维",
	}
	comps = append(comps, &comp)
	issue := jira.Issue{
		Fields: &jira.IssueFields{
			//Assignee: &go_jira.User{
			//	Name:  config.Config.Jira.Assignee,
			//},
			//Reporter: &go_jira.User{
			//	Name:  config.Config.Jira.Reporter,
			//},
			Description: description,
			Type: jira.IssueType{
				Name: "type",
			},
			Project: jira.Project{
				Key: "project",
			},
			Summary:    summary,
			Components: comps,
			Priority: &jira.Priority{
				Name: "priority",
			},
		},
	}
	iss, resp, err := c.client.Issue.Create(&issue)
	fmt.Println(resp.Body)
	fmt.Println(iss.ID)
	return err
}

type JiraIssue struct {
	IssueId       string    `json:"issue_id" bson:"issue_id"`
	Assignee      string    `json:"assignee" bson:"assignee"`
	Reporter      string    `json:"reporter" bson:"reporter"`
	Summary       string    `json:"summary" bson:"summary"`
	Description   string    `json:"description" bson:"description"`
	ComponentName string    `json:"component_name" bson:"component_name"`
	Project       string    `json:"project" bson:"project"`
	Issuetype     string    `json:"issue_type" bson:"issue_type"`
	Priority      string    `json:"priority" bson:"priority"`
	Created       time.Time `json:"created" bson:"created"`
	Updated       time.Time `json:"updated" bson:"updated"`
}
