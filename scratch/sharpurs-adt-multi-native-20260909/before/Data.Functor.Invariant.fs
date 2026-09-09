[<AutoOpen>]
module PureScript_Data_Functor_Invariant

open System
open System.Collections.Generic

let Data_Functor_Invariant_Invariantusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Functor_Invariant_invariantMultiplicative  = (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box f))), (unbox ((box v))), (unbox ((box v1))))) with | (f1, _, x) -> ((sharpurs_apply (box ((box Data_Monoid_Multiplicative_Multiplicative))) (box ((sharpurs_apply (box ((box f1))) (box ((box x))))))))))))))))) Map.empty))))))

let Data_Functor_Invariant_invariantEndo  = (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((box (fun (ab: obj) -> (box (fun (ba: obj) -> (box (fun (v: obj) -> (match (((unbox ((box ab))), (unbox ((box ba))), (unbox ((box v))))) with | (ab1, ba1, f) -> ((sharpurs_apply (box ((box Data_Monoid_Endo_Endo))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box ab1)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f)))))) (box ((box ba1)))))))))))))))))))) Map.empty))))))

let Data_Functor_Invariant_invariantDual  = (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box f))), (unbox ((box v))), (unbox ((box v1))))) with | (f1, _, x) -> ((sharpurs_apply (box ((box Data_Monoid_Dual_Dual))) (box ((sharpurs_apply (box ((box f1))) (box ((box x))))))))))))))))) Map.empty))))))

let Data_Functor_Invariant_invariantDisj  = (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box f))), (unbox ((box v))), (unbox ((box v1))))) with | (f1, _, x) -> ((sharpurs_apply (box ((box Data_Monoid_Disj_Disj))) (box ((sharpurs_apply (box ((box f1))) (box ((box x))))))))))))))))) Map.empty))))))

let Data_Functor_Invariant_invariantConj  = (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box f))), (unbox ((box v))), (unbox ((box v1))))) with | (f1, _, x) -> ((sharpurs_apply (box ((box Data_Monoid_Conj_Conj))) (box ((sharpurs_apply (box ((box f1))) (box ((box x))))))))))))))))) Map.empty))))))

let Data_Functor_Invariant_invariantAdditive  = (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box f))), (unbox ((box v))), (unbox ((box v1))))) with | (f1, _, x) -> ((sharpurs_apply (box ((box Data_Monoid_Additive_Additive))) (box ((sharpurs_apply (box ((box f1))) (box ((box x))))))))))))))))) Map.empty))))))

let Data_Functor_Invariant_imapF  = (box (fun (dictFunctor: obj) -> (box (fun (f: obj) -> (box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((box f))))))))))

let Data_Functor_Invariant_invariantArray  = (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((sharpurs_apply (box ((box Data_Functor_Invariant_imapF))) (box ((box Data_Functor_functorArray)))))) Map.empty))))))

let Data_Functor_Invariant_invariantFn  = (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((sharpurs_apply (box ((box Data_Functor_Invariant_imapF))) (box ((box Data_Functor_functorFn)))))) Map.empty))))))

let Data_Functor_Invariant_imap  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "imap" (unbox<Map<string, obj>> ((box v))))))))

let Data_Functor_Invariant_invariantAlternate  = (box (fun (dictInvariant: obj) -> (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box g))), (unbox ((box v))))) with | (f1, g1, x) -> ((sharpurs_apply (box ((box Data_Monoid_Alternate_Alternate))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_Invariant_imap))) (box ((box dictInvariant)))))) (box ((box f1)))))) (box ((box g1)))))) (box ((box x))))))))))))))))) Map.empty))))))))
