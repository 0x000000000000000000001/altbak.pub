// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_StringConcat1000FFICheatcode

import (
	"strings"
)

func stringConcat1000FFICheatcode_checksum(value string) int {
	total := 0
	for _, cp := range value {
		total += int(cp)
	}
	return total
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := make([]string, n)
		for i := range fixture {
			fixture[i] = "ab"
		}
		return func(count int) int {
			var builder strings.Builder
			builder.Grow(count * len("ab"))
			for _, part := range fixture[:count] {
				builder.WriteString(part)
			}
			result := builder.String()
			return stringConcat1000FFICheatcode_checksum(result)
		}
	}
}
