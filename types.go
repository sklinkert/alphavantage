package alphavantage

// Interval represents the possible interval types
type Interval string

const (
	// IntervalWeekly - Interval weekly
	IntervalWeekly = Interval("weekly")
	// IntervalDaily - Interval daily
	IntervalDaily = Interval("daily")
)
