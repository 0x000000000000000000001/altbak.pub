package Control_MonadPlus

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Plus "gopurs/output/Control.Plus"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_monadPlusArray gopurs_runtime.Value
var once_monadPlusArray sync.Once
func Get_monadPlusArray() gopurs_runtime.Value {
	once_monadPlusArray.Do(func() {
		cache_monadPlusArray = gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Alternative.Get_alternativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad.Get_monadArray()
}))
	})
	return cache_monadPlusArray
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

var cache_alternativeArray__1996030013 gopurs_runtime.Value
var once_alternativeArray__1996030013 sync.Once
func Get_alternativeArray__1996030013() gopurs_runtime.Value {
	once_alternativeArray__1996030013.Do(func() {
		cache_alternativeArray__1996030013 = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Plus.Get_plusArray()
}))
	})
	return cache_alternativeArray__1996030013
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

var cache_bindArray__1650562023 gopurs_runtime.Value
var once_bindArray__1650562023 sync.Once
func Get_bindArray__1650562023() gopurs_runtime.Value {
	once_bindArray__1650562023.Do(func() {
		cache_bindArray__1650562023 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
}), pkg_Control_Bind.Get_arrayBind())
	})
	return cache_bindArray__1650562023
}

var cache_monadArray__2289780851 gopurs_runtime.Value
var once_monadArray__2289780851 sync.Once
func Get_monadArray__2289780851() gopurs_runtime.Value {
	once_monadArray__2289780851.Do(func() {
		cache_monadArray__2289780851 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindArray()
}))
	})
	return cache_monadArray__2289780851
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

type Constructor_MonadPlus[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3236234573] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadPlus[gopurs_runtime.Value])(ptr)
		switch key {
		case "Alternative1": return c.V0
		case "Monad0": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadPlus: " + key)
		}
	}
}



