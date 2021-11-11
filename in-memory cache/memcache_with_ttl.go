package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	updated time.Time
	Value   User
}

type TTLCache struct {
	data map[string]*CacheItem
	ttl  time.Duration
	lock *sync.Mutex
}

func NewTTLCache(ttl time.Duration) (*TTLCache, error) {
	c := &TTLCache{
		data: map[string]*CacheItem{},
		ttl:  ttl,
		lock: &sync.Mutex{},
	}

	t := time.NewTicker(time.Second * 1)
	go func() {
		for range t.C {
			c.reap()
		}
	}()

	return c, nil
}

func (t *TTLCache) Set(key string, val User) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	item := &CacheItem{
		Value:   val,
		updated: time.Now(),
	}
	// update key
	t.data[key] = item
	return nil
}

func (t *TTLCache) Get(key string) *User {
	if item, ok := t.data[key]; ok {
		return &item.Value
	}
	return nil
}

func (t *TTLCache) reap() {
	t.lock.Lock()
	defer t.lock.Unlock()
	for k, v := range t.data {
		elapsed := time.Since(v.updated)
		if elapsed >= t.ttl {
			fmt.Printf("reaping key:" + k)
			delete(t.data, k)
		}
	}
}

type User struct {
	UserName string
}
