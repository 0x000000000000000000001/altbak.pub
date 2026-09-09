open System

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



let (|LitInt|_|) (expected: int) (value: obj) = if value :? int && unbox<int> value = expected then Some() else None
let (|LitBool|_|) (expected: bool) (value: obj) = if value :? bool && unbox<bool> value = expected then Some() else None
let events = ResizeArray<int>()
let sentinel = InvalidOperationException("constructor argument")
let ConstructorTypeApp_track : obj = box (fun (label: obj) -> box (fun (value: obj) -> events.Add(unbox<int> label); value))
let ConstructorTypeApp_explode : obj = box (fun (label: obj) -> events.Add(unbox<int> label); raise sentinel : obj)


type ConstructorNative_Tree =
    | ConstructorNative_Leafusd_Ctor
    | ConstructorNative_Nodeusd_Ctor of int * ConstructorNative_Tree

let ConstructorNative_Leaf_adt_native  : ConstructorNative_Tree = ConstructorNative_Leafusd_Ctor
let ConstructorNative_Leaf : obj = box (ConstructorNative_Leaf_adt_native)

let ConstructorNative_Node_adt_native (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: ConstructorNative_Tree) : ConstructorNative_Tree = ConstructorNative_Nodeusd_Ctor(sharpurs_adt_local_0, sharpurs_adt_local_1)
let ConstructorNative_Node : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (ConstructorNative_Node_adt_native (unbox<int> sharpurs_adt_arg_0) (unbox<ConstructorNative_Tree> sharpurs_adt_arg_1))))

let rec ConstructorNative_depth_adt_native (sharpurs_adt_local_0: ConstructorNative_Tree) : int = (if (match sharpurs_adt_local_0 with | ConstructorNative_Leafusd_Ctor -> true | _ -> false) then (0) else (if (match sharpurs_adt_local_0 with | ConstructorNative_Nodeusd_Ctor(_, _) -> true | _ -> false) then ((1) + (ConstructorNative_depth_adt_native ((match sharpurs_adt_local_0 with | ConstructorNative_Nodeusd_Ctor(_, sharpurs_adt_field) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")))) else (failwith "Failed pattern match")))
let ConstructorNative_depth_adt_native_apply (sharpurs_adt_local_0: ConstructorNative_Tree) : int =
    try ConstructorNative_depth_adt_native sharpurs_adt_local_0
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let ConstructorNative_depth : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (ConstructorNative_depth_adt_native (unbox<ConstructorNative_Tree> sharpurs_adt_arg_0)))
let ConstructorNative_depth_tco (sharpurs_adt_arg_0: obj) : obj = box (ConstructorNative_depth_adt_native (unbox<ConstructorNative_Tree> sharpurs_adt_arg_0))

type ConstructorImported_Envelope =
  | ConstructorImported_Envelopeusd_Ctor of obj

let ConstructorImported_Envelope  = (box ((fun (usd__arg1: obj) -> (box (ConstructorImported_Envelopeusd_Ctor(usd__arg1))))))

module Native =
    type ConstructorTypeApp_Tuple =
      | ConstructorTypeApp_Tupleusd_Ctor of obj * obj
    
    type ConstructorTypeApp_Maybe =
      | ConstructorTypeApp_Nothingusd_Ctor
      | ConstructorTypeApp_Justusd_Ctor of obj
    
    type ConstructorTypeApp_List =
      | ConstructorTypeApp_Nilusd_Ctor
      | ConstructorTypeApp_Consusd_Ctor of obj * obj
    
    type ConstructorTypeApp_Box =
      | ConstructorTypeApp_Boxusd_Ctor of obj
    
    let ConstructorTypeApp_Tuple  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2)))))))
    
    let ConstructorTypeApp_Nothing  = (box ConstructorTypeApp_Nothingusd_Ctor)
    
    let ConstructorTypeApp_Just  = (box ((fun (usd__arg1: obj) -> (box (ConstructorTypeApp_Justusd_Ctor(usd__arg1))))))
    
    let ConstructorTypeApp_Nil  = (box ConstructorTypeApp_Nilusd_Ctor)
    
    let ConstructorTypeApp_Cons  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Consusd_Ctor(usd__arg1, usd__arg2)))))))
    
    let ConstructorTypeApp_Box  = (box ((fun (usd__arg1: obj) -> (box (ConstructorTypeApp_Boxusd_Ctor(usd__arg1))))))
    
    let ConstructorTypeApp_throwSecond  = (box (fun (value: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor((sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 1)))))) (box ((box value)))), (sharpurs_apply (box ((box ConstructorTypeApp_explode))) (box ((box 2)))))))))
    
    let ConstructorTypeApp_throwFirst  = (box (fun (value: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor((sharpurs_apply (box ((box ConstructorTypeApp_explode))) (box ((box 1)))), (sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 2)))))) (box ((box value)))))))))
    
    let ConstructorTypeApp_prepend  = (box (fun (value: obj) -> (box (fun (tail: obj) -> (box (ConstructorTypeApp_Consusd_Ctor((box value), (box tail))))))))
    
    let ConstructorTypeApp_polyPair  = (box (fun (left: obj) -> (box (fun (right: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor((box left), (box right))))))))
    
    let ConstructorTypeApp_partialPair  = (box (fun (left: obj) -> (sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box left))))))
    
    let ConstructorTypeApp_pair  = (box (fun (left: obj) -> (box (fun (right: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor((box left), (box right))))))))
    
    let ConstructorTypeApp_ordinary  = (box (fun (first: obj) -> (box (fun (v: obj) -> (box first)))))
    
    let ConstructorTypeApp_ordinaryCall_direct (first: obj) (second: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_ordinary))) (box ((box first)))))) (box ((box second))))
    
    let ConstructorTypeApp_ordinaryCall_direct_apply (first: obj) (second: obj) : obj =
        try ConstructorTypeApp_ordinaryCall_direct first second
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let ConstructorTypeApp_ordinaryCall = (box (fun (first: obj) -> (box (fun (second: obj) -> (ConstructorTypeApp_ordinaryCall_direct first second)))))
    
    let ConstructorTypeApp_ordered  = (box (fun (left: obj) -> (box (fun (right: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor((sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 1)))))) (box ((box left)))), (sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 2)))))) (box ((box right)))))))))))
    
    let ConstructorTypeApp_nothingInt  = (box ConstructorTypeApp_Nothingusd_Ctor)
    
    let ConstructorTypeApp_nativePair  = (box (fun (value: obj) -> (box (fun (tail: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor((box (ConstructorNative_Node_adt_native (unbox ((box value))) (unbox ((box tail))))), (box tail))))))))
    
    let ConstructorTypeApp_list  = (box (fun (value: obj) -> (box (ConstructorTypeApp_Consusd_Ctor((box value), (box (ConstructorTypeApp_Consusd_Ctor((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box 1)))))), (box ConstructorTypeApp_Nilusd_Ctor)))))))))
    
    let ConstructorTypeApp_justInt  = (box (fun (value: obj) -> (box (ConstructorTypeApp_Justusd_Ctor((box value))))))
    
    let ConstructorTypeApp_importedBox  = (box (fun (value: obj) -> (box (ConstructorImported_Envelopeusd_Ctor((box value))))))
    
    let ConstructorTypeApp_explicitPair  = (box (fun (left: obj) -> (box (fun (right: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor((box left), (box right))))))))
    
    let ConstructorTypeApp_capturedArgument  = (box (fun (first: obj) -> (let saved = (sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 1)))))) (box ((box first))))))) in (box (fun (second: obj) -> (sharpurs_apply (box ((box saved))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 2)))))) (box ((box second))))))))))))
    
    let ConstructorTypeApp_boxInt  = (box (fun (value: obj) -> (box (ConstructorTypeApp_Boxusd_Ctor((box value))))))


module Oracle =
    type ConstructorTypeApp_Tuple =
      | ConstructorTypeApp_Tupleusd_Ctor of obj * obj
    
    type ConstructorTypeApp_Maybe =
      | ConstructorTypeApp_Nothingusd_Ctor
      | ConstructorTypeApp_Justusd_Ctor of obj
    
    type ConstructorTypeApp_List =
      | ConstructorTypeApp_Nilusd_Ctor
      | ConstructorTypeApp_Consusd_Ctor of obj * obj
    
    type ConstructorTypeApp_Box =
      | ConstructorTypeApp_Boxusd_Ctor of obj
    
    let ConstructorTypeApp_Tuple  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2)))))))
    
    let ConstructorTypeApp_Nothing  = (box ConstructorTypeApp_Nothingusd_Ctor)
    
    let ConstructorTypeApp_Just  = (box ((fun (usd__arg1: obj) -> (box (ConstructorTypeApp_Justusd_Ctor(usd__arg1))))))
    
    let ConstructorTypeApp_Nil  = (box ConstructorTypeApp_Nilusd_Ctor)
    
    let ConstructorTypeApp_Cons  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Consusd_Ctor(usd__arg1, usd__arg2)))))))
    
    let ConstructorTypeApp_Box  = (box ((fun (usd__arg1: obj) -> (box (ConstructorTypeApp_Boxusd_Ctor(usd__arg1))))))
    
    let ConstructorTypeApp_throwSecond  = (box (fun (value: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 1)))))) (box ((box value))))))))) (box ((sharpurs_apply (box ((box ConstructorTypeApp_explode))) (box ((box 2)))))))))
    
    let ConstructorTypeApp_throwFirst  = (box (fun (value: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((box ConstructorTypeApp_explode))) (box ((box 1))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 2)))))) (box ((box value)))))))))
    
    let ConstructorTypeApp_prepend  = (box (fun (value: obj) -> (box (fun (tail: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Consusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box value)))))) (box ((box tail))))))))
    
    let ConstructorTypeApp_polyPair  = (box (fun (left: obj) -> (box (fun (right: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box left)))))) (box ((box right))))))))
    
    let ConstructorTypeApp_partialPair  = (box (fun (left: obj) -> (sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box left))))))
    
    let ConstructorTypeApp_pair  = (box (fun (left: obj) -> (box (fun (right: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box left)))))) (box ((box right))))))))
    
    let ConstructorTypeApp_ordinary  = (box (fun (first: obj) -> (box (fun (v: obj) -> (box first)))))
    
    let ConstructorTypeApp_ordinaryCall_direct (first: obj) (second: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_ordinary))) (box ((box first)))))) (box ((box second))))
    
    let ConstructorTypeApp_ordinaryCall_direct_apply (first: obj) (second: obj) : obj =
        try ConstructorTypeApp_ordinaryCall_direct first second
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let ConstructorTypeApp_ordinaryCall = (box (fun (first: obj) -> (box (fun (second: obj) -> (ConstructorTypeApp_ordinaryCall_direct first second)))))
    
    let ConstructorTypeApp_ordered  = (box (fun (left: obj) -> (box (fun (right: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 1)))))) (box ((box left))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 2)))))) (box ((box right)))))))))))
    
    let ConstructorTypeApp_nothingInt  = (box ConstructorTypeApp_Nothingusd_Ctor)
    
    let ConstructorTypeApp_nativePair  = (box (fun (value: obj) -> (box (fun (tail: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box (ConstructorNative_Node_adt_native (unbox ((box value))) (unbox ((box tail)))))))))) (box ((box tail))))))))
    
    let ConstructorTypeApp_list  = (box (fun (value: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Consusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box value)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Consusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box 1))))))))))) (box ((box ConstructorTypeApp_Nilusd_Ctor)))))))))
    
    let ConstructorTypeApp_justInt  = (box (fun (value: obj) -> (sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (ConstructorTypeApp_Justusd_Ctor(usd__arg1)))))))) (box ((box value))))))
    
    let ConstructorTypeApp_importedBox  = (box (fun (value: obj) -> (sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (ConstructorImported_Envelopeusd_Ctor(usd__arg1)))))))) (box ((box value))))))
    
    let ConstructorTypeApp_explicitPair  = (box (fun (left: obj) -> (box (fun (right: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box left)))))) (box ((box right))))))))
    
    let ConstructorTypeApp_capturedArgument  = (box (fun (first: obj) -> (let saved = (sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (ConstructorTypeApp_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 1)))))) (box ((box first))))))) in (box (fun (second: obj) -> (sharpurs_apply (box ((box saved))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ConstructorTypeApp_track))) (box ((box 2)))))) (box ((box second))))))))))))
    
    let ConstructorTypeApp_boxInt  = (box (fun (value: obj) -> (sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (ConstructorTypeApp_Boxusd_Ctor(usd__arg1)))))))) (box ((box value))))))



let mutable checks = 0
let check label condition = if not condition then failwith label else checks <- checks + 1
let apply = sharpurs_apply
let call2 fn a b = apply (apply fn a) b
let pairNative value = match unbox<Native.ConstructorTypeApp_Tuple> value with Native.ConstructorTypeApp_Tupleusd_Ctor (a,b) -> a,b
let pairOracle value = match unbox<Oracle.ConstructorTypeApp_Tuple> value with Oracle.ConstructorTypeApp_Tupleusd_Ctor (a,b) -> a,b
let rec listNative value =
    match unbox<Native.ConstructorTypeApp_List> value with
    | Native.ConstructorTypeApp_Nilusd_Ctor -> []
    | Native.ConstructorTypeApp_Consusd_Ctor (a,tail) -> unbox<int> a :: listNative tail
let rec listOracle value =
    match unbox<Oracle.ConstructorTypeApp_List> value with
    | Oracle.ConstructorTypeApp_Nilusd_Ctor -> []
    | Oracle.ConstructorTypeApp_Consusd_Ctor (a,tail) -> unbox<int> a :: listOracle tail
for value in [System.Int32.MinValue; -1; 0; 1; System.Int32.MaxValue] do
    match unbox<ConstructorImported_Envelope> (apply Native.ConstructorTypeApp_importedBox (box value)) with
    | ConstructorImported_Envelopeusd_Ctor payload -> check "imported polymorphic payload" (unbox<int> payload=value)
    match unbox<ConstructorImported_Envelope> (apply Oracle.ConstructorTypeApp_importedBox (box value)) with
    | ConstructorImported_Envelopeusd_Ctor payload -> check "imported polymorphic oracle payload" (unbox<int> payload=value)
    match unbox<Native.ConstructorTypeApp_Box> (apply Native.ConstructorTypeApp_boxInt (box value)) with
    | Native.ConstructorTypeApp_Boxusd_Ctor payload -> check "single payload" (unbox<int> payload=value)
    match unbox<Oracle.ConstructorTypeApp_Box> (apply Oracle.ConstructorTypeApp_boxInt (box value)) with
    | Oracle.ConstructorTypeApp_Boxusd_Ctor payload -> check "single payload oracle" (unbox<int> payload=value)
    for text in [""; "alpha"; "☃"] do
        for native,baseline in [(Native.ConstructorTypeApp_pair,Oracle.ConstructorTypeApp_pair);(Native.ConstructorTypeApp_polyPair,Oracle.ConstructorTypeApp_polyPair);(Native.ConstructorTypeApp_explicitPair,Oracle.ConstructorTypeApp_explicitPair)] do
            let a,b=pairNative (call2 native (box value) (box text))
            let c,d=pairOracle (call2 baseline (box value) (box text))
            check "two type arguments preserve payloads" (unbox<int> a=value && unbox<string> b=text && a=c && b=d)
        let partial=apply Native.ConstructorTypeApp_partialPair (box value)
        for suffix in [text; text+"!"] do
            let a,b=pairNative (apply partial (box suffix))
            check "partial constructor reused" (unbox<int> a=value && unbox<string> b=suffix)
    match unbox<Native.ConstructorTypeApp_Maybe> (apply Native.ConstructorTypeApp_justInt (box value)) with
    | Native.ConstructorTypeApp_Justusd_Ctor a -> check "Maybe payload" (unbox<int> a=value)
    | _ -> failwith "Expected Just"
    check "List constructor/order/Int32 oracle" (listNative (apply Native.ConstructorTypeApp_list (box value))=listOracle (apply Oracle.ConstructorTypeApp_list (box value)))
    let tail=apply Native.ConstructorTypeApp_list (box value)
    match unbox<Native.ConstructorTypeApp_List> (call2 Native.ConstructorTypeApp_prepend (box 91) tail) with
    | Native.ConstructorTypeApp_Consusd_Ctor (a,b) -> check "recursive payload shares source tail" (unbox<int> a=91 && Object.ReferenceEquals(b,tail))
    | _ -> failwith "Expected Cons"
    check "ordinary polymorphic function unaffected" (call2 Native.ConstructorTypeApp_ordinaryCall (box value) (box 99) |> unbox<int> = value)
match unbox<Native.ConstructorTypeApp_Maybe> Native.ConstructorTypeApp_nothingInt with
| Native.ConstructorTypeApp_Nothingusd_Ctor -> check "nullary constructor remains callable value" true
| _ -> failwith "Expected Nothing"
for fn,decode in [(Native.ConstructorTypeApp_ordered,pairNative);(Oracle.ConstructorTypeApp_ordered,pairOracle);(Native.ConstructorTypeApp_capturedArgument,pairNative);(Oracle.ConstructorTypeApp_capturedArgument,pairOracle)] do
    events.Clear()
    let partial=apply fn (box 7)
    let beforeSecond=List.ofSeq events
    let a,b=decode (apply partial (box 11))
    check "arguments observed once in order" (List.ofSeq events=[1;2] && unbox<int> a=7 && unbox<int> b=11)
    if Object.ReferenceEquals(fn,Native.ConstructorTypeApp_capturedArgument) || Object.ReferenceEquals(fn,Oracle.ConstructorTypeApp_capturedArgument) then
        check "partial captures first argument eagerly" (beforeSecond=[1])
let failure action =
    let caught = try action() |> ignore; None with error -> Some error
    let rec unwrap (error: exn) depth =
        match error with
        | :? System.Reflection.TargetInvocationException as e when not (isNull e.InnerException) -> unwrap e.InnerException (depth+1)
        | e -> e,depth
    match caught with
    | None -> failwith "Expected argument exception"
    | Some e -> unwrap e 0
for native,baseline,expectedEvents in [(Native.ConstructorTypeApp_throwFirst,Oracle.ConstructorTypeApp_throwFirst,[1]);(Native.ConstructorTypeApp_throwSecond,Oracle.ConstructorTypeApp_throwSecond,[1;2])] do
    events.Clear()
    let cause,depth=failure(fun () -> apply native (box 37))
    check "exception argument order" (List.ofSeq events=expectedEvents)
    events.Clear()
    let oldCause,oldDepth=failure(fun () -> apply baseline (box 37))
    check "original exception identity" (Object.ReferenceEquals(cause,sentinel) && Object.ReferenceEquals(oldCause,sentinel))
    check "unchanged exception wrapping depth" (depth=oldDepth && depth=2)
    check "oracle exception argument order" (List.ofSeq events=expectedEvents)
let tail=ConstructorNative_Node_adt_native 13 ConstructorNative_Leaf_adt_native
let boxedTail=box tail
let a,b=pairNative(call2 Native.ConstructorTypeApp_nativePair (box 17) boxedTail)
check "native factory payload type" (a :? ConstructorNative_Tree)
check "native value payload identity" (Object.ReferenceEquals(b,boxedTail))
match unbox<ConstructorNative_Tree> a with
| ConstructorNative_Nodeusd_Ctor (value,child) -> check "native factory field values" (value=17 && Object.ReferenceEquals(child,tail))
| _ -> failwith "Expected native Node"
printfn "constructor-typeapp runtime: %d checks passed" checks
