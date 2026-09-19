[<AutoOpen>]
module PureScript_Control_Monad_Except

open System
open System.Collections.Generic

let Control_Monad_Except_unwrap  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Control_Monad_Except_withExcept  = (sharpurs_apply (box ((box Control_Monad_Except_Trans_withExceptT))) (box ((box Data_Identity_functorIdentity))))

let Control_Monad_Except_runExcept  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Control_Monad_Except_unwrap)))))) (box ((box Control_Monad_Except_Trans_runExceptT))))

let Control_Monad_Except_mapExcept  = (box (fun (f: obj) -> (sharpurs_apply (box ((box Control_Monad_Except_Trans_mapExceptT))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Identity_Identity)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f)))))) (box ((box Control_Monad_Except_unwrap))))))))))))
