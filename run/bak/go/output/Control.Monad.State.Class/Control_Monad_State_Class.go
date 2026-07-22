package Control_Monad_State_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var state gopurs_runtime.Value
var once_state sync.Once
func Get_state() gopurs_runtime.Value {
	once_state.Do(func() {
		state = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["state"]
})
	})
	return state
}

var put gopurs_runtime.Value
var once_put sync.Once
func Get_put() gopurs_runtime.Value {
	once_put.Do(func() {
		put = gopurs_runtime.Func(func(dictMonadState_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadState_0.PtrVal.(map[string]gopurs_runtime.Value)["state"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": pkg_Data_Unit.Get_unit(), "value1": s_1})
}))
})
})
	})
	return put
}

var modify_ gopurs_runtime.Value
var once_modify_ sync.Once
func Get_modify_() gopurs_runtime.Value {
	once_modify_.Do(func() {
		modify_ = gopurs_runtime.Func(func(dictMonadState_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadState_0.PtrVal.(map[string]gopurs_runtime.Value)["state"], gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": pkg_Data_Unit.Get_unit(), "value1": gopurs_runtime.Apply(f_1, s_2)})
}))
})
})
	})
	return modify_
}

var modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		modify = gopurs_runtime.Func(func(dictMonadState_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadState_0.PtrVal.(map[string]gopurs_runtime.Value)["state"], gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_3_0 := gopurs_runtime.Apply(f_1, s_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": s_prime_3_0, "value1": s_prime_3_0})
}))
})
})
	})
	return modify
}

var gets gopurs_runtime.Value
var once_gets sync.Once
func Get_gets() gopurs_runtime.Value {
	once_gets.Do(func() {
		gets = gopurs_runtime.Func(func(dictMonadState_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadState_0.PtrVal.(map[string]gopurs_runtime.Value)["state"], gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_1, s_2), "value1": s_2})
}))
})
})
	})
	return gets
}

var get gopurs_runtime.Value
var once_get sync.Once
func Get_get() gopurs_runtime.Value {
	once_get.Do(func() {
		get = gopurs_runtime.Func(func(dictMonadState_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadState_0.PtrVal.(map[string]gopurs_runtime.Value)["state"], gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": s_1, "value1": s_1})
}))
})
	})
	return get
}


