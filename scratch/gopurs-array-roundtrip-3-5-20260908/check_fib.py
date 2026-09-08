"""Follow up the Fib timing signal without replacing the initial 14-act series."""
import json
from measure import ROOT, BENCH, order, run
from analyze import compare

path = ROOT / "fib-control.jsonl"
assert not path.exists(), "Preserve the existing control series"
for pair in range(1, 11):
    for variant in order(pair):
        text = run([str(ROOT / "bin" / (variant + "-bench")), "-test.run=^$",
                    "-test.bench=^BenchmarkAct/Fib$", "-test.benchtime=1s",
                    "-test.count=1", "-test.cpu=14", "-test.benchmem"],
                   f"fib-control-{pair:02}-{variant}")
        matches = BENCH.findall(text)
        assert len(matches) == 1 and matches[0][0] == "Fib"
        test, iterations, ns, bytes_, allocs = matches[0]
        with path.open("a") as out:
            out.write(json.dumps(dict(kind="act", variant=variant, pair=pair, test=test,
                                      iterations=int(iterations), ns_per_op=float(ns),
                                      bytes_per_op=float(bytes_), allocs_per_op=float(allocs))) + "\n")

rows = [json.loads(line) for line in path.read_text().splitlines()]
grouped = {v: {r["pair"]: r for r in rows if r["variant"] == v} for v in ["before", "after"]}
summary = compare(grouped, ("ns_per_op", "bytes_per_op", "allocs_per_op"), "Fib follow-up")
(ROOT / "fib-control-summary.json").write_text(json.dumps(summary, indent=2) + "\n")
