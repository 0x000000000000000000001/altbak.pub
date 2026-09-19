[<AutoOpen>]
module PureScript_Control_Monad_RWS

open System
open System.Collections.Generic

let Control_Monad_RWS_pure  = (sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Data_Identity_applicativeIdentity))))

let Control_Monad_RWS_unwrap  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Control_Monad_RWS_unwrap1  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Control_Monad_RWS_unwrap2  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Control_Monad_RWS_withRWS  = (box Control_Monad_RWS_Trans_withRWST)

let Control_Monad_RWS_rws  = (box (fun (f: obj) -> (sharpurs_apply (box ((box Control_Monad_RWS_Trans_RWST))) (box ((box (fun (r: obj) -> (box (fun (s: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Control_Monad_RWS_pure)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box f))) (box ((box r)))))) (box ((box s))))))))))))))))

let Control_Monad_RWS_runRWS  = (box (fun (m: obj) -> (box (fun (r: obj) -> (box (fun (s: obj) -> (match ((unbox ((box m)))) with | f -> ((let v = (sharpurs_apply (box ((sharpurs_apply (box ((box f))) (box ((box r)))))) (box ((box s)))) in (match ((unbox ((box v)))) with | x -> ((box x))))))))))))

let Control_Monad_RWS_mapRWS  = (box (fun (f: obj) -> (sharpurs_apply (box ((box Control_Monad_RWS_Trans_mapRWST))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_composeFlipped))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Control_Monad_RWS_unwrap)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_composeFlipped))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f)))))) (box ((box Data_Identity_Identity))))))))))))

let Control_Monad_RWS_execRWS  = (box (fun (m: obj) -> (box (fun (r: obj) -> (box (fun (s: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Control_Monad_RWS_unwrap1)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_RWS_Trans_execRWST))) (box ((box Data_Identity_monadIdentity)))))) (box ((box m)))))) (box ((box r)))))) (box ((box s)))))))))))))

let Control_Monad_RWS_evalRWS  = (box (fun (m: obj) -> (box (fun (r: obj) -> (box (fun (s: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Control_Monad_RWS_unwrap2)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_RWS_Trans_evalRWST))) (box ((box Data_Identity_monadIdentity)))))) (box ((box m)))))) (box ((box r)))))) (box ((box s)))))))))))))
