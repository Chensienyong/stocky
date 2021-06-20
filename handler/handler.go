// Package handler manages the data flow from client to appropriate service.
package handler

import (
	"net/http"

	stck "github.com/chensienyong/stocky"

	"github.com/julienschmidt/httprouter"
)

// Handler controls request flow from client to service.
type Handler struct {
	stocky *stck.Stocky
}

// Meta is used to consolidate all meta statuses.
type Meta struct {
	HTTPStatus int `json:"http_status"`
}

// NewHandler returns a pointer of Handler instance.
func NewHandler(stocky *stck.Stocky) *Handler {
	return &Handler{
		stocky: stocky,
	}
}

// Healthz is used to control the flow of GET /healthz endpoint.
func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("ok"))
}
