[<AutoOpen>]
module PureScript_Control_Monad_State_Class

open System
open System.Collections.Generic

let Control_Monad_State_Class_MonadStateusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Monad_State_Class_state  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "state" (unbox<Map<string, obj>> ((box v))))))))

let Control_Monad_State_Class_put  = (box (fun (dictMonadState: obj) -> (box (fun (s: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_State_Class_state))) (box ((box dictMonadState)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box Data_Unit_unit)))))) (box ((box s)))))))))))))

let Control_Monad_State_Class_modify_  = (box (fun (dictMonadState: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_State_Class_state))) (box ((box dictMonadState)))))) (box ((box (fun (s: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box Data_Unit_unit)))))) (box ((sharpurs_apply (box ((box f))) (box ((box s))))))))))))))))

let Control_Monad_State_Class_modify  = (box (fun (dictMonadState: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_State_Class_state))) (box ((box dictMonadState)))))) (box ((box (fun (s: obj) -> (let s_prime = (sharpurs_apply (box ((box f))) (box ((box s)))) in (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box s_prime)))))) (box ((box s_prime))))))))))))))

let Control_Monad_State_Class_gets  = (box (fun (dictMonadState: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_State_Class_state))) (box ((box dictMonadState)))))) (box ((box (fun (s: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((box f))) (box ((box s))))))))) (box ((box s)))))))))))))

let Control_Monad_State_Class_get  = (box (fun (dictMonadState: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_State_Class_state))) (box ((box dictMonadState)))))) (box ((box (fun (s: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box s)))))) (box ((box s)))))))))))
