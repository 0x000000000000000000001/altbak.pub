module Bench_Extended_FFI

let mutable private extendedResult : obj = null

let consumeResult (result: obj) =
    box (fun (_: obj) ->
        System.Threading.Volatile.Write(&extendedResult, result)
        box ())
