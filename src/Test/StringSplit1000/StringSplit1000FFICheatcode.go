// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_StringSplit1000FFICheatcode

import (
	"strings"
)

func stringSplit1000FFICheatcode_checksum(value string) int {
	total := 0
	for _, cp := range value {
		total += int(cp)
	}
	return total
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		parts := make([]string, n)
		for i := range parts {
			parts[i] = "abc"
		}
		fixture := strings.Join(parts, ",")
		return func(count int) int {
			size := count * 4
			if count > 0 {
				size--
			}
			parts := strings.Split(fixture[:size], ",")
			total := len(parts)
			for _, part := range parts {
				total += stringSplit1000FFICheatcode_checksum(part)
			}
			return total
		}
	}
}
