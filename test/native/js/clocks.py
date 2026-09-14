#!/usr/bin/env python3
"""Check clock units and monotonicity; the 25ms sleeps are not kernel timings."""
import json
from pathlib import Path
import runpy
import shutil
import subprocess
import tempfile

ROOT = Path(__file__).resolve().parents[3]
driver = runpy.run_path(str(ROOT / "bin/js/driver.py"))
environment = driver["environment"]

with tempfile.TemporaryDirectory(prefix="altbak-clock-") as directory:
    work = Path(directory)
    shutil.copyfile(ROOT / "src/Bench.js", work / "Bench.mjs")
    (work / "clock.mjs").write_text('''import assert from 'node:assert/strict';
import {benchNow, opaque} from './Bench.mjs';
const a = benchNow();
for (let i=0, previous=a; i<1000; i++) {
  const now = benchNow(); assert(now >= previous); previous = now;
}
await new Promise(resolve => setTimeout(resolve, 25));
assert(benchNow()-a >= 15000 && benchNow()-a < 2000000);
const f = () => 1; assert.equal(opaque(f)(), f);
console.log('JS clock and opaque passed');
''')
    subprocess.run(["node", "clock.mjs"], cwd=work, check=True)
    package, compatibility = driver["scheme_tool"](work)
    (work / "clock.ss").write_text(f'''(load {json.dumps(str(ROOT / "src/Bench.ss"))})
(import (except (chezscheme) opaque) (Bench foreign)
        (prefix (purescm pstring) ps:))
(define a (benchNow))
(do ([i 0 (+ i 1)] [last a (benchNow)]) ((= i 1000))
  (unless (>= (benchNow) last) (error 'clock "went backwards")))
(sleep (make-time 'time-duration 25000000 0))
(let ([elapsed (- (benchNow) a)])
  (unless (and (>= elapsed 15000) (< elapsed 2000000))
    (error 'clock "wrong microsecond scale" elapsed)))
(unless (= ((opaque 17)) 17) (error 'opaque "changed input"))
(unless (string=? (ps:pstring->string (ps:pstring-upcase (ps:string->pstring "straße"))) "STRASSE")
  (error 'icu "uppercase binding"))
(unless (string=? (ps:pstring->string (ps:pstring-downcase (ps:string->pstring "HELLO"))) "hello")
  (error 'icu "lowercase binding"))
(display "Chez clock and opaque passed\\n")
''')
    subprocess.run(["chez", "--libdirs", str(package / "lib"),
                    "--script", "clock.ss"], cwd=work, env=environment("scm"), check=True)
    shutil.copyfile(ROOT / "src/Bench.erl", work / "bench@foreign.erl")
    subprocess.run(["erlc", "+warnings_as_errors", "bench@foreign.erl"], cwd=work, check=True)
    subprocess.run(["erl", "+S", "1:1", "-noshell", "-pa", str(work), "-eval",
                    "A=(bench@foreign:benchNow())(), timer:sleep(25), "
                    "B=(bench@foreign:benchNow())(), true=B-A>=15000, true=B-A<2000000, "
                    "17=(bench@foreign:opaque(17))(), "
                    'io:format("Erlang clock and opaque passed~n"), halt().'], cwd=work, check=True)
