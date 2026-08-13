package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Ord_Down_Down gopurs_runtime.Value
var once_Data_Ord_Down_Down sync.Once
func Get_Data_Ord_Down_Down() gopurs_runtime.Value {
	once_Data_Ord_Down_Down.Do(func() {
		cache_Data_Ord_Down_Down = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Down_Down(x_0_box)
})
	})
	return cache_Data_Ord_Down_Down
}

var cache_Data_Ord_Down_showDown gopurs_runtime.Value
var once_Data_Ord_Down_showDown sync.Once
func Get_Data_Ord_Down_showDown() gopurs_runtime.Value {
	once_Data_Ord_Down_showDown.Do(func() {
		cache_Data_Ord_Down_showDown = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Down_showDown(dictShow_0_box)
})
	})
	return cache_Data_Ord_Down_showDown
}

var cache_Data_Ord_Down_newtypeDown gopurs_runtime.Value
var once_Data_Ord_Down_newtypeDown sync.Once
func Get_Data_Ord_Down_newtypeDown() gopurs_runtime.Value {
	once_Data_Ord_Down_newtypeDown.Do(func() {
		cache_Data_Ord_Down_newtypeDown = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Ord_Down_newtypeDown
}

var cache_Data_Ord_Down_eqDown gopurs_runtime.Value
var once_Data_Ord_Down_eqDown sync.Once
func Get_Data_Ord_Down_eqDown() gopurs_runtime.Value {
	once_Data_Ord_Down_eqDown.Do(func() {
		cache_Data_Ord_Down_eqDown = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Down_eqDown(dictEq_0_box)
})
	})
	return cache_Data_Ord_Down_eqDown
}

var cache_Data_Ord_Down_ordDown gopurs_runtime.Value
var once_Data_Ord_Down_ordDown sync.Once
func Get_Data_Ord_Down_ordDown() gopurs_runtime.Value {
	once_Data_Ord_Down_ordDown.Do(func() {
		cache_Data_Ord_Down_ordDown = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Down_ordDown(dictOrd_0_box)
})
	})
	return cache_Data_Ord_Down_ordDown
}

var cache_Data_Ord_Down_boundedDown gopurs_runtime.Value
var once_Data_Ord_Down_boundedDown sync.Once
func Get_Data_Ord_Down_boundedDown() gopurs_runtime.Value {
	once_Data_Ord_Down_boundedDown.Do(func() {
		cache_Data_Ord_Down_boundedDown = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_Down_boundedDown(dictBounded_0_box)
})
	})
	return cache_Data_Ord_Down_boundedDown
}

func Call_Data_Ord_Down_Down(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Ord_Down_showDown(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Down ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Ord_Down_eqDown(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_Data_Ord_Down_ordDown(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqDown1_1_0 -> gopurs_runtime.Value
eqDown1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = eqDown1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqDown1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> uint32
__local_var_4_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_2, v1_3).IntVal)
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (__local_var_4_1 == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t2.IntVal)), UnsafePtr: nil}
})
}))
}

func Call_Data_Ord_Down_boundedDown(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqDown1_2_2 -> gopurs_runtime.Value
eqDown1_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Eq0"), gopurs_runtime.Value{})
_ = eqDown1_2_2
// TAST (Let): ordDown1_1_0 -> gopurs_runtime.Value
ordDown1_1_0 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqDown1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> uint32
__local_var_5_3 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "compare"), v_3, v1_4).IntVal)
_ = __local_var_5_3
var __t4 gopurs_runtime.Value
{
if (__local_var_5_3 == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if (__local_var_5_3 == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if (__local_var_5_3 == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t4.IntVal)), UnsafePtr: nil}
})
}))
_ = ordDown1_1_0
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ordDown1_1_0
}), gopurs_runtime.RecordGet(dictBounded_0, "top"), gopurs_runtime.RecordGet(dictBounded_0, "bottom"))
}


