"""Verify final captured sources, runtimes and backend; never executes a benchmark."""
from datetime import datetime, timezone
import json
from pathlib import Path

from common import AUDIT, require, sha

inputs = json.loads((AUDIT / "inputs.json").read_text())
metadata = json.loads((AUDIT / "metadata.json").read_text())
source = Path(inputs["sourceGeneratedDirectory"])
backend = AUDIT.parents[2] / "sharpurs/sharpurs"
require(metadata["inputsSha256"] == sha(AUDIT / "inputs.json"), "Measurement inputs changed")
counts = {}


def verify(directory, hashes):
    for name, expected in hashes.items():
        require(sha(directory / name) == expected, f"Changed file: {directory / name}")
    return len(hashes)


for variant, directory, key in [
    ("before", AUDIT / "before", "beforeSha256"),
    ("after", AUDIT / "after", "afterSha256"),
    ("normal", source, "sourceGeneratedSha256"),
]:
    hashes = inputs[key]
    names = {p.name for p in directory.iterdir() if p.is_file()
             and p.suffix in {".fs", ".cs", ".fsproj", ".csproj", ".props"}}
    require(names == set(hashes) and len(hashes) == 355, "Source file set changed")
    build = json.loads((AUDIT / f"build-{variant}.json").read_text())
    require(build["sourceSha256"] == hashes, "Build inputs differ")
    require(sha(AUDIT / f"build-{variant}.log") == build["buildLogSha256"], "Build log changed")
    counts[variant] = {"sources": verify(directory, hashes),
                       "runtime": verify(directory / "bin/Release/net8.0", build["runtimeSha256"])}

counts["normalBeforeRuntime"] = verify(AUDIT / "normal-before-runtime", inputs["normalRuntimeBeforeSha256"])
counts["backendBeforeBin"] = verify(AUDIT / "backend-before-bin", inputs["backendBeforeBinSha256"])
counts["backendAfterSources"] = verify(backend, inputs["backendAfterSourcesSha256"])
for directory in [AUDIT / "backend-after-bin", backend / "bin"]:
    require(sha(directory / "sharpurs.js") == inputs["backendAfterBundleSha256"], "After backend bundle changed")
require(sha(AUDIT / "README.baseline.md") == inputs["officialReadmeSha256"], "Official baseline changed")
record = {"checkedUtc": datetime.now(timezone.utc).isoformat(), "passed": True,
          "counts": counts, "inputsSha256": sha(AUDIT / "inputs.json"),
          "metadataSha256": sha(AUDIT / "metadata.json"),
          "comparisonSha256": sha(AUDIT / "comparison.json")}
(AUDIT / "integrity-final.json").write_text(json.dumps(record, indent=2) + "\n")
print("Final integrity passed:", json.dumps(counts))
