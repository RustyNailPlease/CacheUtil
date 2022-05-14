# Cache Utils


Caches for simple use in memory.

- LRU
- FIFO (todo)


#### install

``` bash
go get github.com/SomeoneDeng/CacheUtil
```

``` golang
lru := NewLRU[string](20)
lru.Set("a", "s111_a")
lru.Set("b", "s111_b")
```