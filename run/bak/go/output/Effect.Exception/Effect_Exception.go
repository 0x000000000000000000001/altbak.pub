package Effect_Exception

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Effect "gopurs/output/Effect"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_try gopurs_runtime.Value
var once_try sync.Once
func Get_try() gopurs_runtime.Value {
	once_try.Do(func() {
		cache_try = gopurs_runtime.Func(func(action_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_try(action_0_box)
})
	})
	return cache_try
}

var cache_throw gopurs_runtime.Value
var once_throw sync.Once
func Get_throw() gopurs_runtime.Value {
	once_throw.Do(func() {
		cache_throw = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_throw(x_0_box.StrVal())
})
	})
	return cache_throw
}

var cache_stack gopurs_runtime.Value
var once_stack sync.Once
func Get_stack() gopurs_runtime.Value {
	once_stack.Do(func() {
		cache_stack = gopurs_runtime.Apply2(Get_stackImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_stack
}

var cache_showError gopurs_runtime.Value
var once_showError sync.Once
func Get_showError() gopurs_runtime.Value {
	once_showError.Do(func() {
		cache_showError = gopurs_runtime.RecordDict1("show", Get_showErrorImpl())
	})
	return cache_showError
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__113987891 gopurs_runtime.Value
var once_map__113987891 sync.Once
func Get_map__113987891() gopurs_runtime.Value {
	once_map__113987891.Do(func() {
		cache_map__113987891 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__113987891(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_map__113987891
}

var cache_functorEffect__347161653 gopurs_runtime.Value
var once_functorEffect__347161653 sync.Once
func Get_functorEffect__347161653() gopurs_runtime.Value {
	once_functorEffect__347161653.Do(func() {
		cache_functorEffect__347161653 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__347161653
}

func Call_try(action_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var action_0 gopurs_runtime.Value = action_0_loop
_ = action_0
return gopurs_runtime.Apply2(Get_catchException(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_1})})
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), pkg_Data_Either.Get_Right(), action_0))
}

func Call_throw(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_throwException(), gopurs_runtime.Apply(Get_error(), gopurs_runtime.Str(x_0)))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__113987891(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), __eta0_0), __eta1_1)
}

func Get_catchException() gopurs_runtime.Value {
	return _Gopurs_CatchException
}

func Get_error() gopurs_runtime.Value {
	return _Gopurs_Error
}

func Get_errorWithCause() gopurs_runtime.Value {
	return _Gopurs_ErrorWithCause
}

func Get_errorWithName() gopurs_runtime.Value {
	return _Gopurs_ErrorWithName
}

func Get_message() gopurs_runtime.Value {
	return _Gopurs_Message
}

func Get_name() gopurs_runtime.Value {
	return _Gopurs_Name
}

func Get_showErrorImpl() gopurs_runtime.Value {
	return _Gopurs_ShowErrorImpl
}

func Get_stackImpl() gopurs_runtime.Value {
	return _Gopurs_StackImpl
}

func Get_throwException() gopurs_runtime.Value {
	return _Gopurs_ThrowException
}
