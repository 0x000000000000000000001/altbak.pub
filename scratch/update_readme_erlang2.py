import re

log_path_ffi = '/tmp/ffi_out.log'
log_path_fficc = '/tmp/fficc_out.log'
readme_path = 'README.md'

with open(log_path_ffi, 'r') as f:
    log_ffi = f.read()
with open(log_path_fficc, 'r') as f:
    log_fficc = f.read()

times_ffi = {}
times_fficc = {}

pattern = r"\(Test\)\n\n(.*?)\s(FFI|FFICheatcode)(?:\s\([^)]+\))?:\n.*?\n\(Execution time\)\n\n([\d.]+)\sμs"
for match in re.finditer(pattern, log_ffi, re.DOTALL):
    test_name = match.group(1).strip()
    test_type = match.group(2)
    time_us = float(match.group(3))
    
    if time_us >= 1000:
        v = f"{time_us / 1000.0:.3f}"
        if '.' in v:
            v = v.rstrip('0').rstrip('.')
        val = f"~ {v} ms"
    else:
        v = str(time_us)
        if '.' in v:
            v = v.rstrip('0').rstrip('.')
        val = f"~ {v} μs"
    
    if test_type == "FFI" and test_name not in times_ffi:
        times_ffi[test_name] = val

for match in re.finditer(pattern, log_fficc, re.DOTALL):
    test_name = match.group(1).strip()
    test_type = match.group(2)
    time_us = float(match.group(3))
    
    if time_us >= 1000:
        v = f"{time_us / 1000.0:.3f}"
        if '.' in v:
            v = v.rstrip('0').rstrip('.')
        val = f"~ {v} ms"
    else:
        v = str(time_us)
        if '.' in v:
            v = v.rstrip('0').rstrip('.')
        val = f"~ {v} μs"
    
    if test_type == "FFICheatcode" and test_name not in times_fficc:
        times_fficc[test_name] = val

with open(readme_path, 'r') as f:
    lines = f.read().split('\n')

in_erlang_block = False
for i, line in enumerate(lines):
    if line.startswith('#### Erlang'):
        in_erlang_block = True
    elif line.startswith('#### '):
        in_erlang_block = False
        
    if in_erlang_block and '|' in line and not line.startswith('Erlang Benchmark') and not line.startswith('---'):
        cols = line.split('|')
        if len(cols) >= 4:
            name_idx = 1 if line.startswith('|') else 0
            test_name_full = cols[name_idx].strip()
            base_name = test_name_full.split(' (')[0].strip()
            
            if test_name_full == "**Total Execution Time**":
                total_ffi_us = sum(float(v.replace('~ ', '').replace(' μs', '').replace(' ms', '')) * (1000 if 'ms' in v else 1) for v in times_ffi.values())
                total_fficc_us = sum(float(v.replace('~ ', '').replace(' μs', '').replace(' ms', '')) * (1000 if 'ms' in v else 1) for v in times_fficc.values())
                ffi_val = f"~ {total_ffi_us / 1000.0:.2f} ms"
                fficc_val = f"~ {total_fficc_us / 1000.0:.2f} ms"
                cols[name_idx + 2] = " " + ffi_val.ljust(17)
                cols[name_idx + 3] = " " + fficc_val.ljust(27)
                lines[i] = '|'.join(cols)
            elif base_name in times_ffi or base_name in times_fficc:
                ffi_val = times_ffi.get(base_name, "N/A")
                fficc_val = times_fficc.get(base_name, "N/A")
                cols[name_idx + 2] = " " + ffi_val.ljust(17)
                cols[name_idx + 3] = " " + fficc_val.ljust(27)
                lines[i] = '|'.join(cols)

with open(readme_path, 'w') as f:
    f.write('\n'.join(lines))
print("README Erlang block updated!")
