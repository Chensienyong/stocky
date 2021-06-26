package mocks

import (
	"github.com/chensienyong/stocky/entity"

	"github.com/stretchr/testify/mock"
)

// PostgresMock is a mock for database
type PostgresMock struct {
	mock.Mock
}

// GetOrCreateStock is a mock for GetOrCreateStock function
func (pg *PostgresMock) GetOrCreateStock(stockSymbol string) (entity.Stock, error) {
	args := pg.Called(stockSymbol)
	res := args.Get(0)

	if res == nil {
		return entity.Stock{}, args.Error(1)
	}

	return res.(entity.Stock), args.Error(1)
}

// CreateStock is a mock for CreateStock function
func (pg *PostgresMock) CreateStock(stock entity.Stock) (entity.Stock, error) {
	args := pg.Called(stock)
	res := args.Get(0)

	if res == nil {
		return entity.Stock{}, args.Error(1)
	}

	return res.(entity.Stock), args.Error(1)
}

// GetStocks is a mock for GetStocks function
func (pg *PostgresMock) GetStocks() ([]entity.Stock, error) {
	args := pg.Called()
	res := args.Get(0)

	if res == nil {
		return []entity.Stock{}, args.Error(1)
	}

	return res.([]entity.Stock), args.Error(1)
}

// FetchDailySeriesByStock is a mock for FetchDailySeriesByStock function
func (pg *PostgresMock) FetchDailySeriesByStock(stockID int64) ([]entity.Daily, error) {
	args := pg.Called(stockID)
	res := args.Get(0)

	if res == nil {
		return []entity.Daily{}, args.Error(1)
	}

	return res.([]entity.Daily), args.Error(1)
}

// InsertDailies is a mock for InsertDailies function
func (pg *PostgresMock) InsertDailies(dailies []entity.Daily) error {
	args := pg.Called(dailies)
	return args.Error(0)
}

// DeleteDailies is a mock for DeleteDailies function
func (pg *PostgresMock) DeleteDailies(dailyID int64) error {
	args := pg.Called(dailyID)
	return args.Error(0)
}
