package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopurs/output/purescript"
)

type entry struct {
	expr  string
	names string
	funcs string
}

func fnv1a(hash uint32, s string) uint32 {
	for i := 0; i < len(s); i++ {
		hash ^= uint32(s[i])
		hash *= 16777619
	}
	return hash
}

func main() {
	corpusPath := os.Args[1]
	iters := 500
	if len(os.Args) > 2 {
		v, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
		iters = v
	}

	raw, err := os.ReadFile(corpusPath)
	if err != nil {
		panic(err)
	}

	var entries []entry
	for _, line := range strings.Split(string(raw), "\n") {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "\t")
		names := ""
		if parts[1] != "-" {
			names = strings.ReplaceAll(parts[1], ",", "\n")
		}
		funcs := "-"
		if parts[2] != "-" {
			funcs = strings.ReplaceAll(parts[2], ",", "\n")
		}
		entries = append(entries, entry{expr: parts[0], names: names, funcs: funcs})
	}

	var hash uint32 = 2166136261
	ops := 0
	totalBytes := 0

	for it := 0; it < iters; it++ {
		for _, w := range entries {
			c := purescript.Call_Sfpurs_ExpressionLanguage_Api_compileWithConfig(w.expr, w.names, w.funcs, 0)
			l := purescript.Call_Sfpurs_ExpressionLanguage_Api_lintWithConfig(w.expr, w.names, w.funcs, 0)
			d := purescript.Call_Sfpurs_ExpressionLanguage_Api_dumpWithConfig(w.expr, w.names, w.funcs, 0)
			hash = fnv1a(hash, c)
			hash = fnv1a(hash, l)
			hash = fnv1a(hash, d)
			totalBytes += len(c) + len(l) + len(d)
			ops += 3
		}
	}

	fmt.Printf("ops=%d bytes=%d checksum=%d\n", ops, totalBytes, hash)
}
