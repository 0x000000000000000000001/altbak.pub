package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Function_on gopurs_runtime.Value
var once_Data_Function_on sync.Once
func Get_Data_Function_on() gopurs_runtime.Value {
	once_Data_Function_on.Do(func() {
		cache_Data_Function_on = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_on(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_Data_Function_on
}

var cache_Data_Function_flip gopurs_runtime.Value
var once_Data_Function_flip sync.Once
func Get_Data_Function_flip() gopurs_runtime.Value {
	once_Data_Function_flip.Do(func() {
		cache_Data_Function_flip = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip
}

var cache_Data_Function_go__const gopurs_runtime.Value
var once_Data_Function_go__const sync.Once
func Get_Data_Function_go__const() gopurs_runtime.Value {
	once_Data_Function_go__const.Do(func() {
		cache_Data_Function_go__const = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_go__const(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_go__const
}

var cache_Data_Function_applyN gopurs_runtime.Value
var once_Data_Function_applyN sync.Once
func Get_Data_Function_applyN() gopurs_runtime.Value {
	once_Data_Function_applyN.Do(func() {
		cache_Data_Function_applyN = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_applyN(f_0_box)
})
	})
	return cache_Data_Function_applyN
}

var cache_Data_Function_applyFlipped gopurs_runtime.Value
var once_Data_Function_applyFlipped sync.Once
func Get_Data_Function_applyFlipped() gopurs_runtime.Value {
	once_Data_Function_applyFlipped.Do(func() {
		cache_Data_Function_applyFlipped = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_applyFlipped(x_0_box, f_1_box)
})
	})
	return cache_Data_Function_applyFlipped
}

var cache_Data_Function_apply gopurs_runtime.Value
var once_Data_Function_apply sync.Once
func Get_Data_Function_apply() gopurs_runtime.Value {
	once_Data_Function_apply.Do(func() {
		cache_Data_Function_apply = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply(f_0_box, x_1_box)
})
	})
	return cache_Data_Function_apply
}

var cache_Data_Function_apply__458711162 gopurs_runtime.Value
var once_Data_Function_apply__458711162 sync.Once
func Get_Data_Function_apply__458711162() gopurs_runtime.Value {
	once_Data_Function_apply__458711162.Do(func() {
		cache_Data_Function_apply__458711162 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_apply__458711162(f_0_box, x_1_box)
})
	})
	return cache_Data_Function_apply__458711162
}

var cache_Data_Function_const__3524684546 gopurs_runtime.Value
var once_Data_Function_const__3524684546 sync.Once
func Get_Data_Function_const__3524684546() gopurs_runtime.Value {
	once_Data_Function_const__3524684546.Do(func() {
		cache_Data_Function_const__3524684546 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__3524684546(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__3524684546
}

var cache_Data_Function_const__1496134642 gopurs_runtime.Value
var once_Data_Function_const__1496134642 sync.Once
func Get_Data_Function_const__1496134642() gopurs_runtime.Value {
	once_Data_Function_const__1496134642.Do(func() {
		cache_Data_Function_const__1496134642 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__1496134642(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__1496134642
}

var cache_Data_Function_const__1832354163 gopurs_runtime.Value
var once_Data_Function_const__1832354163 sync.Once
func Get_Data_Function_const__1832354163() gopurs_runtime.Value {
	once_Data_Function_const__1832354163.Do(func() {
		cache_Data_Function_const__1832354163 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__1832354163(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__1832354163
}

var cache_Data_Function_const__1426827922 gopurs_runtime.Value
var once_Data_Function_const__1426827922 sync.Once
func Get_Data_Function_const__1426827922() gopurs_runtime.Value {
	once_Data_Function_const__1426827922.Do(func() {
		cache_Data_Function_const__1426827922 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__1426827922(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__1426827922
}

var cache_Data_Function_const__2569866098 gopurs_runtime.Value
var once_Data_Function_const__2569866098 sync.Once
func Get_Data_Function_const__2569866098() gopurs_runtime.Value {
	once_Data_Function_const__2569866098.Do(func() {
		cache_Data_Function_const__2569866098 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__2569866098(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__2569866098
}

var cache_Data_Function_const__4181451586 gopurs_runtime.Value
var once_Data_Function_const__4181451586 sync.Once
func Get_Data_Function_const__4181451586() gopurs_runtime.Value {
	once_Data_Function_const__4181451586.Do(func() {
		cache_Data_Function_const__4181451586 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__4181451586(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__4181451586
}

var cache_Data_Function_const__2390202835 gopurs_runtime.Value
var once_Data_Function_const__2390202835 sync.Once
func Get_Data_Function_const__2390202835() gopurs_runtime.Value {
	once_Data_Function_const__2390202835.Do(func() {
		cache_Data_Function_const__2390202835 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_const__2390202835((a_0_box.IntVal) != (0), v_1_box))
})
	})
	return cache_Data_Function_const__2390202835
}

var cache_Data_Function_const__702735379 gopurs_runtime.Value
var once_Data_Function_const__702735379 sync.Once
func Get_Data_Function_const__702735379() gopurs_runtime.Value {
	once_Data_Function_const__702735379.Do(func() {
		cache_Data_Function_const__702735379 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__702735379(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_0_box), v_1_box))}
})
	})
	return cache_Data_Function_const__702735379
}

var cache_Data_Function_const__3834287059 gopurs_runtime.Value
var once_Data_Function_const__3834287059 sync.Once
func Get_Data_Function_const__3834287059() gopurs_runtime.Value {
	once_Data_Function_const__3834287059.Do(func() {
		cache_Data_Function_const__3834287059 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__3834287059(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_1_box)))}
})
	})
	return cache_Data_Function_const__3834287059
}

var cache_Data_Function_const__1243414737 gopurs_runtime.Value
var once_Data_Function_const__1243414737 sync.Once
func Get_Data_Function_const__1243414737() gopurs_runtime.Value {
	once_Data_Function_const__1243414737.Do(func() {
		cache_Data_Function_const__1243414737 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Function_const__1243414737(a_0_box.IntVal, v_1_box.IntVal))
})
	})
	return cache_Data_Function_const__1243414737
}

var cache_Data_Function_const__3567418756 gopurs_runtime.Value
var once_Data_Function_const__3567418756 sync.Once
func Get_Data_Function_const__3567418756() gopurs_runtime.Value {
	once_Data_Function_const__3567418756.Do(func() {
		cache_Data_Function_const__3567418756 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Function_const__3567418756(a_0_box.IntVal, v_1_box))
})
	})
	return cache_Data_Function_const__3567418756
}

var cache_Data_Function_const__2082174484 gopurs_runtime.Value
var once_Data_Function_const__2082174484 sync.Once
func Get_Data_Function_const__2082174484() gopurs_runtime.Value {
	once_Data_Function_const__2082174484.Do(func() {
		cache_Data_Function_const__2082174484 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Function_const__2082174484(a_0_box.IntVal, v_1_box))
})
	})
	return cache_Data_Function_const__2082174484
}

var cache_Data_Function_const__2632676769 gopurs_runtime.Value
var once_Data_Function_const__2632676769 sync.Once
func Get_Data_Function_const__2632676769() gopurs_runtime.Value {
	once_Data_Function_const__2632676769.Do(func() {
		cache_Data_Function_const__2632676769 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Function_const__2632676769(a_0_box.FloatVal(), v_1_box.FloatVal()))
})
	})
	return cache_Data_Function_const__2632676769
}

var cache_Data_Function_const__2696088948 gopurs_runtime.Value
var once_Data_Function_const__2696088948 sync.Once
func Get_Data_Function_const__2696088948() gopurs_runtime.Value {
	once_Data_Function_const__2696088948.Do(func() {
		cache_Data_Function_const__2696088948 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Function_const__2696088948(a_0_box.FloatVal(), v_1_box))
})
	})
	return cache_Data_Function_const__2696088948
}

var cache_Data_Function_const__4157258135 gopurs_runtime.Value
var once_Data_Function_const__4157258135 sync.Once
func Get_Data_Function_const__4157258135() gopurs_runtime.Value {
	once_Data_Function_const__4157258135.Do(func() {
		cache_Data_Function_const__4157258135 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Function_const__4157258135(a_0_box.StrVal(), v_1_box.StrVal()))
})
	})
	return cache_Data_Function_const__4157258135
}

var cache_Data_Function_const__1562253172 gopurs_runtime.Value
var once_Data_Function_const__1562253172 sync.Once
func Get_Data_Function_const__1562253172() gopurs_runtime.Value {
	once_Data_Function_const__1562253172.Do(func() {
		cache_Data_Function_const__1562253172 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Function_const__1562253172(a_0_box.StrVal(), v_1_box))
})
	})
	return cache_Data_Function_const__1562253172
}

var cache_Data_Function_const__220790420 gopurs_runtime.Value
var once_Data_Function_const__220790420 sync.Once
func Get_Data_Function_const__220790420() gopurs_runtime.Value {
	once_Data_Function_const__220790420.Do(func() {
		cache_Data_Function_const__220790420 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_const__220790420((a_0_box.IntVal) != (0), v_1_box))
})
	})
	return cache_Data_Function_const__220790420
}

var cache_Data_Function_const__4026847508 gopurs_runtime.Value
var once_Data_Function_const__4026847508 sync.Once
func Get_Data_Function_const__4026847508() gopurs_runtime.Value {
	once_Data_Function_const__4026847508.Do(func() {
		cache_Data_Function_const__4026847508 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__4026847508(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__4026847508
}

var cache_Data_Function_const__2857921436 gopurs_runtime.Value
var once_Data_Function_const__2857921436 sync.Once
func Get_Data_Function_const__2857921436() gopurs_runtime.Value {
	once_Data_Function_const__2857921436.Do(func() {
		cache_Data_Function_const__2857921436 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__2857921436(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__2857921436
}

var cache_Data_Function_const__744888567 gopurs_runtime.Value
var once_Data_Function_const__744888567 sync.Once
func Get_Data_Function_const__744888567() gopurs_runtime.Value {
	once_Data_Function_const__744888567.Do(func() {
		cache_Data_Function_const__744888567 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__744888567(a_0_box, v_1_box.StrVal())
})
	})
	return cache_Data_Function_const__744888567
}

var cache_Data_Function_const__2050378404 gopurs_runtime.Value
var once_Data_Function_const__2050378404 sync.Once
func Get_Data_Function_const__2050378404() gopurs_runtime.Value {
	once_Data_Function_const__2050378404.Do(func() {
		cache_Data_Function_const__2050378404 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__2050378404(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__2050378404
}

var cache_Data_Function_const__641934996 gopurs_runtime.Value
var once_Data_Function_const__641934996 sync.Once
func Get_Data_Function_const__641934996() gopurs_runtime.Value {
	once_Data_Function_const__641934996.Do(func() {
		cache_Data_Function_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__641934996
}

var cache_Data_Function_const__2291477660 gopurs_runtime.Value
var once_Data_Function_const__2291477660 sync.Once
func Get_Data_Function_const__2291477660() gopurs_runtime.Value {
	once_Data_Function_const__2291477660.Do(func() {
		cache_Data_Function_const__2291477660 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__2291477660(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__2291477660
}

var cache_Data_Function_const__72843521 gopurs_runtime.Value
var once_Data_Function_const__72843521 sync.Once
func Get_Data_Function_const__72843521() gopurs_runtime.Value {
	once_Data_Function_const__72843521.Do(func() {
		cache_Data_Function_const__72843521 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__72843521(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__72843521
}

var cache_Data_Function_const__3848686068 gopurs_runtime.Value
var once_Data_Function_const__3848686068 sync.Once
func Get_Data_Function_const__3848686068() gopurs_runtime.Value {
	once_Data_Function_const__3848686068.Do(func() {
		cache_Data_Function_const__3848686068 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__3848686068(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__3848686068
}

var cache_Data_Function_const__992538388 gopurs_runtime.Value
var once_Data_Function_const__992538388 sync.Once
func Get_Data_Function_const__992538388() gopurs_runtime.Value {
	once_Data_Function_const__992538388.Do(func() {
		cache_Data_Function_const__992538388 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__992538388(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__992538388
}

var cache_Data_Function_const__1803609684 gopurs_runtime.Value
var once_Data_Function_const__1803609684 sync.Once
func Get_Data_Function_const__1803609684() gopurs_runtime.Value {
	once_Data_Function_const__1803609684.Do(func() {
		cache_Data_Function_const__1803609684 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Function_const__1803609684(uint32(a_0_box.IntVal), v_1_box)), UnsafePtr: nil}
})
	})
	return cache_Data_Function_const__1803609684
}

var cache_Data_Function_const__2568740136 gopurs_runtime.Value
var once_Data_Function_const__2568740136 sync.Once
func Get_Data_Function_const__2568740136() gopurs_runtime.Value {
	once_Data_Function_const__2568740136.Do(func() {
		cache_Data_Function_const__2568740136 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Function_const__2568740136(uint32(a_0_box.IntVal), uint32(v_1_box.IntVal))), UnsafePtr: nil}
})
	})
	return cache_Data_Function_const__2568740136
}

var cache_Data_Function_const__1776115220 gopurs_runtime.Value
var once_Data_Function_const__1776115220 sync.Once
func Get_Data_Function_const__1776115220() gopurs_runtime.Value {
	once_Data_Function_const__1776115220.Do(func() {
		cache_Data_Function_const__1776115220 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__1776115220(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a_0_box), v_1_box))}
})
	})
	return cache_Data_Function_const__1776115220
}

var cache_Data_Function_const__1863643516 gopurs_runtime.Value
var once_Data_Function_const__1863643516 sync.Once
func Get_Data_Function_const__1863643516() gopurs_runtime.Value {
	once_Data_Function_const__1863643516.Do(func() {
		cache_Data_Function_const__1863643516 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__1863643516(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_1_box)))}
})
	})
	return cache_Data_Function_const__1863643516
}

var cache_Data_Function_const__1245418657 gopurs_runtime.Value
var once_Data_Function_const__1245418657 sync.Once
func Get_Data_Function_const__1245418657() gopurs_runtime.Value {
	once_Data_Function_const__1245418657.Do(func() {
		cache_Data_Function_const__1245418657 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__1245418657(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__1245418657
}

var cache_Data_Function_const__1628252324 gopurs_runtime.Value
var once_Data_Function_const__1628252324 sync.Once
func Get_Data_Function_const__1628252324() gopurs_runtime.Value {
	once_Data_Function_const__1628252324.Do(func() {
		cache_Data_Function_const__1628252324 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 953589075, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__1628252324(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](a_0_box), v_1_box))}
})
	})
	return cache_Data_Function_const__1628252324
}

var cache_Data_Function_const__3470624004 gopurs_runtime.Value
var once_Data_Function_const__3470624004 sync.Once
func Get_Data_Function_const__3470624004() gopurs_runtime.Value {
	once_Data_Function_const__3470624004.Do(func() {
		cache_Data_Function_const__3470624004 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__3470624004(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_0_box), v_1_box))}
})
	})
	return cache_Data_Function_const__3470624004
}

var cache_Data_Function_const__587810660 gopurs_runtime.Value
var once_Data_Function_const__587810660 sync.Once
func Get_Data_Function_const__587810660() gopurs_runtime.Value {
	once_Data_Function_const__587810660.Do(func() {
		cache_Data_Function_const__587810660 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__587810660(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_0_box), v_1_box))}
})
	})
	return cache_Data_Function_const__587810660
}

var cache_Data_Function_const__3525124788 gopurs_runtime.Value
var once_Data_Function_const__3525124788 sync.Once
func Get_Data_Function_const__3525124788() gopurs_runtime.Value {
	once_Data_Function_const__3525124788.Do(func() {
		cache_Data_Function_const__3525124788 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__3525124788(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_0_box), v_1_box))}
})
	})
	return cache_Data_Function_const__3525124788
}

var cache_Data_Function_const__153462564 gopurs_runtime.Value
var once_Data_Function_const__153462564 sync.Once
func Get_Data_Function_const__153462564() gopurs_runtime.Value {
	once_Data_Function_const__153462564.Do(func() {
		cache_Data_Function_const__153462564 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__153462564(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_0_box), v_1_box))}
})
	})
	return cache_Data_Function_const__153462564
}

var cache_Data_Function_const__839298276 gopurs_runtime.Value
var once_Data_Function_const__839298276 sync.Once
func Get_Data_Function_const__839298276() gopurs_runtime.Value {
	once_Data_Function_const__839298276.Do(func() {
		cache_Data_Function_const__839298276 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__839298276(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_0_box), v_1_box))}
})
	})
	return cache_Data_Function_const__839298276
}

var cache_Data_Function_const__3168103323 gopurs_runtime.Value
var once_Data_Function_const__3168103323 sync.Once
func Get_Data_Function_const__3168103323() gopurs_runtime.Value {
	once_Data_Function_const__3168103323.Do(func() {
		cache_Data_Function_const__3168103323 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__3168103323(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_0_box), v_1_box))}
})
	})
	return cache_Data_Function_const__3168103323
}

var cache_Data_Function_const__106415396 gopurs_runtime.Value
var once_Data_Function_const__106415396 sync.Once
func Get_Data_Function_const__106415396() gopurs_runtime.Value {
	once_Data_Function_const__106415396.Do(func() {
		cache_Data_Function_const__106415396 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_const__106415396(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_0_box), v_1_box))}
})
	})
	return cache_Data_Function_const__106415396
}

var cache_Data_Function_const__1155968100 gopurs_runtime.Value
var once_Data_Function_const__1155968100 sync.Once
func Get_Data_Function_const__1155968100() gopurs_runtime.Value {
	once_Data_Function_const__1155968100.Do(func() {
		cache_Data_Function_const__1155968100 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__1155968100(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__1155968100
}

var cache_Data_Function_const__73052052 gopurs_runtime.Value
var once_Data_Function_const__73052052 sync.Once
func Get_Data_Function_const__73052052() gopurs_runtime.Value {
	once_Data_Function_const__73052052.Do(func() {
		cache_Data_Function_const__73052052 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__73052052(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__73052052
}

var cache_Data_Function_const__3415939124 gopurs_runtime.Value
var once_Data_Function_const__3415939124 sync.Once
func Get_Data_Function_const__3415939124() gopurs_runtime.Value {
	once_Data_Function_const__3415939124.Do(func() {
		cache_Data_Function_const__3415939124 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__3415939124(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__3415939124
}

var cache_Data_Function_const__2189647754 gopurs_runtime.Value
var once_Data_Function_const__2189647754 sync.Once
func Get_Data_Function_const__2189647754() gopurs_runtime.Value {
	once_Data_Function_const__2189647754.Do(func() {
		cache_Data_Function_const__2189647754 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__2189647754(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__2189647754
}

var cache_Data_Function_const__4270360676 gopurs_runtime.Value
var once_Data_Function_const__4270360676 sync.Once
func Get_Data_Function_const__4270360676() gopurs_runtime.Value {
	once_Data_Function_const__4270360676.Do(func() {
		cache_Data_Function_const__4270360676 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__4270360676(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__4270360676
}

var cache_Data_Function_const__4189285076 gopurs_runtime.Value
var once_Data_Function_const__4189285076 sync.Once
func Get_Data_Function_const__4189285076() gopurs_runtime.Value {
	once_Data_Function_const__4189285076.Do(func() {
		cache_Data_Function_const__4189285076 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__4189285076(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__4189285076
}

var cache_Data_Function_const__3953240484 gopurs_runtime.Value
var once_Data_Function_const__3953240484 sync.Once
func Get_Data_Function_const__3953240484() gopurs_runtime.Value {
	once_Data_Function_const__3953240484.Do(func() {
		cache_Data_Function_const__3953240484 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__3953240484(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__3953240484
}

var cache_Data_Function_const__2557237620 gopurs_runtime.Value
var once_Data_Function_const__2557237620 sync.Once
func Get_Data_Function_const__2557237620() gopurs_runtime.Value {
	once_Data_Function_const__2557237620.Do(func() {
		cache_Data_Function_const__2557237620 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__2557237620(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__2557237620
}

var cache_Data_Function_const__3858583060 gopurs_runtime.Value
var once_Data_Function_const__3858583060 sync.Once
func Get_Data_Function_const__3858583060() gopurs_runtime.Value {
	once_Data_Function_const__3858583060.Do(func() {
		cache_Data_Function_const__3858583060 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__3858583060(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__3858583060
}

var cache_Data_Function_const__3952683620 gopurs_runtime.Value
var once_Data_Function_const__3952683620 sync.Once
func Get_Data_Function_const__3952683620() gopurs_runtime.Value {
	once_Data_Function_const__3952683620.Do(func() {
		cache_Data_Function_const__3952683620 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__3952683620(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__3952683620
}

var cache_Data_Function_const__754826900 gopurs_runtime.Value
var once_Data_Function_const__754826900 sync.Once
func Get_Data_Function_const__754826900() gopurs_runtime.Value {
	once_Data_Function_const__754826900.Do(func() {
		cache_Data_Function_const__754826900 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__754826900(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__754826900
}

var cache_Data_Function_const__4115900574 gopurs_runtime.Value
var once_Data_Function_const__4115900574 sync.Once
func Get_Data_Function_const__4115900574() gopurs_runtime.Value {
	once_Data_Function_const__4115900574.Do(func() {
		cache_Data_Function_const__4115900574 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_const__4115900574(a_0_box, v_1_box)
})
	})
	return cache_Data_Function_const__4115900574
}

var cache_Data_Function_flip__2673533882 gopurs_runtime.Value
var once_Data_Function_flip__2673533882 sync.Once
func Get_Data_Function_flip__2673533882() gopurs_runtime.Value {
	once_Data_Function_flip__2673533882.Do(func() {
		cache_Data_Function_flip__2673533882 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__2673533882(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__2673533882
}

var cache_Data_Function_flip__1826582752 gopurs_runtime.Value
var once_Data_Function_flip__1826582752 sync.Once
func Get_Data_Function_flip__1826582752() gopurs_runtime.Value {
	once_Data_Function_flip__1826582752.Do(func() {
		cache_Data_Function_flip__1826582752 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__1826582752(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](b_1_box), a_2_box.IntVal))}
})
	})
	return cache_Data_Function_flip__1826582752
}

var cache_Data_Function_flip__2176113088 gopurs_runtime.Value
var once_Data_Function_flip__2176113088 sync.Once
func Get_Data_Function_flip__2176113088() gopurs_runtime.Value {
	once_Data_Function_flip__2176113088.Do(func() {
		cache_Data_Function_flip__2176113088 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__2176113088(f_0_box, b_1_box.StrVal(), a_2_box.StrVal())
})
	})
	return cache_Data_Function_flip__2176113088
}

var cache_Data_Function_flip__4032303232 gopurs_runtime.Value
var once_Data_Function_flip__4032303232 sync.Once
func Get_Data_Function_flip__4032303232() gopurs_runtime.Value {
	once_Data_Function_flip__4032303232.Do(func() {
		cache_Data_Function_flip__4032303232 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__4032303232(f_0_box, b_1_box.StrVal(), a_2_box.StrVal()))}
})
	})
	return cache_Data_Function_flip__4032303232
}

var cache_Data_Function_flip__1384624704 gopurs_runtime.Value
var once_Data_Function_flip__1384624704 sync.Once
func Get_Data_Function_flip__1384624704() gopurs_runtime.Value {
	once_Data_Function_flip__1384624704.Do(func() {
		cache_Data_Function_flip__1384624704 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1384624704(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__1384624704
}

var cache_Data_Function_flip__4257609536 gopurs_runtime.Value
var once_Data_Function_flip__4257609536 sync.Once
func Get_Data_Function_flip__4257609536() gopurs_runtime.Value {
	once_Data_Function_flip__4257609536.Do(func() {
		cache_Data_Function_flip__4257609536 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__4257609536(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__4257609536
}

var cache_Data_Function_flip__1760577696 gopurs_runtime.Value
var once_Data_Function_flip__1760577696 sync.Once
func Get_Data_Function_flip__1760577696() gopurs_runtime.Value {
	once_Data_Function_flip__1760577696.Do(func() {
		cache_Data_Function_flip__1760577696 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1760577696(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__1760577696
}

var cache_Data_Function_flip__3525081280 gopurs_runtime.Value
var once_Data_Function_flip__3525081280 sync.Once
func Get_Data_Function_flip__3525081280() gopurs_runtime.Value {
	once_Data_Function_flip__3525081280.Do(func() {
		cache_Data_Function_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3525081280
}

var cache_Data_Function_flip__2226752640 gopurs_runtime.Value
var once_Data_Function_flip__2226752640 sync.Once
func Get_Data_Function_flip__2226752640() gopurs_runtime.Value {
	once_Data_Function_flip__2226752640.Do(func() {
		cache_Data_Function_flip__2226752640 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Function_flip__2226752640(f_0_box, b_1_box, a_2_box))
})
	})
	return cache_Data_Function_flip__2226752640
}

var cache_Data_Function_flip__3110498784 gopurs_runtime.Value
var once_Data_Function_flip__3110498784 sync.Once
func Get_Data_Function_flip__3110498784() gopurs_runtime.Value {
	once_Data_Function_flip__3110498784.Do(func() {
		cache_Data_Function_flip__3110498784 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3110498784(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3110498784
}

var cache_Data_Function_flip__1540101856 gopurs_runtime.Value
var once_Data_Function_flip__1540101856 sync.Once
func Get_Data_Function_flip__1540101856() gopurs_runtime.Value {
	once_Data_Function_flip__1540101856.Do(func() {
		cache_Data_Function_flip__1540101856 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__1540101856(f_0_box, b_1_box, a_2_box))}
})
	})
	return cache_Data_Function_flip__1540101856
}

var cache_Data_Function_flip__3658931456 gopurs_runtime.Value
var once_Data_Function_flip__3658931456 sync.Once
func Get_Data_Function_flip__3658931456() gopurs_runtime.Value {
	once_Data_Function_flip__3658931456.Do(func() {
		cache_Data_Function_flip__3658931456 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3658931456(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3658931456
}

var cache_Data_Function_flip__2974723072 gopurs_runtime.Value
var once_Data_Function_flip__2974723072 sync.Once
func Get_Data_Function_flip__2974723072() gopurs_runtime.Value {
	once_Data_Function_flip__2974723072.Do(func() {
		cache_Data_Function_flip__2974723072 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__2974723072(f_0_box, b_1_box, a_2_box))}
})
	})
	return cache_Data_Function_flip__2974723072
}

var cache_Data_Function_flip__73563104 gopurs_runtime.Value
var once_Data_Function_flip__73563104 sync.Once
func Get_Data_Function_flip__73563104() gopurs_runtime.Value {
	once_Data_Function_flip__73563104.Do(func() {
		cache_Data_Function_flip__73563104 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__73563104(f_0_box, b_1_box, a_2_box))}
})
	})
	return cache_Data_Function_flip__73563104
}

var cache_Data_Function_flip__787327104 gopurs_runtime.Value
var once_Data_Function_flip__787327104 sync.Once
func Get_Data_Function_flip__787327104() gopurs_runtime.Value {
	once_Data_Function_flip__787327104.Do(func() {
		cache_Data_Function_flip__787327104 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__787327104(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__787327104
}

var cache_Data_Function_flip__1117087808 gopurs_runtime.Value
var once_Data_Function_flip__1117087808 sync.Once
func Get_Data_Function_flip__1117087808() gopurs_runtime.Value {
	once_Data_Function_flip__1117087808.Do(func() {
		cache_Data_Function_flip__1117087808 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1117087808(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__1117087808
}

var cache_Data_Function_flip__3021024128 gopurs_runtime.Value
var once_Data_Function_flip__3021024128 sync.Once
func Get_Data_Function_flip__3021024128() gopurs_runtime.Value {
	once_Data_Function_flip__3021024128.Do(func() {
		cache_Data_Function_flip__3021024128 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Function_flip__3021024128(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(b_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), a_2_box))
})
	})
	return cache_Data_Function_flip__3021024128
}

var cache_Data_Function_flip__3904800000 gopurs_runtime.Value
var once_Data_Function_flip__3904800000 sync.Once
func Get_Data_Function_flip__3904800000() gopurs_runtime.Value {
	once_Data_Function_flip__3904800000.Do(func() {
		cache_Data_Function_flip__3904800000 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3904800000(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3904800000
}

var cache_Data_Function_flip__2928880736 gopurs_runtime.Value
var once_Data_Function_flip__2928880736 sync.Once
func Get_Data_Function_flip__2928880736() gopurs_runtime.Value {
	once_Data_Function_flip__2928880736.Do(func() {
		cache_Data_Function_flip__2928880736 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__2928880736(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__2928880736
}

var cache_Data_Function_flip__3019719488 gopurs_runtime.Value
var once_Data_Function_flip__3019719488 sync.Once
func Get_Data_Function_flip__3019719488() gopurs_runtime.Value {
	once_Data_Function_flip__3019719488.Do(func() {
		cache_Data_Function_flip__3019719488 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3019719488(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](b_1_box), a_2_box)
})
	})
	return cache_Data_Function_flip__3019719488
}

var cache_Data_Function_flip__2733110176 gopurs_runtime.Value
var once_Data_Function_flip__2733110176 sync.Once
func Get_Data_Function_flip__2733110176() gopurs_runtime.Value {
	once_Data_Function_flip__2733110176.Do(func() {
		cache_Data_Function_flip__2733110176 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__2733110176(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](b_1_box), a_2_box))}
})
	})
	return cache_Data_Function_flip__2733110176
}

var cache_Data_Function_flip__3019832928 gopurs_runtime.Value
var once_Data_Function_flip__3019832928 sync.Once
func Get_Data_Function_flip__3019832928() gopurs_runtime.Value {
	once_Data_Function_flip__3019832928.Do(func() {
		cache_Data_Function_flip__3019832928 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3019832928(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3019832928
}

var cache_Data_Function_flip__3675729664 gopurs_runtime.Value
var once_Data_Function_flip__3675729664 sync.Once
func Get_Data_Function_flip__3675729664() gopurs_runtime.Value {
	once_Data_Function_flip__3675729664.Do(func() {
		cache_Data_Function_flip__3675729664 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3675729664(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3675729664
}

var cache_Data_Function_flip__3709724320 gopurs_runtime.Value
var once_Data_Function_flip__3709724320 sync.Once
func Get_Data_Function_flip__3709724320() gopurs_runtime.Value {
	once_Data_Function_flip__3709724320.Do(func() {
		cache_Data_Function_flip__3709724320 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3709724320(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_1_box), a_2_box)
})
	})
	return cache_Data_Function_flip__3709724320
}

var cache_Data_Function_flip__3563101792 gopurs_runtime.Value
var once_Data_Function_flip__3563101792 sync.Once
func Get_Data_Function_flip__3563101792() gopurs_runtime.Value {
	once_Data_Function_flip__3563101792.Do(func() {
		cache_Data_Function_flip__3563101792 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__3563101792(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_1_box), a_2_box))}
})
	})
	return cache_Data_Function_flip__3563101792
}

var cache_Data_Function_flip__4284296032 gopurs_runtime.Value
var once_Data_Function_flip__4284296032 sync.Once
func Get_Data_Function_flip__4284296032() gopurs_runtime.Value {
	once_Data_Function_flip__4284296032.Do(func() {
		cache_Data_Function_flip__4284296032 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__4284296032(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](b_1_box), a_2_box))}
})
	})
	return cache_Data_Function_flip__4284296032
}

var cache_Data_Function_flip__1288774720 gopurs_runtime.Value
var once_Data_Function_flip__1288774720 sync.Once
func Get_Data_Function_flip__1288774720() gopurs_runtime.Value {
	once_Data_Function_flip__1288774720.Do(func() {
		cache_Data_Function_flip__1288774720 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__1288774720(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](b_1_box), a_2_box))}
})
	})
	return cache_Data_Function_flip__1288774720
}

var cache_Data_Function_flip__1833071808 gopurs_runtime.Value
var once_Data_Function_flip__1833071808 sync.Once
func Get_Data_Function_flip__1833071808() gopurs_runtime.Value {
	once_Data_Function_flip__1833071808.Do(func() {
		cache_Data_Function_flip__1833071808 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__1833071808(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](b_1_box), a_2_box))}
})
	})
	return cache_Data_Function_flip__1833071808
}

var cache_Data_Function_flip__1350495552 gopurs_runtime.Value
var once_Data_Function_flip__1350495552 sync.Once
func Get_Data_Function_flip__1350495552() gopurs_runtime.Value {
	once_Data_Function_flip__1350495552.Do(func() {
		cache_Data_Function_flip__1350495552 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1350495552(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__1350495552
}

var cache_Data_Function_flip__3858636736 gopurs_runtime.Value
var once_Data_Function_flip__3858636736 sync.Once
func Get_Data_Function_flip__3858636736() gopurs_runtime.Value {
	once_Data_Function_flip__3858636736.Do(func() {
		cache_Data_Function_flip__3858636736 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3858636736(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3858636736
}

var cache_Data_Function_flip__1295484160 gopurs_runtime.Value
var once_Data_Function_flip__1295484160 sync.Once
func Get_Data_Function_flip__1295484160() gopurs_runtime.Value {
	once_Data_Function_flip__1295484160.Do(func() {
		cache_Data_Function_flip__1295484160 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__1295484160(f_0_box, b_1_box, a_2_box))}
})
	})
	return cache_Data_Function_flip__1295484160
}

var cache_Data_Function_flip__1673583840 gopurs_runtime.Value
var once_Data_Function_flip__1673583840 sync.Once
func Get_Data_Function_flip__1673583840() gopurs_runtime.Value {
	once_Data_Function_flip__1673583840.Do(func() {
		cache_Data_Function_flip__1673583840 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1673583840(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__1673583840
}

var cache_Data_Function_flip__534748448 gopurs_runtime.Value
var once_Data_Function_flip__534748448 sync.Once
func Get_Data_Function_flip__534748448() gopurs_runtime.Value {
	once_Data_Function_flip__534748448.Do(func() {
		cache_Data_Function_flip__534748448 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__534748448(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__534748448
}

var cache_Data_Function_flip__2270521984 gopurs_runtime.Value
var once_Data_Function_flip__2270521984 sync.Once
func Get_Data_Function_flip__2270521984() gopurs_runtime.Value {
	once_Data_Function_flip__2270521984.Do(func() {
		cache_Data_Function_flip__2270521984 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__2270521984(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__2270521984
}

var cache_Data_Function_flip__2106496000 gopurs_runtime.Value
var once_Data_Function_flip__2106496000 sync.Once
func Get_Data_Function_flip__2106496000() gopurs_runtime.Value {
	once_Data_Function_flip__2106496000.Do(func() {
		cache_Data_Function_flip__2106496000 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__2106496000(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__2106496000
}

var cache_Data_Function_flip__2857929472 gopurs_runtime.Value
var once_Data_Function_flip__2857929472 sync.Once
func Get_Data_Function_flip__2857929472() gopurs_runtime.Value {
	once_Data_Function_flip__2857929472.Do(func() {
		cache_Data_Function_flip__2857929472 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Function_flip__2857929472(f_0_box, b_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Function_flip__2857929472
}

var cache_Data_Function_flip__3136194560 gopurs_runtime.Value
var once_Data_Function_flip__3136194560 sync.Once
func Get_Data_Function_flip__3136194560() gopurs_runtime.Value {
	once_Data_Function_flip__3136194560.Do(func() {
		cache_Data_Function_flip__3136194560 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Function_flip__3136194560(f_0_box, b_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Function_flip__3136194560
}

var cache_Data_Function_flip__1006079424 gopurs_runtime.Value
var once_Data_Function_flip__1006079424 sync.Once
func Get_Data_Function_flip__1006079424() gopurs_runtime.Value {
	once_Data_Function_flip__1006079424.Do(func() {
		cache_Data_Function_flip__1006079424 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1006079424(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__1006079424
}

var cache_Data_Function_flip__3955399360 gopurs_runtime.Value
var once_Data_Function_flip__3955399360 sync.Once
func Get_Data_Function_flip__3955399360() gopurs_runtime.Value {
	once_Data_Function_flip__3955399360.Do(func() {
		cache_Data_Function_flip__3955399360 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3955399360(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3955399360
}

var cache_Data_Function_flip__4018036032 gopurs_runtime.Value
var once_Data_Function_flip__4018036032 sync.Once
func Get_Data_Function_flip__4018036032() gopurs_runtime.Value {
	once_Data_Function_flip__4018036032.Do(func() {
		cache_Data_Function_flip__4018036032 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__4018036032(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a_2_box))
})
	})
	return cache_Data_Function_flip__4018036032
}

var cache_Data_Function_flip__4154956096 gopurs_runtime.Value
var once_Data_Function_flip__4154956096 sync.Once
func Get_Data_Function_flip__4154956096() gopurs_runtime.Value {
	once_Data_Function_flip__4154956096.Do(func() {
		cache_Data_Function_flip__4154956096 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__4154956096(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a_2_box)))}
})
	})
	return cache_Data_Function_flip__4154956096
}

var cache_Data_Function_flip__682087456 gopurs_runtime.Value
var once_Data_Function_flip__682087456 sync.Once
func Get_Data_Function_flip__682087456() gopurs_runtime.Value {
	once_Data_Function_flip__682087456.Do(func() {
		cache_Data_Function_flip__682087456 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__682087456(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a_2_box)))}
})
	})
	return cache_Data_Function_flip__682087456
}

var cache_Data_Function_flip__2535084352 gopurs_runtime.Value
var once_Data_Function_flip__2535084352 sync.Once
func Get_Data_Function_flip__2535084352() gopurs_runtime.Value {
	once_Data_Function_flip__2535084352.Do(func() {
		cache_Data_Function_flip__2535084352 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__2535084352(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__2535084352
}

var cache_Data_Function_flip__4012165568 gopurs_runtime.Value
var once_Data_Function_flip__4012165568 sync.Once
func Get_Data_Function_flip__4012165568() gopurs_runtime.Value {
	once_Data_Function_flip__4012165568.Do(func() {
		cache_Data_Function_flip__4012165568 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__4012165568(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__4012165568
}

var cache_Data_Function_flip__2422603136 gopurs_runtime.Value
var once_Data_Function_flip__2422603136 sync.Once
func Get_Data_Function_flip__2422603136() gopurs_runtime.Value {
	once_Data_Function_flip__2422603136.Do(func() {
		cache_Data_Function_flip__2422603136 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__2422603136(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__2422603136
}

var cache_Data_Function_flip__1818565536 gopurs_runtime.Value
var once_Data_Function_flip__1818565536 sync.Once
func Get_Data_Function_flip__1818565536() gopurs_runtime.Value {
	once_Data_Function_flip__1818565536.Do(func() {
		cache_Data_Function_flip__1818565536 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1818565536(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__1818565536
}

var cache_Data_Function_flip__816740064 gopurs_runtime.Value
var once_Data_Function_flip__816740064 sync.Once
func Get_Data_Function_flip__816740064() gopurs_runtime.Value {
	once_Data_Function_flip__816740064.Do(func() {
		cache_Data_Function_flip__816740064 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__816740064(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__816740064
}

var cache_Data_Function_flip__1026675680 gopurs_runtime.Value
var once_Data_Function_flip__1026675680 sync.Once
func Get_Data_Function_flip__1026675680() gopurs_runtime.Value {
	once_Data_Function_flip__1026675680.Do(func() {
		cache_Data_Function_flip__1026675680 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1026675680(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__1026675680
}

var cache_Data_Function_flip__141921312 gopurs_runtime.Value
var once_Data_Function_flip__141921312 sync.Once
func Get_Data_Function_flip__141921312() gopurs_runtime.Value {
	once_Data_Function_flip__141921312.Do(func() {
		cache_Data_Function_flip__141921312 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__141921312(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__141921312
}

var cache_Data_Function_flip__2175652032 gopurs_runtime.Value
var once_Data_Function_flip__2175652032 sync.Once
func Get_Data_Function_flip__2175652032() gopurs_runtime.Value {
	once_Data_Function_flip__2175652032.Do(func() {
		cache_Data_Function_flip__2175652032 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__2175652032(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__2175652032
}

var cache_Data_Function_flip__848188896 gopurs_runtime.Value
var once_Data_Function_flip__848188896 sync.Once
func Get_Data_Function_flip__848188896() gopurs_runtime.Value {
	once_Data_Function_flip__848188896.Do(func() {
		cache_Data_Function_flip__848188896 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__848188896(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__848188896
}

var cache_Data_Function_flip__913470112 gopurs_runtime.Value
var once_Data_Function_flip__913470112 sync.Once
func Get_Data_Function_flip__913470112() gopurs_runtime.Value {
	once_Data_Function_flip__913470112.Do(func() {
		cache_Data_Function_flip__913470112 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__913470112(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__913470112
}

var cache_Data_Function_flip__4017289888 gopurs_runtime.Value
var once_Data_Function_flip__4017289888 sync.Once
func Get_Data_Function_flip__4017289888() gopurs_runtime.Value {
	once_Data_Function_flip__4017289888.Do(func() {
		cache_Data_Function_flip__4017289888 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__4017289888(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__4017289888
}

var cache_Data_Function_flip__3036739744 gopurs_runtime.Value
var once_Data_Function_flip__3036739744 sync.Once
func Get_Data_Function_flip__3036739744() gopurs_runtime.Value {
	once_Data_Function_flip__3036739744.Do(func() {
		cache_Data_Function_flip__3036739744 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3036739744(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__3036739744
}

var cache_Data_Function_flip__3539834656 gopurs_runtime.Value
var once_Data_Function_flip__3539834656 sync.Once
func Get_Data_Function_flip__3539834656() gopurs_runtime.Value {
	once_Data_Function_flip__3539834656.Do(func() {
		cache_Data_Function_flip__3539834656 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3539834656(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__3539834656
}

var cache_Data_Function_flip__2097869248 gopurs_runtime.Value
var once_Data_Function_flip__2097869248 sync.Once
func Get_Data_Function_flip__2097869248() gopurs_runtime.Value {
	once_Data_Function_flip__2097869248.Do(func() {
		cache_Data_Function_flip__2097869248 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__2097869248(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__2097869248
}

var cache_Data_Function_flip__3152946016 gopurs_runtime.Value
var once_Data_Function_flip__3152946016 sync.Once
func Get_Data_Function_flip__3152946016() gopurs_runtime.Value {
	once_Data_Function_flip__3152946016.Do(func() {
		cache_Data_Function_flip__3152946016 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3152946016(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__3152946016
}

var cache_Data_Function_flip__1036339264 gopurs_runtime.Value
var once_Data_Function_flip__1036339264 sync.Once
func Get_Data_Function_flip__1036339264() gopurs_runtime.Value {
	once_Data_Function_flip__1036339264.Do(func() {
		cache_Data_Function_flip__1036339264 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1036339264(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__1036339264
}

var cache_Data_Function_flip__754346432 gopurs_runtime.Value
var once_Data_Function_flip__754346432 sync.Once
func Get_Data_Function_flip__754346432() gopurs_runtime.Value {
	once_Data_Function_flip__754346432.Do(func() {
		cache_Data_Function_flip__754346432 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__754346432(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__754346432
}

var cache_Data_Function_flip__3959849920 gopurs_runtime.Value
var once_Data_Function_flip__3959849920 sync.Once
func Get_Data_Function_flip__3959849920() gopurs_runtime.Value {
	once_Data_Function_flip__3959849920.Do(func() {
		cache_Data_Function_flip__3959849920 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3959849920(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__3959849920
}

var cache_Data_Function_flip__4236888928 gopurs_runtime.Value
var once_Data_Function_flip__4236888928 sync.Once
func Get_Data_Function_flip__4236888928() gopurs_runtime.Value {
	once_Data_Function_flip__4236888928.Do(func() {
		cache_Data_Function_flip__4236888928 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__4236888928(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_2_box)))}
})
	})
	return cache_Data_Function_flip__4236888928
}

var cache_Data_Function_flip__1744188480 gopurs_runtime.Value
var once_Data_Function_flip__1744188480 sync.Once
func Get_Data_Function_flip__1744188480() gopurs_runtime.Value {
	once_Data_Function_flip__1744188480.Do(func() {
		cache_Data_Function_flip__1744188480 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1744188480(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](a_2_box))
})
	})
	return cache_Data_Function_flip__1744188480
}

var cache_Data_Function_flip__3468792800 gopurs_runtime.Value
var once_Data_Function_flip__3468792800 sync.Once
func Get_Data_Function_flip__3468792800() gopurs_runtime.Value {
	once_Data_Function_flip__3468792800.Do(func() {
		cache_Data_Function_flip__3468792800 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__3468792800(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](a_2_box)))}
})
	})
	return cache_Data_Function_flip__3468792800
}

var cache_Data_Function_flip__2373571712 gopurs_runtime.Value
var once_Data_Function_flip__2373571712 sync.Once
func Get_Data_Function_flip__2373571712() gopurs_runtime.Value {
	once_Data_Function_flip__2373571712.Do(func() {
		cache_Data_Function_flip__2373571712 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__2373571712(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](a_2_box)))}
})
	})
	return cache_Data_Function_flip__2373571712
}

var cache_Data_Function_flip__437353440 gopurs_runtime.Value
var once_Data_Function_flip__437353440 sync.Once
func Get_Data_Function_flip__437353440() gopurs_runtime.Value {
	once_Data_Function_flip__437353440.Do(func() {
		cache_Data_Function_flip__437353440 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__437353440(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](a_2_box))
})
	})
	return cache_Data_Function_flip__437353440
}

var cache_Data_Function_flip__1172723328 gopurs_runtime.Value
var once_Data_Function_flip__1172723328 sync.Once
func Get_Data_Function_flip__1172723328() gopurs_runtime.Value {
	once_Data_Function_flip__1172723328.Do(func() {
		cache_Data_Function_flip__1172723328 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1172723328(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](a_2_box))
})
	})
	return cache_Data_Function_flip__1172723328
}

var cache_Data_Function_flip__4067779648 gopurs_runtime.Value
var once_Data_Function_flip__4067779648 sync.Once
func Get_Data_Function_flip__4067779648() gopurs_runtime.Value {
	once_Data_Function_flip__4067779648.Do(func() {
		cache_Data_Function_flip__4067779648 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__4067779648(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](a_2_box))
})
	})
	return cache_Data_Function_flip__4067779648
}

var cache_Data_Function_flip__1273449376 gopurs_runtime.Value
var once_Data_Function_flip__1273449376 sync.Once
func Get_Data_Function_flip__1273449376() gopurs_runtime.Value {
	once_Data_Function_flip__1273449376.Do(func() {
		cache_Data_Function_flip__1273449376 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1273449376(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_2_box))
})
	})
	return cache_Data_Function_flip__1273449376
}

var cache_Data_Function_flip__3468501600 gopurs_runtime.Value
var once_Data_Function_flip__3468501600 sync.Once
func Get_Data_Function_flip__3468501600() gopurs_runtime.Value {
	once_Data_Function_flip__3468501600.Do(func() {
		cache_Data_Function_flip__3468501600 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Function_flip__3468501600(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_2_box)))
})
	})
	return cache_Data_Function_flip__3468501600
}

var cache_Data_Function_flip__1570464832 gopurs_runtime.Value
var once_Data_Function_flip__1570464832 sync.Once
func Get_Data_Function_flip__1570464832() gopurs_runtime.Value {
	once_Data_Function_flip__1570464832.Do(func() {
		cache_Data_Function_flip__1570464832 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__1570464832(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](a_2_box)))}
})
	})
	return cache_Data_Function_flip__1570464832
}

var cache_Data_Function_flip__1764632704 gopurs_runtime.Value
var once_Data_Function_flip__1764632704 sync.Once
func Get_Data_Function_flip__1764632704() gopurs_runtime.Value {
	once_Data_Function_flip__1764632704.Do(func() {
		cache_Data_Function_flip__1764632704 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1764632704(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](a_2_box))
})
	})
	return cache_Data_Function_flip__1764632704
}

var cache_Data_Function_flip__362459520 gopurs_runtime.Value
var once_Data_Function_flip__362459520 sync.Once
func Get_Data_Function_flip__362459520() gopurs_runtime.Value {
	once_Data_Function_flip__362459520.Do(func() {
		cache_Data_Function_flip__362459520 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Function_flip__362459520(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](a_2_box)))
})
	})
	return cache_Data_Function_flip__362459520
}

var cache_Data_Function_flip__3192966848 gopurs_runtime.Value
var once_Data_Function_flip__3192966848 sync.Once
func Get_Data_Function_flip__3192966848() gopurs_runtime.Value {
	once_Data_Function_flip__3192966848.Do(func() {
		cache_Data_Function_flip__3192966848 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3192966848(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](a_2_box))
})
	})
	return cache_Data_Function_flip__3192966848
}

var cache_Data_Function_flip__2681340064 gopurs_runtime.Value
var once_Data_Function_flip__2681340064 sync.Once
func Get_Data_Function_flip__2681340064() gopurs_runtime.Value {
	once_Data_Function_flip__2681340064.Do(func() {
		cache_Data_Function_flip__2681340064 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__2681340064(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](a_2_box)))}
})
	})
	return cache_Data_Function_flip__2681340064
}

var cache_Data_Function_flip__3261866592 gopurs_runtime.Value
var once_Data_Function_flip__3261866592 sync.Once
func Get_Data_Function_flip__3261866592() gopurs_runtime.Value {
	once_Data_Function_flip__3261866592.Do(func() {
		cache_Data_Function_flip__3261866592 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3261866592(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3261866592
}

var cache_Data_Function_flip__872296768 gopurs_runtime.Value
var once_Data_Function_flip__872296768 sync.Once
func Get_Data_Function_flip__872296768() gopurs_runtime.Value {
	once_Data_Function_flip__872296768.Do(func() {
		cache_Data_Function_flip__872296768 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__872296768(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__872296768
}

var cache_Data_Function_flip__4091748192 gopurs_runtime.Value
var once_Data_Function_flip__4091748192 sync.Once
func Get_Data_Function_flip__4091748192() gopurs_runtime.Value {
	once_Data_Function_flip__4091748192.Do(func() {
		cache_Data_Function_flip__4091748192 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__4091748192(f_0_box, b_1_box, a_2_box))}
})
	})
	return cache_Data_Function_flip__4091748192
}

var cache_Data_Function_flip__535340480 gopurs_runtime.Value
var once_Data_Function_flip__535340480 sync.Once
func Get_Data_Function_flip__535340480() gopurs_runtime.Value {
	once_Data_Function_flip__535340480.Do(func() {
		cache_Data_Function_flip__535340480 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__535340480(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__535340480
}

var cache_Data_Function_flip__1093730912 gopurs_runtime.Value
var once_Data_Function_flip__1093730912 sync.Once
func Get_Data_Function_flip__1093730912() gopurs_runtime.Value {
	once_Data_Function_flip__1093730912.Do(func() {
		cache_Data_Function_flip__1093730912 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1093730912(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__1093730912
}

var cache_Data_Function_flip__2253242624 gopurs_runtime.Value
var once_Data_Function_flip__2253242624 sync.Once
func Get_Data_Function_flip__2253242624() gopurs_runtime.Value {
	once_Data_Function_flip__2253242624.Do(func() {
		cache_Data_Function_flip__2253242624 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__2253242624(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__2253242624
}

var cache_Data_Function_flip__3659058336 gopurs_runtime.Value
var once_Data_Function_flip__3659058336 sync.Once
func Get_Data_Function_flip__3659058336() gopurs_runtime.Value {
	once_Data_Function_flip__3659058336.Do(func() {
		cache_Data_Function_flip__3659058336 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3659058336(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3659058336
}

var cache_Data_Function_flip__1074184704 gopurs_runtime.Value
var once_Data_Function_flip__1074184704 sync.Once
func Get_Data_Function_flip__1074184704() gopurs_runtime.Value {
	once_Data_Function_flip__1074184704.Do(func() {
		cache_Data_Function_flip__1074184704 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1074184704(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__1074184704
}

var cache_Data_Function_flip__3482198080 gopurs_runtime.Value
var once_Data_Function_flip__3482198080 sync.Once
func Get_Data_Function_flip__3482198080() gopurs_runtime.Value {
	once_Data_Function_flip__3482198080.Do(func() {
		cache_Data_Function_flip__3482198080 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3482198080(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3482198080
}

var cache_Data_Function_flip__1931071680 gopurs_runtime.Value
var once_Data_Function_flip__1931071680 sync.Once
func Get_Data_Function_flip__1931071680() gopurs_runtime.Value {
	once_Data_Function_flip__1931071680.Do(func() {
		cache_Data_Function_flip__1931071680 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__1931071680(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__1931071680
}

var cache_Data_Function_flip__2033550272 gopurs_runtime.Value
var once_Data_Function_flip__2033550272 sync.Once
func Get_Data_Function_flip__2033550272() gopurs_runtime.Value {
	once_Data_Function_flip__2033550272.Do(func() {
		cache_Data_Function_flip__2033550272 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__2033550272(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](b_1_box), a_2_box))}
})
	})
	return cache_Data_Function_flip__2033550272
}

var cache_Data_Function_flip__3047022592 gopurs_runtime.Value
var once_Data_Function_flip__3047022592 sync.Once
func Get_Data_Function_flip__3047022592() gopurs_runtime.Value {
	once_Data_Function_flip__3047022592.Do(func() {
		cache_Data_Function_flip__3047022592 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_flip__3047022592(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_Data_Function_flip__3047022592
}

var cache_Data_Function_on__3122155169 gopurs_runtime.Value
var once_Data_Function_on__3122155169 sync.Once
func Get_Data_Function_on__3122155169() gopurs_runtime.Value {
	once_Data_Function_on__3122155169.Do(func() {
		cache_Data_Function_on__3122155169 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Function_on__3122155169(f_0_box, g_1_box, x_2_box, y_3_box))
})
	})
	return cache_Data_Function_on__3122155169
}

var cache_Data_Function_on__3980724833 gopurs_runtime.Value
var once_Data_Function_on__3980724833 sync.Once
func Get_Data_Function_on__3980724833() gopurs_runtime.Value {
	once_Data_Function_on__3980724833.Do(func() {
		cache_Data_Function_on__3980724833 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Function_on__3980724833(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_Data_Function_on__3980724833
}

var cache_Data_Function_on__3556844193 gopurs_runtime.Value
var once_Data_Function_on__3556844193 sync.Once
func Get_Data_Function_on__3556844193() gopurs_runtime.Value {
	once_Data_Function_on__3556844193.Do(func() {
		cache_Data_Function_on__3556844193 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Function_on__3556844193(f_0_box, g_1_box, x_2_box, y_3_box)), UnsafePtr: nil}
})
	})
	return cache_Data_Function_on__3556844193
}

func Call_Data_Function_on(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_Data_Function_flip(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_go__const(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_applyN(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_0 gopurs_runtime.Value
go__go_1_0_0 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_2_loop int64 = n_2_loop_val.IntVal
var acc_3_loop gopurs_runtime.Value = acc_3_loop_val
go__go_1_0_0:
for {
if false { continue go__go_1_0_0 }
var n_2 int64 = n_2_loop
_ = n_2
var acc_3 gopurs_runtime.Value = acc_3_loop
_ = acc_3
var __t2 gopurs_runtime.Value
{
var __t1 bool
{
if (n_2) > (0) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
if __t1 {
__t2 = acc_3
goto end_branch_2
} else {

}
}
{
n_2_loop = (n_2) - (1)
acc_3_loop = gopurs_runtime.Apply(f_0, acc_3)
continue go__go_1_0_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
return __t2
}
}()
})
})
return go__go_1_0_0
}

func Call_Data_Function_applyFlipped(x_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(f_1, x_0)
}

func Call_Data_Function_apply(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, x_1)
}

func Call_Data_Function_apply__458711162(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, x_1)
}

func Call_Data_Function_const__3524684546(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__1496134642(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__1832354163(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__1426827922(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__2569866098(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__4181451586(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__2390202835(a_0_loop bool, v_1_loop gopurs_runtime.Value) bool {
var a_0 bool = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__702735379(a_0_loop *Constructor_Data_Maybe_Just, v_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var a_0 *Constructor_Data_Maybe_Just = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__3834287059(a_0_loop *Constructor_Data_Maybe_Just, v_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var a_0 *Constructor_Data_Maybe_Just = a_0_loop
_ = a_0
var v_1 *Constructor_Data_Maybe_Just = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__1243414737(a_0_loop int64, v_1_loop int64) int64 {
var a_0 int64 = a_0_loop
_ = a_0
var v_1 int64 = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__3567418756(a_0_loop int64, v_1_loop gopurs_runtime.Value) int64 {
var a_0 int64 = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__2082174484(a_0_loop int64, v_1_loop gopurs_runtime.Value) int64 {
var a_0 int64 = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__2632676769(a_0_loop float64, v_1_loop float64) float64 {
var a_0 float64 = a_0_loop
_ = a_0
var v_1 float64 = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__2696088948(a_0_loop float64, v_1_loop gopurs_runtime.Value) float64 {
var a_0 float64 = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__4157258135(a_0_loop string, v_1_loop string) string {
var a_0 string = a_0_loop
_ = a_0
var v_1 string = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__1562253172(a_0_loop string, v_1_loop gopurs_runtime.Value) string {
var a_0 string = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__220790420(a_0_loop bool, v_1_loop gopurs_runtime.Value) bool {
var a_0 bool = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__4026847508(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__2857921436(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__744888567(a_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 string = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__2050378404(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__2291477660(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__72843521(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__3848686068(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__992538388(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__1803609684(a_0_loop uint32, v_1_loop gopurs_runtime.Value) uint32 {
var a_0 uint32 = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__2568740136(a_0_loop uint32, v_1_loop uint32) uint32 {
var a_0 uint32 = a_0_loop
_ = a_0
var v_1 uint32 = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__1776115220(a_0_loop *Constructor_Data_Date_Date, v_1_loop gopurs_runtime.Value) *Constructor_Data_Date_Date {
var a_0 *Constructor_Data_Date_Date = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__1863643516(a_0_loop *Constructor_Data_Date_Date, v_1_loop *Constructor_Data_Date_Date) *Constructor_Data_Date_Date {
var a_0 *Constructor_Data_Date_Date = a_0_loop
_ = a_0
var v_1 *Constructor_Data_Date_Date = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__1245418657(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__1628252324(a_0_loop *Constructor_Data_Map_Internal_IterNext, v_1_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_IterNext {
var a_0 *Constructor_Data_Map_Internal_IterNext = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__3470624004(a_0_loop *Constructor_Data_Maybe_Just, v_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var a_0 *Constructor_Data_Maybe_Just = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__587810660(a_0_loop *Constructor_Data_Maybe_Just, v_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var a_0 *Constructor_Data_Maybe_Just = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__3525124788(a_0_loop *Constructor_Data_Maybe_Just, v_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var a_0 *Constructor_Data_Maybe_Just = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__153462564(a_0_loop *Constructor_Data_Maybe_Just, v_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var a_0 *Constructor_Data_Maybe_Just = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__839298276(a_0_loop *Constructor_Data_Maybe_Just, v_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var a_0 *Constructor_Data_Maybe_Just = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__3168103323(a_0_loop *Constructor_Data_Maybe_Just, v_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var a_0 *Constructor_Data_Maybe_Just = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__106415396(a_0_loop *Constructor_Data_Maybe_Just, v_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var a_0 *Constructor_Data_Maybe_Just = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__1155968100(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__73052052(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__3415939124(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__2189647754(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__4270360676(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__4189285076(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__3953240484(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__2557237620(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__3858583060(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__3952683620(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__754826900(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_const__4115900574(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_Data_Function_flip__2673533882(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__1826582752(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Date_Date, a_2_loop int64) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Date_Date = b_1_loop
_ = b_1
var a_2 int64 = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(a_2), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__2176113088(f_0_loop gopurs_runtime.Value, b_1_loop string, a_2_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 string = b_1_loop
_ = b_1
var a_2 string = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Str(a_2), gopurs_runtime.Str(b_1))
}

func Call_Data_Function_flip__4032303232(f_0_loop gopurs_runtime.Value, b_1_loop string, a_2_loop string) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 string = b_1_loop
_ = b_1
var a_2 string = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(f_0, gopurs_runtime.Str(a_2), gopurs_runtime.Str(b_1)))
}

func Call_Data_Function_flip__1384624704(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__4257609536(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__1760577696(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__2226752640(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0, a_2, b_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Function_flip__3110498784(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__1540101856(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *Constructor_Data_Generic_Rep_Product {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Product](gopurs_runtime.Apply2(f_0, a_2, b_1))
}

func Call_Data_Function_flip__3658931456(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__2974723072(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(f_0, a_2, b_1))
}

func Call_Data_Function_flip__73563104(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(f_0, a_2, b_1))
}

func Call_Data_Function_flip__787327104(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__1117087808(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__3021024128(f_0_loop gopurs_runtime.Value, b_1_loop []gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 []gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Array(b_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Function_flip__3904800000(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__2928880736(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__3019719488(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Date_Date, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Date_Date = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(b_1)})
}

func Call_Data_Function_flip__2733110176(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Date_Date, a_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Date_Date = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__3019832928(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__3675729664(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__3709724320(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_List_Types_Cons, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_List_Types_Cons = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_1)})
}

func Call_Data_Function_flip__3563101792(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_List_Types_Cons, a_2_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_List_Types_Cons = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__4284296032(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Maybe_Just, a_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Maybe_Just = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__1288774720(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Maybe_Just, a_2_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Maybe_Just = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__1833071808(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_NonEmpty_NonEmpty, a_2_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_NonEmpty_NonEmpty = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__1350495552(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__3858636736(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__1295484160(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(f_0, a_2, b_1))
}

func Call_Data_Function_flip__1673583840(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__534748448(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__2270521984(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__2106496000(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__2857929472(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 []gopurs_runtime.Value = a_2_loop
_ = a_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0, gopurs_runtime.Array(a_2), b_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Function_flip__3136194560(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 []gopurs_runtime.Value = a_2_loop
_ = a_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0, gopurs_runtime.Array(a_2), b_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Function_flip__1006079424(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__3955399360(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__4018036032(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Date_Date, a_2_loop *Constructor_Data_Date_Date) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Date_Date = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Date_Date = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(b_1)})
}

func Call_Data_Function_flip__4154956096(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Date_Date, a_2_loop *Constructor_Data_Date_Date) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Date_Date = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Date_Date = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__682087456(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Time_Time, a_2_loop *Constructor_Data_Date_Date) *Constructor_Data_DateTime_DateTime {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Time_Time = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Date_Date = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__2535084352(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__4012165568(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__2422603136(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__1818565536(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__816740064(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__1026675680(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__141921312(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__2175652032(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__848188896(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__913470112(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__4017289888(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__3036739744(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__3539834656(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__2097869248(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__3152946016(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__1036339264(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__754346432(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Tuple_Tuple, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Tuple_Tuple = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_1)})
}

func Call_Data_Function_flip__3959849920(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Tuple_Tuple, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Tuple_Tuple = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_1)})
}

func Call_Data_Function_flip__4236888928(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Tuple_Tuple, a_2_loop *Constructor_Data_List_Lazy_Types_Cons) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Tuple_Tuple = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Lazy_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__1744188480(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_List_Types_Cons, a_2_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_List_Types_Cons = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_1)})
}

func Call_Data_Function_flip__3468792800(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_List_Types_Cons, a_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_NonEmpty_NonEmpty {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_List_Types_Cons = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__2373571712(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_List_Types_Cons = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_2)}, b_1))
}

func Call_Data_Function_flip__437353440(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Map_Internal_Node = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__1172723328(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Map_Internal_Node = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__4067779648(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Map_Internal_Node = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__1273449376(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Maybe_Just = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__3468501600(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_Maybe_Just) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Maybe_Just = a_2_loop
_ = a_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a_2)}, b_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Function_flip__1570464832(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_NonEmpty_NonEmpty = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(a_2)}, b_1))
}

func Call_Data_Function_flip__1764632704(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Tuple_Tuple = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(a_2)}, b_1)
}

func Call_Data_Function_flip__362459520(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_Tuple_Tuple) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Tuple_Tuple = a_2_loop
_ = a_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(a_2)}, b_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Function_flip__3192966848(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Tuple_Tuple, a_2_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Tuple_Tuple = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Tuple_Tuple = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_1)})
}

func Call_Data_Function_flip__2681340064(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Tuple_Tuple, a_2_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Tuple_Tuple = b_1_loop
_ = b_1
var a_2 *Constructor_Data_Tuple_Tuple = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__3261866592(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__872296768(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__4091748192(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(f_0, a_2, b_1))
}

func Call_Data_Function_flip__535340480(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__1093730912(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__2253242624(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__3659058336(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__1074184704(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__3482198080(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__1931071680(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__2033550272(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Maybe_Just, a_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Maybe_Just = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(b_1)}))
}

func Call_Data_Function_flip__3047022592(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_on__3122155169(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) bool {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return (gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3)).IntVal) != (0)
}

func Call_Data_Function_on__3980724833(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_Data_Function_on__3556844193(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) uint32 {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return uint32(gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3)).IntVal)
}


