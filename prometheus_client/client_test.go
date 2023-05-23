package prometheus_client

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_Query(t *testing.T) {

	client := NewClient("http://10.158.215.90")
	cxt := context.Background()
	var ti int64
	ti = 1682430985867
	query := "sum_over_time(api_data_by_channel_response_code_path{channel=\"ca\",job=\"api_data_by_channel_response_code_path_ca\", response_code=\"429\"}[4m])"
	result := client.Query(cxt, ti, query)
	vaules := client.GetMetricsResultByVector(result)
	fmt.Println(vaules)
}

func TestClient_QueryRange(t *testing.T) {

	client := NewClient("http://10.158.215.90")
	cxt := context.Background()
	var ti int64
	ti = 1684834766000
	query := "api_data_by_channel_response_code_path{channel=\"ca\",response_code=~\"429\"}"
	result := client.QueryRange(cxt, ti, query)
	vaules := client.GetMetricsResultByMatrix(result)
	fmt.Println(vaules)
}
