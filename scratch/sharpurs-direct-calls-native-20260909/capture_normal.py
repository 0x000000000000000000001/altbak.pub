"""Record the regular generated project build and the backend producing it."""
from datetime import datetime, timezone
import json
from pathlib import Path
import subprocess

from common import AUDIT, require, sha

inputs = json.loads((AUDIT / "inputs.json").read_text())
source = Path(inputs["sourceGeneratedDirectory"])
require(inputs["afterSha256"] == inputs["sourceGeneratedSha256"], "After differs from normal input manifest")
require(all(sha(source / name) == expected for name, expected in inputs["afterSha256"].items()),
        "Normal sources changed since capture")
log = AUDIT / "build-normal.log"
content = log.read_text()
require("Build succeeded." in content and "0 Error(s)" in content, "No successful normal build")
runtime = source / "bin/Release/net8.0"
files = sorted(path for path in runtime.iterdir() if path.is_file())
require(runtime / "Program.dll" in files, "Normal Program.dll absent")
backend = AUDIT.parents[2] / "sharpurs/sharpurs"
backend_names = ["bin/sharpurs", "bin/sharpurs.js", "src/Sharpurs/CodeGen.purs", "src/Sharpurs/DirectCall.purs"]
require(all((backend / name).is_file() for name in backend_names), "Backend source or bundle absent")
record = {
    "capturedUtc": datetime.now(timezone.utc).isoformat(),
    "sourceDirectory": str(source), "sourceSha256": inputs["afterSha256"],
    "buildLogSha256": sha(log), "runtimeSha256": {path.name: sha(path) for path in files},
    "backendDirectory": str(backend),
    "backendHead": subprocess.check_output(["git", "rev-parse", "HEAD"], cwd=backend, text=True).strip(),
    "backendSha256": {name: sha(backend / name) for name in backend_names},
}
(AUDIT / "build-normal.json").write_text(json.dumps(record, indent=2) + "\n")
print(f"Recorded normal build: {len(inputs['afterSha256'])} matching inputs and {len(files)} runtime files")
