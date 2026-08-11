package Safe_Coerce

import (
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_coerce gopurs_runtime.Value
var once_coerce sync.Once
func Get_coerce() gopurs_runtime.Value {
	once_coerce.Do(func() {
		cache_coerce = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coerce(_dollar__unused_0_box)
})
	})
	return cache_coerce
}

var cache_coerce__gopurs_runtime_Value_529298068 gopurs_runtime.Value
var once_coerce__gopurs_runtime_Value_529298068 sync.Once
func Get_coerce__gopurs_runtime_Value_529298068() gopurs_runtime.Value {
	once_coerce__gopurs_runtime_Value_529298068.Do(func() {
		cache_coerce__gopurs_runtime_Value_529298068 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coerce__gopurs_runtime_Value_529298068(_dollar__unused_0_box)
})
	})
	return cache_coerce__gopurs_runtime_Value_529298068
}

func Call_coerce(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_coerce__gopurs_runtime_Value_529298068(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}


