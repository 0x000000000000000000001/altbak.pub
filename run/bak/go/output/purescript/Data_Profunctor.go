package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_identity gopurs_runtime.Value
var once_Data_Profunctor_identity sync.Once
func Get_Data_Profunctor_identity() gopurs_runtime.Value {
	once_Data_Profunctor_identity.Do(func() {
		cache_Data_Profunctor_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_identity(x_0_box)
})
	})
	return cache_Data_Profunctor_identity
}

var cache_Data_Profunctor_identity1 gopurs_runtime.Value
var once_Data_Profunctor_identity1 sync.Once
func Get_Data_Profunctor_identity1() gopurs_runtime.Value {
	once_Data_Profunctor_identity1.Do(func() {
		cache_Data_Profunctor_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_identity1(x_0_box)
})
	})
	return cache_Data_Profunctor_identity1
}

var cache_Data_Profunctor_wrap gopurs_runtime.Value
var once_Data_Profunctor_wrap sync.Once
func Get_Data_Profunctor_wrap() gopurs_runtime.Value {
	once_Data_Profunctor_wrap.Do(func() {
		cache_Data_Profunctor_wrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Profunctor_wrap
}

var cache_Data_Profunctor_unwrap gopurs_runtime.Value
var once_Data_Profunctor_unwrap sync.Once
func Get_Data_Profunctor_unwrap() gopurs_runtime.Value {
	once_Data_Profunctor_unwrap.Do(func() {
		cache_Data_Profunctor_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Profunctor_unwrap
}

var cache_Data_Profunctor_Profunctor_dollarDict gopurs_runtime.Value
var once_Data_Profunctor_Profunctor_dollarDict sync.Once
func Get_Data_Profunctor_Profunctor_dollarDict() gopurs_runtime.Value {
	once_Data_Profunctor_Profunctor_dollarDict.Do(func() {
		cache_Data_Profunctor_Profunctor_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Profunctor_dollarDict(x_0_box)
})
	})
	return cache_Data_Profunctor_Profunctor_dollarDict
}

var cache_Data_Profunctor_profunctorFn gopurs_runtime.Value
var once_Data_Profunctor_profunctorFn sync.Once
func Get_Data_Profunctor_profunctorFn() gopurs_runtime.Value {
	once_Data_Profunctor_profunctorFn.Do(func() {
		cache_Data_Profunctor_profunctorFn = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c2d_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c2d_1, gopurs_runtime.Apply(b2c_2, gopurs_runtime.Apply(a2b_0, x_3)))
})
})
})
}))
	})
	return cache_Data_Profunctor_profunctorFn
}

var cache_Data_Profunctor_dimap gopurs_runtime.Value
var once_Data_Profunctor_dimap sync.Once
func Get_Data_Profunctor_dimap() gopurs_runtime.Value {
	once_Data_Profunctor_dimap.Do(func() {
		cache_Data_Profunctor_dimap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_dimap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_dimap
}

var cache_Data_Profunctor_lcmap gopurs_runtime.Value
var once_Data_Profunctor_lcmap sync.Once
func Get_Data_Profunctor_lcmap() gopurs_runtime.Value {
	once_Data_Profunctor_lcmap.Do(func() {
		cache_Data_Profunctor_lcmap = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, a2b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_lcmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), a2b_1_box)
})
	})
	return cache_Data_Profunctor_lcmap
}

var cache_Data_Profunctor_rmap gopurs_runtime.Value
var once_Data_Profunctor_rmap sync.Once
func Get_Data_Profunctor_rmap() gopurs_runtime.Value {
	once_Data_Profunctor_rmap.Do(func() {
		cache_Data_Profunctor_rmap = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, b2c_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_rmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), b2c_1_box)
})
	})
	return cache_Data_Profunctor_rmap
}

var cache_Data_Profunctor_unwrapIso gopurs_runtime.Value
var once_Data_Profunctor_unwrapIso sync.Once
func Get_Data_Profunctor_unwrapIso() gopurs_runtime.Value {
	once_Data_Profunctor_unwrapIso.Do(func() {
		cache_Data_Profunctor_unwrapIso = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_unwrapIso(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_1_box))
})
	})
	return cache_Data_Profunctor_unwrapIso
}

var cache_Data_Profunctor_wrapIso gopurs_runtime.Value
var once_Data_Profunctor_wrapIso sync.Once
func Get_Data_Profunctor_wrapIso() gopurs_runtime.Value {
	once_Data_Profunctor_wrapIso.Do(func() {
		cache_Data_Profunctor_wrapIso = gopurs_runtime.Func3(func(dictProfunctor_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_wrapIso(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_1_box), v_2_box)
})
	})
	return cache_Data_Profunctor_wrapIso
}

var cache_Data_Profunctor_arr gopurs_runtime.Value
var once_Data_Profunctor_arr sync.Once
func Get_Data_Profunctor_arr() gopurs_runtime.Value {
	once_Data_Profunctor_arr.Do(func() {
		cache_Data_Profunctor_arr = gopurs_runtime.Func(func(dictCategory_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_arr(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](dictCategory_0_box))
})
	})
	return cache_Data_Profunctor_arr
}

var cache_Data_Profunctor_dimap__398801877 gopurs_runtime.Value
var once_Data_Profunctor_dimap__398801877 sync.Once
func Get_Data_Profunctor_dimap__398801877() gopurs_runtime.Value {
	once_Data_Profunctor_dimap__398801877.Do(func() {
		cache_Data_Profunctor_dimap__398801877 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_dimap__398801877(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_dimap__398801877
}

var cache_Data_Profunctor_dimap__413237333 gopurs_runtime.Value
var once_Data_Profunctor_dimap__413237333 sync.Once
func Get_Data_Profunctor_dimap__413237333() gopurs_runtime.Value {
	once_Data_Profunctor_dimap__413237333.Do(func() {
		cache_Data_Profunctor_dimap__413237333 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_dimap__413237333(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_dimap__413237333
}

var cache_Data_Profunctor_dimap__1466332548 gopurs_runtime.Value
var once_Data_Profunctor_dimap__1466332548 sync.Once
func Get_Data_Profunctor_dimap__1466332548() gopurs_runtime.Value {
	once_Data_Profunctor_dimap__1466332548.Do(func() {
		cache_Data_Profunctor_dimap__1466332548 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_dimap__1466332548(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_dimap__1466332548
}

var cache_Data_Profunctor_dimap__2995647076 gopurs_runtime.Value
var once_Data_Profunctor_dimap__2995647076 sync.Once
func Get_Data_Profunctor_dimap__2995647076() gopurs_runtime.Value {
	once_Data_Profunctor_dimap__2995647076.Do(func() {
		cache_Data_Profunctor_dimap__2995647076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_dimap__2995647076(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_dimap__2995647076
}

var cache_Data_Profunctor_dimap__109135140 gopurs_runtime.Value
var once_Data_Profunctor_dimap__109135140 sync.Once
func Get_Data_Profunctor_dimap__109135140() gopurs_runtime.Value {
	once_Data_Profunctor_dimap__109135140.Do(func() {
		cache_Data_Profunctor_dimap__109135140 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_dimap__109135140(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_dimap__109135140
}

var cache_Data_Profunctor_dimap__2951692932 gopurs_runtime.Value
var once_Data_Profunctor_dimap__2951692932 sync.Once
func Get_Data_Profunctor_dimap__2951692932() gopurs_runtime.Value {
	once_Data_Profunctor_dimap__2951692932.Do(func() {
		cache_Data_Profunctor_dimap__2951692932 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_dimap__2951692932(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_dimap__2951692932
}

var cache_Data_Profunctor_lcmap__1762133278 gopurs_runtime.Value
var once_Data_Profunctor_lcmap__1762133278 sync.Once
func Get_Data_Profunctor_lcmap__1762133278() gopurs_runtime.Value {
	once_Data_Profunctor_lcmap__1762133278.Do(func() {
		cache_Data_Profunctor_lcmap__1762133278 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, a2b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_lcmap__1762133278(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), a2b_1_box)
})
	})
	return cache_Data_Profunctor_lcmap__1762133278
}

var cache_Data_Profunctor_lcmap__1379840958 gopurs_runtime.Value
var once_Data_Profunctor_lcmap__1379840958 sync.Once
func Get_Data_Profunctor_lcmap__1379840958() gopurs_runtime.Value {
	once_Data_Profunctor_lcmap__1379840958.Do(func() {
		cache_Data_Profunctor_lcmap__1379840958 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, a2b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_lcmap__1379840958(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), a2b_1_box)
})
	})
	return cache_Data_Profunctor_lcmap__1379840958
}

var cache_Data_Profunctor_lcmap__351678174 gopurs_runtime.Value
var once_Data_Profunctor_lcmap__351678174 sync.Once
func Get_Data_Profunctor_lcmap__351678174() gopurs_runtime.Value {
	once_Data_Profunctor_lcmap__351678174.Do(func() {
		cache_Data_Profunctor_lcmap__351678174 = gopurs_runtime.Func(func(a2b_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_lcmap__351678174(a2b_0_box)
})
	})
	return cache_Data_Profunctor_lcmap__351678174
}

var cache_Data_Profunctor_profunctorFn__542207281 gopurs_runtime.Value
var once_Data_Profunctor_profunctorFn__542207281 sync.Once
func Get_Data_Profunctor_profunctorFn__542207281() gopurs_runtime.Value {
	once_Data_Profunctor_profunctorFn__542207281.Do(func() {
		cache_Data_Profunctor_profunctorFn__542207281 = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c2d_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c2d_1, gopurs_runtime.Apply(b2c_2, gopurs_runtime.Apply(a2b_0, x_3)))
})
})
})
}))
	})
	return cache_Data_Profunctor_profunctorFn__542207281
}

var cache_Data_Profunctor_profunctorFn__3736629211 gopurs_runtime.Value
var once_Data_Profunctor_profunctorFn__3736629211 sync.Once
func Get_Data_Profunctor_profunctorFn__3736629211() gopurs_runtime.Value {
	once_Data_Profunctor_profunctorFn__3736629211.Do(func() {
		cache_Data_Profunctor_profunctorFn__3736629211 = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c2d_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c2d_1, gopurs_runtime.Apply(b2c_2, gopurs_runtime.Apply(a2b_0, x_3)))
})
})
})
}))
	})
	return cache_Data_Profunctor_profunctorFn__3736629211
}

var cache_Data_Profunctor_rmap__1762133278 gopurs_runtime.Value
var once_Data_Profunctor_rmap__1762133278 sync.Once
func Get_Data_Profunctor_rmap__1762133278() gopurs_runtime.Value {
	once_Data_Profunctor_rmap__1762133278.Do(func() {
		cache_Data_Profunctor_rmap__1762133278 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, b2c_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_rmap__1762133278(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), b2c_1_box)
})
	})
	return cache_Data_Profunctor_rmap__1762133278
}

var cache_Data_Profunctor_rmap__2905617982 gopurs_runtime.Value
var once_Data_Profunctor_rmap__2905617982 sync.Once
func Get_Data_Profunctor_rmap__2905617982() gopurs_runtime.Value {
	once_Data_Profunctor_rmap__2905617982.Do(func() {
		cache_Data_Profunctor_rmap__2905617982 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, b2c_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_rmap__2905617982(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0_box), b2c_1_box)
})
	})
	return cache_Data_Profunctor_rmap__2905617982
}

type Constructor_Data_Profunctor_Profunctor[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2367018778] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "dimap": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Profunctor_Profunctor: " + key)
		}
	}
}


func Call_Data_Profunctor_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_Profunctor_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_dimap(dict_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Profunctor_lcmap(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], a2b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var a2b_1 gopurs_runtime.Value = a2b_1_loop
_ = a2b_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictProfunctor_0.V0), a2b_1, Get_Data_Profunctor_identity())
}

func Call_Data_Profunctor_rmap(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], b2c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var b2c_1 gopurs_runtime.Value = b2c_1_loop
_ = b2c_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictProfunctor_0.V0), Get_Data_Profunctor_identity1(), b2c_1)
}

func Call_Data_Profunctor_unwrapIso(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], _dollar__unused_1_loop *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var _dollar__unused_1 *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_1_loop
_ = _dollar__unused_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictProfunctor_0.V0), Get_Unsafe_Coerce_unsafeCoerce(), Get_Unsafe_Coerce_unsafeCoerce())
}

func Call_Data_Profunctor_wrapIso(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], _dollar__unused_1_loop *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var _dollar__unused_1 *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictProfunctor_0.V0), Get_Unsafe_Coerce_unsafeCoerce(), Get_Unsafe_Coerce_unsafeCoerce())
}

func Call_Data_Profunctor_arr(dictCategory_0_loop *Constructor_Control_Category_Category[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictCategory_0 *Constructor_Control_Category_Category[gopurs_runtime.Value] = dictCategory_0_loop
_ = dictCategory_0
// TAST (Let): identity2_1_0 -> gopurs_runtime.Value
identity2_1_0 := gopurs_runtime.Box(dictCategory_0.V1)
_ = identity2_1_0
return gopurs_runtime.Func(func(dictProfunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_2, "dimap"), Get_Data_Profunctor_identity1(), f_3, identity2_1_0)
})
})
}

func Call_Data_Profunctor_dimap__398801877(dict_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Profunctor_dimap__413237333(dict_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Profunctor_dimap__1466332548(dict_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Profunctor_dimap__2995647076(dict_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Profunctor_dimap__109135140(dict_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Profunctor_dimap__2951692932(dict_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Profunctor_lcmap__1762133278(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], a2b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var a2b_1 gopurs_runtime.Value = a2b_1_loop
_ = a2b_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictProfunctor_0.V0), a2b_1, Get_Data_Profunctor_identity())
}

func Call_Data_Profunctor_lcmap__1379840958(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], a2b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var a2b_1 gopurs_runtime.Value = a2b_1_loop
_ = a2b_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictProfunctor_0.V0), a2b_1, Get_Data_Profunctor_identity())
}

func Call_Data_Profunctor_lcmap__351678174(a2b_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a2b_0 gopurs_runtime.Value = a2b_0_loop
_ = a2b_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Profunctor_profunctorFn(), "dimap"), a2b_0, Get_Data_Profunctor_identity())
}

func Call_Data_Profunctor_rmap__1762133278(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], b2c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var b2c_1 gopurs_runtime.Value = b2c_1_loop
_ = b2c_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictProfunctor_0.V0), Get_Data_Profunctor_identity1(), b2c_1)
}

func Call_Data_Profunctor_rmap__2905617982(dictProfunctor_0_loop *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value], b2c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 *Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value] = dictProfunctor_0_loop
_ = dictProfunctor_0
var b2c_1 gopurs_runtime.Value = b2c_1_loop
_ = b2c_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictProfunctor_0.V0), Get_Data_Profunctor_identity1(), b2c_1)
}


