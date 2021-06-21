package handler

import (
	"github.com/chensienyong/stocky/entity"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//FetchDailyTimeSeries is used to getting a challenge
func (h *Handler) FetchDailyTimeSeries(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	stockSymbol := params.ByName("stock")
	dailyseries, err := h.stocky.AlphaVantage.GetDaily(stockSymbol)
	if err != nil {
		Error(w, err)
		return err
	}

	stock, err := h.stocky.Postgres.GetOrCreateStock(stockSymbol)
	if err != nil {
		Error(w, err)
		return err
	}
	err = h.stocky.Postgres.DeleteDailies(stock.ID)
	if err != nil {
		Error(w, err)
		return err
	}
	dailies := entity.CreateDailyBatch(stock.ID, dailyseries)
	err = h.stocky.Postgres.InsertDailies(dailies)
	if err != nil {
		Error(w, err)
		return err
	}
	stockDailySeries, err := h.stocky.Postgres.FetchDailySeriesByStock(stock.ID)
	if err != nil {
		Error(w, err)
		return err
	}

	OK(w, entity.NewStockDailyResponse(stock, stockDailySeries), "")
	return nil
}
