package Test_Records

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var updateRec gopurs_runtime.Value
var once_updateRec sync.Once
func Get_updateRec() gopurs_runtime.Value {
	once_updateRec.Do(func() {
		updateRec = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
updateRec:
for {
if false { continue updateRec }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply2(Get_updateRec(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.RecordUpdateDict(v1_1, []string{"a", "b"}, []gopurs_runtime.Value{gopurs_runtime.Int(gopurs_runtime.RecordGet(v1_1, "a").IntVal + gopurs_runtime.Int(1).IntVal), gopurs_runtime.RecordUpdateDict(gopurs_runtime.RecordGet(v1_1, "b"), []string{"c", "d"}, []gopurs_runtime.Value{gopurs_runtime.Int(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "c").IntVal + gopurs_runtime.Int(2).IntVal), gopurs_runtime.RecordUpdateDict(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), []string{"e", "f"}, []gopurs_runtime.Value{gopurs_runtime.Int(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "e").IntVal + gopurs_runtime.Int(3).IntVal), gopurs_runtime.Int(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "f").IntVal + gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), v_0, gopurs_runtime.Int(5)).IntVal)})})}))
}
end_branch_0:
return __t0
}
}()
})
})
	})
	return updateRec
}

var initial gopurs_runtime.Value
var once_initial sync.Once
func Get_initial() gopurs_runtime.Value {
	once_initial.Do(func() {
		initial = gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Int(0), gopurs_runtime.RecordDict2("c", "d", gopurs_runtime.Int(0), gopurs_runtime.RecordDict2("e", "f", gopurs_runtime.Int(0), gopurs_runtime.Int(0))))
	})
	return initial
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Deep Record Updates (10k iterations):"))
	})
	return describe
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(10000))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_updateRec(), dummy_1_1, Get_initial()), "b"), "d"), "f"))), gopurs_runtime.Value{})
})
}()
	})
	return act
}


