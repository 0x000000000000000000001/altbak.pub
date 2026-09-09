[<AutoOpen>]
module PureScript_Data_Profunctor_Split

open System
open System.Collections.Generic

type Data_Profunctor_Split_SplitF =
  | Data_Profunctor_Split_SplitFusd_Ctor of obj * obj * obj

let Data_Profunctor_Split_identity  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Data_Profunctor_Split_SplitF  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (fun (usd__arg3: obj) -> (box (Data_Profunctor_Split_SplitFusd_Ctor(usd__arg1, usd__arg2, usd__arg3))))))))

let Data_Profunctor_Split_Split  = (box (fun (x: obj) -> (box x)))

let Data_Profunctor_Split_unSplit  = (box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, e) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Exists_runExists))) (box ((box (fun (v1: obj) -> (match ((unbox ((box v1)))) with | Data_Profunctor_Split_SplitFusd_Ctor(g, h, fx) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box f1))) (box ((box g)))))) (box ((box h)))))) (box ((box fx))))))))))))) (box ((box e))))))))))

let Data_Profunctor_Split_split  = (box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (fx: obj) -> (sharpurs_apply (box ((box Data_Profunctor_Split_Split))) (box ((sharpurs_apply (box ((box Data_Exists_mkExists))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (fun (usd__arg3: obj) -> (box (Data_Profunctor_Split_SplitFusd_Ctor(usd__arg1, usd__arg2, usd__arg3)))))))))) (box ((box f)))))) (box ((box g)))))) (box ((box fx))))))))))))))))

let Data_Profunctor_Split_profunctorSplit  = (sharpurs_apply (box ((box Data_Profunctor_Profunctorusd_Dict))) (box ((box ((Map.add "dimap" (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (sharpurs_apply (box ((box Data_Profunctor_Split_unSplit))) (box ((box (fun (h: obj) -> (box (fun (i: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_Split_split))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box h)))))) (box ((box f))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box g)))))) (box ((box i)))))))))))))))))))) Map.empty))))))

let Data_Profunctor_Split_lowerSplit  = (box (fun (dictInvariant: obj) -> (sharpurs_apply (box ((box Data_Profunctor_Split_unSplit))) (box ((sharpurs_apply (box ((box Data_Function_flip))) (box ((sharpurs_apply (box ((box Data_Functor_Invariant_imap))) (box ((box dictInvariant))))))))))))

let Data_Profunctor_Split_liftSplit  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_Split_split))) (box ((box Data_Profunctor_Split_identity)))))) (box ((box Data_Profunctor_Split_identity))))

let Data_Profunctor_Split_hoistSplit  = (box (fun (nat: obj) -> (sharpurs_apply (box ((box Data_Profunctor_Split_unSplit))) (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_Split_split))) (box ((box f)))))) (box ((box g))))))))) (box ((box nat)))))))))))))

let Data_Profunctor_Split_functorSplit  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((box Data_Profunctor_Split_unSplit))) (box ((box (fun (g: obj) -> (box (fun (h: obj) -> (box (fun (fx: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_Split_split))) (box ((box g)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f)))))) (box ((box h))))))))) (box ((box fx))))))))))))))))) Map.empty))))))
