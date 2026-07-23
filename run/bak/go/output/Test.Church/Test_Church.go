package Test_Church

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var zeroC gopurs_runtime.Value
var once_zeroC sync.Once
func Get_zeroC() gopurs_runtime.Value {
	once_zeroC.Do(func() {
		zeroC = gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
	})
	return zeroC
}

var toInt gopurs_runtime.Value
var once_toInt sync.Once
func Get_toInt() gopurs_runtime.Value {
	once_toInt.Do(func() {
		toInt = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(n_0, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(x_1.IntVal + gopurs_runtime.Int(1).IntVal)
}), gopurs_runtime.Int(0))
})
	})
	return toInt
}

var succC gopurs_runtime.Value
var once_succC sync.Once
func Get_succC() gopurs_runtime.Value {
	once_succC.Do(func() {
		succC = gopurs_runtime.Func3(func(n_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(n_0, f_1, x_2))
})
	})
	return succC
}

var mulC gopurs_runtime.Value
var once_mulC sync.Once
func Get_mulC() gopurs_runtime.Value {
	once_mulC.Do(func() {
		mulC = gopurs_runtime.Func4(func(m_0 gopurs_runtime.Value, n_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(m_0, gopurs_runtime.Apply(n_1, f_2), x_3)
})
	})
	return mulC
}

var fromInt gopurs_runtime.Value
var once_fromInt sync.Once
func Get_fromInt() gopurs_runtime.Value {
	once_fromInt.Do(func() {
		fromInt = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
fromInt:
for {
if false { continue fromInt }
var v_0 = v_0_loop
_ = v_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t1 = Get_zeroC()
goto end_branch_1
} else {

}
}
{
__local_var_1_0 := gopurs_runtime.Apply(Get_fromInt(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal))
_ = __local_var_1_0
__t1 = gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Apply2(__local_var_1_0, f_2, x_3))
})
}
end_branch_1:
return __t1
}
}()
})
	})
	return fromInt
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Church Numerals (100k Closure Applications):"))
	})
	return describe
}

var c10 gopurs_runtime.Value
var once_c10 sync.Once
func Get_c10() gopurs_runtime.Value {
	once_c10.Do(func() {
		c10 = gopurs_runtime.Apply(Get_fromInt(), gopurs_runtime.Int(10))
	})
	return c10
}

var c100 gopurs_runtime.Value
var once_c100 sync.Once
func Get_c100() gopurs_runtime.Value {
	once_c100.Do(func() {
		c100 = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_c10(), gopurs_runtime.Apply(Get_c10(), f_0), x_1)
})
	})
	return c100
}

var c10k gopurs_runtime.Value
var once_c10k sync.Once
func Get_c10k() gopurs_runtime.Value {
	once_c10k.Do(func() {
		c10k = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_c10(), gopurs_runtime.Apply(Get_c10(), gopurs_runtime.Apply(Get_c100(), f_0)), x_1)
})
	})
	return c10k
}

var c100k gopurs_runtime.Value
var once_c100k sync.Once
func Get_c100k() gopurs_runtime.Value {
	once_c100k.Do(func() {
		c100k = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_c10(), gopurs_runtime.Apply(Get_c10(), gopurs_runtime.Apply(Get_c100(), gopurs_runtime.Apply(Get_c10(), f_0))), x_1)
})
	})
	return c100k
}

var addC gopurs_runtime.Value
var once_addC sync.Once
func Get_addC() gopurs_runtime.Value {
	once_addC.Do(func() {
		addC = gopurs_runtime.Func4(func(m_0 gopurs_runtime.Value, n_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(m_0, f_2, gopurs_runtime.Apply2(n_1, f_2, x_3))
})
	})
	return addC
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply2(Get_c10(), gopurs_runtime.Apply(Get_c10(), gopurs_runtime.Apply(Get_c100(), gopurs_runtime.Apply(Get_c10(), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(x_0.IntVal + gopurs_runtime.Int(1).IntVal)
})))), gopurs_runtime.Int(0))))
	})
	return act
}


