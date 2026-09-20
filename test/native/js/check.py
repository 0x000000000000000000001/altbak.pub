#!/usr/bin/env python3
"""Independent result oracles for the 28 JS, Chez, or Erlang foreign kernels.

No backend build, shared output, or timing is involved. Each native process uses
a temporary directory. Include each published input and small discriminating
inputs so that hardcoded answers and accidentally equal formulas fail.
"""
import argparse
import json
from pathlib import Path
import re
import shutil
import subprocess
import tempfile

ROOT = Path(__file__).resolve().parents[3]


# Shared independently calculated cases, also used by the other native backends.
# run_path avoids creating bytecode alongside the shared oracle.
import runpy
CASES = {}
for name, argument, expected in runpy.run_path(str(ROOT / "test/native/oracle.py"))["cases"]():
    CASES.setdefault(name, []).append((argument, expected))


def run(command, cwd):
    return subprocess.run(command, cwd=cwd, text=True, capture_output=True, timeout=90)


def check(language):
    failures = []
    count = 0
    with tempfile.TemporaryDirectory(prefix="altbak-native-") as directory:
        tmp = Path(directory)
        for name, cases in CASES.items():
            for suffix in ["FFI", "FFICheatcode"]:
                stem = name + suffix
                function = "run" + stem
                if language == "js":
                    shutil.copyfile(ROOT / "src/Test" / (stem + ".js"), tmp / "kernel.mjs")
                    if suffix == "FFI":
                        probes = {
                            "ListOps": '''
if (foldl(a => x => a + x.length)(0)({type:"Cons", value0:"ab", value1:{type:"Cons", value0:"c", value1:{type:"Nil"}}}) !== 3) throw Error("generic list fold");
if (foldl(a => x => a + x)("")(filterEvens(range(1)(6))) !== "642") throw Error("filter order");''',
                            "Primes": '''
const input = {type:"Cons",value0:"a",value1:{type:"Cons",value0:"bb",value1:{type:"Cons",value0:"c",value1:{type:"Nil"}}}};
const selected = filter(x => x.length === 1)(input);
if (selected.value0 !== "a" || selected.value1.value0 !== "c" || input.value1.value0 !== "bb") throw Error("generic persistent filter");''',
                            "Polymorphism": '''
if (polyLoop({mempty_:"x",mappend_:a=>b=>a+b},3,"!") !== "!xxx") throw Error("generic dictionary");''',
                            "Church": '''
const two = succC(succC(zeroC));
if (mulC(two)(two)(s=>s+"x")("!") !== "!xxxx") throw Error("generic Church");''',
                            "StateMonad": '''
const action = bindState(get)(s => bindState(put(s+"!"))(_ => pureState(s.length)));
const result = runState(action)("ab");
if (result.val !== 2 || result.state !== "ab!") throw Error("generic State");''',
                            "RowToList": '''
if (keys(dictNil,{}) !== 0 || keys(dictCons(dictCons(dictNil)),{name:"value",flag:true}) !== 2) throw Error("recursive row dictionary");''',
                        }
                        with (tmp / "kernel.mjs").open("a") as module:
                            module.write("\n" + probes.get(name, "") + "\n")
                    (tmp / "check.mjs").write_text(
                        'import * as kernel from "./kernel.mjs";\n'
                        f'for (const [n, expected] of {json.dumps(cases)}) {{\n'
                        f' const actual = kernel.{function}(n);\n'
                        ' if (actual !== expected) throw new Error(`${n}: ${actual} != ${expected}`);\n'
                        '}\n')
                    result = run(["node", "check.mjs"], tmp)
                elif language == "scheme":
                    source = ROOT / "src/Test" / (stem + ".ss")
                    assertions = "\n".join(
                        f'(unless (= ({function} {n}) {expected}) '
                        f'(error "{stem}" "wrong result" {n} ({function} {n}) {expected}))'
                        for n, expected in cases)
                    (tmp / "check.ss").write_text(
                        f'(load {json.dumps(str(source))})\n'
                        f'(import (chezscheme) (Test.{stem} foreign))\n{assertions}\n')
                    result = run(["chez", "--optimize-level", "3", "--script", "check.ss"], tmp)
                else:
                    source = (ROOT / "src/Test" / (stem + ".erl")).read_text()
                    module = re.search(r"-module\(([^)]+)\)", source)[1]
                    (tmp / (module + ".erl")).write_text(source)
                    result = run(["erlc", "+warnings_as_errors", module + ".erl"], tmp)
                    if result.returncode == 0:
                        assertions = ",".join(
                            f"{expected} = {module}:{function}({n})" for n, expected in cases)
                        result = run(["erl", "+S", "1:1", "-noshell", "-pa", str(tmp),
                                      "-eval", "begin " + assertions + " end,halt()."], tmp)
                if result.returncode:
                    detail = (result.stderr or result.stdout).strip().splitlines()
                    failures.append(f"{stem}: " + " / ".join(detail[:8]))
                else:
                    count += len(cases)
        print(f"{language}: {count} assertions passed, {len(failures)} kernels failed", flush=True)
        for failure in failures:
            print(failure, flush=True)
    return not failures


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("language", choices=["js", "scheme", "erlang"])
    raise SystemExit(0 if check(parser.parse_args().language) else 1)
