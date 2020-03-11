package evaluator

import (
	"testing"
	"time"
)

const testLayout = "2006-01-02 15:04:05"

func formatTime(date time.Time) string {
	return date.Format(testLayout)
}

func testEval(t *testing.T, payload string, initial time.Time) time.Time {
	t.Helper()

	eval := newEvaluator()
	eval.setInitial(initial)
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
	time := testEval(t, "now/s", parseTime("2020-01-01 10:13:23"))
	if !assertTime(t, time, "2020-01-01 10:13:23") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartMinute(t *testing.T) {
	time := testEval(t, "now/m", parseTime("2020-01-01 10:13:23"))
	if !assertTime(t, time, "2020-01-01 10:13:20") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartHour(t *testing.T) {
	time := testEval(t, "now/h", parseTime("2020-01-01 10:13:23"))
	if !assertTime(t, time, "2020-01-01 10:00:00") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartDay(t *testing.T) {
	time := testEval(t, "now/d", parseTime("2020-01-01 10:13:23"))
	if !assertTime(t, time, "2020-01-01 00:00:00") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartWeek(t *testing.T) {
	time := testEval(t, "now/w", parseTime("2020-01-01 10:13:23"))
	if !assertTime(t, time, "2020-01-01 10:00:00") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartMonth(t *testing.T) {
	time := testEval(t, "now/M", parseTime("2020-01-01 10:13:23"))
	if !assertTime(t, time, "2020-01-01 10:00:00") {
		t.FailNow()
	}
}

func TestEvaluator_SnapStartYear(t *testing.T) {
	time := testEval(t, "now/Y", parseTime("2020-03-01 10:13:23"))
	if !assertTime(t, time, "2020-01-01 00:00:00") {
		t.FailNow()
	}
}
