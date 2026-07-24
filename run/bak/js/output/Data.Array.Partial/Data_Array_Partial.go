package Data_Array_Partial

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Array "gopurs/output/Data.Array"
)

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func2(func(_dollar__unused_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_sliceImpl(gopurs_runtime.Int(1), gopurs_runtime.Int(int64(len(xs_1.PtrVal.([]gopurs_runtime.Value)))), xs_1)
})
	})
	return tail
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func2(func(_dollar__unused_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ArrayAccess(xs_1, int(gopurs_runtime.Int(int64(len(xs_1.PtrVal.([]gopurs_runtime.Value)))).IntVal - 1))
})
	})
	return last
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func2(func(_dollar__unused_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Array.Call_sliceImpl(gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_1.PtrVal.([]gopurs_runtime.Value)))).IntVal - 1), xs_1)
})
	})
	return init_
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func2(func(_dollar__unused_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ArrayAccess(xs_1, 0)
})
	})
	return head
}




