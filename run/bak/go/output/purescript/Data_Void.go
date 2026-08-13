package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Void_Void gopurs_runtime.Value
var once_Data_Void_Void sync.Once
func Get_Data_Void_Void() gopurs_runtime.Value {
	once_Data_Void_Void.Do(func() {
		cache_Data_Void_Void = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Void_Void(x_0_box)
})
	})
	return cache_Data_Void_Void
}

var cache_Data_Void_absurd gopurs_runtime.Value
var once_Data_Void_absurd sync.Once
func Get_Data_Void_absurd() gopurs_runtime.Value {
	once_Data_Void_absurd.Do(func() {
		cache_Data_Void_absurd = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Void_absurd(a_0_box)
})
	})
	return cache_Data_Void_absurd
}

var cache_Data_Void_absurd__1771830288 gopurs_runtime.Value
var once_Data_Void_absurd__1771830288 sync.Once
func Get_Data_Void_absurd__1771830288() gopurs_runtime.Value {
	once_Data_Void_absurd__1771830288.Do(func() {
		cache_Data_Void_absurd__1771830288 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Void_absurd__1771830288(a_0_box))
})
	})
	return cache_Data_Void_absurd__1771830288
}

var cache_Data_Void_absurd__1769020947 gopurs_runtime.Value
var once_Data_Void_absurd__1769020947 sync.Once
func Get_Data_Void_absurd__1769020947() gopurs_runtime.Value {
	once_Data_Void_absurd__1769020947.Do(func() {
		cache_Data_Void_absurd__1769020947 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Void_absurd__1769020947(a_0_box)
})
	})
	return cache_Data_Void_absurd__1769020947
}

var cache_Data_Void_absurd__331654555 gopurs_runtime.Value
var once_Data_Void_absurd__331654555 sync.Once
func Get_Data_Void_absurd__331654555() gopurs_runtime.Value {
	once_Data_Void_absurd__331654555.Do(func() {
		cache_Data_Void_absurd__331654555 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Void_absurd__331654555(a_0_box)
})
	})
	return cache_Data_Void_absurd__331654555
}

var cache_Data_Void_absurd__2082956474 gopurs_runtime.Value
var once_Data_Void_absurd__2082956474 sync.Once
func Get_Data_Void_absurd__2082956474() gopurs_runtime.Value {
	once_Data_Void_absurd__2082956474.Do(func() {
		cache_Data_Void_absurd__2082956474 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Void_absurd__2082956474(a_0_box))
})
	})
	return cache_Data_Void_absurd__2082956474
}

var cache_Data_Void_absurd__3279552488 gopurs_runtime.Value
var once_Data_Void_absurd__3279552488 sync.Once
func Get_Data_Void_absurd__3279552488() gopurs_runtime.Value {
	once_Data_Void_absurd__3279552488.Do(func() {
		cache_Data_Void_absurd__3279552488 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Void_absurd__3279552488(a_0_box)
})
	})
	return cache_Data_Void_absurd__3279552488
}

var cache_Data_Void_absurd__977285408 gopurs_runtime.Value
var once_Data_Void_absurd__977285408 sync.Once
func Get_Data_Void_absurd__977285408() gopurs_runtime.Value {
	once_Data_Void_absurd__977285408.Do(func() {
		cache_Data_Void_absurd__977285408 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Void_absurd__977285408(a_0_box)), UnsafePtr: nil}
})
	})
	return cache_Data_Void_absurd__977285408
}

var cache_Data_Void_absurd__943547431 gopurs_runtime.Value
var once_Data_Void_absurd__943547431 sync.Once
func Get_Data_Void_absurd__943547431() gopurs_runtime.Value {
	once_Data_Void_absurd__943547431.Do(func() {
		cache_Data_Void_absurd__943547431 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Void_absurd__943547431(a_0_box)
})
	})
	return cache_Data_Void_absurd__943547431
}

var cache_Data_Void_absurd__701346290 gopurs_runtime.Value
var once_Data_Void_absurd__701346290 sync.Once
func Get_Data_Void_absurd__701346290() gopurs_runtime.Value {
	once_Data_Void_absurd__701346290.Do(func() {
		cache_Data_Void_absurd__701346290 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Void_absurd__701346290(a_0_box)
})
	})
	return cache_Data_Void_absurd__701346290
}

func Call_Data_Void_Void(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Void_absurd(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_0 gopurs_runtime.Value
spin_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_0:
for {
if false { continue spin_1_0_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_0, a_0)
}

func Call_Data_Void_absurd__1771830288(a_0_loop gopurs_runtime.Value) string {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_1 gopurs_runtime.Value
spin_1_0_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_1:
for {
if false { continue spin_1_0_1 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_1
return gopurs_runtime.Str(gopurs_runtime.Value{}.StrVal())
}
}()
})
return gopurs_runtime.Apply(spin_1_0_1, a_0).StrVal()
}

func Call_Data_Void_absurd__1769020947(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_2 gopurs_runtime.Value
spin_1_0_2 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_2:
for {
if false { continue spin_1_0_2 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_2
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_2, a_0)
}

func Call_Data_Void_absurd__331654555(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_3 gopurs_runtime.Value
spin_1_0_3 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_3:
for {
if false { continue spin_1_0_3 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_3
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_3, a_0)
}

func Call_Data_Void_absurd__2082956474(a_0_loop gopurs_runtime.Value) bool {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_4 gopurs_runtime.Value
spin_1_0_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_4:
for {
if false { continue spin_1_0_4 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_4
return gopurs_runtime.Bool((gopurs_runtime.Value{}.IntVal) != (0))
}
}()
})
return (gopurs_runtime.Apply(spin_1_0_4, a_0).IntVal) != (0)
}

func Call_Data_Void_absurd__3279552488(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_5 gopurs_runtime.Value
spin_1_0_5 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_5:
for {
if false { continue spin_1_0_5 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_5
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_5, a_0)
}

func Call_Data_Void_absurd__977285408(a_0_loop gopurs_runtime.Value) uint32 {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_6 gopurs_runtime.Value
spin_1_0_6 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_6:
for {
if false { continue spin_1_0_6 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_6
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{}.IntVal)), UnsafePtr: nil}
}
}()
})
return uint32(gopurs_runtime.Apply(spin_1_0_6, a_0).IntVal)
}

func Call_Data_Void_absurd__943547431(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_7 gopurs_runtime.Value
spin_1_0_7 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_7:
for {
if false { continue spin_1_0_7 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_7
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_7, a_0)
}

func Call_Data_Void_absurd__701346290(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var spin_1_0_8 gopurs_runtime.Value
spin_1_0_8 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
spin_1_0_8:
for {
if false { continue spin_1_0_8 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0_8
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0_8, a_0)
}


