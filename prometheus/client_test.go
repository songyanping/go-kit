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
	client := NewClient("https://prometheus-sre.intranet.local")
	//client := NewClient("http://10.158.215.47:8481/select/0/prometheus/")
	cxt := context.Background()
	var ti int64
	ti = 1718784833172
	query := "k8s_webrequest_requestCount{path=\"/api/addCart\",datasource=\"alicloud\", deploymentType=\"base\"}"
	result := client.QueryRange(cxt, ti, 30, 5, query)
	vaules := client.GetMetricsResultByMatrix(result)
	//fmt.Println(vaules)
	for _, v := range vaules {
		fmt.Println(v.Labels)
		fmt.Println(v.Value)
		fmt.Println(v.Timestamp)
	}
}
