[<AutoOpen>]
module PureScript_Control_Monad_State

open System
open System.Collections.Generic

let Control_Monad_State_unwrap  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Control_Monad_State_withState  = (box Control_Monad_State_Trans_withStateT)

let Control_Monad_State_runState  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | s -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Control_Monad_State_unwrap)))))) (box ((box s))))))))

let Control_Monad_State_mapState  = (box (fun (f: obj) -> (sharpurs_apply (box ((box Control_Monad_State_Trans_mapStateT))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Identity_Identity)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f)))))) (box ((box Control_Monad_State_unwrap))))))))))))

let Control_Monad_State_execState  = (box (fun (v: obj) -> (box (fun (s: obj) -> (match (((unbox ((box v))), (unbox ((box s))))) with | (m, s1) -> ((let v1 = (sharpurs_apply (box ((box m))) (box ((box s1)))) in (match ((unbox ((box v1)))) with | Data_Tuple_Tupleusd_Ctor(_, s_prime) -> ((box s_prime))))))))))

let Control_Monad_State_evalState  = (box (fun (v: obj) -> (box (fun (s: obj) -> (match (((unbox ((box v))), (unbox ((box s))))) with | (m, s1) -> ((let v1 = (sharpurs_apply (box ((box m))) (box ((box s1)))) in (match ((unbox ((box v1)))) with | Data_Tuple_Tupleusd_Ctor(a, _) -> ((box a))))))))))
