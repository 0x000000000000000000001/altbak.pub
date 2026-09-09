"""Validate and summarize official output logs; does not execute either program."""
import argparse
import json

from common import AUDIT, ORDER, RBT, LAZY, baseline, check_outputs, compare, read_run, require, sha, stats

parser = argparse.ArgumentParser(description=__doc__)
parser.add_argument("--probe", action="store_true", help="Analyze the preliminary one-pair probe")
args = parser.parse_args()
prefix, count = ("probe-", 1) if args.probe else ("", 5)
reference = read_run(AUDIT / "expected-output.log")
names = [test["name"] for test in reference["tests"]]
official = baseline()
metadata = json.loads((AUDIT / f"{prefix}metadata.json").read_text())
executions = json.loads((AUDIT / f"{prefix}executions.json").read_text())
require([run["variant"] for run in executions] == (["before", "after"] if args.probe else ORDER),
        "Incorrect number or order of benchmark processes")
require(metadata.get("sourceManifestsUnchanged") is True and "finishedUtc" in metadata,
        "Runner did not finish source provenance validation")
require(metadata["historicalOutputSha256"] == sha(AUDIT / "expected-output.log"), "Historical reference changed")
require(metadata["inputsSha256"] == sha(AUDIT / "inputs.json"), "Input manifest changed")
require(metadata["officialReadme"]["sha256"] == official["sha256"], "README snapshot changed")
for variant, expected in metadata["buildManifestsSha256"].items():
    require(sha(AUDIT / f"build-{variant}.json") == expected, "Build manifest changed")
for run in executions:
    require(run["logSha256"] == sha(AUDIT / f"{run['label']}.log"), "Measurement log changed")
    require(run["assemblySha256"] == metadata["assemblies"][run["variant"]]["sha256"], "Assembly changed between processes")
runs = {variant: [read_run(AUDIT / f"{prefix}official-{variant}-{index}.log") for index in range(1, count + 1)]
        for variant in ("before", "after")}
for series in runs.values():
    for run in series:
        check_outputs(run, reference)
summaries = {variant: {"total_ms": stats([run["total_ms"] for run in series]),
                       "tests_us": {name: stats([run["tests"][index]["best_of_10_us"] for run in series])
                                    for index, name in enumerate(names)}}
             for variant, series in runs.items()}
changes = {"total_ms": compare(summaries["before"]["total_ms"]["median"], summaries["after"]["total_ms"]["median"]),
           "tests_us": {name: compare(summaries["before"]["tests_us"][name]["median"],
                                      summaries["after"]["tests_us"][name]["median"]) for name in names}}
after = summaries["after"]
results = {"validation": {"passed": True, "official_processes_per_variant": count,
                          "tests_per_process": 14, "outputs_and_order_match_historical": True,
                          "sum_rounding_tolerance_ms": 0.02, "provenance_verified": True},
           "method": {"official": f"Unmodified App.main entry point: best of 10 per test, {count} fresh processes per variant.",
                      "order": metadata["order"], "summary": "Median and full range of process-level best-of-10 measurements.",
                      "caution": "Descriptive statistics; per-test medians need not sum to total median. One-pair probe is not the final performance conclusion."},
           "runs": runs, "summaries": summaries, "changes": changes,
           "official_readme_baseline": official,
           "after_vs_official_readme": {"total_ms": compare(official["compiled_total_ms"], after["total_ms"]["median"]),
                                         "lazy_us": compare(official["compiled_lazy_us"], after["tests_us"][LAZY]["median"]),
                                         "rbtree_us": compare(official["compiled_rbtree_us"], after["tests_us"][RBT]["median"])}}
(AUDIT / f"{prefix}comparison.json").write_text(json.dumps(results, indent=2, ensure_ascii=False) + "\n")
print(f"Validated {count * 2} App.main logs and {count * 28} historical outputs.")
for variant, summary in summaries.items():
    total, lazy, rbt = summary["total_ms"], summary["tests_us"][LAZY], summary["tests_us"][RBT]
    print(f"{variant}: total {total['median']:.2f} ms [{total['min']:.2f}, {total['max']:.2f}], "
          f"Lazy {lazy['median'] / 1000:.5f} ms [{lazy['min'] / 1000:.5f}, {lazy['max'] / 1000:.5f}], "
          f"RBTree {rbt['median'] / 1000:.5f} ms [{rbt['min'] / 1000:.5f}, {rbt['max'] / 1000:.5f}]")
for label, change in (("Total", changes["total_ms"]), ("Lazy", changes["tests_us"][LAZY]), ("RBTree (unchanged control)", changes["tests_us"][RBT])):
    print(f"{label}: {change['difference_percent']:+.2f}%, speedup {change['speedup_factor']:.3f}x")
print(f"Official README: {official['compiled_total_ms']:.2f} ms total; "
      f"Lazy compiled {official['compiled_lazy_us'] / 1000:.5f} ms, "
      f"native optimized {official['native_optimized_lazy_us'] / 1000:.3f} ms (different workload/result).")
