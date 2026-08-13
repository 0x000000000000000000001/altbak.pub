package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Apply_Apply_dollarDict gopurs_runtime.Value
var once_Control_Apply_Apply_dollarDict sync.Once
func Get_Control_Apply_Apply_dollarDict() gopurs_runtime.Value {
	once_Control_Apply_Apply_dollarDict.Do(func() {
		cache_Control_Apply_Apply_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_Apply_dollarDict(x_0_box)
})
	})
	return cache_Control_Apply_Apply_dollarDict
}

var cache_Control_Apply_applyProxy gopurs_runtime.Value
var once_Control_Apply_applyProxy sync.Once
func Get_Control_Apply_applyProxy() gopurs_runtime.Value {
	once_Control_Apply_applyProxy.Do(func() {
		cache_Control_Apply_applyProxy = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_Control_Apply_applyProxy
}

var cache_Control_Apply_applyFn gopurs_runtime.Value
var once_Control_Apply_applyFn sync.Once
func Get_Control_Apply_applyFn() gopurs_runtime.Value {
	once_Control_Apply_applyFn.Do(func() {
		cache_Control_Apply_applyFn = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorFn()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_Control_Apply_applyFn
}

var cache_Control_Apply_applyArray gopurs_runtime.Value
var once_Control_Apply_applyArray sync.Once
func Get_Control_Apply_applyArray() gopurs_runtime.Value {
	once_Control_Apply_applyArray.Do(func() {
		cache_Control_Apply_applyArray = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorArray()
}), Get_Control_Apply_arrayApply())
	})
	return cache_Control_Apply_applyArray
}

var cache_Control_Apply_apply gopurs_runtime.Value
var once_Control_Apply_apply sync.Once
func Get_Control_Apply_apply() gopurs_runtime.Value {
	once_Control_Apply_apply.Do(func() {
		cache_Control_Apply_apply = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply
}

var cache_Control_Apply_applyFirst gopurs_runtime.Value
var once_Control_Apply_applyFirst sync.Once
func Get_Control_Apply_applyFirst() gopurs_runtime.Value {
	once_Control_Apply_applyFirst.Do(func() {
		cache_Control_Apply_applyFirst = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_applyFirst(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_applyFirst
}

var cache_Control_Apply_applySecond gopurs_runtime.Value
var once_Control_Apply_applySecond sync.Once
func Get_Control_Apply_applySecond() gopurs_runtime.Value {
	once_Control_Apply_applySecond.Do(func() {
		cache_Control_Apply_applySecond = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_applySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_applySecond
}

var cache_Control_Apply_lift2 gopurs_runtime.Value
var once_Control_Apply_lift2 sync.Once
func Get_Control_Apply_lift2() gopurs_runtime.Value {
	once_Control_Apply_lift2.Do(func() {
		cache_Control_Apply_lift2 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2
}

var cache_Control_Apply_lift3 gopurs_runtime.Value
var once_Control_Apply_lift3 sync.Once
func Get_Control_Apply_lift3() gopurs_runtime.Value {
	once_Control_Apply_lift3.Do(func() {
		cache_Control_Apply_lift3 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift3(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift3
}

var cache_Control_Apply_lift4 gopurs_runtime.Value
var once_Control_Apply_lift4 sync.Once
func Get_Control_Apply_lift4() gopurs_runtime.Value {
	once_Control_Apply_lift4.Do(func() {
		cache_Control_Apply_lift4 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift4(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift4
}

var cache_Control_Apply_lift5 gopurs_runtime.Value
var once_Control_Apply_lift5 sync.Once
func Get_Control_Apply_lift5() gopurs_runtime.Value {
	once_Control_Apply_lift5.Do(func() {
		cache_Control_Apply_lift5 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift5(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift5
}

var cache_Control_Apply_apply__1159579293 gopurs_runtime.Value
var once_Control_Apply_apply__1159579293 sync.Once
func Get_Control_Apply_apply__1159579293() gopurs_runtime.Value {
	once_Control_Apply_apply__1159579293.Do(func() {
		cache_Control_Apply_apply__1159579293 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__1159579293(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__1159579293
}

var cache_Control_Apply_apply__1030762512 gopurs_runtime.Value
var once_Control_Apply_apply__1030762512 sync.Once
func Get_Control_Apply_apply__1030762512() gopurs_runtime.Value {
	once_Control_Apply_apply__1030762512.Do(func() {
		cache_Control_Apply_apply__1030762512 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__1030762512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__1030762512
}

var cache_Control_Apply_apply__2174094052 gopurs_runtime.Value
var once_Control_Apply_apply__2174094052 sync.Once
func Get_Control_Apply_apply__2174094052() gopurs_runtime.Value {
	once_Control_Apply_apply__2174094052.Do(func() {
		cache_Control_Apply_apply__2174094052 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2174094052(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2174094052
}

var cache_Control_Apply_apply__652936900 gopurs_runtime.Value
var once_Control_Apply_apply__652936900 sync.Once
func Get_Control_Apply_apply__652936900() gopurs_runtime.Value {
	once_Control_Apply_apply__652936900.Do(func() {
		cache_Control_Apply_apply__652936900 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__652936900(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__652936900
}

var cache_Control_Apply_apply__3477321027 gopurs_runtime.Value
var once_Control_Apply_apply__3477321027 sync.Once
func Get_Control_Apply_apply__3477321027() gopurs_runtime.Value {
	once_Control_Apply_apply__3477321027.Do(func() {
		cache_Control_Apply_apply__3477321027 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__3477321027(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__3477321027
}

var cache_Control_Apply_apply__350765120 gopurs_runtime.Value
var once_Control_Apply_apply__350765120 sync.Once
func Get_Control_Apply_apply__350765120() gopurs_runtime.Value {
	once_Control_Apply_apply__350765120.Do(func() {
		cache_Control_Apply_apply__350765120 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__350765120(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__350765120
}

var cache_Control_Apply_apply__5413771 gopurs_runtime.Value
var once_Control_Apply_apply__5413771 sync.Once
func Get_Control_Apply_apply__5413771() gopurs_runtime.Value {
	once_Control_Apply_apply__5413771.Do(func() {
		cache_Control_Apply_apply__5413771 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__5413771(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__5413771
}

var cache_Control_Apply_apply__4203183626 gopurs_runtime.Value
var once_Control_Apply_apply__4203183626 sync.Once
func Get_Control_Apply_apply__4203183626() gopurs_runtime.Value {
	once_Control_Apply_apply__4203183626.Do(func() {
		cache_Control_Apply_apply__4203183626 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__4203183626(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__4203183626
}

var cache_Control_Apply_apply__1982519530 gopurs_runtime.Value
var once_Control_Apply_apply__1982519530 sync.Once
func Get_Control_Apply_apply__1982519530() gopurs_runtime.Value {
	once_Control_Apply_apply__1982519530.Do(func() {
		cache_Control_Apply_apply__1982519530 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__1982519530(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__1982519530
}

var cache_Control_Apply_apply__2962221386 gopurs_runtime.Value
var once_Control_Apply_apply__2962221386 sync.Once
func Get_Control_Apply_apply__2962221386() gopurs_runtime.Value {
	once_Control_Apply_apply__2962221386.Do(func() {
		cache_Control_Apply_apply__2962221386 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2962221386(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2962221386
}

var cache_Control_Apply_apply__4140882028 gopurs_runtime.Value
var once_Control_Apply_apply__4140882028 sync.Once
func Get_Control_Apply_apply__4140882028() gopurs_runtime.Value {
	once_Control_Apply_apply__4140882028.Do(func() {
		cache_Control_Apply_apply__4140882028 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__4140882028(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__4140882028
}

var cache_Control_Apply_apply__2008680204 gopurs_runtime.Value
var once_Control_Apply_apply__2008680204 sync.Once
func Get_Control_Apply_apply__2008680204() gopurs_runtime.Value {
	once_Control_Apply_apply__2008680204.Do(func() {
		cache_Control_Apply_apply__2008680204 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2008680204(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2008680204
}

var cache_Control_Apply_apply__2092855244 gopurs_runtime.Value
var once_Control_Apply_apply__2092855244 sync.Once
func Get_Control_Apply_apply__2092855244() gopurs_runtime.Value {
	once_Control_Apply_apply__2092855244.Do(func() {
		cache_Control_Apply_apply__2092855244 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2092855244(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2092855244
}

var cache_Control_Apply_apply__353515660 gopurs_runtime.Value
var once_Control_Apply_apply__353515660 sync.Once
func Get_Control_Apply_apply__353515660() gopurs_runtime.Value {
	once_Control_Apply_apply__353515660.Do(func() {
		cache_Control_Apply_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__353515660(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__353515660
}

var cache_Control_Apply_apply__2080233356 gopurs_runtime.Value
var once_Control_Apply_apply__2080233356 sync.Once
func Get_Control_Apply_apply__2080233356() gopurs_runtime.Value {
	once_Control_Apply_apply__2080233356.Do(func() {
		cache_Control_Apply_apply__2080233356 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2080233356(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2080233356
}

var cache_Control_Apply_apply__197474060 gopurs_runtime.Value
var once_Control_Apply_apply__197474060 sync.Once
func Get_Control_Apply_apply__197474060() gopurs_runtime.Value {
	once_Control_Apply_apply__197474060.Do(func() {
		cache_Control_Apply_apply__197474060 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__197474060(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__197474060
}

var cache_Control_Apply_apply__1526260460 gopurs_runtime.Value
var once_Control_Apply_apply__1526260460 sync.Once
func Get_Control_Apply_apply__1526260460() gopurs_runtime.Value {
	once_Control_Apply_apply__1526260460.Do(func() {
		cache_Control_Apply_apply__1526260460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__1526260460(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__1526260460
}

var cache_Control_Apply_apply__1023214444 gopurs_runtime.Value
var once_Control_Apply_apply__1023214444 sync.Once
func Get_Control_Apply_apply__1023214444() gopurs_runtime.Value {
	once_Control_Apply_apply__1023214444.Do(func() {
		cache_Control_Apply_apply__1023214444 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__1023214444(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__1023214444
}

var cache_Control_Apply_apply__2623550860 gopurs_runtime.Value
var once_Control_Apply_apply__2623550860 sync.Once
func Get_Control_Apply_apply__2623550860() gopurs_runtime.Value {
	once_Control_Apply_apply__2623550860.Do(func() {
		cache_Control_Apply_apply__2623550860 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2623550860(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2623550860
}

var cache_Control_Apply_apply__4046394764 gopurs_runtime.Value
var once_Control_Apply_apply__4046394764 sync.Once
func Get_Control_Apply_apply__4046394764() gopurs_runtime.Value {
	once_Control_Apply_apply__4046394764.Do(func() {
		cache_Control_Apply_apply__4046394764 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__4046394764(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__4046394764
}

var cache_Control_Apply_apply__2021897708 gopurs_runtime.Value
var once_Control_Apply_apply__2021897708 sync.Once
func Get_Control_Apply_apply__2021897708() gopurs_runtime.Value {
	once_Control_Apply_apply__2021897708.Do(func() {
		cache_Control_Apply_apply__2021897708 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2021897708(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2021897708
}

var cache_Control_Apply_apply__2908293516 gopurs_runtime.Value
var once_Control_Apply_apply__2908293516 sync.Once
func Get_Control_Apply_apply__2908293516() gopurs_runtime.Value {
	once_Control_Apply_apply__2908293516.Do(func() {
		cache_Control_Apply_apply__2908293516 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2908293516(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2908293516
}

var cache_Control_Apply_apply__3783667596 gopurs_runtime.Value
var once_Control_Apply_apply__3783667596 sync.Once
func Get_Control_Apply_apply__3783667596() gopurs_runtime.Value {
	once_Control_Apply_apply__3783667596.Do(func() {
		cache_Control_Apply_apply__3783667596 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__3783667596(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__3783667596
}

var cache_Control_Apply_apply__3582447820 gopurs_runtime.Value
var once_Control_Apply_apply__3582447820 sync.Once
func Get_Control_Apply_apply__3582447820() gopurs_runtime.Value {
	once_Control_Apply_apply__3582447820.Do(func() {
		cache_Control_Apply_apply__3582447820 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__3582447820(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__3582447820
}

var cache_Control_Apply_apply__1099902892 gopurs_runtime.Value
var once_Control_Apply_apply__1099902892 sync.Once
func Get_Control_Apply_apply__1099902892() gopurs_runtime.Value {
	once_Control_Apply_apply__1099902892.Do(func() {
		cache_Control_Apply_apply__1099902892 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__1099902892(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__1099902892
}

var cache_Control_Apply_apply__1851858028 gopurs_runtime.Value
var once_Control_Apply_apply__1851858028 sync.Once
func Get_Control_Apply_apply__1851858028() gopurs_runtime.Value {
	once_Control_Apply_apply__1851858028.Do(func() {
		cache_Control_Apply_apply__1851858028 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__1851858028(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__1851858028
}

var cache_Control_Apply_apply__1471954540 gopurs_runtime.Value
var once_Control_Apply_apply__1471954540 sync.Once
func Get_Control_Apply_apply__1471954540() gopurs_runtime.Value {
	once_Control_Apply_apply__1471954540.Do(func() {
		cache_Control_Apply_apply__1471954540 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__1471954540(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__1471954540
}

var cache_Control_Apply_apply__3063957068 gopurs_runtime.Value
var once_Control_Apply_apply__3063957068 sync.Once
func Get_Control_Apply_apply__3063957068() gopurs_runtime.Value {
	once_Control_Apply_apply__3063957068.Do(func() {
		cache_Control_Apply_apply__3063957068 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__3063957068(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__3063957068
}

var cache_Control_Apply_apply__2007181164 gopurs_runtime.Value
var once_Control_Apply_apply__2007181164 sync.Once
func Get_Control_Apply_apply__2007181164() gopurs_runtime.Value {
	once_Control_Apply_apply__2007181164.Do(func() {
		cache_Control_Apply_apply__2007181164 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2007181164(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2007181164
}

var cache_Control_Apply_apply__986161100 gopurs_runtime.Value
var once_Control_Apply_apply__986161100 sync.Once
func Get_Control_Apply_apply__986161100() gopurs_runtime.Value {
	once_Control_Apply_apply__986161100.Do(func() {
		cache_Control_Apply_apply__986161100 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__986161100(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__986161100
}

var cache_Control_Apply_apply__2947753132 gopurs_runtime.Value
var once_Control_Apply_apply__2947753132 sync.Once
func Get_Control_Apply_apply__2947753132() gopurs_runtime.Value {
	once_Control_Apply_apply__2947753132.Do(func() {
		cache_Control_Apply_apply__2947753132 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2947753132(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2947753132
}

var cache_Control_Apply_apply__2234240780 gopurs_runtime.Value
var once_Control_Apply_apply__2234240780 sync.Once
func Get_Control_Apply_apply__2234240780() gopurs_runtime.Value {
	once_Control_Apply_apply__2234240780.Do(func() {
		cache_Control_Apply_apply__2234240780 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2234240780(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2234240780
}

var cache_Control_Apply_apply__2758496428 gopurs_runtime.Value
var once_Control_Apply_apply__2758496428 sync.Once
func Get_Control_Apply_apply__2758496428() gopurs_runtime.Value {
	once_Control_Apply_apply__2758496428.Do(func() {
		cache_Control_Apply_apply__2758496428 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2758496428(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2758496428
}

var cache_Control_Apply_apply__2087590060 gopurs_runtime.Value
var once_Control_Apply_apply__2087590060 sync.Once
func Get_Control_Apply_apply__2087590060() gopurs_runtime.Value {
	once_Control_Apply_apply__2087590060.Do(func() {
		cache_Control_Apply_apply__2087590060 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2087590060(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2087590060
}

var cache_Control_Apply_apply__3302223884 gopurs_runtime.Value
var once_Control_Apply_apply__3302223884 sync.Once
func Get_Control_Apply_apply__3302223884() gopurs_runtime.Value {
	once_Control_Apply_apply__3302223884.Do(func() {
		cache_Control_Apply_apply__3302223884 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__3302223884(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__3302223884
}

var cache_Control_Apply_apply__2285104812 gopurs_runtime.Value
var once_Control_Apply_apply__2285104812 sync.Once
func Get_Control_Apply_apply__2285104812() gopurs_runtime.Value {
	once_Control_Apply_apply__2285104812.Do(func() {
		cache_Control_Apply_apply__2285104812 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2285104812(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2285104812
}

var cache_Control_Apply_apply__75278732 gopurs_runtime.Value
var once_Control_Apply_apply__75278732 sync.Once
func Get_Control_Apply_apply__75278732() gopurs_runtime.Value {
	once_Control_Apply_apply__75278732.Do(func() {
		cache_Control_Apply_apply__75278732 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__75278732(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__75278732
}

var cache_Control_Apply_apply__1941274444 gopurs_runtime.Value
var once_Control_Apply_apply__1941274444 sync.Once
func Get_Control_Apply_apply__1941274444() gopurs_runtime.Value {
	once_Control_Apply_apply__1941274444.Do(func() {
		cache_Control_Apply_apply__1941274444 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__1941274444(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__1941274444
}

var cache_Control_Apply_apply__2462165548 gopurs_runtime.Value
var once_Control_Apply_apply__2462165548 sync.Once
func Get_Control_Apply_apply__2462165548() gopurs_runtime.Value {
	once_Control_Apply_apply__2462165548.Do(func() {
		cache_Control_Apply_apply__2462165548 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2462165548(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2462165548
}

var cache_Control_Apply_apply__4071263818 gopurs_runtime.Value
var once_Control_Apply_apply__4071263818 sync.Once
func Get_Control_Apply_apply__4071263818() gopurs_runtime.Value {
	once_Control_Apply_apply__4071263818.Do(func() {
		cache_Control_Apply_apply__4071263818 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__4071263818(v_0_box, v1_1_box)
})
	})
	return cache_Control_Apply_apply__4071263818
}

var cache_Control_Apply_apply__2140510474 gopurs_runtime.Value
var once_Control_Apply_apply__2140510474 sync.Once
func Get_Control_Apply_apply__2140510474() gopurs_runtime.Value {
	once_Control_Apply_apply__2140510474.Do(func() {
		cache_Control_Apply_apply__2140510474 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2140510474(f_0_box, x_1_box)
})
	})
	return cache_Control_Apply_apply__2140510474
}

var cache_Control_Apply_apply__3620326986 gopurs_runtime.Value
var once_Control_Apply_apply__3620326986 sync.Once
func Get_Control_Apply_apply__3620326986() gopurs_runtime.Value {
	once_Control_Apply_apply__3620326986.Do(func() {
		cache_Control_Apply_apply__3620326986 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__3620326986(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Apply_apply__3620326986
}

var cache_Control_Apply_apply__1636083818 gopurs_runtime.Value
var once_Control_Apply_apply__1636083818 sync.Once
func Get_Control_Apply_apply__1636083818() gopurs_runtime.Value {
	once_Control_Apply_apply__1636083818.Do(func() {
		cache_Control_Apply_apply__1636083818 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__1636083818(f_0_box, x_1_box)
})
	})
	return cache_Control_Apply_apply__1636083818
}

var cache_Control_Apply_apply__2169384906 gopurs_runtime.Value
var once_Control_Apply_apply__2169384906 sync.Once
func Get_Control_Apply_apply__2169384906() gopurs_runtime.Value {
	once_Control_Apply_apply__2169384906.Do(func() {
		cache_Control_Apply_apply__2169384906 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__2169384906(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__2169384906
}

var cache_Control_Apply_apply__1471729482 gopurs_runtime.Value
var once_Control_Apply_apply__1471729482 sync.Once
func Get_Control_Apply_apply__1471729482() gopurs_runtime.Value {
	once_Control_Apply_apply__1471729482.Do(func() {
		cache_Control_Apply_apply__1471729482 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__1471729482(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__1471729482
}

var cache_Control_Apply_apply__420223018 gopurs_runtime.Value
var once_Control_Apply_apply__420223018 sync.Once
func Get_Control_Apply_apply__420223018() gopurs_runtime.Value {
	once_Control_Apply_apply__420223018.Do(func() {
		cache_Control_Apply_apply__420223018 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__420223018(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__420223018
}

var cache_Control_Apply_apply__3882563466 gopurs_runtime.Value
var once_Control_Apply_apply__3882563466 sync.Once
func Get_Control_Apply_apply__3882563466() gopurs_runtime.Value {
	once_Control_Apply_apply__3882563466.Do(func() {
		cache_Control_Apply_apply__3882563466 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__3882563466(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__3882563466
}

var cache_Control_Apply_apply__3867059818 gopurs_runtime.Value
var once_Control_Apply_apply__3867059818 sync.Once
func Get_Control_Apply_apply__3867059818() gopurs_runtime.Value {
	once_Control_Apply_apply__3867059818.Do(func() {
		cache_Control_Apply_apply__3867059818 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__3867059818(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__3867059818
}

var cache_Control_Apply_apply__3489442218 gopurs_runtime.Value
var once_Control_Apply_apply__3489442218 sync.Once
func Get_Control_Apply_apply__3489442218() gopurs_runtime.Value {
	once_Control_Apply_apply__3489442218.Do(func() {
		cache_Control_Apply_apply__3489442218 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__3489442218(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__3489442218
}

var cache_Control_Apply_apply__4016568778 gopurs_runtime.Value
var once_Control_Apply_apply__4016568778 sync.Once
func Get_Control_Apply_apply__4016568778() gopurs_runtime.Value {
	once_Control_Apply_apply__4016568778.Do(func() {
		cache_Control_Apply_apply__4016568778 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__4016568778(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__4016568778
}

var cache_Control_Apply_apply__724144906 gopurs_runtime.Value
var once_Control_Apply_apply__724144906 sync.Once
func Get_Control_Apply_apply__724144906() gopurs_runtime.Value {
	once_Control_Apply_apply__724144906.Do(func() {
		cache_Control_Apply_apply__724144906 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__724144906(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__724144906
}

var cache_Control_Apply_apply__2485508138 gopurs_runtime.Value
var once_Control_Apply_apply__2485508138 sync.Once
func Get_Control_Apply_apply__2485508138() gopurs_runtime.Value {
	once_Control_Apply_apply__2485508138.Do(func() {
		cache_Control_Apply_apply__2485508138 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__2485508138(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__2485508138
}

var cache_Control_Apply_apply__1572009162 gopurs_runtime.Value
var once_Control_Apply_apply__1572009162 sync.Once
func Get_Control_Apply_apply__1572009162() gopurs_runtime.Value {
	once_Control_Apply_apply__1572009162.Do(func() {
		cache_Control_Apply_apply__1572009162 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__1572009162(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__1572009162
}

var cache_Control_Apply_apply__1183285642 gopurs_runtime.Value
var once_Control_Apply_apply__1183285642 sync.Once
func Get_Control_Apply_apply__1183285642() gopurs_runtime.Value {
	once_Control_Apply_apply__1183285642.Do(func() {
		cache_Control_Apply_apply__1183285642 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__1183285642(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__1183285642
}

var cache_Control_Apply_apply__3534390890 gopurs_runtime.Value
var once_Control_Apply_apply__3534390890 sync.Once
func Get_Control_Apply_apply__3534390890() gopurs_runtime.Value {
	once_Control_Apply_apply__3534390890.Do(func() {
		cache_Control_Apply_apply__3534390890 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Apply_apply__3534390890(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v1_1_box)))}
})
	})
	return cache_Control_Apply_apply__3534390890
}

var cache_Control_Apply_apply__3993916842 gopurs_runtime.Value
var once_Control_Apply_apply__3993916842 sync.Once
func Get_Control_Apply_apply__3993916842() gopurs_runtime.Value {
	once_Control_Apply_apply__3993916842.Do(func() {
		cache_Control_Apply_apply__3993916842 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__3993916842(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Control_Apply_apply__3993916842
}

var cache_Control_Apply_apply__2467158700 gopurs_runtime.Value
var once_Control_Apply_apply__2467158700 sync.Once
func Get_Control_Apply_apply__2467158700() gopurs_runtime.Value {
	once_Control_Apply_apply__2467158700.Do(func() {
		cache_Control_Apply_apply__2467158700 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__2467158700(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__2467158700
}

var cache_Control_Apply_apply__3113603660 gopurs_runtime.Value
var once_Control_Apply_apply__3113603660 sync.Once
func Get_Control_Apply_apply__3113603660() gopurs_runtime.Value {
	once_Control_Apply_apply__3113603660.Do(func() {
		cache_Control_Apply_apply__3113603660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply__3113603660(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dict_0_box))
})
	})
	return cache_Control_Apply_apply__3113603660
}

var cache_Control_Apply_applyArray__2998472828 gopurs_runtime.Value
var once_Control_Apply_applyArray__2998472828 sync.Once
func Get_Control_Apply_applyArray__2998472828() gopurs_runtime.Value {
	once_Control_Apply_applyArray__2998472828.Do(func() {
		cache_Control_Apply_applyArray__2998472828 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorArray()
}), Get_Control_Apply_arrayApply())
	})
	return cache_Control_Apply_applyArray__2998472828
}

var cache_Control_Apply_applyFn__2722791663 gopurs_runtime.Value
var once_Control_Apply_applyFn__2722791663 sync.Once
func Get_Control_Apply_applyFn__2722791663() gopurs_runtime.Value {
	once_Control_Apply_applyFn__2722791663.Do(func() {
		cache_Control_Apply_applyFn__2722791663 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorFn()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_Control_Apply_applyFn__2722791663
}

var cache_Control_Apply_applyFn__4042184691 gopurs_runtime.Value
var once_Control_Apply_applyFn__4042184691 sync.Once
func Get_Control_Apply_applyFn__4042184691() gopurs_runtime.Value {
	once_Control_Apply_applyFn__4042184691.Do(func() {
		cache_Control_Apply_applyFn__4042184691 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorFn()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_Control_Apply_applyFn__4042184691
}

var cache_Control_Apply_applyProxy__2261709491 gopurs_runtime.Value
var once_Control_Apply_applyProxy__2261709491 sync.Once
func Get_Control_Apply_applyProxy__2261709491() gopurs_runtime.Value {
	once_Control_Apply_applyProxy__2261709491.Do(func() {
		cache_Control_Apply_applyProxy__2261709491 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_Control_Apply_applyProxy__2261709491
}

var cache_Control_Apply_applyProxy__315643445 gopurs_runtime.Value
var once_Control_Apply_applyProxy__315643445 sync.Once
func Get_Control_Apply_applyProxy__315643445() gopurs_runtime.Value {
	once_Control_Apply_applyProxy__315643445.Do(func() {
		cache_Control_Apply_applyProxy__315643445 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_Control_Apply_applyProxy__315643445
}

var cache_Control_Apply_applySecond__1627424644 gopurs_runtime.Value
var once_Control_Apply_applySecond__1627424644 sync.Once
func Get_Control_Apply_applySecond__1627424644() gopurs_runtime.Value {
	once_Control_Apply_applySecond__1627424644.Do(func() {
		cache_Control_Apply_applySecond__1627424644 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_applySecond__1627424644(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_applySecond__1627424644
}

var cache_Control_Apply_lift2__1942544886 gopurs_runtime.Value
var once_Control_Apply_lift2__1942544886 sync.Once
func Get_Control_Apply_lift2__1942544886() gopurs_runtime.Value {
	once_Control_Apply_lift2__1942544886.Do(func() {
		cache_Control_Apply_lift2__1942544886 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__1942544886(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__1942544886
}

var cache_Control_Apply_lift2__3139828374 gopurs_runtime.Value
var once_Control_Apply_lift2__3139828374 sync.Once
func Get_Control_Apply_lift2__3139828374() gopurs_runtime.Value {
	once_Control_Apply_lift2__3139828374.Do(func() {
		cache_Control_Apply_lift2__3139828374 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__3139828374(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__3139828374
}

var cache_Control_Apply_lift2__3684551766 gopurs_runtime.Value
var once_Control_Apply_lift2__3684551766 sync.Once
func Get_Control_Apply_lift2__3684551766() gopurs_runtime.Value {
	once_Control_Apply_lift2__3684551766.Do(func() {
		cache_Control_Apply_lift2__3684551766 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__3684551766(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__3684551766
}

var cache_Control_Apply_lift2__2286084406 gopurs_runtime.Value
var once_Control_Apply_lift2__2286084406 sync.Once
func Get_Control_Apply_lift2__2286084406() gopurs_runtime.Value {
	once_Control_Apply_lift2__2286084406.Do(func() {
		cache_Control_Apply_lift2__2286084406 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__2286084406(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__2286084406
}

var cache_Control_Apply_lift2__3007543670 gopurs_runtime.Value
var once_Control_Apply_lift2__3007543670 sync.Once
func Get_Control_Apply_lift2__3007543670() gopurs_runtime.Value {
	once_Control_Apply_lift2__3007543670.Do(func() {
		cache_Control_Apply_lift2__3007543670 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__3007543670(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__3007543670
}

var cache_Control_Apply_lift2__2762258480 gopurs_runtime.Value
var once_Control_Apply_lift2__2762258480 sync.Once
func Get_Control_Apply_lift2__2762258480() gopurs_runtime.Value {
	once_Control_Apply_lift2__2762258480.Do(func() {
		cache_Control_Apply_lift2__2762258480 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__2762258480(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__2762258480
}

var cache_Control_Apply_lift2__619377840 gopurs_runtime.Value
var once_Control_Apply_lift2__619377840 sync.Once
func Get_Control_Apply_lift2__619377840() gopurs_runtime.Value {
	once_Control_Apply_lift2__619377840.Do(func() {
		cache_Control_Apply_lift2__619377840 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__619377840(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__619377840
}

var cache_Control_Apply_lift2__2618178704 gopurs_runtime.Value
var once_Control_Apply_lift2__2618178704 sync.Once
func Get_Control_Apply_lift2__2618178704() gopurs_runtime.Value {
	once_Control_Apply_lift2__2618178704.Do(func() {
		cache_Control_Apply_lift2__2618178704 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__2618178704(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__2618178704
}

var cache_Control_Apply_lift2__1605401392 gopurs_runtime.Value
var once_Control_Apply_lift2__1605401392 sync.Once
func Get_Control_Apply_lift2__1605401392() gopurs_runtime.Value {
	once_Control_Apply_lift2__1605401392.Do(func() {
		cache_Control_Apply_lift2__1605401392 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__1605401392(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__1605401392
}

var cache_Control_Apply_lift2__3294332048 gopurs_runtime.Value
var once_Control_Apply_lift2__3294332048 sync.Once
func Get_Control_Apply_lift2__3294332048() gopurs_runtime.Value {
	once_Control_Apply_lift2__3294332048.Do(func() {
		cache_Control_Apply_lift2__3294332048 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__3294332048(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__3294332048
}

var cache_Control_Apply_lift2__2114580400 gopurs_runtime.Value
var once_Control_Apply_lift2__2114580400 sync.Once
func Get_Control_Apply_lift2__2114580400() gopurs_runtime.Value {
	once_Control_Apply_lift2__2114580400.Do(func() {
		cache_Control_Apply_lift2__2114580400 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__2114580400(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__2114580400
}

var cache_Control_Apply_lift2__2273022256 gopurs_runtime.Value
var once_Control_Apply_lift2__2273022256 sync.Once
func Get_Control_Apply_lift2__2273022256() gopurs_runtime.Value {
	once_Control_Apply_lift2__2273022256.Do(func() {
		cache_Control_Apply_lift2__2273022256 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__2273022256(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__2273022256
}

var cache_Control_Apply_lift2__3213187376 gopurs_runtime.Value
var once_Control_Apply_lift2__3213187376 sync.Once
func Get_Control_Apply_lift2__3213187376() gopurs_runtime.Value {
	once_Control_Apply_lift2__3213187376.Do(func() {
		cache_Control_Apply_lift2__3213187376 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__3213187376(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__3213187376
}

var cache_Control_Apply_lift2__1517638032 gopurs_runtime.Value
var once_Control_Apply_lift2__1517638032 sync.Once
func Get_Control_Apply_lift2__1517638032() gopurs_runtime.Value {
	once_Control_Apply_lift2__1517638032.Do(func() {
		cache_Control_Apply_lift2__1517638032 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__1517638032(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__1517638032
}

var cache_Control_Apply_lift2__470376976 gopurs_runtime.Value
var once_Control_Apply_lift2__470376976 sync.Once
func Get_Control_Apply_lift2__470376976() gopurs_runtime.Value {
	once_Control_Apply_lift2__470376976.Do(func() {
		cache_Control_Apply_lift2__470376976 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__470376976(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__470376976
}

var cache_Control_Apply_lift2__650234614 gopurs_runtime.Value
var once_Control_Apply_lift2__650234614 sync.Once
func Get_Control_Apply_lift2__650234614() gopurs_runtime.Value {
	once_Control_Apply_lift2__650234614.Do(func() {
		cache_Control_Apply_lift2__650234614 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__650234614(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Control_Apply_lift2__650234614
}

var cache_Control_Apply_lift2__1424073974 gopurs_runtime.Value
var once_Control_Apply_lift2__1424073974 sync.Once
func Get_Control_Apply_lift2__1424073974() gopurs_runtime.Value {
	once_Control_Apply_lift2__1424073974.Do(func() {
		cache_Control_Apply_lift2__1424073974 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__1424073974(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Control_Apply_lift2__1424073974
}

var cache_Control_Apply_lift2__2401097718 gopurs_runtime.Value
var once_Control_Apply_lift2__2401097718 sync.Once
func Get_Control_Apply_lift2__2401097718() gopurs_runtime.Value {
	once_Control_Apply_lift2__2401097718.Do(func() {
		cache_Control_Apply_lift2__2401097718 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__2401097718(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Control_Apply_lift2__2401097718
}

var cache_Control_Apply_lift2__1465919478 gopurs_runtime.Value
var once_Control_Apply_lift2__1465919478 sync.Once
func Get_Control_Apply_lift2__1465919478() gopurs_runtime.Value {
	once_Control_Apply_lift2__1465919478.Do(func() {
		cache_Control_Apply_lift2__1465919478 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__1465919478(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Control_Apply_lift2__1465919478
}

var cache_Control_Apply_lift2__3315261616 gopurs_runtime.Value
var once_Control_Apply_lift2__3315261616 sync.Once
func Get_Control_Apply_lift2__3315261616() gopurs_runtime.Value {
	once_Control_Apply_lift2__3315261616.Do(func() {
		cache_Control_Apply_lift2__3315261616 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__3315261616(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__3315261616
}

var cache_Control_Apply_lift2__1699116848 gopurs_runtime.Value
var once_Control_Apply_lift2__1699116848 sync.Once
func Get_Control_Apply_lift2__1699116848() gopurs_runtime.Value {
	once_Control_Apply_lift2__1699116848.Do(func() {
		cache_Control_Apply_lift2__1699116848 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2__1699116848(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2__1699116848
}

type Constructor_Control_Apply_Apply struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3032403085] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Apply_Apply)(ptr)
		_ = c
		switch key {
		case "Functor0": return gopurs_runtime.Box(c.V0)
		case "apply": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Apply_Apply: " + key)
		}
	}
}


func Call_Control_Apply_Apply_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Apply_apply(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_applyFirst(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Function_go__const(), a_2), b_3)
})
})
}

func Call_Control_Apply_applySecond(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
})
}), a_2), b_3)
})
})
}

func Call_Control_Apply_lift2(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift3(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4), c_5)
})
})
})
})
}

func Call_Control_Apply_lift4(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4), c_5), d_6)
})
})
})
})
})
}

func Call_Control_Apply_lift5(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4), c_5), d_6), e_7)
})
})
})
})
})
})
}

func Call_Control_Apply_apply__1159579293(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__1030762512(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2174094052(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__652936900(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__3477321027(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__350765120(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__5413771(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__4203183626(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__1982519530(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2962221386(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__4140882028(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2008680204(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2092855244(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__353515660(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2080233356(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__197474060(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__1526260460(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__1023214444(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2623550860(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__4046394764(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2021897708(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2908293516(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__3783667596(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__3582447820(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__1099902892(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__1851858028(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__1471954540(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__3063957068(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2007181164(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__986161100(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2947753132(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2234240780(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2758496428(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2087590060(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__3302223884(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2285104812(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__75278732(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__1941274444(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__2462165548(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__4071263818(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(Get_Data_Either_Left(), (*Constructor_Data_Either_Left)(v_0.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), (*Constructor_Data_Either_Right)(v_0.UnsafePtr).V0, v1_1)
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

func Call_Control_Apply_apply__2140510474(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))
}))
}

func Call_Control_Apply_apply__3620326986(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply3(Get_Control_Monad_ap(), Get_Data_List_Lazy_Types_monadList(), __eta0_0, __eta1_1)
}

func Call_Control_Apply_apply__1636083818(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))
}))
}

func Call_Control_Apply_apply__2169384906(v_0_loop *Constructor_Data_List_Types_Cons, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
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
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorList(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})))}))
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

func Call_Control_Apply_apply__1471729482(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}))
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

func Call_Control_Apply_apply__420223018(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}))
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

func Call_Control_Apply_apply__3882563466(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}))
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

func Call_Control_Apply_apply__3867059818(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}))
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

func Call_Control_Apply_apply__3489442218(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}))
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

func Call_Control_Apply_apply__4016568778(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}))
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

func Call_Control_Apply_apply__724144906(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}))
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

func Call_Control_Apply_apply__2485508138(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}))
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

func Call_Control_Apply_apply__1572009162(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}))
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

func Call_Control_Apply_apply__1183285642(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}))
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

func Call_Control_Apply_apply__3534390890(v_0_loop *Constructor_Data_Maybe_Just, v1_1_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_Maybe_Just = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}))
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

func Call_Control_Apply_apply__3993916842(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply3(Get_Control_Monad_ap(), Get_Effect_Aff_monadAff(), __eta0_0, __eta1_1)
}

func Call_Control_Apply_apply__2467158700(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_apply__3113603660(dict_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Apply_applySecond__1627424644(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
})
}), a_2), b_3)
})
})
}

func Call_Control_Apply_lift2__1942544886(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__3139828374(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__3684551766(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__2286084406(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__3007543670(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__2762258480(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__619377840(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__2618178704(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__1605401392(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__3294332048(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__2114580400(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__2273022256(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__3213187376(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__1517638032(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__470376976(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__650234614(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applyST(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), __eta0_0, __eta1_1), __eta2_2)
}

func Call_Control_Apply_lift2__1424073974(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_applyAff(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_functorAff(), "map"), __eta0_0, __eta1_1), __eta2_2)
}

func Call_Control_Apply_lift2__2401097718(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_applyParAff(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_Aff_functorParAff(), "map"), __eta0_0, __eta1_1), __eta2_2)
}

func Call_Control_Apply_lift2__1465919478(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_applyEffect(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Effect_functorEffect(), "map"), __eta0_0, __eta1_1), __eta2_2)
}

func Call_Control_Apply_lift2__3315261616(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Call_Control_Apply_lift2__1699116848(dictApply_0_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_2, a_3), b_4)
})
})
})
}

func Get_Control_Apply_arrayApply() gopurs_runtime.Value {
	return _Gopurs_Control_Apply_ArrayApply
}
