package entity_test

import (
	"github.com/chensienyong/stocky/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStock(t *testing.T) {
	stockSymbol := "GGRM"
	stock := entity.NewStock(stockSymbol)
	assert.IsType(t, entity.Stock{}, stock)
	assert.Equal(t, stockSymbol, stock.StockSymbol)
}

func TestNewStockDailyResponse(t *testing.T) {
	stockSymbol := "GGRM"
	stock := entity.NewStock(stockSymbol)

	var dailies []entity.Daily
	dailies = append(dailies,
		entity.NewDaily(8, "2021-05-28", 100.5, 150.2, 99.0, 142.5, 508),
		entity.NewDaily(8, "2021-05-27", 98.0, 110.0, 95.0, 100.5, 259))

	stockDailyResponse := entity.NewStockDailyResponse(stock, dailies)
	assert.IsType(t, entity.StockDailyResponse{}, stockDailyResponse)
	assert.Equal(t, stockSymbol, stockDailyResponse.Stock)
	assert.Equal(t, dailies[0].Close, stockDailyResponse.DailySeries[0].Close)
}
