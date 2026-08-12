package Control_Monad_ST_Internal

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_new gopurs_runtime.Value
var once_new sync.Once
func Get_new() gopurs_runtime.Value {
	once_new.Do(func() {
		cache_new = Get_newImpl()
	})
	return cache_new
}

var cache_new__gopurs_runtime_Value_3489595018 gopurs_runtime.Value
var once_new__gopurs_runtime_Value_3489595018 sync.Once
func Get_new__gopurs_runtime_Value_3489595018() gopurs_runtime.Value {
	once_new__gopurs_runtime_Value_3489595018.Do(func() {
		cache_new__gopurs_runtime_Value_3489595018 = Get_newImpl()
	})
	return cache_new__gopurs_runtime_Value_3489595018
}

var cache_new__gopurs_runtime_Value_2010968700 gopurs_runtime.Value
var once_new__gopurs_runtime_Value_2010968700 sync.Once
func Get_new__gopurs_runtime_Value_2010968700() gopurs_runtime.Value {
	once_new__gopurs_runtime_Value_2010968700.Do(func() {
		cache_new__gopurs_runtime_Value_2010968700 = Get_newImpl()
	})
	return cache_new__gopurs_runtime_Value_2010968700
}

var cache_modify_prime gopurs_runtime.Value
var once_modify_prime sync.Once
func Get_modify_prime() gopurs_runtime.Value {
	once_modify_prime.Do(func() {
		cache_modify_prime = Get_modifyImpl()
	})
	return cache_modify_prime
}

var cache_modify_prime__gopurs_runtime_Value_1497736571 gopurs_runtime.Value
var once_modify_prime__gopurs_runtime_Value_1497736571 sync.Once
func Get_modify_prime__gopurs_runtime_Value_1497736571() gopurs_runtime.Value {
	once_modify_prime__gopurs_runtime_Value_1497736571.Do(func() {
		cache_modify_prime__gopurs_runtime_Value_1497736571 = Get_modifyImpl()
	})
	return cache_modify_prime__gopurs_runtime_Value_1497736571
}

var cache_modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		cache_modify = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify(f_0_box)
})
	})
	return cache_modify
}

var cache_modify__gopurs_runtime_Value_781734141 gopurs_runtime.Value
var once_modify__gopurs_runtime_Value_781734141 sync.Once
func Get_modify__gopurs_runtime_Value_781734141() gopurs_runtime.Value {
	once_modify__gopurs_runtime_Value_781734141.Do(func() {
		cache_modify__gopurs_runtime_Value_781734141 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__gopurs_runtime_Value_781734141(f_0_box)
})
	})
	return cache_modify__gopurs_runtime_Value_781734141
}

var cache_functorST gopurs_runtime.Value
var once_functorST sync.Once
func Get_functorST() gopurs_runtime.Value {
	once_functorST.Do(func() {
		cache_functorST = gopurs_runtime.RecordDict1("map", Get_map_())
	})
	return cache_functorST
}

var cache_functorST__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__4062753802 gopurs_runtime.Value
var once_functorST__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__4062753802 sync.Once
func Get_functorST__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__4062753802() gopurs_runtime.Value {
	once_functorST__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__4062753802.Do(func() {
		cache_functorST__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__4062753802 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]{1, Get_map_()})}
	})
	return cache_functorST__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__4062753802
}

var cache_functorST__gopurs_runtime_Value_2441840241 gopurs_runtime.Value
var once_functorST__gopurs_runtime_Value_2441840241 sync.Once
func Get_functorST__gopurs_runtime_Value_2441840241() gopurs_runtime.Value {
	once_functorST__gopurs_runtime_Value_2441840241.Do(func() {
		cache_functorST__gopurs_runtime_Value_2441840241 = gopurs_runtime.RecordDict1("map", Get_map_())
	})
	return cache_functorST__gopurs_runtime_Value_2441840241
}

var cache_go__for gopurs_runtime.Value
var once_go__for sync.Once
func Get_go__for() gopurs_runtime.Value {
	once_go__for.Do(func() {
		cache_go__for = Get_forImpl()
	})
	return cache_go__for
}

var cache_for__gopurs_runtime_Value_956375728 gopurs_runtime.Value
var once_for__gopurs_runtime_Value_956375728 sync.Once
func Get_for__gopurs_runtime_Value_956375728() gopurs_runtime.Value {
	once_for__gopurs_runtime_Value_956375728.Do(func() {
		cache_for__gopurs_runtime_Value_956375728 = Get_forImpl()
	})
	return cache_for__gopurs_runtime_Value_956375728
}

var cache_monadST_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Any gopurs_runtime.Value
var once_monadST_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Any sync.Once
func Get_monadST_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Any() gopurs_runtime.Value {
	once_monadST_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Any.Do(func() {
		cache_monadST_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Any = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindST()
}))
	})
	return cache_monadST_Record_Row_Applicative0_Func_Record_Row__Any_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Bind1_Func_Record_Row__Any_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any_Any
}

var cache_monadST gopurs_runtime.Value
var once_monadST sync.Once
func Get_monadST() gopurs_runtime.Value {
	once_monadST.Do(func() {
		cache_monadST = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindST()
}))
	})
	return cache_monadST
}

var cache_monadST__gopurs_runtime_Value_1413783571 gopurs_runtime.Value
var once_monadST__gopurs_runtime_Value_1413783571 sync.Once
func Get_monadST__gopurs_runtime_Value_1413783571() gopurs_runtime.Value {
	once_monadST__gopurs_runtime_Value_1413783571.Do(func() {
		cache_monadST__gopurs_runtime_Value_1413783571 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindST()
}))
	})
	return cache_monadST__gopurs_runtime_Value_1413783571
}

var cache_bindST_ADT_Control_Bind_Bind_ADT_Control_Monad_ST_Internal_ST_Any gopurs_runtime.Value
var once_bindST_ADT_Control_Bind_Bind_ADT_Control_Monad_ST_Internal_ST_Any sync.Once
func Get_bindST_ADT_Control_Bind_Bind_ADT_Control_Monad_ST_Internal_ST_Any() gopurs_runtime.Value {
	once_bindST_ADT_Control_Bind_Bind_ADT_Control_Monad_ST_Internal_ST_Any.Do(func() {
		cache_bindST_ADT_Control_Bind_Bind_ADT_Control_Monad_ST_Internal_ST_Any = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_bind_())
	})
	return cache_bindST_ADT_Control_Bind_Bind_ADT_Control_Monad_ST_Internal_ST_Any
}

var cache_bindST_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any gopurs_runtime.Value
var once_bindST_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any sync.Once
func Get_bindST_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any() gopurs_runtime.Value {
	once_bindST_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any.Do(func() {
		cache_bindST_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_bind_())
	})
	return cache_bindST_Record_Row_bind_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Any_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any
}

var cache_bindST gopurs_runtime.Value
var once_bindST sync.Once
func Get_bindST() gopurs_runtime.Value {
	once_bindST.Do(func() {
		cache_bindST = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_bind_())
	})
	return cache_bindST
}

var cache_bindST__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2435660861 gopurs_runtime.Value
var once_bindST__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2435660861 sync.Once
func Get_bindST__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2435660861() gopurs_runtime.Value {
	once_bindST__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2435660861.Do(func() {
		cache_bindST__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2435660861 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_bind_()})}
	})
	return cache_bindST__ptrControl_Bind_Constructor_Bind_gopurs_runtime_Value__2435660861
}

var cache_bindST__gopurs_runtime_Value_4187656679 gopurs_runtime.Value
var once_bindST__gopurs_runtime_Value_4187656679 sync.Once
func Get_bindST__gopurs_runtime_Value_4187656679() gopurs_runtime.Value {
	once_bindST__gopurs_runtime_Value_4187656679.Do(func() {
		cache_bindST__gopurs_runtime_Value_4187656679 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_bind_())
	})
	return cache_bindST__gopurs_runtime_Value_4187656679
}

var cache_applyST_ADT_Control_Apply_Apply_ADT_Control_Monad_ST_Internal_ST_Any gopurs_runtime.Value
var once_applyST_ADT_Control_Apply_Apply_ADT_Control_Monad_ST_Internal_ST_Any sync.Once
func Get_applyST_ADT_Control_Apply_Apply_ADT_Control_Monad_ST_Internal_ST_Any() gopurs_runtime.Value {
	once_applyST_ADT_Control_Apply_Apply_ADT_Control_Monad_ST_Internal_ST_Any.Do(func() {
		cache_applyST_ADT_Control_Apply_Apply_ADT_Control_Monad_ST_Internal_ST_Any = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyST_ADT_Control_Apply_Apply_ADT_Control_Monad_ST_Internal_ST_Any
}

var cache_applyST_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any gopurs_runtime.Value
var once_applyST_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any sync.Once
func Get_applyST_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any() gopurs_runtime.Value {
	once_applyST_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any.Do(func() {
		cache_applyST_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyST_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any
}

var cache_applyST gopurs_runtime.Value
var once_applyST sync.Once
func Get_applyST() gopurs_runtime.Value {
	once_applyST.Do(func() {
		cache_applyST = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyST
}

var cache_applyST__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2796778301 gopurs_runtime.Value
var once_applyST__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2796778301 sync.Once
func Get_applyST__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2796778301() gopurs_runtime.Value {
	once_applyST__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2796778301.Do(func() {
		cache_applyST__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2796778301 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
})})}
}()
	})
	return cache_applyST__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2796778301
}

var cache_applyST__gopurs_runtime_Value_2741064779 gopurs_runtime.Value
var once_applyST__gopurs_runtime_Value_2741064779 sync.Once
func Get_applyST__gopurs_runtime_Value_2741064779() gopurs_runtime.Value {
	once_applyST__gopurs_runtime_Value_2741064779.Do(func() {
		cache_applyST__gopurs_runtime_Value_2741064779 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyST__gopurs_runtime_Value_2741064779
}

var cache_applicativeST_ADT_Control_Applicative_Applicative_ADT_Control_Monad_ST_Internal_ST_Any gopurs_runtime.Value
var once_applicativeST_ADT_Control_Applicative_Applicative_ADT_Control_Monad_ST_Internal_ST_Any sync.Once
func Get_applicativeST_ADT_Control_Applicative_Applicative_ADT_Control_Monad_ST_Internal_ST_Any() gopurs_runtime.Value {
	once_applicativeST_ADT_Control_Applicative_Applicative_ADT_Control_Monad_ST_Internal_ST_Any.Do(func() {
		cache_applicativeST_ADT_Control_Applicative_Applicative_ADT_Control_Monad_ST_Internal_ST_Any = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_pure_())
	})
	return cache_applicativeST_ADT_Control_Applicative_Applicative_ADT_Control_Monad_ST_Internal_ST_Any
}

var cache_applicativeST_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any gopurs_runtime.Value
var once_applicativeST_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any sync.Once
func Get_applicativeST_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any() gopurs_runtime.Value {
	once_applicativeST_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any.Do(func() {
		cache_applicativeST_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_pure_())
	})
	return cache_applicativeST_Record_Row_pure_ForAll_a_Func_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Apply0_Func_Record_Row__Any_Record_Row_apply_ForAll_a_b_Func_ADT_Control_Monad_ST_Internal_ST_Any_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_ADT_Control_Monad_ST_Internal_ST_Any_Any_Any_Any_Any
}

var cache_applicativeST gopurs_runtime.Value
var once_applicativeST sync.Once
func Get_applicativeST() gopurs_runtime.Value {
	once_applicativeST.Do(func() {
		cache_applicativeST = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_pure_())
	})
	return cache_applicativeST
}

var cache_applicativeST__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__3091537981 gopurs_runtime.Value
var once_applicativeST__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__3091537981 sync.Once
func Get_applicativeST__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__3091537981() gopurs_runtime.Value {
	once_applicativeST__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__3091537981.Do(func() {
		cache_applicativeST__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__3091537981 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_pure_()})}
	})
	return cache_applicativeST__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__3091537981
}

var cache_applicativeST__gopurs_runtime_Value_2868811880 gopurs_runtime.Value
var once_applicativeST__gopurs_runtime_Value_2868811880 sync.Once
func Get_applicativeST__gopurs_runtime_Value_2868811880() gopurs_runtime.Value {
	once_applicativeST__gopurs_runtime_Value_2868811880.Do(func() {
		cache_applicativeST__gopurs_runtime_Value_2868811880 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyST()
}), Get_pure_())
	})
	return cache_applicativeST__gopurs_runtime_Value_2868811880
}

var cache_semigroupST gopurs_runtime.Value
var once_semigroupST sync.Once
func Get_semigroupST() gopurs_runtime.Value {
	once_semigroupST.Do(func() {
		cache_semigroupST = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupST(dictSemigroup_0_box)
})
	})
	return cache_semigroupST
}

var cache_monadRecST gopurs_runtime.Value
var once_monadRecST sync.Once
func Get_monadRecST() gopurs_runtime.Value {
	once_monadRecST.Do(func() {
		cache_monadRecST = gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadST()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Apply(f_0, a_1), Get_newImpl()), gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), Get_bindST(), gopurs_runtime.Apply2(Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorST(), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 525585346) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return gopurs_runtime.Bool((__t0.IntVal) != (0))
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(r_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(r_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 525585346) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindST(), "bind"), gopurs_runtime.Apply(f_0, (*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorST(), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(r_2.PtrVal().(*interface{})) = e_4
return e_4
}))
}))
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 60402430) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeST(), "pure"), pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorST(), "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 60402430) {
__t2 = (*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(r_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}))
}))
}))
})
}))
	})
	return cache_monadRecST
}

var cache_monoidST gopurs_runtime.Value
var once_monoidST sync.Once
func Get_monoidST() gopurs_runtime.Value {
	once_monoidST.Do(func() {
		cache_monoidST = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidST(dictMonoid_0_box)
})
	})
	return cache_monoidST
}

func Call_modify(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_modify__gopurs_runtime_Value_781734141(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_semigroupST(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyST(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_2_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyST(), "apply"), gopurs_runtime.Apply2(Functor0_1_0.V0, __local_var_2_1, a_3), b_4)
})
}))
}

func Call_monoidST(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
semigroupST1_1_0 := Call_semigroupST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupST1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupST1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeST(), "pure"), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")))
}

func Get_bind_() gopurs_runtime.Value {
	return _Gopurs_Bind_
}

func Get_forImpl() gopurs_runtime.Value {
	return _Gopurs_ForImpl
}

func Get_foreach() gopurs_runtime.Value {
	return _Gopurs_Foreach
}

func Get_map_() gopurs_runtime.Value {
	return _Gopurs_Map_
}

func Get_modifyImpl() gopurs_runtime.Value {
	return _Gopurs_ModifyImpl
}

func Get_newImpl() gopurs_runtime.Value {
	return _Gopurs_NewImpl
}

func Get_pure_() gopurs_runtime.Value {
	return _Gopurs_Pure_
}

func Get_read() gopurs_runtime.Value {
	return _Gopurs_Read
}

func Get_run() gopurs_runtime.Value {
	return _Gopurs_Run
}

func Get_while() gopurs_runtime.Value {
	return _Gopurs_While
}

func Get_write() gopurs_runtime.Value {
	return _Gopurs_Write
}
