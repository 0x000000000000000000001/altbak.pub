#if BASELINE
#r "before/bin/Release/net8.0/Program.dll"
#else
#r "after/bin/Release/net8.0/Program.dll"
#endif

open System
open PureScript_Test_RBTree
open Sharpurs_Prelude

let mutable checks = 0
let check label condition =
    if not condition then failwith label
    checks <- checks + 1

let moduleType = typeof<Test_RBTree_Tree>.Assembly.GetType("PureScript_Test_RBTree", true)
let directMethods =
    moduleType.GetMethods(Reflection.BindingFlags.Public ||| Reflection.BindingFlags.Static)
    |> Array.filter (fun methodInfo -> methodInfo.Name = "Test_RBTree_balance_direct")
check "one public static direct method" (directMethods.Length = 1)
let directMethod = directMethods.[0]
let directParameters = directMethod.GetParameters()
check "direct method has four parameters" (directParameters.Length = 4)
check "direct parameters use System.Object" (directParameters |> Array.forall (fun parameter -> parameter.ParameterType = typeof<obj>))
check "direct method returns System.Object" (directMethod.ReturnType = typeof<obj>)

let applyMethods =
    moduleType.GetMethods(Reflection.BindingFlags.Public ||| Reflection.BindingFlags.Static)
    |> Array.filter (fun methodInfo -> methodInfo.Name = "Test_RBTree_balance_direct_apply")
check "one public static direct application method" (applyMethods.Length = 1)
let applyMethod = applyMethods.[0]
check "direct application method has four object parameters"
    (applyMethod.GetParameters().Length = 4
     && (applyMethod.GetParameters() |> Array.forall (fun parameter -> parameter.ParameterType = typeof<obj>)))
check "direct application method returns System.Object" (applyMethod.ReturnType = typeof<obj>)

let red = Test_RBTree_Rusd_Ctor
let black = Test_RBTree_Busd_Ctor
let empty = Test_RBTree_Eusd_Ctor
let node color left value right = Test_RBTree_Tusd_Ctor(color, left, value, right)
let leaf value = node black empty value empty
let a, b, c, d = leaf 1, leaf 3, leaf 5, leaf 7
let x, y, z = 2, 4, 6

let publicBalance color left value right =
    unbox<Test_RBTree_Tree>
        (sharpurs_apply
            (sharpurs_apply
                (sharpurs_apply
                    (sharpurs_apply Test_RBTree_balance (box color))
                    (box left))
                (box value))
            (box right))

let directBalance color left value right =
    unbox<Test_RBTree_Tree>
        (Test_RBTree_balance_direct (box color) (box left) (box value) (box right))

let directApplyBalance color left value right =
    unbox<Test_RBTree_Tree>
        (Test_RBTree_balance_direct_apply (box color) (box left) (box value) (box right))

let expectedRotation = node red (node black a x b) y (node black c z d)
let checkRotation label actual =
    check (label + ": exact rotated structure") (actual = expectedRotation)
    match actual with
    | Test_RBTree_Tusd_Ctor(_, Test_RBTree_Tusd_Ctor(_, nextA, _, nextB), _, Test_RBTree_Tusd_Ctor(_, nextC, _, nextD)) ->
        for name, previous, current in ["a", a, nextA; "b", b, nextB; "c", c, nextC; "d", d, nextD] do
            check (label + ": shares subtree " + name) (Object.ReferenceEquals(previous, current))
    | _ -> failwith (label + ": missing rotation children")

let rotations = [
    "LL", node red (node red a x b) y c, z, d
    "LR", node red a x (node red b y c), z, d
    "RL", a, x, node red (node red b y c) z d
    "RR", a, x, node red b y (node red c z d)
]

for name, left, value, right in rotations do
    checkRotation (name + " public") (publicBalance black left value right)
    checkRotation (name + " direct") (directBalance black left value right)
    checkRotation (name + " direct application") (directApplyBalance black left value right)

let checkDefault label color left value right actual =
    check (label + ": exact default structure") (actual = node color left value right)
    match actual with
    | Test_RBTree_Tusd_Ctor(_, nextLeft, _, nextRight) ->
        check (label + ": shares left subtree") (Object.ReferenceEquals(left, nextLeft))
        check (label + ": shares right subtree") (Object.ReferenceEquals(right, nextRight))
    | _ -> failwith (label + ": missing default node")

for name, color in ["red", red; "black", black] do
    checkDefault (name + " public") color a y d (publicBalance color a y d)
    checkDefault (name + " direct") color a y d (directBalance color a y d)
    checkDefault (name + " direct application") color a y d (directApplyBalance color a y d)

// Each retained prefix is applied repeatedly, exercising all public ABI stages.
let partial1 = sharpurs_apply Test_RBTree_balance (box black)
let partial1Results =
    rotations
    |> List.map (fun (name, left, value, right) ->
        let result =
            unbox<Test_RBTree_Tree>
                (sharpurs_apply (sharpurs_apply (sharpurs_apply partial1 (box left)) (box value)) (box right))
        checkRotation (name + " reused one-argument prefix") result
        result)
check "one-argument prefix results stay unchanged" (partial1Results |> List.forall ((=) expectedRotation))

let partial2 = sharpurs_apply partial1 (box a)
let partial2Results =
    [y, c; z, d]
    |> List.map (fun (value, right) ->
        let result = unbox<Test_RBTree_Tree> (sharpurs_apply (sharpurs_apply partial2 (box value)) (box right))
        checkDefault "reused two-argument prefix" black a value right result
        result)
check "two-argument prefix results stay independent"
    (partial2Results = [node black a y c; node black a z d])

let partial3 = sharpurs_apply partial2 (box y)
let partial3Results =
    [b; c; d]
    |> List.map (fun right ->
        let result = unbox<Test_RBTree_Tree> (sharpurs_apply partial3 (box right))
        checkDefault "reused three-argument prefix" black a y right result
        result)
check "three-argument prefix results stay independent"
    (partial3Results = [node black a y b; node black a y c; node black a y d])

// Diagnostic only: ill-typed FFI inputs are outside the PureScript type contract.
// An intentionally invalid payload establishes that casts and matching wait
// until saturation, rather than moving into one of the public partial stages.
let deferred1 = sharpurs_apply Test_RBTree_balance (box "invalid color")
check "first partial stage defers body" (deferred1 :? (obj -> obj))
let deferred2 = sharpurs_apply deferred1 (box a)
check "second partial stage defers body" (deferred2 :? (obj -> obj))
let deferred3 = sharpurs_apply deferred2 (box y)
check "third partial stage defers body" (deferred3 :? (obj -> obj))
let rejectedAtSaturation =
    try
        sharpurs_apply deferred3 (box d) |> ignore
        false
    with
    | :? InvalidCastException -> true
    | :? Reflection.TargetInvocationException as failure when (failure.InnerException :? InvalidCastException) -> true
check "invalid payload is rejected at saturation" rejectedAtSaturation

let exceptionTypes call =
    let rec collect (error: exn) =
        error.GetType().FullName :: (if isNull error.InnerException then [] else collect error.InnerException)
    try
        call () |> ignore
        []
    with error -> collect error

let expectedException =
    [typeof<Reflection.TargetInvocationException>.FullName; typeof<InvalidCastException>.FullName]
let publicException = exceptionTypes (fun () -> sharpurs_apply deferred3 (box d))
let directApplyException = exceptionTypes (fun () ->
    Test_RBTree_balance_direct_apply (box "invalid color") (box a) (box y) (box d))
check "public wrapper keeps exactly one exception boundary" (publicException = expectedException)
check "direct application keeps public exception boundary" (directApplyException = publicException)

let publicTrace = ResizeArray<int>()
let publicArgument index value =
    publicTrace.Add index
    box value
let publicTracedResult =
    unbox<Test_RBTree_Tree>
        (sharpurs_apply
            (sharpurs_apply
                (sharpurs_apply
                    (sharpurs_apply Test_RBTree_balance (publicArgument 1 black))
                    (publicArgument 2 a))
                (publicArgument 3 y))
            (publicArgument 4 d))
check "public argument order and single evaluation" (List.ofSeq publicTrace = [1; 2; 3; 4])
checkDefault "public traced call" black a y d publicTracedResult

let directTrace = ResizeArray<int>()
let directArgument index value =
    directTrace.Add index
    box value
let directTracedResult =
    unbox<Test_RBTree_Tree>
        (Test_RBTree_balance_direct
            (directArgument 1 black)
            (directArgument 2 a)
            (directArgument 3 y)
            (directArgument 4 d))
check "direct argument order and single evaluation" (List.ofSeq directTrace = [1; 2; 3; 4])
checkDefault "direct traced call" black a y d directTracedResult
check "public and direct traced calls agree" (publicTracedResult = directTracedResult)

printfn "Generated balance: %d rotation, sharing, partial application, deferral and evaluation-order checks passed" checks

#if !BASELINE
let nativeBalance = Test_RBTree_balance_adt_native
let nativeApplyBalance = Test_RBTree_balance_adt_native_apply
for name, left, value, right in rotations do
    checkRotation (name + " native") (nativeBalance black left value right)
    checkRotation (name + " native application") (nativeApplyBalance black left value right)
for name, color in ["red", red; "black", black] do
    checkDefault (name + " native") color a y d (nativeBalance color a y d)
    checkDefault (name + " native application") color a y d (nativeApplyBalance color a y d)
printfn "Including native entry points: %d balance checks passed" checks
#endif
