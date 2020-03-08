package evaluator

import (
	"fmt"
	"testing"
	"time"
)

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

func parseTime(date string) time.Time {
	res, _ := time.Parse("2006-01-02 15:04:05", date)
	fmt.Println(fmt.Sprintf("time %v(%T)", res, res))
	return res
}

func TestEvaluator_SnapStartSecond(t *testing.T) {
	time := testEval(t, "now/s", parseTime("2020-01-01 10:13:23"))
	if time.Nanosecond() != 0 {
		t.Errorf("expected nanosecond to be '0', got '%d'", time.Second())
	}
}

func TestEvaluator_SnapStartMinute(t *testing.T) {
	time := testEval(t, "now/m", parseTime("2020-01-01 10:13:23"))
	if time.String() != "2020-01-01 10:13:23" {
		t.Errorf("mira mi t %s", time.String())
		t.FailNow()
	}
	if time.Nanosecond() != 0 {
		t.Errorf("expected nanosecond to be '0', got '%d'", time.Second())
	}
	if time.Second() != 0 {
		t.Errorf("expected second to be '0', got '%d'", time.Second())
	}
}

func TestEvaluator_SnapStartHour(t *testing.T) {
	time := testEval(t, "now/h", parseTime("2020-01-01 10:13:23"))
	if time.Nanosecond() != 0 {
		t.Errorf("expected nanosecond to be '0', got '%d'", time.Second())
	}
	if time.Second() != 0 {
		t.Errorf("expected second to be '0', got '%d'", time.Second())
	}
	if time.Second() != 0 {
		t.Errorf("expected minute to be '0', got '%d'", time.Second())
	}
}