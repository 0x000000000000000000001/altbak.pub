package Control_Applicative

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_pure gopurs_runtime.Value
var once_pure sync.Once
func Get_pure() gopurs_runtime.Value {
	once_pure.Do(func() {
		cache_pure = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure(gopurs_runtime.CoerceToStruct[Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure
}

var cache_unless gopurs_runtime.Value
var once_unless sync.Once
func Get_unless() gopurs_runtime.Value {
	once_unless.Do(func() {
		cache_unless = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unless(gopurs_runtime.CoerceToStruct[Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_unless
}

var cache_when gopurs_runtime.Value
var once_when sync.Once
func Get_when() gopurs_runtime.Value {
	once_when.Do(func() {
		cache_when = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_when(gopurs_runtime.CoerceToStruct[Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_when
}

var cache_liftA1 gopurs_runtime.Value
var once_liftA1 sync.Once
func Get_liftA1() gopurs_runtime.Value {
	once_liftA1.Do(func() {
		cache_liftA1 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_liftA1
}

var cache_applicativeProxy gopurs_runtime.Value
var once_applicativeProxy sync.Once
func Get_applicativeProxy() gopurs_runtime.Value {
	once_applicativeProxy.Do(func() {
		cache_applicativeProxy = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
}))
	})
	return cache_applicativeProxy
}

var cache_applicativeFn gopurs_runtime.Value
var once_applicativeFn sync.Once
func Get_applicativeFn() gopurs_runtime.Value {
	once_applicativeFn.Do(func() {
		cache_applicativeFn = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
}))
	})
	return cache_applicativeFn
}

var cache_applicativeArray gopurs_runtime.Value
var once_applicativeArray sync.Once
func Get_applicativeArray() gopurs_runtime.Value {
	once_applicativeArray.Do(func() {
		cache_applicativeArray = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
}))
	})
	return cache_applicativeArray
}

var cache_pure__2935994064 gopurs_runtime.Value
var once_pure__2935994064 sync.Once
func Get_pure__2935994064() gopurs_runtime.Value {
	once_pure__2935994064.Do(func() {
		cache_pure__2935994064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__2935994064(gopurs_runtime.CoerceToStruct[Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__2935994064
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__1449138640 gopurs_runtime.Value
var once_pure__1449138640 sync.Once
func Get_pure__1449138640() gopurs_runtime.Value {
	once_pure__1449138640.Do(func() {
		cache_pure__1449138640 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1449138640(gopurs_runtime.CoerceToStruct[Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__1449138640
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_applyArray__2998472828 gopurs_runtime.Value
var once_applyArray__2998472828 sync.Once
func Get_applyArray__2998472828() gopurs_runtime.Value {
	once_applyArray__2998472828.Do(func() {
		cache_applyArray__2998472828 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), pkg_Control_Apply.Get_arrayApply())
	})
	return cache_applyArray__2998472828
}

var cache_applyFn__4042184691 gopurs_runtime.Value
var once_applyFn__4042184691 sync.Once
func Get_applyFn__4042184691() gopurs_runtime.Value {
	once_applyFn__4042184691.Do(func() {
		cache_applyFn__4042184691 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_applyFn__4042184691
}

var cache_applyProxy__315643445 gopurs_runtime.Value
var once_applyProxy__315643445 sync.Once
func Get_applyProxy__315643445() gopurs_runtime.Value {
	once_applyProxy__315643445.Do(func() {
		cache_applyProxy__315643445 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_applyProxy__315643445
}

var cache_functorArray__361387505 gopurs_runtime.Value
var once_functorArray__361387505 sync.Once
func Get_functorArray__361387505() gopurs_runtime.Value {
	once_functorArray__361387505.Do(func() {
		cache_functorArray__361387505 = gopurs_runtime.RecordDict1("map", pkg_Data_Functor.Get_arrayMap())
	})
	return cache_functorArray__361387505
}

var cache_functorFn__20325936 gopurs_runtime.Value
var once_functorFn__20325936 sync.Once
func Get_functorFn__20325936() gopurs_runtime.Value {
	once_functorFn__20325936.Do(func() {
		cache_functorFn__20325936 = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))
	})
	return cache_functorFn__20325936
}

var cache_functorProxy__1157108209 gopurs_runtime.Value
var once_functorProxy__1157108209 sync.Once
func Get_functorProxy__1157108209() gopurs_runtime.Value {
	once_functorProxy__1157108209.Do(func() {
		cache_functorProxy__1157108209 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_functorProxy__1157108209
}

type Constructor_Applicative[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1459134221] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Applicative[gopurs_runtime.Value])(ptr)
		switch key {
		case "Apply0": return c.V0
		case "pure": return c.V1
		default: panic("Key not found in dictionary Constructor_Applicative: " + key)
		}
	}
}


func Call_pure(dict_0_loop *Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unless(dictApplicative_0_loop *Constructor_Applicative[gopurs_runtime.Value], v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
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
__t0 = gopurs_runtime.Apply(dictApplicative_0.V1, pkg_Data_Unit.Get_unit())
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

func Call_when(dictApplicative_0_loop *Constructor_Applicative[gopurs_runtime.Value], v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
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
__t0 = gopurs_runtime.Apply(dictApplicative_0.V1, pkg_Data_Unit.Get_unit())
}
end_branch_0:
return __t0
}

func Call_liftA1(dictApplicative_0_loop *Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}))
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply(dictApplicative_0.V1, f_2), a_3)
})
})
}

func Call_pure__2935994064(dict_0_loop *Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1449138640(dict_0_loop *Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


