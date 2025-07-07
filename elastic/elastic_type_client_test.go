package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

/*
*
http://127.0.0.1:5601/app/dev_tools#/console

GET /review-1/_search

	{
	  "query": {
	    "bool": {
	      "should":[
	        {"term":{"userID": 147982602}},
	        {"term":{"score": 6}}
	      ]
	    }
	  }
	}

POST /review-1/_create/2

	{
	    "id":1,
	    "userID":147982602,
	    "score":7,
	    "status":2,
	    "publishTime":"2023-09-09T16:07:42.499144+08:00",
	    "content":"差评",
	    "tags":[
	        {
	            "code":1000,
	            "title":"好评"
	        },
	        {
	            "code":2000,
	            "title":"物超所值"
	        },
	        {
	            "code":3000,
	            "title":"有图"
	        }
	    ]
	}
*/
func Test_elastic_search_Typed_Client_get(t *testing.T) {
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	// 创建客户端连接
	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		fmt.Printf("elasticsearch.NewClient failed, err:%v\n", err)
		return
	}

	resp, err := client.Get("review-1", "1").Do(context.Background())
	fmt.Printf("fileds:%s\n", resp.Source_)
}

func Test_elastic_search_Typed_Client_search(t *testing.T) {
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	// 创建客户端连接
	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		fmt.Printf("elasticsearch.NewClient failed, err:%v\n", err)
		return
	}

	query := &types.Query{}

	indexMap := map[string]interface{}{
		"userID": 147982602,
		"score":  3,
	}

	var termQuery []types.Query
	for k, v := range indexMap {
		termQuery = append(termQuery, types.Query{Term: map[string]types.TermQuery{k: {Value: v}}})
	}
	query.Bool = &types.BoolQuery{
		Should: termQuery,
	}

	res, err := client.Search().Index("review-1").Size(10).Query(query).Do(context.Background())

	fmt.Printf("fileds res:%v\n", res)

	for _, hit := range res.Hits.Hits {
		var doc map[string]interface{}
		if err := json.Unmarshal(hit.Source_, &doc); err != nil {
			fmt.Printf("elasticsearch.NewClient json, err:%v\n", err)
		}
		fmt.Printf("fileds hit:%v\n", doc)
	}
}
