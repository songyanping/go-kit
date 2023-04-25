package http_client

import (
	"context"
	"testing"
)

func TestHttpClient(t *testing.T) {
	c := NewClient()
	cxt := context.Background()
	c.Request(cxt, "https://baiduc.com", "GET", nil)
}
