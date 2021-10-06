package alphavantage

import (
	"encoding/json"
	"strconv"
)

// Interval represents the possible interval types
type Interval string

// TimeSeriesInterval is the type for time series data
type TimeSeriesInterval string

// OutPutSize is the type for data output
type OutPutSize string

// Create new type based int64 type so that we can build custom Unmarshaller
type AVInt int

// Custom Unmarshaller for AVInt64 to handle int with JSON value on "N"
func (st *AVInt) UnmarshalJSON(b []byte) error {
	//convert the bytes into an interface
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = AVInt(v)
	case float64:
		*st = AVInt(int(v))
	// attempt string conversion to int and handle "None" special case
	case string:
		///here convert the string into
		///an integer
		i, err := strconv.Atoi(v)
		if err != nil {
			if v == "None" {
				i = 0
			} else {
				///the string might not be of integer type
				///so return an error
				return err
			}
		}
		*st = AVInt(i)

	}
	return nil
}

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
