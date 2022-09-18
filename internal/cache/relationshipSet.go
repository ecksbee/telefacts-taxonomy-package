package cache

import (
	"encoding/json"
)

func MarshalRelationshipSet(namespace string) ([]byte, error) {
	//todo query relationshipsetrepo
	data := Page{
		PageIndicator: 0,
		Items: []struct {
			Display string
			Link    string
		}{},
	}
	return json.Marshal(&data)
}

type RelationshipSetRepo struct {
}

func NewRelationshipSetRepo() (*RelationshipSetRepo, error) {
	//todo inspect relationship sets from discovered base taxonomies
	return nil, nil
}
