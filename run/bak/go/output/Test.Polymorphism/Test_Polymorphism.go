package Test_Polymorphism

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
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
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_5.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t2 = v1_6
goto end_branch_2
} else {

}
}
{
v_5_loop = gopurs_runtime.Int(v_5.IntVal - gopurs_runtime.Int(1).IntVal)
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
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__0_0:
for {
if false { continue go__0_0 }
var v_1 = v_1_loop
_ = v_1
var v1_2 = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t1 = v1_2
goto end_branch_1
} else {

}
}
{
v_1_loop = gopurs_runtime.Int(v_1.IntVal - gopurs_runtime.Int(1).IntVal)
v1_2_loop = gopurs_runtime.Int(v1_2.IntVal + gopurs_runtime.Int(1).IntVal)
continue go__0_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply2(go__0_0, gopurs_runtime.Int(10000000), gopurs_runtime.Int(0))))
}()
	})
	return act
}


