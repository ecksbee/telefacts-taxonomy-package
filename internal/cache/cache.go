package cache

import (
	"sync"

	gocache "github.com/patrickmn/go-cache"
)

var (
	once   sync.Once
	nsrepo *NamespaceRepo
	rsrepo *RelationshipSetRepo
)

func InitRepo(gts string) {
	once.Do(func() {
		appCache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
		nsrepo, _ = NewNamespaceRepo(appCache, gts)
		rsrepo, _ = NewRelationshipSetRepo(appCache, gts)
	})
}

type NamespaceView struct {
	RelationshipSets []string
}

type RelationshipSetView struct {
	Arcs []string
}
