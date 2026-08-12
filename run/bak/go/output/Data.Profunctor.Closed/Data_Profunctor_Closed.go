package Data_Profunctor_Closed

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_closedFunction gopurs_runtime.Value
var once_closedFunction sync.Once
func Get_closedFunction() gopurs_runtime.Value {
	once_closedFunction.Do(func() {
		cache_closedFunction = gopurs_runtime.RecordDict2("Profunctor0", "closed", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Profunctor.Get_profunctorFn()
}), gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))
	})
	return cache_closedFunction
}

var cache_closed gopurs_runtime.Value
var once_closed sync.Once
func Get_closed() gopurs_runtime.Value {
	once_closed.Do(func() {
		cache_closed = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_closed(gopurs_runtime.CoerceToStruct[Constructor_Closed[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_closed
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_composeFlipped__2583068543 gopurs_runtime.Value
var once_composeFlipped__2583068543 sync.Once
func Get_composeFlipped__2583068543() gopurs_runtime.Value {
	once_composeFlipped__2583068543.Do(func() {
		cache_composeFlipped__2583068543 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__2583068543(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__2583068543
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_profunctorFn__3736629211 gopurs_runtime.Value
var once_profunctorFn__3736629211 sync.Once
func Get_profunctorFn__3736629211() gopurs_runtime.Value {
	once_profunctorFn__3736629211.Do(func() {
		cache_profunctorFn__3736629211 = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c2d_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c2d_1, gopurs_runtime.Apply(b2c_2, gopurs_runtime.Apply(a2b_0, x_3)))
})
})
})
}))
	})
	return cache_profunctorFn__3736629211
}

type Constructor_Closed[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[768764671] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Closed[gopurs_runtime.Value])(ptr)
		switch key {
		case "Profunctor0": return c.V0
		case "closed": return c.V1
		default: panic("Key not found in dictionary Constructor_Closed: " + key)
		}
	}
}


func Call_closed(dict_0_loop *Constructor_Closed[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Closed[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_composeFlipped__2583068543(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, g_2, f_1)
}


