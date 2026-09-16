package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Apply_Apply_dollar_Dict gopurs_runtime.Value
var once_Control_Apply_Apply_dollar_Dict sync.Once
func Get_Control_Apply_Apply_dollar_Dict() gopurs_runtime.Value {
	once_Control_Apply_Apply_dollar_Dict.Do(func() {
		cache_Control_Apply_Apply_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Call_Control_Apply_Apply_dollar_Dict(func() struct{
	Functor0 gopurs_runtime.Value
	apply gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Functor0 gopurs_runtime.Value
	apply gopurs_runtime.Value
}{}
					clone.Functor0 = gopurs_runtime.RecordGet(orig, "Functor0")
					clone.apply = gopurs_runtime.RecordGet(orig, "apply")
					return clone
				}()))}
})
	})
	return cache_Control_Apply_Apply_dollar_Dict
}

var cache_Control_Apply_applyProxy gopurs_runtime.Value
var once_Control_Apply_applyProxy sync.Once
func Get_Control_Apply_applyProxy() gopurs_runtime.Value {
	once_Control_Apply_applyProxy.Do(func() {
		cache_Control_Apply_applyProxy = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_1303115796_3741347833((&Constructor_Control_Apply_Apply[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_562069347_2812149806(Rebox_Control_Apply_2812149806_562069347(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorProxy()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})})))}
	})
	return cache_Control_Apply_applyProxy
}

var cache_Control_Apply_applyFn gopurs_runtime.Value
var once_Control_Apply_applyFn sync.Once
func Get_Control_Apply_applyFn() gopurs_runtime.Value {
	once_Control_Apply_applyFn.Do(func() {
		cache_Control_Apply_applyFn = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorFn()))}
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
})}))}
	})
	return cache_Control_Apply_applyFn
}

var cache_Control_Apply_applyArray gopurs_runtime.Value
var once_Control_Apply_applyArray sync.Once
func Get_Control_Apply_applyArray() gopurs_runtime.Value {
	once_Control_Apply_applyArray.Do(func() {
		cache_Control_Apply_applyArray = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
}), Get_Control_Apply_arrayApply()}))}
	})
	return cache_Control_Apply_applyArray
}

var cache_Control_Apply_apply gopurs_runtime.Value
var once_Control_Apply_apply sync.Once
func Get_Control_Apply_apply() gopurs_runtime.Value {
	once_Control_Apply_apply.Do(func() {
		cache_Control_Apply_apply = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Apply_apply
}

var cache_Control_Apply_apply__957400575 gopurs_runtime.Value
var once_Control_Apply_apply__957400575 sync.Once
func Get_Control_Apply_apply__957400575() gopurs_runtime.Value {
	once_Control_Apply_apply__957400575.Do(func() {
		cache_Control_Apply_apply__957400575 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__957400575(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0_box), Rebox_Control_Apply_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Apply_apply__957400575
}

var cache_Control_Apply_apply__3315610751 gopurs_runtime.Value
var once_Control_Apply_apply__3315610751 sync.Once
func Get_Control_Apply_apply__3315610751() gopurs_runtime.Value {
	once_Control_Apply_apply__3315610751.Do(func() {
		cache_Control_Apply_apply__3315610751 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__3315610751(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0_box), Rebox_Control_Apply_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Apply_apply__3315610751
}

var cache_Control_Apply_apply__1880633631 gopurs_runtime.Value
var once_Control_Apply_apply__1880633631 sync.Once
func Get_Control_Apply_apply__1880633631() gopurs_runtime.Value {
	once_Control_Apply_apply__1880633631.Do(func() {
		cache_Control_Apply_apply__1880633631 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__1880633631(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0_box), Rebox_Control_Apply_3094389156_3839235747(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Apply_apply__1880633631
}

var cache_Control_Apply_apply__3188411071 gopurs_runtime.Value
var once_Control_Apply_apply__3188411071 sync.Once
func Get_Control_Apply_apply__3188411071() gopurs_runtime.Value {
	once_Control_Apply_apply__3188411071.Do(func() {
		cache_Control_Apply_apply__3188411071 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__3188411071(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0_box), Rebox_Control_Apply_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Apply_apply__3188411071
}

var cache_Control_Apply_apply__3103948255 gopurs_runtime.Value
var once_Control_Apply_apply__3103948255 sync.Once
func Get_Control_Apply_apply__3103948255() gopurs_runtime.Value {
	once_Control_Apply_apply__3103948255.Do(func() {
		cache_Control_Apply_apply__3103948255 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__3103948255(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0_box), Rebox_Control_Apply_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Apply_apply__3103948255
}

var cache_Control_Apply_apply__3730494975 gopurs_runtime.Value
var once_Control_Apply_apply__3730494975 sync.Once
func Get_Control_Apply_apply__3730494975() gopurs_runtime.Value {
	once_Control_Apply_apply__3730494975.Do(func() {
		cache_Control_Apply_apply__3730494975 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__3730494975(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0_box), Rebox_Control_Apply_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Control_Apply_apply__3730494975
}

var cache_Control_Apply_applyFirst gopurs_runtime.Value
var once_Control_Apply_applyFirst sync.Once
func Get_Control_Apply_applyFirst() gopurs_runtime.Value {
	once_Control_Apply_applyFirst.Do(func() {
		cache_Control_Apply_applyFirst = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_applyFirst(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_Control_Apply_applyFirst
}

var cache_Control_Apply_applySecond gopurs_runtime.Value
var once_Control_Apply_applySecond sync.Once
func Get_Control_Apply_applySecond() gopurs_runtime.Value {
	once_Control_Apply_applySecond.Do(func() {
		cache_Control_Apply_applySecond = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_applySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_Control_Apply_applySecond
}

var cache_Control_Apply_lift2 gopurs_runtime.Value
var once_Control_Apply_lift2 sync.Once
func Get_Control_Apply_lift2() gopurs_runtime.Value {
	once_Control_Apply_lift2.Do(func() {
		cache_Control_Apply_lift2 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift2
}

var cache_Control_Apply_lift3 gopurs_runtime.Value
var once_Control_Apply_lift3 sync.Once
func Get_Control_Apply_lift3() gopurs_runtime.Value {
	once_Control_Apply_lift3.Do(func() {
		cache_Control_Apply_lift3 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift3(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift3
}

var cache_Control_Apply_lift4 gopurs_runtime.Value
var once_Control_Apply_lift4 sync.Once
func Get_Control_Apply_lift4() gopurs_runtime.Value {
	once_Control_Apply_lift4.Do(func() {
		cache_Control_Apply_lift4 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift4(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift4
}

var cache_Control_Apply_lift5 gopurs_runtime.Value
var once_Control_Apply_lift5 sync.Once
func Get_Control_Apply_lift5() gopurs_runtime.Value {
	once_Control_Apply_lift5.Do(func() {
		cache_Control_Apply_lift5 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Apply_lift5(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_Control_Apply_lift5
}

type Constructor_Control_Apply_Apply[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3032403085] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Apply_Apply[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Functor0": return gopurs_runtime.Box(c.V0)
		case "apply": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Apply_Apply: " + key)
		}
	}
}


func Call_Control_Apply_Apply_dollar_Dict(x_0_loop struct{
	Functor0 gopurs_runtime.Value
	apply gopurs_runtime.Value
}) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
var x_0 struct{
	Functor0 gopurs_runtime.Value
	apply gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Functor0", "apply", orig.Functor0, orig.apply)
				}())
}

func Call_Control_Apply_apply(dict_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_Control_Apply_apply__957400575(__eta_norm_1_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[uint32]) struct{V0 gopurs_runtime.Value; V1 bool} {
apply__957400575:
for {
if false { continue apply__957400575 }
var __eta_norm_1_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[uint32] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta_norm_1_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_622082505_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Apply_apply__3315610751(__eta_norm_1_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[uint32]) struct{V0 gopurs_runtime.Value; V1 bool} {
apply__3315610751:
for {
if false { continue apply__3315610751 }
var __eta_norm_1_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[uint32] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta_norm_1_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_622082505_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Apply_apply__1880633631(__eta_norm_1_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]) struct{V0 gopurs_runtime.Value; V1 bool} {
apply__1880633631:
for {
if false { continue apply__1880633631 }
var __eta_norm_1_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta_norm_1_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_3839235747_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Apply_apply__3188411071(__eta_norm_1_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[int64]) struct{V0 gopurs_runtime.Value; V1 bool} {
apply__3188411071:
for {
if false { continue apply__3188411071 }
var __eta_norm_1_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta_norm_1_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_1170268447_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Apply_apply__3103948255(__eta_norm_1_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[int64]) struct{V0 gopurs_runtime.Value; V1 bool} {
apply__3103948255:
for {
if false { continue apply__3103948255 }
var __eta_norm_1_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta_norm_1_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_1170268447_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Apply_apply__3730494975(__eta_norm_1_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[int64]) struct{V0 gopurs_runtime.Value; V1 bool} {
apply__3730494975:
for {
if false { continue apply__3730494975 }
var __eta_norm_1_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta_norm_1_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_1170268447_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Control_Apply_applyFirst(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope29)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, Get_Data_Function_go__const(), a_2), b_3)
})
}

func Call_Control_Apply_applySecond(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope35)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
__local_var_4_1 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_4_1
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_1
}), a_2), b_3)
})
}

func Call_Control_Apply_lift2(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope44)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
}

func Call_Control_Apply_lift3(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope54)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func4(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4), c_5)
})
}

func Call_Control_Apply_lift4(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope66)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func5(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, c_5 gopurs_runtime.Value, d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4), c_5), d_6)
})
}

func Call_Control_Apply_lift5(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope79)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func5(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, c_5 gopurs_runtime.Value, d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4), c_5), d_6), e_7)
})
})
}

func Rebox_Control_Apply_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Control_Apply_1303115796_3741347833(in *Constructor_Control_Apply_Apply[uint32]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Apply_2812149806_562069347(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[uint32]{}
		out.V0 = in.V0
	return out
}

func Rebox_Control_Apply_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Control_Apply_3094389156_3839235747(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V0)
	return out
}

func Rebox_Control_Apply_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[uint32]{}
		out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Control_Apply_3741347833_3552963512(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Apply_3839235747_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Apply_562069347_2812149806(in *Constructor_Data_Functor_Functor[uint32]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Control_Apply_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Get_Control_Apply_arrayApply() gopurs_runtime.Value {
	return _Gopurs_Control_Apply_ArrayApply
}
