package Data_Monoid_Endo

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Endo gopurs_runtime.Value
var once_Endo sync.Once
func Get_Endo() gopurs_runtime.Value {
	once_Endo.Do(func() {
		Endo = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Endo
}

var showEndo gopurs_runtime.Value
var once_showEndo sync.Once
func Get_showEndo() gopurs_runtime.Value {
	once_showEndo.Do(func() {
		showEndo = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Endo " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal + ")")
}))
})
	})
	return showEndo
}

var semigroupEndo gopurs_runtime.Value
var once_semigroupEndo sync.Once
func Get_semigroupEndo() gopurs_runtime.Value {
	once_semigroupEndo.Do(func() {
		semigroupEndo = gopurs_runtime.Func(func(dictSemigroupoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), v_1, v1_2)
}))
})
	})
	return semigroupEndo
}

var ordEndo gopurs_runtime.Value
var once_ordEndo sync.Once
func Get_ordEndo() gopurs_runtime.Value {
	once_ordEndo.Do(func() {
		ordEndo = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordEndo
}

var monoidEndo gopurs_runtime.Value
var once_monoidEndo sync.Once
func Get_monoidEndo() gopurs_runtime.Value {
	once_monoidEndo.Do(func() {
		monoidEndo = gopurs_runtime.Func(func(dictCategory_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCategory_0, "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupEndo1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "compose"), v_2, v1_3)
}))
_ = semigroupEndo1_2_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictCategory_0, "identity"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_2_1
}))
})
	})
	return monoidEndo
}

var eqEndo gopurs_runtime.Value
var once_eqEndo sync.Once
func Get_eqEndo() gopurs_runtime.Value {
	once_eqEndo.Do(func() {
		eqEndo = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqEndo
}

var boundedEndo gopurs_runtime.Value
var once_boundedEndo sync.Once
func Get_boundedEndo() gopurs_runtime.Value {
	once_boundedEndo.Do(func() {
		boundedEndo = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedEndo
}




