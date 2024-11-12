package prometheus

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

type MetricDataVector struct {
	Labels    map[string]string
	Value     float64
	Timestamp int64
}

type MetricDataMatrix struct {
	Labels    map[string]string
	Value     []float64
	Timestamp []int64
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

type Client struct {
	client *api.Client
}

func (c *Client) Query(ctx context.Context, endTime int64, query string) (value model.Value) {
	end := time.Unix(0, endTime*int64(time.Millisecond)).UTC()
	v1api := v1.NewAPI(*c.client)
	result, warnings, err := v1api.Query(ctx, query, end, v1.WithTimeout(120*time.Second))
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
	result, warnings, err := v1api.QueryRange(ctx, query, r, v1.WithTimeout(120*time.Second))
	if err != nil {
		log.Errorf("Error querying Prometheus: %v\n", err)
	}
	if len(warnings) > 0 {
		log.Errorf("Warnings: %v\n", warnings)
	}
	return result
}

func (c *Client) GetMetricsResultByVector(value model.Value) (result []MetricDataVector) {
	var metrics []MetricDataVector
	v, ok := value.(model.Vector)
	if !ok {
		log.Errorf("Model Vector assertion err")
		return metrics
	}

	for _, i := range v {
		log.Infof("%s %s %s\n", i.Metric.String(), i.Value.String(), i.Timestamp.String())
		var metric MetricDataVector
		labelsMap := make(map[string]string)
		for k, v := range i.Metric {
			labelsMap[string(k)] = string(v)
		}

		floatValue, err := strconv.ParseFloat(i.Value.String(), 64)
		if err != nil {
			log.Errorf("strconv parsefloat err:%s", err.Error())
		}
		metric.Labels = labelsMap
		metric.Value = floatValue
		metric.Timestamp = i.Timestamp.Unix()
		metrics = append(metrics, metric)
	}
	return metrics
}

func (c *Client) GetMetricsResultByMatrix(value model.Value) (result []MetricDataMatrix) {
	var metrics []MetricDataMatrix
	v, ok := value.(model.Matrix)
	if !ok {
		log.Errorf("Model Matrix assertion err")
		return metrics
	}

	for _, i := range v {
		log.Printf("%s %s\n", i.Metric.String(), i.Values)

		labelsMap := make(map[string]string)
		for k, v := range i.Metric {
			labelsMap[string(k)] = string(v)
		}

		var valueList []float64
		var timestampList []int64
		for _, j := range i.Values {
			floatValue, _ := strconv.ParseFloat(j.Value.String(), 64)
			valueList = append(valueList, floatValue)

			timestampList = append(timestampList, j.Timestamp.Unix())
		}

		var metric MetricDataMatrix
		metric.Labels = labelsMap
		metric.Value = valueList
		metric.Timestamp = timestampList

		metrics = append(metrics, metric)
	}
	return metrics
}
func ShowQryResult(value model.Value) {
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
		displayVector(v)
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
}

func displayScalar(v *model.Scalar) {
	log.Printf("%s %s\n", v.Timestamp.String(), v.Value.String())
}

func displayVector(v model.Vector) {
	for _, i := range v {
		log.Infof(i.Value.String())
		log.Infof("%s %s %s\n", i.Timestamp.String(), i.Metric.String(), i.Value.String())
	}
}

func displayMatrix(v model.Matrix) {
	for _, i := range v {
		log.Infof("%s\n", i.Metric.String())
		for _, j := range i.Values {
			log.Infof("\t%s %s\n", j.Timestamp.String(), j.Value.String())
		}
	}
}

func displayString(v *model.String) {
	log.Infof("%s %s\n", v.Timestamp.String(), v.Value)
}
