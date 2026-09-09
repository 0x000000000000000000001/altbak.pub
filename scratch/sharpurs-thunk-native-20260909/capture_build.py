"""Record source/assembly provenance immediately after a successful build."""
import argparse
from datetime import datetime, timezone
import json

from common import AUDIT, require, sha

parser = argparse.ArgumentParser(description=__doc__)
parser.add_argument("variant", choices=("before", "after"))
args = parser.parse_args()
manifest = AUDIT / "inputs.json"
inputs = json.loads(manifest.read_text())[f"{args.variant}Sha256"]
directory = AUDIT / args.variant
require(all(sha(directory / name) == expected for name, expected in inputs.items()),
        "Build sources differ from captured inputs")
log = AUDIT / f"build-{args.variant}.log"
content = log.read_text()
require("Build succeeded." in content and "0 Error(s)" in content, "No successful build in log")
runtime = directory / "bin/Release/net8.0"
files = sorted(p for p in runtime.iterdir() if p.is_file())
require((runtime / "Program.dll") in files, "Program.dll absent")
record = {"capturedUtc": datetime.now(timezone.utc).isoformat(),
          "variant": args.variant, "buildLogSha256": sha(log),
          "sourceSha256": inputs,
          "runtimeSha256": {p.name: sha(p) for p in files}}
(AUDIT / f"build-{args.variant}.json").write_text(json.dumps(record, indent=2) + "\n")
print(f"Recorded {args.variant}: {len(inputs)} inputs and {len(files)} runtime files")
