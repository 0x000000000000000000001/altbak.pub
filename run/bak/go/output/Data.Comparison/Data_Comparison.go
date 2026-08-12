package Data_Comparison

import (
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_semigroupFn gopurs_runtime.Value
var once_semigroupFn sync.Once
func Get_semigroupFn() gopurs_runtime.Value {
	once_semigroupFn.Do(func() {
		cache_semigroupFn = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(f_0, x_2)
_ = __local_var_3_0
__local_var_4_1 := gopurs_runtime.Apply(g_1, x_2)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_semigroupOrdering(), "append"), gopurs_runtime.Apply(__local_var_3_0, x_5), gopurs_runtime.Apply(__local_var_4_1, x_5))
})
})
})
})})}
	})
	return cache_semigroupFn
}

var cache_Comparison gopurs_runtime.Value
var once_Comparison sync.Once
func Get_Comparison() gopurs_runtime.Value {
	once_Comparison.Do(func() {
		cache_Comparison = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Comparison(x_0_box)
})
	})
	return cache_Comparison
}

var cache_semigroupComparison gopurs_runtime.Value
var once_semigroupComparison sync.Once
func Get_semigroupComparison() gopurs_runtime.Value {
	once_semigroupComparison.Do(func() {
		cache_semigroupComparison = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](Get_semigroupFn()).V0, v_0, v1_1)
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
		cache_monoidComparison = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupComparison()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
}))
	})
	return cache_monoidComparison
}

var cache_defaultComparison gopurs_runtime.Value
var once_defaultComparison sync.Once
func Get_defaultComparison() gopurs_runtime.Value {
	once_defaultComparison.Do(func() {
		cache_defaultComparison = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defaultComparison(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_defaultComparison
}

var cache_contravariantComparison gopurs_runtime.Value
var once_contravariantComparison sync.Once
func Get_contravariantComparison() gopurs_runtime.Value {
	once_contravariantComparison.Do(func() {
		cache_contravariantComparison = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
})
})
})
}))
	})
	return cache_contravariantComparison
}

var cache_semigroupComparison__1133613061 gopurs_runtime.Value
var once_semigroupComparison__1133613061 sync.Once
func Get_semigroupComparison__1133613061() gopurs_runtime.Value {
	once_semigroupComparison__1133613061.Do(func() {
		cache_semigroupComparison__1133613061 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](Get_semigroupFn()).V0, v_0, v1_1)
})
}))
	})
	return cache_semigroupComparison__1133613061
}

var cache_on__3980724833 gopurs_runtime.Value
var once_on__3980724833 sync.Once
func Get_on__3980724833() gopurs_runtime.Value {
	once_on__3980724833.Do(func() {
		cache_on__3980724833 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_on__3980724833(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_on__3980724833
}

var cache_on__3556844193 gopurs_runtime.Value
var once_on__3556844193 sync.Once
func Get_on__3556844193() gopurs_runtime.Value {
	once_on__3556844193.Do(func() {
		cache_on__3556844193 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_on__3556844193(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_on__3556844193
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_append__2611908184 gopurs_runtime.Value
var once_append__2611908184 sync.Once
func Get_append__2611908184() gopurs_runtime.Value {
	once_append__2611908184.Do(func() {
		cache_append__2611908184 = gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](Get_semigroupFn()).V0
	})
	return cache_append__2611908184
}

func Call_Comparison(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_defaultComparison(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0.V1
}

func Call_on__3980724833(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_on__3556844193(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


