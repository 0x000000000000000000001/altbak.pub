// Verify marshalling and suspension; this performs no benchmark measurement.
import { pathToFileURL } from "node:url";
import { resolve } from "node:path";
import assert from "node:assert/strict";

const { makeMarshal } = await import(pathToFileURL(resolve(process.argv[2])));
let calls = 0;
const effect = { perform: () => { calls++; return { int: 7 }; } };
const runtime = {
  boxInt: (int) => ({ int }),
  unboxInt: (value) => {
    assert.ok(Number.isInteger(value.int), "Effect closure was treated as an Int");
    return value.int;
  },
  applyClo: (closure, unit) => {
    assert.equal(unit.int, 0);
    return closure.perform();
  },
};
const { wrap } = makeMarshal(runtime);
const imported = wrap(count => expected => act => () => {
  assert.equal(calls, 0, "Effect must remain suspended while marshalling arguments");
  for (let i = 0; i < count; i++) assert.equal(act(0), expected);
  return 1.25;
}, { params: ["i", "i", { fn: ["i", "i"] }], result: { eff: "f" } });
assert.equal(imported(3, 7, effect), 1.25);
assert.equal(calls, 3);
console.log("PASS Wasm Effect FFI: suspended callback, repeated invocation, numeric result");
