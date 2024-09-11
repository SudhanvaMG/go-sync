package application

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

func TestKVStore(t *testing.T) {
	t.Run("should test put and get", func(t *testing.T) {
		t.Run("should be able to put and get a key", func(t *testing.T) {
			kv := NewKeyValueStore()
			kv.Put("key", "value")

			value, err := kv.Get("key")
			assert.NoError(t, err)
			require.Equal(t, "value", value)
		})
		t.Run("should handle concurrent access without race conditions", func(t *testing.T) {
			kv := NewKeyValueStore()
			var wg sync.WaitGroup
			const numRoutines = 100

			for i := 0; i < numRoutines; i++ {
				wg.Add(1)
				go func(n int) {
					defer wg.Done()
					key := "key" + string(rune(n))
					kv.Put(key, "value")
				}(i)
			}

			wg.Wait()

			for i := 0; i < numRoutines; i++ {
				wg.Add(1)
				go func(n int) {
					defer wg.Done()
					key := "key" + string(rune(n))
					_, err := kv.Get(key)
					assert.NoError(t, err)
				}(i)
			}

			wg.Wait()
		})
	})
	t.Run("should return an error for a missing key", func(t *testing.T) {
		kv := NewKeyValueStore()
		_, err := kv.Get("missing")
		assert.Equal(t, ErrKeyNotFound, err)
	})
	t.Run("should be able to delete a key", func(t *testing.T) {
		kv := NewKeyValueStore()
		kv.Put("key", "value")

		value, err := kv.Get("key")
		assert.NoError(t, err)
		require.Equal(t, "value", value)

		err = kv.Delete("key")
		assert.NoError(t, err)

		_, err = kv.Get("key")
		assert.Equal(t, ErrKeyNotFound, err)
	})
	t.Run("should return an error when deleting a missing key", func(t *testing.T) {
		kv := NewKeyValueStore()
		err := kv.Delete("missing")
		assert.Equal(t, ErrKeyNotFound, err)
	})
	t.Run("should be able to list all keys", func(t *testing.T) {
		kv := NewKeyValueStore()
		kv.Put("key1", "value1")
		kv.Put("key2", "value2")
		kv.Put("key3", "value3")

		keys := kv.ListKeys()
		require.Len(t, keys, 3)
		assert.Contains(t, keys, "key1")
		assert.Contains(t, keys, "key2")
		assert.Contains(t, keys, "key3")
	})
	t.Run("should test update", func(t *testing.T) {
		t.Run("should update the value for an existing key", func(t *testing.T) {
			kv := NewKeyValueStore()
			kv.Put("key", "value")

			value, err := kv.Get("key")
			assert.NoError(t, err)
			require.Equal(t, "value", value)

			kv.Put("key", "new-value")

			value, err = kv.Get("key")
			assert.NoError(t, err)
			require.Equal(t, "new-value", value)
		})
		t.Run("should handle concurrent read and write operations ", func(t *testing.T) {
			kv := NewKeyValueStore()
			var wg sync.WaitGroup
			const numRoutines = 50

			for i := 0; i < numRoutines; i++ {
				wg.Add(1)
				go func(n int) {
					defer wg.Done()
					key := "key" + string(rune(n))
					kv.Put(key, "value")
				}(i)
			}

			for i := 0; i < numRoutines; i++ {
				wg.Add(1)
				go func(n int) {
					defer wg.Done()
					key := "key" + string(rune(i))
					_, err := kv.Get(key)
					if err != nil && !errors.Is(err, ErrKeyNotFound) {
						assert.NoError(t, err)
					}
				}(i)
			}

			wg.Wait()
		})
	})
}
