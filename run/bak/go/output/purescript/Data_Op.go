package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Op_Op gopurs_runtime.Value
var once_Data_Op_Op sync.Once
func Get_Data_Op_Op() gopurs_runtime.Value {
	once_Data_Op_Op.Do(func() {
		cache_Data_Op_Op = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Op_Op(x_0_box)
})
	})
	return cache_Data_Op_Op
}

var cache_Data_Op_semigroupoidOp gopurs_runtime.Value
var once_Data_Op_semigroupoidOp sync.Once
func Get_Data_Op_semigroupoidOp() gopurs_runtime.Value {
	once_Data_Op_semigroupoidOp.Do(func() {
		cache_Data_Op_semigroupoidOp = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_1, gopurs_runtime.Apply(v_0, x_2))
})
})
}))
	})
	return cache_Data_Op_semigroupoidOp
}

var cache_Data_Op_semigroupOp gopurs_runtime.Value
var once_Data_Op_semigroupOp sync.Once
func Get_Data_Op_semigroupOp() gopurs_runtime.Value {
	once_Data_Op_semigroupOp.Do(func() {
		cache_Data_Op_semigroupOp = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Op_semigroupOp(dictSemigroup_0_box)
})
	})
	return cache_Data_Op_semigroupOp
}

var cache_Data_Op_newtypeOp gopurs_runtime.Value
var once_Data_Op_newtypeOp sync.Once
func Get_Data_Op_newtypeOp() gopurs_runtime.Value {
	once_Data_Op_newtypeOp.Do(func() {
		cache_Data_Op_newtypeOp = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Op_newtypeOp
}

var cache_Data_Op_monoidOp gopurs_runtime.Value
var once_Data_Op_monoidOp sync.Once
func Get_Data_Op_monoidOp() gopurs_runtime.Value {
	once_Data_Op_monoidOp.Do(func() {
		cache_Data_Op_monoidOp = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Op_monoidOp(dictMonoid_0_box)
})
	})
	return cache_Data_Op_monoidOp
}

var cache_Data_Op_contravariantOp gopurs_runtime.Value
var once_Data_Op_contravariantOp sync.Once
func Get_Data_Op_contravariantOp() gopurs_runtime.Value {
	once_Data_Op_contravariantOp.Do(func() {
		cache_Data_Op_contravariantOp = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
}))
	})
	return cache_Data_Op_contravariantOp
}

var cache_Data_Op_categoryOp gopurs_runtime.Value
var once_Data_Op_categoryOp sync.Once
func Get_Data_Op_categoryOp() gopurs_runtime.Value {
	once_Data_Op_categoryOp.Do(func() {
		cache_Data_Op_categoryOp = gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_2, gopurs_runtime.Apply(v_1, x_3))
})
})
}))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Op_categoryOp
}

var cache_Data_Op_contravariantOp__1861723411 gopurs_runtime.Value
var once_Data_Op_contravariantOp__1861723411 sync.Once
func Get_Data_Op_contravariantOp__1861723411() gopurs_runtime.Value {
	once_Data_Op_contravariantOp__1861723411.Do(func() {
		cache_Data_Op_contravariantOp__1861723411 = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
}))
	})
	return cache_Data_Op_contravariantOp__1861723411
}

var cache_Data_Op_semigroupoidOp__3459988149 gopurs_runtime.Value
var once_Data_Op_semigroupoidOp__3459988149 sync.Once
func Get_Data_Op_semigroupoidOp__3459988149() gopurs_runtime.Value {
	once_Data_Op_semigroupoidOp__3459988149.Do(func() {
		cache_Data_Op_semigroupoidOp__3459988149 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_1, gopurs_runtime.Apply(v_0, x_2))
})
})
}))
	})
	return cache_Data_Op_semigroupoidOp__3459988149
}

func Call_Data_Op_Op(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Op_semigroupOp(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Op_monoidOp(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupFn_1_0 -> gopurs_runtime.Value
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


