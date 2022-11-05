package cacheutils

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestLRU(t *testing.T) {
	lru := NewLRU[string](20)
	lru.Set("a", "s111_a")
	lru.Set("b", "s111_b")

	v, err := lru.Get("a")
	if err != nil || v != "s111_a" {
		t.Error(err)
	}
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))

	testData := make(map[string]string)

	for i := 0; i < 2100; i++ {
		k := strconv.Itoa(i)
		v := strconv.Itoa(i + r.Int())
		lru.Set(k, v)
		testData[k] = v

		// t.Log("size: ", lru.GetSize())
	}
	for _, k := range lru.Keys() {
		result, err := lru.Get(k)
		if testData[k] != result || err != nil {
			t.Error("error", err)
		}
	}

	for i := 0; i < 20; i++ {
		k := strconv.Itoa(i)
		v := strconv.Itoa(i + r.Int())
		// expire after 10 second

		if i >= 10 {
			lru.SetWithTimeout(k, v, time.Now().Add(1*time.Second))
		} else {
			lru.SetWithTimeout(k, v, time.Now().Add(3*time.Second))
		}

		testData[k] = v
	}

	fmt.Println("before size: ", lru.GetSize())
	time.Sleep(1 * time.Second)

	lru.Prune()

	fmt.Println("after prune size: ", lru.GetSize())

	if lru.GetSize() != 10 {
		t.Error("timeout test error")
	}

	time.Sleep(2 * time.Second)

	lru.Prune()

	fmt.Println("after prune size: ", lru.GetSize())

	if lru.GetSize() != 0 {
		t.Error("0, timeout test error")
	}
}
