[<AutoOpen>]
module PureScript_Control_Comonad_Env

open System
open System.Collections.Generic

let Control_Comonad_Env_unwrap  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Control_Comonad_Env_withEnv  = (box Control_Comonad_Env_Trans_withEnvT)

let Control_Comonad_Env_runEnv  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | x -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Tuple_functorTuple)))))) (box ((box Control_Comonad_Env_unwrap)))))) (box ((box x))))))))

let Control_Comonad_Env_mapEnv  = (sharpurs_apply (box ((box Data_Functor_map))) (box ((sharpurs_apply (box ((box Control_Comonad_Env_Trans_functorEnvT))) (box ((box Data_Identity_functorIdentity)))))))

let Control_Comonad_Env_env  = (box (fun (e: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Control_Comonad_Env_Trans_EnvT)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box e))))))))) (box ((sharpurs_apply (box ((box Data_Identity_Identity))) (box ((box a))))))))))))))
