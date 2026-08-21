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
		cache_Test_Polymorphism_logShow = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Polymorphism_logShow(a_0_box.IntVal)
		})
	})
	return cache_Test_Polymorphism_logShow
}

var cache_Test_Polymorphism_Monoidish_dollar_Dict gopurs_runtime.Value
var once_Test_Polymorphism_Monoidish_dollar_Dict sync.Once

func Get_Test_Polymorphism_Monoidish_dollar_Dict() gopurs_runtime.Value {
	once_Test_Polymorphism_Monoidish_dollar_Dict.Do(func() {
		cache_Test_Polymorphism_Monoidish_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Polymorphism_Monoidish_dollar_Dict(x_0_box)
		})
	})
	return cache_Test_Polymorphism_Monoidish_dollar_Dict
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

var cache_Test_Polymorphism_polyLoopGo gopurs_runtime.Value
var once_Test_Polymorphism_polyLoopGo sync.Once

func Get_Test_Polymorphism_polyLoopGo() gopurs_runtime.Value {
	once_Test_Polymorphism_polyLoopGo.Do(func() {
		cache_Test_Polymorphism_polyLoopGo = gopurs_runtime.Func3(func(dictMonoidish_0_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, v1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Polymorphism_polyLoopGo(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish](dictMonoidish_0_box), v_2_box.IntVal, v1_3_box)
		})
	})
	return cache_Test_Polymorphism_polyLoopGo
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
		cache_Test_Polymorphism_intMonoidish = gopurs_runtime.Value{Type: 9, IntVal: 459160245, UnsafePtr: unsafe.Pointer((&Constructor_Test_Polymorphism_Monoidish{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((x_0.IntVal) + (y_1.IntVal))
			})
		}), gopurs_runtime.Int(1)}))}
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
		cache_Test_Polymorphism_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(10000000))
			_ = __local_var_0_0
			dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = dummy_1_1
			var __t4 gopurs_runtime.Value
			{
				if (dummy_1_1.IntVal) == (0) {
					__t4 = gopurs_runtime.Int(0)
					goto end_branch_4
				} else {

				}
			}
			{
				__local_var_2_2 := (dummy_1_1.IntVal) - (1)
				_ = __local_var_2_2
				var __t3 gopurs_runtime.Value
				{
					if (__local_var_2_2) == (0) {
						__t3 = gopurs_runtime.Int(1)
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = Call_Test_Polymorphism_polyLoopGo(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish](gopurs_runtime.Value{Type: 9, IntVal: 459160245, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish](Get_Test_Polymorphism_intMonoidish()))}), (__local_var_2_2)-(1), gopurs_runtime.Int(2))
				}
			end_branch_3:
				__t4 = __t3
			}
		end_branch_4:
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(__t4.IntVal)).StrVal())), gopurs_runtime.Value{})
		})
	})
	return cache_Test_Polymorphism_act
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

var cache_Test_Polymorphism_polyLoopGo__2675791634 gopurs_runtime.Value
var once_Test_Polymorphism_polyLoopGo__2675791634 sync.Once

func Get_Test_Polymorphism_polyLoopGo__2675791634() gopurs_runtime.Value {
	once_Test_Polymorphism_polyLoopGo__2675791634.Do(func() {
		cache_Test_Polymorphism_polyLoopGo__2675791634 = gopurs_runtime.Func3(func(dictMonoidish_0_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, v1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Polymorphism_polyLoopGo__2675791634(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish](dictMonoidish_0_box), v_2_box.IntVal, v1_3_box)
		})
	})
	return cache_Test_Polymorphism_polyLoopGo__2675791634
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
		case "mappend_":
			return gopurs_runtime.Box(c.V0)
		case "mempty_":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Test_Polymorphism_Monoidish: " + key)
		}
	}
}

func Call_Test_Polymorphism_logShow(a_0_loop int64) gopurs_runtime.Value {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(a_0)).StrVal()))
}

func Call_Test_Polymorphism_Monoidish_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Test_Polymorphism_polyLoopGo(dictMonoidish_0_loop *Constructor_Test_Polymorphism_Monoidish, v_2_loop int64, v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
polyLoopGo:
	for {
		if false {
			continue polyLoopGo
		}
		var dictMonoidish_0 *Constructor_Test_Polymorphism_Monoidish = dictMonoidish_0_loop
		_ = dictMonoidish_0
		var v_2 int64 = v_2_loop
		_ = v_2
		var v1_3 gopurs_runtime.Value = v1_3_loop
		_ = v1_3
		mempty_1_1_0 := gopurs_runtime.Box(dictMonoidish_0.V1)
		_ = mempty_1_1_0
		var __t1 gopurs_runtime.Value
		{
			if (v_2) == (0) {
				__t1 = v1_3
				goto end_branch_1
			} else {

			}
		}
		{
			dictMonoidish_0_loop = dictMonoidish_0
			v_2_loop = (v_2) - (1)
			v1_3_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), v1_3, mempty_1_1_0)
			continue polyLoopGo
			__t1 = gopurs_runtime.Value{}
		}
	end_branch_1:
		return __t1
	}
}

func Call_Test_Polymorphism_polyLoop(dictMonoidish_0_loop *Constructor_Test_Polymorphism_Monoidish, n_init_1_loop int64, acc_init_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoidish_0 *Constructor_Test_Polymorphism_Monoidish = dictMonoidish_0_loop
	_ = dictMonoidish_0
	var n_init_1 int64 = n_init_1_loop
	_ = n_init_1
	var acc_init_2 gopurs_runtime.Value = acc_init_2_loop
	_ = acc_init_2
	var __t0 gopurs_runtime.Value
	{
		if (n_init_1) == (0) {
			__t0 = acc_init_2
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = Call_Test_Polymorphism_polyLoopGo(dictMonoidish_0, (n_init_1)-(1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), acc_init_2, gopurs_runtime.Box(dictMonoidish_0.V1)))
	}
end_branch_0:
	return __t0
}

func Call_Test_Polymorphism_mappend___3566619927(dict_0_loop *Constructor_Test_Polymorphism_Monoidish) gopurs_runtime.Value {
	var dict_0 *Constructor_Test_Polymorphism_Monoidish = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Test_Polymorphism_polyLoop__2675791634(dictMonoidish_0_loop *Constructor_Test_Polymorphism_Monoidish, n_init_1_loop int64, acc_init_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoidish_0 *Constructor_Test_Polymorphism_Monoidish = dictMonoidish_0_loop
	_ = dictMonoidish_0
	var n_init_1 int64 = n_init_1_loop
	_ = n_init_1
	var acc_init_2 gopurs_runtime.Value = acc_init_2_loop
	_ = acc_init_2
	var __t0 gopurs_runtime.Value
	{
		if (n_init_1) == (0) {
			__t0 = acc_init_2
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = Call_Test_Polymorphism_polyLoopGo(dictMonoidish_0, (n_init_1)-(1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), acc_init_2, gopurs_runtime.Box(dictMonoidish_0.V1)))
	}
end_branch_0:
	return __t0
}

func Call_Test_Polymorphism_polyLoopGo__2675791634(dictMonoidish_0_loop *Constructor_Test_Polymorphism_Monoidish, v_2_loop int64, v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoidish_0 *Constructor_Test_Polymorphism_Monoidish = dictMonoidish_0_loop
	_ = dictMonoidish_0
	var v_2 int64 = v_2_loop
	_ = v_2
	var v1_3 gopurs_runtime.Value = v1_3_loop
	_ = v1_3
	mempty_1_1_0 := gopurs_runtime.Box(dictMonoidish_0.V1)
	_ = mempty_1_1_0
	var __t4 gopurs_runtime.Value
	{
		if (v_2) == (0) {
			__t4 = v1_3
			goto end_branch_4
		} else {

		}
	}
	{
		__local_var_4_1 := (v_2) - (1)
		_ = __local_var_4_1
		__local_var_5_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), v1_3, mempty_1_1_0)
		_ = __local_var_5_2
		var __t3 gopurs_runtime.Value
		{
			if (__local_var_4_1) == (0) {
				__t3 = __local_var_5_2
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = Call_Test_Polymorphism_polyLoopGo(dictMonoidish_0, (__local_var_4_1)-(1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), __local_var_5_2, gopurs_runtime.Box(dictMonoidish_0.V1)))
		}
	end_branch_3:
		__t4 = __t3
	}
end_branch_4:
	return __t4
}
