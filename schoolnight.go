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
		"2020-10-07", // Fall break
		"2020-10-08", // Fall break
		"2020-10-09", // Fall break
		"2020-11-25", // Thanksgiving Break
		"2020-11-26", // Thanksgiving Break
		"2020-11-27", // Thanksgiving Break
		"2020-12-21", // Winter break
		"2020-12-22", // Winter break
		"2020-12-23", // Winter break
		"2020-12-24", // Winter break
		"2020-12-25", // Christmas Day
		"2020-12-28", // Winter Break
		"2020-12-29", // Winter Break
		"2020-12-30", // Winter Break
		"2020-12-31", // Winter Break
		"2021-01-01", // New Year's Day
		"2021-01-04", // Teacher Workday
		"2021-01-18", // MLK Day
		"2021-02-15", // Teacher Workday
		"2021-04-02", // Spring Break
		"2021-04-05", // Spring Break
		"2021-04-06", // Spring Break
		"2021-04-07", // Spring Break
		"2021-04-08", // Spring Break
		"2021-04-09", // Spring Break
		"2021-05-03", // Teacher Workday
		"2021-09-06", // Labor Day
		"2021-10-06", // Fall Break
		"2021-10-07", // Fall Break
		"2021-10-08", // Fall Break
		"2021-11-01", // Teacher Workday
		"2021-11-24", // Thanksgiving
		"2021-11-25", // Thanksgiving
		"2021-11-26", // Thanksgiving
		"2021-12-20", // Winter Break
		"2021-12-21", // Winter Break
		"2021-12-22", // Winter Break
		"2021-12-23", // Winter Break
		"2021-12-24", // Winter Break
		"2021-12-27", // Winter Break
		"2021-12-28", // Winter Break
		"2021-12-29", // Winter Break
		"2021-12-30", // Winter Break
		"2021-12-31", // Winter Break
		"2022-01-03", // Winter Break
		"2022-01-17", // Teacher Workday
		"2022-04-15", // Spring Break
		"2022-04-18", // Spring Break
		"2022-04-19", // Spring Break
		"2022-04-20", // Spring Break
		"2022-04-21", // Spring Break
		"2022-05-02", // Teacher Workday
		"2022-09-05", // Teacher Workday
		"2022-10-05", // Fall Break
		"2022-10-06", // Fall Break
		"2022-10-07", // Fall Break
		"2022-11-01", // Teacher Halloween Safety Workday
		"2022-11-23", // Thanksgiving
		"2022-11-24", // Thanksgiving
		"2022-11-25", // Thanksgiving
		"2022-12-19", // Winter Break
		"2022-12-20", // Winter Break
		"2022-12-21", // Winter Break
		"2022-12-22", // Winter Break
		"2022-12-23", // Winter Break
		"2022-12-26", // Winter Break
		"2022-12-27", // Winter Break
		"2022-12-28", // Winter Break
		"2022-12-29", // Winter Break
		"2022-12-30", // Winter Break
		"2023-01-02", // Teacher Workday
		"2023-01-16", // MLK Teacher Workday
		"2023-02-10", // Teacher Workday
		"2023-04-07", // Teacher Workday
		"2023-04-10", // Teacher Workday
		"2023-04-11", // Teacher Workday
		"2023-04-12", // Teacher Workday
		"2023-04-13", // Teacher Workday
		"2023-04-14", // Teacher Workday

		"2023-08-16", // Teacher Workdays
		"2023-08-17", // Teacher Workdays
		"2023-08-18", // Teacher Workdays
		"2023-08-21", // Mandatory Workdays
		"2023-08-22", // Teacher Workdays
		"2023-08-23", // Teacher Workdays
		"2023-08-24", // Teacher Workdays
		"2023-08-25", // Teacher Workdays
		"2023-09-04", // Holidays
		"2023-10-23", // Teacher Workdays
		"2023-11-07", // Teacher Workdays
		"2023-11-10", // Holidays
		"2023-11-22", // Teacher Workdays
		"2023-11-23", // Holidays
		"2023-11-24", // Holidays
		"2023-12-15", // Excempt teacher exam day
		"2023-12-18", // Excempt teacher exam day
		"2023-12-19", // Excempt teacher exam day
		"2023-12-20", // Excempt teacher exam day
		"2023-12-21", // Excempt teacher exam day
		"2023-12-22", // Excempt teacher exam day
		"2023-12-25", // Holidays
		"2023-12-26", // Holidays
		"2023-12-27", // Holidays
		"2023-12-28", // Vacation
		"2023-12-29", // Vacation
		"2024-01-01", // Holidays
		"2024-01-02", // Vacation
		"2024-01-03", // Vacation
		"2024-01-04", // Vacation
		"2024-01-05", // Teacher Workdays
		"2024-01-08", // Teacher Workdays
		"2024-01-15", // Holidays
		"2024-03-05", // Teacher Workdays
		"2024-03-28", // Teacher Workdays
		"2024-03-29", // Holidays
		"2024-04-01", // Vacation
		"2024-04-02", // Vacation
		"2024-04-03", // Vacation
		"2024-04-04", // Vacation
		"2024-04-05", // Vacation
		"2024-05-03", // Teacher Workdays
		"2024-05-27", // Holidays
		"2024-06-03", // Mandatory Workdays
		"2024-06-04", // Teacher Workdays
		"2024-06-06", // Teacher Workdays
		"2024-06-07", // Teacher Workdays
		"2024-06-10", // Teacher Workdays
		"2024-06-11", // Teacher Workdays
		"2034-06-05", // Teacher Workdays
	}
	summers = map[int]summer{
		2020: {"2020-05-29", "2020-08-09"},
		2021: {"2021-05-29", "2021-08-17"},
		2022: {"2022-05-27", "2022-08-23"},
		2023: {"2023-05-27", "2023-08-27"},
		2024: {"2024-05-31", "2024-08-25"},
	}
	holidayMap map[string]struct{}
)

// summer represents the start and end dates of summer vacation
type summer struct {
	Start string
	End   string
}

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

	// There are no school nights in summer.
	thisYear := time.Now().Year()

	sStart, err := time.Parse("2006-01-02", summers[thisYear].Start)
	if err != nil {
		return false
	}
	sEnd, err := time.Parse("2006-01-02", summers[thisYear].End)
	if err != nil {
		return false
	}
	if when.After(sStart) && when.Before(sEnd) {
		return false
	}

	return true
}
