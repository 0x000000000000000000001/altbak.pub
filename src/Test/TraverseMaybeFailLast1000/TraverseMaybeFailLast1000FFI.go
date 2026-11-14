// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_TraverseMaybeFailLast1000FFI

func traverseMaybeFailLast1000FFI_foldl[A, B any](f func(B) func(A) B, initial B, values []A) B {
	result := initial
	for _, value := range values {
		result = f(result)(value)
	}
	return result
}

func traverseMaybeFailLast1000FFI_weightedChecksum(values []int) int {
	type accumulator struct{ index, total int }
	result := traverseMaybeFailLast1000FFI_foldl(func(acc accumulator) func(int) accumulator {
		return func(value int) accumulator {
			return accumulator{acc.index + 1, acc.total + (1+acc.index%2)*value}
		}
	}, accumulator{}, values)
	return result.total
}

type traverseMaybeFailLast1000FFI_maybe[A any] struct {
	valid bool
	value A
}

func traverseMaybeFailLast1000FFI_mapMaybe[A, B any](f func(A) B, value traverseMaybeFailLast1000FFI_maybe[A]) traverseMaybeFailLast1000FFI_maybe[B] {
	if !value.valid {
		return traverseMaybeFailLast1000FFI_maybe[B]{}
	}
	return traverseMaybeFailLast1000FFI_maybe[B]{true, f(value.value)}
}

func traverseMaybeFailLast1000FFI_applyMaybe[A, B any](f traverseMaybeFailLast1000FFI_maybe[func(A) B], value traverseMaybeFailLast1000FFI_maybe[A]) traverseMaybeFailLast1000FFI_maybe[B] {
	if !f.valid || !value.valid {
		return traverseMaybeFailLast1000FFI_maybe[B]{}
	}
	return traverseMaybeFailLast1000FFI_maybe[B]{true, f.value(value.value)}
}

func traverseMaybeFailLast1000FFI_array1(a int) []int { return []int{a} }
func traverseMaybeFailLast1000FFI_array2(a int) func(int) []int {
	return func(b int) []int { return []int{a, b} }
}
func traverseMaybeFailLast1000FFI_array3(a int) func(int) func(int) []int {
	return func(b int) func(int) []int {
		return func(c int) []int { return []int{a, b, c} }
	}
}
func traverseMaybeFailLast1000FFI_concat2(left []int) func([]int) []int {
	return func(right []int) []int {
		result := make([]int, len(left)+len(right))
		copy(result, left)
		copy(result[len(left):], right)
		return result
	}
}

// Match Data.Traversable.traverseArrayImpl: small leaves and balanced
// applicative concatenation. Both branches evaluate even after a Nothing.
func traverseMaybeFailLast1000FFI_traverse(f func(int) traverseMaybeFailLast1000FFI_maybe[int], values []int) traverseMaybeFailLast1000FFI_maybe[[]int] {
	var goRange func(int, int) traverseMaybeFailLast1000FFI_maybe[[]int]
	goRange = func(bottom, top int) traverseMaybeFailLast1000FFI_maybe[[]int] {
		switch top - bottom {
		case 0:
			return traverseMaybeFailLast1000FFI_maybe[[]int]{true, []int{}}
		case 1:
			return traverseMaybeFailLast1000FFI_mapMaybe(traverseMaybeFailLast1000FFI_array1, f(values[bottom]))
		case 2:
			return traverseMaybeFailLast1000FFI_applyMaybe(traverseMaybeFailLast1000FFI_mapMaybe(traverseMaybeFailLast1000FFI_array2, f(values[bottom])), f(values[bottom+1]))
		case 3:
			return traverseMaybeFailLast1000FFI_applyMaybe(traverseMaybeFailLast1000FFI_applyMaybe(traverseMaybeFailLast1000FFI_mapMaybe(traverseMaybeFailLast1000FFI_array3, f(values[bottom])), f(values[bottom+1])), f(values[bottom+2]))
		default:
			pivot := bottom + ((top-bottom)/4)*2
			return traverseMaybeFailLast1000FFI_applyMaybe(traverseMaybeFailLast1000FFI_mapMaybe(traverseMaybeFailLast1000FFI_concat2, goRange(bottom, pivot)), goRange(pivot, top))
		}
	}
	return goRange(0, len(values))
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		values := make([]int, n)
		for i := range values {
			values[i] = i + 1
		}
		return func(count int) int {
			rejected := count
			result := traverseMaybeFailLast1000FFI_traverse(func(value int) traverseMaybeFailLast1000FFI_maybe[int] {
				if value == rejected {
					return traverseMaybeFailLast1000FFI_maybe[int]{}
				}
				return traverseMaybeFailLast1000FFI_maybe[int]{true, value + 1}
			}, values[:count])
			if !result.valid {
				return -1
			}
			return traverseMaybeFailLast1000FFI_weightedChecksum(result.value)
		}
	}
}
