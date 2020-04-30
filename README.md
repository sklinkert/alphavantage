# alphavantage

**Unofficial** Go/Golang www.alphavantage.co API implementation

**Disclaimer**: This library is not associated with alphavantage or any of its affiliates or subsidiaries. If you use this library, you should contact them to make sure they are okay with how you intend to use it. Use this lib at your own risk.

API doc reference: https://www.alphavantage.co/documentation/

**Note**: Requests are synchronised and throttled automatically to not flood the API servers. One request every 15 seconds is possible at the moment.

## Usage

### New client

```go
package main

import (
	"github.com/sklinkert/alphavantage"
	log "github.com/sirupsen/logrus" // optional
)

func main() {
	avClient := alphavantage.New("MYAPIKEY")
 	// ...
}
```

### TimeSeries (prices)

```go
series, err := avClient.TimeSeries("SIX2.DEX", alphavantage.TimeSeriesDaily, alphavantage.OutPutSizeCompact)
if err != nil {
	log.WithError(err).Fatal("TimeSeries() failed")
}
for date, price := range series.TimeSeriesDaily {
	log.Infof("%s: Open=%f High=%f Low=%f Close=%f Volume=%d", date, price.Open, price.High, price.Low, price.Close, price.Volume)
}
```

### Indicator STOCH

```go
indicators, err := avClient.IndicatorStoch("EURUSD", alphavantage.IntervalDaily)
if err != nil {
	log.WithError(err).Fatal("IndicatorStoch() failed")
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
log.Infof("Today: SlowK=%f SlowD=%f", today.SlowK, today.SlowD)

// By specific date
indicator := indicators.ByDate(time.Now())
if indicator != nil {
	log.Infof("SlowK=%f SlowD=%f", indicator.SlowK, indicator.SlowD)
}

```

### Indicator SMA

```go
perdiod := 200
indicators, err := avClient.IndicatorSMA("EURUSD", alphavantage.IntervalDaily, period)
if err != nil {
	log.WithError(err).Fatal("IndicatorSMA() failed")
}
```
