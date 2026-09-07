#!/usr/bin/env node
import { readFile } from "node:fs/promises";

// Timing, output, and the optimizer barrier may call the host. The measured
// algorithms and their library operations must remain inside WebAssembly.
const allowed = new Map([
  ["Bench", new Set(["formatNumber", "benchNow", "opaque"])],
  ["Effect.Console", new Set(["log"])],
]);

if (process.argv.length !== 3) {
  console.error("Usage: node bin/wasm/check-imports.mjs <module.wasm>");
  process.exitCode = 2;
} else {
  try {
    const bytes = await readFile(process.argv[2]);
    const module = await WebAssembly.compile(bytes);
    const imports = WebAssembly.Module.imports(module);
    const unexpected = imports.filter(
      ({ module, name, kind }) =>
        kind !== "function" || !allowed.get(module)?.has(name),
    );
    if (unexpected.length > 0) {
      throw new Error(
        "Unexpected host imports: " +
          unexpected.map(({ module, name, kind }) => `${module}.${name} (${kind})`).join(", "),
      );
    }
    console.log(
      `[wasm] Import audit passed: ${imports.length} host functions, ${bytes.length} bytes.`,
    );
  } catch (error) {
    console.error(`[wasm] Import audit failed: ${error.message}`);
    process.exitCode = 1;
  }
}
