package Test_Polymorphism

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var cache_mempty_ gopurs_runtime.Value
var once_mempty_ sync.Once
func Get_mempty_() gopurs_runtime.Value {
	once_mempty_.Do(func() {
		cache_mempty_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty_(dict_0_box)
})
	})
	return cache_mempty_
}

var cache_mempty___gopurs_runtime_Value_979774650 gopurs_runtime.Value
var once_mempty___gopurs_runtime_Value_979774650 sync.Once
func Get_mempty___gopurs_runtime_Value_979774650() gopurs_runtime.Value {
	once_mempty___gopurs_runtime_Value_979774650.Do(func() {
		cache_mempty___gopurs_runtime_Value_979774650 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty___gopurs_runtime_Value_979774650(dict_0_box)
})
	})
	return cache_mempty___gopurs_runtime_Value_979774650
}

var cache_mappend_ gopurs_runtime.Value
var once_mappend_ sync.Once
func Get_mappend_() gopurs_runtime.Value {
	once_mappend_.Do(func() {
		cache_mappend_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mappend_(dict_0_box)
})
	})
	return cache_mappend_
}

var cache_mappend___gopurs_runtime_Value_512513274 gopurs_runtime.Value
var once_mappend___gopurs_runtime_Value_512513274 sync.Once
func Get_mappend___gopurs_runtime_Value_512513274() gopurs_runtime.Value {
	once_mappend___gopurs_runtime_Value_512513274.Do(func() {
		cache_mappend___gopurs_runtime_Value_512513274 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mappend___gopurs_runtime_Value_512513274(dict_0_box)
})
	})
	return cache_mappend___gopurs_runtime_Value_512513274
}

var cache_polyLoop gopurs_runtime.Value
var once_polyLoop sync.Once
func Get_polyLoop() gopurs_runtime.Value {
	once_polyLoop.Do(func() {
		cache_polyLoop = gopurs_runtime.Func(func(dictMonoidish_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_polyLoop(dictMonoidish_0_box)
})
	})
	return cache_polyLoop
}

var cache_polyLoop__gopurs_runtime_Value_394168575 gopurs_runtime.Value
var once_polyLoop__gopurs_runtime_Value_394168575 sync.Once
func Get_polyLoop__gopurs_runtime_Value_394168575() gopurs_runtime.Value {
	once_polyLoop__gopurs_runtime_Value_394168575.Do(func() {
		cache_polyLoop__gopurs_runtime_Value_394168575 = gopurs_runtime.Func(func(dictMonoidish_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_polyLoop__gopurs_runtime_Value_394168575(dictMonoidish_0_box)
})
	})
	return cache_polyLoop__gopurs_runtime_Value_394168575
}

var cache_intMonoidish gopurs_runtime.Value
var once_intMonoidish sync.Once
func Get_intMonoidish() gopurs_runtime.Value {
	once_intMonoidish.Do(func() {
		cache_intMonoidish = gopurs_runtime.RecordDict2("mappend_", "mempty_", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((x_0.IntVal) + (y_1.IntVal))
}), gopurs_runtime.Int(1))
	})
	return cache_intMonoidish
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Polymorphism (10M Type Class Dict Lookups):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(10000000)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_0 gopurs_runtime.Value
go__go_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0:
for {
if false { continue go__go_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
v_2_loop = gopurs_runtime.Int((v_2.IntVal) - (1))
v1_3_loop = gopurs_runtime.Int((v1_3.IntVal) + (1))
continue go__go_1_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Apply2(go__go_1_0, dummy_0, gopurs_runtime.Int(0))))
}))
	})
	return cache_act
}

func Call_mempty_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_mempty___gopurs_runtime_Value_979774650(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_mappend_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mappend_")
}

func Call_mappend___gopurs_runtime_Value_512513274(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mappend_")
}

func Call_polyLoop(dictMonoidish_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoidish_0 gopurs_runtime.Value = dictMonoidish_0_loop
_ = dictMonoidish_0
mempty_1_1_0 := gopurs_runtime.RecordGet(dictMonoidish_0, "mempty_")
_ = mempty_1_1_0
return gopurs_runtime.Func2(func(n_init_2 gopurs_runtime.Value, acc_init_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1 gopurs_runtime.Value
go__go_4_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1:
for {
if false { continue go__go_4_1 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v_5.IntVal) == (0) {
__t2 = v1_6
goto end_branch_2
} else {

}
}
{
v_5_loop = gopurs_runtime.Int((v_5.IntVal) - (1))
v1_6_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonoidish_0, "mappend_"), v1_6, mempty_1_1_0)
continue go__go_4_1
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
return __t2
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1, n_init_2, acc_init_3)
})
}

func Call_polyLoop__gopurs_runtime_Value_394168575(dictMonoidish_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoidish_0 gopurs_runtime.Value = dictMonoidish_0_loop
_ = dictMonoidish_0
mempty_1_1_0 := gopurs_runtime.RecordGet(dictMonoidish_0, "mempty_")
_ = mempty_1_1_0
return gopurs_runtime.Func2(func(n_init_2 gopurs_runtime.Value, acc_init_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1 gopurs_runtime.Value
go__go_4_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1:
for {
if false { continue go__go_4_1 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v_5.IntVal) == (0) {
__t2 = v1_6
goto end_branch_2
} else {

}
}
{
v_5_loop = gopurs_runtime.Int((v_5.IntVal) - (1))
v1_6_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonoidish_0, "mappend_"), v1_6, mempty_1_1_0)
continue go__go_4_1
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
return __t2
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1, n_init_2, acc_init_3)
})
}


