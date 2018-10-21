package handlers

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

var storage Cacher

type Cacher interface {
	Get(key string) []byte
	Set(key string, content []byte, duration time.Duration)
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
	return time.Now().UnixNano() > item.Expiration
}

//Storage mecanism for caching strings in memory
type Storage struct {
	items map[string]Item
	mu    *sync.RWMutex
}

//NewStorage creates a new in memory storage
func NewStorage() *Storage {
	return &Storage{
		items: make(map[string]Item),
		mu:    &sync.RWMutex{},
	}
}

//Get a cached content by key
func (s Storage) Get(key string) []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item := s.items[key]
	if item.Expired() {
		delete(s.items, key)
		return nil
	}
	return item.Content
}

//Set a cached content by key
func (s Storage) Set(key string, content []byte, duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = Item{
		Content:    content,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
}

func cached(duration string, handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content := storage.Get(r.RequestURI)
		if content != nil {
			w.Write(content)
		} else {
			c := httptest.NewRecorder()
			handler(c, r)
			for k, v := range c.HeaderMap {
				w.Header()[k] = v
			}
			w.WriteHeader(c.Code)
			content := c.Body.Bytes()
			if d, err := time.ParseDuration(duration); err == nil {
				storage.Set(r.RequestURI, content, d)
			}
			w.Write(content)
		}
	})
}
