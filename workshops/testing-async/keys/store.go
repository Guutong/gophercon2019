package keys

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// notFound is an error returned if the key is not found on a Get lookup
type notFound struct {
	msg string
}

func (nf notFound) Error() string {
	return nf.msg
}

func (notFound) NotFound() {}

func init() {
	rand.Seed(time.Now().Unix())
}

// section: store
// Store is a key/value store that is safe for use in concurrent operations
type Store struct {
	id     int
	values map[string]interface{}
	mu     sync.RWMutex
}

// NewStore will return an instance of store with initialized internal data
func NewStore() *Store {
	s := Store{
		values: map[string]interface{}{},
	}
	return &s
}

// section: store

// section: methods
// Set will store the key/value pair
// If the value exists, it is overwritten
func (vs *Store) Set(key string, value interface{}) {
	// lock to get a new identity
	vs.mu.Lock()
	vs.id++
	vs.mu.Unlock()
	// Make this an asynchronous call
	go func() {
		vs.mu.Lock()
		defer vs.mu.Unlock()

		// take a random amount of time to return to simulate the real world
		t := time.Duration(rand.Intn(3)) * time.Second
		time.Sleep(t)

		vs.values[key] = value
		log.Println("inserted: ", key, " with value of ", value)
	}()
}

// Get will return the specified value.
// If the value is not found it will return a notFound error
func (vs *Store) Get(key string) (interface{}, error) {
	vs.mu.RLock()
	defer vs.mu.RUnlock()
	if v, ok := vs.values[key]; ok {
		return v, nil
	}
	return nil, &notFound{msg: fmt.Sprintf("%s not found", key)}
}

// Count returns the number of entries currently in the store
func (vs *Store) Count() int {
	vs.mu.RLock()
	defer vs.mu.RUnlock()
	return len(vs.values)
}
