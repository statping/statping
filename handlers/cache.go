package handlers

import (
	"github.com/hunterlong/statping/core"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

var CacheStorage Cacher

type Cacher interface {
	Get(key string) []byte
	Delete(key string)
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

//NewStorage creates a new in memory CacheStorage
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
		CacheStorage.Delete(key)
		return nil
	}
	return item.Content
}

func (s Storage) Delete(key string) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	delete(s.items, key)
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

func cached(duration, contentType string, handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content := CacheStorage.Get(r.RequestURI)
		w.Header().Set("Content-Type", contentType)
		if core.Configs == nil {
			handler(w, r)
			return
		}
		if content != nil {
			w.Write(content)
		} else {
			c := httptest.NewRecorder()
			handler(c, r)
			content := c.Body.Bytes()
			result := c.Result()
			if result.StatusCode != 200 {
				w.WriteHeader(result.StatusCode)
				w.Write(content)
				return
			}
			if d, err := time.ParseDuration(duration); err == nil {
				CacheStorage.Set(r.RequestURI, content, d)
			}
			w.Write(content)
		}
	})
}
