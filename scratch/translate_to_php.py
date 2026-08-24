import os
import glob

files = glob.glob('src/Test/*FFI.js') + glob.glob('src/Test/*FFICheatcode.js')
for f in files:
    with open(f, 'r') as file:
        print(f"--- {f} ---")
        print(file.read())
