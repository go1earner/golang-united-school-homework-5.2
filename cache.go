package cache

import (
	"fmt"
	"time"
)

var (
	timeLayout = "2006-1-2 15:4:5"
)

type Cache struct {
	storage map[string]map[string]string
}

func NewCache(storage map[string]map[string]string) Cache {
	return Cache{
		storage: storage,
	}
}

func (cache *Cache) Get(key string) (string, bool) {
	deadline := cache.storage[key]["deadline"]
	if deadline != "" {
		timeNow, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))
		deadlineTime, _ := time.Parse(timeLayout, deadline)
		fmt.Println("The key deadline is", deadlineTime)
		fmt.Println("Time now is", timeNow)
		if deadlineTime.Before(timeNow) {
			delete(cache.storage, key)
		}
	}
	v, ok := cache.storage[key]["value"]
	return v, ok
}

func (cache *Cache) Put(key, value string) {
	cache.storage[key] = map[string]string{
		"value":    value,
		"deadline": "",
	}
}

func (cache Cache) Keys() []string {
	keys := make([]string, 0)
	for key := range cache.storage {
		keys = append(keys, key)
	}
	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.storage[key] = map[string]string{
		"value":    value,
		"deadline": deadline.Format(timeLayout),
	}
}
