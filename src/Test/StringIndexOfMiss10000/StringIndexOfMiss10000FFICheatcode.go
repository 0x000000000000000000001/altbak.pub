// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
package Test_StringIndexOfMiss10000FFICheatcode

import (
	"strings"
)

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := strings.Repeat("a", n) + "needle"
		return func(count int) int {
			return strings.Index(fixture[:count+len("needle")], "absent")
		}
	}
}
