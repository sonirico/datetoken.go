package main

import (
	"fmt"

	"github.com/sonirico/datetoken.go"
)

func main() {
	tokens := []string{
		"now/s",
		"now/m",
		"now/h",
		"now/d",
		"now/w",
		"now/bw",
		"now/M",
		"now/Y",
		"now/Q",
	}
	fmt.Println("Snap to start of units")
	for _, token := range tokens {
		date, _ := datetoken.EvalNaive(token)
		fmt.Printf("\n{token: %s, time: %s}", token, date)
	}
	fmt.Printf("\n")
	fmt.Println("Snap to end of units")
	tokens = []string{
		"now@s",
		"now@m",
		"now@h",
		"now@d",
		"now@w",
		"now@bw",
		"now@M",
		"now@Y",
		"now@Q",
	}
	for _, token := range tokens {
		date, _ := datetoken.EvalNaive(token)
		fmt.Printf("\n{token: %s, time: %s}", token, date)
	}
	fmt.Printf("\n")
}
