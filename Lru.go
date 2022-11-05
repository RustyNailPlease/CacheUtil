package cacheutils

import (
	"errors"
	"time"

	"github.com/RustyNailPlease/CacheUtil/internal"
)

type LRUCache[T interface{}] struct {
	CacheInfo[T]
	// 双向链表
	LinkedList *internal.DoubleLinkedList[T]
}

func NewLRU[T interface{}](cap int64) *LRUCache[T] {
	return &LRUCache[T]{
		CacheInfo: CacheInfo[T]{
			Capacity: int64(cap),
			Size:     0,
			Data:     make(map[string]CacheItem[T]),
		},
		LinkedList: internal.NewDoubleLinkedList[T](),
	}
}

func (c *LRUCache[T]) GetCapacity() int64 {
	return c.Capacity
}

func (c *LRUCache[T]) GetSize() int64 {
	return c.Size
}

func (c *LRUCache[T]) Set(key string, value T) error {
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
	c.Data[key] = CacheItem[T]{key, value, nil}
	c.LinkedList.Add(key, value)
	c.Size++
	// todo
	return nil
}

func (c *LRUCache[T]) SetWithTimeout(key string, value T, timeout time.Time) error {
	if c.Fulled() {
		c.Prune()
		if c.Fulled() {
			first := c.LinkedList.GetFirst()
			if first != nil {
				c.Remove(first.Key)
			}
		}
	}
	c.Data[key] = CacheItem[T]{key, value, &timeout}
	c.LinkedList.Add(key, value)
	c.Size++
	return nil
}

func (c *LRUCache[T]) Get(key string) (t T, err error) {
	if item, ok := c.Data[key]; ok {
		if item.ExpireTime != nil {
			if item.ExpireTime.After(time.Now()) {
				c.LinkedList.Delete(key)
				c.LinkedList.Add(key, c.Data[key].Value)
				return item.Value, nil
			} else {
				c.Remove(key)
				return t, errors.New("expired")
			}
		} else {
			c.LinkedList.Delete(key)
			c.LinkedList.Add(key, c.Data[key].Value)
			return item.Value, nil
		}
	} else {
		return t, errors.New("not found")
	}
}

func (c *LRUCache[T]) Prune() (int, error) {
	count := 0
	for _, v := range c.Data {
		// expired
		if v.ExpireTime != nil && v.ExpireTime.Before(time.Now()) {
			c.Remove(v.Key)
			count++
		}
	}
	return count, nil
}

func (c *LRUCache[T]) Fulled() bool {
	return c.GetCapacity() == c.GetSize()
}

func (c *LRUCache[T]) Remove(key string) error {
	delete(c.Data, key)
	c.LinkedList.Delete(key)
	c.Size--
	return nil
}

func (c *LRUCache[T]) Contains(key string) bool {
	e, ok := c.Data[key]

	expired := false
	if e.ExpireTime != nil && e.ExpireTime.Before(time.Now()) {
		expired = true
	}

	return ok && !expired
}

func (c *LRUCache[T]) Keys() []string {
	keys := make([]string, 0)
	for k := range c.Data {
		keys = append(keys, k)
	}
	return keys
}
