package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToIndicatorStoch(t *testing.T) {
	var buf = `
	{
		"Meta Data": {
			"1: Symbol": "MSFT",
			"2: Indicator": "Stochastic (STOCH)",
			"3: Last Refreshed": "2019-09-20",
			"4: Interval": "daily",
			"5.1: FastK Period": 5,
			"5.2: SlowK Period": 3,
			"5.3: SlowK MA Type": 0,
			"5.4: SlowD Period": 3,
			"5.5: SlowD MA Type": 0,
			"6: Time Zone": "US/Eastern Time"
		},
		"Technical Analysis: STOCH": {
			"2019-09-20": {
				"SlowK": "77.3256",
				"SlowD": "76.3691"
			},
			"2019-09-19": {
				"SlowK": "81.5707",
				"SlowD": "69.3987"
			}
		}
    }
`
	indicator, err := toIndicatorStoch([]byte(buf))
	assert.NoError(t, err)
	assert.EqualStrings(t, "MSFT", indicator.Metadata.Symbol)
	assert.EqualInt(t, 2, len(indicator.TechnicalAnalysis))

	ta1, exists := indicator.TechnicalAnalysis["2019-09-20"]
	if !exists {
		assert.Panic(t, "entry for 2019-09-20 is missing")
	}
	assert.EqualFloat64(t, 77.3256, ta1.SlowK)
	assert.EqualFloat64(t, 76.3691, ta1.SlowD)
}
