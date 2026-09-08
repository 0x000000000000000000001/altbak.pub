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



let AdtConsumer_rootValue  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box 0)) | AdtPilot_Tusd_Ctor(_, _, value, _) -> ((box value)))))

let AdtConsumer_rootColor  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box AdtPilot_Busd_Ctor)) | AdtPilot_Tusd_Ctor(color, _, _, _) -> ((box color)))))

let AdtConsumer_namedChild  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Tusd_Ctor(_, Unbox((AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Rusd_Ctor), _, _, _) as child)), _, _) -> ((box child)) | _ -> ((box AdtPilot_Eusd_Ctor)))))

let AdtConsumer_leftChild  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Eusd_Ctor -> ((box AdtPilot_Eusd_Ctor)) | AdtPilot_Tusd_Ctor(_, child, _, _) -> ((box child)))))

let AdtConsumer_deepPattern  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Busd_Ctor), Unbox(AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Rusd_Ctor), Unbox(AdtPilot_Eusd_Ctor), Unbox(LitInt 7 ()), Unbox(AdtPilot_Eusd_Ctor))), Unbox(LitInt 11 ()), Unbox(AdtPilot_Tusd_Ctor(Unbox(AdtPilot_Busd_Ctor), Unbox(AdtPilot_Eusd_Ctor), value, Unbox(AdtPilot_Eusd_Ctor)))) -> ((box value)) | _ -> ((box 0)))))


let red = AdtPilot_T_adt_native AdtPilot_R_adt_native AdtPilot_E_adt_native 7 AdtPilot_E_adt_native
let black = AdtPilot_T_adt_native AdtPilot_B_adt_native AdtPilot_E_adt_native 99 AdtPilot_E_adt_native
let deep = AdtPilot_T_adt_native AdtPilot_B_adt_native red 11 black
let actual = sharpurs_apply AdtConsumer_deepPattern (box deep) |> unbox<int>
if actual <> 99 then failwithf "pattern mismatch: %d" actual
if not(System.Object.ReferenceEquals(sharpurs_apply AdtConsumer_namedChild (box deep), red)) then failwith "identity mismatch"
printfn "patterns-only runtime: 2 checks passed"
