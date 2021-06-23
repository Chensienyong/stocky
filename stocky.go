package stocky

import (
	"github.com/chensienyong/stocky/entity"
	"time"

	"github.com/chensienyong/stocky/connection"
)

type Stocky struct {
	Postgres      PostgresMethod
	Redis         RedisMethod
	AlphaVantage  connection.AlphaVantageProvider
}

type PostgresMethod interface {
	GetOrCreateStock(string) (entity.Stock, error)
	CreateStock(entity.Stock) (entity.Stock, error)
	GetStocks() ([]entity.Stock, error)

	FetchDailySeriesByStock(int64) ([]entity.Daily, error)
	InsertDailies([]entity.Daily) error
	DeleteDailies(int64) error
}

type RedisMethod interface {
	Get(string) (string, error)
	SetEx(string, interface{}, time.Duration) error
}

func NewStocky(postgres PostgresMethod,
	redis RedisMethod,
	alphaVantage connection.AlphaVantageProvider) *Stocky {
	return &Stocky{
		Postgres:     postgres,
		Redis:        redis,
		AlphaVantage: alphaVantage,
	}
}
