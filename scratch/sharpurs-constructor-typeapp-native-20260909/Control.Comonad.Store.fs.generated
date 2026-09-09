[<AutoOpen>]
module PureScript_Control_Comonad_Store

open System
open System.Collections.Generic

let Control_Comonad_Store_unwrap  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Control_Comonad_Store_store  = (box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Control_Comonad_Store_Trans_StoreT)))))) (box ((box (Data_Tuple_Tupleusd_Ctor((sharpurs_apply (box ((box Data_Identity_Identity))) (box ((box f)))), (box x)))))))))))

let Control_Comonad_Store_runStore  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | s -> ((sharpurs_apply (box ((box Data_Tuple_swap))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Tuple_functorTuple)))))) (box ((box Control_Comonad_Store_unwrap)))))) (box ((sharpurs_apply (box ((box Data_Tuple_swap))) (box ((box s))))))))))))))
