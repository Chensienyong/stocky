package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//GetStocks is used to get list of stock
func (h *Handler) GetStocks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	stocks, err := h.stocky.Postgres.GetStocks()
	if err != nil {
		Error(w, err)
		return err
	}

	OK(w, stocks, "")
	return nil
}
