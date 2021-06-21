package entity_test

import (
	"github.com/chensienyong/stocky/connection"
	"github.com/chensienyong/stocky/entity"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDaily(t *testing.T) {
	open := 100.5
	var volume int64 = 508
	daily := entity.NewDaily(8, "2021-05-28", open, 150.2, 99.0, 142.5, volume)
	assert.IsType(t, entity.Daily{}, daily)
	assert.Equal(t, open, daily.Open)
	assert.Equal(t, volume, daily.Volume)
}

func TestCreateDailyBatch(t *testing.T) {
	datas := make(map[string]connection.Data)
	datas["2021-05-28"] = connection.Data{
		Open: "100.5", High: "150.2", Low: "99.0",
		Close: "142.5", Volume: "508",
	}
	datas["2021-05-27"] = connection.Data{
		Open: "98.0", High: "110.0", Low: "95.0",
		Close: "100.5", Volume: "259",
	}
	dailies := entity.CreateDailyBatch(8, connection.DailyResponse{TimeSeries:datas})
	assert.IsType(t, []entity.Daily{}, dailies)
	open, _   := strconv.ParseFloat(datas["2021-05-28"].Open, 64)
	assert.Equal(t, open, dailies[0].Open)
	assert.Equal(t, "2021-05-27", dailies[1].Date)
}

func TestNewDailyResponse(t *testing.T) {
	daily := entity.NewDaily(8, "2021-05-28", 100.5, 150.2, 99.0, 142.5, 508)
	dailyResponse := entity.NewDailyResponse(daily)
	assert.IsType(t, entity.DailyResponse{}, dailyResponse)
	assert.Equal(t, daily.Open, dailyResponse.Open)
	assert.Equal(t, daily.Volume, dailyResponse.Volume)
}

func TestNewDailiesResponse(t *testing.T) {
	var dailies []entity.Daily
	dailies = append(dailies,
		entity.NewDaily(8, "2021-05-28", 100.5, 150.2, 99.0, 142.5, 508),
		entity.NewDaily(8, "2021-05-27", 98.0, 110.0, 95.0, 100.5, 259))
	dailiesResponse := entity.NewDailiesResponse(dailies)
	assert.IsType(t, []entity.DailyResponse{}, dailiesResponse)
	assert.Equal(t, 100.5, dailiesResponse[0].Open)
	assert.Equal(t, int64(259), dailiesResponse[1].Volume)
}
