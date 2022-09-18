package cache

import (
	"encoding/json"
	"sync"

	gocache "github.com/patrickmn/go-cache"
)

func MarshalRelationshipSet(repo *RelationshipSetRepo, namespace string) ([]byte, error) {
	//todo query relationshipsetrepo
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

type RelationshipSetRepo struct {
	lock sync.RWMutex
}

func NewRelationshipSetRepo(cache *gocache.Cache) (*RelationshipSetRepo, error) {
	//todo inspect relationship sets from discovered base taxonomies
	return nil, nil
}
