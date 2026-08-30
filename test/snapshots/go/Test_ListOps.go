package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_ListOps_add gopurs_runtime.Value
var once_Test_ListOps_add sync.Once

func Get_Test_ListOps_add() gopurs_runtime.Value {
	once_Test_ListOps_add.Do(func() {
		cache_Test_ListOps_add = Get_Data_Semiring_intAdd()
	})
	return cache_Test_ListOps_add
}

var cache_Test_ListOps_Nil gopurs_runtime.Value
var once_Test_ListOps_Nil sync.Once

func Get_Test_ListOps_Nil() gopurs_runtime.Value {
	once_Test_ListOps_Nil.Do(func() {
		cache_Test_ListOps_Nil = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Test_ListOps_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Test_ListOps_Nil
}

var cache_Test_ListOps_Cons gopurs_runtime.Value
var once_Test_ListOps_Cons sync.Once

func Get_Test_ListOps_Cons() gopurs_runtime.Value {
	once_Test_ListOps_Cons.Do(func() {
		cache_Test_ListOps_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((&Constructor_Test_ListOps_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](value1)}))}
			})
		})
	})
	return cache_Test_ListOps_Cons
}

var cache_Test_ListOps_go__range gopurs_runtime.Value
var once_Test_ListOps_go__range sync.Once

func Get_Test_ListOps_go__range() gopurs_runtime.Value {
	once_Test_ListOps_go__range.Do(func() {
		cache_Test_ListOps_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_Test_ListOps_go__range(start_0_box.IntVal, end_1_box.IntVal))}
		})
	})
	return cache_Test_ListOps_go__range
}

var cache_Test_ListOps_foldl gopurs_runtime.Value
var once_Test_ListOps_foldl sync.Once

func Get_Test_ListOps_foldl() gopurs_runtime.Value {
	once_Test_ListOps_foldl.Do(func() {
		cache_Test_ListOps_foldl = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_ListOps_foldl(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](v2_2_box))
		})
	})
	return cache_Test_ListOps_foldl
}

var cache_Test_ListOps_filterEvens gopurs_runtime.Value
var once_Test_ListOps_filterEvens sync.Once

func Get_Test_ListOps_filterEvens() gopurs_runtime.Value {
	once_Test_ListOps_filterEvens.Do(func() {
		cache_Test_ListOps_filterEvens = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_Test_ListOps_filterEvens(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](lst_0_box)))}
		})
	})
	return cache_Test_ListOps_filterEvens
}

var cache_Test_ListOps_sumEvens gopurs_runtime.Value
var once_Test_ListOps_sumEvens sync.Once

func Get_Test_ListOps_sumEvens() gopurs_runtime.Value {
	once_Test_ListOps_sumEvens.Do(func() {
		cache_Test_ListOps_sumEvens = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_ListOps_sumEvens(n_0_box.IntVal))
		})
	})
	return cache_Test_ListOps_sumEvens
}

var cache_Test_ListOps_describe gopurs_runtime.Value
var once_Test_ListOps_describe sync.Once

func Get_Test_ListOps_describe() gopurs_runtime.Value {
	once_Test_ListOps_describe.Do(func() {
		cache_Test_ListOps_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("List Processing (900 elements):"))
	})
	return cache_Test_ListOps_describe
}

var cache_Test_ListOps_act gopurs_runtime.Value
var once_Test_ListOps_act sync.Once

func Get_Test_ListOps_act() gopurs_runtime.Value {
	once_Test_ListOps_act.Do(func() {
		cache_Test_ListOps_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(900))
			_ = __local_var_0_0
			dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = dummy_1_1
			return gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_ListOps_sumEvens(dummy_1_1.IntVal)))
		})
	})
	return cache_Test_ListOps_act
}

type Constructor_Test_ListOps_Nil[T_a any] struct {
	Rc uint32
}

type Constructor_Test_ListOps_Cons[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Test_ListOps_Cons[gopurs_runtime.Value]
}

func Call_Test_ListOps_go__range(start_0_loop int64, end_1_loop int64) *Constructor_Test_ListOps_Cons[int64] {
	var start_0 int64 = start_0_loop
	_ = start_0
	var end_1 int64 = end_1_loop
	_ = end_1
	var Call_local_Test_ListOps_go__go_2_0_0 func(int64, *Constructor_Test_ListOps_Cons[int64]) gopurs_runtime.Value
	_ = Call_local_Test_ListOps_go__go_2_0_0
	var go__go_2_0_0 gopurs_runtime.Value
	_ = go__go_2_0_0
	Call_local_Test_ListOps_go__go_2_0_0 = func(curr_3_loop int64, acc_4_loop *Constructor_Test_ListOps_Cons[int64]) gopurs_runtime.Value {
	go__go_2_0_0:
		for {
			if false {
				continue go__go_2_0_0
			}
			var curr_3 int64 = curr_3_loop
			_ = curr_3
			var acc_4 *Constructor_Test_ListOps_Cons[int64] = acc_4_loop
			_ = acc_4
			var __t1 *Constructor_Test_ListOps_Cons[int64]
			{
				if (curr_3) < (start_0) {
					__t1 = acc_4
					goto end_branch_1
				} else {

				}
			}
			{
				curr_3_loop = (curr_3) - (1)
				acc_4_loop = (&Constructor_Test_ListOps_Cons[int64]{1, gopurs_runtime.Int(curr_3), gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(acc_4)})})
				continue go__go_2_0_0
				__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](gopurs_runtime.Value{})
			}
		end_branch_1:
			return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(__t1)}
		}
	}
	go__go_2_0_0 = gopurs_runtime.Func(func(curr_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(acc_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Test_ListOps_go__go_2_0_0(curr_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](acc_4_loop_val))
		})
	})
	return gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](Call_local_Test_ListOps_go__go_2_0_0(end_1, (*Constructor_Test_ListOps_Cons[int64])(nil)))
}

func Call_Test_ListOps_foldl(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Test_ListOps_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
foldl:
	for {
		if false {
			continue foldl
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var v2_2 *Constructor_Test_ListOps_Cons[gopurs_runtime.Value] = v2_2_loop
		_ = v2_2
		var __t0 gopurs_runtime.Value
		{
			if v2_2 == nil {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			if v2_2 != nil {
				v_0_loop = v_0
				v1_1_loop = gopurs_runtime.Apply2(v_0, v1_1, (v2_2).V0)
				v2_2_loop = (v2_2).V1
				continue foldl
				__t0 = gopurs_runtime.Value{}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Test_ListOps_filterEvens(lst_0_loop *Constructor_Test_ListOps_Cons[int64]) *Constructor_Test_ListOps_Cons[int64] {
	var lst_0 *Constructor_Test_ListOps_Cons[int64] = lst_0_loop
	_ = lst_0
	var Call_local_Test_ListOps_go__go_1_0_1 func(gopurs_runtime.Value, *Constructor_Test_ListOps_Cons[int64]) gopurs_runtime.Value
	_ = Call_local_Test_ListOps_go__go_1_0_1
	var go__go_1_0_1 gopurs_runtime.Value
	_ = go__go_1_0_1
	Call_local_Test_ListOps_go__go_1_0_1 = func(v_2_loop gopurs_runtime.Value, v1_3_loop *Constructor_Test_ListOps_Cons[int64]) gopurs_runtime.Value {
	go__go_1_0_1:
		for {
			if false {
				continue go__go_1_0_1
			}
			var v_2 gopurs_runtime.Value = v_2_loop
			_ = v_2
			var v1_3 *Constructor_Test_ListOps_Cons[int64] = v1_3_loop
			_ = v1_3
			var __t2 gopurs_runtime.Value
			{
				if v_2.Type == 9 && v_2.IntVal == 1127792131 && v_2.UnsafePtr == nil {
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v1_3)}
					goto end_branch_2
				} else {

				}
			}
			{
				if v_2.Type == 9 && v_2.IntVal == 1127792131 && v_2.UnsafePtr != nil {
					var __t1 *Constructor_Test_ListOps_Cons[int64]
					{
						if (gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int((*Constructor_Test_ListOps_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0.IntVal), gopurs_runtime.Int(2)).IntVal) == (0) {
							v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Test_ListOps_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
							v1_3_loop = (&Constructor_Test_ListOps_Cons[int64]{1, (*Constructor_Test_ListOps_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v1_3)})})
							continue go__go_1_0_1
							__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](gopurs_runtime.Value{})
							goto end_branch_1
						} else {

						}
					}
					{
						v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Test_ListOps_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
						v1_3_loop = v1_3
						continue go__go_1_0_1
						__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](gopurs_runtime.Value{})
					}
				end_branch_1:
					__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(__t1)}
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](__t2))}
		}
	}
	go__go_1_0_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Test_ListOps_go__go_1_0_1(v_2_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](v1_3_loop_val))
		})
	})
	return gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](Call_local_Test_ListOps_go__go_1_0_1(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(lst_0)}, (*Constructor_Test_ListOps_Cons[int64])(nil)))
}

func Call_Test_ListOps_sumEvens(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	var Call_local_Test_ListOps_go__go_1_0_2 func(int64, *Constructor_Test_ListOps_Cons[int64]) gopurs_runtime.Value
	_ = Call_local_Test_ListOps_go__go_1_0_2
	var go__go_1_0_2 gopurs_runtime.Value
	_ = go__go_1_0_2
	Call_local_Test_ListOps_go__go_1_0_2 = func(curr_2_loop int64, acc_3_loop *Constructor_Test_ListOps_Cons[int64]) gopurs_runtime.Value {
	go__go_1_0_2:
		for {
			if false {
				continue go__go_1_0_2
			}
			var curr_2 int64 = curr_2_loop
			_ = curr_2
			var acc_3 *Constructor_Test_ListOps_Cons[int64] = acc_3_loop
			_ = acc_3
			var __t1 *Constructor_Test_ListOps_Cons[int64]
			{
				if (curr_2) < (1) {
					__t1 = acc_3
					goto end_branch_1
				} else {

				}
			}
			{
				curr_2_loop = (curr_2) - (1)
				acc_3_loop = (&Constructor_Test_ListOps_Cons[int64]{1, gopurs_runtime.Int(curr_2), gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(acc_3)})})
				continue go__go_1_0_2
				__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](gopurs_runtime.Value{})
			}
		end_branch_1:
			return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(__t1)}
		}
	}
	go__go_1_0_2 = gopurs_runtime.Func(func(curr_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(acc_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Test_ListOps_go__go_1_0_2(curr_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](acc_3_loop_val))
		})
	})
	return Call_Test_ListOps_foldl(Get_Data_Semiring_intAdd(), gopurs_runtime.Int(0), gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_Test_ListOps_filterEvens(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](Call_local_Test_ListOps_go__go_1_0_2(n_0, (*Constructor_Test_ListOps_Cons[int64])(nil)))))})).IntVal
}
