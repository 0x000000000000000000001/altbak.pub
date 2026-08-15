package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_STArray_sumArray gopurs_runtime.Value
var once_Test_STArray_sumArray sync.Once

func Get_Test_STArray_sumArray() gopurs_runtime.Value {
	once_Test_STArray_sumArray.Do(func() {
		cache_Test_STArray_sumArray = gopurs_runtime.Int(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			arr_0_0 := gopurs_runtime.Apply(Get_Data_Array_ST_newImpl(), gopurs_runtime.Value{})
			_ = arr_0_0
			_dollar___unused_1_1 := gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushAllImpl(), func() gopurs_runtime.Value {
				arr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), arr_0_0)
			_ = _dollar___unused_1_1
			x_2_2 := gopurs_runtime.UncurriedApp3(Get_Data_Array_ST_popImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, arr_0_0)
			_ = x_2_2
			var __t3 gopurs_runtime.Value
			{
				if x_2_2.Type == 9 && x_2_2.IntVal == 930809136 && x_2_2.UnsafePtr != nil {
					__t3 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return (*Constructor_Data_Maybe_Just)(x_2_2.UnsafePtr).V0
					})
					goto end_branch_3
				} else {

				}
			}
			{
				if x_2_2.Type == 9 && x_2_2.IntVal == 930809136 && x_2_2.UnsafePtr == nil {
					__t3 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(0)
					})
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_3:
			return gopurs_runtime.Apply(__t3, gopurs_runtime.Value{})
		})).IntVal)
	})
	return cache_Test_STArray_sumArray
}

var cache_Test_STArray_describe gopurs_runtime.Value
var once_Test_STArray_describe sync.Once

func Get_Test_STArray_describe() gopurs_runtime.Value {
	once_Test_STArray_describe.Do(func() {
		cache_Test_STArray_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("STArray Operations:"))
	})
	return cache_Test_STArray_describe
}

var cache_Test_STArray_act gopurs_runtime.Value
var once_Test_STArray_act sync.Once

func Get_Test_STArray_act() gopurs_runtime.Value {
	once_Test_STArray_act.Do(func() {
		cache_Test_STArray_act = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Get_Test_STArray_sumArray().IntVal)).StrVal()))
	})
	return cache_Test_STArray_act
}
