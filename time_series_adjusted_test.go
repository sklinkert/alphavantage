package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToTimeSeriesAdjusted(t *testing.T) {
	var buf = `
		{
			"Meta Data": {
				"1. Information": "Daily Time Series with Splits and Dividend Events",
				"2. Symbol": "MSFT",
				"3. Last Refreshed": "2020-02-11 15:29:30",
				"4. Output Size": "Full size",
				"5. Time Zone": "US/Eastern"
			},
			"Time Series (Daily)": {
				"2020-02-11": {
					"1. open": "190.6500",
					"2. high": "190.7000",
					"3. low": "183.5000",
					"4. close": "184.3900",
					"5. adjusted close": "184.3900",
					"6. volume": "41524766",
					"7. dividend amount": "3.0000",
					"8. split coefficient": "1.0000"
				},
				"2020-02-10": {
					"1. open": "183.5800",
					"2. high": "188.8400",
					"3. low": "183.2500",
					"4. close": "188.7000",
					"5. adjusted close": "188.7000",
					"6. volume": "32625446",
					"7. dividend amount": "0.0000",
					"8. split coefficient": "1.0000"
				}
			}
    }
`
	timeSeries, err := toTimeSeriesAdjusted([]byte(buf))
	assert.NoError(t, err)
	assert.EqualStrings(t, "MSFT", timeSeries.Metadata.Symbol)
	assert.EqualInt(t, 2, len(timeSeries.TimeSeriesAdjusted))

	ta1, exists := timeSeries.TimeSeriesAdjusted["2020-02-11"]
	if !exists {
		assert.Panic(t, "entry for 2020-02-11 is missing")
	}
	assert.EqualFloat64(t, 190.6500, ta1.Open)
	assert.EqualInt(t, 41524766, int(ta1.Volume))
	assert.EqualFloat64(t, 3.00, ta1.DividendAmount)
	assert.EqualFloat64(t, 1.00, ta1.SplitCoefficient)
	assert.EqualFloat64(t, 184.3900, ta1.AdjustedClose)
}
