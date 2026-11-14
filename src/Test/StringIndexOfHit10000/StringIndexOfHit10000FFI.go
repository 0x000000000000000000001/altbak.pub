// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_StringIndexOfHit10000FFI

import (
	"strings"
)

type stringIndexOfHit10000FFI_maybeIndex struct {
	valid bool
	value int
}

func stringIndexOfHit10000FFI_indexOf(pattern, input string) stringIndexOfHit10000FFI_maybeIndex {
	index := strings.Index(input, pattern)
	if index < 0 {
		return stringIndexOfHit10000FFI_maybeIndex{}
	}
	// These fixtures contain ASCII only, so byte and UTF-16 indices coincide.
	return stringIndexOfHit10000FFI_maybeIndex{true, index}
}

func stringIndexOfHit10000FFI_fromMaybe(fallback int, result stringIndexOfHit10000FFI_maybeIndex) int {
	if result.valid {
		return result.value
	}
	return fallback
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := strings.Repeat("a", n) + "needle"
		return func(count int) int {
			return stringIndexOfHit10000FFI_fromMaybe(-1, stringIndexOfHit10000FFI_indexOf("needle", fixture[:count+len("needle")]))
		}
	}
}
