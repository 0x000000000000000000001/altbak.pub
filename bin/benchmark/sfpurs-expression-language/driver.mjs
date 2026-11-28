import { readFileSync } from "node:fs";
import {
  compileWithConfig,
  lintWithConfig,
  dumpWithConfig,
} from "../../../../sfpurs/sfpurs-expression-language/targets/js/output/Sfpurs.ExpressionLanguage.Api/index.js";

const corpusPath = process.argv[2];
const iters = Number.parseInt(process.argv[3] || "1000", 10);

const lines = readFileSync(corpusPath, "utf8")
  .split("\n")
  .filter((l) => l.length > 0)
  .map((line) => {
    const [expr, namesCsv, funcsCsv] = line.split("\t");
    return {
      expr,
      names: namesCsv === "-" ? "" : namesCsv.replace(/,/g, "\n"),
      funcs: funcsCsv === "-" ? "-" : funcsCsv.replace(/,/g, "\n"),
    };
  });

function fnv1a(hash, str) {
  let bytes = 0;
  for (let i = 0; i < str.length; i++) {
    const c = str.charCodeAt(i);
    if (c > 0x7f) {
      const buf = Buffer.from(str.slice(i), "utf8");
      for (let j = 0; j < buf.length; j++) {
        hash ^= buf[j];
        hash = Math.imul(hash, 16777619) >>> 0;
      }
      return [hash, bytes + buf.length];
    }
    hash ^= c;
    hash = Math.imul(hash, 16777619) >>> 0;
    bytes++;
  }
  return [hash, bytes];
}

let hash = 2166136261 >>> 0;
let ops = 0;
let totalBytes = 0;

for (let it = 0; it < iters; it++) {
  for (let i = 0; i < lines.length; i++) {
    const w = lines[i];
    const c = compileWithConfig(w.expr)(w.names)(w.funcs)(0);
    const l = lintWithConfig(w.expr)(w.names)(w.funcs)(0);
    const d = dumpWithConfig(w.expr)(w.names)(w.funcs)(0);
    let b;
    [hash, b] = fnv1a(hash, c);
    totalBytes += b;
    [hash, b] = fnv1a(hash, l);
    totalBytes += b;
    [hash, b] = fnv1a(hash, d);
    totalBytes += b;
    ops += 3;
  }
}

console.log(`ops=${ops} bytes=${totalBytes} checksum=${hash}`);
