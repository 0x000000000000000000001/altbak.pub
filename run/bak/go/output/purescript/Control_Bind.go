package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Bind_identity gopurs_runtime.Value
var once_Control_Bind_identity sync.Once
func Get_Control_Bind_identity() gopurs_runtime.Value {
	once_Control_Bind_identity.Do(func() {
		cache_Control_Bind_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_identity(x_0_box)
})
	})
	return cache_Control_Bind_identity
}

var cache_Control_Bind_Bind_dollarDict gopurs_runtime.Value
var once_Control_Bind_Bind_dollarDict sync.Once
func Get_Control_Bind_Bind_dollarDict() gopurs_runtime.Value {
	once_Control_Bind_Bind_dollarDict.Do(func() {
		cache_Control_Bind_Bind_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_Bind_dollarDict(x_0_box)
})
	})
	return cache_Control_Bind_Bind_dollarDict
}

var cache_Control_Bind_Discard_dollarDict gopurs_runtime.Value
var once_Control_Bind_Discard_dollarDict sync.Once
func Get_Control_Bind_Discard_dollarDict() gopurs_runtime.Value {
	once_Control_Bind_Discard_dollarDict.Do(func() {
		cache_Control_Bind_Discard_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_Discard_dollarDict(x_0_box)
})
	})
	return cache_Control_Bind_Discard_dollarDict
}

var cache_Control_Bind_discard gopurs_runtime.Value
var once_Control_Bind_discard sync.Once
func Get_Control_Bind_discard() gopurs_runtime.Value {
	once_Control_Bind_discard.Do(func() {
		cache_Control_Bind_discard = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Discard](dict_0_box))
})
	})
	return cache_Control_Bind_discard
}

var cache_Control_Bind_bindProxy gopurs_runtime.Value
var once_Control_Bind_bindProxy sync.Once
func Get_Control_Bind_bindProxy() gopurs_runtime.Value {
	once_Control_Bind_bindProxy.Do(func() {
		cache_Control_Bind_bindProxy = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Apply_applyProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_Control_Bind_bindProxy
}

var cache_Control_Bind_bindFn gopurs_runtime.Value
var once_Control_Bind_bindFn sync.Once
func Get_Control_Bind_bindFn() gopurs_runtime.Value {
	once_Control_Bind_bindFn.Do(func() {
		cache_Control_Bind_bindFn = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Apply_applyFn()
}), gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(m_0, x_2), x_2)
})
})
}))
	})
	return cache_Control_Bind_bindFn
}

var cache_Control_Bind_bindArray gopurs_runtime.Value
var once_Control_Bind_bindArray sync.Once
func Get_Control_Bind_bindArray() gopurs_runtime.Value {
	once_Control_Bind_bindArray.Do(func() {
		cache_Control_Bind_bindArray = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Apply_applyArray()
}), Get_Control_Bind_arrayBind())
	})
	return cache_Control_Bind_bindArray
}

var cache_Control_Bind_bind gopurs_runtime.Value
var once_Control_Bind_bind sync.Once
func Get_Control_Bind_bind() gopurs_runtime.Value {
	once_Control_Bind_bind.Do(func() {
		cache_Control_Bind_bind = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind
}

var cache_Control_Bind_bindFlipped gopurs_runtime.Value
var once_Control_Bind_bindFlipped sync.Once
func Get_Control_Bind_bindFlipped() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped.Do(func() {
		cache_Control_Bind_bindFlipped = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bindFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_Control_Bind_bindFlipped
}

var cache_Control_Bind_composeKleisliFlipped gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped sync.Once
func Get_Control_Bind_composeKleisliFlipped() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped.Do(func() {
		cache_Control_Bind_composeKleisliFlipped = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_composeKleisliFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dictBind_0_box), f_1_box, g_2_box, a_3_box)
})
	})
	return cache_Control_Bind_composeKleisliFlipped
}

var cache_Control_Bind_composeKleisli gopurs_runtime.Value
var once_Control_Bind_composeKleisli sync.Once
func Get_Control_Bind_composeKleisli() gopurs_runtime.Value {
	once_Control_Bind_composeKleisli.Do(func() {
		cache_Control_Bind_composeKleisli = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_composeKleisli(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dictBind_0_box), f_1_box, g_2_box, a_3_box)
})
	})
	return cache_Control_Bind_composeKleisli
}

var cache_Control_Bind_discardProxy gopurs_runtime.Value
var once_Control_Bind_discardProxy sync.Once
func Get_Control_Bind_discardProxy() gopurs_runtime.Value {
	once_Control_Bind_discardProxy.Do(func() {
		cache_Control_Bind_discardProxy = gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))
	})
	return cache_Control_Bind_discardProxy
}

var cache_Control_Bind_discardUnit gopurs_runtime.Value
var once_Control_Bind_discardUnit sync.Once
func Get_Control_Bind_discardUnit() gopurs_runtime.Value {
	once_Control_Bind_discardUnit.Do(func() {
		cache_Control_Bind_discardUnit = gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))
	})
	return cache_Control_Bind_discardUnit
}

var cache_Control_Bind_ifM gopurs_runtime.Value
var once_Control_Bind_ifM sync.Once
func Get_Control_Bind_ifM() gopurs_runtime.Value {
	once_Control_Bind_ifM.Do(func() {
		cache_Control_Bind_ifM = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, cond_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_ifM(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dictBind_0_box), cond_1_box, t_2_box, f_3_box)
})
	})
	return cache_Control_Bind_ifM
}

var cache_Control_Bind_join gopurs_runtime.Value
var once_Control_Bind_join sync.Once
func Get_Control_Bind_join() gopurs_runtime.Value {
	once_Control_Bind_join.Do(func() {
		cache_Control_Bind_join = gopurs_runtime.Func2(func(dictBind_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dictBind_0_box), m_1_box)
})
	})
	return cache_Control_Bind_join
}

var cache_Control_Bind_bind__3818858255 gopurs_runtime.Value
var once_Control_Bind_bind__3818858255 sync.Once
func Get_Control_Bind_bind__3818858255() gopurs_runtime.Value {
	once_Control_Bind_bind__3818858255.Do(func() {
		cache_Control_Bind_bind__3818858255 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3818858255(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3818858255
}

var cache_Control_Bind_bind__1276737359 gopurs_runtime.Value
var once_Control_Bind_bind__1276737359 sync.Once
func Get_Control_Bind_bind__1276737359() gopurs_runtime.Value {
	once_Control_Bind_bind__1276737359.Do(func() {
		cache_Control_Bind_bind__1276737359 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1276737359(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1276737359
}

var cache_Control_Bind_bind__3512795567 gopurs_runtime.Value
var once_Control_Bind_bind__3512795567 sync.Once
func Get_Control_Bind_bind__3512795567() gopurs_runtime.Value {
	once_Control_Bind_bind__3512795567.Do(func() {
		cache_Control_Bind_bind__3512795567 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3512795567(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3512795567
}

var cache_Control_Bind_bind__3781138863 gopurs_runtime.Value
var once_Control_Bind_bind__3781138863 sync.Once
func Get_Control_Bind_bind__3781138863() gopurs_runtime.Value {
	once_Control_Bind_bind__3781138863.Do(func() {
		cache_Control_Bind_bind__3781138863 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3781138863(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3781138863
}

var cache_Control_Bind_bind__1739102767 gopurs_runtime.Value
var once_Control_Bind_bind__1739102767 sync.Once
func Get_Control_Bind_bind__1739102767() gopurs_runtime.Value {
	once_Control_Bind_bind__1739102767.Do(func() {
		cache_Control_Bind_bind__1739102767 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1739102767(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1739102767
}

var cache_Control_Bind_bind__3842572251 gopurs_runtime.Value
var once_Control_Bind_bind__3842572251 sync.Once
func Get_Control_Bind_bind__3842572251() gopurs_runtime.Value {
	once_Control_Bind_bind__3842572251.Do(func() {
		cache_Control_Bind_bind__3842572251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3842572251(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3842572251
}

var cache_Control_Bind_bind__556822235 gopurs_runtime.Value
var once_Control_Bind_bind__556822235 sync.Once
func Get_Control_Bind_bind__556822235() gopurs_runtime.Value {
	once_Control_Bind_bind__556822235.Do(func() {
		cache_Control_Bind_bind__556822235 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__556822235(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__556822235
}

var cache_Control_Bind_bind__247450575 gopurs_runtime.Value
var once_Control_Bind_bind__247450575 sync.Once
func Get_Control_Bind_bind__247450575() gopurs_runtime.Value {
	once_Control_Bind_bind__247450575.Do(func() {
		cache_Control_Bind_bind__247450575 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__247450575(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__247450575
}

var cache_Control_Bind_bind__784996047 gopurs_runtime.Value
var once_Control_Bind_bind__784996047 sync.Once
func Get_Control_Bind_bind__784996047() gopurs_runtime.Value {
	once_Control_Bind_bind__784996047.Do(func() {
		cache_Control_Bind_bind__784996047 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__784996047(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__784996047
}

var cache_Control_Bind_bind__830771439 gopurs_runtime.Value
var once_Control_Bind_bind__830771439 sync.Once
func Get_Control_Bind_bind__830771439() gopurs_runtime.Value {
	once_Control_Bind_bind__830771439.Do(func() {
		cache_Control_Bind_bind__830771439 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__830771439(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__830771439
}

var cache_Control_Bind_bind__1397668751 gopurs_runtime.Value
var once_Control_Bind_bind__1397668751 sync.Once
func Get_Control_Bind_bind__1397668751() gopurs_runtime.Value {
	once_Control_Bind_bind__1397668751.Do(func() {
		cache_Control_Bind_bind__1397668751 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1397668751(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1397668751
}

var cache_Control_Bind_bind__2708955087 gopurs_runtime.Value
var once_Control_Bind_bind__2708955087 sync.Once
func Get_Control_Bind_bind__2708955087() gopurs_runtime.Value {
	once_Control_Bind_bind__2708955087.Do(func() {
		cache_Control_Bind_bind__2708955087 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2708955087(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2708955087
}

var cache_Control_Bind_bind__2879969985 gopurs_runtime.Value
var once_Control_Bind_bind__2879969985 sync.Once
func Get_Control_Bind_bind__2879969985() gopurs_runtime.Value {
	once_Control_Bind_bind__2879969985.Do(func() {
		cache_Control_Bind_bind__2879969985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2879969985(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2879969985
}

var cache_Control_Bind_bind__2668924679 gopurs_runtime.Value
var once_Control_Bind_bind__2668924679 sync.Once
func Get_Control_Bind_bind__2668924679() gopurs_runtime.Value {
	once_Control_Bind_bind__2668924679.Do(func() {
		cache_Control_Bind_bind__2668924679 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2668924679(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2668924679
}

var cache_Control_Bind_bind__1858449959 gopurs_runtime.Value
var once_Control_Bind_bind__1858449959 sync.Once
func Get_Control_Bind_bind__1858449959() gopurs_runtime.Value {
	once_Control_Bind_bind__1858449959.Do(func() {
		cache_Control_Bind_bind__1858449959 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1858449959(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1858449959
}

var cache_Control_Bind_bind__3254602343 gopurs_runtime.Value
var once_Control_Bind_bind__3254602343 sync.Once
func Get_Control_Bind_bind__3254602343() gopurs_runtime.Value {
	once_Control_Bind_bind__3254602343.Do(func() {
		cache_Control_Bind_bind__3254602343 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3254602343(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3254602343
}

var cache_Control_Bind_bind__2517150375 gopurs_runtime.Value
var once_Control_Bind_bind__2517150375 sync.Once
func Get_Control_Bind_bind__2517150375() gopurs_runtime.Value {
	once_Control_Bind_bind__2517150375.Do(func() {
		cache_Control_Bind_bind__2517150375 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2517150375(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2517150375
}

var cache_Control_Bind_bind__10508807 gopurs_runtime.Value
var once_Control_Bind_bind__10508807 sync.Once
func Get_Control_Bind_bind__10508807() gopurs_runtime.Value {
	once_Control_Bind_bind__10508807.Do(func() {
		cache_Control_Bind_bind__10508807 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__10508807(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__10508807
}

var cache_Control_Bind_bind__4011257415 gopurs_runtime.Value
var once_Control_Bind_bind__4011257415 sync.Once
func Get_Control_Bind_bind__4011257415() gopurs_runtime.Value {
	once_Control_Bind_bind__4011257415.Do(func() {
		cache_Control_Bind_bind__4011257415 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__4011257415(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__4011257415
}

var cache_Control_Bind_bind__4003908871 gopurs_runtime.Value
var once_Control_Bind_bind__4003908871 sync.Once
func Get_Control_Bind_bind__4003908871() gopurs_runtime.Value {
	once_Control_Bind_bind__4003908871.Do(func() {
		cache_Control_Bind_bind__4003908871 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__4003908871(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__4003908871
}

var cache_Control_Bind_bind__2809271911 gopurs_runtime.Value
var once_Control_Bind_bind__2809271911 sync.Once
func Get_Control_Bind_bind__2809271911() gopurs_runtime.Value {
	once_Control_Bind_bind__2809271911.Do(func() {
		cache_Control_Bind_bind__2809271911 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2809271911(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2809271911
}

var cache_Control_Bind_bind__816807623 gopurs_runtime.Value
var once_Control_Bind_bind__816807623 sync.Once
func Get_Control_Bind_bind__816807623() gopurs_runtime.Value {
	once_Control_Bind_bind__816807623.Do(func() {
		cache_Control_Bind_bind__816807623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__816807623(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__816807623
}

var cache_Control_Bind_bind__4020426567 gopurs_runtime.Value
var once_Control_Bind_bind__4020426567 sync.Once
func Get_Control_Bind_bind__4020426567() gopurs_runtime.Value {
	once_Control_Bind_bind__4020426567.Do(func() {
		cache_Control_Bind_bind__4020426567 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__4020426567(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__4020426567
}

var cache_Control_Bind_bind__268513415 gopurs_runtime.Value
var once_Control_Bind_bind__268513415 sync.Once
func Get_Control_Bind_bind__268513415() gopurs_runtime.Value {
	once_Control_Bind_bind__268513415.Do(func() {
		cache_Control_Bind_bind__268513415 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__268513415(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__268513415
}

var cache_Control_Bind_bind__3225218311 gopurs_runtime.Value
var once_Control_Bind_bind__3225218311 sync.Once
func Get_Control_Bind_bind__3225218311() gopurs_runtime.Value {
	once_Control_Bind_bind__3225218311.Do(func() {
		cache_Control_Bind_bind__3225218311 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3225218311(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3225218311
}

var cache_Control_Bind_bind__4146772295 gopurs_runtime.Value
var once_Control_Bind_bind__4146772295 sync.Once
func Get_Control_Bind_bind__4146772295() gopurs_runtime.Value {
	once_Control_Bind_bind__4146772295.Do(func() {
		cache_Control_Bind_bind__4146772295 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__4146772295(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__4146772295
}

var cache_Control_Bind_bind__2073074151 gopurs_runtime.Value
var once_Control_Bind_bind__2073074151 sync.Once
func Get_Control_Bind_bind__2073074151() gopurs_runtime.Value {
	once_Control_Bind_bind__2073074151.Do(func() {
		cache_Control_Bind_bind__2073074151 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2073074151(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2073074151
}

var cache_Control_Bind_bind__1429695463 gopurs_runtime.Value
var once_Control_Bind_bind__1429695463 sync.Once
func Get_Control_Bind_bind__1429695463() gopurs_runtime.Value {
	once_Control_Bind_bind__1429695463.Do(func() {
		cache_Control_Bind_bind__1429695463 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1429695463(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1429695463
}

var cache_Control_Bind_bind__3043330631 gopurs_runtime.Value
var once_Control_Bind_bind__3043330631 sync.Once
func Get_Control_Bind_bind__3043330631() gopurs_runtime.Value {
	once_Control_Bind_bind__3043330631.Do(func() {
		cache_Control_Bind_bind__3043330631 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3043330631(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3043330631
}

var cache_Control_Bind_bind__1849902087 gopurs_runtime.Value
var once_Control_Bind_bind__1849902087 sync.Once
func Get_Control_Bind_bind__1849902087() gopurs_runtime.Value {
	once_Control_Bind_bind__1849902087.Do(func() {
		cache_Control_Bind_bind__1849902087 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1849902087(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1849902087
}

var cache_Control_Bind_bind__2340925255 gopurs_runtime.Value
var once_Control_Bind_bind__2340925255 sync.Once
func Get_Control_Bind_bind__2340925255() gopurs_runtime.Value {
	once_Control_Bind_bind__2340925255.Do(func() {
		cache_Control_Bind_bind__2340925255 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2340925255(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2340925255
}

var cache_Control_Bind_bind__2601835655 gopurs_runtime.Value
var once_Control_Bind_bind__2601835655 sync.Once
func Get_Control_Bind_bind__2601835655() gopurs_runtime.Value {
	once_Control_Bind_bind__2601835655.Do(func() {
		cache_Control_Bind_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2601835655(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2601835655
}

var cache_Control_Bind_bind__3227627207 gopurs_runtime.Value
var once_Control_Bind_bind__3227627207 sync.Once
func Get_Control_Bind_bind__3227627207() gopurs_runtime.Value {
	once_Control_Bind_bind__3227627207.Do(func() {
		cache_Control_Bind_bind__3227627207 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3227627207(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3227627207
}

var cache_Control_Bind_bind__1881387271 gopurs_runtime.Value
var once_Control_Bind_bind__1881387271 sync.Once
func Get_Control_Bind_bind__1881387271() gopurs_runtime.Value {
	once_Control_Bind_bind__1881387271.Do(func() {
		cache_Control_Bind_bind__1881387271 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1881387271(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1881387271
}

var cache_Control_Bind_bind__3402581063 gopurs_runtime.Value
var once_Control_Bind_bind__3402581063 sync.Once
func Get_Control_Bind_bind__3402581063() gopurs_runtime.Value {
	once_Control_Bind_bind__3402581063.Do(func() {
		cache_Control_Bind_bind__3402581063 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3402581063(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3402581063
}

var cache_Control_Bind_bind__1801470023 gopurs_runtime.Value
var once_Control_Bind_bind__1801470023 sync.Once
func Get_Control_Bind_bind__1801470023() gopurs_runtime.Value {
	once_Control_Bind_bind__1801470023.Do(func() {
		cache_Control_Bind_bind__1801470023 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1801470023(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1801470023
}

var cache_Control_Bind_bind__760543367 gopurs_runtime.Value
var once_Control_Bind_bind__760543367 sync.Once
func Get_Control_Bind_bind__760543367() gopurs_runtime.Value {
	once_Control_Bind_bind__760543367.Do(func() {
		cache_Control_Bind_bind__760543367 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__760543367(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__760543367
}

var cache_Control_Bind_bind__2892370023 gopurs_runtime.Value
var once_Control_Bind_bind__2892370023 sync.Once
func Get_Control_Bind_bind__2892370023() gopurs_runtime.Value {
	once_Control_Bind_bind__2892370023.Do(func() {
		cache_Control_Bind_bind__2892370023 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2892370023(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2892370023
}

var cache_Control_Bind_bind__771821447 gopurs_runtime.Value
var once_Control_Bind_bind__771821447 sync.Once
func Get_Control_Bind_bind__771821447() gopurs_runtime.Value {
	once_Control_Bind_bind__771821447.Do(func() {
		cache_Control_Bind_bind__771821447 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__771821447(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__771821447
}

var cache_Control_Bind_bind__777862183 gopurs_runtime.Value
var once_Control_Bind_bind__777862183 sync.Once
func Get_Control_Bind_bind__777862183() gopurs_runtime.Value {
	once_Control_Bind_bind__777862183.Do(func() {
		cache_Control_Bind_bind__777862183 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__777862183(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__777862183
}

var cache_Control_Bind_bind__1062342183 gopurs_runtime.Value
var once_Control_Bind_bind__1062342183 sync.Once
func Get_Control_Bind_bind__1062342183() gopurs_runtime.Value {
	once_Control_Bind_bind__1062342183.Do(func() {
		cache_Control_Bind_bind__1062342183 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1062342183(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1062342183
}

var cache_Control_Bind_bind__2674370535 gopurs_runtime.Value
var once_Control_Bind_bind__2674370535 sync.Once
func Get_Control_Bind_bind__2674370535() gopurs_runtime.Value {
	once_Control_Bind_bind__2674370535.Do(func() {
		cache_Control_Bind_bind__2674370535 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2674370535(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2674370535
}

var cache_Control_Bind_bind__2453382087 gopurs_runtime.Value
var once_Control_Bind_bind__2453382087 sync.Once
func Get_Control_Bind_bind__2453382087() gopurs_runtime.Value {
	once_Control_Bind_bind__2453382087.Do(func() {
		cache_Control_Bind_bind__2453382087 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2453382087(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2453382087
}

var cache_Control_Bind_bind__889812231 gopurs_runtime.Value
var once_Control_Bind_bind__889812231 sync.Once
func Get_Control_Bind_bind__889812231() gopurs_runtime.Value {
	once_Control_Bind_bind__889812231.Do(func() {
		cache_Control_Bind_bind__889812231 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__889812231(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__889812231
}

var cache_Control_Bind_bind__3652419335 gopurs_runtime.Value
var once_Control_Bind_bind__3652419335 sync.Once
func Get_Control_Bind_bind__3652419335() gopurs_runtime.Value {
	once_Control_Bind_bind__3652419335.Do(func() {
		cache_Control_Bind_bind__3652419335 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3652419335(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3652419335
}

var cache_Control_Bind_bind__1034882759 gopurs_runtime.Value
var once_Control_Bind_bind__1034882759 sync.Once
func Get_Control_Bind_bind__1034882759() gopurs_runtime.Value {
	once_Control_Bind_bind__1034882759.Do(func() {
		cache_Control_Bind_bind__1034882759 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1034882759(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1034882759
}

var cache_Control_Bind_bind__1227450183 gopurs_runtime.Value
var once_Control_Bind_bind__1227450183 sync.Once
func Get_Control_Bind_bind__1227450183() gopurs_runtime.Value {
	once_Control_Bind_bind__1227450183.Do(func() {
		cache_Control_Bind_bind__1227450183 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1227450183(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1227450183
}

var cache_Control_Bind_bind__2931166727 gopurs_runtime.Value
var once_Control_Bind_bind__2931166727 sync.Once
func Get_Control_Bind_bind__2931166727() gopurs_runtime.Value {
	once_Control_Bind_bind__2931166727.Do(func() {
		cache_Control_Bind_bind__2931166727 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2931166727(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2931166727
}

var cache_Control_Bind_bind__1136545735 gopurs_runtime.Value
var once_Control_Bind_bind__1136545735 sync.Once
func Get_Control_Bind_bind__1136545735() gopurs_runtime.Value {
	once_Control_Bind_bind__1136545735.Do(func() {
		cache_Control_Bind_bind__1136545735 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1136545735(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1136545735
}

var cache_Control_Bind_bind__1217453831 gopurs_runtime.Value
var once_Control_Bind_bind__1217453831 sync.Once
func Get_Control_Bind_bind__1217453831() gopurs_runtime.Value {
	once_Control_Bind_bind__1217453831.Do(func() {
		cache_Control_Bind_bind__1217453831 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1217453831(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1217453831
}

var cache_Control_Bind_bind__3465679815 gopurs_runtime.Value
var once_Control_Bind_bind__3465679815 sync.Once
func Get_Control_Bind_bind__3465679815() gopurs_runtime.Value {
	once_Control_Bind_bind__3465679815.Do(func() {
		cache_Control_Bind_bind__3465679815 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3465679815(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3465679815
}

var cache_Control_Bind_bind__1902451271 gopurs_runtime.Value
var once_Control_Bind_bind__1902451271 sync.Once
func Get_Control_Bind_bind__1902451271() gopurs_runtime.Value {
	once_Control_Bind_bind__1902451271.Do(func() {
		cache_Control_Bind_bind__1902451271 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1902451271(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1902451271
}

var cache_Control_Bind_bind__277824519 gopurs_runtime.Value
var once_Control_Bind_bind__277824519 sync.Once
func Get_Control_Bind_bind__277824519() gopurs_runtime.Value {
	once_Control_Bind_bind__277824519.Do(func() {
		cache_Control_Bind_bind__277824519 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__277824519(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__277824519
}

var cache_Control_Bind_bind__925691911 gopurs_runtime.Value
var once_Control_Bind_bind__925691911 sync.Once
func Get_Control_Bind_bind__925691911() gopurs_runtime.Value {
	once_Control_Bind_bind__925691911.Do(func() {
		cache_Control_Bind_bind__925691911 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__925691911(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__925691911
}

var cache_Control_Bind_bind__226726151 gopurs_runtime.Value
var once_Control_Bind_bind__226726151 sync.Once
func Get_Control_Bind_bind__226726151() gopurs_runtime.Value {
	once_Control_Bind_bind__226726151.Do(func() {
		cache_Control_Bind_bind__226726151 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__226726151(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__226726151
}

var cache_Control_Bind_bind__1836438279 gopurs_runtime.Value
var once_Control_Bind_bind__1836438279 sync.Once
func Get_Control_Bind_bind__1836438279() gopurs_runtime.Value {
	once_Control_Bind_bind__1836438279.Do(func() {
		cache_Control_Bind_bind__1836438279 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1836438279(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1836438279
}

var cache_Control_Bind_bind__3306046983 gopurs_runtime.Value
var once_Control_Bind_bind__3306046983 sync.Once
func Get_Control_Bind_bind__3306046983() gopurs_runtime.Value {
	once_Control_Bind_bind__3306046983.Do(func() {
		cache_Control_Bind_bind__3306046983 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3306046983(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3306046983
}

var cache_Control_Bind_bind__3751002439 gopurs_runtime.Value
var once_Control_Bind_bind__3751002439 sync.Once
func Get_Control_Bind_bind__3751002439() gopurs_runtime.Value {
	once_Control_Bind_bind__3751002439.Do(func() {
		cache_Control_Bind_bind__3751002439 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3751002439(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3751002439
}

var cache_Control_Bind_bind__3506215751 gopurs_runtime.Value
var once_Control_Bind_bind__3506215751 sync.Once
func Get_Control_Bind_bind__3506215751() gopurs_runtime.Value {
	once_Control_Bind_bind__3506215751.Do(func() {
		cache_Control_Bind_bind__3506215751 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3506215751(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3506215751
}

var cache_Control_Bind_bind__523053991 gopurs_runtime.Value
var once_Control_Bind_bind__523053991 sync.Once
func Get_Control_Bind_bind__523053991() gopurs_runtime.Value {
	once_Control_Bind_bind__523053991.Do(func() {
		cache_Control_Bind_bind__523053991 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__523053991(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__523053991
}

var cache_Control_Bind_bind__2057121831 gopurs_runtime.Value
var once_Control_Bind_bind__2057121831 sync.Once
func Get_Control_Bind_bind__2057121831() gopurs_runtime.Value {
	once_Control_Bind_bind__2057121831.Do(func() {
		cache_Control_Bind_bind__2057121831 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2057121831(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2057121831
}

var cache_Control_Bind_bind__1459396103 gopurs_runtime.Value
var once_Control_Bind_bind__1459396103 sync.Once
func Get_Control_Bind_bind__1459396103() gopurs_runtime.Value {
	once_Control_Bind_bind__1459396103.Do(func() {
		cache_Control_Bind_bind__1459396103 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1459396103(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1459396103
}

var cache_Control_Bind_bind__367937319 gopurs_runtime.Value
var once_Control_Bind_bind__367937319 sync.Once
func Get_Control_Bind_bind__367937319() gopurs_runtime.Value {
	once_Control_Bind_bind__367937319.Do(func() {
		cache_Control_Bind_bind__367937319 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__367937319(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__367937319
}

var cache_Control_Bind_bind__3229079719 gopurs_runtime.Value
var once_Control_Bind_bind__3229079719 sync.Once
func Get_Control_Bind_bind__3229079719() gopurs_runtime.Value {
	once_Control_Bind_bind__3229079719.Do(func() {
		cache_Control_Bind_bind__3229079719 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3229079719(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3229079719
}

var cache_Control_Bind_bind__991288455 gopurs_runtime.Value
var once_Control_Bind_bind__991288455 sync.Once
func Get_Control_Bind_bind__991288455() gopurs_runtime.Value {
	once_Control_Bind_bind__991288455.Do(func() {
		cache_Control_Bind_bind__991288455 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__991288455(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__991288455
}

var cache_Control_Bind_bind__3078669415 gopurs_runtime.Value
var once_Control_Bind_bind__3078669415 sync.Once
func Get_Control_Bind_bind__3078669415() gopurs_runtime.Value {
	once_Control_Bind_bind__3078669415.Do(func() {
		cache_Control_Bind_bind__3078669415 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3078669415(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3078669415
}

var cache_Control_Bind_bind__961482919 gopurs_runtime.Value
var once_Control_Bind_bind__961482919 sync.Once
func Get_Control_Bind_bind__961482919() gopurs_runtime.Value {
	once_Control_Bind_bind__961482919.Do(func() {
		cache_Control_Bind_bind__961482919 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__961482919(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__961482919
}

var cache_Control_Bind_bind__369905319 gopurs_runtime.Value
var once_Control_Bind_bind__369905319 sync.Once
func Get_Control_Bind_bind__369905319() gopurs_runtime.Value {
	once_Control_Bind_bind__369905319.Do(func() {
		cache_Control_Bind_bind__369905319 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__369905319(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__369905319
}

var cache_Control_Bind_bind__877565287 gopurs_runtime.Value
var once_Control_Bind_bind__877565287 sync.Once
func Get_Control_Bind_bind__877565287() gopurs_runtime.Value {
	once_Control_Bind_bind__877565287.Do(func() {
		cache_Control_Bind_bind__877565287 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__877565287(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__877565287
}

var cache_Control_Bind_bind__1367391623 gopurs_runtime.Value
var once_Control_Bind_bind__1367391623 sync.Once
func Get_Control_Bind_bind__1367391623() gopurs_runtime.Value {
	once_Control_Bind_bind__1367391623.Do(func() {
		cache_Control_Bind_bind__1367391623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1367391623(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1367391623
}

var cache_Control_Bind_bind__3020202567 gopurs_runtime.Value
var once_Control_Bind_bind__3020202567 sync.Once
func Get_Control_Bind_bind__3020202567() gopurs_runtime.Value {
	once_Control_Bind_bind__3020202567.Do(func() {
		cache_Control_Bind_bind__3020202567 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3020202567(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__3020202567
}

var cache_Control_Bind_bind__2669704711 gopurs_runtime.Value
var once_Control_Bind_bind__2669704711 sync.Once
func Get_Control_Bind_bind__2669704711() gopurs_runtime.Value {
	once_Control_Bind_bind__2669704711.Do(func() {
		cache_Control_Bind_bind__2669704711 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2669704711(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2669704711
}

var cache_Control_Bind_bind__2140072327 gopurs_runtime.Value
var once_Control_Bind_bind__2140072327 sync.Once
func Get_Control_Bind_bind__2140072327() gopurs_runtime.Value {
	once_Control_Bind_bind__2140072327.Do(func() {
		cache_Control_Bind_bind__2140072327 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2140072327(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2140072327
}

var cache_Control_Bind_bind__882000455 gopurs_runtime.Value
var once_Control_Bind_bind__882000455 sync.Once
func Get_Control_Bind_bind__882000455() gopurs_runtime.Value {
	once_Control_Bind_bind__882000455.Do(func() {
		cache_Control_Bind_bind__882000455 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__882000455(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__882000455
}

var cache_Control_Bind_bind__324752519 gopurs_runtime.Value
var once_Control_Bind_bind__324752519 sync.Once
func Get_Control_Bind_bind__324752519() gopurs_runtime.Value {
	once_Control_Bind_bind__324752519.Do(func() {
		cache_Control_Bind_bind__324752519 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__324752519(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__324752519
}

var cache_Control_Bind_bind__1492383911 gopurs_runtime.Value
var once_Control_Bind_bind__1492383911 sync.Once
func Get_Control_Bind_bind__1492383911() gopurs_runtime.Value {
	once_Control_Bind_bind__1492383911.Do(func() {
		cache_Control_Bind_bind__1492383911 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1492383911(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1492383911
}

var cache_Control_Bind_bind__2789893127 gopurs_runtime.Value
var once_Control_Bind_bind__2789893127 sync.Once
func Get_Control_Bind_bind__2789893127() gopurs_runtime.Value {
	once_Control_Bind_bind__2789893127.Do(func() {
		cache_Control_Bind_bind__2789893127 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2789893127(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2789893127
}

var cache_Control_Bind_bind__1756310855 gopurs_runtime.Value
var once_Control_Bind_bind__1756310855 sync.Once
func Get_Control_Bind_bind__1756310855() gopurs_runtime.Value {
	once_Control_Bind_bind__1756310855.Do(func() {
		cache_Control_Bind_bind__1756310855 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1756310855(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1756310855
}

var cache_Control_Bind_bind__1907557447 gopurs_runtime.Value
var once_Control_Bind_bind__1907557447 sync.Once
func Get_Control_Bind_bind__1907557447() gopurs_runtime.Value {
	once_Control_Bind_bind__1907557447.Do(func() {
		cache_Control_Bind_bind__1907557447 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1907557447(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1907557447
}

var cache_Control_Bind_bind__2510859911 gopurs_runtime.Value
var once_Control_Bind_bind__2510859911 sync.Once
func Get_Control_Bind_bind__2510859911() gopurs_runtime.Value {
	once_Control_Bind_bind__2510859911.Do(func() {
		cache_Control_Bind_bind__2510859911 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2510859911(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2510859911
}

var cache_Control_Bind_bind__1918470055 gopurs_runtime.Value
var once_Control_Bind_bind__1918470055 sync.Once
func Get_Control_Bind_bind__1918470055() gopurs_runtime.Value {
	once_Control_Bind_bind__1918470055.Do(func() {
		cache_Control_Bind_bind__1918470055 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1918470055(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__1918470055
}

var cache_Control_Bind_bind__2763824711 gopurs_runtime.Value
var once_Control_Bind_bind__2763824711 sync.Once
func Get_Control_Bind_bind__2763824711() gopurs_runtime.Value {
	once_Control_Bind_bind__2763824711.Do(func() {
		cache_Control_Bind_bind__2763824711 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2763824711(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2763824711
}

var cache_Control_Bind_bind__737692327 gopurs_runtime.Value
var once_Control_Bind_bind__737692327 sync.Once
func Get_Control_Bind_bind__737692327() gopurs_runtime.Value {
	once_Control_Bind_bind__737692327.Do(func() {
		cache_Control_Bind_bind__737692327 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__737692327(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__737692327
}

var cache_Control_Bind_bind__2370822215 gopurs_runtime.Value
var once_Control_Bind_bind__2370822215 sync.Once
func Get_Control_Bind_bind__2370822215() gopurs_runtime.Value {
	once_Control_Bind_bind__2370822215.Do(func() {
		cache_Control_Bind_bind__2370822215 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2370822215(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dict_0_box))
})
	})
	return cache_Control_Bind_bind__2370822215
}

var cache_Control_Bind_bind__3352508289 gopurs_runtime.Value
var once_Control_Bind_bind__3352508289 sync.Once
func Get_Control_Bind_bind__3352508289() gopurs_runtime.Value {
	once_Control_Bind_bind__3352508289.Do(func() {
		cache_Control_Bind_bind__3352508289 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3352508289(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3352508289
}

var cache_Control_Bind_bind__1671968481 gopurs_runtime.Value
var once_Control_Bind_bind__1671968481 sync.Once
func Get_Control_Bind_bind__1671968481() gopurs_runtime.Value {
	once_Control_Bind_bind__1671968481.Do(func() {
		cache_Control_Bind_bind__1671968481 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1671968481(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__1671968481
}

var cache_Control_Bind_bind__380208001 gopurs_runtime.Value
var once_Control_Bind_bind__380208001 sync.Once
func Get_Control_Bind_bind__380208001() gopurs_runtime.Value {
	once_Control_Bind_bind__380208001.Do(func() {
		cache_Control_Bind_bind__380208001 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__380208001(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__380208001
}

var cache_Control_Bind_bind__3774783233 gopurs_runtime.Value
var once_Control_Bind_bind__3774783233 sync.Once
func Get_Control_Bind_bind__3774783233() gopurs_runtime.Value {
	once_Control_Bind_bind__3774783233.Do(func() {
		cache_Control_Bind_bind__3774783233 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3774783233(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3774783233
}

var cache_Control_Bind_bind__3115293729 gopurs_runtime.Value
var once_Control_Bind_bind__3115293729 sync.Once
func Get_Control_Bind_bind__3115293729() gopurs_runtime.Value {
	once_Control_Bind_bind__3115293729.Do(func() {
		cache_Control_Bind_bind__3115293729 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3115293729(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3115293729
}

var cache_Control_Bind_bind__3502609729 gopurs_runtime.Value
var once_Control_Bind_bind__3502609729 sync.Once
func Get_Control_Bind_bind__3502609729() gopurs_runtime.Value {
	once_Control_Bind_bind__3502609729.Do(func() {
		cache_Control_Bind_bind__3502609729 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3502609729(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3502609729
}

var cache_Control_Bind_bind__568923361 gopurs_runtime.Value
var once_Control_Bind_bind__568923361 sync.Once
func Get_Control_Bind_bind__568923361() gopurs_runtime.Value {
	once_Control_Bind_bind__568923361.Do(func() {
		cache_Control_Bind_bind__568923361 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__568923361(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__568923361
}

var cache_Control_Bind_bind__1887881377 gopurs_runtime.Value
var once_Control_Bind_bind__1887881377 sync.Once
func Get_Control_Bind_bind__1887881377() gopurs_runtime.Value {
	once_Control_Bind_bind__1887881377.Do(func() {
		cache_Control_Bind_bind__1887881377 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1887881377(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__1887881377
}

var cache_Control_Bind_bind__3897039777 gopurs_runtime.Value
var once_Control_Bind_bind__3897039777 sync.Once
func Get_Control_Bind_bind__3897039777() gopurs_runtime.Value {
	once_Control_Bind_bind__3897039777.Do(func() {
		cache_Control_Bind_bind__3897039777 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3897039777(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3897039777
}

var cache_Control_Bind_bind__2981096353 gopurs_runtime.Value
var once_Control_Bind_bind__2981096353 sync.Once
func Get_Control_Bind_bind__2981096353() gopurs_runtime.Value {
	once_Control_Bind_bind__2981096353.Do(func() {
		cache_Control_Bind_bind__2981096353 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2981096353(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__2981096353
}

var cache_Control_Bind_bind__1491025313 gopurs_runtime.Value
var once_Control_Bind_bind__1491025313 sync.Once
func Get_Control_Bind_bind__1491025313() gopurs_runtime.Value {
	once_Control_Bind_bind__1491025313.Do(func() {
		cache_Control_Bind_bind__1491025313 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1491025313(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__1491025313
}

var cache_Control_Bind_bind__555709729 gopurs_runtime.Value
var once_Control_Bind_bind__555709729 sync.Once
func Get_Control_Bind_bind__555709729() gopurs_runtime.Value {
	once_Control_Bind_bind__555709729.Do(func() {
		cache_Control_Bind_bind__555709729 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__555709729(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__555709729
}

var cache_Control_Bind_bind__2501247745 gopurs_runtime.Value
var once_Control_Bind_bind__2501247745 sync.Once
func Get_Control_Bind_bind__2501247745() gopurs_runtime.Value {
	once_Control_Bind_bind__2501247745.Do(func() {
		cache_Control_Bind_bind__2501247745 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2501247745(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__2501247745
}

var cache_Control_Bind_bind__1160188801 gopurs_runtime.Value
var once_Control_Bind_bind__1160188801 sync.Once
func Get_Control_Bind_bind__1160188801() gopurs_runtime.Value {
	once_Control_Bind_bind__1160188801.Do(func() {
		cache_Control_Bind_bind__1160188801 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1160188801(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__1160188801
}

var cache_Control_Bind_bind__3753603617 gopurs_runtime.Value
var once_Control_Bind_bind__3753603617 sync.Once
func Get_Control_Bind_bind__3753603617() gopurs_runtime.Value {
	once_Control_Bind_bind__3753603617.Do(func() {
		cache_Control_Bind_bind__3753603617 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3753603617(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3753603617
}

var cache_Control_Bind_bind__3937147233 gopurs_runtime.Value
var once_Control_Bind_bind__3937147233 sync.Once
func Get_Control_Bind_bind__3937147233() gopurs_runtime.Value {
	once_Control_Bind_bind__3937147233.Do(func() {
		cache_Control_Bind_bind__3937147233 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3937147233(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3937147233
}

var cache_Control_Bind_bind__187043873 gopurs_runtime.Value
var once_Control_Bind_bind__187043873 sync.Once
func Get_Control_Bind_bind__187043873() gopurs_runtime.Value {
	once_Control_Bind_bind__187043873.Do(func() {
		cache_Control_Bind_bind__187043873 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__187043873(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__187043873
}

var cache_Control_Bind_bind__684765761 gopurs_runtime.Value
var once_Control_Bind_bind__684765761 sync.Once
func Get_Control_Bind_bind__684765761() gopurs_runtime.Value {
	once_Control_Bind_bind__684765761.Do(func() {
		cache_Control_Bind_bind__684765761 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__684765761(xs_0_box, f_1_box)
})
	})
	return cache_Control_Bind_bind__684765761
}

var cache_Control_Bind_bind__4082241 gopurs_runtime.Value
var once_Control_Bind_bind__4082241 sync.Once
func Get_Control_Bind_bind__4082241() gopurs_runtime.Value {
	once_Control_Bind_bind__4082241.Do(func() {
		cache_Control_Bind_bind__4082241 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__4082241(xs_0_box, f_1_box)
})
	})
	return cache_Control_Bind_bind__4082241
}

var cache_Control_Bind_bind__2775489217 gopurs_runtime.Value
var once_Control_Bind_bind__2775489217 sync.Once
func Get_Control_Bind_bind__2775489217() gopurs_runtime.Value {
	once_Control_Bind_bind__2775489217.Do(func() {
		cache_Control_Bind_bind__2775489217 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2775489217(xs_0_box, f_1_box)
})
	})
	return cache_Control_Bind_bind__2775489217
}

var cache_Control_Bind_bind__1872090113 gopurs_runtime.Value
var once_Control_Bind_bind__1872090113 sync.Once
func Get_Control_Bind_bind__1872090113() gopurs_runtime.Value {
	once_Control_Bind_bind__1872090113.Do(func() {
		cache_Control_Bind_bind__1872090113 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Control_Bind_bind__1872090113(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box), v1_1_box))}
})
	})
	return cache_Control_Bind_bind__1872090113
}

var cache_Control_Bind_bind__2207507201 gopurs_runtime.Value
var once_Control_Bind_bind__2207507201 sync.Once
func Get_Control_Bind_bind__2207507201() gopurs_runtime.Value {
	once_Control_Bind_bind__2207507201.Do(func() {
		cache_Control_Bind_bind__2207507201 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Control_Bind_bind__2207507201(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box), v1_1_box))}
})
	})
	return cache_Control_Bind_bind__2207507201
}

var cache_Control_Bind_bind__4062037089 gopurs_runtime.Value
var once_Control_Bind_bind__4062037089 sync.Once
func Get_Control_Bind_bind__4062037089() gopurs_runtime.Value {
	once_Control_Bind_bind__4062037089.Do(func() {
		cache_Control_Bind_bind__4062037089 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Bind_bind__4062037089(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), v1_1_box))}
})
	})
	return cache_Control_Bind_bind__4062037089
}

var cache_Control_Bind_bind__3844717601 gopurs_runtime.Value
var once_Control_Bind_bind__3844717601 sync.Once
func Get_Control_Bind_bind__3844717601() gopurs_runtime.Value {
	once_Control_Bind_bind__3844717601.Do(func() {
		cache_Control_Bind_bind__3844717601 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Bind_bind__3844717601(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), v1_1_box))}
})
	})
	return cache_Control_Bind_bind__3844717601
}

var cache_Control_Bind_bind__1906657537 gopurs_runtime.Value
var once_Control_Bind_bind__1906657537 sync.Once
func Get_Control_Bind_bind__1906657537() gopurs_runtime.Value {
	once_Control_Bind_bind__1906657537.Do(func() {
		cache_Control_Bind_bind__1906657537 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Bind_bind__1906657537(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), v1_1_box))}
})
	})
	return cache_Control_Bind_bind__1906657537
}

var cache_Control_Bind_bind__3444930753 gopurs_runtime.Value
var once_Control_Bind_bind__3444930753 sync.Once
func Get_Control_Bind_bind__3444930753() gopurs_runtime.Value {
	once_Control_Bind_bind__3444930753.Do(func() {
		cache_Control_Bind_bind__3444930753 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Bind_bind__3444930753(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), v1_1_box))}
})
	})
	return cache_Control_Bind_bind__3444930753
}

var cache_Control_Bind_bind__3799579873 gopurs_runtime.Value
var once_Control_Bind_bind__3799579873 sync.Once
func Get_Control_Bind_bind__3799579873() gopurs_runtime.Value {
	once_Control_Bind_bind__3799579873.Do(func() {
		cache_Control_Bind_bind__3799579873 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Bind_bind__3799579873(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), v1_1_box))}
})
	})
	return cache_Control_Bind_bind__3799579873
}

var cache_Control_Bind_bind__1702199617 gopurs_runtime.Value
var once_Control_Bind_bind__1702199617 sync.Once
func Get_Control_Bind_bind__1702199617() gopurs_runtime.Value {
	once_Control_Bind_bind__1702199617.Do(func() {
		cache_Control_Bind_bind__1702199617 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Bind_bind__1702199617(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), v1_1_box))}
})
	})
	return cache_Control_Bind_bind__1702199617
}

var cache_Control_Bind_bind__2389430209 gopurs_runtime.Value
var once_Control_Bind_bind__2389430209 sync.Once
func Get_Control_Bind_bind__2389430209() gopurs_runtime.Value {
	once_Control_Bind_bind__2389430209.Do(func() {
		cache_Control_Bind_bind__2389430209 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Control_Bind_bind__2389430209(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box), f_1_box))}
})
	})
	return cache_Control_Bind_bind__2389430209
}

var cache_Control_Bind_bind__490123073 gopurs_runtime.Value
var once_Control_Bind_bind__490123073 sync.Once
func Get_Control_Bind_bind__490123073() gopurs_runtime.Value {
	once_Control_Bind_bind__490123073.Do(func() {
		cache_Control_Bind_bind__490123073 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__490123073(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__490123073
}

var cache_Control_Bind_bind__1182478273 gopurs_runtime.Value
var once_Control_Bind_bind__1182478273 sync.Once
func Get_Control_Bind_bind__1182478273() gopurs_runtime.Value {
	once_Control_Bind_bind__1182478273.Do(func() {
		cache_Control_Bind_bind__1182478273 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1182478273(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__1182478273
}

var cache_Control_Bind_bind__1451555105 gopurs_runtime.Value
var once_Control_Bind_bind__1451555105 sync.Once
func Get_Control_Bind_bind__1451555105() gopurs_runtime.Value {
	once_Control_Bind_bind__1451555105.Do(func() {
		cache_Control_Bind_bind__1451555105 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1451555105(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__1451555105
}

var cache_Control_Bind_bind__3390533889 gopurs_runtime.Value
var once_Control_Bind_bind__3390533889 sync.Once
func Get_Control_Bind_bind__3390533889() gopurs_runtime.Value {
	once_Control_Bind_bind__3390533889.Do(func() {
		cache_Control_Bind_bind__3390533889 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3390533889(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3390533889
}

var cache_Control_Bind_bind__882999777 gopurs_runtime.Value
var once_Control_Bind_bind__882999777 sync.Once
func Get_Control_Bind_bind__882999777() gopurs_runtime.Value {
	once_Control_Bind_bind__882999777.Do(func() {
		cache_Control_Bind_bind__882999777 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__882999777(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__882999777
}

var cache_Control_Bind_bind__3831761345 gopurs_runtime.Value
var once_Control_Bind_bind__3831761345 sync.Once
func Get_Control_Bind_bind__3831761345() gopurs_runtime.Value {
	once_Control_Bind_bind__3831761345.Do(func() {
		cache_Control_Bind_bind__3831761345 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3831761345(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3831761345
}

var cache_Control_Bind_bind__3550378017 gopurs_runtime.Value
var once_Control_Bind_bind__3550378017 sync.Once
func Get_Control_Bind_bind__3550378017() gopurs_runtime.Value {
	once_Control_Bind_bind__3550378017.Do(func() {
		cache_Control_Bind_bind__3550378017 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3550378017(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3550378017
}

var cache_Control_Bind_bind__1949526049 gopurs_runtime.Value
var once_Control_Bind_bind__1949526049 sync.Once
func Get_Control_Bind_bind__1949526049() gopurs_runtime.Value {
	once_Control_Bind_bind__1949526049.Do(func() {
		cache_Control_Bind_bind__1949526049 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1949526049(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__1949526049
}

var cache_Control_Bind_bind__3328406721 gopurs_runtime.Value
var once_Control_Bind_bind__3328406721 sync.Once
func Get_Control_Bind_bind__3328406721() gopurs_runtime.Value {
	once_Control_Bind_bind__3328406721.Do(func() {
		cache_Control_Bind_bind__3328406721 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3328406721(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3328406721
}

var cache_Control_Bind_bind__2951621345 gopurs_runtime.Value
var once_Control_Bind_bind__2951621345 sync.Once
func Get_Control_Bind_bind__2951621345() gopurs_runtime.Value {
	once_Control_Bind_bind__2951621345.Do(func() {
		cache_Control_Bind_bind__2951621345 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__2951621345(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__2951621345
}

var cache_Control_Bind_bind__3103164513 gopurs_runtime.Value
var once_Control_Bind_bind__3103164513 sync.Once
func Get_Control_Bind_bind__3103164513() gopurs_runtime.Value {
	once_Control_Bind_bind__3103164513.Do(func() {
		cache_Control_Bind_bind__3103164513 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3103164513(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3103164513
}

var cache_Control_Bind_bind__1598467489 gopurs_runtime.Value
var once_Control_Bind_bind__1598467489 sync.Once
func Get_Control_Bind_bind__1598467489() gopurs_runtime.Value {
	once_Control_Bind_bind__1598467489.Do(func() {
		cache_Control_Bind_bind__1598467489 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1598467489(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__1598467489
}

var cache_Control_Bind_bind__899768353 gopurs_runtime.Value
var once_Control_Bind_bind__899768353 sync.Once
func Get_Control_Bind_bind__899768353() gopurs_runtime.Value {
	once_Control_Bind_bind__899768353.Do(func() {
		cache_Control_Bind_bind__899768353 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__899768353(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__899768353
}

var cache_Control_Bind_bind__3119797153 gopurs_runtime.Value
var once_Control_Bind_bind__3119797153 sync.Once
func Get_Control_Bind_bind__3119797153() gopurs_runtime.Value {
	once_Control_Bind_bind__3119797153.Do(func() {
		cache_Control_Bind_bind__3119797153 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3119797153(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3119797153
}

var cache_Control_Bind_bind__1281531809 gopurs_runtime.Value
var once_Control_Bind_bind__1281531809 sync.Once
func Get_Control_Bind_bind__1281531809() gopurs_runtime.Value {
	once_Control_Bind_bind__1281531809.Do(func() {
		cache_Control_Bind_bind__1281531809 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1281531809(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__1281531809
}

var cache_Control_Bind_bind__1325495585 gopurs_runtime.Value
var once_Control_Bind_bind__1325495585 sync.Once
func Get_Control_Bind_bind__1325495585() gopurs_runtime.Value {
	once_Control_Bind_bind__1325495585.Do(func() {
		cache_Control_Bind_bind__1325495585 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1325495585(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__1325495585
}

var cache_Control_Bind_bind__1922668001 gopurs_runtime.Value
var once_Control_Bind_bind__1922668001 sync.Once
func Get_Control_Bind_bind__1922668001() gopurs_runtime.Value {
	once_Control_Bind_bind__1922668001.Do(func() {
		cache_Control_Bind_bind__1922668001 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__1922668001(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__1922668001
}

var cache_Control_Bind_bind__4047544097 gopurs_runtime.Value
var once_Control_Bind_bind__4047544097 sync.Once
func Get_Control_Bind_bind__4047544097() gopurs_runtime.Value {
	once_Control_Bind_bind__4047544097.Do(func() {
		cache_Control_Bind_bind__4047544097 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__4047544097(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__4047544097
}

var cache_Control_Bind_bind__3674668417 gopurs_runtime.Value
var once_Control_Bind_bind__3674668417 sync.Once
func Get_Control_Bind_bind__3674668417() gopurs_runtime.Value {
	once_Control_Bind_bind__3674668417.Do(func() {
		cache_Control_Bind_bind__3674668417 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bind__3674668417(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bind__3674668417
}

var cache_Control_Bind_bindArray__1650562023 gopurs_runtime.Value
var once_Control_Bind_bindArray__1650562023 sync.Once
func Get_Control_Bind_bindArray__1650562023() gopurs_runtime.Value {
	once_Control_Bind_bindArray__1650562023.Do(func() {
		cache_Control_Bind_bindArray__1650562023 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Apply_applyArray()
}), Get_Control_Bind_arrayBind())
	})
	return cache_Control_Bind_bindArray__1650562023
}

var cache_Control_Bind_bindFlipped__804471375 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__804471375 sync.Once
func Get_Control_Bind_bindFlipped__804471375() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__804471375.Do(func() {
		cache_Control_Bind_bindFlipped__804471375 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Bind_bindFlipped__804471375(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dictBind_0_box), b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_2_box)))}
})
	})
	return cache_Control_Bind_bindFlipped__804471375
}

var cache_Control_Bind_bindFlipped__1485397639 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__1485397639 sync.Once
func Get_Control_Bind_bindFlipped__1485397639() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__1485397639.Do(func() {
		cache_Control_Bind_bindFlipped__1485397639 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bindFlipped__1485397639(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_Control_Bind_bindFlipped__1485397639
}

var cache_Control_Bind_bindFlipped__331878215 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__331878215 sync.Once
func Get_Control_Bind_bindFlipped__331878215() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__331878215.Do(func() {
		cache_Control_Bind_bindFlipped__331878215 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bindFlipped__331878215(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_Control_Bind_bindFlipped__331878215
}

var cache_Control_Bind_bindFlipped__3235594689 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__3235594689 sync.Once
func Get_Control_Bind_bindFlipped__3235594689() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__3235594689.Do(func() {
		cache_Control_Bind_bindFlipped__3235594689 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bindFlipped__3235594689(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bindFlipped__3235594689
}

var cache_Control_Bind_bindFlipped__3917280577 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__3917280577 sync.Once
func Get_Control_Bind_bindFlipped__3917280577() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__3917280577.Do(func() {
		cache_Control_Bind_bindFlipped__3917280577 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bindFlipped__3917280577(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bindFlipped__3917280577
}

var cache_Control_Bind_bindFlipped__1454086721 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__1454086721 sync.Once
func Get_Control_Bind_bindFlipped__1454086721() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__1454086721.Do(func() {
		cache_Control_Bind_bindFlipped__1454086721 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bindFlipped__1454086721(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bindFlipped__1454086721
}

var cache_Control_Bind_bindFlipped__3572200705 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__3572200705 sync.Once
func Get_Control_Bind_bindFlipped__3572200705() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__3572200705.Do(func() {
		cache_Control_Bind_bindFlipped__3572200705 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bindFlipped__3572200705(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bindFlipped__3572200705
}

var cache_Control_Bind_bindFlipped__1432323457 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__1432323457 sync.Once
func Get_Control_Bind_bindFlipped__1432323457() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__1432323457.Do(func() {
		cache_Control_Bind_bindFlipped__1432323457 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bindFlipped__1432323457(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bindFlipped__1432323457
}

var cache_Control_Bind_bindFlipped__1317599105 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__1317599105 sync.Once
func Get_Control_Bind_bindFlipped__1317599105() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__1317599105.Do(func() {
		cache_Control_Bind_bindFlipped__1317599105 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_bindFlipped__1317599105(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_bindFlipped__1317599105
}

var cache_Control_Bind_bindFn__1648334822 gopurs_runtime.Value
var once_Control_Bind_bindFn__1648334822 sync.Once
func Get_Control_Bind_bindFn__1648334822() gopurs_runtime.Value {
	once_Control_Bind_bindFn__1648334822.Do(func() {
		cache_Control_Bind_bindFn__1648334822 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Apply_applyFn()
}), gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(m_0, x_2), x_2)
})
})
}))
	})
	return cache_Control_Bind_bindFn__1648334822
}

var cache_Control_Bind_bindProxy__347077479 gopurs_runtime.Value
var once_Control_Bind_bindProxy__347077479 sync.Once
func Get_Control_Bind_bindProxy__347077479() gopurs_runtime.Value {
	once_Control_Bind_bindProxy__347077479.Do(func() {
		cache_Control_Bind_bindProxy__347077479 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Apply_applyProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_Control_Bind_bindProxy__347077479
}

var cache_Control_Bind_composeKleisliFlipped__3637617434 gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped__3637617434 sync.Once
func Get_Control_Bind_composeKleisliFlipped__3637617434() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped__3637617434.Do(func() {
		cache_Control_Bind_composeKleisliFlipped__3637617434 = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_composeKleisliFlipped__3637617434(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dictBind_0_box), f_1_box, g_2_box, a_3_box)
})
	})
	return cache_Control_Bind_composeKleisliFlipped__3637617434
}

var cache_Control_Bind_composeKleisliFlipped__2781497852 gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped__2781497852 sync.Once
func Get_Control_Bind_composeKleisliFlipped__2781497852() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped__2781497852.Do(func() {
		cache_Control_Bind_composeKleisliFlipped__2781497852 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Bind_composeKleisliFlipped__2781497852(f_0_box, g_1_box, a_2_box.StrVal()))}
})
	})
	return cache_Control_Bind_composeKleisliFlipped__2781497852
}

var cache_Control_Bind_discard__439597126 gopurs_runtime.Value
var once_Control_Bind_discard__439597126 sync.Once
func Get_Control_Bind_discard__439597126() gopurs_runtime.Value {
	once_Control_Bind_discard__439597126.Do(func() {
		cache_Control_Bind_discard__439597126 = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__439597126(dictBind_0_box)
})
	})
	return cache_Control_Bind_discard__439597126
}

var cache_Control_Bind_discard__2561459590 gopurs_runtime.Value
var once_Control_Bind_discard__2561459590 sync.Once
func Get_Control_Bind_discard__2561459590() gopurs_runtime.Value {
	once_Control_Bind_discard__2561459590.Do(func() {
		cache_Control_Bind_discard__2561459590 = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2561459590(dictBind_0_box)
})
	})
	return cache_Control_Bind_discard__2561459590
}

var cache_Control_Bind_discard__1876171936 gopurs_runtime.Value
var once_Control_Bind_discard__1876171936 sync.Once
func Get_Control_Bind_discard__1876171936() gopurs_runtime.Value {
	once_Control_Bind_discard__1876171936.Do(func() {
		cache_Control_Bind_discard__1876171936 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1876171936(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_discard__1876171936
}

var cache_Control_Bind_discard__2596713024 gopurs_runtime.Value
var once_Control_Bind_discard__2596713024 sync.Once
func Get_Control_Bind_discard__2596713024() gopurs_runtime.Value {
	once_Control_Bind_discard__2596713024.Do(func() {
		cache_Control_Bind_discard__2596713024 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2596713024(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_discard__2596713024
}

var cache_Control_Bind_discard__3153643456 gopurs_runtime.Value
var once_Control_Bind_discard__3153643456 sync.Once
func Get_Control_Bind_discard__3153643456() gopurs_runtime.Value {
	once_Control_Bind_discard__3153643456.Do(func() {
		cache_Control_Bind_discard__3153643456 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__3153643456(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_discard__3153643456
}

var cache_Control_Bind_discard__203210016 gopurs_runtime.Value
var once_Control_Bind_discard__203210016 sync.Once
func Get_Control_Bind_discard__203210016() gopurs_runtime.Value {
	once_Control_Bind_discard__203210016.Do(func() {
		cache_Control_Bind_discard__203210016 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__203210016(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_discard__203210016
}

var cache_Control_Bind_discard__1979268384 gopurs_runtime.Value
var once_Control_Bind_discard__1979268384 sync.Once
func Get_Control_Bind_discard__1979268384() gopurs_runtime.Value {
	once_Control_Bind_discard__1979268384.Do(func() {
		cache_Control_Bind_discard__1979268384 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1979268384(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_discard__1979268384
}

var cache_Control_Bind_discard__2966453088 gopurs_runtime.Value
var once_Control_Bind_discard__2966453088 sync.Once
func Get_Control_Bind_discard__2966453088() gopurs_runtime.Value {
	once_Control_Bind_discard__2966453088.Do(func() {
		cache_Control_Bind_discard__2966453088 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2966453088(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_discard__2966453088
}

var cache_Control_Bind_discard__2110164512 gopurs_runtime.Value
var once_Control_Bind_discard__2110164512 sync.Once
func Get_Control_Bind_discard__2110164512() gopurs_runtime.Value {
	once_Control_Bind_discard__2110164512.Do(func() {
		cache_Control_Bind_discard__2110164512 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2110164512(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_discard__2110164512
}

var cache_Control_Bind_discard__2399711136 gopurs_runtime.Value
var once_Control_Bind_discard__2399711136 sync.Once
func Get_Control_Bind_discard__2399711136() gopurs_runtime.Value {
	once_Control_Bind_discard__2399711136.Do(func() {
		cache_Control_Bind_discard__2399711136 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2399711136(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_discard__2399711136
}

var cache_Control_Bind_discard__2520179008 gopurs_runtime.Value
var once_Control_Bind_discard__2520179008 sync.Once
func Get_Control_Bind_discard__2520179008() gopurs_runtime.Value {
	once_Control_Bind_discard__2520179008.Do(func() {
		cache_Control_Bind_discard__2520179008 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__2520179008(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_discard__2520179008
}

var cache_Control_Bind_discard__317162198 gopurs_runtime.Value
var once_Control_Bind_discard__317162198 sync.Once
func Get_Control_Bind_discard__317162198() gopurs_runtime.Value {
	once_Control_Bind_discard__317162198.Do(func() {
		cache_Control_Bind_discard__317162198 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__317162198(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Discard](dict_0_box))
})
	})
	return cache_Control_Bind_discard__317162198
}

var cache_Control_Bind_discard__1341268336 gopurs_runtime.Value
var once_Control_Bind_discard__1341268336 sync.Once
func Get_Control_Bind_discard__1341268336() gopurs_runtime.Value {
	once_Control_Bind_discard__1341268336.Do(func() {
		cache_Control_Bind_discard__1341268336 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_discard__1341268336(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Bind_discard__1341268336
}

var cache_Control_Bind_discardUnit__2687062302 gopurs_runtime.Value
var once_Control_Bind_discardUnit__2687062302 sync.Once
func Get_Control_Bind_discardUnit__2687062302() gopurs_runtime.Value {
	once_Control_Bind_discardUnit__2687062302.Do(func() {
		cache_Control_Bind_discardUnit__2687062302 = gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))
	})
	return cache_Control_Bind_discardUnit__2687062302
}

var cache_Control_Bind_join__1635241211 gopurs_runtime.Value
var once_Control_Bind_join__1635241211 sync.Once
func Get_Control_Bind_join__1635241211() gopurs_runtime.Value {
	once_Control_Bind_join__1635241211.Do(func() {
		cache_Control_Bind_join__1635241211 = gopurs_runtime.Func2(func(dictBind_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Bind_join__1635241211(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](dictBind_0_box), m_1_box)
})
	})
	return cache_Control_Bind_join__1635241211
}

var cache_Control_Bind_join__880516349 gopurs_runtime.Value
var once_Control_Bind_join__880516349 sync.Once
func Get_Control_Bind_join__880516349() gopurs_runtime.Value {
	once_Control_Bind_join__880516349.Do(func() {
		cache_Control_Bind_join__880516349 = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Bind_join__880516349(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](m_0_box)))}
})
	})
	return cache_Control_Bind_join__880516349
}

type Constructor_Control_Bind_Bind struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4032919565] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Bind_Bind)(ptr)
		_ = c
		switch key {
		case "Apply0": return gopurs_runtime.Box(c.V0)
		case "bind": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Bind_Bind: " + key)
		}
	}
}


type Constructor_Control_Bind_Discard struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2260728934] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Bind_Discard)(ptr)
		_ = c
		switch key {
		case "discard": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Control_Bind_Discard: " + key)
		}
	}
}


func Call_Control_Bind_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Bind_Bind_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Bind_Discard_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Bind_discard(dict_0_loop *Constructor_Control_Bind_Discard) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Discard = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Bind_bind(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bindFlipped(dictBind_0_loop *Constructor_Control_Bind_Bind, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), a_2, b_1)
}

func Call_Control_Bind_composeKleisliFlipped(dictBind_0_loop *Constructor_Control_Bind_Bind, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind = dictBind_0_loop
_ = dictBind_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), gopurs_runtime.Apply(g_2, a_3), f_1)
}

func Call_Control_Bind_composeKleisli(dictBind_0_loop *Constructor_Control_Bind_Bind, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind = dictBind_0_loop
_ = dictBind_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), gopurs_runtime.Apply(f_1, a_3), g_2)
}

func Call_Control_Bind_ifM(dictBind_0_loop *Constructor_Control_Bind_Bind, cond_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind = dictBind_0_loop
_ = dictBind_0
var cond_1 gopurs_runtime.Value = cond_1_loop
_ = cond_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), cond_1, gopurs_runtime.Func(func(cond_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (cond_prime_4.IntVal) != (0) {
__t0 = t_2
goto end_branch_0
} else {

}
}
{
__t0 = f_3
}
end_branch_0:
return __t0
}))
}

func Call_Control_Bind_join(dictBind_0_loop *Constructor_Control_Bind_Bind, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind = dictBind_0_loop
_ = dictBind_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), m_1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}

func Call_Control_Bind_bind__3818858255(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1276737359(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3512795567(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3781138863(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1739102767(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3842572251(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__556822235(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__247450575(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__784996047(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__830771439(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1397668751(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2708955087(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2879969985(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2668924679(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1858449959(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3254602343(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2517150375(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__10508807(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__4011257415(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__4003908871(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2809271911(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__816807623(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__4020426567(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__268513415(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3225218311(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__4146772295(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2073074151(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1429695463(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3043330631(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1849902087(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2340925255(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2601835655(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3227627207(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1881387271(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3402581063(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1801470023(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__760543367(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2892370023(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__771821447(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__777862183(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1062342183(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2674370535(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2453382087(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__889812231(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3652419335(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1034882759(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1227450183(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2931166727(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1136545735(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1217453831(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3465679815(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1902451271(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__277824519(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__925691911(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__226726151(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1836438279(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3306046983(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3751002439(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3506215751(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__523053991(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2057121831(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1459396103(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__367937319(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3229079719(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__991288455(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3078669415(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__961482919(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__369905319(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__877565287(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1367391623(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3020202567(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2669704711(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2140072327(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__882000455(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__324752519(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1492383911(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2789893127(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1756310855(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1907557447(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2510859911(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__1918470055(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2763824711(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__737692327(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__2370822215(dict_0_loop *Constructor_Control_Bind_Bind) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Bind = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Bind_bind__3352508289(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__1671968481(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__380208001(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__3774783233(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__3115293729(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__3502609729(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__568923361(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__1887881377(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__3897039777(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__2981096353(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__1491025313(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__555709729(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__2501247745(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__1160188801(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__3753603617(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__3937147233(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__187043873(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__684765761(xs_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](Get_Data_List_Lazy_Types_Nil())
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_bindList(), "bind"), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1, f_1))))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t0)}
}), gopurs_runtime.Apply2(Get_Data_Newtype_unwrap(), gopurs_runtime.Value{}, xs_0))
}

func Call_Control_Bind_bind__4082241(xs_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](Get_Data_List_Lazy_Types_Nil())
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_bindList(), "bind"), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1, f_1))))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t0)}
}), gopurs_runtime.Apply2(Get_Data_Newtype_unwrap(), gopurs_runtime.Value{}, xs_0))
}

func Call_Control_Bind_bind__2775489217(xs_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](Get_Data_List_Lazy_Types_Nil())
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0))}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_bindList(), "bind"), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1, f_1))))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t0)}
}), gopurs_runtime.Apply2(Get_Data_Newtype_unwrap(), gopurs_runtime.Value{}, xs_0))
}

func Call_Control_Bind_bind__1872090113(v_0_loop *Constructor_Data_List_Types_Cons, v1_1_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_List_Types_Cons
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](Get_Data_List_Types_Nil())
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_1, (v_0).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}, v1_1)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Control_Bind_bind__2207507201(v_0_loop *Constructor_Data_List_Types_Cons, v1_1_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_List_Types_Cons
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](Get_Data_List_Types_Nil())
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V0))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}, v1_1)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Control_Bind_bind__4062037089(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, gopurs_runtime.Int((v_0).V0.IntVal)))
goto end_branch_0
} else {

}
}
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Control_Bind_bind__3844717601(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, (v_0).V0))
goto end_branch_0
} else {

}
}
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Control_Bind_bind__1906657537(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, (v_0).V0))
goto end_branch_0
} else {

}
}
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Control_Bind_bind__3444930753(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((v_0).V0))}))
goto end_branch_0
} else {

}
}
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Control_Bind_bind__3799579873(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, (v_0).V0))
goto end_branch_0
} else {

}
}
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Control_Bind_bind__1702199617(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, (v_0).V0))
goto end_branch_0
} else {

}
}
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Control_Bind_bind__2389430209(v_0_loop *Constructor_Data_NonEmpty_NonEmpty, f_1_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
// TAST (Let): v1_2_0 -> *Constructor_Data_NonEmpty_NonEmpty
v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V0))}))
_ = v1_2_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_2_0).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_2_0).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_List_Types_toList(), gopurs_runtime.Apply(f_1, x_3))
}))))})))}))
}

func Call_Control_Bind_bind__490123073(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), __eta0_0, __eta1_1)
}

func Call_Control_Bind_bind__1182478273(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), __eta0_0, __eta1_1)
}

func Call_Control_Bind_bind__1451555105(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), __eta0_0, __eta1_1)
}

func Call_Control_Bind_bind__3390533889(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), __eta0_0, __eta1_1)
}

func Call_Control_Bind_bind__882999777(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), __eta0_0, __eta1_1)
}

func Call_Control_Bind_bind__3831761345(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), __eta0_0, __eta1_1)
}

func Call_Control_Bind_bind__3550378017(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__1949526049(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__3328406721(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__2951621345(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__3103164513(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__1598467489(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__899768353(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__3119797153(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__1281531809(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__1325495585(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__1922668001(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__4047544097(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bind__3674668417(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_Control_Bind_bindFlipped__804471375(dictBind_0_loop *Constructor_Control_Bind_Bind, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var dictBind_0 *Constructor_Control_Bind_Bind = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Maybe_Just = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a_2)}, b_1))
}

func Call_Control_Bind_bindFlipped__1485397639(dictBind_0_loop *Constructor_Control_Bind_Bind, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), a_2, b_1)
}

func Call_Control_Bind_bindFlipped__331878215(dictBind_0_loop *Constructor_Control_Bind_Bind, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), a_2, b_1)
}

func Call_Control_Bind_bindFlipped__3235594689(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), __eta1_1, __eta0_0)
}

func Call_Control_Bind_bindFlipped__3917280577(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_bindMaybe(), "bind"), __eta1_1, __eta0_0)
}

func Call_Control_Bind_bindFlipped__1454086721(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_bindMaybe(), "bind"), __eta1_1, __eta0_0)
}

func Call_Control_Bind_bindFlipped__3572200705(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_bindAff(), "bind"), __eta1_1, __eta0_0)
}

func Call_Control_Bind_bindFlipped__1432323457(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_bindAff(), "bind"), __eta1_1, __eta0_0)
}

func Call_Control_Bind_bindFlipped__1317599105(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), __eta1_1, __eta0_0)
}

func Call_Control_Bind_composeKleisliFlipped__3637617434(dictBind_0_loop *Constructor_Control_Bind_Bind, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind = dictBind_0_loop
_ = dictBind_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), gopurs_runtime.Apply(g_2, a_3), f_1)
}

func Call_Control_Bind_composeKleisliFlipped__2781497852(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop string) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var a_2 string = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(g_1, gopurs_runtime.Str(a_2))))}, f_0))
}

func Call_Control_Bind_discard__439597126(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}

func Call_Control_Bind_discard__2561459590(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}

func Call_Control_Bind_discard__1876171936(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), __eta0_0, __eta1_1)
}

func Call_Control_Bind_discard__2596713024(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_bindAff(), "bind"), __eta0_0, __eta1_1)
}

func Call_Control_Bind_discard__3153643456(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_bindAff(), "bind"), __eta0_0, __eta1_1)
}

func Call_Control_Bind_discard__203210016(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), __eta0_0, __eta1_1)
}

func Call_Control_Bind_discard__1979268384(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), __eta0_0, __eta1_1)
}

func Call_Control_Bind_discard__2966453088(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), __eta0_0, __eta1_1)
}

func Call_Control_Bind_discard__2110164512(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), __eta0_0, __eta1_1)
}

func Call_Control_Bind_discard__2399711136(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), __eta0_0, __eta1_1)
}

func Call_Control_Bind_discard__2520179008(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_bindEffect(), "bind"), __eta0_0, __eta1_1)
}

func Call_Control_Bind_discard__317162198(dict_0_loop *Constructor_Control_Bind_Discard) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Bind_Discard = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Bind_discard__1341268336(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), __eta0_0, __eta1_1)
}

func Call_Control_Bind_join__1635241211(dictBind_0_loop *Constructor_Control_Bind_Bind, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *Constructor_Control_Bind_Bind = dictBind_0_loop
_ = dictBind_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBind_0.V1), m_1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}

func Call_Control_Bind_join__880516349(m_0_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var m_0 *Constructor_Data_Maybe_Just = m_0_loop
_ = m_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(m_0)}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})))
}

func Get_Control_Bind_arrayBind() gopurs_runtime.Value {
	return _Gopurs_Control_Bind_ArrayBind
}
