package Data_Functor

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_go__map gopurs_runtime.Value
var once_go__map sync.Once
func Get_go__map() gopurs_runtime.Value {
	once_go__map.Do(func() {
		cache_go__map = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_go__map(gopurs_runtime.CoerceToStruct[Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_go__map
}

var cache_mapFlipped gopurs_runtime.Value
var once_mapFlipped sync.Once
func Get_mapFlipped() gopurs_runtime.Value {
	once_mapFlipped.Do(func() {
		cache_mapFlipped = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped(gopurs_runtime.CoerceToStruct[Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped
}

var cache_void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		cache_void = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_void(gopurs_runtime.CoerceToStruct[Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box))
})
	})
	return cache_void
}

var cache_voidLeft gopurs_runtime.Value
var once_voidLeft sync.Once
func Get_voidLeft() gopurs_runtime.Value {
	once_voidLeft.Do(func() {
		cache_voidLeft = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidLeft(gopurs_runtime.CoerceToStruct[Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, x_2_box)
})
	})
	return cache_voidLeft
}

var cache_voidRight gopurs_runtime.Value
var once_voidRight sync.Once
func Get_voidRight() gopurs_runtime.Value {
	once_voidRight.Do(func() {
		cache_voidRight = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidRight(gopurs_runtime.CoerceToStruct[Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), x_1_box)
})
	})
	return cache_voidRight
}

var cache_functorProxy gopurs_runtime.Value
var once_functorProxy sync.Once
func Get_functorProxy() gopurs_runtime.Value {
	once_functorProxy.Do(func() {
		cache_functorProxy = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_functorProxy
}

var cache_functorFn gopurs_runtime.Value
var once_functorFn sync.Once
func Get_functorFn() gopurs_runtime.Value {
	once_functorFn.Do(func() {
		cache_functorFn = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))
	})
	return cache_functorFn
}

var cache_functorArray gopurs_runtime.Value
var once_functorArray sync.Once
func Get_functorArray() gopurs_runtime.Value {
	once_functorArray.Do(func() {
		cache_functorArray = gopurs_runtime.RecordDict1("map", Get_arrayMap())
	})
	return cache_functorArray
}

var cache_flap gopurs_runtime.Value
var once_flap sync.Once
func Get_flap() gopurs_runtime.Value {
	once_flap.Do(func() {
		cache_flap = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, ff_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flap(gopurs_runtime.CoerceToStruct[Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), ff_1_box, x_2_box)
})
	})
	return cache_flap
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

var cache_map__3240628980 gopurs_runtime.Value
var once_map__3240628980 sync.Once
func Get_map__3240628980() gopurs_runtime.Value {
	once_map__3240628980.Do(func() {
		cache_map__3240628980 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3240628980(gopurs_runtime.CoerceToStruct[Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3240628980
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__1938733460 gopurs_runtime.Value
var once_map__1938733460 sync.Once
func Get_map__1938733460() gopurs_runtime.Value {
	once_map__1938733460.Do(func() {
		cache_map__1938733460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1938733460(gopurs_runtime.CoerceToStruct[Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1938733460
}

type Constructor_Functor[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[929368378] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Functor[gopurs_runtime.Value])(ptr)
		switch key {
		case "map": return c.V0
		default: panic("Key not found in dictionary Constructor_Functor: " + key)
		}
	}
}


func Call_go__map(dict_0_loop *Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mapFlipped(dictFunctor_0_loop *Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_void(dictFunctor_0_loop *Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
}

func Call_voidLeft(dictFunctor_0_loop *Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), f_1)
}

func Call_voidRight(dictFunctor_0_loop *Constructor_Functor[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_flap(dictFunctor_0_loop *Constructor_Functor[gopurs_runtime.Value], ff_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var ff_1 gopurs_runtime.Value = ff_1_loop
_ = ff_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, x_2)
}), ff_1)
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

func Call_map__3240628980(dict_0_loop *Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1938733460(dict_0_loop *Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get_arrayMap() gopurs_runtime.Value {
	return _Gopurs_ArrayMap
}
