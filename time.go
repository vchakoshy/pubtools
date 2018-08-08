package pubtools

import (
	"time"
)

// TimeLib struct
type TimeLib struct{}

// Time returns TimeLib
func Time() *TimeLib {
	return &TimeLib{}
}

func (t *TimeLib) GetDefaultLocation() *time.Location {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		loc = time.FixedZone("Asia/Tehran", 12600)
	}
	return loc
}

func (t *TimeLib) GetNowLocation() (time.Time, *time.Location) {
	loc := t.GetDefaultLocation()
	n := time.Now()
	return n, loc
}

func (t *TimeLib) GetNowYear() time.Time {
	n, loc := t.GetNowLocation()
	return time.Date(n.Year(), 0, 0, 0, 0, 0, 0, loc)
}

func (t *TimeLib) GetNowMonth() time.Time {
	n, loc := t.GetNowLocation()
	return time.Date(n.Year(), n.Month(), 0, 0, 0, 0, 0, loc)
}

func (t *TimeLib) GetNowToday() time.Time {
	n, loc := t.GetNowLocation()
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, loc)
}

func (t *TimeLib) GetNowAsHour() time.Time {
	n, loc := t.GetNowLocation()
	return time.Date(n.Year(), n.Month(), n.Day(), n.Hour(), 0, 0, 0, loc)
}

func (t *TimeLib) GetNowAsMinute() time.Time {
	n, loc := t.GetNowLocation()
	return time.Date(n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), 0, 0, loc)
}

func (t *TimeLib) GetInDefaultLocation(tt time.Time) time.Time {
	loc := t.GetDefaultLocation()
	return tt.In(loc)
}

func (t *TimeLib) NowInDefault() time.Time {
	n, loc := t.GetNowLocation()
	return n.In(loc)
}

func (t *TimeLib) GetYesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

func (t *TimeLib) GetLastMonth() time.Time {
	return time.Now().AddDate(0, 0, -30)
}
