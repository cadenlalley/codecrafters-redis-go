package cache

import (
	"sync"
)

var internalCache cache

type cache struct {
	kv map[string]Item
	mu sync.RWMutex
}

type Item struct {
	Value any
}

func init() {
	internalCache = cache{
		kv: map[string]Item{},
		mu: sync.RWMutex{},
	}
}

func Set(k string, v Item) {
	internalCache.mu.Lock()
	internalCache.kv[k] = v
	internalCache.mu.Unlock()
}

func Get(k string) (Item, bool) {
	value, ok := internalCache.kv[k]

	return value, ok
}

func Delete(k string) {
	// TODO: might need to check if the item isnt already there and return a non-ok value
	delete(internalCache.kv, k)
}
