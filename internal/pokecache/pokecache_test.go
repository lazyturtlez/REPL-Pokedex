package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cahce is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		intputVal []byte
	}{
		{
			inputKey: "key1",
			intputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			intputVal: []byte("val2"),
		},
		{
			inputKey: "",
			intputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.intputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
		}
		if string(actual) != string(cas.intputVal) {
			t.Errorf("%s does not match %s",
			string(actual),
			cas.intputVal,
			)
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond*10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))
	time.Sleep(interval + time.Millisecond)
	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}