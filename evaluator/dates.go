package evaluator

import (
	"time"
)

func BeginningOfSecond(date time.Time) time.Time {
	return date.Truncate(time.Second)
}

// BeginningOfMinute beginning of minute
func BeginningOfMinute(date time.Time) time.Time {
	return date.Truncate(time.Minute)
}

// BeginningOfHour beginning of hour
func BeginningOfHour(date time.Time) time.Time {
	y, m, d := date.Date()
	return time.Date(y, m, d, date.Hour(), 0, 0, 0, date.Location())
}

// BeginningOfDay beginning of day
func BeginningOfDay(date time.Time) time.Time {
	y, m, d := date.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, date.Location())
}

// BeginningOfWeek beginning of week
func BeginningOfWeek(date time.Time, weekStartDay time.Weekday) time.Time {
	t := BeginningOfDay(date)
	weekday := int(t.Weekday())

	if weekStartDay != time.Sunday {
		weekStartDayInt := int(weekStartDay)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	return t.AddDate(0, 0, -weekday)
}

// BeginningOfMonth beginning of month
func BeginningOfMonth(date time.Time) time.Time {
	y, m, _ := date.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, date.Location())
}

// BeginningOfYear BeginningOfYear beginning of year
func BeginningOfYear(date time.Time) time.Time {
	y, _, _ := date.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, date.Location())
}

// EndOfMinute end of minute
func EndOfMinute(date time.Time) time.Time {
	return BeginningOfMinute(date).Add(time.Minute - time.Nanosecond)
}

// EndOfHour end of hour
func EndOfHour(date time.Time) time.Time {
	return BeginningOfHour(date).Add(time.Hour - time.Nanosecond)
}

// EndOfDay end of day
func EndOfDay(date time.Time) time.Time {
	y, m, d := date.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), date.Location())
}

// EndOfWeek end of week
func EndOfWeek(date time.Time, weekStartAt time.Weekday) time.Time {
	return BeginningOfWeek(date, weekStartAt).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// EndOfMonth end of month
func EndOfMonth(date time.Time) time.Time {
	return BeginningOfMonth(date).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfYear end of year
func EndOfYear(date time.Time) time.Time {
	return BeginningOfYear(date).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// Monday monday
func PreviousMonday(date time.Time) time.Time {
	t := BeginningOfDay(date)
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return t.AddDate(0, 0, -weekday+1)
}

// Sunday sunday
func PreviousSunday(date time.Time) time.Time {
	t := BeginningOfDay(date)
	weekday := int(t.Weekday())
	if weekday == 0 {
		return t
	}
	return t.AddDate(0, 0, 7-weekday)
}

// Next sunday
func EndOfSunday(date time.Time) time.Time {
	return EndOfDay(PreviousSunday(date))
}
