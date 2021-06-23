package redis_test

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/chensienyong/stocky/pkg/redis"
	goredis "github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedis_Get_NotFound(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	stockKey := "stocky_STOCK"

	redis, _ := redis.NewRedis(redis.Options{
		Network:  "tcp",
		Addr:     s.Addr(),
	})

	_, err = redis.Get(stockKey)
	assert.Equal(t, goredis.Nil, err)
}

func TestRedis_Get_Found(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	stockKey := "stocky_STOCK"

	s.Set(stockKey, "exists")

	redis, _ := redis.NewRedis(redis.Options{
		Network:  "tcp",
		Addr:     s.Addr(),
	})

	value, err := redis.Get(stockKey)
	assert.Nil(t, err)
	assert.Equal(t, "exists", value)
}

func TestRedis_SetEx(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	stockKey := "stocky_STOCK"

	redis, _ := redis.NewRedis(redis.Options{
		Network:  "tcp",
		Addr:     s.Addr(),
	})

	err = redis.SetEx(stockKey, "exists", time.Second)
	assert.Nil(t, err)
}
