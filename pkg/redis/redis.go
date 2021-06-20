package redis

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

// Redis holds the redis client.
type Redis struct {
	client *redis.Client
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

	return &Redis{client: client}, nil
}

// Get fetches a value in the redis based on the key
func (c *Redis) Get(ctx context.Context, key string) (string, error) {
	cmd := c.client.Get(key)

	return cmd.Val(), cmd.Err()
}

// Set sets a value to an assigned key
func (c *Redis) Set(ctx context.Context, key string, val interface{}) error {
	newVal, err := json.Marshal(val)
	if err != nil {
		return err
	}
	status := c.client.Set(key, newVal, 0)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

// SetEx sets a value to an assigned key with expired time
func (c *Redis) SetEx(ctx context.Context, key string, val interface{}, duration time.Duration) error {
	newVal, err := json.Marshal(val)
	if err != nil {
		return err
	}
	status := c.client.Set(key, newVal, duration)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

// DeleteByKey deletes a value based on its key.
func (c *Redis) DeleteByKey(ctx context.Context, key string) error {
	status := c.client.Del(key)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

// Delete deletes a value based on its key.
func (c *Redis) Delete(key string) error {
	status := c.client.Del(key)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

// Exists Check exists of value based on its key.
func (c *Redis) Exists(ctx context.Context, key string) int64 {
	status := c.client.Exists(key)
	if status.Err() != nil {
		return 0
	}
	return status.Val()
}

// HSet implements redis' HSET.
func (c *Redis) HSet(ctx context.Context, key string, field string, value string) error {
	if res := c.client.HSet(key, field, value); res.Err() != nil {
		return res.Err()
	}

	return nil
}

// HSetJSON implements redis' HSET.
func (c *Redis) HSetJSON(ctx context.Context, key string, field string, value interface{}) error {
	newVal, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if res := c.client.HSet(key, field, newVal); res.Err() != nil {
		return res.Err()
	}

	return nil
}

// HMSet implements redis' HMSET.
func (c *Redis) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	if res := c.client.HMSet(key, fields); res.Err() != nil {
		return res.Err()
	}

	return nil
}

// HGet implements redis' HGET.
func (c *Redis) HGet(ctx context.Context, key string, field string) (string, error) {
	res := c.client.HGet(key, field)
	if res.Err() != nil {
		return "", res.Err()
	}

	result, _ := res.Result()

	return result, nil
}

// HExists implements redis' HExists.
func (c *Redis) HExists(ctx context.Context, key string, field string) (bool, error) {
	res := c.client.HExists(key, field)
	if res.Err() != nil {
		return false, res.Err()
	}

	result, _ := res.Result()

	return result, nil
}

// HDel implements redis' HDel.
func (c *Redis) HDel(ctx context.Context, key string, field string) error {
	res := c.client.HDel(key, field)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

// SAdd implements redis' SADD.
func (c *Redis) SAdd(key string, member interface{}) error {
	member, err := json.Marshal(member)
	if err != nil {
		return errors.Wrap(err, "failed to marshal json before sadd")
	}

	if res := c.client.SAdd(key, member); res.Err() != nil {
		return errors.Wrap(res.Err(), "failed to add to set")
	}

	return nil
}

// SMembers implements redis' SMEMBERS.
func (c *Redis) SMembers(key string) ([]string, error) {
	res := c.client.SMembers(key)
	if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "failed to add to set")
	}

	return res.Val(), nil
}

// Write writes the item for given key.
func (c *Redis) Write(key string, value []byte, expiration time.Duration) error {
	cmd := c.client.Set(key, value, expiration)
	return cmd.Err()
}

// Read reads the item for given key.
func (c *Redis) Read(key string) ([]byte, error) {
	cmd := c.client.Get(key)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	return cmd.Bytes()
}
