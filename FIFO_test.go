package cacheutils

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestFIFO(t *testing.T) {
	fifo := NewFIFO[string](20)
	fifo.Set("a", "s111_a")
	fifo.Set("b", "s111_b")

	v, err := fifo.Get("a")
	if err != nil || v != "s111_a" {
		t.Error(err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixMilli()))

	testData := make(map[string]string)

	for i := 0; i < 2100; i++ {
		k := strconv.Itoa(i)
		v := strconv.Itoa(i + r.Int())
		fifo.Set(k, v)
		testData[k] = v

		// t.Log("size: ", fifo.GetSize())
	}
	for _, k := range fifo.Keys() {
		result, err := fifo.Get(k)
		if testData[k] != result || err != nil {
			t.Error("error", err)
		}
	}

	for i := 0; i < 20; i++ {
		k := strconv.Itoa(i)
		v := strconv.Itoa(i + r.Int())
		// expire after 10 second

		if i >= 10 {
			fifo.SetWithTimeout(k, v, time.Now().Add(1*time.Second))
		} else {
			fifo.SetWithTimeout(k, v, time.Now().Add(3*time.Second))
		}

		testData[k] = v
	}

	fmt.Println("before size: ", fifo.GetSize())
	time.Sleep(1 * time.Second)

	fifo.Prune()

	fmt.Println("after prune size: ", fifo.GetSize())

	if fifo.GetSize() != 10 {
		t.Error("timeout test error")
	}

	time.Sleep(2 * time.Second)

	fifo.Prune()

	fmt.Println("after prune size: ", fifo.GetSize())

	if fifo.GetSize() != 0 {
		t.Error("0, timeout test error")
	}
}
