package Bench

import (
	"time"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func BenchNow() float64 {
	return float64(time.Now().UnixNano()) / 1e3
}

func Opaque(a any) func() any {
	return func() any {
		return a
	}
}

func FormatNumber(n float64) string {
	return fmt.Sprintf("%.2f", n)
}

func KeepAlive() any {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	return nil
}
