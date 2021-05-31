package schoolnight

import (
	"testing"
	"time"
)

type testSet struct {
	Name  string
	Month time.Month
	Day   int
	want  bool
}

func TestCheck(t *testing.T) {
	// Using the current year will have the nice side effect that the test will
	// start to fail if this package isn't kept up to date.
	year := time.Now().Year()

	testData := []testSet{
		{
			Name:  "Christmas Eve",
			Month: time.December,
			Day:   24,
			want:  false,
		},
		{
			Name:  "New Year's Eve",
			Month: time.December,
			Day:   31,
			want:  false,
		},
		{
			Name:  "Thanksgiving Monday",
			Month: time.November,
			Day:   thanksgivingDay().AddDate(0, 0, -3).Day(),
			want:  true,
		},
		{
			Name:  "A Summer Wed in July",
			Month: time.July,
			Day:   summerWednesday().Day(),
			want:  false,
		},
	}
	for _, ts := range testData {
		testDay := time.Date(year, ts.Month, ts.Day, 0, 0, 0, 0, time.Local)
		t.Logf("Testing day: %s", testDay)
		got := Check(testDay)
		if ts.want != got {
			t.Errorf("%s: Check(%s), want: %t, got: %t", ts.Name, testDay, ts.want, got)
		}
	}
}

func TestSummersParse(t *testing.T) {
	for year, s := range summers {
		if _, err := time.Parse("2006-01-02", summers[year].Start); err != nil {
			t.Errorf("Summer Start for %d of %q does not parse: %s", year, s.Start, err)
		}
		if _, err := time.Parse("2006-01-02", summers[year].End); err != nil {
			t.Errorf("Summer End for %d of %q does not parse: %s", year, s.End, err)
		}
	}
}

// returns the time.Time of Thanksgiving Day, the 4th Thursday in November
func thanksgivingDay() time.Time {
	o := time.Date(time.Now().Year(), time.November, 1, 0, 0, 0, 0, time.Local)
	d := ((11 - int(o.Weekday())) % 7) // first Thursday in November
	return o.AddDate(0, 0, 21+d)       // 3 thursdays later
}

// summerWednesday returns the first Wed in July of the current year, as a
// representative day in Summer
func summerWednesday() time.Time {
	o := time.Date(time.Now().Year(), time.July, 1, 0, 0, 0, 0, time.Local)
	d := ((10 - int(o.Weekday())) % 7) // first Thursday in November
	return o.AddDate(0, 0, d)
}
