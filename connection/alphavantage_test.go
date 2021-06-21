package connection_test

import (
	"fmt"
	"github.com/chensienyong/stocky/connection"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewAlphaVantage(t *testing.T) {
	apiKey  := "ABC123"
	host    := "localhost"
	timeout := 1*time.Second

	alphaVantage := connection.NewAlphaVantage(apiKey, host, timeout)
	assert.IsType(t, &connection.AlphaVantage{}, alphaVantage)
}

func TestAlphaVantage_GetDaily_Success(t *testing.T) {
	stockSymbol := "GGRM"
	apiKey      := "ABC123"
	host        := "localhost"
	url         := fmt.Sprintf("%s/query?function=TIME_SERIES_DAILY&outputsize=compact&symbol=%s&apikey=%s", host, stockSymbol, apiKey)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, `
	  {
		"Meta Data": {
		  "1. Information": "Daily Prices (open, high, low, close) and Volumes",
		  "2. Symbol": "IBM",
		  "3. Last Refreshed": "2021-06-18",
		  "4. Output Size": "Compact",
		  "5. Time Zone": "US/Eastern"
		},
		"Time Series (Daily)": {
	      "2021-06-18": {
	        "1. open": "144.4800",
	        "2. high": "144.6800",
	        "3. low": "143.0400",
	        "4. close": "143.1200",
	        "5. volume": "9156505"
	      }
	    }
	  }
	`))

	alphaVantage := connection.NewAlphaVantage(host, apiKey, 1*time.Second)
	resp, err := alphaVantage.GetDaily(stockSymbol)

	assert.Nil(t, err)
	assert.IsType(t, connection.DailyResponse{}, resp)
}

func TestAlphaVantage_GetDaily_FailFetching(t *testing.T) {
	stockSymbol := "GGRM"
	apiKey      := "ABC123"
	host        := "localhost"
	url         := fmt.Sprintf("%s/query?function=TIME_SERIES_DAILY&outputsize=compact&symbol=%s&apikey=%s", host, stockSymbol, apiKey)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(500, "Fail"))

	alphaVantage := connection.NewAlphaVantage(host, apiKey, 1*time.Second)
	resp, _ := alphaVantage.GetDaily(stockSymbol)

	assert.Equal(t, connection.DailyResponse{}, resp)
}
