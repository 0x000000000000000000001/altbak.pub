[<AutoOpen>]
module PureScript_Data_Semigroup_Generic

open System
open System.Collections.Generic

let Data_Semigroup_Generic_GenericSemigroupusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Semigroup_Generic_genericSemigroupNoConstructors  = (sharpurs_apply (box ((box Data_Semigroup_Generic_GenericSemigroupusd_Dict))) (box ((box ((Map.add "genericAppend'" (box ((box (fun (a: obj) -> (box (fun (v: obj) -> (box a))))))) Map.empty))))))

let Data_Semigroup_Generic_genericSemigroupNoArguments  = (sharpurs_apply (box ((box Data_Semigroup_Generic_GenericSemigroupusd_Dict))) (box ((box ((Map.add "genericAppend'" (box ((box (fun (a: obj) -> (box (fun (v: obj) -> (box a))))))) Map.empty))))))

let Data_Semigroup_Generic_genericSemigroupArgument  = (box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Generic_GenericSemigroupusd_Dict))) (box ((box ((Map.add "genericAppend'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a1, a2) -> ((sharpurs_apply (box ((box Data_Generic_Rep_Argument))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup)))))) (box ((box a1)))))) (box ((box a2))))))))))))))) Map.empty))))))))

let Data_Semigroup_Generic_genericAppend_prime  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "genericAppend'" (unbox<Map<string, obj>> ((box v))))))))

let Data_Semigroup_Generic_genericSemigroupConstructor  = (box (fun (dictGenericSemigroup: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Generic_GenericSemigroupusd_Dict))) (box ((box ((Map.add "genericAppend'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a1, a2) -> ((sharpurs_apply (box ((box Data_Generic_Rep_Constructor))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_Generic_genericAppend_prime))) (box ((box dictGenericSemigroup)))))) (box ((box a1)))))) (box ((box a2))))))))))))))) Map.empty))))))))

let Data_Semigroup_Generic_genericSemigroupProduct  = (box (fun (dictGenericSemigroup: obj) -> (box (fun (dictGenericSemigroup1: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Generic_GenericSemigroupusd_Dict))) (box ((box ((Map.add "genericAppend'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (Data_Generic_Rep_Productusd_Ctor(a1, b1), Data_Generic_Rep_Productusd_Ctor(a2, b2)) -> ((box (Data_Generic_Rep_Productusd_Ctor((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_Generic_genericAppend_prime))) (box ((box dictGenericSemigroup)))))) (box ((box a1)))))) (box ((box a2)))), (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_Generic_genericAppend_prime))) (box ((box dictGenericSemigroup1)))))) (box ((box b1)))))) (box ((box b2))))))))))))))) Map.empty))))))))))

let Data_Semigroup_Generic_genericAppend  = (box (fun (dictGeneric: obj) -> (box (fun (dictGenericSemigroup: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_to))) (box ((box dictGeneric)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_Generic_genericAppend_prime))) (box ((box dictGenericSemigroup)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_from))) (box ((box dictGeneric)))))) (box ((box x))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_from))) (box ((box dictGeneric)))))) (box ((box y))))))))))))))))))
