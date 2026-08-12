package Control_Applicative

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
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

var cache_pure__gopurs_runtime_Value_3215807376 gopurs_runtime.Value
var once_pure__gopurs_runtime_Value_3215807376 sync.Once
func Get_pure__gopurs_runtime_Value_3215807376() gopurs_runtime.Value {
	once_pure__gopurs_runtime_Value_3215807376.Do(func() {
		cache_pure__gopurs_runtime_Value_3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__gopurs_runtime_Value_3215807376(gopurs_runtime.CoerceToStruct[Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__gopurs_runtime_Value_3215807376
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

var cache_unless__gopurs_runtime_Value_1954875249 gopurs_runtime.Value
var once_unless__gopurs_runtime_Value_1954875249 sync.Once
func Get_unless__gopurs_runtime_Value_1954875249() gopurs_runtime.Value {
	once_unless__gopurs_runtime_Value_1954875249.Do(func() {
		cache_unless__gopurs_runtime_Value_1954875249 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unless__gopurs_runtime_Value_1954875249(gopurs_runtime.CoerceToStruct[Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_unless__gopurs_runtime_Value_1954875249
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

var cache_when__gopurs_runtime_Value_1954875249 gopurs_runtime.Value
var once_when__gopurs_runtime_Value_1954875249 sync.Once
func Get_when__gopurs_runtime_Value_1954875249() gopurs_runtime.Value {
	once_when__gopurs_runtime_Value_1954875249.Do(func() {
		cache_when__gopurs_runtime_Value_1954875249 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_when__gopurs_runtime_Value_1954875249(gopurs_runtime.CoerceToStruct[Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_when__gopurs_runtime_Value_1954875249
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
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
}))
	})
	return cache_applicativeProxy
}

var cache_applicativeProxy__gopurs_runtime_Value_1913125352 gopurs_runtime.Value
var once_applicativeProxy__gopurs_runtime_Value_1913125352 sync.Once
func Get_applicativeProxy__gopurs_runtime_Value_1913125352() gopurs_runtime.Value {
	once_applicativeProxy__gopurs_runtime_Value_1913125352.Do(func() {
		cache_applicativeProxy__gopurs_runtime_Value_1913125352 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyProxy__gopurs_runtime_Value_315643445()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
}))
	})
	return cache_applicativeProxy__gopurs_runtime_Value_1913125352
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

var cache_applicativeFn__gopurs_runtime_Value_3751223912 gopurs_runtime.Value
var once_applicativeFn__gopurs_runtime_Value_3751223912 sync.Once
func Get_applicativeFn__gopurs_runtime_Value_3751223912() gopurs_runtime.Value {
	once_applicativeFn__gopurs_runtime_Value_3751223912.Do(func() {
		cache_applicativeFn__gopurs_runtime_Value_3751223912 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyFn__gopurs_runtime_Value_4042184691()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
}))
	})
	return cache_applicativeFn__gopurs_runtime_Value_3751223912
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

var cache_applicativeArray__gopurs_runtime_Value_1604836744 gopurs_runtime.Value
var once_applicativeArray__gopurs_runtime_Value_1604836744 sync.Once
func Get_applicativeArray__gopurs_runtime_Value_1604836744() gopurs_runtime.Value {
	once_applicativeArray__gopurs_runtime_Value_1604836744.Do(func() {
		cache_applicativeArray__gopurs_runtime_Value_1604836744 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray__gopurs_runtime_Value_2998472828()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
}))
	})
	return cache_applicativeArray__gopurs_runtime_Value_1604836744
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

func Call_pure__gopurs_runtime_Value_3215807376(dict_0_loop *Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_unless__gopurs_runtime_Value_1954875249(dictApplicative_0_loop *Constructor_Applicative[gopurs_runtime.Value], v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_when__gopurs_runtime_Value_1954875249(dictApplicative_0_loop *Constructor_Applicative[gopurs_runtime.Value], v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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


