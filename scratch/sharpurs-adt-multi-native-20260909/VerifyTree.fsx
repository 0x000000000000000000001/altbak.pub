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

let empty = unbox<Test_RBTree_Tree> Test_RBTree_E
let insert value tree =
    unbox<Test_RBTree_Tree> (sharpurs_apply (sharpurs_apply Test_RBTree_insert (box value)) (box tree))
let depth tree = unbox<int> (sharpurs_apply Test_RBTree_depth (box tree))
let blacken tree = unbox<Test_RBTree_Tree> (sharpurs_apply Test_RBTree_makeBlack (box tree))

let rec values tree =
    match tree with
    | Test_RBTree_Eusd_Ctor -> []
    | Test_RBTree_Tusd_Ctor (_, left, value, right) -> values left @ [value] @ values right

let rec referenceDepth tree =
    match tree with
    | Test_RBTree_Eusd_Ctor -> 0
    | Test_RBTree_Tusd_Ctor (_, left, _, right) -> 1 + max (referenceDepth left) (referenceDepth right)

let rec validateNodes parentRed low high tree =
    match tree with
    | Test_RBTree_Eusd_Ctor -> 1
    | Test_RBTree_Tusd_Ctor (color, left, value, right) ->
        check "BST lower bound" (Option.forall (fun lower -> lower < value) low)
        check "BST upper bound" (Option.forall (fun upper -> value < upper) high)
        let red = color = Test_RBTree_Rusd_Ctor
        check "red node has black parent" (not (parentRed && red))
        let leftHeight = validateNodes red low (Some value) left
        let rightHeight = validateNodes red (Some value) high right
        check "equal black heights" (leftHeight = rightHeight)
        leftHeight + (if red then 0 else 1)

let validate expected tree =
    match tree with
    | Test_RBTree_Eusd_Ctor -> ()
    | Test_RBTree_Tusd_Ctor (color, _, _, _) -> check "black root" (color = Test_RBTree_Busd_Ctor)
    validateNodes false None None tree |> ignore
    check "sorted unique contents" (values tree = (expected |> List.distinct |> List.sort))
    check "source depth matches independent traversal" (depth tree = referenceDepth tree)
    check "native depth agrees with public wrapper" (Test_RBTree_depth_adt_native tree = depth tree)
    check "recursive ABI bridge agrees" (unbox<int> (Test_RBTree_depth_tco (box tree)) = depth tree)

let orders = [
    "empty", []
    "singleton", [5]
    "ascending", [1 .. 64]
    "descending", [64 .. -1 .. 1]
    "alternating extremes", List.collect (fun n -> [n; 65 - n]) [1 .. 32]
    "duplicates", List.replicate 4 [7; 3; 11; 3; 7; 1; 13; 11] |> List.concat
    "signed limits", [Int32.MinValue; 0; Int32.MaxValue; -1; 1; Int32.MinValue; Int32.MaxValue]
]

for label, order in orders do
    let mutable tree = empty
    let mutable expected = []
    validate expected tree
    for value in order do
        let previous = tree
        let previousValues = values previous
        tree <- insert value tree
        expected <- value :: expected
        validate expected tree
        check (label + ": old tree remains persistent") (values previous = previousValues)
    printfn "%s: %d insertions, final depth %d" label order.Length (depth tree)

let seed = List.fold (fun tree value -> insert value tree) empty [1 .. 31]
let originalValues = values seed
match seed with
| Test_RBTree_Eusd_Ctor -> failwith "Expected a nonempty seed"
| Test_RBTree_Tusd_Ctor (_, left, value, right) ->
    let redRoot = Test_RBTree_Tusd_Ctor (Test_RBTree_Rusd_Ctor, left, value, right)
    for label, transformed in ["public", blacken redRoot; "native", Test_RBTree_makeBlack_adt_native redRoot] do
        match transformed with
        | Test_RBTree_Tusd_Ctor (color, nextLeft, nextValue, nextRight) ->
            check (label + ": black color") (color = Test_RBTree_Busd_Ctor)
            check (label + ": payload preserved") (nextValue = value)
            check (label + ": left subtree shared") (Object.ReferenceEquals(left, nextLeft))
            check (label + ": right subtree shared") (Object.ReferenceEquals(right, nextRight))
        | _ -> failwith "makeBlack discarded a nonempty root"
    let partial = sharpurs_apply (sharpurs_apply Test_RBTree_T Test_RBTree_B) (box left)
    for nextValue in [value; value + 1] do
        let rebuilt = unbox<Test_RBTree_Tree> (sharpurs_apply (sharpurs_apply partial (box nextValue)) (box right))
        match rebuilt with
        | Test_RBTree_Tusd_Ctor (color, nextLeft, actual, nextRight) ->
            check "partial native constructor color" (color = Test_RBTree_Busd_Ctor)
            check "partial native constructor payload" (actual = nextValue)
            check "partial native constructor shares left" (Object.ReferenceEquals(left, nextLeft))
            check "partial native constructor shares right" (Object.ReferenceEquals(right, nextRight))
        | _ -> failwith "Constructor wrapper discarded a nonempty root"
check "seed remains unchanged" (values seed = originalValues)
check "makeBlack empty" (blacken empty = empty)
check "native makeBlack empty" (Test_RBTree_makeBlack_adt_native empty = empty)
printfn "Real generated RBTree: %d invariant, depth, ABI and sharing checks passed" checks
