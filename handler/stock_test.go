package handler_test

import (
	"github.com/chensienyong/stocky"
	"github.com/chensienyong/stocky/customerror"
	"github.com/chensienyong/stocky/handler"
	"github.com/chensienyong/stocky/mocks"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetStocks(t *testing.T) {
	pg := mocks.PostgresMock{}
	pg.On("GetStocks").Return(nil, nil).Once()
	stck := stocky.NewStocky(&pg, &mocks.RedisMock{}, &mocks.AlphaVantageMock{})
	h := handler.NewHandler(stck)
	w := httptest.NewRecorder()
	err := h.GetStocks(w, nil, nil)

	assert.Nil(t, err)
	assert.Equal(t, 200, w.Code)
}

func TestHandler_GetStocks_ErrorPostgres(t *testing.T) {
	pg := mocks.PostgresMock{}
	pg.On("GetStocks").Return(nil, customerror.DBError).Once()
	stck := stocky.NewStocky(&pg, &mocks.RedisMock{}, &mocks.AlphaVantageMock{})
	h := handler.NewHandler(stck)
	w := httptest.NewRecorder()
	err := h.GetStocks(w, nil, nil)

	assert.Equal(t, customerror.DBError, err)
	assert.Equal(t, customerror.DBError.Code, w.Code)
}
