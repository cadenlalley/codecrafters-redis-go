package cache

import "sync"

var internalCache cache

type cache struct {
	kv map[string]string
	mu sync.Mutex
}

func init() {
	internalCache = cache{
		kv: map[string]string{},
		mu: sync.Mutex{},
	}
}

func Set(k string, v string) error {
	internalCache.mu.Lock()
	internalCache.kv[k] = v
	internalCache.mu.Unlock()

	return nil
}

func Get(k string) (string, error) {
	return internalCache.kv[k], nil
}
