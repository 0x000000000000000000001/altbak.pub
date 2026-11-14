"""Shared source and oracle contract for the library benchmarks."""
import hashlib
import json
from pathlib import Path
import re
import shutil
from functools import lru_cache

ROOT = Path(__file__).resolve().parents[2]
SOURCE_ROOT = ROOT / "src"
MANIFEST = ROOT / "bin/benchmark/library-cases.json"
SUPPORT_SOURCES = ("AppAll.purs", "AppLibrary.purs",
                   "Bench/ArrayFixtures.purs", "Bench/StringFixtures.purs")
DEPENDENCIES = ("foldable-traversable", "refs", "enums")


def is_library_source(path):
    """Identify optional WIP files by their path relative to src/."""
    path = Path(path)
    return path.as_posix() in SUPPORT_SOURCES or (
        len(path.parts) >= 2 and path.parts[0] == "Test"
        and path.parts[1] in library_test_names())


@lru_cache(maxsize=1)
def library_test_names():
    return frozenset(case["module"].removeprefix("Test.") for case in library_cases())


def library_sources():
    return [*(SOURCE_ROOT / name for name in SUPPORT_SOURCES),
            *(path for name in sorted(library_test_names())
              for path in (SOURCE_ROOT / "Test" / name).rglob("*") if path.is_file())]


def library_cases():
    manifest = json.loads(MANIFEST.read_text())
    if not isinstance(manifest, list) or not manifest:
        raise ValueError("library-cases.json must contain a nonempty list")
    cases = []
    labels = set()
    for case in manifest:
        if not isinstance(case, dict):
            raise ValueError("Each library case must be an object")
        module = case.get("module", "")
        prepare = case.get("prepare", "")
        label = case.get("label", "")
        expected = case.get("expected")
        if not isinstance(module, str) or not re.fullmatch(r"Test\.[A-Z][A-Za-z0-9_]*", module):
            raise ValueError(f"Invalid library module: {module!r}")
        if not isinstance(prepare, str) or not re.fullmatch(r"[a-z][A-Za-z0-9_]*", prepare):
            raise ValueError(f"Invalid library preparation function: {prepare!r}")
        if not isinstance(label, str) or not label or label.strip() != label or "\n" in label or label in labels:
            raise ValueError(f"Missing, repeated or invalid library label: {label!r}")
        if type(expected) is not int or not -(2 ** 31) <= expected < 2 ** 31:
            raise ValueError(f"Library oracle must be a signed 32-bit integer: {label}")
        labels.add(label)
        cases.append({"module": module, "prepare": prepare, "label": label, "value": str(expected)})
    return cases


def library_inputs():
    paths = [MANIFEST, *library_sources()]
    return {str(path.relative_to(ROOT)): hashlib.sha256(path.read_bytes()).hexdigest()
            for path in sorted(paths)}


def copy_library_sources(destination):
    for source in library_sources():
        target = Path(destination) / source.relative_to(SOURCE_ROOT)
        target.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(source, target)


def library_config(config):
    """Add direct dependencies to a private workspace's existing YAML text."""
    match = re.search(r"(?m)^  dependencies:\n((?:    - [^\n]+\n)*)", config)
    if match is None:
        raise ValueError("Cannot find package dependencies in the backend configuration")
    existing = set(re.findall(r"(?m)^    - ([^\n]+)$", match[1]))
    additions = "".join(f"    - {name}\n" for name in DEPENDENCIES if name not in existing)
    return config[:match.end()] + additions + config[match.end():]
