package Effect_Exception_Unsafe

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Effect_Exception "gopurs/output/Effect.Exception"
	pkg_Effect_Unsafe "gopurs/output/Effect.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_unsafeThrowException gopurs_runtime.Value
var once_unsafeThrowException sync.Once
func Get_unsafeThrowException() gopurs_runtime.Value {
	once_unsafeThrowException.Do(func() {
		cache_unsafeThrowException = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeThrowException(x_0_box)
})
	})
	return cache_unsafeThrowException
}

var cache_unsafeThrow gopurs_runtime.Value
var once_unsafeThrow sync.Once
func Get_unsafeThrow() gopurs_runtime.Value {
	once_unsafeThrow.Do(func() {
		cache_unsafeThrow = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeThrow(x_0_box.StrVal())
})
	})
	return cache_unsafeThrow
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

var cache_unsafeThrowException__748557467 gopurs_runtime.Value
var once_unsafeThrowException__748557467 sync.Once
func Get_unsafeThrowException__748557467() gopurs_runtime.Value {
	once_unsafeThrowException__748557467.Do(func() {
		cache_unsafeThrowException__748557467 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeThrowException__748557467(x_0_box)
})
	})
	return cache_unsafeThrowException__748557467
}

func Call_unsafeThrowException(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.Apply(pkg_Effect_Exception.Get_throwException(), x_0))
}

func Call_unsafeThrow(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.Apply(pkg_Effect_Exception.Get_throwException(), gopurs_runtime.Apply(pkg_Effect_Exception.Get_error(), gopurs_runtime.Str(x_0))))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_unsafeThrowException__748557467(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.Apply(pkg_Effect_Exception.Get_throwException(), x_0))
}


