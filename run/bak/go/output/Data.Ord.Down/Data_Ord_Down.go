package Data_Ord_Down

import (
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Down gopurs_runtime.Value
var once_Down sync.Once
func Get_Down() gopurs_runtime.Value {
	once_Down.Do(func() {
		cache_Down = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Down(x_0_box)
})
	})
	return cache_Down
}

var cache_showDown gopurs_runtime.Value
var once_showDown sync.Once
func Get_showDown() gopurs_runtime.Value {
	once_showDown.Do(func() {
		cache_showDown = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showDown(dictShow_0_box)
})
	})
	return cache_showDown
}

var cache_newtypeDown gopurs_runtime.Value
var once_newtypeDown sync.Once
func Get_newtypeDown() gopurs_runtime.Value {
	once_newtypeDown.Do(func() {
		cache_newtypeDown = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeDown
}

var cache_eqDown gopurs_runtime.Value
var once_eqDown sync.Once
func Get_eqDown() gopurs_runtime.Value {
	once_eqDown.Do(func() {
		cache_eqDown = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqDown(dictEq_0_box)
})
	})
	return cache_eqDown
}

var cache_ordDown gopurs_runtime.Value
var once_ordDown sync.Once
func Get_ordDown() gopurs_runtime.Value {
	once_ordDown.Do(func() {
		cache_ordDown = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordDown(dictOrd_0_box)
})
	})
	return cache_ordDown
}

var cache_boundedDown gopurs_runtime.Value
var once_boundedDown sync.Once
func Get_boundedDown() gopurs_runtime.Value {
	once_boundedDown.Do(func() {
		cache_boundedDown = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedDown(dictBounded_0_box)
})
	})
	return cache_boundedDown
}

var cache_bottom__338427193 gopurs_runtime.Value
var once_bottom__338427193 sync.Once
func Get_bottom__338427193() gopurs_runtime.Value {
	once_bottom__338427193.Do(func() {
		cache_bottom__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottom__338427193(dict_0_box)
})
	})
	return cache_bottom__338427193
}

var cache_top__338427193 gopurs_runtime.Value
var once_top__338427193 sync.Once
func Get_top__338427193() gopurs_runtime.Value {
	once_top__338427193.Do(func() {
		cache_top__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_top__338427193(dict_0_box)
})
	})
	return cache_top__338427193
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

func Call_Down(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showDown(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(Down "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_eqDown(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_ordDown(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqDown1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = eqDown1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqDown1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_boundedDown(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_1
eqDown1_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Eq0"), gopurs_runtime.Value{})
_ = eqDown1_2_2
ordDown1_1_0 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqDown1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_bottom__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottom")
}

func Call_top__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


