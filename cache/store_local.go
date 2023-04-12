package cache

import (
	"encoding/json"
	"time"

	"github.com/patrickmn/go-cache"
)

type LocalStore struct {
	GoCacheClient *cache.Cache
	KeyPrefix     string
}

func (s *LocalStore) Set(key string, value string, expireTime time.Duration) {
	s.GoCacheClient.Set(s.KeyPrefix+key, value, expireTime)
}

func (s *LocalStore) Get(key string) string {
	res, has := s.GoCacheClient.Get(s.KeyPrefix + key)
	if !has {
		return ""
	}

	by, err := json.Marshal(res)
	if err != nil {
		return ""
	}

	return string(by)
}

func (s *LocalStore) Has(key string) bool {
	_, has := s.GoCacheClient.Get(s.KeyPrefix + key)

	return has
}

func (s *LocalStore) Forget(key string) {
	s.GoCacheClient.Delete(s.KeyPrefix + key)
}

func (s *LocalStore) Forever(key string, value string) {
	s.GoCacheClient.Set(s.KeyPrefix+key, value, 0)
}

func (s *LocalStore) Flush() {
	s.GoCacheClient.Flush()
}

func (s *LocalStore) Incr(key string) {
	_ = s.GoCacheClient.Increment(s.KeyPrefix+key, 1)
}

func (s *LocalStore) IncrBy(key string, value int64) {
	_ = s.GoCacheClient.Increment(s.KeyPrefix+key, value)
}

func (s *LocalStore) Decr(key string) {
	_ = s.GoCacheClient.Decrement(s.KeyPrefix+key, 1)
}

func (s *LocalStore) DecrBy(key string, value int64) {
	_ = s.GoCacheClient.Decrement(s.KeyPrefix+key, value)
}
