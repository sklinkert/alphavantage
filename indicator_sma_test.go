package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToIndicatorSMA(t *testing.T) {
	var buf = `
	{
		"Meta Data": {
			"1: Symbol": "USDEUR",
			"2: Indicator": "Simple Moving Average (SMA)",
			"3: Last Refreshed": "2020-04-30",
			"4: Interval": "weekly",
			"5: Time Period": 10,
			"6: Series Type": "open",
			"7: Time Zone": "US/Eastern"
		},
		"Technical Analysis: SMA": {
			"2020-04-30": {
				"SMA": "0.9120"
			},
			"2020-04-24": {
				"SMA": "0.9118"
			},
			"2020-04-17": {
				"SMA": "0.9111"
			},
			"2020-04-10": {
				"SMA": "0.9098"
			},
			"2020-04-03": {
				"SMA": "0.9080"
			}
		}
    }
`
	indicator, err := toIndicatorSMA([]byte(buf))
	assert.NoError(t, err)
	assert.EqualStrings(t, "USDEUR", indicator.Metadata.Symbol)
	assert.EqualInt(t, 5, len(indicator.TechnicalAnalysis))

	ta1, exists := indicator.TechnicalAnalysis["2020-04-24"]
	if !exists {
		assert.Panic(t, "entry for 2020-04-24 is missing")
	}
	assert.EqualFloat64(t, 0.9118, ta1.SMA)
}
