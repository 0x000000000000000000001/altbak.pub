[<AutoOpen>]
module PureScript_Data_Traversable_Accum_Internal

open System
open System.Collections.Generic

let Data_Traversable_Accum_Internal_StateR  = (box (fun (x: obj) -> (box x)))

let Data_Traversable_Accum_Internal_StateL  = (box (fun (x: obj) -> (box x)))

let Data_Traversable_Accum_Internal_stateR  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | k -> ((box k)))))

let Data_Traversable_Accum_Internal_stateL  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | k -> ((box k)))))

let Data_Traversable_Accum_Internal_functorStateR  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (f: obj) -> (box (fun (k: obj) -> (sharpurs_apply (box ((box Data_Traversable_Accum_Internal_StateR))) (box ((box (fun (s: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Traversable_Accum_Internal_stateR))) (box ((box k)))))) (box ((box s)))) in (match ((unbox ((box v)))) with | (HasProp "accum" (s1) & HasProp "value" (a)) -> ((box ((Map.add "accum" (box ((box s1))) (Map.add "value" (box ((sharpurs_apply (box ((box f))) (box ((box a)))))) Map.empty)))))))))))))))))) Map.empty))))))

let Data_Traversable_Accum_Internal_functorStateL  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (f: obj) -> (box (fun (k: obj) -> (sharpurs_apply (box ((box Data_Traversable_Accum_Internal_StateL))) (box ((box (fun (s: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Traversable_Accum_Internal_stateL))) (box ((box k)))))) (box ((box s)))) in (match ((unbox ((box v)))) with | (HasProp "accum" (s1) & HasProp "value" (a)) -> ((box ((Map.add "accum" (box ((box s1))) (Map.add "value" (box ((sharpurs_apply (box ((box f))) (box ((box a)))))) Map.empty)))))))))))))))))) Map.empty))))))

let Data_Traversable_Accum_Internal_applyStateR  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((box Data_Traversable_Accum_Internal_StateR))) (box ((box (fun (s: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Traversable_Accum_Internal_stateR))) (box ((box x)))))) (box ((box s)))) in (match ((unbox ((box v)))) with | (HasProp "accum" (s1) & HasProp "value" (x_prime)) -> ((let v1 = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Traversable_Accum_Internal_stateR))) (box ((box f)))))) (box ((box s1)))) in (match ((unbox ((box v1)))) with | (HasProp "accum" (s2) & HasProp "value" (f_prime)) -> ((box ((Map.add "accum" (box ((box s2))) (Map.add "value" (box ((sharpurs_apply (box ((box f_prime))) (box ((box x_prime)))))) Map.empty))))))))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Traversable_Accum_Internal_functorStateR))))) Map.empty)))))))

let Data_Traversable_Accum_Internal_applyStateL  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((box Data_Traversable_Accum_Internal_StateL))) (box ((box (fun (s: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Traversable_Accum_Internal_stateL))) (box ((box f)))))) (box ((box s)))) in (match ((unbox ((box v)))) with | (HasProp "accum" (s1) & HasProp "value" (f_prime)) -> ((let v1 = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Traversable_Accum_Internal_stateL))) (box ((box x)))))) (box ((box s1)))) in (match ((unbox ((box v1)))) with | (HasProp "accum" (s2) & HasProp "value" (x_prime)) -> ((box ((Map.add "accum" (box ((box s2))) (Map.add "value" (box ((sharpurs_apply (box ((box f_prime))) (box ((box x_prime)))))) Map.empty))))))))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Traversable_Accum_Internal_functorStateL))))) Map.empty)))))))

let Data_Traversable_Accum_Internal_applicativeStateR  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box (fun (a: obj) -> (sharpurs_apply (box ((box Data_Traversable_Accum_Internal_StateR))) (box ((box (fun (s: obj) -> (box ((Map.add "accum" (box ((box s))) (Map.add "value" (box ((box a))) Map.empty))))))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Traversable_Accum_Internal_applyStateR))))) Map.empty)))))))

let Data_Traversable_Accum_Internal_applicativeStateL  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box (fun (a: obj) -> (sharpurs_apply (box ((box Data_Traversable_Accum_Internal_StateL))) (box ((box (fun (s: obj) -> (box ((Map.add "accum" (box ((box s))) (Map.add "value" (box ((box a))) Map.empty))))))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Traversable_Accum_Internal_applyStateL))))) Map.empty)))))))
