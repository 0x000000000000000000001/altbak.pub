package Control_Monad_Cont

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var withCont gopurs_runtime.Value
var once_withCont sync.Once
func Get_withCont() gopurs_runtime.Value {
	once_withCont.Do(func() {
		withCont = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_2, x_3)
})))
})
	})
	return withCont
}

var runCont gopurs_runtime.Value
var once_runCont sync.Once
func Get_runCont() gopurs_runtime.Value {
	once_runCont.Do(func() {
		runCont = gopurs_runtime.Func2(func(cc_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(cc_0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, x_2)
}))
})
	})
	return runCont
}

var mapCont gopurs_runtime.Value
var once_mapCont sync.Once
func Get_mapCont() gopurs_runtime.Value {
	once_mapCont.Do(func() {
		mapCont = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, k_2))
})
	})
	return mapCont
}

var cont gopurs_runtime.Value
var once_cont sync.Once
func Get_cont() gopurs_runtime.Value {
	once_cont.Do(func() {
		cont = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, c_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_1, x_2)
}))
})
	})
	return cont
}


