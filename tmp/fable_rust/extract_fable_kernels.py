#!/usr/bin/env python3
"""Slice sharpurs' generated F# declarations without changing their bodies.

This deliberately targets sharpurs' regular generated formatting, not arbitrary
F#. It retains entire recursive binding/type groups, resolves generated global
names, and slices four-space-indented FFI modules by member. It does not analyze
dictionary fields, optimize algorithms, translate C#, or repair Fable output.
F# type checking is still required to validate standard-library/local references.
"""
from __future__ import annotations

import argparse
from collections import defaultdict
from dataclasses import dataclass
import hashlib
import json
from pathlib import Path
import re
import sys
import xml.etree.ElementTree as ET


WRAPPERS = {
    "AstTree": '''let runAstTree (n: int) : int =
    unbox<int> (sharpurs_apply Test_AstTree_eval (sharpurs_apply Test_AstTree_buildTree (box n)))''',
    "Fib": '''let runFib (n: int) : int =
    unbox<int> (sharpurs_apply Test_Fib_fib (box n))''',
    "ListOps": '''let runListOps (n: int) : int =
    unbox<int> (sharpurs_apply Test_ListOps_sumEvens (box n))''',
    "TCO": '''let runTCO (n: int) : int =
    unbox<int> (sharpurs_apply (sharpurs_apply Test_TCO_deepTailRec (box n)) (box 0))''',
    "Records": '''let runRecords (n: int) : int =
    let updated = sharpurs_apply (sharpurs_apply Test_Records_updateRec (box n)) Test_Records_initial
    let b = Map.find "b" (unbox<Map<string, obj>> updated)
    let d = Map.find "d" (unbox<Map<string, obj>> b)
    unbox<int> (Map.find "f" (unbox<Map<string, obj>> d))''',
    "Ackermann": '''let runAckermann (n: int) : int =
    unbox<int> (sharpurs_apply (sharpurs_apply Test_Ackermann_ackermann (box n)) (box 4))''',
    "Church": '''let runChurch (n: int) : int =
    unbox<int> (sharpurs_apply Test_Church_toInt (sharpurs_apply Test_Church_c100k (box n)))''',
    "Primes": '''let runPrimes (n: int) : int =
    unbox<int> (sharpurs_apply Test_Primes_sumList (sharpurs_apply Test_Primes_sieve (sharpurs_apply (sharpurs_apply Test_Primes_range (box 2)) (box n))))''',
    "RBTree": '''let runRBTree (n: int) : int =
    unbox<int> (sharpurs_apply Test_RBTree_depth (sharpurs_apply (sharpurs_apply Test_RBTree_buildTree (box n)) Test_RBTree_E))''',
    "Polymorphism": '''let runPolymorphism (n: int) : int =
    unbox<int> (sharpurs_apply (sharpurs_apply (sharpurs_apply Test_Polymorphism_polyLoop Test_Polymorphism_intMonoidish) (box n)) (box 0))''',
    "StateMonad": '''let runStateMonad (n: int) : int =
    unbox<int> (sharpurs_apply (sharpurs_apply Test_StateMonad_runManyTimes (box n)) (box 0))''',
    "LazyEvaluation": '''let runLazyEvaluation (n: int) : int =
    unbox<int> (sharpurs_apply (sharpurs_apply Test_LazyEvaluation_runManyTimes (box n)) (box 0))''',
    "ArrayOps": '''let runArrayOps (n: int) : int =
    unbox<int> (sharpurs_apply Test_ArrayOps_sumEvens (box n))''',
    "RowToList": '''let runRowToList (_: int) : int =
    let record = box (Map.empty<string, obj> |> Map.add "a" (box 1) |> Map.add "b" (box "two") |> Map.add "c" (box true) |> Map.add "d" (box 4.0) |> Map.add "e" (box "five"))
    unbox<int> (sharpurs_apply (sharpurs_apply (sharpurs_apply Test_RowToList_keys Prim_undefined) Test_RowToList_keysCons1) record)''',
}


def mask_literals(source: str) -> str:
    """Hide comments/string content, retaining line offsets and backtick names."""
    result = list(source)
    i = 0
    while i < len(source):
        start = i
        if source.startswith('//', i):
            i = source.find('\n', i)
            if i < 0:
                i = len(source)
        elif source.startswith('(*', i):
            depth = 1
            i += 2
            while i < len(source) and depth:
                if source.startswith('(*', i):
                    depth += 1
                    i += 2
                elif source.startswith('*)', i):
                    depth -= 1
                    i += 2
                else:
                    i += 1
            if depth:
                raise ValueError('Unclosed F# comment')
        elif source.startswith('@"', i) or source[i] == '"':
            verbatim = source.startswith('@"', i)
            i += 2 if verbatim else 1
            while i < len(source):
                if verbatim and source.startswith('""', i):
                    i += 2
                elif source[i] == '"':
                    i += 1
                    break
                elif not verbatim and source[i] == '\\':
                    i += 2
                else:
                    i += 1
        else:
            i += 1
            continue
        for j in range(start, min(i, len(source))):
            if source[j] != '\n':
                result[j] = ' '
    return ''.join(result)


def tokens(source: str) -> set[str]:
    return set(re.findall(r'\b[A-Za-z_]\w*(?:\.[A-Za-z_]\w*)*', mask_literals(source).replace('``', '')))


def alias_separation_advice(source: str) -> dict | None:
    """Report only: do not modify mutually recursive declaration groups."""
    primary = re.match(r'let rec (\w+_tco)\b', source)
    aliases = list(re.finditer(r'^and (\w+)\b', source, re.M))
    if not primary or not aliases:
        return None
    alias_names = [match.group(1) for match in aliases]
    primary_refs = sorted(tokens(source[:aliases[0].start()]) & set(alias_names))
    simple_alias = (len(aliases) == 1 and bool(re.match(
        r'and ' + re.escape(alias_names[0]) + r'\s*=\s*box\s*\(fun\b',
        source[aliases[0].start():],
    )))
    safe = simple_alias and not primary_refs
    return {
        'primary': primary.group(1), 'aliases': alias_names,
        'primary_references_to_aliases': primary_refs,
        'may_replace_and_with_let': safe,
        'reason': ('One boxed function alias, declared after a recursive implementation which does not reference the alias.'
                   if safe else 'Keep the group intact unless a stronger dependency analysis proves separation valid.'),
        'applied': False,
    }


@dataclass
class Block:
    path: Path
    start: int
    text: str
    names: list[str]
    scope: str | None = None


def declaration_names(body: str, indent: str) -> list[str]:
    names = []
    modifier = r'(?:(?:rec|private|internal|inline|mutable)\s+)*'
    for m in re.finditer(r'^' + indent + r'(?:let|type|and)\s+' + modifier + r'(\w+|``[^`]+``|\(\|[^\n]+?\|\))', mask_literals(body), re.M):
        name = m.group(1).replace('``', '')
        if name.startswith('(|'):
            names.extend(part for part in name[2:-2].split('|') if part != '_')
        else:
            names.append(name)
    if body.lstrip().startswith('type '):
        names.extend(re.findall(r'^\s*\|\s*(\w+)', mask_literals(body), re.M))
    if not names:
        raise ValueError('Cannot identify declaration: ' + body[:100])
    return names


def starts(source: str, indent: str) -> list[int]:
    pattern = r'^' + indent + r'(?:let\s|type\s|module\s+\w+\s*=)'
    positions = []
    for match in re.finditer(pattern, mask_literals(source), re.M):
        start = match.start()
        # Attributes belong to the following declaration, if any.
        while start > 0:
            previous_end = start - 1
            previous_start = source.rfind('\n', 0, previous_end) + 1
            previous = source[previous_start:previous_end]
            if previous.startswith(indent + '[<') and previous.rstrip().endswith('>]'):
                start = previous_start
            else:
                break
        positions.append(start)
    return positions


class Extractor:
    def __init__(self, source: Path):
        self.source = source
        self.blocks: list[Block] = []
        self.symbols: dict[str, list[int]] = defaultdict(list)
        self.headers: dict[str, str] = {}
        self.module_headers: dict[tuple[str, str], str] = {}
        self.order = [n.attrib['Include'] for n in ET.parse(source / 'Program.fsproj').getroot().iter('Compile')]
        for filename in self.order:
            path = source / filename
            if path.suffix != '.fs':
                continue
            text = path.read_text()
            bounds = starts(text, '')
            if not bounds:
                continue
            self.headers[filename] = text[:bounds[0]]
            bounds.append(len(text))
            for a, b in zip(bounds, bounds[1:]):
                body = text[a:b]
                module = re.match(r'module (\w+)\s*=', body)
                if module:
                    scope = module.group(1)
                    members = starts(body, '    ')
                    if not members:
                        continue
                    self.module_headers[(filename, scope)] = body[:members[0]]
                    members.append(len(body))
                    for x, y in zip(members, members[1:]):
                        self.add(Block(path, a+x, body[x:y], declaration_names(body[x:y], '    '), scope))
                else:
                    self.add(Block(path, a, body, declaration_names(body, '')))

    def add(self, block: Block) -> None:
        index = len(self.blocks)
        self.blocks.append(block)
        for name in set(block.names):
            qualified = block.scope + '.' + name if block.scope else name
            self.symbols[qualified].append(index)

    def resolve(self, body: str, scope: str | None = None) -> tuple[set[int], set[str]]:
        found, missing = set(), set()
        for token in tokens(body):
            scoped = scope + '.' + token if scope else token
            name = scoped if scoped in self.symbols else token
            if name in self.symbols:
                found.update(self.symbols[name])
            elif re.match(r'^(?:Data|Control|Test|Type|Record|Prim|Effect|Unsafe|Partial|Bench)_', token):
                missing.add(token)
        return found, missing

    def closure(self, wrappers: str) -> set[int]:
        stack, missing = self.resolve(wrappers)
        seen = set()
        while stack:
            index = stack.pop()
            if index in seen:
                continue
            seen.add(index)
            block = self.blocks[index]
            dependencies, unknown = self.resolve(block.text, block.scope)
            missing.update(unknown)
            stack.update(dependencies - seen)
        if missing:
            raise ValueError('Unresolved generated references: ' + ', '.join(sorted(missing)))
        for index in seen:
            block = self.blocks[index]
            if any(name.endswith(('_act', '_describe')) for name in block.names):
                raise ValueError('Unexpected benchmark effect dependency: ' + ', '.join(block.names))
            if re.search(r'\b(?:FFI\.CSharp|global\.|[A-Z]\w*(?:\.[A-Z]\w*)*\.FFI\.)', mask_literals(block.text)):
                raise ValueError('External C# dependency: ' + str(block.path) + ': ' + ', '.join(block.names))
        return seen

    def emit(self, chosen: set[int], target: Path) -> tuple[list[str], list[dict]]:
        files, manifest = [], []
        for filename in self.order:
            selected = sorted((self.blocks[i] for i in chosen if self.blocks[i].path.name == filename), key=lambda b: b.start)
            if not selected:
                continue
            emitted = self.headers[filename]
            current_scope = None
            for block in selected:
                if block.scope and block.scope != current_scope:
                    emitted += self.module_headers[(filename, block.scope)]
                current_scope = block.scope
                emitted += block.text
                manifest.append({
                    'file': filename, 'scope': block.scope, 'names': block.names,
                    'source_line': block.path.read_text()[:block.start].count('\n') + 1,
                    'sha256': hashlib.sha256(block.text.encode()).hexdigest(),
                    'recursive_alias_advice': alias_separation_advice(block.text),
                })
            (target / filename).write_text(emitted)
            files.append(filename)
        return files, manifest


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('source', type=Path, help='sharpurs output/Main containing Program.fsproj')
    parser.add_argument('target', type=Path, help='Separate scratch project directory')
    parser.add_argument('--benchmark', action='append', choices=list(WRAPPERS), help='Repeat to select kernels; default all 14')
    parser.add_argument('--target-framework', default='net8.0')
    args = parser.parse_args()
    source, target = args.source.resolve(), args.target.resolve()
    if source == target or source in target.parents:
        parser.error('Target must be outside the generated source directory')
    selected = args.benchmark or list(WRAPPERS)
    wrapper = 'module BenchFableKernels\n\n' + '\n\n'.join(WRAPPERS[name] for name in selected) + '\n'
    extractor = Extractor(source)
    chosen = extractor.closure(wrapper)
    target.mkdir(parents=True, exist_ok=True)
    files, declarations = extractor.emit(chosen, target)
    (target / 'BenchFableKernels.fs').write_text(wrapper)
    project = ET.Element('Project', Sdk='Microsoft.NET.Sdk')
    properties = ET.SubElement(project, 'PropertyGroup')
    for key, value in [('TargetFramework', args.target_framework), ('OutputType', 'Library'), ('GenerateAssemblyInfo', 'false'), ('EnableDefaultCompileItems', 'false')]:
        ET.SubElement(properties, key).text = value
    includes = ET.SubElement(project, 'ItemGroup')
    for filename in files + ['BenchFableKernels.fs']:
        ET.SubElement(includes, 'Compile', Include=filename)
    ET.indent(project)
    ET.ElementTree(project).write(target / 'Kernels.fsproj', encoding='unicode')
    manifest = {
        'source': str(source), 'benchmarks': selected, 'files': files,
        'declarations': declarations,
        'limitations': [
            'Declaration-level reachability only: dictionary fields and recursive groups remain intact.',
            'The parser targets the regular sharpurs generated F# layout, not arbitrary F#.',
            'Generated global names are checked; F# type checking must validate local and .NET references.',
            'Generated declaration bodies, including runtime reflection/exception code, are preserved verbatim.',
            'Wrapper dictionaries are the concrete Int Monoidish and five-field RowToList instances used by these benchmark acts.',
        ],
    }
    (target / 'extraction-manifest.json').write_text(json.dumps(manifest, indent=2) + '\n')
    print(json.dumps({'target': str(target), 'project': str(target / 'Kernels.fsproj'), 'benchmarks': selected, 'source_files': len(files), 'declaration_groups': len(declarations), 'ffi_members': [d['scope'] + '.' + n for d in declarations if d['scope'] for n in d['names']]}, indent=2))


if __name__ == '__main__':
    try:
        main()
    except (ValueError, OSError, ET.ParseError) as error:
        print('Extraction failed: ' + str(error), file=sys.stderr)
        sys.exit(1)
