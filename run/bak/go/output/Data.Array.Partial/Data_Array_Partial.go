package Data_Array_Partial

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Array "gopurs/output/Data.Array"
)

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_tail(_dollar__unused_0_box, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_tail
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_last(_dollar__unused_0_box, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_last
}

var cache_init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		cache_init_ = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_init_(_dollar__unused_0_box, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_init_
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_head(_dollar__unused_0_box, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_head
}

func Call_tail(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []interface{}) []interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_sliceImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_last(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []interface{}) interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
return gopurs_runtime.UnboxAny(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), int((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))).IntVal) - (1))))
}

func Call_init_(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []interface{}) []interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))).IntVal) - (1)), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_head(_dollar__unused_0_loop gopurs_runtime.Value, xs_1_loop []interface{}) interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
return gopurs_runtime.UnboxAny(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), 0))
}
