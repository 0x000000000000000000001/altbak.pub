import re
import subprocess

# Restore fresh
subprocess.run(["./bin/go/run", "--clean"])

with open('output/purescript/Data_Map_Internal.go', 'r') as f:
    content = f.read()

# We need to extract the original body of go__go_3_0_73
match = re.search(r'go__go_3_0_73 = gopurs_runtime\.Func\(func\(v1_4 gopurs_runtime\.Value\) gopurs_runtime\.Value \{(.*?)\n\}\)', content, re.DOTALL)
body = match.group(1)

# Now we rewrite the body to call Call_Data_Map_Internal_insert_go
body_rewritten = body.replace('gopurs_runtime.Apply(go__go_3_0_73, ', 'Call_Data_Map_Internal_insert_go(dictOrd_0, k_1, v_2, ')

# Now we define Call_Data_Map_Internal_insert_go at the top level
top_level_func = f"""
func Call_Data_Map_Internal_insert_go(dictOrd_0 *Constructor_Data_Ord_Ord, k_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {{{body_rewritten}
}}
"""

# Replace the original closure with a call to the top-level func
new_closure = """go__go_3_0_73 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Internal_insert_go(dictOrd_0, k_1, v_2, v1_4)
})"""

content = content.replace(match.group(0), new_closure)

# Inject the top-level func right before the Call_Data_Map_Internal_insert__4289641298 declaration
content = content.replace("func Call_Data_Map_Internal_insert__4289641298(", top_level_func + "\nfunc Call_Data_Map_Internal_insert__4289641298(")


with open('output/purescript/Data_Map_Internal.go', 'w') as f:
    f.write(content)

print("Patched successfully")
