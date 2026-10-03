// Explicit Cargo/Rust flags for isolated, reproducible host experiments.
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { existsSync, readFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { hash } from '../gopurs-aff/common.mjs';

export function readProfile(argument) {
  if (!argument || ['false', 'thin'].includes(argument)) return { lto: argument === 'thin' ? 'thin' : false };
  return JSON.parse(readFileSync(resolve(argument), 'utf8'));
}

export function cargoProfile(configuration, environment, rust, target) {
  assert([false, 'thin'].includes(configuration.lto));
  const profile = { opt_level: 3, debug: false, lto: configuration.lto, threaded: true, allocator: 'mimalloc' };
  const env = { ...environment };
  const args = ['--release', '--target-dir', target,
    '--config', `profile.release.lto=${JSON.stringify(profile.lto)}`,
    '--config', 'profile.release.opt-level=3', '--config', 'profile.release.debug=false',
    '--manifest-path', join(rust, 'Cargo.toml')];
  if (configuration.codegen_units !== undefined) {
    assert(Number.isInteger(configuration.codegen_units) && configuration.codegen_units > 0);
    args.push('--config', `profile.release.codegen-units=${configuration.codegen_units}`);
    profile.codegen_units = configuration.codegen_units;
  }
  const flags = [...(configuration.rustflags ?? [])];
  assert(flags.every(flag => typeof flag === 'string' && !flag.includes('\x1f')));
  if (configuration.target_cpu) { profile.target_cpu = configuration.target_cpu; flags.push('-C', 'target-cpu=' + configuration.target_cpu); }
  if (flags.length) env.CARGO_ENCODED_RUSTFLAGS = flags.join('\x1f');
  let linker = null;
  if (configuration.linker === 'rust-lld') {
    assert.equal(process.platform, 'darwin');
    const sysroot = execFileSync('rustc', ['--print', 'sysroot'], { env, encoding: 'utf8' }).trim();
    const host = execFileSync('rustc', ['-vV'], { env, encoding: 'utf8' }).match(/^host: (.+)$/m)[1];
    const path = join(sysroot, 'lib/rustlib', host, 'bin/gcc-ld/ld64.lld');
    assert(existsSync(path));
    linker = { path, sha256: hash(readFileSync(path)), driver: 'cc', option: '-fuse-ld=' + path };
    args.push('--bin', 'purust_output', '--', '-C', 'link-arg=' + linker.option);
  } else assert(!configuration.linker);
  return { args: [linker ? 'rustc' : 'build', ...args], env, profile, linker, rustflags: flags };
}
