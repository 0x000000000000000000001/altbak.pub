package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_RBTree_logShow gopurs_runtime.Value
var once_Test_RBTree_logShow sync.Once

func Get_Test_RBTree_logShow() gopurs_runtime.Value {
	once_Test_RBTree_logShow.Do(func() {
		cache_Test_RBTree_logShow = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_RBTree_logShow(a_0_box.IntVal)
		})
	})
	return cache_Test_RBTree_logShow
}

var cache_Test_RBTree_R gopurs_runtime.Value
var once_Test_RBTree_R sync.Once

func Get_Test_RBTree_R() gopurs_runtime.Value {
	once_Test_RBTree_R.Do(func() {
		cache_Test_RBTree_R = gopurs_runtime.Value{Type: 9, IntVal: int64(3668501016), UnsafePtr: nil}
	})
	return cache_Test_RBTree_R
}

var cache_Test_RBTree_B gopurs_runtime.Value
var once_Test_RBTree_B sync.Once

func Get_Test_RBTree_B() gopurs_runtime.Value {
	once_Test_RBTree_B.Do(func() {
		cache_Test_RBTree_B = gopurs_runtime.Value{Type: 9, IntVal: int64(1583507464), UnsafePtr: nil}
	})
	return cache_Test_RBTree_B
}

var cache_Test_RBTree_E gopurs_runtime.Value
var once_Test_RBTree_E sync.Once

func Get_Test_RBTree_E() gopurs_runtime.Value {
	once_Test_RBTree_E.Do(func() {
		cache_Test_RBTree_E = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_Test_RBTree_T)(nil))}
	})
	return cache_Test_RBTree_E
}

var cache_Test_RBTree_T gopurs_runtime.Value
var once_Test_RBTree_T sync.Once

func Get_Test_RBTree_T() gopurs_runtime.Value {
	once_Test_RBTree_T.Do(func() {
		cache_Test_RBTree_T = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((&Constructor_Test_RBTree_T{1, uint32(value0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](value1), value2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](value3)}))}
					})
				})
			})
		})
	})
	return cache_Test_RBTree_T
}

var cache_Test_RBTree_max gopurs_runtime.Value
var once_Test_RBTree_max sync.Once

func Get_Test_RBTree_max() gopurs_runtime.Value {
	once_Test_RBTree_max.Do(func() {
		cache_Test_RBTree_max = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_RBTree_max(x_0_box.IntVal, y_1_box.IntVal))
		})
	})
	return cache_Test_RBTree_max
}

var cache_Test_RBTree_makeBlack gopurs_runtime.Value
var once_Test_RBTree_makeBlack sync.Once

func Get_Test_RBTree_makeBlack() gopurs_runtime.Value {
	once_Test_RBTree_makeBlack.Do(func() {
		cache_Test_RBTree_makeBlack = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_Test_RBTree_makeBlack(gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v_0_box)))}
		})
	})
	return cache_Test_RBTree_makeBlack
}

var cache_Test_RBTree_describe gopurs_runtime.Value
var once_Test_RBTree_describe sync.Once

func Get_Test_RBTree_describe() gopurs_runtime.Value {
	once_Test_RBTree_describe.Do(func() {
		cache_Test_RBTree_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Red-Black Tree (100k Worst-Case Insertions):"))
	})
	return cache_Test_RBTree_describe
}

var cache_Test_RBTree_depth gopurs_runtime.Value
var once_Test_RBTree_depth sync.Once

func Get_Test_RBTree_depth() gopurs_runtime.Value {
	once_Test_RBTree_depth.Do(func() {
		cache_Test_RBTree_depth = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_RBTree_depth(gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v_0_box)))
		})
	})
	return cache_Test_RBTree_depth
}

var cache_Test_RBTree_balance gopurs_runtime.Value
var once_Test_RBTree_balance sync.Once

func Get_Test_RBTree_balance() gopurs_runtime.Value {
	once_Test_RBTree_balance.Do(func() {
		cache_Test_RBTree_balance = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_Test_RBTree_balance(uint32(v_0_box.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v1_1_box), v2_2_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v3_3_box)))}
		})
	})
	return cache_Test_RBTree_balance
}

var cache_Test_RBTree_ins gopurs_runtime.Value
var once_Test_RBTree_ins sync.Once

func Get_Test_RBTree_ins() gopurs_runtime.Value {
	once_Test_RBTree_ins.Do(func() {
		cache_Test_RBTree_ins = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_Test_RBTree_ins(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v1_1_box)))}
		})
	})
	return cache_Test_RBTree_ins
}

var cache_Test_RBTree_insert gopurs_runtime.Value
var once_Test_RBTree_insert sync.Once

func Get_Test_RBTree_insert() gopurs_runtime.Value {
	once_Test_RBTree_insert.Do(func() {
		cache_Test_RBTree_insert = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_Test_RBTree_insert(x_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](s_1_box)))}
		})
	})
	return cache_Test_RBTree_insert
}

var cache_Test_RBTree_buildTree gopurs_runtime.Value
var once_Test_RBTree_buildTree sync.Once

func Get_Test_RBTree_buildTree() gopurs_runtime.Value {
	once_Test_RBTree_buildTree.Do(func() {
		cache_Test_RBTree_buildTree = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_Test_RBTree_buildTree(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](v1_1_box)))}
		})
	})
	return cache_Test_RBTree_buildTree
}

var cache_Test_RBTree_act gopurs_runtime.Value
var once_Test_RBTree_act sync.Once

func Get_Test_RBTree_act() gopurs_runtime.Value {
	once_Test_RBTree_act.Do(func() {
		cache_Test_RBTree_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(100000))
			_ = __local_var_0_0
			dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = dummy_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_RBTree_depth(Call_Test_RBTree_buildTree(dummy_1_1.IntVal, (*Constructor_Test_RBTree_T)(nil))))).StrVal())), gopurs_runtime.Value{})
		})
	})
	return cache_Test_RBTree_act
}

type Constructor_Test_RBTree_R struct {
	Rc uint32
}

type Constructor_Test_RBTree_B struct {
	Rc uint32
}

type Constructor_Test_RBTree_E struct {
	Rc uint32
}

type Constructor_Test_RBTree_T struct {
	Rc uint32
	V0 uint32
	V1 *Constructor_Test_RBTree_T
	V2 int64
	V3 *Constructor_Test_RBTree_T
}

func Call_Test_RBTree_logShow(a_0_loop int64) gopurs_runtime.Value {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(a_0)).StrVal()))
}

func Call_Test_RBTree_max(x_0_loop int64, y_1_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	var y_1 int64 = y_1_loop
	_ = y_1
	var __t0 int64
	{
		if (x_0) > (y_1) {
			__t0 = x_0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = y_1
	}
end_branch_0:
	return __t0
}

func Call_Test_RBTree_makeBlack(v_0_loop *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
	var v_0 *Constructor_Test_RBTree_T = v_0_loop
	_ = v_0
	var __t0 *Constructor_Test_RBTree_T
	{
		if v_0 != nil {
			__t0 = (&Constructor_Test_RBTree_T{1, 1583507464, (v_0).V1, (v_0).V2, (v_0).V3})
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == nil {
			__t0 = (*Constructor_Test_RBTree_T)(nil)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](func() gopurs_runtime.Value { panic("Failed pattern match") }())
	}
end_branch_0:
	return __t0
}

func Call_Test_RBTree_depth(v_0_loop *Constructor_Test_RBTree_T) int64 {
depth:
	for {
		if false {
			continue depth
		}
		var v_0 *Constructor_Test_RBTree_T = v_0_loop
		_ = v_0
		var __t3 int64
		{
			if v_0 == nil {
				__t3 = 0
				goto end_branch_3
			} else {

			}
		}
		{
			if v_0 != nil {
				__local_var_1_0 := Call_Test_RBTree_depth((v_0).V1)
				_ = __local_var_1_0
				__local_var_2_1 := Call_Test_RBTree_depth((v_0).V3)
				_ = __local_var_2_1
				var __t2 int64
				{
					if (__local_var_1_0) > (__local_var_2_1) {
						__t2 = __local_var_1_0
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = __local_var_2_1
				}
			end_branch_2:
				__t3 = (1) + (__t2)
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
		}
	end_branch_3:
		return __t3
	}
}

func Call_Test_RBTree_balance(v_0_loop uint32, v1_1_loop *Constructor_Test_RBTree_T, v2_2_loop int64, v3_3_loop *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var v1_1 *Constructor_Test_RBTree_T = v1_1_loop
	_ = v1_1
	var v2_2 int64 = v2_2_loop
	_ = v2_2
	var v3_3 *Constructor_Test_RBTree_T = v3_3_loop
	_ = v3_3
	var __t85 *Constructor_Test_RBTree_T
	{
		if v_0 == 1583507464 {
			var __t84 *Constructor_Test_RBTree_T
			{
				if v1_1 != nil {
					var __t71 *Constructor_Test_RBTree_T
					{
						var __t_tag_0 uint32 = (v1_1).V0
						if uint32(__t_tag_0) == 3668501016 {
							var __t58 *Constructor_Test_RBTree_T
							{
								var __t_tag_1 *Constructor_Test_RBTree_T = (v1_1).V1
								if __t_tag_1 != nil {
									var __t30 *Constructor_Test_RBTree_T
									{
										var __t_tag_2 uint32 = ((v1_1).V1).V0
										if uint32(__t_tag_2) == 3668501016 {
											__t30 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, ((v1_1).V1).V1, ((v1_1).V1).V2, ((v1_1).V1).V3}), (v1_1).V2, (&Constructor_Test_RBTree_T{1, 1583507464, (v1_1).V3, v2_2, v3_3})})
											goto end_branch_30
										} else {

										}
									}
									{
										var __t_tag_3 *Constructor_Test_RBTree_T = (v1_1).V3
										if __t_tag_3 != nil {
											var __t17 *Constructor_Test_RBTree_T
											{
												var __t_tag_4 uint32 = ((v1_1).V3).V0
												if uint32(__t_tag_4) == 3668501016 {
													__t17 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, (v1_1).V1, (v1_1).V2, ((v1_1).V3).V1}), ((v1_1).V3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v1_1).V3).V3, v2_2, v3_3})})
													goto end_branch_17
												} else {

												}
											}
											{
												var __t_and_6 bool = false
												if v3_3 != nil {

													var __t_tag_5 uint32 = (v3_3).V0
													__t_and_6 = (uint32(__t_tag_5) == 3668501016)
												}
												if __t_and_6 {
													var __t16 *Constructor_Test_RBTree_T
													{
														var __t_tag_7 *Constructor_Test_RBTree_T = (v3_3).V1
														if __t_tag_7 != nil {
															var __t12 *Constructor_Test_RBTree_T
															{
																var __t_tag_8 uint32 = ((v3_3).V1).V0
																if uint32(__t_tag_8) == 3668501016 {
																	__t12 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
																	goto end_branch_12
																} else {

																}
															}
															{
																var __t_tag_9 *Constructor_Test_RBTree_T = (v3_3).V3
																var __t_and_11 bool = false
																if __t_tag_9 != nil {

																	var __t_tag_10 uint32 = ((v3_3).V3).V0
																	__t_and_11 = (uint32(__t_tag_10) == 3668501016)
																}
																if __t_and_11 {
																	__t12 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
																	goto end_branch_12
																} else {

																}
															}
															{
																__t12 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
															}
														end_branch_12:
															__t16 = __t12
															goto end_branch_16
														} else {

														}
													}
													{
														var __t_tag_13 *Constructor_Test_RBTree_T = (v3_3).V3
														var __t_and_15 bool = false
														if __t_tag_13 != nil {

															var __t_tag_14 uint32 = ((v3_3).V3).V0
															__t_and_15 = (uint32(__t_tag_14) == 3668501016)
														}
														if __t_and_15 {
															__t16 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
															goto end_branch_16
														} else {

														}
													}
													{
														__t16 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
													}
												end_branch_16:
													__t17 = __t16
													goto end_branch_17
												} else {

												}
											}
											{
												__t17 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
											}
										end_branch_17:
											__t30 = __t17
											goto end_branch_30
										} else {

										}
									}
									{
										var __t_and_19 bool = false
										if v3_3 != nil {

											var __t_tag_18 uint32 = (v3_3).V0
											__t_and_19 = (uint32(__t_tag_18) == 3668501016)
										}
										if __t_and_19 {
											var __t29 *Constructor_Test_RBTree_T
											{
												var __t_tag_20 *Constructor_Test_RBTree_T = (v3_3).V1
												if __t_tag_20 != nil {
													var __t25 *Constructor_Test_RBTree_T
													{
														var __t_tag_21 uint32 = ((v3_3).V1).V0
														if uint32(__t_tag_21) == 3668501016 {
															__t25 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
															goto end_branch_25
														} else {

														}
													}
													{
														var __t_tag_22 *Constructor_Test_RBTree_T = (v3_3).V3
														var __t_and_24 bool = false
														if __t_tag_22 != nil {

															var __t_tag_23 uint32 = ((v3_3).V3).V0
															__t_and_24 = (uint32(__t_tag_23) == 3668501016)
														}
														if __t_and_24 {
															__t25 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
															goto end_branch_25
														} else {

														}
													}
													{
														__t25 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
													}
												end_branch_25:
													__t29 = __t25
													goto end_branch_29
												} else {

												}
											}
											{
												var __t_tag_26 *Constructor_Test_RBTree_T = (v3_3).V3
												var __t_and_28 bool = false
												if __t_tag_26 != nil {

													var __t_tag_27 uint32 = ((v3_3).V3).V0
													__t_and_28 = (uint32(__t_tag_27) == 3668501016)
												}
												if __t_and_28 {
													__t29 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
													goto end_branch_29
												} else {

												}
											}
											{
												__t29 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
											}
										end_branch_29:
											__t30 = __t29
											goto end_branch_30
										} else {

										}
									}
									{
										__t30 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
									}
								end_branch_30:
									__t58 = __t30
									goto end_branch_58
								} else {

								}
							}
							{
								var __t_tag_31 *Constructor_Test_RBTree_T = (v1_1).V3
								if __t_tag_31 != nil {
									var __t45 *Constructor_Test_RBTree_T
									{
										var __t_tag_32 uint32 = ((v1_1).V3).V0
										if uint32(__t_tag_32) == 3668501016 {
											__t45 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, (v1_1).V1, (v1_1).V2, ((v1_1).V3).V1}), ((v1_1).V3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v1_1).V3).V3, v2_2, v3_3})})
											goto end_branch_45
										} else {

										}
									}
									{
										var __t_and_34 bool = false
										if v3_3 != nil {

											var __t_tag_33 uint32 = (v3_3).V0
											__t_and_34 = (uint32(__t_tag_33) == 3668501016)
										}
										if __t_and_34 {
											var __t44 *Constructor_Test_RBTree_T
											{
												var __t_tag_35 *Constructor_Test_RBTree_T = (v3_3).V1
												if __t_tag_35 != nil {
													var __t40 *Constructor_Test_RBTree_T
													{
														var __t_tag_36 uint32 = ((v3_3).V1).V0
														if uint32(__t_tag_36) == 3668501016 {
															__t40 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
															goto end_branch_40
														} else {

														}
													}
													{
														var __t_tag_37 *Constructor_Test_RBTree_T = (v3_3).V3
														var __t_and_39 bool = false
														if __t_tag_37 != nil {

															var __t_tag_38 uint32 = ((v3_3).V3).V0
															__t_and_39 = (uint32(__t_tag_38) == 3668501016)
														}
														if __t_and_39 {
															__t40 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
															goto end_branch_40
														} else {

														}
													}
													{
														__t40 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
													}
												end_branch_40:
													__t44 = __t40
													goto end_branch_44
												} else {

												}
											}
											{
												var __t_tag_41 *Constructor_Test_RBTree_T = (v3_3).V3
												var __t_and_43 bool = false
												if __t_tag_41 != nil {

													var __t_tag_42 uint32 = ((v3_3).V3).V0
													__t_and_43 = (uint32(__t_tag_42) == 3668501016)
												}
												if __t_and_43 {
													__t44 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
													goto end_branch_44
												} else {

												}
											}
											{
												__t44 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
											}
										end_branch_44:
											__t45 = __t44
											goto end_branch_45
										} else {

										}
									}
									{
										__t45 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
									}
								end_branch_45:
									__t58 = __t45
									goto end_branch_58
								} else {

								}
							}
							{
								var __t_and_47 bool = false
								if v3_3 != nil {

									var __t_tag_46 uint32 = (v3_3).V0
									__t_and_47 = (uint32(__t_tag_46) == 3668501016)
								}
								if __t_and_47 {
									var __t57 *Constructor_Test_RBTree_T
									{
										var __t_tag_48 *Constructor_Test_RBTree_T = (v3_3).V1
										if __t_tag_48 != nil {
											var __t53 *Constructor_Test_RBTree_T
											{
												var __t_tag_49 uint32 = ((v3_3).V1).V0
												if uint32(__t_tag_49) == 3668501016 {
													__t53 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
													goto end_branch_53
												} else {

												}
											}
											{
												var __t_tag_50 *Constructor_Test_RBTree_T = (v3_3).V3
												var __t_and_52 bool = false
												if __t_tag_50 != nil {

													var __t_tag_51 uint32 = ((v3_3).V3).V0
													__t_and_52 = (uint32(__t_tag_51) == 3668501016)
												}
												if __t_and_52 {
													__t53 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
													goto end_branch_53
												} else {

												}
											}
											{
												__t53 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
											}
										end_branch_53:
											__t57 = __t53
											goto end_branch_57
										} else {

										}
									}
									{
										var __t_tag_54 *Constructor_Test_RBTree_T = (v3_3).V3
										var __t_and_56 bool = false
										if __t_tag_54 != nil {

											var __t_tag_55 uint32 = ((v3_3).V3).V0
											__t_and_56 = (uint32(__t_tag_55) == 3668501016)
										}
										if __t_and_56 {
											__t57 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
											goto end_branch_57
										} else {

										}
									}
									{
										__t57 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
									}
								end_branch_57:
									__t58 = __t57
									goto end_branch_58
								} else {

								}
							}
							{
								__t58 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
							}
						end_branch_58:
							__t71 = __t58
							goto end_branch_71
						} else {

						}
					}
					{
						var __t_and_60 bool = false
						if v3_3 != nil {

							var __t_tag_59 uint32 = (v3_3).V0
							__t_and_60 = (uint32(__t_tag_59) == 3668501016)
						}
						if __t_and_60 {
							var __t70 *Constructor_Test_RBTree_T
							{
								var __t_tag_61 *Constructor_Test_RBTree_T = (v3_3).V1
								if __t_tag_61 != nil {
									var __t66 *Constructor_Test_RBTree_T
									{
										var __t_tag_62 uint32 = ((v3_3).V1).V0
										if uint32(__t_tag_62) == 3668501016 {
											__t66 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
											goto end_branch_66
										} else {

										}
									}
									{
										var __t_tag_63 *Constructor_Test_RBTree_T = (v3_3).V3
										var __t_and_65 bool = false
										if __t_tag_63 != nil {

											var __t_tag_64 uint32 = ((v3_3).V3).V0
											__t_and_65 = (uint32(__t_tag_64) == 3668501016)
										}
										if __t_and_65 {
											__t66 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
											goto end_branch_66
										} else {

										}
									}
									{
										__t66 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
									}
								end_branch_66:
									__t70 = __t66
									goto end_branch_70
								} else {

								}
							}
							{
								var __t_tag_67 *Constructor_Test_RBTree_T = (v3_3).V3
								var __t_and_69 bool = false
								if __t_tag_67 != nil {

									var __t_tag_68 uint32 = ((v3_3).V3).V0
									__t_and_69 = (uint32(__t_tag_68) == 3668501016)
								}
								if __t_and_69 {
									__t70 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
									goto end_branch_70
								} else {

								}
							}
							{
								__t70 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
							}
						end_branch_70:
							__t71 = __t70
							goto end_branch_71
						} else {

						}
					}
					{
						__t71 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
					}
				end_branch_71:
					__t84 = __t71
					goto end_branch_84
				} else {

				}
			}
			{
				var __t_and_73 bool = false
				if v3_3 != nil {

					var __t_tag_72 uint32 = (v3_3).V0
					__t_and_73 = (uint32(__t_tag_72) == 3668501016)
				}
				if __t_and_73 {
					var __t83 *Constructor_Test_RBTree_T
					{
						var __t_tag_74 *Constructor_Test_RBTree_T = (v3_3).V1
						if __t_tag_74 != nil {
							var __t79 *Constructor_Test_RBTree_T
							{
								var __t_tag_75 uint32 = ((v3_3).V1).V0
								if uint32(__t_tag_75) == 3668501016 {
									__t79 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
									goto end_branch_79
								} else {

								}
							}
							{
								var __t_tag_76 *Constructor_Test_RBTree_T = (v3_3).V3
								var __t_and_78 bool = false
								if __t_tag_76 != nil {

									var __t_tag_77 uint32 = ((v3_3).V3).V0
									__t_and_78 = (uint32(__t_tag_77) == 3668501016)
								}
								if __t_and_78 {
									__t79 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
									goto end_branch_79
								} else {

								}
							}
							{
								__t79 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
							}
						end_branch_79:
							__t83 = __t79
							goto end_branch_83
						} else {

						}
					}
					{
						var __t_tag_80 *Constructor_Test_RBTree_T = (v3_3).V3
						var __t_and_82 bool = false
						if __t_tag_80 != nil {

							var __t_tag_81 uint32 = ((v3_3).V3).V0
							__t_and_82 = (uint32(__t_tag_81) == 3668501016)
						}
						if __t_and_82 {
							__t83 = (&Constructor_Test_RBTree_T{1, 3668501016, (&Constructor_Test_RBTree_T{1, 1583507464, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Test_RBTree_T{1, 1583507464, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
							goto end_branch_83
						} else {

						}
					}
					{
						__t83 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
					}
				end_branch_83:
					__t84 = __t83
					goto end_branch_84
				} else {

				}
			}
			{
				__t84 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
			}
		end_branch_84:
			__t85 = __t84
			goto end_branch_85
		} else {

		}
	}
	{
		__t85 = (&Constructor_Test_RBTree_T{1, v_0, v1_1, v2_2, v3_3})
	}
end_branch_85:
	return __t85
}

func Call_Test_RBTree_ins(v_0_loop int64, v1_1_loop *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
ins:
	for {
		if false {
			continue ins
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Test_RBTree_T = v1_1_loop
		_ = v1_1
		var __t2 *Constructor_Test_RBTree_T
		{
			if v1_1 == nil {
				__t2 = (&Constructor_Test_RBTree_T{1, 3668501016, (*Constructor_Test_RBTree_T)(nil), v_0, (*Constructor_Test_RBTree_T)(nil)})
				goto end_branch_2
			} else {

			}
		}
		{
			if v1_1 != nil {
				var __t1 *Constructor_Test_RBTree_T
				{
					if (v_0) < ((v1_1).V2) {
						__t1 = Call_Test_RBTree_balance((v1_1).V0, Call_Test_RBTree_ins(v_0, (v1_1).V1), (v1_1).V2, (v1_1).V3)
						goto end_branch_1
					} else {

					}
				}
				{
					var __t0 *Constructor_Test_RBTree_T
					{
						if (v_0) > ((v1_1).V2) {
							__t0 = Call_Test_RBTree_balance((v1_1).V0, (v1_1).V1, (v1_1).V2, Call_Test_RBTree_ins(v_0, (v1_1).V3))
							goto end_branch_0
						} else {

						}
					}
					{
						__t0 = (&Constructor_Test_RBTree_T{1, (v1_1).V0, (v1_1).V1, (v1_1).V2, (v1_1).V3})
					}
				end_branch_0:
					__t1 = __t0
				}
			end_branch_1:
				__t2 = __t1
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](func() gopurs_runtime.Value { panic("Failed pattern match") }())
		}
	end_branch_2:
		return __t2
	}
}

func Call_Test_RBTree_insert(x_0_loop int64, s_1_loop *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
	var x_0 int64 = x_0_loop
	_ = x_0
	var s_1 *Constructor_Test_RBTree_T = s_1_loop
	_ = s_1
	__local_var_2_0 := Call_Test_RBTree_ins(x_0, s_1)
	_ = __local_var_2_0
	var __t1 *Constructor_Test_RBTree_T
	{
		if __local_var_2_0 != nil {
			__t1 = (&Constructor_Test_RBTree_T{1, 1583507464, (__local_var_2_0).V1, (__local_var_2_0).V2, (__local_var_2_0).V3})
			goto end_branch_1
		} else {

		}
	}
	{
		if __local_var_2_0 == nil {
			__t1 = (*Constructor_Test_RBTree_T)(nil)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](func() gopurs_runtime.Value { panic("Failed pattern match") }())
	}
end_branch_1:
	return __t1
}

func Call_Test_RBTree_buildTree(v_0_loop int64, v1_1_loop *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
buildTree:
	for {
		if false {
			continue buildTree
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Test_RBTree_T = v1_1_loop
		_ = v1_1
		var __t0 *Constructor_Test_RBTree_T
		{
			if (v_0) == (0) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (1)
			v1_1_loop = Call_Test_RBTree_insert(v_0, v1_1)
			continue buildTree
			__t0 = gopurs_runtime.CoerceToStruct[Constructor_Test_RBTree_T](gopurs_runtime.Value{})
		}
	end_branch_0:
		return __t0
	}
}
