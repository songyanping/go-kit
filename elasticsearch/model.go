package elasticsearch

type Alert struct {
	Uuid  string `json:"uuid"`
	Time  string `json:"time"`
	Name  string `json:"name"`
	Event Event  `json:"event"`
}

type Event struct {
	Time    string                 `json:"time"`
	Service string                 `json:"service"`
	Title   string                 `json:"title"`
	Details map[string]interface{} `json:"details"`
	Labels  map[string]string      `json:"labels"`
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
