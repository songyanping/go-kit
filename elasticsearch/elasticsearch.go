package elasticsearch

import (
	"context"
	"fmt"
	es8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func NewEsConfig(host string, port int64, username, password, protocol string) EsConfig {
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
	Port     int64
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

func (es *EsClient) Insert(ctx context.Context, index string, documentID string, body []byte) (err error) {
	if documentID == "" {
		documentID = uuid.NewString()
	}
	// 创建 Index 请求
	indexReq := esapi.IndexRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       strings.NewReader(string(body)),
		Refresh:    "true",
	}

	// 发送 Index 请求
	indexRes, err := indexReq.Do(ctx, es.client)
	if err != nil {
		log.Error("Error insert document")
		return err
	}
	defer indexRes.Body.Close()
	if indexRes.IsError() {
		log.Errorf("Error indexing document: %s", indexRes.Status())
		return err
	} else {
		log.Println("Document indexed successfully!")
	}
	//log.Println(indexRes.String())
	return nil
}

func (es *EsClient) Update(ctx context.Context, index string, documentID string, body []byte) (err error) {
	// 创建 Index 请求
	indexReq := esapi.UpdateRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       strings.NewReader(string(body)),
		Refresh:    "true",
	}

	// 发送 Index 请求
	indexRes, err := indexReq.Do(ctx, es.client)
	if err != nil {
		log.Error("Error update document")
		return err
	}
	defer indexRes.Body.Close()
	if indexRes.IsError() {
		log.Errorf("Error update document: %s", indexRes.Status())
		return err
	} else {
		log.Println("Document update successfully!")
	}
	//log.Println(indexRes.String())
	return nil
}
