#if BASELINE
#r "before/bin/Release/net8.0/Program.dll"
#else
#if NORMAL
#r "../../run/bak/sharp/output/Main/bin/Release/net8.0/Program.dll"
#else
#r "after/bin/Release/net8.0/Program.dll"
#endif
#endif

open System
open System.Collections.Generic
open Sharpurs_Prelude
open PureScript_Test_LazyEvaluation

let mutable checks = 0
let check label condition =
    if not condition then failwith label
    checks <- checks + 1

let force lazyValue = unbox<int> (sharpurs_apply Test_LazyEvaluation_force lazyValue)
let defer thunk = sharpurs_apply Test_LazyEvaluation_defer (box thunk)
let build depth seed =
    sharpurs_apply (sharpurs_apply Test_LazyEvaluation_buildThunks (box depth)) seed
let runMany count seed =
    unbox<int> (sharpurs_apply (sharpurs_apply Test_LazyEvaluation_runManyTimes (box count)) (box seed))

let rec exceptionCause (error: exn) envelopes =
    match error with
    | :? System.Reflection.TargetInvocationException as wrapper when not (isNull wrapper.InnerException) ->
        exceptionCause wrapper.InnerException (envelopes + 1)
    | cause -> cause, envelopes

// Independent arithmetic oracle: compute in Int64, then explicitly wrap to Int32.
// The tested depths are nonnegative, as required by buildThunks' terminating domain.
let expected (seed: int) (increments: int) =
    let wide = int64 seed + int64 increments
    let normalized =
        if wide > int64 Int32.MaxValue then wide - 4294967296L
        elif wide < int64 Int32.MinValue then wide + 4294967296L
        else wide
    int normalized

for depth in [0; 1; 3; 1000] do
    for value in [Int32.MinValue; Int32.MinValue + 1; -1; 0; 1; Int32.MaxValue - 1; Int32.MaxValue] do
        let mutable calls = 0
        let seed = defer (fun (_: obj) -> calls <- calls + 1; box value)
        let built = build depth seed
        check "Building does not force the seed" (calls = 0)
        if depth = 0 then check "Depth zero returns the original seed" (Object.ReferenceEquals(seed, built))
        for iteration in [1 .. 3] do
            check "Forced value matches independent wrapping arithmetic" (force built = expected value depth)
            check "Each force calls the seed exactly once; no memoization" (calls = iteration)
        let direct = Test_LazyEvaluation_buildThunks_tco (box depth) seed
        check "Direct ABI construction is lazy" (calls = 3)
        check "Direct and curried ABI agree" (force direct = expected value depth)
        check "Direct ABI calls the seed once" (calls = 4)

let partial = sharpurs_apply Test_LazyEvaluation_buildThunks (box 3)
let mutable leftCalls = 0
let mutable rightCalls = 0
let leftSeed = defer (fun (_: obj) -> leftCalls <- leftCalls + 1; box 17)
let rightSeed = defer (fun (_: obj) -> rightCalls <- rightCalls + 1; box -23)
let left = sharpurs_apply partial leftSeed
let right = sharpurs_apply partial rightSeed
check "Reusing a partial application does not force either seed" (leftCalls = 0 && rightCalls = 0)
check "First captured seed is independent" (force left = 20 && leftCalls = 1 && rightCalls = 0)
check "Second captured seed is independent" (force right = -20 && leftCalls = 1 && rightCalls = 1)
check "Reusing the first chain preserves its captured seed" (force left = 20 && leftCalls = 2 && rightCalls = 1)

let mutable changing = 40
let changingSeed = defer (fun (_: obj) -> changing <- changing + 1; box changing)
let changingChain = build 3 changingSeed
check "Building preserves delayed reads" (changing = 40)
check "First force reads the current seed" (force changingChain = 44 && changing = 41)
check "Second force observes the next seed value" (force changingChain = 45 && changing = 42)

for depth in [0; 1; 3; 1000] do
    let sentinel = InvalidOperationException("deferred seed failure")
    let mutable throws = 0
    let failingSeed = defer (fun (_: obj) -> throws <- throws + 1; raise sentinel : obj)
    let failingChain = build depth failingSeed
    check "A failing seed is not evaluated during construction" (throws = 0)
    for iteration in [1 .. 2] do
        let mutable caught = None
        try force failingChain |> ignore with error -> caught <- Some error
        let cause, envelopes =
            match caught with
            | Some error -> exceptionCause error 0
            | None -> failwith "Expected the deferred seed failure"
        check "Forcing preserves the original exception cause identity" (Object.ReferenceEquals(cause, sentinel))
        // Each force calls the public force wrapper and its thunk via sharpurs_apply.
        check "Forcing preserves the generic-application exception boundaries" (envelopes = 2 * (depth + 1))
        check "Repeated forcing re-evaluates the failing seed once" (throws = iteration)

for useDirect in [false; true] do
    let events = ResizeArray<string>()
    let depthArgument () = events.Add("depth"); box 3
    let seedArgument () =
        events.Add("seed")
        defer (fun (_: obj) -> events.Add("force"); box 11)
    let built =
        if useDirect then Test_LazyEvaluation_buildThunks_tco (depthArgument ()) (seedArgument ())
        else sharpurs_apply (sharpurs_apply Test_LazyEvaluation_buildThunks (depthArgument ())) (seedArgument ())
    check "Arguments are evaluated left to right, without early forcing" (List.ofSeq events = ["depth"; "seed"])
    check "Forcing produces the expected value after argument evaluation" (force built = 14)
    check "Only the delayed seed runs while forcing" (List.ofSeq events = ["depth"; "seed"; "force"])

for count, seed in [0, Int32.MaxValue; 1, 7; 3, -11; 1, Int32.MaxValue] do
    check "runManyTimes agrees with independent arithmetic" (runMany count seed = expected seed (count * 1000))
check "Full one-million-thunk workload returns one million" (runMany 1000 0 = 1000000)

printfn "Real generated LazyEvaluation: %d arithmetic, laziness, exception, ABI and workload checks passed" checks
