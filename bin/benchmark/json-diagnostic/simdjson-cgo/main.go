// Audit driver: simdjson through cgo.
//
// Reports, over the JSON Decoding corpus:
//   parse_us  - one cgo call per document (parse + node count in C);
//   ping_ns   - one cgo crossing of a trivial call (per-call overhead);
//   values    - number of JSON values (nodes) in the corpus, counted in Go,
//               used to price a naive per-value conversion;
//   counts_match - the C-side node counts equal the Go-side counts.
package main

/*
#cgo CXXFLAGS: -std=c++17 -O3 -I/opt/homebrew/opt/simdjson/include
#cgo LDFLAGS: -L/opt/homebrew/opt/simdjson/lib -lsimdjson
#include <stdlib.h>
#include <stdint.h>
uint64_t simdjson_parse_count(const char *data, size_t len);
void simdjson_release(void);
int cgo_ping(int value);
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"
	"unsafe"
)

type corpusCase struct {
	Name      string `json:"name"`
	Benchmark bool   `json:"benchmark"`
	Contents  string `json:"contents"`
}

func nowUS() float64 { return float64(time.Now().UnixNano()) / 1000.0 }

func countNodes(value any) uint64 {
	switch typed := value.(type) {
	case map[string]any:
		total := uint64(1)
		for _, entry := range typed {
			total += 1 + countNodes(entry)
		}
		return total
	case []any:
		total := uint64(1)
		for _, entry := range typed {
			total += countNodes(entry)
		}
		return total
	default:
		return 1
	}
}

func main() {
	raw, err := os.ReadFile(os.Getenv("DIAG_CORPUS"))
	if err != nil {
		panic(err)
	}
	var cases []corpusCase
	if err := json.Unmarshal(raw, &cases); err != nil {
		panic(err)
	}
	timed := []int{}
	for index, entry := range cases {
		if entry.Benchmark {
			timed = append(timed, index)
		}
	}

	// Go-side node counts double as the parse correctness check.
	goCounts := make([]uint64, len(cases))
	var corpusValues uint64
	for index, entry := range cases {
		var dom any
		if err := json.Unmarshal([]byte(entry.Contents), &dom); err != nil {
			panic(err)
		}
		goCounts[index] = countNodes(dom)
		if entry.Benchmark {
			corpusValues += goCounts[index]
		}
	}

	measure := func(runs int, body func() bool) float64 {
		var best float64
		for run := 0; run < runs; run++ {
			begin := nowUS()
			ok := body()
			elapsed := nowUS() - begin
			if !ok {
				panic("parse failed")
			}
			if run == 0 || elapsed < best {
				best = elapsed
			}
		}
		return best
	}

	parseUS := measure(7, func() bool {
		ok := true
		for _, index := range timed {
			text := cases[index].Contents
			cText := C.CString(text)
			count := C.simdjson_parse_count(cText, C.size_t(len(text)))
			C.free(unsafe.Pointer(cText))
			if uint64(count) != goCounts[index] {
				ok = false
			}
		}
		return ok
	})

	const pings = 1000000
	pingBegin := nowUS()
	accumulator := 0
	for index := 0; index < pings; index++ {
		accumulator = int(C.cgo_ping(C.int(accumulator)))
	}
	pingUS := nowUS() - pingBegin
	runtime.KeepAlive(accumulator)
	C.simdjson_release()

	report := map[string]any{
		"backend":      "simdjson-cgo",
		"parse_us":     parseUS,
		"values":       corpusValues,
		"ping_ns":      pingUS * 1000 / pings,
		"pings":        pings,
		"go":           runtime.Version(),
		"gomaxprocs":   runtime.GOMAXPROCS(0),
	}
	encoded, err := json.Marshal(report)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))
}
