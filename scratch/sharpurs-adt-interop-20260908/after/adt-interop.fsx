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


let (|LitInt|_|) (expected: int) (value: obj) = if value :? int && unbox value = expected then Some() else None


type AdtPilot_Color =
    | AdtPilot_Rusd_Ctor
    | AdtPilot_Busd_Ctor
and AdtPilot_Tree =
    | AdtPilot_Eusd_Ctor
    | AdtPilot_Tusd_Ctor of AdtPilot_Color * AdtPilot_Tree * int * AdtPilot_Tree

let AdtPilot_R_adt_native  : AdtPilot_Color = AdtPilot_Rusd_Ctor
let AdtPilot_R : obj = box (AdtPilot_R_adt_native)

let AdtPilot_B_adt_native  : AdtPilot_Color = AdtPilot_Busd_Ctor
let AdtPilot_B : obj = box (AdtPilot_B_adt_native)

let AdtPilot_E_adt_native  : AdtPilot_Tree = AdtPilot_Eusd_Ctor
let AdtPilot_E : obj = box (AdtPilot_E_adt_native)

let AdtPilot_T_adt_native (sharpurs_adt_local_0: AdtPilot_Color) (sharpurs_adt_local_1: AdtPilot_Tree) (sharpurs_adt_local_2: int) (sharpurs_adt_local_3: AdtPilot_Tree) : AdtPilot_Tree = AdtPilot_Tusd_Ctor(sharpurs_adt_local_0, sharpurs_adt_local_1, sharpurs_adt_local_2, sharpurs_adt_local_3)
let AdtPilot_T : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (fun (sharpurs_adt_arg_2: obj) -> box (fun (sharpurs_adt_arg_3: obj) -> box (AdtPilot_T_adt_native (unbox<AdtPilot_Color> sharpurs_adt_arg_0) (unbox<AdtPilot_Tree> sharpurs_adt_arg_1) (unbox<int> sharpurs_adt_arg_2) (unbox<AdtPilot_Tree> sharpurs_adt_arg_3))))))

let AdtPilot_singletonWith_adt_native (sharpurs_adt_local_0: AdtPilot_Color) (sharpurs_adt_local_1: int) : AdtPilot_Tree = AdtPilot_Tusd_Ctor(sharpurs_adt_local_0, AdtPilot_Eusd_Ctor, sharpurs_adt_local_1, AdtPilot_Eusd_Ctor)
let AdtPilot_singletonWith : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (AdtPilot_singletonWith_adt_native (unbox<AdtPilot_Color> sharpurs_adt_arg_0) (unbox<int> sharpurs_adt_arg_1))))

let AdtPilot_singleton_adt_native  : AdtPilot_Tree = AdtPilot_Tusd_Ctor(AdtPilot_Busd_Ctor, AdtPilot_Eusd_Ctor, (2147483647), AdtPilot_Eusd_Ctor)
let AdtPilot_singleton : obj = box (AdtPilot_singleton_adt_native)

let AdtPilot_rootValue_adt_native (sharpurs_adt_local_0: AdtPilot_Tree) : int = (if (match sharpurs_adt_local_0 with | AdtPilot_Eusd_Ctor -> true | _ -> false) then (0) else (if (match sharpurs_adt_local_0 with | AdtPilot_Tusd_Ctor(_, _, _, _) -> true | _ -> false) then (match sharpurs_adt_local_0 with | AdtPilot_Tusd_Ctor(_, _, sharpurs_adt_field, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor") else (failwith "Failed pattern match")))
let AdtPilot_rootValue : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (AdtPilot_rootValue_adt_native (unbox<AdtPilot_Tree> sharpurs_adt_arg_0)))

let AdtPilot_rootColor_adt_native (sharpurs_adt_local_0: AdtPilot_Tree) : AdtPilot_Color = (if (match sharpurs_adt_local_0 with | AdtPilot_Eusd_Ctor -> true | _ -> false) then AdtPilot_Busd_Ctor else (if (match sharpurs_adt_local_0 with | AdtPilot_Tusd_Ctor(_, _, _, _) -> true | _ -> false) then (match sharpurs_adt_local_0 with | AdtPilot_Tusd_Ctor(sharpurs_adt_field, _, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor") else (failwith "Failed pattern match")))
let AdtPilot_rootColor : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (AdtPilot_rootColor_adt_native (unbox<AdtPilot_Tree> sharpurs_adt_arg_0)))

let AdtPilot_max_adt_native (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: int) : int = (if (sharpurs_adt_local_0 > sharpurs_adt_local_1) then sharpurs_adt_local_0 else sharpurs_adt_local_1)
let AdtPilot_max : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (AdtPilot_max_adt_native (unbox<int> sharpurs_adt_arg_0) (unbox<int> sharpurs_adt_arg_1))))

let AdtPilot_leftChild_adt_native (sharpurs_adt_local_0: AdtPilot_Tree) : AdtPilot_Tree = (if (match sharpurs_adt_local_0 with | AdtPilot_Eusd_Ctor -> true | _ -> false) then AdtPilot_Eusd_Ctor else (if (match sharpurs_adt_local_0 with | AdtPilot_Tusd_Ctor(_, _, _, _) -> true | _ -> false) then (match sharpurs_adt_local_0 with | AdtPilot_Tusd_Ctor(_, sharpurs_adt_field, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor") else (failwith "Failed pattern match")))
let AdtPilot_leftChild : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (AdtPilot_leftChild_adt_native (unbox<AdtPilot_Tree> sharpurs_adt_arg_0)))

let AdtPilot_isRed_adt_native (sharpurs_adt_local_0: AdtPilot_Color) : bool = (if (match sharpurs_adt_local_0 with | AdtPilot_Rusd_Ctor -> true | _ -> false) then true else (if (match sharpurs_adt_local_0 with | AdtPilot_Busd_Ctor -> true | _ -> false) then false else (failwith "Failed pattern match")))
let AdtPilot_isRed : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (AdtPilot_isRed_adt_native (unbox<AdtPilot_Color> sharpurs_adt_arg_0)))

let AdtPilot_empty_adt_native  : AdtPilot_Tree = AdtPilot_Eusd_Ctor
let AdtPilot_empty : obj = box (AdtPilot_empty_adt_native)

let rec AdtPilot_depth_adt_native (sharpurs_adt_local_0: AdtPilot_Tree) : int = (if (match sharpurs_adt_local_0 with | AdtPilot_Eusd_Ctor -> true | _ -> false) then (0) else (if (match sharpurs_adt_local_0 with | AdtPilot_Tusd_Ctor(_, _, _, _) -> true | _ -> false) then ((1) + (let sharpurs_adt_local_1: int = (AdtPilot_depth_adt_native ((match sharpurs_adt_local_0 with | AdtPilot_Tusd_Ctor(_, sharpurs_adt_field, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor"))) in (let sharpurs_adt_local_2: int = (AdtPilot_depth_adt_native ((match sharpurs_adt_local_0 with | AdtPilot_Tusd_Ctor(_, _, _, sharpurs_adt_field) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor"))) in (if (sharpurs_adt_local_1 > sharpurs_adt_local_2) then sharpurs_adt_local_1 else sharpurs_adt_local_2)))) else (failwith "Failed pattern match")))
let AdtPilot_depth : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (AdtPilot_depth_adt_native (unbox<AdtPilot_Tree> sharpurs_adt_arg_0)))

let AdtPilot_asymmetric_adt_native  : AdtPilot_Tree = AdtPilot_Tusd_Ctor(AdtPilot_Busd_Ctor, AdtPilot_Tusd_Ctor(AdtPilot_Rusd_Ctor, AdtPilot_Eusd_Ctor, (-2147483648), AdtPilot_Eusd_Ctor), (0), AdtPilot_Tusd_Ctor(AdtPilot_Busd_Ctor, AdtPilot_Eusd_Ctor, (42), AdtPilot_Tusd_Ctor(AdtPilot_Rusd_Ctor, AdtPilot_Eusd_Ctor, (42), AdtPilot_Eusd_Ctor)))
let AdtPilot_asymmetric : obj = box (AdtPilot_asymmetric_adt_native)




let events = ResizeArray<int>()
let track : obj = box (fun (label: obj) -> box (fun (value: obj) -> events.Add(unbox<int> label); value))
let AdtConsumer_trackColor = track
let AdtConsumer_trackTree = track
let AdtConsumer_trackInt = track


type AdtConsumer_ConsumerBox =
  | AdtConsumer_ConsumerBoxusd_Ctor of obj * obj

let AdtConsumer_ConsumerBox  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (AdtConsumer_ConsumerBoxusd_Ctor(usd__arg1, usd__arg2)))))))

let AdtConsumer_wrap  = (box (fun (child: obj) -> (box (AdtConsumer_ConsumerBoxusd_Ctor((box child), (box 42))))))

let AdtConsumer_unwrap  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtConsumer_ConsumerBoxusd_Ctor(child, _) -> ((box child)))))

let AdtConsumer_shared  = (box (fun (child: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((box AdtPilot_B)))))) (box ((box child)))))) (box ((box 11)))))) (box ((box child))))))

let AdtConsumer_saturated  = (box (fun (color: obj) -> (box (fun (left: obj) -> (box (fun (value: obj) -> (box (fun (right: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((box color)))))) (box ((box left)))))) (box ((box value)))))) (box ((box right))))))))))))

let AdtConsumer_roundTrip  = (box (fun (child: obj) -> (sharpurs_apply (box ((box AdtPilot_depth))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((box AdtPilot_B)))))) (box ((box child)))))) (box ((box 11)))))) (box ((box child)))))))))

let AdtConsumer_rootValue  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box 0)) | AdtPilot_Tusd_Ctor(_, _, value, _) -> ((box value)))))

let AdtConsumer_rootColor  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box AdtPilot_B)) | AdtPilot_Tusd_Ctor(color, _, _, _) -> ((box color)))))

let AdtConsumer_partial  = (sharpurs_apply (box ((box AdtPilot_T))) (box ((box AdtPilot_B))))

let AdtConsumer_orderedPartial  = (box (fun (child: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackColor))) (box ((box 1)))))) (box ((box AdtPilot_B))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackTree))) (box ((box 2)))))) (box ((box child)))))))))

let AdtConsumer_orderedConstruction  = (box (fun (child: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackColor))) (box ((box 1)))))) (box ((box AdtPilot_B))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackTree))) (box ((box 2)))))) (box ((box child))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackInt))) (box ((box 3)))))) (box ((box 11))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_trackTree))) (box ((box 4)))))) (box ((box child)))))))))

let AdtConsumer_namedChild  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Tusd_Ctor(_, Unbox((AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Rusd_Ctor), _, _, _) as child)), _, _) -> ((box child)) | _ -> ((box AdtPilot_E)))))

let AdtConsumer_leftChild  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box AdtPilot_E)) | AdtPilot_Tusd_Ctor(_, child, _, _) -> ((box child)))))

let AdtConsumer_deepPattern  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Busd_Ctor), Unbox(AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Rusd_Ctor), Unbox(AdtPilot_Eusd_Ctor), Unbox(LitInt 7 ()), Unbox(AdtPilot_Eusd_Ctor))), Unbox(LitInt 11 ()), Unbox(AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Busd_Ctor), Unbox(AdtPilot_Eusd_Ctor), value, Unbox(AdtPilot_Eusd_Ctor)))) -> ((box value)) | _ -> ((box 0)))))

let AdtConsumer_boxedValue  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtConsumer_ConsumerBoxusd_Ctor(_, value) -> ((box value)))))

let AdtConsumer_boxedPattern  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtConsumer_ConsumerBoxusd_Ctor(Unbox(AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Busd_Ctor), _, Unbox(LitInt 11 ()), _)), Unbox(LitInt 42 ())) -> ((box 7)) | _ -> ((box 0)))))

let AdtConsumer_applyValue  = (box (fun (f: obj) -> (box (fun (value: obj) -> (sharpurs_apply (box ((box f))) (box ((box value))))))))

let AdtConsumer_throughGeneric  = (box (fun (value: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_applyValue))) (box ((box AdtPilot_T)))))) (box ((box AdtPilot_R)))))) (box ((box AdtPilot_E)))))) (box ((box value)))))) (box ((box AdtPilot_E))))))

let AdtConsumer_throughPartial  = (box (fun (value: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtConsumer_applyValue))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtPilot_T))) (box ((box AdtPilot_B)))))) (box ((box AdtPilot_E))))))))) (box ((box value)))))) (box ((box AdtPilot_E))))))


let mutable checks = 0
let check label ok = if not ok then failwith label else checks <- checks + 1
let apply = sharpurs_apply
let make fn color left value right = apply (apply (apply (apply fn color) (box left)) (box value)) (box right) |> unbox<AdtPilot_Tree>
let leaf value = AdtPilot_T_adt_native AdtPilot_R_adt_native AdtPilot_E_adt_native value AdtPilot_E_adt_native
let red = leaf 7
let made = make AdtConsumer_saturated AdtPilot_B red 11 red
check "saturated consumer native depth" (AdtPilot_depth_adt_native made = 2)
check "saturated consumer Int" (AdtPilot_rootValue_adt_native made = 11)
check "saturated consumer left identity" (System.Object.ReferenceEquals(AdtPilot_leftChild_adt_native made, red))
let partial = apply AdtConsumer_partial (box red)
let make11 = apply partial (box 11)
let first = apply make11 (box red) |> unbox<AdtPilot_Tree>
let second = apply make11 AdtPilot_E |> unbox<AdtPilot_Tree>
check "partial constructor reusable first" (AdtPilot_depth_adt_native first = 2)
check "partial constructor reusable second" (AdtPilot_depth_adt_native second = 2)
check "partial constructor retained identity" (System.Object.ReferenceEquals(AdtPilot_leftChild_adt_native second, red))
for value in [System.Int32.MinValue; -1; 0; 7; System.Int32.MaxValue] do
    for fn in [AdtConsumer_throughGeneric; AdtConsumer_throughPartial] do
        let tree = apply fn (box value) |> unbox<AdtPilot_Tree>
        check "constructor as polymorphic value Int" (AdtPilot_rootValue_adt_native tree = value)
        check "consumer projection native Int" (unbox<int> (apply AdtConsumer_rootValue (box tree)) = value)
        check "constructor as polymorphic value depth" (AdtPilot_depth_adt_native tree = 1)
check "consumer left projection identity" (System.Object.ReferenceEquals(apply AdtConsumer_leftChild (box made), red))
check "consumer Color projection" (unbox<AdtPilot_Color> (apply AdtConsumer_rootColor (box made)) = AdtPilot_B_adt_native)
let black value = AdtPilot_T_adt_native AdtPilot_B_adt_native AdtPilot_E_adt_native value AdtPilot_E_adt_native
let deep = AdtPilot_T_adt_native AdtPilot_B_adt_native red 11 (black 99)
check "nested color/tree/Int pattern hit" (unbox<int> (apply AdtConsumer_deepPattern (box deep)) = 99)
check "nested Int pattern miss" (unbox<int> (apply AdtConsumer_deepPattern (box (AdtPilot_T_adt_native AdtPilot_B_adt_native (leaf 8) 11 (black 99)))) = 0)
check "nested color pattern miss" (unbox<int> (apply AdtConsumer_deepPattern (box (AdtPilot_T_adt_native AdtPilot_R_adt_native red 11 (black 99)))) = 0)
check "nested tree pattern empty miss" (unbox<int> (apply AdtConsumer_deepPattern AdtPilot_E) = 0)
check "named child projection identity" (System.Object.ReferenceEquals(apply AdtConsumer_namedChild (box deep), red))
check "named child fallback" (unbox<AdtPilot_Tree> (apply AdtConsumer_namedChild AdtPilot_E) = AdtPilot_E_adt_native)
let shared = apply AdtConsumer_shared (box red) |> unbox<AdtPilot_Tree>
match shared with
| AdtPilot_Tusd_Ctor(_, left, _, right) ->
    check "shared children identity" (System.Object.ReferenceEquals(left, right))
    check "shared input identity" (System.Object.ReferenceEquals(left, red))
| _ -> failwith "shared consumer value was not a node"
check "producer native depth accepts consumer result" (AdtPilot_depth_adt_native shared = 2)
check "consumer calls producer wrapper" (unbox<int> (apply AdtConsumer_roundTrip (box red)) = 2)
check "input retains original value" (AdtPilot_rootValue_adt_native red = 7)
let boxed = apply AdtConsumer_wrap (box made)
check "local boxed ADT retains native child identity" (System.Object.ReferenceEquals(apply AdtConsumer_unwrap boxed, made))
check "local boxed ADT retains boxed Int" (unbox<int> (apply AdtConsumer_boxedValue boxed) = 42)
check "boxed outer and native inner pattern hit" (unbox<int> (apply AdtConsumer_boxedPattern boxed) = 7)
check "boxed outer and native inner pattern miss" (unbox<int> (apply AdtConsumer_boxedPattern (apply AdtConsumer_wrap (box red))) = 0)
match unbox<AdtConsumer_ConsumerBox> boxed with
| AdtConsumer_ConsumerBoxusd_Ctor(child, value) ->
    check "local boxed constructor retains object fields" ((child :? AdtPilot_Tree) && (value :? int))
events.Clear()
let ordered = apply AdtConsumer_orderedConstruction (box red) |> unbox<AdtPilot_Tree>
check "saturated arguments evaluate left to right once" (List.ofSeq events = [1; 2; 3; 4])
check "ordered construction native depth" (AdtPilot_depth_adt_native ordered = 2)
events.Clear()
let orderedPartial = apply AdtConsumer_orderedPartial (box red)
check "supplied partial arguments evaluate eagerly" (List.ofSeq events = [1; 2])
let reusedPartial = apply orderedPartial (box 17)
check "partial retains arguments without reevaluation" (List.ofSeq events = [1; 2])
let orderedFirst = apply reusedPartial (box red) |> unbox<AdtPilot_Tree>
let orderedSecond = apply reusedPartial AdtPilot_E |> unbox<AdtPilot_Tree>
check "repeated partial application keeps captured argument evaluations" (List.ofSeq events = [1; 2])
check "partial first result retains native Int" (AdtPilot_rootValue_adt_native orderedFirst = 17)
check "partial second result retains native Int" (AdtPilot_rootValue_adt_native orderedSecond = 17)
check "partial result retains left input identity" (System.Object.ReferenceEquals(AdtPilot_leftChild_adt_native orderedSecond, red))
printfn "adt-interop runtime: %d checks passed" checks
