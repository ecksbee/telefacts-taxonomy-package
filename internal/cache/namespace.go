package cache

import (
	"encoding/json"
	"sync"

	gocache "github.com/patrickmn/go-cache"
)

func MarshalNamespace(repo *NamespaceRepo, namespace string) ([]byte, error) {
	items := nsrepo.Query(namespace)
	data := Page{
		PageIndicator: 0,
		Items:         items,
	}
	return json.Marshal(&data)
}

type NamespaceRepo struct {
	lock  sync.RWMutex
	cache *gocache.Cache
}

func (repo *NamespaceRepo) Query(namespace string) []PageItem {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	if x, found := repo.cache.Get("ns:" + namespace); found {
		ret := x.([]PageItem)
		return ret
	}
	return make([]PageItem, 0)
}

func NewNamespaceRepo(cache *gocache.Cache) (*NamespaceRepo, error) {
	//todo inspect namespaces from discovered base taxonomies
	return nil, nil
}
