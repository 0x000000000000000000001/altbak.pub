package Data_Op

import (
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
		cache_semigroupoidOp = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_1, gopurs_runtime.Apply(v_0, x_2))
}))))
	})
	return cache_semigroupoidOp
}

var cache_semigroupOp gopurs_runtime.Value
var once_semigroupOp sync.Once
func Get_semigroupOp() gopurs_runtime.Value {
	once_semigroupOp.Do(func() {
		cache_semigroupOp = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_semigroupOp(dictSemigroup_0_box))
})
	})
	return cache_semigroupOp
}

var cache_newtypeOp gopurs_runtime.Value
var once_newtypeOp sync.Once
func Get_newtypeOp() gopurs_runtime.Value {
	once_newtypeOp.Do(func() {
		cache_newtypeOp = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))))
	})
	return cache_newtypeOp
}

var cache_monoidOp gopurs_runtime.Value
var once_monoidOp sync.Once
func Get_monoidOp() gopurs_runtime.Value {
	once_monoidOp.Do(func() {
		cache_monoidOp = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monoidOp(dictMonoid_0_box))
})
	})
	return cache_monoidOp
}

var cache_contravariantOp gopurs_runtime.Value
var once_contravariantOp sync.Once
func Get_contravariantOp() gopurs_runtime.Value {
	once_contravariantOp.Do(func() {
		cache_contravariantOp = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
}))))
	})
	return cache_contravariantOp
}

var cache_categoryOp gopurs_runtime.Value
var once_categoryOp sync.Once
func Get_categoryOp() gopurs_runtime.Value {
	once_categoryOp.Do(func() {
		cache_categoryOp = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupoidOp()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))))
	})
	return cache_categoryOp
}

func Call_Op(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_semigroupOp(dictSemigroup_0_loop gopurs_runtime.Value) interface{} {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
})))
}

func Call_monoidOp(dictMonoid_0_loop gopurs_runtime.Value) interface{} {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty1_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupFn_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), gopurs_runtime.Apply(f_3, x_5), gopurs_runtime.Apply(g_4, x_5))
}))
_ = semigroupFn_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty1_1_0
})))
}
