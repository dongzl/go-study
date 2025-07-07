package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func Test_elastic_new_client_get(t *testing.T) {
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	// 创建客户端连接
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("elasticsearch.NewClient failed, err:%v\n", err)
		return
	}

	resp, err := client.Get("review-1", "1")
	if err != nil {
		fmt.Errorf("error getting mapping: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error...")
		}
	}(resp.Body)
	//fmt.Printf("fileds:%s\n", resp)
	if resp.IsError() {
		fmt.Errorf("error response: %s", resp.String())
	}

	// Read and parse the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Errorf("failed to make request to %s", err)
	}
	fmt.Printf("fileds:%s\n", result["_source"].(map[string]interface{})["content"])
}

func Test_elastic_new_client_search(t *testing.T) {
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	// 创建客户端连接
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("elasticsearch.NewClient failed, err:%v\n", err)
		return
	}

	resp, err := client.Search(client.Search.WithIndex("review-1"), client.Search.WithBody(strings.NewReader(`{"query":{"match_phrase":{"content": "差评"}}}`)))
	if err != nil {
		fmt.Errorf("error getting mapping: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error...")
		}
	}(resp.Body)
	//fmt.Printf("fileds:%s\n", resp)
	if resp.IsError() {
		fmt.Errorf("error response: %s", resp.String())
	}

	// Read and parse the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Errorf("failed to make request to %s", err)
	}
	fmt.Printf("fileds:%s\n", result["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})["content"])
}

func Test_elastic_new_client_get_index(t *testing.T) {
	es, err := createEsDefaultClient()
	if err != nil {
		fmt.Errorf("failed to create Elasticsearch client: %w", err)
	}

	// Build the request to get the mappings
	req := esapi.IndicesGetMappingRequest{
		Index: []string{"review-1"}, // Specify the index name
	}

	// Execute the request
	res, err := req.Do(context.Background(), es)
	if err != nil {
		fmt.Errorf("error getting mapping: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Errorf("error closing body")
		}
	}(res.Body)

	// Check if the response is successful
	if res.IsError() {
		fmt.Errorf("error response: %s", res.String())
	}

	// Read and parse the response
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		fmt.Errorf("failed to make request to %s", err)
	}

	fmt.Errorf("error getting mapping: %v", result)

	// Call method to extract field names
	indexFields := extractFields(result)

	fmt.Errorf("error getting mapping: %v", indexFields)
}

func Test_elastic_new_client_get_index_2(t *testing.T) {
	es, err := createEsDefaultClient()
	if err != nil {
		fmt.Errorf("failed to create Elasticsearch client: %w", err)
	}

	res, err := es.Indices.GetMapping().Index("review-1").Do(context.Background())

	if err != nil {
		fmt.Errorf("error getting mapping: %s", err)
	}
	fmt.Errorf("error getting mapping: %v", res)
}

// getFieldNames recursively extracts field names under the "fields" key
func getFieldNames(properties map[string]interface{}, prefix string) []string {
	fieldNames := []string{}

	for fieldName, fieldValue := range properties {
		// Generate the full path for the current field
		fullName := fieldName
		if prefix != "" {
			fullName = prefix + "." + fieldName
		}

		// Check if the field contains a "fields" node
		if fieldMap, ok := fieldValue.(map[string]interface{}); ok {
			if fields, ok := fieldMap["fields"].(map[string]interface{}); ok {
				// If there is a "fields" node, iterate through its fields and add to the result
				for subField := range fields {
					fieldNames = append(fieldNames, fullName+"."+subField)
				}
			}

			// If the field contains nested "properties", recursively parse subfields
			if nestedProperties, ok := fieldMap["properties"].(map[string]interface{}); ok {
				fieldNames = append(fieldNames, getFieldNames(nestedProperties, fullName)...)
			}
		}
	}

	return fieldNames
}

// extractFields extracts all field names from a nested map structure
func extractFields(data map[string]interface{}) []string {
	var allFields []string

	// Iterate through each index to get its field names
	for _, indexData := range data {
		if indexMap, ok := indexData.(map[string]interface{}); ok {
			if mappings, ok := indexMap["mappings"].(map[string]interface{}); ok {
				if properties, ok := mappings["properties"].(map[string]interface{}); ok {
					// Get all field names under "fields" and merge into the result
					allFields = append(allFields, getFieldNames(properties, "")...)
				}
			}
		}
	}

	return allFields
}

// createEsDefaultClient configures and creates a new Elasticsearch client
func createEsDefaultClient() (*elasticsearch.TypedClient, error) {
	// Configure the Elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	// Create the client instance
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating the client: %s", err)
	}
	return es, nil
}
