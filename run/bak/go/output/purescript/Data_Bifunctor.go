package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Bifunctor_identity gopurs_runtime.Value
var once_Data_Bifunctor_identity sync.Once
func Get_Data_Bifunctor_identity() gopurs_runtime.Value {
	once_Data_Bifunctor_identity.Do(func() {
		cache_Data_Bifunctor_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_identity(x_0_box)
})
	})
	return cache_Data_Bifunctor_identity
}

var cache_Data_Bifunctor_identity1 gopurs_runtime.Value
var once_Data_Bifunctor_identity1 sync.Once
func Get_Data_Bifunctor_identity1() gopurs_runtime.Value {
	once_Data_Bifunctor_identity1.Do(func() {
		cache_Data_Bifunctor_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_identity1(x_0_box)
})
	})
	return cache_Data_Bifunctor_identity1
}

var cache_Data_Bifunctor_Bifunctor_dollarDict gopurs_runtime.Value
var once_Data_Bifunctor_Bifunctor_dollarDict sync.Once
func Get_Data_Bifunctor_Bifunctor_dollarDict() gopurs_runtime.Value {
	once_Data_Bifunctor_Bifunctor_dollarDict.Do(func() {
		cache_Data_Bifunctor_Bifunctor_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_Bifunctor_dollarDict(x_0_box)
})
	})
	return cache_Data_Bifunctor_Bifunctor_dollarDict
}

var cache_Data_Bifunctor_bimap gopurs_runtime.Value
var once_Data_Bifunctor_bimap sync.Once
func Get_Data_Bifunctor_bimap() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap.Do(func() {
		cache_Data_Bifunctor_bimap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](dict_0_box))
})
	})
	return cache_Data_Bifunctor_bimap
}

var cache_Data_Bifunctor_bivoid gopurs_runtime.Value
var once_Data_Bifunctor_bivoid sync.Once
func Get_Data_Bifunctor_bivoid() gopurs_runtime.Value {
	once_Data_Bifunctor_bivoid.Do(func() {
		cache_Data_Bifunctor_bivoid = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bivoid(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](dictBifunctor_0_box))
})
	})
	return cache_Data_Bifunctor_bivoid
}

var cache_Data_Bifunctor_lmap gopurs_runtime.Value
var once_Data_Bifunctor_lmap sync.Once
func Get_Data_Bifunctor_lmap() gopurs_runtime.Value {
	once_Data_Bifunctor_lmap.Do(func() {
		cache_Data_Bifunctor_lmap = gopurs_runtime.Func2(func(dictBifunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_lmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](dictBifunctor_0_box), f_1_box)
})
	})
	return cache_Data_Bifunctor_lmap
}

var cache_Data_Bifunctor_rmap gopurs_runtime.Value
var once_Data_Bifunctor_rmap sync.Once
func Get_Data_Bifunctor_rmap() gopurs_runtime.Value {
	once_Data_Bifunctor_rmap.Do(func() {
		cache_Data_Bifunctor_rmap = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_rmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](dictBifunctor_0_box))
})
	})
	return cache_Data_Bifunctor_rmap
}

var cache_Data_Bifunctor_bifunctorTuple gopurs_runtime.Value
var once_Data_Bifunctor_bifunctorTuple sync.Once
func Get_Data_Bifunctor_bifunctorTuple() gopurs_runtime.Value {
	once_Data_Bifunctor_bifunctorTuple.Do(func() {
		cache_Data_Bifunctor_bifunctorTuple = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)})}
})
})
}))
	})
	return cache_Data_Bifunctor_bifunctorTuple
}

var cache_Data_Bifunctor_bifunctorEither gopurs_runtime.Value
var once_Data_Bifunctor_bifunctorEither sync.Once
func Get_Data_Bifunctor_bifunctorEither() gopurs_runtime.Value {
	once_Data_Bifunctor_bifunctorEither.Do(func() {
		cache_Data_Bifunctor_bifunctorEither = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})
}))
	})
	return cache_Data_Bifunctor_bifunctorEither
}

var cache_Data_Bifunctor_bifunctorConst gopurs_runtime.Value
var once_Data_Bifunctor_bifunctorConst sync.Once
func Get_Data_Bifunctor_bifunctorConst() gopurs_runtime.Value {
	once_Data_Bifunctor_bifunctorConst.Do(func() {
		cache_Data_Bifunctor_bifunctorConst = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_Data_Bifunctor_bifunctorConst
}

var cache_Data_Bifunctor_bifunctorConst__3156923452 gopurs_runtime.Value
var once_Data_Bifunctor_bifunctorConst__3156923452 sync.Once
func Get_Data_Bifunctor_bifunctorConst__3156923452() gopurs_runtime.Value {
	once_Data_Bifunctor_bifunctorConst__3156923452.Do(func() {
		cache_Data_Bifunctor_bifunctorConst__3156923452 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_Data_Bifunctor_bifunctorConst__3156923452
}

var cache_Data_Bifunctor_bifunctorEither__3558063994 gopurs_runtime.Value
var once_Data_Bifunctor_bifunctorEither__3558063994 sync.Once
func Get_Data_Bifunctor_bifunctorEither__3558063994() gopurs_runtime.Value {
	once_Data_Bifunctor_bifunctorEither__3558063994.Do(func() {
		cache_Data_Bifunctor_bifunctorEither__3558063994 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})
}))
	})
	return cache_Data_Bifunctor_bifunctorEither__3558063994
}

var cache_Data_Bifunctor_bifunctorEither__1585706332 gopurs_runtime.Value
var once_Data_Bifunctor_bifunctorEither__1585706332 sync.Once
func Get_Data_Bifunctor_bifunctorEither__1585706332() gopurs_runtime.Value {
	once_Data_Bifunctor_bifunctorEither__1585706332.Do(func() {
		cache_Data_Bifunctor_bifunctorEither__1585706332 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})
}))
	})
	return cache_Data_Bifunctor_bifunctorEither__1585706332
}

var cache_Data_Bifunctor_bifunctorTuple__3421321530 gopurs_runtime.Value
var once_Data_Bifunctor_bifunctorTuple__3421321530 sync.Once
func Get_Data_Bifunctor_bifunctorTuple__3421321530() gopurs_runtime.Value {
	once_Data_Bifunctor_bifunctorTuple__3421321530.Do(func() {
		cache_Data_Bifunctor_bifunctorTuple__3421321530 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)})}
})
})
}))
	})
	return cache_Data_Bifunctor_bifunctorTuple__3421321530
}

var cache_Data_Bifunctor_bifunctorTuple__553376860 gopurs_runtime.Value
var once_Data_Bifunctor_bifunctorTuple__553376860 sync.Once
func Get_Data_Bifunctor_bifunctorTuple__553376860() gopurs_runtime.Value {
	once_Data_Bifunctor_bifunctorTuple__553376860.Do(func() {
		cache_Data_Bifunctor_bifunctorTuple__553376860 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)})}
})
})
}))
	})
	return cache_Data_Bifunctor_bifunctorTuple__553376860
}

var cache_Data_Bifunctor_bimap__4044928099 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__4044928099 sync.Once
func Get_Data_Bifunctor_bimap__4044928099() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__4044928099.Do(func() {
		cache_Data_Bifunctor_bimap__4044928099 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap__4044928099(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](dict_0_box))
})
	})
	return cache_Data_Bifunctor_bimap__4044928099
}

var cache_Data_Bifunctor_bimap__3304410354 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__3304410354 sync.Once
func Get_Data_Bifunctor_bimap__3304410354() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__3304410354.Do(func() {
		cache_Data_Bifunctor_bimap__3304410354 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap__3304410354(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](dict_0_box))
})
	})
	return cache_Data_Bifunctor_bimap__3304410354
}

var cache_Data_Bifunctor_bimap__132457202 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__132457202 sync.Once
func Get_Data_Bifunctor_bimap__132457202() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__132457202.Do(func() {
		cache_Data_Bifunctor_bimap__132457202 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap__132457202(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](dict_0_box))
})
	})
	return cache_Data_Bifunctor_bimap__132457202
}

var cache_Data_Bifunctor_bimap__2848069618 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__2848069618 sync.Once
func Get_Data_Bifunctor_bimap__2848069618() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__2848069618.Do(func() {
		cache_Data_Bifunctor_bimap__2848069618 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap__2848069618(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Bifunctor_bimap__2848069618
}

var cache_Data_Bifunctor_bimap__4141710674 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__4141710674 sync.Once
func Get_Data_Bifunctor_bimap__4141710674() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__4141710674.Do(func() {
		cache_Data_Bifunctor_bimap__4141710674 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap__4141710674(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](dict_0_box))
})
	})
	return cache_Data_Bifunctor_bimap__4141710674
}

var cache_Data_Bifunctor_bimap__3643630450 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__3643630450 sync.Once
func Get_Data_Bifunctor_bimap__3643630450() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__3643630450.Do(func() {
		cache_Data_Bifunctor_bimap__3643630450 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap__3643630450(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](dict_0_box))
})
	})
	return cache_Data_Bifunctor_bimap__3643630450
}

var cache_Data_Bifunctor_bimap__861179602 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__861179602 sync.Once
func Get_Data_Bifunctor_bimap__861179602() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__861179602.Do(func() {
		cache_Data_Bifunctor_bimap__861179602 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap__861179602(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](dict_0_box))
})
	})
	return cache_Data_Bifunctor_bimap__861179602
}

var cache_Data_Bifunctor_bimap__2801350668 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__2801350668 sync.Once
func Get_Data_Bifunctor_bimap__2801350668() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__2801350668.Do(func() {
		cache_Data_Bifunctor_bimap__2801350668 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap__2801350668(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Bifunctor_bimap__2801350668
}

var cache_Data_Bifunctor_bimap__1783967194 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__1783967194 sync.Once
func Get_Data_Bifunctor_bimap__1783967194() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__1783967194.Do(func() {
		cache_Data_Bifunctor_bimap__1783967194 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap__1783967194(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Bifunctor_bimap__1783967194
}

var cache_Data_Bifunctor_bimap__1727657434 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__1727657434 sync.Once
func Get_Data_Bifunctor_bimap__1727657434() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__1727657434.Do(func() {
		cache_Data_Bifunctor_bimap__1727657434 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_bimap__1727657434(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Bifunctor_bimap__1727657434
}

var cache_Data_Bifunctor_bimap__214753306 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__214753306 sync.Once
func Get_Data_Bifunctor_bimap__214753306() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__214753306.Do(func() {
		cache_Data_Bifunctor_bimap__214753306 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Bifunctor_bimap__214753306(f_0_box, g_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_2_box)))}
})
	})
	return cache_Data_Bifunctor_bimap__214753306
}

var cache_Data_Bifunctor_bimap__298925978 gopurs_runtime.Value
var once_Data_Bifunctor_bimap__298925978 sync.Once
func Get_Data_Bifunctor_bimap__298925978() gopurs_runtime.Value {
	once_Data_Bifunctor_bimap__298925978.Do(func() {
		cache_Data_Bifunctor_bimap__298925978 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Bifunctor_bimap__298925978(f_0_box, g_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_2_box)))}
})
	})
	return cache_Data_Bifunctor_bimap__298925978
}

var cache_Data_Bifunctor_lmap__2196160232 gopurs_runtime.Value
var once_Data_Bifunctor_lmap__2196160232 sync.Once
func Get_Data_Bifunctor_lmap__2196160232() gopurs_runtime.Value {
	once_Data_Bifunctor_lmap__2196160232.Do(func() {
		cache_Data_Bifunctor_lmap__2196160232 = gopurs_runtime.Func2(func(dictBifunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifunctor_lmap__2196160232(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](dictBifunctor_0_box), f_1_box)
})
	})
	return cache_Data_Bifunctor_lmap__2196160232
}

type Constructor_Data_Bifunctor_Bifunctor struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4141114362] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bifunctor_Bifunctor)(ptr)
		_ = c
		switch key {
		case "bimap": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Bifunctor_Bifunctor: " + key)
		}
	}
}


func Call_Data_Bifunctor_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bifunctor_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bifunctor_Bifunctor_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bifunctor_bimap(dict_0_loop *Constructor_Data_Bifunctor_Bifunctor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifunctor_Bifunctor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifunctor_bivoid(dictBifunctor_0_loop *Constructor_Data_Bifunctor_Bifunctor) gopurs_runtime.Value {
var dictBifunctor_0 *Constructor_Data_Bifunctor_Bifunctor = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBifunctor_0.V0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
}

func Call_Data_Bifunctor_lmap(dictBifunctor_0_loop *Constructor_Data_Bifunctor_Bifunctor, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 *Constructor_Data_Bifunctor_Bifunctor = dictBifunctor_0_loop
_ = dictBifunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBifunctor_0.V0), f_1, Get_Data_Bifunctor_identity())
}

func Call_Data_Bifunctor_rmap(dictBifunctor_0_loop *Constructor_Data_Bifunctor_Bifunctor) gopurs_runtime.Value {
var dictBifunctor_0 *Constructor_Data_Bifunctor_Bifunctor = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictBifunctor_0.V0), Get_Data_Bifunctor_identity1())
}

func Call_Data_Bifunctor_bimap__4044928099(dict_0_loop *Constructor_Data_Bifunctor_Bifunctor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifunctor_Bifunctor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifunctor_bimap__3304410354(dict_0_loop *Constructor_Data_Bifunctor_Bifunctor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifunctor_Bifunctor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifunctor_bimap__132457202(dict_0_loop *Constructor_Data_Bifunctor_Bifunctor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifunctor_Bifunctor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifunctor_bimap__2848069618(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 525585346) {
__t0 = gopurs_runtime.Apply(Get_Control_Monad_Rec_Class_Loop(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v2_2.UnsafePtr).V0)))})
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 60402430) {
__t0 = gopurs_runtime.Apply(Get_Control_Monad_Rec_Class_Done(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_1, (*Constructor_Control_Monad_Rec_Class_Done)(v2_2.UnsafePtr).V0)))})
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

func Call_Data_Bifunctor_bimap__4141710674(dict_0_loop *Constructor_Data_Bifunctor_Bifunctor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifunctor_Bifunctor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifunctor_bimap__3643630450(dict_0_loop *Constructor_Data_Bifunctor_Bifunctor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifunctor_Bifunctor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifunctor_bimap__861179602(dict_0_loop *Constructor_Data_Bifunctor_Bifunctor) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifunctor_Bifunctor = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifunctor_bimap__2801350668(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 525585346) {
__t0 = gopurs_runtime.Apply(Get_Control_Monad_Rec_Class_Loop(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v_0, (*Constructor_Control_Monad_Rec_Class_Loop)(v2_2.UnsafePtr).V0)))})
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 60402430) {
__t0 = gopurs_runtime.Apply(Get_Control_Monad_Rec_Class_Done(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_1, (*Constructor_Control_Monad_Rec_Class_Done)(v2_2.UnsafePtr).V0)))})
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

func Call_Data_Bifunctor_bimap__1783967194(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right)(v2_2.UnsafePtr).V0)})}
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

func Call_Data_Bifunctor_bimap__1727657434(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Apply2(Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V1))
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Apply2(Get_Data_Interval_DurationEnd(), gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V1))
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Apply2(Get_Data_Interval_StartDuration(), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V1))
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Apply(Get_Data_Interval_DurationOnly(), gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationOnly)(v2_2.UnsafePtr).V0))
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

func Call_Data_Bifunctor_bimap__214753306(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var v_2 *Constructor_Data_Tuple_Tuple = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((v_2).V0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(g_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((v_2).V1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
}

func Call_Data_Bifunctor_bimap__298925978(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var v_2 *Constructor_Data_Tuple_Tuple = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_0, (v_2).V0), gopurs_runtime.Apply(g_1, (v_2).V1)})})
}

func Call_Data_Bifunctor_lmap__2196160232(dictBifunctor_0_loop *Constructor_Data_Bifunctor_Bifunctor, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 *Constructor_Data_Bifunctor_Bifunctor = dictBifunctor_0_loop
_ = dictBifunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBifunctor_0.V0), f_1, Get_Data_Bifunctor_identity())
}


