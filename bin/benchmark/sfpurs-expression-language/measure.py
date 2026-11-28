#!/usr/bin/env python3
"""Measure sfpurs-expression-language application benchmark: official JS vs gopurs Go."""
import argparse
from datetime import datetime, timezone
import hashlib
import json
import os
from pathlib import Path
import platform
import shutil
import statistics
import subprocess
import sys
import time

HERE = Path(__file__).resolve().parent
SFPURS_ROOT = HERE.parents[3] / "sfpurs" / "sfpurs-expression-language"
CORPUS = HERE / "corpus.tsv"
DRIVER_JS = HERE / "driver.mjs"
DRIVER_GO = HERE / "driver.go"
GO_OUTPUT = SFPURS_ROOT / "targets" / "go" / "output"
JS_OUTPUT = SFPURS_ROOT / "targets" / "js" / "output"


def sha256_file(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def run_once(cmd: list[str]) -> tuple[float, str]:
    t0 = time.perf_counter_ns()
    proc = subprocess.run(cmd, check=True, capture_output=True, text=True)
    dt_ms = (time.perf_counter_ns() - t0) / 1e6
    return dt_ms, proc.stdout.strip()


def build_go_binary(bin_path: Path) -> None:
    tmp_driver = GO_OUTPUT / "bench_app_driver.go"
    shutil.copy2(DRIVER_GO, tmp_driver)
    try:
        subprocess.run(
            ["go", "build", "-o", str(bin_path), "bench_app_driver.go"],
            cwd=GO_OUTPUT,
            check=True,
        )
    finally:
        tmp_driver.unlink(missing_ok=True)


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--runs", type=int, default=24)
    parser.add_argument("--iters", type=int, default=1000)
    parser.add_argument("--output", type=Path, default=HERE / "results.json")
    args = parser.parse_args()

    go_bin = HERE / ".bench-go-bin"
    build_go_binary(go_bin)

    try:
        js_cmd = ["node", str(DRIVER_JS), str(CORPUS), str(args.iters)]
        go_cmd = [str(go_bin), str(CORPUS), str(args.iters)]

        # Warmup (1 excluded run per backend) + oracle verification
        _, js_oracle = run_once(js_cmd)
        _, go_oracle = run_once(go_cmd)
        if js_oracle != go_oracle:
            print(f"ORACLE MISMATCH!\n  JS: {js_oracle}\n  Go: {go_oracle}", file=sys.stderr)
            return 1
        print(f"Oracle verified (identical JS & Go): {go_oracle}")

        samples = {"JS": [], "Go": []}
        orders = [("JS", "Go"), ("Go", "JS")]
        for r in range(args.runs):
            for variant in orders[r % 2]:
                cmd = js_cmd if variant == "JS" else go_cmd
                dt_ms, out = run_once(cmd)
                if out != js_oracle:
                    print(f"Run {r} {variant} checksum mismatch: {out} != {js_oracle}", file=sys.stderr)
                    return 1
                samples[variant].append(dt_ms)

        js_med = statistics.median(samples["JS"])
        go_med = statistics.median(samples["Go"])
        ratio = go_med / js_med

        node_ver = subprocess.check_output(["node", "-v"], text=True).strip()
        go_ver = subprocess.check_output(["go", "version"], text=True).strip()
        js_header = (JS_OUTPUT / "Sfpurs.ExpressionLanguage.Api" / "index.js").read_text().splitlines()[0]

        report = {
            "schema": "sfpurs-expression-language-runtime-v1",
            "created_at_utc": datetime.now(timezone.utc).isoformat(),
            "host": {
                "system": f"{platform.system()}-{platform.release()}-{platform.machine()}",
                "machine": platform.machine(),
                "logical_cpus": os.cpu_count(),
            },
            "versions": {"node": node_ver, "go": go_ver},
            "frontends": {"JS": js_header, "Go_TAST": "0.15.17"},
            "protocol": {
                "corpus_expressions": 16,
                "ops_per_expression": 3,
                "iterations_per_process": args.iters,
                "total_ops_per_process": 16 * 3 * args.iters,
                "runs_per_backend": args.runs,
                "timing": "whole process including startup; compile + lint + dump over corpus.tsv with FNV-1a checksum over all UTF-8 output bytes",
                "order": "alternating JS/Go and Go/JS; 1 excluded warmup per backend; no sample exclusions",
            },
            "oracle": js_oracle,
            "corpus_sha256": sha256_file(CORPUS),
            "summary": {
                "JS_median_ms": js_med,
                "Go_median_ms": go_med,
                "Go_over_JS": ratio,
                "JS_min_ms": min(samples["JS"]),
                "Go_min_ms": min(samples["Go"]),
                "JS_max_ms": max(samples["JS"]),
                "Go_max_ms": max(samples["Go"]),
            },
            "samples_ms": samples,
        }

        args.output.write_text(json.dumps(report, indent=2) + "\n")
        print(f"JS median ({args.runs} runs): {js_med:.2f} ms (min {min(samples['JS']):.2f}, max {max(samples['JS']):.2f})")
        print(f"Go median ({args.runs} runs): {go_med:.2f} ms (min {min(samples['Go']):.2f}, max {max(samples['Go']):.2f})")
        print(f"Go / JS ratio: {ratio:.4f}x ({1.0/ratio:.2f}x faster)")
        print(f"Saved report to {args.output}")
        return 0
    finally:
        go_bin.unlink(missing_ok=True)


if __name__ == "__main__":
    raise SystemExit(main())
