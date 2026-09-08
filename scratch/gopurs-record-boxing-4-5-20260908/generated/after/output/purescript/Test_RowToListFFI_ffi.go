package purescript

import "gopurs/output/gopurs_runtime"



type RecordKeys interface {
	keysImpl(interface{}) int
}

type dictNil struct{}
func (dictNil) keysImpl(_ interface{}) int {
	return 0
}

type dictCons struct {
	tail RecordKeys
}
func (d dictCons) keysImpl(_ interface{}) int {
	return 1 + d.tail.keysImpl(nil)
}

func Test_RowToListFFI_RunRowToListFFI(limit int) int {
	// dummy := limit
	// rec is not even used in keysImpl, it's just for the type signature
	dict := dictCons{tail: dictCons{tail: dictCons{tail: dictCons{tail: dictCons{tail: dictNil{}}}}}}
	return (dict.keysImpl(nil))
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_RowToListFFI_RunRowToListFFI = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_RowToListFFI_RunRowToListFFI(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})