module Bench_FFI

open System

let benchNow (_unit: obj) =
    let ticks = float (System.Diagnostics.Stopwatch.GetTimestamp())
    let freq = float System.Diagnostics.Stopwatch.Frequency
    box ((ticks / freq) * 1000000.0)

let opaque (a: obj) =
    box (fun (_: obj) -> a)

let formatNumber (a: obj) =
    let num = unbox<float> a
    box (num.ToString("0.00"))
