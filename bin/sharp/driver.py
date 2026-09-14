#!/usr/bin/env python3
"""Build and validate F# benchmarks without modifying the shared workspace."""
import argparse
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
HELPERS = runpy.run_path(str(ROOT / "bin/js/driver.py"))
VALIDATOR = runpy.run_path(str(ROOT / "bin/benchmark/validate.py"))


def digest(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def options():
    parser = argparse.ArgumentParser(description=__doc__)
    modes = parser.add_mutually_exclusive_group()
    modes.add_argument("--ffi", action="store_true")
    modes.add_argument("--fficc", action="store_true")
    modes.add_argument("--test")
    modes.add_argument("--x", action="store_true")
    actions = parser.add_mutually_exclusive_group()
    actions.add_argument("--build-only", action="store_true")
    actions.add_argument("--run-only", action="store_true")
    parser.add_argument("--artifact", type=Path)
    parser.add_argument("--expected")
    parser.add_argument("-c", "--clean", action="store_true", help="All builds already use a fresh directory")
    args = parser.parse_args()
    if args.clean and args.run_only:
        parser.error("--clean conflicts with --run-only")
    if args.artifact and not args.run_only:
        parser.error("--artifact requires --run-only")
    args.test = args.test.removeprefix("Test.") if args.test else None
    args.mode = "test" if args.test else "x" if args.x else "ffi" if args.ffi else "fficc" if args.fficc else "pure"
    VALIDATOR["expected_cases"](args.mode, args.test, args.expected)
    return args


def snapshot(work, args):
    HELPERS["snapshot_sources"](work, args)
    directories = ["src", "srx"] if args.mode == "x" else ["src"]
    for directory in directories:
        for source in (ROOT / directory).rglob("*"):
            if source.is_file() and source.suffix in {".fs", ".cs"}:
                target = work / "src" / source.relative_to(ROOT / directory)
                target.parent.mkdir(parents=True, exist_ok=True)
                shutil.copyfile(source, target)
    dependencies = work / ".spago/p"
    shutil.copytree(ROOT / "run/bak/sharp/spago/p", dependencies)
    config = (ROOT / "run/bak/sharp/spago.sharp.yaml").read_text()
    providers = re.findall(r'(?m)^    ([a-z0-9-]+):\n      path: "([^"\n]+)"', config)
    for name, relative in providers:
        original = (ROOT / relative).resolve()
        target = dependencies / ("local-" + name)
        shutil.copytree(original / "src", target / "src")
        config = config.replace('path: "' + relative + '"',
                                'path: "' + str(target.relative_to(work)) + '"')
    (work / "spago.yaml").write_text(config)


def build(work, args, env, dotnet):
    snapshot(work, args)
    entry = {"pure": "App", "ffi": "AppFFI", "fficc": "AppFFICheatcode",
             "test": "AppX", "x": "AppX"}[args.mode]
    commands = []
    sources = HELPERS["source_closure"](work, entry, "sharp")
    command = HELPERS["command"]
    command(["purs", "compile", "--codegen", "corefn", *sources], work, env, "purs", commands)
    typed = json.loads((work / "output" / entry / "corefn.json").read_text())
    if "dataDecls" not in typed or "classDecls" not in typed:
        raise RuntimeError("Sharpurs requires the fork's typed CoreFn; ordinary CoreFn was generated")
    backend = ROOT.parent / "sharpurs/sharpurs/bin/sharpurs"
    command([str(backend), "--main", entry], work, env, "sharpurs", commands)
    project = work / "output/Main/Program.fsproj"
    if not project.is_file():
        raise RuntimeError("Sharpurs did not generate Program.fsproj")
    command([dotnet, "build", str(project), "-c", "Release", "-p:Optimize=true",
             "-p:NuGetAudit=false", "--ignore-failed-sources", "-o", str(work / "binary")],
            work, env, "dotnet", commands)
    program = work / "binary/Program.dll"
    if not program.is_file():
        raise RuntimeError("No compiled Program.dll")
    manifest = {"mode": args.mode, "test": args.test, "expected": args.expected,
                "entry": entry, "commands": commands, "profile": "Release; Optimize=true",
                "dotnet": subprocess.check_output([dotnet, "--version"], text=True).strip(),
                "backend_sha256": digest(backend.with_suffix(".js")),
                "source_sha256": {str(path.relative_to(work)): digest(path)
                                  for directory in [work / "src", work / ".spago", work / "output/Main"]
                                  for path in directory.rglob("*")
                                  if path.is_file() and path.suffix in {".purs", ".fs", ".cs", ".fsproj", ".csproj"}},
                "binary_sha256": {str(path.relative_to(work)): digest(path)
                                  for path in (work / "binary").rglob("*") if path.is_file()}}
    (work / "manifest.json").write_text(json.dumps(manifest, indent=2) + "\n")


def execute(work, args, env, dotnet):
    manifest = json.loads((work / "manifest.json").read_text())
    for key in ["mode", "test", "expected"]:
        if manifest[key] != getattr(args, key):
            raise RuntimeError(f"Artifact {key} differs from the requested run")
    if manifest["dotnet"] != subprocess.check_output([dotnet, "--version"], text=True).strip():
        raise RuntimeError("dotnet version changed since build")
    for path, expected in manifest["binary_sha256"].items():
        if digest(work / path) != expected:
            raise RuntimeError(f"Compiled artifact changed: {path}")
    result = subprocess.run([dotnet, str(work / "binary/Program.dll")], cwd=work,
                            env=env, text=True, capture_output=True)
    (work / "run.stdout").write_text(result.stdout)
    (work / "run.stderr").write_text(result.stderr)
    print(result.stdout, end="", flush=True)
    if result.stderr:
        print(result.stderr, end="", file=sys.stderr)
    if result.returncode:
        raise RuntimeError(f"F# benchmark exited {result.returncode}")
    checked = VALIDATOR["validate_output"](result.stdout, args.mode, args.test, args.expected)
    (work / "validated.json").write_text(json.dumps(checked, indent=2) + "\n")
    print(f"Validated output: {work / 'validated.json'}", file=sys.stderr)


def main():
    args = options()
    env = os.environ.copy()
    env["PATH"] = str(ROOT / "run/bak/js/node_modules/.bin") + os.pathsep + env.get("PATH", "")
    env["DOTNET_CLI_TELEMETRY_OPTOUT"] = "1"
    env["DOTNET_SKIP_FIRST_TIME_EXPERIENCE"] = "1"
    dotnet = env.get("DOTNET") or shutil.which("dotnet") or str(Path.home() / ".dotnet/dotnet")
    store = ROOT / "run/bak/sharp/output-runs"
    store.mkdir(parents=True, exist_ok=True)
    pointer = store / ("latest-" + args.mode + ("-" + args.test if args.test else "") + ".json")
    if args.run_only:
        work = args.artifact.resolve() if args.artifact else Path(json.loads(pointer.read_text())["artifact"])
    else:
        work = Path(tempfile.mkdtemp(prefix=args.mode + "-", dir=store))
        print(f"Artifact: {work}", flush=True)
        build(work, args, env, dotnet)
        pointer.write_text(json.dumps({"artifact": str(work)}) + "\n")
    if not args.build_only:
        execute(work, args, env, dotnet)


if __name__ == "__main__":
    try:
        main()
    except (OSError, ValueError, RuntimeError, subprocess.SubprocessError) as error:
        print(f"error: {error}", file=sys.stderr)
        raise SystemExit(1)
