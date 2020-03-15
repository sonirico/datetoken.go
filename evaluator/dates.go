package evaluator

import (
	"time"
)

// BeginningOfSecond returns the given date with nanoseconds truncated to 0
func BeginningOfSecond(date time.Time) time.Time {
	return date.Truncate(time.Second)
}

// BeginningOfMinute returns the given date with seconds truncated to 0
func BeginningOfMinute(date time.Time) time.Time {
	return date.Truncate(time.Minute)
}

// BeginningOfHour returns the given date with minutes and seconds truncated 0
func BeginningOfHour(date time.Time) time.Time {
	y, m, d := date.Date()
	return time.Date(y, m, d, date.Hour(), 0, 0, 0, date.Location())
}

// BeginningOfDay returns the given date with hours, minutes and seconds truncated 0
func BeginningOfDay(date time.Time) time.Time {
	y, m, d := date.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, date.Location())
}

// BeginningOfWeek returns the given date with hours, minutes and seconds truncated to 0. The day
// will be set to the day of the week at which the week has started, defined by 'weekStartDay'
func BeginningOfWeek(date time.Time, weekStartDay time.Weekday) time.Time {
	return PreviousWeekDay(BeginningOfDay(date), weekStartDay)
}

// BeginningOfMonth returns the given date to the start of the month, with day as 1, and hours, minutes
// and seconds truncated to 0
func BeginningOfMonth(date time.Time) time.Time {
	y, m, _ := date.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, date.Location())
}

// BeginningOfYear returns the given date to the start of the year, with month and day set to 1. Hours, minutes
// and seconds are truncated to 0
func BeginningOfYear(date time.Time) time.Time {
	y, _, _ := date.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, date.Location())
}

// EndOfMinute returns the given day to the end of the minute, with seconds truncated to 59
func EndOfMinute(date time.Time) time.Time {
	return BeginningOfMinute(date).Add(time.Minute - time.Nanosecond)
}

// EndOfHour returns the given day to the end of the hour, with minutes and seconds truncated to 59
func EndOfHour(date time.Time) time.Time {
	return BeginningOfHour(date).Add(time.Hour - time.Nanosecond)
}

// EndOfDay returns the given day to the end of the day, with minutes and seconds truncated to 59 and the
// hour truncated to 23
func EndOfDay(date time.Time) time.Time {
	y, m, d := date.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), date.Location())
}

// EndOfWeek returns the given date snapped to end of the week, taking in account that the start of the week
// is configurable
func EndOfWeek(date time.Time, weekStartAt time.Weekday) time.Time {
	return BeginningOfWeek(date, weekStartAt).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// EndOfBusinessWeek as EndOfWeek returns the date snapped to its end, but instead of adding 7 days, only 5
// will be add. THOUGHT: configurable too?
func EndOfBusinessWeek(date time.Time, weekStartAt time.Weekday) time.Time {
	return BeginningOfWeek(date, weekStartAt).AddDate(0, 0, 5).Add(-time.Nanosecond)
}

// EndOfMonth returns the given date snapped to the end of the month, being the day 28 or 29 for February, 30
// for April, June, September and November, and 31 for the rest. Hours are truncated to 23, whereas minutes and
// seconds will be truncated to 59
func EndOfMonth(date time.Time) time.Time {
	return BeginningOfMonth(date).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfYear returns the given date snapped to the end of the year, being the day 31th of December. Hours are
// truncated to 23, whereas minutes and seconds will be truncated to 59
func EndOfYear(date time.Time) time.Time {
	return BeginningOfYear(date).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// AddSeconds adds any amount of seconds to the given date
func AddSeconds(date time.Time, amount int) time.Time {
	return date.Add(time.Duration(amount) * time.Second)
}

// AddMinutes adds any amount of minutes to the given date
func AddMinutes(date time.Time, amount int) time.Time {
	return date.Add(time.Duration(amount) * time.Minute)
}

// AddHours adds any amount of hours to the given date
func AddHours(date time.Time, amount int) time.Time {
	return date.Add(time.Duration(amount) * time.Hour)
}

// AddDays adds any amount of days to the given date
func AddDays(date time.Time, amount int) time.Time {
	return date.Add(time.Duration(amount) * 24 * time.Hour)
}

// AddWeeks adds any amount of weeks to the given date
func AddWeeks(date time.Time, amount int) time.Time {
	return date.Add(time.Duration(amount) * 24 * 7 * time.Hour)
}

// AddMonths adds any amount of months to the given date
func AddMonths(date time.Time, amount int) time.Time {
	return date.AddDate(0, amount, 0)
}

// AddYears adds any amount of years to the given date
func AddYears(date time.Time, amount int) time.Time {
	return date.AddDate(amount, 0, 0)
}

// PreviousWeekDay returns the previous start of the week according to what it is understood to be as the
// start of the week, defined by 'targetWeekDay' parameter. If the target start of the week is the same as
// the current week, the current date is considered to be the current start of the week
func PreviousWeekDay(date time.Time, targetWeekDay time.Weekday) time.Time {
	wd := int(date.Weekday())
	twd := int(targetWeekDay)
	return date.AddDate(0, 0, -(((wd-twd)%7 + 7) % 7))
}

// PreviousMonday returns the date of the previous Monday
func PreviousMonday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Monday)
}

// PreviousTuesday returns the date of the previous Tuesday
func PreviousTuesday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Tuesday)
}

// PreviousWednesday returns the date of the previous Wednesday
func PreviousWednesday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Wednesday)
}

// PreviousThursday returns the date of the previous Thursday
func PreviousThursday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Thursday)
}

// PreviousFriday returns the date of the previous Friday
func PreviousFriday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Friday)
}

// PreviousSaturday returns the date of the previous Saturday
func PreviousSaturday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Saturday)
}

// PreviousSunday returns the date of the previous Sunday
func PreviousSunday(date time.Time) time.Time {
	return PreviousWeekDay(date, time.Sunday)
}

// NextWeekDay returns the next start of the week according to what it is understood to be as the
// start of the week, defined by 'targetWeekDay' parameter. If the target start of the week is the same as
// the current week, the current date is considered to be the current start of the week
func NextWeekDay(date time.Time, targetWeekDay time.Weekday) time.Time {
	wd := int(date.Weekday())
	twd := int(targetWeekDay)
	return date.AddDate(0, 0, ((twd-wd)%7+7)%7)
}

// NextMonday returns the date of the next Monday
func NextMonday(date time.Time) time.Time {
	return NextWeekDay(date, time.Monday)
}

// NextTuesday returns the date of the next Tuesday
func NextTuesday(date time.Time) time.Time {
	return NextWeekDay(date, time.Tuesday)
}

// NextWednesday returns the date of the next Wednesday
func NextWednesday(date time.Time) time.Time {
	return NextWeekDay(date, time.Wednesday)
}

// NextThursday returns the date of the next Thursday
func NextThursday(date time.Time) time.Time {
	return NextWeekDay(date, time.Thursday)
}

// NextFriday returns the date of the next Friday
func NextFriday(date time.Time) time.Time {
	return NextWeekDay(date, time.Friday)
}

// NextSaturday returns the date of the next Saturday
func NextSaturday(date time.Time) time.Time {
	return NextWeekDay(date, time.Saturday)
}

// NextSunday returns the date of the next Sunday
func NextSunday(date time.Time) time.Time {
	return NextWeekDay(date, time.Sunday)
}
