package application

import (
	"errors"
	"sync"
)

var ErrKeyNotFound = errors.New("key not found")

type KVStore struct {
	store sync.Map
}

func NewKeyValueStore() *KVStore {
	return &KVStore{}
}

// Get retrieves the value for a given key.
func (kv *KVStore) Get(key string) (string, error) {
	if value, ok := kv.store.Load(key); ok {
		return value.(string), nil
	}
	return "", ErrKeyNotFound
}

// Put sets the value for the given key.
func (kv *KVStore) Put(key string, value string) {
	kv.store.Store(key, value)
}

// Delete removes the key-value pair from the store.
func (kv *KVStore) Delete(key string) error {
	if _, ok := kv.store.Load(key); ok {
		kv.store.Delete(key)
		return nil
	}
	return ErrKeyNotFound
}

// ListKeys returns a slice of all keys in the store.
func (kv *KVStore) ListKeys() []string {
	var keys []string
	kv.store.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}
