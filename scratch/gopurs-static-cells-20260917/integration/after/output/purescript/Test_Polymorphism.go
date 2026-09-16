package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_Polymorphism_Monoidish_dollar_Dict gopurs_runtime.Value
var once_Test_Polymorphism_Monoidish_dollar_Dict sync.Once
func Get_Test_Polymorphism_Monoidish_dollar_Dict() gopurs_runtime.Value {
	once_Test_Polymorphism_Monoidish_dollar_Dict.Do(func() {
		cache_Test_Polymorphism_Monoidish_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 459160245, UnsafePtr: unsafe.Pointer(Call_Test_Polymorphism_Monoidish_dollar_Dict(func() struct{
	mappend_ gopurs_runtime.Value
	mempty_ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	mappend_ gopurs_runtime.Value
	mempty_ gopurs_runtime.Value
}{}
					clone.mappend_ = gopurs_runtime.RecordGet(orig, "mappend_")
					clone.mempty_ = gopurs_runtime.RecordGet(orig, "mempty_")
					return clone
				}()))}
})
	})
	return cache_Test_Polymorphism_Monoidish_dollar_Dict
}

var cache_Test_Polymorphism_Monoidish_dollar_Dict__256243535 gopurs_runtime.Value
var once_Test_Polymorphism_Monoidish_dollar_Dict__256243535 sync.Once
func Get_Test_Polymorphism_Monoidish_dollar_Dict__256243535() gopurs_runtime.Value {
	once_Test_Polymorphism_Monoidish_dollar_Dict__256243535.Do(func() {
		cache_Test_Polymorphism_Monoidish_dollar_Dict__256243535 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 459160245, UnsafePtr: unsafe.Pointer(Rebox_Test_Polymorphism_1832189818_1317435201(Call_Test_Polymorphism_Monoidish_dollar_Dict__256243535(func() struct{
	mappend_ gopurs_runtime.Value
	mempty_ int64
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	mappend_ gopurs_runtime.Value
	mempty_ int64
}{}
					clone.mappend_ = gopurs_runtime.RecordGet(orig, "mappend_")
					clone.mempty_ = gopurs_runtime.RecordGet(orig, "mempty_").IntVal
					return clone
				}())))}
})
	})
	return cache_Test_Polymorphism_Monoidish_dollar_Dict__256243535
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

var cache_Test_Polymorphism_intMonoidish gopurs_runtime.Value
var once_Test_Polymorphism_intMonoidish sync.Once
func Get_Test_Polymorphism_intMonoidish() gopurs_runtime.Value {
	once_Test_Polymorphism_intMonoidish.Do(func() {
		cache_Test_Polymorphism_intMonoidish = gopurs_runtime.Value{Type: 9, IntVal: 459160245, UnsafePtr: unsafe.Pointer(Rebox_Test_Polymorphism_1832189818_1317435201((&Constructor_Test_Polymorphism_Monoidish[int64]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((x_0.IntVal) + (y_1.IntVal))
}), int64(1)})))}
	})
	return cache_Test_Polymorphism_intMonoidish
}

var cache_Test_Polymorphism_mempty___647844308 gopurs_runtime.Value
var once_Test_Polymorphism_mempty___647844308 sync.Once
func Get_Test_Polymorphism_mempty___647844308() gopurs_runtime.Value {
	once_Test_Polymorphism_mempty___647844308.Do(func() {
		cache_Test_Polymorphism_mempty___647844308 = gopurs_runtime.Int(gopurs_runtime.Int(int64(1)).IntVal)
	})
	return cache_Test_Polymorphism_mempty___647844308
}

var cache_Test_Polymorphism_mappend_ gopurs_runtime.Value
var once_Test_Polymorphism_mappend_ sync.Once
func Get_Test_Polymorphism_mappend_() gopurs_runtime.Value {
	once_Test_Polymorphism_mappend_.Do(func() {
		cache_Test_Polymorphism_mappend_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Polymorphism_mappend_(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Test_Polymorphism_mappend_
}

var cache_Test_Polymorphism_mappend___65806726 gopurs_runtime.Value
var once_Test_Polymorphism_mappend___65806726 sync.Once
func Get_Test_Polymorphism_mappend___65806726() gopurs_runtime.Value {
	once_Test_Polymorphism_mappend___65806726.Do(func() {
		cache_Test_Polymorphism_mappend___65806726 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Polymorphism_mappend___65806726(__eta_norm_1_0_box.IntVal, __eta_norm_0_unused_1_box.IntVal))
})
	})
	return cache_Test_Polymorphism_mappend___65806726
}

var cache_Test_Polymorphism_mappend___854306779 gopurs_runtime.Value
var once_Test_Polymorphism_mappend___854306779 sync.Once
func Get_Test_Polymorphism_mappend___854306779() gopurs_runtime.Value {
	once_Test_Polymorphism_mappend___854306779.Do(func() {
		cache_Test_Polymorphism_mappend___854306779 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Polymorphism_mappend___854306779(__eta_norm_1_0_box.IntVal, __eta_norm_0_1_box.IntVal))
})
	})
	return cache_Test_Polymorphism_mappend___854306779
}

var cache_Test_Polymorphism_polyLoop gopurs_runtime.Value
var once_Test_Polymorphism_polyLoop sync.Once
func Get_Test_Polymorphism_polyLoop() gopurs_runtime.Value {
	once_Test_Polymorphism_polyLoop.Do(func() {
		cache_Test_Polymorphism_polyLoop = gopurs_runtime.Func3(func(dictMonoidish_0_box gopurs_runtime.Value, n_init_1_box gopurs_runtime.Value, acc_init_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Polymorphism_polyLoop(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value]](dictMonoidish_0_box), n_init_1_box.IntVal, acc_init_2_box)
})
	})
	return cache_Test_Polymorphism_polyLoop
}

var cache_Test_Polymorphism_polyLoop__854306779 gopurs_runtime.Value
var once_Test_Polymorphism_polyLoop__854306779 sync.Once
func Get_Test_Polymorphism_polyLoop__854306779() gopurs_runtime.Value {
	once_Test_Polymorphism_polyLoop__854306779.Do(func() {
		cache_Test_Polymorphism_polyLoop__854306779 = gopurs_runtime.Func2(func(n_init_0_box gopurs_runtime.Value, acc_init_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Polymorphism_polyLoop__854306779(n_init_0_box.IntVal, acc_init_1_box.IntVal))
})
	})
	return cache_Test_Polymorphism_polyLoop__854306779
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
		cache_Test_Polymorphism_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(int64(10000000)))
_ = __local_var_0_0
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
var Call_local_Test_Polymorphism_go__467072791_2_2_3 func(int64, int64) int64
_ = Call_local_Test_Polymorphism_go__467072791_2_2_3
var go__467072791_2_2_3 gopurs_runtime.Value
_ = go__467072791_2_2_3
var Call_local_Test_Polymorphism_go__go_2_3_4 func(int64, int64) int64
_ = Call_local_Test_Polymorphism_go__go_2_3_4
var go__go_2_3_4 gopurs_runtime.Value
_ = go__go_2_3_4
Call_local_Test_Polymorphism_go__467072791_2_2_3 = func(v_3_loop int64, v1_4_loop int64) int64 {
go__467072791_2_2_3:
for {
if false { continue go__467072791_2_2_3 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 int64 = v1_4_loop
_ = v1_4
var __t4 int64
{
if (v_3) == (int64(0)) {
__t4 = v1_4
goto end_branch_4
} else {

}
}
{
v_3_loop = (v_3) - (int64(1))
v1_4_loop = (v1_4) + (int64(1))
continue go__467072791_2_2_3
__t4 = func() int64 { panic("unreachable") }()
}
end_branch_4:
return __t4
}
}
go__467072791_2_2_3 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Test_Polymorphism_go__467072791_2_2_3(v_3_loop_val.IntVal, v1_4_loop_val.IntVal))
})
})
Call_local_Test_Polymorphism_go__go_2_3_4 = func(v_3_loop int64, v1_4_loop int64) int64 {
go__go_2_3_4:
for {
if false { continue go__go_2_3_4 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 int64 = v1_4_loop
_ = v1_4
var __t5 int64
{
if (v_3) == (int64(0)) {
__t5 = v1_4
goto end_branch_5
} else {

}
}
{
__t5 = Call_local_Test_Polymorphism_go__467072791_2_2_3((v_3) - (int64(1)), (v1_4) + (int64(1)))
}
end_branch_5:
return __t5
}
}
go__go_2_3_4 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Test_Polymorphism_go__go_2_3_4(v_3_loop_val.IntVal, v1_4_loop_val.IntVal))
})
})
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_local_Test_Polymorphism_go__467072791_2_2_3(__local_var_1_1.IntVal, int64(0)))).StrVal())
})
	})
	return cache_Test_Polymorphism_act
}

type Constructor_Test_Polymorphism_Monoidish[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_a
}


func init() {
	gopurs_runtime.StructGetters[459160245] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "mappend_": return gopurs_runtime.Box(c.V0)
		case "mempty_": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Test_Polymorphism_Monoidish: " + key)
		}
	}
}


func Call_Test_Polymorphism_Monoidish_dollar_Dict(x_0_loop struct{
	mappend_ gopurs_runtime.Value
	mempty_ gopurs_runtime.Value
}) *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value] {
var x_0 struct{
	mappend_ gopurs_runtime.Value
	mempty_ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("mappend_", "mempty_", orig.mappend_, orig.mempty_)
				}())
}

func Call_Test_Polymorphism_Monoidish_dollar_Dict__256243535(x_0_loop struct{
	mappend_ gopurs_runtime.Value
	mempty_ int64
}) *Constructor_Test_Polymorphism_Monoidish[int64] {
Monoidish_dollar_Dict__256243535:
for {
if false { continue Monoidish_dollar_Dict__256243535 }
var x_0 struct{
	mappend_ gopurs_runtime.Value
	mempty_ int64
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish[int64]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("mappend_", "mempty_", orig.mappend_, gopurs_runtime.Int(orig.mempty_))
				}())
}
}

func Call_Test_Polymorphism_mempty_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_Test_Polymorphism_mappend_(dict_0_loop *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_Test_Polymorphism_mappend___65806726(__eta_norm_1_0_loop int64, __eta_norm_0_unused_1_loop int64) int64 {
mappend___65806726:
for {
if false { continue mappend___65806726 }
var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 int64 = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return (__eta_norm_1_0) + (int64(1))
}
}

func Call_Test_Polymorphism_mappend___854306779(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop int64) int64 {
mappend___854306779:
for {
if false { continue mappend___854306779 }
var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return (__eta_norm_1_0) + (__eta_norm_0_1)
}
}

func Call_Test_Polymorphism_polyLoop(dictMonoidish_0_loop *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value], n_init_1_loop int64, acc_init_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoidish_0 *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value] = dictMonoidish_0_loop
_ = dictMonoidish_0
var n_init_1 int64 = n_init_1_loop
_ = n_init_1
var acc_init_2 gopurs_runtime.Value = acc_init_2_loop
_ = acc_init_2
var Call_local_Test_Polymorphism_go__go_3_0_0 func(int64, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Test_Polymorphism_go__go_3_0_0
var go__go_3_0_0 gopurs_runtime.Value
_ = go__go_3_0_0
Call_local_Test_Polymorphism_go__go_3_0_0 = func(v_4_loop int64, v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_0_0:
for {
if false { continue go__go_3_0_0 }
var v_4 int64 = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (v_4) == (int64(0)) {
__t1 = v1_5
goto end_branch_1
} else {

}
}
{
v_4_loop = (v_4) - (int64(1))
v1_5_loop = gopurs_runtime.Apply2(dictMonoidish_0.V0, v1_5, dictMonoidish_0.V1)
continue go__go_3_0_0
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_1:
return __t1
}
}
go__go_3_0_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Test_Polymorphism_go__go_3_0_0(v_4_loop_val.IntVal, v1_5_loop_val)
})
})
return Call_local_Test_Polymorphism_go__go_3_0_0(n_init_1, acc_init_2)
}

func Call_Test_Polymorphism_polyLoop__854306779(n_init_0_loop int64, acc_init_1_loop int64) int64 {
polyLoop__854306779:
for {
if false { continue polyLoop__854306779 }
var n_init_0 int64 = n_init_0_loop
_ = n_init_0
var acc_init_1 int64 = acc_init_1_loop
_ = acc_init_1
var Call_local_Test_Polymorphism_go__467072791_2_0_1 func(int64, int64) int64
_ = Call_local_Test_Polymorphism_go__467072791_2_0_1
var go__467072791_2_0_1 gopurs_runtime.Value
_ = go__467072791_2_0_1
var Call_local_Test_Polymorphism_go__go_2_1_2 func(int64, int64) int64
_ = Call_local_Test_Polymorphism_go__go_2_1_2
var go__go_2_1_2 gopurs_runtime.Value
_ = go__go_2_1_2
Call_local_Test_Polymorphism_go__467072791_2_0_1 = func(v_3_loop int64, v1_4_loop int64) int64 {
go__467072791_2_0_1:
for {
if false { continue go__467072791_2_0_1 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 int64 = v1_4_loop
_ = v1_4
var __t2 int64
{
if (v_3) == (int64(0)) {
__t2 = v1_4
goto end_branch_2
} else {

}
}
{
v_3_loop = (v_3) - (int64(1))
v1_4_loop = (v1_4) + (int64(1))
continue go__467072791_2_0_1
__t2 = func() int64 { panic("unreachable") }()
}
end_branch_2:
return __t2
}
}
go__467072791_2_0_1 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Test_Polymorphism_go__467072791_2_0_1(v_3_loop_val.IntVal, v1_4_loop_val.IntVal))
})
})
Call_local_Test_Polymorphism_go__go_2_1_2 = func(v_3_loop int64, v1_4_loop int64) int64 {
go__go_2_1_2:
for {
if false { continue go__go_2_1_2 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 int64 = v1_4_loop
_ = v1_4
var __t3 int64
{
if (v_3) == (int64(0)) {
__t3 = v1_4
goto end_branch_3
} else {

}
}
{
__t3 = Call_local_Test_Polymorphism_go__467072791_2_0_1((v_3) - (int64(1)), (v1_4) + (int64(1)))
}
end_branch_3:
return __t3
}
}
go__go_2_1_2 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Test_Polymorphism_go__go_2_1_2(v_3_loop_val.IntVal, v1_4_loop_val.IntVal))
})
})
return Call_local_Test_Polymorphism_go__467072791_2_0_1(n_init_0, acc_init_1)
}
}

func Rebox_Test_Polymorphism_1832189818_1317435201(in *Constructor_Test_Polymorphism_Monoidish[int64]) *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Int(in.V1)
	return out
}


