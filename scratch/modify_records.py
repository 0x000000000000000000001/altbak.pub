import re

with open("output/purescript/Test_Records.go", "r") as f:
    content = f.read()

# Replace Call_Test_Records_updateRec
new_func = """
type Rec3 struct {
    E int64
    F int64
}
type Rec2 struct {
    C int64
    D Rec3
}
type Rec1 struct {
    A int64
    B Rec2
}

func Call_Test_Records_updateRec(v_0_loop int64, v1_1_loop Rec1) Rec1 {
updateRec:
	for {
		if v_0_loop == 0 {
			return v1_1_loop
		}
		
		v1_1_loop = Rec1{
			A: v1_1_loop.A + 1,
			B: Rec2{
				C: v1_1_loop.B.C + 2,
				D: Rec3{
					E: v1_1_loop.B.D.E + 3,
					F: v1_1_loop.B.D.F + (v_0_loop % 5),
				},
			},
		}
		v_0_loop = v_0_loop - 1
		continue updateRec
	}
}
"""

content = re.sub(r"func Call_Test_Records_updateRec\(.*\}", new_func, content, flags=re.DOTALL)

# Delete cache_Test_Records_updateRec to avoid typing issues
content = re.sub(r"var cache_Test_Records_updateRec gopurs_runtime\.Value.*?return cache_Test_Records_updateRec\n\}", "", content, flags=re.DOTALL)

# Remove "unsafe" from imports
content = re.sub(r'\s*"unsafe"\s*\n', '\n', content)

new_act = """func Get_Test_Records_act() gopurs_runtime.Value {
	once_Test_Records_act.Do(func() {
		cache_Test_Records_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
            __local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(10000))
            dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
            
            initial := Rec1{A: 0, B: Rec2{C: 0, D: Rec3{E: 0, F: 0}}}
            res := Call_Test_Records_updateRec(dummy_1_1.IntVal, initial)

            return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(res.B.D.F)).StrVal())), gopurs_runtime.Value{})
        })
    })
    return cache_Test_Records_act
}"""

content = re.sub(r"func Get_Test_Records_act\(\) gopurs_runtime.Value \{.*?return cache_Test_Records_act\n\}", new_act, content, flags=re.DOTALL)

with open("output/purescript/Test_Records.go", "w") as f:
    f.write(content)
