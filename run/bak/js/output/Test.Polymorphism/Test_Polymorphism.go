package Test_Polymorphism

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var mempty_ gopurs_runtime.Value
var once_mempty_ sync.Once
func Get_mempty_() gopurs_runtime.Value {
	once_mempty_.Do(func() {
		mempty_ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "mempty_")
})
	})
	return mempty_
}

var mappend_ gopurs_runtime.Value
var once_mappend_ sync.Once
func Get_mappend_() gopurs_runtime.Value {
	once_mappend_.Do(func() {
		mappend_ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "mappend_")
})
	})
	return mappend_
}

var polyLoop gopurs_runtime.Value
var once_polyLoop sync.Once
func Get_polyLoop() gopurs_runtime.Value {
	once_polyLoop.Do(func() {
		polyLoop = gopurs_runtime.Func(func(dictMonoidish_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_1_0 := gopurs_runtime.RecordGet(dictMonoidish_0, "mempty_")
_ = mempty_1_1_0
return gopurs_runtime.Func2(func(n_init_2 gopurs_runtime.Value, acc_init_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_1 gopurs_runtime.Value
go__4_1 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_1:
for {
if false { continue go__4_1 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if v_5.IntVal == 0 {
__t2 = v1_6
goto end_branch_2
} else {

}
}
{
v_5_loop = v_5.IntVal - 1
v1_6_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonoidish_0, "mappend_"), v1_6, mempty_1_1_0)
continue go__4_1
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
return __t2
}
}()
})
})
return gopurs_runtime.Apply2(go__4_1, n_init_2, acc_init_3)
})
})
	})
	return polyLoop
}

var intMonoidish gopurs_runtime.Value
var once_intMonoidish sync.Once
func Get_intMonoidish() gopurs_runtime.Value {
	once_intMonoidish.Do(func() {
		intMonoidish = gopurs_runtime.RecordDict2("mempty_", "mappend_", gopurs_runtime.Int(1), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(x_0.IntVal + y_1.IntVal)
}))
	})
	return intMonoidish
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Polymorphism (10M Type Class Dict Lookups):"))
	})
	return describe
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(10000000))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
var go__2_2 gopurs_runtime.Value
go__2_2 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_2:
for {
if false { continue go__2_2 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t3 gopurs_runtime.Value
{
if v_3.IntVal == 0 {
__t3 = v1_4
goto end_branch_3
} else {

}
}
{
v_3_loop = v_3.IntVal - 1
v1_4_loop = v1_4.IntVal + 1
continue go__2_2
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
return __t3
}
}()
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply2(go__2_2, dummy_1_1, gopurs_runtime.Int(0)))), gopurs_runtime.Value{})
})
}()
	})
	return act
}




