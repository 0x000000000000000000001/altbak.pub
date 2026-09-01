package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_Records_updateRec gopurs_runtime.Value
var once_Test_Records_updateRec sync.Once

func Get_Test_Records_updateRec() gopurs_runtime.Value {
	once_Test_Records_updateRec.Do(func() {
		cache_Test_Records_updateRec = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Records_updateRec(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Test_Records_updateRec
}

var cache_Test_Records_initial gopurs_runtime.Value
var once_Test_Records_initial sync.Once

func Get_Test_Records_initial() gopurs_runtime.Value {
	once_Test_Records_initial.Do(func() {
		cache_Test_Records_initial = gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Int(0), gopurs_runtime.RecordDict2("c", "d", gopurs_runtime.Int(0), gopurs_runtime.RecordDict2("e", "f", gopurs_runtime.Int(0), gopurs_runtime.Int(0))))
	})
	return cache_Test_Records_initial
}

var cache_Test_Records_describe gopurs_runtime.Value
var once_Test_Records_describe sync.Once

func Get_Test_Records_describe() gopurs_runtime.Value {
	once_Test_Records_describe.Do(func() {
		cache_Test_Records_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Deep Record Updates (10k iterations):"))
	})
	return cache_Test_Records_describe
}

var cache_Test_Records_act gopurs_runtime.Value
var once_Test_Records_act sync.Once

func Get_Test_Records_act() gopurs_runtime.Value {
	once_Test_Records_act.Do(func() {
		cache_Test_Records_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(10000))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(Call_Test_Records_updateRec(__local_var_1_1.IntVal, Get_Test_Records_initial()), "b"), "d"), "f")).StrVal())
		})
	})
	return cache_Test_Records_act
}

func Call_Test_Records_updateRec(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
updateRec:
	for {
		if false {
			continue updateRec
		}
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
			v1_1_loop = gopurs_runtime.RecordUpdate2(v1_1, "a", gopurs_runtime.Int((gopurs_runtime.RecordGet(v1_1, "a").IntVal)+(1)), "b", gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(v1_1, "b"), "c", gopurs_runtime.Int((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "c").IntVal)+(2)), "d", gopurs_runtime.RecordUpdate2(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "e", gopurs_runtime.Int((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "e").IntVal)+(3)), "f", gopurs_runtime.Int((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_1, "b"), "d"), "f").IntVal)+((v_0)%(5))))))
			continue updateRec
			__t0 = gopurs_runtime.Value{}
		}
	end_branch_0:
		return __t0
	}
}
