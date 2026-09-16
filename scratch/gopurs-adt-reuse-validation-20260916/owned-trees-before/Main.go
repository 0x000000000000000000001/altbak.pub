package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Crimson gopurs_runtime.Value
var once_Main_Crimson sync.Once

func Get_Main_Crimson() gopurs_runtime.Value {
	once_Main_Crimson.Do(func() {
		cache_Main_Crimson = gopurs_runtime.Value{Type: 9, IntVal: int64(2247809753), UnsafePtr: nil}
	})
	return cache_Main_Crimson
}

var cache_Main_Onyx gopurs_runtime.Value
var once_Main_Onyx sync.Once

func Get_Main_Onyx() gopurs_runtime.Value {
	once_Main_Onyx.Do(func() {
		cache_Main_Onyx = gopurs_runtime.Value{Type: 9, IntVal: int64(1685833310), UnsafePtr: nil}
	})
	return cache_Main_Onyx
}

var cache_Main_Tip gopurs_runtime.Value
var once_Main_Tip sync.Once

func Get_Main_Tip() gopurs_runtime.Value {
	once_Main_Tip.Do(func() {
		cache_Main_Tip = gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer((*Constructor_Main_Branch)(nil))}
	})
	return cache_Main_Tip
}

var cache_Main_Branch gopurs_runtime.Value
var once_Main_Branch sync.Once

func Get_Main_Branch() gopurs_runtime.Value {
	once_Main_Branch.Do(func() {
		cache_Main_Branch = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer((&Constructor_Main_Branch{1, uint32(value0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](value1), value2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](value3)}))}
					})
				})
			})
		})
	})
	return cache_Main_Branch
}

var cache_Main_render gopurs_runtime.Value
var once_Main_render sync.Once

func Get_Main_render() gopurs_runtime.Value {
	once_Main_render.Do(func() {
		cache_Main_render = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v_0_box)))
		})
	})
	return cache_Main_render
}

var cache_Main_rebalance gopurs_runtime.Value
var once_Main_rebalance sync.Once

func Get_Main_rebalance() gopurs_runtime.Value {
	once_Main_rebalance.Do(func() {
		cache_Main_rebalance = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_rebalance(uint32(v_0_box.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v1_1_box), v2_2_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v3_3_box)))}
		})
	})
	return cache_Main_rebalance
}

var cache_Main_leftChild gopurs_runtime.Value
var once_Main_leftChild sync.Once

func Get_Main_leftChild() gopurs_runtime.Value {
	once_Main_leftChild.Do(func() {
		cache_Main_leftChild = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_leftChild(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v_0_box)))}
		})
	})
	return cache_Main_leftChild
}

var cache_Main_isRed gopurs_runtime.Value
var once_Main_isRed sync.Once

func Get_Main_isRed() gopurs_runtime.Value {
	once_Main_isRed.Do(func() {
		cache_Main_isRed = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_isRed(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v_0_box)))
		})
	})
	return cache_Main_isRed
}

var cache_Main_descend gopurs_runtime.Value
var once_Main_descend sync.Once

func Get_Main_descend() gopurs_runtime.Value {
	once_Main_descend.Do(func() {
		cache_Main_descend = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_descend(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v1_1_box)))}
		})
	})
	return cache_Main_descend
}

var cache_Main_blacken gopurs_runtime.Value
var once_Main_blacken sync.Once

func Get_Main_blacken() gopurs_runtime.Value {
	once_Main_blacken.Do(func() {
		cache_Main_blacken = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_blacken(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v_0_box)))}
		})
	})
	return cache_Main_blacken
}

var cache_Main_put gopurs_runtime.Value
var once_Main_put sync.Once

func Get_Main_put() gopurs_runtime.Value {
	once_Main_put.Do(func() {
		cache_Main_put = gopurs_runtime.Func2(func(value_0_box gopurs_runtime.Value, tree_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_put(value_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](tree_1_box)))}
		})
	})
	return cache_Main_put
}

var cache_Main_build gopurs_runtime.Value
var once_Main_build sync.Once

func Get_Main_build() gopurs_runtime.Value {
	once_Main_build.Do(func() {
		cache_Main_build = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_build(v_0_box.IntVal, v1_1_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v2_2_box)))}
		})
	})
	return cache_Main_build
}

var cache_Main_graftAndPut gopurs_runtime.Value
var once_Main_graftAndPut sync.Once

func Get_Main_graftAndPut() gopurs_runtime.Value {
	once_Main_graftAndPut.Do(func() {
		cache_Main_graftAndPut = gopurs_runtime.Func(func(child_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_graftAndPut(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](child_0_box)))}
		})
	})
	return cache_Main_graftAndPut
}

var cache_Main_mixed gopurs_runtime.Value
var once_Main_mixed sync.Once

func Get_Main_mixed() gopurs_runtime.Value {
	once_Main_mixed.Do(func() {
		cache_Main_mixed = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_mixed(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v1_1_box)))}
		})
	})
	return cache_Main_mixed
}

var cache_Main_audit gopurs_runtime.Value
var once_Main_audit sync.Once

func Get_Main_audit() gopurs_runtime.Value {
	once_Main_audit.Do(func() {
		cache_Main_audit = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_audit(v_0_box.IntVal, v1_1_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v2_2_box))
				_ = orig
				return gopurs_runtime.RecordDict3("blackHeight", "size", "valid", gopurs_runtime.Int(orig.blackHeight), gopurs_runtime.Int(orig.size), gopurs_runtime.Bool(orig.valid))
			}()
		})
	})
	return cache_Main_audit
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func2(func(expectedSize_0_box gopurs_runtime.Value, tree_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(expectedSize_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](tree_1_box))
		})
	})
	return cache_Main_check
}

var cache_Main_snapshots gopurs_runtime.Value
var once_Main_snapshots sync.Once

func Get_Main_snapshots() gopurs_runtime.Value {
	once_Main_snapshots.Do(func() {
		cache_Main_snapshots = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_snapshots(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v1_1_box))
		})
	})
	return cache_Main_snapshots
}

var cache_Main_three gopurs_runtime.Value
var once_Main_three sync.Once

func Get_Main_three() gopurs_runtime.Value {
	once_Main_three.Do(func() {
		cache_Main_three = gopurs_runtime.Func3(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_three(a_0_box.IntVal, b_1_box.IntVal, c_2_box.IntVal)
		})
	})
	return cache_Main_three
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_three(int64(3), int64(2), int64(1)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_three(int64(3), int64(1), int64(2)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_three(int64(1), int64(3), int64(2)), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_three(int64(1), int64(2), int64(3)), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=Any
							__local_var_4_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(127)))
							_ = __local_var_4_0
							__local_var_5_1 := gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Value{})
							_ = __local_var_5_1
							__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_1), gopurs_runtime.Value{})
							_ = __local_var_6_2
							__local_var_7_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_build(__local_var_6_2.IntVal, int64(-1), (*Constructor_Main_Branch)(nil)))}), gopurs_runtime.Value{})
							_ = __local_var_7_3
							__local_var_8_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_build(__local_var_6_2.IntVal, int64(1), (*Constructor_Main_Branch)(nil)))}), gopurs_runtime.Value{})
							_ = __local_var_8_4
							__local_var_9_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_mixed(__local_var_6_2.IntVal, (*Constructor_Main_Branch)(nil)))}), gopurs_runtime.Value{})
							_ = __local_var_9_5
							__local_var_10_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_7_3), gopurs_runtime.Value{})
							_ = __local_var_10_6
							__local_var_11_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_8_4), gopurs_runtime.Value{})
							_ = __local_var_11_7
							__local_var_12_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_9_5), gopurs_runtime.Value{})
							_ = __local_var_12_8
							return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(__local_var_6_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_10_6)), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(__local_var_6_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_11_7)), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(__local_var_6_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_12_8)), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											__local_var_16_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_12_8)))), gopurs_runtime.Value{})
											_ = __local_var_16_9
											__local_var_17_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_16_9), gopurs_runtime.Value{})
											_ = __local_var_17_10
											__local_var_18_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_mixed(__local_var_6_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_12_8)))}), gopurs_runtime.Value{})
											_ = __local_var_18_11
											__local_var_19_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_18_11), gopurs_runtime.Value{})
											_ = __local_var_19_12
											return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(__local_var_6_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_19_12)), gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
													actual   string
													expected string
												}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_19_12)), __local_var_17_10.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
														actual   string
														expected string
													}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_12_8)), __local_var_17_10.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_snapshots(int64(32), (*Constructor_Main_Branch)(nil)), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																// TAST (Let): __local_var_24_13 shape=App(Var) bindingType=Any
																__local_var_24_13 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer((&Constructor_Main_Branch{1, 1685833310, (&Constructor_Main_Branch{1, 1685833310, (*Constructor_Main_Branch)(nil), int64(10), (*Constructor_Main_Branch)(nil)}), int64(50), (&Constructor_Main_Branch{1, 1685833310, (*Constructor_Main_Branch)(nil), int64(90), (*Constructor_Main_Branch)(nil)})}))})
																_ = __local_var_24_13
																__local_var_25_14 := gopurs_runtime.Apply(__local_var_24_13, gopurs_runtime.Value{})
																_ = __local_var_25_14
																__local_var_26_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_25_14), gopurs_runtime.Value{})
																_ = __local_var_26_15
																var __t19 *Constructor_Main_Branch
																{
																	var __t_tag_17 *Constructor_Main_Branch = gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)
																	_ = __t_tag_17
																	if __t_tag_17 == nil {
																		__t19 = (*Constructor_Main_Branch)(nil)
																		goto end_branch_19
																	} else {

																	}
																}
																{
																	var __t_tag_18 *Constructor_Main_Branch = gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)
																	_ = __t_tag_18
																	if __t_tag_18 != nil {
																		__t19 = (*Constructor_Main_Branch)(__local_var_26_15.UnsafePtr).V1
																		goto end_branch_19
																	} else {

																	}
																}
																{
																	__t19 = func() *Constructor_Main_Branch { panic("Failed pattern match") }()
																}
															end_branch_19:
																__local_var_27_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(__t19)}), gopurs_runtime.Value{})
																_ = __local_var_27_16
																__local_var_28_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_27_16), gopurs_runtime.Value{})
																_ = __local_var_28_20
																__local_var_29_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)))), gopurs_runtime.Value{})
																_ = __local_var_29_21
																__local_var_30_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_28_20)))), gopurs_runtime.Value{})
																_ = __local_var_30_22
																__local_var_31_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_29_21), gopurs_runtime.Value{})
																_ = __local_var_31_23
																__local_var_32_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_30_22), gopurs_runtime.Value{})
																_ = __local_var_32_24
																__local_var_33_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_put(int64(5), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)))}), gopurs_runtime.Value{})
																_ = __local_var_33_25
																__local_var_34_26 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_33_25), gopurs_runtime.Value{})
																_ = __local_var_34_26
																return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(int64(4), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_34_26)), gopurs_runtime.Func(func(_dollar___unused_35 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																		actual   string
																		expected string
																	}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)), __local_var_31_23.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_36 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																			actual   string
																			expected string
																		}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_28_20)), __local_var_32_24.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_37 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																				__local_var_38_27 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_graftAndPut(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_28_20)))}), gopurs_runtime.Value{})
																				_ = __local_var_38_27
																				__local_var_39_28 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_38_27), gopurs_runtime.Value{})
																				_ = __local_var_39_28
																				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(int64(4), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_39_28)), gopurs_runtime.Func(func(_dollar___unused_40 gopurs_runtime.Value) gopurs_runtime.Value {
																					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																						actual   string
																						expected string
																					}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_28_20)), __local_var_32_24.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_41 gopurs_runtime.Value) gopurs_runtime.Value {
																						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																							actual   string
																							expected string
																						}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)), __local_var_31_23.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_42 gopurs_runtime.Value) gopurs_runtime.Value {
																							return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
																						}))
																					}))
																				})), gopurs_runtime.Value{})
																			})
																		}))
																	}))
																})), gopurs_runtime.Value{})
															})
														}))
													}))
												}))
											})), gopurs_runtime.Value{})
										})
									}))
								}))
							})), gopurs_runtime.Value{})
						})
					}))
				}))
			}))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_Crimson struct {
	Rc uint32
}

type Constructor_Main_Onyx struct {
	Rc uint32
}

type Constructor_Main_Tip struct {
	Rc uint32
}

type Constructor_Main_Branch struct {
	Rc uint32
	V0 uint32
	V1 *Constructor_Main_Branch
	V2 int64
	V3 *Constructor_Main_Branch
}

func Call_Main_render(v_0_loop *Constructor_Main_Branch) string {
render:
	for {
		if false {
			continue render
		}
		var v_0 *Constructor_Main_Branch = v_0_loop
		_ = v_0
		var __t3 string
		{
			if v_0 == nil {
				__t3 = "E"
				goto end_branch_3
			} else {

			}
		}
		{
			if v_0 != nil {
				var __t2 string
				{
					var __t_tag_0 uint32 = (v_0).V0
					_ = __t_tag_0
					if uint32(__t_tag_0) == 2247809753 {
						__t2 = "R"
						goto end_branch_2
					} else {

					}
				}
				{
					var __t_tag_1 uint32 = (v_0).V0
					_ = __t_tag_1
					if uint32(__t_tag_1) == 1685833310 {
						__t2 = "B"
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
				}
			end_branch_2:
				__t3 = ((((((__t2) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((v_0).V2)).StrVal())) + ("(")) + (Call_Main_render((v_0).V1))) + (")(")) + (Call_Main_render((v_0).V3))) + (")")
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = func() string { panic("Failed pattern match") }()
		}
	end_branch_3:
		return __t3
	}
}

func Call_Main_rebalance(v_0_loop uint32, v1_1_loop *Constructor_Main_Branch, v2_2_loop int64, v3_3_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var v1_1 *Constructor_Main_Branch = v1_1_loop
	_ = v1_1
	var v2_2 int64 = v2_2_loop
	_ = v2_2
	var v3_3 *Constructor_Main_Branch = v3_3_loop
	_ = v3_3
	var __t308 *Constructor_Main_Branch
	{
		if v_0 == 1685833310 {
			var __t307 *Constructor_Main_Branch
			{
				if v1_1 != nil {
					var __t265 *Constructor_Main_Branch
					{
						var __t_tag_12 uint32 = (v1_1).V0
						_ = __t_tag_12
						if uint32(__t_tag_12) == 2247809753 {
							var __t223 *Constructor_Main_Branch
							{
								var __t_tag_17 *Constructor_Main_Branch = (v1_1).V1
								_ = __t_tag_17
								if __t_tag_17 != nil {
									var __t126 *Constructor_Main_Branch
									{
										var __t_tag_22 uint32 = ((v1_1).V1).V0
										_ = __t_tag_22
										if uint32(__t_tag_22) == 2247809753 {
											// TAST (Let): __local_var_4_23 shape=Other bindingType=Any
											__local_var_4_23 := ((v1_1).V1).V1
											_ = __local_var_4_23
											// TAST (Let): __local_var_5_24 shape=Other bindingType=Any
											__local_var_5_24 := ((v1_1).V1).V3
											_ = __local_var_5_24
											// TAST (Let): __local_var_6_25 shape=Other bindingType=Any
											__local_var_6_25 := (v1_1).V3
											_ = __local_var_6_25
											// TAST (Let): __local_var_7_26 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
											__local_var_7_26 := v3_3
											_ = __local_var_7_26
											// TAST (Let): __local_var_8_27 shape=Other bindingType=Any
											__local_var_8_27 := ((v1_1).V1).V2
											_ = __local_var_8_27
											// TAST (Let): __local_var_9_28 shape=Other bindingType=Any
											__local_var_9_28 := (v1_1).V2
											_ = __local_var_9_28
											// TAST (Let): __local_var_10_29 shape=Other bindingType=Int
											__local_var_10_29 := v2_2
											_ = __local_var_10_29
											__t126 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_23, __local_var_8_27, __local_var_5_24}), __local_var_9_28, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_25, __local_var_10_29, __local_var_7_26})})
											goto end_branch_126
										} else {

										}
									}
									{
										var __t_tag_30 *Constructor_Main_Branch = (v1_1).V3
										_ = __t_tag_30
										if __t_tag_30 != nil {
											var __t84 *Constructor_Main_Branch
											{
												var __t_tag_35 uint32 = ((v1_1).V3).V0
												_ = __t_tag_35
												if uint32(__t_tag_35) == 2247809753 {
													// TAST (Let): __local_var_4_36 shape=Other bindingType=Any
													__local_var_4_36 := (v1_1).V1
													_ = __local_var_4_36
													// TAST (Let): __local_var_5_37 shape=Other bindingType=Any
													__local_var_5_37 := ((v1_1).V3).V1
													_ = __local_var_5_37
													// TAST (Let): __local_var_6_38 shape=Other bindingType=Any
													__local_var_6_38 := ((v1_1).V3).V3
													_ = __local_var_6_38
													// TAST (Let): __local_var_7_39 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
													__local_var_7_39 := v3_3
													_ = __local_var_7_39
													// TAST (Let): __local_var_8_40 shape=Other bindingType=Any
													__local_var_8_40 := (v1_1).V2
													_ = __local_var_8_40
													// TAST (Let): __local_var_9_41 shape=Other bindingType=Any
													__local_var_9_41 := ((v1_1).V3).V2
													_ = __local_var_9_41
													// TAST (Let): __local_var_10_42 shape=Other bindingType=Int
													__local_var_10_42 := v2_2
													_ = __local_var_10_42
													__t84 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_36, __local_var_8_40, __local_var_5_37}), __local_var_9_41, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_38, __local_var_10_42, __local_var_7_39})})
													goto end_branch_84
												} else {

												}
											}
											{
												var __t_and_44 bool = false
												if v3_3 != nil {

													var __t_tag_43 uint32 = (v3_3).V0
													_ = __t_tag_43
													__t_and_44 = (uint32(__t_tag_43) == 2247809753)
												}
												if __t_and_44 {
													var __t83 *Constructor_Main_Branch
													{
														var __t_tag_49 *Constructor_Main_Branch = (v3_3).V1
														_ = __t_tag_49
														if __t_tag_49 != nil {
															var __t72 *Constructor_Main_Branch
															{
																var __t_tag_54 uint32 = ((v3_3).V1).V0
																_ = __t_tag_54
																if uint32(__t_tag_54) == 2247809753 {
																	// TAST (Let): __local_var_4_55 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
																	__local_var_4_55 := v1_1
																	_ = __local_var_4_55
																	// TAST (Let): __local_var_5_56 shape=Other bindingType=Any
																	__local_var_5_56 := ((v3_3).V1).V1
																	_ = __local_var_5_56
																	// TAST (Let): __local_var_6_57 shape=Other bindingType=Any
																	__local_var_6_57 := ((v3_3).V1).V3
																	_ = __local_var_6_57
																	// TAST (Let): __local_var_7_58 shape=Other bindingType=Any
																	__local_var_7_58 := (v3_3).V3
																	_ = __local_var_7_58
																	// TAST (Let): __local_var_8_59 shape=Other bindingType=Int
																	__local_var_8_59 := v2_2
																	_ = __local_var_8_59
																	// TAST (Let): __local_var_9_60 shape=Other bindingType=Any
																	__local_var_9_60 := ((v3_3).V1).V2
																	_ = __local_var_9_60
																	// TAST (Let): __local_var_10_61 shape=Other bindingType=Any
																	__local_var_10_61 := (v3_3).V2
																	_ = __local_var_10_61
																	__t72 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_55, __local_var_8_59, __local_var_5_56}), __local_var_9_60, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_57, __local_var_10_61, __local_var_7_58})})
																	goto end_branch_72
																} else {

																}
															}
															{
																var __t_tag_62 *Constructor_Main_Branch = (v3_3).V3
																_ = __t_tag_62
																var __t_and_64 bool = false
																if __t_tag_62 != nil {

																	var __t_tag_63 uint32 = ((v3_3).V3).V0
																	_ = __t_tag_63
																	__t_and_64 = (uint32(__t_tag_63) == 2247809753)
																}
																if __t_and_64 {
																	// TAST (Let): __local_var_4_65 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
																	__local_var_4_65 := v1_1
																	_ = __local_var_4_65
																	// TAST (Let): __local_var_5_66 shape=Other bindingType=Any
																	__local_var_5_66 := (v3_3).V1
																	_ = __local_var_5_66
																	// TAST (Let): __local_var_6_67 shape=Other bindingType=Any
																	__local_var_6_67 := ((v3_3).V3).V1
																	_ = __local_var_6_67
																	// TAST (Let): __local_var_7_68 shape=Other bindingType=Any
																	__local_var_7_68 := ((v3_3).V3).V3
																	_ = __local_var_7_68
																	// TAST (Let): __local_var_8_69 shape=Other bindingType=Int
																	__local_var_8_69 := v2_2
																	_ = __local_var_8_69
																	// TAST (Let): __local_var_9_70 shape=Other bindingType=Any
																	__local_var_9_70 := (v3_3).V2
																	_ = __local_var_9_70
																	// TAST (Let): __local_var_10_71 shape=Other bindingType=Any
																	__local_var_10_71 := ((v3_3).V3).V2
																	_ = __local_var_10_71
																	__t72 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_65, __local_var_8_69, __local_var_5_66}), __local_var_9_70, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_67, __local_var_10_71, __local_var_7_68})})
																	goto end_branch_72
																} else {

																}
															}
															{
																// TAST (Let): __local_var_4_50 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
																__local_var_4_50 := v1_1
																_ = __local_var_4_50
																// TAST (Let): __local_var_5_51 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
																__local_var_5_51 := v3_3
																_ = __local_var_5_51
																// TAST (Let): __local_var_6_52 shape=Other bindingType=(ADT ["Main","Paint"] [])
																__local_var_6_52 := v_0
																_ = __local_var_6_52
																// TAST (Let): __local_var_7_53 shape=Other bindingType=Int
																__local_var_7_53 := v2_2
																_ = __local_var_7_53
																__t72 = (&Constructor_Main_Branch{1, __local_var_6_52, __local_var_4_50, __local_var_7_53, __local_var_5_51})
															}
														end_branch_72:
															__t83 = __t72
															goto end_branch_83
														} else {

														}
													}
													{
														var __t_tag_73 *Constructor_Main_Branch = (v3_3).V3
														_ = __t_tag_73
														var __t_and_75 bool = false
														if __t_tag_73 != nil {

															var __t_tag_74 uint32 = ((v3_3).V3).V0
															_ = __t_tag_74
															__t_and_75 = (uint32(__t_tag_74) == 2247809753)
														}
														if __t_and_75 {
															// TAST (Let): __local_var_4_76 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
															__local_var_4_76 := v1_1
															_ = __local_var_4_76
															// TAST (Let): __local_var_5_77 shape=Other bindingType=Any
															__local_var_5_77 := (v3_3).V1
															_ = __local_var_5_77
															// TAST (Let): __local_var_6_78 shape=Other bindingType=Any
															__local_var_6_78 := ((v3_3).V3).V1
															_ = __local_var_6_78
															// TAST (Let): __local_var_7_79 shape=Other bindingType=Any
															__local_var_7_79 := ((v3_3).V3).V3
															_ = __local_var_7_79
															// TAST (Let): __local_var_8_80 shape=Other bindingType=Int
															__local_var_8_80 := v2_2
															_ = __local_var_8_80
															// TAST (Let): __local_var_9_81 shape=Other bindingType=Any
															__local_var_9_81 := (v3_3).V2
															_ = __local_var_9_81
															// TAST (Let): __local_var_10_82 shape=Other bindingType=Any
															__local_var_10_82 := ((v3_3).V3).V2
															_ = __local_var_10_82
															__t83 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_76, __local_var_8_80, __local_var_5_77}), __local_var_9_81, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_78, __local_var_10_82, __local_var_7_79})})
															goto end_branch_83
														} else {

														}
													}
													{
														// TAST (Let): __local_var_4_45 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_4_45 := v1_1
														_ = __local_var_4_45
														// TAST (Let): __local_var_5_46 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_5_46 := v3_3
														_ = __local_var_5_46
														// TAST (Let): __local_var_6_47 shape=Other bindingType=(ADT ["Main","Paint"] [])
														__local_var_6_47 := v_0
														_ = __local_var_6_47
														// TAST (Let): __local_var_7_48 shape=Other bindingType=Int
														__local_var_7_48 := v2_2
														_ = __local_var_7_48
														__t83 = (&Constructor_Main_Branch{1, __local_var_6_47, __local_var_4_45, __local_var_7_48, __local_var_5_46})
													}
												end_branch_83:
													__t84 = __t83
													goto end_branch_84
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_31 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_4_31 := v1_1
												_ = __local_var_4_31
												// TAST (Let): __local_var_5_32 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_5_32 := v3_3
												_ = __local_var_5_32
												// TAST (Let): __local_var_6_33 shape=Other bindingType=(ADT ["Main","Paint"] [])
												__local_var_6_33 := v_0
												_ = __local_var_6_33
												// TAST (Let): __local_var_7_34 shape=Other bindingType=Int
												__local_var_7_34 := v2_2
												_ = __local_var_7_34
												__t84 = (&Constructor_Main_Branch{1, __local_var_6_33, __local_var_4_31, __local_var_7_34, __local_var_5_32})
											}
										end_branch_84:
											__t126 = __t84
											goto end_branch_126
										} else {

										}
									}
									{
										var __t_and_86 bool = false
										if v3_3 != nil {

											var __t_tag_85 uint32 = (v3_3).V0
											_ = __t_tag_85
											__t_and_86 = (uint32(__t_tag_85) == 2247809753)
										}
										if __t_and_86 {
											var __t125 *Constructor_Main_Branch
											{
												var __t_tag_91 *Constructor_Main_Branch = (v3_3).V1
												_ = __t_tag_91
												if __t_tag_91 != nil {
													var __t114 *Constructor_Main_Branch
													{
														var __t_tag_96 uint32 = ((v3_3).V1).V0
														_ = __t_tag_96
														if uint32(__t_tag_96) == 2247809753 {
															// TAST (Let): __local_var_4_97 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
															__local_var_4_97 := v1_1
															_ = __local_var_4_97
															// TAST (Let): __local_var_5_98 shape=Other bindingType=Any
															__local_var_5_98 := ((v3_3).V1).V1
															_ = __local_var_5_98
															// TAST (Let): __local_var_6_99 shape=Other bindingType=Any
															__local_var_6_99 := ((v3_3).V1).V3
															_ = __local_var_6_99
															// TAST (Let): __local_var_7_100 shape=Other bindingType=Any
															__local_var_7_100 := (v3_3).V3
															_ = __local_var_7_100
															// TAST (Let): __local_var_8_101 shape=Other bindingType=Int
															__local_var_8_101 := v2_2
															_ = __local_var_8_101
															// TAST (Let): __local_var_9_102 shape=Other bindingType=Any
															__local_var_9_102 := ((v3_3).V1).V2
															_ = __local_var_9_102
															// TAST (Let): __local_var_10_103 shape=Other bindingType=Any
															__local_var_10_103 := (v3_3).V2
															_ = __local_var_10_103
															__t114 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_97, __local_var_8_101, __local_var_5_98}), __local_var_9_102, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_99, __local_var_10_103, __local_var_7_100})})
															goto end_branch_114
														} else {

														}
													}
													{
														var __t_tag_104 *Constructor_Main_Branch = (v3_3).V3
														_ = __t_tag_104
														var __t_and_106 bool = false
														if __t_tag_104 != nil {

															var __t_tag_105 uint32 = ((v3_3).V3).V0
															_ = __t_tag_105
															__t_and_106 = (uint32(__t_tag_105) == 2247809753)
														}
														if __t_and_106 {
															// TAST (Let): __local_var_4_107 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
															__local_var_4_107 := v1_1
															_ = __local_var_4_107
															// TAST (Let): __local_var_5_108 shape=Other bindingType=Any
															__local_var_5_108 := (v3_3).V1
															_ = __local_var_5_108
															// TAST (Let): __local_var_6_109 shape=Other bindingType=Any
															__local_var_6_109 := ((v3_3).V3).V1
															_ = __local_var_6_109
															// TAST (Let): __local_var_7_110 shape=Other bindingType=Any
															__local_var_7_110 := ((v3_3).V3).V3
															_ = __local_var_7_110
															// TAST (Let): __local_var_8_111 shape=Other bindingType=Int
															__local_var_8_111 := v2_2
															_ = __local_var_8_111
															// TAST (Let): __local_var_9_112 shape=Other bindingType=Any
															__local_var_9_112 := (v3_3).V2
															_ = __local_var_9_112
															// TAST (Let): __local_var_10_113 shape=Other bindingType=Any
															__local_var_10_113 := ((v3_3).V3).V2
															_ = __local_var_10_113
															__t114 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_107, __local_var_8_111, __local_var_5_108}), __local_var_9_112, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_109, __local_var_10_113, __local_var_7_110})})
															goto end_branch_114
														} else {

														}
													}
													{
														// TAST (Let): __local_var_4_92 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_4_92 := v1_1
														_ = __local_var_4_92
														// TAST (Let): __local_var_5_93 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_5_93 := v3_3
														_ = __local_var_5_93
														// TAST (Let): __local_var_6_94 shape=Other bindingType=(ADT ["Main","Paint"] [])
														__local_var_6_94 := v_0
														_ = __local_var_6_94
														// TAST (Let): __local_var_7_95 shape=Other bindingType=Int
														__local_var_7_95 := v2_2
														_ = __local_var_7_95
														__t114 = (&Constructor_Main_Branch{1, __local_var_6_94, __local_var_4_92, __local_var_7_95, __local_var_5_93})
													}
												end_branch_114:
													__t125 = __t114
													goto end_branch_125
												} else {

												}
											}
											{
												var __t_tag_115 *Constructor_Main_Branch = (v3_3).V3
												_ = __t_tag_115
												var __t_and_117 bool = false
												if __t_tag_115 != nil {

													var __t_tag_116 uint32 = ((v3_3).V3).V0
													_ = __t_tag_116
													__t_and_117 = (uint32(__t_tag_116) == 2247809753)
												}
												if __t_and_117 {
													// TAST (Let): __local_var_4_118 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
													__local_var_4_118 := v1_1
													_ = __local_var_4_118
													// TAST (Let): __local_var_5_119 shape=Other bindingType=Any
													__local_var_5_119 := (v3_3).V1
													_ = __local_var_5_119
													// TAST (Let): __local_var_6_120 shape=Other bindingType=Any
													__local_var_6_120 := ((v3_3).V3).V1
													_ = __local_var_6_120
													// TAST (Let): __local_var_7_121 shape=Other bindingType=Any
													__local_var_7_121 := ((v3_3).V3).V3
													_ = __local_var_7_121
													// TAST (Let): __local_var_8_122 shape=Other bindingType=Int
													__local_var_8_122 := v2_2
													_ = __local_var_8_122
													// TAST (Let): __local_var_9_123 shape=Other bindingType=Any
													__local_var_9_123 := (v3_3).V2
													_ = __local_var_9_123
													// TAST (Let): __local_var_10_124 shape=Other bindingType=Any
													__local_var_10_124 := ((v3_3).V3).V2
													_ = __local_var_10_124
													__t125 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_118, __local_var_8_122, __local_var_5_119}), __local_var_9_123, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_120, __local_var_10_124, __local_var_7_121})})
													goto end_branch_125
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_87 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_4_87 := v1_1
												_ = __local_var_4_87
												// TAST (Let): __local_var_5_88 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_5_88 := v3_3
												_ = __local_var_5_88
												// TAST (Let): __local_var_6_89 shape=Other bindingType=(ADT ["Main","Paint"] [])
												__local_var_6_89 := v_0
												_ = __local_var_6_89
												// TAST (Let): __local_var_7_90 shape=Other bindingType=Int
												__local_var_7_90 := v2_2
												_ = __local_var_7_90
												__t125 = (&Constructor_Main_Branch{1, __local_var_6_89, __local_var_4_87, __local_var_7_90, __local_var_5_88})
											}
										end_branch_125:
											__t126 = __t125
											goto end_branch_126
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_18 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_4_18 := v1_1
										_ = __local_var_4_18
										// TAST (Let): __local_var_5_19 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_5_19 := v3_3
										_ = __local_var_5_19
										// TAST (Let): __local_var_6_20 shape=Other bindingType=(ADT ["Main","Paint"] [])
										__local_var_6_20 := v_0
										_ = __local_var_6_20
										// TAST (Let): __local_var_7_21 shape=Other bindingType=Int
										__local_var_7_21 := v2_2
										_ = __local_var_7_21
										__t126 = (&Constructor_Main_Branch{1, __local_var_6_20, __local_var_4_18, __local_var_7_21, __local_var_5_19})
									}
								end_branch_126:
									__t223 = __t126
									goto end_branch_223
								} else {

								}
							}
							{
								var __t_tag_127 *Constructor_Main_Branch = (v1_1).V3
								_ = __t_tag_127
								if __t_tag_127 != nil {
									var __t181 *Constructor_Main_Branch
									{
										var __t_tag_132 uint32 = ((v1_1).V3).V0
										_ = __t_tag_132
										if uint32(__t_tag_132) == 2247809753 {
											// TAST (Let): __local_var_4_133 shape=Other bindingType=Any
											__local_var_4_133 := (v1_1).V1
											_ = __local_var_4_133
											// TAST (Let): __local_var_5_134 shape=Other bindingType=Any
											__local_var_5_134 := ((v1_1).V3).V1
											_ = __local_var_5_134
											// TAST (Let): __local_var_6_135 shape=Other bindingType=Any
											__local_var_6_135 := ((v1_1).V3).V3
											_ = __local_var_6_135
											// TAST (Let): __local_var_7_136 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
											__local_var_7_136 := v3_3
											_ = __local_var_7_136
											// TAST (Let): __local_var_8_137 shape=Other bindingType=Any
											__local_var_8_137 := (v1_1).V2
											_ = __local_var_8_137
											// TAST (Let): __local_var_9_138 shape=Other bindingType=Any
											__local_var_9_138 := ((v1_1).V3).V2
											_ = __local_var_9_138
											// TAST (Let): __local_var_10_139 shape=Other bindingType=Int
											__local_var_10_139 := v2_2
											_ = __local_var_10_139
											__t181 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_133, __local_var_8_137, __local_var_5_134}), __local_var_9_138, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_135, __local_var_10_139, __local_var_7_136})})
											goto end_branch_181
										} else {

										}
									}
									{
										var __t_and_141 bool = false
										if v3_3 != nil {

											var __t_tag_140 uint32 = (v3_3).V0
											_ = __t_tag_140
											__t_and_141 = (uint32(__t_tag_140) == 2247809753)
										}
										if __t_and_141 {
											var __t180 *Constructor_Main_Branch
											{
												var __t_tag_146 *Constructor_Main_Branch = (v3_3).V1
												_ = __t_tag_146
												if __t_tag_146 != nil {
													var __t169 *Constructor_Main_Branch
													{
														var __t_tag_151 uint32 = ((v3_3).V1).V0
														_ = __t_tag_151
														if uint32(__t_tag_151) == 2247809753 {
															// TAST (Let): __local_var_4_152 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
															__local_var_4_152 := v1_1
															_ = __local_var_4_152
															// TAST (Let): __local_var_5_153 shape=Other bindingType=Any
															__local_var_5_153 := ((v3_3).V1).V1
															_ = __local_var_5_153
															// TAST (Let): __local_var_6_154 shape=Other bindingType=Any
															__local_var_6_154 := ((v3_3).V1).V3
															_ = __local_var_6_154
															// TAST (Let): __local_var_7_155 shape=Other bindingType=Any
															__local_var_7_155 := (v3_3).V3
															_ = __local_var_7_155
															// TAST (Let): __local_var_8_156 shape=Other bindingType=Int
															__local_var_8_156 := v2_2
															_ = __local_var_8_156
															// TAST (Let): __local_var_9_157 shape=Other bindingType=Any
															__local_var_9_157 := ((v3_3).V1).V2
															_ = __local_var_9_157
															// TAST (Let): __local_var_10_158 shape=Other bindingType=Any
															__local_var_10_158 := (v3_3).V2
															_ = __local_var_10_158
															__t169 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_152, __local_var_8_156, __local_var_5_153}), __local_var_9_157, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_154, __local_var_10_158, __local_var_7_155})})
															goto end_branch_169
														} else {

														}
													}
													{
														var __t_tag_159 *Constructor_Main_Branch = (v3_3).V3
														_ = __t_tag_159
														var __t_and_161 bool = false
														if __t_tag_159 != nil {

															var __t_tag_160 uint32 = ((v3_3).V3).V0
															_ = __t_tag_160
															__t_and_161 = (uint32(__t_tag_160) == 2247809753)
														}
														if __t_and_161 {
															// TAST (Let): __local_var_4_162 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
															__local_var_4_162 := v1_1
															_ = __local_var_4_162
															// TAST (Let): __local_var_5_163 shape=Other bindingType=Any
															__local_var_5_163 := (v3_3).V1
															_ = __local_var_5_163
															// TAST (Let): __local_var_6_164 shape=Other bindingType=Any
															__local_var_6_164 := ((v3_3).V3).V1
															_ = __local_var_6_164
															// TAST (Let): __local_var_7_165 shape=Other bindingType=Any
															__local_var_7_165 := ((v3_3).V3).V3
															_ = __local_var_7_165
															// TAST (Let): __local_var_8_166 shape=Other bindingType=Int
															__local_var_8_166 := v2_2
															_ = __local_var_8_166
															// TAST (Let): __local_var_9_167 shape=Other bindingType=Any
															__local_var_9_167 := (v3_3).V2
															_ = __local_var_9_167
															// TAST (Let): __local_var_10_168 shape=Other bindingType=Any
															__local_var_10_168 := ((v3_3).V3).V2
															_ = __local_var_10_168
															__t169 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_162, __local_var_8_166, __local_var_5_163}), __local_var_9_167, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_164, __local_var_10_168, __local_var_7_165})})
															goto end_branch_169
														} else {

														}
													}
													{
														// TAST (Let): __local_var_4_147 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_4_147 := v1_1
														_ = __local_var_4_147
														// TAST (Let): __local_var_5_148 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_5_148 := v3_3
														_ = __local_var_5_148
														// TAST (Let): __local_var_6_149 shape=Other bindingType=(ADT ["Main","Paint"] [])
														__local_var_6_149 := v_0
														_ = __local_var_6_149
														// TAST (Let): __local_var_7_150 shape=Other bindingType=Int
														__local_var_7_150 := v2_2
														_ = __local_var_7_150
														__t169 = (&Constructor_Main_Branch{1, __local_var_6_149, __local_var_4_147, __local_var_7_150, __local_var_5_148})
													}
												end_branch_169:
													__t180 = __t169
													goto end_branch_180
												} else {

												}
											}
											{
												var __t_tag_170 *Constructor_Main_Branch = (v3_3).V3
												_ = __t_tag_170
												var __t_and_172 bool = false
												if __t_tag_170 != nil {

													var __t_tag_171 uint32 = ((v3_3).V3).V0
													_ = __t_tag_171
													__t_and_172 = (uint32(__t_tag_171) == 2247809753)
												}
												if __t_and_172 {
													// TAST (Let): __local_var_4_173 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
													__local_var_4_173 := v1_1
													_ = __local_var_4_173
													// TAST (Let): __local_var_5_174 shape=Other bindingType=Any
													__local_var_5_174 := (v3_3).V1
													_ = __local_var_5_174
													// TAST (Let): __local_var_6_175 shape=Other bindingType=Any
													__local_var_6_175 := ((v3_3).V3).V1
													_ = __local_var_6_175
													// TAST (Let): __local_var_7_176 shape=Other bindingType=Any
													__local_var_7_176 := ((v3_3).V3).V3
													_ = __local_var_7_176
													// TAST (Let): __local_var_8_177 shape=Other bindingType=Int
													__local_var_8_177 := v2_2
													_ = __local_var_8_177
													// TAST (Let): __local_var_9_178 shape=Other bindingType=Any
													__local_var_9_178 := (v3_3).V2
													_ = __local_var_9_178
													// TAST (Let): __local_var_10_179 shape=Other bindingType=Any
													__local_var_10_179 := ((v3_3).V3).V2
													_ = __local_var_10_179
													__t180 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_173, __local_var_8_177, __local_var_5_174}), __local_var_9_178, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_175, __local_var_10_179, __local_var_7_176})})
													goto end_branch_180
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_142 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_4_142 := v1_1
												_ = __local_var_4_142
												// TAST (Let): __local_var_5_143 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_5_143 := v3_3
												_ = __local_var_5_143
												// TAST (Let): __local_var_6_144 shape=Other bindingType=(ADT ["Main","Paint"] [])
												__local_var_6_144 := v_0
												_ = __local_var_6_144
												// TAST (Let): __local_var_7_145 shape=Other bindingType=Int
												__local_var_7_145 := v2_2
												_ = __local_var_7_145
												__t180 = (&Constructor_Main_Branch{1, __local_var_6_144, __local_var_4_142, __local_var_7_145, __local_var_5_143})
											}
										end_branch_180:
											__t181 = __t180
											goto end_branch_181
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_128 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_4_128 := v1_1
										_ = __local_var_4_128
										// TAST (Let): __local_var_5_129 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_5_129 := v3_3
										_ = __local_var_5_129
										// TAST (Let): __local_var_6_130 shape=Other bindingType=(ADT ["Main","Paint"] [])
										__local_var_6_130 := v_0
										_ = __local_var_6_130
										// TAST (Let): __local_var_7_131 shape=Other bindingType=Int
										__local_var_7_131 := v2_2
										_ = __local_var_7_131
										__t181 = (&Constructor_Main_Branch{1, __local_var_6_130, __local_var_4_128, __local_var_7_131, __local_var_5_129})
									}
								end_branch_181:
									__t223 = __t181
									goto end_branch_223
								} else {

								}
							}
							{
								var __t_and_183 bool = false
								if v3_3 != nil {

									var __t_tag_182 uint32 = (v3_3).V0
									_ = __t_tag_182
									__t_and_183 = (uint32(__t_tag_182) == 2247809753)
								}
								if __t_and_183 {
									var __t222 *Constructor_Main_Branch
									{
										var __t_tag_188 *Constructor_Main_Branch = (v3_3).V1
										_ = __t_tag_188
										if __t_tag_188 != nil {
											var __t211 *Constructor_Main_Branch
											{
												var __t_tag_193 uint32 = ((v3_3).V1).V0
												_ = __t_tag_193
												if uint32(__t_tag_193) == 2247809753 {
													// TAST (Let): __local_var_4_194 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
													__local_var_4_194 := v1_1
													_ = __local_var_4_194
													// TAST (Let): __local_var_5_195 shape=Other bindingType=Any
													__local_var_5_195 := ((v3_3).V1).V1
													_ = __local_var_5_195
													// TAST (Let): __local_var_6_196 shape=Other bindingType=Any
													__local_var_6_196 := ((v3_3).V1).V3
													_ = __local_var_6_196
													// TAST (Let): __local_var_7_197 shape=Other bindingType=Any
													__local_var_7_197 := (v3_3).V3
													_ = __local_var_7_197
													// TAST (Let): __local_var_8_198 shape=Other bindingType=Int
													__local_var_8_198 := v2_2
													_ = __local_var_8_198
													// TAST (Let): __local_var_9_199 shape=Other bindingType=Any
													__local_var_9_199 := ((v3_3).V1).V2
													_ = __local_var_9_199
													// TAST (Let): __local_var_10_200 shape=Other bindingType=Any
													__local_var_10_200 := (v3_3).V2
													_ = __local_var_10_200
													__t211 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_194, __local_var_8_198, __local_var_5_195}), __local_var_9_199, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_196, __local_var_10_200, __local_var_7_197})})
													goto end_branch_211
												} else {

												}
											}
											{
												var __t_tag_201 *Constructor_Main_Branch = (v3_3).V3
												_ = __t_tag_201
												var __t_and_203 bool = false
												if __t_tag_201 != nil {

													var __t_tag_202 uint32 = ((v3_3).V3).V0
													_ = __t_tag_202
													__t_and_203 = (uint32(__t_tag_202) == 2247809753)
												}
												if __t_and_203 {
													// TAST (Let): __local_var_4_204 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
													__local_var_4_204 := v1_1
													_ = __local_var_4_204
													// TAST (Let): __local_var_5_205 shape=Other bindingType=Any
													__local_var_5_205 := (v3_3).V1
													_ = __local_var_5_205
													// TAST (Let): __local_var_6_206 shape=Other bindingType=Any
													__local_var_6_206 := ((v3_3).V3).V1
													_ = __local_var_6_206
													// TAST (Let): __local_var_7_207 shape=Other bindingType=Any
													__local_var_7_207 := ((v3_3).V3).V3
													_ = __local_var_7_207
													// TAST (Let): __local_var_8_208 shape=Other bindingType=Int
													__local_var_8_208 := v2_2
													_ = __local_var_8_208
													// TAST (Let): __local_var_9_209 shape=Other bindingType=Any
													__local_var_9_209 := (v3_3).V2
													_ = __local_var_9_209
													// TAST (Let): __local_var_10_210 shape=Other bindingType=Any
													__local_var_10_210 := ((v3_3).V3).V2
													_ = __local_var_10_210
													__t211 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_204, __local_var_8_208, __local_var_5_205}), __local_var_9_209, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_206, __local_var_10_210, __local_var_7_207})})
													goto end_branch_211
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_189 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_4_189 := v1_1
												_ = __local_var_4_189
												// TAST (Let): __local_var_5_190 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_5_190 := v3_3
												_ = __local_var_5_190
												// TAST (Let): __local_var_6_191 shape=Other bindingType=(ADT ["Main","Paint"] [])
												__local_var_6_191 := v_0
												_ = __local_var_6_191
												// TAST (Let): __local_var_7_192 shape=Other bindingType=Int
												__local_var_7_192 := v2_2
												_ = __local_var_7_192
												__t211 = (&Constructor_Main_Branch{1, __local_var_6_191, __local_var_4_189, __local_var_7_192, __local_var_5_190})
											}
										end_branch_211:
											__t222 = __t211
											goto end_branch_222
										} else {

										}
									}
									{
										var __t_tag_212 *Constructor_Main_Branch = (v3_3).V3
										_ = __t_tag_212
										var __t_and_214 bool = false
										if __t_tag_212 != nil {

											var __t_tag_213 uint32 = ((v3_3).V3).V0
											_ = __t_tag_213
											__t_and_214 = (uint32(__t_tag_213) == 2247809753)
										}
										if __t_and_214 {
											// TAST (Let): __local_var_4_215 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
											__local_var_4_215 := v1_1
											_ = __local_var_4_215
											// TAST (Let): __local_var_5_216 shape=Other bindingType=Any
											__local_var_5_216 := (v3_3).V1
											_ = __local_var_5_216
											// TAST (Let): __local_var_6_217 shape=Other bindingType=Any
											__local_var_6_217 := ((v3_3).V3).V1
											_ = __local_var_6_217
											// TAST (Let): __local_var_7_218 shape=Other bindingType=Any
											__local_var_7_218 := ((v3_3).V3).V3
											_ = __local_var_7_218
											// TAST (Let): __local_var_8_219 shape=Other bindingType=Int
											__local_var_8_219 := v2_2
											_ = __local_var_8_219
											// TAST (Let): __local_var_9_220 shape=Other bindingType=Any
											__local_var_9_220 := (v3_3).V2
											_ = __local_var_9_220
											// TAST (Let): __local_var_10_221 shape=Other bindingType=Any
											__local_var_10_221 := ((v3_3).V3).V2
											_ = __local_var_10_221
											__t222 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_215, __local_var_8_219, __local_var_5_216}), __local_var_9_220, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_217, __local_var_10_221, __local_var_7_218})})
											goto end_branch_222
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_184 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_4_184 := v1_1
										_ = __local_var_4_184
										// TAST (Let): __local_var_5_185 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_5_185 := v3_3
										_ = __local_var_5_185
										// TAST (Let): __local_var_6_186 shape=Other bindingType=(ADT ["Main","Paint"] [])
										__local_var_6_186 := v_0
										_ = __local_var_6_186
										// TAST (Let): __local_var_7_187 shape=Other bindingType=Int
										__local_var_7_187 := v2_2
										_ = __local_var_7_187
										__t222 = (&Constructor_Main_Branch{1, __local_var_6_186, __local_var_4_184, __local_var_7_187, __local_var_5_185})
									}
								end_branch_222:
									__t223 = __t222
									goto end_branch_223
								} else {

								}
							}
							{
								// TAST (Let): __local_var_4_13 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_4_13 := v1_1
								_ = __local_var_4_13
								// TAST (Let): __local_var_5_14 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_5_14 := v3_3
								_ = __local_var_5_14
								// TAST (Let): __local_var_6_15 shape=Other bindingType=(ADT ["Main","Paint"] [])
								__local_var_6_15 := v_0
								_ = __local_var_6_15
								// TAST (Let): __local_var_7_16 shape=Other bindingType=Int
								__local_var_7_16 := v2_2
								_ = __local_var_7_16
								__t223 = (&Constructor_Main_Branch{1, __local_var_6_15, __local_var_4_13, __local_var_7_16, __local_var_5_14})
							}
						end_branch_223:
							__t265 = __t223
							goto end_branch_265
						} else {

						}
					}
					{
						var __t_and_225 bool = false
						if v3_3 != nil {

							var __t_tag_224 uint32 = (v3_3).V0
							_ = __t_tag_224
							__t_and_225 = (uint32(__t_tag_224) == 2247809753)
						}
						if __t_and_225 {
							var __t264 *Constructor_Main_Branch
							{
								var __t_tag_230 *Constructor_Main_Branch = (v3_3).V1
								_ = __t_tag_230
								if __t_tag_230 != nil {
									var __t253 *Constructor_Main_Branch
									{
										var __t_tag_235 uint32 = ((v3_3).V1).V0
										_ = __t_tag_235
										if uint32(__t_tag_235) == 2247809753 {
											// TAST (Let): __local_var_4_236 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
											__local_var_4_236 := v1_1
											_ = __local_var_4_236
											// TAST (Let): __local_var_5_237 shape=Other bindingType=Any
											__local_var_5_237 := ((v3_3).V1).V1
											_ = __local_var_5_237
											// TAST (Let): __local_var_6_238 shape=Other bindingType=Any
											__local_var_6_238 := ((v3_3).V1).V3
											_ = __local_var_6_238
											// TAST (Let): __local_var_7_239 shape=Other bindingType=Any
											__local_var_7_239 := (v3_3).V3
											_ = __local_var_7_239
											// TAST (Let): __local_var_8_240 shape=Other bindingType=Int
											__local_var_8_240 := v2_2
											_ = __local_var_8_240
											// TAST (Let): __local_var_9_241 shape=Other bindingType=Any
											__local_var_9_241 := ((v3_3).V1).V2
											_ = __local_var_9_241
											// TAST (Let): __local_var_10_242 shape=Other bindingType=Any
											__local_var_10_242 := (v3_3).V2
											_ = __local_var_10_242
											__t253 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_236, __local_var_8_240, __local_var_5_237}), __local_var_9_241, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_238, __local_var_10_242, __local_var_7_239})})
											goto end_branch_253
										} else {

										}
									}
									{
										var __t_tag_243 *Constructor_Main_Branch = (v3_3).V3
										_ = __t_tag_243
										var __t_and_245 bool = false
										if __t_tag_243 != nil {

											var __t_tag_244 uint32 = ((v3_3).V3).V0
											_ = __t_tag_244
											__t_and_245 = (uint32(__t_tag_244) == 2247809753)
										}
										if __t_and_245 {
											// TAST (Let): __local_var_4_246 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
											__local_var_4_246 := v1_1
											_ = __local_var_4_246
											// TAST (Let): __local_var_5_247 shape=Other bindingType=Any
											__local_var_5_247 := (v3_3).V1
											_ = __local_var_5_247
											// TAST (Let): __local_var_6_248 shape=Other bindingType=Any
											__local_var_6_248 := ((v3_3).V3).V1
											_ = __local_var_6_248
											// TAST (Let): __local_var_7_249 shape=Other bindingType=Any
											__local_var_7_249 := ((v3_3).V3).V3
											_ = __local_var_7_249
											// TAST (Let): __local_var_8_250 shape=Other bindingType=Int
											__local_var_8_250 := v2_2
											_ = __local_var_8_250
											// TAST (Let): __local_var_9_251 shape=Other bindingType=Any
											__local_var_9_251 := (v3_3).V2
											_ = __local_var_9_251
											// TAST (Let): __local_var_10_252 shape=Other bindingType=Any
											__local_var_10_252 := ((v3_3).V3).V2
											_ = __local_var_10_252
											__t253 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_246, __local_var_8_250, __local_var_5_247}), __local_var_9_251, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_248, __local_var_10_252, __local_var_7_249})})
											goto end_branch_253
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_231 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_4_231 := v1_1
										_ = __local_var_4_231
										// TAST (Let): __local_var_5_232 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_5_232 := v3_3
										_ = __local_var_5_232
										// TAST (Let): __local_var_6_233 shape=Other bindingType=(ADT ["Main","Paint"] [])
										__local_var_6_233 := v_0
										_ = __local_var_6_233
										// TAST (Let): __local_var_7_234 shape=Other bindingType=Int
										__local_var_7_234 := v2_2
										_ = __local_var_7_234
										__t253 = (&Constructor_Main_Branch{1, __local_var_6_233, __local_var_4_231, __local_var_7_234, __local_var_5_232})
									}
								end_branch_253:
									__t264 = __t253
									goto end_branch_264
								} else {

								}
							}
							{
								var __t_tag_254 *Constructor_Main_Branch = (v3_3).V3
								_ = __t_tag_254
								var __t_and_256 bool = false
								if __t_tag_254 != nil {

									var __t_tag_255 uint32 = ((v3_3).V3).V0
									_ = __t_tag_255
									__t_and_256 = (uint32(__t_tag_255) == 2247809753)
								}
								if __t_and_256 {
									// TAST (Let): __local_var_4_257 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
									__local_var_4_257 := v1_1
									_ = __local_var_4_257
									// TAST (Let): __local_var_5_258 shape=Other bindingType=Any
									__local_var_5_258 := (v3_3).V1
									_ = __local_var_5_258
									// TAST (Let): __local_var_6_259 shape=Other bindingType=Any
									__local_var_6_259 := ((v3_3).V3).V1
									_ = __local_var_6_259
									// TAST (Let): __local_var_7_260 shape=Other bindingType=Any
									__local_var_7_260 := ((v3_3).V3).V3
									_ = __local_var_7_260
									// TAST (Let): __local_var_8_261 shape=Other bindingType=Int
									__local_var_8_261 := v2_2
									_ = __local_var_8_261
									// TAST (Let): __local_var_9_262 shape=Other bindingType=Any
									__local_var_9_262 := (v3_3).V2
									_ = __local_var_9_262
									// TAST (Let): __local_var_10_263 shape=Other bindingType=Any
									__local_var_10_263 := ((v3_3).V3).V2
									_ = __local_var_10_263
									__t264 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_257, __local_var_8_261, __local_var_5_258}), __local_var_9_262, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_259, __local_var_10_263, __local_var_7_260})})
									goto end_branch_264
								} else {

								}
							}
							{
								// TAST (Let): __local_var_4_226 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_4_226 := v1_1
								_ = __local_var_4_226
								// TAST (Let): __local_var_5_227 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_5_227 := v3_3
								_ = __local_var_5_227
								// TAST (Let): __local_var_6_228 shape=Other bindingType=(ADT ["Main","Paint"] [])
								__local_var_6_228 := v_0
								_ = __local_var_6_228
								// TAST (Let): __local_var_7_229 shape=Other bindingType=Int
								__local_var_7_229 := v2_2
								_ = __local_var_7_229
								__t264 = (&Constructor_Main_Branch{1, __local_var_6_228, __local_var_4_226, __local_var_7_229, __local_var_5_227})
							}
						end_branch_264:
							__t265 = __t264
							goto end_branch_265
						} else {

						}
					}
					{
						// TAST (Let): __local_var_4_8 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
						__local_var_4_8 := v1_1
						_ = __local_var_4_8
						// TAST (Let): __local_var_5_9 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
						__local_var_5_9 := v3_3
						_ = __local_var_5_9
						// TAST (Let): __local_var_6_10 shape=Other bindingType=(ADT ["Main","Paint"] [])
						__local_var_6_10 := v_0
						_ = __local_var_6_10
						// TAST (Let): __local_var_7_11 shape=Other bindingType=Int
						__local_var_7_11 := v2_2
						_ = __local_var_7_11
						__t265 = (&Constructor_Main_Branch{1, __local_var_6_10, __local_var_4_8, __local_var_7_11, __local_var_5_9})
					}
				end_branch_265:
					__t307 = __t265
					goto end_branch_307
				} else {

				}
			}
			{
				var __t_and_267 bool = false
				if v3_3 != nil {

					var __t_tag_266 uint32 = (v3_3).V0
					_ = __t_tag_266
					__t_and_267 = (uint32(__t_tag_266) == 2247809753)
				}
				if __t_and_267 {
					var __t306 *Constructor_Main_Branch
					{
						var __t_tag_272 *Constructor_Main_Branch = (v3_3).V1
						_ = __t_tag_272
						if __t_tag_272 != nil {
							var __t295 *Constructor_Main_Branch
							{
								var __t_tag_277 uint32 = ((v3_3).V1).V0
								_ = __t_tag_277
								if uint32(__t_tag_277) == 2247809753 {
									// TAST (Let): __local_var_4_278 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
									__local_var_4_278 := v1_1
									_ = __local_var_4_278
									// TAST (Let): __local_var_5_279 shape=Other bindingType=Any
									__local_var_5_279 := ((v3_3).V1).V1
									_ = __local_var_5_279
									// TAST (Let): __local_var_6_280 shape=Other bindingType=Any
									__local_var_6_280 := ((v3_3).V1).V3
									_ = __local_var_6_280
									// TAST (Let): __local_var_7_281 shape=Other bindingType=Any
									__local_var_7_281 := (v3_3).V3
									_ = __local_var_7_281
									// TAST (Let): __local_var_8_282 shape=Other bindingType=Int
									__local_var_8_282 := v2_2
									_ = __local_var_8_282
									// TAST (Let): __local_var_9_283 shape=Other bindingType=Any
									__local_var_9_283 := ((v3_3).V1).V2
									_ = __local_var_9_283
									// TAST (Let): __local_var_10_284 shape=Other bindingType=Any
									__local_var_10_284 := (v3_3).V2
									_ = __local_var_10_284
									__t295 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_278, __local_var_8_282, __local_var_5_279}), __local_var_9_283, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_280, __local_var_10_284, __local_var_7_281})})
									goto end_branch_295
								} else {

								}
							}
							{
								var __t_tag_285 *Constructor_Main_Branch = (v3_3).V3
								_ = __t_tag_285
								var __t_and_287 bool = false
								if __t_tag_285 != nil {

									var __t_tag_286 uint32 = ((v3_3).V3).V0
									_ = __t_tag_286
									__t_and_287 = (uint32(__t_tag_286) == 2247809753)
								}
								if __t_and_287 {
									// TAST (Let): __local_var_4_288 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
									__local_var_4_288 := v1_1
									_ = __local_var_4_288
									// TAST (Let): __local_var_5_289 shape=Other bindingType=Any
									__local_var_5_289 := (v3_3).V1
									_ = __local_var_5_289
									// TAST (Let): __local_var_6_290 shape=Other bindingType=Any
									__local_var_6_290 := ((v3_3).V3).V1
									_ = __local_var_6_290
									// TAST (Let): __local_var_7_291 shape=Other bindingType=Any
									__local_var_7_291 := ((v3_3).V3).V3
									_ = __local_var_7_291
									// TAST (Let): __local_var_8_292 shape=Other bindingType=Int
									__local_var_8_292 := v2_2
									_ = __local_var_8_292
									// TAST (Let): __local_var_9_293 shape=Other bindingType=Any
									__local_var_9_293 := (v3_3).V2
									_ = __local_var_9_293
									// TAST (Let): __local_var_10_294 shape=Other bindingType=Any
									__local_var_10_294 := ((v3_3).V3).V2
									_ = __local_var_10_294
									__t295 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_288, __local_var_8_292, __local_var_5_289}), __local_var_9_293, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_290, __local_var_10_294, __local_var_7_291})})
									goto end_branch_295
								} else {

								}
							}
							{
								// TAST (Let): __local_var_4_273 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_4_273 := v1_1
								_ = __local_var_4_273
								// TAST (Let): __local_var_5_274 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_5_274 := v3_3
								_ = __local_var_5_274
								// TAST (Let): __local_var_6_275 shape=Other bindingType=(ADT ["Main","Paint"] [])
								__local_var_6_275 := v_0
								_ = __local_var_6_275
								// TAST (Let): __local_var_7_276 shape=Other bindingType=Int
								__local_var_7_276 := v2_2
								_ = __local_var_7_276
								__t295 = (&Constructor_Main_Branch{1, __local_var_6_275, __local_var_4_273, __local_var_7_276, __local_var_5_274})
							}
						end_branch_295:
							__t306 = __t295
							goto end_branch_306
						} else {

						}
					}
					{
						var __t_tag_296 *Constructor_Main_Branch = (v3_3).V3
						_ = __t_tag_296
						var __t_and_298 bool = false
						if __t_tag_296 != nil {

							var __t_tag_297 uint32 = ((v3_3).V3).V0
							_ = __t_tag_297
							__t_and_298 = (uint32(__t_tag_297) == 2247809753)
						}
						if __t_and_298 {
							// TAST (Let): __local_var_4_299 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
							__local_var_4_299 := v1_1
							_ = __local_var_4_299
							// TAST (Let): __local_var_5_300 shape=Other bindingType=Any
							__local_var_5_300 := (v3_3).V1
							_ = __local_var_5_300
							// TAST (Let): __local_var_6_301 shape=Other bindingType=Any
							__local_var_6_301 := ((v3_3).V3).V1
							_ = __local_var_6_301
							// TAST (Let): __local_var_7_302 shape=Other bindingType=Any
							__local_var_7_302 := ((v3_3).V3).V3
							_ = __local_var_7_302
							// TAST (Let): __local_var_8_303 shape=Other bindingType=Int
							__local_var_8_303 := v2_2
							_ = __local_var_8_303
							// TAST (Let): __local_var_9_304 shape=Other bindingType=Any
							__local_var_9_304 := (v3_3).V2
							_ = __local_var_9_304
							// TAST (Let): __local_var_10_305 shape=Other bindingType=Any
							__local_var_10_305 := ((v3_3).V3).V2
							_ = __local_var_10_305
							__t306 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_299, __local_var_8_303, __local_var_5_300}), __local_var_9_304, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_301, __local_var_10_305, __local_var_7_302})})
							goto end_branch_306
						} else {

						}
					}
					{
						// TAST (Let): __local_var_4_268 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
						__local_var_4_268 := v1_1
						_ = __local_var_4_268
						// TAST (Let): __local_var_5_269 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
						__local_var_5_269 := v3_3
						_ = __local_var_5_269
						// TAST (Let): __local_var_6_270 shape=Other bindingType=(ADT ["Main","Paint"] [])
						__local_var_6_270 := v_0
						_ = __local_var_6_270
						// TAST (Let): __local_var_7_271 shape=Other bindingType=Int
						__local_var_7_271 := v2_2
						_ = __local_var_7_271
						__t306 = (&Constructor_Main_Branch{1, __local_var_6_270, __local_var_4_268, __local_var_7_271, __local_var_5_269})
					}
				end_branch_306:
					__t307 = __t306
					goto end_branch_307
				} else {

				}
			}
			{
				// TAST (Let): __local_var_4_4 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
				__local_var_4_4 := v1_1
				_ = __local_var_4_4
				// TAST (Let): __local_var_5_5 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
				__local_var_5_5 := v3_3
				_ = __local_var_5_5
				// TAST (Let): __local_var_6_6 shape=Other bindingType=(ADT ["Main","Paint"] [])
				__local_var_6_6 := v_0
				_ = __local_var_6_6
				// TAST (Let): __local_var_7_7 shape=Other bindingType=Int
				__local_var_7_7 := v2_2
				_ = __local_var_7_7
				__t307 = (&Constructor_Main_Branch{1, __local_var_6_6, __local_var_4_4, __local_var_7_7, __local_var_5_5})
			}
		end_branch_307:
			__t308 = __t307
			goto end_branch_308
		} else {

		}
	}
	{
		// TAST (Let): __local_var_4_0 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
		__local_var_4_0 := v1_1
		_ = __local_var_4_0
		// TAST (Let): __local_var_5_1 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
		__local_var_5_1 := v3_3
		_ = __local_var_5_1
		// TAST (Let): __local_var_6_2 shape=Other bindingType=(ADT ["Main","Paint"] [])
		__local_var_6_2 := v_0
		_ = __local_var_6_2
		// TAST (Let): __local_var_7_3 shape=Other bindingType=Int
		__local_var_7_3 := v2_2
		_ = __local_var_7_3
		__t308 = (&Constructor_Main_Branch{1, __local_var_6_2, __local_var_4_0, __local_var_7_3, __local_var_5_1})
	}
end_branch_308:
	return __t308
}

func Call_Main_leftChild(v_0_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
	var v_0 *Constructor_Main_Branch = v_0_loop
	_ = v_0
	var __t0 *Constructor_Main_Branch
	{
		if v_0 == nil {
			__t0 = (*Constructor_Main_Branch)(nil)
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 != nil {
			__t0 = (v_0).V1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() *Constructor_Main_Branch { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Main_isRed(v_0_loop *Constructor_Main_Branch) bool {
	var v_0 *Constructor_Main_Branch = v_0_loop
	_ = v_0
	var __t_and_1 bool = false
	if v_0 != nil {

		var __t_tag_0 uint32 = (v_0).V0
		_ = __t_tag_0
		__t_and_1 = (uint32(__t_tag_0) == 2247809753)
	}
	return __t_and_1
}

func Call_Main_descend(v_0_loop int64, v1_1_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
descend:
	for {
		if false {
			continue descend
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Main_Branch = v1_1_loop
		_ = v1_1
		var __t1 *Constructor_Main_Branch
		{
			if v1_1 == nil {
				__t1 = (&Constructor_Main_Branch{1, 2247809753, (*Constructor_Main_Branch)(nil), v_0, (*Constructor_Main_Branch)(nil)})
				goto end_branch_1
			} else {

			}
		}
		{
			if v1_1 != nil {
				var __t0 *Constructor_Main_Branch
				{
					if (v_0) < ((v1_1).V2) {
						__t0 = Call_Main_rebalance((v1_1).V0, Call_Main_descend(v_0, (v1_1).V1), (v1_1).V2, (v1_1).V3)
						goto end_branch_0
					} else {

					}
				}
				{
					if (v_0) > ((v1_1).V2) {
						__t0 = Call_Main_rebalance((v1_1).V0, (v1_1).V1, (v1_1).V2, Call_Main_descend(v_0, (v1_1).V3))
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = (&Constructor_Main_Branch{1, (v1_1).V0, (v1_1).V1, (v1_1).V2, (v1_1).V3})
				}
			end_branch_0:
				__t1 = __t0
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() *Constructor_Main_Branch { panic("Failed pattern match") }()
		}
	end_branch_1:
		return __t1
	}
}

func Call_Main_blacken(v_0_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
	var v_0 *Constructor_Main_Branch = v_0_loop
	_ = v_0
	var __t1 *Constructor_Main_Branch
	{
		if v_0 == nil {
			__t1 = (*Constructor_Main_Branch)(nil)
			goto end_branch_1
		} else {

		}
	}
	{
		if v_0 != nil {
			var __reuse_0 *Constructor_Main_Branch
			if ((v_0) != (nil)) && (((v_0).V0) == (1685833310)) {
				__reuse_0 = v_0
			} else {
				__reuse_0 = (&Constructor_Main_Branch{1, 1685833310, (v_0).V1, (v_0).V2, (v_0).V3})
			}
			__t1 = __reuse_0
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() *Constructor_Main_Branch { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Main_put(value_0_loop int64, tree_1_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
	var value_0 int64 = value_0_loop
	_ = value_0
	var tree_1 *Constructor_Main_Branch = tree_1_loop
	_ = tree_1
	return Call_Main_blacken(Call_Main_descend(value_0, tree_1))
}

func Call_Main_build(v_0_loop int64, v1_1_loop int64, v2_2_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
build:
	for {
		if false {
			continue build
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var v2_2 *Constructor_Main_Branch = v2_2_loop
		_ = v2_2
		var __t0 *Constructor_Main_Branch
		{
			if (v_0) == (int64(0)) {
				__t0 = v2_2
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = v1_1
			v2_2_loop = Call_Main_put((v_0)*(v1_1), v2_2)
			continue build
			__t0 = func() *Constructor_Main_Branch { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_graftAndPut(child_0_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
	var child_0 *Constructor_Main_Branch = child_0_loop
	_ = child_0
	return Call_Main_put(int64(5), (&Constructor_Main_Branch{1, 1685833310, child_0, int64(50), (&Constructor_Main_Branch{1, 1685833310, (*Constructor_Main_Branch)(nil), int64(90), (*Constructor_Main_Branch)(nil)})}))
}

func Call_Main_mixed(v_0_loop int64, v1_1_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
mixed:
	for {
		if false {
			continue mixed
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Main_Branch = v1_1_loop
		_ = v1_1
		var __t0 *Constructor_Main_Branch
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = Call_Main_put(gopurs_runtime.IntMod((v_0)*(int64(37)), int64(127)), v1_1)
			continue mixed
			__t0 = func() *Constructor_Main_Branch { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_audit(v_0_loop int64, v1_1_loop int64, v2_2_loop *Constructor_Main_Branch) struct {
	blackHeight int64
	size        int64
	valid       bool
} {
audit:
	for {
		if false {
			continue audit
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var v2_2 *Constructor_Main_Branch = v2_2_loop
		_ = v2_2
		var __t20 struct {
			blackHeight int64
			size        int64
			valid       bool
		}
		{
			if v2_2 == nil {
				__t20 = struct {
					blackHeight int64
					size        int64
					valid       bool
				}{int64(1), int64(0), true}
				goto end_branch_20
			} else {

			}
		}
		{
			if v2_2 != nil {
				// TAST (Let): r_3_0 shape=App(Var) bindingType=(Record (Row [valid: Boolean, blackHeight: Int, size: Int] Empty))
				r_3_0 := Call_Main_audit((v2_2).V2, v1_1, (v2_2).V3)
				_ = r_3_0
				// TAST (Let): l_4_1 shape=App(Var) bindingType=(Record (Row [valid: Boolean, blackHeight: Int, size: Int] Empty))
				l_4_1 := Call_Main_audit(v_0, (v2_2).V2, (v2_2).V1)
				_ = l_4_1
				var __t4 int64
				{
					var __t_tag_2 uint32 = (v2_2).V0
					_ = __t_tag_2
					if uint32(__t_tag_2) == 1685833310 {
						__t4 = (l_4_1.blackHeight) + (int64(1))
						goto end_branch_4
					} else {

					}
				}
				{
					var __t_tag_3 uint32 = (v2_2).V0
					_ = __t_tag_3
					if uint32(__t_tag_3) == 2247809753 {
						__t4 = (l_4_1.blackHeight) + (int64(0))
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() int64 { panic("Failed pattern match") }()
				}
			end_branch_4:
				var __t_and_19 bool = false
				if l_4_1.valid {

					var __t_and_18 bool = false
					if r_3_0.valid {

						var __t_and_17 bool = false
						if (v_0) < ((v2_2).V2) {

							var __t_and_16 bool = false
							if ((v2_2).V2) < (v1_1) {

								var __t_and_15 bool = false
								if (l_4_1.blackHeight) == (r_3_0.blackHeight) {

									var __t_and_6 bool = false
									if v2_2 != nil {

										var __t_tag_5 uint32 = (v2_2).V0
										_ = __t_tag_5
										__t_and_6 = (uint32(__t_tag_5) == 2247809753)
									}
									var __t_or_14 bool = true
									if !((__t_and_6) != (true)) {

										var __t_tag_7 *Constructor_Main_Branch = (v2_2).V1
										_ = __t_tag_7
										var __t_and_9 bool = false
										if __t_tag_7 != nil {

											var __t_tag_8 uint32 = ((v2_2).V1).V0
											_ = __t_tag_8
											__t_and_9 = (uint32(__t_tag_8) == 2247809753)
										}
										var __t_and_13 bool = false
										if (__t_and_9) != (true) {

											var __t_tag_10 *Constructor_Main_Branch = (v2_2).V3
											_ = __t_tag_10
											var __t_and_12 bool = false
											if __t_tag_10 != nil {

												var __t_tag_11 uint32 = ((v2_2).V3).V0
												_ = __t_tag_11
												__t_and_12 = (uint32(__t_tag_11) == 2247809753)
											}
											__t_and_13 = (__t_and_12) != (true)
										}
										__t_or_14 = __t_and_13
									}
									__t_and_15 = __t_or_14
								}
								__t_and_16 = __t_and_15
							}
							__t_and_17 = __t_and_16
						}
						__t_and_18 = __t_and_17
					}
					__t_and_19 = __t_and_18
				}
				__t20 = struct {
					blackHeight int64
					size        int64
					valid       bool
				}{__t4, ((l_4_1.size) + (r_3_0.size)) + (int64(1)), __t_and_19}
				goto end_branch_20
			} else {

			}
		}
		{
			__t20 = func() struct {
				blackHeight int64
				size        int64
				valid       bool
			} {
				panic("Failed pattern match")
			}()
		}
	end_branch_20:
		return __t20
	}
}

func Call_Main_check(expectedSize_0_loop int64, tree_1_loop *Constructor_Main_Branch) gopurs_runtime.Value {
	var expectedSize_0 int64 = expectedSize_0_loop
	_ = expectedSize_0
	var tree_1 *Constructor_Main_Branch = tree_1_loop
	_ = tree_1
	// TAST (Let): result_2_0 shape=App(Var) bindingType=(Record (Row [valid: Boolean, blackHeight: Int, size: Int] Empty))
	result_2_0 := Call_Main_audit(int64(-1000000), int64(1000000), tree_1)
	_ = result_2_0
	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
		actual   bool
		expected bool
	}{result_2_0.valid, true}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t_and_2 bool = false
		if tree_1 != nil {

			var __t_tag_1 uint32 = (tree_1).V0
			_ = __t_tag_1
			__t_and_2 = (uint32(__t_tag_1) == 2247809753)
		}
		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
			actual   bool
			expected bool
		}{__t_and_2, false}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{result_2_0.size, expectedSize_0})
		}))
	}))
}

func Call_Main_snapshots(v_0_loop int64, v1_1_loop *Constructor_Main_Branch) gopurs_runtime.Value {
snapshots:
	for {
		if false {
			continue snapshots
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Main_Branch = v1_1_loop
		_ = v1_1
		var __t7 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t7 = Call_Main_check(int64(32), v1_1)
				goto end_branch_7
			} else {

			}
		}
		{
			__t7 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
				__local_var_2_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(v1_1)})
				_ = __local_var_2_0
				__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
				_ = __local_var_3_1
				__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_3_1), gopurs_runtime.Value{})
				_ = __local_var_4_2
				__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_4_2)))), gopurs_runtime.Value{})
				_ = __local_var_5_3
				__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_3), gopurs_runtime.Value{})
				_ = __local_var_6_4
				__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_put(v_0, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_4_2)))}), gopurs_runtime.Value{})
				_ = __local_var_7_5
				__local_var_8_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_7_5), gopurs_runtime.Value{})
				_ = __local_var_8_6
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check((int64(33))-(v_0), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_8_6)), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_snapshots((v_0)-(int64(1)), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_8_6)), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_Test_Assert_assertEqual_prime___4176622598("", struct {
							actual   string
							expected string
						}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_4_2)), __local_var_6_4.StrVal()})
					}))
				})), gopurs_runtime.Value{})
			})
		}
	end_branch_7:
		return __t7
	}
}

func Call_Main_three(a_0_loop int64, b_1_loop int64, c_2_loop int64) gopurs_runtime.Value {
	var a_0 int64 = a_0_loop
	_ = a_0
	var b_1 int64 = b_1_loop
	_ = b_1
	var c_2 int64 = c_2_loop
	_ = c_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
			orig := struct {
				a int64
				b int64
				c int64
			}{a_0, b_1, c_2}
			_ = orig
			return gopurs_runtime.RecordDict3("a", "b", "c", gopurs_runtime.Int(orig.a), gopurs_runtime.Int(orig.b), gopurs_runtime.Int(orig.c))
		}())
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
		_ = __local_var_5_2
		__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_put(gopurs_runtime.RecordGet(__local_var_5_2, "c").IntVal, Call_Main_put(gopurs_runtime.RecordGet(__local_var_5_2, "b").IntVal, Call_Main_put(gopurs_runtime.RecordGet(__local_var_5_2, "a").IntVal, (*Constructor_Main_Branch)(nil)))))}), gopurs_runtime.Value{})
		_ = __local_var_6_3
		__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_6_3), gopurs_runtime.Value{})
		_ = __local_var_7_4
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(int64(3), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_7_4)), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Assert_assertEqual_prime___4176622598("", struct {
				actual   string
				expected string
			}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_7_4)), "B2(B1(E)(E))(B3(E)(E))"})
		})), gopurs_runtime.Value{})
	})
}
