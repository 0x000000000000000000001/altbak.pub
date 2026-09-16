#!/usr/bin/env python3
"""Compile once; time sequentially in randomized, paired rounds. Scratch only."""
from pathlib import Path
import argparse
import hashlib
import json
import platform
import random
import statistics
import subprocess
import time

ROOT = Path(__file__).resolve().parent
OUT = ROOT / "timing"
VARIANTS = {
    "generated": (ROOT / "borrow/kernel-baseline.rs", False),
    "borrow_depth": (ROOT / "borrow/kernel-borrow.rs", False),
    "borrow_ins": (ROOT / "borrow/kernel-borrow-ins.rs", False),
    "borrow_both": (ROOT / "borrow/kernel-borrow-ins-depth.rs", False),
    "unique_checks": (ROOT / "unique/kernel_unique.rs", False),
}
FLAGS = ["--edition=2021", "-C", "opt-level=3", "-C", "codegen-units=1"]
MAIN = r'''
fn main() {
    let args: Vec<String> = std::env::args().collect();
    let n: i64 = args[1].parse().unwrap();
    let count: usize = args[2].parse().unwrap();
    for sample in 0..count+3 {
        let begin = std::time::Instant::now();
        let value = run(std::hint::black_box(n));
        std::hint::black_box(value);
        let elapsed = begin.elapsed().as_nanos();
        assert_eq!(value, 22);
        if sample >= 3 { println!("{}", elapsed); }
    }
}
'''
RUN = r'''
#[inline(never)]
fn run(n: i64) -> i64 {
    let tree = Test_RBTree_buildTree(n, std::rc::Rc::new(Tree::E));
    Test_RBTree_depth(std::hint::black_box(tree))
}
'''

def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()

def main():
    global OUT
    parser = argparse.ArgumentParser()
    parser.add_argument("--compile-only", action="store_true")
    parser.add_argument("--skip-compile", action="store_true")
    parser.add_argument("--rounds", type=int, default=21)
    parser.add_argument("--samples", type=int, default=10)
    parser.add_argument("--with-owned", action="store_true")
    parser.add_argument("--final-checks", action="store_true")
    args = parser.parse_args()
    if args.with_owned:
        VARIANTS.update({
            "owned_rc": (ROOT / "unique/owned/kernel.rs", True),
            "owned_box": (ROOT / "unique/owned/kernel.rs", True),
        })
    if args.final_checks:
        VARIANTS.clear()
        VARIANTS.update({
            "generated": (ROOT / "borrow/kernel-baseline.rs", False),
            "borrow_ins": (ROOT / "borrow/kernel-borrow-ins.rs", False),
            "borrow_ins_hoisted": (ROOT / "borrow/kernel-borrow-ins-hoisted.rs", False),
            "unique_checks": (ROOT / "unique/kernel_unique.rs", False),
            "unique_sites": (ROOT / "unique/kernel_unique_sites.rs", False),
        })
        OUT = ROOT / "timing-checks"
    OUT.mkdir(exist_ok=True)
    manifest = {"rustc": subprocess.check_output(["rustc", "-Vv"], text=True),
                "platform": platform.platform(), "flags": FLAGS,
                "n": 100000, "warmups": 3, "samples": args.samples,
                "rounds": args.rounds, "seed": 16092026,
                "timed_scope": "build descending 100k + depth + full destruction",
                "variants": {}}
    for name, (kernel, own_run) in VARIANTS.items():
        source, binary = OUT / (name + ".rs"), OUT / name
        extra_flags = ["--cfg", "rc_owner"] if name == "owned_rc" else []
        if not args.skip_compile:
            source.write_text('#![allow(warnings)]\n#![recursion_limit="512"]\n'
                              + 'include!(' + json.dumps(str(kernel)) + ');\n'
                              + ("" if own_run else RUN) + MAIN)
            subprocess.run(["rustc", *FLAGS, *extra_flags, str(source), "-o", str(binary)],
                           check=True, capture_output=True, text=True, timeout=60)
        manifest["variants"][name] = {"kernel": str(kernel), "kernel_sha": sha(kernel),
                                      "extra_flags": extra_flags,
                                      "harness_sha": sha(source), "binary_sha": sha(binary)}
    (OUT / "manifest.json").write_text(json.dumps(manifest, indent=2) + "\n")
    if args.compile_only:
        print("compiled: " + ", ".join(VARIANTS))
        return
    rounds = []
    rng = random.Random(16092026)
    for idx in range(args.rounds):
        order = list(VARIANTS)
        rng.shuffle(order)
        samples = {}
        for name in order:
            result = subprocess.run([str(OUT / name), "100000", str(args.samples)],
                                    capture_output=True, text=True, check=True, timeout=30)
            values = [int(s) / 1000.0 for s in result.stdout.splitlines()]
            assert len(values) == args.samples
            samples[name] = {"us": values, "min_us": min(values),
                             "median_us": statistics.median(values)}
        rounds.append({"index": idx, "order": order, "samples": samples})
        print("round", idx+1, {k: round(v["min_us"], 2) for k, v in samples.items()}, flush=True)
    summary = {}
    for name in VARIANTS:
        mins = [r["samples"][name]["min_us"] for r in rounds]
        medians = [r["samples"][name]["median_us"] for r in rounds]
        gains = [100*(1-r["samples"][name]["min_us"]/r["samples"]["generated"]["min_us"])
                 for r in rounds]
        summary[name] = {"median_best10_us": statistics.median(mins),
                         "median_median10_us": statistics.median(medians),
                         "best10_range_us": [min(mins), max(mins)],
                         "median_paired_gain_pct": statistics.median(gains),
                         "paired_gain_range_pct": [min(gains), max(gains)]}
    if "owned_box" in summary:
        ratios = [100*(1-r["samples"]["owned_box"]["min_us"]/r["samples"]["owned_rc"]["min_us"])
                  for r in rounds]
        summary["owned_box"]["paired_gain_vs_owned_rc_pct"] = statistics.median(ratios)
    result = {"manifest": manifest, "rounds": rounds, "summary": summary,
              "finished_unix": time.time()}
    (OUT / "results.json").write_text(json.dumps(result, indent=2) + "\n")
    print(json.dumps(summary, indent=2))

if __name__ == "__main__":
    main()
