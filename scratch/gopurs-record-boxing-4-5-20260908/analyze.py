#!/usr/bin/env python3
"""Summarize paired StateMonad 4.5 measurements; never run builds or benchmarks.

Run from the measurement directory: python3 analyze.py
Input: measures.jsonl. Outputs: summary.json and summary.md.
Quartiles use the inclusive convention. IQR describes dispersion, not a CI.
"""

import argparse
import json
import math
import statistics
from collections import defaultdict
from pathlib import Path


TESTS = [
    "AstTree", "Fib", "ListOps", "TCO", "Records", "Ackermann", "Church",
    "Primes", "RBTree", "Polymorphism", "StateMonad", "LazyEvaluation",
    "ArrayOps", "RowToList",
]
SIZES = ["n20"]
VARIANTS = {"before", "after"}
COSTS = ("ns_per_op", "bytes_per_op", "allocs_per_op")


def number(value, label, positive=False):
    assert isinstance(value, (int, float)) and not isinstance(value, bool), label
    assert math.isfinite(value), f"{label}: non-finite value"
    assert value > 0 if positive else value >= 0, f"{label}: invalid sign"
    return float(value)


def stats(values):
    values = list(values)
    if not values:
        return {"n": 0, "median": None, "min": None, "max": None,
                "q1": None, "q3": None, "iqr": None}
    q1, _, q3 = (statistics.quantiles(values, n=4, method="inclusive")
                 if len(values) > 1 else [values[0]] * 3)
    return {"n": len(values), "median": statistics.median(values),
            "min": min(values), "max": max(values), "q1": q1,
            "q3": q3, "iqr": q3 - q1}


def check_pairs(rows, label):
    assert set(rows) == VARIANTS, f"{label}: before/after missing"
    before, after = set(rows["before"]), set(rows["after"])
    assert before == after, (
        f"{label}: unmatched pairs; before-only={sorted(before - after)}, "
        f"after-only={sorted(after - before)}")
    assert before, f"{label}: empty group"
    pairs = sorted(before)
    assert pairs == list(range(1, pairs[-1] + 1)), (
        f"{label}: pair numbering must be complete from 1; got {pairs}")
    return pairs


def compare(rows, metrics, label):
    pairs = check_pairs(rows, label)
    result = {"pair_count": len(pairs), "pairs": pairs, "metrics": {}}
    for metric in metrics:
        before = [rows["before"][p][metric] for p in pairs]
        after = [rows["after"][p][metric] for p in pairs]
        bs, ass = stats(before), stats(after)
        ratios = [a / b for b, a in zip(before, after) if b > 0]
        result["metrics"][metric] = {
            "before": bs, "after": ass,
            "paired_delta_after_minus_before": stats(
                a - b for b, a in zip(before, after)),
            "paired_ratio_after_over_before": stats(ratios),
            "paired_percent_change": stats((r - 1) * 100 for r in ratios),
            "ratio_of_medians": ass["median"] / bs["median"]
            if bs["median"] > 0 else None,
            "ratio_undefined_pairs": [p for p, b in zip(pairs, before) if b == 0],
        }
    return result


def insert(group, row, label):
    variant, pair = row["variant"], row["pair"]
    assert variant in VARIANTS, f"{label}: invalid variant {variant!r}"
    assert isinstance(pair, int) and not isinstance(pair, bool) and pair >= 1, (
        f"{label}: invalid pair {pair!r}")
    target = group.setdefault(variant, {})
    assert pair not in target, f"{label}: duplicate {variant} pair {pair}"
    target[pair] = row


def load(input_path):
    groups = {kind: defaultdict(dict) for kind in ("core", "act", "rss")}
    app = {}
    expected_results = {}
    reference_path = input_path.with_name(
        "expected-results.json" if input_path.name == "measures.jsonl"
        else input_path.stem + "-expected-results.json")
    checked_results = json.loads(reference_path.read_text())
    assert list(checked_results) == TESTS, "checked result reference has an invalid inventory"
    rounded_zero_times = []
    lines = 0
    for line_no, line in enumerate(input_path.read_text().splitlines(), 1):
        if not line.strip():
            continue
        row = json.loads(line)
        assert isinstance(row, dict), f"line {line_no}: expected object"
        kind = row.get("kind")
        assert kind in {"core", "act", "app", "rss"}, (
            f"line {line_no}: unknown kind {kind!r}")
        label = f"line {line_no} ({kind})"
        if kind == "app":
            number(row["total_ms"], f"{label}.total_ms", positive=True)
            tests = row["tests"]
            assert isinstance(tests, list) and len(tests) == len(TESTS), (
                f"{label}: expected exactly 14 tests")
            names = [test["test"] for test in tests]
            assert len(set(names)) == len(TESTS) and set(names) == set(TESTS), (
                f"{label}: missing, duplicate or unexpected app tests: {names}")
            for test in tests:
                name = test["test"]
                # App prints rounded times: a displayed 0.00 us is valid, but
                # cannot supply a nonzero baseline for a relative comparison.
                number(test["time_us"], f"{label}.{name}.time_us")
                if test["time_us"] == 0:
                    rounded_zero_times.append({"variant": row["variant"],
                                               "pair": row["pair"], "test": name})
                assert isinstance(test["result"], str), f"{label}.{name}: result not string"
                assert test["result"] == checked_results[name], (
                    f"{label}.{name}: output differs from checked A/B reference")
                if name in expected_results:
                    assert test["result"] == expected_results[name], (
                        f"{label}.{name}: result changed: {test['result']!r} "
                        f"!= {expected_results[name]!r}")
                expected_results[name] = test["result"]
            insert(app, row, label)
        else:
            test = row["test"]
            valid = TESTS if kind == "act" else SIZES
            assert test in valid, f"{label}: unexpected test {test!r}"
            if kind == "rss":
                number(row["rss_bytes"], f"{label}.rss_bytes", positive=True)
                assert isinstance(row["iterations"], int) and not isinstance(
                    row["iterations"], bool) and row["iterations"] >= 1, (
                        f"{label}: invalid iterations")
                assert row["iterations"] == 10000 and row.get("result") == 1200, (
                    f"{label}: expected 10000 iterations with result 1200")
            else:
                assert isinstance(row["iterations"], int) and not isinstance(
                    row["iterations"], bool) and row["iterations"] >= 1, (
                        f"{label}: invalid benchmark iterations")
                assert isinstance(row.get("benchtime"), str) and row["benchtime"], (
                    f"{label}: missing benchmark duration")
                for metric in COSTS:
                    number(row[metric], f"{label}.{metric}", positive=metric == "ns_per_op")
            insert(groups[kind][test], row, label)
        lines += 1

    summary = {
        "schema_version": 1, "input": str(input_path), "measurement_lines": lines,
        "checked_results": str(reference_path),
        "experiment": "StateMonad record boxing 4.5", "core": {},
        "statistics": {
            "quartiles": "inclusive", "iqr": "q3 - q1; descriptive, not a confidence interval",
            "paired_ratio": "after / before for the same pair; median of these ratios",
            "paired_delta": "after - before for the same pair; median of these deltas",
            "ratio_of_medians": "reported separately; not the median paired ratio",
            "zero_baseline": "delta retained; ratio/percent change omitted for that pair",
        },
        "act": {}, "rss": {}, "app": {},
    }
    for kind in ("core", "act", "rss"):
        expected = TESTS if kind == "act" else SIZES
        assert set(groups[kind]) == set(expected), f"{kind}: incomplete test inventory"
        common_pairs = None
        for test in expected:
            rows = groups[kind][test]
            metrics = ("rss_bytes",) if kind == "rss" else COSTS
            data = compare(rows, metrics, f"{kind}/{test}")
            if common_pairs is None:
                common_pairs = data["pairs"]
            assert data["pairs"] == common_pairs, (
                f"{kind}/{test}: pair set differs from other tests of this kind")
            if kind == "rss":
                for pair in data["pairs"]:
                    assert rows["before"][pair]["iterations"] == rows["after"][pair]["iterations"], (
                        f"rss/{test}: iteration count differs within pair {pair}")
                data["iterations_by_pair"] = {
                    str(p): rows["before"][p]["iterations"] for p in data["pairs"]}
                data["decimal_MB"] = {
                    variant: stats(rows[variant][p]["rss_bytes"] / 1_000_000
                                   for p in data["pairs"])
                    for variant in ("before", "after")}
                data["binary_MiB"] = {
                    variant: stats(rows[variant][p]["rss_bytes"] / 1_048_576
                                   for p in data["pairs"])
                    for variant in ("before", "after")}
            else:
                durations = {row["benchtime"] for variant_rows in rows.values()
                             for row in variant_rows.values()}
                assert len(durations) == 1, (
                    f"{kind}/{test}: mixed benchmark durations; keep pilot/control separate")
                data["benchtime"] = next(iter(durations))
            summary[kind][test] = data

    summary["app"]["total"] = compare(app, ("total_ms",), "app/total")
    summary["app"]["results_identical"] = True
    summary["app"]["results"] = expected_results
    summary["app"]["rounded_zero_times"] = rounded_zero_times
    summary["app"]["tests"] = {}
    for test in TESTS:
        rows = {variant: {
            pair: next(item for item in row["tests"] if item["test"] == test)
            for pair, row in app[variant].items()} for variant in VARIANTS}
        summary["app"]["tests"][test] = compare(rows, ("time_us",), f"app/{test}")
    return summary


def fmt(value):
    return "—" if value is None else f"{value:.6g}"


def spread(value, scale=1):
    if value["n"] == 0:
        return "non défini"
    return (f"{fmt(value['median'] / scale)} [{fmt(value['min'] / scale)} ; "
            f"{fmt(value['max'] / scale)}] ; {fmt(value['iqr'] / scale)}")


def markdown(summary):
    out = [
        "# Boxing des records — StateMonad, mesures intégrées 4.5", "",
        "Les valeurs sont présentées sous la forme **médiane [minimum ; maximum] ; IQR**. "
        "L’IQR utilise les quartiles inclusifs et décrit la dispersion ; ce n’est pas un intervalle de confiance.", "",
        "Les ratios et variations sont calculés **après/avant pour chaque paire**, puis résumés. "
        "Un ratio de temps inférieur à 1 indique moins de temps. La médiane des différences appariées "
        "est distincte de la différence entre les deux médianes.", "",
        "Les écarts faibles peuvent relever du bruit de mesure. Aucun test de significativité n’est effectué ; "
        "un gain local ne permet pas de conclure à une accélération générale.", "",
        "**Allocations et mémoire résidente sont distinctes** : les octets/op sont cumulés pendant une exécution ; "
        "le RSS décrit la mémoire résidente rapportée pour le processus et ne doit pas être divisé par les itérations. "
        "Les Mo/MB ci-dessous valent 1 000 000 octets ; les Mio/MiB sont aussi disponibles dans le JSON.", "",
    ]

    def table(title, groups, metrics):
        out.extend([f"## {title}", "",
                    "| Test | Mesure | Avant | Après | Δ apparié médian | Ratio apparié | Variation appariée médiane |",
                    "|---|---|---:|---:|---:|---:|---:|"])
        for name, data in groups:
            for metric, label, scale in metrics:
                m = data["metrics"][metric]
                delta = m["paired_delta_after_minus_before"]["median"] / scale
                percent = m["paired_percent_change"]["median"]
                ratio = spread(m["paired_ratio_after_over_before"])
                if m["ratio_undefined_pairs"]:
                    ratio += (f" ({m['paired_ratio_after_over_before']['n']}/"
                              f"{data['pair_count']} paires)")
                percent_text = "non définie" if percent is None else f"{fmt(percent)} %"
                out.append(
                    f"| {name} ({data['pair_count']} paires) | {label} | "
                    f"{spread(m['before'], scale)} | {spread(m['after'], scale)} | "
                    f"{fmt(delta)} | {ratio} | {percent_text} |")
        out.append("")

    costs = [("ns_per_op", "µs/op", 1000), ("bytes_per_op", "octets/op", 1),
             ("allocs_per_op", "allocations/op", 1)]
    table("Noyau StateMonad — runManyTimes(20, 0), résultat 1200", summary["core"].items(), costs)
    table("Les 14 actes (avec leur enveloppe)", summary["act"].items(), costs)
    table("Application : temps de chaque test", summary["app"]["tests"].items(),
          [("time_us", "µs", 1)])
    zero_times = summary["app"]["rounded_zero_times"]
    if zero_times:
        affected = ", ".join(test for test in TESTS
                             if any(row["test"] == test for row in zero_times))
        out.extend([
            f"**Temps affichés arrondis à zéro : {len(zero_times)} observations ({affected}).** "
            "Un affichage de 0,00 µs ne démontre pas un coût nul. Les valeurs restent dans les "
            "statistiques absolues ; les paires dont le temps avant vaut zéro sont exclues des ratios "
            "et pourcentages. Le nombre de paires utilisables est indiqué dans leur cellule. "
            "Les ratios impliquant un temps après arrondi à zéro ne prouvent pas une suppression totale du coût.",
            "",
        ])
    table("Application : total rapporté", [("Total", summary["app"]["total"])],
          [("total_ms", "ms", 1)])
    out.extend([
        "Le total de l’application est celui fourni par son protocole : il n’est pas reconstruit "
        "en additionnant des médianes. S’il provient de Bench.runBench, il additionne les minima "
        "de dix mesures par test, et ne représente pas le temps mural de tout le processus.", "",
        "**Les 14 résultats de l’application sont identiques dans toutes les lignes avant et après.**", "",
        "| Test | Résultat |", "|---|---|",
    ])
    for test in TESTS:
        value = summary["app"]["results"][test].replace("|", "\\|").replace("\n", "<br>")
        out.append(f"| {test} | {value} |")
    out.append("")
    table("RSS rapporté par processus", summary["rss"].items(), [("rss_bytes", "Mo / MB", 1_000_000)])
    for test, data in summary["rss"].items():
        iterations = sorted(set(data["iterations_by_pair"].values()))
        out.append(f"- {test} : itérations par processus {', '.join(map(str, iterations))} ; "
                   "même nombre avant/après au sein de chaque paire.")
    out.extend(["", "## Lecture prudente des temps", ""])
    for kind, groups in (("Noyau", summary["core"]), ("Acte", summary["act"]),
                         ("Application", {"total": summary["app"]["total"]})):
        metric = "total_ms" if kind == "Application" else "ns_per_op"
        for test, data in groups.items():
            pc = data["metrics"][metric]["paired_percent_change"]
            notes = []
            if abs(pc["median"]) < 5:
                notes.append("écart médian inférieur à 5 %, à ne pas surinterpréter")
            if pc["q1"] <= 0 <= pc["q3"]:
                notes.append("les quartiles des variations encadrent zéro")
            if notes:
                out.append(f"- {kind} {test} : {fmt(pc['median'])} % ; {' ; '.join(notes)}.")
    out.extend(["", "Le seuil de 5 % est un repère de lecture, pas un seuil statistique. "
                "Les statistiques complètes, différences, ratios et numéros des paires figurent dans summary.json.", ""])
    return "\n".join(out)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--input", type=Path, default=Path(__file__).resolve().parent / "measures.jsonl")
    parser.add_argument("--output-prefix", type=Path,
                        help="output prefix; defaults to summary beside primary input, or INPUT-summary")
    args = parser.parse_args()
    input_path = args.input.resolve()
    prefix = args.output_prefix or input_path.with_name(
        "summary" if input_path.name == "measures.jsonl" else input_path.stem + "-summary")
    summary = load(input_path)  # Validate everything before writing either output.
    rendered = markdown(summary)
    json_path = Path(str(prefix) + ".json")
    markdown_path = Path(str(prefix) + ".md")
    json_path.parent.mkdir(parents=True, exist_ok=True)
    json_path.write_text(json.dumps(summary, indent=2, ensure_ascii=False, allow_nan=False) + "\n")
    markdown_path.write_text(rendered)
    print(f"Validated {summary['measurement_lines']} measurement lines; wrote {json_path} and {markdown_path}.")


if __name__ == "__main__":
    main()
