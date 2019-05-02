package keys

import (
	"log"
	"sync"
)

type Store struct {
	id     int
	values map[string]interface{}
	mu     sync.RWMutex
}

func NewStore() *Store {
	s := Store{}
	s.values = make(map[string]interface{})
	return &s
}

func (vs *Store) Set(key string, value interface{}) error {
	vs.mu.Lock()
	defer vs.mu.Unlock()

	vs.id++

	vs.values[key] = value
	log.Println("inserted: ", key, " with value of ", value)

	return nil
}

func (vs *Store) Get(key string) (interface{}, error) {
	if v, ok := vs.values[key]; ok {
		return v, nil
	}
	return nil, notFound{key: key}
}

func (vs *Store) Count() int {
	return len(vs.values)
}
