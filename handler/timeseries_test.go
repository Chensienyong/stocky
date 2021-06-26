package handler_test

import (
	"github.com/chensienyong/stocky"
	"github.com/chensienyong/stocky/customerror"
	"github.com/chensienyong/stocky/entity"
	"github.com/chensienyong/stocky/handler"
	"github.com/chensienyong/stocky/mocks"
	rds "github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

func TestHandler_FetchDailyTimeSeries_DBError(t *testing.T) {
	stockSymbol := "FB"
	pg := mocks.PostgresMock{}
	pg.On("GetOrCreateStock", stockSymbol).Return(nil, customerror.DBError).Once()
	stck := stocky.NewStocky(&pg, &mocks.RedisMock{}, &mocks.AlphaVantageMock{})
	h := handler.NewHandler(stck)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{ Key: "stock", Value: stockSymbol }}
	err := h.FetchDailyTimeSeries(w, nil, params)

	assert.Equal(t, customerror.DBError, err)
	assert.Equal(t, 500, w.Code)
}

func TestHandler_FetchDailyTimeSeries_RedisError(t *testing.T) {
	stockSymbol := "FB"
	stock := entity.Stock{ ID: 1, StockSymbol: stockSymbol }
	pg := mocks.PostgresMock{}
	pg.On("GetOrCreateStock", stockSymbol).Return(stock, nil).Once()
	redis := mocks.RedisMock{}
	redis.On("Get", mock.Anything).Return(nil, customerror.RedisError).Once()
	stck := stocky.NewStocky(&pg, &redis, &mocks.AlphaVantageMock{})
	h := handler.NewHandler(stck)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{ Key: "stock", Value: stockSymbol }}
	err := h.FetchDailyTimeSeries(w, nil, params)

	assert.Equal(t, customerror.RedisError, err)
	assert.Equal(t, 500, w.Code)
}

func TestHandler_FetchDailyTimeSeries_SuccessWithExisting(t *testing.T) {
	stockSymbol := "FB"
	stock := entity.Stock{ ID: 1, StockSymbol: stockSymbol }
	pg := mocks.PostgresMock{}
	pg.On("GetOrCreateStock", stockSymbol).Return(stock, nil).Once()
	pg.On("FetchDailySeriesByStock", stock.ID).Return(nil, nil).Once()
	redis := mocks.RedisMock{}
	redis.On("Get", mock.Anything).Return(nil, nil).Once()
	stck := stocky.NewStocky(&pg, &redis, &mocks.AlphaVantageMock{})
	h := handler.NewHandler(stck)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{ Key: "stock", Value: stockSymbol }}
	err := h.FetchDailyTimeSeries(w, nil, params)

	assert.Nil(t, err)
	assert.Equal(t, 200, w.Code)
}

func TestHandler_FetchDailyTimeSeries_FailInFetching(t *testing.T) {
	stockSymbol := "FB"
	stock := entity.Stock{ ID: 1, StockSymbol: stockSymbol }
	pg := mocks.PostgresMock{}
	pg.On("GetOrCreateStock", stockSymbol).Return(stock, nil).Once()
	pg.On("FetchDailySeriesByStock", stock.ID).Return(nil, customerror.DBError).Once()
	redis := mocks.RedisMock{}
	redis.On("Get", mock.Anything).Return(nil, nil).Once()
	stck := stocky.NewStocky(&pg, &redis, &mocks.AlphaVantageMock{})
	h := handler.NewHandler(stck)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{ Key: "stock", Value: stockSymbol }}
	err := h.FetchDailyTimeSeries(w, nil, params)

	assert.Equal(t, customerror.DBError, err)
	assert.Equal(t, 500, w.Code)
}

func TestHandler_FetchDailyTimeSeries_SuccessWithNewEntry(t *testing.T) {
	stockSymbol := "FB"
	stock := entity.Stock{ ID: 1, StockSymbol: stockSymbol }
	// Mocking
	pg := mocks.PostgresMock{}
	pg.On("GetOrCreateStock", stockSymbol).Return(stock, nil).Once()
	pg.On("FetchDailySeriesByStock", stock.ID).Return(nil, nil).Once()
	pg.On("DeleteDailies", stock.ID).Return(nil).Once()
	pg.On("InsertDailies", mock.Anything).Return(nil).Once()
	redis := mocks.RedisMock{}
	redis.On("Get", mock.Anything).Return(nil, rds.Nil).Once()
	redis.On("SetEx", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	alpha := mocks.AlphaVantageMock{}
	alpha.On("GetDaily", stockSymbol).Return(nil, nil).Once()
	stck := stocky.NewStocky(&pg, &redis, &alpha)
	h := handler.NewHandler(stck)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{ Key: "stock", Value: stockSymbol }}

	err := h.FetchDailyTimeSeries(w, nil, params)

	assert.Nil(t, err)
	assert.Equal(t, 200, w.Code)
}

func TestHandler_FetchDailyTimeSeries_FailFetchingFromAlpha(t *testing.T) {
	stockSymbol := "FB"
	stock := entity.Stock{ ID: 1, StockSymbol: stockSymbol }
	// Mocking
	pg := mocks.PostgresMock{}
	pg.On("GetOrCreateStock", stockSymbol).Return(stock, nil).Once()
	pg.On("FetchDailySeriesByStock", stock.ID).Return(nil, nil).Once()
	redis := mocks.RedisMock{}
	redis.On("Get", mock.Anything).Return(nil, rds.Nil).Once()
	alpha := mocks.AlphaVantageMock{}
	alpha.On("GetDaily", stockSymbol).Return(nil, customerror.RecordNotFound).Once()
	stck := stocky.NewStocky(&pg, &redis, &alpha)
	h := handler.NewHandler(stck)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{ Key: "stock", Value: stockSymbol }}

	err := h.FetchDailyTimeSeries(w, nil, params)

	assert.Equal(t, customerror.RecordNotFound, err)
	assert.Equal(t, 404, w.Code)
}

func TestHandler_FetchDailyTimeSeries_FailOnDeleteDailies(t *testing.T) {
	stockSymbol := "FB"
	stock := entity.Stock{ ID: 1, StockSymbol: stockSymbol }
	// Mocking
	pg := mocks.PostgresMock{}
	pg.On("GetOrCreateStock", stockSymbol).Return(stock, nil).Once()
	pg.On("FetchDailySeriesByStock", stock.ID).Return(nil, nil).Once()
	pg.On("DeleteDailies", stock.ID).Return(customerror.DBError).Once()
	redis := mocks.RedisMock{}
	redis.On("Get", mock.Anything).Return(nil, rds.Nil).Once()
	alpha := mocks.AlphaVantageMock{}
	alpha.On("GetDaily", stockSymbol).Return(nil, nil).Once()
	stck := stocky.NewStocky(&pg, &redis, &alpha)
	h := handler.NewHandler(stck)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{ Key: "stock", Value: stockSymbol }}

	err := h.FetchDailyTimeSeries(w, nil, params)

	assert.Equal(t, customerror.DBError, err)
	assert.Equal(t, 500, w.Code)
}

func TestHandler_FetchDailyTimeSeries_FailOnInsertDailies(t *testing.T) {
	stockSymbol := "FB"
	stock := entity.Stock{ ID: 1, StockSymbol: stockSymbol }
	// Mocking
	pg := mocks.PostgresMock{}
	pg.On("GetOrCreateStock", stockSymbol).Return(stock, nil).Once()
	pg.On("FetchDailySeriesByStock", stock.ID).Return(nil, nil).Once()
	pg.On("DeleteDailies", stock.ID).Return(nil).Once()
	pg.On("InsertDailies", mock.Anything).Return(customerror.DBError).Once()
	redis := mocks.RedisMock{}
	redis.On("Get", mock.Anything).Return(nil, rds.Nil).Once()
	alpha := mocks.AlphaVantageMock{}
	alpha.On("GetDaily", stockSymbol).Return(nil, nil).Once()
	stck := stocky.NewStocky(&pg, &redis, &alpha)
	h := handler.NewHandler(stck)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{ Key: "stock", Value: stockSymbol }}

	err := h.FetchDailyTimeSeries(w, nil, params)

	assert.Equal(t, customerror.DBError, err)
	assert.Equal(t, 500, w.Code)
}

func TestHandler_FetchDailyTimeSeries_FailOnRedisSet(t *testing.T) {
	stockSymbol := "FB"
	stock := entity.Stock{ ID: 1, StockSymbol: stockSymbol }
	// Mocking
	pg := mocks.PostgresMock{}
	pg.On("GetOrCreateStock", stockSymbol).Return(stock, nil).Once()
	pg.On("FetchDailySeriesByStock", stock.ID).Return(nil, nil).Once()
	pg.On("DeleteDailies", stock.ID).Return(nil).Once()
	pg.On("InsertDailies", mock.Anything).Return(nil).Once()
	redis := mocks.RedisMock{}
	redis.On("Get", mock.Anything).Return(nil, rds.Nil).Once()
	redis.On("SetEx", mock.Anything, mock.Anything, mock.Anything).Return(customerror.RedisError).Once()
	alpha := mocks.AlphaVantageMock{}
	alpha.On("GetDaily", stockSymbol).Return(nil, nil).Once()
	stck := stocky.NewStocky(&pg, &redis, &alpha)
	h := handler.NewHandler(stck)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{ Key: "stock", Value: stockSymbol }}

	err := h.FetchDailyTimeSeries(w, nil, params)

	assert.Equal(t, customerror.RedisError, err)
	assert.Equal(t, 500, w.Code)
}
