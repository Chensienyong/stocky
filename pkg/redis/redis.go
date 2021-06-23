package redis

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// Redis holds the redis client.
type Redis struct {
	Client *redis.Client
}

// Options stores the redis connection options.
type Options struct {
	Network  string
	Addr     string
	Password string
	DB       string
}

// NewRedis initializes a new redis object.
func NewRedis(opt Options) (*Redis, error) {
	db, _ := strconv.Atoi(opt.DB)
	client := redis.NewClient(&redis.Options{
		Network:  opt.Network,
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Redis{Client: client}, nil
}

// Get fetches a value in the redis based on the key
func (c *Redis) Get(key string) (string, error) {
	cmd := c.Client.Get(key)

	return cmd.Val(), cmd.Err()
}

// SetEx sets a value to an assigned key with expired time
func (c *Redis) SetEx(key string, val interface{}, duration time.Duration) error {
	newVal, err := json.Marshal(val)
	if err != nil {
		return err
	}
	status := c.Client.Set(key, newVal, duration)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}
