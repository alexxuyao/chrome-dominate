package chromedominate

import (
	"sync"
	"time"
)

type CacheItem struct {
	Id      int64
	AddTime int64
	Data    []byte
}

type ResultCache struct {
	cMap       map[int64]*CacheItem
	ticker     *time.Ticker
	tickerStop chan bool
	mutex      *sync.RWMutex
	keepTime   time.Duration
}

func NewDefaultResultCache() *ResultCache {
	return NewResultCache(1*time.Minute, 1*time.Minute)
}

func NewResultCache(tickerTime time.Duration, keepTime time.Duration) *ResultCache {
	cache := &ResultCache{
		cMap:       make(map[int64]*CacheItem),
		ticker:     time.NewTicker(tickerTime),
		tickerStop: make(chan bool),
		mutex:      new(sync.RWMutex),
		keepTime:   keepTime,
	}

	go func(cache *ResultCache) {
		defer cache.ticker.Stop()

		for {
			select {
			case <-cache.ticker.C:
				cache.clear()
			case stop := <-cache.tickerStop:
				if stop {
					return
				}
			}
		}

	}(cache)

	return cache
}

func (cache *ResultCache) SetKeepTime(keepTime time.Duration) {
	cache.keepTime = keepTime
}

func (cache *ResultCache) Put(id int64, data []byte) {
	item := &CacheItem{
		Id:      id,
		AddTime: time.Now().UnixNano(),
		Data:    data,
	}

	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.cMap[id] = item
}

func (cache *ResultCache) Pop(id int64, timeout time.Duration) (*CacheItem, bool) {

	startTime := time.Now().UnixNano()

	for {
		now := time.Now().UnixNano()
		if now-startTime > timeout.Nanoseconds() {
			return nil, false
		}

		ret, find := cache.pop(id)
		if find {
			return ret, find
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func (cache *ResultCache) Close() {
	cache.tickerStop <- true
}

func (cache *ResultCache) pop(id int64) (*CacheItem, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	r, ok := cache.cMap[id]
	if ok {
		delete(cache.cMap, r.Id)
	}

	return r, ok
}

func (cache *ResultCache) clear() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	now := time.Now().UnixNano()

	for _, item := range cache.cMap {
		if now-item.AddTime > cache.keepTime.Nanoseconds() {
			delete(cache.cMap, item.Id)
		}
	}

}
