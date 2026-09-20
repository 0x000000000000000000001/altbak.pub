module Bench_Extended_FFI

let mutable private extendedResult : obj = null

let consumeResult (expected: obj) (result: obj) =
    box (fun (_: obj) ->
        System.Threading.Volatile.Write(&extendedResult, result)
        if unbox<string> result <> unbox<string> expected then
            failwith "Unstable extended benchmark result"
        box ())
