package repository

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Cache struct {
	cache *memcache.Client
}

func NewCache() *Cache {
	cacheServer := memcache.New("localhost:11211")
	cacheServer.Timeout = 5 * time.Second
	return &Cache{
		cache: cacheServer,
	}
}

func (c *Cache) SetCache(key string, value []byte, expiration int32) error {
	item := &memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}

	return c.cache.Set(item)
}

func (c *Cache) GetCache(key string) ([]byte, error) {
	item, err := c.cache.Get(key)

	if err != nil {
		return nil, err
	}

	return item.Value, err

}
