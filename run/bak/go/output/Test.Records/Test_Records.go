package Test_Records

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Effect_Console "gopurs/output/Effect.Console"
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
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_updateRec(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(1))), gopurs_runtime.RecordUpdate(v1_1, map[string]gopurs_runtime.Value{"a": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["a"]), gopurs_runtime.Int(1)), "b": gopurs_runtime.RecordUpdate(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["b"], map[string]gopurs_runtime.Value{"c": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["b"].PtrVal.(map[string]gopurs_runtime.Value)["c"]), gopurs_runtime.Int(2)), "d": gopurs_runtime.RecordUpdate(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["b"].PtrVal.(map[string]gopurs_runtime.Value)["d"], map[string]gopurs_runtime.Value{"e": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["b"].PtrVal.(map[string]gopurs_runtime.Value)["d"].PtrVal.(map[string]gopurs_runtime.Value)["e"]), gopurs_runtime.Int(3)), "f": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), v1_1.PtrVal.(map[string]gopurs_runtime.Value)["b"].PtrVal.(map[string]gopurs_runtime.Value)["d"].PtrVal.(map[string]gopurs_runtime.Value)["f"]), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intMod(), v_0), gopurs_runtime.Int(5)))})})}))
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
		initial = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"a": gopurs_runtime.Int(0), "b": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"c": gopurs_runtime.Int(0), "d": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"e": gopurs_runtime.Int(0), "f": gopurs_runtime.Int(0)})})})
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
		act = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_updateRec(), gopurs_runtime.Int(10000)), Get_initial()).PtrVal.(map[string]gopurs_runtime.Value)["b"].PtrVal.(map[string]gopurs_runtime.Value)["d"].PtrVal.(map[string]gopurs_runtime.Value)["f"]))
	})
	return act
}


