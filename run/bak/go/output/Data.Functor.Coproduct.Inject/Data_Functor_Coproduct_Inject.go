package Data_Functor_Coproduct_Inject

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Either "gopurs/output/Data.Either"
	unsafe "unsafe"
)

var cache_prj gopurs_runtime.Value
var once_prj sync.Once
func Get_prj() gopurs_runtime.Value {
	once_prj.Do(func() {
		cache_prj = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "prj")
}()
})
	})
	return cache_prj
}

var cache_injectReflexive gopurs_runtime.Value
var once_injectReflexive sync.Once
func Get_injectReflexive() gopurs_runtime.Value {
	once_injectReflexive.Do(func() {
		cache_injectReflexive = gopurs_runtime.RecordDict2("inj", "prj", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), pkg_Data_Maybe.Get_Just())
	})
	return cache_injectReflexive
}

var cache_injectLeft gopurs_runtime.Value
var once_injectLeft sync.Once
func Get_injectLeft() gopurs_runtime.Value {
	once_injectLeft.Do(func() {
		cache_injectLeft = gopurs_runtime.RecordDict2("inj", "prj", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{x_0})}
}), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Either.Data_Data_Either_Left)(v2_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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
	return cache_injectLeft
}

var cache_inj gopurs_runtime.Value
var once_inj sync.Once
func Get_inj() gopurs_runtime.Value {
	once_inj.Do(func() {
		cache_inj = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "inj")
}()
})
	})
	return cache_inj
}

var cache_injectRight gopurs_runtime.Value
var once_injectRight sync.Once
func Get_injectRight() gopurs_runtime.Value {
	once_injectRight.Do(func() {
		cache_injectRight = gopurs_runtime.Func(func(dictInject_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictInject_0 gopurs_runtime.Value = dictInject_0_loop
_ = dictInject_0
return gopurs_runtime.RecordDict2("inj", "prj", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictInject_0, "inj"), x_1)})}
}), gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_1.Type == 9 && v2_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictInject_0, "prj"), (*pkg_Data_Either.Data_Data_Either_Right)(v2_1.UnsafePtr).V0)
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
	return cache_injectRight
}




