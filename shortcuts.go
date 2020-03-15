package datetoken

import (
	"github.com/sonirico/datetoken.go/evaluator"
	"time"
)

// EvalNaive evaluates a token with default settings (TZ=UTC, weeks start on Sunday)
func EvalNaive(payload string) (time.Time, error) {
	eva := evaluator.New()
	return eva.Eval(payload)
}

// Eval evaluates a token for a give time zone and weekday, being the start of weeks configurable
func Eval(payload string, loc *time.Location, weekday time.Weekday) (time.Time, error) {
	eva := evaluator.New()
	eva.SetTZ(loc)
	eva.SetWeeksStartDay(weekday)
	return eva.Eval(payload)
}

// EvalIn evaluates a token for a given time zone, being Sunday the start of weeks
func EvalIn(payload string, loc *time.Location) (time.Time, error) {
	eva := evaluator.New()
	eva.SetTZ(loc)
	return eva.Eval(payload)
}

// EvalWeekDay evaluates a token in UTC, being the start of a week configurable
func EvalWeekDay(payload string, weekday time.Weekday) (time.Time, error) {
	eva := evaluator.New()
	eva.SetWeeksStartDay(weekday)
	return eva.Eval(payload)
}

// EvalCfg evaluates a token, allowing configuration by using any struct compliant with Config
func EvalCfg(payload string, cfg Config) (time.Time, error) {
	eva := evaluator.New()
	eva.SetTZ(cfg.Tz())
	eva.SetWeeksStartDay(cfg.WeeksStartAt())
	return eva.Eval(payload)
}
