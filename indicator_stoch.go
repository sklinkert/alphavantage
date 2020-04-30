package alphavantage

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

// IndicatorStoch represents the overall struct for stochastics indicator
// Example https://www.alphavantage.co/query?function=STOCH&symbol=MSFT&interval=daily&apikey=demo
type IndicatorStoch struct {
	Metadata          IndicatorStochMetadata            `json:"Meta Data"`
	TechnicalAnalysis map[string]TechnicalStochAnalysis `json:"Technical Analysis: STOCH"`
}

// IndicatorStochMetadata is the metadata subset of IndicatorStoch
type IndicatorStochMetadata struct {
	Symbol        string `json:"1: Symbol"`
	Indicator     string `json:"2: Indicator"`
	LastRefreshed string `json:"3: Last Refreshed"`
	Interval      string `json:"4: Interval"`
	FastKPeriod   int    `json:"5.1: FastK Period"`
	SlowKPeriod   int    `json:"5.2: SlowK Period"`
	SlowKMAType   int    `json:"5.3: SlowK MA Type"`
	SlowDPeriod   int    `json:"5.4: SlowD Period"`
	SlowDMAType   int    `json:"5.5: SlowD MA Type"`
	TimeZone      string `json:"6: Time Zone"`
}

// TechnicalStochAnalysis is the stoch indicator subset of IndicatorStoch
type TechnicalStochAnalysis struct {
	SlowK float64 `json:",string"`
	SlowD float64 `json:",string"`
}

func toIndicatorStoch(buf []byte) (*IndicatorStoch, error) {
	indicatorStoch := &IndicatorStoch{}
	if err := json.Unmarshal(buf, indicatorStoch); err != nil {
		return nil, err
	}
	return indicatorStoch, nil
}

// IndicatorStoch fetches the "STOCH" indicators for given symbol from API.
// The order of dates in TechnicalAnalysis is random because it's a map.
func (c *Client) IndicatorStoch(symbol string, interval Interval) (*IndicatorStoch, error) {
	const functionName = "STOCH"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&interval=%s&apikey=%s",
		baseURL, functionName, symbol, interval, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	indicator, err := toIndicatorStoch(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return indicator, nil
}

// Latest returns the most recent TechnicalStochAnalysis for given stoch.
func (stoch *IndicatorStoch) Latest() (date string, latest *TechnicalStochAnalysis) {
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

// Today returns TechnicalStochAnalysis for today.
func (stoch *IndicatorStoch) Today() *TechnicalStochAnalysis {
	today := time.Now()
	return stoch.ByDate(today)
}

// ByDate returns TechnicalStochAnalysis for the given date.
func (stoch *IndicatorStoch) ByDate(date time.Time) *TechnicalStochAnalysis {
	day := date.Format(DateFormat)
	item, exists := stoch.TechnicalAnalysis[day]
	if !exists {
		return nil
	}
	return &item
}
