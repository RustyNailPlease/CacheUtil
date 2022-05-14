package cacheutils

import (
	"errors"
	"time"

	"github.com/SomeoneDeng/CacheUtil/internal"
)

type FIFOCache[T interface{}] struct {
	CacheInfo[T]

	LinkedList *internal.DoubleLinkedList[T]
}

func NewFIFO[T interface{}](cap int64) *FIFOCache[T] {
	return &FIFOCache[T]{
		CacheInfo: CacheInfo[T]{
			Capacity: int64(cap),
			Size:     0,
			Data:     make(map[string]CacheItem[T]),
		},
		LinkedList: internal.NewDoubleLinkedList[T](),
	}
}

func (c *FIFOCache[T]) GetCapacity() int64 {
	return c.Capacity
}

func (c *FIFOCache[T]) GetSize() int64 {
	return c.Size
}

func (c *FIFOCache[T]) Set(key string, value T) error {
	if c.Fulled() {
		c.Prune()
		if c.Fulled() {
			first := c.LinkedList.GetFirst()
			if first != nil {
				c.Remove(first.Key)
			}
		}
	}
	// infinity
	c.Data[key] = CacheItem[T]{key, value, -1}
	c.LinkedList.Add(key, value)
	c.Size++
	// todo
	return nil
}

func (c *FIFOCache[T]) SetWithTimeout(key string, value T, timeout int64) error {
	if c.Fulled() {
		c.Prune()
		if c.Fulled() {
			first := c.LinkedList.GetFirst()
			if first != nil {
				c.Remove(first.Key)
			}
		}
	}
	c.Data[key] = CacheItem[T]{key, value, timeout}
	c.LinkedList.Add(key, value)
	c.Size++
	return nil
}

func (c *FIFOCache[T]) Get(key string) (*T, error) {
	if item, ok := c.Data[key]; ok {
		if item.ExpireTime > 0 {
			if item.ExpireTime > time.Now().UnixMilli() {
				return &item.Value, nil
			} else {
				c.Remove(key)
				return nil, errors.New("expired")
			}
		} else {
			return &item.Value, nil
		}
	} else {
		return nil, errors.New("not found")
	}
}

func (c *FIFOCache[T]) Prune() (int, error) {
	count := 0
	for _, v := range c.Data {
		// expired
		if v.ExpireTime != -1 && v.ExpireTime <= time.Now().UnixMilli() {
			c.Remove(v.Key)
			count++
		}
	}
	return count, nil
}

func (c *FIFOCache[T]) Fulled() bool {
	return c.GetCapacity() == c.GetSize()
}

func (c *FIFOCache[T]) Remove(key string) error {
	delete(c.Data, key)
	c.LinkedList.Delete(key)
	c.Size--
	return nil
}

func (c *FIFOCache[T]) Contains(key string) bool {
	e, ok := c.Data[key]

	expired := false
	if e.ExpireTime != -1 {
		if e.ExpireTime <= time.Now().UnixMilli() {
			expired = true
		}
	}

	return ok && expired
}

func (c *FIFOCache[T]) Keys() []string {
	keys := make([]string, 0)
	for k := range c.Data {
		keys = append(keys, k)
	}
	return keys
}
