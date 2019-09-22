package alphavantage

// Interval represents the possible interval types
type Interval string

// TimeSeriesInterval is the type for time series data
type TimeSeriesInterval string

const (
	// IntervalWeekly - Interval weekly
	IntervalWeekly = Interval("weekly")
	// IntervalDaily - Interval daily
	IntervalDaily = Interval("daily")

	dateLayout = "2006-01-02"

	// TimeSeriesDaily - Time series for daily prices
	TimeSeriesDaily = TimeSeriesInterval("TIME_SERIES_DAILY")
	// TimeSeriesWeekly - Time series for weekly prices
	TimeSeriesWeekly = TimeSeriesInterval("TIME_SERIES_WEEKLY")
)
