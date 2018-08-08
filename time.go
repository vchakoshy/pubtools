package pubtools

import (
	"time"
)

func GetTimeDefaultLocation() *time.Location {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		loc = time.FixedZone("Asia/Tehran", 12600)
	}
	return loc
}

func GetTimeNowAndLocation() (time.Time, *time.Location) {
	loc := GetTimeDefaultLocation()
	n := time.Now()
	return n, loc
}

func GetTimeNowYear() time.Time {
	n, loc := GetTimeNowAndLocation()
	t := time.Date(n.Year(), 0, 0, 0, 0, 0, 0, loc)
	return t
}

func GetTimeNowMonth() time.Time {
	n, loc := GetTimeNowAndLocation()
	t := time.Date(n.Year(), n.Month(), 0, 0, 0, 0, 0, loc)
	return t
}

func GetTimeNowToday() time.Time {
	n, loc := GetTimeNowAndLocation()
	t := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, loc)
	return t
}

func GetTimeNowAsHour() time.Time {
	n, loc := GetTimeNowAndLocation()
	t := time.Date(n.Year(), n.Month(), n.Day(), n.Hour(), 0, 0, 0, loc)
	return t
}

func GetTimeNowAsMinute() time.Time {
	n, loc := GetTimeNowAndLocation()
	t := time.Date(n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), 0, 0, loc)
	return t
}

func GetTimeInDefaultLocation(t time.Time) time.Time {
	loc := GetTimeDefaultLocation()
	return t.In(loc)
}

func NowInDefault() time.Time {
	n, loc := GetTimeNowAndLocation()
	return n.In(loc)
}

func GetTimeYesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

func GetTimeLastMonth() time.Time {
	return time.Now().AddDate(0, 0, -30)
}
