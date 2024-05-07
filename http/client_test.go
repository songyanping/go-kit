package http

import (
	"context"
	"testing"
)

func TestHttpClient(t *testing.T) {
	cxt := context.Background()
	c := NewClient()
	c.RequestWithBody(cxt, "http://xxx.com", "GET", "")
}
