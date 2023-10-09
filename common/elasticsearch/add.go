package elasticsearch

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

var EsConn *elasticsearch.Client

func init() {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
	}

	EsConn, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println("ElasticSearch初始化链接失败")
		return
	}
}

// AddDocumentToES 添加数据至es
func AddDocumentToES(index string, documentID string, sta interface{}) error {
	// 序列化文档数据
	marshal, err := json.Marshal(sta)
	if err != nil {
		return err
	}
	// 构建索引请求
	req := esapi.IndexRequest{
		Index:      index,      //索引
		DocumentID: documentID, //文档ID
		Body:       strings.NewReader(string(marshal)),
	}

	// 发送索引请求
	res, err := req.Do(context.Background(), EsConn)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
