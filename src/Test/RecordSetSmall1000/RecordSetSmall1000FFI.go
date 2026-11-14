// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_RecordSetSmall1000FFI

type recordSetSmall1000FFI_record struct{ a, b int }

func recordSetSmall1000FFI_sumRecord(value recordSetSmall1000FFI_record) int {
	return value.a + value.b
}

func recordSetSmall1000FFI_iterate(f func(recordSetSmall1000FFI_record) recordSetSmall1000FFI_record, count int, initial recordSetSmall1000FFI_record) recordSetSmall1000FFI_record {
	result := initial
	for i := 0; i < count; i++ {
		result = f(result)
	}
	return result
}

func recordSetSmall1000FFI_incrementLast(value recordSetSmall1000FFI_record) recordSetSmall1000FFI_record {
	result := value
	result.b++
	return result
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := recordSetSmall1000FFI_record{a: 1, b: 2}
		return func(count int) int {
			result := recordSetSmall1000FFI_iterate(recordSetSmall1000FFI_incrementLast, count, fixture)
			return recordSetSmall1000FFI_sumRecord(result) + recordSetSmall1000FFI_sumRecord(fixture)
		}
	}
}
