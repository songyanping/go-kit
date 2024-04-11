package prometheus

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

	client := NewClient("https://prometheus-sre-dev.intranet.local")
	cxt := context.Background()
	var ti int64
	ti = 1712831135000
	query := "k8s_webrequest_requestCount{path=\"/api/addCart\",datasource=\"alicloud\", deploymentType=\"base\"}"
	result := client.QueryRange(cxt, ti, 5, 1, query)
	vaules := client.GetMetricsResultByMatrix(result)
	//fmt.Println(vaules)
	for _, v := range vaules {
		fmt.Println(v.Labels)
		fmt.Println(v.Value)
		fmt.Println(v.Timestamp)
	}
}
