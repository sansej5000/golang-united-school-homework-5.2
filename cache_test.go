package cache

import (
	"fmt"
	"testing"
	"time"
)

func cacheCreate(delta time.Duration) Cache {
	timer_date := time.Now()
	tmr := timer_date.Add(time.Minute * delta)
	return Cache{storage: map[string]Item{"a": Item{value: "aa", timer: tmr}}}
}

func Test_NewCache(t *testing.T) {
	src := NewCache()
	cache := cacheCreate(1)
	a := fmt.Sprintf("%T\n", src)
	b := fmt.Sprintf("%T\n", cache)
	if a != b {
		t.Errorf("Expected 'Cache', got %v", src)
	}
}

func Test_Get_OneValue(t *testing.T) {
	cache := cacheCreate(10)
	data, ok := cache.Get("a")
	if ok != true {
		t.Errorf("Expected 'aa', got %s", data)
	}
}

func Test_Get_Empty(t *testing.T) {
	cache := cacheCreate(10)
	data, ok := cache.Get("b")
	if ok != false {
		t.Errorf("Expected '', got %s", data)
	}
}

func Test_Get_TimeIsOver(t *testing.T) {
	cache := cacheCreate(-10)
	data, ok := cache.Get("a")
	if ok != false {
		t.Errorf("Expected '', got %s", data)
	}
}

func Test_Put(t *testing.T) {
	cache := cacheCreate(-10)
	cache.Put("b", "bb")
	data, ok := cache.Get("b")
	if ok != true {
		t.Errorf("Expected 'bb', got %s", data)
	}
}

func Test_PutOverwriteValue(t *testing.T) {
	cache := cacheCreate(-10)
	cache.Put("b", "bb")
	cache.Put("b", "cc")
	data, _ := cache.Get("b")
	if data != "cc" {
		t.Errorf("Expected 'cc', got %s", data)
	}
}

func Test_Keys(t *testing.T) {
	cache := cacheCreate(-10)
	cache.Put("b", "bb")
	data := cache.Keys()
	if data[0] != "b" {
		t.Errorf("Expected '[b]', got %s", data)
	}
}

func Test_PutTill(t *testing.T) {
	cache := cacheCreate(10)
	tmr := time.Now()
	cache.PutTill("c", "cc", tmr)
	data, _ := cache.Get("c")
	if data == "cc" {
		t.Errorf("Expected 'false', got %s", data)
	}
}
