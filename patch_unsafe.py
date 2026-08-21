import re
import subprocess

# Clean first
subprocess.run(["./bin/go/run", "--clean"])

with open('output/purescript/Data_Map_Internal.go', 'r') as f:
    content = f.read()

# The pattern is: gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), arg1, arg2, arg3, arg4))
# We want to replace it with: Call_Data_Map_Internal_unsafeBalancedNode(arg1, arg2, arg3, arg4)

# Wait, UncurriedApp4 takes 5 arguments: the function, and 4 arguments.
# Let's match gopurs_runtime.CoerceToStruct\[Constructor_Data_Map_Internal_Node\]\(gopurs_runtime.UncurriedApp4\(Get_Data_Map_Internal_unsafeBalancedNode\(\), (.*?)\)\)

def replacer(m):
    args = m.group(1)
    return f"Call_Data_Map_Internal_unsafeBalancedNode({args})"

# We also need to remove the CoerceToStruct!
content = re.sub(r'gopurs_runtime\.CoerceToStruct\[Constructor_Data_Map_Internal_Node\]\(gopurs_runtime\.UncurriedApp4\(Get_Data_Map_Internal_unsafeBalancedNode\(\), (.*?)\)\)', replacer, content)

with open('output/purescript/Data_Map_Internal.go', 'w') as f:
    f.write(content)

print("Patched successfully")
