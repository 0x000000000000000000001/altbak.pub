module Sharpurs_EntryPoint

open System
open System.Diagnostics
open System.Runtime
open System.Runtime.InteropServices
open System.Text.Json
open System.Threading

let emit value = Console.WriteLine(JsonSerializer.Serialize(value))
let warmupCount = 3
let globalWarmupCount = 3
let measurementCount = 10
let stackSizeBytes = 1024 * 1024 * 1024

let validate (value: obj) =
    let actual = unbox<string> value
    if actual <> "100000" then failwithf "TCO: expected 100000, got %A" actual
    actual

let measure () =
    // Resolve the real action before timing. Its opaque input, public wrapper,
    // effects and conversion to String are all still part of each invocation.
    let act = unbox<obj -> obj> Test_TCO_act
    let globalWarmup = unbox<obj -> obj> App_warmup
    for _ in 1 .. globalWarmupCount do globalWarmup null |> ignore
    let warmups = Array.init warmupCount (fun _ -> act null |> validate)
    let measurements = Array.init measurementCount (fun index ->
        let gc0Before = GC.CollectionCount(0)
        let gc1Before = GC.CollectionCount(1)
        let gc2Before = GC.CollectionCount(2)
        let bytesBefore = GC.GetAllocatedBytesForCurrentThread()
        let start = Stopwatch.GetTimestamp()
        let result = act null
        let finish = Stopwatch.GetTimestamp()
        let allocatedBytes = GC.GetAllocatedBytesForCurrentThread() - bytesBefore
        let gc0 = GC.CollectionCount(0) - gc0Before
        let gc1 = GC.CollectionCount(1) - gc1Before
        let gc2 = GC.CollectionCount(2) - gc2Before
        // Validation, record allocation and serialization are outside timings
        // and allocation snapshots. No explicit GC or concurrent worker runs.
        let output = validate result
        {| kind = "measurement"; iteration = index + 1; output = output
           elapsedUs = float (finish - start) * 1000000.0 / float Stopwatch.Frequency
           allocatedBytes = allocatedBytes; gc0 = gc0; gc1 = gc1; gc2 = gc2 |})
    emit {| kind = "metadata"; runtime = RuntimeInformation.FrameworkDescription
            architecture = string RuntimeInformation.ProcessArchitecture
            os = RuntimeInformation.OSDescription; serverGc = GCSettings.IsServerGC
            gcLatency = string GCSettings.LatencyMode; stopwatchFrequency = Stopwatch.Frequency
            warmups = warmupCount; globalWarmups = globalWarmupCount
            measurements = measurementCount; stackSizeBytes = stackSizeBytes
            allocationScope = "current worker thread"; gcScope = "process"; forcedGc = false
            action = "Test.TCO.act"; input = 100000; pid = Environment.ProcessId
            tieredCompilationOverride = Environment.GetEnvironmentVariable("DOTNET_TieredCompilation")
            tieredPgoOverride = Environment.GetEnvironmentVariable("DOTNET_TieredPGO") |}
    warmups |> Array.iteri (fun index output -> emit {| kind = "warmup"; iteration = index + 1; output = output |})
    measurements |> Array.iter emit
    let times = measurements |> Array.map (fun item -> item.elapsedUs) |> Array.sort
    emit {| kind = "summary"; bestUs = times.[0]; medianUs = (times.[4] + times.[5]) / 2.0
            minAllocatedBytes = measurements |> Array.map (fun item -> item.allocatedBytes) |> Array.min
            maxAllocatedBytes = measurements |> Array.map (fun item -> item.allocatedBytes) |> Array.max
            gc0 = measurements |> Array.sumBy (fun item -> item.gc0)
            gc1 = measurements |> Array.sumBy (fun item -> item.gc1)
            gc2 = measurements |> Array.sumBy (fun item -> item.gc2); output = "100000" |}

[<EntryPoint>]
let main argv =
    let mutable failure: exn option = None
    let worker = Thread(ThreadStart(fun () ->
        try
            if argv = [| "--official" |] then
                (unbox<obj -> obj> App_main) null |> ignore
            else measure ()
        with error -> failure <- Some error), stackSizeBytes)
    worker.Start()
    worker.Join()
    match failure with
    | Some error -> eprintfn "%O" error; 1
    | None -> 0
