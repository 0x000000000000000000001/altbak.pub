"""Instrument a copy of the regenerated program; do not time this variant."""
from pathlib import Path
import shutil

audit = Path(__file__).resolve().parent
source = audit.parent.parent / "output/Main"
target = audit / "dispatch-profile"
target.mkdir(exist_ok=True)
for file in source.iterdir():
    if file.is_file() and file.suffix in {".fs", ".cs", ".fsproj", ".csproj", ".props"}:
        shutil.copy2(file, target / file.name)

prelude = target / "Sharpurs_Prelude.fs"
code = prelude.read_text()
replacements = {
    "let sharpurs_apply (func: obj)":
        "let mutable sharpurs_directCalls = 0L\nlet mutable sharpurs_reflectedCalls = 0L\n\nlet sharpurs_apply (func: obj)",
    "    | :? (obj -> obj) as invoke ->\n":
        "    | :? (obj -> obj) as invoke ->\n        sharpurs_directCalls <- sharpurs_directCalls + 1L\n",
    "    | _ ->\n        let method = func.GetType()":
        "    | _ ->\n        sharpurs_reflectedCalls <- sharpurs_reflectedCalls + 1L\n        let method = func.GetType()",
}
for old, new in replacements.items():
    assert code.count(old) == 1, old
    code = code.replace(old, new)
prelude.write_text(code)
(target / "EntryPoint.fs").write_text('''module Sharpurs_EntryPoint
open System
open System.Threading
open System.Text.Json

[<EntryPoint>]
let main argv =
    let mutable failure: exn option = None
    let worker = Thread(ThreadStart(fun () ->
        try
            for name, action, expected in [
                "Polymorphism", Test_Polymorphism_act, "10000000"
                "RBTree", Test_RBTree_act, "22"
                "LazyEvaluation", Test_LazyEvaluation_act, "1000000"
            ] do
                sharpurs_directCalls <- 0L
                sharpurs_reflectedCalls <- 0L
                let output = string ((unbox<obj -> obj> action) null)
                let direct, reflected = sharpurs_directCalls, sharpurs_reflectedCalls
                if output <> expected then failwithf "%s: unexpected result %s" name output
                Console.WriteLine(JsonSerializer.Serialize({| caseName = name; direct = direct; reflected = reflected; output = output |}))
        with ex -> failure <- Some ex
    ), 1024 * 1024 * 1024)
    worker.Start()
    worker.Join()
    match failure with
    | Some ex -> raise ex
    | None -> 0
''')
print(f"Prepared {target}; normal generated program unchanged.")
