package entity

import (
	"github.com/chensienyong/stocky/connection"
	"strconv"
)

type Daily struct {
	ID          int64   `json:"id"`
	StockID     int64   `json:"stock_id"`
	Date        string  `json:"date"`
	Open        float64 `json:"open"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	Close       float64 `json:"close"`
	Volume      int64   `json:"volume"`
}

type DailyResponse struct {
	Date        string  `json:"date,omitempty"`
	Open        float64 `json:"open,omitempty"`
	High        float64 `json:"high,omitempty"`
	Low         float64 `json:"low,omitempty"`
	Close       float64 `json:"close,omitempty"`
	Volume      int64   `json:"volume,omitempty"`
}

func NewDaily(stockID int64, date string, open float64, high float64, low float64, close float64, volume int64) Daily {
	return Daily{
		StockID: stockID,
		Date:    date,
		Open:    open,
		High:    high,
		Low:     low,
		Close:   close,
		Volume:  volume,
	}
}

func CreateDailyBatch(stockID int64, dailyResponse connection.DailyResponse) []Daily {
	var dailies []Daily
	for date, v := range dailyResponse.TimeSeries {
		open, _   := strconv.ParseFloat(v.Open, 64)
		high, _   := strconv.ParseFloat(v.High, 64)
		low, _    := strconv.ParseFloat(v.Low, 64)
		close, _  := strconv.ParseFloat(v.Close, 64)
		volume, _ := strconv.ParseInt(v.Volume, 10, 64)
		dailies = append(dailies,
			NewDaily(stockID, date, open, high, low, close, volume),
		)
	}
	return dailies
}

func NewDailyResponse(daily Daily) DailyResponse {
	return DailyResponse{
		Date:        daily.Date,
		Open:        daily.Open,
		High:        daily.High,
		Low:         daily.Low,
		Close:       daily.Close,
		Volume:      daily.Volume,
	}
}

func NewDailiesResponse(dailies []Daily) []DailyResponse {
	var dailiesResponse []DailyResponse
	for _, daily := range dailies {
		dailiesResponse = append(dailiesResponse, NewDailyResponse(daily))
	}
	return dailiesResponse
}
