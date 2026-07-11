package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.org",
			val: []byte("moredata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case for %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key %s in cache, but it wasn't found", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected value %s for key %s, but got %s", c.val, c.key, val)
				return
			}
		})
	}
}

func TestRealLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key https://example.com in cache, but it wasn't found")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected key https://example.com to be removed from cache, but it was still found")
		return
	}
}
