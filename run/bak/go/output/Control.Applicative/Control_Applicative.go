package Control_Applicative

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Control_Apply "gopurs/output/Control.Apply"
)

var pure gopurs_runtime.Value
var once_pure sync.Once
func Get_pure() gopurs_runtime.Value {
	once_pure.Do(func() {
		pure = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "pure")
})
	})
	return pure
}

var unless gopurs_runtime.Value
var once_unless sync.Once
func Get_unless() gopurs_runtime.Value {
	once_unless.Do(func() {
		unless = gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1.IntVal == 0)).IntVal != 0 {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
if (v_1).IntVal != 0 {
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
})
	})
	return unless
}

var when gopurs_runtime.Value
var once_when sync.Once
func Get_when() gopurs_runtime.Value {
	once_when.Do(func() {
		when = gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1).IntVal != 0 {
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
})
	})
	return when
}

var liftA1 gopurs_runtime.Value
var once_liftA1 sync.Once
func Get_liftA1() gopurs_runtime.Value {
	once_liftA1.Do(func() {
		liftA1 = gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), f_1), a_2)
})
	})
	return liftA1
}

var applicativeProxy gopurs_runtime.Value
var once_applicativeProxy sync.Once
func Get_applicativeProxy() gopurs_runtime.Value {
	once_applicativeProxy.Do(func() {
		applicativeProxy = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy"))
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


