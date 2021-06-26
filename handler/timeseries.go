package handler

import (
	"fmt"
	"github.com/chensienyong/stocky/entity"
	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

//FetchDailyTimeSeries is used to getting a challenge
func (h *Handler) FetchDailyTimeSeries(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	stockSymbol := params.ByName("stock")

	stock, err := h.Stocky.Postgres.GetOrCreateStock(stockSymbol)
	if err != nil {
		Error(w, err)
		return err
	}

	redisKey := fmt.Sprintf("stocky_%s", stockSymbol)
	_, err = h.Stocky.Redis.Get(redisKey)
	if err == redis.Nil {
		err = h.updateDaily(stockSymbol, redisKey, stock.ID)
		if err != nil {
			Error(w, err)
			return err
		}
	} else if err != nil {
		Error(w, err)
		return err
	}

	stockDailySeries, err := h.Stocky.Postgres.FetchDailySeriesByStock(stock.ID)
	if err != nil {
		Error(w, err)
		return err
	}

	OK(w, entity.NewStockDailyResponse(stock, stockDailySeries), "")
	return nil
}

func (h *Handler) updateDaily(stockSymbol, redisKey string, stockID int64) error {
	dailyseries, err := h.Stocky.AlphaVantage.GetDaily(stockSymbol)
	if err != nil {
		return err
	}

	err = h.Stocky.Postgres.DeleteDailies(stockID)
	if err != nil {
		return err
	}

	dailies := entity.CreateDailyBatch(stockID, dailyseries)
	err = h.Stocky.Postgres.InsertDailies(dailies)
	if err != nil {
		return err
	}

	err = h.Stocky.Redis.SetEx(redisKey, "exists", time.Hour)
	if err != nil {
		return err
	}

	return nil
}