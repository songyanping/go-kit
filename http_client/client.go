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
	log.Printf("Request parameters: url=%s,method=%s,params=%s", url, method, params)
	data, err := json.Marshal(params)
	if err != nil {
		log.Printf("Params Marshal err:%s", err.Error())
		return nil, err
	}
	req, err := http.NewRequestWithContext(context, method, url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("NewRequestWithContext error: %s", err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		log.Printf("Do error: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Request response body: %s", string(body))
	return body, err
}

func (c *Client) RequestStr(context context.Context, url string, method string, params string) (result []byte, err error) {
	log.Printf("RequestStr parameters: url=%s,method=%s,params=%s", url, method, params)
	req, err := http.NewRequestWithContext(context, method, url, bytes.NewBuffer([]byte(params)))
	if err != nil {
		log.Printf("NewRequestWithContext error: %s", err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		log.Printf("Do error: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Request response body: %s", string(body))
	return body, err
}
func (c *Client) RequestAuth(context context.Context, url string, method string, params interface{}, username string, password string) (result []byte, err error) {
	log.Printf("RequestAuth parameters: url=%s,method=%s,params=%s", url, method, params)
	data, err := json.Marshal(params)
	if err != nil {
		log.Printf("Params Marshal err:%s", err.Error())
		return nil, err
	}
	req, err := http.NewRequestWithContext(context, method, url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("NewRequestWithContext error: %s", err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(username, password)

	resp, err := c.client.Do(req)
	if err != nil {
		log.Printf("Do error: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("RequestAuth response body: %s", string(body))
	return body, err
}
