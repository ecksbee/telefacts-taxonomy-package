package cache

import (
	"encoding/json"
)

func MarshalNamespace(namespace string) ([]byte, error) {
	//todo query namespacerepo
	data := Page{
		PageIndicator: 0,
		Items: []struct {
			Display string
			Link    string
		}{},
	}
	return json.Marshal(&data)
}

type NamespaceRepo struct {
}

func NewNamespaceRepo() (*NamespaceRepo, error) {
	//todo inspect namespaces from discovered base taxonomies
	return nil, nil
}
