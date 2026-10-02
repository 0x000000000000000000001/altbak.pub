import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync, writeFileSync } from 'node:fs';
import { dirname, join, relative, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { unusedFfi } from './unused-ffi.mjs';

const here = dirname(fileURLToPath(import.meta.url)), htdocs = resolve(here, '../../../..');
const workspace = resolve(process.argv[2]), prepared = JSON.parse(readFileSync(join(workspace, 'prepared.json'), 'utf8'));
const ports = new Map();
const hash = bytes => createHash('sha256').update(bytes).digest('hex');
function walk(directory) {
  if (!existsSync(directory)) return [];
  return readdirSync(directory, { withFileTypes: true }).flatMap(entry => {
    const path = join(directory, entry.name);
    return entry.isDirectory() ? walk(path) : entry.isFile() ? [path] : [];
  });
}
for (const packageName of readdirSync(join(htdocs, 'purust')).filter(name => name.startsWith('purust-')).sort()) {
  const source = join(htdocs, 'purust', packageName, 'src');
  for (const file of walk(source).filter(file => file.endsWith('.rs'))) {
    const module = relative(source, file).slice(0, -3).split('/').join('.');
    ports.set(module, [...(ports.get(module) ?? []), file]);
  }
}
const directory = join(workspace, 'rust-ffi'); mkdirSync(directory, { recursive: true });
const base = join(workspace, 'rust-ffi-originals'); mkdirSync(base, { recursive: true });
const records = [];
const replace = (source, before, after) => {
  assert.equal(source.split(before).length, 2, before); return source.replace(before, after);
};
for (const module of prepared.tast.files) {
  const candidates = ports.get(module.module) ?? [];
  const sourcePackage = module.source.match(/(?:^|\/)(gopurs-[^/]+)\//)?.[1]?.replace(/^gopurs-/, 'purust-');
  const preferred = candidates.filter(path => sourcePackage && path.includes('/' + sourcePackage + '/'));
  const choices = preferred.length ? preferred : candidates;
  assert(choices.length < 2 || new Set(choices.map(path => hash(readFileSync(path)))).size === 1,
    `Ambiguous Rust FFI for ${module.module}: ${choices.join(', ')}`);
  const port = choices[0];
  if (!port) {
    if (module.foreign.length) {
      const source = unusedFfi[module.module];
      assert(source, `No FFI or explicit unused-module trap for ${module.module}`);
      const target = join(directory, module.module + '.rs'); writeFileSync(target, source);
      records.push({ module: module.module, path: target, sha256: hash(source), unused_trap: true, foreign: module.foreign });
    }
    continue;
  }
  const original = readFileSync(port, 'utf8');
  copyFileSync(port, join(base, module.module + '.rs'));
  let source = original, adapter = null;
  switch (module.module) {
    case 'Effect.Aff':
      source = replace(source, `    fn new(value: AffValue) -> Self {
        Self {
            is_left: value.get_isLeft(),
            from_left: value.get_fromLeft(),
            from_right: value.get_fromRight(),
            left: value.get_left(),
            right: value.get_right(),
        }
    }`, '    fn new(_value: AffValue) -> Self { gopurs_aff_util() }');
      source = source.replaceAll('Effect_Aff__makeSupervisedFiber', 'gopurs_base_makeSupervisedFiber')
        .replaceAll('Effect_Aff_makeAff', 'gopurs_base_makeAff');
      adapter = 'aff-compat.rs';
      source += '\n' + readFileSync(join(here, adapter), 'utf8');
      break;
    case 'Effect.AVar':
      for (const name of ['putVar', 'takeVar', 'readVar', 'killVar', 'tryPutVar', 'tryTakeVar', 'tryReadVar', 'status']) {
        source = source.replaceAll('Effect_AVar__' + name, 'gopurs_base_AVar__' + name);
      }
      source = source.replace(/util\.get_(\w+)\(\)/g, (_, field) => `gopurs_avar_field(&util, "${field}")`);
      adapter = 'avar-compat.rs';
      source += '\n' + readFileSync(join(here, adapter), 'utf8');
      break;
    case 'Data.Array.ST':
      source = source.replaceAll('Data_Array_ST_new()', 'Data_Array_ST_newImpl()');
      adapter = 'newImpl alias'; break;
    case 'Control.Monad.ST.Internal':
      source = source.replaceAll('Control_Monad_ST_Internal_for(', 'Control_Monad_ST_Internal_forImpl(')
        .replaceAll('Control_Monad_ST_Internal_new(', 'Control_Monad_ST_Internal_newImpl(');
      adapter = 'forImpl/newImpl aliases'; break;
    case 'Data.Traversable':
      source = replace(source, '    pure: TraversalFunction,\n    function:',
        '    pure: TraversalFunction,\n    _append: TraversalApply,\n    function:');
      adapter = 'explicit array append argument'; break;
    case 'Effect.Ref':
      source += `
pub fn Effect_Ref_modify_(update: purust_core::Func1<crate::UnknownType, crate::UnknownType>, reference: crate::UnknownType) -> crate::UnknownType {
    purust_ref_effect(move || {
        let mut guard = reference.unwrap_class::<PurustRef>().lock();
        *guard = update(guard.clone());
        crate::Value::Unit
    })
}
`;
      adapter = 'atomic modify_'; break;
  }
  const target = join(directory, module.module + '.rs'); writeFileSync(target, source);
  const record = { module: module.module, origin: port, original_sha256: hash(original), path: target, sha256: hash(source), adapter };
  if (existsSync(port + '.cargo.json')) {
    copyFileSync(port + '.cargo.json', target + '.cargo.json');
    record.cargo = { path: target + '.cargo.json', sha256: hash(readFileSync(target + '.cargo.json')) };
  }
  records.push(record);
}
writeFileSync(join(workspace, 'rust-ffi.json'), JSON.stringify(records, null, 2) + '\n');
console.log(JSON.stringify({ rust_ffi: records.filter(record => !record.missing).length,
  adapted: records.filter(record => record.adapter).map(record => record.module),
  missing: records.filter(record => record.missing) }, null, 2));
