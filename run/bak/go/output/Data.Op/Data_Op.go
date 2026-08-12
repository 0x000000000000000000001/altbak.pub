package Data_Op

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Op gopurs_runtime.Value
var once_Op sync.Once
func Get_Op() gopurs_runtime.Value {
	once_Op.Do(func() {
		cache_Op = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Op(x_0_box)
})
	})
	return cache_Op
}

var cache_semigroupoidOp gopurs_runtime.Value
var once_semigroupoidOp sync.Once
func Get_semigroupoidOp() gopurs_runtime.Value {
	once_semigroupoidOp.Do(func() {
		cache_semigroupoidOp = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_1, gopurs_runtime.Apply(v_0, x_2))
})
})
}))
	})
	return cache_semigroupoidOp
}

var cache_semigroupOp gopurs_runtime.Value
var once_semigroupOp sync.Once
func Get_semigroupOp() gopurs_runtime.Value {
	once_semigroupOp.Do(func() {
		cache_semigroupOp = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupOp(dictSemigroup_0_box)
})
	})
	return cache_semigroupOp
}

var cache_newtypeOp gopurs_runtime.Value
var once_newtypeOp sync.Once
func Get_newtypeOp() gopurs_runtime.Value {
	once_newtypeOp.Do(func() {
		cache_newtypeOp = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeOp
}

var cache_monoidOp gopurs_runtime.Value
var once_monoidOp sync.Once
func Get_monoidOp() gopurs_runtime.Value {
	once_monoidOp.Do(func() {
		cache_monoidOp = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidOp(dictMonoid_0_box)
})
	})
	return cache_monoidOp
}

var cache_contravariantOp gopurs_runtime.Value
var once_contravariantOp sync.Once
func Get_contravariantOp() gopurs_runtime.Value {
	once_contravariantOp.Do(func() {
		cache_contravariantOp = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
}))
	})
	return cache_contravariantOp
}

var cache_categoryOp gopurs_runtime.Value
var once_categoryOp sync.Once
func Get_categoryOp() gopurs_runtime.Value {
	once_categoryOp.Do(func() {
		cache_categoryOp = gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupoidOp()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_categoryOp
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_semigroupoidOp__3459988149 gopurs_runtime.Value
var once_semigroupoidOp__3459988149 sync.Once
func Get_semigroupoidOp__3459988149() gopurs_runtime.Value {
	once_semigroupoidOp__3459988149.Do(func() {
		cache_semigroupoidOp__3459988149 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_1, gopurs_runtime.Apply(v_0, x_2))
})
})
}))
	})
	return cache_semigroupoidOp__3459988149
}

func Call_Op(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_semigroupOp(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
})
})
}))
}

func Call_monoidOp(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_1
semigroupFn_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "append"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})
})
}))
_ = semigroupFn_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
}))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


