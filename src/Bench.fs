module Bench_FFI

open System

let mutable private benchmarkInput : obj = null
let mutable private benchmarkResult : obj = null

let benchNow (_unit: obj) =
    let ticks = float (System.Diagnostics.Stopwatch.GetTimestamp())
    let freq = float System.Diagnostics.Stopwatch.Frequency
    box ((ticks / freq) * 1000000.0)

let opaque (a: obj) =
    box (fun (_: obj) ->
        System.Threading.Volatile.Write(&benchmarkInput, a)
        System.Threading.Volatile.Read(&benchmarkInput))

let formatNumber (a: obj) =
    let num = unbox<float> a
    box (num.ToString("0.000000", System.Globalization.CultureInfo.InvariantCulture))

let measureBatch (iterations: obj) (expected: obj) (action: obj) =
    box (fun (_: obj) ->
        let act = unbox<obj -> obj> action
        let mutable result : obj = null
        let start = System.Diagnostics.Stopwatch.GetTimestamp()
        for _ in 1 .. unbox<int> iterations do
            result <- act null
            System.Threading.Volatile.Write(&benchmarkResult, result)
        let elapsed = float (System.Diagnostics.Stopwatch.GetTimestamp() - start) * 1000000.0 / float System.Diagnostics.Stopwatch.Frequency
        if unbox<int> result <> unbox<int> expected then failwith "Unstable benchmark result"
        box elapsed)
