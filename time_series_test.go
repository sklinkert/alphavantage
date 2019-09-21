package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToTimeSeries(t *testing.T) {
	var buf = `
	{
			"Meta Data": {
				"1. Information": "Weekly Prices (open, high, low, close) and Volumes",
				"2. Symbol": "SIX2.DEX",
				"3. Last Refreshed": "2019-09-20",
				"4. Time Zone": "US/Eastern"
			},
			"Weekly Time Series": {
				"2019-09-20": {
					"1. open": "93.2500",
					"2. high": "94.2000",
					"3. low": "89.5500",
					"4. close": "90.3500",
					"5. volume": "199054"
				},
				"2019-09-13": {
					"1. open": "92.3000",
					"2. high": "95.4000",
					"3. low": "91.5000",
					"4. close": "94.4000",
					"5. volume": "254033"
				}
			}
    }
`
	timeSeries, err := toTimeSeries([]byte(buf))
	assert.NoError(t, err)
	assert.EqualStrings(t, "SIX2.DEX", timeSeries.Metadata.Symbol)
	assert.EqualInt(t, 2, len(timeSeries.TimeSeriesWeekly))

	ta1, exists := timeSeries.TimeSeriesWeekly["2019-09-20"]
	if !exists {
		assert.Panic(t, "entry for 2019-09-20 is missing")
	}
	assert.EqualFloat64(t, 93.2500, ta1.Open)
	assert.EqualInt(t, 199054, int(ta1.Volume))
}
