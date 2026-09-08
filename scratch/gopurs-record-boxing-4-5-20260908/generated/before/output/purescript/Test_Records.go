package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_Records_updateRec gopurs_runtime.Value
var once_Test_Records_updateRec sync.Once
func Get_Test_Records_updateRec() gopurs_runtime.Value {
	once_Test_Records_updateRec.Do(func() {
		cache_Test_Records_updateRec = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Test_Records_updateRec(v_0_box.IntVal, func() struct{
	a int64
	b struct{
	c int64
	d struct{
	e int64
	f int64
}
}
} {
					orig := v1_1_box
					_ = orig
					clone := struct{
	a int64
	b struct{
	c int64
	d struct{
	e int64
	f int64
}
}
}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
					clone.b = func() struct{
	c int64
	d struct{
	e int64
	f int64
}
} {
					orig := gopurs_runtime.RecordGet(orig, "b")
					_ = orig
					clone := struct{
	c int64
	d struct{
	e int64
	f int64
}
}{}
					clone.c = gopurs_runtime.RecordGet(orig, "c").IntVal
					clone.d = func() struct{
	e int64
	f int64
} {
					orig := gopurs_runtime.RecordGet(orig, "d")
					_ = orig
					clone := struct{
	e int64
	f int64
}{}
					clone.e = gopurs_runtime.RecordGet(orig, "e").IntVal
					clone.f = gopurs_runtime.RecordGet(orig, "f").IntVal
					return clone
				}()
					return clone
				}()
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "b"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.a), func() gopurs_runtime.Value {
				orig := orig.b
				_ = orig
				return gopurs_runtime.RecordDict([]string{"c", "d"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.c), func() gopurs_runtime.Value {
				orig := orig.d
				_ = orig
				return gopurs_runtime.RecordDict([]string{"e", "f"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.e), gopurs_runtime.Int(orig.f)})
				}()})
				}()})
				}()
})
	})
	return cache_Test_Records_updateRec
}

var cache_Test_Records_initial gopurs_runtime.Value
var once_Test_Records_initial sync.Once
func Get_Test_Records_initial() gopurs_runtime.Value {
	once_Test_Records_initial.Do(func() {
		cache_Test_Records_initial = func() gopurs_runtime.Value {
				orig := struct{
	a int64
	b struct{
	c int64
	d struct{
	e int64
	f int64
}
}
}{int64(0), struct{
	c int64
	d struct{
	e int64
	f int64
}
}{int64(0), struct{
	e int64
	f int64
}{int64(0), int64(0)}}}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "b"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.a), func() gopurs_runtime.Value {
				orig := orig.b
				_ = orig
				return gopurs_runtime.RecordDict([]string{"c", "d"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.c), func() gopurs_runtime.Value {
				orig := orig.d
				_ = orig
				return gopurs_runtime.RecordDict([]string{"e", "f"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.e), gopurs_runtime.Int(orig.f)})
				}()})
				}()})
				}()
	})
	return cache_Test_Records_initial
}

var cache_Test_Records_describe gopurs_runtime.Value
var once_Test_Records_describe sync.Once
func Get_Test_Records_describe() gopurs_runtime.Value {
	once_Test_Records_describe.Do(func() {
		cache_Test_Records_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Deep Record Updates (10k iterations):"))
	})
	return cache_Test_Records_describe
}

var cache_Test_Records_act gopurs_runtime.Value
var once_Test_Records_act sync.Once
func Get_Test_Records_act() gopurs_runtime.Value {
	once_Test_Records_act.Do(func() {
		cache_Test_Records_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(int64(10000)))
_ = __local_var_0_0
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_Records_updateRec(__local_var_1_1.IntVal, func() struct{
	a int64
	b struct{
	c int64
	d struct{
	e int64
	f int64
}
}
} {
					orig := Get_Test_Records_initial()
					_ = orig
					clone := struct{
	a int64
	b struct{
	c int64
	d struct{
	e int64
	f int64
}
}
}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
					clone.b = func() struct{
	c int64
	d struct{
	e int64
	f int64
}
} {
					orig := gopurs_runtime.RecordGet(orig, "b")
					_ = orig
					clone := struct{
	c int64
	d struct{
	e int64
	f int64
}
}{}
					clone.c = gopurs_runtime.RecordGet(orig, "c").IntVal
					clone.d = func() struct{
	e int64
	f int64
} {
					orig := gopurs_runtime.RecordGet(orig, "d")
					_ = orig
					clone := struct{
	e int64
	f int64
}{}
					clone.e = gopurs_runtime.RecordGet(orig, "e").IntVal
					clone.f = gopurs_runtime.RecordGet(orig, "f").IntVal
					return clone
				}()
					return clone
				}()
					return clone
				}()).b.d.f)).StrVal())
})
	})
	return cache_Test_Records_act
}

func Call_Test_Records_updateRec(v_0_loop int64, v1_1_loop struct{
	a int64
	b struct{
	c int64
	d struct{
	e int64
	f int64
}
}
}) struct{
	a int64
	b struct{
	c int64
	d struct{
	e int64
	f int64
}
}
} {
updateRec:
for {
if false { continue updateRec }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 struct{
	a int64
	b struct{
	c int64
	d struct{
	e int64
	f int64
}
}
} = v1_1_loop
_ = v1_1
var __t0 struct{
	a int64
	b struct{
	c int64
	d struct{
	e int64
	f int64
}
}
}
{
if (v_0) == (int64(0)) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (int64(1))
v1_1_loop = func() struct{
	a int64
	b struct{
	c int64
	d struct{
	e int64
	f int64
}
}
} {
clone := v1_1
clone.a = (v1_1.a) + (int64(1))
clone.b = func() struct{
	c int64
	d struct{
	e int64
	f int64
}
} {
clone := v1_1.b
clone.c = (v1_1.b.c) + (int64(2))
clone.d = func() struct{
	e int64
	f int64
} {
clone := v1_1.b.d
clone.e = (v1_1.b.d.e) + (int64(3))
clone.f = (v1_1.b.d.f) + ((v_0) % (int64(5)))
return clone
}()
return clone
}()
return clone
}()
continue updateRec
__t0 = func() struct{
	a int64
	b struct{
	c int64
	d struct{
	e int64
	f int64
}
}
} { panic("unreachable") }()
}
end_branch_0:
return __t0
}
}


