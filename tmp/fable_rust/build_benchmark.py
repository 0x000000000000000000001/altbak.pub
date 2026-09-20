#!/usr/bin/env python3
"""Build the 14 sharpurs/Fable Rust benchmarks without running them.

Keep this script beside extract_fable_kernels.py, adapt_fsharp.py, runtime.fs,
bench_fable.rs, and fable-rust-benchmark.patch. Supply an empty build directory;
all generated files, compiler changes, logs, and caches stay inside it.
"""
from __future__ import annotations

import argparse
from datetime import datetime, timezone
import hashlib
import json
import os
from pathlib import Path
import re
import shutil
import subprocess
import sys
import time
import xml.etree.ElementTree as ET

from adapt_fsharp import chunks, propagate_reflection_exceptions


ASSETS = (
    "extract_fable_kernels.py", "adapt_fsharp.py", "runtime.fs",
    "bench_fable.rs", "fable-rust-benchmark.patch",
)
REFERENCES = (
    "Fable.AST", "Fable.Transforms", "Rust.AST", "FSharp.Core",
    "FSharp.Compiler.Service", "Fable.Core",
)
NUGET_CONFIG = "<configuration><packageSources><clear/></packageSources></configuration>\n"


def sha256(path: Path) -> str:
    digest = hashlib.sha256()
    with path.open("rb") as stream:
        for block in iter(lambda: stream.read(1024 * 1024), b""):
            digest.update(block)
    return digest.hexdigest()


def tree_hashes(root: Path, suffixes: set[str] | None = None) -> dict[str, str]:
    return {
        path.relative_to(root).as_posix(): sha256(path)
        for path in sorted(root.rglob("*"))
        if path.is_file() and (suffixes is None or path.suffix in suffixes)
    }


def write_project(path: Path, assembly: str, sources: list[str], references: Path) -> None:
    project = ET.Element("Project", Sdk="Microsoft.NET.Sdk")
    properties = ET.SubElement(project, "PropertyGroup")
    for name, value in {
        "TargetFramework": "net10.0",
        "AssemblyName": assembly,
        "DisableImplicitFSharpCoreReference": "true",
        "EnableDefaultCompileItems": "false",
        "OtherFlags": "--nowarn:3536",
    }.items():
        ET.SubElement(properties, name).text = value
    files = ET.SubElement(project, "ItemGroup")
    for source in sources:
        ET.SubElement(files, "Compile", Include=source)
    refs = ET.SubElement(project, "ItemGroup")
    for name in REFERENCES:
        ref = ET.SubElement(refs, "Reference", Include=name)
        ET.SubElement(ref, "HintPath").text = str(references / f"{name}.dll")
    ET.indent(project)
    ET.ElementTree(project).write(path, encoding="unicode")


def set_explicit_fsharp_core(project_path: Path, core: Path) -> None:
    tree = ET.parse(project_path)
    project = tree.getroot()
    properties = project.find("PropertyGroup")
    if properties is None:
        raise ValueError(f"Missing PropertyGroup: {project_path}")
    ET.SubElement(properties, "DisableImplicitFSharpCoreReference").text = "true"
    group = ET.SubElement(project, "ItemGroup")
    reference = ET.SubElement(group, "Reference", Include="FSharp.Core")
    ET.SubElement(reference, "HintPath").text = str(core)
    ET.indent(tree)
    tree.write(project_path, encoding="unicode")


class Build:
    def __init__(self, directory: Path, environment: dict[str, str], manifest: dict):
        self.directory = directory
        self.environment = environment
        self.manifest = manifest
        self.logs = directory / "logs"
        self.logs.mkdir()
        self.save()

    def save(self) -> None:
        (self.directory / "manifest.json").write_text(
            json.dumps(self.manifest, indent=2, sort_keys=True) + "\n"
        )

    def run(self, phase: str, command: list[str | Path], cwd: Path, timeout: int) -> None:
        argv = [str(value) for value in command]
        entry = {
            "phase": phase, "argv": argv, "cwd": str(cwd),
            "timeout_seconds": timeout, "log": str(self.logs / f"{phase}.log"),
        }
        self.manifest["commands"].append(entry)
        self.save()
        print(f"[{phase}]", flush=True)
        started = time.monotonic()
        try:
            with Path(entry["log"]).open("wb") as log:
                result = subprocess.run(
                    argv, cwd=cwd, env=self.environment, stdout=log,
                    stderr=subprocess.STDOUT, timeout=timeout, check=False,
                )
            entry["returncode"] = result.returncode
            if result.returncode:
                raise RuntimeError(f"{phase} failed ({result.returncode}); see {entry['log']}")
        except subprocess.TimeoutExpired as exc:
            entry["timed_out"] = True
            raise RuntimeError(f"{phase} exceeded {timeout}s; see {entry['log']}") from exc
        finally:
            entry["elapsed_seconds"] = round(time.monotonic() - started, 3)
            self.save()


def adapt_compatibility_sources(directory: Path, runtime: Path) -> list[dict]:
    changes = []
    target_runtime = directory / "Sharpurs_Prelude.fs"
    # Keep the generated Euclidean Int implementation when the typed sharpurs
    # path needs it. The compatibility overlay changes function dispatch, not
    # integer arithmetic; this declaration is copied byte-for-byte.
    retained = [(name, body) for name, body, _ in chunks(target_runtime.read_text())
                if name == "sharpurs_int_mod"]
    if len(retained) > 1:
        raise ValueError("Duplicate generated sharpurs_int_mod runtime declarations")
    shutil.copy2(runtime, target_runtime)
    changes.append({"file": target_runtime.name, "action": "overlay_runtime", "sha256": sha256(runtime)})
    for name, body in retained:
        with target_runtime.open("a") as output:
            output.write("\n" + body)
        changes.append({"file": target_runtime.name, "action": "retain_generated_runtime_helper",
                        "helper": name, "sha256": hashlib.sha256(body.encode()).hexdigest(),
                        "proof": "Generated declaration copied verbatim from the extracted source runtime."})
    array = directory / "Data.Array.fs"
    before = array.read_text()
    after, count = re.subn(r"^[ \t]*open System\.Linq[ \t]*(?:\n|$)", "", before, flags=re.M)
    array.write_text(after)
    changes.append({"file": array.name, "action": "remove_unused_open_System_Linq", "count": count})
    tree = directory / "Test.RBTree.fs"
    before = tree.read_text()
    after, adapters = propagate_reflection_exceptions(before)
    tree.write_text(after)
    changes.append({"file": tree.name, "action": "propagate_original_exception",
                    "count": len(adapters), "sha256_before": hashlib.sha256(before.encode()).hexdigest(),
                    "sha256_after": sha256(tree), "adapters": adapters})
    return changes


def perform_build(args: argparse.Namespace, assets: Path, build: Build) -> None:
    root = build.directory
    compiler = root / "compiler"
    compiler.mkdir()
    package = compiler / "package"
    source = args.fable_source / "src" / "Fable.Transforms" / "Rust"
    for name in ("RustPrinter.fs", "Fable2Rust.fs"):
        shutil.copy2(source / name, compiler / name)
    shutil.copytree(args.fable_package / "tools" / "net10.0" / "any", package / "tools" / "net10.0" / "any")
    shutil.copytree(args.fable_package / "fable-library-rust", package / "fable-library-rust")
    tools = package / "tools" / "net10.0" / "any"
    (compiler / "NuGet.Config").write_text(NUGET_CONFIG)
    build.run("01-patch-compiler", ["patch", "--batch", "--forward", "-p4", "-i", assets / "fable-rust-benchmark.patch"], compiler, 30)
    project = compiler / "Fable.Transforms.Rust.fsproj"
    write_project(project, "Fable.Transforms.Rust", ["RustPrinter.fs", "Fable2Rust.fs"], tools)
    build.run("02-build-compiler", [args.dotnet, "build", project, "--nologo", "--configuration", "Release", "--output", compiler / "bin", "--configfile", compiler / "NuGet.Config"], compiler, 180)
    shutil.copy2(compiler / "bin" / "Fable.Transforms.Rust.dll", tools / "Fable.Transforms.Rust.dll")
    raw, fsharp, rust = root / "raw", root / "fsharp", root / "rust"
    build.run("03-extract-kernels", [sys.executable, assets / "extract_fable_kernels.py", args.source, raw, "--target-framework", "net10.0"], root, 60)
    build.run("04-adapt-fsharp", [sys.executable, assets / "adapt_fsharp.py", raw, fsharp], root, 60)
    build.manifest["compatibility_changes"] = adapt_compatibility_sources(fsharp, assets / "runtime.fs")
    (fsharp / "NuGet.Config").write_text(NUGET_CONFIG)
    set_explicit_fsharp_core(fsharp / "Kernels.fsproj", tools / "FSharp.Core.dll")
    build.manifest["adapted_fsharp_sha256"] = tree_hashes(fsharp, {".fs", ".fsproj"})
    build.save()
    build.run("05-fable-rust", [args.dotnet, tools / "fable.dll", fsharp / "Kernels.fsproj", "--noCache", "--lang", "Rust", "--outDir", rust], root, 180)
    harness = (assets / "bench_fable.rs").read_text()
    harness, replacements = re.subn(r"\bBenchFable\b", "BenchFableKernels", harness)
    if replacements == 0 and "BenchFableKernels" not in harness:
        raise ValueError("Harness does not reference BenchFable or BenchFableKernels")
    generated = (rust / "BenchFableKernels.rs").read_text()
    allocator = "#[global_allocator]\nstatic FABLE_BENCH_ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;\n"
    (rust / "main.rs").write_text(generated.rstrip() + "\n\n" + allocator + "\n" + harness)
    (root / "Cargo.toml").write_text('''[package]
name = "altbak-fable"
version = "0.1.0"
edition = "2024"

[[bin]]
name = "altbak-fable"
path = "rust/main.rs"

[dependencies]
fable_library_rust = { path = "rust/fable_modules/fable-library-rust", default-features = false }
mimalloc = "0.1.32"

[profile.release]
opt-level = 3
debug = false
''')
    build.run("06-cargo-release", ["cargo", "build", "--offline", "--release", "--manifest-path", root / "Cargo.toml"], root, 600)
    executable = root / "target" / "release" / ("altbak-fable.exe" if os.name == "nt" else "altbak-fable")
    build.manifest.update({
        "status": "built", "benchmark_executed": False,
        "executable": str(executable), "executable_sha256": sha256(executable),
        "cargo_lock_sha256": sha256(root / "Cargo.lock"),
        "patched_compiler_sha256": sha256(compiler / "Fable2Rust.fs"),
        "patched_assembly_sha256": sha256(tools / "Fable.Transforms.Rust.dll"),
        "generated_rust_sha256": tree_hashes(rust, {".rs", ".toml"}),
    })
    build.save()
    print(f"Built {executable}\nManifest: {root / 'manifest.json'}\nBenchmarks were not run.")


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--source", required=True, type=Path, help="sharpurs Main_Fsharp directory containing Program.fsproj")
    parser.add_argument("--build-dir", required=True, type=Path, help="New or empty output directory")
    parser.add_argument("--dotnet", required=True, type=Path, help=".NET 10 dotnet executable")
    parser.add_argument("--fable-package", required=True, type=Path, help="Fable NuGet package directory, e.g. ~/.nuget/packages/fable/5.17.2")
    parser.add_argument("--fable-source", required=True, type=Path, help="Matching Fable source checkout")
    args = parser.parse_args()
    for name in ("source", "build_dir", "dotnet", "fable_package", "fable_source"):
        setattr(args, name, getattr(args, name).expanduser().resolve())
    assets = Path(__file__).resolve().parent
    required = [assets / name for name in ASSETS] + [args.dotnet, args.source / "Program.fsproj"]
    compiler_sources = args.fable_source / "src" / "Fable.Transforms" / "Rust"
    required += [compiler_sources / name for name in ("RustPrinter.fs", "Fable2Rust.fs")]
    package_tools = args.fable_package / "tools" / "net10.0" / "any"
    required += [package_tools / f"{name}.dll" for name in REFERENCES]
    required += [package_tools / "fable.dll", args.fable_package / "fable-library-rust" / "Cargo.toml"]
    for path in required:
        if not path.is_file():
            parser.error(f"Required file not found: {path}")
    for command in ("cargo", "rustc", "patch"):
        if shutil.which(command) is None:
            parser.error(f"Executable not found on PATH: {command}")
    if args.build_dir.exists() and (not args.build_dir.is_dir() or any(args.build_dir.iterdir())):
        parser.error("--build-dir must be new or empty; existing build outputs are never overwritten")
    for protected in (args.source, args.fable_source, args.fable_package, assets):
        if args.build_dir == protected or protected in args.build_dir.parents:
            parser.error(f"Build directory must be outside {protected}")
    args.build_dir.mkdir(parents=True, exist_ok=True)
    environment = os.environ.copy()
    environment.update({
        "DOTNET_CLI_HOME": str(args.build_dir / "dotnet-home"),
        "DOTNET_ROOT": str(args.dotnet.parent),
        "DOTNET_SKIP_FIRST_TIME_EXPERIENCE": "1", "DOTNET_CLI_TELEMETRY_OPTOUT": "1",
        "DOTNET_NOLOGO": "1", "DOTNET_GENERATE_ASPNET_CERTIFICATE": "false", "CARGO_TARGET_DIR": str(args.build_dir / "target"),
        "PATH": str(args.dotnet.parent) + os.pathsep + environment.get("PATH", ""),
    })
    manifest = {
        "status": "building", "started_utc": datetime.now(timezone.utc).isoformat(),
        "arguments": {key: str(value) for key, value in vars(args).items()},
        "commands": [], "benchmark_executed": False,
        "orchestrator_sha256": sha256(Path(__file__).resolve()),
        "asset_sha256": {name: sha256(assets / name) for name in ASSETS},
        "source_sha256": tree_hashes(args.source, {".fs", ".fsproj", ".cs", ".csproj"}),
        "compiler_source_sha256": {name: sha256(compiler_sources / name) for name in ("RustPrinter.fs", "Fable2Rust.fs")},
        "package_tools_sha256": tree_hashes(package_tools),
        "package_runtime_sha256": tree_hashes(args.fable_package / "fable-library-rust"),
        "tool_sha256": {"dotnet": sha256(args.dotnet), "python": sha256(Path(sys.executable)), **{name: sha256(Path(shutil.which(name))) for name in ("cargo", "rustc", "patch")}},
    }
    build = Build(args.build_dir, environment, manifest)
    try:
        for name, command in (("dotnet", [args.dotnet, "--version"]), ("cargo", ["cargo", "--version"]), ("rustc", ["rustc", "--version"]), ("python", [sys.executable, "--version"])):
            build.run(f"00-version-{name}", command, args.build_dir, 30)
        perform_build(args, assets, build)
    except Exception as error:
        manifest.update({"status": "failed", "error": str(error)})
        build.save()
        raise


if __name__ == "__main__":
    try:
        main()
    except (OSError, ValueError, RuntimeError) as error:
        print(f"Build failed: {error}", file=sys.stderr)
        sys.exit(1)
