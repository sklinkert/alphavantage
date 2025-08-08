package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToCurrencyExchangeRate(t *testing.T) {
	var buf = `
{
    "Realtime Currency Exchange Rate": {
        "1. From_Currency Code": "BTC",
        "2. From_Currency Name": "Bitcoin",
        "3. To_Currency Code": "EUR",
        "4. To_Currency Name": "Euro",
        "5. Exchange Rate": "98897.68000000",
        "6. Last Refreshed": "2025-08-06 20:31:02",
        "7. Time Zone": "UTC",
        "8. Bid Price": "98893.39700000",
        "9. Ask Price": "98900.98300000"
    }
}
`
	exchangeRate, err := toCurrencyExchangeRate([]byte(buf))
	assert.NoError(t.Fatalf, err)

	assert.EqualStrings(t, "BTC", exchangeRate.FromCurrencyCode)
	assert.EqualStrings(t, "Bitcoin", exchangeRate.FromCurrencyName)
	assert.EqualStrings(t, "EUR", exchangeRate.ToCurrencyCode)
	assert.EqualStrings(t, "Euro", exchangeRate.ToCurrencyName)
	assert.EqualStrings(t, "98897.68000000", exchangeRate.ExchangeRate)
	assert.EqualStrings(t, "2025-08-06 20:31:02", exchangeRate.LastRefreshed)
	assert.EqualStrings(t, "UTC", exchangeRate.TimeZone)
	assert.EqualStrings(t, "98893.39700000", exchangeRate.BidPrice)
	assert.EqualStrings(t, "98900.98300000", exchangeRate.AskPrice)
}
