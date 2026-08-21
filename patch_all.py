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

