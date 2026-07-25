package Control_Monad_ST_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Effect "gopurs/output/Effect"
)

var cache_monadSTST gopurs_runtime.Value
var once_monadSTST sync.Once
func Get_monadSTST() gopurs_runtime.Value {
	once_monadSTST.Do(func() {
		cache_monadSTST = gopurs_runtime.RecordDict2("liftST", "Monad0", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_monadST()
}))
	})
	return cache_monadSTST
}

var cache_monadSTEffect gopurs_runtime.Value
var once_monadSTEffect sync.Once
func Get_monadSTEffect() gopurs_runtime.Value {
	once_monadSTEffect.Do(func() {
		cache_monadSTEffect = gopurs_runtime.RecordDict2("liftST", "Monad0", pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
}))
	})
	return cache_monadSTEffect
}

var cache_liftST gopurs_runtime.Value
var once_liftST sync.Once
func Get_liftST() gopurs_runtime.Value {
	once_liftST.Do(func() {
		cache_liftST = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "liftST")
}()
})
	})
	return cache_liftST
}




