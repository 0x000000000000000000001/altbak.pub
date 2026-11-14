#!/usr/bin/env python3
"""Expose diagnostic protocols through the usual Go, JS and Rust runners."""
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
    parser.add_argument('-d', '--dev', action='store_true', help='Go: use the gopurs development profile')
    actions = parser.add_mutually_exclusive_group()
    actions.add_argument("--build-only", action="store_true")
    actions.add_argument("--run-only", action="store_true")
    parser.add_argument("--build-dir", type=Path, help="isolated build and campaign directory")
    parser.add_argument("--output", type=Path, help="new measurement campaign directory")
    args = parser.parse_args(arguments)
    if args.dev and runtime != 'go':
        parser.error('--dev is Go-only')
    args.test = args.test.removeprefix("Test.")
    if args.test not in DIAGNOSTICS:
        parser.error("choose exactly one diagnostic test")
    if args.build_only and args.output:
        parser.error("--output requires a measurement, not --build-only")
    suffix = '-dev' if args.dev else ''
    selected = args.build_dir or ROOT / f"run/bak/{runtime}/modes/test-{args.test}{suffix}"
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


def json_workspace(base, runtime, run_only, suite='JsonTypedAst', dev=False):
    pointer = base / "latest-build.json"
    if not run_only:
        return base / "builds" / stamp()
    if not pointer.exists() and (base / "manifest.json").is_file():
        return base
    previous = json.loads(pointer.read_text())
    relative = Path(previous["workspace"])
    if (previous.get("runtime") != runtime or previous.get('suite', 'JsonTypedAst') != suite
            or previous.get('go_backend_profile', 'release') != ('dev' if dev else 'release')
            or len(relative.parts) != 2
            or relative.parts[0] != "builds" or relative.parts[1] in {".", ".."}):
        raise ValueError(f"Invalid JSON diagnostic build pointer: {pointer}")
    return base / relative


def execute(runtime, args):
    if args.dev:
        print('[altbak/go] Build mode: DEV (--dev); generated Go packages use -N -l.', flush=True)
    base = args.build_dir
    script = ROOT / "bin/benchmark" / DIAGNOSTICS[args.test]
    common = [sys.executable, script]
    if args.dev:
        common.append('--dev')
    output = args.output or base / "campaigns" / stamp()
    if args.test == "ArrayIndexing":
        common += ["--runtime", runtime, "--build-dir", base]
        if not args.run_only:
            run([*common, "--build-only"])
        if not args.build_only:
            run([*common, "--run-only", "--output", output])
    else:
        work = json_workspace(base, runtime, args.run_only, args.test, args.dev)
        common += ["--runtime", runtime, "--suite", args.test, "--workspace", work]
        if not args.run_only:
            # This runner requires a new workspace; retain previous builds and
            # publish the pointer only after the replacement build succeeds.
            run([*common, "build"])
            pointer = base / "latest-build.json"
            temporary = pointer.with_suffix(".tmp")
            temporary.write_text(json.dumps({"runtime": runtime, "suite": args.test,
                "go_backend_profile": 'dev' if args.dev else 'release',
                "workspace": str(work.relative_to(base))}, indent=2) + "\n")
            temporary.replace(pointer)
        if not args.build_only:
            run([*common, "measure", "--output", output])
    if not args.build_only:
        print(f"Validated {runtime} {args.test} results: {output / 'results.json'}", flush=True)


def ordinary_driver(runtime):
    return ROOT / ("bin/native/driver.py" if runtime in {"go", "rust"} else "bin/js/driver.py")


def execute_with_wip(runtime, arguments):
    parser = argparse.ArgumentParser(
        description="Run the core suite plus all available excluded WIP benchmarks.",
        allow_abbrev=False)
    parser.add_argument("--wip-included", action="store_true", required=True)
    actions = parser.add_mutually_exclusive_group()
    actions.add_argument("--build-only", action="store_true")
    actions.add_argument("--run-only", action="store_true")
    parser.add_argument("-d", "--dev", action="store_true")
    parser.add_argument("-c", "--clean", action="store_true",
                        help="Rebuild the native backend compiler; diagnostic builds are already fresh")
    args = parser.parse_args(arguments)
    if args.dev and runtime != "go":
        parser.error("--dev is Go-only")
    if args.clean and args.run_only:
        parser.error("--clean cannot be combined with --run-only")
    profile_args = ["--dev"] if args.dev else []
    phases = (["--build-only"] if args.build_only else ["--run-only"] if args.run_only
              else ["--build-only", "--run-only"])
    for phase in phases:
        # Build the ordinary workspace first, so --clean refreshes the backend
        # compiler before any diagnostic artifact records its identity.
        native_args = ["--wip-included", phase, *profile_args]
        if args.clean and phase == "--build-only":
            native_args.append("--clean")
        print(f"=== {runtime}: core and library WIP cases / {phase[2:]} ===", flush=True)
        run([sys.executable, ordinary_driver(runtime), runtime, *native_args])
        for name in DIAGNOSTICS:
            print(f"=== {runtime}: {name} (excluded WIP) / {phase[2:]} ===", flush=True)
            execute(runtime, options(runtime, ["--test", name, phase, *profile_args]))


def main():
    runtime, *arguments = sys.argv[1:]
    if runtime not in {"go", "js", "rust"}:
        raise ValueError("diagnostic dispatch supports Go, JS and Rust")
    if "--wip-included" in arguments:
        execute_with_wip(runtime, arguments)
        return
    if diagnostic_name(arguments) is None:
        ordinary = ordinary_driver(runtime)
        os.execv(sys.executable, [sys.executable, str(ordinary), runtime, *arguments])
    execute(runtime, options(runtime, arguments))


if __name__ == "__main__":
    try:
        main()
    except (OSError, ValueError, KeyError, subprocess.CalledProcessError) as error:
        print(f"ERROR: {error}", file=sys.stderr)
        sys.exit(1)
