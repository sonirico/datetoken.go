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
func (e *evaluator) BeginningOfHour() time.Time {
	y, m, d := e.Date()
	return time.Date(y, m, d, e.Time.Hour(), 0, 0, 0, e.Time.Location())
}

// BeginningOfDay beginning of day
func (e *evaluator) BeginningOfDay() time.Time {
	y, m, d := e.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, e.Time.Location())
}

// BeginningOfWeek beginning of week
func (e *evaluator) BeginningOfWeek() time.Time {
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
func (e *evaluator) BeginningOfMonth() time.Time {
	y, m, _ := e.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, e.Location())
}

// BeginningOfQuarter beginning of quarter
func (e *evaluator) BeginningOfQuarter() time.Time {
	month := e.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

// BeginningOfHalf beginning of half year
func (e *evaluator) BeginningOfHalf() time.Time {
	month := e.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

// BeginningOfYear BeginningOfYear beginning of year
func (e *evaluator) BeginningOfYear() time.Time {
	y, _, _ := e.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, e.Location())
}

// EndOfMinute end of minute
func (e *evaluator) EndOfMinute() time.Time {
	return e.beginningOfMinute().Add(time.Minute - time.Nanosecond)
}

// EndOfHour end of hour
func (e *evaluator) EndOfHour() time.Time {
	return e.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

// EndOfDay end of day
func (e *evaluator) EndOfDay() time.Time {
	y, m, d := e.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), e.Location())
}

// EndOfWeek end of week
func (e *evaluator) EndOfWeek() time.Time {
	return e.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// EndOfMonth end of month
func (e *evaluator) EndOfMonth() time.Time {
	return e.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfQuarter end of quarter
func (e *evaluator) EndOfQuarter() time.Time {
	return e.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

// EndOfHalf end of half year
func (e *evaluator) EndOfHalf() time.Time {
	return e.BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond)
}

// EndOfYear end of year
func (e *evaluator) EndOfYear() time.Time {
	return e.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// Monday monday
func (e *evaluator) Monday() time.Time {
	t := e.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return t.AddDate(0, 0, -weekday+1)
}

// Sunday sunday
func (e *evaluator) Sunday() time.Time {
	t := e.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		return t
	}
	return t.AddDate(0, 0, (7 - weekday))
}

// TODO!! EndOfSunday end of sunday
func (e *evaluator) EndOfSunday() time.Time {
	//  return New(e.Sunday()).EndOfDay()
	return time.Time{}
}
