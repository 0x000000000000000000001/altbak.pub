#r "/Users/0x1/Documents/htdocs/altbak.pub-sharpurs/run/bak/sharp/output/Main/bin/Release/net8.0/Program.dll"
open System
open System.Diagnostics
open PureScript_Test_RBTree
open Sharpurs_Prelude

module GeneratedBaseline =
    let Test_RBTree_balance_direct (v: obj) (v1: obj) (v2: obj) (v3: obj) : obj = (match (((unbox ((box v))), (unbox ((box v1))), (unbox ((box v2))), (unbox ((box v3))))) with | (Test_RBTree_Busd_Ctor, Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), Unbox(Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), a, x, b)), y, c), z, d) -> ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_R_adt_native)))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box a))) (unbox ((box x))) (unbox ((box b))))))) (unbox ((box y))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box c))) (unbox ((box z))) (unbox ((box d)))))))))) | (Test_RBTree_Busd_Ctor, Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), a, x, Unbox(Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), b, y, c))), z, d) -> ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_R_adt_native)))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box a))) (unbox ((box x))) (unbox ((box b))))))) (unbox ((box y))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box c))) (unbox ((box z))) (unbox ((box d)))))))))) | (Test_RBTree_Busd_Ctor, a, x, Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), Unbox(Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), b, y, c)), z, d)) -> ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_R_adt_native)))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box a))) (unbox ((box x))) (unbox ((box b))))))) (unbox ((box y))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box c))) (unbox ((box z))) (unbox ((box d)))))))))) | (Test_RBTree_Busd_Ctor, a, x, Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), b, y, Unbox(Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), c, z, d)))) -> ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_R_adt_native)))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box a))) (unbox ((box x))) (unbox ((box b))))))) (unbox ((box y))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box c))) (unbox ((box z))) (unbox ((box d)))))))))) | (color, a, x, b) -> ((box (Test_RBTree_T_adt_native (unbox ((box color))) (unbox ((box a))) (unbox ((box x))) (unbox ((box b)))))))
    
    let Test_RBTree_balance_direct_apply (v: obj) (v1: obj) (v2: obj) (v3: obj) : obj =
        try Test_RBTree_balance_direct v v1 v2 v3
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let Test_RBTree_balance = (box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (fun (v3: obj) -> (Test_RBTree_balance_direct v v1 v2 v3)))))))))
    
    let rec Test_RBTree_ins_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (x, Test_RBTree_Eusd_Ctor) -> ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_R_adt_native)))) (unbox ((box (Test_RBTree_E_adt_native)))) (unbox ((box x))) (unbox ((box (Test_RBTree_E_adt_native))))))) | (x, Test_RBTree_Tusd_Ctor(color, a, y, b)) -> ((match ((unbox ((box ((unbox<int> (box ((box x)))) < (unbox<int> (box ((box y))))))))) with | LitBool true () -> ((Test_RBTree_balance_direct_apply ((box ((box color)))) ((box ((Test_RBTree_ins_tco ((box x)) ((box a)))))) ((box ((box y)))) ((box ((box b)))))) | _ -> ((match ((unbox ((box ((unbox<int> (box ((box x)))) > (unbox<int> (box ((box y))))))))) with | LitBool true () -> ((Test_RBTree_balance_direct_apply ((box ((box color)))) ((box ((box a)))) ((box ((box y)))) ((box ((Test_RBTree_ins_tco ((box x)) ((box b)))))))) | _ -> ((box (Test_RBTree_T_adt_native (unbox ((box color))) (unbox ((box a))) (unbox ((box y))) (unbox ((box b))))))))))))
    and Test_RBTree_ins = box (fun (v: obj) ->  (fun (v1: obj) -> Test_RBTree_ins_tco v v1))
    
    
    let Test_RBTree_insert_direct (x: obj) (s: obj) : obj = (sharpurs_apply (box ((box Test_RBTree_makeBlack))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_RBTree_ins))) (box ((box x)))))) (box ((box s)))))))
    
    let Test_RBTree_insert_direct_apply (x: obj) (s: obj) : obj =
        try Test_RBTree_insert_direct x s
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let Test_RBTree_insert = (box (fun (x: obj) -> (box (fun (s: obj) -> (Test_RBTree_insert_direct x s)))))
    
    let rec Test_RBTree_buildTree_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), acc) -> ((box acc)) | (n, acc) -> ((Test_RBTree_buildTree_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((Test_RBTree_insert_direct_apply ((box ((box n)))) ((box ((box acc))))))))))
    and Test_RBTree_buildTree = box (fun (v: obj) ->  (fun (v1: obj) -> Test_RBTree_buildTree_tco v v1))
    
    

let originalBuild = GeneratedBaseline.Test_RBTree_buildTree_tco
let originalInsert = GeneratedBaseline.Test_RBTree_insert_direct_apply
let originalBalance = GeneratedBaseline.Test_RBTree_balance_direct_apply

// Same four Okasaki rotations, same persistent DU and constructor layout.
let typedBalance (color: Test_RBTree_Color) (left: Test_RBTree_Tree) (value: int) (right: Test_RBTree_Tree) : Test_RBTree_Tree =
    match color, left, value, right with
    | Test_RBTree_Busd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, a, x, b), y, c), z, d
    | Test_RBTree_Busd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, a, x, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, b, y, c)), z, d
    | Test_RBTree_Busd_Ctor, a, x, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, b, y, c), z, d)
    | Test_RBTree_Busd_Ctor, a, x, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, b, y, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, c, z, d)) ->
        Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor,
            Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, a, x, b), y,
            Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, c, z, d))
    | color, a, x, b -> Test_RBTree_Tusd_Ctor(color, a, x, b)

let typedBalanceApply color left value right =
    try typedBalance color left value right
    with ex -> raise (Reflection.TargetInvocationException(ex))

let rec typedIns (x: int) (tree: Test_RBTree_Tree) : Test_RBTree_Tree =
    match tree with
    | Test_RBTree_Eusd_Ctor -> Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, Test_RBTree_Eusd_Ctor, x, Test_RBTree_Eusd_Ctor)
    | Test_RBTree_Tusd_Ctor(color, a, y, b) ->
        if x < y then typedBalanceApply color (typedIns x a) y b
        elif x > y then typedBalanceApply color a y (typedIns x b)
        else Test_RBTree_Tusd_Ctor(color, a, y, b)

let Test_RBTree_ins_tco (v: obj) (v1: obj) : obj = box (typedIns (unbox<int> v) (unbox<Test_RBTree_Tree> v1))
let Test_RBTree_ins = box (fun (v: obj) -> (fun (v1: obj) -> Test_RBTree_ins_tco v v1))

// The following generated insert and outer build loop are copied byte-for-byte below.

let Test_RBTree_insert_direct (x: obj) (s: obj) : obj = (sharpurs_apply (box ((box Test_RBTree_makeBlack))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_RBTree_ins))) (box ((box x)))))) (box ((box s)))))))

let Test_RBTree_insert_direct_apply (x: obj) (s: obj) : obj =
    try Test_RBTree_insert_direct x s
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let Test_RBTree_insert = (box (fun (x: obj) -> (box (fun (s: obj) -> (Test_RBTree_insert_direct x s)))))

let rec Test_RBTree_buildTree_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), acc) -> ((box acc)) | (n, acc) -> ((Test_RBTree_buildTree_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((Test_RBTree_insert_direct_apply ((box ((box n)))) ((box ((box acc))))))))))
and Test_RBTree_buildTree = box (fun (v: obj) ->  (fun (v1: obj) -> Test_RBTree_buildTree_tco v v1))

let mutable checks = 0
let check label passed =
    if not passed then failwith label
    checks <- checks + 1
let empty = Test_RBTree_Eusd_Ctor
let rec values tree =
    match tree with
    | Test_RBTree_Eusd_Ctor -> []
    | Test_RBTree_Tusd_Ctor(_, l, x, r) -> values l @ [x] @ values r
let rec invariant parentRed low high tree =
    match tree with
    | Test_RBTree_Eusd_Ctor -> 1
    | Test_RBTree_Tusd_Ctor(c, l, x, r) ->
        check "lower BST bound" (Option.forall (fun lo -> lo < x) low)
        check "upper BST bound" (Option.forall (fun hi -> x < hi) high)
        let red = c = Test_RBTree_Rusd_Ctor
        check "red parent" (not (parentRed && red))
        let lb = invariant red low (Some x) l
        let rb = invariant red (Some x) high r
        check "black height" (lb = rb)
        lb + (if red then 0 else 1)
let validate expected tree =
    invariant false None None tree |> ignore
    match tree with
    | Test_RBTree_Eusd_Ctor -> ()
    | Test_RBTree_Tusd_Ctor(c, _, _, _) -> check "root black" (c = Test_RBTree_Busd_Ctor)
    check "values sorted and unique" (values tree = List.sort (List.distinct expected))

let orders = [
    []; [5]; [1..64]; [64.. -1 ..1]
    List.collect (fun x -> [x;65-x]) [1..32]
    List.replicate 4 [7;3;11;3;7;1;13;11] |> List.concat
    [Int32.MinValue;0;Int32.MaxValue;-1;1;Int32.MinValue;Int32.MaxValue]
]
for order in orders do
    let mutable original, candidate = empty, empty
    let mutable seen = []
    for x in order do
        let previous, contents = candidate, values candidate
        original <- unbox<Test_RBTree_Tree> (originalInsert (box x) (box original))
        candidate <- unbox<Test_RBTree_Tree> (Test_RBTree_insert_direct_apply (box x) (box candidate))
        seen <- x::seen
        validate seen candidate
        check "identical full tree" (candidate = original)
        check "persistent prior tree" (values previous = contents)

let node c a x b = Test_RBTree_Tusd_Ctor(c,a,x,b)
let red, black = Test_RBTree_Rusd_Ctor, Test_RBTree_Busd_Ctor
let a,b,c,d = node black empty 1 empty,node black empty 3 empty,node black empty 5 empty,node black empty 7 empty
let rotations = [
    node red (node red a 2 b) 4 c, 6, d
    node red a 2 (node red b 4 c), 6, d
    a, 2, node red (node red b 4 c) 6 d
    a, 2, node red b 4 (node red c 6 d)
]
for left,x,right in rotations do
    check "rotation exact" (typedBalance black left x right = unbox<Test_RBTree_Tree> (originalBalance (box black) (box left) (box x) (box right)))
for color in [red;black] do
    check "default exact" (typedBalance color a 4 d = unbox<Test_RBTree_Tree> (originalBalance (box color) (box a) (box 4) (box d)))
for n in [0;1;2;3;7;31;127;1024;100000] do
    let original = unbox<Test_RBTree_Tree> (originalBuild (box n) (box empty))
    let candidate = unbox<Test_RBTree_Tree> (Test_RBTree_buildTree_tco (box n) (box empty))
    check "full build identical" (original = candidate)
    validate [1..n] candidate
    if n = 100000 then check "depth22" (Test_RBTree_depth_adt_native candidate = 22)
printfn "VALIDATION %d assertions passed" checks

let run build =
    let tree = unbox<Test_RBTree_Tree> (build (box 100000) (box empty))
    let result = Test_RBTree_depth_adt_native tree
    if result <> 22 then failwithf "Bad depth %d" result
    result
let variants = [| originalBuild; Test_RBTree_buildTree_tco |]
for _ in 1..3 do
    for variant in variants do run variant |> ignore
let samples = [|ResizeArray<float>();ResizeArray<float>()|]
for pair in 0..4 do
    for which in (if pair % 2 = 0 then [0;1] else [1;0]) do
        let timer = Stopwatch.StartNew()
        let result = run variants[which]
        timer.Stop()
        let ms = timer.Elapsed.TotalMilliseconds
        samples[which].Add ms
        printfn "SAMPLE pair=%d variant=%s ms=%.5f depth=%d" pair (if which=0 then "current" else "typed-ins-balance") ms result
let median (sample: ResizeArray<float>) = sample.ToArray() |> Array.sort |> fun sorted -> sorted[sorted.Length/2]
let before, after = median samples[0], median samples[1]
printfn "MEDIAN current=%.5f typed=%.5f delta=%.2f%%" before after ((after/before-1.0)*100.0)


