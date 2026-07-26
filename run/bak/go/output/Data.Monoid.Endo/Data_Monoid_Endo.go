package Data_Monoid_Endo

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_Endo gopurs_runtime.Value
var once_Endo sync.Once
func Get_Endo() gopurs_runtime.Value {
	once_Endo.Do(func() {
		cache_Endo = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Endo(x_0_box)
})
	})
	return cache_Endo
}

var cache_showEndo gopurs_runtime.Value
var once_showEndo sync.Once
func Get_showEndo() gopurs_runtime.Value {
	once_showEndo.Do(func() {
		cache_showEndo = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showEndo(dictShow_0_box)
})
	})
	return cache_showEndo
}

var cache_semigroupEndo gopurs_runtime.Value
var once_semigroupEndo sync.Once
func Get_semigroupEndo() gopurs_runtime.Value {
	once_semigroupEndo.Do(func() {
		cache_semigroupEndo = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupEndo(dictSemigroupoid_0_box)
})
	})
	return cache_semigroupEndo
}

var cache_ordEndo gopurs_runtime.Value
var once_ordEndo sync.Once
func Get_ordEndo() gopurs_runtime.Value {
	once_ordEndo.Do(func() {
		cache_ordEndo = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordEndo(dictOrd_0_box)
})
	})
	return cache_ordEndo
}

var cache_monoidEndo gopurs_runtime.Value
var once_monoidEndo sync.Once
func Get_monoidEndo() gopurs_runtime.Value {
	once_monoidEndo.Do(func() {
		cache_monoidEndo = gopurs_runtime.Func(func(dictCategory_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidEndo(dictCategory_0_box)
})
	})
	return cache_monoidEndo
}

var cache_eqEndo gopurs_runtime.Value
var once_eqEndo sync.Once
func Get_eqEndo() gopurs_runtime.Value {
	once_eqEndo.Do(func() {
		cache_eqEndo = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqEndo(dictEq_0_box)
})
	})
	return cache_eqEndo
}

var cache_boundedEndo gopurs_runtime.Value
var once_boundedEndo sync.Once
func Get_boundedEndo() gopurs_runtime.Value {
	once_boundedEndo.Do(func() {
		cache_boundedEndo = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedEndo(dictBounded_0_box)
})
	})
	return cache_boundedEndo
}

func Call_Endo(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showEndo(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Endo "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow_0.UnsafePtr)).V0, v_1), gopurs_runtime.Str(")")))
}))
}

func Call_semigroupEndo(dictSemigroupoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictSemigroupoid_0.UnsafePtr)).V0, v_1, v1_2)
}))
}

func Call_ordEndo(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_monoidEndo(dictCategory_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCategory_0 gopurs_runtime.Value = dictCategory_0_loop
_ = dictCategory_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCategory_0, "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupEndo1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "compose"), v_2, v1_3)
}))
_ = semigroupEndo1_2_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_2_1
}), ((*gopurs_runtime.RecordData1)(dictCategory_0.UnsafePtr)).V0)
}

func Call_eqEndo(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_boundedEndo(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}


