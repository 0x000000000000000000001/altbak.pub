package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_TestDPE_Monoidish_dollarDict gopurs_runtime.Value
var once_TestDPE_Monoidish_dollarDict sync.Once
func Get_TestDPE_Monoidish_dollarDict() gopurs_runtime.Value {
	once_TestDPE_Monoidish_dollarDict.Do(func() {
		cache_TestDPE_Monoidish_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_TestDPE_Monoidish_dollarDict(x_0_box)
})
	})
	return cache_TestDPE_Monoidish_dollarDict
}

var cache_TestDPE_mempty_ gopurs_runtime.Value
var once_TestDPE_mempty_ sync.Once
func Get_TestDPE_mempty_() gopurs_runtime.Value {
	once_TestDPE_mempty_.Do(func() {
		cache_TestDPE_mempty_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_TestDPE_mempty_(dict_0_box)
})
	})
	return cache_TestDPE_mempty_
}

var cache_TestDPE_mappend_ gopurs_runtime.Value
var once_TestDPE_mappend_ sync.Once
func Get_TestDPE_mappend_() gopurs_runtime.Value {
	once_TestDPE_mappend_.Do(func() {
		cache_TestDPE_mappend_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_TestDPE_mappend_(gopurs_runtime.CoerceToStruct[Constructor_TestDPE_Monoidish](dict_0_box))
})
	})
	return cache_TestDPE_mappend_
}

var cache_TestDPE_polyLoop gopurs_runtime.Value
var once_TestDPE_polyLoop sync.Once
func Get_TestDPE_polyLoop() gopurs_runtime.Value {
	once_TestDPE_polyLoop.Do(func() {
		cache_TestDPE_polyLoop = gopurs_runtime.Func(func(dictMonoidish_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_TestDPE_polyLoop(gopurs_runtime.CoerceToStruct[Constructor_TestDPE_Monoidish](dictMonoidish_0_box))
})
	})
	return cache_TestDPE_polyLoop
}

var cache_TestDPE_intMonoidish gopurs_runtime.Value
var once_TestDPE_intMonoidish sync.Once
func Get_TestDPE_intMonoidish() gopurs_runtime.Value {
	once_TestDPE_intMonoidish.Do(func() {
		cache_TestDPE_intMonoidish = gopurs_runtime.RecordDict2("mappend_", "mempty_", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((a_0.IntVal) + (b_1.IntVal))
})
}), gopurs_runtime.Int(0))
	})
	return cache_TestDPE_intMonoidish
}

var cache_TestDPE_test gopurs_runtime.Value
var once_TestDPE_test sync.Once
func Get_TestDPE_test() gopurs_runtime.Value {
	once_TestDPE_test.Do(func() {
		cache_TestDPE_test = gopurs_runtime.Int(Call_TestDPE_polyLoop__1285086188(gopurs_runtime.Int(10000000), gopurs_runtime.Int(0)).IntVal)
	})
	return cache_TestDPE_test
}

var cache_TestDPE_mappend___2927892844 gopurs_runtime.Value
var once_TestDPE_mappend___2927892844 sync.Once
func Get_TestDPE_mappend___2927892844() gopurs_runtime.Value {
	once_TestDPE_mappend___2927892844.Do(func() {
		cache_TestDPE_mappend___2927892844 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_TestDPE_mappend___2927892844(gopurs_runtime.CoerceToStruct[Constructor_TestDPE_Monoidish](dict_0_box))
})
	})
	return cache_TestDPE_mappend___2927892844
}

var cache_TestDPE_mappend___1285086188 gopurs_runtime.Value
var once_TestDPE_mappend___1285086188 sync.Once
func Get_TestDPE_mappend___1285086188() gopurs_runtime.Value {
	once_TestDPE_mappend___1285086188.Do(func() {
		cache_TestDPE_mappend___1285086188 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_TestDPE_mappend___1285086188(gopurs_runtime.CoerceToStruct[Constructor_TestDPE_Monoidish](dict_0_box))
})
	})
	return cache_TestDPE_mappend___1285086188
}

var cache_TestDPE_mappend___3425151628 gopurs_runtime.Value
var once_TestDPE_mappend___3425151628 sync.Once
func Get_TestDPE_mappend___3425151628() gopurs_runtime.Value {
	once_TestDPE_mappend___3425151628.Do(func() {
		cache_TestDPE_mappend___3425151628 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_TestDPE_mappend___3425151628(gopurs_runtime.CoerceToStruct[Constructor_TestDPE_Monoidish](dict_0_box))
})
	})
	return cache_TestDPE_mappend___3425151628
}

var cache_TestDPE_polyLoop__2927892844 gopurs_runtime.Value
var once_TestDPE_polyLoop__2927892844 sync.Once
func Get_TestDPE_polyLoop__2927892844() gopurs_runtime.Value {
	once_TestDPE_polyLoop__2927892844.Do(func() {
		cache_TestDPE_polyLoop__2927892844 = gopurs_runtime.Func(func(dictMonoidish_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_TestDPE_polyLoop__2927892844(gopurs_runtime.CoerceToStruct[Constructor_TestDPE_Monoidish](dictMonoidish_0_box))
})
	})
	return cache_TestDPE_polyLoop__2927892844
}

var cache_TestDPE_polyLoop__1285086188 gopurs_runtime.Value
var once_TestDPE_polyLoop__1285086188 sync.Once
func Get_TestDPE_polyLoop__1285086188() gopurs_runtime.Value {
	once_TestDPE_polyLoop__1285086188.Do(func() {
		cache_TestDPE_polyLoop__1285086188 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_TestDPE_polyLoop__1285086188(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_TestDPE_polyLoop__1285086188
}

var cache_TestDPE_polyLoop__11377801 gopurs_runtime.Value
var once_TestDPE_polyLoop__11377801 sync.Once
func Get_TestDPE_polyLoop__11377801() gopurs_runtime.Value {
	once_TestDPE_polyLoop__11377801.Do(func() {
		cache_TestDPE_polyLoop__11377801 = gopurs_runtime.Func(func(dictMonoidish_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_TestDPE_polyLoop__11377801(gopurs_runtime.CoerceToStruct[Constructor_TestDPE_Monoidish](dictMonoidish_0_box))
})
	})
	return cache_TestDPE_polyLoop__11377801
}

type Constructor_TestDPE_Monoidish struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2768669742] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_TestDPE_Monoidish)(ptr)
		_ = c
		switch key {
		case "mappend_": return gopurs_runtime.Box(c.V0)
		case "mempty_": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_TestDPE_Monoidish: " + key)
		}
	}
}


func Call_TestDPE_Monoidish_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_TestDPE_mempty_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_TestDPE_mappend_(dict_0_loop *Constructor_TestDPE_Monoidish) gopurs_runtime.Value {
var dict_0 *Constructor_TestDPE_Monoidish = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_TestDPE_polyLoop(dictMonoidish_0_loop *Constructor_TestDPE_Monoidish) gopurs_runtime.Value {
polyLoop:
for {
if false { continue polyLoop }
var dictMonoidish_0 *Constructor_TestDPE_Monoidish = dictMonoidish_0_loop
_ = dictMonoidish_0
// TAST (Let): mempty_1_1_0 -> gopurs_runtime.Value
mempty_1_1_0 := gopurs_runtime.Box(dictMonoidish_0.V1)
_ = mempty_1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(Call_TestDPE_polyLoop(dictMonoidish_0), gopurs_runtime.Int((v_2.IntVal) - (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), v1_3, mempty_1_1_0))
}
end_branch_1:
return __t1
})
})
}
}

func Call_TestDPE_mappend___2927892844(dict_0_loop *Constructor_TestDPE_Monoidish) gopurs_runtime.Value {
var dict_0 *Constructor_TestDPE_Monoidish = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_TestDPE_mappend___1285086188(dict_0_loop *Constructor_TestDPE_Monoidish) gopurs_runtime.Value {
var dict_0 *Constructor_TestDPE_Monoidish = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_TestDPE_mappend___3425151628(dict_0_loop *Constructor_TestDPE_Monoidish) gopurs_runtime.Value {
var dict_0 *Constructor_TestDPE_Monoidish = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_TestDPE_polyLoop__2927892844(dictMonoidish_0_loop *Constructor_TestDPE_Monoidish) gopurs_runtime.Value {
var dictMonoidish_0 *Constructor_TestDPE_Monoidish = dictMonoidish_0_loop
_ = dictMonoidish_0
// TAST (Let): mempty_1_1_0 -> int64
mempty_1_1_0 := gopurs_runtime.Box(dictMonoidish_0.V1).IntVal
_ = mempty_1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 int64
{
if (v_2.IntVal) == (0) {
__t4 = v1_3.IntVal
goto end_branch_4
} else {

}
}
{
// TAST (Let): __local_var_4_1 -> int64
__local_var_4_1 := (v_2.IntVal) - (1)
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 -> int64
__local_var_5_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), gopurs_runtime.Int(v1_3.IntVal), gopurs_runtime.Int(mempty_1_1_0)).IntVal
_ = __local_var_5_2
var __t3 gopurs_runtime.Value
{
if (__local_var_4_1) == (0) {
__t3 = gopurs_runtime.Int(__local_var_5_2)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(Call_TestDPE_polyLoop(gopurs_runtime.CoerceToStruct[Constructor_TestDPE_Monoidish](gopurs_runtime.Value{Type: 9, IntVal: 2768669742, UnsafePtr: unsafe.Pointer(dictMonoidish_0)})), gopurs_runtime.Int((__local_var_4_1) - (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), gopurs_runtime.Int(__local_var_5_2), gopurs_runtime.Box(dictMonoidish_0.V1)))
}
end_branch_3:
__t4 = __t3.IntVal
}
end_branch_4:
return gopurs_runtime.Int(__t4)
})
})
}

func Call_TestDPE_polyLoop__1285086188(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __t4 int64
{
if (__eta0_0.IntVal) == (0) {
__t4 = __eta1_1.IntVal
goto end_branch_4
} else {

}
}
{
// TAST (Let): __local_var_2_0 -> *Constructor_TestDPE_Monoidish
__local_var_2_0 := &Constructor_TestDPE_Monoidish{1, gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((a_2.IntVal) + (b_3.IntVal))
})
}), gopurs_runtime.Int(0)}
_ = __local_var_2_0
// TAST (Let): __local_var_3_1 -> int64
__local_var_3_1 := (__eta0_0.IntVal) - (1)
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> int64
__local_var_4_2 := (__eta1_1.IntVal) + (0)
_ = __local_var_4_2
var __t3 gopurs_runtime.Value
{
if (__local_var_3_1) == (0) {
__t3 = gopurs_runtime.Int(__local_var_4_2)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(Call_TestDPE_polyLoop(gopurs_runtime.CoerceToStruct[Constructor_TestDPE_Monoidish](gopurs_runtime.Value{Type: 9, IntVal: 2768669742, UnsafePtr: unsafe.Pointer(__local_var_2_0)})), gopurs_runtime.Int((__local_var_3_1) - (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_2_0.V0), gopurs_runtime.Int(__local_var_4_2), gopurs_runtime.Box(__local_var_2_0.V1)))
}
end_branch_3:
__t4 = __t3.IntVal
}
end_branch_4:
return gopurs_runtime.Int(__t4)
}

func Call_TestDPE_polyLoop__11377801(dictMonoidish_0_loop *Constructor_TestDPE_Monoidish) gopurs_runtime.Value {
var dictMonoidish_0 *Constructor_TestDPE_Monoidish = dictMonoidish_0_loop
_ = dictMonoidish_0
// TAST (Let): mempty_1_1_0 -> gopurs_runtime.Value
mempty_1_1_0 := gopurs_runtime.Box(dictMonoidish_0.V1)
_ = mempty_1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t4 = v1_3
goto end_branch_4
} else {

}
}
{
// TAST (Let): __local_var_4_1 -> int64
__local_var_4_1 := (v_2.IntVal) - (1)
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), v1_3, mempty_1_1_0)
_ = __local_var_5_2
var __t3 gopurs_runtime.Value
{
if (__local_var_4_1) == (0) {
__t3 = __local_var_5_2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(Call_TestDPE_polyLoop(dictMonoidish_0), gopurs_runtime.Int((__local_var_4_1) - (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), __local_var_5_2, gopurs_runtime.Box(dictMonoidish_0.V1)))
}
end_branch_3:
__t4 = __t3
}
end_branch_4:
return __t4
})
})
}


