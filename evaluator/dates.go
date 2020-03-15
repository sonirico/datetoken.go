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
	return PreviousWeekDay(BeginningOfDay(date), weekStartDay)
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

// EndOfBusinessWeek end of week
func EndOfBusinessWeek(date time.Time, weekStartAt time.Weekday) time.Time {
	return BeginningOfWeek(date, weekStartAt).AddDate(0, 0, 5).Add(-time.Nanosecond)
}

// EndOfMonth end of month
func EndOfMonth(date time.Time) time.Time {
	return BeginningOfMonth(date).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfYear end of year
func EndOfYear(date time.Time) time.Time {
	return BeginningOfYear(date).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

func AddSeconds(date time.Time, amount int) time.Time {
	return date.Add(time.Duration(amount) * time.Second)
}

func AddMinutes(date time.Time, amount int) time.Time {
	return date.Add(time.Duration(amount) * time.Minute)
}

func AddHours(date time.Time, amount int) time.Time {
	return date.Add(time.Duration(amount) * time.Hour)
}

func AddDays(date time.Time, amount int) time.Time {
	return date.Add(time.Duration(amount) * 24 * time.Hour)
}

func AddWeeks(date time.Time, amount int) time.Time {
	return date.Add(time.Duration(amount) * 24 * 7 * time.Hour)
}

func AddMonths(date time.Time, amount int) time.Time {
	return date.AddDate(0, amount, 0)
}

func AddYears(date time.Time, amount int) time.Time {
	return date.AddDate(amount, 0, 0)
}

func PreviousWeekDay(date time.Time, targetWeekDay time.Weekday) time.Time {
	wd := int(date.Weekday())
	twd := int(targetWeekDay)
	return date.AddDate(0, 0, -(((wd-twd)%7 + 7) % 7))
}

func PreviousMonday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Monday)
}

func PreviousTuesday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Tuesday)
}

func PreviousWednesday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Wednesday)
}

func PreviousThursday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Thursday)
}

func PreviousFriday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Friday)
}

func PreviousSaturday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Saturday)
}

func PreviousSunday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Sunday)
}

func NextWeekDay(date time.Time, targetWeekDay time.Weekday) time.Time {
	wd := int(date.Weekday())
	twd := int(targetWeekDay)
	return date.AddDate(0, 0, ((twd-wd)%7+7)%7)
}

func NextMonday(date time.Time) time.Time {
	return NextWeekDay(date, time.Monday)
}

func NextTuesday(date time.Time) time.Time {
	return NextWeekDay(date, time.Tuesday)
}

func NextWednesday(date time.Time) time.Time {
	return NextWeekDay(date, time.Wednesday)
}

func NextThursday(date time.Time) time.Time {
	return NextWeekDay(date, time.Thursday)
}

func NextFriday(date time.Time) time.Time {
	return NextWeekDay(date, time.Friday)
}

func NextSaturday(date time.Time) time.Time {
	return NextWeekDay(date, time.Saturday)
}

func NextSunday(date time.Time) time.Time {
	return NextWeekDay(date, time.Sunday)
}
