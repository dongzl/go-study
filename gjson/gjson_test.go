package gjson

import (
	"fmt"
	"github.com/tidwall/gjson"
	"testing"
)

// https://darjun.github.io/2020/03/22/godailylib/gjson/

func TestName(t *testing.T) {
	json := `{
    "metadata":"metadata",
    "data.listeners":"listeners",
    "data.filters":"filters",
    "data.clusters":"dataSourceClusters",
    "data.sharding_rule":"shardingRule",
    "data.tenants":"tenants"
}
`
	metadata := gjson.Get(json, "metadata")
	fmt.Println("metadata:", metadata.String())

	filter := gjson.Get(json, `data\.filters`)
	fmt.Println("filters:", filter)
}
