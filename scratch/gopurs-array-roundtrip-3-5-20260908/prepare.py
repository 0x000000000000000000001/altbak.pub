"""Freeze compiler/input variants, regenerate both Go trees, then build once."""
import difflib
import hashlib
import json
import os
from pathlib import Path
import shutil
import subprocess

ROOT = Path(__file__).resolve().parent
ALTBAK = ROOT.parents[1]
GOPURS = ALTBAK.parent / "gopurs/gopurs"
OUTPUT = ALTBAK / "run/bak/go/output"
ENV = dict(os.environ, GOCACHE="/private/tmp/altbak-solod-gocache",
           GOMODCACHE="/private/tmp/altbak-solod-modcache", GOGC="800",
           GOMAXPROCS="14", GOMEMLIMIT="off", GOFLAGS="-pgo=off", PPROF="0")


def digest(data):
    return hashlib.sha256(data).hexdigest()


def hashes(root, pattern):
    return {str(p.relative_to(root)): digest(p.read_bytes())
            for p in sorted(root.glob(pattern)) if p.is_file()}


def run(args, cwd, log):
    print("Preparing:", log, flush=True)
    with (ROOT / "logs" / log).open("w") as out:
        subprocess.run(args, cwd=cwd, env=ENV, stdout=out,
                       stderr=subprocess.STDOUT, check=True)


def main():
    for name in ["logs", "bin", "compiler/bin", "generated"]:
        (ROOT / name).mkdir(parents=True, exist_ok=True)
    for name, target in [("compiler/tools", GOPURS / "tools"),
                         ("generated/gopurs", GOPURS.parent)]:
        link = ROOT / name
        if not link.exists():
            link.symlink_to(target, target_is_directory=True)
    original = (GOPURS / "bin/gopurs.js").read_text()
    start = original.index("var isIntArrayFold = function(fn) {")
    end = original.index("\nvar ", start + 1)
    guard = original[start:end]
    assert guard.count("return true;") == 1
    disabled = original[:start] + guard.replace("return true;", "return false;") + original[end:]
    (ROOT / "compiler/bin/after.mjs").write_text(original)
    (ROOT / "compiler/bin/before.mjs").write_text(disabled)
    (ROOT / "compiler-variant.patch").write_text("".join(difflib.unified_diff(
        original.splitlines(True), disabled.splitlines(True),
        fromfile="current-compiler", tofile="normalization-disabled")))

    corefns = sorted(OUTPUT.glob("*/corefn.json"))
    assert len(corefns) > 200
    input_hashes = hashes(OUTPUT, "*/corefn.json")
    # PBO reads implementations from .purmeta during inlining. Freeze the same
    # initial metadata for both variants; empty metadata changes unrelated Go.
    cache_snapshot = ROOT / "generated/initial-purmeta"
    assert not cache_snapshot.exists(), "Use a fresh metadata snapshot"
    shutil.copytree(ALTBAK / ".purmeta", cache_snapshot)
    cache_hashes = hashes(cache_snapshot, "*.purmeta")
    source_hashes = {}
    for corefn in corefns:
        data = json.loads(corefn.read_text())
        assert "typeTable" in data
        path = ALTBAK / data["modulePath"]
        for suffix in [".purs", ".go"]:
            source = path.with_suffix(suffix).resolve()
            if source.is_file():
                source_hashes[str(source)] = digest(source.read_bytes())
    (ROOT / "inputs-sha256.json").write_text(json.dumps(
        {"corefn": input_hashes, "sources": source_hashes, "purmeta": cache_hashes}, indent=2) + "\n")

    metadata = {
        "go": subprocess.check_output(["go", "version"], text=True).strip(),
        "node": subprocess.check_output(["node", "--version"], text=True).strip(),
        "purs": subprocess.check_output([str(ALTBAK / "run/bak/js/node_modules/.bin/purs"), "--version"], text=True).strip(),
        "gopurs_head": subprocess.check_output(["git", "rev-parse", "HEAD"], cwd=GOPURS, text=True).strip(),
        "pbo_head": subprocess.check_output(["git", "rev-parse", "HEAD"], cwd=ALTBAK.parent / "purescript-backend-optimizer", text=True).strip(),
        "settings": {key: ENV[key] for key in ["GOGC", "GOMAXPROCS", "GOMEMLIMIT", "GOFLAGS", "PPROF"]},
        "compiler_sha256": {"before": digest(disabled.encode()), "after": digest(original.encode())},
        "corefn_count": len(corefns),
    }
    (ROOT / "gopurs-source.patch").write_bytes(subprocess.check_output(
        ["git", "diff", "--", "src/Gopurs/CodeGen.purs", "src/Gopurs/GoAst.purs", "src/Gopurs/Printer.purs"], cwd=GOPURS))

    generated = {}
    for variant in ["before", "after"]:
        work = ROOT / "generated" / variant
        out = work / "output"
        out.mkdir(parents=True, exist_ok=True)
        assert not (out / "purescript").exists(), "Use a fresh generated directory"
        shutil.copytree(cache_snapshot, work / ".purmeta")
        for name, target in [("src", ALTBAK / "src"), (".spago", ALTBAK / ".spago")]:
            (work / name).symlink_to(target, target_is_directory=True)
        for corefn in corefns:
            dest = out / corefn.relative_to(OUTPUT)
            dest.parent.mkdir(parents=True, exist_ok=True)
            shutil.copy2(corefn, dest)
        run(["node", "--expose-gc", "--stack-size=65536", "--max-old-space-size=16384",
             str(ROOT / "compiler/bin" / (variant + ".mjs")), "--main", "App"],
            work, variant + "-generation.log")
        generated[variant] = hashes(out, "**/*.go")
        assert (out / "main/main.go").is_file()
        assert (out / "purescript/Test_ArrayOps.go").is_file()
    assert generated["before"].keys() == generated["after"].keys()
    changed = [p for p in generated["before"] if generated["before"][p] != generated["after"][p]]
    assert changed == ["purescript/Test_ArrayOps.go"], changed
    # Rebuild production from the same still-unchanged initial metadata. A
    # previous production run may itself have updated its .purmeta files.
    assert cache_hashes == hashes(ALTBAK / ".purmeta", "*.purmeta")
    run([str(GOPURS / "bin/gopurs"), "--main", "App"], ALTBAK, "production-generation.log")
    assert generated["after"] == hashes(OUTPUT, "**/*.go"), "Production generation differs"
    assert input_hashes == hashes(OUTPUT, "*/corefn.json")
    assert all(digest(Path(p).read_bytes()) == h for p, h in source_hashes.items())
    metadata["generated_go_sha256"] = generated
    metadata["changed_go_files"] = changed
    before = (ROOT / "generated/before/output" / changed[0]).read_text()
    after = (ROOT / "generated/after/output" / changed[0]).read_text()
    (ROOT / "generated-go.patch").write_text("".join(difflib.unified_diff(
        before.splitlines(True), after.splitlines(True), fromfile="before/" + changed[0], tofile="after/" + changed[0])))

    for variant in ["before", "after"]:
        out = ROOT / "generated" / variant / "output"
        shutil.copy2(ROOT / "bench_test.go", out / "purescript/array_roundtrip_bench_test.go")
        run(["go", "test", "-c", "-pgo=off", "-trimpath", "-o", str(ROOT / "bin" / (variant + "-bench")), "./purescript"], out, variant + "-bench-build.log")
        run(["go", "build", "-pgo=off", "-trimpath", "-o", str(ROOT / "bin" / (variant + "-app")), "./main"], out, variant + "-app-build.log")
    run(["go", "build", "-pgo=off", "-trimpath", "-o", "go_app", "./main"], OUTPUT, "production-app-build.log")
    metadata["binaries_sha256"] = hashes(ROOT / "bin", "*")
    (ROOT / "metadata.json").write_text(json.dumps(metadata, indent=2) + "\n")
    print("Ready; only generated change:", changed, flush=True)


if __name__ == "__main__":
    main()
