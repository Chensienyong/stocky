package entity

type Stock struct {
	ID          int64  `json:"id"`
	StockSymbol string `json:"stock_symbol"`
}

type StockDailyResponse struct {
	Stock       string          `json:"stock"`
	DailySeries []DailyResponse `json:"daily_series,omitempty"`
}

// NewStock is a constructor for the Stock object
func NewStock(stockSymbol string) Stock {
	return Stock{
		StockSymbol: stockSymbol,
	}
}

func NewStockDailyResponse(stock Stock, dailies []Daily) StockDailyResponse {
	return StockDailyResponse{
		Stock:       stock.StockSymbol,
		DailySeries: NewDailiesResponse(dailies),
	}
}
