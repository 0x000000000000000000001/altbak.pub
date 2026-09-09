[<AutoOpen>]
module PureScript_Test_StateMonad

open System
open System.Collections.Generic

let Test_StateMonad_State  = (box (fun (x: obj) -> (box x)))

let Test_StateMonad_runState  = (box (fun (v: obj) -> (box (fun (s: obj) -> (match (((unbox ((box v))), (unbox ((box s))))) with | (f, s1) -> ((sharpurs_apply (box ((box f))) (box ((box s1))))))))))

let Test_StateMonad_put  = (box (fun (s: obj) -> (sharpurs_apply (box ((box Test_StateMonad_State))) (box ((box (fun (v: obj) -> (box ((Map.add "val" (box ((box Data_Unit_unit))) (Map.add "state" (box ((box s))) Map.empty)))))))))))

let Test_StateMonad_pureState  = (box (fun (a: obj) -> (sharpurs_apply (box ((box Test_StateMonad_State))) (box ((box (fun (s: obj) -> (box ((Map.add "val" (box ((box a))) (Map.add "state" (box ((box s))) Map.empty)))))))))))

let Test_StateMonad_get  = (sharpurs_apply (box ((box Test_StateMonad_State))) (box ((box (fun (s: obj) -> (box ((Map.add "val" (box ((box s))) (Map.add "state" (box ((box s))) Map.empty)))))))))

let Test_StateMonad_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "State Monad (1.2k Binds, 60 Stack Depth):"))))

let Test_StateMonad_bindState  = (box (fun (v: obj) -> (box (fun (g: obj) -> (match (((unbox ((box v))), (unbox ((box g))))) with | (f, g1) -> ((sharpurs_apply (box ((box Test_StateMonad_State))) (box ((box (fun (s: obj) -> (let r1 = (sharpurs_apply (box ((box f))) (box ((box s)))) in (let v1 = (sharpurs_apply (box ((box g1))) (box ((Map.find "val" (unbox<Map<string, obj>> ((box r1))))))) in (match ((unbox ((box v1)))) with | g_prime -> ((sharpurs_apply (box ((box g_prime))) (box ((Map.find "state" (unbox<Map<string, obj>> ((box r1))))))))))))))))))))))

let Test_StateMonad_modify  = (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Test_StateMonad_bindState))) (box ((box Test_StateMonad_get)))))) (box ((box (fun (s: obj) -> (sharpurs_apply (box ((box Test_StateMonad_put))) (box ((sharpurs_apply (box ((box f))) (box ((box s))))))))))))))

let rec Test_StateMonad_chainModifications_tco (v: obj) : obj = ((match ((unbox ((box v)))) with | LitInt 0 () -> ((sharpurs_apply (box ((box Test_StateMonad_pureState))) (box ((box Data_Unit_unit))))) | n -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_StateMonad_bindState))) (box ((sharpurs_apply (box ((box Test_StateMonad_modify))) (box ((box (fun (x: obj) -> (box ((unbox<int> (box ((box x)))) + (unbox<int> (box ((box 1)))))))))))))))) (box ((box (fun (v1: obj) -> (Test_StateMonad_chainModifications_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))))))))))))
and Test_StateMonad_chainModifications = box (fun (v: obj) -> Test_StateMonad_chainModifications_tco v)


let rec Test_StateMonad_runManyTimes_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), acc) -> ((box acc)) | (n, acc) -> ((Test_StateMonad_runManyTimes_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((box ((unbox<int> (box ((box acc)))) + (unbox<int> (box ((Map.find "state" (unbox<Map<string, obj>> ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_StateMonad_runState))) (box ((sharpurs_apply (box ((box Test_StateMonad_chainModifications))) (box ((box 60))))))))) (box ((box 0)))))))))))))))))
and Test_StateMonad_runManyTimes = box (fun (v: obj) ->  (fun (v1: obj) -> Test_StateMonad_runManyTimes_tco v v1))


let Test_StateMonad_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 20))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_StateMonad_runManyTimes))) (box ((box dummy)))))) (box ((box 0)))))))))))))))
