import re

log_path = '/tmp/scheme_ffi.log'
readme_path = 'README.md'

with open(log_path, 'r') as f:
    log = f.read()

times_ffi = {}
times_fficc = {}

pattern = r"\(Test\)\n\n(.*?)\s(FFI|FFICheatcode)(?:\s\([^)]+\))?:\n.*?\n\(Execution time\)\n\n([\d.]+)\sμs"
for match in re.finditer(pattern, log, re.DOTALL):
    test_name = match.group(1).strip()
    test_type = match.group(2)
    time_us = float(match.group(3))
    
    if time_us >= 1000:
        v = f"{time_us / 1000.0:.3f}"
        if '.' in v:
            v = v.rstrip('0').rstrip('.')
        val = f"~ {v} ms"
    else:
        val = f"~ {time_us} μs"
    
    if test_type == "FFI" and test_name not in times_ffi:
        times_ffi[test_name] = val
    elif test_type == "FFICheatcode" and test_name not in times_fficc:
        times_fficc[test_name] = val

with open(readme_path, 'r') as f:
    lines = f.read().split('\n')

in_scheme_block = False
for i, line in enumerate(lines):
    if line.startswith('#### Scheme'):
        in_scheme_block = True
    elif line.startswith('#### Erlang'):
        in_scheme_block = False
        
    if in_scheme_block and '|' in line and not line.startswith('Scheme Benchmark') and not line.startswith('---'):
        cols = line.split('|')
        if len(cols) >= 4:
            test_name_full = cols[0].strip()
            base_name = test_name_full.split(' (')[0].strip()
            
            if test_name_full == "**Total Execution Time**":
                total_ffi_us = sum(float(v.replace('~ ', '').replace(' μs', '').replace(' ms', '')) * (1000 if 'ms' in v else 1) for v in times_ffi.values())
                total_fficc_us = sum(float(v.replace('~ ', '').replace(' μs', '').replace(' ms', '')) * (1000 if 'ms' in v else 1) for v in times_fficc.values())
                ffi_val = f"~ {total_ffi_us / 1000.0:.2f} ms"
                fficc_val = f"~ {total_fficc_us / 1000.0:.2f} ms"
                cols[2] = " " + ffi_val.ljust(17)
                cols[3] = " " + fficc_val.ljust(27)
                lines[i] = '|'.join(cols)
            elif base_name in times_ffi or base_name in times_fficc:
                ffi_val = times_ffi.get(base_name, "N/A")
                fficc_val = times_fficc.get(base_name, "N/A")
                cols[2] = " " + ffi_val.ljust(17)
                cols[3] = " " + fficc_val.ljust(27)
                lines[i] = '|'.join(cols)

with open(readme_path, 'w') as f:
    f.write('\n'.join(lines))
