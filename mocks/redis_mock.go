package mocks

import (
	"github.com/stretchr/testify/mock"
	"time"
)

// RedisMock is a mock for redis
type RedisMock struct {
	mock.Mock
}

// Get is a mock for redis Get function
func (rm *RedisMock) Get(key string) (string, error) {
	args := rm.Called(key)
	res := args.Get(0)

	if res == nil {
		return "", args.Error(1)
	}
	return res.(string), args.Error(1)
}

// SetEx is a mock for redis SetEx function
func (rm *RedisMock) SetEx(key string, value interface{}, duration time.Duration) error {
	args := rm.Called(key, value, duration)
	return args.Error(0)
}
