// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_StringTakeCodePointsUnicode1500FFI

import (
	"strings"
)

func stringTakeCodePointsUnicode1500FFI_foldl[A, B any](f func(B) func(A) B, initial B, values []A) B {
	result := initial
	for _, value := range values {
		result = f(result)(value)
	}
	return result
}

func stringTakeCodePointsUnicode1500FFI_checksum(value string) int {
	return stringTakeCodePointsUnicode1500FFI_foldl(func(total int) func(rune) int {
		return func(cp rune) int { return total + int(cp) }
	}, 0, []rune(value))
}

func stringTakeCodePointsUnicode1500FFI_takeCodePoints(count int, value string) string {
	result, seen := "", 0
	for _, cp := range value {
		if seen == count {
			break
		}
		result += string(cp)
		seen++
	}
	return result
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := strings.Repeat("a😀é", n)
		return func(count int) int {
			result := stringTakeCodePointsUnicode1500FFI_takeCodePoints(count+count/2, fixture[:count*len("a😀é")])
			return stringTakeCodePointsUnicode1500FFI_checksum(result)
		}
	}
}
