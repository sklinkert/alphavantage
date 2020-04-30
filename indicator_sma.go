package alphavantage

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

const (
	// SeriesTypeOpen - Series Type Open
	SeriesTypeOpen = "open"
	// SeriesTypeHigh - Series Type High
	SeriesTypeHigh = "high"
	// SeriesTypeLow - Series Type Low
	SeriesTypeLow = "low"
	// SeriesTypeClose - Series Type Close
	SeriesTypeClose = "close"
)

// IndicatorSMA represents the overall struct for stochastics indicator
// Example https://www.alphavantage.co/query?function=STOCH&symbol=MSFT&interval=daily&apikey=demo
type IndicatorSMA struct {
	Metadata          IndicatorSMAMetadata            `json:"Meta Data"`
	TechnicalAnalysis map[string]TechnicalSMAAnalysis `json:"Technical Analysis: SMA"`
}

// IndicatorSMAMetadata is the metadata subset of IndicatorSMA
type IndicatorSMAMetadata struct {
	Symbol        string `json:"1: Symbol"`
	Indicator     string `json:"2: Indicator"`
	LastRefreshed string `json:"3: Last Refreshed"`
	Interval      string `json:"4: Interval"`
	TimePeriod    int    `json:"5: Time Period"`
	SeriesType    string `json:"6: Series Type"`
	TimeZone      string `json:"7: Time Zone"`
}

// TechnicalSMAAnalysis is the SMA indicator subset of IndicatorSMA
type TechnicalSMAAnalysis struct {
	SMA float64 `json:",string"`
}

func toIndicatorSMA(buf []byte) (*IndicatorSMA, error) {
	indicatorSMA := &IndicatorSMA{}
	if err := json.Unmarshal(buf, indicatorSMA); err != nil {
		return nil, err
	}
	return indicatorSMA, nil
}

// IndicatorSMA fetches the "SMA" indicators for given symbol from API.
// The order of dates in TechnicalAnalysis is random because it's a map.
func (c *Client) IndicatorSMA(symbol string, interval Interval, seriesType string, timePeriod int) (*IndicatorSMA, error) {
	const functionName = "SMA"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&interval=%s&time_period=%d&series_type=%s&apikey=%s",
		baseURL, functionName, symbol, interval, timePeriod, seriesType, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	indicator, err := toIndicatorSMA(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return indicator, nil
}

// Latest returns the most recent TechnicalSMAAnalysis for given stoch.
func (stoch *IndicatorSMA) Latest() (date string, latest *TechnicalSMAAnalysis) {
	if len(stoch.TechnicalAnalysis) == 0 {
		return "", nil
	}
	dates := make([]string, len(stoch.TechnicalAnalysis))
	for date := range stoch.TechnicalAnalysis {
		dates = append(dates, date)
	}
	sort.Strings(dates)
	date = dates[len(dates)-1]
	latestVal, _ := stoch.TechnicalAnalysis[date]
	latest = &latestVal
	return
}

// Today returns TechnicalSMAAnalysis for today.
func (stoch *IndicatorSMA) Today() *TechnicalSMAAnalysis {
	today := time.Now()
	return stoch.ByDate(today)
}

// ByDate returns TechnicalSMAAnalysis for the given date.
func (stoch *IndicatorSMA) ByDate(date time.Time) *TechnicalSMAAnalysis {
	day := date.Format(DateFormat)
	item, exists := stoch.TechnicalAnalysis[day]
	if !exists {
		return nil
	}
	return &item
}
