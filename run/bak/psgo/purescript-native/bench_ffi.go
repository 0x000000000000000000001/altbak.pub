package ffi_loader

import (
	"time"
	"fmt"
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Bench")
	exports["benchNow"] = func() Any {
		return float64(time.Now().UnixNano()) / 1e3
	}
	exports["formatNumber"] = func(n Any) Any {
		return fmt.Sprintf("%.2f", n.(float64))
	}
	exports["opaque"] = func(a Any) Any {
		return func() Any { return a }
	}
}
