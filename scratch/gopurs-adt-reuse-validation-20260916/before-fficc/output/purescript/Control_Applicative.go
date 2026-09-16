package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Applicative_Applicative_dollar_Dict gopurs_runtime.Value
var once_Control_Applicative_Applicative_dollar_Dict sync.Once
func Get_Control_Applicative_Applicative_dollar_Dict() gopurs_runtime.Value {
	once_Control_Applicative_Applicative_dollar_Dict.Do(func() {
		cache_Control_Applicative_Applicative_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Call_Control_Applicative_Applicative_dollar_Dict(func() struct{
	Apply0 gopurs_runtime.Value
	pure gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Apply0 gopurs_runtime.Value
	pure gopurs_runtime.Value
}{}
					clone.Apply0 = gopurs_runtime.RecordGet(orig, "Apply0")
					clone.pure = gopurs_runtime.RecordGet(orig, "pure")
					return clone
				}()))}
})
	})
	return cache_Control_Applicative_Applicative_dollar_Dict
}

var cache_Control_Applicative_pure gopurs_runtime.Value
var once_Control_Applicative_pure sync.Once
func Get_Control_Applicative_pure() gopurs_runtime.Value {
	once_Control_Applicative_pure.Do(func() {
		cache_Control_Applicative_pure = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Applicative_pure
}

var cache_Control_Applicative_pure__2646035056 gopurs_runtime.Value
var once_Control_Applicative_pure__2646035056 sync.Once
func Get_Control_Applicative_pure__2646035056() gopurs_runtime.Value {
	once_Control_Applicative_pure__2646035056.Do(func() {
		cache_Control_Applicative_pure__2646035056 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Applicative_pure__2646035056(uint32(__eta_norm_0_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Applicative_pure__2646035056
}

var cache_Control_Applicative_pure__3381183341 gopurs_runtime.Value
var once_Control_Applicative_pure__3381183341 sync.Once
func Get_Control_Applicative_pure__3381183341() gopurs_runtime.Value {
	once_Control_Applicative_pure__3381183341.Do(func() {
		cache_Control_Applicative_pure__3381183341 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Control_Applicative_pure__3381183341(__eta_norm_0_unused_0_box))}
})
	})
	return cache_Control_Applicative_pure__3381183341
}

var cache_Control_Applicative_pure__987087589 gopurs_runtime.Value
var once_Control_Applicative_pure__987087589 sync.Once
func Get_Control_Applicative_pure__987087589() gopurs_runtime.Value {
	once_Control_Applicative_pure__987087589.Do(func() {
		cache_Control_Applicative_pure__987087589 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Control_Applicative_pure__987087589(__eta_norm_0_unused_0_box))}
})
	})
	return cache_Control_Applicative_pure__987087589
}

var cache_Control_Applicative_pure__3634220417 gopurs_runtime.Value
var once_Control_Applicative_pure__3634220417 sync.Once
func Get_Control_Applicative_pure__3634220417() gopurs_runtime.Value {
	once_Control_Applicative_pure__3634220417.Do(func() {
		cache_Control_Applicative_pure__3634220417 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Control_Applicative_pure__3634220417(__eta_norm_0_0_box))}
})
	})
	return cache_Control_Applicative_pure__3634220417
}

var cache_Control_Applicative_pure__107423425 gopurs_runtime.Value
var once_Control_Applicative_pure__107423425 sync.Once
func Get_Control_Applicative_pure__107423425() gopurs_runtime.Value {
	once_Control_Applicative_pure__107423425.Do(func() {
		cache_Control_Applicative_pure__107423425 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__107423425((__eta_norm_0_0_box.IntVal) != (0))
})
	})
	return cache_Control_Applicative_pure__107423425
}

var cache_Control_Applicative_pure__2758158054 gopurs_runtime.Value
var once_Control_Applicative_pure__2758158054 sync.Once
func Get_Control_Applicative_pure__2758158054() gopurs_runtime.Value {
	once_Control_Applicative_pure__2758158054.Do(func() {
		cache_Control_Applicative_pure__2758158054 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2758158054(__eta_norm_0_unused_0_box)
})
	})
	return cache_Control_Applicative_pure__2758158054
}

var cache_Control_Applicative_pure__3637276289 gopurs_runtime.Value
var once_Control_Applicative_pure__3637276289 sync.Once
func Get_Control_Applicative_pure__3637276289() gopurs_runtime.Value {
	once_Control_Applicative_pure__3637276289.Do(func() {
		cache_Control_Applicative_pure__3637276289 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3637276289(__eta_norm_0_0_box)
})
	})
	return cache_Control_Applicative_pure__3637276289
}

var cache_Control_Applicative_pure__1953105552 gopurs_runtime.Value
var once_Control_Applicative_pure__1953105552 sync.Once
func Get_Control_Applicative_pure__1953105552() gopurs_runtime.Value {
	once_Control_Applicative_pure__1953105552.Do(func() {
		cache_Control_Applicative_pure__1953105552 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Applicative_pure__1953105552(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Applicative_pure__1953105552
}

var cache_Control_Applicative_pure__2352921904 gopurs_runtime.Value
var once_Control_Applicative_pure__2352921904 sync.Once
func Get_Control_Applicative_pure__2352921904() gopurs_runtime.Value {
	once_Control_Applicative_pure__2352921904.Do(func() {
		cache_Control_Applicative_pure__2352921904 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2352921904(__eta_norm_0_0_box.FloatVal())
})
	})
	return cache_Control_Applicative_pure__2352921904
}

var cache_Control_Applicative_pure__3111615553 gopurs_runtime.Value
var once_Control_Applicative_pure__3111615553 sync.Once
func Get_Control_Applicative_pure__3111615553() gopurs_runtime.Value {
	once_Control_Applicative_pure__3111615553.Do(func() {
		cache_Control_Applicative_pure__3111615553 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3111615553(__eta_norm_0_0_box.FloatVal())
})
	})
	return cache_Control_Applicative_pure__3111615553
}

var cache_Control_Applicative_pure__723319553 gopurs_runtime.Value
var once_Control_Applicative_pure__723319553 sync.Once
func Get_Control_Applicative_pure__723319553() gopurs_runtime.Value {
	once_Control_Applicative_pure__723319553.Do(func() {
		cache_Control_Applicative_pure__723319553 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__723319553(__eta_norm_0_0_box.StrVal())
})
	})
	return cache_Control_Applicative_pure__723319553
}

var cache_Control_Applicative_pure__3482448870 gopurs_runtime.Value
var once_Control_Applicative_pure__3482448870 sync.Once
func Get_Control_Applicative_pure__3482448870() gopurs_runtime.Value {
	once_Control_Applicative_pure__3482448870.Do(func() {
		cache_Control_Applicative_pure__3482448870 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Applicative_pure__3482448870(__eta_norm_0_unused_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
})
	})
	return cache_Control_Applicative_pure__3482448870
}

var cache_Control_Applicative_pure__2745689897 gopurs_runtime.Value
var once_Control_Applicative_pure__2745689897 sync.Once
func Get_Control_Applicative_pure__2745689897() gopurs_runtime.Value {
	once_Control_Applicative_pure__2745689897.Do(func() {
		cache_Control_Applicative_pure__2745689897 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Applicative_pure__2745689897(__eta_norm_0_unused_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Applicative_pure__2745689897
}

var cache_Control_Applicative_pure__2353832556 gopurs_runtime.Value
var once_Control_Applicative_pure__2353832556 sync.Once
func Get_Control_Applicative_pure__2353832556() gopurs_runtime.Value {
	once_Control_Applicative_pure__2353832556.Do(func() {
		cache_Control_Applicative_pure__2353832556 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2353832556(__eta0_0_box)
})
	})
	return cache_Control_Applicative_pure__2353832556
}

var cache_Control_Applicative_pure__3122131407 gopurs_runtime.Value
var once_Control_Applicative_pure__3122131407 sync.Once
func Get_Control_Applicative_pure__3122131407() gopurs_runtime.Value {
	once_Control_Applicative_pure__3122131407.Do(func() {
		cache_Control_Applicative_pure__3122131407 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__3122131407(__eta_norm_0_unused_0_box)
})
	})
	return cache_Control_Applicative_pure__3122131407
}

var cache_Control_Applicative_pure__1085681566 gopurs_runtime.Value
var once_Control_Applicative_pure__1085681566 sync.Once
func Get_Control_Applicative_pure__1085681566() gopurs_runtime.Value {
	once_Control_Applicative_pure__1085681566.Do(func() {
		cache_Control_Applicative_pure__1085681566 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__1085681566(__eta_norm_0_unused_0_box)
})
	})
	return cache_Control_Applicative_pure__1085681566
}

var cache_Control_Applicative_pure__2097684664 gopurs_runtime.Value
var once_Control_Applicative_pure__2097684664 sync.Once
func Get_Control_Applicative_pure__2097684664() gopurs_runtime.Value {
	once_Control_Applicative_pure__2097684664.Do(func() {
		cache_Control_Applicative_pure__2097684664 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_pure__2097684664(__eta_norm_0_unused_0_box)
})
	})
	return cache_Control_Applicative_pure__2097684664
}

var cache_Control_Applicative_unless gopurs_runtime.Value
var once_Control_Applicative_unless sync.Once
func Get_Control_Applicative_unless() gopurs_runtime.Value {
	once_Control_Applicative_unless.Do(func() {
		cache_Control_Applicative_unless = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_unless(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_Control_Applicative_unless
}

var cache_Control_Applicative_when gopurs_runtime.Value
var once_Control_Applicative_when sync.Once
func Get_Control_Applicative_when() gopurs_runtime.Value {
	once_Control_Applicative_when.Do(func() {
		cache_Control_Applicative_when = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_when(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_Control_Applicative_when
}

var cache_Control_Applicative_liftA1 gopurs_runtime.Value
var once_Control_Applicative_liftA1 sync.Once
func Get_Control_Applicative_liftA1() gopurs_runtime.Value {
	once_Control_Applicative_liftA1.Do(func() {
		cache_Control_Applicative_liftA1 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Control_Applicative_liftA1
}

var cache_Control_Applicative_applicativeProxy gopurs_runtime.Value
var once_Control_Applicative_applicativeProxy sync.Once
func Get_Control_Applicative_applicativeProxy() gopurs_runtime.Value {
	once_Control_Applicative_applicativeProxy.Do(func() {
		cache_Control_Applicative_applicativeProxy = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Control_Applicative_3924528404_1439734649((&Constructor_Control_Applicative_Applicative[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Control_Applicative_1303115796_3741347833(Rebox_Control_Applicative_3741347833_1303115796(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyProxy()))))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})})))}
	})
	return cache_Control_Applicative_applicativeProxy
}

var cache_Control_Applicative_applicativeFn gopurs_runtime.Value
var once_Control_Applicative_applicativeFn sync.Once
func Get_Control_Applicative_applicativeFn() gopurs_runtime.Value {
	once_Control_Applicative_applicativeFn.Do(func() {
		cache_Control_Applicative_applicativeFn = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyFn()))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})}))}
	})
	return cache_Control_Applicative_applicativeFn
}

var cache_Control_Applicative_applicativeArray gopurs_runtime.Value
var once_Control_Applicative_applicativeArray sync.Once
func Get_Control_Applicative_applicativeArray() gopurs_runtime.Value {
	once_Control_Applicative_applicativeArray.Do(func() {
		cache_Control_Applicative_applicativeArray = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyArray()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})}))}
	})
	return cache_Control_Applicative_applicativeArray
}

type Constructor_Control_Applicative_Applicative[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1459134221] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Applicative_Applicative[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Apply0": return gopurs_runtime.Box(c.V0)
		case "pure": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Applicative_Applicative: " + key)
		}
	}
}


func Call_Control_Applicative_Applicative_dollar_Dict(x_0_loop struct{
	Apply0 gopurs_runtime.Value
	pure gopurs_runtime.Value
}) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
var x_0 struct{
	Apply0 gopurs_runtime.Value
	pure gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Apply0", "pure", orig.Apply0, orig.pure)
				}())
}

func Call_Control_Applicative_pure(dict_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_Control_Applicative_pure__2646035056(__eta_norm_0_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
pure__2646035056:
for {
if false { continue pure__2646035056 }
var __eta_norm_0_0 uint32 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(Rebox_Control_Applicative_1439734649_649684152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Maybe_applicativeMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta_norm_0_0), UnsafePtr: nil})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Applicative_pure__3381183341(__eta_norm_0_unused_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
pure__3381183341:
for {
if false { continue pure__3381183341 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Control_Applicative_1439734649_2870828117(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Types_applicativeList())).V1, Get_Data_Interval_Duration_Iso_InvalidWeekComponentUsage()))
}
}

func Call_Control_Applicative_pure__987087589(__eta_norm_0_unused_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
pure__987087589:
for {
if false { continue pure__987087589 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Control_Applicative_1439734649_2870828117(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Types_applicativeList())).V1, Get_Data_Interval_Duration_Iso_IsEmpty()))
}
}

func Call_Control_Applicative_pure__3634220417(__eta_norm_0_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
pure__3634220417:
for {
if false { continue pure__3634220417 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Control_Applicative_1439734649_2870828117(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Types_applicativeList())).V1, __eta_norm_0_0))
}
}

func Call_Control_Applicative_pure__107423425(__eta_norm_0_0_loop bool) gopurs_runtime.Value {
pure__107423425:
for {
if false { continue pure__107423425 }
var __eta_norm_0_0 bool = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, gopurs_runtime.Bool(__eta_norm_0_0))
}
}

func Call_Control_Applicative_pure__2758158054(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pure__2758158054:
for {
if false { continue pure__2758158054 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, Get_Effect_Aff_nonCanceler())
}
}

func Call_Control_Applicative_pure__3637276289(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pure__3637276289:
for {
if false { continue pure__3637276289 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, __eta_norm_0_0)
}
}

func Call_Control_Applicative_pure__1953105552(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
pure__1953105552:
for {
if false { continue pure__1953105552 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(Rebox_Control_Applicative_1439734649_649684152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Maybe_applicativeMaybe())).V1, gopurs_runtime.Int(__eta_norm_0_0))
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Applicative_pure__2352921904(__eta_norm_0_0_loop float64) gopurs_runtime.Value {
pure__2352921904:
for {
if false { continue pure__2352921904 }
var __eta_norm_0_0 float64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeAff()).V1, gopurs_runtime.Float(__eta_norm_0_0))
}
}

func Call_Control_Applicative_pure__3111615553(__eta_norm_0_0_loop float64) gopurs_runtime.Value {
pure__3111615553:
for {
if false { continue pure__3111615553 }
var __eta_norm_0_0 float64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, gopurs_runtime.Float(__eta_norm_0_0))
}
}

func Call_Control_Applicative_pure__723319553(__eta_norm_0_0_loop string) gopurs_runtime.Value {
pure__723319553:
for {
if false { continue pure__723319553 }
var __eta_norm_0_0 string = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, gopurs_runtime.Str(__eta_norm_0_0))
}
}

func Call_Control_Applicative_pure__3482448870(__eta_norm_0_unused_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
pure__3482448870:
for {
if false { continue pure__3482448870 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
				_v := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_applicativeEither(), "pure"), Get_Data_Unit_unit())
				if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: gopurs_runtime.Value{}, V1: (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V2: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: gopurs_runtime.Value{}, V2: false}
			}()
}
}

func Call_Control_Applicative_pure__2745689897(__eta_norm_0_unused_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
pure__2745689897:
for {
if false { continue pure__2745689897 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Rebox_Control_Applicative_2307501113_673668088(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Data_Maybe_alternativeMaybe())).V0, gopurs_runtime.Value{}), "pure"), Get_Data_Unit_unit())
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Applicative_pure__2353832556(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pure__2353832556:
for {
if false { continue pure__2353832556 }
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeAff()).V1, __eta0_0)
}
}

func Call_Control_Applicative_pure__3122131407(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pure__3122131407:
for {
if false { continue pure__3122131407 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeAff()).V1, Get_Data_Unit_unit())
}
}

func Call_Control_Applicative_pure__1085681566(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pure__1085681566:
for {
if false { continue pure__1085681566 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, Get_Data_Unit_unit())
}
}

func Call_Control_Applicative_pure__2097684664(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pure__2097684664:
for {
if false { continue pure__2097684664 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Control_Alternative_alternativeArray()).V0, gopurs_runtime.Value{}), "pure"), Get_Data_Unit_unit())
}
}

func Call_Control_Applicative_unless(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
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
__t0 = gopurs_runtime.Apply(dictApplicative_0.V1, Get_Data_Unit_unit())
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

func Call_Control_Applicative_when(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
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
__t0 = gopurs_runtime.Apply(dictApplicative_0.V1, Get_Data_Unit_unit())
}
end_branch_0:
return __t0
}

func Call_Control_Applicative_liftA1(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f$scope9)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}))
_ = Apply0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply(dictApplicative_0.V1, f_2), a_3)
})
}

func Rebox_Control_Applicative_1303115796_3741347833(in *Constructor_Control_Apply_Apply[uint32]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Applicative_1439734649_2870828117(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Applicative_1439734649_649684152(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Applicative_2307501113_673668088(in *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) *Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Applicative_3741347833_1303115796(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[uint32] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[uint32]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Applicative_3924528404_1439734649(in *Constructor_Control_Applicative_Applicative[uint32]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}


