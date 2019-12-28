package handlers

import (
	"github.com/hunterlong/statping/utils"
	"sync"
	"time"
)

var CacheStorage Cacher

type Cacher interface {
	Get(key string) []byte
	Delete(key string)
	Set(key string, content []byte, duration time.Duration)
	List() map[string]Item
}

// Item is a cached reference
type Item struct {
	Content    []byte
	Expiration int64
}

// Expired returns true if the item has expired.
func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return utils.Now().UnixNano() > item.Expiration
}

//Storage mecanism for caching strings in memory
type Storage struct {
	items map[string]Item
	mu    *sync.RWMutex
}

//NewStorage creates a new in memory CacheStorage
func NewStorage() *Storage {
	return &Storage{
		items: make(map[string]Item),
		mu:    &sync.RWMutex{},
	}
}

func (s Storage) List() map[string]Item {
	return s.items
}

//Get a cached content by key
func (s Storage) Get(key string) []byte {
	item := s.items[key]
	if item.Expired() {
		CacheStorage.Delete(key)
		return nil
	}
	return item.Content
}

func (s Storage) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, key)
}

//Set a cached content by key
func (s Storage) Set(key string, content []byte, duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = Item{
		Content:    content,
		Expiration: utils.Now().Add(duration).UnixNano(),
	}
}
