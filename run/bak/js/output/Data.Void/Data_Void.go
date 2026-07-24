package Data_Void

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var absurd gopurs_runtime.Value
var once_absurd sync.Once
func Get_absurd() gopurs_runtime.Value {
	once_absurd.Do(func() {
		absurd = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
spin_1_0:
for {
if false { continue spin_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v_2_loop = v_2
continue spin_1_0
return gopurs_runtime.Value{}
}
}()
})
return gopurs_runtime.Apply(spin_1_0, a_0)
})
	})
	return absurd
}




