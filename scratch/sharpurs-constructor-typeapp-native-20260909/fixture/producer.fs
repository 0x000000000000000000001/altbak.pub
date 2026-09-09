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