package Test_Records

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
)

var cache_updateRec gopurs_runtime.Value
var once_updateRec sync.Once
func Get_updateRec() gopurs_runtime.Value {
	once_updateRec.Do(func() {
		cache_updateRec = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_updateRec(v_0_box.IntVal, v1_1_box)
})
	})
	return cache_updateRec
}

var cache_initial gopurs_runtime.Value
var once_initial sync.Once
func Get_initial() gopurs_runtime.Value {
	once_initial.Do(func() {
		cache_initial = gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Int(0), gopurs_runtime.RecordDict2("c", "d", gopurs_runtime.Int(0), gopurs_runtime.RecordDict2("e", "f", gopurs_runtime.Int(0), gopurs_runtime.Int(0))))
	})
	return cache_initial
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Deep Record Updates (10k iterations):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(10000)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Call_updateRec(dummy_0.IntVal, Get_initial()), "b"), "d"), "f")))
}))
	})
	return cache_act
}

func Call_updateRec(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
updateRec:
for {
if false { continue updateRec }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (1)
v1_1_loop = gopurs_runtime.RecordUpdate2(v1_1, "a", gopurs_runtime.Int((gopurs_runtime.RecordGet(v1_1, "a").IntVal) + (1)), "b", gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(v1_1, "b"), "c", gopurs_runtime.Int((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "c").IntVal) + (2)), "d", gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "e", gopurs_runtime.Int((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "e").IntVal) + (3)), "f", gopurs_runtime.Int((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "f").IntVal) + (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), gopurs_runtime.Int(v_0), gopurs_runtime.Int(5)).IntVal)))))
continue updateRec
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return __t0
}
}


