"""Shared source and oracle contract for the library benchmarks."""
import hashlib
import json
from pathlib import Path
import re
import shutil

ROOT = Path(__file__).resolve().parents[2]
LIBRARY_ROOT = ROOT / "benchmarks/library"
DEPENDENCIES = ("foldable-traversable", "refs", "enums")


def library_cases():
    manifest = json.loads((LIBRARY_ROOT / "cases.json").read_text())
    if not isinstance(manifest, list) or not manifest:
        raise ValueError("Library cases.json must contain a nonempty list")
    cases = []
    labels = set()
    for case in manifest:
        if not isinstance(case, dict):
            raise ValueError("Each library case must be an object")
        module = case.get("module", "")
        prepare = case.get("prepare", "")
        label = case.get("label", "")
        expected = case.get("expected")
        if not isinstance(module, str) or not re.fullmatch(r"Test\.Library\.[A-Z][A-Za-z0-9_]*(?:\.[A-Z][A-Za-z0-9_]*)*", module):
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
    paths = [LIBRARY_ROOT / "cases.json", *(LIBRARY_ROOT / "src").rglob("*.purs")]
    return {str(path.relative_to(ROOT)): hashlib.sha256(path.read_bytes()).hexdigest()
            for path in sorted(paths)}


def copy_library_sources(destination):
    shutil.copytree(LIBRARY_ROOT / "src", destination, dirs_exist_ok=True)


def library_config(config):
    """Add direct dependencies to a private workspace's existing YAML text."""
    match = re.search(r"(?m)^  dependencies:\n((?:    - [^\n]+\n)*)", config)
    if match is None:
        raise ValueError("Cannot find package dependencies in the backend configuration")
    existing = set(re.findall(r"(?m)^    - ([^\n]+)$", match[1]))
    additions = "".join(f"    - {name}\n" for name in DEPENDENCIES if name not in existing)
    return config[:match.end()] + additions + config[match.end():]
