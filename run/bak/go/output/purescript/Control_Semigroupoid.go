package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Semigroupoid_Semigroupoid_dollarDict gopurs_runtime.Value
var once_Control_Semigroupoid_Semigroupoid_dollarDict sync.Once
func Get_Control_Semigroupoid_Semigroupoid_dollarDict() gopurs_runtime.Value {
	once_Control_Semigroupoid_Semigroupoid_dollarDict.Do(func() {
		cache_Control_Semigroupoid_Semigroupoid_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_Semigroupoid_dollarDict(x_0_box)
})
	})
	return cache_Control_Semigroupoid_Semigroupoid_dollarDict
}

var cache_Control_Semigroupoid_semigroupoidFn gopurs_runtime.Value
var once_Control_Semigroupoid_semigroupoidFn sync.Once
func Get_Control_Semigroupoid_semigroupoidFn() gopurs_runtime.Value {
	once_Control_Semigroupoid_semigroupoidFn.Do(func() {
		cache_Control_Semigroupoid_semigroupoidFn = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_Control_Semigroupoid_semigroupoidFn
}

var cache_Control_Semigroupoid_compose gopurs_runtime.Value
var once_Control_Semigroupoid_compose sync.Once
func Get_Control_Semigroupoid_compose() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose.Do(func() {
		cache_Control_Semigroupoid_compose = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose
}

var cache_Control_Semigroupoid_composeFlipped gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped sync.Once
func Get_Control_Semigroupoid_composeFlipped() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped.Do(func() {
		cache_Control_Semigroupoid_composeFlipped = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_Control_Semigroupoid_composeFlipped
}

var cache_Control_Semigroupoid_compose__1987728071 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1987728071 sync.Once
func Get_Control_Semigroupoid_compose__1987728071() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1987728071.Do(func() {
		cache_Control_Semigroupoid_compose__1987728071 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1987728071(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1987728071
}

var cache_Control_Semigroupoid_compose__1636404804 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1636404804 sync.Once
func Get_Control_Semigroupoid_compose__1636404804() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1636404804.Do(func() {
		cache_Control_Semigroupoid_compose__1636404804 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1636404804(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1636404804
}

var cache_Control_Semigroupoid_compose__346034828 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__346034828 sync.Once
func Get_Control_Semigroupoid_compose__346034828() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__346034828.Do(func() {
		cache_Control_Semigroupoid_compose__346034828 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__346034828(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__346034828
}

var cache_Control_Semigroupoid_compose__4141960292 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__4141960292 sync.Once
func Get_Control_Semigroupoid_compose__4141960292() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__4141960292.Do(func() {
		cache_Control_Semigroupoid_compose__4141960292 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__4141960292(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__4141960292
}

var cache_Control_Semigroupoid_compose__1254722180 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1254722180 sync.Once
func Get_Control_Semigroupoid_compose__1254722180() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1254722180.Do(func() {
		cache_Control_Semigroupoid_compose__1254722180 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1254722180(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1254722180
}

var cache_Control_Semigroupoid_compose__1362634137 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1362634137 sync.Once
func Get_Control_Semigroupoid_compose__1362634137() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1362634137.Do(func() {
		cache_Control_Semigroupoid_compose__1362634137 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1362634137(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1362634137
}

var cache_Control_Semigroupoid_compose__3628566261 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3628566261 sync.Once
func Get_Control_Semigroupoid_compose__3628566261() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3628566261.Do(func() {
		cache_Control_Semigroupoid_compose__3628566261 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3628566261(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3628566261
}

var cache_Control_Semigroupoid_compose__706970832 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__706970832 sync.Once
func Get_Control_Semigroupoid_compose__706970832() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__706970832.Do(func() {
		cache_Control_Semigroupoid_compose__706970832 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__706970832(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__706970832
}

var cache_Control_Semigroupoid_compose__3249109794 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3249109794 sync.Once
func Get_Control_Semigroupoid_compose__3249109794() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3249109794.Do(func() {
		cache_Control_Semigroupoid_compose__3249109794 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3249109794(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3249109794
}

var cache_Control_Semigroupoid_compose__1774599291 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1774599291 sync.Once
func Get_Control_Semigroupoid_compose__1774599291() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1774599291.Do(func() {
		cache_Control_Semigroupoid_compose__1774599291 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1774599291(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1774599291
}

var cache_Control_Semigroupoid_compose__1543665403 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1543665403 sync.Once
func Get_Control_Semigroupoid_compose__1543665403() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1543665403.Do(func() {
		cache_Control_Semigroupoid_compose__1543665403 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1543665403(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1543665403
}

var cache_Control_Semigroupoid_compose__2755435211 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2755435211 sync.Once
func Get_Control_Semigroupoid_compose__2755435211() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2755435211.Do(func() {
		cache_Control_Semigroupoid_compose__2755435211 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2755435211(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2755435211
}

var cache_Control_Semigroupoid_compose__618460331 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__618460331 sync.Once
func Get_Control_Semigroupoid_compose__618460331() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__618460331.Do(func() {
		cache_Control_Semigroupoid_compose__618460331 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__618460331(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__618460331
}

var cache_Control_Semigroupoid_compose__1555187646 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1555187646 sync.Once
func Get_Control_Semigroupoid_compose__1555187646() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1555187646.Do(func() {
		cache_Control_Semigroupoid_compose__1555187646 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1555187646(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1555187646
}

var cache_Control_Semigroupoid_compose__2527254334 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2527254334 sync.Once
func Get_Control_Semigroupoid_compose__2527254334() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2527254334.Do(func() {
		cache_Control_Semigroupoid_compose__2527254334 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2527254334(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2527254334
}

var cache_Control_Semigroupoid_compose__564147166 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__564147166 sync.Once
func Get_Control_Semigroupoid_compose__564147166() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__564147166.Do(func() {
		cache_Control_Semigroupoid_compose__564147166 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__564147166(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__564147166
}

var cache_Control_Semigroupoid_compose__2532574046 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2532574046 sync.Once
func Get_Control_Semigroupoid_compose__2532574046() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2532574046.Do(func() {
		cache_Control_Semigroupoid_compose__2532574046 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2532574046(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2532574046
}

var cache_Control_Semigroupoid_compose__794534846 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__794534846 sync.Once
func Get_Control_Semigroupoid_compose__794534846() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__794534846.Do(func() {
		cache_Control_Semigroupoid_compose__794534846 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__794534846(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__794534846
}

var cache_Control_Semigroupoid_compose__1604328382 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1604328382 sync.Once
func Get_Control_Semigroupoid_compose__1604328382() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1604328382.Do(func() {
		cache_Control_Semigroupoid_compose__1604328382 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1604328382(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1604328382
}

var cache_Control_Semigroupoid_compose__4254807102 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__4254807102 sync.Once
func Get_Control_Semigroupoid_compose__4254807102() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__4254807102.Do(func() {
		cache_Control_Semigroupoid_compose__4254807102 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__4254807102(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__4254807102
}

var cache_Control_Semigroupoid_compose__2995688990 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2995688990 sync.Once
func Get_Control_Semigroupoid_compose__2995688990() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2995688990.Do(func() {
		cache_Control_Semigroupoid_compose__2995688990 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2995688990(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2995688990
}

var cache_Control_Semigroupoid_compose__2595061246 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2595061246 sync.Once
func Get_Control_Semigroupoid_compose__2595061246() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2595061246.Do(func() {
		cache_Control_Semigroupoid_compose__2595061246 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2595061246(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2595061246
}

var cache_Control_Semigroupoid_compose__3140790526 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3140790526 sync.Once
func Get_Control_Semigroupoid_compose__3140790526() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3140790526.Do(func() {
		cache_Control_Semigroupoid_compose__3140790526 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3140790526(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3140790526
}

var cache_Control_Semigroupoid_compose__3384557662 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3384557662 sync.Once
func Get_Control_Semigroupoid_compose__3384557662() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3384557662.Do(func() {
		cache_Control_Semigroupoid_compose__3384557662 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3384557662(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3384557662
}

var cache_Control_Semigroupoid_compose__4020612094 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__4020612094 sync.Once
func Get_Control_Semigroupoid_compose__4020612094() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__4020612094.Do(func() {
		cache_Control_Semigroupoid_compose__4020612094 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__4020612094(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__4020612094
}

var cache_Control_Semigroupoid_compose__2710321297 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2710321297 sync.Once
func Get_Control_Semigroupoid_compose__2710321297() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2710321297.Do(func() {
		cache_Control_Semigroupoid_compose__2710321297 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2710321297(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2710321297
}

var cache_Control_Semigroupoid_compose__1933206353 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1933206353 sync.Once
func Get_Control_Semigroupoid_compose__1933206353() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1933206353.Do(func() {
		cache_Control_Semigroupoid_compose__1933206353 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1933206353(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1933206353
}

var cache_Control_Semigroupoid_compose__491672472 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__491672472 sync.Once
func Get_Control_Semigroupoid_compose__491672472() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__491672472.Do(func() {
		cache_Control_Semigroupoid_compose__491672472 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__491672472(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__491672472
}

var cache_Control_Semigroupoid_compose__3311306424 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3311306424 sync.Once
func Get_Control_Semigroupoid_compose__3311306424() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3311306424.Do(func() {
		cache_Control_Semigroupoid_compose__3311306424 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3311306424(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3311306424
}

var cache_Control_Semigroupoid_compose__3249366840 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3249366840 sync.Once
func Get_Control_Semigroupoid_compose__3249366840() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3249366840.Do(func() {
		cache_Control_Semigroupoid_compose__3249366840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3249366840(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3249366840
}

var cache_Control_Semigroupoid_compose__858342840 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__858342840 sync.Once
func Get_Control_Semigroupoid_compose__858342840() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__858342840.Do(func() {
		cache_Control_Semigroupoid_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__858342840(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__858342840
}

var cache_Control_Semigroupoid_compose__187692920 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__187692920 sync.Once
func Get_Control_Semigroupoid_compose__187692920() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__187692920.Do(func() {
		cache_Control_Semigroupoid_compose__187692920 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__187692920(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__187692920
}

var cache_Control_Semigroupoid_compose__2323083928 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2323083928 sync.Once
func Get_Control_Semigroupoid_compose__2323083928() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2323083928.Do(func() {
		cache_Control_Semigroupoid_compose__2323083928 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2323083928(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2323083928
}

var cache_Control_Semigroupoid_compose__3009541624 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3009541624 sync.Once
func Get_Control_Semigroupoid_compose__3009541624() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3009541624.Do(func() {
		cache_Control_Semigroupoid_compose__3009541624 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3009541624(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3009541624
}

var cache_Control_Semigroupoid_compose__2451652760 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2451652760 sync.Once
func Get_Control_Semigroupoid_compose__2451652760() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2451652760.Do(func() {
		cache_Control_Semigroupoid_compose__2451652760 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2451652760(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2451652760
}

var cache_Control_Semigroupoid_compose__1210748600 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1210748600 sync.Once
func Get_Control_Semigroupoid_compose__1210748600() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1210748600.Do(func() {
		cache_Control_Semigroupoid_compose__1210748600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1210748600(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1210748600
}

var cache_Control_Semigroupoid_compose__1871585528 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1871585528 sync.Once
func Get_Control_Semigroupoid_compose__1871585528() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1871585528.Do(func() {
		cache_Control_Semigroupoid_compose__1871585528 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1871585528(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1871585528
}

var cache_Control_Semigroupoid_compose__2679722072 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2679722072 sync.Once
func Get_Control_Semigroupoid_compose__2679722072() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2679722072.Do(func() {
		cache_Control_Semigroupoid_compose__2679722072 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2679722072(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2679722072
}

var cache_Control_Semigroupoid_compose__3219295544 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3219295544 sync.Once
func Get_Control_Semigroupoid_compose__3219295544() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3219295544.Do(func() {
		cache_Control_Semigroupoid_compose__3219295544 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3219295544(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3219295544
}

var cache_Control_Semigroupoid_compose__4205385912 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__4205385912 sync.Once
func Get_Control_Semigroupoid_compose__4205385912() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__4205385912.Do(func() {
		cache_Control_Semigroupoid_compose__4205385912 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__4205385912(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__4205385912
}

var cache_Control_Semigroupoid_compose__3919088440 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3919088440 sync.Once
func Get_Control_Semigroupoid_compose__3919088440() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3919088440.Do(func() {
		cache_Control_Semigroupoid_compose__3919088440 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3919088440(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3919088440
}

var cache_Control_Semigroupoid_compose__2249324024 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2249324024 sync.Once
func Get_Control_Semigroupoid_compose__2249324024() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2249324024.Do(func() {
		cache_Control_Semigroupoid_compose__2249324024 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2249324024(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2249324024
}

var cache_Control_Semigroupoid_compose__2053777400 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2053777400 sync.Once
func Get_Control_Semigroupoid_compose__2053777400() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2053777400.Do(func() {
		cache_Control_Semigroupoid_compose__2053777400 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2053777400(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2053777400
}

var cache_Control_Semigroupoid_compose__2426682552 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2426682552 sync.Once
func Get_Control_Semigroupoid_compose__2426682552() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2426682552.Do(func() {
		cache_Control_Semigroupoid_compose__2426682552 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2426682552(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2426682552
}

var cache_Control_Semigroupoid_compose__3432522072 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3432522072 sync.Once
func Get_Control_Semigroupoid_compose__3432522072() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3432522072.Do(func() {
		cache_Control_Semigroupoid_compose__3432522072 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3432522072(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3432522072
}

var cache_Control_Semigroupoid_compose__1247829080 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1247829080 sync.Once
func Get_Control_Semigroupoid_compose__1247829080() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1247829080.Do(func() {
		cache_Control_Semigroupoid_compose__1247829080 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1247829080(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1247829080
}

var cache_Control_Semigroupoid_compose__1408335838 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1408335838 sync.Once
func Get_Control_Semigroupoid_compose__1408335838() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1408335838.Do(func() {
		cache_Control_Semigroupoid_compose__1408335838 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1408335838(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1408335838
}

var cache_Control_Semigroupoid_compose__2437406270 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2437406270 sync.Once
func Get_Control_Semigroupoid_compose__2437406270() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2437406270.Do(func() {
		cache_Control_Semigroupoid_compose__2437406270 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2437406270(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2437406270
}

var cache_Control_Semigroupoid_compose__3328715582 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3328715582 sync.Once
func Get_Control_Semigroupoid_compose__3328715582() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3328715582.Do(func() {
		cache_Control_Semigroupoid_compose__3328715582 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3328715582(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3328715582
}

var cache_Control_Semigroupoid_compose__1863173246 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1863173246 sync.Once
func Get_Control_Semigroupoid_compose__1863173246() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1863173246.Do(func() {
		cache_Control_Semigroupoid_compose__1863173246 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1863173246(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1863173246
}

var cache_Control_Semigroupoid_compose__1796733246 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1796733246 sync.Once
func Get_Control_Semigroupoid_compose__1796733246() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1796733246.Do(func() {
		cache_Control_Semigroupoid_compose__1796733246 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1796733246(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1796733246
}

var cache_Control_Semigroupoid_compose__2493224702 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2493224702 sync.Once
func Get_Control_Semigroupoid_compose__2493224702() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2493224702.Do(func() {
		cache_Control_Semigroupoid_compose__2493224702 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2493224702(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2493224702
}

var cache_Control_Semigroupoid_compose__3977568382 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3977568382 sync.Once
func Get_Control_Semigroupoid_compose__3977568382() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3977568382.Do(func() {
		cache_Control_Semigroupoid_compose__3977568382 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3977568382(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3977568382
}

var cache_Control_Semigroupoid_compose__1152238654 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1152238654 sync.Once
func Get_Control_Semigroupoid_compose__1152238654() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1152238654.Do(func() {
		cache_Control_Semigroupoid_compose__1152238654 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1152238654(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1152238654
}

var cache_Control_Semigroupoid_compose__1515786046 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1515786046 sync.Once
func Get_Control_Semigroupoid_compose__1515786046() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1515786046.Do(func() {
		cache_Control_Semigroupoid_compose__1515786046 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1515786046(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1515786046
}

var cache_Control_Semigroupoid_compose__379849950 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__379849950 sync.Once
func Get_Control_Semigroupoid_compose__379849950() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__379849950.Do(func() {
		cache_Control_Semigroupoid_compose__379849950 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__379849950(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__379849950
}

var cache_Control_Semigroupoid_compose__167393694 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__167393694 sync.Once
func Get_Control_Semigroupoid_compose__167393694() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__167393694.Do(func() {
		cache_Control_Semigroupoid_compose__167393694 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__167393694(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__167393694
}

var cache_Control_Semigroupoid_compose__3324445208 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__3324445208 sync.Once
func Get_Control_Semigroupoid_compose__3324445208() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__3324445208.Do(func() {
		cache_Control_Semigroupoid_compose__3324445208 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__3324445208(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__3324445208
}

var cache_Control_Semigroupoid_compose__1005376280 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__1005376280 sync.Once
func Get_Control_Semigroupoid_compose__1005376280() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__1005376280.Do(func() {
		cache_Control_Semigroupoid_compose__1005376280 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__1005376280(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__1005376280
}

var cache_Control_Semigroupoid_compose__2298713176 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__2298713176 sync.Once
func Get_Control_Semigroupoid_compose__2298713176() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__2298713176.Do(func() {
		cache_Control_Semigroupoid_compose__2298713176 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__2298713176(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__2298713176
}

var cache_Control_Semigroupoid_compose__299469432 gopurs_runtime.Value
var once_Control_Semigroupoid_compose__299469432 sync.Once
func Get_Control_Semigroupoid_compose__299469432() gopurs_runtime.Value {
	once_Control_Semigroupoid_compose__299469432.Do(func() {
		cache_Control_Semigroupoid_compose__299469432 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_compose__299469432(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dict_0_box))
})
	})
	return cache_Control_Semigroupoid_compose__299469432
}

var cache_Control_Semigroupoid_composeFlipped__2583068543 gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped__2583068543 sync.Once
func Get_Control_Semigroupoid_composeFlipped__2583068543() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped__2583068543.Do(func() {
		cache_Control_Semigroupoid_composeFlipped__2583068543 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_composeFlipped__2583068543(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_Control_Semigroupoid_composeFlipped__2583068543
}

var cache_Control_Semigroupoid_composeFlipped__4057377183 gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped__4057377183 sync.Once
func Get_Control_Semigroupoid_composeFlipped__4057377183() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped__4057377183.Do(func() {
		cache_Control_Semigroupoid_composeFlipped__4057377183 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_composeFlipped__4057377183(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_Control_Semigroupoid_composeFlipped__4057377183
}

var cache_Control_Semigroupoid_composeFlipped__1872684191 gopurs_runtime.Value
var once_Control_Semigroupoid_composeFlipped__1872684191 sync.Once
func Get_Control_Semigroupoid_composeFlipped__1872684191() gopurs_runtime.Value {
	once_Control_Semigroupoid_composeFlipped__1872684191.Do(func() {
		cache_Control_Semigroupoid_composeFlipped__1872684191 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Semigroupoid_composeFlipped__1872684191(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_Control_Semigroupoid_composeFlipped__1872684191
}

var cache_Control_Semigroupoid_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_Control_Semigroupoid_semigroupoidFn__2387483462 sync.Once
func Get_Control_Semigroupoid_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_Control_Semigroupoid_semigroupoidFn__2387483462.Do(func() {
		cache_Control_Semigroupoid_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_Control_Semigroupoid_semigroupoidFn__2387483462
}

var cache_Control_Semigroupoid_semigroupoidFn__3002128382 gopurs_runtime.Value
var once_Control_Semigroupoid_semigroupoidFn__3002128382 sync.Once
func Get_Control_Semigroupoid_semigroupoidFn__3002128382() gopurs_runtime.Value {
	once_Control_Semigroupoid_semigroupoidFn__3002128382.Do(func() {
		cache_Control_Semigroupoid_semigroupoidFn__3002128382 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_Control_Semigroupoid_semigroupoidFn__3002128382
}

type Constructor_Control_Semigroupoid_Semigroupoid struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[350442445] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Semigroupoid_Semigroupoid)(ptr)
		_ = c
		switch key {
		case "compose": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Control_Semigroupoid_Semigroupoid: " + key)
		}
	}
}


func Call_Control_Semigroupoid_Semigroupoid_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Semigroupoid_compose(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_composeFlipped(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroupoid_0.V0), g_2, f_1)
}

func Call_Control_Semigroupoid_compose__1987728071(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1636404804(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__346034828(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__4141960292(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1254722180(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1362634137(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3628566261(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__706970832(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3249109794(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1774599291(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1543665403(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2755435211(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__618460331(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1555187646(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2527254334(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__564147166(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2532574046(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__794534846(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1604328382(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__4254807102(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2995688990(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2595061246(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3140790526(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3384557662(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__4020612094(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2710321297(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1933206353(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__491672472(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3311306424(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3249366840(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__858342840(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__187692920(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2323083928(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3009541624(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2451652760(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1210748600(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1871585528(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2679722072(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3219295544(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__4205385912(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3919088440(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2249324024(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2053777400(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2426682552(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3432522072(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1247829080(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1408335838(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2437406270(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3328715582(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1863173246(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1796733246(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2493224702(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3977568382(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1152238654(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1515786046(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__379849950(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__167393694(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__3324445208(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__1005376280(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__2298713176(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_compose__299469432(dict_0_loop *Constructor_Control_Semigroupoid_Semigroupoid) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Semigroupoid_Semigroupoid = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Semigroupoid_composeFlipped__2583068543(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroupoid_0.V0), g_2, f_1)
}

func Call_Control_Semigroupoid_composeFlipped__4057377183(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroupoid_0.V0), g_2, f_1)
}

func Call_Control_Semigroupoid_composeFlipped__1872684191(dictSemigroupoid_0_loop *Constructor_Control_Semigroupoid_Semigroupoid, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *Constructor_Control_Semigroupoid_Semigroupoid = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroupoid_0.V0), g_2, f_1)
}


