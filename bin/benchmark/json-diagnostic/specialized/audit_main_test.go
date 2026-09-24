package purescript

import (
	"fmt"
	"strings"
	"testing"

	"gopurs/output/gopurs_runtime"
)

// A setup-only check or a last-result sink misses the bad first timed output.
// The validator must inspect both retained outputs without rerunning decode.
func TestAuditValidatesEveryTimedOutput(t *testing.T) {
	for _, phase := range []string{"parse", "decode", "combined"} {
		t.Run(phase, func(t *testing.T) {
			calls := 0
			identity := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value { return value })
			badFirst := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
				calls++
				if calls == 1 {
					return gopurs_runtime.Str(`"bad"`)
				}
				return value
			})
			parse, decode := identity, badFirst
			if phase == "parse" {
				parse, decode = badFirst, identity
			}
			files := []zzCorpusFile{{Name: "first", Contents: `"good"`}, {Name: "last", Contents: `"good"`}}
			parsed := []gopurs_runtime.Value{gopurs_runtime.Str(`"good"`), gopurs_runtime.Str(`"good"`)}
			hashes := []string{canonicalHash(`"good"`), canonicalHash(`"good"`)}
			defer func() {
				failure := recover()
				if failure == nil || !strings.Contains(fmt.Sprint(failure), "output for first") {
					t.Fatalf("first timed result was not validated: %v", failure)
				}
				if calls != 2 {
					t.Fatalf("validation reran the measured operation: %d calls", calls)
				}
			}()
			zzRunSegment(phase, decode, parse, identity, identity, files, parsed, []int{0, 1},
				zzOracle{Fingerprints: hashes, JSONFingerprints: hashes})
		})
	}
}
