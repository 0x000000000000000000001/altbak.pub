// Reclaim only the completed experiment's regenerable Cargo cache.
import assert from 'node:assert/strict';
import { existsSync, readFileSync, rmSync, statSync, statfsSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { hash, walk, writeJson } from '../gopurs-aff/common.mjs';
const archive = resolve(process.argv[2]), target = join(archive, 'work/target');
const report = join(archive, 'reclaimed-experiment-target.json'); assert(!existsSync(report));
const selection = JSON.parse(readFileSync(join(archive, 'final-selection.json')));
assert.equal(selection.status, 'passed');
const build = JSON.parse(readFileSync(join(archive, 'candidates', selection.candidate, 'build.json')));
const binary = join(archive, 'candidates', selection.candidate, 'gopurs-rust');
assert.equal(hash(readFileSync(binary)), build.binary_sha256);
assert.equal(hash(readFileSync(join(target, 'release/purust_output'))), build.binary_sha256);
assert(existsSync(join(target, '.rustc_info.json')));
const files = walk(target);
assert(!files.some(path => path.includes('/.purmeta/')));
const available = () => { const fs = statfsSync(archive); return fs.bavail * fs.bsize; };
const result = { status: 'pending', started_at: new Date().toISOString(), target,
  retained_binary: { path: binary, sha256: build.binary_sha256 },
  files: files.length, logical_bytes: files.reduce((total, file) => total + statSync(file).size, 0),
  available_before: available(), reason: 'Completed PGO experiment; binaries, generated sources, profiles and raw diagnostics retained outside this Cargo target.' };
writeJson(report, result);
rmSync(target, { recursive: true });
assert.equal(hash(readFileSync(binary)), build.binary_sha256);
result.available_after = available(); result.status = 'passed'; result.finished_at = new Date().toISOString();
writeJson(report, result); console.log(JSON.stringify(result, null, 2));
