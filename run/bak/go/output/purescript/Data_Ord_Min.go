package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ord_Min_Min gopurs_runtime.Value
var once_Data_Ord_Min_Min sync.Once
func Get_Data_Ord_Min_Min() gopurs_runtime.Value {
	once_Data_Ord_Min_Min.Do(func() {
		cache_Data_Ord_Min_Min = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Min_Min(x_0_box)
})
	})
	return cache_Data_Ord_Min_Min
}

var cache_Data_Ord_Min_showMin gopurs_runtime.Value
var once_Data_Ord_Min_showMin sync.Once
func Get_Data_Ord_Min_showMin() gopurs_runtime.Value {
	once_Data_Ord_Min_showMin.Do(func() {
		cache_Data_Ord_Min_showMin = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Min_showMin(dictShow_0_box)
})
	})
	return cache_Data_Ord_Min_showMin
}

var cache_Data_Ord_Min_semigroupMin gopurs_runtime.Value
var once_Data_Ord_Min_semigroupMin sync.Once
func Get_Data_Ord_Min_semigroupMin() gopurs_runtime.Value {
	once_Data_Ord_Min_semigroupMin.Do(func() {
		cache_Data_Ord_Min_semigroupMin = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Min_semigroupMin(dictOrd_0_box)
})
	})
	return cache_Data_Ord_Min_semigroupMin
}

var cache_Data_Ord_Min_newtypeMin gopurs_runtime.Value
var once_Data_Ord_Min_newtypeMin sync.Once
func Get_Data_Ord_Min_newtypeMin() gopurs_runtime.Value {
	once_Data_Ord_Min_newtypeMin.Do(func() {
		cache_Data_Ord_Min_newtypeMin = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_Ord_Min_newtypeMin
}

var cache_Data_Ord_Min_monoidMin gopurs_runtime.Value
var once_Data_Ord_Min_monoidMin sync.Once
func Get_Data_Ord_Min_monoidMin() gopurs_runtime.Value {
	once_Data_Ord_Min_monoidMin.Do(func() {
		cache_Data_Ord_Min_monoidMin = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Min_monoidMin(dictBounded_0_box)
})
	})
	return cache_Data_Ord_Min_monoidMin
}

var cache_Data_Ord_Min_eqMin gopurs_runtime.Value
var once_Data_Ord_Min_eqMin sync.Once
func Get_Data_Ord_Min_eqMin() gopurs_runtime.Value {
	once_Data_Ord_Min_eqMin.Do(func() {
		cache_Data_Ord_Min_eqMin = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Min_eqMin(dictEq_0_box)
})
	})
	return cache_Data_Ord_Min_eqMin
}

var cache_Data_Ord_Min_ordMin gopurs_runtime.Value
var once_Data_Ord_Min_ordMin sync.Once
func Get_Data_Ord_Min_ordMin() gopurs_runtime.Value {
	once_Data_Ord_Min_ordMin.Do(func() {
		cache_Data_Ord_Min_ordMin = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Min_ordMin(dictOrd_0_box)
})
	})
	return cache_Data_Ord_Min_ordMin
}

func Call_Data_Ord_Min_Min(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Ord_Min_showMin(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Min ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
})})}
}

func Call_Data_Ord_Min_semigroupMin(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_1, v1_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
__t1 = v1_2
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
})})}
}

func Call_Data_Ord_Min_monoidMin(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupMin1_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupMin1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_2 -> gopurs_runtime.Value
v_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "compare"), v_2, v1_3)
_ = v_4_2
var __t3 gopurs_runtime.Value
{
if (uint32(v_4_2.IntVal) == 1527465420) {
__t3 = v_2
goto end_branch_3
} else {

}
}
{
if (uint32(v_4_2.IntVal) == 902936544) {
__t3 = v_2
goto end_branch_3
} else {

}
}
{
if (uint32(v_4_2.IntVal) == 380165415) {
__t3 = v1_3
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
})))
_ = semigroupMin1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupMin1_1_0)}
}), gopurs_runtime.RecordGet(dictBounded_0, "top")})}
}

func Call_Data_Ord_Min_eqMin(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0))}
}

func Call_Data_Ord_Min_ordMin(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqMin1_1_0 -> *Constructor_Data_Eq_Eq
eqMin1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqMin1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMin1_1_0)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_2, v1_3).IntVal)), UnsafePtr: nil}
})
})})}
}


