package Control_Plus

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Alt "gopurs/output/Control.Alt"
)

var cache_plusArray gopurs_runtime.Value
var once_plusArray sync.Once
func Get_plusArray() gopurs_runtime.Value {
	once_plusArray.Do(func() {
		cache_plusArray = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Alt.Get_altArray()
}), gopurs_runtime.Array([]gopurs_runtime.Value{}))))
	})
	return cache_plusArray
}

var cache_empty gopurs_runtime.Value
var once_empty sync.Once
func Get_empty() gopurs_runtime.Value {
	once_empty.Do(func() {
		cache_empty = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_empty(dict_0_box))
})
	})
	return cache_empty
}

var cache_empty__func_gopurs_runtime_Value__interface___134048739 gopurs_runtime.Value
var once_empty__func_gopurs_runtime_Value__interface___134048739 sync.Once
func Get_empty__func_gopurs_runtime_Value__interface___134048739() gopurs_runtime.Value {
	once_empty__func_gopurs_runtime_Value__interface___134048739.Do(func() {
		cache_empty__func_gopurs_runtime_Value__interface___134048739 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_empty__func_gopurs_runtime_Value__interface___134048739(dict_0_box))
})
	})
	return cache_empty__func_gopurs_runtime_Value__interface___134048739
}

func Call_empty(dict_0_loop gopurs_runtime.Value) interface{} {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(dict_0, "empty"))
}

func Call_empty__func_gopurs_runtime_Value__interface___134048739(dict_0_loop gopurs_runtime.Value) interface{} {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(dict_0, "empty"))
}
