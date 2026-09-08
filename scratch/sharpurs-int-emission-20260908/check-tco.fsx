// Correctness only: dotnet fsi --exec check-tco.fsx /absolute/path/Program.dll
open System
open System.IO
open System.Reflection

let assemblyPath = Path.GetFullPath(fsi.CommandLineArgs.[1])
let directory = Path.GetDirectoryName assemblyPath
AppDomain.CurrentDomain.add_AssemblyResolve(ResolveEventHandler(fun _ args ->
    let path = Path.Combine(directory, AssemblyName(args.Name).Name + ".dll")
    if File.Exists path then Assembly.LoadFrom path else null))
Assembly.LoadFrom(Path.Combine(directory, "FFI.CSharp.dll")) |> ignore
let program = Assembly.LoadFrom assemblyPath
let tco = program.GetType("PureScript_Test_TCO", true)
let get name = tco.GetProperty(name, BindingFlags.Public ||| BindingFlags.Static).GetValue(null)
let apply (fn: obj) (arg: obj) = (unbox<obj -> obj> fn) arg
let expect label expected actual =
    if expected <> actual then failwithf "%s: expected %A, got %A" label expected actual
    printfn "%s: %A" label actual

// This is the real act, including Bench.opaque, the generated public wrapper,
// the existing generic apply path, and conversion of the answer to String.
expect "Test.TCO.act" "100000" (string (apply (get "Test_TCO_act") null))
let loop = get "Test_TCO_deepTailRec"
let call n acc = unbox<int> (apply (apply loop (box n)) (box acc))
expect "zero iterations / negative accumulator" -7 (call 0 -7)
expect "one million native iterations" 1000000 (call 1000000 0)
let partial = apply loop (box 3)
expect "partial application" 20 (unbox<int> (apply partial (box 17)))
expect "reused partial application" 0 (unbox<int> (apply partial (box -3)))
printfn "Generated TCO: 5 checks passed; no timings collected."
