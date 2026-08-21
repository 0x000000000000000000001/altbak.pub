import re

with open('output/purescript/Data_Map_Internal.go', 'r') as f:
    content = f.read()

# Replace Apply with native call
content = content.replace("gopurs_runtime.Apply(go__go_3_0_73, ", "go__go_3_0_73_native(")
content = content.replace("gopurs_runtime.Apply(go__go_3_0_17, ", "go__go_3_0_17_native(")

# Replace function declaration
content = re.sub(r'var (go__go_3_0_\d+) gopurs_runtime\.Value\n_ = \1\n\1 = gopurs_runtime\.Func\(func\((v1_\d+) gopurs_runtime\.Value\) gopurs_runtime\.Value \{', r'var \1 gopurs_runtime.Value\n_ = \1\nvar \1_native func(\2 gopurs_runtime.Value) gopurs_runtime.Value\n\1_native = func(\2 gopurs_runtime.Value) gopurs_runtime.Value {', content)

# Wrap before return
# For 17
content = content.replace("end_branch_3:\nreturn gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}\n})\nreturn go__go_3_0_17\n}", "end_branch_3:\nreturn gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}\n}\ngo__go_3_0_17 = gopurs_runtime.Func(go__go_3_0_17_native)\nreturn go__go_3_0_17\n}")

# For 73
content = content.replace("end_branch_3:\nreturn gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}\n})\nreturn go__go_3_0_73\n}", "end_branch_3:\nreturn gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}\n}\ngo__go_3_0_73 = gopurs_runtime.Func(go__go_3_0_73_native)\nreturn go__go_3_0_73\n}")

with open('output/purescript/Data_Map_Internal.go', 'w') as f:
    f.write(content)

print("Patched successfully")
