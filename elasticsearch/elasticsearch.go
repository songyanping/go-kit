package elasticsearch

import (
	"fmt"
	es8 "github.com/elastic/go-elasticsearch/v8"
	"strings"
	"time"
)

func NewEsConfig(host string, port int, username, password, protocol string) EsConfig {
	return EsConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Protocol: protocol,
	}
}

type EsConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Protocol string
}

func NewSearchClient(config EsConfig) (*EsClient, error) {
	cfg := es8.Config{
		Addresses: []string{
			fmt.Sprintf("%s://%s:%d", config.Protocol, config.Host, config.Port),
		},
		Username: config.Username,
		Password: config.Password,
	}
	client, err := es8.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("Error creating the client: %s", err)
	}

	return &EsClient{
		client: client,
	}, nil
}

type EsClient struct {
	client *es8.Client
}

func (es *EsClient) SearchContent(index, query string) (result string, err error) {
	res, err := es.client.Search(
		es.client.Search.WithIndex(index),
		es.client.Search.WithBody(strings.NewReader(query)),
		es.client.Search.WithPretty(),
		es.client.Search.WithTimeout(60*time.Second),
	)
	if err != nil {
		fmt.Errorf("Error connect search: %s", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Errorf("[%s] Error getting response: %s", res.Status(), res.String())
		return
	}
	result = res.String()
	return
}