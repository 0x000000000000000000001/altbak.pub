// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_StringTakeCodePointsUnicode1500FFICheatcode

import (
	"strings"
)

func stringTakeCodePointsUnicode1500FFICheatcode_checksum(value string) int {
	total := 0
	for _, cp := range value {
		total += int(cp)
	}
	return total
}

func stringTakeCodePointsUnicode1500FFICheatcode_takeCodePoints(count int, value string) string {
	end, seen := len(value), 0
	for index := range value {
		if seen == count {
			end = index
			break
		}
		seen++
	}
	return strings.Clone(value[:end])
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := strings.Repeat("a😀é", n)
		return func(count int) int {
			result := stringTakeCodePointsUnicode1500FFICheatcode_takeCodePoints(count+count/2, fixture[:count*len("a😀é")])
			return stringTakeCodePointsUnicode1500FFICheatcode_checksum(result)
		}
	}
}
