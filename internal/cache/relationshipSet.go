package cache

import (
	"encoding/json"
	"sync"

	gocache "github.com/patrickmn/go-cache"
)

func MarshalRelationshipSet(roleuri string) ([]byte, error) {
	items := rsrepo.Query(roleuri)
	data := Page{
		PageIndicator: 0,
		Items:         items,
	}
	return json.Marshal(&data)
}

type RelationshipSetRepo struct {
	lock  sync.RWMutex
	cache *gocache.Cache
}

func (repo *RelationshipSetRepo) Query(roleuri string) []PageItem {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	if x, found := repo.cache.Get("rs:" + roleuri); found {
		ret := x.([]PageItem)
		return ret
	}
	return make([]PageItem, 0)
}

func NewRelationshipSetRepo(cache *gocache.Cache, gts string) (*RelationshipSetRepo, error) {
	//todo inspect relationship sets from discovered base taxonomies
	return nil, nil
}
