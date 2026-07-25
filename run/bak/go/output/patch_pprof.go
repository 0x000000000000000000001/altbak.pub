package main
import (
    "os"
    "log"
    "runtime/pprof"
)
func startPprof() func() {
    f, err := os.Create("cpu.prof")
    if err != nil {
        log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
    return func() {
        pprof.StopCPUProfile()
        f.Close()
    }
}
