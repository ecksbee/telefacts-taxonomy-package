package cache

import (
	"encoding/json"
	"sync"
)

func MarshalNamespace(namespace string) ([]byte, error) {
	//todo query namespacerepo
	data := Page{
		PageIndicator: 0,
		Items: []struct {
			Display string
			Link    string
		}{
			struct {
				Display string
				Link    string
			}{
				Display: "test",
			},
		},
	}
	return json.Marshal(&data)
}

type NamespaceRepo struct {
	lock sync.RWMutex
}

func NewNamespaceRepo() (*NamespaceRepo, error) {
	//todo inspect namespaces from discovered base taxonomies
	return nil, nil
}
