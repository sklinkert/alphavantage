package alphavantage

import (
	"encoding/json"
	"fmt"
)

type CompanyOverview struct {
	Symbol                     string  `json:"Symbol"`
	AssetType                  string  `json:"AssetType"`
	Name                       string  `json:"Name"`
	Description                string  `json:"Description"`
	CIK                        string  `json:"CIK"`
	Exchange                   string  `json:"Exchange"`
	Currency                   string  `json:"Currency"`
	Country                    string  `json:"Country"`
	Sector                     string  `json:"Sector"`
	Industry                   string  `json:"Industry"`
	Address                    string  `json:"Address"`
	FiscalYearEnd              string  `json:"FiscalYearEnd"`
	LatestQuarter              string  `json:"LatestQuarter"` // "2023-03-31"
	MarketCapitalization       int     `json:"MarketCapitalization,string"`
	EBITDA                     int     `json:"EBITDA,string"`
	PERatio                    float64 `json:"PERatio,string"`
	PEGRatio                   float64 `json:"PEGRatio,string"`
	BookValue                  float64 `json:"BookValue,string"`
	DividendPerShare           float64 `json:"DividendPerShare,string"`
	DividendYield              float64 `json:"DividendYield,string"`
	EPS                        float64 `json:"EPS,string"`
	RevenuePerShareTTM         float64 `json:"RevenuePerShareTTM,string"`
	ProfitMargin               float64 `json:"ProfitMargin,string"`
	OperatingMarginTTM         float64 `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM          float64 `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM          float64 `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                 int     `json:"RevenueTTM,string"`
	GrossProfitTTM             int     `json:"GrossProfitTTM,string"`
	DilutedEPSTTM              float64 `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY float64 `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY  float64 `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice         float64 `json:"AnalystTargetPrice,string"`
	TrailingPE                 float64 `json:"TrailingPE,string"`
	ForwardPE                  float64 `json:"ForwardPE,string"`
	PriceToSalesRatioTTM       float64 `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio           float64 `json:"PriceToBookRatio,string"`
	EVToRevenue                float64 `json:"EVToRevenue,string"`
	EVToEBITDA                 float64 `json:"EVToEBITDA,string"`
	Beta                       float64 `json:"Beta,string"`
	Week52High                 float64 `json:"52WeekHigh,string"`
	Week52Low                  float64 `json:"52WeekLow,string"`
	MovingAverage50Day         float64 `json:"50DayMovingAverage,string"`
	MovingAverage200Day        float64 `json:"200DayMovingAverage,string"`
	SharesOutstanding          int     `json:"SharesOutstanding,string"`
	DividendDate               string  `json:"DividendDate"`   //  "2023-06-10"
	ExDividendDate             string  `json:"ExDividendDate"` // "2023-05-09"
}

func toCompanyOverview(buf []byte) (*CompanyOverview, error) {
	var overview CompanyOverview
	err := json.Unmarshal(buf, &overview)
	if err != nil {
		return nil, err
	}
	return &overview, nil
}

func (c *Client) CompanyOverview(symbol string) (*CompanyOverview, error) {
	const function = "OVERVIEW"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toCompanyOverview(body)
}
