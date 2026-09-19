#!/usr/bin/env python3
"""Run the PureScript benchmark using sharpurs -> Fable -> Rust."""

import argparse
from datetime import datetime, timezone
import os
from pathlib import Path
import re
import shutil
import subprocess
import sys
import tempfile

sys.dont_write_bytecode = True
import run_benchmarks as references

ROOT = references.ROOT

def patch_fsharp_code(main_dir):
    """Modify the generated F# code to remove C# FFI and fix Fable unsupported methods."""
    fsproj = main_dir / "Program.fsproj"
    proj_content = fsproj.read_text()
    
    # Remove unsupported projects and files from fsproj
    removals = [
        r'<ProjectReference Include="FFI.CSharp.csproj" />',
        r'<Compile Include="Effect.Aff.fs" />',
        r'<Compile Include="Effect.Now.fs" />',
        r'<Compile Include="Effect.Exception.fs" />',
        r'<Compile Include="Sharpurs_Prelude.fs" />',
        r'<Compile Include="Control.Bind.fs" />',
        r'<Compile Include="Data.Array.ST.fs" />',
        r'<Compile Include="Control.Monad.ST.Internal.fs" />',
        r'<Compile Include="Data.String.NonEmpty.Internal.fs" />',
        r'<Compile Include="Data.String.NonEmpty.CodePoints.fs" />',
        r'<Compile Include="Data.String.NonEmpty.fs" />',
        r'<Compile Include="Data.String.NonEmpty.CaseInsensitive.fs" />',
        r'<Compile Include="Data.String.NonEmpty.CodeUnits.fs" />',

    ]
    for r in removals:
        proj_content = re.sub(r'^\s*' + re.escape(r) + r'\s*$', '', proj_content, flags=re.MULTILINE)


    
    # Add Fable.Core
    proj_content = proj_content.replace(
        '</ItemGroup>', 
        '  <PackageReference Include="Fable.Core" Version="4.3.0" />\n  </ItemGroup>'
    )
    fsproj.write_text(proj_content)

    # Patch Bench.fs (Stopwatch not supported)
    bench_fs = main_dir / "Bench.fs"
    bench_content = bench_fs.read_text()
    bench_content = bench_content.replace("System.Diagnostics.Stopwatch.GetTimestamp()", "DateTime.UtcNow.Ticks")
    bench_content = bench_content.replace("System.Diagnostics.Stopwatch.Frequency", "10000000L")
    bench_fs.write_text(bench_content)
    
    # Patch String FFI (Missing open System)
    for f in ["Data.String.Common.fs", "Data.String.CodePoints.fs", "Data.String.CodeUnits.fs"]:
        path = main_dir / f
        if path.exists():
            path.write_text("open System\n" + path.read_text())
            
    # Patch Number FFI (Math.Round expects different overload or uses System)
    num_fs = main_dir / "Data.Number.fs"
    if num_fs.exists():
        num_content = num_fs.read_text()
        num_content = num_content.replace("System.Math.Round", "Math.Round")
        num_fs.write_text(num_content)


def execute_and_parse(binary):
    """Execute the Rust binary and parse its stdout for benchmark results."""
    print("Running benchmark...", flush=True)
    env = os.environ.copy()
    output = subprocess.check_output([str(binary)], env=env, text=True)
    print(output)
    
    results = {}
    total = 0.0
    
    matches = re.finditer(r'\(Test\)\s+([^:]+):\s+\(Output & Warm-up\)\s+(.*?)\s+\(Execution time - best of 10\)\s+([0-9.]+)\s+μs', output, re.S)
    
    for match in matches:
        test_name = match.group(1).strip()
        # Some mappings to match README format
        if "List Processing" in test_name: test_name = "List Processing"
        elif "Tail Call Optimization" in test_name: test_name = "Tail Call Optimization"
        elif "Deep Record Updates" in test_name: test_name = "Deep Record Updates"
        elif "Ackermann" in test_name: test_name = "Ackermann"
        elif "Church Numerals" in test_name: test_name = "Church Numerals"
        elif "Prime Sieve" in test_name: test_name = "Prime Sieve"
        elif "Red-Black Tree" in test_name: test_name = "Red-Black Tree"
        elif "State Monad" in test_name: test_name = "State Monad"
        elif "Lazy Evaluation" in test_name: test_name = "Lazy Evaluation"
        elif "Array Processing" in test_name: test_name = "Array Processing"
        elif "RowToList" in test_name: test_name = "RowToList"
        elif "Fibonacci" in test_name: test_name = "Fibonacci"
        elif "AST Evaluation" in test_name: test_name = "AST Evaluation"
        elif "Polymorphism" in test_name: test_name = "Polymorphism"
        
        time_us = float(match.group(3))
        results[test_name] = time_us
        total += time_us / 1000.0
        
    return results, total

def render_readme(original, fable_results, fable_total):
    match = re.search(r"(#### Rust\n)(.*?)(?=\n#### )", original, re.S)
    section = match[2]
    
    # We update the "Compiled Rust (Fable, WIP)" column (which is the second data column)
    # The columns: 
    # [1] Name 
    # [2] purust column |
    # [3] Fable column |
    # [4] Rust FFI |
    # [5] Rust FFI |
    
    def replace_row(row_match):
        name = row_match.group(1).strip()
        if name in fable_results:
            new_val = f"~ {fable_results[name]:.3f} μs"
            # row_match groups:
            # 1: name
            # 2: separator
            # 3: purust cell
            # 4: separator
            # 5: fable cell
            # 6: separator
            # 7: rest
            
            # Since markdown regex can be tricky, let's split the line by pipe
            parts = row_match.group(0).split('|')
            # parts[0] is name
            # parts[1] is purust
            # parts[2] is fable
            # parts[3] is rustc
            parts[2] = f" {new_val:<26} "
            return "|".join(parts)
        return row_match.group(0)
    
    # Regex to capture a full table row
    section = re.sub(r"^([^|\n]+)\|([^|\n]+)\|([^|\n]+)\|([^|\n]+)\|([^|\n]+)\|", replace_row, section, flags=re.M)
    
    def replace_total(row_match):
        parts = row_match.group(0).split('|')
        parts[2] = f" ~ {fable_total:.2f} ms{' ':>17} "
        return "|".join(parts)
        
    section = re.sub(r"^(\*\*Total Execution Time\*\*\s*\|[^|\n]+\|)([^|\n]+)(\|)", replace_total, section, flags=re.M)
    
    return original[:match.start(2)] + section + original[match.end(2):]

def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--update-readme", action="store_true")
    args = parser.parse_args()
    
    original_readme = (ROOT / "README.md").read_text()
    
    env = os.environ.copy()
    env["PATH"] = str(ROOT / "run/bak/js/node_modules/.bin") + os.pathsep + env["PATH"] + os.pathsep + os.path.expanduser("~/.dotnet") + os.pathsep + os.path.expanduser("~/.dotnet/tools")
    env["DOTNET_ROOT"] = os.path.expanduser("~/.dotnet")
    
    parent = ROOT / "var/benchmark/fable-rust"
    parent.mkdir(parents=True, exist_ok=True)
    stamp = datetime.now(timezone.utc).strftime("%Y%m%dT%H%M%SZ-")
    directory = Path(tempfile.mkdtemp(prefix=stamp, dir=parent))
    print(f"Builds and results: {directory}", flush=True)
    
    main_dir = directory / "Main"
    
    # 1. We assume sharpurs has been run, grab its output
    sharpurs_out = ROOT / "run/bak/sharp/output/Main"
    if not sharpurs_out.exists():
        print("Running bin/sharp/run to generate F# code...")
        subprocess.check_call(["./bin/sharp/run", "--pure"], cwd=ROOT, env=env)
        
    shutil.copytree(sharpurs_out, main_dir)
    
    # 2. Patch F#
    patch_fsharp_code(main_dir)
    
    # 3. Run Fable
    print("Transpiling to Rust using Fable...")
    fable_out = directory / "fable_output"
    subprocess.check_call(["fable", "Program.fsproj", "--lang", "Rust", "-o", str(fable_out)], cwd=main_dir, env=env)
    
    # 4. Cargo Build
    print("Building Rust binary...")
    subprocess.check_call(["cargo", "build", "--release"], cwd=fable_out, env=env)
    
    binary = fable_out / "target" / "release" / "Main"
    if not binary.exists():
        binary = fable_out / "target" / "release" / "Program" # depends on project name
    
    # 5. Execute and Parse
    results, total = execute_and_parse(binary)
    print("Results:")
    for k, v in results.items():
        print(f"  {k}: {v} us")
    print(f"Total: {total} ms")
    
    # 6. Update README
    if args.update_readme:
        updated = render_readme(original_readme, results, total)
        (ROOT / "README.md").write_text(updated)
        print("README.md updated.")

if __name__ == "__main__":
    main()
