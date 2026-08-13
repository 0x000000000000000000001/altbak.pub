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
		cache_Data_Comparison_semigroupFn = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(f_0, x_2)
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(g_1, x_2)
_ = __local_var_4_2
var __t3 uint32
{
if (uint32(__local_var_3_1.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(__local_var_3_1.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
if (uint32(__local_var_3_1.IntVal) == 902936544) {
__t3 = uint32(__local_var_4_2.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
})
}))
_ = __local_var_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
})
})
}))))}
}()
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
		cache_Data_Comparison_semigroupComparison = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Apply(f_2, x_4)
_ = __local_var_5_1
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(g_3, x_4)
_ = __local_var_6_2
var __t3 uint32
{
if (uint32(__local_var_5_1.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(__local_var_5_1.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
if (uint32(__local_var_5_1.IntVal) == 902936544) {
__t3 = uint32(__local_var_6_2.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
})
}))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), gopurs_runtime.Apply(v_0, x_3), gopurs_runtime.Apply(v1_1, x_3))
})
})
}))
	})
	return cache_Data_Comparison_semigroupComparison
}

var cache_Data_Comparison_newtypeComparison gopurs_runtime.Value
var once_Data_Comparison_newtypeComparison sync.Once
func Get_Data_Comparison_newtypeComparison() gopurs_runtime.Value {
	once_Data_Comparison_newtypeComparison.Do(func() {
		cache_Data_Comparison_newtypeComparison = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Comparison_newtypeComparison
}

var cache_Data_Comparison_monoidComparison gopurs_runtime.Value
var once_Data_Comparison_monoidComparison sync.Once
func Get_Data_Comparison_monoidComparison() gopurs_runtime.Value {
	once_Data_Comparison_monoidComparison.Do(func() {
		cache_Data_Comparison_monoidComparison = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Comparison_semigroupComparison()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Comparison_monoidComparison
}

var cache_Data_Comparison_defaultComparison gopurs_runtime.Value
var once_Data_Comparison_defaultComparison sync.Once
func Get_Data_Comparison_defaultComparison() gopurs_runtime.Value {
	once_Data_Comparison_defaultComparison.Do(func() {
		cache_Data_Comparison_defaultComparison = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Comparison_defaultComparison(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Comparison_defaultComparison
}

var cache_Data_Comparison_contravariantComparison gopurs_runtime.Value
var once_Data_Comparison_contravariantComparison sync.Once
func Get_Data_Comparison_contravariantComparison() gopurs_runtime.Value {
	once_Data_Comparison_contravariantComparison.Do(func() {
		cache_Data_Comparison_contravariantComparison = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
})
})
})
}))
	})
	return cache_Data_Comparison_contravariantComparison
}

var cache_Data_Comparison_contravariantComparison__1065380147 gopurs_runtime.Value
var once_Data_Comparison_contravariantComparison__1065380147 sync.Once
func Get_Data_Comparison_contravariantComparison__1065380147() gopurs_runtime.Value {
	once_Data_Comparison_contravariantComparison__1065380147.Do(func() {
		cache_Data_Comparison_contravariantComparison__1065380147 = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
})
})
})
}))
	})
	return cache_Data_Comparison_contravariantComparison__1065380147
}

var cache_Data_Comparison_semigroupComparison__1133613061 gopurs_runtime.Value
var once_Data_Comparison_semigroupComparison__1133613061 sync.Once
func Get_Data_Comparison_semigroupComparison__1133613061() gopurs_runtime.Value {
	once_Data_Comparison_semigroupComparison__1133613061.Do(func() {
		cache_Data_Comparison_semigroupComparison__1133613061 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Apply(f_2, x_4)
_ = __local_var_5_1
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(g_3, x_4)
_ = __local_var_6_2
var __t3 uint32
{
if (uint32(__local_var_5_1.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(__local_var_5_1.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
if (uint32(__local_var_5_1.IntVal) == 902936544) {
__t3 = uint32(__local_var_6_2.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
})
}))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), gopurs_runtime.Apply(v_0, x_3), gopurs_runtime.Apply(v1_1, x_3))
})
})
}))
	})
	return cache_Data_Comparison_semigroupComparison__1133613061
}

func Call_Data_Comparison_Comparison(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Comparison_defaultComparison(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Box(dictOrd_0.V1)
}


