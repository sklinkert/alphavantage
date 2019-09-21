# alphavantage

Unofficial Golang www.alphavantage.co API implementation

Disclaimer: This library is not associated with alphavantage or any of its affiliates or subsidiaries. If you use this library, you should contact them to make sure they are okay with how you intend to use it. Use this lib at your own risk.

API doc reference: https://www.alphavantage.co/documentation/

## Usage

### New client

```go
package main

import (
	"github.com/sklinkert/alphavantage"
)

func main() {
  avClient := alphavantage.New("MYAPIKEY")
  ...
}
```



Indicator STOCH

```go
indicators, err := avClient.IndicatorStoch("EURUSD", alphavantage.IntervalDaily)
if err != nil {
	log.WithError(err).Error("IndicatorStoch() failed")
}

// Loop over all indicators
for date, indicator := range indicators.TechnicalAnalysis {
	log.Infof("%s: SlowK=%f SlowD=%f", date, indicator.SlowK, indicator.SlowD)
}

// Get the most recent one
latestDate, latest := indicators.Latest()
log.Infof("Latest: %s: SlowK=%f SlowD=%f", latestDate, latest.SlowK, latest.SlowD)

// Get today only
today := indicators.Today()
log.Infof("Todat:SlowK=%f SlowD=%f", today.SlowK, today.SlowD)

```

