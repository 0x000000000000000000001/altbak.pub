// Compare the complete frozen declaration corpus to the pre-prime scanner oracle.
import assert from 'node:assert/strict';
import { constants, copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { pathToFileURL } from 'node:url';
import { hash, writeJson } from '../gopurs-aff/common.mjs';
import { foreignTypeForwards, foreignUnboundTypes } from '../../../../purust/purust/src/Purust/ForeignTypes.js';

const archive = resolve(process.argv[2]), directory = resolve(process.argv[3]);
assert(!existsSync(directory)); mkdirSync(directory, { recursive: true });
const previous = await import(pathToFileURL(join(archive, 'diagnostics/http-primed-before/src/Purust/ForeignTypes.js')));
const results = ['libraries-final-results.json', 'b8x-results.json'].flatMap(name =>
  JSON.parse(readFileSync(join(archive, 'revision1', name))).results);
assert.equal(results.length, 58);
const seen = new Set(), rows = [];
for (const row of results) {
  const definition = JSON.parse(readFileSync(row.definition));
  const files = new Map(definition.input_manifest.map(file => [file.path, file]));
  for (const file of definition.input_manifest.filter(file => file.path.endsWith('.purs'))) {
    const native = files.get(file.path.replace(/\.purs$/, '.rs'));
    const key = file.sha256 + (native?.sha256 ?? ''); if (seen.has(key)) continue; seen.add(key);
    const path = join(definition.input, file.path), source = readFileSync(path, 'utf8');
    assert.equal(hash(source), file.sha256);
    const old = previous.foreignUnboundTypes(source)(''), current = foreignUnboundTypes(source)('');
    const primedHttp = /^foreign\s+import\s+data\s+HttpServer'\s*(?:::|∷)/m.test(source);
    const expected = primedHttp ? old.map(name => name === 'HttpServer' ? "HttpServer'" : name) : old;
    assert.deepEqual(current, expected, path);
    const rust = native ? readFileSync(join(definition.input, native.path), 'utf8') : '';
    if (native) assert.equal(hash(rust), native.sha256);
    if (!primedHttp) {
      assert.deepEqual(foreignUnboundTypes(source)(rust), previous.foreignUnboundTypes(source)(rust), path);
      assert.equal(foreignTypeForwards(source)(rust), previous.foreignTypeForwards(source)(rust), path);
    }
    if (!current.length) continue;
    const destination = join(directory, key + '.purs'); copyFileSync(path, destination, constants.COPYFILE_FICLONE);
    rows.push({ original: path, source_sha256: file.sha256, ffi_sha256: native?.sha256 ?? null,
      declarations: current, primed_http: primedHttp, unicode_kind: /^foreign\s+import\s+data\s+[^\n]*∷/m.test(source) });
  }
}
assert(rows.some(row => row.primed_http)); assert(rows.some(row => row.unicode_kind));
const report = { status: 'passed', unique_source_ffi_pairs: seen.size, declaration_sources: rows.length,
  declarations: rows.reduce((sum, row) => sum + row.declarations.length, 0),
  unicode_sources: rows.filter(row => row.unicode_kind).length, rows };
writeJson(join(directory, 'corpus.json'), report);
console.log(JSON.stringify({ ...report, rows: undefined }, null, 2));
