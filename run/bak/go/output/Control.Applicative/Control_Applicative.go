package Control_Applicative

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var pure gopurs_runtime.Value
var once_pure sync.Once
func Get_pure() gopurs_runtime.Value {
	once_pure.Do(func() {
		pure = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "pure")
}()
})
	})
	return pure
}

var unless gopurs_runtime.Value
var once_unless sync.Once
func Get_unless() gopurs_runtime.Value {
	once_unless.Do(func() {
		unless = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unless(dictApplicative_0_box, v_1_box, v1_2_box)
})
	})
	return unless
}

var when gopurs_runtime.Value
var once_when sync.Once
func Get_when() gopurs_runtime.Value {
	once_when.Do(func() {
		when = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_when(dictApplicative_0_box, v_1_box, v1_2_box)
})
	})
	return when
}

var liftA1 gopurs_runtime.Value
var once_liftA1 sync.Once
func Get_liftA1() gopurs_runtime.Value {
	once_liftA1.Do(func() {
		liftA1 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftA1(dictApplicative_0_box, f_1_box, a_2_box)
})
	})
	return liftA1
}

var applicativeProxy gopurs_runtime.Value
var once_applicativeProxy sync.Once
func Get_applicativeProxy() gopurs_runtime.Value {
	once_applicativeProxy.Do(func() {
		applicativeProxy = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Proxy")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyProxy()
}))
	})
	return applicativeProxy
}

var applicativeFn gopurs_runtime.Value
var once_applicativeFn sync.Once
func Get_applicativeFn() gopurs_runtime.Value {
	once_applicativeFn.Do(func() {
		applicativeFn = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyFn()
}))
	})
	return applicativeFn
}

var applicativeArray gopurs_runtime.Value
var once_applicativeArray sync.Once
func Get_applicativeArray() gopurs_runtime.Value {
	once_applicativeArray.Do(func() {
		applicativeArray = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array([]gopurs_runtime.Value{x_0})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
}))
	})
	return applicativeArray
}

func Call_unless(dictApplicative_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t0 gopurs_runtime.Value
{
if v_1.IntVal != 0 != true {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
if v_1.IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_Unit.Get_unit())
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

func Call_when(dictApplicative_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t0 gopurs_runtime.Value
{
if v_1.IntVal != 0 {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_Unit.Get_unit())
}
end_branch_0:
return __t0
}

func Call_liftA1(dictApplicative_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), f_1), a_2)
}


