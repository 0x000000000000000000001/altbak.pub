package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Exception_pure gopurs_runtime.Value
var once_Effect_Exception_pure sync.Once
func Get_Effect_Exception_pure() gopurs_runtime.Value {
	once_Effect_Exception_pure.Do(func() {
		cache_Effect_Exception_pure = Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))
	})
	return cache_Effect_Exception_pure
}

var cache_Effect_Exception_try gopurs_runtime.Value
var once_Effect_Exception_try sync.Once
func Get_Effect_Exception_try() gopurs_runtime.Value {
	once_Effect_Exception_try.Do(func() {
		cache_Effect_Exception_try = gopurs_runtime.Func(func(action_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Exception_try(action_0_box)
})
	})
	return cache_Effect_Exception_try
}

var cache_Effect_Exception_throw gopurs_runtime.Value
var once_Effect_Exception_throw sync.Once
func Get_Effect_Exception_throw() gopurs_runtime.Value {
	once_Effect_Exception_throw.Do(func() {
		cache_Effect_Exception_throw = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Exception_throwException(), Get_Effect_Exception_error())
	})
	return cache_Effect_Exception_throw
}

var cache_Effect_Exception_stack gopurs_runtime.Value
var once_Effect_Exception_stack sync.Once
func Get_Effect_Exception_stack() gopurs_runtime.Value {
	once_Effect_Exception_stack.Do(func() {
		cache_Effect_Exception_stack = gopurs_runtime.Apply2(Get_Effect_Exception_stackImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))})
	})
	return cache_Effect_Exception_stack
}

var cache_Effect_Exception_showError gopurs_runtime.Value
var once_Effect_Exception_showError sync.Once
func Get_Effect_Exception_showError() gopurs_runtime.Value {
	once_Effect_Exception_showError.Do(func() {
		cache_Effect_Exception_showError = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, Get_Effect_Exception_showErrorImpl()}))}
	})
	return cache_Effect_Exception_showError
}

func Call_Effect_Exception_try(action_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var action_0 gopurs_runtime.Value = action_0_loop
_ = action_0
return gopurs_runtime.Apply2(Get_Effect_Exception_catchException(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect())), Get_Data_Either_Left()), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, Get_Data_Either_Right(), action_0))
}

func Get_Effect_Exception_catchException() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_CatchException
}

func Get_Effect_Exception_error() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_Error
}

func Get_Effect_Exception_errorWithCause() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_ErrorWithCause
}

func Get_Effect_Exception_errorWithName() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_ErrorWithName
}

func Get_Effect_Exception_message() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_Message
}

func Get_Effect_Exception_name() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_Name
}

func Get_Effect_Exception_showErrorImpl() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_ShowErrorImpl
}

func Get_Effect_Exception_stackImpl() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_StackImpl
}

func Get_Effect_Exception_throwException() gopurs_runtime.Value {
	return _Gopurs_Effect_Exception_ThrowException
}
