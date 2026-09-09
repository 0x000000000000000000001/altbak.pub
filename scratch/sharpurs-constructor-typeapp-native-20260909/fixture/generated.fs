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