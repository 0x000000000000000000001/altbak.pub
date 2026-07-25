package Data_Array_ST_Partial

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_ST_Uncurried "gopurs/output/Control.Monad.ST.Uncurried"
)

var cache_poke gopurs_runtime.Value
var once_poke sync.Once
func Get_poke() gopurs_runtime.Value {
	once_poke.Do(func() {
		cache_poke = gopurs_runtime.Func(func(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), Get_pokeImpl())
}()
})
	})
	return cache_poke
}

var cache_peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		cache_peek = gopurs_runtime.Func(func(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), Get_peekImpl())
}()
})
	})
	return cache_peek
}



func Get_peekImpl() gopurs_runtime.Value {
	return _Gopurs_PeekImpl
}

func Get_pokeImpl() gopurs_runtime.Value {
	return _Gopurs_PokeImpl
}
