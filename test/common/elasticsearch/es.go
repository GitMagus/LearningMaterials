package elasticsearch

import (
	"2102Amonth/common/database"
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strconv"
)

var esClient *elasticsearch.Client
var err error

func init() {
	// 创建链接
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	esClient, err = elasticsearch.NewClient(cfg)

	if err != nil {
		log.Println("链接es错误")
		log.Println(err)
	}
}

func InsertData(data database.Order) {
	d, _ := json.Marshal(data)
	req := esapi.IndexRequest{
		Index:      "order",
		DocumentID: strconv.Itoa(data.Id),
		Body:       bytes.NewReader(d),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
}
