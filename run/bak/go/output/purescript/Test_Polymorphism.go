package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_Polymorphism_logShow gopurs_runtime.Value
var once_Test_Polymorphism_logShow sync.Once
func Get_Test_Polymorphism_logShow() gopurs_runtime.Value {
	once_Test_Polymorphism_logShow.Do(func() {
		cache_Test_Polymorphism_logShow = gopurs_runtime.Apply(Get_Effect_Console_logShow(), Get_Data_Show_showInt())
	})
	return cache_Test_Polymorphism_logShow
}

var cache_Test_Polymorphism_Monoidish_dollarDict gopurs_runtime.Value
var once_Test_Polymorphism_Monoidish_dollarDict sync.Once
func Get_Test_Polymorphism_Monoidish_dollarDict() gopurs_runtime.Value {
	once_Test_Polymorphism_Monoidish_dollarDict.Do(func() {
		cache_Test_Polymorphism_Monoidish_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Polymorphism_Monoidish_dollarDict(x_0_box)
})
	})
	return cache_Test_Polymorphism_Monoidish_dollarDict
}

var cache_Test_Polymorphism_mempty_ gopurs_runtime.Value
var once_Test_Polymorphism_mempty_ sync.Once
func Get_Test_Polymorphism_mempty_() gopurs_runtime.Value {
	once_Test_Polymorphism_mempty_.Do(func() {
		cache_Test_Polymorphism_mempty_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Polymorphism_mempty_(dict_0_box)
})
	})
	return cache_Test_Polymorphism_mempty_
}

var cache_Test_Polymorphism_mappend_ gopurs_runtime.Value
var once_Test_Polymorphism_mappend_ sync.Once
func Get_Test_Polymorphism_mappend_() gopurs_runtime.Value {
	once_Test_Polymorphism_mappend_.Do(func() {
		cache_Test_Polymorphism_mappend_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Polymorphism_mappend_(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish](dict_0_box))
})
	})
	return cache_Test_Polymorphism_mappend_
}

var cache_Test_Polymorphism_polyLoop gopurs_runtime.Value
var once_Test_Polymorphism_polyLoop sync.Once
func Get_Test_Polymorphism_polyLoop() gopurs_runtime.Value {
	once_Test_Polymorphism_polyLoop.Do(func() {
		cache_Test_Polymorphism_polyLoop = gopurs_runtime.Func3(func(dictMonoidish_0_box gopurs_runtime.Value, n_init_1_box gopurs_runtime.Value, acc_init_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Polymorphism_polyLoop(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish](dictMonoidish_0_box), n_init_1_box.IntVal, acc_init_2_box)
})
	})
	return cache_Test_Polymorphism_polyLoop
}

var cache_Test_Polymorphism_intMonoidish gopurs_runtime.Value
var once_Test_Polymorphism_intMonoidish sync.Once
func Get_Test_Polymorphism_intMonoidish() gopurs_runtime.Value {
	once_Test_Polymorphism_intMonoidish.Do(func() {
		cache_Test_Polymorphism_intMonoidish = gopurs_runtime.RecordDict2("mappend_", "mempty_", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((x_0.IntVal) + (y_1.IntVal))
})
}), gopurs_runtime.Int(1))
	})
	return cache_Test_Polymorphism_intMonoidish
}

var cache_Test_Polymorphism_describe gopurs_runtime.Value
var once_Test_Polymorphism_describe sync.Once
func Get_Test_Polymorphism_describe() gopurs_runtime.Value {
	once_Test_Polymorphism_describe.Do(func() {
		cache_Test_Polymorphism_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Polymorphism (10M Type Class Dict Lookups):"))
	})
	return cache_Test_Polymorphism_describe
}

var cache_Test_Polymorphism_act gopurs_runtime.Value
var once_Test_Polymorphism_act sync.Once
func Get_Test_Polymorphism_act() gopurs_runtime.Value {
	once_Test_Polymorphism_act.Do(func() {
		cache_Test_Polymorphism_act = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(10000000))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_Polymorphism_polyLoop__1533381815(__local_var_1_1.IntVal, 0))).StrVal())), gopurs_runtime.Value{})
})
}()
	})
	return cache_Test_Polymorphism_act
}

var cache_Test_Polymorphism_mappend___2927892844 gopurs_runtime.Value
var once_Test_Polymorphism_mappend___2927892844 sync.Once
func Get_Test_Polymorphism_mappend___2927892844() gopurs_runtime.Value {
	once_Test_Polymorphism_mappend___2927892844.Do(func() {
		cache_Test_Polymorphism_mappend___2927892844 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Polymorphism_mappend___2927892844(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish](dict_0_box))
})
	})
	return cache_Test_Polymorphism_mappend___2927892844
}

var cache_Test_Polymorphism_mappend___3566619927 gopurs_runtime.Value
var once_Test_Polymorphism_mappend___3566619927 sync.Once
func Get_Test_Polymorphism_mappend___3566619927() gopurs_runtime.Value {
	once_Test_Polymorphism_mappend___3566619927.Do(func() {
		cache_Test_Polymorphism_mappend___3566619927 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Polymorphism_mappend___3566619927(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish](dict_0_box))
})
	})
	return cache_Test_Polymorphism_mappend___3566619927
}

var cache_Test_Polymorphism_mempty___3961231853 gopurs_runtime.Value
var once_Test_Polymorphism_mempty___3961231853 sync.Once
func Get_Test_Polymorphism_mempty___3961231853() gopurs_runtime.Value {
	once_Test_Polymorphism_mempty___3961231853.Do(func() {
		cache_Test_Polymorphism_mempty___3961231853 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Polymorphism_mempty___3961231853(dict_0_box)
})
	})
	return cache_Test_Polymorphism_mempty___3961231853
}

var cache_Test_Polymorphism_mempty___1540866998 gopurs_runtime.Value
var once_Test_Polymorphism_mempty___1540866998 sync.Once
func Get_Test_Polymorphism_mempty___1540866998() gopurs_runtime.Value {
	once_Test_Polymorphism_mempty___1540866998.Do(func() {
		cache_Test_Polymorphism_mempty___1540866998 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Polymorphism_mempty___1540866998(dict_0_box)
})
	})
	return cache_Test_Polymorphism_mempty___1540866998
}

var cache_Test_Polymorphism_polyLoop__1533381815 gopurs_runtime.Value
var once_Test_Polymorphism_polyLoop__1533381815 sync.Once
func Get_Test_Polymorphism_polyLoop__1533381815() gopurs_runtime.Value {
	once_Test_Polymorphism_polyLoop__1533381815.Do(func() {
		cache_Test_Polymorphism_polyLoop__1533381815 = gopurs_runtime.Func2(func(n_init_0_box gopurs_runtime.Value, acc_init_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Polymorphism_polyLoop__1533381815(n_init_0_box.IntVal, acc_init_1_box.IntVal))
})
	})
	return cache_Test_Polymorphism_polyLoop__1533381815
}

var cache_Test_Polymorphism_polyLoop__2675791634 gopurs_runtime.Value
var once_Test_Polymorphism_polyLoop__2675791634 sync.Once
func Get_Test_Polymorphism_polyLoop__2675791634() gopurs_runtime.Value {
	once_Test_Polymorphism_polyLoop__2675791634.Do(func() {
		cache_Test_Polymorphism_polyLoop__2675791634 = gopurs_runtime.Func3(func(dictMonoidish_0_box gopurs_runtime.Value, n_init_1_box gopurs_runtime.Value, acc_init_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Polymorphism_polyLoop__2675791634(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish](dictMonoidish_0_box), n_init_1_box.IntVal, acc_init_2_box)
})
	})
	return cache_Test_Polymorphism_polyLoop__2675791634
}

type Constructor_Test_Polymorphism_Monoidish struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[459160245] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Test_Polymorphism_Monoidish)(ptr)
		_ = c
		switch key {
		case "mappend_": return gopurs_runtime.Box(c.V0)
		case "mempty_": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Test_Polymorphism_Monoidish: " + key)
		}
	}
}


func Call_Test_Polymorphism_Monoidish_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Test_Polymorphism_mempty_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_Test_Polymorphism_mappend_(dict_0_loop *Constructor_Test_Polymorphism_Monoidish) gopurs_runtime.Value {
var dict_0 *Constructor_Test_Polymorphism_Monoidish = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Test_Polymorphism_polyLoop(dictMonoidish_0_loop *Constructor_Test_Polymorphism_Monoidish, n_init_1_loop int64, acc_init_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoidish_0 *Constructor_Test_Polymorphism_Monoidish = dictMonoidish_0_loop
_ = dictMonoidish_0
var n_init_1 int64 = n_init_1_loop
_ = n_init_1
var acc_init_2 gopurs_runtime.Value = acc_init_2_loop
_ = acc_init_2
var go__go_3_0_0 gopurs_runtime.Value
go__go_3_0_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop int64 = v_4_loop_val.IntVal
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_0:
for {
if false { continue go__go_3_0_0 }
var v_4 int64 = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (v_4) == (0) {
__t1 = v1_5
goto end_branch_1
} else {

}
}
{
v_4_loop = (v_4) - (1)
v1_5_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), v1_5, gopurs_runtime.Box(dictMonoidish_0.V1))
continue go__go_3_0_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply2(go__go_3_0_0, gopurs_runtime.Int(n_init_1), acc_init_2)
}

func Call_Test_Polymorphism_mappend___2927892844(dict_0_loop *Constructor_Test_Polymorphism_Monoidish) gopurs_runtime.Value {
var dict_0 *Constructor_Test_Polymorphism_Monoidish = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Test_Polymorphism_mappend___3566619927(dict_0_loop *Constructor_Test_Polymorphism_Monoidish) gopurs_runtime.Value {
var dict_0 *Constructor_Test_Polymorphism_Monoidish = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Test_Polymorphism_mempty___3961231853(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_Test_Polymorphism_mempty___1540866998(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_Test_Polymorphism_polyLoop__1533381815(n_init_0_loop int64, acc_init_1_loop int64) int64 {
var n_init_0 int64 = n_init_0_loop
_ = n_init_0
var acc_init_1 int64 = acc_init_1_loop
_ = acc_init_1
var go__go_2_0_1 gopurs_runtime.Value
go__go_2_0_1 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop int64 = v_3_loop_val.IntVal
var v1_4_loop int64 = v1_4_loop_val.IntVal
go__go_2_0_1:
for {
if false { continue go__go_2_0_1 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 int64 = v1_4_loop
_ = v1_4
var __t1 int64
{
if (v_3) == (0) {
__t1 = v1_4
goto end_branch_1
} else {

}
}
{
v_3_loop = (v_3) - (1)
v1_4_loop = (v1_4) + (1)
continue go__go_2_0_1
__t1 = gopurs_runtime.Value{}.IntVal
}
end_branch_1:
return gopurs_runtime.Int(__t1)
}
}()
})
})
return gopurs_runtime.Apply2(go__go_2_0_1, gopurs_runtime.Int(n_init_0), gopurs_runtime.Int(acc_init_1)).IntVal
}

func Call_Test_Polymorphism_polyLoop__2675791634(dictMonoidish_0_loop *Constructor_Test_Polymorphism_Monoidish, n_init_1_loop int64, acc_init_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoidish_0 *Constructor_Test_Polymorphism_Monoidish = dictMonoidish_0_loop
_ = dictMonoidish_0
var n_init_1 int64 = n_init_1_loop
_ = n_init_1
var acc_init_2 gopurs_runtime.Value = acc_init_2_loop
_ = acc_init_2
var go__go_3_0_2 gopurs_runtime.Value
go__go_3_0_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop int64 = v_4_loop_val.IntVal
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_2:
for {
if false { continue go__go_3_0_2 }
var v_4 int64 = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (v_4) == (0) {
__t1 = v1_5
goto end_branch_1
} else {

}
}
{
v_4_loop = (v_4) - (1)
v1_5_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), v1_5, gopurs_runtime.Box(dictMonoidish_0.V1))
continue go__go_3_0_2
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply2(go__go_3_0_2, gopurs_runtime.Int(n_init_1), acc_init_2)
}


