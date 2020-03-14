package main

import (
	"github.com/sonirico/datetoken.go/evaluator"
	"time"
)

func evaltoken(payload string) (time.Time, error) {
	evaluatorImpl := evaluator.New()
	return evaluatorImpl.Eval(payload)
}

func Datetoken(payload string) (time.Time, error) {
	return evaltoken(payload)
}
