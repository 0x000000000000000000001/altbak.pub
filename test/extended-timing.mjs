// Instrument the generated harness, using trivial actions rather than benchmarks.
// Usage: node test/extended-timing.mjs /absolute/generated/output [...]
import assert from "node:assert/strict";
import path from "node:path";
import { pathToFileURL } from "node:url";

assert.ok(process.argv.length > 2, "Pass at least one generated module directory");
for (const output of process.argv.slice(2)) {
  const load = (name, file = "index.js") => import(pathToFileURL(path.resolve(output, name, file)));
  const [bench, foreign, aff] = await Promise.all([
    load("Bench.Extended"), load("Bench.Extended", "foreign.js"), load("Effect.Aff")]);
  const runAff = action => new Promise((resolve, reject) => {
    aff.runAff(result => () => {
      const tag = result.tag ?? result.constructor.name;
      const value = result.tag ? result._1 : result.value0;
      if (tag === "Right") resolve(value); else reject(value);
    })(action)();
  });

  globalThis.altbakExtendedResult = "initial";
  const deferredMismatch = foreign.consumeResult("expected")("incorrect");
  assert.equal(globalThis.altbakExtendedResult, "initial");
  assert.throws(deferredMismatch, /Unstable extended benchmark result/);
  assert.equal(globalThis.altbakExtendedResult, "incorrect");

  for (const kind of ["sync", "aff"]) {
    for (const mismatch of [false, true]) {
      const events = [];
      let calls = 0, ticks = 0, sink = "initial";
      const clockDescriptor = Object.getOwnPropertyDescriptor(performance, "now");
      const sinkDescriptor = Object.getOwnPropertyDescriptor(globalThis, "altbakExtendedResult");
      const log = console.log;
      Object.defineProperty(performance, "now", { configurable: true, value: () => {
        events.push("clock"); return ticks++;
      } });
      Object.defineProperty(globalThis, "altbakExtendedResult", { configurable: true,
        get: () => sink, set: value => { events.push("sink"); sink = value; } });
      console.log = () => {};
      try {
        const act = () => {
          events.push("act"); calls++;
          return mismatch && calls === 4 ? "incorrect" : "expected";
        };
        if (kind === "sync") {
          const action = bench.runBenchSync(() => {})(act);
          assert.deepEqual(events, [], "Constructing Effect must not execute it");
          if (mismatch) assert.throws(action, /Unstable extended benchmark result/);
          else assert.equal(action(), 1000);
        } else {
          const action = bench.runBenchAff(() => {})(aff.monadEffectAff.liftEffect(act));
          assert.deepEqual(events, [], "Constructing Aff must not execute it");
          if (mismatch) await assert.rejects(runAff(action), /Unstable extended benchmark result/);
          else assert.equal(await runAff(action), 1000);
        }
        const warmups = Array.from({ length: 3 }, () => ["act", "sink"]).flat();
        const measured = mismatch ? ["clock", "act", "sink"]
          : Array.from({ length: 10 }, () => ["clock", "act", "sink", "clock"]).flat();
        assert.deepEqual(events, [...warmups, ...measured]);
        assert.equal(calls, mismatch ? 4 : 13);
        assert.equal(sink, mismatch ? "incorrect" : "expected");
      } finally {
        console.log = log;
        if (clockDescriptor) Object.defineProperty(performance, "now", clockDescriptor);
        else delete performance.now;
        Object.defineProperty(globalThis, "altbakExtendedResult", sinkDescriptor);
      }
    }
  }
  console.log(JSON.stringify({ output, suspendedEffects: true, syncAndAff: true,
    calls: 13, warmupConsumptions: 3, timedConsumptions: 10, timedMismatchRejected: true }));
}
