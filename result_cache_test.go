package chromedominate

import (
	"testing"
	"time"
)

func TestResultCache(t *testing.T) {
	cache := NewResultCache(1*time.Second, 1*time.Second)

	go func(cache *ResultCache) {
		cache.Put(1, []byte("first"))
		cache.Put(3, []byte("3"))
	}(cache)

	ret, find := cache.Pop(1, time.Second*1)

	if !find {
		panic("find is false")
	}

	if string(ret.Data) != "first" {
		panic("find is not first")
	}

	ret, find = cache.Pop(2, time.Second*1)

	if find {
		panic("find is true")
	}

	time.Sleep(time.Second * 1)

	ret, find = cache.Pop(3, time.Second*1)

	if find {
		panic("3 data error")
	}

	cache.Close()

}
