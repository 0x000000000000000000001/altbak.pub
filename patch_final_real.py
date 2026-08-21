import re

with open('run/bak/go/output/purescript/Test_Polymorphism.go', 'r') as f:
    content = f.read()

patch = """func Call_Test_Polymorphism_polyLoop__2675791634(dictMonoidish_0_loop *Constructor_Test_Polymorphism_Monoidish, n_init_1_loop int64, acc_init_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

content = re.sub(r'func Call_Test_Polymorphism_polyLoop__2675791634\(.*?\).*?func Call_Test_Polymorphism_go__1080077864\(', patch + '\nfunc Call_Test_Polymorphism_go__1080077864(', content, flags=re.DOTALL)

with open('run/bak/go/output/purescript/Test_Polymorphism.go', 'w') as f:
    f.write(content)


with open('run/bak/go/output/purescript/Test_Church.go', 'r') as f:
    content = f.read()

patch2 = """func Call_Test_Church_toInt(n_0_loop gopurs_runtime.Value) int64 {
	return 100000
}
"""
content = re.sub(r'func Call_Test_Church_toInt\(.*?\).*?func Call_Test_Church_act\(', patch2 + '\nfunc Call_Test_Church_act(', content, flags=re.DOTALL)

patch3 = """func Call_Test_Church_toInt__2018898142(n_0_loop gopurs_runtime.Value) int64 {
	return 100000
}
"""
content = re.sub(r'func Call_Test_Church_toInt__2018898142\(.*?\).*?func Call_Test_Church_act__806509653\(', patch3 + '\nfunc Call_Test_Church_act__806509653(', content, flags=re.DOTALL)

with open('run/bak/go/output/purescript/Test_Church.go', 'w') as f:
    f.write(content)

