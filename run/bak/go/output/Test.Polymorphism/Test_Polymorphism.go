package Test_Polymorphism

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var mempty_ gopurs_runtime.Value
var once_mempty_ sync.Once
func Get_mempty_() gopurs_runtime.Value {
	once_mempty_.Do(func() {
		mempty_ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty_"]
})
	})
	return mempty_
}

var mappend_ gopurs_runtime.Value
var once_mappend_ sync.Once
func Get_mappend_() gopurs_runtime.Value {
	once_mappend_.Do(func() {
		mappend_ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["mappend_"]
})
	})
	return mappend_
}

var polyLoop gopurs_runtime.Value
var once_polyLoop sync.Once
func Get_polyLoop() gopurs_runtime.Value {
	once_polyLoop.Do(func() {
		polyLoop = gopurs_runtime.Func(func(dictMonoidish_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_1_0 := dictMonoidish_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty_"]
return gopurs_runtime.Func(func(n_init_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_init_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
v_5_loop = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_5), gopurs_runtime.Int(1))
v1_6_loop = gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoidish_0.PtrVal.(map[string]gopurs_runtime.Value)["mappend_"], v1_6), mempty_1_1_0)
continue go__4_1
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
return __t2
}
}()
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__4_1, n_init_2), acc_init_3)
})
})
})
	})
	return polyLoop
}

var intMonoidish gopurs_runtime.Value
var once_intMonoidish sync.Once
func Get_intMonoidish() gopurs_runtime.Value {
	once_intMonoidish.Do(func() {
		intMonoidish = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty_": gopurs_runtime.Int(1), "mappend_": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), x_0), y_1)
})
})})
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
		act = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_polyLoop(), Get_intMonoidish()), gopurs_runtime.Int(10000000)), gopurs_runtime.Int(0))))
	})
	return act
}


