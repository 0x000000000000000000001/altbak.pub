// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_StringSplit1000FFI

import (
	"strings"
)

func stringSplit1000FFI_foldl[A, B any](f func(B) func(A) B, initial B, values []A) B {
	result := initial
	for _, value := range values {
		result = f(result)(value)
	}
	return result
}

func stringSplit1000FFI_checksum(value string) int {
	return stringSplit1000FFI_foldl(func(total int) func(rune) int {
		return func(cp rune) int { return total + int(cp) }
	}, 0, []rune(value))
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
			return stringSplit1000FFI_foldl(func(total int) func(string) int {
				return func(part string) int { return total + stringSplit1000FFI_checksum(part) }
			}, len(parts), parts)
		}
	}
}
