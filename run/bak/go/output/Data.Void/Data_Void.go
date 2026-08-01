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
return Call_absurd(a_0_box)
})
	})
	return cache_absurd
}

var cache_absurd__gopurs_runtime_Value_1771830288 gopurs_runtime.Value
var once_absurd__gopurs_runtime_Value_1771830288 sync.Once
func Get_absurd__gopurs_runtime_Value_1771830288() gopurs_runtime.Value {
	once_absurd__gopurs_runtime_Value_1771830288.Do(func() {
		cache_absurd__gopurs_runtime_Value_1771830288 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_absurd__gopurs_runtime_Value_1771830288(a_0_box))
})
	})
	return cache_absurd__gopurs_runtime_Value_1771830288
}

var cache_absurd__gopurs_runtime_Value_2182866177 gopurs_runtime.Value
var once_absurd__gopurs_runtime_Value_2182866177 sync.Once
func Get_absurd__gopurs_runtime_Value_2182866177() gopurs_runtime.Value {
	once_absurd__gopurs_runtime_Value_2182866177.Do(func() {
		cache_absurd__gopurs_runtime_Value_2182866177 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_absurd__gopurs_runtime_Value_2182866177(a_0_box))
})
	})
	return cache_absurd__gopurs_runtime_Value_2182866177
}

var cache_absurd__gopurs_runtime_Value_1499175369 gopurs_runtime.Value
var once_absurd__gopurs_runtime_Value_1499175369 sync.Once
func Get_absurd__gopurs_runtime_Value_1499175369() gopurs_runtime.Value {
	once_absurd__gopurs_runtime_Value_1499175369.Do(func() {
		cache_absurd__gopurs_runtime_Value_1499175369 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_absurd__gopurs_runtime_Value_1499175369(a_0_box)
})
	})
	return cache_absurd__gopurs_runtime_Value_1499175369
}

var cache_absurd__gopurs_runtime_Value_1133204955 gopurs_runtime.Value
var once_absurd__gopurs_runtime_Value_1133204955 sync.Once
func Get_absurd__gopurs_runtime_Value_1133204955() gopurs_runtime.Value {
	once_absurd__gopurs_runtime_Value_1133204955.Do(func() {
		cache_absurd__gopurs_runtime_Value_1133204955 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_absurd__gopurs_runtime_Value_1133204955(a_0_box)
})
	})
	return cache_absurd__gopurs_runtime_Value_1133204955
}

var cache_absurd__gopurs_runtime_Value_331654555 gopurs_runtime.Value
var once_absurd__gopurs_runtime_Value_331654555 sync.Once
func Get_absurd__gopurs_runtime_Value_331654555() gopurs_runtime.Value {
	once_absurd__gopurs_runtime_Value_331654555.Do(func() {
		cache_absurd__gopurs_runtime_Value_331654555 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_absurd__gopurs_runtime_Value_331654555(a_0_box)
})
	})
	return cache_absurd__gopurs_runtime_Value_331654555
}

var cache_absurd__gopurs_runtime_Value_1769020947 gopurs_runtime.Value
var once_absurd__gopurs_runtime_Value_1769020947 sync.Once
func Get_absurd__gopurs_runtime_Value_1769020947() gopurs_runtime.Value {
	once_absurd__gopurs_runtime_Value_1769020947.Do(func() {
		cache_absurd__gopurs_runtime_Value_1769020947 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_absurd__gopurs_runtime_Value_1769020947(a_0_box)
})
	})
	return cache_absurd__gopurs_runtime_Value_1769020947
}

func Call_absurd(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0, a_0)
}

func Call_absurd__gopurs_runtime_Value_1771830288(a_0_loop gopurs_runtime.Value) string {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0, a_0).StrVal()
}

func Call_absurd__gopurs_runtime_Value_2182866177(a_0_loop gopurs_runtime.Value) bool {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return (gopurs_runtime.Apply(spin_1_0, a_0).IntVal) != (0)
}

func Call_absurd__gopurs_runtime_Value_1499175369(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0, a_0)
}

func Call_absurd__gopurs_runtime_Value_1133204955(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0, a_0)
}

func Call_absurd__gopurs_runtime_Value_331654555(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0, a_0)
}

func Call_absurd__gopurs_runtime_Value_1769020947(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0, a_0)
}


