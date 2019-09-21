package alphavantage

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

// TimeSeries represents the overall struct for time series
type TimeSeries struct {
	Metadata         TimeSeriesMetadata        `json:"Meta Data"`
	TimeSeriesDaily  map[string]TimeSeriesData `json:"Time Series (Daily)"`
	TimeSeriesWeekly map[string]TimeSeriesData `json:"Weekly Time Series"`
}

// TimeSeriesMetadata is the metadata subset of TimeSeries
type TimeSeriesMetadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"4. Time Zone"`
}

// TimeSeriesData is a subset of TimeSeries
type TimeSeriesData struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume uint64  `json:"5. volume,string"`
}

func toTimeSeries(buf []byte) (*TimeSeries, error) {
	timeSeries := &TimeSeries{}
	if err := json.Unmarshal(buf, timeSeries); err != nil {
		return nil, err
	}
	return timeSeries, nil
}

// TimeSeries fetches the time series for given symbol from API.
// The order of dates in returned object is random because it's a map.
func (c *Client) TimeSeries(symbol string, interval TimeSeriesInterval) (*TimeSeries, error) {
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, interval, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	timeSeries, err := toTimeSeries(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return timeSeries, nil
}

// getFilledData returns the data subset for the filled interval
func (ts *TimeSeries) getFilledData() map[string]TimeSeriesData {
	if len(ts.TimeSeriesDaily) > 0 {
		return ts.TimeSeriesDaily
	} else if len(ts.TimeSeriesWeekly) > 0 {
		return ts.TimeSeriesWeekly
	}
	return nil
}

// Latest returns the most recent item
func (ts *TimeSeries) Latest() (date string, latest *TimeSeriesData) {
	datasets := ts.getFilledData()
	dates := make([]string, len(datasets))
	for date := range datasets {
		dates = append(dates, date)
	}
	sort.Strings(dates)
	date = dates[len(dates)-1]
	latestVal, _ := datasets[date]
	latest = &latestVal
	return
}

// Today returns dataset for today.
func (ts *TimeSeries) Today() *TimeSeriesData {
	today := time.Now()
	return ts.ByDate(today)
}

// ByDate returns the dataset for the given date.
func (ts *TimeSeries) ByDate(date time.Time) *TimeSeriesData {
	dataset := ts.getFilledData()
	day := date.Format("2006-02-01")
	item, exists := dataset[day]
	if !exists {
		return nil
	}
	return &item
}
