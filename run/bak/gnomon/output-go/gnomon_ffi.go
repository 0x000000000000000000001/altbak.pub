package main

import (
	"time"
	"fmt"
)

var Bench_benchNow any = func() any {
	return float64(time.Now().UnixNano()) / 1e3
}

var Bench_formatNumber any = func(n any) any {
	return fmt.Sprintf("%.2f", n.(float64))
}

var Bench_opaque any = func(a any) any {
	return func() any { return a }
}
