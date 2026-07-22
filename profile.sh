#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"

./bin/go/run

cat << 'GOEOF' > run/bak/go/output/main.go
package main

import (
	"gopurs/output/App"
	"gopurs/output/gopurs_runtime"
	"os"
	"runtime/pprof"
)

func main() {
	f, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	gopurs_runtime.Apply(App.Get_main(), gopurs_runtime.Value{})
}
GOEOF

cd run/bak/go/output
go build -o go_app .
./go_app
go tool pprof -top cpu.prof | head -n 30
