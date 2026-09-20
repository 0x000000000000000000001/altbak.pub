// Check generated JS/ES workloads without running the calibrated benchmark.
// Usage: node test/extended-workloads.mjs /absolute/path/to/output [other/output]
import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { pathToFileURL } from "node:url";

assert.ok(process.argv.length > 2, "Pass at least one generated module directory");
for (const output of process.argv.slice(2)) {
  const load = name => import(pathToFileURL(path.resolve(output, name, "index.js")));
  const [strings, arrays, files, parallel, delay, aff] = await Promise.all(
    ["Test.StringOps", "Test.STArray", "Test.FileOps", "Test.Parallelism",
      "Test.AffOperations", "Effect.Aff"].map(load));

  // Patch only after module initialization: a cached import-time answer would
  // return the right value while executing none of these operations.
  const replace = String.prototype.replace;
  const split = String.prototype.split;
  let replacements = 0, splits = 0;
  String.prototype.replace = function (...args) {
    replacements++;
    return replace.apply(this, args);
  };
  String.prototype.split = function (...args) {
    splits++;
    return split.apply(this, args);
  };
  let stringValues;
  try { stringValues = [strings.act(), strings.act()]; }
  finally { String.prototype.replace = replace; String.prototype.split = split; }
  assert.deepEqual(stringValues, ["2000", "2000"]);
  assert.equal(replacements, 2000);
  assert.equal(splits, 2000);

  const pop = Array.prototype.pop;
  let pops = 0;
  Array.prototype.pop = function () { pops++; return pop.call(this); };
  let arrayValues;
  try { arrayValues = [arrays.act(), arrays.act()]; }
  finally { Array.prototype.pop = pop; }
  assert.deepEqual(arrayValues, ["10", "10"]);
  assert.equal(pops, 2);
  assert.equal(arrays.sumArray(23), 23);

  const oldCwd = process.cwd();
  const scratch = fs.mkdtempSync(path.join(os.tmpdir(), "altbak-extended-check-"));
  const write = fs.writeFileSync, read = fs.readFileSync;
  let writes = 0, reads = 0, fileValue;
  fs.mkdirSync(path.join(scratch, "var"));
  process.chdir(scratch);
  fs.writeFileSync = (...args) => { writes++; return write(...args); };
  fs.readFileSync = (...args) => { reads++; return read(...args); };
  try {
    fileValue = files.act();
    // Detect discarded or unchecked file contents independently of the nominal run.
    fs.readFileSync = () => "incorrect content";
    assert.equal(files.loopIO(3)(), 0);
  } finally {
    fs.writeFileSync = write;
    fs.readFileSync = read;
    process.chdir(oldCwd);
    fs.rmSync(scratch, { recursive: true });
  }
  assert.equal(fileValue, "10000");
  assert.equal(writes, 10003);
  assert.equal(reads, 10000);

  const runAff = action => new Promise((resolve, reject) => {
    aff.runAff(result => () => {
      const tag = result.tag ?? result.constructor.name;
      const value = result.tag ? result._1 : result.value0;
      if (tag === "Right") resolve(value);
      else reject(value);
    })(action)();
  });
  const start = performance.now();
  assert.equal(await runAff(delay.act), "10");
  assert.ok(performance.now() - start >= 9, "The 10 ms delay must actually run");
  assert.equal(await runAff(parallel.heavyTask(12)), 144);
  assert.equal(parallel.checksum(Array(10).fill(267914296)), 679142946);
  console.log(JSON.stringify({ output, stringCalls: 2, replacements, splits,
    arrayCalls: 2, pops, fileWrites: writes - 3, fileReads: reads,
    delayResult: "10", parallelSmallInput: 12, parallelChecksum: 679142946 }));
}
