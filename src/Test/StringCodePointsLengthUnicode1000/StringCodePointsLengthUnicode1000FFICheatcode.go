// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
package Test_StringCodePointsLengthUnicode1000FFICheatcode

import (
	"strings"
	"unicode/utf8"
)

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := strings.Repeat("a😀é", n)
		return func(count int) int {
			return utf8.RuneCountInString(fixture[:count*len("a😀é")])
		}
	}
}
