package stocky

import (
	"github.com/chensienyong/stocky/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStocky(t *testing.T) {
	pg    := &mocks.PostgresMock{}
	redis := &mocks.RedisMock{}
	alpha := &mocks.AlphaVantage{}

	stocky := NewStocky(pg, redis, alpha)

	assert.IsType(t, &Stocky{}, stocky)
}
