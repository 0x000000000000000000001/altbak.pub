package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Enum_bottom gopurs_runtime.Value
var once_Data_Enum_bottom sync.Once
func Get_Data_Enum_bottom() gopurs_runtime.Value {
	once_Data_Enum_bottom.Do(func() {
		cache_Data_Enum_bottom = gopurs_runtime.Int(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3764732725_2094947566(Rebox_Data_Enum_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedInt()))))}).IntVal)
	})
	return cache_Data_Enum_bottom
}

var cache_Data_Enum_fromJust gopurs_runtime.Value
var once_Data_Enum_fromJust sync.Once
func Get_Data_Enum_fromJust() gopurs_runtime.Value {
	once_Data_Enum_fromJust.Do(func() {
		cache_Data_Enum_fromJust = gopurs_runtime.Apply(Get_Data_Maybe_fromJust(), gopurs_runtime.Value{})
	})
	return cache_Data_Enum_fromJust
}

var cache_Data_Enum_bottom1 gopurs_runtime.Value
var once_Data_Enum_bottom1 sync.Once
func Get_Data_Enum_bottom1() gopurs_runtime.Value {
	once_Data_Enum_bottom1.Do(func() {
		cache_Data_Enum_bottom1 = gopurs_runtime.Str(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}).StrVal())
	})
	return cache_Data_Enum_bottom1
}

var cache_Data_Enum_top gopurs_runtime.Value
var once_Data_Enum_top sync.Once
func Get_Data_Enum_top() gopurs_runtime.Value {
	once_Data_Enum_top.Do(func() {
		cache_Data_Enum_top = gopurs_runtime.Str(Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}).StrVal())
	})
	return cache_Data_Enum_top
}

var cache_Data_Enum_Enum_dollar_Dict gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict sync.Once
func Get_Data_Enum_Enum_dollar_Dict() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Call_Data_Enum_Enum_dollar_Dict(func() struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}{}
					clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
					clone.pred = gopurs_runtime.RecordGet(orig, "pred")
					clone.succ = gopurs_runtime.RecordGet(orig, "succ")
					return clone
				}()))}
})
	})
	return cache_Data_Enum_Enum_dollar_Dict
}

var cache_Data_Enum_Enum_dollar_Dict__2787884439 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__2787884439 sync.Once
func Get_Data_Enum_Enum_dollar_Dict__2787884439() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__2787884439.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__2787884439 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094(Call_Data_Enum_Enum_dollar_Dict__2787884439(func() struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}{}
					clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
					clone.pred = gopurs_runtime.RecordGet(orig, "pred")
					clone.succ = gopurs_runtime.RecordGet(orig, "succ")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_Enum_dollar_Dict__2787884439
}

var cache_Data_Enum_Enum_dollar_Dict__479058327 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__479058327 sync.Once
func Get_Data_Enum_Enum_dollar_Dict__479058327() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__479058327.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__479058327 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094(Call_Data_Enum_Enum_dollar_Dict__479058327(func() struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}{}
					clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
					clone.pred = gopurs_runtime.RecordGet(orig, "pred")
					clone.succ = gopurs_runtime.RecordGet(orig, "succ")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_Enum_dollar_Dict__479058327
}

var cache_Data_Enum_Enum_dollar_Dict__3712690007 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__3712690007 sync.Once
func Get_Data_Enum_Enum_dollar_Dict__3712690007() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__3712690007.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__3712690007 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3494563625_556578094(Call_Data_Enum_Enum_dollar_Dict__3712690007(func() struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}{}
					clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
					clone.pred = gopurs_runtime.RecordGet(orig, "pred")
					clone.succ = gopurs_runtime.RecordGet(orig, "succ")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_Enum_dollar_Dict__3712690007
}

var cache_Data_Enum_Enum_dollar_Dict__2123951319 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__2123951319 sync.Once
func Get_Data_Enum_Enum_dollar_Dict__2123951319() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__2123951319.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__2123951319 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094(Call_Data_Enum_Enum_dollar_Dict__2123951319(func() struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}{}
					clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
					clone.pred = gopurs_runtime.RecordGet(orig, "pred")
					clone.succ = gopurs_runtime.RecordGet(orig, "succ")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_Enum_dollar_Dict__2123951319
}

var cache_Data_Enum_Enum_dollar_Dict__1439164631 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__1439164631 sync.Once
func Get_Data_Enum_Enum_dollar_Dict__1439164631() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__1439164631.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__1439164631 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_342165130_556578094(Call_Data_Enum_Enum_dollar_Dict__1439164631(func() struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}{}
					clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
					clone.pred = gopurs_runtime.RecordGet(orig, "pred")
					clone.succ = gopurs_runtime.RecordGet(orig, "succ")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_Enum_dollar_Dict__1439164631
}

var cache_Data_Enum_Enum_dollar_Dict__2220371927 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__2220371927 sync.Once
func Get_Data_Enum_Enum_dollar_Dict__2220371927() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__2220371927.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__2220371927 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3409577041_556578094(Call_Data_Enum_Enum_dollar_Dict__2220371927(func() struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}{}
					clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
					clone.pred = gopurs_runtime.RecordGet(orig, "pred")
					clone.succ = gopurs_runtime.RecordGet(orig, "succ")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_Enum_dollar_Dict__2220371927
}

var cache_Data_Enum_Enum_dollar_Dict__3674609207 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__3674609207 sync.Once
func Get_Data_Enum_Enum_dollar_Dict__3674609207() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__3674609207.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__3674609207 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4060049525_556578094(Call_Data_Enum_Enum_dollar_Dict__3674609207(func() struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}{}
					clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
					clone.pred = gopurs_runtime.RecordGet(orig, "pred")
					clone.succ = gopurs_runtime.RecordGet(orig, "succ")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_Enum_dollar_Dict__3674609207
}

var cache_Data_Enum_Enum_dollar_Dict__4138904023 gopurs_runtime.Value
var once_Data_Enum_Enum_dollar_Dict__4138904023 sync.Once
func Get_Data_Enum_Enum_dollar_Dict__4138904023() gopurs_runtime.Value {
	once_Data_Enum_Enum_dollar_Dict__4138904023.Do(func() {
		cache_Data_Enum_Enum_dollar_Dict__4138904023 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Call_Data_Enum_Enum_dollar_Dict__4138904023(func() struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}{}
					clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
					clone.pred = gopurs_runtime.RecordGet(orig, "pred")
					clone.succ = gopurs_runtime.RecordGet(orig, "succ")
					return clone
				}()))}
})
	})
	return cache_Data_Enum_Enum_dollar_Dict__4138904023
}

var cache_Data_Enum_Cardinality gopurs_runtime.Value
var once_Data_Enum_Cardinality sync.Once
func Get_Data_Enum_Cardinality() gopurs_runtime.Value {
	once_Data_Enum_Cardinality.Do(func() {
		cache_Data_Enum_Cardinality = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality
}

var cache_Data_Enum_Cardinality__1400572777 gopurs_runtime.Value
var once_Data_Enum_Cardinality__1400572777 sync.Once
func Get_Data_Enum_Cardinality__1400572777() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__1400572777.Do(func() {
		cache_Data_Enum_Cardinality__1400572777 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__1400572777(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__1400572777
}

var cache_Data_Enum_Cardinality__2274214265 gopurs_runtime.Value
var once_Data_Enum_Cardinality__2274214265 sync.Once
func Get_Data_Enum_Cardinality__2274214265() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__2274214265.Do(func() {
		cache_Data_Enum_Cardinality__2274214265 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__2274214265(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__2274214265
}

var cache_Data_Enum_Cardinality__123016565 gopurs_runtime.Value
var once_Data_Enum_Cardinality__123016565 sync.Once
func Get_Data_Enum_Cardinality__123016565() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__123016565.Do(func() {
		cache_Data_Enum_Cardinality__123016565 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__123016565(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__123016565
}

var cache_Data_Enum_Cardinality__1784762877 gopurs_runtime.Value
var once_Data_Enum_Cardinality__1784762877 sync.Once
func Get_Data_Enum_Cardinality__1784762877() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__1784762877.Do(func() {
		cache_Data_Enum_Cardinality__1784762877 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__1784762877(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__1784762877
}

var cache_Data_Enum_Cardinality__4276004359 gopurs_runtime.Value
var once_Data_Enum_Cardinality__4276004359 sync.Once
func Get_Data_Enum_Cardinality__4276004359() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__4276004359.Do(func() {
		cache_Data_Enum_Cardinality__4276004359 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__4276004359(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__4276004359
}

var cache_Data_Enum_Cardinality__3890726945 gopurs_runtime.Value
var once_Data_Enum_Cardinality__3890726945 sync.Once
func Get_Data_Enum_Cardinality__3890726945() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__3890726945.Do(func() {
		cache_Data_Enum_Cardinality__3890726945 = gopurs_runtime.Func(func(x_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__3890726945(x_unused_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__3890726945
}

var cache_Data_Enum_Cardinality__2587234619 gopurs_runtime.Value
var once_Data_Enum_Cardinality__2587234619 sync.Once
func Get_Data_Enum_Cardinality__2587234619() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__2587234619.Do(func() {
		cache_Data_Enum_Cardinality__2587234619 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__2587234619(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__2587234619
}

var cache_Data_Enum_Cardinality__4110088112 gopurs_runtime.Value
var once_Data_Enum_Cardinality__4110088112 sync.Once
func Get_Data_Enum_Cardinality__4110088112() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__4110088112.Do(func() {
		cache_Data_Enum_Cardinality__4110088112 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__4110088112(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__4110088112
}

var cache_Data_Enum_Cardinality__4100296677 gopurs_runtime.Value
var once_Data_Enum_Cardinality__4100296677 sync.Once
func Get_Data_Enum_Cardinality__4100296677() gopurs_runtime.Value {
	once_Data_Enum_Cardinality__4100296677.Do(func() {
		cache_Data_Enum_Cardinality__4100296677 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Cardinality__4100296677(x_0_box.IntVal)
})
	})
	return cache_Data_Enum_Cardinality__4100296677
}

var cache_Data_Enum_BoundedEnum_dollar_Dict gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict sync.Once
func Get_Data_Enum_BoundedEnum_dollar_Dict() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Call_Data_Enum_BoundedEnum_dollar_Dict(func() struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}{}
					clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
					clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
					clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
					clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
					clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
					return clone
				}()))}
})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__4237532871 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__4237532871 sync.Once
func Get_Data_Enum_BoundedEnum_dollar_Dict__4237532871() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__4237532871.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__4237532871 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4021906832_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__4237532871(func() struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}{}
					clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
					clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
					clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
					clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
					clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__4237532871
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__2081071559 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__2081071559 sync.Once
func Get_Data_Enum_BoundedEnum_dollar_Dict__2081071559() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__2081071559.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__2081071559 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4021906832_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__2081071559(func() struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}{}
					clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
					clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
					clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
					clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
					clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__2081071559
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__2130790279 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__2130790279 sync.Once
func Get_Data_Enum_BoundedEnum_dollar_Dict__2130790279() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__2130790279.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__2130790279 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4021906832_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__2130790279(func() struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}{}
					clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
					clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
					clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
					clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
					clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__2130790279
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__1429463943 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__1429463943 sync.Once
func Get_Data_Enum_BoundedEnum_dollar_Dict__1429463943() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__1429463943.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__1429463943 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3083804185_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__1429463943(func() struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}{}
					clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
					clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
					clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
					clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
					clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__1429463943
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__882291527 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__882291527 sync.Once
func Get_Data_Enum_BoundedEnum_dollar_Dict__882291527() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__882291527.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__882291527 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3569500834_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__882291527(func() struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}{}
					clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
					clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
					clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
					clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
					clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__882291527
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__2897860935 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__2897860935 sync.Once
func Get_Data_Enum_BoundedEnum_dollar_Dict__2897860935() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__2897860935.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__2897860935 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1306125126_123048125(Call_Data_Enum_BoundedEnum_dollar_Dict__2897860935(func() struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}{}
					clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
					clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
					clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
					clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
					clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
					return clone
				}())))}
})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__2897860935
}

var cache_Data_Enum_BoundedEnum_dollar_Dict__3240687239 gopurs_runtime.Value
var once_Data_Enum_BoundedEnum_dollar_Dict__3240687239 sync.Once
func Get_Data_Enum_BoundedEnum_dollar_Dict__3240687239() gopurs_runtime.Value {
	once_Data_Enum_BoundedEnum_dollar_Dict__3240687239.Do(func() {
		cache_Data_Enum_BoundedEnum_dollar_Dict__3240687239 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Call_Data_Enum_BoundedEnum_dollar_Dict__3240687239(func() struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}{}
					clone.Bounded0 = gopurs_runtime.RecordGet(orig, "Bounded0")
					clone.Enum1 = gopurs_runtime.RecordGet(orig, "Enum1")
					clone.cardinality = gopurs_runtime.RecordGet(orig, "cardinality")
					clone.fromEnum = gopurs_runtime.RecordGet(orig, "fromEnum")
					clone.toEnum = gopurs_runtime.RecordGet(orig, "toEnum")
					return clone
				}()))}
})
	})
	return cache_Data_Enum_BoundedEnum_dollar_Dict__3240687239
}

var cache_Data_Enum_toEnum gopurs_runtime.Value
var once_Data_Enum_toEnum sync.Once
func Get_Data_Enum_toEnum() gopurs_runtime.Value {
	once_Data_Enum_toEnum.Do(func() {
		cache_Data_Enum_toEnum = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Enum_toEnum
}

var cache_Data_Enum_toEnum__3016447460 gopurs_runtime.Value
var once_Data_Enum_toEnum__3016447460 sync.Once
func Get_Data_Enum_toEnum__3016447460() gopurs_runtime.Value {
	once_Data_Enum_toEnum__3016447460.Do(func() {
		cache_Data_Enum_toEnum__3016447460 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__3016447460(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_toEnum__3016447460
}

var cache_Data_Enum_toEnum__3403550180 gopurs_runtime.Value
var once_Data_Enum_toEnum__3403550180 sync.Once
func Get_Data_Enum_toEnum__3403550180() gopurs_runtime.Value {
	once_Data_Enum_toEnum__3403550180.Do(func() {
		cache_Data_Enum_toEnum__3403550180 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__3403550180(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_toEnum__3403550180
}

var cache_Data_Enum_charToEnum gopurs_runtime.Value
var once_Data_Enum_charToEnum sync.Once
func Get_Data_Enum_charToEnum() gopurs_runtime.Value {
	once_Data_Enum_charToEnum.Do(func() {
		cache_Data_Enum_charToEnum = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum(v_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_charToEnum
}

var cache_Data_Enum_defaultSucc__3855932694 gopurs_runtime.Value
var once_Data_Enum_defaultSucc__3855932694 sync.Once
func Get_Data_Enum_defaultSucc__3855932694() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc__3855932694.Do(func() {
		cache_Data_Enum_defaultSucc__3855932694 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultSucc__3855932694(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_defaultSucc__3855932694
}

var cache_Data_Enum_defaultPred__3855932694 gopurs_runtime.Value
var once_Data_Enum_defaultPred__3855932694 sync.Once
func Get_Data_Enum_defaultPred__3855932694() gopurs_runtime.Value {
	once_Data_Enum_defaultPred__3855932694.Do(func() {
		cache_Data_Enum_defaultPred__3855932694 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultPred__3855932694(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_defaultPred__3855932694
}

var cache_Data_Enum_enumChar gopurs_runtime.Value
var once_Data_Enum_enumChar sync.Once
func Get_Data_Enum_enumChar() gopurs_runtime.Value {
	once_Data_Enum_enumChar.Do(func() {
		cache_Data_Enum_enumChar = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3409577041_556578094((&Constructor_Data_Enum_Enum[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2406510097_4177771502(Rebox_Data_Enum_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordChar()))))}
}), gopurs_runtime.Apply2(Get_Data_Enum_defaultPred__3855932694(), Get_Data_Enum_charToEnum(), Get_Data_Enum_toCharCode()), gopurs_runtime.Apply2(Get_Data_Enum_defaultSucc__3855932694(), Get_Data_Enum_charToEnum(), Get_Data_Enum_toCharCode())})))}
	})
	return cache_Data_Enum_enumChar
}

var cache_Data_Enum_boundedEnumChar gopurs_runtime.Value
var once_Data_Enum_boundedEnumChar sync.Once
func Get_Data_Enum_boundedEnumChar() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumChar.Do(func() {
		cache_Data_Enum_boundedEnumChar = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3569500834_123048125((&Constructor_Data_Enum_BoundedEnum[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3409577041_556578094(Rebox_Data_Enum_556578094_3409577041(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Enum_enumChar()))))}
}), gopurs_runtime.Int((gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}).StrVal())).IntVal) - (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}).StrVal())).IntVal)).IntVal, Get_Data_Enum_toCharCode(), Get_Data_Enum_charToEnum()})))}
	})
	return cache_Data_Enum_boundedEnumChar
}

var cache_Data_Enum_toEnum__1707062200 gopurs_runtime.Value
var once_Data_Enum_toEnum__1707062200 sync.Once
func Get_Data_Enum_toEnum__1707062200() gopurs_runtime.Value {
	once_Data_Enum_toEnum__1707062200.Do(func() {
		cache_Data_Enum_toEnum__1707062200 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__1707062200(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_toEnum__1707062200
}

var cache_Data_Enum_toEnum__488919121 gopurs_runtime.Value
var once_Data_Enum_toEnum__488919121 sync.Once
func Get_Data_Enum_toEnum__488919121() gopurs_runtime.Value {
	once_Data_Enum_toEnum__488919121.Do(func() {
		cache_Data_Enum_toEnum__488919121 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__488919121(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_toEnum__488919121
}

var cache_Data_Enum_toEnum__3448535138 gopurs_runtime.Value
var once_Data_Enum_toEnum__3448535138 sync.Once
func Get_Data_Enum_toEnum__3448535138() gopurs_runtime.Value {
	once_Data_Enum_toEnum__3448535138.Do(func() {
		cache_Data_Enum_toEnum__3448535138 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__3448535138(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_toEnum__3448535138
}

var cache_Data_Enum_toEnum__622773164 gopurs_runtime.Value
var once_Data_Enum_toEnum__622773164 sync.Once
func Get_Data_Enum_toEnum__622773164() gopurs_runtime.Value {
	once_Data_Enum_toEnum__622773164.Do(func() {
		cache_Data_Enum_toEnum__622773164 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__622773164(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_toEnum__622773164
}

var cache_Data_Enum_toEnum__969921489 gopurs_runtime.Value
var once_Data_Enum_toEnum__969921489 sync.Once
func Get_Data_Enum_toEnum__969921489() gopurs_runtime.Value {
	once_Data_Enum_toEnum__969921489.Do(func() {
		cache_Data_Enum_toEnum__969921489 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__969921489(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_toEnum__969921489
}

var cache_Data_Enum_toEnum__1521182434 gopurs_runtime.Value
var once_Data_Enum_toEnum__1521182434 sync.Once
func Get_Data_Enum_toEnum__1521182434() gopurs_runtime.Value {
	once_Data_Enum_toEnum__1521182434.Do(func() {
		cache_Data_Enum_toEnum__1521182434 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__1521182434(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_toEnum__1521182434
}

var cache_Data_Enum_toEnum__2807452764 gopurs_runtime.Value
var once_Data_Enum_toEnum__2807452764 sync.Once
func Get_Data_Enum_toEnum__2807452764() gopurs_runtime.Value {
	once_Data_Enum_toEnum__2807452764.Do(func() {
		cache_Data_Enum_toEnum__2807452764 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_toEnum__2807452764(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_toEnum__2807452764
}

var cache_Data_Enum_succ gopurs_runtime.Value
var once_Data_Enum_succ sync.Once
func Get_Data_Enum_succ() gopurs_runtime.Value {
	once_Data_Enum_succ.Do(func() {
		cache_Data_Enum_succ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_succ(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Enum_succ
}

var cache_Data_Enum_succ__3880853294 gopurs_runtime.Value
var once_Data_Enum_succ__3880853294 sync.Once
func Get_Data_Enum_succ__3880853294() gopurs_runtime.Value {
	once_Data_Enum_succ__3880853294.Do(func() {
		cache_Data_Enum_succ__3880853294 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_succ__3880853294(uint32(__eta_norm_0_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_succ__3880853294
}

var cache_Data_Enum_succ__40000523 gopurs_runtime.Value
var once_Data_Enum_succ__40000523 sync.Once
func Get_Data_Enum_succ__40000523() gopurs_runtime.Value {
	once_Data_Enum_succ__40000523.Do(func() {
		cache_Data_Enum_succ__40000523 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_succ__40000523(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](__eta_norm_0_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_succ__40000523
}

var cache_Data_Enum_succ__2747750594 gopurs_runtime.Value
var once_Data_Enum_succ__2747750594 sync.Once
func Get_Data_Enum_succ__2747750594() gopurs_runtime.Value {
	once_Data_Enum_succ__2747750594.Do(func() {
		cache_Data_Enum_succ__2747750594 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_succ__2747750594(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_succ__2747750594
}

var cache_Data_Enum_succ__3243829713 gopurs_runtime.Value
var once_Data_Enum_succ__3243829713 sync.Once
func Get_Data_Enum_succ__3243829713() gopurs_runtime.Value {
	once_Data_Enum_succ__3243829713.Do(func() {
		cache_Data_Enum_succ__3243829713 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_succ__3243829713(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_succ__3243829713
}

var cache_Data_Enum_upFromIncluding gopurs_runtime.Value
var once_Data_Enum_upFromIncluding sync.Once
func Get_Data_Enum_upFromIncluding() gopurs_runtime.Value {
	once_Data_Enum_upFromIncluding.Do(func() {
		cache_Data_Enum_upFromIncluding = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_upFromIncluding(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
})
	})
	return cache_Data_Enum_upFromIncluding
}

var cache_Data_Enum_showCardinality gopurs_runtime.Value
var once_Data_Enum_showCardinality sync.Once
func Get_Data_Enum_showCardinality() gopurs_runtime.Value {
	once_Data_Enum_showCardinality.Do(func() {
		cache_Data_Enum_showCardinality = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Cardinality ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), v_0).StrVal())) + (")"))
})}))}
	})
	return cache_Data_Enum_showCardinality
}

var cache_Data_Enum_pred gopurs_runtime.Value
var once_Data_Enum_pred sync.Once
func Get_Data_Enum_pred() gopurs_runtime.Value {
	once_Data_Enum_pred.Do(func() {
		cache_Data_Enum_pred = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_pred(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Enum_pred
}

var cache_Data_Enum_pred__3880853294 gopurs_runtime.Value
var once_Data_Enum_pred__3880853294 sync.Once
func Get_Data_Enum_pred__3880853294() gopurs_runtime.Value {
	once_Data_Enum_pred__3880853294.Do(func() {
		cache_Data_Enum_pred__3880853294 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_pred__3880853294(uint32(__eta_norm_0_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_pred__3880853294
}

var cache_Data_Enum_pred__2747750594 gopurs_runtime.Value
var once_Data_Enum_pred__2747750594 sync.Once
func Get_Data_Enum_pred__2747750594() gopurs_runtime.Value {
	once_Data_Enum_pred__2747750594.Do(func() {
		cache_Data_Enum_pred__2747750594 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_pred__2747750594(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_pred__2747750594
}

var cache_Data_Enum_pred__3243829713 gopurs_runtime.Value
var once_Data_Enum_pred__3243829713 sync.Once
func Get_Data_Enum_pred__3243829713() gopurs_runtime.Value {
	once_Data_Enum_pred__3243829713.Do(func() {
		cache_Data_Enum_pred__3243829713 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_pred__3243829713(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_pred__3243829713
}

var cache_Data_Enum_ordCardinality gopurs_runtime.Value
var once_Data_Enum_ordCardinality sync.Once
func Get_Data_Enum_ordCardinality() gopurs_runtime.Value {
	once_Data_Enum_ordCardinality.Do(func() {
		cache_Data_Enum_ordCardinality = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3308271157_4177771502(Rebox_Data_Enum_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
	})
	return cache_Data_Enum_ordCardinality
}

var cache_Data_Enum_newtypeCardinality gopurs_runtime.Value
var once_Data_Enum_newtypeCardinality sync.Once
func Get_Data_Enum_newtypeCardinality() gopurs_runtime.Value {
	once_Data_Enum_newtypeCardinality.Do(func() {
		cache_Data_Enum_newtypeCardinality = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2621187955_385277032((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})))}
	})
	return cache_Data_Enum_newtypeCardinality
}

var cache_Data_Enum_fromEnum gopurs_runtime.Value
var once_Data_Enum_fromEnum sync.Once
func Get_Data_Enum_fromEnum() gopurs_runtime.Value {
	once_Data_Enum_fromEnum.Do(func() {
		cache_Data_Enum_fromEnum = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_fromEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Enum_fromEnum
}

var cache_Data_Enum_fromEnum__3859924869 gopurs_runtime.Value
var once_Data_Enum_fromEnum__3859924869 sync.Once
func Get_Data_Enum_fromEnum__3859924869() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__3859924869.Do(func() {
		cache_Data_Enum_fromEnum__3859924869 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__3859924869(uint32(__eta_norm_0_0_box.IntVal)))
})
	})
	return cache_Data_Enum_fromEnum__3859924869
}

var cache_Data_Enum_fromEnum__3683097841 gopurs_runtime.Value
var once_Data_Enum_fromEnum__3683097841 sync.Once
func Get_Data_Enum_fromEnum__3683097841() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__3683097841.Do(func() {
		cache_Data_Enum_fromEnum__3683097841 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__3683097841(__eta_norm_0_unused_0_box.StrVal()))
})
	})
	return cache_Data_Enum_fromEnum__3683097841
}

var cache_Data_Enum_fromEnum__2502474905 gopurs_runtime.Value
var once_Data_Enum_fromEnum__2502474905 sync.Once
func Get_Data_Enum_fromEnum__2502474905() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__2502474905.Do(func() {
		cache_Data_Enum_fromEnum__2502474905 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__2502474905(__eta_norm_0_0_box.StrVal()))
})
	})
	return cache_Data_Enum_fromEnum__2502474905
}

var cache_Data_Enum_fromEnum__2908037520 gopurs_runtime.Value
var once_Data_Enum_fromEnum__2908037520 sync.Once
func Get_Data_Enum_fromEnum__2908037520() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__2908037520.Do(func() {
		cache_Data_Enum_fromEnum__2908037520 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__2908037520(__eta_norm_0_0_box.IntVal))
})
	})
	return cache_Data_Enum_fromEnum__2908037520
}

var cache_Data_Enum_fromEnum__2720130627 gopurs_runtime.Value
var once_Data_Enum_fromEnum__2720130627 sync.Once
func Get_Data_Enum_fromEnum__2720130627() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__2720130627.Do(func() {
		cache_Data_Enum_fromEnum__2720130627 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__2720130627(__eta_norm_0_0_box.IntVal))
})
	})
	return cache_Data_Enum_fromEnum__2720130627
}

var cache_Data_Enum_fromEnum__3706906321 gopurs_runtime.Value
var once_Data_Enum_fromEnum__3706906321 sync.Once
func Get_Data_Enum_fromEnum__3706906321() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__3706906321.Do(func() {
		cache_Data_Enum_fromEnum__3706906321 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__3706906321(__eta_norm_0_0_box.IntVal))
})
	})
	return cache_Data_Enum_fromEnum__3706906321
}

var cache_Data_Enum_fromEnum__1519505997 gopurs_runtime.Value
var once_Data_Enum_fromEnum__1519505997 sync.Once
func Get_Data_Enum_fromEnum__1519505997() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__1519505997.Do(func() {
		cache_Data_Enum_fromEnum__1519505997 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__1519505997(__eta_norm_0_0_box.IntVal))
})
	})
	return cache_Data_Enum_fromEnum__1519505997
}

var cache_Data_Enum_fromEnum__2871620880 gopurs_runtime.Value
var once_Data_Enum_fromEnum__2871620880 sync.Once
func Get_Data_Enum_fromEnum__2871620880() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__2871620880.Do(func() {
		cache_Data_Enum_fromEnum__2871620880 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__2871620880(__eta_norm_0_0_box.IntVal))
})
	})
	return cache_Data_Enum_fromEnum__2871620880
}

var cache_Data_Enum_fromEnum__992265411 gopurs_runtime.Value
var once_Data_Enum_fromEnum__992265411 sync.Once
func Get_Data_Enum_fromEnum__992265411() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__992265411.Do(func() {
		cache_Data_Enum_fromEnum__992265411 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__992265411(__eta_norm_0_0_box.IntVal))
})
	})
	return cache_Data_Enum_fromEnum__992265411
}

var cache_Data_Enum_fromEnum__3188633789 gopurs_runtime.Value
var once_Data_Enum_fromEnum__3188633789 sync.Once
func Get_Data_Enum_fromEnum__3188633789() gopurs_runtime.Value {
	once_Data_Enum_fromEnum__3188633789.Do(func() {
		cache_Data_Enum_fromEnum__3188633789 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Enum_fromEnum__3188633789(__eta_norm_0_0_box.IntVal))
})
	})
	return cache_Data_Enum_fromEnum__3188633789
}

var cache_Data_Enum_toEnumWithDefaults gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults sync.Once
func Get_Data_Enum_toEnumWithDefaults() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults.Do(func() {
		cache_Data_Enum_toEnumWithDefaults = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_toEnumWithDefaults(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dictBoundedEnum_0_box))
})
	})
	return cache_Data_Enum_toEnumWithDefaults
}

var cache_Data_Enum_toEnumWithDefaults__4115932335 gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults__4115932335 sync.Once
func Get_Data_Enum_toEnumWithDefaults__4115932335() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults__4115932335.Do(func() {
		cache_Data_Enum_toEnumWithDefaults__4115932335 = gopurs_runtime.Func3(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Enum_toEnumWithDefaults__4115932335(__eta_norm_1_unused_0_box.StrVal(), __eta_norm_0_unused_1_box.StrVal(), x_2_box.IntVal))
})
	})
	return cache_Data_Enum_toEnumWithDefaults__4115932335
}

var cache_Data_Enum_toEnumWithDefaults__3037086219 gopurs_runtime.Value
var once_Data_Enum_toEnumWithDefaults__3037086219 sync.Once
func Get_Data_Enum_toEnumWithDefaults__3037086219() gopurs_runtime.Value {
	once_Data_Enum_toEnumWithDefaults__3037086219.Do(func() {
		cache_Data_Enum_toEnumWithDefaults__3037086219 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Enum_toEnumWithDefaults__3037086219(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box.StrVal(), x_2_box.IntVal))
})
	})
	return cache_Data_Enum_toEnumWithDefaults__3037086219
}

var cache_Data_Enum_eqCardinality gopurs_runtime.Value
var once_Data_Enum_eqCardinality sync.Once
func Get_Data_Enum_eqCardinality() gopurs_runtime.Value {
	once_Data_Enum_eqCardinality.Do(func() {
		cache_Data_Enum_eqCardinality = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1053099733_3790796878(Rebox_Data_Enum_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}
	})
	return cache_Data_Enum_eqCardinality
}

var cache_Data_Enum_enumUnit gopurs_runtime.Value
var once_Data_Enum_enumUnit sync.Once
func Get_Data_Enum_enumUnit() gopurs_runtime.Value {
	once_Data_Enum_enumUnit.Do(func() {
		cache_Data_Enum_enumUnit = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Enum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordUnit()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}
})}))}
	})
	return cache_Data_Enum_enumUnit
}

var cache_Data_Enum_enumTuple gopurs_runtime.Value
var once_Data_Enum_enumTuple sync.Once
func Get_Data_Enum_enumTuple() gopurs_runtime.Value {
	once_Data_Enum_enumTuple.Do(func() {
		cache_Data_Enum_enumTuple = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumTuple(dictEnum_0_box)
})
	})
	return cache_Data_Enum_enumTuple
}

var cache_Data_Enum_enumOrdering gopurs_runtime.Value
var once_Data_Enum_enumOrdering sync.Once
func Get_Data_Enum_enumOrdering() gopurs_runtime.Value {
	once_Data_Enum_enumOrdering.Do(func() {
		cache_Data_Enum_enumOrdering = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094((&Constructor_Data_Enum_Enum[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3730953251_4177771502(Rebox_Data_Enum_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordOrdering()))))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just[uint32]
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
_ = __t_tag_0
if (uint32(__t_tag_0) == 1527465420) {
__t3 = Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_3
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_0.IntVal)
_ = __t_tag_1
if (uint32(__t_tag_1) == 902936544) {
__t3 = Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_3
} else {

}
}
{
var __t_tag_2 uint32 = uint32(v_0.IntVal)
_ = __t_tag_2
if (uint32(__t_tag_2) == 380165415) {
__t3 = Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Maybe_Just[uint32] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(__t3))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Maybe_Just[uint32]
{
var __t_tag_4 uint32 = uint32(v_0.IntVal)
_ = __t_tag_4
if (uint32(__t_tag_4) == 1527465420) {
__t7 = Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_7
} else {

}
}
{
var __t_tag_5 uint32 = uint32(v_0.IntVal)
_ = __t_tag_5
if (uint32(__t_tag_5) == 902936544) {
__t7 = Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_7
} else {

}
}
{
var __t_tag_6 uint32 = uint32(v_0.IntVal)
_ = __t_tag_6
if (uint32(__t_tag_6) == 380165415) {
__t7 = Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_Maybe_Just[uint32] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(__t7))}
})})))}
	})
	return cache_Data_Enum_enumOrdering
}

var cache_Data_Enum_enumMaybe gopurs_runtime.Value
var once_Data_Enum_enumMaybe sync.Once
func Get_Data_Enum_enumMaybe() gopurs_runtime.Value {
	once_Data_Enum_enumMaybe.Do(func() {
		cache_Data_Enum_enumMaybe = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumMaybe(dictBoundedEnum_0_box)
})
	})
	return cache_Data_Enum_enumMaybe
}

var cache_Data_Enum_enumInt gopurs_runtime.Value
var once_Data_Enum_enumInt sync.Once
func Get_Data_Enum_enumInt() gopurs_runtime.Value {
	once_Data_Enum_enumInt.Do(func() {
		cache_Data_Enum_enumInt = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4060049525_556578094((&Constructor_Data_Enum_Enum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3308271157_4177771502(Rebox_Data_Enum_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just[int64]
{
if (n_0.IntVal) > (Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3764732725_2094947566(Rebox_Data_Enum_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedInt()))))}).IntVal) {
__t0 = Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int((n_0.IntVal) - (int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_0
} else {

}
}
{
__t0 = Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(__t0))}
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[int64]
{
if (n_0.IntVal) < (gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal).IntVal) {
__t1 = Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int((n_0.IntVal) + (int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_1
} else {

}
}
{
__t1 = Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(__t1))}
})})))}
	})
	return cache_Data_Enum_enumInt
}

var cache_Data_Enum_enumFromTo gopurs_runtime.Value
var once_Data_Enum_enumFromTo sync.Once
func Get_Data_Enum_enumFromTo() gopurs_runtime.Value {
	once_Data_Enum_enumFromTo.Do(func() {
		cache_Data_Enum_enumFromTo = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumFromTo(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
})
	})
	return cache_Data_Enum_enumFromTo
}

var cache_Data_Enum_enumFromThenTo gopurs_runtime.Value
var once_Data_Enum_enumFromThenTo sync.Once
func Get_Data_Enum_enumFromThenTo() gopurs_runtime.Value {
	once_Data_Enum_enumFromThenTo.Do(func() {
		cache_Data_Enum_enumFromThenTo = gopurs_runtime.Func3(func(dictUnfoldable_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictBoundedEnum_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumFromThenTo(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](dictBoundedEnum_2_box))
})
	})
	return cache_Data_Enum_enumFromThenTo
}

var cache_Data_Enum_enumEither gopurs_runtime.Value
var once_Data_Enum_enumEither sync.Once
func Get_Data_Enum_enumEither() gopurs_runtime.Value {
	once_Data_Enum_enumEither.Do(func() {
		cache_Data_Enum_enumEither = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_enumEither(dictBoundedEnum_0_box)
})
	})
	return cache_Data_Enum_enumEither
}

var cache_Data_Enum_enumBoolean gopurs_runtime.Value
var once_Data_Enum_enumBoolean sync.Once
func Get_Data_Enum_enumBoolean() gopurs_runtime.Value {
	once_Data_Enum_enumBoolean.Do(func() {
		cache_Data_Enum_enumBoolean = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_342165130_556578094((&Constructor_Data_Enum_Enum[bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_219188042_4177771502(Rebox_Data_Enum_4177771502_219188042(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordBoolean()))))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just[bool]
{
if (v_0.IntVal) != (0) {
__t0 = Rebox_Data_Enum_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Bool(false), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_0
} else {

}
}
{
__t0 = Rebox_Data_Enum_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1044698560_3094389156(__t0))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[bool]
{
if ((v_0.IntVal) != (0)) != (true) {
__t1 = Rebox_Data_Enum_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Bool(true), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_1
} else {

}
}
{
__t1 = Rebox_Data_Enum_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1044698560_3094389156(__t1))}
})})))}
	})
	return cache_Data_Enum_enumBoolean
}

var cache_Data_Enum_downFromIncluding gopurs_runtime.Value
var once_Data_Enum_downFromIncluding sync.Once
func Get_Data_Enum_downFromIncluding() gopurs_runtime.Value {
	once_Data_Enum_downFromIncluding.Do(func() {
		cache_Data_Enum_downFromIncluding = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_downFromIncluding(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
})
	})
	return cache_Data_Enum_downFromIncluding
}

var cache_Data_Enum_diag gopurs_runtime.Value
var once_Data_Enum_diag sync.Once
func Get_Data_Enum_diag() gopurs_runtime.Value {
	once_Data_Enum_diag.Do(func() {
		cache_Data_Enum_diag = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_diag(a_0_box)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
})
	})
	return cache_Data_Enum_diag
}

var cache_Data_Enum_downFrom gopurs_runtime.Value
var once_Data_Enum_downFrom sync.Once
func Get_Data_Enum_downFrom() gopurs_runtime.Value {
	once_Data_Enum_downFrom.Do(func() {
		cache_Data_Enum_downFrom = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_downFrom(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
})
	})
	return cache_Data_Enum_downFrom
}

var cache_Data_Enum_upFrom gopurs_runtime.Value
var once_Data_Enum_upFrom sync.Once
func Get_Data_Enum_upFrom() gopurs_runtime.Value {
	once_Data_Enum_upFrom.Do(func() {
		cache_Data_Enum_upFrom = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_upFrom(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
})
	})
	return cache_Data_Enum_upFrom
}

var cache_Data_Enum_defaultToEnum gopurs_runtime.Value
var once_Data_Enum_defaultToEnum sync.Once
func Get_Data_Enum_defaultToEnum() gopurs_runtime.Value {
	once_Data_Enum_defaultToEnum.Do(func() {
		cache_Data_Enum_defaultToEnum = gopurs_runtime.Func3(func(dictBounded_0_box gopurs_runtime.Value, dictEnum_1_box gopurs_runtime.Value, i_prime__2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultToEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](dictBounded_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_1_box), i_prime__2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_defaultToEnum
}

var cache_Data_Enum_defaultSucc gopurs_runtime.Value
var once_Data_Enum_defaultSucc sync.Once
func Get_Data_Enum_defaultSucc() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc.Do(func() {
		cache_Data_Enum_defaultSucc = gopurs_runtime.Func3(func(toEnum_prime__0_box gopurs_runtime.Value, fromEnum_prime__1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultSucc(toEnum_prime__0_box, fromEnum_prime__1_box, a_2_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_defaultSucc
}

var cache_Data_Enum_defaultSucc__2406132389 gopurs_runtime.Value
var once_Data_Enum_defaultSucc__2406132389 sync.Once
func Get_Data_Enum_defaultSucc__2406132389() gopurs_runtime.Value {
	once_Data_Enum_defaultSucc__2406132389.Do(func() {
		cache_Data_Enum_defaultSucc__2406132389 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultSucc__2406132389(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_defaultSucc__2406132389
}

var cache_Data_Enum_defaultPred gopurs_runtime.Value
var once_Data_Enum_defaultPred sync.Once
func Get_Data_Enum_defaultPred() gopurs_runtime.Value {
	once_Data_Enum_defaultPred.Do(func() {
		cache_Data_Enum_defaultPred = gopurs_runtime.Func3(func(toEnum_prime__0_box gopurs_runtime.Value, fromEnum_prime__1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultPred(toEnum_prime__0_box, fromEnum_prime__1_box, a_2_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_defaultPred
}

var cache_Data_Enum_defaultPred__2406132389 gopurs_runtime.Value
var once_Data_Enum_defaultPred__2406132389 sync.Once
func Get_Data_Enum_defaultPred__2406132389() gopurs_runtime.Value {
	once_Data_Enum_defaultPred__2406132389.Do(func() {
		cache_Data_Enum_defaultPred__2406132389 = gopurs_runtime.Func3(func(toEnum_prime__unused_0_box gopurs_runtime.Value, fromEnum_prime__unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Enum_defaultPred__2406132389(toEnum_prime__unused_0_box, fromEnum_prime__unused_1_box, a_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Enum_defaultPred__2406132389
}

var cache_Data_Enum_defaultFromEnum gopurs_runtime.Value
var once_Data_Enum_defaultFromEnum sync.Once
func Get_Data_Enum_defaultFromEnum() gopurs_runtime.Value {
	once_Data_Enum_defaultFromEnum.Do(func() {
		cache_Data_Enum_defaultFromEnum = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_defaultFromEnum(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](dictEnum_0_box))
})
	})
	return cache_Data_Enum_defaultFromEnum
}

var cache_Data_Enum_defaultCardinality gopurs_runtime.Value
var once_Data_Enum_defaultCardinality sync.Once
func Get_Data_Enum_defaultCardinality() gopurs_runtime.Value {
	once_Data_Enum_defaultCardinality.Do(func() {
		cache_Data_Enum_defaultCardinality = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_defaultCardinality(dictBounded_0_box)
})
	})
	return cache_Data_Enum_defaultCardinality
}

var cache_Data_Enum_cardinality gopurs_runtime.Value
var once_Data_Enum_cardinality sync.Once
func Get_Data_Enum_cardinality() gopurs_runtime.Value {
	once_Data_Enum_cardinality.Do(func() {
		cache_Data_Enum_cardinality = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_cardinality(dict_0_box)
})
	})
	return cache_Data_Enum_cardinality
}

var cache_Data_Enum_boundedEnumUnit gopurs_runtime.Value
var once_Data_Enum_boundedEnumUnit sync.Once
func Get_Data_Enum_boundedEnumUnit() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumUnit.Do(func() {
		cache_Data_Enum_boundedEnumUnit = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedUnit()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Enum_enumUnit()))}
}), gopurs_runtime.Int(int64(1)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(int64(0))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_0.IntVal) == (int64(0)) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{Get_Data_Unit_unit(), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})}))}
	})
	return cache_Data_Enum_boundedEnumUnit
}

var cache_Data_Enum_boundedEnumOrdering gopurs_runtime.Value
var once_Data_Enum_boundedEnumOrdering sync.Once
func Get_Data_Enum_boundedEnumOrdering() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumOrdering.Do(func() {
		cache_Data_Enum_boundedEnumOrdering = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4021906832_123048125((&Constructor_Data_Enum_BoundedEnum[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_832288803_2094947566(Rebox_Data_Enum_2094947566_832288803(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedOrdering()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2359585123_556578094(Rebox_Data_Enum_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Enum_enumOrdering()))))}
}), gopurs_runtime.Int(int64(3)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 int64
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
_ = __t_tag_0
if (uint32(__t_tag_0) == 1527465420) {
__t3 = int64(0)
goto end_branch_3
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_0.IntVal)
_ = __t_tag_1
if (uint32(__t_tag_1) == 902936544) {
__t3 = int64(1)
goto end_branch_3
} else {

}
}
{
var __t_tag_2 uint32 = uint32(v_0.IntVal)
_ = __t_tag_2
if (uint32(__t_tag_2) == 380165415) {
__t3 = int64(2)
goto end_branch_3
} else {

}
}
{
__t3 = func() int64 { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Int(__t3)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just[uint32]
{
if (v_0.IntVal) == (int64(0)) {
__t4 = Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_4
} else {

}
}
{
if (v_0.IntVal) == (int64(1)) {
__t4 = Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_4
} else {

}
}
{
if (v_0.IntVal) == (int64(2)) {
__t4 = Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_4
} else {

}
}
{
__t4 = Rebox_Data_Enum_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_622082505_3094389156(__t4))}
})})))}
	})
	return cache_Data_Enum_boundedEnumOrdering
}

var cache_Data_Enum_boundedEnumBoolean gopurs_runtime.Value
var once_Data_Enum_boundedEnumBoolean sync.Once
func Get_Data_Enum_boundedEnumBoolean() gopurs_runtime.Value {
	once_Data_Enum_boundedEnumBoolean.Do(func() {
		cache_Data_Enum_boundedEnumBoolean = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3083804185_123048125((&Constructor_Data_Enum_BoundedEnum[bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3728870730_2094947566(Rebox_Data_Enum_2094947566_3728870730(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedBoolean()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_342165130_556578094(Rebox_Data_Enum_556578094_342165130(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Enum_enumBoolean()))))}
}), gopurs_runtime.Int(int64(2)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 int64
{
if ((v_0.IntVal) != (0)) != (true) {
__t0 = int64(0)
goto end_branch_0
} else {

}
}
{
if (v_0.IntVal) != (0) {
__t0 = int64(1)
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[bool]
{
if (v_0.IntVal) == (int64(0)) {
__t1 = Rebox_Data_Enum_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Bool(false), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_1
} else {

}
}
{
if (v_0.IntVal) == (int64(1)) {
__t1 = Rebox_Data_Enum_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Bool(true), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_1
} else {

}
}
{
__t1 = Rebox_Data_Enum_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1044698560_3094389156(__t1))}
})})))}
	})
	return cache_Data_Enum_boundedEnumBoolean
}

type Constructor_Data_Enum_Enum[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4075786298] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_Enum[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Ord0": return gopurs_runtime.Box(c.V0)
		case "pred": return gopurs_runtime.Box(c.V1)
		case "succ": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Enum_Enum: " + key)
		}
	}
}


type Constructor_Data_Enum_BoundedEnum[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 int64
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[287434377] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Bounded0": return gopurs_runtime.Box(c.V0)
		case "Enum1": return gopurs_runtime.Box(c.V1)
		case "cardinality": return gopurs_runtime.Box(c.V2)
		case "fromEnum": return gopurs_runtime.Box(c.V3)
		case "toEnum": return gopurs_runtime.Box(c.V4)
		default: panic("Key not found in dictionary Constructor_Data_Enum_BoundedEnum: " + key)
		}
	}
}


func Call_Data_Enum_Enum_dollar_Dict(x_0_loop struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
var x_0 struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
				}())
}

func Call_Data_Enum_Enum_dollar_Dict__2787884439(x_0_loop struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}) *Constructor_Data_Enum_Enum[uint32] {
Enum_dollar_Dict__2787884439:
for {
if false { continue Enum_dollar_Dict__2787884439 }
var x_0 struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[uint32]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
				}())
}
}

func Call_Data_Enum_Enum_dollar_Dict__479058327(x_0_loop struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}) *Constructor_Data_Enum_Enum[uint32] {
Enum_dollar_Dict__479058327:
for {
if false { continue Enum_dollar_Dict__479058327 }
var x_0 struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[uint32]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
				}())
}
}

func Call_Data_Enum_Enum_dollar_Dict__3712690007(x_0_loop struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}) *Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date] {
Enum_dollar_Dict__3712690007:
for {
if false { continue Enum_dollar_Dict__3712690007 }
var x_0 struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
				}())
}
}

func Call_Data_Enum_Enum_dollar_Dict__2123951319(x_0_loop struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}) *Constructor_Data_Enum_Enum[uint32] {
Enum_dollar_Dict__2123951319:
for {
if false { continue Enum_dollar_Dict__2123951319 }
var x_0 struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[uint32]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
				}())
}
}

func Call_Data_Enum_Enum_dollar_Dict__1439164631(x_0_loop struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}) *Constructor_Data_Enum_Enum[bool] {
Enum_dollar_Dict__1439164631:
for {
if false { continue Enum_dollar_Dict__1439164631 }
var x_0 struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[bool]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
				}())
}
}

func Call_Data_Enum_Enum_dollar_Dict__2220371927(x_0_loop struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}) *Constructor_Data_Enum_Enum[string] {
Enum_dollar_Dict__2220371927:
for {
if false { continue Enum_dollar_Dict__2220371927 }
var x_0 struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[string]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
				}())
}
}

func Call_Data_Enum_Enum_dollar_Dict__3674609207(x_0_loop struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}) *Constructor_Data_Enum_Enum[int64] {
Enum_dollar_Dict__3674609207:
for {
if false { continue Enum_dollar_Dict__3674609207 }
var x_0 struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[int64]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
				}())
}
}

func Call_Data_Enum_Enum_dollar_Dict__4138904023(x_0_loop struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
Enum_dollar_Dict__4138904023:
for {
if false { continue Enum_dollar_Dict__4138904023 }
var x_0 struct{
	Ord0 gopurs_runtime.Value
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Ord0", "pred", "succ", orig.Ord0, orig.pred, orig.succ)
				}())
}
}

func Call_Data_Enum_Cardinality(x_0_loop int64) gopurs_runtime.Value {
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}

func Call_Data_Enum_Cardinality__1400572777(x_0_loop int64) gopurs_runtime.Value {
Cardinality__1400572777:
for {
if false { continue Cardinality__1400572777 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__2274214265(x_0_loop int64) gopurs_runtime.Value {
Cardinality__2274214265:
for {
if false { continue Cardinality__2274214265 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__123016565(x_0_loop int64) gopurs_runtime.Value {
Cardinality__123016565:
for {
if false { continue Cardinality__123016565 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__1784762877(x_0_loop int64) gopurs_runtime.Value {
Cardinality__1784762877:
for {
if false { continue Cardinality__1784762877 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__4276004359(x_0_loop int64) gopurs_runtime.Value {
Cardinality__4276004359:
for {
if false { continue Cardinality__4276004359 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__3890726945(x_unused_0_loop int64) gopurs_runtime.Value {
Cardinality__3890726945:
for {
if false { continue Cardinality__3890726945 }
var x_unused_0 int64 = x_unused_0_loop
_ = x_unused_0
return gopurs_runtime.Int((gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}).StrVal())).IntVal) - (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}).StrVal())).IntVal))
}
}

func Call_Data_Enum_Cardinality__2587234619(x_0_loop int64) gopurs_runtime.Value {
Cardinality__2587234619:
for {
if false { continue Cardinality__2587234619 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__4110088112(x_0_loop int64) gopurs_runtime.Value {
Cardinality__4110088112:
for {
if false { continue Cardinality__4110088112 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_Cardinality__4100296677(x_0_loop int64) gopurs_runtime.Value {
Cardinality__4100296677:
for {
if false { continue Cardinality__4100296677 }
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Int(x_0)
}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict(x_0_loop struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
var x_0 struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
				}())
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__4237532871(x_0_loop struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}) *Constructor_Data_Enum_BoundedEnum[uint32] {
BoundedEnum_dollar_Dict__4237532871:
for {
if false { continue BoundedEnum_dollar_Dict__4237532871 }
var x_0 struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
				}())
}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__2081071559(x_0_loop struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}) *Constructor_Data_Enum_BoundedEnum[uint32] {
BoundedEnum_dollar_Dict__2081071559:
for {
if false { continue BoundedEnum_dollar_Dict__2081071559 }
var x_0 struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
				}())
}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__2130790279(x_0_loop struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}) *Constructor_Data_Enum_BoundedEnum[uint32] {
BoundedEnum_dollar_Dict__2130790279:
for {
if false { continue BoundedEnum_dollar_Dict__2130790279 }
var x_0 struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
				}())
}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__1429463943(x_0_loop struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}) *Constructor_Data_Enum_BoundedEnum[bool] {
BoundedEnum_dollar_Dict__1429463943:
for {
if false { continue BoundedEnum_dollar_Dict__1429463943 }
var x_0 struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[bool]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
				}())
}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__882291527(x_0_loop struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}) *Constructor_Data_Enum_BoundedEnum[string] {
BoundedEnum_dollar_Dict__882291527:
for {
if false { continue BoundedEnum_dollar_Dict__882291527 }
var x_0 struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[string]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
				}())
}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__2897860935(x_0_loop struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}) *Constructor_Data_Enum_BoundedEnum[int64] {
BoundedEnum_dollar_Dict__2897860935:
for {
if false { continue BoundedEnum_dollar_Dict__2897860935 }
var x_0 struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
				}())
}
}

func Call_Data_Enum_BoundedEnum_dollar_Dict__3240687239(x_0_loop struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
BoundedEnum_dollar_Dict__3240687239:
for {
if false { continue BoundedEnum_dollar_Dict__3240687239 }
var x_0 struct{
	Bounded0 gopurs_runtime.Value
	Enum1 gopurs_runtime.Value
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", orig.Bounded0, orig.Enum1, orig.cardinality, orig.fromEnum, orig.toEnum)
				}())
}
}

func Call_Data_Enum_toEnum(dict_0_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_Data_Enum_toEnum__3016447460(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
toEnum__3016447460:
for {
if false { continue toEnum__3016447460 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumMonth())).V4, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_toEnum__3403550180(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
toEnum__3403550180:
for {
if false { continue toEnum__3403550180 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumWeekday())).V4, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_charToEnum(v_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just[string]
{
if ((v_0) >= (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}).StrVal())).IntVal)) && ((v_0) <= (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}).StrVal())).IntVal)) {
__t0 = Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Enum_fromCharCode(), gopurs_runtime.Int(v_0)).StrVal()), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_0
} else {

}
}
{
__t0 = Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(__t0))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_defaultSucc__3855932694(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
defaultSucc__3855932694:
for {
if false { continue defaultSucc__3855932694 }
var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
_ = toEnum_prime__unused_0
var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
_ = fromEnum_prime__unused_1
var a_2 string = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum((gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(a_2)).IntVal) + (int64(1)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_defaultPred__3855932694(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
defaultPred__3855932694:
for {
if false { continue defaultPred__3855932694 }
var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
_ = toEnum_prime__unused_0
var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
_ = fromEnum_prime__unused_1
var a_2 string = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum((gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(a_2)).IntVal) - (int64(1)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_toEnum__1707062200(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
toEnum__1707062200:
for {
if false { continue toEnum__1707062200 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_742090555_3094389156(Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum(__eta_norm_0_0)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_toEnum__488919121(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
toEnum__488919121:
for {
if false { continue toEnum__488919121 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumDay())).V4, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_toEnum__3448535138(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
toEnum__3448535138:
for {
if false { continue toEnum__3448535138 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumYear())).V4, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_toEnum__622773164(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
toEnum__622773164:
for {
if false { continue toEnum__622773164 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumHour())).V4, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_toEnum__969921489(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
toEnum__969921489:
for {
if false { continue toEnum__969921489 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMillisecond())).V4, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_toEnum__1521182434(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
toEnum__1521182434:
for {
if false { continue toEnum__1521182434 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMinute())).V4, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_toEnum__2807452764(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
toEnum__2807452764:
for {
if false { continue toEnum__2807452764 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumSecond())).V4, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_succ(dict_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_Data_Enum_succ__3880853294(__eta_norm_0_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
succ__3880853294:
for {
if false { continue succ__3880853294 }
var __eta_norm_0_0 uint32 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumMonth())).V2, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta_norm_0_0), UnsafePtr: nil})))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_succ__40000523(__eta_norm_0_0_loop *Constructor_Data_Date_Date) struct{V0 gopurs_runtime.Value; V1 bool} {
succ__40000523:
for {
if false { continue succ__40000523 }
var __eta_norm_0_0 *Constructor_Data_Date_Date = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_3494563625(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_enumDate())).V2, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(__eta_norm_0_0)})))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_succ__2747750594(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
succ__2747750594:
for {
if false { continue succ__2747750594 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumDay())).V2, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_succ__3243829713(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
succ__3243829713:
for {
if false { continue succ__3243829713 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V2, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_upFromIncluding(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
// TAST (Let): succ1_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope4)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope4)]))
succ1_1_0 := Call_Data_Enum_succ(dictEnum_0)
_ = succ1_1_0
return gopurs_runtime.Func(func(dictUnfoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_3, gopurs_runtime.Apply(succ1_1_0, x_3)}))}
}))
})
}

func Call_Data_Enum_pred(dict_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_Data_Enum_pred__3880853294(__eta_norm_0_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
pred__3880853294:
for {
if false { continue pred__3880853294 }
var __eta_norm_0_0 uint32 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumMonth())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta_norm_0_0), UnsafePtr: nil})))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_pred__2747750594(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
pred__2747750594:
for {
if false { continue pred__2747750594 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumDay())).V1, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_pred__3243829713(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
pred__3243829713:
for {
if false { continue pred__3243829713 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Enum_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V1, gopurs_runtime.Int(__eta_norm_0_0))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_fromEnum(dict_0_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_Data_Enum_fromEnum__3859924869(__eta_norm_0_0_loop uint32) int64 {
fromEnum__3859924869:
for {
if false { continue fromEnum__3859924869 }
var __eta_norm_0_0 uint32 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_4021906832(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumMonth())).V3, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta_norm_0_0), UnsafePtr: nil}).IntVal
}
}

func Call_Data_Enum_fromEnum__3683097841(__eta_norm_0_unused_0_loop string) int64 {
fromEnum__3683097841:
for {
if false { continue fromEnum__3683097841 }
var __eta_norm_0_unused_0 string = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}).StrVal())).IntVal
}
}

func Call_Data_Enum_fromEnum__2502474905(__eta_norm_0_0_loop string) int64 {
fromEnum__2502474905:
for {
if false { continue fromEnum__2502474905 }
var __eta_norm_0_0 string = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(__eta_norm_0_0)).IntVal
}
}

func Call_Data_Enum_fromEnum__2908037520(__eta_norm_0_0_loop int64) int64 {
fromEnum__2908037520:
for {
if false { continue fromEnum__2908037520 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumDay())).V3, gopurs_runtime.Int(__eta_norm_0_0)).IntVal
}
}

func Call_Data_Enum_fromEnum__2720130627(__eta_norm_0_0_loop int64) int64 {
fromEnum__2720130627:
for {
if false { continue fromEnum__2720130627 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumYear())).V3, gopurs_runtime.Int(__eta_norm_0_0)).IntVal
}
}

func Call_Data_Enum_fromEnum__3706906321(__eta_norm_0_0_loop int64) int64 {
fromEnum__3706906321:
for {
if false { continue fromEnum__3706906321 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())).V3, gopurs_runtime.Int(__eta_norm_0_0)).IntVal
}
}

func Call_Data_Enum_fromEnum__1519505997(__eta_norm_0_0_loop int64) int64 {
fromEnum__1519505997:
for {
if false { continue fromEnum__1519505997 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumHour())).V3, gopurs_runtime.Int(__eta_norm_0_0)).IntVal
}
}

func Call_Data_Enum_fromEnum__2871620880(__eta_norm_0_0_loop int64) int64 {
fromEnum__2871620880:
for {
if false { continue fromEnum__2871620880 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMillisecond())).V3, gopurs_runtime.Int(__eta_norm_0_0)).IntVal
}
}

func Call_Data_Enum_fromEnum__992265411(__eta_norm_0_0_loop int64) int64 {
fromEnum__992265411:
for {
if false { continue fromEnum__992265411 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumMinute())).V3, gopurs_runtime.Int(__eta_norm_0_0)).IntVal
}
}

func Call_Data_Enum_fromEnum__3188633789(__eta_norm_0_0_loop int64) int64 {
fromEnum__3188633789:
for {
if false { continue fromEnum__3188633789 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Time_Component_boundedEnumSecond())).V3, gopurs_runtime.Int(__eta_norm_0_0)).IntVal
}
}

func Call_Data_Enum_toEnumWithDefaults(dictBoundedEnum_0_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBoundedEnum_0 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
// TAST (Let): bottom2_1_0 shape=App(Var) bindingType=(TypeVar a$scope21)
bottom2_1_0 := Call_Data_Bounded_bottom(gopurs_runtime.Apply(dictBoundedEnum_0.V0, gopurs_runtime.Value{}))
_ = bottom2_1_0
return gopurs_runtime.Func3(func(low_2 gopurs_runtime.Value, high_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope21)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(dictBoundedEnum_0.V4, gopurs_runtime.Int(x_4.IntVal)))
_ = v_5_1
var __t3 gopurs_runtime.Value
{
if (v_5_1 != nil) {
__t3 = (v_5_1).V0
goto end_branch_3
} else {

}
}
{
if (v_5_1 == nil) {
var __t2 gopurs_runtime.Value
{
if (x_4.IntVal) < (gopurs_runtime.Apply(dictBoundedEnum_0.V3, bottom2_1_0).IntVal) {
__t2 = low_2
goto end_branch_2
} else {

}
}
{
__t2 = high_3
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
}

func Call_Data_Enum_toEnumWithDefaults__4115932335(__eta_norm_1_unused_0_loop string, __eta_norm_0_unused_1_loop string, x_2_loop int64) string {
toEnumWithDefaults__4115932335:
for {
if false { continue toEnumWithDefaults__4115932335 }
var __eta_norm_1_unused_0 string = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_unused_1 string = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
var x_2 int64 = x_2_loop
_ = x_2
// TAST (Let): v_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Char])
v_3_0 := Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum(x_2)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
_ = v_3_0
var __t2 string
{
if (v_3_0 != nil) {
__t2 = (v_3_0).V0
goto end_branch_2
} else {

}
}
{
if (v_3_0 == nil) {
var __t1 string
{
if Call_Data_Ord_lessThan__4092919914(x_2, gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}).StrVal())).IntVal) {
__t1 = gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()).StrVal()
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal()).StrVal()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() string { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}

func Call_Data_Enum_toEnumWithDefaults__3037086219(__eta_norm_1_0_loop string, __eta_norm_0_1_loop string, x_2_loop int64) string {
toEnumWithDefaults__3037086219:
for {
if false { continue toEnumWithDefaults__3037086219 }
var __eta_norm_1_0 string = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 string = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var x_2 int64 = x_2_loop
_ = x_2
// TAST (Let): v_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Char])
v_3_0 := Rebox_Data_Enum_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_Enum_charToEnum(x_2)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
_ = v_3_0
var __t2 string
{
if (v_3_0 != nil) {
__t2 = (v_3_0).V0
goto end_branch_2
} else {

}
}
{
if (v_3_0 == nil) {
var __t1 string
{
if Call_Data_Ord_lessThan__4092919914(x_2, gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2420955921_2094947566(Rebox_Data_Enum_2094947566_2420955921(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Bounded_boundedChar()))))}).StrVal())).IntVal) {
__t1 = __eta_norm_1_0
goto end_branch_1
} else {

}
}
{
__t1 = __eta_norm_0_1
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() string { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}

func Call_Data_Enum_enumTuple(dictEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEnum_0 gopurs_runtime.Value = dictEnum_0_loop
_ = dictEnum_0
// TAST (Let): ordTuple_1_0 shape=App(Var) bindingType=Any
ordTuple_1_0 := Call_Data_Tuple_ordTuple(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "Ord0"), gopurs_runtime.Value{}))
_ = ordTuple_1_0
return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bounded0_3_1 shape=App(Other) bindingType=(ADT ["Data","Bounded","Bounded"] [(TypeVar b$scope30)])
Bounded0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_3_1
// TAST (Let): Enum1_4_2 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar b$scope30)])
Enum1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_4_2
// TAST (Let): ordTuple1_5_3 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar b$scope30)])])
ordTuple1_5_3 := Rebox_Data_Enum_4177771502_1535415139(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordTuple_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{}))))
_ = ordTuple1_5_3
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3481421219_556578094((&Constructor_Data_Enum_Enum[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1535415139_4177771502(ordTuple1_5_3))}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope29)])
__local_var_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))
_ = __local_var_7_5
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_7_5 != nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_7_5).V0, Bounded0_3_1.V2}))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_6:
// TAST (Let): __local_var_7_4 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar b$scope30)])])
var __local_var_7_4 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Data_Enum_3094389156_4010058633(__t6)
// TAST (Let): __local_var_8_7 shape=App(Var) bindingType=(Func [(TypeVar b$scope30)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar b$scope30)])]))
__local_var_8_7 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))
_ = __local_var_8_7
// TAST (Let): __local_var_9_8 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope30)])
__local_var_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum1_4_2.V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
_ = __local_var_9_8
var __t9 gopurs_runtime.Value
{
if (__local_var_9_8 == nil) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4010058633_3094389156(__local_var_7_4))}
goto end_branch_9
} else {

}
}
{
if (__local_var_9_8 != nil) {
__t9 = gopurs_runtime.Apply(__local_var_8_7, (__local_var_9_8).V0)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4010058633_3094389156(Rebox_Data_Enum_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t9))))}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_11 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope29)])
__local_var_7_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))
_ = __local_var_7_11
var __t12 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_7_11 != nil) {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_7_11).V0, Bounded0_3_1.V1}))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_12:
// TAST (Let): __local_var_7_10 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar b$scope30)])])
var __local_var_7_10 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = Rebox_Data_Enum_3094389156_4010058633(__t12)
// TAST (Let): __local_var_8_13 shape=App(Var) bindingType=(Func [(TypeVar b$scope30)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope29), (TypeVar b$scope30)])]))
__local_var_8_13 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0))
_ = __local_var_8_13
// TAST (Let): __local_var_9_14 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope30)])
__local_var_9_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum1_4_2.V2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
_ = __local_var_9_14
var __t15 gopurs_runtime.Value
{
if (__local_var_9_14 == nil) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4010058633_3094389156(__local_var_7_10))}
goto end_branch_15
} else {

}
}
{
if (__local_var_9_14 != nil) {
__t15 = gopurs_runtime.Apply(__local_var_8_13, (__local_var_9_14).V0)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_4010058633_3094389156(Rebox_Data_Enum_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t15))))}
})})))}
})
}

func Call_Data_Enum_enumMaybe(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
// TAST (Let): Bounded0_1_0 shape=App(Other) bindingType=(ADT ["Data","Bounded","Bounded"] [(TypeVar a$scope52)])
Bounded0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_1_0
// TAST (Let): Enum1_2_1 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar a$scope52)])
Enum1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_2_1
// TAST (Let): ordMaybe_3_2 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope52)])])
ordMaybe_3_2 := Rebox_Data_Enum_4177771502_2155612431(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Maybe_ordMaybe(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{}))))
_ = ordMaybe_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2522609743_556578094((&Constructor_Data_Enum_Enum[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_2155612431_4177771502(ordMaybe_3_2))}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4)
_ = __t_tag_3
if (__t_tag_3 == nil) {
__t5 = Rebox_Data_Enum_3094389156_215731685(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4)
_ = __t_tag_4
if (__t_tag_4 != nil) {
__t5 = Rebox_Data_Enum_3094389156_215731685(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(Enum1_2_1.V1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_215731685_3094389156(__t5))}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4)
_ = __t_tag_6
if (__t_tag_6 == nil) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_215731685_3094389156(Rebox_Data_Enum_3094389156_215731685(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, Bounded0_1_0.V1}))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
goto end_branch_10
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_4)
_ = __t_tag_7
if (__t_tag_7 != nil) {
// TAST (Let): __local_var_5_8 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope52)])
__local_var_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum1_2_1.V2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_4.UnsafePtr).V0))
_ = __local_var_5_8
var __t9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_8 != nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (__local_var_5_8).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_9:
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_215731685_3094389156(func() *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] { panic("Failed pattern match") }()))}
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_215731685_3094389156(Rebox_Data_Enum_3094389156_215731685(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t10))))}
})})))}
}

func Call_Data_Enum_enumFromTo(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
// TAST (Let): Ord0_1_0 shape=App(Other) bindingType=Any
Ord0_1_0 := gopurs_runtime.Apply(dictEnum_0.V0, gopurs_runtime.Value{})
_ = Ord0_1_0
// TAST (Let): Eq0_2_1 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeVar a$scope66)])
Eq0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_1_0, "Eq0"), gopurs_runtime.Value{}))
_ = Eq0_2_1
// TAST (Let): Ord01_3_2 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(TypeVar a$scope66)])
Ord01_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(dictEnum_0.V0, gopurs_runtime.Value{}))
_ = Ord01_3_2
// TAST (Let): succ1_4_3 shape=App(Var) bindingType=(Func [(TypeVar a$scope66)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope66)]))
succ1_4_3 := Call_Data_Enum_succ(dictEnum_0)
_ = succ1_4_3
// TAST (Let): pred1_5_4 shape=App(Var) bindingType=(Func [(TypeVar a$scope66)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope66)]))
pred1_5_4 := Call_Data_Enum_pred(dictEnum_0)
_ = pred1_5_4
return gopurs_runtime.Func3(func(dictUnfoldable1_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Eq0_2_1.V0, v_7, v1_8).IntVal) != (0) {
__t15 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_6, "unfoldr1"), gopurs_runtime.Func(func(i_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]
{
if (i_9.IntVal) <= (int64(0)) {
__t9 = Rebox_Data_Enum_138441832_3418986898(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{v_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))
goto end_branch_9
} else {

}
}
{
__t9 = Rebox_Data_Enum_138441832_3418986898(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{v_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int((i_9.IntVal) - (int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3418986898_138441832(__t9))}
}), gopurs_runtime.Int(int64(0)))
goto end_branch_15
} else {

}
}
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Apply2(Ord01_3_2.V1, v_7, v1_8)
_ = __t_tag_10
if (uint32(__t_tag_10.IntVal) == 1527465420) {
__t15 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_6, "unfoldr1"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 shape=App(Other) bindingType=Any
__local_var_10_11 := gopurs_runtime.Apply(succ1_4_3, a_9)
_ = __local_var_10_11
var __t14 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_10_11.Type == 9 && __local_var_10_11.IntVal == 930809136 && __local_var_10_11.UnsafePtr != nil) {
var __t13 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_12 struct{V0 gopurs_runtime.Value; V1 bool} = Call_Control_Alternative_guard__3603952853(Call_Data_Ord_lessThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Ord0_1_0), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_10_11.UnsafePtr).V0, v1_8))
_ = __t_tag_12
if __t_tag_12.V1 {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_10_11.UnsafePtr).V0, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
if (__local_var_10_11.Type == 9 && __local_var_10_11.IntVal == 930809136 && __local_var_10_11.UnsafePtr == nil) {
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_14
} else {

}
}
{
__t14 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3804580809_138441832(Rebox_Data_Enum_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_9, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t14)}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
}), v_7)
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_6, "unfoldr1"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_5 shape=App(Other) bindingType=Any
__local_var_10_5 := gopurs_runtime.Apply(pred1_5_4, a_9)
_ = __local_var_10_5
var __t8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_10_5.Type == 9 && __local_var_10_5.IntVal == 930809136 && __local_var_10_5.UnsafePtr != nil) {
var __t7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_6 struct{V0 gopurs_runtime.Value; V1 bool} = Call_Control_Alternative_guard__3603952853(Call_Data_Ord_greaterThanOrEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Ord0_1_0), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_10_5.UnsafePtr).V0, v1_8))
_ = __t_tag_6
if __t_tag_6.V1 {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_10_5.UnsafePtr).V0, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
if (__local_var_10_5.Type == 9 && __local_var_10_5.IntVal == 930809136 && __local_var_10_5.UnsafePtr == nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3804580809_138441832(Rebox_Data_Enum_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_9, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
}), v_7)
}
end_branch_15:
return __t15
})
}

func Call_Data_Enum_enumFromThenTo(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value], dictFunctor_1_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], dictBoundedEnum_2_loop *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var dictFunctor_1 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
var dictBoundedEnum_2 *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] = dictBoundedEnum_2_loop
_ = dictBoundedEnum_2
// TAST (Let): toEnum1_3_0 shape=App(Var) bindingType=(Func [Int] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope73)]))
toEnum1_3_0 := Call_Data_Enum_toEnum(dictBoundedEnum_2)
_ = toEnum1_3_0
return gopurs_runtime.Func3(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value, c_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): a_prime__7_1 shape=App(Other) bindingType=Int
a_prime__7_1 := gopurs_runtime.Apply(dictBoundedEnum_2.V3, a_4).IntVal
_ = a_prime__7_1
// TAST (Let): __local_var_8_2 shape=Other bindingType=Int
__local_var_8_2 := (gopurs_runtime.Apply(dictBoundedEnum_2.V3, b_5).IntVal) - (a_prime__7_1)
_ = __local_var_8_2
// TAST (Let): __local_var_9_3 shape=App(Other) bindingType=Int
__local_var_9_3 := gopurs_runtime.Apply(dictBoundedEnum_2.V3, c_6).IntVal
_ = __local_var_9_3
return gopurs_runtime.Apply2(dictFunctor_1.V0, Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), toEnum1_3_0, gopurs_runtime.Apply(Get_Data_Maybe_fromJust(), gopurs_runtime.Value{})), gopurs_runtime.Apply2(dictUnfoldable_0.V1, gopurs_runtime.Func(func(e_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]
{
if (e_10.IntVal) <= (__local_var_9_3) {
__t4 = Rebox_Data_Enum_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3363075976_138441832(Rebox_Data_Enum_138441832_3363075976(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(e_10.IntVal), gopurs_runtime.Int((e_10.IntVal) + (__local_var_8_2))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_4
} else {

}
}
{
__t4 = Rebox_Data_Enum_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1415037225_3094389156(__t4))}
}), gopurs_runtime.Int(a_prime__7_1)))
})
}

func Call_Data_Enum_enumEither(dictBoundedEnum_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBoundedEnum_0 gopurs_runtime.Value = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
// TAST (Let): Enum1_1_0 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar a$scope81)])
Enum1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_1_0
// TAST (Let): Bounded0_2_1 shape=App(Other) bindingType=(ADT ["Data","Bounded","Bounded"] [(TypeVar a$scope81)])
Bounded0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_2_1
// TAST (Let): ordEither_3_2 shape=App(Var) bindingType=Any
ordEither_3_2 := Call_Data_Either_ordEither(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{}))
_ = ordEither_3_2
return gopurs_runtime.Func(func(dictBoundedEnum1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bounded01_5_3 shape=App(Other) bindingType=(ADT ["Data","Bounded","Bounded"] [(TypeVar b$scope82)])
Bounded01_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded01_5_3
// TAST (Let): Enum11_6_4 shape=App(Other) bindingType=(ADT ["Data","Enum","Enum"] [(TypeVar b$scope82)])
Enum11_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Enum1"), gopurs_runtime.Value{}))
_ = Enum11_6_4
// TAST (Let): ordEither1_7_5 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope81), (TypeVar b$scope82)])])
ordEither1_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordEither_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Enum1"), gopurs_runtime.Value{}), "Ord0"), gopurs_runtime.Value{})))
_ = ordEither1_7_5
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Enum_Enum[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordEither1_7_5)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_8.Type == 9 && v_8.IntVal == 3711209382) {
// TAST (Let): __local_var_9_6 shape=App(Var) bindingType=(Func [(TypeVar a$scope81)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope81), (TypeVar b$scope82)])]))
__local_var_9_6 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Left())
_ = __local_var_9_6
// TAST (Let): __local_var_10_7 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope81)])
__local_var_10_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum1_1_0.V1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
_ = __local_var_10_7
var __t8 gopurs_runtime.Value
{
if (__local_var_10_7 == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_8
} else {

}
}
{
if (__local_var_10_7 != nil) {
__t8 = gopurs_runtime.Apply(__local_var_9_6, (__local_var_10_7).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t8)
goto end_branch_12
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 2465973597) {
// TAST (Let): __local_var_9_9 shape=App(Var) bindingType=(Func [(TypeVar b$scope82)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope81), (TypeVar b$scope82)])]))
__local_var_9_9 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Right())
_ = __local_var_9_9
// TAST (Let): __local_var_10_10 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope82)])
__local_var_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum11_6_4.V1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
_ = __local_var_10_10
var __t11 gopurs_runtime.Value
{
if (__local_var_10_10 == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, Bounded0_2_1.V2}))}}))}
goto end_branch_11
} else {

}
}
{
if (__local_var_10_10 != nil) {
__t11 = gopurs_runtime.Apply(__local_var_9_9, (__local_var_10_10).V0)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t11)
goto end_branch_12
} else {

}
}
{
__t12 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t12)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_8.Type == 9 && v_8.IntVal == 3711209382) {
// TAST (Let): __local_var_9_13 shape=App(Var) bindingType=(Func [(TypeVar a$scope81)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope81), (TypeVar b$scope82)])]))
__local_var_9_13 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Left())
_ = __local_var_9_13
// TAST (Let): __local_var_10_14 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope81)])
__local_var_10_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum1_1_0.V2, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
_ = __local_var_10_14
var __t15 gopurs_runtime.Value
{
if (__local_var_10_14 == nil) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, Bounded01_5_3.V1}))}}))}
goto end_branch_15
} else {

}
}
{
if (__local_var_10_14 != nil) {
__t15 = gopurs_runtime.Apply(__local_var_9_13, (__local_var_10_14).V0)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t15)
goto end_branch_19
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 2465973597) {
// TAST (Let): __local_var_9_16 shape=App(Var) bindingType=(Func [(TypeVar b$scope82)] (ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Either","Either"] [(TypeVar a$scope81), (TypeVar b$scope82)])]))
__local_var_9_16 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Right())
_ = __local_var_9_16
// TAST (Let): __local_var_10_17 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope82)])
__local_var_10_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Enum11_6_4.V2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
_ = __local_var_10_17
var __t18 gopurs_runtime.Value
{
if (__local_var_10_17 == nil) {
__t18 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_18
} else {

}
}
{
if (__local_var_10_17 != nil) {
__t18 = gopurs_runtime.Apply(__local_var_9_16, (__local_var_10_17).V0)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t18)
goto end_branch_19
} else {

}
}
{
__t19 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t19)}
})}))}
})
}

func Call_Data_Enum_downFromIncluding(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
// TAST (Let): pred1_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope115)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope115)]))
pred1_1_0 := Call_Data_Enum_pred(dictEnum_0)
_ = pred1_1_0
return gopurs_runtime.Func(func(dictUnfoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_3, gopurs_runtime.Apply(pred1_1_0, x_3)}))}
}))
})
}

func Call_Data_Enum_diag(a_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_0, a_0}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Data_Enum_downFrom(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
// TAST (Let): pred1_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope125)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope125)]))
pred1_1_0 := Call_Data_Enum_pred(dictEnum_0)
_ = pred1_1_0
return gopurs_runtime.Func(func(dictUnfoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_2, "unfoldr"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{func() gopurs_runtime.Value {
				_v := Call_Data_Enum_diag((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}(), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), pred1_1_0))
})
}

func Call_Data_Enum_upFrom(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
// TAST (Let): succ1_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope130)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope130)]))
succ1_1_0 := Call_Data_Enum_succ(dictEnum_0)
_ = succ1_1_0
return gopurs_runtime.Func(func(dictUnfoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_2, "unfoldr"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{func() gopurs_runtime.Value {
				_v := Call_Data_Enum_diag((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}(), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), succ1_1_0))
})
}

func Call_Data_Enum_defaultToEnum(dictBounded_0_loop *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value], dictEnum_1_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value], i_prime__2_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictBounded_0 *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] = dictBounded_0_loop
_ = dictBounded_0
var dictEnum_1 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_1_loop
_ = dictEnum_1
var i_prime__2 int64 = i_prime__2_loop
_ = i_prime__2
var Call_local_Data_Enum_go__go_3_0_0 func(int64, gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_Enum_go__go_3_0_0
var go__go_3_0_0 gopurs_runtime.Value
_ = go__go_3_0_0
Call_local_Data_Enum_go__go_3_0_0 = func(i_4_loop int64, x_5_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__go_3_0_0:
for {
if false { continue go__go_3_0_0 }
var i_4 int64 = i_4_loop
_ = i_4
var x_5 gopurs_runtime.Value = x_5_loop
_ = x_5
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (i_4) == (int64(0)) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{x_5, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_3
} else {

}
}
{
// TAST (Let): v_6_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope134)])
v_6_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(dictEnum_1.V2, x_5))
_ = v_6_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_6_1 != nil) {
i_4_loop = (i_4) - (int64(1))
x_5_loop = (v_6_1).V0
continue go__go_3_0_0
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v_6_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return __t3
}
}
go__go_3_0_0 = gopurs_runtime.Func(func(i_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_Enum_go__go_3_0_0(i_4_loop_val.IntVal, x_5_loop_val))}
})
})
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (i_prime__2) < (int64(0)) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_4
} else {

}
}
{
__t4 = Call_local_Data_Enum_go__go_3_0_0(i_prime__2, dictBounded_0.V1)
}
end_branch_4:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_defaultSucc(toEnum_prime__0_loop gopurs_runtime.Value, fromEnum_prime__1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var toEnum_prime__0 gopurs_runtime.Value = toEnum_prime__0_loop
_ = toEnum_prime__0
var fromEnum_prime__1 gopurs_runtime.Value = fromEnum_prime__1_loop
_ = fromEnum_prime__1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime__0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime__1, a_2).IntVal) + (int64(1))))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_defaultSucc__2406132389(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
defaultSucc__2406132389:
for {
if false { continue defaultSucc__2406132389 }
var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
_ = toEnum_prime__unused_0
var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
_ = fromEnum_prime__unused_1
var a_2 int64 = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Enum_toEnum(Rebox_Data_Enum_1306125126_123048125(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())))), gopurs_runtime.Int((gopurs_runtime.Apply(Call_Data_Enum_fromEnum(Rebox_Data_Enum_1306125126_123048125(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())))), gopurs_runtime.Int(a_2)).IntVal) + (int64(1))))))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_defaultPred(toEnum_prime__0_loop gopurs_runtime.Value, fromEnum_prime__1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var toEnum_prime__0 gopurs_runtime.Value = toEnum_prime__0_loop
_ = toEnum_prime__0
var fromEnum_prime__1 gopurs_runtime.Value = fromEnum_prime__1_loop
_ = fromEnum_prime__1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime__0, gopurs_runtime.Int((gopurs_runtime.Apply(fromEnum_prime__1, a_2).IntVal) - (int64(1))))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Enum_defaultPred__2406132389(toEnum_prime__unused_0_loop gopurs_runtime.Value, fromEnum_prime__unused_1_loop gopurs_runtime.Value, a_2_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
defaultPred__2406132389:
for {
if false { continue defaultPred__2406132389 }
var toEnum_prime__unused_0 gopurs_runtime.Value = toEnum_prime__unused_0_loop
_ = toEnum_prime__unused_0
var fromEnum_prime__unused_1 gopurs_runtime.Value = fromEnum_prime__unused_1_loop
_ = fromEnum_prime__unused_1
var a_2 int64 = a_2_loop
_ = a_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Enum_toEnum(Rebox_Data_Enum_1306125126_123048125(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())))), gopurs_runtime.Int((gopurs_runtime.Apply(Call_Data_Enum_fromEnum(Rebox_Data_Enum_1306125126_123048125(Rebox_Data_Enum_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_String_CodePoints_boundedEnumCodePoint())))), gopurs_runtime.Int(a_2)).IntVal) - (int64(1))))))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Enum_defaultFromEnum(dictEnum_0_loop *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *Constructor_Data_Enum_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
var Call_local_Data_Enum_go__118764402_1_0_1 func(int64, gopurs_runtime.Value) int64
_ = Call_local_Data_Enum_go__118764402_1_0_1
var go__118764402_1_0_1 gopurs_runtime.Value
_ = go__118764402_1_0_1
var Call_local_Data_Enum_go__go_1_1_2 func(int64, gopurs_runtime.Value) int64
_ = Call_local_Data_Enum_go__go_1_1_2
var go__go_1_1_2 gopurs_runtime.Value
_ = go__go_1_1_2
Call_local_Data_Enum_go__118764402_1_0_1 = func(i_2_loop int64, x_3_loop gopurs_runtime.Value) int64 {
go__118764402_1_0_1:
for {
if false { continue go__118764402_1_0_1 }
var i_2 int64 = i_2_loop
_ = i_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
// TAST (Let): v_4_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope143)])
v_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(dictEnum_0.V1, x_3))
_ = v_4_2
var __t3 int64
{
if (v_4_2 != nil) {
i_2_loop = (i_2) + (int64(1))
x_3_loop = (v_4_2).V0
continue go__118764402_1_0_1
__t3 = func() int64 { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
if (v_4_2 == nil) {
__t3 = i_2
goto end_branch_3
} else {

}
}
{
__t3 = func() int64 { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__118764402_1_0_1 = gopurs_runtime.Func(func(i_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Data_Enum_go__118764402_1_0_1(i_2_loop_val.IntVal, x_3_loop_val))
})
})
Call_local_Data_Enum_go__go_1_1_2 = func(i_2_loop int64, x_3_loop gopurs_runtime.Value) int64 {
go__go_1_1_2:
for {
if false { continue go__go_1_1_2 }
var i_2 int64 = i_2_loop
_ = i_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
// TAST (Let): v_4_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope143)])
v_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(dictEnum_0.V1, x_3))
_ = v_4_4
var __t5 int64
{
if (v_4_4 != nil) {
__t5 = Call_local_Data_Enum_go__118764402_1_0_1((i_2) + (int64(1)), (v_4_4).V0)
goto end_branch_5
} else {

}
}
{
if (v_4_4 == nil) {
__t5 = i_2
goto end_branch_5
} else {

}
}
{
__t5 = func() int64 { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_1_1_2 = gopurs_runtime.Func(func(i_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Data_Enum_go__go_1_1_2(i_2_loop_val.IntVal, x_3_loop_val))
})
})
return gopurs_runtime.Apply(go__118764402_1_0_1, gopurs_runtime.Int(int64(0)))
}

func Call_Data_Enum_defaultCardinality(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): bottom2_1_0 shape=App(Var) bindingType=(TypeVar a$scope145)
bottom2_1_0 := Call_Data_Bounded_bottom(dictBounded_0)
_ = bottom2_1_0
return gopurs_runtime.Func(func(dictEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Enum_go__118764402_3_1_3 func(int64, gopurs_runtime.Value) int64
_ = Call_local_Data_Enum_go__118764402_3_1_3
var go__118764402_3_1_3 gopurs_runtime.Value
_ = go__118764402_3_1_3
var Call_local_Data_Enum_go__go_3_2_4 func(int64, gopurs_runtime.Value) int64
_ = Call_local_Data_Enum_go__go_3_2_4
var go__go_3_2_4 gopurs_runtime.Value
_ = go__go_3_2_4
Call_local_Data_Enum_go__118764402_3_1_3 = func(i_4_loop int64, x_5_loop gopurs_runtime.Value) int64 {
go__118764402_3_1_3:
for {
if false { continue go__118764402_3_1_3 }
var i_4 int64 = i_4_loop
_ = i_4
var x_5 gopurs_runtime.Value = x_5_loop
_ = x_5
// TAST (Let): v_6_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Any])
v_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_5))
_ = v_6_3
var __t4 int64
{
if (v_6_3 != nil) {
i_4_loop = (i_4) + (int64(1))
x_5_loop = (v_6_3).V0
continue go__118764402_3_1_3
__t4 = func() int64 { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
if (v_6_3 == nil) {
__t4 = i_4
goto end_branch_4
} else {

}
}
{
__t4 = func() int64 { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__118764402_3_1_3 = gopurs_runtime.Func(func(i_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Data_Enum_go__118764402_3_1_3(i_4_loop_val.IntVal, x_5_loop_val))
})
})
Call_local_Data_Enum_go__go_3_2_4 = func(i_4_loop int64, x_5_loop gopurs_runtime.Value) int64 {
go__go_3_2_4:
for {
if false { continue go__go_3_2_4 }
var i_4 int64 = i_4_loop
_ = i_4
var x_5 gopurs_runtime.Value = x_5_loop
_ = x_5
// TAST (Let): v_6_5 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope145)])
v_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_5))
_ = v_6_5
var __t6 int64
{
if (v_6_5 != nil) {
__t6 = Call_local_Data_Enum_go__118764402_3_1_3((i_4) + (int64(1)), (v_6_5).V0)
goto end_branch_6
} else {

}
}
{
if (v_6_5 == nil) {
__t6 = i_4
goto end_branch_6
} else {

}
}
{
__t6 = func() int64 { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}
go__go_3_2_4 = gopurs_runtime.Func(func(i_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Data_Enum_go__go_3_2_4(i_4_loop_val.IntVal, x_5_loop_val))
})
})
return gopurs_runtime.Int(Call_local_Data_Enum_go__118764402_3_1_3(int64(1), bottom2_1_0))
})
}

func Call_Data_Enum_cardinality(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.Int(gopurs_runtime.RecordGet(dict_0, "cardinality").IntVal)
}

func Rebox_Data_Enum_1044698560_3094389156(in *Constructor_Data_Maybe_Just[bool]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Bool(in.V0)
	return out
}

func Rebox_Data_Enum_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Enum_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Enum_123048125_1306125126(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Enum_123048125_4021906832(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[uint32]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Enum_1306125126_123048125(in *Constructor_Data_Enum_BoundedEnum[int64]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Enum_138441832_3363075976(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[int64, int64]{}
		out.V0 = in.V0.IntVal
		out.V1 = in.V1.IntVal
	return out
}

func Rebox_Data_Enum_138441832_3418986898(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]{}
		out.V0 = in.V0
		out.V1 = Rebox_Data_Enum_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1))
	return out
}

func Rebox_Data_Enum_138441832_3804580809(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Data_Enum_1415037225_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_3363075976_138441832(in.V0))}
	return out
}

func Rebox_Data_Enum_1535415139_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_2094947566_2420955921(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[string]{}
		out.V0 = in.V0
		out.V1 = in.V1.StrVal()
		out.V2 = in.V2.StrVal()
	return out
}

func Rebox_Data_Enum_2094947566_3728870730(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[bool] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[bool]{}
		out.V0 = in.V0
		out.V1 = (in.V1.IntVal) != (0)
		out.V2 = (in.V2.IntVal) != (0)
	return out
}

func Rebox_Data_Enum_2094947566_3764732725(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1.IntVal
		out.V2 = in.V2.IntVal
	return out
}

func Rebox_Data_Enum_2094947566_832288803(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[uint32]{}
		out.V0 = in.V0
		out.V1 = uint32(in.V1.IntVal)
		out.V2 = uint32(in.V2.IntVal)
	return out
}

func Rebox_Data_Enum_2155612431_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_215731685_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Enum_219188042_4177771502(in *Constructor_Data_Ord_Ord[bool]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_2359585123_556578094(in *Constructor_Data_Enum_Enum[uint32]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_2420955921_2094947566(in *Constructor_Data_Bounded_Bounded[string]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Str(in.V1)
		out.V2 = gopurs_runtime.Str(in.V2)
	return out
}

func Rebox_Data_Enum_2522609743_556578094(in *Constructor_Data_Enum_Enum[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_2621187955_385277032(in *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, int64]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Enum_3083804185_123048125(in *Constructor_Data_Enum_BoundedEnum[bool]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Enum_3094389156_1044698560(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[bool] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[bool]{}
		out.V0 = (in.V0.IntVal) != (0)
	return out
}

func Rebox_Data_Enum_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Enum_3094389156_1415037225(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]{}
		out.V0 = Rebox_Data_Enum_138441832_3363075976(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Enum_3094389156_215731685(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_Enum_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_Enum_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[uint32]{}
		out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Data_Enum_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[string]{}
		out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_Enum_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_3363075976_138441832(in *Constructor_Data_Tuple_Tuple[int64, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Data_Enum_3409577041_556578094(in *Constructor_Data_Enum_Enum[string]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_3418986898_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Enum_1170268447_3094389156(in.V1))}
	return out
}

func Rebox_Data_Enum_342165130_556578094(in *Constructor_Data_Enum_Enum[bool]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_3481421219_556578094(in *Constructor_Data_Enum_Enum[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_3494563625_556578094(in *Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_3569500834_123048125(in *Constructor_Data_Enum_BoundedEnum[string]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Enum_3728870730_2094947566(in *Constructor_Data_Bounded_Bounded[bool]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Bool(in.V1)
		out.V2 = gopurs_runtime.Bool(in.V2)
	return out
}

func Rebox_Data_Enum_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_3764732725_2094947566(in *Constructor_Data_Bounded_Bounded[int64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Int(in.V1)
		out.V2 = gopurs_runtime.Int(in.V2)
	return out
}

func Rebox_Data_Enum_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[int64]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Enum_3804580809_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Enum_4010058633_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Enum_4021906832_123048125(in *Constructor_Data_Enum_BoundedEnum[uint32]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Enum_4060049525_556578094(in *Constructor_Data_Enum_Enum[int64]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_4177771502_1535415139(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_4177771502_2155612431(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_4177771502_219188042(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[bool] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[bool]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_4177771502_2406510097(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[string]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_4177771502_3730953251(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[uint32]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Enum_556578094_2359585123(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[uint32]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_556578094_3409577041(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[string]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_556578094_342165130(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[bool] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[bool]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_556578094_3494563625(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_556578094_4060049525(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Enum_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Data_Enum_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Rebox_Data_Enum_832288803_2094947566(in *Constructor_Data_Bounded_Bounded[uint32]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
	return out
}

func Get_Data_Enum_fromCharCode() gopurs_runtime.Value {
	return _Gopurs_Data_Enum_FromCharCode
}

func Get_Data_Enum_toCharCode() gopurs_runtime.Value {
	return _Gopurs_Data_Enum_ToCharCode
}
