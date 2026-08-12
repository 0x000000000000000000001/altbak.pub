package Control_Monad_Reader

import (
	pkg_Control_Monad_Reader_Trans "gopurs/output/Control.Monad.Reader.Trans"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_withReader gopurs_runtime.Value
var once_withReader sync.Once
func Get_withReader() gopurs_runtime.Value {
	once_withReader.Do(func() {
		cache_withReader = pkg_Control_Monad_Reader_Trans.Get_withReaderT()
	})
	return cache_withReader
}

var cache_runReader gopurs_runtime.Value
var once_runReader sync.Once
func Get_runReader() gopurs_runtime.Value {
	once_runReader.Do(func() {
		cache_runReader = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runReader(v_0_box, x_1_box)
})
	})
	return cache_runReader
}

var cache_mapReader gopurs_runtime.Value
var once_mapReader sync.Once
func Get_mapReader() gopurs_runtime.Value {
	once_mapReader.Do(func() {
		cache_mapReader = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapReader(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_mapReader
}

var cache_mapReaderT__552640602 gopurs_runtime.Value
var once_mapReaderT__552640602 sync.Once
func Get_mapReaderT__552640602() gopurs_runtime.Value {
	once_mapReaderT__552640602.Do(func() {
		cache_mapReaderT__552640602 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapReaderT__552640602(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_mapReaderT__552640602
}

var cache_withReaderT__552640602 gopurs_runtime.Value
var once_withReaderT__552640602 sync.Once
func Get_withReaderT__552640602() gopurs_runtime.Value {
	once_withReaderT__552640602.Do(func() {
		cache_withReaderT__552640602 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withReaderT__552640602(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_withReaderT__552640602
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

func Call_runReader(v_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(v_0, x_1)
}

func Call_mapReader(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_mapReaderT__552640602(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_withReaderT__552640602(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


