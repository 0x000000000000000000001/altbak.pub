package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Array_Partial_tail gopurs_runtime.Value
var once_Data_Array_Partial_tail sync.Once
func Get_Data_Array_Partial_tail() gopurs_runtime.Value {
	once_Data_Array_Partial_tail.Do(func() {
		cache_Data_Array_Partial_tail = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_Partial_tail(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_Partial_tail
}

var cache_Data_Array_Partial_last gopurs_runtime.Value
var once_Data_Array_Partial_last sync.Once
func Get_Data_Array_Partial_last() gopurs_runtime.Value {
	once_Data_Array_Partial_last.Do(func() {
		cache_Data_Array_Partial_last = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_Partial_last(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_Partial_last
}

var cache_Data_Array_Partial_init gopurs_runtime.Value
var once_Data_Array_Partial_init sync.Once
func Get_Data_Array_Partial_init() gopurs_runtime.Value {
	once_Data_Array_Partial_init.Do(func() {
		cache_Data_Array_Partial_init = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_Partial_init(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_Partial_init
}

var cache_Data_Array_Partial_head gopurs_runtime.Value
var once_Data_Array_Partial_head sync.Once
func Get_Data_Array_Partial_head() gopurs_runtime.Value {
	once_Data_Array_Partial_head.Do(func() {
		cache_Data_Array_Partial_head = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_Partial_head(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_Partial_head
}

func Call_Data_Array_Partial_tail(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_1))).IntVal), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_Partial_last(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return xs_1[(gopurs_runtime.Int(int64(len(xs_1))).IntVal) - (1)]
}

func Call_Data_Array_Partial_init(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_1))).IntVal) - (1)), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_Partial_head(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return xs_1[0]
}


