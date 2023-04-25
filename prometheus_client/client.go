package prometheus_client

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Client struct {
	client *api.Client
}

func NewClient(url string) (client *Client) {
	// skip cert
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	config := api.Config{
		Address:      url,
		RoundTripper: tr,
	}
	prometheusClient, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		client: &prometheusClient,
	}

}

func (c *Client) Query(ctx context.Context, endTime int64, query string) (value model.Value) {
	end := time.Unix(0, endTime*int64(time.Millisecond)).UTC()
	v1api := v1.NewAPI(*c.client)
	result, warnings, err := v1api.Query(ctx, query, end, v1.WithTimeout(5*time.Second))
	if err != nil {
		log.Errorf("Error querying Prometheus: %v\n", err)
		return
	}
	if len(warnings) > 0 {
		log.Errorf("Warnings: %v\n", warnings)
	}
	return result
}

func (c *Client) QueryRange(ctx context.Context, endTime int64, query string) (value model.Value) {

	end := time.Unix(0, endTime*int64(time.Millisecond)).UTC()
	r := v1.Range{
		Start: end.Add(-4 * time.Minute),
		End:   end,
		Step:  time.Minute,
	}
	v1api := v1.NewAPI(*c.client)
	result, warnings, err := v1api.QueryRange(ctx, query, r, v1.WithTimeout(5*time.Second))
	if err != nil {
		log.Errorf("Error querying Prometheus: %v\n", err)
	}
	if len(warnings) > 0 {
		log.Errorf("Warnings: %v\n", warnings)
	}
	return result
}

func (c *Client) GetApiMetricDetails(value model.Value) (result []ApiMetric) {
	var values []ApiMetric
	switch value.Type() {
	case model.ValNone:
		fmt.Println("None Type")
	case model.ValScalar:
		fmt.Println("Scalar Type")
		v, _ := value.(*model.Scalar)
		displayScalar(v)
	case model.ValVector:
		fmt.Println("Vector Type")
		v, _ := value.(model.Vector)
		values = getApiMetricDetailsForVector(v)
		break
	case model.ValMatrix:
		fmt.Println("Matrix Type")
		v, _ := value.(model.Matrix)
		displayMatrix(v)
	case model.ValString:
		fmt.Println("String Type")
		v, _ := value.(*model.String)
		displayString(v)
	default:
		fmt.Printf("Unknow Type")
	}
	return values
}

func getApiMetricDetailsForVector(v model.Vector) (result []ApiMetric) {
	var values []ApiMetric
	for _, i := range v {
		fmt.Printf("%s %s %s\n", i.Metric.String(), i.Value.String(), i.Timestamp.String())
		var apiMetric ApiMetric
		for k, v := range i.Metric {
			switch k {
			case "channel":
				apiMetric.Channel = string(v)
				continue
			case "response_code":
				apiMetric.ResponseCode = string(v)
				continue
			case "url_path":
				apiMetric.UrlPath = string(v)
				continue
			}
		}
		apiMetric.Value = i.Value.String()
		values = append(values, apiMetric)
	}
	return values
}
