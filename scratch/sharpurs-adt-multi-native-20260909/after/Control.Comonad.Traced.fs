[<AutoOpen>]
module PureScript_Control_Comonad_Traced

open System
open System.Collections.Generic

let Control_Comonad_Traced_traced  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_composeFlipped))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Identity_Identity)))))) (box ((box Control_Comonad_Traced_Trans_TracedT))))

let Control_Comonad_Traced_runTraced  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | t -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined)))))) (box ((box t))))))))
