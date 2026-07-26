package Control_Comonad_Traced

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Control_Comonad_Traced_Trans "gopurs/output/Control.Comonad.Traced.Trans"
	pkg_Data_Identity "gopurs/output/Data.Identity"
)

var cache_traced gopurs_runtime.Value
var once_traced sync.Once
func Get_traced() gopurs_runtime.Value {
	once_traced.Do(func() {
		cache_traced = gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Control_Comonad_Traced_Trans.Get_TracedT(), pkg_Data_Identity.Get_Identity())
	})
	return cache_traced
}

var cache_runTraced gopurs_runtime.Value
var once_runTraced sync.Once
func Get_runTraced() gopurs_runtime.Value {
	once_runTraced.Do(func() {
		cache_runTraced = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runTraced(v_0_box)
})
	})
	return cache_runTraced
}

func Call_runTraced(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}


