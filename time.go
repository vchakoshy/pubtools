package pubtools

import (
	"time"
)

// TimeLib struct
type TimeLib struct {
	loc *time.Location
	now time.Time
}

// Time returns TimeLib
func Time() *TimeLib {
	return &TimeLib{
		now: time.Now(),
		loc: getTimeDefaultLocation(),
	}
}

func getTimeDefaultLocation() *time.Location {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		loc = time.FixedZone("Asia/Tehran", 12600)
	}
	return loc
}

func (t *TimeLib) GetNowYear() time.Time {
	return time.Date(t.now.Year(), 0, 0, 0, 0, 0, 0, t.loc)
}

func (t *TimeLib) GetNowMonth() time.Time {
	return time.Date(t.now.Year(), t.now.Month(), 0, 0, 0, 0, 0, t.loc)
}

func (t *TimeLib) GetNowToday() time.Time {
	return time.Date(t.now.Year(), t.now.Month(), t.now.Day(), 0, 0, 0, 0, t.loc)
}

func (t *TimeLib) GetNowAsHour() time.Time {
	return time.Date(t.now.Year(), t.now.Month(), t.now.Day(), t.now.Hour(), 0, 0, 0, t.loc)
}

func (t *TimeLib) GetNowAsMinute() time.Time {
	return time.Date(t.now.Year(), t.now.Month(), t.now.Day(), t.now.Hour(), t.now.Minute(), 0, 0, t.loc)
}

func (t *TimeLib) GetInDefaultLocation(tt time.Time) time.Time {
	return tt.In(t.loc)
}

func (t *TimeLib) NowInDefault() time.Time {
	return t.now.In(t.loc)
}

func (t *TimeLib) GetYesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

func (t *TimeLib) GetLastMonth() time.Time {
	return time.Now().AddDate(0, 0, -30)
}
