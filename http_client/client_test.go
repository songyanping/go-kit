package http_client

import (
	"context"
	"testing"
)

func TestHttpClient(t *testing.T) {
	c := NewClient()
	cxt := context.Background()
	c.Request(cxt, "https://amwaychina.codefactori.com/jira/rest/api/2/issue/NEOC-56804", "GET", nil)
}
