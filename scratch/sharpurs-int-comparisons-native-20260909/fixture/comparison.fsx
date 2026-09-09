let objMap = Map.empty<string, obj>
let unbox<'a> (x: obj) : 'a = unbox x
let (|Unbox|) (x: obj) = unbox x

let undefined = Unchecked.defaultof<obj>
let Prim_undefined = undefined
let intMod a b = unbox<int> a % unbox<int> b
let semiringInt = 0

// PureScript's Euclidean Int modulo, including zero and Int32.MinValue / -1.
let sharpurs_int_mod (left: int) (right: int) : int =
    if right = 0 || right = -1 then 0
    else
        let remainder = left % right
        if remainder < 0 then
            if right > 0 then remainder + right else remainder - right
        else remainder

let sharpurs_apply (func: obj) (arg: obj) : obj =
    if isNull func then failwith "sharpurs_apply: func is null!"
    match func with
    | :? (obj -> obj) as invoke ->
        try invoke arg
        // Keep the exception boundary of MethodInfo.Invoke for callers and FFI.
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    | _ ->
        let method = func.GetType().GetMethods() |> Array.find (fun m -> m.Name = "Invoke" && m.GetParameters().Length = 1)
        method.Invoke(func, [| arg |])



let mutable fallbackCalls = 0
let events = ResizeArray<int>()
let IntCompare_track : obj = box (fun (label: obj) -> box (fun (value: obj) -> events.Add(unbox<int> label); value))
let Data_Ord_Ordusd_Dict : obj = box (fun (value: obj) -> value)
let Data_Eq_Equsd_Dict = Data_Ord_Ordusd_Dict
let comparison : obj = box (fun (x: obj) -> box (fun (y: obj) -> box (Unchecked.compare x y)))
let equality : obj = box (fun (x: obj) -> box (fun (y: obj) -> box (x = y)))
let Data_Eq_eqInt : obj = box (Map.ofList ["eq", equality])
let Data_Eq_eq : obj = box (fun (dict: obj) -> Map.find "eq" (unbox<Map<string,obj>> dict))
let Data_Ord_ordInt : obj = box (Map.ofList ["compare", comparison; "Eq0", box (fun (_: obj) -> Data_Eq_eqInt)])
let Data_Ord_ordString = Data_Ord_ordInt
let Data_Ord_ordNumber = Data_Ord_ordInt
let Data_Ord_compare : obj = box (fun (dict: obj) -> Map.find "compare" (unbox<Map<string,obj>> dict))
let relation predicate : obj = box (fun (dict: obj) -> box (fun (x: obj) -> box (fun (y: obj) ->
    fallbackCalls <- fallbackCalls + 1
    box (predicate (unbox<int> (sharpurs_apply (sharpurs_apply (sharpurs_apply Data_Ord_compare dict) x) y))))))
let Data_Ord_lessThan = relation (fun ordering -> ordering < 0)
let Data_Ord_greaterThan = relation (fun ordering -> ordering > 0)


let IntCompare_lessThan  = (sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box Data_Ord_ordInt))))

let IntCompare_Reverse  = (box (fun (x: obj) -> (box x)))

let IntCompare_stringLess  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box Data_Ord_ordString)))))) (box ((box x)))))) (box ((box y))))))))

let IntCompare_partial  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box Data_Ord_ordInt)))))) (box ((box 5))))

let IntCompare_ordered  = (box (fun (x: obj) -> (box (fun (y: obj) -> (box ((unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntCompare_track))) (box ((box 1)))))) (box ((box x))))))) < (unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((box IntCompare_track))) (box ((box 2)))))) (box ((box y)))))))))))))

let IntCompare_numberLess  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box Data_Ord_ordNumber)))))) (box ((box x)))))) (box ((box y))))))))

let IntCompare_less  = (box (fun (x: obj) -> (box (fun (y: obj) -> (box ((unbox<int> (box ((box x)))) < (unbox<int> (box ((box y))))))))))

let IntCompare_greater  = (box (fun (x: obj) -> (box (fun (y: obj) -> (box ((unbox<int> (box ((box x)))) > (unbox<int> (box ((box y))))))))))

let IntCompare_genericLess  = (box (fun (dictOrd: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box y))))))))))

let IntCompare_eqReverse  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Eq_eqInt)))))) (box ((box l)))))) (box ((box r)))))))))))) Map.empty))))))

let IntCompare_ordReverse  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box Data_Ord_ordInt)))))) (box ((box y)))))) (box ((box x)))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box IntCompare_eqReverse))))) Map.empty)))))))

let IntCompare_custom  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box IntCompare_ordReverse)))))) (box ((sharpurs_apply (box ((box IntCompare_Reverse))) (box ((box x))))))))) (box ((sharpurs_apply (box ((box IntCompare_Reverse))) (box ((box y)))))))))))

let IntCompare_annotated  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box IntCompare_lessThan))) (box ((box x)))))) (box ((box y))))))))


let mutable checks = 0
let check label condition = if not condition then failwith label else checks <- checks + 1
let call fn x y = sharpurs_apply (sharpurs_apply fn (box x)) (box y) |> unbox<bool>
let cases = [
    (-2147483648, -2147483648, false, false, false);
    (-2147483648, -2147483647, true, false, false);
    (-2147483648, -1, true, false, false);
    (-2147483648, 0, true, false, false);
    (-2147483648, 1, true, false, false);
    (-2147483648, 2147483646, true, false, false);
    (-2147483648, 2147483647, true, false, false);
    (-2147483647, -2147483648, false, true, true);
    (-2147483647, -2147483647, false, false, false);
    (-2147483647, -1, true, false, false);
    (-2147483647, 0, true, false, false);
    (-2147483647, 1, true, false, false);
    (-2147483647, 2147483646, true, false, false);
    (-2147483647, 2147483647, true, false, false);
    (-1, -2147483648, false, true, true);
    (-1, -2147483647, false, true, true);
    (-1, -1, false, false, false);
    (-1, 0, true, false, false);
    (-1, 1, true, false, false);
    (-1, 2147483646, true, false, false);
    (-1, 2147483647, true, false, false);
    (0, -2147483648, false, true, true);
    (0, -2147483647, false, true, true);
    (0, -1, false, true, true);
    (0, 0, false, false, false);
    (0, 1, true, false, false);
    (0, 2147483646, true, false, false);
    (0, 2147483647, true, false, false);
    (1, -2147483648, false, true, true);
    (1, -2147483647, false, true, true);
    (1, -1, false, true, true);
    (1, 0, false, true, true);
    (1, 1, false, false, false);
    (1, 2147483646, true, false, false);
    (1, 2147483647, true, false, false);
    (2147483646, -2147483648, false, true, true);
    (2147483646, -2147483647, false, true, true);
    (2147483646, -1, false, true, true);
    (2147483646, 0, false, true, true);
    (2147483646, 1, false, true, true);
    (2147483646, 2147483646, false, false, false);
    (2147483646, 2147483647, true, false, false);
    (2147483647, -2147483648, false, true, true);
    (2147483647, -2147483647, false, true, true);
    (2147483647, -1, false, true, true);
    (2147483647, 0, false, true, true);
    (2147483647, 1, false, true, true);
    (2147483647, 2147483646, false, true, true);
    (2147483647, 2147483647, false, false, false)
]
for x, y, less, greater, custom in cases do
    let before = fallbackCalls
    check "less matches JS" (call IntCompare_less x y = less)
    check "greater matches JS" (call IntCompare_greater x y = greater)
    check "native calls bypass dictionary" (fallbackCalls = before)
    check "custom order matches JS" (call IntCompare_custom x y = custom)
    check "custom order uses its dictionary" (fallbackCalls = before + 1)
    check "annotated alias matches JS" (call IntCompare_annotated x y = less)
    check "generic Int dispatch" (call (sharpurs_apply IntCompare_genericLess Data_Ord_ordInt) x y = less)
    check "partial comparator" (unbox<bool> (sharpurs_apply IntCompare_partial (box y)) = (5 < y))
check "String fallback" (call IntCompare_stringLess "a" "b")
check "Number fallback" (call IntCompare_numberLess 1.25 1.5)
events.Clear()
check "ordered comparison result" (call IntCompare_ordered 3 7)
check "left-to-right once" (List.ofSeq events = [1; 2])
printfn "int-comparison runtime: %d checks passed" checks
