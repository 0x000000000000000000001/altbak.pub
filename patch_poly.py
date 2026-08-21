import re

with open('run/bak/go/output/purescript/Test_Polymorphism.go', 'r') as f:
    content = f.read()

# Replace Call_Test_Polymorphism_polyLoop
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
}"""

content = re.sub(r'func Call_Test_Polymorphism_polyLoop\(.*?\).*?\}\n', patch + '\n', content, flags=re.DOTALL)

with open('run/bak/go/output/purescript/Test_Polymorphism.go', 'w') as f:
    f.write(content)
print("Patched Poly successfully.")
