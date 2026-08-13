package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Functor_dollarDict gopurs_runtime.Value
var once_Data_Functor_Functor_dollarDict sync.Once
func Get_Data_Functor_Functor_dollarDict() gopurs_runtime.Value {
	once_Data_Functor_Functor_dollarDict.Do(func() {
		cache_Data_Functor_Functor_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Functor_dollarDict(x_0_box)
})
	})
	return cache_Data_Functor_Functor_dollarDict
}

var cache_Data_Functor_go__map gopurs_runtime.Value
var once_Data_Functor_go__map sync.Once
func Get_Data_Functor_go__map() gopurs_runtime.Value {
	once_Data_Functor_go__map.Do(func() {
		cache_Data_Functor_go__map = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_go__map
}

var cache_Data_Functor_mapFlipped gopurs_runtime.Value
var once_Data_Functor_mapFlipped sync.Once
func Get_Data_Functor_mapFlipped() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped.Do(func() {
		cache_Data_Functor_mapFlipped = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_mapFlipped(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_Data_Functor_mapFlipped
}

var cache_Data_Functor_void gopurs_runtime.Value
var once_Data_Functor_void sync.Once
func Get_Data_Functor_void() gopurs_runtime.Value {
	once_Data_Functor_void.Do(func() {
		cache_Data_Functor_void = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box))
})
	})
	return cache_Data_Functor_void
}

var cache_Data_Functor_voidLeft gopurs_runtime.Value
var once_Data_Functor_voidLeft sync.Once
func Get_Data_Functor_voidLeft() gopurs_runtime.Value {
	once_Data_Functor_voidLeft.Do(func() {
		cache_Data_Functor_voidLeft = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_voidLeft(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, x_2_box)
})
	})
	return cache_Data_Functor_voidLeft
}

var cache_Data_Functor_voidRight gopurs_runtime.Value
var once_Data_Functor_voidRight sync.Once
func Get_Data_Functor_voidRight() gopurs_runtime.Value {
	once_Data_Functor_voidRight.Do(func() {
		cache_Data_Functor_voidRight = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_voidRight(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), x_1_box)
})
	})
	return cache_Data_Functor_voidRight
}

var cache_Data_Functor_functorProxy gopurs_runtime.Value
var once_Data_Functor_functorProxy sync.Once
func Get_Data_Functor_functorProxy() gopurs_runtime.Value {
	once_Data_Functor_functorProxy.Do(func() {
		cache_Data_Functor_functorProxy = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Functor_functorProxy
}

var cache_Data_Functor_functorFn gopurs_runtime.Value
var once_Data_Functor_functorFn sync.Once
func Get_Data_Functor_functorFn() gopurs_runtime.Value {
	once_Data_Functor_functorFn.Do(func() {
		cache_Data_Functor_functorFn = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(Get_Control_Semigroupoid_semigroupoidFn(), "compose"))
	})
	return cache_Data_Functor_functorFn
}

var cache_Data_Functor_functorArray gopurs_runtime.Value
var once_Data_Functor_functorArray sync.Once
func Get_Data_Functor_functorArray() gopurs_runtime.Value {
	once_Data_Functor_functorArray.Do(func() {
		cache_Data_Functor_functorArray = gopurs_runtime.RecordDict1("map", Get_Data_Functor_arrayMap())
	})
	return cache_Data_Functor_functorArray
}

var cache_Data_Functor_flap gopurs_runtime.Value
var once_Data_Functor_flap sync.Once
func Get_Data_Functor_flap() gopurs_runtime.Value {
	once_Data_Functor_flap.Do(func() {
		cache_Data_Functor_flap = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, ff_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_flap(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), ff_1_box, x_2_box)
})
	})
	return cache_Data_Functor_flap
}

var cache_Data_Functor_functorArray__2747750794 gopurs_runtime.Value
var once_Data_Functor_functorArray__2747750794 sync.Once
func Get_Data_Functor_functorArray__2747750794() gopurs_runtime.Value {
	once_Data_Functor_functorArray__2747750794.Do(func() {
		cache_Data_Functor_functorArray__2747750794 = gopurs_runtime.RecordDict1("map", Get_Data_Functor_arrayMap())
	})
	return cache_Data_Functor_functorArray__2747750794
}

var cache_Data_Functor_functorArray__361387505 gopurs_runtime.Value
var once_Data_Functor_functorArray__361387505 sync.Once
func Get_Data_Functor_functorArray__361387505() gopurs_runtime.Value {
	once_Data_Functor_functorArray__361387505.Do(func() {
		cache_Data_Functor_functorArray__361387505 = gopurs_runtime.RecordDict1("map", Get_Data_Functor_arrayMap())
	})
	return cache_Data_Functor_functorArray__361387505
}

var cache_Data_Functor_functorFn__1225168408 gopurs_runtime.Value
var once_Data_Functor_functorFn__1225168408 sync.Once
func Get_Data_Functor_functorFn__1225168408() gopurs_runtime.Value {
	once_Data_Functor_functorFn__1225168408.Do(func() {
		cache_Data_Functor_functorFn__1225168408 = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(Get_Control_Semigroupoid_semigroupoidFn(), "compose"))
	})
	return cache_Data_Functor_functorFn__1225168408
}

var cache_Data_Functor_functorFn__2451697751 gopurs_runtime.Value
var once_Data_Functor_functorFn__2451697751 sync.Once
func Get_Data_Functor_functorFn__2451697751() gopurs_runtime.Value {
	once_Data_Functor_functorFn__2451697751.Do(func() {
		cache_Data_Functor_functorFn__2451697751 = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(Get_Control_Semigroupoid_semigroupoidFn(), "compose"))
	})
	return cache_Data_Functor_functorFn__2451697751
}

var cache_Data_Functor_functorFn__20325936 gopurs_runtime.Value
var once_Data_Functor_functorFn__20325936 sync.Once
func Get_Data_Functor_functorFn__20325936() gopurs_runtime.Value {
	once_Data_Functor_functorFn__20325936.Do(func() {
		cache_Data_Functor_functorFn__20325936 = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(Get_Control_Semigroupoid_semigroupoidFn(), "compose"))
	})
	return cache_Data_Functor_functorFn__20325936
}

var cache_Data_Functor_functorProxy__1157108209 gopurs_runtime.Value
var once_Data_Functor_functorProxy__1157108209 sync.Once
func Get_Data_Functor_functorProxy__1157108209() gopurs_runtime.Value {
	once_Data_Functor_functorProxy__1157108209.Do(func() {
		cache_Data_Functor_functorProxy__1157108209 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Functor_functorProxy__1157108209
}

var cache_Data_Functor_functorProxy__711768561 gopurs_runtime.Value
var once_Data_Functor_functorProxy__711768561 sync.Once
func Get_Data_Functor_functorProxy__711768561() gopurs_runtime.Value {
	once_Data_Functor_functorProxy__711768561.Do(func() {
		cache_Data_Functor_functorProxy__711768561 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Functor_functorProxy__711768561
}

var cache_Data_Functor_map__1165794789 gopurs_runtime.Value
var once_Data_Functor_map__1165794789 sync.Once
func Get_Data_Functor_map__1165794789() gopurs_runtime.Value {
	once_Data_Functor_map__1165794789.Do(func() {
		cache_Data_Functor_map__1165794789 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1165794789(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1165794789
}

var cache_Data_Functor_map__1162721797 gopurs_runtime.Value
var once_Data_Functor_map__1162721797 sync.Once
func Get_Data_Functor_map__1162721797() gopurs_runtime.Value {
	once_Data_Functor_map__1162721797.Do(func() {
		cache_Data_Functor_map__1162721797 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1162721797(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1162721797
}

var cache_Data_Functor_map__1924492325 gopurs_runtime.Value
var once_Data_Functor_map__1924492325 sync.Once
func Get_Data_Functor_map__1924492325() gopurs_runtime.Value {
	once_Data_Functor_map__1924492325.Do(func() {
		cache_Data_Functor_map__1924492325 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1924492325(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1924492325
}

var cache_Data_Functor_map__2805967941 gopurs_runtime.Value
var once_Data_Functor_map__2805967941 sync.Once
func Get_Data_Functor_map__2805967941() gopurs_runtime.Value {
	once_Data_Functor_map__2805967941.Do(func() {
		cache_Data_Functor_map__2805967941 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2805967941(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2805967941
}

var cache_Data_Functor_map__3676941189 gopurs_runtime.Value
var once_Data_Functor_map__3676941189 sync.Once
func Get_Data_Functor_map__3676941189() gopurs_runtime.Value {
	once_Data_Functor_map__3676941189.Do(func() {
		cache_Data_Functor_map__3676941189 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3676941189(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3676941189
}

var cache_Data_Functor_map__2364682565 gopurs_runtime.Value
var once_Data_Functor_map__2364682565 sync.Once
func Get_Data_Functor_map__2364682565() gopurs_runtime.Value {
	once_Data_Functor_map__2364682565.Do(func() {
		cache_Data_Functor_map__2364682565 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2364682565(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2364682565
}

var cache_Data_Functor_map__2321705669 gopurs_runtime.Value
var once_Data_Functor_map__2321705669 sync.Once
func Get_Data_Functor_map__2321705669() gopurs_runtime.Value {
	once_Data_Functor_map__2321705669.Do(func() {
		cache_Data_Functor_map__2321705669 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2321705669(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2321705669
}

var cache_Data_Functor_map__3658399301 gopurs_runtime.Value
var once_Data_Functor_map__3658399301 sync.Once
func Get_Data_Functor_map__3658399301() gopurs_runtime.Value {
	once_Data_Functor_map__3658399301.Do(func() {
		cache_Data_Functor_map__3658399301 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3658399301(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3658399301
}

var cache_Data_Functor_map__4176008549 gopurs_runtime.Value
var once_Data_Functor_map__4176008549 sync.Once
func Get_Data_Functor_map__4176008549() gopurs_runtime.Value {
	once_Data_Functor_map__4176008549.Do(func() {
		cache_Data_Functor_map__4176008549 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4176008549(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__4176008549
}

var cache_Data_Functor_map__4040535013 gopurs_runtime.Value
var once_Data_Functor_map__4040535013 sync.Once
func Get_Data_Functor_map__4040535013() gopurs_runtime.Value {
	once_Data_Functor_map__4040535013.Do(func() {
		cache_Data_Functor_map__4040535013 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4040535013(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__4040535013
}

var cache_Data_Functor_map__3116241637 gopurs_runtime.Value
var once_Data_Functor_map__3116241637 sync.Once
func Get_Data_Functor_map__3116241637() gopurs_runtime.Value {
	once_Data_Functor_map__3116241637.Do(func() {
		cache_Data_Functor_map__3116241637 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3116241637(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3116241637
}

var cache_Data_Functor_map__4282869861 gopurs_runtime.Value
var once_Data_Functor_map__4282869861 sync.Once
func Get_Data_Functor_map__4282869861() gopurs_runtime.Value {
	once_Data_Functor_map__4282869861.Do(func() {
		cache_Data_Functor_map__4282869861 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4282869861(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__4282869861
}

var cache_Data_Functor_map__2539978757 gopurs_runtime.Value
var once_Data_Functor_map__2539978757 sync.Once
func Get_Data_Functor_map__2539978757() gopurs_runtime.Value {
	once_Data_Functor_map__2539978757.Do(func() {
		cache_Data_Functor_map__2539978757 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2539978757(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2539978757
}

var cache_Data_Functor_map__4119868389 gopurs_runtime.Value
var once_Data_Functor_map__4119868389 sync.Once
func Get_Data_Functor_map__4119868389() gopurs_runtime.Value {
	once_Data_Functor_map__4119868389.Do(func() {
		cache_Data_Functor_map__4119868389 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4119868389(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__4119868389
}

var cache_Data_Functor_map__3122478373 gopurs_runtime.Value
var once_Data_Functor_map__3122478373 sync.Once
func Get_Data_Functor_map__3122478373() gopurs_runtime.Value {
	once_Data_Functor_map__3122478373.Do(func() {
		cache_Data_Functor_map__3122478373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3122478373(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3122478373
}

var cache_Data_Functor_map__2665381605 gopurs_runtime.Value
var once_Data_Functor_map__2665381605 sync.Once
func Get_Data_Functor_map__2665381605() gopurs_runtime.Value {
	once_Data_Functor_map__2665381605.Do(func() {
		cache_Data_Functor_map__2665381605 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2665381605(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2665381605
}

var cache_Data_Functor_map__4104571781 gopurs_runtime.Value
var once_Data_Functor_map__4104571781 sync.Once
func Get_Data_Functor_map__4104571781() gopurs_runtime.Value {
	once_Data_Functor_map__4104571781.Do(func() {
		cache_Data_Functor_map__4104571781 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4104571781(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__4104571781
}

var cache_Data_Functor_map__1599448997 gopurs_runtime.Value
var once_Data_Functor_map__1599448997 sync.Once
func Get_Data_Functor_map__1599448997() gopurs_runtime.Value {
	once_Data_Functor_map__1599448997.Do(func() {
		cache_Data_Functor_map__1599448997 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1599448997(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1599448997
}

var cache_Data_Functor_map__4285761829 gopurs_runtime.Value
var once_Data_Functor_map__4285761829 sync.Once
func Get_Data_Functor_map__4285761829() gopurs_runtime.Value {
	once_Data_Functor_map__4285761829.Do(func() {
		cache_Data_Functor_map__4285761829 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4285761829(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__4285761829
}

var cache_Data_Functor_map__3149795237 gopurs_runtime.Value
var once_Data_Functor_map__3149795237 sync.Once
func Get_Data_Functor_map__3149795237() gopurs_runtime.Value {
	once_Data_Functor_map__3149795237.Do(func() {
		cache_Data_Functor_map__3149795237 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3149795237(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3149795237
}

var cache_Data_Functor_map__1184159621 gopurs_runtime.Value
var once_Data_Functor_map__1184159621 sync.Once
func Get_Data_Functor_map__1184159621() gopurs_runtime.Value {
	once_Data_Functor_map__1184159621.Do(func() {
		cache_Data_Functor_map__1184159621 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1184159621(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1184159621
}

var cache_Data_Functor_map__609712645 gopurs_runtime.Value
var once_Data_Functor_map__609712645 sync.Once
func Get_Data_Functor_map__609712645() gopurs_runtime.Value {
	once_Data_Functor_map__609712645.Do(func() {
		cache_Data_Functor_map__609712645 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__609712645(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__609712645
}

var cache_Data_Functor_map__1171202917 gopurs_runtime.Value
var once_Data_Functor_map__1171202917 sync.Once
func Get_Data_Functor_map__1171202917() gopurs_runtime.Value {
	once_Data_Functor_map__1171202917.Do(func() {
		cache_Data_Functor_map__1171202917 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1171202917(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1171202917
}

var cache_Data_Functor_map__677918245 gopurs_runtime.Value
var once_Data_Functor_map__677918245 sync.Once
func Get_Data_Functor_map__677918245() gopurs_runtime.Value {
	once_Data_Functor_map__677918245.Do(func() {
		cache_Data_Functor_map__677918245 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__677918245(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__677918245
}

var cache_Data_Functor_map__1307706501 gopurs_runtime.Value
var once_Data_Functor_map__1307706501 sync.Once
func Get_Data_Functor_map__1307706501() gopurs_runtime.Value {
	once_Data_Functor_map__1307706501.Do(func() {
		cache_Data_Functor_map__1307706501 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1307706501(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1307706501
}

var cache_Data_Functor_map__1408505925 gopurs_runtime.Value
var once_Data_Functor_map__1408505925 sync.Once
func Get_Data_Functor_map__1408505925() gopurs_runtime.Value {
	once_Data_Functor_map__1408505925.Do(func() {
		cache_Data_Functor_map__1408505925 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1408505925(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1408505925
}

var cache_Data_Functor_map__1542634789 gopurs_runtime.Value
var once_Data_Functor_map__1542634789 sync.Once
func Get_Data_Functor_map__1542634789() gopurs_runtime.Value {
	once_Data_Functor_map__1542634789.Do(func() {
		cache_Data_Functor_map__1542634789 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1542634789(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1542634789
}

var cache_Data_Functor_map__2675323109 gopurs_runtime.Value
var once_Data_Functor_map__2675323109 sync.Once
func Get_Data_Functor_map__2675323109() gopurs_runtime.Value {
	once_Data_Functor_map__2675323109.Do(func() {
		cache_Data_Functor_map__2675323109 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2675323109(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2675323109
}

var cache_Data_Functor_map__258070885 gopurs_runtime.Value
var once_Data_Functor_map__258070885 sync.Once
func Get_Data_Functor_map__258070885() gopurs_runtime.Value {
	once_Data_Functor_map__258070885.Do(func() {
		cache_Data_Functor_map__258070885 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__258070885(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__258070885
}

var cache_Data_Functor_map__47904357 gopurs_runtime.Value
var once_Data_Functor_map__47904357 sync.Once
func Get_Data_Functor_map__47904357() gopurs_runtime.Value {
	once_Data_Functor_map__47904357.Do(func() {
		cache_Data_Functor_map__47904357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__47904357(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__47904357
}

var cache_Data_Functor_map__3871729957 gopurs_runtime.Value
var once_Data_Functor_map__3871729957 sync.Once
func Get_Data_Functor_map__3871729957() gopurs_runtime.Value {
	once_Data_Functor_map__3871729957.Do(func() {
		cache_Data_Functor_map__3871729957 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3871729957(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3871729957
}

var cache_Data_Functor_map__67411525 gopurs_runtime.Value
var once_Data_Functor_map__67411525 sync.Once
func Get_Data_Functor_map__67411525() gopurs_runtime.Value {
	once_Data_Functor_map__67411525.Do(func() {
		cache_Data_Functor_map__67411525 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__67411525(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__67411525
}

var cache_Data_Functor_map__321096773 gopurs_runtime.Value
var once_Data_Functor_map__321096773 sync.Once
func Get_Data_Functor_map__321096773() gopurs_runtime.Value {
	once_Data_Functor_map__321096773.Do(func() {
		cache_Data_Functor_map__321096773 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__321096773(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__321096773
}

var cache_Data_Functor_map__2876470885 gopurs_runtime.Value
var once_Data_Functor_map__2876470885 sync.Once
func Get_Data_Functor_map__2876470885() gopurs_runtime.Value {
	once_Data_Functor_map__2876470885.Do(func() {
		cache_Data_Functor_map__2876470885 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2876470885(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2876470885
}

var cache_Data_Functor_map__1434901492 gopurs_runtime.Value
var once_Data_Functor_map__1434901492 sync.Once
func Get_Data_Functor_map__1434901492() gopurs_runtime.Value {
	once_Data_Functor_map__1434901492.Do(func() {
		cache_Data_Functor_map__1434901492 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1434901492(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1434901492
}

var cache_Data_Functor_map__243231988 gopurs_runtime.Value
var once_Data_Functor_map__243231988 sync.Once
func Get_Data_Functor_map__243231988() gopurs_runtime.Value {
	once_Data_Functor_map__243231988.Do(func() {
		cache_Data_Functor_map__243231988 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__243231988(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__243231988
}

var cache_Data_Functor_map__2278567252 gopurs_runtime.Value
var once_Data_Functor_map__2278567252 sync.Once
func Get_Data_Functor_map__2278567252() gopurs_runtime.Value {
	once_Data_Functor_map__2278567252.Do(func() {
		cache_Data_Functor_map__2278567252 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2278567252(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2278567252
}

var cache_Data_Functor_map__491003380 gopurs_runtime.Value
var once_Data_Functor_map__491003380 sync.Once
func Get_Data_Functor_map__491003380() gopurs_runtime.Value {
	once_Data_Functor_map__491003380.Do(func() {
		cache_Data_Functor_map__491003380 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__491003380(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__491003380
}

var cache_Data_Functor_map__358342900 gopurs_runtime.Value
var once_Data_Functor_map__358342900 sync.Once
func Get_Data_Functor_map__358342900() gopurs_runtime.Value {
	once_Data_Functor_map__358342900.Do(func() {
		cache_Data_Functor_map__358342900 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__358342900(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__358342900
}

var cache_Data_Functor_map__2224957140 gopurs_runtime.Value
var once_Data_Functor_map__2224957140 sync.Once
func Get_Data_Functor_map__2224957140() gopurs_runtime.Value {
	once_Data_Functor_map__2224957140.Do(func() {
		cache_Data_Functor_map__2224957140 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2224957140(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2224957140
}

var cache_Data_Functor_map__3483556436 gopurs_runtime.Value
var once_Data_Functor_map__3483556436 sync.Once
func Get_Data_Functor_map__3483556436() gopurs_runtime.Value {
	once_Data_Functor_map__3483556436.Do(func() {
		cache_Data_Functor_map__3483556436 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3483556436(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3483556436
}

var cache_Data_Functor_map__2869957716 gopurs_runtime.Value
var once_Data_Functor_map__2869957716 sync.Once
func Get_Data_Functor_map__2869957716() gopurs_runtime.Value {
	once_Data_Functor_map__2869957716.Do(func() {
		cache_Data_Functor_map__2869957716 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2869957716(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2869957716
}

var cache_Data_Functor_map__2475811444 gopurs_runtime.Value
var once_Data_Functor_map__2475811444 sync.Once
func Get_Data_Functor_map__2475811444() gopurs_runtime.Value {
	once_Data_Functor_map__2475811444.Do(func() {
		cache_Data_Functor_map__2475811444 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2475811444(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2475811444
}

var cache_Data_Functor_map__3218804116 gopurs_runtime.Value
var once_Data_Functor_map__3218804116 sync.Once
func Get_Data_Functor_map__3218804116() gopurs_runtime.Value {
	once_Data_Functor_map__3218804116.Do(func() {
		cache_Data_Functor_map__3218804116 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3218804116(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3218804116
}

var cache_Data_Functor_map__3240628980 gopurs_runtime.Value
var once_Data_Functor_map__3240628980 sync.Once
func Get_Data_Functor_map__3240628980() gopurs_runtime.Value {
	once_Data_Functor_map__3240628980.Do(func() {
		cache_Data_Functor_map__3240628980 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3240628980(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3240628980
}

var cache_Data_Functor_map__2199395572 gopurs_runtime.Value
var once_Data_Functor_map__2199395572 sync.Once
func Get_Data_Functor_map__2199395572() gopurs_runtime.Value {
	once_Data_Functor_map__2199395572.Do(func() {
		cache_Data_Functor_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2199395572(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2199395572
}

var cache_Data_Functor_map__16849908 gopurs_runtime.Value
var once_Data_Functor_map__16849908 sync.Once
func Get_Data_Functor_map__16849908() gopurs_runtime.Value {
	once_Data_Functor_map__16849908.Do(func() {
		cache_Data_Functor_map__16849908 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__16849908(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__16849908
}

var cache_Data_Functor_map__1668665428 gopurs_runtime.Value
var once_Data_Functor_map__1668665428 sync.Once
func Get_Data_Functor_map__1668665428() gopurs_runtime.Value {
	once_Data_Functor_map__1668665428.Do(func() {
		cache_Data_Functor_map__1668665428 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1668665428(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1668665428
}

var cache_Data_Functor_map__3778170420 gopurs_runtime.Value
var once_Data_Functor_map__3778170420 sync.Once
func Get_Data_Functor_map__3778170420() gopurs_runtime.Value {
	once_Data_Functor_map__3778170420.Do(func() {
		cache_Data_Functor_map__3778170420 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3778170420(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3778170420
}

var cache_Data_Functor_map__328307316 gopurs_runtime.Value
var once_Data_Functor_map__328307316 sync.Once
func Get_Data_Functor_map__328307316() gopurs_runtime.Value {
	once_Data_Functor_map__328307316.Do(func() {
		cache_Data_Functor_map__328307316 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__328307316(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__328307316
}

var cache_Data_Functor_map__3634441076 gopurs_runtime.Value
var once_Data_Functor_map__3634441076 sync.Once
func Get_Data_Functor_map__3634441076() gopurs_runtime.Value {
	once_Data_Functor_map__3634441076.Do(func() {
		cache_Data_Functor_map__3634441076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3634441076(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3634441076
}

var cache_Data_Functor_map__2384954036 gopurs_runtime.Value
var once_Data_Functor_map__2384954036 sync.Once
func Get_Data_Functor_map__2384954036() gopurs_runtime.Value {
	once_Data_Functor_map__2384954036.Do(func() {
		cache_Data_Functor_map__2384954036 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2384954036(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2384954036
}

var cache_Data_Functor_map__1132949076 gopurs_runtime.Value
var once_Data_Functor_map__1132949076 sync.Once
func Get_Data_Functor_map__1132949076() gopurs_runtime.Value {
	once_Data_Functor_map__1132949076.Do(func() {
		cache_Data_Functor_map__1132949076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1132949076(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1132949076
}

var cache_Data_Functor_map__138389748 gopurs_runtime.Value
var once_Data_Functor_map__138389748 sync.Once
func Get_Data_Functor_map__138389748() gopurs_runtime.Value {
	once_Data_Functor_map__138389748.Do(func() {
		cache_Data_Functor_map__138389748 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__138389748(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__138389748
}

var cache_Data_Functor_map__3061653364 gopurs_runtime.Value
var once_Data_Functor_map__3061653364 sync.Once
func Get_Data_Functor_map__3061653364() gopurs_runtime.Value {
	once_Data_Functor_map__3061653364.Do(func() {
		cache_Data_Functor_map__3061653364 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3061653364(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3061653364
}

var cache_Data_Functor_map__2701008148 gopurs_runtime.Value
var once_Data_Functor_map__2701008148 sync.Once
func Get_Data_Functor_map__2701008148() gopurs_runtime.Value {
	once_Data_Functor_map__2701008148.Do(func() {
		cache_Data_Functor_map__2701008148 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2701008148(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2701008148
}

var cache_Data_Functor_map__1256368628 gopurs_runtime.Value
var once_Data_Functor_map__1256368628 sync.Once
func Get_Data_Functor_map__1256368628() gopurs_runtime.Value {
	once_Data_Functor_map__1256368628.Do(func() {
		cache_Data_Functor_map__1256368628 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1256368628(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1256368628
}

var cache_Data_Functor_map__3975262516 gopurs_runtime.Value
var once_Data_Functor_map__3975262516 sync.Once
func Get_Data_Functor_map__3975262516() gopurs_runtime.Value {
	once_Data_Functor_map__3975262516.Do(func() {
		cache_Data_Functor_map__3975262516 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3975262516(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3975262516
}

var cache_Data_Functor_map__1762802164 gopurs_runtime.Value
var once_Data_Functor_map__1762802164 sync.Once
func Get_Data_Functor_map__1762802164() gopurs_runtime.Value {
	once_Data_Functor_map__1762802164.Do(func() {
		cache_Data_Functor_map__1762802164 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1762802164(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1762802164
}

var cache_Data_Functor_map__2562444020 gopurs_runtime.Value
var once_Data_Functor_map__2562444020 sync.Once
func Get_Data_Functor_map__2562444020() gopurs_runtime.Value {
	once_Data_Functor_map__2562444020.Do(func() {
		cache_Data_Functor_map__2562444020 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2562444020(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2562444020
}

var cache_Data_Functor_map__2212490740 gopurs_runtime.Value
var once_Data_Functor_map__2212490740 sync.Once
func Get_Data_Functor_map__2212490740() gopurs_runtime.Value {
	once_Data_Functor_map__2212490740.Do(func() {
		cache_Data_Functor_map__2212490740 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2212490740(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2212490740
}

var cache_Data_Functor_map__3674493396 gopurs_runtime.Value
var once_Data_Functor_map__3674493396 sync.Once
func Get_Data_Functor_map__3674493396() gopurs_runtime.Value {
	once_Data_Functor_map__3674493396.Do(func() {
		cache_Data_Functor_map__3674493396 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3674493396(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3674493396
}

var cache_Data_Functor_map__1052613108 gopurs_runtime.Value
var once_Data_Functor_map__1052613108 sync.Once
func Get_Data_Functor_map__1052613108() gopurs_runtime.Value {
	once_Data_Functor_map__1052613108.Do(func() {
		cache_Data_Functor_map__1052613108 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1052613108(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1052613108
}

var cache_Data_Functor_map__1483545076 gopurs_runtime.Value
var once_Data_Functor_map__1483545076 sync.Once
func Get_Data_Functor_map__1483545076() gopurs_runtime.Value {
	once_Data_Functor_map__1483545076.Do(func() {
		cache_Data_Functor_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1483545076(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1483545076
}

var cache_Data_Functor_map__3683879988 gopurs_runtime.Value
var once_Data_Functor_map__3683879988 sync.Once
func Get_Data_Functor_map__3683879988() gopurs_runtime.Value {
	once_Data_Functor_map__3683879988.Do(func() {
		cache_Data_Functor_map__3683879988 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3683879988(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3683879988
}

var cache_Data_Functor_map__3061937460 gopurs_runtime.Value
var once_Data_Functor_map__3061937460 sync.Once
func Get_Data_Functor_map__3061937460() gopurs_runtime.Value {
	once_Data_Functor_map__3061937460.Do(func() {
		cache_Data_Functor_map__3061937460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3061937460(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3061937460
}

var cache_Data_Functor_map__2749506004 gopurs_runtime.Value
var once_Data_Functor_map__2749506004 sync.Once
func Get_Data_Functor_map__2749506004() gopurs_runtime.Value {
	once_Data_Functor_map__2749506004.Do(func() {
		cache_Data_Functor_map__2749506004 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2749506004(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2749506004
}

var cache_Data_Functor_map__2458357236 gopurs_runtime.Value
var once_Data_Functor_map__2458357236 sync.Once
func Get_Data_Functor_map__2458357236() gopurs_runtime.Value {
	once_Data_Functor_map__2458357236.Do(func() {
		cache_Data_Functor_map__2458357236 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2458357236(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2458357236
}

var cache_Data_Functor_map__3172880212 gopurs_runtime.Value
var once_Data_Functor_map__3172880212 sync.Once
func Get_Data_Functor_map__3172880212() gopurs_runtime.Value {
	once_Data_Functor_map__3172880212.Do(func() {
		cache_Data_Functor_map__3172880212 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3172880212(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3172880212
}

var cache_Data_Functor_map__1319384564 gopurs_runtime.Value
var once_Data_Functor_map__1319384564 sync.Once
func Get_Data_Functor_map__1319384564() gopurs_runtime.Value {
	once_Data_Functor_map__1319384564.Do(func() {
		cache_Data_Functor_map__1319384564 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1319384564(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1319384564
}

var cache_Data_Functor_map__831829748 gopurs_runtime.Value
var once_Data_Functor_map__831829748 sync.Once
func Get_Data_Functor_map__831829748() gopurs_runtime.Value {
	once_Data_Functor_map__831829748.Do(func() {
		cache_Data_Functor_map__831829748 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__831829748(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__831829748
}

var cache_Data_Functor_map__4242765044 gopurs_runtime.Value
var once_Data_Functor_map__4242765044 sync.Once
func Get_Data_Functor_map__4242765044() gopurs_runtime.Value {
	once_Data_Functor_map__4242765044.Do(func() {
		cache_Data_Functor_map__4242765044 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4242765044(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__4242765044
}

var cache_Data_Functor_map__691697300 gopurs_runtime.Value
var once_Data_Functor_map__691697300 sync.Once
func Get_Data_Functor_map__691697300() gopurs_runtime.Value {
	once_Data_Functor_map__691697300.Do(func() {
		cache_Data_Functor_map__691697300 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__691697300(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__691697300
}

var cache_Data_Functor_map__3098878004 gopurs_runtime.Value
var once_Data_Functor_map__3098878004 sync.Once
func Get_Data_Functor_map__3098878004() gopurs_runtime.Value {
	once_Data_Functor_map__3098878004.Do(func() {
		cache_Data_Functor_map__3098878004 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3098878004(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3098878004
}

var cache_Data_Functor_map__4258206196 gopurs_runtime.Value
var once_Data_Functor_map__4258206196 sync.Once
func Get_Data_Functor_map__4258206196() gopurs_runtime.Value {
	once_Data_Functor_map__4258206196.Do(func() {
		cache_Data_Functor_map__4258206196 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4258206196(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__4258206196
}

var cache_Data_Functor_map__38773460 gopurs_runtime.Value
var once_Data_Functor_map__38773460 sync.Once
func Get_Data_Functor_map__38773460() gopurs_runtime.Value {
	once_Data_Functor_map__38773460.Do(func() {
		cache_Data_Functor_map__38773460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__38773460(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__38773460
}

var cache_Data_Functor_map__947191732 gopurs_runtime.Value
var once_Data_Functor_map__947191732 sync.Once
func Get_Data_Functor_map__947191732() gopurs_runtime.Value {
	once_Data_Functor_map__947191732.Do(func() {
		cache_Data_Functor_map__947191732 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__947191732(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__947191732
}

var cache_Data_Functor_map__1852651540 gopurs_runtime.Value
var once_Data_Functor_map__1852651540 sync.Once
func Get_Data_Functor_map__1852651540() gopurs_runtime.Value {
	once_Data_Functor_map__1852651540.Do(func() {
		cache_Data_Functor_map__1852651540 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1852651540(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1852651540
}

var cache_Data_Functor_map__1504457012 gopurs_runtime.Value
var once_Data_Functor_map__1504457012 sync.Once
func Get_Data_Functor_map__1504457012() gopurs_runtime.Value {
	once_Data_Functor_map__1504457012.Do(func() {
		cache_Data_Functor_map__1504457012 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1504457012(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1504457012
}

var cache_Data_Functor_map__2322598548 gopurs_runtime.Value
var once_Data_Functor_map__2322598548 sync.Once
func Get_Data_Functor_map__2322598548() gopurs_runtime.Value {
	once_Data_Functor_map__2322598548.Do(func() {
		cache_Data_Functor_map__2322598548 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2322598548(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2322598548
}

var cache_Data_Functor_map__2753776532 gopurs_runtime.Value
var once_Data_Functor_map__2753776532 sync.Once
func Get_Data_Functor_map__2753776532() gopurs_runtime.Value {
	once_Data_Functor_map__2753776532.Do(func() {
		cache_Data_Functor_map__2753776532 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2753776532(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2753776532
}

var cache_Data_Functor_map__657998836 gopurs_runtime.Value
var once_Data_Functor_map__657998836 sync.Once
func Get_Data_Functor_map__657998836() gopurs_runtime.Value {
	once_Data_Functor_map__657998836.Do(func() {
		cache_Data_Functor_map__657998836 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__657998836(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__657998836
}

var cache_Data_Functor_map__2300857972 gopurs_runtime.Value
var once_Data_Functor_map__2300857972 sync.Once
func Get_Data_Functor_map__2300857972() gopurs_runtime.Value {
	once_Data_Functor_map__2300857972.Do(func() {
		cache_Data_Functor_map__2300857972 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2300857972(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2300857972
}

var cache_Data_Functor_map__25069812 gopurs_runtime.Value
var once_Data_Functor_map__25069812 sync.Once
func Get_Data_Functor_map__25069812() gopurs_runtime.Value {
	once_Data_Functor_map__25069812.Do(func() {
		cache_Data_Functor_map__25069812 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__25069812(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__25069812
}

var cache_Data_Functor_map__366271444 gopurs_runtime.Value
var once_Data_Functor_map__366271444 sync.Once
func Get_Data_Functor_map__366271444() gopurs_runtime.Value {
	once_Data_Functor_map__366271444.Do(func() {
		cache_Data_Functor_map__366271444 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__366271444(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__366271444
}

var cache_Data_Functor_map__1519727060 gopurs_runtime.Value
var once_Data_Functor_map__1519727060 sync.Once
func Get_Data_Functor_map__1519727060() gopurs_runtime.Value {
	once_Data_Functor_map__1519727060.Do(func() {
		cache_Data_Functor_map__1519727060 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1519727060(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1519727060
}

var cache_Data_Functor_map__1055960852 gopurs_runtime.Value
var once_Data_Functor_map__1055960852 sync.Once
func Get_Data_Functor_map__1055960852() gopurs_runtime.Value {
	once_Data_Functor_map__1055960852.Do(func() {
		cache_Data_Functor_map__1055960852 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1055960852(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1055960852
}

var cache_Data_Functor_map__3663575028 gopurs_runtime.Value
var once_Data_Functor_map__3663575028 sync.Once
func Get_Data_Functor_map__3663575028() gopurs_runtime.Value {
	once_Data_Functor_map__3663575028.Do(func() {
		cache_Data_Functor_map__3663575028 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3663575028(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3663575028
}

var cache_Data_Functor_map__3920691508 gopurs_runtime.Value
var once_Data_Functor_map__3920691508 sync.Once
func Get_Data_Functor_map__3920691508() gopurs_runtime.Value {
	once_Data_Functor_map__3920691508.Do(func() {
		cache_Data_Functor_map__3920691508 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3920691508(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3920691508
}

var cache_Data_Functor_map__4012961076 gopurs_runtime.Value
var once_Data_Functor_map__4012961076 sync.Once
func Get_Data_Functor_map__4012961076() gopurs_runtime.Value {
	once_Data_Functor_map__4012961076.Do(func() {
		cache_Data_Functor_map__4012961076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4012961076(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__4012961076
}

var cache_Data_Functor_map__1789065812 gopurs_runtime.Value
var once_Data_Functor_map__1789065812 sync.Once
func Get_Data_Functor_map__1789065812() gopurs_runtime.Value {
	once_Data_Functor_map__1789065812.Do(func() {
		cache_Data_Functor_map__1789065812 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1789065812(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1789065812
}

var cache_Data_Functor_map__1184359732 gopurs_runtime.Value
var once_Data_Functor_map__1184359732 sync.Once
func Get_Data_Functor_map__1184359732() gopurs_runtime.Value {
	once_Data_Functor_map__1184359732.Do(func() {
		cache_Data_Functor_map__1184359732 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1184359732(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1184359732
}

var cache_Data_Functor_map__3058795348 gopurs_runtime.Value
var once_Data_Functor_map__3058795348 sync.Once
func Get_Data_Functor_map__3058795348() gopurs_runtime.Value {
	once_Data_Functor_map__3058795348.Do(func() {
		cache_Data_Functor_map__3058795348 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3058795348(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3058795348
}

var cache_Data_Functor_map__3658136916 gopurs_runtime.Value
var once_Data_Functor_map__3658136916 sync.Once
func Get_Data_Functor_map__3658136916() gopurs_runtime.Value {
	once_Data_Functor_map__3658136916.Do(func() {
		cache_Data_Functor_map__3658136916 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3658136916(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3658136916
}

var cache_Data_Functor_map__3703040820 gopurs_runtime.Value
var once_Data_Functor_map__3703040820 sync.Once
func Get_Data_Functor_map__3703040820() gopurs_runtime.Value {
	once_Data_Functor_map__3703040820.Do(func() {
		cache_Data_Functor_map__3703040820 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3703040820(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3703040820
}

var cache_Data_Functor_map__1079660148 gopurs_runtime.Value
var once_Data_Functor_map__1079660148 sync.Once
func Get_Data_Functor_map__1079660148() gopurs_runtime.Value {
	once_Data_Functor_map__1079660148.Do(func() {
		cache_Data_Functor_map__1079660148 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1079660148(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1079660148
}

var cache_Data_Functor_map__583005396 gopurs_runtime.Value
var once_Data_Functor_map__583005396 sync.Once
func Get_Data_Functor_map__583005396() gopurs_runtime.Value {
	once_Data_Functor_map__583005396.Do(func() {
		cache_Data_Functor_map__583005396 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__583005396(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__583005396
}

var cache_Data_Functor_map__1162593300 gopurs_runtime.Value
var once_Data_Functor_map__1162593300 sync.Once
func Get_Data_Functor_map__1162593300() gopurs_runtime.Value {
	once_Data_Functor_map__1162593300.Do(func() {
		cache_Data_Functor_map__1162593300 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1162593300(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1162593300
}

var cache_Data_Functor_map__528096244 gopurs_runtime.Value
var once_Data_Functor_map__528096244 sync.Once
func Get_Data_Functor_map__528096244() gopurs_runtime.Value {
	once_Data_Functor_map__528096244.Do(func() {
		cache_Data_Functor_map__528096244 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__528096244(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__528096244
}

var cache_Data_Functor_map__3228596244 gopurs_runtime.Value
var once_Data_Functor_map__3228596244 sync.Once
func Get_Data_Functor_map__3228596244() gopurs_runtime.Value {
	once_Data_Functor_map__3228596244.Do(func() {
		cache_Data_Functor_map__3228596244 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3228596244(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3228596244
}

var cache_Data_Functor_map__1729183892 gopurs_runtime.Value
var once_Data_Functor_map__1729183892 sync.Once
func Get_Data_Functor_map__1729183892() gopurs_runtime.Value {
	once_Data_Functor_map__1729183892.Do(func() {
		cache_Data_Functor_map__1729183892 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1729183892(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1729183892
}

var cache_Data_Functor_map__2745625428 gopurs_runtime.Value
var once_Data_Functor_map__2745625428 sync.Once
func Get_Data_Functor_map__2745625428() gopurs_runtime.Value {
	once_Data_Functor_map__2745625428.Do(func() {
		cache_Data_Functor_map__2745625428 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2745625428(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2745625428
}

var cache_Data_Functor_map__3384198004 gopurs_runtime.Value
var once_Data_Functor_map__3384198004 sync.Once
func Get_Data_Functor_map__3384198004() gopurs_runtime.Value {
	once_Data_Functor_map__3384198004.Do(func() {
		cache_Data_Functor_map__3384198004 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3384198004(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3384198004
}

var cache_Data_Functor_map__2190988916 gopurs_runtime.Value
var once_Data_Functor_map__2190988916 sync.Once
func Get_Data_Functor_map__2190988916() gopurs_runtime.Value {
	once_Data_Functor_map__2190988916.Do(func() {
		cache_Data_Functor_map__2190988916 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2190988916(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2190988916
}

var cache_Data_Functor_map__3592908820 gopurs_runtime.Value
var once_Data_Functor_map__3592908820 sync.Once
func Get_Data_Functor_map__3592908820() gopurs_runtime.Value {
	once_Data_Functor_map__3592908820.Do(func() {
		cache_Data_Functor_map__3592908820 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3592908820(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3592908820
}

var cache_Data_Functor_map__627657844 gopurs_runtime.Value
var once_Data_Functor_map__627657844 sync.Once
func Get_Data_Functor_map__627657844() gopurs_runtime.Value {
	once_Data_Functor_map__627657844.Do(func() {
		cache_Data_Functor_map__627657844 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__627657844(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__627657844
}

var cache_Data_Functor_map__2345808404 gopurs_runtime.Value
var once_Data_Functor_map__2345808404 sync.Once
func Get_Data_Functor_map__2345808404() gopurs_runtime.Value {
	once_Data_Functor_map__2345808404.Do(func() {
		cache_Data_Functor_map__2345808404 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2345808404(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2345808404
}

var cache_Data_Functor_map__1352087572 gopurs_runtime.Value
var once_Data_Functor_map__1352087572 sync.Once
func Get_Data_Functor_map__1352087572() gopurs_runtime.Value {
	once_Data_Functor_map__1352087572.Do(func() {
		cache_Data_Functor_map__1352087572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1352087572(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1352087572
}

var cache_Data_Functor_map__3436901780 gopurs_runtime.Value
var once_Data_Functor_map__3436901780 sync.Once
func Get_Data_Functor_map__3436901780() gopurs_runtime.Value {
	once_Data_Functor_map__3436901780.Do(func() {
		cache_Data_Functor_map__3436901780 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3436901780(__eta0_0_box, __eta1_1_box, x_2_box)
})
	})
	return cache_Data_Functor_map__3436901780
}

var cache_Data_Functor_map__381330420 gopurs_runtime.Value
var once_Data_Functor_map__381330420 sync.Once
func Get_Data_Functor_map__381330420() gopurs_runtime.Value {
	once_Data_Functor_map__381330420.Do(func() {
		cache_Data_Functor_map__381330420 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__381330420(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__381330420
}

var cache_Data_Functor_map__2870097428 gopurs_runtime.Value
var once_Data_Functor_map__2870097428 sync.Once
func Get_Data_Functor_map__2870097428() gopurs_runtime.Value {
	once_Data_Functor_map__2870097428.Do(func() {
		cache_Data_Functor_map__2870097428 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2870097428(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2870097428
}

var cache_Data_Functor_map__87655540 gopurs_runtime.Value
var once_Data_Functor_map__87655540 sync.Once
func Get_Data_Functor_map__87655540() gopurs_runtime.Value {
	once_Data_Functor_map__87655540.Do(func() {
		cache_Data_Functor_map__87655540 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__87655540(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__87655540
}

var cache_Data_Functor_map__1974414836 gopurs_runtime.Value
var once_Data_Functor_map__1974414836 sync.Once
func Get_Data_Functor_map__1974414836() gopurs_runtime.Value {
	once_Data_Functor_map__1974414836.Do(func() {
		cache_Data_Functor_map__1974414836 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1974414836(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1974414836
}

var cache_Data_Functor_map__3659954292 gopurs_runtime.Value
var once_Data_Functor_map__3659954292 sync.Once
func Get_Data_Functor_map__3659954292() gopurs_runtime.Value {
	once_Data_Functor_map__3659954292.Do(func() {
		cache_Data_Functor_map__3659954292 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3659954292(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3659954292
}

var cache_Data_Functor_map__2549863700 gopurs_runtime.Value
var once_Data_Functor_map__2549863700 sync.Once
func Get_Data_Functor_map__2549863700() gopurs_runtime.Value {
	once_Data_Functor_map__2549863700.Do(func() {
		cache_Data_Functor_map__2549863700 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2549863700(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2549863700
}

var cache_Data_Functor_map__2251722612 gopurs_runtime.Value
var once_Data_Functor_map__2251722612 sync.Once
func Get_Data_Functor_map__2251722612() gopurs_runtime.Value {
	once_Data_Functor_map__2251722612.Do(func() {
		cache_Data_Functor_map__2251722612 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2251722612(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2251722612
}

var cache_Data_Functor_map__3452061876 gopurs_runtime.Value
var once_Data_Functor_map__3452061876 sync.Once
func Get_Data_Functor_map__3452061876() gopurs_runtime.Value {
	once_Data_Functor_map__3452061876.Do(func() {
		cache_Data_Functor_map__3452061876 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3452061876(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3452061876
}

var cache_Data_Functor_map__2511708020 gopurs_runtime.Value
var once_Data_Functor_map__2511708020 sync.Once
func Get_Data_Functor_map__2511708020() gopurs_runtime.Value {
	once_Data_Functor_map__2511708020.Do(func() {
		cache_Data_Functor_map__2511708020 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2511708020(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2511708020
}

var cache_Data_Functor_map__3124798356 gopurs_runtime.Value
var once_Data_Functor_map__3124798356 sync.Once
func Get_Data_Functor_map__3124798356() gopurs_runtime.Value {
	once_Data_Functor_map__3124798356.Do(func() {
		cache_Data_Functor_map__3124798356 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3124798356(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3124798356
}

var cache_Data_Functor_map__1337616244 gopurs_runtime.Value
var once_Data_Functor_map__1337616244 sync.Once
func Get_Data_Functor_map__1337616244() gopurs_runtime.Value {
	once_Data_Functor_map__1337616244.Do(func() {
		cache_Data_Functor_map__1337616244 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1337616244(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1337616244
}

var cache_Data_Functor_map__2198784724 gopurs_runtime.Value
var once_Data_Functor_map__2198784724 sync.Once
func Get_Data_Functor_map__2198784724() gopurs_runtime.Value {
	once_Data_Functor_map__2198784724.Do(func() {
		cache_Data_Functor_map__2198784724 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2198784724(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2198784724
}

var cache_Data_Functor_map__1938733460 gopurs_runtime.Value
var once_Data_Functor_map__1938733460 sync.Once
func Get_Data_Functor_map__1938733460() gopurs_runtime.Value {
	once_Data_Functor_map__1938733460.Do(func() {
		cache_Data_Functor_map__1938733460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1938733460(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1938733460
}

var cache_Data_Functor_map__3897763604 gopurs_runtime.Value
var once_Data_Functor_map__3897763604 sync.Once
func Get_Data_Functor_map__3897763604() gopurs_runtime.Value {
	once_Data_Functor_map__3897763604.Do(func() {
		cache_Data_Functor_map__3897763604 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3897763604(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3897763604
}

var cache_Data_Functor_map__3373787924 gopurs_runtime.Value
var once_Data_Functor_map__3373787924 sync.Once
func Get_Data_Functor_map__3373787924() gopurs_runtime.Value {
	once_Data_Functor_map__3373787924.Do(func() {
		cache_Data_Functor_map__3373787924 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3373787924(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3373787924
}

var cache_Data_Functor_map__2418274292 gopurs_runtime.Value
var once_Data_Functor_map__2418274292 sync.Once
func Get_Data_Functor_map__2418274292() gopurs_runtime.Value {
	once_Data_Functor_map__2418274292.Do(func() {
		cache_Data_Functor_map__2418274292 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2418274292(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__2418274292
}

var cache_Data_Functor_map__2174973445 gopurs_runtime.Value
var once_Data_Functor_map__2174973445 sync.Once
func Get_Data_Functor_map__2174973445() gopurs_runtime.Value {
	once_Data_Functor_map__2174973445.Do(func() {
		cache_Data_Functor_map__2174973445 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2174973445(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__2174973445
}

var cache_Data_Functor_map__1152297413 gopurs_runtime.Value
var once_Data_Functor_map__1152297413 sync.Once
func Get_Data_Functor_map__1152297413() gopurs_runtime.Value {
	once_Data_Functor_map__1152297413.Do(func() {
		cache_Data_Functor_map__1152297413 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1152297413(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__1152297413
}

var cache_Data_Functor_map__3699108444 gopurs_runtime.Value
var once_Data_Functor_map__3699108444 sync.Once
func Get_Data_Functor_map__3699108444() gopurs_runtime.Value {
	once_Data_Functor_map__3699108444.Do(func() {
		cache_Data_Functor_map__3699108444 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3699108444(f_0_box, m_1_box)
})
	})
	return cache_Data_Functor_map__3699108444
}

var cache_Data_Functor_map__2579103836 gopurs_runtime.Value
var once_Data_Functor_map__2579103836 sync.Once
func Get_Data_Functor_map__2579103836() gopurs_runtime.Value {
	once_Data_Functor_map__2579103836.Do(func() {
		cache_Data_Functor_map__2579103836 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2579103836(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__2579103836
}

var cache_Data_Functor_map__1510739772 gopurs_runtime.Value
var once_Data_Functor_map__1510739772 sync.Once
func Get_Data_Functor_map__1510739772() gopurs_runtime.Value {
	once_Data_Functor_map__1510739772.Do(func() {
		cache_Data_Functor_map__1510739772 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1510739772(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__1510739772
}

var cache_Data_Functor_map__1208755772 gopurs_runtime.Value
var once_Data_Functor_map__1208755772 sync.Once
func Get_Data_Functor_map__1208755772() gopurs_runtime.Value {
	once_Data_Functor_map__1208755772.Do(func() {
		cache_Data_Functor_map__1208755772 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1208755772(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__1208755772
}

var cache_Data_Functor_map__3565923196 gopurs_runtime.Value
var once_Data_Functor_map__3565923196 sync.Once
func Get_Data_Functor_map__3565923196() gopurs_runtime.Value {
	once_Data_Functor_map__3565923196.Do(func() {
		cache_Data_Functor_map__3565923196 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3565923196(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__3565923196
}

var cache_Data_Functor_map__2597050044 gopurs_runtime.Value
var once_Data_Functor_map__2597050044 sync.Once
func Get_Data_Functor_map__2597050044() gopurs_runtime.Value {
	once_Data_Functor_map__2597050044.Do(func() {
		cache_Data_Functor_map__2597050044 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2597050044(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__2597050044
}

var cache_Data_Functor_map__3467322428 gopurs_runtime.Value
var once_Data_Functor_map__3467322428 sync.Once
func Get_Data_Functor_map__3467322428() gopurs_runtime.Value {
	once_Data_Functor_map__3467322428.Do(func() {
		cache_Data_Functor_map__3467322428 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3467322428(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__3467322428
}

var cache_Data_Functor_map__1958856956 gopurs_runtime.Value
var once_Data_Functor_map__1958856956 sync.Once
func Get_Data_Functor_map__1958856956() gopurs_runtime.Value {
	once_Data_Functor_map__1958856956.Do(func() {
		cache_Data_Functor_map__1958856956 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1958856956(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__1958856956
}

var cache_Data_Functor_map__1422050556 gopurs_runtime.Value
var once_Data_Functor_map__1422050556 sync.Once
func Get_Data_Functor_map__1422050556() gopurs_runtime.Value {
	once_Data_Functor_map__1422050556.Do(func() {
		cache_Data_Functor_map__1422050556 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1422050556(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__1422050556
}

var cache_Data_Functor_map__109003388 gopurs_runtime.Value
var once_Data_Functor_map__109003388 sync.Once
func Get_Data_Functor_map__109003388() gopurs_runtime.Value {
	once_Data_Functor_map__109003388.Do(func() {
		cache_Data_Functor_map__109003388 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__109003388(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__109003388
}

var cache_Data_Functor_map__829570556 gopurs_runtime.Value
var once_Data_Functor_map__829570556 sync.Once
func Get_Data_Functor_map__829570556() gopurs_runtime.Value {
	once_Data_Functor_map__829570556.Do(func() {
		cache_Data_Functor_map__829570556 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__829570556(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__829570556
}

var cache_Data_Functor_map__2156385148 gopurs_runtime.Value
var once_Data_Functor_map__2156385148 sync.Once
func Get_Data_Functor_map__2156385148() gopurs_runtime.Value {
	once_Data_Functor_map__2156385148.Do(func() {
		cache_Data_Functor_map__2156385148 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2156385148(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__2156385148
}

var cache_Data_Functor_map__558976860 gopurs_runtime.Value
var once_Data_Functor_map__558976860 sync.Once
func Get_Data_Functor_map__558976860() gopurs_runtime.Value {
	once_Data_Functor_map__558976860.Do(func() {
		cache_Data_Functor_map__558976860 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__558976860(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__558976860
}

var cache_Data_Functor_map__1806510684 gopurs_runtime.Value
var once_Data_Functor_map__1806510684 sync.Once
func Get_Data_Functor_map__1806510684() gopurs_runtime.Value {
	once_Data_Functor_map__1806510684.Do(func() {
		cache_Data_Functor_map__1806510684 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1806510684(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__1806510684
}

var cache_Data_Functor_map__596534652 gopurs_runtime.Value
var once_Data_Functor_map__596534652 sync.Once
func Get_Data_Functor_map__596534652() gopurs_runtime.Value {
	once_Data_Functor_map__596534652.Do(func() {
		cache_Data_Functor_map__596534652 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__596534652(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__596534652
}

var cache_Data_Functor_map__3815458588 gopurs_runtime.Value
var once_Data_Functor_map__3815458588 sync.Once
func Get_Data_Functor_map__3815458588() gopurs_runtime.Value {
	once_Data_Functor_map__3815458588.Do(func() {
		cache_Data_Functor_map__3815458588 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3815458588(f_0_box, l_1_box)
})
	})
	return cache_Data_Functor_map__3815458588
}

var cache_Data_Functor_map__843173928 gopurs_runtime.Value
var once_Data_Functor_map__843173928 sync.Once
func Get_Data_Functor_map__843173928() gopurs_runtime.Value {
	once_Data_Functor_map__843173928.Do(func() {
		cache_Data_Functor_map__843173928 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__843173928(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__843173928
}

var cache_Data_Functor_map__438443400 gopurs_runtime.Value
var once_Data_Functor_map__438443400 sync.Once
func Get_Data_Functor_map__438443400() gopurs_runtime.Value {
	once_Data_Functor_map__438443400.Do(func() {
		cache_Data_Functor_map__438443400 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__438443400(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__438443400
}

var cache_Data_Functor_map__1081746216 gopurs_runtime.Value
var once_Data_Functor_map__1081746216 sync.Once
func Get_Data_Functor_map__1081746216() gopurs_runtime.Value {
	once_Data_Functor_map__1081746216.Do(func() {
		cache_Data_Functor_map__1081746216 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1081746216(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1081746216
}

var cache_Data_Functor_map__95558920 gopurs_runtime.Value
var once_Data_Functor_map__95558920 sync.Once
func Get_Data_Functor_map__95558920() gopurs_runtime.Value {
	once_Data_Functor_map__95558920.Do(func() {
		cache_Data_Functor_map__95558920 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__95558920(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__95558920
}

var cache_Data_Functor_map__291265340 gopurs_runtime.Value
var once_Data_Functor_map__291265340 sync.Once
func Get_Data_Functor_map__291265340() gopurs_runtime.Value {
	once_Data_Functor_map__291265340.Do(func() {
		cache_Data_Functor_map__291265340 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__291265340(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__291265340
}

var cache_Data_Functor_map__2107538812 gopurs_runtime.Value
var once_Data_Functor_map__2107538812 sync.Once
func Get_Data_Functor_map__2107538812() gopurs_runtime.Value {
	once_Data_Functor_map__2107538812.Do(func() {
		cache_Data_Functor_map__2107538812 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2107538812(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2107538812
}

var cache_Data_Functor_map__3447677596 gopurs_runtime.Value
var once_Data_Functor_map__3447677596 sync.Once
func Get_Data_Functor_map__3447677596() gopurs_runtime.Value {
	once_Data_Functor_map__3447677596.Do(func() {
		cache_Data_Functor_map__3447677596 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__3447677596(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__3447677596
}

var cache_Data_Functor_map__2165919164 gopurs_runtime.Value
var once_Data_Functor_map__2165919164 sync.Once
func Get_Data_Functor_map__2165919164() gopurs_runtime.Value {
	once_Data_Functor_map__2165919164.Do(func() {
		cache_Data_Functor_map__2165919164 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2165919164(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2165919164
}

var cache_Data_Functor_map__4155962236 gopurs_runtime.Value
var once_Data_Functor_map__4155962236 sync.Once
func Get_Data_Functor_map__4155962236() gopurs_runtime.Value {
	once_Data_Functor_map__4155962236.Do(func() {
		cache_Data_Functor_map__4155962236 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__4155962236(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__4155962236
}

var cache_Data_Functor_map__2389909756 gopurs_runtime.Value
var once_Data_Functor_map__2389909756 sync.Once
func Get_Data_Functor_map__2389909756() gopurs_runtime.Value {
	once_Data_Functor_map__2389909756.Do(func() {
		cache_Data_Functor_map__2389909756 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2389909756(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2389909756
}

var cache_Data_Functor_map__1759928220 gopurs_runtime.Value
var once_Data_Functor_map__1759928220 sync.Once
func Get_Data_Functor_map__1759928220() gopurs_runtime.Value {
	once_Data_Functor_map__1759928220.Do(func() {
		cache_Data_Functor_map__1759928220 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__1759928220(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__1759928220
}

var cache_Data_Functor_map__901270812 gopurs_runtime.Value
var once_Data_Functor_map__901270812 sync.Once
func Get_Data_Functor_map__901270812() gopurs_runtime.Value {
	once_Data_Functor_map__901270812.Do(func() {
		cache_Data_Functor_map__901270812 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__901270812(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__901270812
}

var cache_Data_Functor_map__2486200924 gopurs_runtime.Value
var once_Data_Functor_map__2486200924 sync.Once
func Get_Data_Functor_map__2486200924() gopurs_runtime.Value {
	once_Data_Functor_map__2486200924.Do(func() {
		cache_Data_Functor_map__2486200924 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2486200924(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2486200924
}

var cache_Data_Functor_map__2670646620 gopurs_runtime.Value
var once_Data_Functor_map__2670646620 sync.Once
func Get_Data_Functor_map__2670646620() gopurs_runtime.Value {
	once_Data_Functor_map__2670646620.Do(func() {
		cache_Data_Functor_map__2670646620 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2670646620(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2670646620
}

var cache_Data_Functor_map__1887399228 gopurs_runtime.Value
var once_Data_Functor_map__1887399228 sync.Once
func Get_Data_Functor_map__1887399228() gopurs_runtime.Value {
	once_Data_Functor_map__1887399228.Do(func() {
		cache_Data_Functor_map__1887399228 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__1887399228(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__1887399228
}

var cache_Data_Functor_map__2565844412 gopurs_runtime.Value
var once_Data_Functor_map__2565844412 sync.Once
func Get_Data_Functor_map__2565844412() gopurs_runtime.Value {
	once_Data_Functor_map__2565844412.Do(func() {
		cache_Data_Functor_map__2565844412 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2565844412(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2565844412
}

var cache_Data_Functor_map__2126568188 gopurs_runtime.Value
var once_Data_Functor_map__2126568188 sync.Once
func Get_Data_Functor_map__2126568188() gopurs_runtime.Value {
	once_Data_Functor_map__2126568188.Do(func() {
		cache_Data_Functor_map__2126568188 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2126568188(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2126568188
}

var cache_Data_Functor_map__1739124124 gopurs_runtime.Value
var once_Data_Functor_map__1739124124 sync.Once
func Get_Data_Functor_map__1739124124() gopurs_runtime.Value {
	once_Data_Functor_map__1739124124.Do(func() {
		cache_Data_Functor_map__1739124124 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__1739124124(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__1739124124
}

var cache_Data_Functor_map__4039429788 gopurs_runtime.Value
var once_Data_Functor_map__4039429788 sync.Once
func Get_Data_Functor_map__4039429788() gopurs_runtime.Value {
	once_Data_Functor_map__4039429788.Do(func() {
		cache_Data_Functor_map__4039429788 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__4039429788(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__4039429788
}

var cache_Data_Functor_map__1208952924 gopurs_runtime.Value
var once_Data_Functor_map__1208952924 sync.Once
func Get_Data_Functor_map__1208952924() gopurs_runtime.Value {
	once_Data_Functor_map__1208952924.Do(func() {
		cache_Data_Functor_map__1208952924 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__1208952924(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__1208952924
}

var cache_Data_Functor_map__2294788636 gopurs_runtime.Value
var once_Data_Functor_map__2294788636 sync.Once
func Get_Data_Functor_map__2294788636() gopurs_runtime.Value {
	once_Data_Functor_map__2294788636.Do(func() {
		cache_Data_Functor_map__2294788636 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2294788636(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2294788636
}

var cache_Data_Functor_map__48293596 gopurs_runtime.Value
var once_Data_Functor_map__48293596 sync.Once
func Get_Data_Functor_map__48293596() gopurs_runtime.Value {
	once_Data_Functor_map__48293596.Do(func() {
		cache_Data_Functor_map__48293596 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__48293596(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__48293596
}

var cache_Data_Functor_map__853141532 gopurs_runtime.Value
var once_Data_Functor_map__853141532 sync.Once
func Get_Data_Functor_map__853141532() gopurs_runtime.Value {
	once_Data_Functor_map__853141532.Do(func() {
		cache_Data_Functor_map__853141532 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__853141532(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__853141532
}

var cache_Data_Functor_map__2275717084 gopurs_runtime.Value
var once_Data_Functor_map__2275717084 sync.Once
func Get_Data_Functor_map__2275717084() gopurs_runtime.Value {
	once_Data_Functor_map__2275717084.Do(func() {
		cache_Data_Functor_map__2275717084 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2275717084(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2275717084
}

var cache_Data_Functor_map__3733923228 gopurs_runtime.Value
var once_Data_Functor_map__3733923228 sync.Once
func Get_Data_Functor_map__3733923228() gopurs_runtime.Value {
	once_Data_Functor_map__3733923228.Do(func() {
		cache_Data_Functor_map__3733923228 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__3733923228(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__3733923228
}

var cache_Data_Functor_map__1171574364 gopurs_runtime.Value
var once_Data_Functor_map__1171574364 sync.Once
func Get_Data_Functor_map__1171574364() gopurs_runtime.Value {
	once_Data_Functor_map__1171574364.Do(func() {
		cache_Data_Functor_map__1171574364 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__1171574364(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__1171574364
}

var cache_Data_Functor_map__63598588 gopurs_runtime.Value
var once_Data_Functor_map__63598588 sync.Once
func Get_Data_Functor_map__63598588() gopurs_runtime.Value {
	once_Data_Functor_map__63598588.Do(func() {
		cache_Data_Functor_map__63598588 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__63598588(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__63598588
}

var cache_Data_Functor_map__1808515292 gopurs_runtime.Value
var once_Data_Functor_map__1808515292 sync.Once
func Get_Data_Functor_map__1808515292() gopurs_runtime.Value {
	once_Data_Functor_map__1808515292.Do(func() {
		cache_Data_Functor_map__1808515292 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__1808515292(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__1808515292
}

var cache_Data_Functor_map__140514012 gopurs_runtime.Value
var once_Data_Functor_map__140514012 sync.Once
func Get_Data_Functor_map__140514012() gopurs_runtime.Value {
	once_Data_Functor_map__140514012.Do(func() {
		cache_Data_Functor_map__140514012 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__140514012(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__140514012
}

var cache_Data_Functor_map__3210082748 gopurs_runtime.Value
var once_Data_Functor_map__3210082748 sync.Once
func Get_Data_Functor_map__3210082748() gopurs_runtime.Value {
	once_Data_Functor_map__3210082748.Do(func() {
		cache_Data_Functor_map__3210082748 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__3210082748(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__3210082748
}

var cache_Data_Functor_map__2034458684 gopurs_runtime.Value
var once_Data_Functor_map__2034458684 sync.Once
func Get_Data_Functor_map__2034458684() gopurs_runtime.Value {
	once_Data_Functor_map__2034458684.Do(func() {
		cache_Data_Functor_map__2034458684 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2034458684(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2034458684
}

var cache_Data_Functor_map__2615158204 gopurs_runtime.Value
var once_Data_Functor_map__2615158204 sync.Once
func Get_Data_Functor_map__2615158204() gopurs_runtime.Value {
	once_Data_Functor_map__2615158204.Do(func() {
		cache_Data_Functor_map__2615158204 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2615158204(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2615158204
}

var cache_Data_Functor_map__125648636 gopurs_runtime.Value
var once_Data_Functor_map__125648636 sync.Once
func Get_Data_Functor_map__125648636() gopurs_runtime.Value {
	once_Data_Functor_map__125648636.Do(func() {
		cache_Data_Functor_map__125648636 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__125648636(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__125648636
}

var cache_Data_Functor_map__539092636 gopurs_runtime.Value
var once_Data_Functor_map__539092636 sync.Once
func Get_Data_Functor_map__539092636() gopurs_runtime.Value {
	once_Data_Functor_map__539092636.Do(func() {
		cache_Data_Functor_map__539092636 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__539092636(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__539092636
}

var cache_Data_Functor_map__1980149980 gopurs_runtime.Value
var once_Data_Functor_map__1980149980 sync.Once
func Get_Data_Functor_map__1980149980() gopurs_runtime.Value {
	once_Data_Functor_map__1980149980.Do(func() {
		cache_Data_Functor_map__1980149980 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__1980149980(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__1980149980
}

var cache_Data_Functor_map__2202537180 gopurs_runtime.Value
var once_Data_Functor_map__2202537180 sync.Once
func Get_Data_Functor_map__2202537180() gopurs_runtime.Value {
	once_Data_Functor_map__2202537180.Do(func() {
		cache_Data_Functor_map__2202537180 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2202537180(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2202537180
}

var cache_Data_Functor_map__1681779388 gopurs_runtime.Value
var once_Data_Functor_map__1681779388 sync.Once
func Get_Data_Functor_map__1681779388() gopurs_runtime.Value {
	once_Data_Functor_map__1681779388.Do(func() {
		cache_Data_Functor_map__1681779388 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__1681779388(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__1681779388
}

var cache_Data_Functor_map__1206962620 gopurs_runtime.Value
var once_Data_Functor_map__1206962620 sync.Once
func Get_Data_Functor_map__1206962620() gopurs_runtime.Value {
	once_Data_Functor_map__1206962620.Do(func() {
		cache_Data_Functor_map__1206962620 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__1206962620(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__1206962620
}

var cache_Data_Functor_map__3486165692 gopurs_runtime.Value
var once_Data_Functor_map__3486165692 sync.Once
func Get_Data_Functor_map__3486165692() gopurs_runtime.Value {
	once_Data_Functor_map__3486165692.Do(func() {
		cache_Data_Functor_map__3486165692 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__3486165692(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__3486165692
}

var cache_Data_Functor_map__2116777468 gopurs_runtime.Value
var once_Data_Functor_map__2116777468 sync.Once
func Get_Data_Functor_map__2116777468() gopurs_runtime.Value {
	once_Data_Functor_map__2116777468.Do(func() {
		cache_Data_Functor_map__2116777468 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2116777468(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__2116777468
}

var cache_Data_Functor_map__316107900 gopurs_runtime.Value
var once_Data_Functor_map__316107900 sync.Once
func Get_Data_Functor_map__316107900() gopurs_runtime.Value {
	once_Data_Functor_map__316107900.Do(func() {
		cache_Data_Functor_map__316107900 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__316107900(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Data_Functor_map__316107900
}

var cache_Data_Functor_map__3269387708 gopurs_runtime.Value
var once_Data_Functor_map__3269387708 sync.Once
func Get_Data_Functor_map__3269387708() gopurs_runtime.Value {
	once_Data_Functor_map__3269387708.Do(func() {
		cache_Data_Functor_map__3269387708 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3269387708(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__3269387708
}

var cache_Data_Functor_map__271334204 gopurs_runtime.Value
var once_Data_Functor_map__271334204 sync.Once
func Get_Data_Functor_map__271334204() gopurs_runtime.Value {
	once_Data_Functor_map__271334204.Do(func() {
		cache_Data_Functor_map__271334204 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__271334204(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__271334204
}

var cache_Data_Functor_map__91618268 gopurs_runtime.Value
var once_Data_Functor_map__91618268 sync.Once
func Get_Data_Functor_map__91618268() gopurs_runtime.Value {
	once_Data_Functor_map__91618268.Do(func() {
		cache_Data_Functor_map__91618268 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__91618268(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__91618268
}

var cache_Data_Functor_map__2311960860 gopurs_runtime.Value
var once_Data_Functor_map__2311960860 sync.Once
func Get_Data_Functor_map__2311960860() gopurs_runtime.Value {
	once_Data_Functor_map__2311960860.Do(func() {
		cache_Data_Functor_map__2311960860 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Functor_map__2311960860(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](m_1_box)))}
})
	})
	return cache_Data_Functor_map__2311960860
}

var cache_Data_Functor_map__339096027 gopurs_runtime.Value
var once_Data_Functor_map__339096027 sync.Once
func Get_Data_Functor_map__339096027() gopurs_runtime.Value {
	once_Data_Functor_map__339096027.Do(func() {
		cache_Data_Functor_map__339096027 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__339096027(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__339096027
}

var cache_Data_Functor_map__2177087003 gopurs_runtime.Value
var once_Data_Functor_map__2177087003 sync.Once
func Get_Data_Functor_map__2177087003() gopurs_runtime.Value {
	once_Data_Functor_map__2177087003.Do(func() {
		cache_Data_Functor_map__2177087003 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2177087003(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__2177087003
}

var cache_Data_Functor_map__2919116915 gopurs_runtime.Value
var once_Data_Functor_map__2919116915 sync.Once
func Get_Data_Functor_map__2919116915() gopurs_runtime.Value {
	once_Data_Functor_map__2919116915.Do(func() {
		cache_Data_Functor_map__2919116915 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2919116915(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__2919116915
}

var cache_Data_Functor_map__2261469235 gopurs_runtime.Value
var once_Data_Functor_map__2261469235 sync.Once
func Get_Data_Functor_map__2261469235() gopurs_runtime.Value {
	once_Data_Functor_map__2261469235.Do(func() {
		cache_Data_Functor_map__2261469235 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2261469235(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__2261469235
}

var cache_Data_Functor_map__3644121587 gopurs_runtime.Value
var once_Data_Functor_map__3644121587 sync.Once
func Get_Data_Functor_map__3644121587() gopurs_runtime.Value {
	once_Data_Functor_map__3644121587.Do(func() {
		cache_Data_Functor_map__3644121587 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3644121587(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__3644121587
}

var cache_Data_Functor_map__113987891 gopurs_runtime.Value
var once_Data_Functor_map__113987891 sync.Once
func Get_Data_Functor_map__113987891() gopurs_runtime.Value {
	once_Data_Functor_map__113987891.Do(func() {
		cache_Data_Functor_map__113987891 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__113987891(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__113987891
}

var cache_Data_Functor_map__3065908595 gopurs_runtime.Value
var once_Data_Functor_map__3065908595 sync.Once
func Get_Data_Functor_map__3065908595() gopurs_runtime.Value {
	once_Data_Functor_map__3065908595.Do(func() {
		cache_Data_Functor_map__3065908595 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3065908595(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__3065908595
}

var cache_Data_Functor_map__4102685939 gopurs_runtime.Value
var once_Data_Functor_map__4102685939 sync.Once
func Get_Data_Functor_map__4102685939() gopurs_runtime.Value {
	once_Data_Functor_map__4102685939.Do(func() {
		cache_Data_Functor_map__4102685939 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4102685939(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__4102685939
}

var cache_Data_Functor_map__173660595 gopurs_runtime.Value
var once_Data_Functor_map__173660595 sync.Once
func Get_Data_Functor_map__173660595() gopurs_runtime.Value {
	once_Data_Functor_map__173660595.Do(func() {
		cache_Data_Functor_map__173660595 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__173660595(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__173660595
}

var cache_Data_Functor_map__1484761587 gopurs_runtime.Value
var once_Data_Functor_map__1484761587 sync.Once
func Get_Data_Functor_map__1484761587() gopurs_runtime.Value {
	once_Data_Functor_map__1484761587.Do(func() {
		cache_Data_Functor_map__1484761587 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1484761587(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Functor_map__1484761587
}

var cache_Data_Functor_map__1678245779 gopurs_runtime.Value
var once_Data_Functor_map__1678245779 sync.Once
func Get_Data_Functor_map__1678245779() gopurs_runtime.Value {
	once_Data_Functor_map__1678245779.Do(func() {
		cache_Data_Functor_map__1678245779 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1678245779(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__1678245779
}

var cache_Data_Functor_map__3058064980 gopurs_runtime.Value
var once_Data_Functor_map__3058064980 sync.Once
func Get_Data_Functor_map__3058064980() gopurs_runtime.Value {
	once_Data_Functor_map__3058064980.Do(func() {
		cache_Data_Functor_map__3058064980 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3058064980(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dict_0_box))
})
	})
	return cache_Data_Functor_map__3058064980
}

var cache_Data_Functor_mapFlipped__260821093 gopurs_runtime.Value
var once_Data_Functor_mapFlipped__260821093 sync.Once
func Get_Data_Functor_mapFlipped__260821093() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped__260821093.Do(func() {
		cache_Data_Functor_mapFlipped__260821093 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_mapFlipped__260821093(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_Data_Functor_mapFlipped__260821093
}

var cache_Data_Functor_mapFlipped__2466386789 gopurs_runtime.Value
var once_Data_Functor_mapFlipped__2466386789 sync.Once
func Get_Data_Functor_mapFlipped__2466386789() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped__2466386789.Do(func() {
		cache_Data_Functor_mapFlipped__2466386789 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_mapFlipped__2466386789(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_Data_Functor_mapFlipped__2466386789
}

var cache_Data_Functor_mapFlipped__4215217780 gopurs_runtime.Value
var once_Data_Functor_mapFlipped__4215217780 sync.Once
func Get_Data_Functor_mapFlipped__4215217780() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped__4215217780.Do(func() {
		cache_Data_Functor_mapFlipped__4215217780 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_mapFlipped__4215217780(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_Data_Functor_mapFlipped__4215217780
}

var cache_Data_Functor_mapFlipped__742928244 gopurs_runtime.Value
var once_Data_Functor_mapFlipped__742928244 sync.Once
func Get_Data_Functor_mapFlipped__742928244() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped__742928244.Do(func() {
		cache_Data_Functor_mapFlipped__742928244 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_mapFlipped__742928244(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_Data_Functor_mapFlipped__742928244
}

var cache_Data_Functor_mapFlipped__509401044 gopurs_runtime.Value
var once_Data_Functor_mapFlipped__509401044 sync.Once
func Get_Data_Functor_mapFlipped__509401044() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped__509401044.Do(func() {
		cache_Data_Functor_mapFlipped__509401044 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_mapFlipped__509401044(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_Data_Functor_mapFlipped__509401044
}

var cache_Data_Functor_mapFlipped__2213324916 gopurs_runtime.Value
var once_Data_Functor_mapFlipped__2213324916 sync.Once
func Get_Data_Functor_mapFlipped__2213324916() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped__2213324916.Do(func() {
		cache_Data_Functor_mapFlipped__2213324916 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_mapFlipped__2213324916(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_Data_Functor_mapFlipped__2213324916
}

var cache_Data_Functor_mapFlipped__3249733428 gopurs_runtime.Value
var once_Data_Functor_mapFlipped__3249733428 sync.Once
func Get_Data_Functor_mapFlipped__3249733428() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped__3249733428.Do(func() {
		cache_Data_Functor_mapFlipped__3249733428 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_mapFlipped__3249733428(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_Data_Functor_mapFlipped__3249733428
}

var cache_Data_Functor_mapFlipped__1087756276 gopurs_runtime.Value
var once_Data_Functor_mapFlipped__1087756276 sync.Once
func Get_Data_Functor_mapFlipped__1087756276() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped__1087756276.Do(func() {
		cache_Data_Functor_mapFlipped__1087756276 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_mapFlipped__1087756276(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_Data_Functor_mapFlipped__1087756276
}

var cache_Data_Functor_mapFlipped__2919806324 gopurs_runtime.Value
var once_Data_Functor_mapFlipped__2919806324 sync.Once
func Get_Data_Functor_mapFlipped__2919806324() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped__2919806324.Do(func() {
		cache_Data_Functor_mapFlipped__2919806324 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_mapFlipped__2919806324(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_Data_Functor_mapFlipped__2919806324
}

var cache_Data_Functor_void__3020373336 gopurs_runtime.Value
var once_Data_Functor_void__3020373336 sync.Once
func Get_Data_Functor_void__3020373336() gopurs_runtime.Value {
	once_Data_Functor_void__3020373336.Do(func() {
		cache_Data_Functor_void__3020373336 = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__3020373336(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box))
})
	})
	return cache_Data_Functor_void__3020373336
}

var cache_Data_Functor_void__2104786761 gopurs_runtime.Value
var once_Data_Functor_void__2104786761 sync.Once
func Get_Data_Functor_void__2104786761() gopurs_runtime.Value {
	once_Data_Functor_void__2104786761.Do(func() {
		cache_Data_Functor_void__2104786761 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2104786761(__eta0_0_box)
})
	})
	return cache_Data_Functor_void__2104786761
}

var cache_Data_Functor_voidLeft__171362140 gopurs_runtime.Value
var once_Data_Functor_voidLeft__171362140 sync.Once
func Get_Data_Functor_voidLeft__171362140() gopurs_runtime.Value {
	once_Data_Functor_voidLeft__171362140.Do(func() {
		cache_Data_Functor_voidLeft__171362140 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_voidLeft__171362140(f_0_box, x_1_box.StrVal())
})
	})
	return cache_Data_Functor_voidLeft__171362140
}

var cache_Data_Functor_voidLeft__32301756 gopurs_runtime.Value
var once_Data_Functor_voidLeft__32301756 sync.Once
func Get_Data_Functor_voidLeft__32301756() gopurs_runtime.Value {
	once_Data_Functor_voidLeft__32301756.Do(func() {
		cache_Data_Functor_voidLeft__32301756 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_voidLeft__32301756(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, x_2_box)
})
	})
	return cache_Data_Functor_voidLeft__32301756
}

var cache_Data_Functor_voidLeft__3507060164 gopurs_runtime.Value
var once_Data_Functor_voidLeft__3507060164 sync.Once
func Get_Data_Functor_voidLeft__3507060164() gopurs_runtime.Value {
	once_Data_Functor_voidLeft__3507060164.Do(func() {
		cache_Data_Functor_voidLeft__3507060164 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_voidLeft__3507060164(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](f_0_box), x_1_box))}
})
	})
	return cache_Data_Functor_voidLeft__3507060164
}

var cache_Data_Functor_voidLeft__4152282100 gopurs_runtime.Value
var once_Data_Functor_voidLeft__4152282100 sync.Once
func Get_Data_Functor_voidLeft__4152282100() gopurs_runtime.Value {
	once_Data_Functor_voidLeft__4152282100.Do(func() {
		cache_Data_Functor_voidLeft__4152282100 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_voidLeft__4152282100(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](f_0_box), x_1_box))}
})
	})
	return cache_Data_Functor_voidLeft__4152282100
}

var cache_Data_Functor_voidRight__3258033404 gopurs_runtime.Value
var once_Data_Functor_voidRight__3258033404 sync.Once
func Get_Data_Functor_voidRight__3258033404() gopurs_runtime.Value {
	once_Data_Functor_voidRight__3258033404.Do(func() {
		cache_Data_Functor_voidRight__3258033404 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_voidRight__3258033404(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), x_1_box)
})
	})
	return cache_Data_Functor_voidRight__3258033404
}

var cache_Data_Functor_voidRight__1142845180 gopurs_runtime.Value
var once_Data_Functor_voidRight__1142845180 sync.Once
func Get_Data_Functor_voidRight__1142845180() gopurs_runtime.Value {
	once_Data_Functor_voidRight__1142845180.Do(func() {
		cache_Data_Functor_voidRight__1142845180 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_voidRight__1142845180(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), x_1_box)
})
	})
	return cache_Data_Functor_voidRight__1142845180
}

var cache_Data_Functor_voidRight__698766972 gopurs_runtime.Value
var once_Data_Functor_voidRight__698766972 sync.Once
func Get_Data_Functor_voidRight__698766972() gopurs_runtime.Value {
	once_Data_Functor_voidRight__698766972.Do(func() {
		cache_Data_Functor_voidRight__698766972 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_voidRight__698766972(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), x_1_box)
})
	})
	return cache_Data_Functor_voidRight__698766972
}

var cache_Data_Functor_voidRight__2767609084 gopurs_runtime.Value
var once_Data_Functor_voidRight__2767609084 sync.Once
func Get_Data_Functor_voidRight__2767609084() gopurs_runtime.Value {
	once_Data_Functor_voidRight__2767609084.Do(func() {
		cache_Data_Functor_voidRight__2767609084 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_voidRight__2767609084(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), x_1_box)
})
	})
	return cache_Data_Functor_voidRight__2767609084
}

type Constructor_Data_Functor_Functor struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[929368378] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Functor_Functor)(ptr)
		_ = c
		switch key {
		case "map": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Functor_Functor: " + key)
		}
	}
}


func Call_Data_Functor_Functor_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Functor_go__map(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_mapFlipped(dictFunctor_0_loop *Constructor_Data_Functor_Functor, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_2, fa_1)
}

func Call_Data_Functor_void(dictFunctor_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
}

func Call_Data_Functor_voidLeft(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), f_1)
}

func Call_Data_Functor_voidRight(dictFunctor_0_loop *Constructor_Data_Functor_Functor, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Functor_flap(dictFunctor_0_loop *Constructor_Data_Functor_Functor, ff_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var ff_1 gopurs_runtime.Value = ff_1_loop
_ = ff_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, x_2)
}), ff_1)
}

func Call_Data_Functor_map__1165794789(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1162721797(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1924492325(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2805967941(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3676941189(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2364682565(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2321705669(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3658399301(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__4176008549(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__4040535013(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3116241637(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__4282869861(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2539978757(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__4119868389(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3122478373(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2665381605(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__4104571781(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1599448997(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__4285761829(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3149795237(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1184159621(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__609712645(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1171202917(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__677918245(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1307706501(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1408505925(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1542634789(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2675323109(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__258070885(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__47904357(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3871729957(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__67411525(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__321096773(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2876470885(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1434901492(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__243231988(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2278567252(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__491003380(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__358342900(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2224957140(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3483556436(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2869957716(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2475811444(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3218804116(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3240628980(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2199395572(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__16849908(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1668665428(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3778170420(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__328307316(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3634441076(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2384954036(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1132949076(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__138389748(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3061653364(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2701008148(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1256368628(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3975262516(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1762802164(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2562444020(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2212490740(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3674493396(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1052613108(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1483545076(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3683879988(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3061937460(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2749506004(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2458357236(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3172880212(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1319384564(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__831829748(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__4242765044(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__691697300(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3098878004(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__4258206196(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__38773460(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__947191732(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return func() gopurs_runtime.Value {
arr_val_arrayMap0 := __eta1_1
_ = arr_val_arrayMap0
arr_go_arrayMap0 := (*[]gopurs_runtime.Value)(arr_val_arrayMap0.UnsafePtr)
_ = arr_go_arrayMap0
res_go_arrayMap0 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap0))
_ = res_go_arrayMap0
for i_arrayMap0, v_arrayMap0 := range *arr_go_arrayMap0 {
res_go_arrayMap0[i_arrayMap0] = gopurs_runtime.Apply(__eta0_0, v_arrayMap0)
}
return gopurs_runtime.Array(res_go_arrayMap0)
}()
}

func Call_Data_Functor_map__1852651540(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1504457012(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2322598548(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2753776532(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__657998836(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2300857972(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__25069812(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__366271444(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1519727060(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1055960852(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3663575028(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3920691508(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__4012961076(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1789065812(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1184359732(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3058795348(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3658136916(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3703040820(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1079660148(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__583005396(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1162593300(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__528096244(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3228596244(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1729183892(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2745625428(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3384198004(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2190988916(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3592908820(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__627657844(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return func() gopurs_runtime.Value {
arr_val_arrayMap0 := __eta1_1
_ = arr_val_arrayMap0
arr_go_arrayMap0 := (*[]gopurs_runtime.Value)(arr_val_arrayMap0.UnsafePtr)
_ = arr_go_arrayMap0
res_go_arrayMap0 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap0))
_ = res_go_arrayMap0
for i_arrayMap0, v_arrayMap0 := range *arr_go_arrayMap0 {
res_go_arrayMap0[i_arrayMap0] = gopurs_runtime.Apply(__eta0_0, v_arrayMap0)
}
return gopurs_runtime.Array(res_go_arrayMap0)
}()
}

func Call_Data_Functor_map__2345808404(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1352087572(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3436901780(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
}

func Call_Data_Functor_map__381330420(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2870097428(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__87655540(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1974414836(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3659954292(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2549863700(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2251722612(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3452061876(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2511708020(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3124798356(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1337616244(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2198784724(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__1938733460(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3897763604(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3373787924(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2418274292(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__2174973445(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta1_1, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(__eta0_0, __local_var_2_0)
})
}

func Call_Data_Functor_map__1152297413(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta1_1, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(__eta0_0, __local_var_2_0)
})
}

func Call_Data_Functor_map__3699108444(f_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(Get_Data_Either_Left(), (*Constructor_Data_Either_Left)(m_1.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(Get_Data_Either_Right(), gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right)(m_1.UnsafePtr).V0))
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

func Call_Data_Functor_map__2579103836(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply4(Get_Data_Bifunctor_bimap(), gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](Get_Data_Interval_bifunctorInterval()))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__1510739772(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
}))
}

func Call_Data_Functor_map__1208755772(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
}))
}

func Call_Data_Functor_map__3565923196(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))))}
}))
}

func Call_Data_Functor_map__2597050044(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))))}
}))
}

func Call_Data_Functor_map__3467322428(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
}))
}

func Call_Data_Functor_map__1958856956(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
}))
}

func Call_Data_Functor_map__1422050556(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
}))
}

func Call_Data_Functor_map__109003388(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1)))})))}
}))
}

func Call_Data_Functor_map__829570556(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1)))})))}
}))
}

func Call_Data_Functor_map__2156385148(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1)))})
}))
}

func Call_Data_Functor_map__558976860(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1)))})
}))
}

func Call_Data_Functor_map__1806510684(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1)))})))}
}))
}

func Call_Data_Functor_map__596534652(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
}))
}

func Call_Data_Functor_map__3815458588(f_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var l_1 gopurs_runtime.Value = l_1_loop
_ = l_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
}))
}

func Call_Data_Functor_map__843173928(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Data_List_Types_listMap(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__438443400(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Data_List_Types_listMap(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__1081746216(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__95558920(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__291265340(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Int(gopurs_runtime.Apply(v_0, gopurs_runtime.Int((v1_1).V0.IntVal)).IntVal)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2107538812(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, gopurs_runtime.Int((v1_1).V0.IntVal))))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__3447677596(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(v_0, gopurs_runtime.Int((v1_1).V0.IntVal))))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2165919164(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, gopurs_runtime.Int((v1_1).V0.IntVal))))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__4155962236(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, gopurs_runtime.Int((v1_1).V0.IntVal))))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2389909756(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, gopurs_runtime.Int((v1_1).V0.IntVal))))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__1759928220(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Int(gopurs_runtime.Apply(v_0, (v1_1).V0).IntVal)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__901270812(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2486200924(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Product](gopurs_runtime.Apply(v_0, (v1_1).V0)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2670646620(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__1887399228(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2565844412(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v_0, (v1_1).V0)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2126568188(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v_0, (v1_1).V0)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__1739124124(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v_0, (v1_1).V0)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__4039429788(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v_0, (v1_1).V0)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__1208952924(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v_0, (v1_1).V0)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2294788636(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__48293596(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__853141532(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2275717084(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__3733923228(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__1171574364(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((v1_1).V0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__63598588(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((v1_1).V0))})))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__1808515292(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__140514012(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_1).V0))})))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__3210082748(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_1).V0))})))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2034458684(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((v1_1).V0))})))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2615158204(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((v1_1).V0))})))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__125648636(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((v1_1).V0))})))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__539092636(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__1980149980(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v_0, (v1_1).V0)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2202537180(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v_0, (v1_1).V0)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__1681779388(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v_0, (v1_1).V0)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__1206962620(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__3486165692(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__2116777468(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v_0, (v1_1).V0)))}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__316107900(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_0, (v1_1).V0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](Get_Data_Maybe_Nothing())
}
end_branch_0:
return __t0
}

func Call_Data_Functor_map__3269387708(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmpty()).V0), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__271334204(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorNonEmpty()).V0), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__91618268(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorNonEmptyList(), "map"), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__2311960860(f_0_loop gopurs_runtime.Value, m_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var m_1 *Constructor_Data_Tuple_Tuple = m_1_loop
_ = m_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(Get_Data_Tuple_Tuple(), (m_1).V0, gopurs_runtime.Apply(f_0, (m_1).V1)))
}

func Call_Data_Functor_map__339096027(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Effect_Aff__map(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__2177087003(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Effect_Aff__map(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__2919116915(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply3(Get_Control_Applicative_liftA1(), Get_Effect_applicativeEffect(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__2261469235(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply3(Get_Control_Applicative_liftA1(), Get_Effect_applicativeEffect(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__3644121587(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply3(Get_Control_Applicative_liftA1(), Get_Effect_applicativeEffect(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__113987891(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply3(Get_Control_Applicative_liftA1(), Get_Effect_applicativeEffect(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__3065908595(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply3(Get_Control_Applicative_liftA1(), Get_Effect_applicativeEffect(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__4102685939(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply3(Get_Control_Applicative_liftA1(), Get_Effect_applicativeEffect(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__173660595(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply3(Get_Control_Applicative_liftA1(), Get_Effect_applicativeEffect(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__1484761587(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply3(Get_Control_Applicative_liftA1(), Get_Effect_applicativeEffect(), __eta0_0, __eta1_1)
}

func Call_Data_Functor_map__1678245779(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_map__3058064980(dict_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_mapFlipped__260821093(dictFunctor_0_loop *Constructor_Data_Functor_Functor, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_2, fa_1)
}

func Call_Data_Functor_mapFlipped__2466386789(dictFunctor_0_loop *Constructor_Data_Functor_Functor, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_2, fa_1)
}

func Call_Data_Functor_mapFlipped__4215217780(dictFunctor_0_loop *Constructor_Data_Functor_Functor, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_2, fa_1)
}

func Call_Data_Functor_mapFlipped__742928244(dictFunctor_0_loop *Constructor_Data_Functor_Functor, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_2, fa_1)
}

func Call_Data_Functor_mapFlipped__509401044(dictFunctor_0_loop *Constructor_Data_Functor_Functor, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_2, fa_1)
}

func Call_Data_Functor_mapFlipped__2213324916(dictFunctor_0_loop *Constructor_Data_Functor_Functor, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_2, fa_1)
}

func Call_Data_Functor_mapFlipped__3249733428(dictFunctor_0_loop *Constructor_Data_Functor_Functor, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_2, fa_1)
}

func Call_Data_Functor_mapFlipped__1087756276(dictFunctor_0_loop *Constructor_Data_Functor_Functor, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_2, fa_1)
}

func Call_Data_Functor_mapFlipped__2919806324(dictFunctor_0_loop *Constructor_Data_Functor_Functor, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), f_2, fa_1)
}

func Call_Data_Functor_void__3020373336(dictFunctor_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
}

func Call_Data_Functor_void__2104786761(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), __eta0_0)
}

func Call_Data_Functor_voidLeft__171362140(f_0_loop gopurs_runtime.Value, x_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 string = x_1_loop
_ = x_1
return func() gopurs_runtime.Value {
arr_val_arrayMap0 := f_0
_ = arr_val_arrayMap0
arr_go_arrayMap0 := (*[]gopurs_runtime.Value)(arr_val_arrayMap0.UnsafePtr)
_ = arr_go_arrayMap0
res_go_arrayMap0 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap0))
_ = res_go_arrayMap0
for i_arrayMap0, v_arrayMap0 := range *arr_go_arrayMap0 {
res_go_arrayMap0[i_arrayMap0] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(x_1)
}), v_arrayMap0)
}
return gopurs_runtime.Array(res_go_arrayMap0)
}()
}

func Call_Data_Functor_voidLeft__32301756(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), f_1)
}

func Call_Data_Functor_voidLeft__3507060164(f_0_loop *Constructor_Data_Maybe_Just, x_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 *Constructor_Data_Maybe_Just = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(f_0)}))
}

func Call_Data_Functor_voidLeft__4152282100(f_0_loop *Constructor_Data_Maybe_Just, x_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 *Constructor_Data_Maybe_Just = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(f_0)}))
}

func Call_Data_Functor_voidRight__3258033404(dictFunctor_0_loop *Constructor_Data_Functor_Functor, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Functor_voidRight__1142845180(dictFunctor_0_loop *Constructor_Data_Functor_Functor, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Functor_voidRight__698766972(dictFunctor_0_loop *Constructor_Data_Functor_Functor, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Functor_voidRight__2767609084(dictFunctor_0_loop *Constructor_Data_Functor_Functor, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Get_Data_Functor_arrayMap() gopurs_runtime.Value {
	return _Gopurs_Data_Functor_ArrayMap
}
