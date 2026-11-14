// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_StringCodeUnitsLengthUnicode1000FFICheatcode

import (
	"strings"
)

func stringCodeUnitsLengthUnicode1000FFICheatcode_codeUnitLength(value string) int {
	count := 0
	for _, cp := range value {
		count++
		if cp > 0xffff {
			count++
		}
	}
	return count
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := strings.Repeat("a😀é", n)
		return func(count int) int {
			return stringCodeUnitsLengthUnicode1000FFICheatcode_codeUnitLength(fixture[:count*len("a😀é")])
		}
	}
}
