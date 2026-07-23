package Data_Op

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var Op gopurs_runtime.Value
var once_Op sync.Once
func Get_Op() gopurs_runtime.Value {
	once_Op.Do(func() {
		Op = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Op
}

var semigroupoidOp gopurs_runtime.Value
var once_semigroupoidOp sync.Once
func Get_semigroupoidOp() gopurs_runtime.Value {
	once_semigroupoidOp.Do(func() {
		semigroupoidOp = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_1, gopurs_runtime.Apply(v_0, x_2))
}))
	})
	return semigroupoidOp
}

var semigroupOp gopurs_runtime.Value
var once_semigroupOp sync.Once
func Get_semigroupOp() gopurs_runtime.Value {
	once_semigroupOp.Do(func() {
		semigroupOp = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
}))
})
	})
	return semigroupOp
}

var newtypeOp gopurs_runtime.Value
var once_newtypeOp sync.Once
func Get_newtypeOp() gopurs_runtime.Value {
	once_newtypeOp.Do(func() {
		newtypeOp = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeOp
}

var monoidOp gopurs_runtime.Value
var once_monoidOp sync.Once
func Get_monoidOp() gopurs_runtime.Value {
	once_monoidOp.Do(func() {
		monoidOp = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty1_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupFn_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), gopurs_runtime.Apply(f_3, x_5), gopurs_runtime.Apply(g_4, x_5))
}))
_ = semigroupFn_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_3_2
}))
})
	})
	return monoidOp
}

var contravariantOp gopurs_runtime.Value
var once_contravariantOp sync.Once
func Get_contravariantOp() gopurs_runtime.Value {
	once_contravariantOp.Do(func() {
		contravariantOp = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
}))
	})
	return contravariantOp
}

var categoryOp gopurs_runtime.Value
var once_categoryOp sync.Once
func Get_categoryOp() gopurs_runtime.Value {
	once_categoryOp.Do(func() {
		categoryOp = gopurs_runtime.RecordDict2("identity", "Semigroupoid0", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupoidOp()
}))
	})
	return categoryOp
}


