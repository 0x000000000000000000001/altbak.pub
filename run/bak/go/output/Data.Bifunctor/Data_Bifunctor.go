package Data_Bifunctor

import (
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_identity1 gopurs_runtime.Value
var once_identity1 sync.Once
func Get_identity1() gopurs_runtime.Value {
	once_identity1.Do(func() {
		cache_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity1(x_0_box)
})
	})
	return cache_identity1
}

var cache_bimap gopurs_runtime.Value
var once_bimap sync.Once
func Get_bimap() gopurs_runtime.Value {
	once_bimap.Do(func() {
		cache_bimap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap(gopurs_runtime.CoerceToStruct[Constructor_Bifunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bimap
}

var cache_bivoid gopurs_runtime.Value
var once_bivoid sync.Once
func Get_bivoid() gopurs_runtime.Value {
	once_bivoid.Do(func() {
		cache_bivoid = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bivoid(gopurs_runtime.CoerceToStruct[Constructor_Bifunctor[gopurs_runtime.Value]](dictBifunctor_0_box))
})
	})
	return cache_bivoid
}

var cache_lmap gopurs_runtime.Value
var once_lmap sync.Once
func Get_lmap() gopurs_runtime.Value {
	once_lmap.Do(func() {
		cache_lmap = gopurs_runtime.Func2(func(dictBifunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lmap(gopurs_runtime.CoerceToStruct[Constructor_Bifunctor[gopurs_runtime.Value]](dictBifunctor_0_box), f_1_box)
})
	})
	return cache_lmap
}

var cache_rmap gopurs_runtime.Value
var once_rmap sync.Once
func Get_rmap() gopurs_runtime.Value {
	once_rmap.Do(func() {
		cache_rmap = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rmap(gopurs_runtime.CoerceToStruct[Constructor_Bifunctor[gopurs_runtime.Value]](dictBifunctor_0_box))
})
	})
	return cache_rmap
}

var cache_bifunctorTuple gopurs_runtime.Value
var once_bifunctorTuple sync.Once
func Get_bifunctorTuple() gopurs_runtime.Value {
	once_bifunctorTuple.Do(func() {
		cache_bifunctorTuple = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
})
})
}))
	})
	return cache_bifunctorTuple
}

var cache_bifunctorEither gopurs_runtime.Value
var once_bifunctorEither sync.Once
func Get_bifunctorEither() gopurs_runtime.Value {
	once_bifunctorEither.Do(func() {
		cache_bifunctorEither = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
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
	return cache_bifunctorEither
}

var cache_bifunctorConst gopurs_runtime.Value
var once_bifunctorConst sync.Once
func Get_bifunctorConst() gopurs_runtime.Value {
	once_bifunctorConst.Do(func() {
		cache_bifunctorConst = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_bifunctorConst
}

var cache_bimap__3304410354 gopurs_runtime.Value
var once_bimap__3304410354 sync.Once
func Get_bimap__3304410354() gopurs_runtime.Value {
	once_bimap__3304410354.Do(func() {
		cache_bimap__3304410354 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__3304410354(gopurs_runtime.CoerceToStruct[Constructor_Bifunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bimap__3304410354
}

var cache_bimap__132457202 gopurs_runtime.Value
var once_bimap__132457202 sync.Once
func Get_bimap__132457202() gopurs_runtime.Value {
	once_bimap__132457202.Do(func() {
		cache_bimap__132457202 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__132457202(gopurs_runtime.CoerceToStruct[Constructor_Bifunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bimap__132457202
}

var cache_const__4026847508 gopurs_runtime.Value
var once_const__4026847508 sync.Once
func Get_const__4026847508() gopurs_runtime.Value {
	once_const__4026847508.Do(func() {
		cache_const__4026847508 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__4026847508(a_0_box, v_1_box)
})
	})
	return cache_const__4026847508
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

type Constructor_Bifunctor[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4141114362] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Bifunctor[gopurs_runtime.Value])(ptr)
		switch key {
		case "bimap": return c.V0
		default: panic("Key not found in dictionary Constructor_Bifunctor: " + key)
		}
	}
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_bimap(dict_0_loop *Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bivoid(dictBifunctor_0_loop *Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifunctor_0 *Constructor_Bifunctor[gopurs_runtime.Value] = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.Apply2(dictBifunctor_0.V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
}

func Call_lmap(dictBifunctor_0_loop *Constructor_Bifunctor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 *Constructor_Bifunctor[gopurs_runtime.Value] = dictBifunctor_0_loop
_ = dictBifunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(dictBifunctor_0.V0, f_1, Get_identity())
}

func Call_rmap(dictBifunctor_0_loop *Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifunctor_0 *Constructor_Bifunctor[gopurs_runtime.Value] = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.Apply(dictBifunctor_0.V0, Get_identity1())
}

func Call_bimap__3304410354(dict_0_loop *Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bimap__132457202(dict_0_loop *Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_const__4026847508(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}


