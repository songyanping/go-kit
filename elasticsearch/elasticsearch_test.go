package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

type Alert struct {
	Time  time.Time
	Name  string
	Event Event
}

type Event struct {
	Time        time.Time              `json:"time"`
	StartsAt    string                 `json:"starts_at"`
	EndsAt      string                 `json:"ends_at"`
	Duration    int64                  `json:"duration"` //持续时间
	EventId     string                 `json:"event_id"`
	Severity    string                 `json:"severity"`
	Status      string                 `json:"status"`
	Service     string                 `json:"service" mapping:"type:keyword"`
	Title       string                 `json:"title" mapping:"type:keyword"`
	Env         string                 `json:"env"`
	AppId       string                 `json:"app_id" mapping:"type:keyword"`
	DataSource  string                 `json:"data_source" mapping:"type:keyword"`
	Path        string                 `json:"path" mapping:"type:keyword"`
	Reason      string                 `json:"reason" yaml:"reason" mapping:"type:keyword"`
	Description string                 `json:"description" yaml:"description" mapping:"type:keyword"`
	RuleType    string                 `json:"rule_type" yaml:"rule_type" mapping:"type:keyword"`
	InstanceId  string                 `json:"instance_id" yaml:"instance_id" mapping:"type:keyword"`
	Details     map[string]interface{} `json:"details" yaml:"details" mapping:"type:object"`
	Labels      map[string]string      `json:"labels" yaml:"labels" mapping:"type:object"`
}

type SearchResultEvent struct {
	Took      int  `json:"took"`
	Timed_out bool `json:"timed_out"`
	Shards    struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		Max_score float64 `json:"max_score"`
		Hits      []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			Id     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source Event   `json:"_source"`
			Sort   []int   `json:"sort"`
		} `json:"hits"`
	} `json:"hits"`
	Summary string `json:"_summary"`
}

func queryEvents(client *EsClient, index, query string) (result []Event, err error) {
	//var qry = `{"query": {"bool": {"filter": [{"term":{"summary_date": "SUMMARYDATE"}}]}},"sort": [{"time": "desc"}]}`
	//qry = strings.Replace(qry, "SUMMARYDATE", query, -1)
	//fmt.Println("Elasticsearch eventSummary query: %s", qry)

	res, err := client.SearchContent(index, query)
	if err != nil {
		fmt.Errorf("Error getting elasticsearch response: %s", err.Error())
		return result, err
	}

	var searchResultEvent SearchResultEvent
	err = json.Unmarshal([]byte(strings.Replace(res, "[200 OK]", "", -1)), &searchResultEvent)
	if err != nil {
		fmt.Sprintf("Error json unmarshal: %s", err.Error())
		return result, err
	}
	for _, hit := range searchResultEvent.Hits.Hits {
		result = append(result, hit.Source)
	}

	return result, nil
}

func TestElasticsearchClient(t *testing.T) {
	Conf := NewEsConfig("10.158.215.42", 9200, "", "", "http")
	client, err := NewSearchClient(Conf)
	if err != nil {
		t.Error(err)
	}

	var qry = `{"query": {"bool": {"filter": [{"term":{"env": "pd"}}]}},"size":10,"sort": [{"time": "desc"}]}`
	events, err := queryEvents(client, "events", qry)
	if err != nil {
		t.Error(err)
	}

	for _, i := range events {
		fmt.Println(i.Service)
	}

}

func TestElasticsearchInsert(t *testing.T) {
	Conf := NewEsConfig("10.158.215.42", 9200, "", "", "http")
	client, err := NewSearchClient(Conf)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	fmt.Println(time.Local)
	fmt.Println(time.Now())
	fmt.Println(time.Now().Local())
	event := Event{
		Time:    time.Now().Local(),
		Title:   "test es",
		Service: "demo",
		Env:     "ft1",
	}

	alert := Alert{
		Time:  time.Now().Local(),
		Name:  "alert",
		Event: event,
	}
	body, _ := json.Marshal(alert)
	client.Insert(ctx, "events", "", body)
}

func TestElasticsearchUpdate(t *testing.T) {
	Conf := NewEsConfig("10.158.215.42", 9200, "", "", "http")
	client, err := NewSearchClient(Conf)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	fmt.Println(time.Local)
	fmt.Println(time.Now())
	fmt.Println(time.Now().Local())
	//rfc := time.Now().Format(time.RFC3339)
	event := Event{
		Time:    time.Now().Local(),
		Title:   "update test es",
		Service: "demo",
		Env:     "ft1",
	}

	alert := Alert{
		Time:  time.Now().Local(),
		Name:  "alert",
		Event: event,
	}
	body, _ := json.Marshal(alert)
	client.Insert(ctx, "events", "", body)
}
