package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Comparison_semigroupFn gopurs_runtime.Value
var once_Data_Comparison_semigroupFn sync.Once
func Get_Data_Comparison_semigroupFn() gopurs_runtime.Value {
	once_Data_Comparison_semigroupFn.Do(func() {
		cache_Data_Comparison_semigroupFn = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeVar s')
__local_var_3_0 := gopurs_runtime.Apply(f_0, x_2)
_ = __local_var_3_0
// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=(TypeVar s')
__local_var_4_1 := gopurs_runtime.Apply(g_1, x_2)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 shape=App(Other) bindingType=Any
__local_var_6_2 := gopurs_runtime.Apply(__local_var_3_0, x_5)
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 shape=App(Other) bindingType=(TypeVar s')
__local_var_7_3 := gopurs_runtime.Apply(__local_var_4_1, x_5)
_ = __local_var_7_3
var __t4 uint32
{
if (uint32(__local_var_6_2.IntVal) == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (uint32(__local_var_6_2.IntVal) == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if (uint32(__local_var_6_2.IntVal) == 902936544) {
__t4 = uint32(__local_var_7_3.IntVal)
goto end_branch_4
} else {

}
}
{
__t4 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
})
})}))}
	})
	return cache_Data_Comparison_semigroupFn
}

var cache_Data_Comparison_Comparison gopurs_runtime.Value
var once_Data_Comparison_Comparison sync.Once
func Get_Data_Comparison_Comparison() gopurs_runtime.Value {
	once_Data_Comparison_Comparison.Do(func() {
		cache_Data_Comparison_Comparison = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Comparison_Comparison(x_0_box)
})
	})
	return cache_Data_Comparison_Comparison
}

var cache_Data_Comparison_semigroupComparison gopurs_runtime.Value
var once_Data_Comparison_semigroupComparison sync.Once
func Get_Data_Comparison_semigroupComparison() gopurs_runtime.Value {
	once_Data_Comparison_semigroupComparison.Do(func() {
		cache_Data_Comparison_semigroupComparison = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeVar s')
__local_var_3_0 := gopurs_runtime.Apply(v_0, x_2)
_ = __local_var_3_0
// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=(TypeVar s')
__local_var_4_1 := gopurs_runtime.Apply(v1_1, x_2)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 shape=App(Other) bindingType=Any
__local_var_6_2 := gopurs_runtime.Apply(__local_var_3_0, x_5)
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 shape=App(Other) bindingType=(TypeVar s')
__local_var_7_3 := gopurs_runtime.Apply(__local_var_4_1, x_5)
_ = __local_var_7_3
var __t4 uint32
{
if (uint32(__local_var_6_2.IntVal) == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (uint32(__local_var_6_2.IntVal) == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if (uint32(__local_var_6_2.IntVal) == 902936544) {
__t4 = uint32(__local_var_7_3.IntVal)
goto end_branch_4
} else {

}
}
{
__t4 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
})
})}))}
	})
	return cache_Data_Comparison_semigroupComparison
}

var cache_Data_Comparison_newtypeComparison gopurs_runtime.Value
var once_Data_Comparison_newtypeComparison sync.Once
func Get_Data_Comparison_newtypeComparison() gopurs_runtime.Value {
	once_Data_Comparison_newtypeComparison.Do(func() {
		cache_Data_Comparison_newtypeComparison = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
	})
	return cache_Data_Comparison_newtypeComparison
}

var cache_Data_Comparison_monoidComparison gopurs_runtime.Value
var once_Data_Comparison_monoidComparison sync.Once
func Get_Data_Comparison_monoidComparison() gopurs_runtime.Value {
	once_Data_Comparison_monoidComparison.Do(func() {
		cache_Data_Comparison_monoidComparison = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Comparison_semigroupComparison()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})}))}
	})
	return cache_Data_Comparison_monoidComparison
}

var cache_Data_Comparison_defaultComparison gopurs_runtime.Value
var once_Data_Comparison_defaultComparison sync.Once
func Get_Data_Comparison_defaultComparison() gopurs_runtime.Value {
	once_Data_Comparison_defaultComparison.Do(func() {
		cache_Data_Comparison_defaultComparison = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Comparison_defaultComparison(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Comparison_defaultComparison
}

var cache_Data_Comparison_contravariantComparison gopurs_runtime.Value
var once_Data_Comparison_contravariantComparison sync.Once
func Get_Data_Comparison_contravariantComparison() gopurs_runtime.Value {
	once_Data_Comparison_contravariantComparison.Do(func() {
		cache_Data_Comparison_contravariantComparison = gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3)).IntVal)), UnsafePtr: nil}
})}))}
	})
	return cache_Data_Comparison_contravariantComparison
}

func Call_Data_Comparison_Comparison(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Comparison_defaultComparison(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Box(dictOrd_0.V1)
}


