package Data_Void

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_absurd gopurs_runtime.Value
var once_absurd sync.Once
func Get_absurd() gopurs_runtime.Value {
	once_absurd.Do(func() {
		cache_absurd = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_absurd(a_0_box))
})
	})
	return cache_absurd
}

var cache_absurd__func_gopurs_runtime_Value__string_1771830288 gopurs_runtime.Value
var once_absurd__func_gopurs_runtime_Value__string_1771830288 sync.Once
func Get_absurd__func_gopurs_runtime_Value__string_1771830288() gopurs_runtime.Value {
	once_absurd__func_gopurs_runtime_Value__string_1771830288.Do(func() {
		cache_absurd__func_gopurs_runtime_Value__string_1771830288 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_absurd__func_gopurs_runtime_Value__string_1771830288(a_0_box))
})
	})
	return cache_absurd__func_gopurs_runtime_Value__string_1771830288
}

var cache_absurd__func_gopurs_runtime_Value__gopurs_runtime_Value_331654555 gopurs_runtime.Value
var once_absurd__func_gopurs_runtime_Value__gopurs_runtime_Value_331654555 sync.Once
func Get_absurd__func_gopurs_runtime_Value__gopurs_runtime_Value_331654555() gopurs_runtime.Value {
	once_absurd__func_gopurs_runtime_Value__gopurs_runtime_Value_331654555.Do(func() {
		cache_absurd__func_gopurs_runtime_Value__gopurs_runtime_Value_331654555 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_absurd__func_gopurs_runtime_Value__gopurs_runtime_Value_331654555(a_0_box)
})
	})
	return cache_absurd__func_gopurs_runtime_Value__gopurs_runtime_Value_331654555
}

var cache_absurd__func_gopurs_runtime_Value__interface___1344965310 gopurs_runtime.Value
var once_absurd__func_gopurs_runtime_Value__interface___1344965310 sync.Once
func Get_absurd__func_gopurs_runtime_Value__interface___1344965310() gopurs_runtime.Value {
	once_absurd__func_gopurs_runtime_Value__interface___1344965310.Do(func() {
		cache_absurd__func_gopurs_runtime_Value__interface___1344965310 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_absurd__func_gopurs_runtime_Value__interface___1344965310(a_0_box))
})
	})
	return cache_absurd__func_gopurs_runtime_Value__interface___1344965310
}

var cache_absurd__func_gopurs_runtime_Value__interface___1234480061 gopurs_runtime.Value
var once_absurd__func_gopurs_runtime_Value__interface___1234480061 sync.Once
func Get_absurd__func_gopurs_runtime_Value__interface___1234480061() gopurs_runtime.Value {
	once_absurd__func_gopurs_runtime_Value__interface___1234480061.Do(func() {
		cache_absurd__func_gopurs_runtime_Value__interface___1234480061 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_absurd__func_gopurs_runtime_Value__interface___1234480061(a_0_box))
})
	})
	return cache_absurd__func_gopurs_runtime_Value__interface___1234480061
}

func Call_absurd(a_0_loop gopurs_runtime.Value) interface{} {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop interface{} = gopurs_runtime.UnboxAny(v_2_loop_val)
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 interface{} = v_2_loop
_ = v_2
v_2_loop = gopurs_runtime.UnboxAny(v_2)
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Any(spin_1_0), a_0))
}

func Call_absurd__func_gopurs_runtime_Value__string_1771830288(a_0_loop gopurs_runtime.Value) string {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop interface{} = gopurs_runtime.UnboxAny(v_2_loop_val)
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 interface{} = v_2_loop
_ = v_2
v_2_loop = gopurs_runtime.UnboxAny(v_2)
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(gopurs_runtime.Any(spin_1_0), a_0).StrVal()
}

func Call_absurd__func_gopurs_runtime_Value__gopurs_runtime_Value_331654555(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop interface{} = gopurs_runtime.UnboxAny(v_2_loop_val)
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 interface{} = v_2_loop
_ = v_2
v_2_loop = gopurs_runtime.UnboxAny(v_2)
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(gopurs_runtime.Any(spin_1_0), a_0)
}

func Call_absurd__func_gopurs_runtime_Value__interface___1344965310(a_0_loop gopurs_runtime.Value) interface{} {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop interface{} = gopurs_runtime.UnboxAny(v_2_loop_val)
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 interface{} = v_2_loop
_ = v_2
v_2_loop = gopurs_runtime.UnboxAny(v_2)
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Any(spin_1_0), a_0))
}

func Call_absurd__func_gopurs_runtime_Value__interface___1234480061(a_0_loop gopurs_runtime.Value) interface{} {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop interface{} = gopurs_runtime.UnboxAny(v_2_loop_val)
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 interface{} = v_2_loop
_ = v_2
v_2_loop = gopurs_runtime.UnboxAny(v_2)
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Any(spin_1_0), a_0))
}
