package Control_Monad_Gen_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_sized gopurs_runtime.Value
var once_sized sync.Once
func Get_sized() gopurs_runtime.Value {
	once_sized.Do(func() {
		cache_sized = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sized(dict_0_box)
})
	})
	return cache_sized
}

var cache_resize gopurs_runtime.Value
var once_resize sync.Once
func Get_resize() gopurs_runtime.Value {
	once_resize.Do(func() {
		cache_resize = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_resize(dict_0_box)
})
	})
	return cache_resize
}

var cache_chooseInt gopurs_runtime.Value
var once_chooseInt sync.Once
func Get_chooseInt() gopurs_runtime.Value {
	once_chooseInt.Do(func() {
		cache_chooseInt = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseInt(dict_0_box)
})
	})
	return cache_chooseInt
}

var cache_chooseFloat gopurs_runtime.Value
var once_chooseFloat sync.Once
func Get_chooseFloat() gopurs_runtime.Value {
	once_chooseFloat.Do(func() {
		cache_chooseFloat = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseFloat(dict_0_box)
})
	})
	return cache_chooseFloat
}

var cache_chooseBool gopurs_runtime.Value
var once_chooseBool sync.Once
func Get_chooseBool() gopurs_runtime.Value {
	once_chooseBool.Do(func() {
		cache_chooseBool = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseBool(dict_0_box)
})
	})
	return cache_chooseBool
}

func Call_sized(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "sized")
}

func Call_resize(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "resize")
}

func Call_chooseInt(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "chooseInt")
}

func Call_chooseFloat(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "chooseFloat")
}

func Call_chooseBool(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "chooseBool")
}


