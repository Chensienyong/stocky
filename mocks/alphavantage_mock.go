package mocks

import (
	"github.com/chensienyong/stocky/connection"
	"github.com/stretchr/testify/mock"
)

// AlphaVantage is a mock struct for alphavantage
type AlphaVantage struct {
	mock.Mock
}

// GetDaily is a mock for alphavantage service GetDaily function
func (av *AlphaVantage) GetDaily(stockSymbol string) (connection.DailyResponse, error) {
	args := av.Called(stockSymbol)
	res := args.Get(0)

	if res == nil {
		return connection.DailyResponse{}, args.Error(1)
	}

	return res.(connection.DailyResponse), args.Error(1)
}
