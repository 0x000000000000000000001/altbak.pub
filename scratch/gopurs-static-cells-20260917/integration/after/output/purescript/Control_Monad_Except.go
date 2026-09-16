package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Except_unwrap gopurs_runtime.Value
var once_Control_Monad_Except_unwrap sync.Once
func Get_Control_Monad_Except_unwrap() gopurs_runtime.Value {
	once_Control_Monad_Except_unwrap.Do(func() {
		cache_Control_Monad_Except_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Control_Monad_Except_unwrap
}

var cache_Control_Monad_Except_withExcept gopurs_runtime.Value
var once_Control_Monad_Except_withExcept sync.Once
func Get_Control_Monad_Except_withExcept() gopurs_runtime.Value {
	once_Control_Monad_Except_withExcept.Do(func() {
		cache_Control_Monad_Except_withExcept = gopurs_runtime.Apply(Get_Control_Monad_Except_Trans_withExceptT(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Identity_functorIdentity()))})
	})
	return cache_Control_Monad_Except_withExcept
}

var cache_Control_Monad_Except_runExcept gopurs_runtime.Value
var once_Control_Monad_Except_runExcept sync.Once
func Get_Control_Monad_Except_runExcept() gopurs_runtime.Value {
	once_Control_Monad_Except_runExcept.Do(func() {
		cache_Control_Monad_Except_runExcept = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), Get_Control_Monad_Except_Trans_runExceptT())
	})
	return cache_Control_Monad_Except_runExcept
}

var cache_Control_Monad_Except_mapExcept gopurs_runtime.Value
var once_Control_Monad_Except_mapExcept sync.Once
func Get_Control_Monad_Except_mapExcept() gopurs_runtime.Value {
	once_Control_Monad_Except_mapExcept.Do(func() {
		cache_Control_Monad_Except_mapExcept = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_mapExcept(f_0_box)
})
	})
	return cache_Control_Monad_Except_mapExcept
}

func Call_Control_Monad_Except_mapExcept(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
__local_var_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Identity_Identity(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, v_2)
})
}


