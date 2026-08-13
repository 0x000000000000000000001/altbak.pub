package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Applicative_Applicative_dollarDict gopurs_runtime.Value
var once_Control_Applicative_Applicative_dollarDict sync.Once
func Get_Control_Applicative_Applicative_dollarDict() gopurs_runtime.Value {
	once_Control_Applicative_Applicative_dollarDict.Do(func() {
		cache_Control_Applicative_Applicative_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_Applicative_dollarDict(x_0_box)
})
	})
	return cache_Control_Applicative_Applicative_dollarDict
}

var cache_Control_Applicative_pure gopurs_runtime.Value
var once_Control_Applicative_pure sync.Once
func Get_Control_Applicative_pure() gopurs_runtime.Value {
	once_Control_Applicative_pure.Do(func() {
		cache_Control_Applicative_pure = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure
}

var cache_Control_Applicative_unless gopurs_runtime.Value
var once_Control_Applicative_unless sync.Once
func Get_Control_Applicative_unless() gopurs_runtime.Value {
	once_Control_Applicative_unless.Do(func() {
		cache_Control_Applicative_unless = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_unless(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_Control_Applicative_unless
}

var cache_Control_Applicative_when gopurs_runtime.Value
var once_Control_Applicative_when sync.Once
func Get_Control_Applicative_when() gopurs_runtime.Value {
	once_Control_Applicative_when.Do(func() {
		cache_Control_Applicative_when = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_when(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_Control_Applicative_when
}

var cache_Control_Applicative_liftA1 gopurs_runtime.Value
var once_Control_Applicative_liftA1 sync.Once
func Get_Control_Applicative_liftA1() gopurs_runtime.Value {
	once_Control_Applicative_liftA1.Do(func() {
		cache_Control_Applicative_liftA1 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Control_Applicative_liftA1
}

var cache_Control_Applicative_applicativeProxy gopurs_runtime.Value
var once_Control_Applicative_applicativeProxy sync.Once
func Get_Control_Applicative_applicativeProxy() gopurs_runtime.Value {
	once_Control_Applicative_applicativeProxy.Do(func() {
		cache_Control_Applicative_applicativeProxy = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Control_Apply_applyProxy()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})})}
	})
	return cache_Control_Applicative_applicativeProxy
}

var cache_Control_Applicative_applicativeFn gopurs_runtime.Value
var once_Control_Applicative_applicativeFn sync.Once
func Get_Control_Applicative_applicativeFn() gopurs_runtime.Value {
	once_Control_Applicative_applicativeFn.Do(func() {
		cache_Control_Applicative_applicativeFn = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Control_Apply_applyFn()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
})})}
	})
	return cache_Control_Applicative_applicativeFn
}

var cache_Control_Applicative_applicativeArray gopurs_runtime.Value
var once_Control_Applicative_applicativeArray sync.Once
func Get_Control_Applicative_applicativeArray() gopurs_runtime.Value {
	once_Control_Applicative_applicativeArray.Do(func() {
		cache_Control_Applicative_applicativeArray = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Control_Apply_applyArray()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})})}
	})
	return cache_Control_Applicative_applicativeArray
}

var cache_Control_Applicative_applicativeFn__1805574895 gopurs_runtime.Value
var once_Control_Applicative_applicativeFn__1805574895 sync.Once
func Get_Control_Applicative_applicativeFn__1805574895() gopurs_runtime.Value {
	once_Control_Applicative_applicativeFn__1805574895.Do(func() {
		cache_Control_Applicative_applicativeFn__1805574895 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Control_Apply_applyFn()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
})})}
	})
	return cache_Control_Applicative_applicativeFn__1805574895
}

var cache_Control_Applicative_pure__2382779990 gopurs_runtime.Value
var once_Control_Applicative_pure__2382779990 sync.Once
func Get_Control_Applicative_pure__2382779990() gopurs_runtime.Value {
	once_Control_Applicative_pure__2382779990.Do(func() {
		cache_Control_Applicative_pure__2382779990 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2382779990(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__2382779990
}

var cache_Control_Applicative_pure__1812053569 gopurs_runtime.Value
var once_Control_Applicative_pure__1812053569 sync.Once
func Get_Control_Applicative_pure__1812053569() gopurs_runtime.Value {
	once_Control_Applicative_pure__1812053569.Do(func() {
		cache_Control_Applicative_pure__1812053569 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1812053569(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__1812053569
}

var cache_Control_Applicative_pure__2491902360 gopurs_runtime.Value
var once_Control_Applicative_pure__2491902360 sync.Once
func Get_Control_Applicative_pure__2491902360() gopurs_runtime.Value {
	once_Control_Applicative_pure__2491902360.Do(func() {
		cache_Control_Applicative_pure__2491902360 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2491902360(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__2491902360
}

var cache_Control_Applicative_pure__149572504 gopurs_runtime.Value
var once_Control_Applicative_pure__149572504 sync.Once
func Get_Control_Applicative_pure__149572504() gopurs_runtime.Value {
	once_Control_Applicative_pure__149572504.Do(func() {
		cache_Control_Applicative_pure__149572504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__149572504(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__149572504
}

var cache_Control_Applicative_pure__189931222 gopurs_runtime.Value
var once_Control_Applicative_pure__189931222 sync.Once
func Get_Control_Applicative_pure__189931222() gopurs_runtime.Value {
	once_Control_Applicative_pure__189931222.Do(func() {
		cache_Control_Applicative_pure__189931222 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__189931222(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__189931222
}

var cache_Control_Applicative_pure__779048344 gopurs_runtime.Value
var once_Control_Applicative_pure__779048344 sync.Once
func Get_Control_Applicative_pure__779048344() gopurs_runtime.Value {
	once_Control_Applicative_pure__779048344.Do(func() {
		cache_Control_Applicative_pure__779048344 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__779048344(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__779048344
}

var cache_Control_Applicative_pure__4037597590 gopurs_runtime.Value
var once_Control_Applicative_pure__4037597590 sync.Once
func Get_Control_Applicative_pure__4037597590() gopurs_runtime.Value {
	once_Control_Applicative_pure__4037597590.Do(func() {
		cache_Control_Applicative_pure__4037597590 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__4037597590(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__4037597590
}

var cache_Control_Applicative_pure__3432560598 gopurs_runtime.Value
var once_Control_Applicative_pure__3432560598 sync.Once
func Get_Control_Applicative_pure__3432560598() gopurs_runtime.Value {
	once_Control_Applicative_pure__3432560598.Do(func() {
		cache_Control_Applicative_pure__3432560598 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3432560598(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__3432560598
}

var cache_Control_Applicative_pure__3236307030 gopurs_runtime.Value
var once_Control_Applicative_pure__3236307030 sync.Once
func Get_Control_Applicative_pure__3236307030() gopurs_runtime.Value {
	once_Control_Applicative_pure__3236307030.Do(func() {
		cache_Control_Applicative_pure__3236307030 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3236307030(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__3236307030
}

var cache_Control_Applicative_pure__2331489366 gopurs_runtime.Value
var once_Control_Applicative_pure__2331489366 sync.Once
func Get_Control_Applicative_pure__2331489366() gopurs_runtime.Value {
	once_Control_Applicative_pure__2331489366.Do(func() {
		cache_Control_Applicative_pure__2331489366 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2331489366(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__2331489366
}

var cache_Control_Applicative_pure__2302010678 gopurs_runtime.Value
var once_Control_Applicative_pure__2302010678 sync.Once
func Get_Control_Applicative_pure__2302010678() gopurs_runtime.Value {
	once_Control_Applicative_pure__2302010678.Do(func() {
		cache_Control_Applicative_pure__2302010678 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2302010678(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__2302010678
}

var cache_Control_Applicative_pure__993904534 gopurs_runtime.Value
var once_Control_Applicative_pure__993904534 sync.Once
func Get_Control_Applicative_pure__993904534() gopurs_runtime.Value {
	once_Control_Applicative_pure__993904534.Do(func() {
		cache_Control_Applicative_pure__993904534 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__993904534(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__993904534
}

var cache_Control_Applicative_pure__2935994064 gopurs_runtime.Value
var once_Control_Applicative_pure__2935994064 sync.Once
func Get_Control_Applicative_pure__2935994064() gopurs_runtime.Value {
	once_Control_Applicative_pure__2935994064.Do(func() {
		cache_Control_Applicative_pure__2935994064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2935994064(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__2935994064
}

var cache_Control_Applicative_pure__3215807376 gopurs_runtime.Value
var once_Control_Applicative_pure__3215807376 sync.Once
func Get_Control_Applicative_pure__3215807376() gopurs_runtime.Value {
	once_Control_Applicative_pure__3215807376.Do(func() {
		cache_Control_Applicative_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3215807376(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__3215807376
}

var cache_Control_Applicative_pure__2960485136 gopurs_runtime.Value
var once_Control_Applicative_pure__2960485136 sync.Once
func Get_Control_Applicative_pure__2960485136() gopurs_runtime.Value {
	once_Control_Applicative_pure__2960485136.Do(func() {
		cache_Control_Applicative_pure__2960485136 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2960485136(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__2960485136
}

var cache_Control_Applicative_pure__160425008 gopurs_runtime.Value
var once_Control_Applicative_pure__160425008 sync.Once
func Get_Control_Applicative_pure__160425008() gopurs_runtime.Value {
	once_Control_Applicative_pure__160425008.Do(func() {
		cache_Control_Applicative_pure__160425008 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__160425008(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__160425008
}

var cache_Control_Applicative_pure__3286817552 gopurs_runtime.Value
var once_Control_Applicative_pure__3286817552 sync.Once
func Get_Control_Applicative_pure__3286817552() gopurs_runtime.Value {
	once_Control_Applicative_pure__3286817552.Do(func() {
		cache_Control_Applicative_pure__3286817552 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3286817552(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__3286817552
}

var cache_Control_Applicative_pure__763072784 gopurs_runtime.Value
var once_Control_Applicative_pure__763072784 sync.Once
func Get_Control_Applicative_pure__763072784() gopurs_runtime.Value {
	once_Control_Applicative_pure__763072784.Do(func() {
		cache_Control_Applicative_pure__763072784 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__763072784(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__763072784
}

var cache_Control_Applicative_pure__871290128 gopurs_runtime.Value
var once_Control_Applicative_pure__871290128 sync.Once
func Get_Control_Applicative_pure__871290128() gopurs_runtime.Value {
	once_Control_Applicative_pure__871290128.Do(func() {
		cache_Control_Applicative_pure__871290128 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__871290128(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__871290128
}

var cache_Control_Applicative_pure__1748760400 gopurs_runtime.Value
var once_Control_Applicative_pure__1748760400 sync.Once
func Get_Control_Applicative_pure__1748760400() gopurs_runtime.Value {
	once_Control_Applicative_pure__1748760400.Do(func() {
		cache_Control_Applicative_pure__1748760400 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1748760400(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__1748760400
}

var cache_Control_Applicative_pure__3012389648 gopurs_runtime.Value
var once_Control_Applicative_pure__3012389648 sync.Once
func Get_Control_Applicative_pure__3012389648() gopurs_runtime.Value {
	once_Control_Applicative_pure__3012389648.Do(func() {
		cache_Control_Applicative_pure__3012389648 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3012389648(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__3012389648
}

var cache_Control_Applicative_pure__566620048 gopurs_runtime.Value
var once_Control_Applicative_pure__566620048 sync.Once
func Get_Control_Applicative_pure__566620048() gopurs_runtime.Value {
	once_Control_Applicative_pure__566620048.Do(func() {
		cache_Control_Applicative_pure__566620048 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__566620048(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__566620048
}

var cache_Control_Applicative_pure__3197665392 gopurs_runtime.Value
var once_Control_Applicative_pure__3197665392 sync.Once
func Get_Control_Applicative_pure__3197665392() gopurs_runtime.Value {
	once_Control_Applicative_pure__3197665392.Do(func() {
		cache_Control_Applicative_pure__3197665392 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3197665392(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__3197665392
}

var cache_Control_Applicative_pure__1475749520 gopurs_runtime.Value
var once_Control_Applicative_pure__1475749520 sync.Once
func Get_Control_Applicative_pure__1475749520() gopurs_runtime.Value {
	once_Control_Applicative_pure__1475749520.Do(func() {
		cache_Control_Applicative_pure__1475749520 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1475749520(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__1475749520
}

var cache_Control_Applicative_pure__1253336208 gopurs_runtime.Value
var once_Control_Applicative_pure__1253336208 sync.Once
func Get_Control_Applicative_pure__1253336208() gopurs_runtime.Value {
	once_Control_Applicative_pure__1253336208.Do(func() {
		cache_Control_Applicative_pure__1253336208 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1253336208(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__1253336208
}

var cache_Control_Applicative_pure__355615152 gopurs_runtime.Value
var once_Control_Applicative_pure__355615152 sync.Once
func Get_Control_Applicative_pure__355615152() gopurs_runtime.Value {
	once_Control_Applicative_pure__355615152.Do(func() {
		cache_Control_Applicative_pure__355615152 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__355615152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__355615152
}

var cache_Control_Applicative_pure__1741581456 gopurs_runtime.Value
var once_Control_Applicative_pure__1741581456 sync.Once
func Get_Control_Applicative_pure__1741581456() gopurs_runtime.Value {
	once_Control_Applicative_pure__1741581456.Do(func() {
		cache_Control_Applicative_pure__1741581456 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1741581456(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__1741581456
}

var cache_Control_Applicative_pure__243192752 gopurs_runtime.Value
var once_Control_Applicative_pure__243192752 sync.Once
func Get_Control_Applicative_pure__243192752() gopurs_runtime.Value {
	once_Control_Applicative_pure__243192752.Do(func() {
		cache_Control_Applicative_pure__243192752 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__243192752(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__243192752
}

var cache_Control_Applicative_pure__4233214992 gopurs_runtime.Value
var once_Control_Applicative_pure__4233214992 sync.Once
func Get_Control_Applicative_pure__4233214992() gopurs_runtime.Value {
	once_Control_Applicative_pure__4233214992.Do(func() {
		cache_Control_Applicative_pure__4233214992 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__4233214992(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__4233214992
}

var cache_Control_Applicative_pure__154576880 gopurs_runtime.Value
var once_Control_Applicative_pure__154576880 sync.Once
func Get_Control_Applicative_pure__154576880() gopurs_runtime.Value {
	once_Control_Applicative_pure__154576880.Do(func() {
		cache_Control_Applicative_pure__154576880 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__154576880(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__154576880
}

var cache_Control_Applicative_pure__3820067664 gopurs_runtime.Value
var once_Control_Applicative_pure__3820067664 sync.Once
func Get_Control_Applicative_pure__3820067664() gopurs_runtime.Value {
	once_Control_Applicative_pure__3820067664.Do(func() {
		cache_Control_Applicative_pure__3820067664 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3820067664(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__3820067664
}

var cache_Control_Applicative_pure__1953455120 gopurs_runtime.Value
var once_Control_Applicative_pure__1953455120 sync.Once
func Get_Control_Applicative_pure__1953455120() gopurs_runtime.Value {
	once_Control_Applicative_pure__1953455120.Do(func() {
		cache_Control_Applicative_pure__1953455120 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1953455120(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__1953455120
}

var cache_Control_Applicative_pure__778206864 gopurs_runtime.Value
var once_Control_Applicative_pure__778206864 sync.Once
func Get_Control_Applicative_pure__778206864() gopurs_runtime.Value {
	once_Control_Applicative_pure__778206864.Do(func() {
		cache_Control_Applicative_pure__778206864 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__778206864(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__778206864
}

var cache_Control_Applicative_pure__983529968 gopurs_runtime.Value
var once_Control_Applicative_pure__983529968 sync.Once
func Get_Control_Applicative_pure__983529968() gopurs_runtime.Value {
	once_Control_Applicative_pure__983529968.Do(func() {
		cache_Control_Applicative_pure__983529968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__983529968(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__983529968
}

var cache_Control_Applicative_pure__2812254544 gopurs_runtime.Value
var once_Control_Applicative_pure__2812254544 sync.Once
func Get_Control_Applicative_pure__2812254544() gopurs_runtime.Value {
	once_Control_Applicative_pure__2812254544.Do(func() {
		cache_Control_Applicative_pure__2812254544 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2812254544(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__2812254544
}

var cache_Control_Applicative_pure__1304937360 gopurs_runtime.Value
var once_Control_Applicative_pure__1304937360 sync.Once
func Get_Control_Applicative_pure__1304937360() gopurs_runtime.Value {
	once_Control_Applicative_pure__1304937360.Do(func() {
		cache_Control_Applicative_pure__1304937360 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1304937360(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__1304937360
}

var cache_Control_Applicative_pure__1670386480 gopurs_runtime.Value
var once_Control_Applicative_pure__1670386480 sync.Once
func Get_Control_Applicative_pure__1670386480() gopurs_runtime.Value {
	once_Control_Applicative_pure__1670386480.Do(func() {
		cache_Control_Applicative_pure__1670386480 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1670386480(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__1670386480
}

var cache_Control_Applicative_pure__1449138640 gopurs_runtime.Value
var once_Control_Applicative_pure__1449138640 sync.Once
func Get_Control_Applicative_pure__1449138640() gopurs_runtime.Value {
	once_Control_Applicative_pure__1449138640.Do(func() {
		cache_Control_Applicative_pure__1449138640 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1449138640(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__1449138640
}

var cache_Control_Applicative_pure__3079134646 gopurs_runtime.Value
var once_Control_Applicative_pure__3079134646 sync.Once
func Get_Control_Applicative_pure__3079134646() gopurs_runtime.Value {
	once_Control_Applicative_pure__3079134646.Do(func() {
		cache_Control_Applicative_pure__3079134646 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3079134646(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__3079134646
}

var cache_Control_Applicative_pure__2516715542 gopurs_runtime.Value
var once_Control_Applicative_pure__2516715542 sync.Once
func Get_Control_Applicative_pure__2516715542() gopurs_runtime.Value {
	once_Control_Applicative_pure__2516715542.Do(func() {
		cache_Control_Applicative_pure__2516715542 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2516715542(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__2516715542
}

var cache_Control_Applicative_pure__3145599862 gopurs_runtime.Value
var once_Control_Applicative_pure__3145599862 sync.Once
func Get_Control_Applicative_pure__3145599862() gopurs_runtime.Value {
	once_Control_Applicative_pure__3145599862.Do(func() {
		cache_Control_Applicative_pure__3145599862 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3145599862(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__3145599862
}

var cache_Control_Applicative_pure__1715998582 gopurs_runtime.Value
var once_Control_Applicative_pure__1715998582 sync.Once
func Get_Control_Applicative_pure__1715998582() gopurs_runtime.Value {
	once_Control_Applicative_pure__1715998582.Do(func() {
		cache_Control_Applicative_pure__1715998582 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1715998582(a_0_box)
})
	})
	return cache_Control_Applicative_pure__1715998582
}

var cache_Control_Applicative_pure__2154981942 gopurs_runtime.Value
var once_Control_Applicative_pure__2154981942 sync.Once
func Get_Control_Applicative_pure__2154981942() gopurs_runtime.Value {
	once_Control_Applicative_pure__2154981942.Do(func() {
		cache_Control_Applicative_pure__2154981942 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Control_Applicative_pure__2154981942(a_0_box))}
})
	})
	return cache_Control_Applicative_pure__2154981942
}

var cache_Control_Applicative_pure__577281046 gopurs_runtime.Value
var once_Control_Applicative_pure__577281046 sync.Once
func Get_Control_Applicative_pure__577281046() gopurs_runtime.Value {
	once_Control_Applicative_pure__577281046.Do(func() {
		cache_Control_Applicative_pure__577281046 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__577281046(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__577281046
}

var cache_Control_Applicative_pure__1646981750 gopurs_runtime.Value
var once_Control_Applicative_pure__1646981750 sync.Once
func Get_Control_Applicative_pure__1646981750() gopurs_runtime.Value {
	once_Control_Applicative_pure__1646981750.Do(func() {
		cache_Control_Applicative_pure__1646981750 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1646981750(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__1646981750
}

var cache_Control_Applicative_pure__3181299446 gopurs_runtime.Value
var once_Control_Applicative_pure__3181299446 sync.Once
func Get_Control_Applicative_pure__3181299446() gopurs_runtime.Value {
	once_Control_Applicative_pure__3181299446.Do(func() {
		cache_Control_Applicative_pure__3181299446 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3181299446(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__3181299446
}

var cache_Control_Applicative_pure__771998518 gopurs_runtime.Value
var once_Control_Applicative_pure__771998518 sync.Once
func Get_Control_Applicative_pure__771998518() gopurs_runtime.Value {
	once_Control_Applicative_pure__771998518.Do(func() {
		cache_Control_Applicative_pure__771998518 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__771998518(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__771998518
}

var cache_Control_Applicative_pure__575667894 gopurs_runtime.Value
var once_Control_Applicative_pure__575667894 sync.Once
func Get_Control_Applicative_pure__575667894() gopurs_runtime.Value {
	once_Control_Applicative_pure__575667894.Do(func() {
		cache_Control_Applicative_pure__575667894 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__575667894(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__575667894
}

var cache_Control_Applicative_pure__1895379222 gopurs_runtime.Value
var once_Control_Applicative_pure__1895379222 sync.Once
func Get_Control_Applicative_pure__1895379222() gopurs_runtime.Value {
	once_Control_Applicative_pure__1895379222.Do(func() {
		cache_Control_Applicative_pure__1895379222 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1895379222(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__1895379222
}

var cache_Control_Applicative_pure__3514127574 gopurs_runtime.Value
var once_Control_Applicative_pure__3514127574 sync.Once
func Get_Control_Applicative_pure__3514127574() gopurs_runtime.Value {
	once_Control_Applicative_pure__3514127574.Do(func() {
		cache_Control_Applicative_pure__3514127574 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3514127574(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__3514127574
}

var cache_Control_Applicative_pure__2195681590 gopurs_runtime.Value
var once_Control_Applicative_pure__2195681590 sync.Once
func Get_Control_Applicative_pure__2195681590() gopurs_runtime.Value {
	once_Control_Applicative_pure__2195681590.Do(func() {
		cache_Control_Applicative_pure__2195681590 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2195681590(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__2195681590
}

var cache_Control_Applicative_pure__3229300374 gopurs_runtime.Value
var once_Control_Applicative_pure__3229300374 sync.Once
func Get_Control_Applicative_pure__3229300374() gopurs_runtime.Value {
	once_Control_Applicative_pure__3229300374.Do(func() {
		cache_Control_Applicative_pure__3229300374 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3229300374(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__3229300374
}

var cache_Control_Applicative_pure__3527452822 gopurs_runtime.Value
var once_Control_Applicative_pure__3527452822 sync.Once
func Get_Control_Applicative_pure__3527452822() gopurs_runtime.Value {
	once_Control_Applicative_pure__3527452822.Do(func() {
		cache_Control_Applicative_pure__3527452822 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3527452822(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__3527452822
}

var cache_Control_Applicative_pure__629383158 gopurs_runtime.Value
var once_Control_Applicative_pure__629383158 sync.Once
func Get_Control_Applicative_pure__629383158() gopurs_runtime.Value {
	once_Control_Applicative_pure__629383158.Do(func() {
		cache_Control_Applicative_pure__629383158 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__629383158(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__629383158
}

var cache_Control_Applicative_pure__4209427318 gopurs_runtime.Value
var once_Control_Applicative_pure__4209427318 sync.Once
func Get_Control_Applicative_pure__4209427318() gopurs_runtime.Value {
	once_Control_Applicative_pure__4209427318.Do(func() {
		cache_Control_Applicative_pure__4209427318 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__4209427318(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__4209427318
}

var cache_Control_Applicative_pure__3540891798 gopurs_runtime.Value
var once_Control_Applicative_pure__3540891798 sync.Once
func Get_Control_Applicative_pure__3540891798() gopurs_runtime.Value {
	once_Control_Applicative_pure__3540891798.Do(func() {
		cache_Control_Applicative_pure__3540891798 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3540891798(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__3540891798
}

var cache_Control_Applicative_pure__1641029622 gopurs_runtime.Value
var once_Control_Applicative_pure__1641029622 sync.Once
func Get_Control_Applicative_pure__1641029622() gopurs_runtime.Value {
	once_Control_Applicative_pure__1641029622.Do(func() {
		cache_Control_Applicative_pure__1641029622 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1641029622(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__1641029622
}

var cache_Control_Applicative_pure__2106705590 gopurs_runtime.Value
var once_Control_Applicative_pure__2106705590 sync.Once
func Get_Control_Applicative_pure__2106705590() gopurs_runtime.Value {
	once_Control_Applicative_pure__2106705590.Do(func() {
		cache_Control_Applicative_pure__2106705590 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2106705590(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__2106705590
}

var cache_Control_Applicative_pure__2644984438 gopurs_runtime.Value
var once_Control_Applicative_pure__2644984438 sync.Once
func Get_Control_Applicative_pure__2644984438() gopurs_runtime.Value {
	once_Control_Applicative_pure__2644984438.Do(func() {
		cache_Control_Applicative_pure__2644984438 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2644984438(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__2644984438
}

var cache_Control_Applicative_pure__3453203222 gopurs_runtime.Value
var once_Control_Applicative_pure__3453203222 sync.Once
func Get_Control_Applicative_pure__3453203222() gopurs_runtime.Value {
	once_Control_Applicative_pure__3453203222.Do(func() {
		cache_Control_Applicative_pure__3453203222 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3453203222(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__3453203222
}

var cache_Control_Applicative_pure__4169392240 gopurs_runtime.Value
var once_Control_Applicative_pure__4169392240 sync.Once
func Get_Control_Applicative_pure__4169392240() gopurs_runtime.Value {
	once_Control_Applicative_pure__4169392240.Do(func() {
		cache_Control_Applicative_pure__4169392240 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__4169392240(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__4169392240
}

var cache_Control_Applicative_pure__1476406672 gopurs_runtime.Value
var once_Control_Applicative_pure__1476406672 sync.Once
func Get_Control_Applicative_pure__1476406672() gopurs_runtime.Value {
	once_Control_Applicative_pure__1476406672.Do(func() {
		cache_Control_Applicative_pure__1476406672 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1476406672(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dict_0_box))
})
	})
	return cache_Control_Applicative_pure__1476406672
}

var cache_Control_Applicative_unless__1954875249 gopurs_runtime.Value
var once_Control_Applicative_unless__1954875249 sync.Once
func Get_Control_Applicative_unless__1954875249() gopurs_runtime.Value {
	once_Control_Applicative_unless__1954875249.Do(func() {
		cache_Control_Applicative_unless__1954875249 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_unless__1954875249(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_Control_Applicative_unless__1954875249
}

var cache_Control_Applicative_when__1954875249 gopurs_runtime.Value
var once_Control_Applicative_when__1954875249 sync.Once
func Get_Control_Applicative_when__1954875249() gopurs_runtime.Value {
	once_Control_Applicative_when__1954875249.Do(func() {
		cache_Control_Applicative_when__1954875249 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_when__1954875249(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_Control_Applicative_when__1954875249
}

var cache_Control_Applicative_when__4245282944 gopurs_runtime.Value
var once_Control_Applicative_when__4245282944 sync.Once
func Get_Control_Applicative_when__4245282944() gopurs_runtime.Value {
	once_Control_Applicative_when__4245282944.Do(func() {
		cache_Control_Applicative_when__4245282944 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_when__4245282944((v_0_box.IntVal) != (0), v1_1_box)
})
	})
	return cache_Control_Applicative_when__4245282944
}

type Constructor_Control_Applicative_Applicative struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1459134221] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Applicative_Applicative)(ptr)
		_ = c
		switch key {
		case "Apply0": return gopurs_runtime.Box(c.V0)
		case "pure": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Applicative_Applicative: " + key)
		}
	}
}


func Call_Control_Applicative_Applicative_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Applicative_pure(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_unless(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 bool = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t0 gopurs_runtime.Value
{
if (v_1) != (true) {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
if v_1 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit())
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_Control_Applicative_when(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 bool = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t0 gopurs_runtime.Value
{
if v_1 {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit())
}
end_branch_0:
return __t0
}

func Call_Control_Applicative_liftA1(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}))
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), f_2), a_3)
})
})
}

func Call_Control_Applicative_pure__2382779990(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__1812053569(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__2491902360(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__149572504(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__189931222(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__779048344(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__4037597590(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__3432560598(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__3236307030(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__2331489366(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__2302010678(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__993904534(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__2935994064(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__3215807376(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__2960485136(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__160425008(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__3286817552(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__763072784(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__871290128(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__1748760400(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__3012389648(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__566620048(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__3197665392(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__1475749520(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__1253336208(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__355615152(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__1741581456(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__243192752(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__4233214992(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__154576880(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__3820067664(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__1953455120(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__778206864(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__983529968(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__2812254544(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__1304937360(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__1670386480(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__1449138640(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__3079134646(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __eta0_0
})
}

func Call_Control_Applicative_pure__2516715542(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __eta0_0
})
}

func Call_Control_Applicative_pure__3145599862(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Data_Either_Right(), __eta0_0)
}

func Call_Control_Applicative_pure__1715998582(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
}

func Call_Control_Applicative_pure__2154981942(a_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), a_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](Get_Data_List_Types_Nil()))}))
}

func Call_Control_Applicative_pure__577281046(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), __eta0_0)))}
}

func Call_Control_Applicative_pure__1646981750(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), __eta0_0)))}
}

func Call_Control_Applicative_pure__3181299446(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Maybe_Just(), __eta0_0)))}
}

func Call_Control_Applicative_pure__771998518(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_List_Types_NonEmptyList(), gopurs_runtime.Apply2(Get_Data_NonEmpty_singleton(), gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](Get_Data_List_Types_plusList()))}, __eta0_0))))}
}

func Call_Control_Applicative_pure__575667894(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Data_List_Types_NonEmptyList(), gopurs_runtime.Apply2(Get_Data_NonEmpty_singleton(), gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](Get_Data_List_Types_plusList()))}, __eta0_0))))}
}

func Call_Control_Applicative_pure__1895379222(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Effect_Aff__pure(), __eta0_0)
}

func Call_Control_Applicative_pure__3514127574(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Effect_Aff__pure(), __eta0_0)
}

func Call_Control_Applicative_pure__2195681590(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Effect_Aff__pure(), __eta0_0)
}

func Call_Control_Applicative_pure__3229300374(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Effect_Aff__pure(), __eta0_0)
}

func Call_Control_Applicative_pure__3527452822(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply2(Get_Control_Parallel_Class_parallel(), gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel](Get_Effect_Aff_parallelAff()))}, gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Effect_Aff_applicativeAff()).V1), __eta0_0))
}

func Call_Control_Applicative_pure__629383158(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __eta0_0
})
}

func Call_Control_Applicative_pure__4209427318(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __eta0_0
})
}

func Call_Control_Applicative_pure__3540891798(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __eta0_0
})
}

func Call_Control_Applicative_pure__1641029622(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __eta0_0
})
}

func Call_Control_Applicative_pure__2106705590(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __eta0_0
})
}

func Call_Control_Applicative_pure__2644984438(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __eta0_0
})
}

func Call_Control_Applicative_pure__3453203222(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __eta0_0
})
}

func Call_Control_Applicative_pure__4169392240(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_pure__1476406672(dict_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Applicative_unless__1954875249(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 bool = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t0 gopurs_runtime.Value
{
if (v_1) != (true) {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
if v_1 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit())
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_Control_Applicative_when__1954875249(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 bool = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t0 gopurs_runtime.Value
{
if v_1 {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit())
}
end_branch_0:
return __t0
}

func Call_Control_Applicative_when__4245282944(v_0_loop bool, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if v_0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Monad_ST_Internal_applicativeST()).V1), Get_Data_Unit_unit())
}
end_branch_0:
return __t0
}


