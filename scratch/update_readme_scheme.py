import re

readme_path = 'README.md'
log_path = '/tmp/scheme_ffi.log'

with open(readme_path, 'r') as f:
    readme = f.read()

with open(log_path, 'r') as f:
    log = f.read()

sections = log.split("Building...")
if len(sections) >= 3:
    ffi_log = sections[1]
    fficc_log = sections[2]
else:
    ffi_log = log
    fficc_log = ""

def extract_times(log_text):
    times = {}
    pattern = r"\(Test\)\n\n(.*?)(?:\sFFI|\sFFICheatcode|):\n.*?\(Execution time\)\n\n([\d.]+)\sμs"
    for match in re.finditer(pattern, log_text, re.DOTALL):
        test_name = match.group(1).strip()
        time_us = float(match.group(2))
        times[test_name] = f"~ {time_us / 1000.0:.3f} ms".rstrip('0').rstrip('.') + " ms" if time_us >= 1000 else f"~ {time_us} μs"
    return times

ffi_times = extract_times(ffi_log)
fficc_times = extract_times(fficc_log)

test_names = [
    "AST Evaluation",
    "Fibonacci",
    "List Processing (900 elements)",
    "Tail Call Optimization (100k calls)",
    "Deep Record Updates (10k iterations)",
    "Ackermann (3, 4)",
    "Church Numerals (100k Closure Applications)",
    "Prime Sieve (sum primes up to 500)",
    "Red-Black Tree (100k Worst-Case Insertions)",
    "Polymorphism (10M Type Class Dict Lookups)",
    "State Monad (1.2k Binds, 60 Stack Depth)",
    "Lazy Evaluation (1M Thunks Forced, 1k Depth)",
    "Array Processing (900 elements)",
    "RowToList (Keys Count)"
]

lines = readme.split('\n')
in_scheme_block = False
for i, line in enumerate(lines):
    if line.startswith('#### Scheme'):
        in_scheme_block = True
    elif line.startswith('#### Erlang'):
        in_scheme_block = False
        
    if in_scheme_block and line.startswith('| ') and not line.startswith('| Scheme Benchmark'):
        cols = line.split('|')
        if len(cols) >= 5:
            test_name = cols[1].strip()
            if test_name == "**Total Execution Time**":
                # Calculate total execution time for FFI and FFICheatcode
                total_ffi_us = sum(float(v.replace('~ ', '').replace(' μs', '').replace(' ms', '')) * (1000 if 'ms' in v else 1) for v in ffi_times.values() if v != "N/A")
                total_fficc_us = sum(float(v.replace('~ ', '').replace(' μs', '').replace(' ms', '')) * (1000 if 'ms' in v else 1) for v in fficc_times.values() if v != "N/A")
                ffi_val = f"~ {total_ffi_us / 1000.0:.2f} ms"
                fficc_val = f"~ {total_fficc_us / 1000.0:.2f} ms"
            elif test_name in test_names:
                ffi_val = ffi_times.get(test_name, "N/A")
                fficc_val = fficc_times.get(test_name, "N/A")
            else:
                continue
                
            cols[3] = " " + ffi_val.ljust(17)
            cols[4] = " " + fficc_val.ljust(27)
            lines[i] = '|'.join(cols)

with open(readme_path, 'w') as f:
    f.write('\n'.join(lines))
print("README updated for Scheme FFI and FFICheatcode")
