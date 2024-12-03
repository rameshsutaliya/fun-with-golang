package cache

import (
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"time"
)

// fixed size cache

func GetFixedSizeCache(size int) (*lru.Cache[string, string], error) {
	return lru.New[string, string](size)
}

func GetExpirableCache(expireTime int, size int) *expirable.LRU[string, string] {
	return expirable.NewLRU[string, string](size, nil, time.Second*time.Duration(expireTime))
}
