// Exercise the compiled converter on the actual optimized binding captured in 2.1.
import assert from "node:assert/strict";
import { readFileSync, writeFileSync } from "node:fs";
import { dirname, resolve } from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";

const audit = dirname(fileURLToPath(import.meta.url));
const output = resolve(audit, "../../..", "sharpurs/sharpurs/output");
const load = (name) => import(pathToFileURL(resolve(output, name, "index.js")));
const [Syntax, CoreFn, DataMaybe, DataTuple, Kernel] = await Promise.all([
  "PureScript.Backend.Optimizer.Syntax", "PureScript.Backend.Optimizer.CoreFn",
  "Data.Maybe", "Data.Tuple", "Sharpurs.IntKernel",
].map(load));
const namespaces = { Syntax, CoreFn, DataMaybe, DataTuple };
const input = JSON.parse(readFileSync(resolve(audit, "../sharpurs-tco-ir-20260908/deepTailRec.optimized.json"), "utf8"), (_key, value) => {
  if (value && typeof value === "object" && value.$tag) {
    const [module, name] = value.$tag.split("$");
    const constructor = namespaces[module]?.[name];
    assert.equal(typeof constructor, "function", value.$tag);
    const revived = Object.assign(Object.create(constructor.prototype), value);
    delete revived.$tag;
    return revived;
  }
  return value;
});
const name = new CoreFn.Qualified(new DataMaybe.Just(input.module), input.ident);
const result = Kernel.fromBinding(name)(input.expression);
assert.ok(result instanceof DataMaybe.Just, "Real optimized binding must be accepted");
assert.equal(result.value0.args.length, 2);
const encoded = JSON.stringify(result.value0, (_key, value) => {
  if (value && typeof value === "object" && !Array.isArray(value) && value.constructor.name !== "Object") {
    return { $tag: value.constructor.name, ...value };
  }
  return value;
}, 2) + "\n";
writeFileSync(resolve(audit, "deepTailRec.int-kernel.json"), encoded);
console.log("Recorded optimized deepTailRec accepted: 2 Int parameters; typed conditional and terminal recursive call.");
