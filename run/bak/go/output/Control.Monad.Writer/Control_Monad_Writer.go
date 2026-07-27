package Control_Monad_Writer

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var cache_writer gopurs_runtime.Value
var once_writer sync.Once
func Get_writer() gopurs_runtime.Value {
	once_writer.Do(func() {
		cache_writer = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_writer(x_0_box)
})
	})
	return cache_writer
}

var cache_runWriter gopurs_runtime.Value
var once_runWriter sync.Once
func Get_runWriter() gopurs_runtime.Value {
	once_runWriter.Do(func() {
		cache_runWriter = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runWriter(x_0_box)
})
	})
	return cache_runWriter
}

var cache_mapWriter gopurs_runtime.Value
var once_mapWriter sync.Once
func Get_mapWriter() gopurs_runtime.Value {
	once_mapWriter.Do(func() {
		cache_mapWriter = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWriter(func(inner_arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0_box, inner_arg0)
}, v_1_box)
})
	})
	return cache_mapWriter
}

var cache_execWriter gopurs_runtime.Value
var once_execWriter sync.Once
func Get_execWriter() gopurs_runtime.Value {
	once_execWriter.Do(func() {
		cache_execWriter = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_execWriter(m_0_box))
})
	})
	return cache_execWriter
}

func Call_writer(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_applicativeIdentity(), "pure"), x_0)
}

func Call_runWriter(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_mapWriter(f_0_loop func(gopurs_runtime.Value) gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 func(gopurs_runtime.Value) gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return f_0(v_1)
}

func Call_execWriter(m_0_loop gopurs_runtime.Value) interface{} {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_0.UnsafePtr).V1))
}
