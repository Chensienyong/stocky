package stocky

import (
	"context"
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

	FetchDailySeriesByStock(int64) ([]entity.Daily, error)
	InsertDailies([]entity.Daily) ([]entity.Daily, error)
	DeleteDailies(int64) error
}

type RedisMethod interface {
	Get(context.Context, string) (string, error)
	Set(context.Context, string, interface{}) error
	Exists(context.Context, string) int64
	SetEx(context.Context, string, interface{}, time.Duration) error
	DeleteByKey(context.Context, string) error
	HSet(context.Context, string, string, string) error
	HSetJSON(context.Context, string, string, interface{}) error
	HMSet(context.Context, string, map[string]interface{}) error
	HGet(context.Context, string, string) (string, error)
	HExists(context.Context, string, string) (bool, error)
	HDel(context.Context, string, string) error
	SAdd(string, interface{}) error
	SMembers(string) ([]string, error)
	Delete(string) error
	Read(string) ([]byte, error)
	Write(string, []byte, time.Duration) error
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
