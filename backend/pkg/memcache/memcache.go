package memcache

import (
	"sync"
	"time"
)

type MemCache struct {
	data map[string]*entry
	once sync.Once
	mu   sync.Mutex
}

func (m *MemCache) lazyinit() {
	m.once.Do(func() {
		m.data = make(map[string]*entry)
		go m.gc()
	})
}

func (m *MemCache) Set(key string, value any) {
	m.lazyinit()
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = &entry{k: key, v: value}
}

func (m *MemCache) SetWithExpire(key string, value any, ttl time.Duration) {
	m.lazyinit()
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = &entry{k: key, v: value, expire: time.Now().Add(ttl)}
}

func (m *MemCache) Get(key string) (any, bool) {
	m.lazyinit()
	m.mu.Lock()
	defer m.mu.Unlock()
	if v, ok := m.data[key]; ok {
		if v.expired() {
			delete(m.data, key)
			return nil, false
		}
		return v.v, true
	}
	return nil, false
}

func (m *MemCache) delete(key string) {
	m.lazyinit()
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func (m *MemCache) Delete(key string) {
	m.delete(key)
}

func (m *MemCache) gc() {
	for {
		time.Sleep(time.Minute * 5)
		for k, v := range m.data {
			if v.expired() {
				go m.delete(k)
			}
		}
	}
}
