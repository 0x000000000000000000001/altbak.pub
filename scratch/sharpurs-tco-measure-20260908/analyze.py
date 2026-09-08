"""Validate all recorded answers and summarize observations without dropping runs."""
from pathlib import Path
import json
import re
from statistics import median

audit = Path(__file__).resolve().parent
read = lambda path: json.loads(path.read_text())
result = {"historicalReadmeTcoUs": 184283, "variants": {}}
for variant in ["before", "after"]:
    runs = []
    combined = []
    for index in range(1, 4):
        rows = [json.loads(line) for line in (audit / f"{variant}-{index}.jsonl").read_text().splitlines()]
        data = [row for row in rows if row["kind"] == "measurement"]
        warmups = [row for row in rows if row["kind"] == "warmup"]
        metadata = next(row for row in rows if row["kind"] == "metadata")
        assert len(data) == 10 and len(warmups) == 3 and metadata["globalWarmups"] == 3
        assert metadata["serverGc"] is False
        assert all(row["output"] == "100000" for row in data + warmups)
        times = [row["elapsedUs"] for row in data]
        runs.append({"process": index, "bestUs": min(times), "medianUs": median(times),
                     "gc": [sum(row[f"gc{generation}"] for row in data) for generation in range(3)]})
        combined.extend(data)
    allocated = [row["allocatedBytes"] for row in combined]
    result["variants"][variant] = {
        "runs": runs, "medianBestUs": median(row["bestUs"] for row in runs),
        "medianProcessMedianUs": median(row["medianUs"] for row in runs),
        "pooledMedianUs": median(row["elapsedUs"] for row in combined),
        "allocationBytesPerActionRange": [min(allocated), max(allocated)],
        "gcAcross30Measurements": [sum(row[f"gc{generation}"] for row in combined) for generation in range(3)],
    }
before, after = [result["variants"][variant] for variant in ["before", "after"]]
result["speedupMedianBest"] = before["medianBestUs"] / after["medianBestUs"]
result["speedupMedianProcessMedian"] = before["medianProcessMedianUs"] / after["medianProcessMedianUs"]

expected = {row["case"]: row["output"] for row in read(audit.parent / "sharpurs-apply-step1-20260908/comparison.json")["cases"]}
pattern = re.compile(r"\(Test\)\s*\n(.*?)\n\s*\(Output & Warm-up\)\s*\n(.*?)\n\s*\(Execution time - best of 10\)\s*\n([\d,.]+) μs", re.S)
def official(path):
    text = path.read_text()
    cases = [{"case": label.strip(), "output": output.strip(), "us": float(time.replace(",", ""))}
             for label, output, time in pattern.findall(text)]
    assert len(cases) == 14 and {row["case"]: row["output"] for row in cases} == expected
    total = float(re.search(r"Total exec time: ([\d,.]+) ms", text).group(1).replace(",", ""))
    assert abs(sum(row["us"] for row in cases) / 1000 - total) < 0.02
    tco = next(row["us"] for row in cases if row["case"].startswith("Tail Call"))
    dominant = [row for row in cases if row["case"].startswith(("Polymorphism", "Red-Black", "Lazy Evaluation"))]
    return {"cases": cases, "tcoUs": tco, "totalMs": total, "dominantCasesMs": sum(row["us"] for row in dominant) / 1000,
            "dominantSharePercent": sum(row["us"] for row in dominant) / (total * 10)}

result["officialPairs"] = []
for directory in [audit / "isolated", audit]:
    pair = {variant: official(directory / f"official-{variant}.log") for variant in ["before", "after"]}
    pair["speedupTco"] = pair["before"]["tcoUs"] / pair["after"]["tcoUs"]
    pair["tcoSavedMs"] = (pair["before"]["tcoUs"] - pair["after"]["tcoUs"]) / 1000
    result["officialPairs"].append(pair)

(audit / "comparison.json").write_text(json.dumps(result, indent=2) + "\n")
compact = {key: value for key, value in result.items() if key != "officialPairs"}
compact["officialPairs"] = [
    {"beforeTcoUs": pair["before"]["tcoUs"], "afterTcoUs": pair["after"]["tcoUs"],
     "speedupTco": pair["speedupTco"], "tcoSavedMs": pair["tcoSavedMs"],
     "beforeTotalMs": pair["before"]["totalMs"], "afterTotalMs": pair["after"]["totalMs"],
     "afterDominantSharePercent": pair["after"]["dominantSharePercent"]}
    for pair in result["officialPairs"]]
print(json.dumps(compact, indent=2))
