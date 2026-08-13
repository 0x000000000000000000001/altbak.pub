package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Monoid_Endo_Endo gopurs_runtime.Value
var once_Data_Monoid_Endo_Endo sync.Once
func Get_Data_Monoid_Endo_Endo() gopurs_runtime.Value {
	once_Data_Monoid_Endo_Endo.Do(func() {
		cache_Data_Monoid_Endo_Endo = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Endo_Endo(x_0_box)
})
	})
	return cache_Data_Monoid_Endo_Endo
}

var cache_Data_Monoid_Endo_showEndo gopurs_runtime.Value
var once_Data_Monoid_Endo_showEndo sync.Once
func Get_Data_Monoid_Endo_showEndo() gopurs_runtime.Value {
	once_Data_Monoid_Endo_showEndo.Do(func() {
		cache_Data_Monoid_Endo_showEndo = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Endo_showEndo(dictShow_0_box)
})
	})
	return cache_Data_Monoid_Endo_showEndo
}

var cache_Data_Monoid_Endo_semigroupEndo gopurs_runtime.Value
var once_Data_Monoid_Endo_semigroupEndo sync.Once
func Get_Data_Monoid_Endo_semigroupEndo() gopurs_runtime.Value {
	once_Data_Monoid_Endo_semigroupEndo.Do(func() {
		cache_Data_Monoid_Endo_semigroupEndo = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Endo_semigroupEndo(dictSemigroupoid_0_box)
})
	})
	return cache_Data_Monoid_Endo_semigroupEndo
}

var cache_Data_Monoid_Endo_ordEndo gopurs_runtime.Value
var once_Data_Monoid_Endo_ordEndo sync.Once
func Get_Data_Monoid_Endo_ordEndo() gopurs_runtime.Value {
	once_Data_Monoid_Endo_ordEndo.Do(func() {
		cache_Data_Monoid_Endo_ordEndo = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Endo_ordEndo(dictOrd_0_box)
})
	})
	return cache_Data_Monoid_Endo_ordEndo
}

var cache_Data_Monoid_Endo_monoidEndo gopurs_runtime.Value
var once_Data_Monoid_Endo_monoidEndo sync.Once
func Get_Data_Monoid_Endo_monoidEndo() gopurs_runtime.Value {
	once_Data_Monoid_Endo_monoidEndo.Do(func() {
		cache_Data_Monoid_Endo_monoidEndo = gopurs_runtime.Func(func(dictCategory_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Endo_monoidEndo(dictCategory_0_box)
})
	})
	return cache_Data_Monoid_Endo_monoidEndo
}

var cache_Data_Monoid_Endo_eqEndo gopurs_runtime.Value
var once_Data_Monoid_Endo_eqEndo sync.Once
func Get_Data_Monoid_Endo_eqEndo() gopurs_runtime.Value {
	once_Data_Monoid_Endo_eqEndo.Do(func() {
		cache_Data_Monoid_Endo_eqEndo = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Endo_eqEndo(dictEq_0_box)
})
	})
	return cache_Data_Monoid_Endo_eqEndo
}

var cache_Data_Monoid_Endo_boundedEndo gopurs_runtime.Value
var once_Data_Monoid_Endo_boundedEndo sync.Once
func Get_Data_Monoid_Endo_boundedEndo() gopurs_runtime.Value {
	once_Data_Monoid_Endo_boundedEndo.Do(func() {
		cache_Data_Monoid_Endo_boundedEndo = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Endo_boundedEndo(dictBounded_0_box)
})
	})
	return cache_Data_Monoid_Endo_boundedEndo
}

func Call_Data_Monoid_Endo_Endo(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Monoid_Endo_showEndo(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Endo ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
})})}
}

func Call_Data_Monoid_Endo_semigroupEndo(dictSemigroupoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), v_1, v1_2)
})
})})}
}

func Call_Data_Monoid_Endo_ordEndo(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0))}
}

func Call_Data_Monoid_Endo_monoidEndo(dictCategory_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCategory_0 gopurs_runtime.Value = dictCategory_0_loop
_ = dictCategory_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCategory_0, "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupEndo1_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupEndo1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "compose"), v_2, v1_3)
})
})))
_ = semigroupEndo1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupEndo1_1_0)}
}), gopurs_runtime.RecordGet(dictCategory_0, "identity")})}
}

func Call_Data_Monoid_Endo_eqEndo(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0))}
}

func Call_Data_Monoid_Endo_boundedEndo(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](dictBounded_0))}
}


