#!/usr/bin/env python3
"""Expose diagnostic protocols through the usual Go and JS runners."""
import argparse
from datetime import datetime, timezone
import json
import os
from pathlib import Path
import subprocess
import sys

ROOT = Path(__file__).resolve().parents[2]
DIAGNOSTICS = {"ArrayIndexing": "array-indexing.py", "JsonTypedAst": "json-diagnostic.py",
               "JsonDecoding": "json-diagnostic.py"}


def diagnostic_name(arguments):
    names = []
    for index, value in enumerate(arguments):
        if value == "--test" and index + 1 < len(arguments):
            names.append(arguments[index + 1].removeprefix("Test."))
        elif value.startswith("--test="):
            names.append(value.partition("=")[2].removeprefix("Test."))
    return next((name for name in names if name in DIAGNOSTICS), None)


def options(runtime, arguments):
    parser = argparse.ArgumentParser(
        description="Run ArrayIndexing, JsonTypedAst or JsonDecoding with its recorded diagnostic protocol.",
        allow_abbrev=False)
    parser.add_argument("--test", required=True)
    actions = parser.add_mutually_exclusive_group()
    actions.add_argument("--build-only", action="store_true")
    actions.add_argument("--run-only", action="store_true")
    parser.add_argument("--build-dir", type=Path, help="isolated build and campaign directory")
    parser.add_argument("--output", type=Path, help="new measurement campaign directory")
    args = parser.parse_args(arguments)
    args.test = args.test.removeprefix("Test.")
    if args.test not in DIAGNOSTICS:
        parser.error("choose exactly one diagnostic test")
    if args.build_only and args.output:
        parser.error("--output requires a measurement, not --build-only")
    selected = args.build_dir or ROOT / f"run/bak/{runtime}/modes/test-{args.test}"
    args.build_dir = selected.resolve()
    if selected.is_symlink() or args.build_dir == ROOT or args.build_dir in ROOT.parents:
        parser.error("unsafe build directory")
    if args.output:
        args.output = args.output.resolve()
        if args.output.exists():
            parser.error("--output must be a new campaign directory")
        if args.output == args.build_dir or args.output in args.build_dir.parents:
            parser.error("--output cannot contain the build directory")
    return args


def stamp():
    return datetime.now(timezone.utc).strftime("%Y%m%dT%H%M%S.%fZ") + f"-{os.getpid()}"


def run(command):
    subprocess.run([str(part) for part in command], check=True)


def json_workspace(base, runtime, run_only, suite='JsonTypedAst'):
    pointer = base / "latest-build.json"
    if not run_only:
        return base / "builds" / stamp()
    if not pointer.exists() and (base / "manifest.json").is_file():
        return base
    previous = json.loads(pointer.read_text())
    relative = Path(previous["workspace"])
    if (previous.get("runtime") != runtime or previous.get('suite', 'JsonTypedAst') != suite or len(relative.parts) != 2
            or relative.parts[0] != "builds" or relative.parts[1] in {".", ".."}):
        raise ValueError(f"Invalid JSON diagnostic build pointer: {pointer}")
    return base / relative


def execute(runtime, args):
    base = args.build_dir
    script = ROOT / "bin/benchmark" / DIAGNOSTICS[args.test]
    common = [sys.executable, script]
    output = args.output or base / "campaigns" / stamp()
    if args.test == "ArrayIndexing":
        common += ["--runtime", runtime, "--build-dir", base]
        if not args.run_only:
            run([*common, "--build-only"])
        if not args.build_only:
            run([*common, "--run-only", "--output", output])
    else:
        work = json_workspace(base, runtime, args.run_only, args.test)
        common += ["--runtime", runtime, "--suite", args.test, "--workspace", work]
        if not args.run_only:
            # This runner requires a new workspace; retain previous builds and
            # publish the pointer only after the replacement build succeeds.
            run([*common, "build"])
            pointer = base / "latest-build.json"
            temporary = pointer.with_suffix(".tmp")
            temporary.write_text(json.dumps({"runtime": runtime, "suite": args.test,
                "workspace": str(work.relative_to(base))}, indent=2) + "\n")
            temporary.replace(pointer)
        if not args.build_only:
            run([*common, "measure", "--output", output])
    if not args.build_only:
        print(f"Validated {runtime} {args.test} results: {output / 'results.json'}", flush=True)


def main():
    runtime, *arguments = sys.argv[1:]
    if runtime not in {"go", "js"}:
        raise ValueError("diagnostic dispatch supports Go and JS")
    if diagnostic_name(arguments) is None:
        ordinary = ROOT / ("bin/native/driver.py" if runtime == "go" else "bin/js/driver.py")
        os.execv(sys.executable, [sys.executable, str(ordinary), runtime, *arguments])
    execute(runtime, options(runtime, arguments))


if __name__ == "__main__":
    try:
        main()
    except (OSError, ValueError, KeyError, subprocess.CalledProcessError) as error:
        print(f"ERROR: {error}", file=sys.stderr)
        sys.exit(1)
