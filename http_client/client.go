package http_client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient() (client *Client) {

	timeout := 60 * time.Second
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
	return &Client{
		client: &httpClient,
	}
}

func (c *Client) Request(context context.Context, url string, method string, params interface{}) (result []byte, err error) {
	data, _ := json.Marshal(params)
	req, err := http.NewRequestWithContext(context, method, url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("NewRequestWithContext error: %s", err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response body: ", string(body))
	return body, err
}