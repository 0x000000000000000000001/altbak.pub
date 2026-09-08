module AdtPilotGenerated

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

