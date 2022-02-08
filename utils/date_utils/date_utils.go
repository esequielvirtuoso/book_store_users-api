// Package dateUtils is an util package responsible for handling dates.
// Every time it is necessary to work with dates, this module must be called
// so we can keep the date times standardized and a single point of failure.
package dateUtils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

// GetNow returns the current date and time in Cordinated Universal Time (UTC).
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns the current date and time in UTC with the apiDateLayout format.
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDbFormat returns the current date and time in UTC with the apiDbLayout format.
func GetNowDbFormat() string {
	return GetNow().Format(apiDbLayout)
}
