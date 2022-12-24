package decoretor

import (
	"fmt"
	"time"
)

type MultiplyFunc func(int, int) int

func Multiply(a, b int) int {
	return a * b
}

func execTime(f MultiplyFunc) MultiplyFunc {
	return func(a, b int) int {
		start := time.Now()
		c := f(a, b)
		end := time.Since(start)
		fmt.Printf("Execution time: %d\n", end)
		return c
	}
}
