package database

import (
	"errors"
	"github.com/chensienyong/stocky/entity"
	"gorm.io/gorm"
)

const TableStock = "stocks"

func (pg *Postgres) GetOrCreateStock(stockSymbol string) (entity.Stock, error) {
	var stock entity.Stock

	err := pg.Db.Table(TableStock).Select("stocks.id, stocks.stock_symbol").Where("stocks.stock_symbol = ?", stockSymbol).First(&stock)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return pg.CreateStock(entity.NewStock(stockSymbol))
		}
		return stock, err.Error
	}

	return stock, nil
}

func (pg *Postgres) CreateStock(stock entity.Stock) (entity.Stock, error) {
	var newStock entity.Stock

	err := pg.Db.Table(TableStock).Create(&stock).Scan(&newStock)
	if err.Error != nil {
		return stock, err.Error
	}

	return stock, nil
}
