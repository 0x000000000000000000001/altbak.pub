package main

import (
    "testing"
    "time"
    . "github.com/purescript-native/go-runtime"
    _ "project.localhost/purescript-native/ffi-loader"
)

func TestBenchmarkClockAndOpaque(t *testing.T) {
    exports := Foreign("Bench")
    now := exports["benchNow"].(func() Any)
    previous := now().(float64)
    for i := 0; i < 100; i++ {
        current := now().(float64)
        if current < previous { t.Fatal("clock moved backwards") }
        previous = current
    }
    begin := now().(float64)
    time.Sleep(25 * time.Millisecond)
    elapsed := now().(float64) - begin
    if elapsed < 20000 || elapsed > 500000 {
        t.Fatalf("expected microseconds around a 25ms sleep, received %f", elapsed)
    }
    opaque := exports["opaque"].(func(Any) Any)
    for _, value := range []int{-17, 0, 93} {
        got := opaque(value).(func() Any)()
        if got != value { t.Fatalf("opaque(%d) = %v", value, got) }
    }
}
