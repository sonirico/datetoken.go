package evaluator

import (
	"testing"
	"time"
)

const testLayout = "2006-01-02 15:04:05"

func formatTime(date time.Time) string {
	return date.Format(testLayout)
}

func testEval(t *testing.T, payload string, initial string) time.Time {
	t.Helper()

	eval := New()
	eval.setInitial(parseTime(initial))
	time, err := eval.Eval(payload)
	if err != nil {
		t.Fatalf("eval error %s", err.Error())
	}
	return time
}

func testEvalWithWeekday(t *testing.T, payload string, initial string, weekday time.Weekday) time.Time {
	t.Helper()

	eva := New()
	eva.SetWeeksStartDay(weekday)
	eva.setInitial(parseTime(initial))
	time, err := eva.Eval(payload)
	if err != nil {
		t.Fatalf("eval error %s", err.Error())
	}
	return time
}

func assertTime(t *testing.T, date time.Time, expected string) bool {
	if formatTime(date) != expected {
		t.Errorf("unexpected result. want '%s', have '%s'",
			expected, formatTime(date))
		return false
	}
	return true
}

func parseTime(date string) time.Time {
	res, _ := time.Parse(testLayout, date)
	return res
}

func TestEvaluator_SnapStartSecond(t *testing.T) {
	time := testEval(t, "now/s", "2020-01-01 10:13:23")
	if !assertTime(t, time, "2020-01-01 10:13:23") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartMinute(t *testing.T) {
	time := testEval(t, "now/m", "2020-01-01 10:13:23")
	if !assertTime(t, time, "2020-01-01 10:13:00") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartHour(t *testing.T) {
	time := testEval(t, "now/h", "2020-01-01 10:13:23")
	if !assertTime(t, time, "2020-01-01 10:00:00") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartDay(t *testing.T) {
	time := testEval(t, "now/d", "2020-01-01 10:13:23")
	if !assertTime(t, time, "2020-01-01 00:00:00") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartWeek_DefaultsSunday(t *testing.T) {
	time := testEval(t, "now/w", "2020-03-11 00:00:00")
	if !assertTime(t, time, "2020-03-08 00:00:00") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartWeek(t *testing.T) {
	tests := []struct {
		Payload     string
		Initial     string
		WeekStartAt time.Weekday
		Expected    string
	}{
		{"now/w", "2020-03-11 00:00:00", time.Sunday, "2020-03-08 00:00:00"},
		{"now/w", "2020-03-11 00:00:00", time.Monday, "2020-03-09 00:00:00"},
		{"now/w", "2020-03-11 00:00:00", time.Tuesday, "2020-03-10 00:00:00"},
		{"now/w", "2020-03-11 00:00:00", time.Wednesday, "2020-03-11 00:00:00"},
		{"now/w", "2020-03-11 00:00:00", time.Thursday, "2020-03-05 00:00:00"},
		{"now/w", "2020-03-11 00:00:00", time.Friday, "2020-03-06 00:00:00"},
		{"now/w", "2020-03-11 00:00:00", time.Saturday, "2020-03-07 00:00:00"},
	}
	for _, test := range tests {
		time := testEvalWithWeekday(t, test.Payload, test.Initial, test.WeekStartAt)
		if !assertTime(t, time, test.Expected) {
			t.FailNow()
		}
	}

}

func TestEvaluator_SnapStartBusinessWeek_DefaultsSunday(t *testing.T) {
	time := testEval(t, "now/bw", "2020-03-11 00:00:00")
	if !assertTime(t, time, "2020-03-08 00:00:00") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartBusinessWeek(t *testing.T) {
	tests := []struct {
		Payload     string
		Initial     string
		WeekStartAt time.Weekday
		Expected    string
	}{
		{"now/bw", "2020-03-11 00:00:00", time.Sunday, "2020-03-08 00:00:00"},
		{"now/bw", "2020-03-11 00:00:00", time.Monday, "2020-03-09 00:00:00"},
		{"now/bw", "2020-03-11 00:00:00", time.Tuesday, "2020-03-10 00:00:00"},
		{"now/bw", "2020-03-11 00:00:00", time.Wednesday, "2020-03-11 00:00:00"},
		{"now/bw", "2020-03-11 00:00:00", time.Thursday, "2020-03-05 00:00:00"},
		{"now/bw", "2020-03-11 00:00:00", time.Friday, "2020-03-06 00:00:00"},
		{"now/bw", "2020-03-11 00:00:00", time.Saturday, "2020-03-07 00:00:00"},
	}
	for _, test := range tests {
		time := testEvalWithWeekday(t, test.Payload, test.Initial, test.WeekStartAt)
		if !assertTime(t, time, test.Expected) {
			t.FailNow()
		}
	}

}

func TestEvaluator_SnapStartMonth(t *testing.T) {
	time := testEval(t, "now/M", "2020-02-11 10:13:23")
	if !assertTime(t, time, "2020-02-01 00:00:00") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartYear(t *testing.T) {
	time := testEval(t, "now/Y", "2020-03-01 10:13:23")
	if !assertTime(t, time, "2020-01-01 00:00:00") {
		t.FailNow()
	}
}

// End snap

func TestEvaluator_SnapEndSecond(t *testing.T) {
	time := testEval(t, "now@s", "2020-01-01 10:13:23")
	if !assertTime(t, time, "2020-01-01 10:13:23") {
		t.FailNow()
	}
}

func TestEvaluator_SnapEndMinute(t *testing.T) {
	time := testEval(t, "now@m", "2020-01-01 10:13:23")
	if !assertTime(t, time, "2020-01-01 10:13:59") {
		t.FailNow()
	}
}

func TestEvaluator_SnapEndHour(t *testing.T) {
	time := testEval(t, "now@h", "2020-01-01 10:13:23")
	if !assertTime(t, time, "2020-01-01 10:59:59") {
		t.FailNow()
	}
}

func TestEvaluator_SnapEndDay(t *testing.T) {
	time := testEval(t, "now@d", "2020-01-01 10:13:23")
	if !assertTime(t, time, "2020-01-01 23:59:59") {
		t.FailNow()
	}
}

func TestEvaluator_SnapEndWeek_DefaultsSunday(t *testing.T) {
	time := testEval(t, "now@w", "2020-03-11 00:00:00")
	if !assertTime(t, time, "2020-03-14 23:59:59") {
		t.FailNow()
	}
}

func TestEvaluator_SnapEndWeek(t *testing.T) {
	tests := []struct {
		Payload     string
		Initial     string
		WeekStartAt time.Weekday
		Expected    string
	}{
		{"now@w", "2020-03-11 00:00:00", time.Sunday, "2020-03-14 23:59:59"},
		{"now@w", "2020-03-11 00:00:00", time.Monday, "2020-03-15 23:59:59"},
		{"now@w", "2020-03-11 00:00:00", time.Tuesday, "2020-03-16 23:59:59"},
		{"now@w", "2020-03-11 00:00:00", time.Wednesday, "2020-03-17 23:59:59"},
		{"now@w", "2020-03-11 00:00:00", time.Thursday, "2020-03-11 23:59:59"},
		{"now@w", "2020-03-11 00:00:00", time.Friday, "2020-03-12 23:59:59"},
		{"now@w", "2020-03-11 00:00:00", time.Saturday, "2020-03-13 23:59:59"},
	}
	for _, test := range tests {
		time := testEvalWithWeekday(t, test.Payload, test.Initial, test.WeekStartAt)
		if !assertTime(t, time, test.Expected) {
			t.FailNow()
		}
	}
}

func TestEvaluator_SnapEndBusinessWeek_DefaultsSunday(t *testing.T) {
	time := testEval(t, "now@bw", "2020-03-11 00:00:00")
	if !assertTime(t, time, "2020-03-12 23:59:59") {
		t.FailNow()
	}
}

func TestEvaluator_SnapEndBusinessWeek(t *testing.T) {
	tests := []struct {
		Payload     string
		Initial     string
		WeekStartAt time.Weekday
		Expected    string
	}{
		{"now@bw", "2020-03-11 00:00:00", time.Sunday, "2020-03-12 23:59:59"},
		{"now@bw", "2020-03-11 00:00:00", time.Monday, "2020-03-13 23:59:59"},
		{"now@bw", "2020-03-11 00:00:00", time.Tuesday, "2020-03-14 23:59:59"},
		{"now@bw", "2020-03-11 00:00:00", time.Wednesday, "2020-03-15 23:59:59"},
		{"now@bw", "2020-03-11 00:00:00", time.Thursday, "2020-03-09 23:59:59"},
		{"now@bw", "2020-03-11 00:00:00", time.Friday, "2020-03-10 23:59:59"},
		{"now@bw", "2020-03-11 00:00:00", time.Saturday, "2020-03-11 23:59:59"},
	}
	for _, test := range tests {
		time := testEvalWithWeekday(t, test.Payload, test.Initial, test.WeekStartAt)
		if !assertTime(t, time, test.Expected) {
			t.FailNow()
		}
	}
}

func TestEvaluator_SnapEndMonth(t *testing.T) {
	time := testEval(t, "now@M", "2020-02-11 10:13:23")
	if !assertTime(t, time, "2020-02-29 23:59:59") {
		t.FailNow()
	}
}

func TestEvaluator_SnapEndYear(t *testing.T) {
	time := testEval(t, "now@Y", "2020-02-11 10:13:23")
	if !assertTime(t, time, "2020-12-31 23:59:59") {
		t.FailNow()
	}
}

// Arithmetic

func TestEvaluator_AddSeconds(t *testing.T) {
	tests := []struct {
		Token    string
		Initial  string
		Expected string
	}{
		// Plus
		{"now+s", "2020-02-11 10:13:23", "2020-02-11 10:13:24"},
		{"now+40s", "2020-02-11 10:13:23", "2020-02-11 10:14:03"},
		{"now+1s", "2020-02-11 10:59:59", "2020-02-11 11:00:00"},
		{"now+10s", "2020-02-11 23:59:50", "2020-02-12 00:00:00"},
		{"now+10s", "2020-01-31 23:59:50", "2020-02-01 00:00:00"},
		{"now+10s", "2020-12-31 23:59:59", "2021-01-01 00:00:09"},
		// Minus
		{"now-s", "2020-02-11 10:13:23", "2020-02-11 10:13:22"},
		{"now-40s", "2020-02-11 10:13:23", "2020-02-11 10:12:43"},
		{"now-1s", "2020-02-11 10:00:00", "2020-02-11 09:59:59"},
		{"now-10s", "2020-03-01 00:00:09", "2020-02-29 23:59:59"},
		{"now-10s", "2020-02-01 00:00:00", "2020-01-31 23:59:50"},
		{"now-10s", "2021-01-01 00:00:09", "2020-12-31 23:59:59"},
	}

	for _, test := range tests {
		actual := testEval(t, test.Token, test.Initial)
		if !assertTime(t, actual, test.Expected) {
			t.FailNow()
		}
	}
}

func TestEvaluator_AddMinutes(t *testing.T) {
	tests := []struct {
		Token    string
		Initial  string
		Expected string
	}{
		// Plus
		{"now+m", "2020-02-11 10:13:23", "2020-02-11 10:14:23"},
		{"now+40m", "2020-02-11 10:13:23", "2020-02-11 10:53:23"},
		{"now+1m", "2020-02-11 10:59:59", "2020-02-11 11:00:59"},
		{"now+10m", "2020-02-11 23:59:50", "2020-02-12 00:09:50"},
		{"now+10m", "2020-01-31 23:59:50", "2020-02-01 00:09:50"},
		{"now+10m", "2020-12-31 23:59:59", "2021-01-01 00:09:59"},
		// Minus
		{"now-m", "2020-02-11 10:14:23", "2020-02-11 10:13:23"},
		{"now-40m", "2020-02-11 10:53:23", "2020-02-11 10:13:23"},
		{"now-1m", "2020-02-11 11:00:59", "2020-02-11 10:59:59"},
		{"now-10m", "2020-02-12 00:09:50", "2020-02-11 23:59:50"},
		{"now-10m", "2020-02-01 00:09:50", "2020-01-31 23:59:50"},
		{"now-10m", "2021-01-01 00:09:59", "2020-12-31 23:59:59"},
	}

	for _, test := range tests {
		actual := testEval(t, test.Token, test.Initial)
		if !assertTime(t, actual, test.Expected) {
			t.FailNow()
		}
	}
}

func TestEvaluator_AddHours(t *testing.T) {
	tests := []struct {
		Token    string
		Initial  string
		Expected string
	}{
		// Plus
		{"now+h", "2020-02-11 10:13:23", "2020-02-11 11:13:23"},
		{"now+40h", "2020-02-11 10:13:23", "2020-02-13 02:13:23"},
		{"now+1h", "2020-02-11 10:59:59", "2020-02-11 11:59:59"},
		{"now+10h", "2020-02-11 23:59:50", "2020-02-12 09:59:50"},
		{"now+10h", "2020-01-31 23:59:50", "2020-02-01 09:59:50"},
		{"now+10h", "2020-12-31 23:59:59", "2021-01-01 09:59:59"},
		// Minus
		{"now-h", "2020-02-11 11:13:23", "2020-02-11 10:13:23"},
		{"now-40h", "2020-02-13 02:13:23", "2020-02-11 10:13:23"},
		{"now-1h", "2020-02-11 11:59:59", "2020-02-11 10:59:59"},
		{"now-10h", "2020-02-12 09:59:50", "2020-02-11 23:59:50"},
		{"now-10h", "2020-02-01 09:59:50", "2020-01-31 23:59:50"},
		{"now-10h", "2021-01-01 09:59:59", "2020-12-31 23:59:59"},
	}

	for _, test := range tests {
		actual := testEval(t, test.Token, test.Initial)
		if !assertTime(t, actual, test.Expected) {
			t.FailNow()
		}
	}
}

func TestEvaluator_AddDays(t *testing.T) {
	tests := []struct {
		Token    string
		Initial  string
		Expected string
	}{
		// Plus
		{"now+d", "2020-02-11 00:00:00", "2020-02-12 00:00:00"},
		{"now+50d", "2020-02-11 00:00:00", "2020-04-01 00:00:00"},
		{"now+1d", "2020-02-28 00:00:00", "2020-02-29 00:00:00"},
		{"now+1d", "2020-02-29 00:00:00", "2020-03-01 00:00:00"},
		{"now+10d", "2020-12-31 00:00:00", "2021-01-10 00:00:00"},
		// 30 dias tiene septiembre, junto con abril, junio y noviembre.
		{"now+1d", "2020-04-30 00:00:00", "2020-05-01 00:00:00"},
		{"now+1d", "2020-06-30 00:00:00", "2020-07-01 00:00:00"},
		{"now+1d", "2020-09-30 00:00:00", "2020-10-01 00:00:00"},
		{"now+1d", "2020-11-30 00:00:00", "2020-12-01 00:00:00"},
		// Minus
		{"now-d", "2020-02-12 00:00:00", "2020-02-11 00:00:00"},
		{"now-50d", "2020-04-01 00:00:00", "2020-02-11 00:00:00"},
		{"now-1d", "2020-02-29 00:00:00", "2020-02-28 00:00:00"},
		{"now-1d", "2020-03-01 00:00:00", "2020-02-29 00:00:00"},
		{"now-10d", "2021-01-10 00:00:00", "2020-12-31 00:00:00"},
		// 30 dias tiene septiembre, junto con abril, junio y noviembre.
		{"now-1d", "2020-05-01 00:00:00", "2020-04-30 00:00:00"},
		{"now-1d", "2020-07-01 00:00:00", "2020-06-30 00:00:00"},
		{"now-1d", "2020-10-01 00:00:00", "2020-09-30 00:00:00"},
		{"now-1d", "2020-12-01 00:00:00", "2020-11-30 00:00:00"},
	}

	for _, test := range tests {
		actual := testEval(t, test.Token, test.Initial)
		if !assertTime(t, actual, test.Expected) {
			t.FailNow()
		}
	}
}

func TestEvaluator_AddWeeks(t *testing.T) {
	tests := []struct {
		Token    string
		Initial  string
		Expected string
	}{
		// Plus
		{"now+w", "2020-04-30 00:00:00", "2020-05-07 00:00:00"},
		{"now+2w", "2020-06-30 00:00:00", "2020-07-14 00:00:00"},
		{"now+3w", "2020-09-30 00:00:00", "2020-10-21 00:00:00"},
		{"now+5w", "2020-11-30 00:00:00", "2021-01-04 00:00:00"},
		// Minus
		{"now-w", "2020-05-07 00:00:00", "2020-04-30 00:00:00"},
		{"now-2w", "2020-07-14 00:00:00", "2020-06-30 00:00:00"},
		{"now-3w", "2020-10-21 00:00:00", "2020-09-30 00:00:00"},
		{"now-5w", "2021-01-04 00:00:00", "2020-11-30 00:00:00"},
	}

	for _, test := range tests {
		actual := testEval(t, test.Token, test.Initial)
		if !assertTime(t, actual, test.Expected) {
			t.FailNow()
		}
	}
}

func TestEvaluator_AddMonths(t *testing.T) {
	tests := []struct {
		Token    string
		Initial  string
		Expected string
	}{
		// Plus
		{"now+M", "2020-02-11 10:13:23", "2020-03-11 10:13:23"},
		{"now+3M", "2020-02-11 10:13:23", "2020-05-11 10:13:23"},
		{"now+5M", "2020-08-11 10:59:59", "2021-01-11 10:59:59"},
		{"now+20M", "2020-09-11 23:59:50", "2022-05-11 23:59:50"},
		// Minus
		{"now-M", "2020-03-11 10:13:23", "2020-02-11 10:13:23"},
		{"now-3M", "2020-05-11 10:13:23", "2020-02-11 10:13:23"},
		{"now-5M", "2021-01-11 10:59:59", "2020-08-11 10:59:59"},
		{"now-20M", "2022-05-11 23:59:50", "2020-09-11 23:59:50"},
	}

	for _, test := range tests {
		actual := testEval(t, test.Token, test.Initial)
		if !assertTime(t, actual, test.Expected) {
			t.FailNow()
		}
	}
}

func TestEvaluator_AddYears(t *testing.T) {
	tests := []struct {
		Token    string
		Initial  string
		Expected string
	}{
		// Plus
		{"now+Y", "2020-02-11 10:13:23", "2021-02-11 10:13:23"},
		{"now+1Y", "2020-02-11 10:13:23", "2021-02-11 10:13:23"},
		{"now+2Y", "2020-08-11 10:59:59", "2022-08-11 10:59:59"},
		{"now+10Y", "2020-09-11 23:59:50", "2030-09-11 23:59:50"},
		// Minus
		{"now-Y", "2020-02-11 10:13:23", "2019-02-11 10:13:23"},
		{"now-1Y", "2020-02-11 10:13:23", "2019-02-11 10:13:23"},
		{"now-2Y", "2020-08-11 10:59:59", "2018-08-11 10:59:59"},
		{"now-10Y", "2020-09-11 23:59:50", "2010-09-11 23:59:50"},
	}

	for _, test := range tests {
		actual := testEval(t, test.Token, test.Initial)
		if !assertTime(t, actual, test.Expected) {
			t.FailNow()
		}
	}
}

// Week days
func TestEvaluator_PreviousWeekday(t *testing.T) {
	tests := []struct {
		Token    string
		Initial  string
		Expected string
	}{
		{"now/mon", "2020-03-11 00:00:00", "2020-03-09 00:00:00"},
		{"now/tue", "2020-03-11 00:00:00", "2020-03-10 00:00:00"},
		{"now/wed", "2020-03-11 00:00:00", "2020-03-11 00:00:00"},
		{"now/thu", "2020-03-11 00:00:00", "2020-03-05 00:00:00"},
		{"now/fri", "2020-03-11 00:00:00", "2020-03-06 00:00:00"},
		{"now/sat", "2020-03-11 00:00:00", "2020-03-07 00:00:00"},
		{"now/sun", "2020-03-11 00:00:00", "2020-03-08 00:00:00"},
	}

	for _, test := range tests {
		actual := testEval(t, test.Token, test.Initial)
		if !assertTime(t, actual, test.Expected) {
			t.FailNow()
		}
	}
}

func TestEvaluator_NextWeekday(t *testing.T) {
	tests := []struct {
		Token    string
		Initial  string
		Expected string
	}{
		{"now@mon", "2020-03-11 00:00:00", "2020-03-16 00:00:00"},
		{"now@tue", "2020-03-11 00:00:00", "2020-03-17 00:00:00"},
		{"now@wed", "2020-03-11 00:00:00", "2020-03-11 00:00:00"},
		{"now@thu", "2020-03-11 00:00:00", "2020-03-12 00:00:00"},
		{"now@fri", "2020-03-11 00:00:00", "2020-03-13 00:00:00"},
		{"now@sat", "2020-03-11 00:00:00", "2020-03-14 00:00:00"},
		{"now@sun", "2020-03-11 00:00:00", "2020-03-15 00:00:00"},
	}

	for _, test := range tests {
		actual := testEval(t, test.Token, test.Initial)
		if !assertTime(t, actual, test.Expected) {
			t.FailNow()
		}
	}
}

func TestEvaluator_InvalidToken(t *testing.T) {
	tests := []string{
		"",
		"now/pepe",
		"now@paco",
		"now-abc",
		"now+xyz",
		"now-second",
		"now-w/bw+2h@null",
	}
	for _, payload := range tests {
		eva := New()
		if _, err := eva.Eval(payload); err == nil {
			t.Errorf("unexpected valid token. want error, have token '%s'",
				payload)
		}
	}
}
