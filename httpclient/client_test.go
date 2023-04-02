package httpclient

import "testing"

func TestHttpClient(t *testing.T) {
	c := NewClient()
	c.Request("https://baiduc.com", "GET", nil)
}
