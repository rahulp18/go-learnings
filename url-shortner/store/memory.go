package store

import "sync"

type MemoryStore struct {
	data  map[string]string
	mutex sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]string),
	}
}
func (m *MemoryStore) Save(code string, longURL string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data[code] = longURL
}
func (m *MemoryStore) Get(code string) (string, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	data, ok := m.data[code]
	return data, ok
}
