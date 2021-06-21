package connection

import (
	"encoding/json"
	"fmt"
	"github.com/chensienyong/stocky/pkg/logger"
	"net/http"
	"time"
)

//AlphaVantageProvider is a contract for stock data.
type AlphaVantageProvider interface {
	GetDaily(string) (DailyResponse, error)
}

// AlphaVantage holds connection data
type AlphaVantage struct {
	apiKey   string
	host     string
	client   *http.Client
}

type metadata struct {
	Information   string `json:"1. Information,omitempty"`
	Symbol 		  string `json:"2. Symbol,omitempty"`
	LastRefreshed string `json:"3. Last Refreshed,omitempty"`
	OutputSize 	  string `json:"4. Output Size,omitempty"`
	TimeZone 	  string `json:"5. Time Zone,omitempty"`
}

type Data struct {
	Open   string `json:"1. open,omitempty"`
	High   string `json:"2. high,omitempty"`
	Low    string `json:"3. low,omitempty"`
	Close  string `json:"4. close,omitempty"`
	Volume string `json:"5. volume,omitempty"`
}

type DailyResponse struct {
	Metadata   metadata        `json:"Meta Data,omitempty"`
	TimeSeries map[string]Data `json:"Time Series (Daily),omitempty"`
}

// NewAlphaVantage return a pointer of AlphaVantage instance
func NewAlphaVantage(host string, apiKey string, timeout time.Duration) *AlphaVantage {
	return &AlphaVantage{
		apiKey: apiKey,
		host:   host,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// GetDaily retrieve daily series of stock
func (alpha *AlphaVantage) GetDaily(stockSymbol string) (dailyResponse DailyResponse, err error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/query?function=TIME_SERIES_DAILY&outputsize=compact&symbol=%s&apikey=%s", alpha.host, stockSymbol, alpha.apiKey),
		nil)
	if err != nil {
		return DailyResponse{}, err
	}

	startTime := time.Now()
	resp, err := alpha.client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		logger.Infof("Fail on fetching daily series of stock %s. Time taken: %v", stockSymbol, float64(time.Since(startTime)/time.Millisecond))
		return DailyResponse{}, err
	}
	logger.Infof("Success on fetching daily series of stock %s. Time taken: %v", stockSymbol, float64(time.Since(startTime)/time.Millisecond))

	dailyResp := DailyResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&dailyResp); err != nil {
		return DailyResponse{}, err
	}

	return dailyResp, nil
}
