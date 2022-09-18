package cache

import (
	"sync"

	gocache "github.com/patrickmn/go-cache"
)

var (
	lock     sync.RWMutex
	once     sync.Once
	appCache *gocache.Cache
)

func NewCache() *gocache.Cache {
	once.Do(func() {
		appCache = gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	})
	return appCache
}

type Page struct {
	PageIndicator int
	Items         []struct {
		Display string
		Link    string
	}
}

type PageItem struct {
	PageIndicator int
	Items         []PageItem
}
