// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_StringConcat1000FFI

func stringConcat1000FFI_foldl[A, B any](f func(B) func(A) B, initial B, values []A) B {
	result := initial
	for _, value := range values {
		result = f(result)(value)
	}
	return result
}

func stringConcat1000FFI_checksum(value string) int {
	return stringConcat1000FFI_foldl(func(total int) func(rune) int {
		return func(cp rune) int { return total + int(cp) }
	}, 0, []rune(value))
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := make([]string, n)
		for i := range fixture {
			fixture[i] = "ab"
		}
		return func(count int) int {
			result := stringConcat1000FFI_foldl(func(left string) func(string) string {
				return func(right string) string { return left + right }
			}, "", fixture[:count])
			return stringConcat1000FFI_checksum(result)
		}
	}
}
