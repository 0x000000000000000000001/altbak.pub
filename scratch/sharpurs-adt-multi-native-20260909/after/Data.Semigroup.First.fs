[<AutoOpen>]
module PureScript_Data_Semigroup_First

open System
open System.Collections.Generic

let Data_Semigroup_First_First  = (box (fun (x: obj) -> (box x)))

let Data_Semigroup_First_showFirst  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(First ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Semigroup_First_semigroupFirst  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (x: obj) -> (box (fun (v: obj) -> (box x))))))) Map.empty))))))

let Data_Semigroup_First_ordFirst  = (box (fun (dictOrd: obj) -> (box dictOrd)))

let Data_Semigroup_First_functorFirst  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (f: obj) -> (box (fun (m: obj) -> (match ((unbox ((box m)))) with | v -> ((sharpurs_apply (box ((box Data_Semigroup_First_First))) (box ((sharpurs_apply (box ((box f))) (box ((box v))))))))))))))) Map.empty))))))

let Data_Semigroup_First_eqFirst  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Semigroup_First_eq1First  = (sharpurs_apply (box ((box Data_Eq_Eq1usd_Dict))) (box ((box ((Map.add "eq1" (box ((box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_eq))) (box ((sharpurs_apply (box ((box Data_Semigroup_First_eqFirst))) (box ((box dictEq))))))))))) Map.empty))))))

let Data_Semigroup_First_ord1First  = (sharpurs_apply (box ((box Data_Ord_Ord1usd_Dict))) (box ((box ((Map.add "compare1" (box ((box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_compare))) (box ((sharpurs_apply (box ((box Data_Semigroup_First_ordFirst))) (box ((box dictOrd))))))))))) (Map.add "Eq10" (box ((box (fun (usd__unused: obj) -> (box Data_Semigroup_First_eq1First))))) Map.empty)))))))

let Data_Semigroup_First_boundedFirst  = (box (fun (dictBounded: obj) -> (box dictBounded)))

let Data_Semigroup_First_applyFirst  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (f, x) -> ((sharpurs_apply (box ((box Data_Semigroup_First_First))) (box ((sharpurs_apply (box ((box f))) (box ((box x))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Semigroup_First_functorFirst))))) Map.empty)))))))

let Data_Semigroup_First_bindFirst  = (sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((box (fun (v: obj) -> (box (fun (f: obj) -> (match (((unbox ((box v))), (unbox ((box f))))) with | (x, f1) -> ((sharpurs_apply (box ((box f1))) (box ((box x)))))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Semigroup_First_applyFirst))))) Map.empty)))))))

let Data_Semigroup_First_applicativeFirst  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box Data_Semigroup_First_First))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Semigroup_First_applyFirst))))) Map.empty)))))))

let Data_Semigroup_First_monadFirst  = (sharpurs_apply (box ((box Control_Monad_Monadusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Data_Semigroup_First_applicativeFirst))))) (Map.add "Bind1" (box ((box (fun (usd__unused: obj) -> (box Data_Semigroup_First_bindFirst))))) Map.empty)))))))
