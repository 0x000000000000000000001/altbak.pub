with open("README.md", "r") as f:
    lines = f.readlines()

ffi = {
    "AST Evaluation": "685",
    "Fibonacci": "271",
    "List Processing": "226",
    "Tail Call Optimization": "298",
    "Deep Record Updates": "1543",
    "Ackermann": "237",
    "Church Numerals": "1507",
    "Prime Sieve": "306",
    "Red-Black Tree": "56485",
    "Polymorphism": "209291",
    "State Monad": "407",
    "Lazy Evaluation": "526",
    "Array Processing": "302",
    "RowToList": "285" # RowToList is not in the original table but let's just ignore it if missing
}

fficc = {
    "AST Evaluation": "854",
    "Fibonacci": "179",
    "List Processing": "259",
    "Tail Call Optimization": "269",
    "Deep Record Updates": "220",
    "Ackermann": "251",
    "Church Numerals": "196",
    "Prime Sieve": "1476",
    "Red-Black Tree": "41541",
    "Polymorphism": "21875",
    "State Monad": "264",
    "Lazy Evaluation": "314",
    "Array Processing": "229",
    "RowToList": "221"
}

in_erlang = False
for i, line in enumerate(lines):
    if "#### Erlang" in line:
        in_erlang = True
    if "#### Go" in line:
        in_erlang = False
    
    if in_erlang and "|" in line and "---" not in line and "Erlang Benchmark" not in line and "Total Execution Time" not in line:
        parts = line.split("|")
        # if no leading |, parts[0] is the name
        name_idx = 1 if line.startswith("|") else 0
        name = parts[name_idx].strip()
        
        if name in ffi:
            parts[name_idx + 2] = f" ~ {ffi[name]} μs "
            parts[name_idx + 3] = f" ~ {fficc[name]} μs "
            # if line had trailing \n, parts[-1] might have it, but join will preserve it if we replace the middle parts
            
            # The line ended with `|\n` so parts[-1] is `\n`.
            lines[i] = "|".join(parts)

with open("README.md", "w") as f:
    f.writelines(lines)
