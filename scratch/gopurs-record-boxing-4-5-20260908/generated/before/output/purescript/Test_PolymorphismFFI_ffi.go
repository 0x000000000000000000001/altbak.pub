package purescript

import "gopurs/output/gopurs_runtime"



type Monoidish interface {
	mempty_() int
	mappend_(int) func(int) int
}

type IntMonoidish struct{}

func (IntMonoidish) mempty_() int {
	return 1
}

func (IntMonoidish) mappend_(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func polyLoop(dict Monoidish, n_init int, acc_init int) int {
	n := n_init
	acc := acc_init
	for n > 0 {
		acc = dict.mappend_(acc)(dict.mempty_())
		n--
	}
	return acc
}

func Test_PolymorphismFFI_RunPolymorphismFFI(limit int) int {
	dummy := limit
	return (polyLoop(IntMonoidish{}, dummy, 0))
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_PolymorphismFFI_RunPolymorphismFFI = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_PolymorphismFFI_RunPolymorphismFFI(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})