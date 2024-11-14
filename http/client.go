package http

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
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

func (c *Client) RequestWithBody(context context.Context, url string, method string, body string) (result []byte, err error) {
	fmt.Printf("Request parameters: url=%s,method=%s,body=%s\n", url, method, body)
	req, err := http.NewRequestWithContext(context, method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Printf("NewRequestWithContext error: %s", err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Printf("Do error: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Io readAll error: %s", err.Error())
		return nil, err
	}
	fmt.Printf("Request response body: %s\n", string(respBody))
	return respBody, err
}
func (c *Client) RequestWithAuth(context context.Context, url string, method string, body string, username string, password string) (result []byte, err error) {
	fmt.Printf("Request parameters: url=%s,method=%s,body=%s\n", url, method, body)
	req, err := http.NewRequestWithContext(context, method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Printf("NewRequestWithContext error: %s", err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(username, password)

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Printf("Do error: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Io readAll error: %s", err.Error())
		return nil, err
	}
	fmt.Printf("Request response body: %s\n", string(respBody))
	return respBody, err
}
