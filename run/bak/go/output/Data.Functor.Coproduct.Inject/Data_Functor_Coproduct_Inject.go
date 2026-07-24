package Data_Functor_Coproduct_Inject

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
)

var prj gopurs_runtime.Value
var once_prj sync.Once
func Get_prj() gopurs_runtime.Value {
	once_prj.Do(func() {
		prj = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "prj")
}()
})
	})
	return prj
}

var injectReflexive gopurs_runtime.Value
var once_injectReflexive sync.Once
func Get_injectReflexive() gopurs_runtime.Value {
	once_injectReflexive.Do(func() {
		injectReflexive = gopurs_runtime.RecordDict2("inj", "prj", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), pkg_Data_Maybe.Get_Just())
	})
	return injectReflexive
}

var injectLeft gopurs_runtime.Value
var once_injectLeft sync.Once
func Get_injectLeft() gopurs_runtime.Value {
	once_injectLeft.Do(func() {
		injectLeft = gopurs_runtime.RecordDict2("inj", "prj", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Left", x_0)
}), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_0.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v2_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_0.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return injectLeft
}

var inj gopurs_runtime.Value
var once_inj sync.Once
func Get_inj() gopurs_runtime.Value {
	once_inj.Do(func() {
		inj = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "inj")
}()
})
	})
	return inj
}

var injectRight gopurs_runtime.Value
var once_injectRight sync.Once
func Get_injectRight() gopurs_runtime.Value {
	once_injectRight.Do(func() {
		injectRight = gopurs_runtime.Func(func(dictInject_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictInject_0 gopurs_runtime.Value = dictInject_0_loop
_ = dictInject_0
return gopurs_runtime.RecordDict2("inj", "prj", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictInject_0_loop, "inj"), x_1))
}), gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_1.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_1.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictInject_0_loop, "prj"), (*[1024]gopurs_runtime.Value)(v2_1.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}()
})
	})
	return injectRight
}




