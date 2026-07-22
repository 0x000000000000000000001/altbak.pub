package Control_Monad_Writer_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var tell gopurs_runtime.Value
var once_tell sync.Once
func Get_tell() gopurs_runtime.Value {
	once_tell.Do(func() {
		tell = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["tell"]
})
	})
	return tell
}

var pass gopurs_runtime.Value
var once_pass sync.Once
func Get_pass() gopurs_runtime.Value {
	once_pass.Do(func() {
		pass = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["pass"]
})
	})
	return pass
}

var listen gopurs_runtime.Value
var once_listen sync.Once
func Get_listen() gopurs_runtime.Value {
	once_listen.Do(func() {
		listen = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["listen"]
})
	})
	return listen
}

var listens gopurs_runtime.Value
var once_listens sync.Once
func Get_listens() gopurs_runtime.Value {
	once_listens.Do(func() {
		listens = gopurs_runtime.Func(func(dictMonadWriter_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadTell1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Monad1"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["listen"], m_3)), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_2, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])}))
}))
})
})
})
	})
	return listens
}

var censor gopurs_runtime.Value
var once_censor sync.Once
func Get_censor() gopurs_runtime.Value {
	once_censor.Do(func() {
		censor = gopurs_runtime.Func(func(dictMonadWriter_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadTell1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Monad1"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["pass"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], m_3), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_4, "value1": f_2}))
})))
})
})
})
	})
	return censor
}


