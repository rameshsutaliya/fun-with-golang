package service

import (
	"fun-with-golang/pocs/lru/cache"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/pkg/errors"
)

type (
	Service interface {
		Get(key string) (string, bool)
		Put(key, val string) bool
		Post(key, val string) bool
		RemoveFromCache(key string) bool
		EraseCache()
		GetAllCache() []string
	}
	service struct {
		val      map[string]string
		expCache *expirable.LRU[string, string]
		fixCache *lru.Cache[string, string]
	}
)

func NewService(isExpirable bool) Service {
	s := &service{val: make(map[string]string)}
	var err error
	if isExpirable {
		s.expCache = cache.GetExpirableCache(10, 5)
	} else {
		s.fixCache, err = cache.GetFixedSizeCache(10)
		if err != nil {
			err = errors.Wrap(err, "Creating the fix cache")
			panic(err)
		}
	}
	return s
}

func (s *service) Get(key string) (string, bool) {
	if s.expCache != nil {
		return s.expCache.Get(key)
	} else {
		return s.fixCache.Get(key)
	}
}

func (s *service) Put(key, val string) bool {
	if s.expCache != nil {
		s.expCache.Remove(key)
		return s.expCache.Add(key, val)
	} else {
		s.fixCache.Remove(key)
		return s.fixCache.Add(key, val)
	}
}

func (s *service) Post(key, val string) bool {
	if s.expCache != nil {
		return s.expCache.Add(key, val)
	} else {
		return s.fixCache.Add(key, val)
	}
}

func (s *service) RemoveFromCache(key string) bool {
	if s.expCache != nil {
		return s.expCache.Remove(key)
	} else {
		return s.fixCache.Remove(key)
	}
}

func (s *service) EraseCache() {
	if s.expCache != nil {
		s.expCache.Purge()
	} else {
		s.fixCache.Purge()
	}
}

func (s *service) GetAllCache() []string {
	if s.expCache != nil {
		return s.expCache.Keys()
	} else {
		return s.fixCache.Keys()
	}
}
