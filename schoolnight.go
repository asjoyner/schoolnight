package schoolnight

import (
	"fmt"
	"time"
)

var (
	holidays = []string{
		"2020-01-01", // New Years
		"2020-01-02",
		"2020-01-03",
		"2020-01-20", // Workday
		"2020-02-10", // Workday
		"2020-04-10", // Spring Break
		"2020-05-08", // Workday
		"2020-05-25", // Memorial Day
		"2020-06-04", // First day of summer
		"2020-12-25", // Christmas Day
		"2021-01-01", // New Year's Day
	}
	summerStart  = "2020-05-29"
	summerEnd    = "2020-08-09"
	sStart, sEnd time.Time
	holidayMap   map[string]struct{}
)

func init() {
	// calculate the day before holidays as "not school nights", and convert the
	// list into a map for faster lookups along the way
	holidayMap = make(map[string]struct{}, len(holidays))
	for _, hs := range holidays {
		t, err := time.Parse("2006-01-02", hs)
		if err != nil {
			panic(err)
		}
		b := t.AddDate(0, 0, -1)
		beforeStr := fmt.Sprintf("%d-%d-%d", b.Year(), b.Month(), b.Day())
		holidayMap[beforeStr] = struct{}{}
	}
	var err error
	sStart, err = time.Parse("2006-01-02", summerStart)
	if err != nil {
		panic(err)
	}
	sEnd, err = time.Parse("2006-01-02", summerEnd)
	if err != nil {
		panic(err)
	}
}

// Check returns true if 'when' falls on a school night.
func Check(when time.Time) bool {
	// weekends are never school nights
	if when.Weekday() >= 5 {
		return false
	}

	// if the next day is a holiday, it's not a school night
	wt := fmt.Sprintf("%d-%d-%d", when.Year(), when.Month(), when.Day())
	if _, ok := holidayMap[wt]; ok {
		return false
	}

	if when.After(sStart) && when.Before(sEnd) {
		return false
	}

	return true
}
