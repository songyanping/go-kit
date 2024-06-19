package elasticsearch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/songyanping/go-kit/elasticsearch"
	"strings"
	"testing"
	"time"
)

func queryEvents(client *elasticsearch.EsClient, index, query string) (result []elasticsearch.Event, err error) {
	//var qry = `{"query": {"bool": {"filter": [{"term":{"summary_date": "SUMMARYDATE"}}]}},"sort": [{"time": "desc"}]}`
	//qry = strings.Replace(qry, "SUMMARYDATE", query, -1)
	//fmt.Println("Elasticsearch eventSummary query: %s", qry)

	res, err := client.SearchContent(index, query)
	if err != nil {
		fmt.Errorf("Error getting elasticsearch response: %s", err.Error())
		return result, err
	}

	var searchResultEvent elasticsearch.SearchResultEvent
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

func TestElasticsearchEvent(t *testing.T) {
	conf := elasticsearch.NewEsConfig("10.158.215.42", 9200, "", "", "http")
	client, err := elasticsearch.NewSearchClient(conf)
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
	conf := elasticsearch.NewEsConfig("10.158.215.42", 9200, "", "", "http")
	client, err := elasticsearch.NewSearchClient(conf)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	now := time.Now().Local().Format(time.RFC3339)
	event := elasticsearch.Event{
		Time:    now,
		Title:   "test es",
		Service: "demo",
	}
	uuid := uuid.NewString()
	alert := elasticsearch.Alert{
		Uuid:  uuid,
		Time:  now,
		Name:  "alert",
		Event: event,
	}
	body, _ := json.Marshal(alert)
	client.Insert(ctx, "alert-test", uuid, body)
}

func TestElasticsearchUpdate(t *testing.T) {
	conf := elasticsearch.NewEsConfig("10.158.215.42", 9200, "", "", "http")
	client, err := elasticsearch.NewSearchClient(conf)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	now := time.Now().Local().Format(time.RFC3339)
	event := elasticsearch.Event{
		Time:    now,
		Title:   "update test es",
		Service: "demo",
	}

	alert := elasticsearch.Alert{
		Time:  now,
		Name:  "alert",
		Event: event,
	}
	docMap := make(map[string]interface{})
	docMap["doc"] = alert
	body, err := json.Marshal(docMap)
	if err != nil {
		t.Error(err)
	}

	client.Update(ctx, "", "d394b077-1a4b-4d6d-a9cb-5692c1184077", body)
}
