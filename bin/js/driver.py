#!/usr/bin/env python3
"""Build JS, ES, Chez and Erlang in independent workspaces, then validate output."""
import argparse
import ctypes
import ctypes.util
import hashlib
import json
import os
from pathlib import Path
import re
import runpy
import shutil
import subprocess
import sys
import tempfile

ROOT = Path(__file__).resolve().parents[2]


def options():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("backend", choices=["js", "es", "scm", "erl"])
    modes = parser.add_mutually_exclusive_group()
    modes.add_argument("--ffi", action="store_true")
    modes.add_argument("--fficc", action="store_true")
    modes.add_argument("--test", metavar="MODULE")
    modes.add_argument("--x", action="store_true")
    actions = parser.add_mutually_exclusive_group()
    actions.add_argument("--build-only", action="store_true")
    actions.add_argument("--run-only", action="store_true")
    parser.add_argument("--artifact", type=Path, help="Existing isolated build directory (run-only)")
    parser.add_argument("--expected", help="Expected output for a custom --test")
    parser.add_argument("-c", "--clean", action="store_true",
                        help="Build afresh (all builds already use a fresh isolated directory)")
    args = parser.parse_args()
    if args.artifact and not args.run_only:
        parser.error("--artifact requires --run-only")
    if args.clean and args.run_only:
        parser.error("--clean cannot be combined with --run-only")
    if args.expected is not None and args.test is None:
        parser.error("--expected requires --test")
    if args.test and not re.fullmatch(r"[A-Z][A-Za-z0-9_]*(?:\.[A-Z][A-Za-z0-9_]*)*", args.test):
        parser.error("--test expects a PureScript module name")
    args.mode = "test" if args.test else "x" if args.x else "ffi" if args.ffi else "fficc" if args.fficc else "pure"
    args.test = args.test.removeprefix("Test.") if args.test else None
    return args


def environment(backend):
    env = os.environ.copy()
    paths = [ROOT / "bin" / backend]
    if backend in ["scm", "erl"]:
        paths.append(ROOT / "bin/purs-vanilla/node_modules/.bin")
    paths.append(ROOT / "run/bak/js/node_modules/.bin")
    env["PATH"] = os.pathsep.join(map(str, paths)) + os.pathsep + env.get("PATH", "")
    if backend == "scm" and sys.platform == "darwin":
        libraries = ["/opt/homebrew/lib", "/opt/homebrew/opt/icu4c/lib", "/opt/homebrew/opt/pcre2/lib"]
        env["DYLD_LIBRARY_PATH"] = os.pathsep.join(libraries + [env.get("DYLD_LIBRARY_PATH", "")])
    return env


def command(argv, work, env, name, commands):
    commands.append(argv)
    display = argv if len(argv) <= 12 else argv[:4] + [f"… ({len(argv) - 4} source arguments)"]
    print(f"[{name}] {' '.join(display)}", flush=True)
    with (work / (name + ".log")).open("w") as log:
        result = subprocess.run(argv, cwd=work, env=env, stdout=log, stderr=subprocess.STDOUT)
    if result.returncode:
        print((work / (name + ".log")).read_text()[-16000:], file=sys.stderr)
        raise RuntimeError(f"{name} failed ({result.returncode}); log: {work / (name + '.log')}")


def snapshot_sources(work, args):
    suffixes = {".purs", ".js", ".ss", ".erl"}
    for source in (ROOT / "src").rglob("*"):
        if source.is_file() and source.suffix in suffixes:
            relative = source.relative_to(ROOT / "src")
            # These diagnostics use independent workspaces and optional dependencies.
            if any(relative.parts[:2] == ('Test', diagnostic)
                   or relative.as_posix() in {f'Test/{diagnostic}.{ext}' for ext in ['purs', 'go', 'js']}
                   for diagnostic in ['JsonTypedAst', 'JsonDecoding']):
                continue
            target = work / source.relative_to(ROOT)
            target.parent.mkdir(parents=True, exist_ok=True)
            shutil.copyfile(source, target)
    if args.mode == "test":
        # Only the isolated copy is replaced; a user's src/AppX.purs is untouched.
        (work / "src/AppX.purs").write_text(
            "module AppX where\nimport Prelude\nimport Effect (Effect)\n"
            "import Bench (runBench)\n"
            f"import Test.{args.test} as Selected\n"
            "main :: Effect Unit\nmain = void $ runBench Selected.describe Selected.act\n")
    elif args.mode == "x":
        for source in (ROOT / "srx").rglob("*"):
            if source.is_file() and source.suffix in suffixes:
                target = work / "src" / source.relative_to(ROOT / "srx")
                target.parent.mkdir(parents=True, exist_ok=True)
                shutil.copyfile(source, target)


def source_closure(work, entry, backend):
    modules = {}
    for directory in [work / ".spago/p", work / "src"]:
        for path in sorted(directory.rglob("*.purs")):
            if directory.name != "src" and "src" not in path.relative_to(directory).parts:
                continue
            source = path.read_text()
            declaration = re.search(r"(?m)^module\s+([A-Z][\w.]*)\b", source)
            if not declaration:
                continue
            name = declaration[1]
            if name in modules and modules[name].read_text() != source:
                raise RuntimeError(f"Ambiguous source module {name}: {modules[name]}, {path}")
            modules[name] = path
    selected = {}
    # purerl inserts runFn3 calls for the compiler's lazy dictionary bindings.
    pending = [entry] + (["Data.Function.Uncurried"] if backend == "erl" else [])
    while pending:
        name = pending.pop()
        if name in selected or name == "Prim" or name.startswith("Prim."):
            continue
        if name not in modules:
            raise RuntimeError(f"Missing dependency {name} in the isolated package snapshot")
        path = modules[name]
        selected[name] = path
        pending.extend(re.findall(r"(?m)^import\s+([A-Z][\w.]*)\b", path.read_text()))
    return [str(path.relative_to(work)) for _, path in sorted(selected.items())]


def scheme_tool(work):
    """Copy the runtime and bind its two ICU calls to the installed ICU ABI."""
    package = work / "toolchain/purescm"
    shutil.copytree(ROOT / "run/bak/js/node_modules/purescm", package)
    changes = []
    candidates = [Path("/opt/homebrew/opt/icu4c/lib/libicuuc.dylib")]
    library = next((str(path) for path in candidates if path.is_file()), None)
    library = library or ctypes.util.find_library("icuuc")
    if not library:
        raise RuntimeError("Chez runtime requires ICU; no libicuuc found")
    icu = ctypes.CDLL(library)
    if not hasattr(icu, "u_strToUpper_74") or not hasattr(icu, "u_strToLower_74"):
        version = next((n for n in range(100, 39, -1)
                        if hasattr(icu, f"u_strToUpper_{n}") and hasattr(icu, f"u_strToLower_{n}")), None)
        if version is None:
            raise RuntimeError(f"No matching ICU string case symbols found in {library}")
        source = package / "lib/purescm/pstring.ss"
        content = source.read_text()
        for name in ["u_strToUpper", "u_strToLower"]:
            old, new = name + "_74", name + "_" + str(version)
            if content.count('"' + old + '"') != 1:
                raise RuntimeError("Unexpected purescm ICU binding; refusing an unverified runtime patch")
            content = content.replace('"' + old + '"', '"' + new + '"')
        source.write_text(content)
        changes.append(f"Private purescm ICU bindings _74 -> _{version} ({library})")
    return package, changes


def build(args, store, pointer):
    work = Path(tempfile.mkdtemp(prefix=args.mode + "-", dir=store))
    print(f"Build directory: {work}", flush=True)
    snapshot_sources(work, args)
    cache_backend = "js" if args.backend == "es" else args.backend
    base = ROOT / "run/bak" / cache_backend
    shutil.copyfile(base / f"spago.{cache_backend}.yaml", work / "spago.yaml")
    if (base / "spago.lock").is_file():
        shutil.copyfile(base / "spago.lock", work / "spago.lock")
    # Dependencies are private copies; no mutation of root links or backend caches.
    if (base / "spago").is_dir():
        shutil.copytree(base / "spago", work / ".spago", symlinks=False)
    env = environment(args.backend)
    commands = []
    entry = {"pure": "App", "ffi": "AppFFI", "fficc": "AppFFICheatcode",
             "test": "AppX", "x": "AppX"}[args.mode]
    compatibility = []
    if args.backend == "erl":
        # Actual identity foreigns required by this old purerl package set.
        # Install only in this build's private dependency snapshot.
        for name in ["Reflectable", "Symbol"]:
            for source in (work / ".spago").rglob(name + ".purs"):
                if source.parent.name != "Data":
                    continue
                foreign = source.with_suffix(".erl")
                if not foreign.exists():
                    imports = re.findall(r"(?m)^foreign import (\w+)\s*::", source.read_text())
                    if imports != ["unsafeCoerce"]:
                        raise RuntimeError(f"Unexpected foreign interface in {source}")
                    module = "data_" + name.lower() + "@foreign"
                    foreign.write_text(f"-module({module}).\n-export([unsafeCoerce/1]).\nunsafeCoerce(X) -> X.\n")
                    compatibility.append(str(foreign.relative_to(work)))
        if compatibility:
            print("Identity compatibility foreigns: " + ", ".join(compatibility), flush=True)
    # Compile the reachable modules from the copied packages. This requires no
    # writable global Spago database, network, or fake JavaScript foreigns.
    sources = source_closure(work, entry, args.backend)
    codegen = "corefn" if args.backend in ["scm", "erl"] else "js,corefn"
    # Arista consumes standard CoreFn, unlike the local typed CoreFn backends.
    compiler = str(ROOT / "run/bak/js/node_modules/purescript/purs.bin") if args.backend == "es" else "purs"
    build_command = [compiler, "compile", "--codegen", codegen, *sources]
    command(build_command, work, env, "purs", commands)
    if args.backend == "es":
        command(["purs-backend-es", "build", "--output-dir", "es-output"], work, env, "es", commands)
        for foreign in (work / "output").glob("*/foreign.js"):
            target = work / "es-output" / foreign.parent.name / "foreign.js"
            target.parent.mkdir(parents=True, exist_ok=True)
            shutil.copyfile(foreign, target)
    if args.backend in ["js", "es"]:
        output = "es-output" if args.backend == "es" else "output"
        run = ["node", "--input-type=module", "-e",
               f"import {{ main }} from './{output}/{entry}/index.js'; main();"]
    elif args.backend == "scm":
        package, changes = scheme_tool(work)
        compatibility.extend(changes)
        for change in changes:
            print(change, flush=True)
        tool = ["node", str(package / "index.js")]
        command([*tool, "build"], work, env, "scheme", commands)
        command([*tool, "bundle-app", "--main", entry, "--output", "bundle"],
                work, env, "bundle", commands)
        run = ["chez", "--optimize-level", "3", "--program", "bundle/main"]
    else:
        command(["purerl"], work, env, "purerl", commands)
        beams = work / "beams"
        beams.mkdir()
        files = sorted(str(path.relative_to(work)) for path in (work / "output").glob("*/*.erl"))
        if not files:
            raise RuntimeError("purerl produced no Erlang source")
        command(["erlc", "-o", "beams", *files], work, env, "erlc", commands)
        module = entry[0].lower() + entry[1:] + "@ps"
        run = ["erl", "-noshell", "-pa", "beams", "-eval", f"({module}:main())(),halt()."]
    hashes = {str(path.relative_to(work)): hashlib.sha256(path.read_bytes()).hexdigest()
              for directory in [work / "src", work / ".spago", work / "toolchain"]
              for path in directory.rglob("*")
              if path.is_file() and path.suffix in {".purs", ".js", ".ss", ".erl"}}
    manifest = {"backend": args.backend, "mode": args.mode, "test": args.test,
                "expected": args.expected, "entry": entry, "run": run,
                "commands": commands, "source_sha256": hashes, "compatibility": compatibility}
    (work / "manifest.json").write_text(json.dumps(manifest, indent=2) + "\n")
    pointer.write_text(json.dumps({"artifact": str(work)}) + "\n")
    print(f"Artifact: {work}\nCommand: {' '.join(run)}", flush=True)
    return work


def execute(args, work):
    manifest = json.loads((work / "manifest.json").read_text())
    for key in ["backend", "mode", "test", "expected"]:
        if manifest[key] != getattr(args, key):
            raise RuntimeError(f"Artifact {key}={manifest[key]!r} does not match requested {getattr(args, key)!r}")
    if args.mode == "x":
        (work / "var").mkdir(exist_ok=True)
    result = subprocess.run(manifest["run"], cwd=work, env=environment(args.backend),
                            text=True, capture_output=True)
    (work / "run.stdout").write_text(result.stdout)
    (work / "run.stderr").write_text(result.stderr)
    print(result.stdout, end="", flush=True)
    if result.stderr:
        print(result.stderr, end="", file=sys.stderr)
    if result.returncode:
        raise RuntimeError(f"Native program failed ({result.returncode})")
    validator = runpy.run_path(str(ROOT / "bin/benchmark/validate.py"))
    checked = validator["validate_output"](result.stdout, args.mode, args.test, args.expected)
    (work / "validated.json").write_text(json.dumps(checked, indent=2) + "\n")
    print(f"Validated output: {work / 'validated.json'}", file=sys.stderr)


def main():
    args = options()
    store = ROOT / "run/bak" / args.backend / "output-runs"
    store.mkdir(parents=True, exist_ok=True)
    key = args.mode + ("-" + args.test if args.test else "")
    pointer = store / ("latest-" + key + ".json")
    work = ((args.artifact.resolve() if args.artifact else Path(json.loads(pointer.read_text())["artifact"]))
            if args.run_only else build(args, store, pointer))
    if not args.build_only:
        execute(args, work)


if __name__ == "__main__":
    try:
        main()
    except (OSError, ValueError, RuntimeError, subprocess.SubprocessError) as error:
        print(f"error: {error}", file=sys.stderr)
        raise SystemExit(1)
