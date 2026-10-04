#!/usr/bin/env node
// Execute the existing Docker-oriented regression commands on this macOS host.
// This is an explicit test adapter, not a Docker/platform equivalence claim.
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
const root = resolve(dirname(fileURLToPath(import.meta.url)), '../../../..');
const translate = text => text.replaceAll('/var/www/', root + '/');
const args = process.argv.slice(2), env = { ...process.env };
assert.equal(args.shift(), 'exec');
let cwd = process.cwd();
while (args[0]?.startsWith('-')) {
  const option = args.shift(), value = args.shift();
  if (option === '-w') cwd = translate(value);
  else if (option === '-e') { const index = value.indexOf('='); assert(index > 0); env[value.slice(0, index)] = translate(value.slice(index + 1)); }
  else throw new Error('Unsupported local adapter option: ' + option);
}
assert.equal(args.shift(), 'core-api-cli-1');
let timeout = 180000;
if (args[0] === 'timeout') {
  args.shift();
  if (args[0] === '-k') { args.shift(); args.shift(); }
  const duration = args.shift(); assert(/^\d+(\.\d+)?s$/.test(duration)); timeout = parseFloat(duration) * 1000;
}
const command = translate(args.shift()), actualArgs = args.map(translate);
const result = spawnSync(command, actualArgs, { cwd, env, timeout, maxBuffer: 64 * 1024 * 1024 });
const logs = process.env.NIGHT_HOST_ADAPTER_LOGS;
assert(logs, 'Record the actual local commands and platform'); mkdirSync(logs, { recursive: true });
writeFileSync(join(logs, `${process.pid}.json`), JSON.stringify({ platform: process.platform, architecture: process.arch,
  original: process.argv.slice(2), command: [command, ...actualArgs], cwd, timeout_ms: timeout,
  status: result.status, signal: result.signal, error: result.error?.message }, null, 2));
if (result.stdout) writeFileSync(1, result.stdout);
if (result.stderr) writeFileSync(2, result.stderr);
if (result.error) writeFileSync(2, result.error.message + '\n');
process.exit(result.status ?? (result.error?.code === 'ETIMEDOUT' ? 124 : 1));
