package Control_Biapplicative

import (
	pkg_Control_Biapply "gopurs/output/Control.Biapply"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_bipure gopurs_runtime.Value
var once_bipure sync.Once
func Get_bipure() gopurs_runtime.Value {
	once_bipure.Do(func() {
		cache_bipure = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bipure(gopurs_runtime.CoerceToStruct[Constructor_Biapplicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bipure
}

var cache_biapplicativeTuple gopurs_runtime.Value
var once_biapplicativeTuple sync.Once
func Get_biapplicativeTuple() gopurs_runtime.Value {
	once_biapplicativeTuple.Do(func() {
		cache_biapplicativeTuple = gopurs_runtime.RecordDict2("Biapply0", "bipure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Biapply.Get_biapplyTuple()
}), pkg_Data_Tuple.Get_Tuple())
	})
	return cache_biapplicativeTuple
}

var cache_biapplyTuple__355763440 gopurs_runtime.Value
var once_biapplyTuple__355763440 sync.Once
func Get_biapplyTuple__355763440() gopurs_runtime.Value {
	once_biapplyTuple__355763440.Do(func() {
		cache_biapplyTuple__355763440 = gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorTuple()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_biapplyTuple__355763440
}

var cache_bifunctorTuple__553376860 gopurs_runtime.Value
var once_bifunctorTuple__553376860 sync.Once
func Get_bifunctorTuple__553376860() gopurs_runtime.Value {
	once_bifunctorTuple__553376860.Do(func() {
		cache_bifunctorTuple__553376860 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
})
})
}))
	})
	return cache_bifunctorTuple__553376860
}

type Constructor_Biapplicative[T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3949191309] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Biapplicative[gopurs_runtime.Value])(ptr)
		switch key {
		case "Biapply0": return c.V0
		case "bipure": return c.V1
		default: panic("Key not found in dictionary Constructor_Biapplicative: " + key)
		}
	}
}


func Call_bipure(dict_0_loop *Constructor_Biapplicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Biapplicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


