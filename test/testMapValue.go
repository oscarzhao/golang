package main

import (
	"fmt"
	"encoding/json"
	
	"k8s.io/kubernetes/pkg/api"
)

type PodListParams struct {
	Namespace      string       `json:"namespace"`
	NamespaceSet   bool         `json:"namespace_set"`
	Phase          api.PodPhase `json:"phase"`
	PhaseSet       bool         `json:"phase_set"`
	ClusterName    string       `json:"cluster_name"`
	ClusterNameSet bool         `json:"cluster_name_set"`
	Name           string       `json:"name"`
	NameSet        bool         `json:"name_set"`
	SelectorSet    bool    `json:"selectorSet"`
	Selector map[string]string `json:"selector"`
}

func main() {
	// getSlice()
	parseJsonMap()
}

// 判断json解析，能否解析深次的map
func parseJsonMap() {
	rawString := `{"namespace":"default"}`
	param := PodListParams{}
	if err := json.Unmarshal([]byte(rawString), &param); err != nil {
		fmt.Printf("error unmarshal: %s\n", err)
	} else {
		fmt.Printf("parse success, data:%#v\n", param)
	}
	
}

// pageSize 's value is ZERO by default
func mapInitialValue() {
	str2int := make(map[string]int)
	pageSize, ok := str2int["page_size"]
	fmt.Printf("pageSize: %d, ok: %v\n", pageSize, ok)
}

// 判断 是否包含尾部
func getSlice() {
	slice := []string{"aaaa", "bbbb", "cccc", "dddd", "eeee", "ffff", "gggg"}
	fmt.Printf("slice[:1]: %v\n", slice[:1])

}
