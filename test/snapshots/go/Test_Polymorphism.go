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

var cache_Test_Polymorphism_mempty___1382332941 gopurs_runtime.Value
var once_Test_Polymorphism_mempty___1382332941 sync.Once

func Get_Test_Polymorphism_mempty___1382332941() gopurs_runtime.Value {
	once_Test_Polymorphism_mempty___1382332941.Do(func() {
		cache_Test_Polymorphism_mempty___1382332941 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Polymorphism_mempty___1382332941(dict_0_box)
		})
	})
	return cache_Test_Polymorphism_mempty___1382332941
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

var cache_Test_Polymorphism_mappend___4026513452 gopurs_runtime.Value
var once_Test_Polymorphism_mappend___4026513452 sync.Once

func Get_Test_Polymorphism_mappend___4026513452() gopurs_runtime.Value {
	once_Test_Polymorphism_mappend___4026513452.Do(func() {
		cache_Test_Polymorphism_mappend___4026513452 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Polymorphism_mappend___4026513452(gopurs_runtime.CoerceToStruct[Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Test_Polymorphism_mappend___4026513452
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

var cache_Test_Polymorphism_polyLoop__3052583336 gopurs_runtime.Value
var once_Test_Polymorphism_polyLoop__3052583336 sync.Once

func Get_Test_Polymorphism_polyLoop__3052583336() gopurs_runtime.Value {
	once_Test_Polymorphism_polyLoop__3052583336.Do(func() {
		cache_Test_Polymorphism_polyLoop__3052583336 = gopurs_runtime.Func2(func(n_init_0_box gopurs_runtime.Value, acc_init_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_Polymorphism_polyLoop__3052583336(n_init_0_box.IntVal, acc_init_1_box.IntVal))
		})
	})
	return cache_Test_Polymorphism_polyLoop__3052583336
}

var cache_Test_Polymorphism_intMonoidish gopurs_runtime.Value
var once_Test_Polymorphism_intMonoidish sync.Once

func Get_Test_Polymorphism_intMonoidish() gopurs_runtime.Value {
	once_Test_Polymorphism_intMonoidish.Do(func() {
		cache_Test_Polymorphism_intMonoidish = gopurs_runtime.Value{Type: 9, IntVal: 459160245, UnsafePtr: unsafe.Pointer((&Constructor_Test_Polymorphism_Monoidish[int64]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((x_0.IntVal) + (y_1.IntVal))
			})
		}), 1}))}
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
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(10000000))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			var Call_local_Test_Polymorphism_go__go_2_2_2 func(gopurs_runtime.Value, int64) int64
			_ = Call_local_Test_Polymorphism_go__go_2_2_2
			var go__go_2_2_2 gopurs_runtime.Value
			_ = go__go_2_2_2
			Call_local_Test_Polymorphism_go__go_2_2_2 = func(v_3_loop gopurs_runtime.Value, v1_4_loop int64) int64 {
			go__go_2_2_2:
				for {
					if false {
						continue go__go_2_2_2
					}
					var v_3 gopurs_runtime.Value = v_3_loop
					_ = v_3
					var v1_4 int64 = v1_4_loop
					_ = v1_4
					var __t3 int64
					{
						if (v_3.IntVal) == (0) {
							__t3 = v1_4
							goto end_branch_3
						} else {

						}
					}
					{
						v_3_loop = gopurs_runtime.Int((v_3.IntVal) - (1))
						v1_4_loop = (v1_4) + (1)
						continue go__go_2_2_2
						__t3 = gopurs_runtime.Value{}.IntVal
					}
				end_branch_3:
					return __t3
				}
			}
			go__go_2_2_2 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Test_Polymorphism_go__go_2_2_2(v_3_loop_val, v1_4_loop_val.IntVal))
				})
			})
			return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_local_Test_Polymorphism_go__go_2_2_2(gopurs_runtime.Int(__local_var_1_1.IntVal), 0))).StrVal())
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
		c := (*Constructor_Test_Polymorphism_Monoidish[any])(ptr)
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

func Call_Test_Polymorphism_mempty___1382332941(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
mempty___1382332941:
	for {
		if false {
			continue mempty___1382332941
		}
		var dict_0 gopurs_runtime.Value = dict_0_loop
		_ = dict_0
		return gopurs_runtime.RecordGet(dict_0, "mempty_")
	}
}

func Call_Test_Polymorphism_mappend_(dict_0_loop *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Test_Polymorphism_mappend___4026513452(dict_0_loop *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
mappend___4026513452:
	for {
		if false {
			continue mappend___4026513452
		}
		var dict_0 *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value] = dict_0_loop
		_ = dict_0
		return gopurs_runtime.Box(dict_0.V0)
	}
}

func Call_Test_Polymorphism_polyLoop(dictMonoidish_0_loop *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value], n_init_1_loop int64, acc_init_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoidish_0 *Constructor_Test_Polymorphism_Monoidish[gopurs_runtime.Value] = dictMonoidish_0_loop
	_ = dictMonoidish_0
	var n_init_1 int64 = n_init_1_loop
	_ = n_init_1
	var acc_init_2 gopurs_runtime.Value = acc_init_2_loop
	_ = acc_init_2
	var Call_local_Test_Polymorphism_go__go_3_0_0 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
	_ = Call_local_Test_Polymorphism_go__go_3_0_0
	var go__go_3_0_0 gopurs_runtime.Value
	_ = go__go_3_0_0
	Call_local_Test_Polymorphism_go__go_3_0_0 = func(v_4_loop gopurs_runtime.Value, v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
	go__go_3_0_0:
		for {
			if false {
				continue go__go_3_0_0
			}
			var v_4 gopurs_runtime.Value = v_4_loop
			_ = v_4
			var v1_5 gopurs_runtime.Value = v1_5_loop
			_ = v1_5
			var __t1 gopurs_runtime.Value
			{
				if (v_4.IntVal) == (0) {
					__t1 = v1_5
					goto end_branch_1
				} else {

				}
			}
			{
				v_4_loop = gopurs_runtime.Int((v_4.IntVal) - (1))
				v1_5_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonoidish_0.V0), v1_5, gopurs_runtime.Box(dictMonoidish_0.V1))
				continue go__go_3_0_0
				__t1 = gopurs_runtime.Value{}
			}
		end_branch_1:
			return __t1
		}
	}
	go__go_3_0_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Test_Polymorphism_go__go_3_0_0(v_4_loop_val, v1_5_loop_val)
		})
	})
	return Call_local_Test_Polymorphism_go__go_3_0_0(gopurs_runtime.Int(n_init_1), acc_init_2)
}

func Call_Test_Polymorphism_polyLoop__3052583336(n_init_0_loop int64, acc_init_1_loop int64) int64 {
polyLoop__3052583336:
	for {
		if false {
			continue polyLoop__3052583336
		}
		var n_init_0 int64 = n_init_0_loop
		_ = n_init_0
		var acc_init_1 int64 = acc_init_1_loop
		_ = acc_init_1
		var Call_local_Test_Polymorphism_go__go_2_0_1 func(gopurs_runtime.Value, int64) int64
		_ = Call_local_Test_Polymorphism_go__go_2_0_1
		var go__go_2_0_1 gopurs_runtime.Value
		_ = go__go_2_0_1
		Call_local_Test_Polymorphism_go__go_2_0_1 = func(v_3_loop gopurs_runtime.Value, v1_4_loop int64) int64 {
		go__go_2_0_1:
			for {
				if false {
					continue go__go_2_0_1
				}
				var v_3 gopurs_runtime.Value = v_3_loop
				_ = v_3
				var v1_4 int64 = v1_4_loop
				_ = v1_4
				var __t1 int64
				{
					if (v_3.IntVal) == (0) {
						__t1 = v1_4
						goto end_branch_1
					} else {

					}
				}
				{
					v_3_loop = gopurs_runtime.Int((v_3.IntVal) - (1))
					v1_4_loop = (v1_4) + (1)
					continue go__go_2_0_1
					__t1 = gopurs_runtime.Value{}.IntVal
				}
			end_branch_1:
				return __t1
			}
		}
		go__go_2_0_1 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(Call_local_Test_Polymorphism_go__go_2_0_1(v_3_loop_val, v1_4_loop_val.IntVal))
			})
		})
		return Call_local_Test_Polymorphism_go__go_2_0_1(gopurs_runtime.Int(n_init_0), acc_init_1)
	}
}
