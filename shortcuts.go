package main

import (
	"github.com/sonirico/datetoken.go/evaluator"
	"time"
)

func NaiveDatetoken(payload string) (time.Time, error) {
	eva := evaluator.New()
	return eva.Eval(payload)
}

func Datetoken(payload string, loc *time.Location, weekday time.Weekday) (time.Time, error) {
	eva := evaluator.New()
	eva.SetTZ(loc)
	eva.SetWeeksStartDay(weekday)
	return eva.Eval(payload)
}

func DatetokenIn(payload string, loc *time.Location) (time.Time, error) {
	eva := evaluator.New()
	eva.SetTZ(loc)
	return eva.Eval(payload)
}

func DatetokenWd(payload string, weekday time.Weekday) (time.Time, error) {
	eva := evaluator.New()
	eva.SetWeeksStartDay(weekday)
	return eva.Eval(payload)
}

func DatetokenCfg(payload string, cfg Config) (time.Time, error) {
	eva := evaluator.New()
	eva.SetTZ(cfg.Tz())
	eva.SetWeeksStartDay(cfg.WeeksStartAt())
	return eva.Eval(payload)
}
