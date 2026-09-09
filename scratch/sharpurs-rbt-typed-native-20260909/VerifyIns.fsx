#if BASELINE
#r "before/bin/Release/net8.0/Program.dll"
#else
#r "after/bin/Release/net8.0/Program.dll"
#endif

open System
open System.Collections.Generic
open System.IO
open System.Text
open System.Text.Json
open PureScript_Test_RBTree
open Sharpurs_Prelude

let mutable checks = 0
let check label condition =
    if not condition then failwith label
    checks <- checks + 1

let red = Test_RBTree_Rusd_Ctor
let black = Test_RBTree_Busd_Ctor
let empty = Test_RBTree_Eusd_Ctor
let node color left value right = Test_RBTree_Tusd_Ctor(color, left, value, right)
let leaf value = node black empty value empty
let publicIns value tree =
    unbox<Test_RBTree_Tree> (sharpurs_apply (sharpurs_apply Test_RBTree_ins (box value)) (box tree))
let bridgeIns value tree = unbox<Test_RBTree_Tree> (Test_RBTree_ins_tco (box value) (box tree))
let insert value tree =
    unbox<Test_RBTree_Tree> (sharpurs_apply (sharpurs_apply Test_RBTree_insert (box value)) (box tree))
let blacken = Test_RBTree_makeBlack_adt_native
let depth = Test_RBTree_depth_adt_native
let paths =
    ["public", publicIns; "bridge", bridgeIns
#if BASELINE
#else
     "native", Test_RBTree_ins_adt_native
#endif
    ]

// Exact prefix encoding of constructor, color and signed payload. The before
// process records the actual generated implementation; after checks every byte.
let encode tree =
    let text = StringBuilder()
    let rec loop current =
        match current with
        | Test_RBTree_Eusd_Ctor -> text.Append('E') |> ignore
        | Test_RBTree_Tusd_Ctor(color, left, value, right) ->
            text.Append(if color = red then 'R' else 'B').Append(value).Append('(') |> ignore
            loop left
            text.Append(',') |> ignore
            loop right
            text.Append(')') |> ignore
    loop tree
    text.ToString()

let observations = ResizeArray<string array>()
let observe label value = observations.Add([|label; value|])
let observeTree label tree = observe label (encode tree)

let rec contents tree =
    match tree with
    | Test_RBTree_Eusd_Ctor -> []
    | Test_RBTree_Tusd_Ctor(_, left, value, right) -> contents left @ (value :: contents right)

let rec validate parentRed low high tree =
    match tree with
    | Test_RBTree_Eusd_Ctor -> 1, 0, 0
    | Test_RBTree_Tusd_Ctor(color, left, value, right) ->
        check "lower BST bound" (Option.forall (fun lower -> lower < value) low)
        check "upper BST bound" (Option.forall (fun upper -> value < upper) high)
        check "no adjacent red nodes" (not (parentRed && color = red))
        let lh, ld, lc = validate (color = red) low (Some value) left
        let rh, rd, rc = validate (color = red) (Some value) high right
        check "identical black heights" (lh = rh)
        lh + (if color = black then 1 else 0), 1 + max ld rd, 1 + lc + rc

let random = Random(20260909)
let orders =
    [ "ascending", [1 .. 127]
      "descending", [128 .. -1 .. 1]
      "alternating", List.collect (fun n -> [n; 129 - n]) [1 .. 64]
      "duplicates", List.replicate 5 [7; 3; 11; 3; 7; 1; 13; 11] |> List.concat
      "bounds", [Int32.MinValue; Int32.MaxValue; 0; -1; 1; Int32.MinValue; Int32.MaxValue]
      "seeded signed values", [for _ in 1 .. 96 -> random.Next(-10000, 10001)] ]

for label, order in orders do
    let mutable tree = empty
    let mutable expected = Set.empty
    observeTree (label + "/empty") tree
    for index, value in List.indexed order do
        let previous = tree
        let previousEncoding = encode previous
        let raw = publicIns value previous
        for name, path in paths do
            let actual = path value previous
            check (label + "/" + name + ": exact ins structure") (actual = raw)
        tree <- insert value previous
        expected <- Set.add value expected
        check "insert only blackens ins" (tree = blacken raw)
        check "previous tree remains unchanged" (encode previous = previousEncoding)
        let _, actualDepth, count = validate false None None tree
        check "expected sorted values" (contents tree = Set.toList expected)
        check "correct node count" (count = Set.count expected)
        check "native depth agrees with traversal" (depth tree = actualDepth)
        match tree with
        | Test_RBTree_Tusd_Ctor(color, _, _, _) -> check "root is black" (color = black)
        | _ -> failwith "Insertion returned an empty tree"
        observeTree (sprintf "%s/%d/raw" label index) raw
        observeTree (sprintf "%s/%d/insert" label index) tree

let seed = node black (leaf 1) 4 (leaf 8)
let seedEncoding = encode seed
for name, path in paths do
    let repeated = path 4 seed
    check (name + ": root duplicate has identical structure") (repeated = seed)
    match seed, repeated with
    | Test_RBTree_Tusd_Ctor(_, left, _, right), Test_RBTree_Tusd_Ctor(_, nextLeft, _, nextRight) ->
        check (name + ": duplicate shares left") (Object.ReferenceEquals(left, nextLeft))
        check (name + ": duplicate shares right") (Object.ReferenceEquals(right, nextRight))
    | _ -> failwith "Expected nonempty duplicate result"
    for value, sharesLeft in [2, false; 6, true] do
        let grown = path value seed
        match seed, grown with
        | Test_RBTree_Tusd_Ctor(_, left, _, right), Test_RBTree_Tusd_Ctor(_, nextLeft, _, nextRight) ->
            check (name + ": untouched branch is shared")
                (if sharesLeft then Object.ReferenceEquals(left, nextLeft)
                 else Object.ReferenceEquals(right, nextRight))
        | _ -> failwith "Expected nonempty grown tree"
        check (name + ": original remains persistent") (encode seed = seedEncoding)
        check (name + ": inserted value is present") (contents grown = List.sort [1; 4; 8; value])

let partial = sharpurs_apply Test_RBTree_ins (box 2)
let partialInputs = [empty; seed; insert 2 seed; node red empty Int32.MaxValue empty]
let partialResults =
    partialInputs |> List.map (fun input ->
        let result = unbox<Test_RBTree_Tree> (sharpurs_apply partial (box input))
        check "reused partial equals raw bridge" (result = bridgeIns 2 input)
        result)
for index, input in List.indexed partialInputs do
    check "previous partial results are independent" (partialResults.[index] = publicIns 2 input)
    observeTree (sprintf "partial/%d" index) partialResults.[index]

let exceptionTypes call =
    let rec collect (error: exn) =
        error.GetType().FullName :: (if isNull error.InnerException then [] else collect error.InnerException)
    try call (); [] with error -> collect error

// Valid ABI arguments. Failures belong to argument evaluation, before entering
// the total ins body, so both value order and exception chain must be unchanged.
for name, path in paths do
    let trace = ResizeArray<int>()
    let valueArgument () = trace.Add(1); 2
    let treeArgument () = trace.Add(2); seed
    let result = path (valueArgument ()) (treeArgument ())
    check (name + ": arguments each evaluated once, in order") (List.ofSeq trace = [1; 2])
    check (name + ": traced result") (result = publicIns 2 seed)
    for failAt in [1; 2] do
        trace.Clear()
        let sentinel = InvalidOperationException(sprintf "argument %d failed" failAt)
        let mutable captured: exn option = None
        let valueArgument () =
            trace.Add(1)
            if failAt = 1 then raise sentinel
            2
        let treeArgument () =
            trace.Add(2)
            if failAt = 2 then raise sentinel
            seed
        let chain = exceptionTypes (fun () ->
            try path (valueArgument ()) (treeArgument ()) |> ignore
            with error -> captured <- Some error; reraise())
        check (name + ": correct failure trace") (List.ofSeq trace = [1 .. failAt])
        check (name + ": original argument exception identity")
            (captured |> Option.exists (fun error -> Object.ReferenceEquals(error, sentinel)))
        check (name + ": no extra exception wrapper on arguments")
            (chain = [typeof<InvalidOperationException>.FullName])
        if name = "public" || name = "bridge" then
            observe (sprintf "argument-failure/%s/%d" name failAt) (String.concat "/" chain)

// The exact benchmark workload uses unchanged generated buildTree and depth.
let large =
    unbox<Test_RBTree_Tree>
        (sharpurs_apply (sharpurs_apply Test_RBTree_buildTree (box 100000)) (box empty))
let _, largeDepth, largeCount = validate false None None large
check "full workload has 100000 unique nodes" (largeCount = 100000)
check "full workload has depth 22" (largeDepth = 22 && depth large = 22)
check "full workload contents are precisely 1..100000" (contents large = [1 .. 100000])
observeTree "descending/100000" large

let oraclePath = Path.Combine(__SOURCE_DIRECTORY__, "validation-oracle.before.json")
#if BASELINE
File.WriteAllText(oraclePath, JsonSerializer.Serialize(observations.ToArray()))
printfn "Recorded %d exact generated-before observations" observations.Count
#else
let expected = JsonSerializer.Deserialize<string array array>(File.ReadAllText(oraclePath))
check "before/after observation counts agree" (expected.Length = observations.Count)
for index in 0 .. expected.Length - 1 do
    check (sprintf "oracle label %d" index) (expected.[index].[0] = observations.[index].[0])
    check ("exact baseline result: " + observations.[index].[0]) (expected.[index].[1] = observations.[index].[1])
printfn "Compared %d exact generated-before observations" observations.Count
#endif

printfn "Generated ins: %d structure, invariants, persistence, ABI, evaluation and full-workload checks passed" checks

// Out-of-contract diagnostics: the language cannot supply a String as Int or a
// null native union. Print consequences without making those values a semantic
// acceptance gate. This also locates rejection at saturation for both bridges.
let invalidPartial = sharpurs_apply Test_RBTree_ins (box "invalid Int")
printfn "DIAGNOSTIC ill-typed Int: partial produced %s; saturated chain %A"
    (invalidPartial.GetType().FullName)
    (exceptionTypes (fun () -> sharpurs_apply invalidPartial (box seed) |> ignore))
printfn "DIAGNOSTIC null tree: public chain %A; bridge chain %A"
    (exceptionTypes (fun () -> sharpurs_apply (sharpurs_apply Test_RBTree_ins (box 2)) null |> ignore))
    (exceptionTypes (fun () -> Test_RBTree_ins_tco (box 2) null |> ignore))
