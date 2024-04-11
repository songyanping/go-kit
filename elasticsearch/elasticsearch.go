package elasticsearch

import (
	"fmt"
	es8 "github.com/elastic/go-elasticsearch/v8"
)

func NewEsConfig(host, port, username, password, protocol string) EsConfig {
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
	Port     string
	Username string
	Password string
	Protocol string
}

func NewSearchClient(config EsConfig) (*Client, error) {
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
	return &Client{
		client: client,
	}, nil
}

type Client struct {
	client *es8.Client
}
