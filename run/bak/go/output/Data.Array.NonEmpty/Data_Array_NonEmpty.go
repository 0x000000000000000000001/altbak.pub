package Data_Array_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Array_NonEmpty_Internal "gopurs/output/Data.Array.NonEmpty.Internal"
	pkg_Data_Array "gopurs/output/Data.Array"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Data_Array_ST "gopurs/output/Data.Array.ST"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	unsafe "unsafe"
)

var cache_max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		cache_max = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}))
_ = __local_var_0_0
return gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply2(__local_var_0_0, x_1, y_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 1527465420) {
__t2 = y_2
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 902936544) {
__t2 = x_1
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 380165415) {
__t2 = x_1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
}()
	})
	return cache_max
}

var cache_intercalate1 gopurs_runtime.Value
var once_intercalate1 sync.Once
func Get_intercalate1() gopurs_runtime.Value {
	once_intercalate1.Do(func() {
		cache_intercalate1 = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate1(dictSemigroup_0_box)
})
	})
	return cache_intercalate1
}

var cache_greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		cache_greaterThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}))
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415))
})
}()
	})
	return cache_greaterThan
}

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}))
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420))
})
}()
	})
	return cache_lessThan
}

var cache_transpose gopurs_runtime.Value
var once_transpose sync.Once
func Get_transpose() gopurs_runtime.Value {
	once_transpose.Do(func() {
		cache_transpose = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_transpose(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_transpose
}

var cache_toArray gopurs_runtime.Value
var once_toArray sync.Once
func Get_toArray() gopurs_runtime.Value {
	once_toArray.Do(func() {
		cache_toArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toArray(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
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
	return cache_toArray
}

var cache_toArray__func_arrint64__arrint64_2781090619 gopurs_runtime.Value
var once_toArray__func_arrint64__arrint64_2781090619 sync.Once
func Get_toArray__func_arrint64__arrint64_2781090619() gopurs_runtime.Value {
	once_toArray__func_arrint64__arrint64_2781090619.Do(func() {
		cache_toArray__func_arrint64__arrint64_2781090619 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toArray__func_arrint64__arrint64_2781090619(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_toArray__func_arrint64__arrint64_2781090619
}

var cache_toArray__func_arrstring__arrstring_1949224283 gopurs_runtime.Value
var once_toArray__func_arrstring__arrstring_1949224283 sync.Once
func Get_toArray__func_arrstring__arrstring_1949224283() gopurs_runtime.Value {
	once_toArray__func_arrstring__arrstring_1949224283.Do(func() {
		cache_toArray__func_arrstring__arrstring_1949224283 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toArray__func_arrstring__arrstring_1949224283(func() []string {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_toArray__func_arrstring__arrstring_1949224283
}

var cache_toArray__func_arrarrinterface____arrarrinterface___293675387 gopurs_runtime.Value
var once_toArray__func_arrarrinterface____arrarrinterface___293675387 sync.Once
func Get_toArray__func_arrarrinterface____arrarrinterface___293675387() gopurs_runtime.Value {
	once_toArray__func_arrarrinterface____arrarrinterface___293675387.Do(func() {
		cache_toArray__func_arrarrinterface____arrarrinterface___293675387 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toArray__func_arrarrinterface____arrarrinterface___293675387(func() [][]interface{} {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([][]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}() }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}() }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_toArray__func_arrarrinterface____arrarrinterface___293675387
}

var cache_toArray__func_arrgopurs_runtime_Value__arrgopurs_runtime_Value_2580695675 gopurs_runtime.Value
var once_toArray__func_arrgopurs_runtime_Value__arrgopurs_runtime_Value_2580695675 sync.Once
func Get_toArray__func_arrgopurs_runtime_Value__arrgopurs_runtime_Value_2580695675() gopurs_runtime.Value {
	once_toArray__func_arrgopurs_runtime_Value__arrgopurs_runtime_Value_2580695675.Do(func() {
		cache_toArray__func_arrgopurs_runtime_Value__arrgopurs_runtime_Value_2580695675 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_toArray__func_arrgopurs_runtime_Value__arrgopurs_runtime_Value_2580695675(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_toArray__func_arrgopurs_runtime_Value__arrgopurs_runtime_Value_2580695675
}

var cache_toArray__func_arrinterface____arrinterface___3233560571 gopurs_runtime.Value
var once_toArray__func_arrinterface____arrinterface___3233560571 sync.Once
func Get_toArray__func_arrinterface____arrinterface___3233560571() gopurs_runtime.Value {
	once_toArray__func_arrinterface____arrinterface___3233560571.Do(func() {
		cache_toArray__func_arrinterface____arrinterface___3233560571 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toArray__func_arrinterface____arrinterface___3233560571(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
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
	return cache_toArray__func_arrinterface____arrinterface___3233560571
}

var cache_toArray__func_arrinterface____arrinterface___1797011707 gopurs_runtime.Value
var once_toArray__func_arrinterface____arrinterface___1797011707 sync.Once
func Get_toArray__func_arrinterface____arrinterface___1797011707() gopurs_runtime.Value {
	once_toArray__func_arrinterface____arrinterface___1797011707.Do(func() {
		cache_toArray__func_arrinterface____arrinterface___1797011707 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toArray__func_arrinterface____arrinterface___1797011707(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
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
	return cache_toArray__func_arrinterface____arrinterface___1797011707
}

var cache_unionBy_prime gopurs_runtime.Value
var once_unionBy_prime sync.Once
func Get_unionBy_prime() gopurs_runtime.Value {
	once_unionBy_prime.Do(func() {
		cache_unionBy_prime = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_unionBy_prime(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(eq_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_unionBy_prime
}

var cache_unionBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 gopurs_runtime.Value
var once_unionBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 sync.Once
func Get_unionBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147() gopurs_runtime.Value {
	once_unionBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147.Do(func() {
		cache_unionBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_unionBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(eq_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_unionBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147
}

var cache_union_prime gopurs_runtime.Value
var once_union_prime sync.Once
func Get_union_prime() gopurs_runtime.Value {
	once_union_prime.Do(func() {
		cache_union_prime = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_union_prime(dictEq_0_box)
})
	})
	return cache_union_prime
}

var cache_unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		cache_unionBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_unionBy(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(eq_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_unionBy
}

var cache_unionBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 gopurs_runtime.Value
var once_unionBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 sync.Once
func Get_unionBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147() gopurs_runtime.Value {
	once_unionBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147.Do(func() {
		cache_unionBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_unionBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(eq_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_unionBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147
}

var cache_union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		cache_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_union(dictEq_0_box)
})
	})
	return cache_union
}

var cache_unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		cache_unzip = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Bifunctor.Get_bifunctorTuple(), "bimap"), pkg_Data_Array_NonEmpty_Internal.Get_NonEmptyArray(), pkg_Data_Array_NonEmpty_Internal.Get_NonEmptyArray())
_ = __local_var_0_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func(inner_arg0 []gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Apply(pkg_Data_Array.Get_unzip(), x_1))
}), gopurs_runtime.Array(inner_arg0))
}(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
}()
	})
	return cache_unzip
}

var cache_updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		cache_updateAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_updateAt(i_0_box.IntVal, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_updateAt
}

var cache_zip gopurs_runtime.Value
var once_zip sync.Once
func Get_zip() gopurs_runtime.Value {
	once_zip.Do(func() {
		cache_zip = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_zip(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(ys_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_zip
}

var cache_zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		cache_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_zipWith(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr)
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
	return cache_zipWith
}

var cache_zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		cache_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWithA(dictApplicative_0_box)
})
	})
	return cache_zipWithA
}

var cache_splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		cache_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_splitAt(i_0_box.IntVal, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_splitAt
}

var cache_some gopurs_runtime.Value
var once_some sync.Once
func Get_some() gopurs_runtime.Value {
	once_some.Do(func() {
		cache_some = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_some(dictAlternative_0_box, dictLazy_1_box, gopurs_runtime.UnboxAny(x_2_box)))
})
	})
	return cache_some
}

var cache_snoc_prime gopurs_runtime.Value
var once_snoc_prime sync.Once
func Get_snoc_prime() gopurs_runtime.Value {
	once_snoc_prime.Do(func() {
		cache_snoc_prime = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_snoc_prime(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), gopurs_runtime.UnboxAny(x_1_box))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_snoc_prime
}

var cache_snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		cache_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_snoc(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), gopurs_runtime.UnboxAny(x_1_box))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_snoc
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_singleton(gopurs_runtime.UnboxAny(x_0_box))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_singleton
}

var cache_replicate gopurs_runtime.Value
var once_replicate sync.Once
func Get_replicate() gopurs_runtime.Value {
	once_replicate.Do(func() {
		cache_replicate = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_replicate(i_0_box.IntVal, gopurs_runtime.UnboxAny(x_1_box))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_replicate
}

var cache_range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		cache_range_ = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_range_(x_0_box.IntVal, y_1_box.IntVal)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_range_
}

var cache_prependArray gopurs_runtime.Value
var once_prependArray sync.Once
func Get_prependArray() gopurs_runtime.Value {
	once_prependArray.Do(func() {
		cache_prependArray = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_prependArray(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(ys_1_box.UnsafePtr)
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
	return cache_prependArray
}

var cache_modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		cache_modifyAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_modifyAt(i_0_box.IntVal, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_modifyAt
}

var cache_intersectBy_prime gopurs_runtime.Value
var once_intersectBy_prime sync.Once
func Get_intersectBy_prime() gopurs_runtime.Value {
	once_intersectBy_prime.Do(func() {
		cache_intersectBy_prime = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersectBy_prime(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(eq_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())
})
	})
	return cache_intersectBy_prime
}

var cache_intersectBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 gopurs_runtime.Value
var once_intersectBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 sync.Once
func Get_intersectBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147() gopurs_runtime.Value {
	once_intersectBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147.Do(func() {
		cache_intersectBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersectBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(eq_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())
})
	})
	return cache_intersectBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147
}

var cache_intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		cache_intersectBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_intersectBy(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(eq_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_intersectBy
}

var cache_intersectBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 gopurs_runtime.Value
var once_intersectBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 sync.Once
func Get_intersectBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147() gopurs_runtime.Value {
	once_intersectBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147.Do(func() {
		cache_intersectBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_intersectBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(eq_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_intersectBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147
}

var cache_intersect_prime gopurs_runtime.Value
var once_intersect_prime sync.Once
func Get_intersect_prime() gopurs_runtime.Value {
	once_intersect_prime.Do(func() {
		cache_intersect_prime = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersect_prime(dictEq_0_box)
})
	})
	return cache_intersect_prime
}

var cache_intersect gopurs_runtime.Value
var once_intersect sync.Once
func Get_intersect() gopurs_runtime.Value {
	once_intersect.Do(func() {
		cache_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersect(dictEq_0_box)
})
	})
	return cache_intersect
}

var cache_intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		cache_intercalate = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate(dictSemigroup_0_box)
})
	})
	return cache_intercalate
}

var cache_insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		cache_insertAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_insertAt(i_0_box.IntVal, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_insertAt
}

var cache_fromFoldable1 gopurs_runtime.Value
var once_fromFoldable1 sync.Once
func Get_fromFoldable1() gopurs_runtime.Value {
	once_fromFoldable1.Do(func() {
		cache_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable1(dictFoldable1_0_box)
})
	})
	return cache_fromFoldable1
}

var cache_fromArray gopurs_runtime.Value
var once_fromArray sync.Once
func Get_fromArray() gopurs_runtime.Value {
	once_fromArray.Do(func() {
		cache_fromArray = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromArray(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_fromArray
}

var cache_fromArray__func_arrint64__ptrData_Maybe_Constructor_Just[arrint64]_2195001498 gopurs_runtime.Value
var once_fromArray__func_arrint64__ptrData_Maybe_Constructor_Just[arrint64]_2195001498 sync.Once
func Get_fromArray__func_arrint64__ptrData_Maybe_Constructor_Just[arrint64]_2195001498() gopurs_runtime.Value {
	once_fromArray__func_arrint64__ptrData_Maybe_Constructor_Just[arrint64]_2195001498.Do(func() {
		cache_fromArray__func_arrint64__ptrData_Maybe_Constructor_Just[arrint64]_2195001498 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromArray__func_arrint64__ptrData_Maybe_Constructor_Just[arrint64]_2195001498(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()))}
})
	})
	return cache_fromArray__func_arrint64__ptrData_Maybe_Constructor_Just[arrint64]_2195001498
}

var cache_fromArray__func_arrstring__ptrData_Maybe_Constructor_Just[arrstring]_260997498 gopurs_runtime.Value
var once_fromArray__func_arrstring__ptrData_Maybe_Constructor_Just[arrstring]_260997498 sync.Once
func Get_fromArray__func_arrstring__ptrData_Maybe_Constructor_Just[arrstring]_260997498() gopurs_runtime.Value {
	once_fromArray__func_arrstring__ptrData_Maybe_Constructor_Just[arrstring]_260997498.Do(func() {
		cache_fromArray__func_arrstring__ptrData_Maybe_Constructor_Just[arrstring]_260997498 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromArray__func_arrstring__ptrData_Maybe_Constructor_Just[arrstring]_260997498(func() []string {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()))}
})
	})
	return cache_fromArray__func_arrstring__ptrData_Maybe_Constructor_Just[arrstring]_260997498
}

var cache_fromArray__func_arrarrinterface____ptrData_Maybe_Constructor_Just[arrarrinterface__]_1152949338 gopurs_runtime.Value
var once_fromArray__func_arrarrinterface____ptrData_Maybe_Constructor_Just[arrarrinterface__]_1152949338 sync.Once
func Get_fromArray__func_arrarrinterface____ptrData_Maybe_Constructor_Just[arrarrinterface__]_1152949338() gopurs_runtime.Value {
	once_fromArray__func_arrarrinterface____ptrData_Maybe_Constructor_Just[arrarrinterface__]_1152949338.Do(func() {
		cache_fromArray__func_arrarrinterface____ptrData_Maybe_Constructor_Just[arrarrinterface__]_1152949338 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromArray__func_arrarrinterface____ptrData_Maybe_Constructor_Just[arrarrinterface__]_1152949338(func() [][]interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([][]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}() }
					return unboxed
				}()))}
})
	})
	return cache_fromArray__func_arrarrinterface____ptrData_Maybe_Constructor_Just[arrarrinterface__]_1152949338
}

var cache_fromArray__func_arrinterface____ptrData_Maybe_Constructor_Just[arrinterface__]_1288807770 gopurs_runtime.Value
var once_fromArray__func_arrinterface____ptrData_Maybe_Constructor_Just[arrinterface__]_1288807770 sync.Once
func Get_fromArray__func_arrinterface____ptrData_Maybe_Constructor_Just[arrinterface__]_1288807770() gopurs_runtime.Value {
	once_fromArray__func_arrinterface____ptrData_Maybe_Constructor_Just[arrinterface__]_1288807770.Do(func() {
		cache_fromArray__func_arrinterface____ptrData_Maybe_Constructor_Just[arrinterface__]_1288807770 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromArray__func_arrinterface____ptrData_Maybe_Constructor_Just[arrinterface__]_1288807770(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_fromArray__func_arrinterface____ptrData_Maybe_Constructor_Just[arrinterface__]_1288807770
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(dictFoldable_0_box)
})
	})
	return cache_fromFoldable
}

var cache_transpose_prime gopurs_runtime.Value
var once_transpose_prime sync.Once
func Get_transpose_prime() gopurs_runtime.Value {
	once_transpose_prime.Do(func() {
		cache_transpose_prime = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_transpose_prime(func() [][]interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([][]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}() }
					return unboxed
				}()))}
})
	})
	return cache_transpose_prime
}

var cache_foldr1 gopurs_runtime.Value
var once_foldr1 sync.Once
func Get_foldr1() gopurs_runtime.Value {
	once_foldr1.Do(func() {
		cache_foldr1 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(interface{}, interface{}) interface{}, inner_arg1 []interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Array_NonEmpty_Internal.Get_foldable1NonEmptyArray(), "foldr1"), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := inner_arg1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))
}(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_foldr1
}

var cache_foldl1 gopurs_runtime.Value
var once_foldl1 sync.Once
func Get_foldl1() gopurs_runtime.Value {
	once_foldl1.Do(func() {
		cache_foldl1 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(interface{}, interface{}) interface{}, inner_arg1 []interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Array_NonEmpty_Internal.Get_foldable1NonEmptyArray(), "foldl1"), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := inner_arg1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))
}(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_foldl1
}

var cache_foldMap1 gopurs_runtime.Value
var once_foldMap1 sync.Once
func Get_foldMap1() gopurs_runtime.Value {
	once_foldMap1.Do(func() {
		cache_foldMap1 = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1(dictSemigroup_0_box)
})
	})
	return cache_foldMap1
}

var cache_fold1 gopurs_runtime.Value
var once_fold1 sync.Once
func Get_fold1() gopurs_runtime.Value {
	once_fold1.Do(func() {
		cache_fold1 = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fold1(dictSemigroup_0_box)
})
	})
	return cache_fold1
}

var cache_difference_prime gopurs_runtime.Value
var once_difference_prime sync.Once
func Get_difference_prime() gopurs_runtime.Value {
	once_difference_prime.Do(func() {
		cache_difference_prime = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference_prime(dictEq_0_box)
})
	})
	return cache_difference_prime
}

var cache_difference_prime__func_gopurs_runtime_Value__arrinterface____arrinterface____arrinterface___3309683994 gopurs_runtime.Value
var once_difference_prime__func_gopurs_runtime_Value__arrinterface____arrinterface____arrinterface___3309683994 sync.Once
func Get_difference_prime__func_gopurs_runtime_Value__arrinterface____arrinterface____arrinterface___3309683994() gopurs_runtime.Value {
	once_difference_prime__func_gopurs_runtime_Value__arrinterface____arrinterface____arrinterface___3309683994.Do(func() {
		cache_difference_prime__func_gopurs_runtime_Value__arrinterface____arrinterface____arrinterface___3309683994 = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference_prime__func_gopurs_runtime_Value__arrinterface____arrinterface____arrinterface___3309683994(dictEq_0_box)
})
	})
	return cache_difference_prime__func_gopurs_runtime_Value__arrinterface____arrinterface____arrinterface___3309683994
}

var cache_cons_prime gopurs_runtime.Value
var once_cons_prime sync.Once
func Get_cons_prime() gopurs_runtime.Value {
	once_cons_prime.Do(func() {
		cache_cons_prime = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_cons_prime(gopurs_runtime.UnboxAny(x_0_box), func() []interface{} {
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
	return cache_cons_prime
}

var cache_cons_prime__func_interface____arrinterface____arrinterface___2578448511 gopurs_runtime.Value
var once_cons_prime__func_interface____arrinterface____arrinterface___2578448511 sync.Once
func Get_cons_prime__func_interface____arrinterface____arrinterface___2578448511() gopurs_runtime.Value {
	once_cons_prime__func_interface____arrinterface____arrinterface___2578448511.Do(func() {
		cache_cons_prime__func_interface____arrinterface____arrinterface___2578448511 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_cons_prime__func_interface____arrinterface____arrinterface___2578448511(gopurs_runtime.UnboxAny(x_0_box), func() []interface{} {
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
	return cache_cons_prime__func_interface____arrinterface____arrinterface___2578448511
}

var cache_fromNonEmpty gopurs_runtime.Value
var once_fromNonEmpty sync.Once
func Get_fromNonEmpty() gopurs_runtime.Value {
	once_fromNonEmpty.Do(func() {
		cache_fromNonEmpty = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_fromNonEmpty(v_0_box)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_fromNonEmpty
}

var cache_concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		cache_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_concatMap(func(inner_arg0 interface{}) []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(b_0_box, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(a_1_box.UnsafePtr)
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
	return cache_concatMap
}

var cache_concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		cache_concat = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), Get_toArray())
_ = __local_var_0_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func(inner_arg0 []gopurs_runtime.Value) []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Array.Get_concat(), gopurs_runtime.Apply(__local_var_0_0, x_1))
}), gopurs_runtime.Array(inner_arg0)).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
}()
	})
	return cache_concat
}

var cache_appendArray gopurs_runtime.Value
var once_appendArray sync.Once
func Get_appendArray() gopurs_runtime.Value {
	once_appendArray.Do(func() {
		cache_appendArray = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_appendArray(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(ys_1_box.UnsafePtr)
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
	return cache_appendArray
}

var cache_alterAt gopurs_runtime.Value
var once_alterAt sync.Once
func Get_alterAt() gopurs_runtime.Value {
	once_alterAt.Do(func() {
		cache_alterAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_alterAt(i_0_box.IntVal, func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_alterAt
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_head(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_head
}

var cache_init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		cache_init_ = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_init_(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
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

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_last(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_last
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_tail(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
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

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncons(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_uncons
}

var cache_uncons__func_arrinterface____interface___1551523312 gopurs_runtime.Value
var once_uncons__func_arrinterface____interface___1551523312 sync.Once
func Get_uncons__func_arrinterface____interface___1551523312() gopurs_runtime.Value {
	once_uncons__func_arrinterface____interface___1551523312.Do(func() {
		cache_uncons__func_arrinterface____interface___1551523312 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncons__func_arrinterface____interface___1551523312(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_uncons__func_arrinterface____interface___1551523312
}

var cache_toNonEmpty gopurs_runtime.Value
var once_toNonEmpty sync.Once
func Get_toNonEmpty() gopurs_runtime.Value {
	once_toNonEmpty.Do(func() {
		cache_toNonEmpty = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toNonEmpty(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())
})
	})
	return cache_toNonEmpty
}

var cache_unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		cache_unsnoc = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unsnoc(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_unsnoc
}

var cache_all gopurs_runtime.Value
var once_all sync.Once
func Get_all() gopurs_runtime.Value {
	once_all.Do(func() {
		cache_all = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_all(func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(p_0_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_all
}

var cache_any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		cache_any = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_any(func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(p_0_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_any
}

var cache_catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		cache_catMaybes = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_catMaybes(func() []*pkg_Data_Maybe.Constructor_Just[interface{}] {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]*pkg_Data_Maybe.Constructor_Just[interface{}], len(arr))
					for i, v := range arr { unboxed[i] = (*pkg_Data_Maybe.Constructor_Just[interface{}])(v.UnsafePtr) }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_catMaybes
}

var cache_delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		cache_delete_ = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_delete_(dictEq_0_box, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_delete_
}

var cache_deleteAt gopurs_runtime.Value
var once_deleteAt sync.Once
func Get_deleteAt() gopurs_runtime.Value {
	once_deleteAt.Do(func() {
		cache_deleteAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_deleteAt(i_0_box.IntVal, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_deleteAt
}

var cache_deleteBy gopurs_runtime.Value
var once_deleteBy sync.Once
func Get_deleteBy() gopurs_runtime.Value {
	once_deleteBy.Do(func() {
		cache_deleteBy = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_deleteBy(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_deleteBy
}

var cache_difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		cache_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference(dictEq_0_box)
})
	})
	return cache_difference
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_drop(i_0_box.IntVal, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_drop
}

var cache_dropEnd gopurs_runtime.Value
var once_dropEnd sync.Once
func Get_dropEnd() gopurs_runtime.Value {
	once_dropEnd.Do(func() {
		cache_dropEnd = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_dropEnd(i_0_box.IntVal, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_dropEnd
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_dropWhile(func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_dropWhile
}

var cache_elem gopurs_runtime.Value
var once_elem sync.Once
func Get_elem() gopurs_runtime.Value {
	once_elem.Do(func() {
		cache_elem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_elem(dictEq_0_box, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_elem
}

var cache_elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		cache_elemIndex = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_elemIndex(dictEq_0_box, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_elemIndex
}

var cache_elemLastIndex gopurs_runtime.Value
var once_elemLastIndex sync.Once
func Get_elemLastIndex() gopurs_runtime.Value {
	once_elemLastIndex.Do(func() {
		cache_elemLastIndex = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_elemLastIndex(dictEq_0_box, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_elemLastIndex
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_filter(func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_filter
}

var cache_filterA gopurs_runtime.Value
var once_filterA sync.Once
func Get_filterA() gopurs_runtime.Value {
	once_filterA.Do(func() {
		cache_filterA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterA(dictApplicative_0_box)
})
	})
	return cache_filterA
}

var cache_find gopurs_runtime.Value
var once_find sync.Once
func Get_find() gopurs_runtime.Value {
	once_find.Do(func() {
		cache_find = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_find(func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(p_0_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_find
}

var cache_findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		cache_findIndex = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findIndex(func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(p_0_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_findIndex
}

var cache_findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		cache_findLastIndex = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findLastIndex(func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(x_0_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_findLastIndex
}

var cache_findMap gopurs_runtime.Value
var once_findMap sync.Once
func Get_findMap() gopurs_runtime.Value {
	once_findMap.Do(func() {
		cache_findMap = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMap(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(p_0_box, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))}
})
	})
	return cache_findMap
}

var cache_foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		cache_foldM = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, acc_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_foldM(dictMonad_0_box, func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_1_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(acc_2_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_3_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_foldM
}

var cache_foldRecM gopurs_runtime.Value
var once_foldRecM sync.Once
func Get_foldRecM() gopurs_runtime.Value {
	once_foldRecM.Do(func() {
		cache_foldRecM = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldRecM(dictMonadRec_0_box)
})
	})
	return cache_foldRecM
}

var cache_index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		cache_index = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_index(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())
})
	})
	return cache_index
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_length(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_length
}

var cache_length__func_arrinterface____int64_4085341461 gopurs_runtime.Value
var once_length__func_arrinterface____int64_4085341461 sync.Once
func Get_length__func_arrinterface____int64_4085341461() gopurs_runtime.Value {
	once_length__func_arrinterface____int64_4085341461.Do(func() {
		cache_length__func_arrinterface____int64_4085341461 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_length__func_arrinterface____int64_4085341461(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_length__func_arrinterface____int64_4085341461
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_mapMaybe(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_mapMaybe
}

var cache_notElem gopurs_runtime.Value
var once_notElem sync.Once
func Get_notElem() gopurs_runtime.Value {
	once_notElem.Do(func() {
		cache_notElem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notElem(dictEq_0_box, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_notElem
}

var cache_partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		cache_partition = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_partition(func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_partition
}

var cache_slice gopurs_runtime.Value
var once_slice sync.Once
func Get_slice() gopurs_runtime.Value {
	once_slice.Do(func() {
		cache_slice = gopurs_runtime.Func3(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_slice(start_0_box.IntVal, end_1_box.IntVal, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_slice
}

var cache_span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		cache_span = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_span(func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_span
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_take(i_0_box.IntVal, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_take
}

var cache_takeEnd gopurs_runtime.Value
var once_takeEnd sync.Once
func Get_takeEnd() gopurs_runtime.Value {
	once_takeEnd.Do(func() {
		cache_takeEnd = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_takeEnd(i_0_box.IntVal, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_takeEnd
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_takeWhile(func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_takeWhile
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func2(func(dictUnfoldable_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_toUnfoldable(dictUnfoldable_0_box, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_toUnfoldable
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_cons(gopurs_runtime.UnboxAny(x_0_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_cons
}

var cache_group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		cache_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_group(dictEq_0_box)
})
	})
	return cache_group
}

var cache_groupAllBy gopurs_runtime.Value
var once_groupAllBy sync.Once
func Get_groupAllBy() gopurs_runtime.Value {
	once_groupAllBy.Do(func() {
		cache_groupAllBy = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupAllBy(func(inner_arg0 interface{}, inner_arg1 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1))
})
})
	})
	return cache_groupAllBy
}

var cache_groupAll gopurs_runtime.Value
var once_groupAll sync.Once
func Get_groupAll() gopurs_runtime.Value {
	once_groupAll.Do(func() {
		cache_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupAll(dictOrd_0_box)
})
	})
	return cache_groupAll
}

var cache_groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		cache_groupBy = gopurs_runtime.Func2(func(op_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_groupBy(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(op_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_groupBy
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_insert(dictOrd_0_box, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_insert
}

var cache_insertBy gopurs_runtime.Value
var once_insertBy sync.Once
func Get_insertBy() gopurs_runtime.Value {
	once_insertBy.Do(func() {
		cache_insertBy = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_insertBy(func(inner_arg0 interface{}, inner_arg1 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1))
}, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_insertBy
}

var cache_intersperse gopurs_runtime.Value
var once_intersperse sync.Once
func Get_intersperse() gopurs_runtime.Value {
	once_intersperse.Do(func() {
		cache_intersperse = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_intersperse(gopurs_runtime.UnboxAny(x_0_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_intersperse
}

var cache_mapWithIndex gopurs_runtime.Value
var once_mapWithIndex sync.Once
func Get_mapWithIndex() gopurs_runtime.Value {
	once_mapWithIndex.Do(func() {
		cache_mapWithIndex = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex(func(inner_arg0 int64, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Int(inner_arg0), gopurs_runtime.Any(inner_arg1)))
})
})
	})
	return cache_mapWithIndex
}

var cache_modifyAtIndices gopurs_runtime.Value
var once_modifyAtIndices sync.Once
func Get_modifyAtIndices() gopurs_runtime.Value {
	once_modifyAtIndices.Do(func() {
		cache_modifyAtIndices = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyAtIndices(dictFoldable_0_box)
})
	})
	return cache_modifyAtIndices
}

var cache_nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		cache_nub = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_nub(dictOrd_0_box, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_nub
}

var cache_nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		cache_nubBy = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_nubBy(func(inner_arg0 interface{}, inner_arg1 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1))
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_nubBy
}

var cache_nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		cache_nubByEq = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_nubByEq(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_nubByEq
}

var cache_nubEq gopurs_runtime.Value
var once_nubEq sync.Once
func Get_nubEq() gopurs_runtime.Value {
	once_nubEq.Do(func() {
		cache_nubEq = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_nubEq(dictEq_0_box, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_nubEq
}

var cache_reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		cache_reverse = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_reverse(func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
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
	return cache_reverse
}

var cache_scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		cache_scanl = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_scanl(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_scanl
}

var cache_scanr gopurs_runtime.Value
var once_scanr sync.Once
func Get_scanr() gopurs_runtime.Value {
	once_scanr.Do(func() {
		cache_scanr = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_scanr(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(x_1_box), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_scanr
}

var cache_sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		cache_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sort(dictOrd_0_box)
})
	})
	return cache_sort
}

var cache_sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		cache_sortBy = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_sortBy(func(inner_arg0 interface{}, inner_arg1 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1))
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
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
	return cache_sortBy
}

var cache_sortWith gopurs_runtime.Value
var once_sortWith sync.Once
func Get_sortWith() gopurs_runtime.Value {
	once_sortWith.Do(func() {
		cache_sortWith = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_sortWith(dictOrd_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
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
	return cache_sortWith
}

var cache_updateAtIndices gopurs_runtime.Value
var once_updateAtIndices sync.Once
func Get_updateAtIndices() gopurs_runtime.Value {
	once_updateAtIndices.Do(func() {
		cache_updateAtIndices = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_updateAtIndices(dictFoldable_0_box)
})
	})
	return cache_updateAtIndices
}

var cache_unsafeIndex gopurs_runtime.Value
var once_unsafeIndex sync.Once
func Get_unsafeIndex() gopurs_runtime.Value {
	once_unsafeIndex.Do(func() {
		cache_unsafeIndex = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unsafeIndex(_dollar__unused_0_box, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), __local_var_2_box.IntVal))
})
	})
	return cache_unsafeIndex
}

var cache_unsafeIndex__func_gopurs_runtime_Value__arrinterface____int64__interface___3957585020 gopurs_runtime.Value
var once_unsafeIndex__func_gopurs_runtime_Value__arrinterface____int64__interface___3957585020 sync.Once
func Get_unsafeIndex__func_gopurs_runtime_Value__arrinterface____int64__interface___3957585020() gopurs_runtime.Value {
	once_unsafeIndex__func_gopurs_runtime_Value__arrinterface____int64__interface___3957585020.Do(func() {
		cache_unsafeIndex__func_gopurs_runtime_Value__arrinterface____int64__interface___3957585020 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unsafeIndex__func_gopurs_runtime_Value__arrinterface____int64__interface___3957585020(_dollar__unused_0_box, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), __local_var_2_box.IntVal))
})
	})
	return cache_unsafeIndex__func_gopurs_runtime_Value__arrinterface____int64__interface___3957585020
}

var cache_toUnfoldable1 gopurs_runtime.Value
var once_toUnfoldable1 sync.Once
func Get_toUnfoldable1() gopurs_runtime.Value {
	once_toUnfoldable1.Do(func() {
		cache_toUnfoldable1 = gopurs_runtime.Func2(func(dictUnfoldable1_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_toUnfoldable1(dictUnfoldable1_0_box, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_toUnfoldable1
}

func Call_intercalate1(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
foldMap12_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Array_NonEmpty_Internal.Get_foldable1NonEmptyArray(), "foldMap1"), gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(v_1, j_3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), j_3, gopurs_runtime.Apply(v1_2, j_3)))
})))
_ = foldMap12_1_0
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, foldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(foldMap12_1_0, gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}), foldable_3, a_2)
})
}

func Call_transpose(x_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Data_Array.Get_transpose(), gopurs_runtime.Array(x_0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_toArray(v_0_loop []interface{}) []interface{} {
var v_0 []interface{} = v_0_loop
_ = v_0
return v_0
}

func Call_toArray__func_arrint64__arrint64_2781090619(v_0_loop []int64) []int64 {
var v_0 []int64 = v_0_loop
_ = v_0
return v_0
}

func Call_toArray__func_arrstring__arrstring_1949224283(v_0_loop []string) []string {
var v_0 []string = v_0_loop
_ = v_0
return v_0
}

func Call_toArray__func_arrarrinterface____arrarrinterface___293675387(v_0_loop [][]interface{}) [][]interface{} {
var v_0 [][]interface{} = v_0_loop
_ = v_0
return v_0
}

func Call_toArray__func_arrgopurs_runtime_Value__arrgopurs_runtime_Value_2580695675(v_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var v_0 []gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_toArray__func_arrinterface____arrinterface___3233560571(v_0_loop []interface{}) []interface{} {
var v_0 []interface{} = v_0_loop
_ = v_0
return v_0
}

func Call_toArray__func_arrinterface____arrinterface___1797011707(v_0_loop []interface{}) []interface{} {
var v_0 []interface{} = v_0_loop
_ = v_0
return v_0
}

func Call_unionBy_prime(eq_0_loop func(interface{}, interface{}) bool, xs_1_loop []interface{}, x_2_loop []interface{}) []interface{} {
var eq_0 func(interface{}, interface{}) bool = eq_0_loop
_ = eq_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Array.Get_unionBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(eq_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_unionBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147(eq_0_loop func(interface{}, interface{}) bool, xs_1_loop []interface{}, x_2_loop []interface{}) []interface{} {
var eq_0 func(interface{}, interface{}) bool = eq_0_loop
_ = eq_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Array.Get_unionBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(eq_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_union_prime(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_unionBy_prime(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_unionBy(eq_0_loop func(interface{}, interface{}) bool, xs_1_loop []interface{}, x_2_loop []interface{}) []interface{} {
var eq_0 func(interface{}, interface{}) bool = eq_0_loop
_ = eq_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Array.Get_unionBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(eq_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_unionBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147(eq_0_loop func(interface{}, interface{}) bool, xs_1_loop []interface{}, x_2_loop []interface{}) []interface{} {
var eq_0 func(interface{}, interface{}) bool = eq_0_loop
_ = eq_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Array.Get_unionBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(eq_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_union(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_unionBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_updateAt(i_0_loop int64, x_1_loop interface{}, x_2_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[[]interface{}] {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return (*pkg_Data_Maybe.Constructor_Just[[]interface{}])(gopurs_runtime.UncurriedApp5(pkg_Data_Array.Get__updateAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), gopurs_runtime.Int(i_0), gopurs_runtime.Any(x_1), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
}

func Call_zip(xs_0_loop []interface{}, ys_1_loop []interface{}) []gopurs_runtime.Value {
var xs_0 []interface{} = xs_0_loop
_ = xs_0
var ys_1 []interface{} = ys_1_loop
_ = ys_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_zipWithImpl(), pkg_Data_Tuple.Get_Tuple(), func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := ys_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_zipWith(f_0_loop func(interface{}, interface{}) interface{}, xs_1_loop []interface{}, ys_2_loop []interface{}) []interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
var ys_2 []interface{} = ys_2_loop
_ = ys_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_zipWithImpl(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := ys_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_zipWithA(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply(pkg_Data_Array.Get_zipWithA(), dictApplicative_0)
}

func Call_splitAt(i_0_loop int64, xs_1_loop []interface{}) interface{} {
var i_0 int64 = i_0_loop
_ = i_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(pkg_Data_Array.Get_splitAt(), gopurs_runtime.Int(i_0), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))
}

func Call_some(dictAlternative_0_loop gopurs_runtime.Value, dictLazy_1_loop gopurs_runtime.Value, x_2_loop interface{}) interface{} {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
var dictLazy_1 gopurs_runtime.Value = dictLazy_1_loop
_ = dictLazy_1
var x_2 interface{} = x_2_loop
_ = x_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply3(pkg_Data_Array.Get_some(), dictAlternative_0, dictLazy_1, gopurs_runtime.Any(x_2)))
}

func Call_snoc_prime(xs_0_loop []interface{}, x_1_loop interface{}) []interface{} {
var xs_0 []interface{} = xs_0_loop
_ = xs_0
var x_1 interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), gopurs_runtime.Any(x_1)), func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_snoc(xs_0_loop []interface{}, x_1_loop interface{}) []interface{} {
var xs_0 []interface{} = xs_0_loop
_ = xs_0
var x_1 interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), gopurs_runtime.Any(x_1)), func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_singleton(x_0_loop interface{}) []interface{} {
var x_0 interface{} = x_0_loop
_ = x_0
return []interface{}{x_0}
}

func Call_replicate(i_0_loop int64, x_1_loop interface{}) []interface{} {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_replicateImpl(), gopurs_runtime.Apply2(Get_max(), gopurs_runtime.Int(1), gopurs_runtime.Int(i_0)), gopurs_runtime.Any(x_1)).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_range_(x_0_loop int64, y_1_loop int64) []int64 {
var x_0 int64 = x_0_loop
_ = x_0
var y_1 int64 = y_1_loop
_ = y_1
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_rangeImpl(), gopurs_runtime.Int(x_0), gopurs_runtime.Int(y_1)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_prependArray(xs_0_loop []interface{}, ys_1_loop []interface{}) []interface{} {
var xs_0 []interface{} = xs_0_loop
_ = xs_0
var ys_1 []interface{} = ys_1_loop
_ = ys_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := ys_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_modifyAt(i_0_loop int64, f_1_loop func(interface{}) interface{}, x_2_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[[]interface{}] {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var x_2 []interface{} = x_2_loop
_ = x_2
return (*pkg_Data_Maybe.Constructor_Just[[]interface{}])(gopurs_runtime.Apply3(pkg_Data_Array.Get_modifyAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
}

func Call_intersectBy_prime(eq_0_loop func(interface{}, interface{}) bool, xs_1_loop []interface{}) gopurs_runtime.Value {
var eq_0 func(interface{}, interface{}) bool = eq_0_loop
_ = eq_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(pkg_Data_Array.Get_intersectBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(eq_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
}

func Call_intersectBy_prime__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147(eq_0_loop func(interface{}, interface{}) bool, xs_1_loop []interface{}) gopurs_runtime.Value {
var eq_0 func(interface{}, interface{}) bool = eq_0_loop
_ = eq_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(pkg_Data_Array.Get_intersectBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(eq_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
}

func Call_intersectBy(eq_0_loop func(interface{}, interface{}) bool, xs_1_loop []interface{}, x_2_loop []interface{}) []interface{} {
var eq_0 func(interface{}, interface{}) bool = eq_0_loop
_ = eq_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Array.Get_intersectBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(eq_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_intersectBy__func_func_interface____interface____bool__arrinterface____arrinterface____arrinterface___3048774147(eq_0_loop func(interface{}, interface{}) bool, xs_1_loop []interface{}, x_2_loop []interface{}) []interface{} {
var eq_0 func(interface{}, interface{}) bool = eq_0_loop
_ = eq_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Array.Get_intersectBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(eq_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_intersect_prime(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_intersectBy_prime(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_intersect(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_intersectBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_intercalate(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return Call_intercalate1(dictSemigroup_0)
}

func Call_insertAt(i_0_loop int64, x_1_loop interface{}, x_2_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[[]interface{}] {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return (*pkg_Data_Maybe.Constructor_Just[[]interface{}])(gopurs_runtime.UncurriedApp5(pkg_Data_Array.Get__insertAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), gopurs_runtime.Int(i_0), gopurs_runtime.Any(x_1), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
}

func Call_fromFoldable1(dictFoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
__local_var_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "Foldable0"), gopurs_runtime.Value{}), "foldr")
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_fromFoldableImpl(), __local_var_1_0, x_2)
})
}

func Call_fromArray(xs_0_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[[]interface{}] {
var xs_0 []interface{} = xs_0_loop
_ = xs_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))), gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[[]interface{}])(__t0.UnsafePtr)
}

func Call_fromArray__func_arrint64__ptrData_Maybe_Constructor_Just[arrint64]_2195001498(xs_0_loop []int64) *pkg_Data_Maybe.Constructor_Just[[]int64] {
var xs_0 []int64 = xs_0_loop
_ = xs_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()))), gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}())})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[[]int64])(__t0.UnsafePtr)
}

func Call_fromArray__func_arrstring__ptrData_Maybe_Constructor_Just[arrstring]_260997498(xs_0_loop []string) *pkg_Data_Maybe.Constructor_Just[[]string] {
var xs_0 []string = xs_0_loop
_ = xs_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()))), gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}())})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[[]string])(__t0.UnsafePtr)
}

func Call_fromArray__func_arrarrinterface____ptrData_Maybe_Constructor_Just[arrarrinterface__]_1152949338(xs_0_loop [][]interface{}) *pkg_Data_Maybe.Constructor_Just[[][]interface{}] {
var xs_0 [][]interface{} = xs_0_loop
_ = xs_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}() }
					return gopurs_runtime.Array(boxed)
				}()))), gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}() }
					return gopurs_runtime.Array(boxed)
				}())})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[[][]interface{}])(__t0.UnsafePtr)
}

func Call_fromArray__func_arrinterface____ptrData_Maybe_Constructor_Just[arrinterface__]_1288807770(xs_0_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[[]interface{}] {
var xs_0 []interface{} = xs_0_loop
_ = xs_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))), gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[[]interface{}])(__t0.UnsafePtr)
}

func Call_fromFoldable(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictFoldable_0, "foldr")
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_fromFoldableImpl(), __local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_3_1))), gopurs_runtime.Int(0)).IntVal) != (0) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(__local_var_3_1)})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_2:
return __t2
})
}

func Call_transpose_prime(x_0_loop [][]interface{}) *pkg_Data_Maybe.Constructor_Just[[][]interface{}] {
var x_0 [][]interface{} = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_Array.Get_transpose(), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}() }
					return gopurs_runtime.Array(boxed)
				}())
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_1_0))), gopurs_runtime.Int(0)).IntVal) != (0) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(__local_var_1_0)})})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_1:
return (*pkg_Data_Maybe.Constructor_Just[[][]interface{}])(__t1.UnsafePtr)
}

func Call_foldMap1(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Array_NonEmpty_Internal.Get_foldable1NonEmptyArray(), "foldMap1"), dictSemigroup_0)
}

func Call_fold1(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Array_NonEmpty_Internal.Get_foldable1NonEmptyArray(), "foldMap1"), dictSemigroup_0, pkg_Data_Semigroup_Foldable.Get_identity())
}

func Call_difference_prime(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Apply(pkg_Data_Array.Get_delete_(), dictEq_0))
}

func Call_difference_prime__func_gopurs_runtime_Value__arrinterface____arrinterface____arrinterface___3309683994(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Apply(pkg_Data_Array.Get_delete_(), dictEq_0))
}

func Call_cons_prime(x_0_loop interface{}, xs_1_loop []interface{}) []interface{} {
var x_0 interface{} = x_0_loop
_ = x_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), func() gopurs_runtime.Value {
					arr := []interface{}{x_0}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
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

func Call_cons_prime__func_interface____arrinterface____arrinterface___2578448511(x_0_loop interface{}, xs_1_loop []interface{}) []interface{} {
var x_0 interface{} = x_0_loop
_ = x_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), func() gopurs_runtime.Value {
					arr := []interface{}{x_0}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
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

func Call_fromNonEmpty(v_0_loop gopurs_runtime.Value) []interface{} {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Any((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0)}), gopurs_runtime.Any((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_concatMap(b_0_loop func(interface{}) []interface{}, a_1_loop []interface{}) []interface{} {
var b_0 func(interface{}) []interface{} = b_0_loop
_ = b_0
var a_1 []interface{} = a_1_loop
_ = a_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_bindArray(), "bind"), func() gopurs_runtime.Value {
					arr := a_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := b_0(gopurs_runtime.UnboxAny(arg0))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_appendArray(xs_0_loop []interface{}, ys_1_loop []interface{}) []interface{} {
var xs_0 []interface{} = xs_0_loop
_ = xs_0
var ys_1 []interface{} = ys_1_loop
_ = ys_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := ys_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_alterAt(i_0_loop int64, f_1_loop func(interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}], x_2_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[[]interface{}] {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 func(interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] = f_1_loop
_ = f_1
var x_2 []interface{} = x_2_loop
_ = x_2
return (*pkg_Data_Maybe.Constructor_Just[[]interface{}])(gopurs_runtime.Apply3(pkg_Data_Array.Get_alterAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(f_1(gopurs_runtime.UnboxAny(arg0)))}
}), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
}

func Call_head(x_0_loop []interface{}) interface{} {
var x_0 []interface{} = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.UncurriedApp4(pkg_Data_Array.Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(0))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.UnboxAny(__t1)
}

func Call_init_(x_0_loop []interface{}) []interface{} {
var x_0 []interface{} = x_0_loop
_ = x_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))).IntVal) == (0) {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))).IntVal) - (1)), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
}
end_branch_0:
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(__t0.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_last(x_0_loop []interface{}) interface{} {
var x_0 []interface{} = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.UncurriedApp4(pkg_Data_Array.Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))).IntVal) - (1)))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.UnboxAny(__t1)
}

func Call_tail(x_0_loop []interface{}) []interface{} {
var x_0 []interface{} = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(xs_2)})})
}), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(__t1.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_uncons(x_0_loop []interface{}) interface{} {
var x_0 []interface{} = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("head", "tail", x_1, xs_2))})})
}), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.UnboxAny(__t1)
}

func Call_uncons__func_arrinterface____interface___1551523312(x_0_loop []interface{}) interface{} {
var x_0 []interface{} = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("head", "tail", x_1, xs_2))})})
}), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.UnboxAny(__t1)
}

func Call_toNonEmpty(x_0_loop []interface{}) gopurs_runtime.Value {
var x_0 []interface{} = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Any(Call_uncons(x_0))
_ = __local_var_1_0
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(__local_var_1_0, "head")), gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(__local_var_1_0, "tail"))})})
}

func Call_unsnoc(x_0_loop []interface{}) interface{} {
var x_0 []interface{} = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_Array.Get_unsnoc(), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.UnboxAny(__t1)
}

func Call_all(p_0_loop func(interface{}) bool, x_1_loop []interface{}) bool {
var p_0 func(interface{}) bool = p_0_loop
_ = p_0
var x_1 []interface{} = x_1_loop
_ = x_1
return (gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_allImpl(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(p_0(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0)
}

func Call_any(p_0_loop func(interface{}) bool, x_1_loop []interface{}) bool {
var p_0 func(interface{}) bool = p_0_loop
_ = p_0
var x_1 []interface{} = x_1_loop
_ = x_1
return (gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_anyImpl(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(p_0(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0)
}

func Call_catMaybes(x_0_loop []*pkg_Data_Maybe.Constructor_Just[interface{}]) []interface{} {
var x_0 []*pkg_Data_Maybe.Constructor_Just[interface{}] = x_0_loop
_ = x_0
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(pkg_Data_Array.Get_mapMaybe(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_delete_(dictEq_0_loop gopurs_runtime.Value, x_1_loop interface{}, x_2_loop []interface{}) []interface{} {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Array.Get_deleteBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.Any(x_1), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_deleteAt(i_0_loop int64, x_1_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[[]interface{}] {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []interface{} = x_1_loop
_ = x_1
return (*pkg_Data_Maybe.Constructor_Just[[]interface{}])(gopurs_runtime.UncurriedApp4(pkg_Data_Array.Get__deleteAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), gopurs_runtime.Int(i_0), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
}

func Call_deleteBy(f_0_loop func(interface{}, interface{}) bool, x_1_loop interface{}, x_2_loop []interface{}) []interface{} {
var f_0 func(interface{}, interface{}) bool = f_0_loop
_ = f_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Array.Get_deleteBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(f_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), gopurs_runtime.Any(x_1), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_difference(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Apply(pkg_Data_Array.Get_delete_(), dictEq_0))
}

func Call_drop(i_0_loop int64, x_1_loop []interface{}) []interface{} {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []interface{} = x_1_loop
_ = x_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(pkg_Data_Array.Get_lessThan(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(1)).IntVal) != (0) {
__t0 = func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_sliceImpl(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
}
end_branch_0:
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(__t0.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_dropEnd(i_0_loop int64, x_1_loop []interface{}) []interface{} {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(pkg_Data_Array.Get_dropEnd(), gopurs_runtime.Int(i_0), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_dropWhile(f_0_loop func(interface{}) bool, x_1_loop []interface{}) []interface{} {
var f_0 func(interface{}) bool = f_0_loop
_ = f_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Array.Get_span(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(f_0(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()), "rest").UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_elem(dictEq_0_loop gopurs_runtime.Value, x_1_loop interface{}, x_2_loop []interface{}) bool {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return (gopurs_runtime.Apply3(pkg_Data_Array.Get_elem(), dictEq_0, gopurs_runtime.Any(x_1), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0)
}

func Call_elemIndex(dictEq_0_loop gopurs_runtime.Value, x_1_loop interface{}, x_2_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[int64] {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.UncurriedApp4(pkg_Data_Array.Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_3, gopurs_runtime.Any(x_1))
}), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
}

func Call_elemLastIndex(dictEq_0_loop gopurs_runtime.Value, x_1_loop interface{}, x_2_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[int64] {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.UncurriedApp4(pkg_Data_Array.Get_findLastIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_3, gopurs_runtime.Any(x_1))
}), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
}

func Call_filter(f_0_loop func(interface{}) bool, x_1_loop []interface{}) []interface{} {
var f_0 func(interface{}) bool = f_0_loop
_ = f_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_filterImpl(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(f_0(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_filterA(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply(pkg_Data_Array.Get_filterA(), dictApplicative_0)
}

func Call_find(p_0_loop func(interface{}) bool, x_1_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
var p_0 func(interface{}) bool = p_0_loop
_ = p_0
var x_1 []interface{} = x_1_loop
_ = x_1
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), int(__local_var_2.IntVal))
}), gopurs_runtime.UncurriedApp4(pkg_Data_Array.Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(p_0(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())).UnsafePtr)
}

func Call_findIndex(p_0_loop func(interface{}) bool, x_1_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[int64] {
var p_0 func(interface{}) bool = p_0_loop
_ = p_0
var x_1 []interface{} = x_1_loop
_ = x_1
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.UncurriedApp4(pkg_Data_Array.Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(p_0(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
}

func Call_findLastIndex(x_0_loop func(interface{}) bool, x_1_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[int64] {
var x_0 func(interface{}) bool = x_0_loop
_ = x_0
var x_1 []interface{} = x_1_loop
_ = x_1
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.UncurriedApp4(pkg_Data_Array.Get_findLastIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(x_0(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
}

func Call_findMap(p_0_loop func(interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}], x_1_loop []interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
var p_0 func(interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] = p_0_loop
_ = p_0
var x_1 []interface{} = x_1_loop
_ = x_1
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.UncurriedApp4(pkg_Data_Array.Get_findMapImpl(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}), pkg_Data_Maybe.Get_isJust(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(p_0(gopurs_runtime.UnboxAny(arg0)))}
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
}

func Call_foldM(dictMonad_0_loop gopurs_runtime.Value, f_1_loop func(interface{}, interface{}) interface{}, acc_2_loop interface{}, x_3_loop []interface{}) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 func(interface{}, interface{}) interface{} = f_1_loop
_ = f_1
var acc_2 interface{} = acc_2_loop
_ = acc_2
var x_3 []interface{} = x_3_loop
_ = x_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply4(pkg_Data_Array.Get_foldM(), dictMonad_0, gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), gopurs_runtime.Any(acc_2), func() gopurs_runtime.Value {
					arr := x_3
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))
}

func Call_foldRecM(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
return gopurs_runtime.Apply(pkg_Data_Array.Get_foldRecM(), dictMonadRec_0)
}

func Call_index(x_0_loop []interface{}) gopurs_runtime.Value {
var x_0 []interface{} = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_Array.Get_index(), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
}

func Call_length(x_0_loop []interface{}) int64 {
var x_0 []interface{} = x_0_loop
_ = x_0
return gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))).IntVal
}

func Call_length__func_arrinterface____int64_4085341461(x_0_loop []interface{}) int64 {
var x_0 []interface{} = x_0_loop
_ = x_0
return gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))).IntVal
}

func Call_mapMaybe(f_0_loop func(interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}], x_1_loop []interface{}) []interface{} {
var f_0 func(interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] = f_0_loop
_ = f_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(pkg_Data_Array.Get_mapMaybe(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(f_0(gopurs_runtime.UnboxAny(arg0)))}
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_notElem(dictEq_0_loop gopurs_runtime.Value, x_1_loop interface{}, x_2_loop []interface{}) bool {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return (gopurs_runtime.Apply3(pkg_Data_Array.Get_notElem(), dictEq_0, gopurs_runtime.Any(x_1), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0)
}

func Call_partition(f_0_loop func(interface{}) bool, x_1_loop []interface{}) interface{} {
var f_0 func(interface{}) bool = f_0_loop
_ = f_0
var x_1 []interface{} = x_1_loop
_ = x_1
return gopurs_runtime.UnboxAny(gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_partitionImpl(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(f_0(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))
}

func Call_slice(start_0_loop int64, end_1_loop int64, x_2_loop []interface{}) []interface{} {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_sliceImpl(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_span(f_0_loop func(interface{}) bool, x_1_loop []interface{}) interface{} {
var f_0 func(interface{}) bool = f_0_loop
_ = f_0
var x_1 []interface{} = x_1_loop
_ = x_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(pkg_Data_Array.Get_span(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(f_0(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))
}

func Call_take(i_0_loop int64, x_1_loop []interface{}) []interface{} {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []interface{} = x_1_loop
_ = x_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(pkg_Data_Array.Get_lessThan(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(1)).IntVal) != (0) {
__t0 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(i_0), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())
}
end_branch_0:
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(__t0.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_takeEnd(i_0_loop int64, x_1_loop []interface{}) []interface{} {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(pkg_Data_Array.Get_takeEnd(), gopurs_runtime.Int(i_0), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_takeWhile(f_0_loop func(interface{}) bool, x_1_loop []interface{}) []interface{} {
var f_0 func(interface{}) bool = f_0_loop
_ = f_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Array.Get_span(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(f_0(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()), "init").UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_toUnfoldable(dictUnfoldable_0_loop gopurs_runtime.Value, x_1_loop []interface{}) interface{} {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var x_1 []interface{} = x_1_loop
_ = x_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(pkg_Data_Array.Get_toUnfoldable(), dictUnfoldable_0, func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))
}

func Call_cons(x_0_loop interface{}, x_1_loop []interface{}) []interface{} {
var x_0 interface{} = x_0_loop
_ = x_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), func() gopurs_runtime.Value {
					arr := []interface{}{x_0}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_group(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
eq2_1_0 := gopurs_runtime.RecordGet(dictEq_0, "eq")
_ = eq2_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Array.Get_groupBy(), eq2_1_0, x_2)
})
}

func Call_groupAllBy(op_0_loop func(interface{}, interface{}) gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 func(interface{}, interface{}) gopurs_runtime.Value = op_0_loop
_ = op_0
return gopurs_runtime.Apply(pkg_Data_Array.Get_groupAllBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return op_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1))
}))
}

func Call_groupAll(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Array.Get_groupAllBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
}

func Call_groupBy(op_0_loop func(interface{}, interface{}) bool, x_1_loop []interface{}) []gopurs_runtime.Value {
var op_0 func(interface{}, interface{}) bool = op_0_loop
_ = op_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(pkg_Data_Array.Get_groupBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(op_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_insert(dictOrd_0_loop gopurs_runtime.Value, x_1_loop interface{}, x_2_loop []interface{}) []interface{} {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Array.Get_insertBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Any(x_1), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_insertBy(f_0_loop func(interface{}, interface{}) gopurs_runtime.Value, x_1_loop interface{}, x_2_loop []interface{}) []interface{} {
var f_0 func(interface{}, interface{}) gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Array.Get_insertBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return f_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1))
}), gopurs_runtime.Any(x_1), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_intersperse(x_0_loop interface{}, x_1_loop []interface{}) []interface{} {
var x_0 interface{} = x_0_loop
_ = x_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(pkg_Data_Array.Get_intersperse(), gopurs_runtime.Any(x_0), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_mapWithIndex(f_0_loop func(int64, interface{}) interface{}) gopurs_runtime.Value {
var f_0 func(int64, interface{}) interface{} = f_0_loop
_ = f_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex"), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(arg0.IntVal, gopurs_runtime.UnboxAny(arg1)))
}))
}

func Call_modifyAtIndices(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply(pkg_Data_Array.Get_modifyAtIndices(), dictFoldable_0)
}

func Call_nub(dictOrd_0_loop gopurs_runtime.Value, x_1_loop []interface{}) []interface{} {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(pkg_Data_Array.Get_nubBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_nubBy(f_0_loop func(interface{}, interface{}) gopurs_runtime.Value, x_1_loop []interface{}) []interface{} {
var f_0 func(interface{}, interface{}) gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(pkg_Data_Array.Get_nubBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return f_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_nubByEq(f_0_loop func(interface{}, interface{}) bool, x_1_loop []interface{}) []interface{} {
var f_0 func(interface{}, interface{}) bool = f_0_loop
_ = f_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(pkg_Data_Array.Get_nubByEq(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(f_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_nubEq(dictEq_0_loop gopurs_runtime.Value, x_1_loop []interface{}) []interface{} {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(pkg_Data_Array.Get_nubByEq(), gopurs_runtime.RecordGet(dictEq_0, "eq"), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_reverse(x_0_loop []interface{}) []interface{} {
var x_0 []interface{} = x_0_loop
_ = x_0
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Data_Array.Get_reverse(), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_scanl(f_0_loop func(interface{}, interface{}) interface{}, x_1_loop interface{}, x_2_loop []interface{}) []interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_scanlImpl(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), gopurs_runtime.Any(x_1), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_scanr(f_0_loop func(interface{}, interface{}) interface{}, x_1_loop interface{}, x_2_loop []interface{}) []interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var x_1 interface{} = x_1_loop
_ = x_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(pkg_Data_Array.Get_scanrImpl(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1)))
}), gopurs_runtime.Any(x_1), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_sort(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Array.Get_sortBy(), compare_1_0, x_2)
})
}

func Call_sortBy(f_0_loop func(interface{}, interface{}) gopurs_runtime.Value, x_1_loop []interface{}) []interface{} {
var f_0 func(interface{}, interface{}) gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []interface{} = x_1_loop
_ = x_1
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(pkg_Data_Array.Get_sortBy(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return f_0(gopurs_runtime.UnboxAny(arg0), gopurs_runtime.UnboxAny(arg1))
}), func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_sortWith(dictOrd_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}, x_2_loop []interface{}) []interface{} {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var x_2 []interface{} = x_2_loop
_ = x_2
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(pkg_Data_Array.Get_sortWith(), dictOrd_0, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(arg0)))
}), func() gopurs_runtime.Value {
					arr := x_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}

func Call_updateAtIndices(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply(pkg_Data_Array.Get_updateAtIndices(), dictFoldable_0)
}

func Call_unsafeIndex(_dollar__unused_0_loop gopurs_runtime.Value, x_1_loop []interface{}, __local_var_2_loop int64) interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var x_1 []interface{} = x_1_loop
_ = x_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.UnboxAny(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), int(__local_var_2)))
}

func Call_unsafeIndex__func_gopurs_runtime_Value__arrinterface____int64__interface___3957585020(_dollar__unused_0_loop gopurs_runtime.Value, x_1_loop []interface{}, __local_var_2_loop int64) interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var x_1 []interface{} = x_1_loop
_ = x_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.UnboxAny(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
					arr := x_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), int(__local_var_2)))
}

func Call_toUnfoldable1(dictUnfoldable1_0_loop gopurs_runtime.Value, xs_1_loop []interface{}) interface{} {
var dictUnfoldable1_0 gopurs_runtime.Value = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var xs_1 []interface{} = xs_1_loop
_ = xs_1
len_2_0 := gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}())))
_ = len_2_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_0, "unfoldr1"), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), i_3, gopurs_runtime.Int((len_2_0.IntVal) - (1))).IntVal) != (0) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Int((i_3.IntVal) + (1)))})})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_1:
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
					arr := xs_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}(), int(i_3.IntVal))), gopurs_runtime.UnboxAny(__t1)})})
}), gopurs_runtime.Int(0)))
}
