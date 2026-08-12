package Control_Alternative

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Plus "gopurs/output/Control.Plus"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		cache_guard = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_guard(gopurs_runtime.CoerceToStruct[Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_guard
}

var cache_alternativeArray gopurs_runtime.Value
var once_alternativeArray sync.Once
func Get_alternativeArray() gopurs_runtime.Value {
	once_alternativeArray.Do(func() {
		cache_alternativeArray = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Plus.Get_plusArray()
}))
	})
	return cache_alternativeArray
}

var cache_altArray__2010533188 gopurs_runtime.Value
var once_altArray__2010533188 sync.Once
func Get_altArray__2010533188() gopurs_runtime.Value {
	once_altArray__2010533188.Do(func() {
		cache_altArray__2010533188 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"))
	})
	return cache_altArray__2010533188
}

var cache_applicativeArray__1604836744 gopurs_runtime.Value
var once_applicativeArray__1604836744 sync.Once
func Get_applicativeArray__1604836744() gopurs_runtime.Value {
	once_applicativeArray__1604836744.Do(func() {
		cache_applicativeArray__1604836744 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
}))
	})
	return cache_applicativeArray__1604836744
}

var cache_pure__2935994064 gopurs_runtime.Value
var once_pure__2935994064 sync.Once
func Get_pure__2935994064() gopurs_runtime.Value {
	once_pure__2935994064.Do(func() {
		cache_pure__2935994064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__2935994064(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__2935994064
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_applyArray__2998472828 gopurs_runtime.Value
var once_applyArray__2998472828 sync.Once
func Get_applyArray__2998472828() gopurs_runtime.Value {
	once_applyArray__2998472828.Do(func() {
		cache_applyArray__2998472828 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), pkg_Control_Apply.Get_arrayApply())
	})
	return cache_applyArray__2998472828
}

var cache_plusArray__4260531026 gopurs_runtime.Value
var once_plusArray__4260531026 sync.Once
func Get_plusArray__4260531026() gopurs_runtime.Value {
	once_plusArray__4260531026.Do(func() {
		cache_plusArray__4260531026 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Alt.Get_altArray()
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
	})
	return cache_plusArray__4260531026
}

var cache_functorArray__361387505 gopurs_runtime.Value
var once_functorArray__361387505 sync.Once
func Get_functorArray__361387505() gopurs_runtime.Value {
	once_functorArray__361387505.Do(func() {
		cache_functorArray__361387505 = gopurs_runtime.RecordDict1("map", pkg_Data_Functor.Get_arrayMap())
	})
	return cache_functorArray__361387505
}

type Constructor_Alternative[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[397869517] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Alternative[gopurs_runtime.Value])(ptr)
		switch key {
		case "Applicative0": return c.V0
		case "Plus1": return c.V1
		default: panic("Key not found in dictionary Constructor_Alternative: " + key)
		}
	}
}


func Call_guard(dictAlternative_0_loop *Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
empty_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "empty")
_ = empty_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.IntVal) != (0) {
__t2 = gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_Unit.Get_unit())
goto end_branch_2
} else {

}
}
{
__t2 = empty_2_1
}
end_branch_2:
return __t2
})
}

func Call_pure__2935994064(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


