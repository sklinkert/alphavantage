package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToGlobalQuote(t *testing.T) {
	var buf = `
	{
		"Global Quote": {
			"01. symbol": "IBM",
			"02. open": "126.5200",
			"03. high": "127.2700",
			"04. low": "125.2200",
			"05. price": "125.8400",
			"06. volume": "3792418",
			"07. latest trading day": "2020-04-30",
			"08. previous close": "128.6900",
			"09. change": "-2.8500",
			"10. change percent": "-2.2146%"
		}
	}
`
	globalQuote, err := toGlobalQuote([]byte(buf))
	assert.NoError(t.Fatalf, err)
	assert.EqualStrings(t, "IBM", globalQuote.Symbol)
	assert.EqualStrings(t, "2020-04-30", globalQuote.LatestTradingDay)
	assert.EqualFloat64(t, 125.8400, globalQuote.Price)
	assert.EqualFloat64(t, 125.2200, globalQuote.Low)
	assert.EqualFloat64(t, 127.2700, globalQuote.High)
	assert.EqualFloat64(t, 126.5200, globalQuote.Open)
	assert.EqualFloat64(t, -2.8500, globalQuote.Change)
	// assert.EqualFloat64(t, -2.2146, globalQuote.ChangePercent)
	assert.EqualInt(t, 3792418, globalQuote.Volume)
}
