package alphavantage

// Interval represents the possible interval types
type Interval string

// TimeSeriesInterval is the type for time series data
type TimeSeriesInterval string

// OutPutSize is the type for data output
type OutPutSize string

const (
	// IntervalWeekly - Interval weekly
	IntervalWeekly = Interval("weekly")
	// IntervalDaily - Interval daily
	IntervalDaily = Interval("daily")

	// DateFormat is the date format used by the API
	DateFormat = "2006-01-02"
	// DateTimeFormat datetime format used by the API
	DateTimeFormat = "2006-01-02 15:04:05"

	// TimeSeriesDaily - Time series for daily prices
	TimeSeriesDaily = TimeSeriesInterval("TIME_SERIES_DAILY")
	// TimeSeriesWeekly - Time series for weekly prices
	TimeSeriesWeekly = TimeSeriesInterval("TIME_SERIES_WEEKLY")

	// OutputSizeCompact is for the latest 100 items
	OutputSizeCompact = OutPutSize("compact")
	// OutputSizeFull is for the full-length time series
	OutputSizeFull = OutPutSize("full")
)
