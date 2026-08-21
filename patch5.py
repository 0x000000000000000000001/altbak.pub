import re

with open('output/purescript/Data_Map_Internal.go', 'r') as f:
    content = f.read()

# Restore from backup since patch4 corrupted it
import subprocess
subprocess.run(["./bin/go/run", "--clean"])

with open('output/purescript/Data_Map_Internal.go', 'r') as f:
    content = f.read()

# We need to extract the closure body and put it at the top level!
# It's easier to just do string replacements for the specific functions.

# For insert:
# The body is:
"""
var go__go_3_0_73 gopurs_runtime.Value
_ = go__go_3_0_73
go__go_3_0_73 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
...
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
"""

# Let's write a python script that finds the closure, extracts the body, replaces 'gopurs_runtime.Apply(go__go_3_0_73, arg)' with 'Call_insert_go(dictOrd_0_loop, k_1, v_2, arg)'
# And puts Call_insert_go at the top of the file!

# ACTUALLY, for a quick test, I'll just write the entire patched Data_Map_Internal.go myself? No, it's 450KB.
pass
