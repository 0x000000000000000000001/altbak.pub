[<AutoOpen>]
module PureScript_Control_Monad_Reader

open System
open System.Collections.Generic

let Control_Monad_Reader_unwrap  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Control_Monad_Reader_withReader  = (box Control_Monad_Reader_Trans_withReaderT)

let Control_Monad_Reader_runReader  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | m -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Control_Monad_Reader_unwrap)))))) (box ((box m))))))))

let Control_Monad_Reader_mapReader  = (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Control_Monad_Reader_Trans_mapReaderT)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Identity_Identity)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f)))))) (box ((box Control_Monad_Reader_unwrap))))))))))))
