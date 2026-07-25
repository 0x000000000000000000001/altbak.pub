package main

import (
	"os"
	"runtime/pprof"
	"gopurs/output/App"
	"gopurs/output/gopurs_runtime"
)

func main() { defer startPprof()() 
	if os.Getenv("PPROF") == "1" {
		f, err := os.Create("cpu.prof")
		if err != nil { panic(err) }
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	gopurs_runtime.Apply(App.Get_main(), gopurs_runtime.Value{})

	if os.Getenv("PPROF") == "1" {
		mf, err := os.Create("mem.prof")
		if err != nil { panic(err) }
		pprof.WriteHeapProfile(mf)
		mf.Close()
	}
}
