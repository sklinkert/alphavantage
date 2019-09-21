package alphavantage

import (
	"encoding/json"
	"fmt"
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

// IndicatorStoch fetches the
func (c *Client) IndicatorStoch(symbol string, interval Interval) (*IndicatorStoch, error) {
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&interval=%s&apikey=%s",
		baseURL, "STOCH", symbol, interval, c.apiKey)
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
