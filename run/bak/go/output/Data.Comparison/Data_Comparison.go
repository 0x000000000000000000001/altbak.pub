package Data_Comparison

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	unsafe "unsafe"
)

var cache_Comparison gopurs_runtime.Value
var once_Comparison sync.Once
func Get_Comparison() gopurs_runtime.Value {
	once_Comparison.Do(func() {
		cache_Comparison = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Comparison
}

var cache_semigroupComparison gopurs_runtime.Value
var once_semigroupComparison sync.Once
func Get_semigroupComparison() gopurs_runtime.Value {
	once_semigroupComparison.Do(func() {
		cache_semigroupComparison = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(v_0, x_2)
_ = __local_var_3_0
__local_var_4_1 := gopurs_runtime.Apply(v1_1, x_2)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(__local_var_3_0, x_5)
_ = __local_var_6_2
__local_var_7_3 := gopurs_runtime.Apply(__local_var_4_1, x_5)
_ = __local_var_7_3
var __t4 gopurs_runtime.Value
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
goto end_branch_4
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_4
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 902936544) {
__t4 = __local_var_7_3
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
}))
	})
	return cache_semigroupComparison
}

var cache_newtypeComparison gopurs_runtime.Value
var once_newtypeComparison sync.Once
func Get_newtypeComparison() gopurs_runtime.Value {
	once_newtypeComparison.Do(func() {
		cache_newtypeComparison = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeComparison
}

var cache_monoidComparison gopurs_runtime.Value
var once_monoidComparison sync.Once
func Get_monoidComparison() gopurs_runtime.Value {
	once_monoidComparison.Do(func() {
		cache_monoidComparison = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupComparison()
}))
	})
	return cache_monoidComparison
}

var cache_defaultComparison gopurs_runtime.Value
var once_defaultComparison sync.Once
func Get_defaultComparison() gopurs_runtime.Value {
	once_defaultComparison.Do(func() {
		cache_defaultComparison = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}()
})
	})
	return cache_defaultComparison
}

var cache_contravariantComparison gopurs_runtime.Value
var once_contravariantComparison sync.Once
func Get_contravariantComparison() gopurs_runtime.Value {
	once_contravariantComparison.Do(func() {
		cache_contravariantComparison = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
}))
	})
	return cache_contravariantComparison
}




