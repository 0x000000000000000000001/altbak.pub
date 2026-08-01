package Control_Semigroupoid

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_semigroupoidFn gopurs_runtime.Value
var once_semigroupoidFn sync.Once
func Get_semigroupoidFn() gopurs_runtime.Value {
	once_semigroupoidFn.Do(func() {
		cache_semigroupoidFn = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
}))
	})
	return cache_semigroupoidFn
}

var cache_compose gopurs_runtime.Value
var once_compose sync.Once
func Get_compose() gopurs_runtime.Value {
	once_compose.Do(func() {
		cache_compose = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose(dict_0_box)
})
	})
	return cache_compose
}

var cache_compose__gopurs_runtime_Value_1458035199 gopurs_runtime.Value
var once_compose__gopurs_runtime_Value_1458035199 sync.Once
func Get_compose__gopurs_runtime_Value_1458035199() gopurs_runtime.Value {
	once_compose__gopurs_runtime_Value_1458035199.Do(func() {
		cache_compose__gopurs_runtime_Value_1458035199 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__gopurs_runtime_Value_1458035199(dict_0_box)
})
	})
	return cache_compose__gopurs_runtime_Value_1458035199
}

var cache_compose__gopurs_runtime_Value_1992203551 gopurs_runtime.Value
var once_compose__gopurs_runtime_Value_1992203551 sync.Once
func Get_compose__gopurs_runtime_Value_1992203551() gopurs_runtime.Value {
	once_compose__gopurs_runtime_Value_1992203551.Do(func() {
		cache_compose__gopurs_runtime_Value_1992203551 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__gopurs_runtime_Value_1992203551(dict_0_box)
})
	})
	return cache_compose__gopurs_runtime_Value_1992203551
}

var cache_compose__gopurs_runtime_Value_3874593599 gopurs_runtime.Value
var once_compose__gopurs_runtime_Value_3874593599 sync.Once
func Get_compose__gopurs_runtime_Value_3874593599() gopurs_runtime.Value {
	once_compose__gopurs_runtime_Value_3874593599.Do(func() {
		cache_compose__gopurs_runtime_Value_3874593599 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__gopurs_runtime_Value_3874593599(dict_0_box)
})
	})
	return cache_compose__gopurs_runtime_Value_3874593599
}

var cache_composeFlipped gopurs_runtime.Value
var once_composeFlipped sync.Once
func Get_composeFlipped() gopurs_runtime.Value {
	once_composeFlipped.Do(func() {
		cache_composeFlipped = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped(dictSemigroupoid_0_box, f_1_box, g_2_box)
})
	})
	return cache_composeFlipped
}

var cache_composeFlipped__gopurs_runtime_Value_4011228415 gopurs_runtime.Value
var once_composeFlipped__gopurs_runtime_Value_4011228415 sync.Once
func Get_composeFlipped__gopurs_runtime_Value_4011228415() gopurs_runtime.Value {
	once_composeFlipped__gopurs_runtime_Value_4011228415.Do(func() {
		cache_composeFlipped__gopurs_runtime_Value_4011228415 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__gopurs_runtime_Value_4011228415(dictSemigroupoid_0_box, f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__gopurs_runtime_Value_4011228415
}

var cache_composeFlipped__gopurs_runtime_Value_770785887 gopurs_runtime.Value
var once_composeFlipped__gopurs_runtime_Value_770785887 sync.Once
func Get_composeFlipped__gopurs_runtime_Value_770785887() gopurs_runtime.Value {
	once_composeFlipped__gopurs_runtime_Value_770785887.Do(func() {
		cache_composeFlipped__gopurs_runtime_Value_770785887 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__gopurs_runtime_Value_770785887(dictSemigroupoid_0_box, f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__gopurs_runtime_Value_770785887
}

var cache_composeFlipped__gopurs_runtime_Value_3153195711 gopurs_runtime.Value
var once_composeFlipped__gopurs_runtime_Value_3153195711 sync.Once
func Get_composeFlipped__gopurs_runtime_Value_3153195711() gopurs_runtime.Value {
	once_composeFlipped__gopurs_runtime_Value_3153195711.Do(func() {
		cache_composeFlipped__gopurs_runtime_Value_3153195711 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__gopurs_runtime_Value_3153195711(dictSemigroupoid_0_box, f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__gopurs_runtime_Value_3153195711
}

var cache_composeFlipped__gopurs_runtime_Value_3874593599 gopurs_runtime.Value
var once_composeFlipped__gopurs_runtime_Value_3874593599 sync.Once
func Get_composeFlipped__gopurs_runtime_Value_3874593599() gopurs_runtime.Value {
	once_composeFlipped__gopurs_runtime_Value_3874593599.Do(func() {
		cache_composeFlipped__gopurs_runtime_Value_3874593599 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__gopurs_runtime_Value_3874593599(dictSemigroupoid_0_box, f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__gopurs_runtime_Value_3874593599
}

func Call_compose(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "compose")
}

func Call_compose__gopurs_runtime_Value_1458035199(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "compose")
}

func Call_compose__gopurs_runtime_Value_1992203551(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "compose")
}

func Call_compose__gopurs_runtime_Value_3874593599(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "compose")
}

func Call_composeFlipped(dictSemigroupoid_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), g_2, f_1)
}

func Call_composeFlipped__gopurs_runtime_Value_4011228415(dictSemigroupoid_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), g_2, f_1)
}

func Call_composeFlipped__gopurs_runtime_Value_770785887(dictSemigroupoid_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), g_2, f_1)
}

func Call_composeFlipped__gopurs_runtime_Value_3153195711(dictSemigroupoid_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), g_2, f_1)
}

func Call_composeFlipped__gopurs_runtime_Value_3874593599(dictSemigroupoid_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), g_2, f_1)
}


