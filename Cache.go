package cacheutils

import "time"

// Cache interface
type Cache[T interface{}] interface {
	// GetCapacity returns the capacity of the cache.
	GetCapacity() int64
	// GetSize returns the size of the cache.
	GetSize() int64
	// Set sets the value of the key.
	Set(key string, value T) error
	// Get gets the value of the key with timeout.
	SetWithTimeout(key string, value T, timeout int64) error
	// Get gets the value of the key.
	Get(key string) (T, error)
	// Prune remove the expired items
	Prune() (int, error)
	// Fulled returns true if the cache is full.
	Fulled() bool
	// Remove removes the key.
	Remove(key string) error
	// Contains returns true if the key exists.
	Contains(key string) bool
	// Keys returns all the keys.
	Keys() []string
}

type CacheInfo[T interface{}] struct {
	Capacity int64
	Size     int64
	Data     map[string]CacheItem[T]
}

type CacheItem[T interface{}] struct {
	Key        string
	Value      T
	ExpireTime *time.Time // milliseconds
}
