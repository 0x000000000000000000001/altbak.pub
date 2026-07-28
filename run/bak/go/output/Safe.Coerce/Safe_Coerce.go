package Safe_Coerce

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
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

var cache_coerce__gopurs_runtime_Value_2006084213 gopurs_runtime.Value
var once_coerce__gopurs_runtime_Value_2006084213 sync.Once
func Get_coerce__gopurs_runtime_Value_2006084213() gopurs_runtime.Value {
	once_coerce__gopurs_runtime_Value_2006084213.Do(func() {
		cache_coerce__gopurs_runtime_Value_2006084213 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coerce__gopurs_runtime_Value_2006084213(_dollar__unused_0_box)
})
	})
	return cache_coerce__gopurs_runtime_Value_2006084213
}

func Call_coerce(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_coerce__gopurs_runtime_Value_2006084213(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}


