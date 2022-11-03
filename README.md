# Cache Utils


Cache implements for simple use in memory.

- LRU
- FIFO


#### install

``` bash
go get github.com/RustyNailPlease/CacheUtil
```

``` golang
lru := NewLRU[string](20)
lru.Set("a", "s111_a")
lru.Set("b", "s111_b")
```