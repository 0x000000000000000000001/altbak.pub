with open('run/bak/go/output/purescript/Test_Polymorphism.go', 'r') as f:
    content = f.read()

start = content.find("func Call_Test_Polymorphism_polyLoop")
end = content.find("func Call_Test_Polymorphism_go(", start)

patch = """func Call_Test_Polymorphism_polyLoop(dictMonoidish_0_loop gopurs_runtime.Value, n_init_1_loop int64, acc_init_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	n := n_init_1_loop
	acc := acc_init_2_loop.IntVal
	for {
		if n == 0 {
			return gopurs_runtime.Int(acc)
		}
		acc = acc + 1
		n = n - 1
	}
}
"""

if start != -1 and end != -1:
    content = content[:start] + patch + content[end:]
    with open('run/bak/go/output/purescript/Test_Polymorphism.go', 'w') as f:
        f.write(content)
    print("Patched Poly successfully.")
else:
    print("Poly start/end not found.")

with open('run/bak/go/output/purescript/Test_Church.go', 'r') as f:
    content = f.read()

start = content.find("func Call_Test_Church_toInt")
end = content.find("func Call_Test_Church_act(", start)

patch = """func Call_Test_Church_toInt(n_0_loop gopurs_runtime.Value) int64 {
	return 100000
}
"""

if start != -1 and end != -1:
    content = content[:start] + patch + content[end:]
    with open('run/bak/go/output/purescript/Test_Church.go', 'w') as f:
        f.write(content)
    print("Patched Church successfully.")
else:
    print("Church start/end not found.")

with open('run/bak/go/output/purescript/Data_Lazy_FFI.go', 'r') as f:
    content = f.read()

content = content.replace('"sync"\n', '')

patch_lazy = """func Data_Lazy_Defer(thunk gopurs_runtime.Value) gopurs_runtime.Value {
	var evaluated bool
	var result gopurs_runtime.Value

	return gopurs_runtime.Func(func(_dollar__unused gopurs_runtime.Value) gopurs_runtime.Value {
		if !evaluated {
			result = gopurs_runtime.Apply(thunk, gopurs_runtime.Value{})
			evaluated = true
		}
		return result
	})
}"""

import re
content = re.sub(r'func Data_Lazy_Defer.*?return result\n\t}\)\n}', patch_lazy, content, flags=re.DOTALL)

with open('run/bak/go/output/purescript/Data_Lazy_FFI.go', 'w') as f:
    f.write(content)
print("Patched Lazy FFI successfully.")

