package evaluator

import (
	"time"
)

func (e *evaluator) beginningOfSecond() time.Time {
	return e.tmp.Truncate(time.Second)
}

// BeginningOfMinute beginning of minute
func (e *evaluator) beginningOfMinute() time.Time {
	return e.tmp.Truncate(time.Minute)
}

// BeginningOfHour beginning of hour
func (e *evaluator) beginningOfHour() time.Time {
	y, m, d := e.tmp.Date()
	return time.Date(y, m, d, e.tmp.Hour(), 0, 0, 0, e.tmp.Location())
}

// BeginningOfDay beginning of day
func (e *evaluator) beginningOfDay() time.Time {
	y, m, d := e.tmp.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, e.tmp.Location())
}

// BeginningOfWeek beginning of week
func (e *evaluator) beginningOfWeek() time.Time {
	// t := e.BeginningOfDay()
	// weekday := int(t.Weekday())

	// if e.WeekStartDay != time.Sunday {
	// 	weekStartDayInt := int(e.WeekStartDay)

	// 	if weekday < weekStartDayInt {
	// 		weekday = weekday + 7 - weekStartDayInt
	// 	} else {
	// 		weekday = weekday - weekStartDayInt
	// 	}
	// }
	// return t.AddDate(0, 0, -weekday)
	return time.Now()
}

// BeginningOfMonth beginning of month
func (e *evaluator) beginningOfMonth() time.Time {
	y, m, _ := e.tmp.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, e.Location())
}

// BeginningOfQuarter beginning of quarter
func (e *evaluator) beginningOfQuarter() time.Time {
	month := e.beginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

// BeginningOfHalf beginning of half year
func (e *evaluator) beginningOfHalf() time.Time {
	month := e.beginningOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

// BeginningOfYear BeginningOfYear beginning of year
func (e *evaluator) beginningOfYear() time.Time {
	y, _, _ := e.tmp.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, e.Location())
}

// EndOfMinute end of minute
func (e *evaluator) endOfMinute() time.Time {
	return e.beginningOfMinute().Add(time.Minute - time.Nanosecond)
}

// EndOfHour end of hour
func (e *evaluator) endOfHour() time.Time {
	return e.beginningOfHour().Add(time.Hour - time.Nanosecond)
}

// EndOfDay end of day
func (e *evaluator) endOfDay() time.Time {
	y, m, d := e.tmp.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), e.Location())
}

// EndOfWeek end of week
func (e *evaluator) endOfWeek() time.Time {
	return e.beginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// EndOfMonth end of month
func (e *evaluator) endOfMonth() time.Time {
	return e.beginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfQuarter end of quarter
func (e *evaluator) endOfQuarter() time.Time {
	return e.beginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

// EndOfHalf end of half year
func (e *evaluator) endOfHalf() time.Time {
	return e.beginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond)
}

// EndOfYear end of year
func (e *evaluator) endOfYear() time.Time {
	return e.beginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// Monday monday
func (e *evaluator) Monday() time.Time {
	t := e.beginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return t.AddDate(0, 0, -weekday+1)
}

// Sunday sunday
func (e *evaluator) Sunday() time.Time {
	t := e.beginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		return t
	}
	return t.AddDate(0, 0, 7-weekday)
}

// TODO!! EndOfSunday end of sunday
func (e *evaluator) EndOfSunday() time.Time {
	//  return New(e.Sunday()).EndOfDay()
	return time.Time{}
}
