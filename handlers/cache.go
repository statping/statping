package handlers

import (
	"github.com/statping/statping/utils"
	"sync"
	"time"
)

var CacheStorage Cacher

type Cacher interface {
	Get(key string) []byte
	Delete(key string)
	Set(key string, content []byte, duration time.Duration)
	List() map[string]Item
	Lock()
	Unlock()
	StopRoutine()
}

// Item is a cached reference
type Item struct {
	Content    []byte
	Expiration int64
}

// cleanRoutine is a go routine to automatically remove expired caches that haven't been hit recently
func cleanRoutine(s *Storage) {
	duration := 5 * time.Second

CacheRoutine:
	for {
		select {
		case <-s.running:
			break CacheRoutine
		case <-time.After(duration):
			for k, v := range s.List() {
				if v.Expired() {
					s.Delete(k)
				}
			}
			duration = 5 * time.Second
		}
	}
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
	items   map[string]Item
	mu      *sync.RWMutex
	running chan bool
}

//NewStorage creates a new in memory CacheStorage
func NewStorage() *Storage {
	storage := &Storage{
		items:   make(map[string]Item),
		mu:      &sync.RWMutex{},
		running: make(chan bool),
	}
	go cleanRoutine(storage)
	return storage
}

func (s Storage) StopRoutine() {
	close(s.running)
}

func (s Storage) Lock() {
	s.mu.Lock()
}

func (s Storage) Unlock() {
	s.mu.Unlock()
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
