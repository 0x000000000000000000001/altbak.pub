#!/usr/bin/env python3
"""Adapt audited sharpurs source shapes in a separate scratch copy.

Remove unused two-argument boxed aliases and extract the typed unit-thunk
factory without changing its work. A separate helper validates reflection
forwarders for the build orchestrator. Optionally split unreferenced top-level
aliases out of let-rec groups. Every adaptation has a strict source-shape guard.
"""
from __future__ import annotations
import argparse
import hashlib
import json
from pathlib import Path
import re
import shutil
import sys


def digest(text: str) -> str:
    return hashlib.sha256(text.encode()).hexdigest()


def propagate_reflection_exceptions(text: str) -> tuple[str, list[dict]]:
    """Adapt only the two audited sharpurs RBTree forwarding-adapter variants.

    Older sharpurs emits boxed max/insert adapters; its typed ADT path emits
    balance/ins/insert adapters. Check complete signatures and forwarding calls,
    rather than accepting arbitrary exception occurrences based on their count.
    """
    wrapper = "raise (System.Reflection.TargetInvocationException(ex))"
    variants = (
        (
            ("Test_RBTree_max_direct", (("x", "obj"), ("y", "obj")), "obj"),
            ("Test_RBTree_insert_direct", (("x", "obj"), ("s", "obj")), "obj"),
        ),
        (
            ("Test_RBTree_balance_adt_native", tuple(zip(
                (f"sharpurs_adt_local_{i}" for i in range(4)),
                ("Test_RBTree_Color", "Test_RBTree_Tree", "int", "Test_RBTree_Tree"))), "Test_RBTree_Tree"),
            ("Test_RBTree_ins_adt_native", (("sharpurs_adt_local_0", "int"),
                ("sharpurs_adt_local_1", "Test_RBTree_Tree")), "Test_RBTree_Tree"),
            ("Test_RBTree_insert_adt_native", (("sharpurs_adt_local_0", "int"),
                ("sharpurs_adt_local_1", "Test_RBTree_Tree")), "Test_RBTree_Tree"),
        ),
    )
    occurrences = text.count("System.Reflection.TargetInvocationException")
    for variant in variants:
        if occurrences != len(variant):
            continue
        blocks = []
        for callee, parameters, result_type in variant:
            name = callee + "_apply"
            signature = " ".join(f"({arg}: {typ})" for arg, typ in parameters)
            success = "    try " + callee + " " + " ".join(arg for arg, _ in parameters) + "\n"
            block = f"let {name} {signature} : {result_type} =\n{success}    with ex -> {wrapper}\n"
            pattern = re.compile(r"^" + re.escape(block), re.M)
            matches = list(pattern.finditer(text))
            declarations = re.findall(r"^let (?:rec )?" + re.escape(callee) + r"\b", text, re.M)
            if len(matches) != 1 or len(declarations) != 1:
                break
            blocks.append((name, matches[0], success))
        else:
            actions = []
            rewritten = text
            for name, match, success in sorted(blocks, key=lambda item: item[1].start(), reverse=True):
                replacement = match.group().replace(wrapper, "raise ex")
                rewritten = rewritten[:match.start()] + replacement + rewritten[match.end():]
                actions.append({
                    "action": "propagate_original_exception", "adapter": name,
                    "source_line": text[:match.start()].count("\n") + 1,
                    "success_path_sha256_before": digest(success),
                    "success_path_sha256_after": digest(success),
                    "proof": "Exact typed forwarding signature and call checked; only the exception rethrow changes.",
                })
            return rewritten, list(reversed(actions))
    raise ValueError(f"Unsupported RBTree reflection adapters: found {occurrences} wrappers, "
                     "but neither audited set of exact signatures and forwarding calls matched")


def adapt_typed_thunk_capture(text: str) -> tuple[str, list[dict]]:
    """Give Fable an explicit value parameter to capture in each thunk step."""
    name = "Test_LazyEvaluation_buildThunks_thunk_native"
    if not re.search(r"^let rec private " + name + r"\b", text, re.M):
        return text, []
    helper = "sharpurs_fable_lazyStep"
    if helper in identifiers(text):
        raise ValueError("Typed thunk adaptation helper name already exists")
    closure = "((fun (sharpurs_thunk_local_2: unit) -> ((sharpurs_thunk_local_1 ()) + (1))))"
    declaration = (f"let rec private {name} (sharpurs_thunk_local_0: int) "
        "(sharpurs_thunk_local_1: (unit -> int)) : (unit -> int) = "
        "(if (sharpurs_thunk_local_0 = (0)) then sharpurs_thunk_local_1 else "
        f"({name} ((sharpurs_thunk_local_0 - (1))) {closure}))\n")
    matches = list(re.finditer(r"^" + re.escape(declaration), text, re.M))
    if len(matches) != 1:
        raise ValueError("Unsupported typed LazyEvaluation thunk builder shape")
    factory = (f"let private {helper} (previous: unit -> int) : unit -> int =\n"
               "    fun () -> previous () + 1\n\n")
    replacement = factory + declaration.replace(closure, f"({helper} sharpurs_thunk_local_1)")
    match = matches[0]
    rewritten = text[:match.start()] + replacement + text[match.end():]
    return rewritten, [{
        "action": "extract_typed_thunk_factory", "enclosing": name, "helper": helper,
        "source_line": text[:match.start()].count("\n") + 1,
        "sha256_before": digest(declaration), "sha256_after": digest(replacement),
        "proof": "Exact builder checked. Each recursive step still allocates a non-memoizing unit thunk that calls the captured previous thunk and adds one. The factory captures a value before Fable reassigns its tail-call parameters.",
    }]


def mask(text: str) -> str:
    # Generated benchmark declarations have ordinary strings and // comments.
    # Refuse more complicated lexical forms instead of guessing at scopes.
    if '(*' in text or '@"' in text:
        raise ValueError('Unsupported F# lexical form (block comment/verbatim string)')
    return re.sub(r'"(?:\\.|[^"\\])*"|//[^\n]*', lambda m: '\n' * m.group().count('\n') + ' ', text)


def identifiers(text: str) -> set[str]:
    return set(re.findall(r'\b\w+\b', mask(text)))


def chunks(text: str):
    starts = [m.start() for m in re.finditer(r'^(?:let|type)\s', text, re.M)]
    if not starts:
        return [('', text, 0)]
    result = [('', text[:starts[0]], 0)]
    starts.append(len(text))
    for a, b in zip(starts, starts[1:]):
        name = re.match(r'(?:let (?:rec )?|type )(\w+)', text[a:b])
        result.append((name.group(1) if name else '', text[a:b], a))
    return result


def transform(text: str, split_top: bool) -> tuple[str, list[dict]]:
    rewritten, actions = [], []
    exact_local_alias = re.compile(
        r'^[ \t]+and (?P<alias>\w+) = box \(\(fun \((?P<a>\w+): obj\) -> '
        r'\(fun \((?P<b>\w+): obj\) -> (?P<primary>\w+_tco) '
        r'(?P=a) (?P=b)\)\)\)[ \t]*(?:\n|$)', re.M,
    )
    for enclosing, original, offset in chunks(text):
        body = original
        local_candidates = list(re.finditer(r'^[ \t]+and \w+\s*=', body, re.M))
        matches = list(exact_local_alias.finditer(body))
        if local_candidates and len(matches) != len(local_candidates):
            raise ValueError(f'{enclosing}: unsupported local recursive alias shape')
        if len(matches) > 1:
            raise ValueError(f'{enclosing}: multiple local groups require deeper scope analysis')
        for match in matches:
            alias, primary = match.group('alias'), match.group('primary')
            remaining = body[:match.start()] + body[match.end():]
            if alias in identifiers(remaining):
                raise ValueError(f'{enclosing}: local alias {alias} is referenced; cannot delete')
            primary_declarations = list(re.finditer(r'^[ \t]+let rec ' + re.escape(primary) + r'\b[^\n]*', body, re.M))
            if len(primary_declarations) != 1 or primary_declarations[0].start() >= match.start():
                raise ValueError(f'{enclosing}: could not prove preceding local recursive implementation')
            if not re.match(r'\s*in\b', body[match.end():]):
                raise ValueError(f'{enclosing}: alias is not the last binding in its recursive group')
            primary_line = primary_declarations[0].group()
            if primary_line not in remaining:
                raise ValueError(f'{enclosing}: recursive implementation changed unexpectedly')
            actions.append({
                'action': 'remove_unused_local_alias', 'enclosing': enclosing,
                'alias': alias, 'primary': primary,
                'source_line': text[:offset + match.start()].count('\n') + 1,
                'removed': match.group().strip(),
                'primary_line_sha256_before': digest(primary_line),
                'primary_line_sha256_after': digest(primary_line),
                'proof': 'The alias is an effect-free boxed two-argument forwarding lambda; its identifier is absent everywhere else in the enclosing declaration.',
            })
            body = remaining
        if split_top and body.startswith('let rec '):
            primary = re.match(r'let rec (\w+_tco)\b', body)
            aliases = list(re.finditer(r'^and (\w+)\b', body, re.M))
            if primary and len(aliases) == 1:
                alias = aliases[0]
                name = alias.group(1)
                if (name not in identifiers(body[:alias.start()])
                        and re.match(r'and ' + re.escape(name) + r'\s*=\s*box\s*\(fun\b', body[alias.start():])):
                    actions.append({
                        'action': 'split_top_level_alias', 'primary': primary.group(1), 'alias': name,
                        'proof': 'The preceding recursive implementation does not reference its boxed forwarding alias.',
                    })
                    body = body[:alias.start()] + 'let' + body[alias.start()+3:]
        rewritten.append(body)
    return ''.join(rewritten), actions


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('source', type=Path)
    parser.add_argument('target', type=Path)
    parser.add_argument('--split-top-level-aliases', action='store_true')
    args = parser.parse_args()
    source, target = args.source.resolve(), args.target.resolve()
    if source == target or source in target.parents:
        parser.error('Target must be a separate directory outside source')
    planned = []
    for path in sorted(source.iterdir()):
        if path.suffix == '.fs':
            before = path.read_text()
            after, actions = transform(before, args.split_top_level_aliases)
            after, thunk_actions = adapt_typed_thunk_capture(after)
            actions.extend(thunk_actions)
            planned.append((path, before, after, actions))
    # Complete all analyses before writing the scratch copy.
    target.mkdir(parents=True, exist_ok=True)
    for path in source.iterdir():
        if path.is_file():
            shutil.copy2(path, target / path.name)
    report = []
    for path, before, after, actions in planned:
        (target/path.name).write_text(after)
        if actions:
            report.append({'file': path.name, 'sha256_before': digest(before), 'sha256_after': digest(after), 'actions': actions})
    manifest = {'source': str(source), 'target': str(target), 'split_top_level_aliases': args.split_top_level_aliases, 'changes': report}
    (target/'fsharp-adaptation-manifest.json').write_text(json.dumps(manifest, indent=2)+'\n')
    print(json.dumps({'target': str(target), 'files_changed': len(report), 'local_aliases_removed': sum(a['action']=='remove_unused_local_alias' for c in report for a in c['actions']), 'top_aliases_split': sum(a['action']=='split_top_level_alias' for c in report for a in c['actions'])}, indent=2))


if __name__ == '__main__':
    try:
        main()
    except (ValueError, OSError) as exc:
        print('Adaptation failed: ' + str(exc), file=sys.stderr)
        sys.exit(1)
