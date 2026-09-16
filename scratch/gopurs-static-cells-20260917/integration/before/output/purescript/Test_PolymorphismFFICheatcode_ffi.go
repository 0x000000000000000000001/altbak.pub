package purescript

import "gopurs/output/gopurs_runtime"


type Monoidish_Cheatcode interface {
	mempty_() int
	mappend_(int, int) int
}
type IntMonoidish_Cheatcode struct{}
func (IntMonoidish_Cheatcode) mempty_() int { return 1 }
func (IntMonoidish_Cheatcode) mappend_(x, y int) int { return x + y }

func Test_PolymorphismFFICheatcode_RunPolymorphismFFICheatcode(limit int) int {
	n := int(limit)
	acc := 0
	var m Monoidish_Cheatcode = IntMonoidish_Cheatcode{}
	for i := 0; i < n; i++ {
		acc = m.mappend_(acc, m.mempty_())
	}
	return acc
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_PolymorphismFFICheatcode_RunPolymorphismFFICheatcode = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_PolymorphismFFICheatcode_RunPolymorphismFFICheatcode(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})