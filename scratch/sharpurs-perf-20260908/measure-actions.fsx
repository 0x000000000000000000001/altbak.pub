// Usage: dotnet fsi --exec measure-actions.fsx /absolute/path/Program.dll [Polymorphism RBTree LazyEvaluation]
// Run one Program.dll per process. This measures original generated actions at their original sizes.
// The cached MethodInfo.Invoke adds one fixed reflection call per action, outside inner loops.
open System
open System.Diagnostics
open System.IO
open System.Reflection
open System.Text.Json
open System.Threading

let emit value = Console.WriteLine(JsonSerializer.Serialize(value))

let args = fsi.CommandLineArgs |> Array.skip 1
if args.Length = 0 then
    failwith "Usage: dotnet fsi --exec measure-actions.fsx /absolute/path/Program.dll [Polymorphism RBTree LazyEvaluation]"
if not (Path.IsPathFullyQualified args.[0]) then
    failwith "Program.dll must be an absolute path."

let assemblyPath = Path.GetFullPath args.[0]
if not (File.Exists assemblyPath) then failwithf "Assembly not found: %s" assemblyPath
let assemblyDir = Path.GetDirectoryName assemblyPath
let allCases = [| "Polymorphism", "10000000"; "RBTree", "22"; "LazyEvaluation", "1000000" |]
let requestedCases = args |> Array.skip 1
let selectedCases =
    if requestedCases.Length = 0 then allCases
    else
        requestedCases
        |> Array.map (fun name ->
            allCases
            |> Array.tryFind (fun (candidate, _) -> String.Equals(name, candidate, StringComparison.OrdinalIgnoreCase))
            |> Option.defaultWith (fun () -> failwithf "Unknown case: %s. Expected Polymorphism, RBTree, or LazyEvaluation." name))
        |> Array.distinct

// Resolve generated dependencies beside the selected Program.dll, including FFI.CSharp.dll.
AppDomain.CurrentDomain.add_AssemblyResolve(ResolveEventHandler(fun _ resolveArgs ->
    let requested = AssemblyName(resolveArgs.Name)
    let localPath = Path.Combine(assemblyDir, requested.Name + ".dll")
    if File.Exists localPath then Assembly.LoadFrom localPath else null))
let ffiPath = Path.Combine(assemblyDir, "FFI.CSharp.dll")
if not (File.Exists ffiPath) then failwithf "Generated dependency not found: %s" ffiPath
Assembly.LoadFrom ffiPath |> ignore
let program = Assembly.LoadFrom assemblyPath

let warmupCount = 3
let measureCount = 5
let stackSizeBytes = 1024 * 1024 * 1024

let resolveAction caseName =
    let moduleName = "PureScript_Test_" + caseName
    let propertyName = "Test_" + caseName + "_act"
    let moduleType = program.GetType(moduleName, true)
    let property = moduleType.GetProperty(propertyName, BindingFlags.Public ||| BindingFlags.Static)
    if isNull property then failwithf "Property not found: %s.%s" moduleName propertyName
    let action = property.GetValue(null)
    if isNull action then failwithf "Action is null: %s" propertyName
    let invokeMethod = action.GetType().GetMethod("Invoke", BindingFlags.Public ||| BindingFlags.Instance, null, [| typeof<obj> |], null)
    if isNull invokeMethod then failwithf "Invoke(obj) not found for %s" propertyName
    let invocationArgs: obj array = [| null |]
    fun () -> invokeMethod.Invoke(action, invocationArgs)

let validate caseName expected (result: obj) =
    let actual = string result
    if actual <> expected then
        failwithf "%s returned %A; expected %A" caseName actual expected
    actual

let run () =
    emit {| kind = "metadata"; assembly = assemblyPath; runtime = System.Runtime.InteropServices.RuntimeInformation.FrameworkDescription;
            stopwatchFrequency = Stopwatch.Frequency; warmups = warmupCount; measurements = measureCount;
            stackSizeBytes = stackSizeBytes; cases = selectedCases |> Array.map fst;
            allocationScope = "current worker thread"; gcScope = "process"; forcedGc = false |}
    for (caseName, expected) in selectedCases do
        // Property lookup and static initialization happen before all timings and allocation snapshots.
        let invokeAction = resolveAction caseName
        for iteration in 1 .. warmupCount do
            let output = invokeAction () |> validate caseName expected
            emit {| kind = "warmup"; caseName = caseName; iteration = iteration; output = output |}

        let measurements =
            [| for iteration in 1 .. measureCount do
                let gc0Before = GC.CollectionCount(0)
                let gc1Before = GC.CollectionCount(1)
                let gc2Before = GC.CollectionCount(2)
                let bytesBefore = GC.GetAllocatedBytesForCurrentThread()
                let ticksBefore = Stopwatch.GetTimestamp()
                let result = invokeAction ()
                let ticksAfter = Stopwatch.GetTimestamp()
                let allocatedBytes = GC.GetAllocatedBytesForCurrentThread() - bytesBefore
                let gc0 = GC.CollectionCount(0) - gc0Before
                let gc1 = GC.CollectionCount(1) - gc1Before
                let gc2 = GC.CollectionCount(2) - gc2Before
                let elapsedUs = float (ticksAfter - ticksBefore) * 1000000.0 / float Stopwatch.Frequency
                // Validation and JSON serialization stay outside the measured interval.
                let output = validate caseName expected result
                emit {| kind = "measurement"; caseName = caseName; iteration = iteration; elapsedUs = elapsedUs;
                        allocatedBytes = allocatedBytes; gc0 = gc0; gc1 = gc1; gc2 = gc2; output = output |}
                yield elapsedUs, allocatedBytes, gc0, gc1, gc2 |]
        let times = measurements |> Array.map (fun (elapsedUs, _, _, _, _) -> elapsedUs) |> Array.sort
        let allocations = measurements |> Array.map (fun (_, allocatedBytes, _, _, _) -> allocatedBytes)
        emit {| kind = "summary"; caseName = caseName; bestUs = times.[0]; medianUs = times.[times.Length / 2];
                minAllocatedBytes = Array.min allocations; maxAllocatedBytes = Array.max allocations;
                gc0 = measurements |> Array.sumBy (fun (_, _, count, _, _) -> count);
                gc1 = measurements |> Array.sumBy (fun (_, _, _, count, _) -> count);
                gc2 = measurements |> Array.sumBy (fun (_, _, _, _, count) -> count); output = expected |}

// Match the generated EntryPoint's 1 GiB stack: original RBTree/Lazy code can be deeply recursive.
let mutable failure: exn option = None
let worker = Thread(ThreadStart(fun () ->
    try run ()
    with error -> failure <- Some error), stackSizeBytes)
worker.Start()
worker.Join()
match failure with
| Some error -> raise error
| None -> ()
