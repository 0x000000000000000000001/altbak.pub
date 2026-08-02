package main

import (
	"os"
	"runtime/pprof"
	"gopurs/output/AppX"
	"gopurs/output/gopurs_runtime"
)

func main() {
	if os.Getenv("PPROF") == "1" {
		f, err := os.Create("cpu.prof")
		if err != nil { panic(err) }
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	gopurs_runtime.Apply(AppX.Get_main(), gopurs_runtime.Value{})

	gopurs_runtime.EventLoopWait()

	if os.Getenv("PPROF") == "1" {
		mf, err := os.Create("mem.prof")
		if err != nil { panic(err) }
		pprof.WriteHeapProfile(mf)
		mf.Close()
	}
}
