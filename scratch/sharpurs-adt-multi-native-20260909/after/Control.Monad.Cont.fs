[<AutoOpen>]
module PureScript_Control_Monad_Cont

open System
open System.Collections.Generic

let Control_Monad_Cont_unwrap  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Control_Monad_Cont_withCont  = (box (fun (f: obj) -> (sharpurs_apply (box ((box Control_Monad_Cont_Trans_withContT))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Identity_Identity))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Control_Monad_Cont_unwrap)))))))))))))))

let Control_Monad_Cont_runCont  = (box (fun (cc: obj) -> (box (fun (k: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_Cont_Trans_runContT))) (box ((box cc)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Identity_Identity)))))) (box ((box k))))))))))))))

let Control_Monad_Cont_mapCont  = (box (fun (f: obj) -> (sharpurs_apply (box ((box Control_Monad_Cont_Trans_mapContT))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Identity_Identity)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f)))))) (box ((box Control_Monad_Cont_unwrap))))))))))))

let Control_Monad_Cont_cont  = (box (fun (f: obj) -> (sharpurs_apply (box ((box Control_Monad_Cont_Trans_ContT))) (box ((box (fun (c: obj) -> (sharpurs_apply (box ((box Data_Identity_Identity))) (box ((sharpurs_apply (box ((box f))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Control_Monad_Cont_unwrap)))))) (box ((box c)))))))))))))))))
