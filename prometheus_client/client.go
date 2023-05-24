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
	"strconv"
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

func (c *Client) QueryRange(ctx context.Context, endTime int64, beforeMinute int64, setp int64, query string) (value model.Value) {

	end := time.Unix(0, endTime*int64(time.Millisecond)).UTC()
	before := time.Duration(-(beforeMinute - 1)) * time.Minute // -(5-1)=-4,如查询5分钟之前指标
	r := v1.Range{
		Start: end.Add(before),
		End:   end,
		Step:  time.Duration(setp) * time.Minute,
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

func (c *Client) GetMetricsResultByVector(value model.Value) (result []MetricsModel) {
	var metricsModels []MetricsModel
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
		metricsModels = getMetricsByVector(v)
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
	return metricsModels
}

func (c *Client) GetMetricsResultByMatrix(value model.Value) (result []MetricsMatrixModel) {
	v, ok := value.(model.Matrix)
	if !ok {
		fmt.Errorf("Model Matrix assertion err:%s", value.String())
		return nil
	}

	var metricsMatrixModels []MetricsMatrixModel
	for _, i := range v {
		fmt.Printf("%s %s\n", i.Metric.String(), i.Values)

		labelsMap := make(map[string]string)
		for k, v := range i.Metric {
			labelsMap[string(k)] = string(v)
		}

		var valueList []float64
		for _, j := range i.Values {
			floatValue, _ := strconv.ParseFloat(j.Value.String(), 64)
			valueList = append(valueList, floatValue)
		}

		var metrics MetricsMatrixModel
		metrics.Labels = labelsMap
		metrics.Value = valueList

		metricsMatrixModels = append(metricsMatrixModels, metrics)
	}
	return metricsMatrixModels
}

func getMetricsByVector(v model.Vector) (result []MetricsModel) {
	var metricsModels []MetricsModel
	for _, i := range v {
		fmt.Printf("%s %s %s\n", i.Metric.String(), i.Value.String(), i.Timestamp.String())
		var metricsModel MetricsModel

		labelsMap := make(map[string]string)
		for k, v := range i.Metric {
			labelsMap[string(k)] = string(v)
		}

		floatValue, err := strconv.ParseFloat(i.Value.String(), 64)
		if err != nil {
			log.Errorf("strconv parsefloat err:%s", err.Error())
		}
		metricsModel.Labels = labelsMap
		metricsModel.Value = floatValue
		metricsModel.Timestamp = i.Timestamp.Unix()
		metricsModels = append(metricsModels, metricsModel)
	}
	return metricsModels
}
