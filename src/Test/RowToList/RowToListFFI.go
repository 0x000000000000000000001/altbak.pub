package Test_RowToListFFI

// A typed heterogeneous record spine models PureScript's row list. The labels
// and values are present, while the dictionary supplies the type-level count.
type rowNil struct{}
type rowCons[A, Tail any] struct {
	label string
	value A
	tail  Tail
}
type proxy[R any] struct{}
type RecordKeys[R any] interface{ keysImpl(proxy[R]) int }
type keysNil struct{}

func (keysNil) keysImpl(_ proxy[rowNil]) int { return 0 }

type keysCons[A, Tail any] struct{ tail RecordKeys[Tail] }

func (dict keysCons[A, Tail]) keysImpl(_ proxy[rowCons[A, Tail]]) int {
	return 1 + dict.tail.keysImpl(proxy[Tail]{})
}

func keys[R any](dict RecordKeys[R], _ R) int { return dict.keysImpl(proxy[R]{}) }

func RunRowToListFFI(limit int) int {
	type E = rowCons[string, rowNil]
	type D = rowCons[float64, E]
	type C = rowCons[bool, D]
	type B = rowCons[string, C]
	type A = rowCons[int, B]
	record := A{"a", 1, B{"b", "two", C{"c", true, D{"d", 4.0, E{"e", "five", rowNil{}}}}}}
	dict := keysCons[int, B]{keysCons[string, C]{keysCons[bool, D]{keysCons[float64, E]{keysCons[string, rowNil]{keysNil{}}}}}}
	return keys[A](dict, record)
}
