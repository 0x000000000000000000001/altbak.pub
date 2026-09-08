"""Collect exact allocation profiles outside every timing/RSS measurement."""
import json
import os
from pathlib import Path
import re
import subprocess

ROOT = Path(__file__).resolve().parent
ENV = dict(os.environ, GOGC="800", GOMAXPROCS="14", GOMEMLIMIT="off", PPROF="0",
           GOCACHE="/private/tmp/altbak-solod-gocache", GOMODCACHE="/private/tmp/altbak-solod-modcache")
ITERATIONS = 20


def main():
    dest = ROOT / "profiles"
    dest.mkdir(exist_ok=False)
    result = {}
    for variant in ["before", "after"]:
        binary = ROOT / "bin" / (variant + "-bench")
        profile = dest / (variant + ".pprof")
        env = dict(ENV, GOPURS_ALLOC_PROFILE=str(profile), GOPURS_ALLOC_ITERATIONS=str(ITERATIONS))
        output = subprocess.check_output([str(binary), "-test.run=^TestAllocationProfile$",
                                          "-test.memprofilerate=1", "-test.v"], env=env, cwd=ROOT, text=True)
        (dest / (variant + "-profile.log")).write_text(output)
        match = re.search(r"ALLOC StateMonad iterations=(\d+) result=(\d+) bytes=(\d+) allocs=(\d+) gcs=(\d+) bytes/op=([\d.]+) allocs/op=([\d.]+)", output)
        assert match and int(match[1]) == ITERATIONS and match[2] == "1200"
        result[variant] = dict(zip(["iterations", "result", "bytes", "allocs", "gcs", "bytes_per_op", "allocs_per_op"], map(float, match.groups())))
        for metric in ["alloc_space", "alloc_objects"]:
            common = ["go", "tool", "pprof", "-sample_index=" + metric,
                      "-base=" + str(profile) + ".before", "-focus=Test_StateMonad",
                      "-trim_path=gopurs/output", "-source_path=" + str(ROOT / "generated" / variant / "output")]
            for name, option in [("top", "-top"), ("lines", "-list=RecordDict")]:
                text = subprocess.check_output(common + [option, "-nodecount=40", str(binary), str(profile)], env=ENV, cwd=ROOT, stderr=subprocess.STDOUT, text=True)
                (dest / (variant + "-" + metric + "-" + name + ".txt")).write_text(text)
        print(variant, result[variant], flush=True)
    (dest / "summary.json").write_text(json.dumps(result, indent=2) + "\n")


if __name__ == "__main__":
    main()
