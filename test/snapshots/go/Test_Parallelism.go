package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_Parallelism_liftEffect gopurs_runtime.Value
var once_Test_Parallelism_liftEffect sync.Once

func Get_Test_Parallelism_liftEffect() gopurs_runtime.Value {
	once_Test_Parallelism_liftEffect.Do(func() {
		cache_Test_Parallelism_liftEffect = Get_Effect_Aff__liftEffect()
	})
	return cache_Test_Parallelism_liftEffect
}

var cache_Test_Parallelism_fib gopurs_runtime.Value
var once_Test_Parallelism_fib sync.Once

func Get_Test_Parallelism_fib() gopurs_runtime.Value {
	once_Test_Parallelism_fib.Do(func() {
		cache_Test_Parallelism_fib = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_Parallelism_fib(v_0_box.IntVal))
		})
	})
	return cache_Test_Parallelism_fib
}

var cache_Test_Parallelism_heavyTask gopurs_runtime.Value
var once_Test_Parallelism_heavyTask sync.Once

func Get_Test_Parallelism_heavyTask() gopurs_runtime.Value {
	once_Test_Parallelism_heavyTask.Do(func() {
		cache_Test_Parallelism_heavyTask = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Parallelism_heavyTask(n_0_box.IntVal)
		})
	})
	return cache_Test_Parallelism_heavyTask
}

var cache_Test_Parallelism_describe gopurs_runtime.Value
var once_Test_Parallelism_describe sync.Once

func Get_Test_Parallelism_describe() gopurs_runtime.Value {
	once_Test_Parallelism_describe.Do(func() {
		cache_Test_Parallelism_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Parallelism (4 x Fib 42)"))
	})
	return cache_Test_Parallelism_describe
}

var cache_Test_Parallelism_act gopurs_runtime.Value
var once_Test_Parallelism_act sync.Once

func Get_Test_Parallelism_act() gopurs_runtime.Value {
	once_Test_Parallelism_act.Do(func() {
		cache_Test_Parallelism_act = gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), f_0, gopurs_runtime.Func(func(f_prime__2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), a_1, gopurs_runtime.Func(func(a_prime__3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(Get_Effect_Aff__pure(), gopurs_runtime.Apply(f_prime__2, a_prime__3))
					}))
				}))
			})
		}), Get_Effect_Aff__map(), Get_Effect_Aff__pure(), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Aff_forkAff(), Call_Test_Parallelism_heavyTask(42))
		}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_replicateImpl(), gopurs_runtime.Int(4), Get_Data_Unit_unit()).UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}())), gopurs_runtime.Func(func(fibers_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), f_1, gopurs_runtime.Func(func(f_prime__3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), a_2, gopurs_runtime.Func(func(a_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(Get_Effect_Aff__pure(), gopurs_runtime.Apply(f_prime__3, a_prime__4))
						}))
					}))
				})
			}), Get_Effect_Aff__map(), Get_Effect_Aff__pure(), Get_Data_Semigroup_concatArray(), Get_Effect_Aff_joinFiber(), func() gopurs_runtime.Value {
				arr := func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(fibers_0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = v
				}
				return gopurs_runtime.Array(boxed)
			}()), gopurs_runtime.Func(func(results_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(("Sum of results: ")+(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(func() gopurs_runtime.Value {
					arr_val_foldlArray8 := func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(results_1.UnsafePtr)
							unboxed := make([]int64, len(arr))
							for i, v := range arr {
								unboxed[i] = v.IntVal
							}
							return unboxed
						}()
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}()
					_ = arr_val_foldlArray8
					res_go_foldlArray8 := gopurs_runtime.Int(0)
					_ = res_go_foldlArray8
					arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
					_ = arr_go_foldlArray8
					for _, v_foldlArray8 := range *arr_go_foldlArray8 {
						res_go_foldlArray8 = gopurs_runtime.Apply2(Get_Data_Semiring_intAdd(), res_go_foldlArray8, v_foldlArray8)
					}
					return res_go_foldlArray8
				}().IntVal)).StrVal()))))
			}))
		}))
	})
	return cache_Test_Parallelism_act
}

func Call_Test_Parallelism_fib(v_0_loop int64) int64 {
fib:
	for {
		if false {
			continue fib
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var __t0 int64
		{
			if (v_0) == (0) {
				__t0 = 0
				goto end_branch_0
			} else {

			}
		}
		{
			if (v_0) == (1) {
				__t0 = 1
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = (Call_Test_Parallelism_fib((v_0) - (1))) + (Call_Test_Parallelism_fib((v_0) - (2)))
		}
	end_branch_0:
		return __t0
	}
}

func Call_Test_Parallelism_heavyTask(n_0_loop int64) gopurs_runtime.Value {
	var n_0 int64 = n_0_loop
	_ = n_0
	return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.UncurriedApp2(Get_Effect_Aff__delay(), Get_Data_Either_Right(), gopurs_runtime.Float(0.0)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Get_Effect_Aff__pure(), gopurs_runtime.Int(Call_Test_Parallelism_fib(n_0)))
	}))
}
