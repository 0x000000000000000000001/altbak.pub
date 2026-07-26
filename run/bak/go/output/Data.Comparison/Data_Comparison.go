package Data_Comparison

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
)

var cache_append_ gopurs_runtime.Value
var once_append_ sync.Once
func Get_append_() gopurs_runtime.Value {
	once_append_.Do(func() {
		cache_append_ = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append_(f_0_box, g_1_box, x_2_box)
})
	})
	return cache_append_
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
		cache_semigroupComparison = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_append_(), v_0, v1_1)
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
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
}))
	})
	return cache_monoidComparison
}

var cache_defaultComparison gopurs_runtime.Value
var once_defaultComparison sync.Once
func Get_defaultComparison() gopurs_runtime.Value {
	once_defaultComparison.Do(func() {
		cache_defaultComparison = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defaultComparison(dictOrd_0_box)
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

func Call_append_(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
__local_var_3_0 := gopurs_runtime.Apply(f_0, x_2)
_ = __local_var_3_0
__local_var_4_1 := gopurs_runtime.Apply(g_1, x_2)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_semigroupOrdering(), "append"), gopurs_runtime.Apply(__local_var_3_0, x_5), gopurs_runtime.Apply(__local_var_4_1, x_5))
})
}

func Call_Comparison(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_defaultComparison(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}


