package main

import (
	"fmt"
	"time"
)

type Lazy func() int

func buildThunks(n int, acc Lazy) Lazy {
	for n > 0 {
		oldAcc := acc
		acc = func() int {
			return oldAcc() + 1
		}
		n--
	}
	return acc
}

func runManyTimes(n int, acc int) int {
	for n > 0 {
		acc += buildThunks(1000, func() int { return 0 })()
		n--
	}
	return acc
}

func main() {
	start := time.Now()
	res := runManyTimes(1000, 0)
	dt := time.Since(start).Microseconds()
	fmt.Printf("Lazy Evaluation: %d, Execution time: %d μs\n", res, dt)
}
