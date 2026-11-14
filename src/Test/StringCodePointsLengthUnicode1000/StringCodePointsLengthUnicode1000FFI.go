// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_StringCodePointsLengthUnicode1000FFI

import (
	"strings"
)

func stringCodePointsLengthUnicode1000FFI_toCodePointArray(value string) []int {
	result := make([]int, 0)
	for _, cp := range value {
		result = append(result, int(cp))
	}
	return result
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := strings.Repeat("a😀é", n)
		return func(count int) int {
			return len(stringCodePointsLengthUnicode1000FFI_toCodePointArray(fixture[:count*len("a😀é")]))
		}
	}
}
