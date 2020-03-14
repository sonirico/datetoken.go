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

	eval := newEvaluator()
	eval.setInitial(parseTime(initial))
	time, err := eval.Eval(payload)
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

func TestEvaluator_SnapStartWeek(t *testing.T) {
	time := testEval(t, "now/w", "2020-01-01 10:13:23")
	if !assertTime(t, time, "2020-01-01 10:00:00") {
		t.FailNow()
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

func TestEvaluator_SnapEndWeek(t *testing.T) {
	time := testEval(t, "now@w", "2020-01-01 10:13:23")
	if !assertTime(t, time, "2020-01-01 10:00:00") {
		t.FailNow()
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
		{"now+s", "2020-02-11 10:13:23", "2020-02-11 10:13:24"},
		{"now+40s", "2020-02-11 10:13:23", "2020-02-11 10:14:03"},
		{"now+1s", "2020-02-11 10:59:59", "2020-02-11 11:00:00"},
		{"now+10s", "2020-02-11 23:59:50", "2020-02-12 00:00:00"},
		{"now+10s", "2020-01-31 23:59:50", "2020-02-01 00:00:00"},
		{"now+10s", "2020-12-31 23:59:59", "2021-01-01 00:00:09"},
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
		{"now+m", "2020-02-11 10:13:23", "2020-02-11 10:14:23"},
		{"now+40m", "2020-02-11 10:13:23", "2020-02-11 10:53:23"},
		{"now+1m", "2020-02-11 10:59:59", "2020-02-11 11:00:59"},
		{"now+10m", "2020-02-11 23:59:50", "2020-02-12 00:09:50"},
		{"now+10m", "2020-01-31 23:59:50", "2020-02-01 00:09:50"},
		{"now+10m", "2020-12-31 23:59:59", "2021-01-01 00:09:59"},
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
		{"now+h", "2020-02-11 10:13:23", "2020-02-11 11:13:23"},
		{"now+40h", "2020-02-11 10:13:23", "2020-02-13 02:13:23"},
		{"now+1h", "2020-02-11 10:59:59", "2020-02-11 11:59:59"},
		{"now+10h", "2020-02-11 23:59:50", "2020-02-12 09:59:50"},
		{"now+10h", "2020-01-31 23:59:50", "2020-02-01 09:59:50"},
		{"now+10h", "2020-12-31 23:59:59", "2021-01-01 09:59:59"},
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
		// 30 dias tiene septiembre, junto con abril, junio y noviembre.
		{"now+w", "2020-04-30 00:00:00", "2020-05-07 00:00:00"},
		{"now+2w", "2020-06-30 00:00:00", "2020-07-14 00:00:00"},
		{"now+3w", "2020-09-30 00:00:00", "2020-10-21 00:00:00"},
		{"now+5w", "2020-11-30 00:00:00", "2021-01-04 00:00:00"},
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
		{"now+M", "2020-02-11 10:13:23", "2020-03-11 10:13:23"},
		{"now+3M", "2020-02-11 10:13:23", "2020-05-11 10:13:23"},
		{"now+5M", "2020-08-11 10:59:59", "2021-01-11 10:59:59"},
		{"now+20M", "2020-09-11 23:59:50", "2022-05-11 23:59:50"},
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
		{"now+Y", "2020-02-11 10:13:23", "2021-02-11 10:13:23"},
		{"now+1Y", "2020-02-11 10:13:23", "2021-02-11 10:13:23"},
		{"now+2Y", "2020-08-11 10:59:59", "2022-08-11 10:59:59"},
		{"now+10Y", "2020-09-11 23:59:50", "2030-09-11 23:59:50"},
	}

	for _, test := range tests {
		actual := testEval(t, test.Token, test.Initial)
		if !assertTime(t, actual, test.Expected) {
			t.FailNow()
		}
	}
}
