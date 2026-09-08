package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Enum_bottom gopurs_runtime.Value
var once_Data_Enum_bottom sync.Once
func Get_Data_Enum_bottom() gopurs_runtime.Value {
	once_Data_Enum_bottom.Do(func() {
		cache_Data_Enum_bottom = gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal)
	})
	return cache_Data_Enum_bottom
}

var cache_Data_Enum_fromJust gopurs_runtime.Value
var once_Data_Enum_fromJust sync.Once
func Get_Data_Enum_fromJust() gopurs_runtime.Value {
	once_Data_Enum_fromJust.Do(func() {
		cache_Data_Enum_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_Data_Enum_fromJust
}

var cache_Data_Enum_bottom1 gopurs_runtime.Value
var once_Data_Enum_bottom1 sync.Once
func Get_Data_Enum_bottom1() gopurs_runtime.Value {
	once_Data_Enum_bottom1.Do(func() {
		cache_Data_Enum_bottom1 = gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())
	})
	return cache_Data_Enum_bottom1
}

var cache_Data_Enum_top gopurs_runtime.Value
var once_Data_Enum_top sync.Once
func Get_Data_Enum_top() gopurs_runtime.Value {
	once_Data_Enum_top.Do(func() {
		cache_Data_Enum_top = gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())
	})
	return cache_Data_Enum_top
}

var cache_Data_Enum_Enum_dollar_Dict gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict sync.Once
func Get_Data_Enum_Enum_dollar_Dict() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Call_Data_Enum_Enum_dollar_Dict(func() struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}{}
					clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
					clone.pred = gopurs_runtime.RecordGet(orig, "pred")
					clone.succ = gopurs_runtime.RecordGet(orig, "succ")
					return clone
				}()))}
})
	})
	return cache_Data_Enum_Enum_dollar_Dict
}

var cache_Data_Enum_Cardinality gopurs_runtime.Value
var once_Data_Enum_Cardinality sync.Once
func Get_Data_Enum_Cardinality() gopurs_runtime.Value {
	once_Data_Enum_Cardinality.Do(func() {
		cache_Data_Enum_Cardinality = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality
}

var cache_Data_Enum_Cardinality__11842167 gopurs_runtime.Value
var once_Data_Enum_Cardinality__11842167 sync.Once
func Get_Data_Enum_Cardinality__11842167() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__11842167.Do(func() {
		cache_Data_Enum_Cardinality__11842167 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__11842167(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__11842167
}

var cache_Data_Enum_Cardinality__3304734055 gopurs_runtime.Value
var once_Data_Enum_Cardinality__3304734055 sync.Once
func Get_Data_Enum_Cardinality__3304734055() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__3304734055.Do(func() {
		cache_Data_Enum_Cardinality__3304734055 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__3304734055(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__3304734055
}

var cache_Data_Enum_Cardinality__4193071403 gopurs_runtime.Value
var once_Data_Enum_Cardinality__4193071403 sync.Once
func Get_Data_Enum_Cardinality__4193071403() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__4193071403.Do(func() {
		cache_Data_Enum_Cardinality__4193071403 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__4193071403(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__4193071403
}

var cache_Data_Enum_Cardinality__485476899 gopurs_runtime.Value
var once_Data_Enum_Cardinality__485476899 sync.Once
func Get_Data_Enum_Cardinality__485476899() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__485476899.Do(func() {
		cache_Data_Enum_Cardinality__485476899 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__485476899(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__485476899
}

var cache_Data_Enum_Cardinality__2083952025 gopurs_runtime.Value
var once_Data_Enum_Cardinality__2083952025 sync.Once
func Get_Data_Enum_Cardinality__2083952025() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__2083952025.Do(func() {
		cache_Data_Enum_Cardinality__2083952025 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__2083952025(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__2083952025
}

var cache_Data_Enum_Cardinality__115045349 gopurs_runtime.Value
var once_Data_Enum_Cardinality__115045349 sync.Once
func Get_Data_Enum_Cardinality__115045349() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__115045349.Do(func() {
		cache_Data_Enum_Cardinality__115045349 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__115045349(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__115045349
}

var cache_Data_Enum_Cardinality__2170980526 gopurs_runtime.Value
var once_Data_Enum_Cardinality__2170980526 sync.Once
func Get_Data_Enum_Cardinality__2170980526() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__2170980526.Do(func() {
		cache_Data_Enum_Cardinality__2170980526 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__2170980526(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__2170980526
}

var cache_Data_Enum_Cardinality__2518865467 gopurs_runtime.Value
var once_Data_Enum_Cardinality__2518865467 sync.Once
func Get_Data_Enum_Cardinality__2518865467() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__2518865467.Do(func() {
		cache_Data_Enum_Cardinality__2518865467 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__2518865467(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__2518865467
}

var cache_Data_Enum_BoundedEnum_dollar_Dict gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict sync.Once
func Get_Data_Enum_BoundedEnum_dollar_Dict() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Call_Data_Enum_BoundedEnum_dollar_Dict(func() struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}{}
					clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
					clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
					clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
					clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
					clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
					return clone
				}()))}
})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict
}

var cache_Data_Enum_toEnum gopurs_runtime.Value
var once_Data_Enum_toEnum sync.Once
func Get_Data_Enum_toEnum() gopurs_runtime.Value {
	once_Data_Enum_toEnum.Do(func() {
		cache_Data_Enum_toEnum = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_toEnum
}

var cache_Data_Enum_toEnum__1948670652 gopurs_runtime.Value
var once_Data_Enum_toEnum__1948670652 sync.Once
func Get_Data_Enum_toEnum__1948670652() gopurs_runtime.Value {
	once_Data_Enum_toEnum__1948670652.Do(func() {
		cache_Data_Enum_toEnum__1948670652 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnum__1948670652(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_toEnum__1948670652
}

var cache_Data_Enum_toEnum__3403136940 gopurs_runtime.Value
var once_Data_Enum_toEnum__3403136940 sync.Once
func Get_Data_Enum_toEnum__3403136940() gopurs_runtime.Value {
	once_Data_Enum_toEnum__3403136940.Do(func() {
		cache_Data_Enum_toEnum__3403136940 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnum__3403136940(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_toEnum__3403136940
}

var cache_Data_Enum_toEnum__106342766 gopurs_runtime.Value
var once_Data_Enum_toEnum__106342766 sync.Once
func Get_Data_Enum_toEnum__106342766() gopurs_runtime.Value {
	once_Data_Enum_toEnum__106342766.Do(func() {
		cache_Data_Enum_toEnum__106342766 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnum__106342766(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_toEnum__106342766
}

var cache_Data_Enum_toEnum__983269893 gopurs_runtime.Value
var once_Data_Enum_toEnum__983269893 sync.Once
func Get_Data_Enum_toEnum__983269893() gopurs_runtime.Value {
	once_Data_Enum_toEnum__983269893.Do(func() {
		cache_Data_Enum_toEnum__983269893 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnum__983269893(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_toEnum__983269893
}

var cache_Data_Enum_succ gopurs_runtime.Value
var once_Data_Enum_succ sync.Once
func Get_Data_Enum_succ() gopurs_runtime.Value {
	once_Data_Enum_succ.Do(func() {
		cache_Data_Enum_succ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_succ(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dict_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_succ
}

var cache_Data_Enum_succ__4151667621 gopurs_runtime.Value
var once_Data_Enum_succ__4151667621 sync.Once
func Get_Data_Enum_succ__4151667621() gopurs_runtime.Value {
	once_Data_Enum_succ__4151667621.Do(func() {
		cache_Data_Enum_succ__4151667621 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_succ__4151667621(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_succ__4151667621
}

var cache_Data_Enum_succ__3996651813 gopurs_runtime.Value
var once_Data_Enum_succ__3996651813 sync.Once
func Get_Data_Enum_succ__3996651813() gopurs_runtime.Value {
	once_Data_Enum_succ__3996651813.Do(func() {
		cache_Data_Enum_succ__3996651813 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_succ__3996651813(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_succ__3996651813
}

var cache_Data_Enum_succ__983269893 gopurs_runtime.Value
var once_Data_Enum_succ__983269893 sync.Once
func Get_Data_Enum_succ__983269893() gopurs_runtime.Value {
	once_Data_Enum_succ__983269893.Do(func() {
		cache_Data_Enum_succ__983269893 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_succ__983269893(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_succ__983269893
}

var cache_Data_Enum_upFromIncluding gopurs_runtime.Value
var once_Data_Enum_upFromIncluding sync.Once
func Get_Data_Enum_upFromIncluding() gopurs_runtime.Value {
	once_Data_Enum_upFromIncluding.Do(func() {
		cache_Data_Enum_upFromIncluding = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_upFromIncluding(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_1_box))
})
	})
	return cache_Data_Enum_upFromIncluding
}

var cache_Data_Enum_showCardinality gopurs_runtime.Value
var once_Data_Enum_showCardinality sync.Once
func Get_Data_Enum_showCardinality() gopurs_runtime.Value {
	once_Data_Enum_showCardinality.Do(func() {
		cache_Data_Enum_showCardinality = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Cardinality ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), v_0).StrVal())) + (")"))
})}))}
	})
	return cache_Data_Enum_showCardinality
}

var cache_Data_Enum_pred gopurs_runtime.Value
var once_Data_Enum_pred sync.Once
func Get_Data_Enum_pred() gopurs_runtime.Value {
	once_Data_Enum_pred.Do(func() {
		cache_Data_Enum_pred = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_pred(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dict_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_pred
}

var cache_Data_Enum_pred__4151667621 gopurs_runtime.Value
var once_Data_Enum_pred__4151667621 sync.Once
func Get_Data_Enum_pred__4151667621() gopurs_runtime.Value {
	once_Data_Enum_pred__4151667621.Do(func() {
		cache_Data_Enum_pred__4151667621 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_pred__4151667621(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_pred__4151667621
}

var cache_Data_Enum_pred__983269893 gopurs_runtime.Value
var once_Data_Enum_pred__983269893 sync.Once
func Get_Data_Enum_pred__983269893() gopurs_runtime.Value {
	once_Data_Enum_pred__983269893.Do(func() {
		cache_Data_Enum_pred__983269893 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_pred__983269893(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_pred__983269893
}

var cache_Data_Enum_ordCardinality gopurs_runtime.Value
var once_Data_Enum_ordCardinality sync.Once
func Get_Data_Enum_ordCardinality() gopurs_runtime.Value {
	once_Data_Enum_ordCardinality.Do(func() {
		cache_Data_Enum_ordCardinality = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3308271157_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt())))}
	})
	return cache_Data_Enum_ordCardinality
}

var cache_Data_Enum_newtypeCardinality gopurs_runtime.Value
var once_Data_Enum_newtypeCardinality sync.Once
func Get_Data_Enum_newtypeCardinality() gopurs_runtime.Value {
	once_Data_Enum_newtypeCardinality.Do(func() {
		cache_Data_Enum_newtypeCardinality = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2621187955_385277032((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})))}
	})
	return cache_Data_Enum_newtypeCardinality
}

var cache_Data_Enum_fromEnum gopurs_runtime.Value
var once_Data_Enum_fromEnum sync.Once
func Get_Data_Enum_fromEnum() gopurs_runtime.Value {
	once_Data_Enum_fromEnum.Do(func() {
		cache_Data_Enum_fromEnum = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Enum_fromEnum
}

var cache_Data_Enum_fromEnum__557243613 gopurs_runtime.Value
var once_Data_Enum_fromEnum__557243613 sync.Once
func Get_Data_Enum_fromEnum__557243613() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__557243613.Do(func() {
		cache_Data_Enum_fromEnum__557243613 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromEnum__557243613(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_fromEnum__557243613
}

var cache_Data_Enum_fromEnum__4291755151 gopurs_runtime.Value
var once_Data_Enum_fromEnum__4291755151 sync.Once
func Get_Data_Enum_fromEnum__4291755151() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__4291755151.Do(func() {
		cache_Data_Enum_fromEnum__4291755151 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromEnum__4291755151(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_fromEnum__4291755151
}

var cache_Data_Enum_fromEnum__4163941923 gopurs_runtime.Value
var once_Data_Enum_fromEnum__4163941923 sync.Once
func Get_Data_Enum_fromEnum__4163941923() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__4163941923.Do(func() {
		cache_Data_Enum_fromEnum__4163941923 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromEnum__4163941923(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Enum_fromEnum__4163941923
}

var cache_Data_Enum_fromEnum__3883318852 gopurs_runtime.Value
var once_Data_Enum_fromEnum__3883318852 sync.Once
func Get_Data_Enum_fromEnum__3883318852() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__3883318852.Do(func() {
		cache_Data_Enum_fromEnum__3883318852 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromEnum__3883318852(__eta_norm_0_0_box)
})
	})
	return cache_Data_Enum_fromEnum__3883318852
}

var cache_Data_Enum_toEnumWithDefaults gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults sync.Once
func Get_Data_Enum_toEnumWithDefaults() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults.Do(func() {
		cache_Data_Enum_toEnumWithDefaults = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnumWithDefaults(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dictBoundedEnum_0_box))
})
	})
	return cache_Data_Enum_toEnumWithDefaults
}

var cache_Data_Enum_toEnumWithDefaults__4181113775 gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults__4181113775 sync.Once
func Get_Data_Enum_toEnumWithDefaults__4181113775() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults__4181113775.Do(func() {
		cache_Data_Enum_toEnumWithDefaults__4181113775 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnumWithDefaults__4181113775(__eta_norm_1_0_box, __eta_norm_0_1_box, x_2_box)
})
	})
	return cache_Data_Enum_toEnumWithDefaults__4181113775
}

var cache_Data_Enum_eqCardinality gopurs_runtime.Value
var once_Data_Enum_eqCardinality sync.Once
func Get_Data_Enum_eqCardinality() gopurs_runtime.Value {
	once_Data_Enum_eqCardinality.Do(func() {
		cache_Data_Enum_eqCardinality = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}
	})
	return cache_Data_Enum_eqCardinality
}

var cache_Data_Enum_enumUnit gopurs_runtime.Value
var once_Data_Enum_enumUnit sync.Once
func Get_Data_Enum_enumUnit() gopurs_runtime.Value {
	once_Data_Enum_enumUnit.Do(func() {
		cache_Data_Enum_enumUnit = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Enum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordUnit()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}
})}))}
	})
	return cache_Data_Enum_enumUnit
}

var cache_Data_Enum_enumTuple gopurs_runtime.Value
var once_Data_Enum_enumTuple sync.Once
func Get_Data_Enum_enumTuple() gopurs_runtime.Value {
	once_Data_Enum_enumTuple.Do(func() {
		cache_Data_Enum_enumTuple = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumTuple(dictEnum_0_box)
})
	})
	return cache_Data_Enum_enumTuple
}

var cache_Data_Enum_enumOrdering gopurs_runtime.Value
var once_Data_Enum_enumOrdering sync.Once
func Get_Data_Enum_enumOrdering() gopurs_runtime.Value {
	once_Data_Enum_enumOrdering.Do(func() {
		cache_Data_Enum_enumOrdering = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094((&Constructor_Data_Enum_Enum[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3730953251_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Ord_ordOrdering())))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_0) == 1527465420) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_3
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_1) == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}}))}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_2) == 380165415) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}}))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(func() *Constructor_Data_Maybe_Just[uint32] { panic("Failed pattern match") }()))}
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_tag_4 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_4) == 1527465420) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}}))}
goto end_branch_7
} else {

}
}
{
var __t_tag_5 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_5) == 902936544) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}}))}
goto end_branch_7
} else {

}
}
{
var __t_tag_6 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_6) == 380165415) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(func() *Constructor_Data_Maybe_Just[uint32] { panic("Failed pattern match") }()))}
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t7))))}
})})))}
	})
	return cache_Data_Enum_enumOrdering
}

var cache_Data_Enum_enumMaybe gopurs_runtime.Value
var once_Data_Enum_enumMaybe sync.Once
func Get_Data_Enum_enumMaybe() gopurs_runtime.Value {
	once_Data_Enum_enumMaybe.Do(func() {
		cache_Data_Enum_enumMaybe = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumMaybe(dictBoundedEnum_0_box)
})
	})
	return cache_Data_Enum_enumMaybe
}

var cache_Data_Enum_enumInt gopurs_runtime.Value
var once_Data_Enum_enumInt sync.Once
func Get_Data_Enum_enumInt() gopurs_runtime.Value {
	once_Data_Enum_enumInt.Do(func() {
		cache_Data_Enum_enumInt = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4060049525_556578094((&Constructor_Data_Enum_Enum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3308271157_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt())))}
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (n_0.IntVal) > (Get_Data_Bounded_bottomInt().IntVal) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((n_0.IntVal) - (int64(1)))}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (n_0.IntVal) < (Get_Data_Bounded_topInt().IntVal) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((n_0.IntVal) + (int64(1)))}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
})})))}
	})
	return cache_Data_Enum_enumInt
}

var cache_Data_Enum_enumFromTo gopurs_runtime.Value
var once_Data_Enum_enumFromTo sync.Once
func Get_Data_Enum_enumFromTo() gopurs_runtime.Value {
	once_Data_Enum_enumFromTo.Do(func() {
		cache_Data_Enum_enumFromTo = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumFromTo(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
})
	})
	return cache_Data_Enum_enumFromTo
}

var cache_Data_Enum_enumFromThenTo gopurs_runtime.Value
var once_Data_Enum_enumFromThenTo sync.Once
func Get_Data_Enum_enumFromThenTo() gopurs_runtime.Value {
	once_Data_Enum_enumFromThenTo.Do(func() {
		cache_Data_Enum_enumFromThenTo = gopurs_runtime.Func6(func(dictUnfoldable_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictBoundedEnum_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value, b_4_box gopurs_runtime.Value, c_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumFromThenTo(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dictBoundedEnum_2_box), a_3_box, b_4_box, c_5_box)
})
	})
	return cache_Data_Enum_enumFromThenTo
}

var cache_Data_Enum_enumEither gopurs_runtime.Value
var once_Data_Enum_enumEither sync.Once
func Get_Data_Enum_enumEither() gopurs_runtime.Value {
	once_Data_Enum_enumEither.Do(func() {
		cache_Data_Enum_enumEither = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumEither(dictBoundedEnum_0_box)
})
	})
	return cache_Data_Enum_enumEither
}

var cache_Data_Enum_enumBoolean gopurs_runtime.Value
var once_Data_Enum_enumBoolean sync.Once
func Get_Data_Enum_enumBoolean() gopurs_runtime.Value {
	once_Data_Enum_enumBoolean.Do(func() {
		cache_Data_Enum_enumBoolean = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_342165130_556578094((&Constructor_Data_Enum_Enum[bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_219188042_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[bool]](Get_Data_Ord_ordBoolean())))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Bool(false)}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1044698560_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[bool]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1044698560_3094389156(Rebox_Data_Enum_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if ((v_0.IntVal) != (0)) != (true) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Bool(true)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1044698560_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[bool]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1044698560_3094389156(Rebox_Data_Enum_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
})})))}
	})
	return cache_Data_Enum_enumBoolean
}

var cache_Data_Enum_downFromIncluding gopurs_runtime.Value
var once_Data_Enum_downFromIncluding sync.Once
func Get_Data_Enum_downFromIncluding() gopurs_runtime.Value {
	once_Data_Enum_downFromIncluding.Do(func() {
		cache_Data_Enum_downFromIncluding = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_downFromIncluding(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_1_box))
})
	})
	return cache_Data_Enum_downFromIncluding
}

var cache_Data_Enum_diag gopurs_runtime.Value
var once_Data_Enum_diag sync.Once
func Get_Data_Enum_diag() gopurs_runtime.Value {
	once_Data_Enum_diag.Do(func() {
		cache_Data_Enum_diag = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_diag(a_0_box)
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Data_Enum_diag
}

var cache_Data_Enum_downFrom gopurs_runtime.Value
var once_Data_Enum_downFrom sync.Once
func Get_Data_Enum_downFrom() gopurs_runtime.Value {
	once_Data_Enum_downFrom.Do(func() {
		cache_Data_Enum_downFrom = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_downFrom(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_1_box))
})
	})
	return cache_Data_Enum_downFrom
}

var cache_Data_Enum_upFrom gopurs_runtime.Value
var once_Data_Enum_upFrom sync.Once
func Get_Data_Enum_upFrom() gopurs_runtime.Value {
	once_Data_Enum_upFrom.Do(func() {
		cache_Data_Enum_upFrom = gopurs_runtime.Func2(func(dictEnum_0_box gopurs_runtime.Value, dictUnfoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_upFrom(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_1_box))
})
	})
	return cache_Data_Enum_upFrom
}

var cache_Data_Enum_defaultToEnum gopurs_runtime.Value
var once_Data_Enum_defaultToEnum sync.Once
func Get_Data_Enum_defaultToEnum() gopurs_runtime.Value {
	once_Data_Enum_defaultToEnum.Do(func() {
		cache_Data_Enum_defaultToEnum = gopurs_runtime.Func3(func(dictBounded_0_box gopurs_runtime.Value, dictEnum_1_box gopurs_runtime.Value, i_prime__2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultToEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](dictBounded_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_1_box), i_prime__2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_defaultToEnum
}

var cache_Data_Enum_defaultSucc gopurs_runtime.Value
var once_Data_Enum_defaultSucc sync.Once
func Get_Data_Enum_defaultSucc() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc.Do(func() {
		cache_Data_Enum_defaultSucc = gopurs_runtime.Func3(func(toEnum_prime__0_box gopurs_runtime.Value, fromEnum_prime__1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultSucc(toEnum_prime__0_box, fromEnum_prime__1_box, a_2_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_defaultSucc
}

var cache_Data_Enum_defaultSucc__2686793765 gopurs_runtime.Value
var once_Data_Enum_defaultSucc__2686793765 sync.Once
func Get_Data_Enum_defaultSucc__2686793765() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc__2686793765.Do(func() {
		cache_Data_Enum_defaultSucc__2686793765 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultSucc__2686793765(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_defaultSucc__2686793765
}

var cache_Data_Enum_defaultSucc__2536258710 gopurs_runtime.Value
var once_Data_Enum_defaultSucc__2536258710 sync.Once
func Get_Data_Enum_defaultSucc__2536258710() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc__2536258710.Do(func() {
		cache_Data_Enum_defaultSucc__2536258710 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultSucc__2536258710(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_defaultSucc__2536258710
}

var cache_Data_Enum_defaultPred gopurs_runtime.Value
var once_Data_Enum_defaultPred sync.Once
func Get_Data_Enum_defaultPred() gopurs_runtime.Value {
	once_Data_Enum_defaultPred.Do(func() {
		cache_Data_Enum_defaultPred = gopurs_runtime.Func3(func(toEnum_prime__0_box gopurs_runtime.Value, fromEnum_prime__1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultPred(toEnum_prime__0_box, fromEnum_prime__1_box, a_2_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_defaultPred
}

var cache_Data_Enum_defaultPred__2686793765 gopurs_runtime.Value
var once_Data_Enum_defaultPred__2686793765 sync.Once
func Get_Data_Enum_defaultPred__2686793765() gopurs_runtime.Value {
	once_Data_Enum_defaultPred__2686793765.Do(func() {
		cache_Data_Enum_defaultPred__2686793765 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultPred__2686793765(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_defaultPred__2686793765
}

var cache_Data_Enum_defaultPred__2536258710 gopurs_runtime.Value
var once_Data_Enum_defaultPred__2536258710 sync.Once
func Get_Data_Enum_defaultPred__2536258710() gopurs_runtime.Value {
	once_Data_Enum_defaultPred__2536258710.Do(func() {
		cache_Data_Enum_defaultPred__2536258710 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultPred__2536258710(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_defaultPred__2536258710
}

var cache_Data_Enum_defaultFromEnum gopurs_runtime.Value
var once_Data_Enum_defaultFromEnum sync.Once
func Get_Data_Enum_defaultFromEnum() gopurs_runtime.Value {
	once_Data_Enum_defaultFromEnum.Do(func() {
		cache_Data_Enum_defaultFromEnum = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_defaultFromEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
})
	})
	return cache_Data_Enum_defaultFromEnum
}

var cache_Data_Enum_defaultCardinality gopurs_runtime.Value
var once_Data_Enum_defaultCardinality sync.Once
func Get_Data_Enum_defaultCardinality() gopurs_runtime.Value {
	once_Data_Enum_defaultCardinality.Do(func() {
		cache_Data_Enum_defaultCardinality = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_defaultCardinality(dictBounded_0_box)
})
	})
	return cache_Data_Enum_defaultCardinality
}

var cache_Data_Enum_charToEnum gopurs_runtime.Value
var once_Data_Enum_charToEnum sync.Once
func Get_Data_Enum_charToEnum() gopurs_runtime.Value {
	once_Data_Enum_charToEnum.Do(func() {
		cache_Data_Enum_charToEnum = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum(v_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Enum_charToEnum
}

var cache_Data_Enum_enumChar gopurs_runtime.Value
var once_Data_Enum_enumChar sync.Once
func Get_Data_Enum_enumChar() gopurs_runtime.Value {
	once_Data_Enum_enumChar.Do(func() {
		cache_Data_Enum_enumChar = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3409577041_556578094((&Constructor_Data_Enum_Enum[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2406510097_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[string]](Get_Data_Ord_ordChar())))}
}), gopurs_runtime.Apply2(Get_Data_Enum_defaultPred__2686793765(), Get_Data_Enum_charToEnum(), Get_Data_Enum_toCharCode()), gopurs_runtime.Apply2(Get_Data_Enum_defaultSucc__2686793765(), Get_Data_Enum_charToEnum(), Get_Data_Enum_toCharCode())})))}
	})
	return cache_Data_Enum_enumChar
}

var cache_Data_Enum_cardinality gopurs_runtime.Value
var once_Data_Enum_cardinality sync.Once
func Get_Data_Enum_cardinality() gopurs_runtime.Value {
	once_Data_Enum_cardinality.Do(func() {
		cache_Data_Enum_cardinality = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_cardinality(dict_0_box)
})
	})
	return cache_Data_Enum_cardinality
}

var cache_Data_Enum_boundedEnumUnit gopurs_runtime.Value
var once_Data_Enum_boundedEnumUnit sync.Once
func Get_Data_Enum_boundedEnumUnit() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumUnit.Do(func() {
		cache_Data_Enum_boundedEnumUnit = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedUnit()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Enum_enumUnit()))}
}), gopurs_runtime.Int(int64(1)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(int64(0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_0.IntVal) == (int64(0)) {
__t0 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, Get_Data_Unit_unit()})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})}))}
	})
	return cache_Data_Enum_boundedEnumUnit
}

var cache_Data_Enum_boundedEnumOrdering gopurs_runtime.Value
var once_Data_Enum_boundedEnumOrdering sync.Once
func Get_Data_Enum_boundedEnumOrdering() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumOrdering.Do(func() {
		cache_Data_Enum_boundedEnumOrdering = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4021906832_123048125((&Constructor_Data_Enum_BoundedEnum[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_832288803_2094947566(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[uint32]](Get_Data_Bounded_boundedOrdering())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[uint32]](Get_Data_Enum_enumOrdering())))}
}), gopurs_runtime.Int(int64(3)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 int64
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_0) == 1527465420) {
__t3 = int64(0)
goto end_branch_3
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_1) == 902936544) {
__t3 = int64(1)
goto end_branch_3
} else {

}
}
{
var __t_tag_2 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_2) == 380165415) {
__t3 = int64(2)
goto end_branch_3
} else {

}
}
{
__t3 = func() int64 { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Int(__t3)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_0.IntVal) == (int64(0)) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}}))}
goto end_branch_4
} else {

}
}
{
if (v_0.IntVal) == (int64(1)) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}}))}
goto end_branch_4
} else {

}
}
{
if (v_0.IntVal) == (int64(2)) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}}))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t4))))}
})})))}
	})
	return cache_Data_Enum_boundedEnumOrdering
}

var cache_Data_Enum_boundedEnumChar gopurs_runtime.Value
var once_Data_Enum_boundedEnumChar sync.Once
func Get_Data_Enum_boundedEnumChar() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumChar.Do(func() {
		cache_Data_Enum_boundedEnumChar = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3569500834_123048125((&Constructor_Data_Enum_BoundedEnum[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[string]](Get_Data_Bounded_boundedChar())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3409577041_556578094(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[string]](Get_Data_Enum_enumChar())))}
}), gopurs_runtime.Int((gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())).IntVal) - (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal)).IntVal, Get_Data_Enum_toCharCode(), Get_Data_Enum_charToEnum()})))}
	})
	return cache_Data_Enum_boundedEnumChar
}

var cache_Data_Enum_boundedEnumBoolean gopurs_runtime.Value
var once_Data_Enum_boundedEnumBoolean sync.Once
func Get_Data_Enum_boundedEnumBoolean() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumBoolean.Do(func() {
		cache_Data_Enum_boundedEnumBoolean = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3083804185_123048125((&Constructor_Data_Enum_BoundedEnum[bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3728870730_2094947566(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[bool]](Get_Data_Bounded_boundedBoolean())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_342165130_556578094(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[bool]](Get_Data_Enum_enumBoolean())))}
}), gopurs_runtime.Int(int64(2)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 int64
{
if ((v_0.IntVal) != (0)) != (true) {
__t0 = int64(0)
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) != (0) {
__t0 = int64(1)
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_0.IntVal) == (int64(0)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Bool(false)}))}
goto end_branch_1
} else {

}
}
{
if (v_0.IntVal) == (int64(1)) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Bool(true)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1044698560_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[bool]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1044698560_3094389156(Rebox_Data_Enum_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
})})))}
	})
	return cache_Data_Enum_boundedEnumBoolean
}

type Constructor_Data_Enum_Enum[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4075786298] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_Enum[any])(ptr)
		_ = c
		switch key {
		case "Ord0": return gopurs_runtime.Box(c.V0)
		case "pred": return gopurs_runtime.Box(c.V1)
		case "succ": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Enum_Enum: " + key)
		}
	}
}


type Constructor_Data_Enum_BoundedEnum[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 int64
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[287434377] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_BoundedEnum[any])(ptr)
		_ = c
		switch key {
		case "Bounded0": return gopurs_runtime.Box(c.V0)
		case "Enum1": return gopurs_runtime.Box(c.V1)
		case "cardinality": return gopurs_runtime.Box(c.V2)
		case "fromEnum": return gopurs_runtime.Box(c.V3)
		case "toEnum": return gopurs_runtime.Box(c.V4)
		default: panic("Key not found in dictionary Constructor_Data_Enum_BoundedEnum: " + key)
		}
	}
}


func Call_Data_Enum_fromJust(v_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 != nil) {
__t0 = (v_0).V0
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

func Call_Data_Enum_Enum_dollar_Dict(x_0_loop struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
var x_0 struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
				}())
}

func Call_Data_Enum_Cardinality(x_0_loop int64) gopurs_runtime.Value {
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}

func Call_Data_Enum_Cardinality__11842167(x_0_loop int64) gopurs_runtime.Value {
Cardinality__11842167:
for {
if false { continue Cardinality__11842167 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__3304734055(x_0_loop int64) gopurs_runtime.Value {
Cardinality__3304734055:
for {
if false { continue Cardinality__3304734055 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__4193071403(x_0_loop int64) gopurs_runtime.Value {
Cardinality__4193071403:
for {
if false { continue Cardinality__4193071403 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__485476899(x_0_loop int64) gopurs_runtime.Value {
Cardinality__485476899:
for {
if false { continue Cardinality__485476899 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__2083952025(x_0_loop int64) gopurs_runtime.Value {
Cardinality__2083952025:
for {
if false { continue Cardinality__2083952025 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__115045349(x_0_loop int64) gopurs_runtime.Value {
Cardinality__115045349:
for {
if false { continue Cardinality__115045349 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__2170980526(x_0_loop int64) gopurs_runtime.Value {
Cardinality__2170980526:
for {
if false { continue Cardinality__2170980526 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__2518865467(x_0_loop int64) gopurs_runtime.Value {
Cardinality__2518865467:
for {
if false { continue Cardinality__2518865467 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict(x_0_loop struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
var x_0 struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
				}())
}

func Call_Data_Enum_toEnum(dict_0_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dict_0 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Box(dict_0.V4)
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_toEnum__1948670652(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
toEnum__1948670652:
for {
if false { continue toEnum__1948670652 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_0.IntVal) == (int64(1)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(2)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(3)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(4)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(5)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(6)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(7)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(8)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(9)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(10)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(11)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(12)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
}
}

func Call_Data_Enum_toEnum__3403136940(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
toEnum__3403136940:
for {
if false { continue toEnum__3403136940 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_0.IntVal) == (int64(1)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2900196686), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(2)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(20457557), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(3)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4227105004), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(4)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3818857258), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(5)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2946274527), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(6)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1070786179), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0.IntVal) == (int64(7)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1326716170), UnsafePtr: nil}}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
}
}

func Call_Data_Enum_toEnum__106342766(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
toEnum__106342766:
for {
if false { continue toEnum__106342766 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum(__eta_norm_0_0.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
}

func Call_Data_Enum_toEnum__983269893(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
toEnum__983269893:
for {
if false { continue toEnum__983269893 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if ((__eta_norm_0_0.IntVal) >= (int64(1))) && ((__eta_norm_0_0.IntVal) <= (int64(31))) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__eta_norm_0_0.IntVal)}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
}
}

func Call_Data_Enum_succ(dict_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dict_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Box(dict_0.V2)
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_succ__4151667621(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
succ__4151667621:
for {
if false { continue succ__4151667621 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t13 int64
{
var __t_tag_1 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_1) == 1908470532) {
__t13 = int64(2)
goto end_branch_13
} else {

}
}
{
var __t_tag_2 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_2) == 2455627378) {
__t13 = int64(3)
goto end_branch_13
} else {

}
}
{
var __t_tag_3 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_3) == 4162469099) {
__t13 = int64(4)
goto end_branch_13
} else {

}
}
{
var __t_tag_4 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_4) == 1692989816) {
__t13 = int64(5)
goto end_branch_13
} else {

}
}
{
var __t_tag_5 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_5) == 330658827) {
__t13 = int64(6)
goto end_branch_13
} else {

}
}
{
var __t_tag_6 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_6) == 4067355978) {
__t13 = int64(7)
goto end_branch_13
} else {

}
}
{
var __t_tag_7 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_7) == 2276710548) {
__t13 = int64(8)
goto end_branch_13
} else {

}
}
{
var __t_tag_8 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_8) == 243771071) {
__t13 = int64(9)
goto end_branch_13
} else {

}
}
{
var __t_tag_9 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_9) == 215731793) {
__t13 = int64(10)
goto end_branch_13
} else {

}
}
{
var __t_tag_10 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_10) == 8639228) {
__t13 = int64(11)
goto end_branch_13
} else {

}
}
{
var __t_tag_11 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_11) == 49471444) {
__t13 = int64(12)
goto end_branch_13
} else {

}
}
{
var __t_tag_12 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_12) == 3889233761) {
__t13 = int64(13)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_13:
// TAST (Let): __local_var_1_0 shape=Branch(LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, def=Other) bindingType=Any
__local_var_1_0 := __t13
_ = __local_var_1_0
var __t14 gopurs_runtime.Value
{
if (__local_var_1_0) == (int64(1)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(2)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(3)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(4)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(5)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(6)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(7)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(8)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(9)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(10)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(11)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(12)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t14))))}
}
}

func Call_Data_Enum_succ__3996651813(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
succ__3996651813:
for {
if false { continue succ__3996651813 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2280409795_3094389156(Rebox_Data_Enum_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date]](Get_Data_Date_enumDate()).V2), __eta_norm_0_0)))))}
}
}

func Call_Data_Enum_succ__983269893(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
succ__983269893:
for {
if false { continue succ__983269893 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
// TAST (Let): __local_var_1_0 shape=Other bindingType=Int
__local_var_1_0 := (__eta_norm_0_0.IntVal) + (int64(1))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if ((__local_var_1_0) >= (int64(1))) && ((__local_var_1_0) <= (int64(31))) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_1_0)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
}
}

func Call_Data_Enum_upFromIncluding(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value], dictUnfoldable1_1_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable1_1 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable1_1.V0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V2), x_2)))}}))}
}))
}

func Call_Data_Enum_pred(dict_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dict_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Box(dict_0.V1)
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_pred__4151667621(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pred__4151667621:
for {
if false { continue pred__4151667621 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t13 int64
{
var __t_tag_1 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_1) == 1908470532) {
__t13 = int64(0)
goto end_branch_13
} else {

}
}
{
var __t_tag_2 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_2) == 2455627378) {
__t13 = int64(1)
goto end_branch_13
} else {

}
}
{
var __t_tag_3 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_3) == 4162469099) {
__t13 = int64(2)
goto end_branch_13
} else {

}
}
{
var __t_tag_4 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_4) == 1692989816) {
__t13 = int64(3)
goto end_branch_13
} else {

}
}
{
var __t_tag_5 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_5) == 330658827) {
__t13 = int64(4)
goto end_branch_13
} else {

}
}
{
var __t_tag_6 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_6) == 4067355978) {
__t13 = int64(5)
goto end_branch_13
} else {

}
}
{
var __t_tag_7 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_7) == 2276710548) {
__t13 = int64(6)
goto end_branch_13
} else {

}
}
{
var __t_tag_8 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_8) == 243771071) {
__t13 = int64(7)
goto end_branch_13
} else {

}
}
{
var __t_tag_9 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_9) == 215731793) {
__t13 = int64(8)
goto end_branch_13
} else {

}
}
{
var __t_tag_10 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_10) == 8639228) {
__t13 = int64(9)
goto end_branch_13
} else {

}
}
{
var __t_tag_11 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_11) == 49471444) {
__t13 = int64(10)
goto end_branch_13
} else {

}
}
{
var __t_tag_12 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_12) == 3889233761) {
__t13 = int64(11)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_13:
// TAST (Let): __local_var_1_0 shape=Branch(LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, def=Other) bindingType=Any
__local_var_1_0 := __t13
_ = __local_var_1_0
var __t14 gopurs_runtime.Value
{
if (__local_var_1_0) == (int64(1)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(2)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(3)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(4)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(5)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(6)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(7)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(8)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(9)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(10)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(11)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
if (__local_var_1_0) == (int64(12)) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}}))}
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t14))))}
}
}

func Call_Data_Enum_pred__983269893(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pred__983269893:
for {
if false { continue pred__983269893 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
// TAST (Let): __local_var_1_0 shape=Other bindingType=Int
__local_var_1_0 := (gopurs_runtime.Int(__eta_norm_0_0.IntVal).IntVal) - (int64(1))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if ((__local_var_1_0) >= (int64(1))) && ((__local_var_1_0) <= (int64(31))) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_1_0)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
}
}

func Call_Data_Enum_fromEnum(dict_0_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Enum_fromEnum__557243613(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
fromEnum__557243613:
for {
if false { continue fromEnum__557243613 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t12 int64
{
var __t_tag_0 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_0) == 1908470532) {
__t12 = int64(1)
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_1) == 2455627378) {
__t12 = int64(2)
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_2) == 4162469099) {
__t12 = int64(3)
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_3) == 1692989816) {
__t12 = int64(4)
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_4) == 330658827) {
__t12 = int64(5)
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_5) == 4067355978) {
__t12 = int64(6)
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_6) == 2276710548) {
__t12 = int64(7)
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_7) == 243771071) {
__t12 = int64(8)
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_8) == 215731793) {
__t12 = int64(9)
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_9) == 8639228) {
__t12 = int64(10)
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_10) == 49471444) {
__t12 = int64(11)
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = uint32(__eta_norm_0_0.IntVal)
if (uint32(__t_tag_11) == 3889233761) {
__t12 = int64(12)
goto end_branch_12
} else {

}
}
{
__t12 = func() int64 { panic("Failed pattern match") }()
}
end_branch_12:
return gopurs_runtime.Int(__t12)
}
}

func Call_Data_Enum_fromEnum__4291755151(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
fromEnum__4291755151:
for {
if false { continue fromEnum__4291755151 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), __eta_norm_0_0).IntVal)
}
}

func Call_Data_Enum_fromEnum__4163941923(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
fromEnum__4163941923:
for {
if false { continue fromEnum__4163941923 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()).StrVal())).IntVal)
}
}

func Call_Data_Enum_fromEnum__3883318852(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
fromEnum__3883318852:
for {
if false { continue fromEnum__3883318852 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Int(__eta_norm_0_0.IntVal)
}
}

func Call_Data_Enum_toEnumWithDefaults(dictBoundedEnum_0_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBoundedEnum_0 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
// TAST (Let): bottom2_1_0 shape=Other bindingType=(TypeVar a)
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_0.V0), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func3(func(low_2 gopurs_runtime.Value, high_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_0.V4), gopurs_runtime.Int(x_4.IntVal)))
_ = v_5_1
var __t3 gopurs_runtime.Value
{
if (v_5_1 != nil) {
__t3 = (v_5_1).V0
goto end_branch_3
} else {

}
}
{
if (v_5_1 == nil) {
var __t2 gopurs_runtime.Value
{
if (x_4.IntVal) < (gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_0.V3), bottom2_1_0).IntVal) {
__t2 = low_2
goto end_branch_2
} else {

}
}
{
__t2 = high_3
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
}

func Call_Data_Enum_toEnumWithDefaults__4181113775(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
toEnumWithDefaults__4181113775:
for {
if false { continue toEnumWithDefaults__4181113775 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
// TAST (Let): v_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Char])
v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum(x_2.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = v_3_0
var __t2 string
{
if (v_3_0 != nil) {
__t2 = (v_3_0).V0
goto end_branch_2
} else {

}
}
{
if (v_3_0 == nil) {
var __t1 string
{
if (x_2.IntVal) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()).StrVal())).IntVal) {
__t1 = __eta_norm_1_0.StrVal()
goto end_branch_1
} else {

}
}
{
__t1 = __eta_norm_0_1.StrVal()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() string { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Str(__t2)
}
}

func Call_Data_Enum_enumTuple(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
// TAST (Let): ordTuple__193435443_1_0 shape=App(Var) bindingType=Any
ordTuple__193435443_1_0 := gopurs_runtime.Apply(Get_Data_Tuple_ordTuple(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "Ord0"), gopurs_runtime.Value{}))
_ = ordTuple__193435443_1_0
return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bounded0_3_1 shape=App(Other) bindingType=(ADT ["Data","Bounded","Bounded"] [(TypeVar b)])
Bounded0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_3_1
// TAST (Let): Enum1_4_2 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar b)])
Enum1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_4_2
// TAST (Let): ordTuple1_5_3 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])])
ordTuple1_5_3 := Rebox_Data_Enum_4177771502_1535415139(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordTuple__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{}))))
_ = ordTuple1_5_3
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3481421219_556578094((&Constructor_Data_Enum_Enum[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1535415139_4177771502(ordTuple1_5_3))}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))
_ = __local_var_7_4
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_7_4 != nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_7_4).V0, gopurs_runtime.Box(Bounded0_3_1.V2)}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_6:
// TAST (Let): __local_var_8_5 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_8_5 := __t6
_ = __local_var_8_5
// TAST (Let): __local_var_9_8 shape=App(Var) bindingType=(Func [(TypeVar b)] (ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)]))
__local_var_9_8 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
_ = __local_var_9_8
// TAST (Let): __local_var_9_7 shape=Let(Abs(Other)) bindingType=(Func [(TypeVar b)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])]))
__local_var_9_7 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_9_8, x_10)}))}
})
_ = __local_var_9_7
// TAST (Let): __local_var_10_9 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_10_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_4_2.V1), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
_ = __local_var_10_9
var __t10 gopurs_runtime.Value
{
if (__local_var_10_9 == nil) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_8_5)}
goto end_branch_10
} else {

}
}
{
if (__local_var_10_9 != nil) {
__t10 = gopurs_runtime.Apply(__local_var_9_7, (__local_var_10_9).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4010058633_3094389156(Rebox_Data_Enum_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t10))))}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_11 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_7_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))
_ = __local_var_7_11
var __t13 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_7_11 != nil) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_7_11).V0, gopurs_runtime.Box(Bounded0_3_1.V1)}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_13:
// TAST (Let): __local_var_8_12 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_8_12 := __t13
_ = __local_var_8_12
// TAST (Let): __local_var_9_15 shape=App(Var) bindingType=(Func [(TypeVar b)] (ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)]))
__local_var_9_15 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
_ = __local_var_9_15
// TAST (Let): __local_var_9_14 shape=Let(Abs(Other)) bindingType=(Func [(TypeVar b)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])]))
__local_var_9_14 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_9_15, x_10)}))}
})
_ = __local_var_9_14
// TAST (Let): __local_var_10_16 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_10_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_4_2.V2), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
_ = __local_var_10_16
var __t17 gopurs_runtime.Value
{
if (__local_var_10_16 == nil) {
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_8_12)}
goto end_branch_17
} else {

}
}
{
if (__local_var_10_16 != nil) {
__t17 = gopurs_runtime.Apply(__local_var_9_14, (__local_var_10_16).V0)
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4010058633_3094389156(Rebox_Data_Enum_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t17))))}
})})))}
})
}

func Call_Data_Enum_enumMaybe(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
// TAST (Let): Bounded0_1_0 shape=App(Other) bindingType=(ADT ["Data","Bounded","Bounded"] [(TypeVar a)])
Bounded0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_1_0
// TAST (Let): Enum1_2_1 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar a)])
Enum1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_2_1
// TAST (Let): __local_var_3_3 shape=App(Other) bindingType=Any
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=Any
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_4_5
// TAST (Let): eqMaybe1_4_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])
eqMaybe1_4_4 := (&Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 bool
{
var __t_tag_9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_5)
if (__t_tag_9 == nil) {
var __t_tag_10 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_6)
__t11 = (__t_tag_10 == nil)
goto end_branch_11
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_5)
var __t_and_8 bool = false
if (__t_tag_6 != nil) {

var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_6)
__t_and_8 = ((__t_tag_7 != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_5, "eq"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_6.UnsafePtr).V0).IntVal) != (0))
}
__t11 = __t_and_8
}
end_branch_11:
return gopurs_runtime.Bool(__t11)
})})
_ = eqMaybe1_4_4
// TAST (Let): ordMaybe_3_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])
ordMaybe_3_2 := (&Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3508461103_3790796878(eqMaybe1_4_4))}
}), gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 uint32
{
var __t_tag_12 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_5)
if (__t_tag_12 == nil) {
var __t14 uint32
{
var __t_tag_13 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_6)
if (__t_tag_13 == nil) {
__t14 = 902936544
goto end_branch_14
} else {

}
}
{
__t14 = 1527465420
}
end_branch_14:
__t19 = __t14
goto end_branch_19
} else {

}
}
{
var __t_tag_15 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_6)
if (__t_tag_15 == nil) {
__t19 = 380165415
goto end_branch_19
} else {

}
}
{
var __t_tag_16 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_5)
var __t_and_18 bool = false
if (__t_tag_16 != nil) {

var __t_tag_17 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_6)
__t_and_18 = (__t_tag_17 != nil)
}
if __t_and_18 {
__t19 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "compare"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_6.UnsafePtr).V0).IntVal)
goto end_branch_19
} else {

}
}
{
__t19 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t19), UnsafePtr: nil}
})})
_ = ordMaybe_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2522609743_556578094((&Constructor_Data_Enum_Enum[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2155612431_4177771502(ordMaybe_3_2))}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t22 *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]
{
var __t_tag_20 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4)
if (__t_tag_20 == nil) {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_22
} else {

}
}
{
var __t_tag_21 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4)
if (__t_tag_21 != nil) {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_2_1.V1), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_22
} else {

}
}
{
__t22 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] { panic("Failed pattern match") }()
}
end_branch_22:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_215731685_3094389156(__t22))}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 gopurs_runtime.Value
{
var __t_tag_23 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4)
if (__t_tag_23 == nil) {
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_215731685_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Box(Bounded0_1_0.V1)}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_27
} else {

}
}
{
var __t_tag_24 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4)
if (__t_tag_24 != nil) {
// TAST (Let): __local_var_5_25 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_2_1.V2), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0))
_ = __local_var_5_25
var __t26 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_25 != nil) {
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (__local_var_5_25).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_26:
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t26)}
goto end_branch_27
} else {

}
}
{
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_215731685_3094389156(func() *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] { panic("Failed pattern match") }()))}
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_215731685_3094389156(Rebox_Data_Enum_3094389156_215731685(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t27))))}
})})))}
}

func Call_Data_Enum_enumFromTo(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
// TAST (Let): Ord0_1_0 shape=App(Other) bindingType=Any
Ord0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V0), gopurs_runtime.Value{})
_ = Ord0_1_0
// TAST (Let): Eq0_2_1 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeVar a)])
Eq0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_1_0, "Eq0"), gopurs_runtime.Value{}))
_ = Eq0_2_1
// TAST (Let): Ord01_3_2 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(TypeVar a)])
Ord01_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V0), gopurs_runtime.Value{}))
_ = Ord01_3_2
return gopurs_runtime.Func3(func(dictUnfoldable1_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(Eq0_2_1.V0), v_5, v1_6).IntVal) != (0) {
__t13 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_4, "unfoldr1"), gopurs_runtime.Func(func(i_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]
{
if (i_7.IntVal) <= (int64(0)) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{v_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{v_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((i_7.IntVal) - (int64(1)))}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3418986898_138441832(__t7))}
}), gopurs_runtime.Int(int64(0)))
goto end_branch_13
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(Ord01_3_2.V1), v_5, v1_6)
if (uint32(__t_tag_8.IntVal) == 1527465420) {
__t13 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_4, "unfoldr1"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_9 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V2), a_7))
_ = __local_var_8_9
var __t12 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_8_9 != nil) {
var __t11 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), (__local_var_8_9).V0, v1_6)
if ((uint32(__t_tag_10.IntVal) == 380165415)) != (true) {
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_8_9).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_11:
__t12 = __t11
goto end_branch_12
} else {

}
}
{
if (__local_var_8_9 == nil) {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_12
} else {

}
}
{
__t12 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3804580809_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t12)}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
}), v_5)
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_4, "unfoldr1"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_8_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V1), a_7))
_ = __local_var_8_3
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_8_3 != nil) {
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), (__local_var_8_3).V0, v1_6)
if ((uint32(__t_tag_4.IntVal) == 1527465420)) != (true) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_8_3).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (__local_var_8_3 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3804580809_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
}), v_5)
}
end_branch_13:
return __t13
})
}

func Call_Data_Enum_enumFromThenTo(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value], dictFunctor_1_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], dictBoundedEnum_2_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value], a_3_loop gopurs_runtime.Value, b_4_loop gopurs_runtime.Value, c_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var dictFunctor_1 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
var dictBoundedEnum_2 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dictBoundedEnum_2_loop
_ = dictBoundedEnum_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var c_5 gopurs_runtime.Value = c_5_loop
_ = c_5
// TAST (Let): a_prime__6_0 shape=App(Other) bindingType=Int
a_prime__6_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_2.V3), a_3).IntVal
_ = a_prime__6_0
// TAST (Let): __local_var_7_3 shape=Other bindingType=Int
__local_var_7_3 := (gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_2.V3), b_4).IntVal) - (a_prime__6_0)
_ = __local_var_7_3
// TAST (Let): __local_var_8_4 shape=App(Other) bindingType=Int
__local_var_8_4 := gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_2.V3), c_5).IntVal
_ = __local_var_8_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_1.V0), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_1 shape=App(Other) bindingType=(TypeVar c)
__local_var_8_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictBoundedEnum_2.V4), x_7))
_ = __local_var_8_1
var __t2 gopurs_runtime.Value
{
if (__local_var_8_1 != nil) {
__t2 = (__local_var_8_1).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(e_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]
{
if (e_9.IntVal) <= (__local_var_8_4) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(e_9.IntVal), gopurs_runtime.Int((e_9.IntVal) + (__local_var_7_3))}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_5
} else {

}
}
{
__t5 = Rebox_Data_Enum_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1415037225_3094389156(__t5))}
}), gopurs_runtime.Int(a_prime__6_0)))
}

func Call_Data_Enum_enumEither(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
// TAST (Let): Enum1_1_0 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar a)])
Enum1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_1_0
// TAST (Let): Bounded0_2_1 shape=App(Other) bindingType=(ADT ["Data","Bounded","Bounded"] [(TypeVar a)])
Bounded0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_2_1
// TAST (Let): ordEither__193435443_3_2 shape=App(Var) bindingType=Any
ordEither__193435443_3_2 := gopurs_runtime.Apply(Get_Data_Either_ordEither(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{}))
_ = ordEither__193435443_3_2
return gopurs_runtime.Func(func(dictBoundedEnum1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bounded01_5_3 shape=App(Other) bindingType=(ADT ["Data","Bounded","Bounded"] [(TypeVar b)])
Bounded01_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded01_5_3
// TAST (Let): Enum11_6_4 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar b)])
Enum11_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Enum1"), gopurs_runtime.Value{}))
_ = Enum11_6_4
// TAST (Let): ordEither1_7_5 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Either","Either"] [(TypeVar a), (TypeVar b)])])
ordEither1_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordEither__193435443_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{})))
_ = ordEither1_7_5
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Enum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordEither1_7_5)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_8.Type == 9 && v_8.IntVal == 3711209382) {
// TAST (Let): __local_var_9_6 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_9_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_1_0.V1), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
_ = __local_var_9_6
var __t7 gopurs_runtime.Value
{
if (__local_var_9_6 == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_7
} else {

}
}
{
if (__local_var_9_6 != nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(__local_var_9_6).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()}))}
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t7)
goto end_branch_10
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 2465973597) {
// TAST (Let): __local_var_9_8 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(Enum11_6_4.V1), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
_ = __local_var_9_8
var __t9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_9_8 == nil) {
__t9 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Box(Bounded0_2_1.V2)}))}})
goto end_branch_9
} else {

}
}
{
if (__local_var_9_8 != nil) {
__t9 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (__local_var_9_8).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()})
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
__t10 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_8.Type == 9 && v_8.IntVal == 3711209382) {
// TAST (Let): __local_var_9_11 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_9_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_1_0.V2), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
_ = __local_var_9_11
var __t12 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_9_11 == nil) {
__t12 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Box(Bounded01_5_3.V1)}))}})
goto end_branch_12
} else {

}
}
{
if (__local_var_9_11 != nil) {
__t12 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(__local_var_9_11).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()})
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_12:
__t15 = __t12
goto end_branch_15
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 2465973597) {
// TAST (Let): __local_var_9_13 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_9_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(Enum11_6_4.V2), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
_ = __local_var_9_13
var __t14 gopurs_runtime.Value
{
if (__local_var_9_13 == nil) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_14
} else {

}
}
{
if (__local_var_9_13 != nil) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (__local_var_9_13).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()}))}
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t14)
goto end_branch_15
} else {

}
}
{
__t15 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t15)}
})}))}
})
}

func Call_Data_Enum_downFromIncluding(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value], dictUnfoldable1_1_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable1_1 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_1_loop
_ = dictUnfoldable1_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable1_1.V0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V1), x_2)))}}))}
}))
}

func Call_Data_Enum_diag(a_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_0, a_0}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Data_Enum_downFrom(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value], dictUnfoldable_1_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable_1 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_1_loop
_ = dictUnfoldable_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_1.V1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeVar c)
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V1), x_2))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(__local_var_3_0).V0, (__local_var_3_0).V0}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))
}

func Call_Data_Enum_upFrom(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value], dictUnfoldable_1_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
var dictUnfoldable_1 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_1_loop
_ = dictUnfoldable_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_1.V1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeVar c)
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V2), x_2))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(__local_var_3_0).V0, (__local_var_3_0).V0}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))
}

func Call_Data_Enum_defaultToEnum(dictBounded_0_loop *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value], dictEnum_1_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value], i_prime__2_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictBounded_0 *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] = dictBounded_0_loop
_ = dictBounded_0
var dictEnum_1 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_1_loop
_ = dictEnum_1
var i_prime__2 int64 = i_prime__2_loop
_ = i_prime__2
var Call_local_Data_Enum_go__go_3_0_0 func(int64, gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_Enum_go__go_3_0_0
var go__go_3_0_0 gopurs_runtime.Value
_ = go__go_3_0_0
Call_local_Data_Enum_go__go_3_0_0 = func(i_4_loop int64, x_5_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__go_3_0_0:
for {
if false { continue go__go_3_0_0 }
var i_4 int64 = i_4_loop
_ = i_4
var x_5 gopurs_runtime.Value = x_5_loop
_ = x_5
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (i_4) == (int64(0)) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{x_5, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_3
} else {

}
}
{
// TAST (Let): v_6_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
v_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_1.V2), x_5))
_ = v_6_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_6_1 != nil) {
i_4_loop = (i_4) - (int64(1))
x_5_loop = (v_6_1).V0
continue go__go_3_0_0
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v_6_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return __t3
}
}
go__go_3_0_0 = gopurs_runtime.Func(func(i_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_Enum_go__go_3_0_0(i_4_loop_val.IntVal, x_5_loop_val))}
})
})
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (i_prime__2) < (int64(0)) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_4
} else {

}
}
{
__t4 = Call_local_Data_Enum_go__go_3_0_0(i_prime__2, gopurs_runtime.Box(dictBounded_0.V1))
}
end_branch_4:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_defaultSucc(toEnum_prime__0_loop gopurs_runtime.Value, fromEnum_prime__1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var toEnum_prime__0 gopurs_runtime.Value = toEnum_prime__0_loop
_ = toEnum_prime__0
var fromEnum_prime__1 gopurs_runtime.Value = fromEnum_prime__1_loop
_ = fromEnum_prime__1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime__0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime__1, a_2).IntVal) + (int64(1))))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_defaultSucc__2686793765(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
defaultSucc__2686793765:
for {
if false { continue defaultSucc__2686793765 }
var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
_ = toEnum_prime__unused_0
var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
_ = fromEnum_prime__unused_1
var a_2 string = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum((gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(a_2)).IntVal) + (int64(1)))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_defaultSucc__2536258710(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
defaultSucc__2536258710:
for {
if false { continue defaultSucc__2536258710 }
var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
_ = toEnum_prime__unused_0
var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
_ = fromEnum_prime__unused_1
var a_2 int64 = a_2_loop
_ = a_2
// TAST (Let): __local_var_3_0 shape=Other bindingType=Int
__local_var_3_0 := (a_2) + (int64(1))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if ((__local_var_3_0) >= (int64(0))) && ((__local_var_3_0) <= (int64(1114111))) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_3_0)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_defaultPred(toEnum_prime__0_loop gopurs_runtime.Value, fromEnum_prime__1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var toEnum_prime__0 gopurs_runtime.Value = toEnum_prime__0_loop
_ = toEnum_prime__0
var fromEnum_prime__1 gopurs_runtime.Value = fromEnum_prime__1_loop
_ = fromEnum_prime__1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime__0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime__1, a_2).IntVal) - (int64(1))))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_defaultPred__2686793765(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
defaultPred__2686793765:
for {
if false { continue defaultPred__2686793765 }
var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
_ = toEnum_prime__unused_0
var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
_ = fromEnum_prime__unused_1
var a_2 string = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum((gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(a_2)).IntVal) - (int64(1)))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_defaultPred__2536258710(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
defaultPred__2536258710:
for {
if false { continue defaultPred__2536258710 }
var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
_ = toEnum_prime__unused_0
var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
_ = fromEnum_prime__unused_1
var a_2 int64 = a_2_loop
_ = a_2
// TAST (Let): __local_var_3_0 shape=Other bindingType=Int
__local_var_3_0 := (a_2) - (int64(1))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if ((__local_var_3_0) >= (int64(0))) && ((__local_var_3_0) <= (int64(1114111))) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_3_0)}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_defaultFromEnum(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
var Call_local_Data_Enum_go__118764402_1_0_1 func(int64, gopurs_runtime.Value) int64
_ = Call_local_Data_Enum_go__118764402_1_0_1
var go__118764402_1_0_1 gopurs_runtime.Value
_ = go__118764402_1_0_1
Call_local_Data_Enum_go__118764402_1_0_1 = func(i_2_loop int64, x_3_loop gopurs_runtime.Value) int64 {
go__118764402_1_0_1:
for {
if false { continue go__118764402_1_0_1 }
var i_2 int64 = i_2_loop
_ = i_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
// TAST (Let): v_4_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V1), x_3))
_ = v_4_1
var __t2 int64
{
if (v_4_1 != nil) {
i_2_loop = (i_2) + (int64(1))
x_3_loop = (v_4_1).V0
continue go__118764402_1_0_1
__t2 = func() int64 { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v_4_1 == nil) {
__t2 = i_2
goto end_branch_2
} else {

}
}
{
__t2 = func() int64 { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__118764402_1_0_1 = gopurs_runtime.Func(func(i_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Data_Enum_go__118764402_1_0_1(i_2_loop_val.IntVal, x_3_loop_val))
})
})
var go__go_2_3_2 gopurs_runtime.Value
_ = go__go_2_3_2
// FALLBACK TCO: isLoop=false len=1
go__go_2_3_2 = gopurs_runtime.Func2(func(i_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
v_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictEnum_0.V1), x_4))
_ = v_5_4
var __t5 int64
{
if (v_5_4 != nil) {
__t5 = Call_local_Data_Enum_go__118764402_1_0_1((i_3.IntVal) + (int64(1)), (v_5_4).V0)
goto end_branch_5
} else {

}
}
{
if (v_5_4 == nil) {
__t5 = i_3.IntVal
goto end_branch_5
} else {

}
}
{
__t5 = func() int64 { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Int(__t5)
})
return gopurs_runtime.Apply(go__118764402_1_0_1, gopurs_runtime.Int(int64(0)))
}

func Call_Data_Enum_defaultCardinality(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): bottom2_1_0 shape=Other bindingType=(TypeVar a)
bottom2_1_0 := gopurs_runtime.RecordGet(dictBounded_0, "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func(func(dictEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Enum_go__118764402_3_1_3 func(int64, gopurs_runtime.Value) int64
_ = Call_local_Data_Enum_go__118764402_3_1_3
var go__118764402_3_1_3 gopurs_runtime.Value
_ = go__118764402_3_1_3
Call_local_Data_Enum_go__118764402_3_1_3 = func(i_4_loop int64, x_5_loop gopurs_runtime.Value) int64 {
go__118764402_3_1_3:
for {
if false { continue go__118764402_3_1_3 }
var i_4 int64 = i_4_loop
_ = i_4
var x_5 gopurs_runtime.Value = x_5_loop
_ = x_5
// TAST (Let): v_6_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Any])
v_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_5))
_ = v_6_2
var __t3 int64
{
if (v_6_2 != nil) {
i_4_loop = (i_4) + (int64(1))
x_5_loop = (v_6_2).V0
continue go__118764402_3_1_3
__t3 = func() int64 { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
if (v_6_2 == nil) {
__t3 = i_4
goto end_branch_3
} else {

}
}
{
__t3 = func() int64 { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__118764402_3_1_3 = gopurs_runtime.Func(func(i_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Data_Enum_go__118764402_3_1_3(i_4_loop_val.IntVal, x_5_loop_val))
})
})
var go__go_4_4_4 gopurs_runtime.Value
_ = go__go_4_4_4
// FALLBACK TCO: isLoop=false len=1
go__go_4_4_4 = gopurs_runtime.Func2(func(i_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_7_5 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
v_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_6))
_ = v_7_5
var __t6 int64
{
if (v_7_5 != nil) {
__t6 = Call_local_Data_Enum_go__118764402_3_1_3((i_5.IntVal) + (int64(1)), (v_7_5).V0)
goto end_branch_6
} else {

}
}
{
if (v_7_5 == nil) {
__t6 = i_5.IntVal
goto end_branch_6
} else {

}
}
{
__t6 = func() int64 { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Int(__t6)
})
return gopurs_runtime.Int(Call_local_Data_Enum_go__118764402_3_1_3(int64(1), bottom2_1_0))
})
}

func Call_Data_Enum_charToEnum(v_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if ((gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal) >= (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal)) && ((gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())).IntVal) <= (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())).IntVal)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Enum_fromCharCode(), gopurs_runtime.Int(v_0))}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_cardinality(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.Int(gopurs_runtime.RecordGet(dict_0, "cardinality").IntVal)
}

func Rebox_Data_Enum_1044698560_3094389156(in *Constructor_Data_Maybe_Just[bool]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Bool(in.V0)
	return out
}

func Rebox_Data_Enum_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Enum_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Enum_138441832_3363075976(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[int64, int64]{}
		out.V0 = in.V0.IntVal
		out.V1 = in.V1.IntVal
	return out
}

func Rebox_Data_Enum_1415037225_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3363075976_138441832(in.V0))}
	return out
}

func Rebox_Data_Enum_1535415139_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_2155612431_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_215731685_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Enum_219188042_4177771502(in *Constructor_Data_Ord_Ord[bool]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Enum_2359585123_556578094(in *Constructor_Data_Enum_Enum[uint32]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_2420955921_2094947566(in *Constructor_Data_Bounded_Bounded[string]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Str(in.V1)
		out.V2 = gopurs_runtime.Str(in.V2)
	return out
}

func Rebox_Data_Enum_2522609743_556578094(in *Constructor_Data_Enum_Enum[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_2621187955_385277032(in *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, int64]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Enum_3083804185_123048125(in *Constructor_Data_Enum_BoundedEnum[bool]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Enum_3094389156_1044698560(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[bool] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[bool]{}
		out.V0 = (in.V0.IntVal) != (0)
	return out
}

func Rebox_Data_Enum_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Enum_3094389156_1415037225(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]{}
		out.V0 = Rebox_Data_Enum_138441832_3363075976(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Enum_3094389156_215731685(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_Enum_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Data_Enum_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_Enum_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[uint32]{}
		out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Data_Enum_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[string]{}
		out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_Enum_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_3363075976_138441832(in *Constructor_Data_Tuple_Tuple[int64, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Data_Enum_3409577041_556578094(in *Constructor_Data_Enum_Enum[string]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_3418986898_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(in.V1))}
	return out
}

func Rebox_Data_Enum_342165130_556578094(in *Constructor_Data_Enum_Enum[bool]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_3481421219_556578094(in *Constructor_Data_Enum_Enum[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_3508461103_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Enum_3569500834_123048125(in *Constructor_Data_Enum_BoundedEnum[string]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Enum_3728870730_2094947566(in *Constructor_Data_Bounded_Bounded[bool]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Bool(in.V1)
		out.V2 = gopurs_runtime.Bool(in.V2)
	return out
}

func Rebox_Data_Enum_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_3804580809_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Enum_4010058633_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Enum_4021906832_123048125(in *Constructor_Data_Enum_BoundedEnum[uint32]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Enum_4060049525_556578094(in *Constructor_Data_Enum_Enum[int64]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_4177771502_1535415139(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Data_Enum_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Rebox_Data_Enum_832288803_2094947566(in *Constructor_Data_Bounded_Bounded[uint32]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
	return out
}

func Get_Data_Enum_fromCharCode() gopurs_runtime.Value {
	return _Gopurs_Data_Enum_FromCharCode
}

func Get_Data_Enum_toCharCode() gopurs_runtime.Value {
	return _Gopurs_Data_Enum_ToCharCode
}
