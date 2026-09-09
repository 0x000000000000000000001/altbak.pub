[<AutoOpen>]
module PureScript_Data_Ring_Generic

open System
open System.Collections.Generic

let Data_Ring_Generic_GenericRingusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Ring_Generic_genericSub_prime  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "genericSub'" (unbox<Map<string, obj>> ((box v))))))))

let Data_Ring_Generic_genericSub  = (box (fun (dictGeneric: obj) -> (let to_var = (sharpurs_apply (box ((box Data_Generic_Rep_to))) (box ((box dictGeneric)))) in (box (fun (dictGenericRing: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box to_var)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_Generic_genericSub_prime))) (box ((box dictGenericRing)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_from))) (box ((box dictGeneric)))))) (box ((box x))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_from))) (box ((box dictGeneric)))))) (box ((box y)))))))))))))))))))

let Data_Ring_Generic_genericRingProduct  = (box (fun (dictGenericRing: obj) -> (box (fun (dictGenericRing1: obj) -> (sharpurs_apply (box ((box Data_Ring_Generic_GenericRingusd_Dict))) (box ((box ((Map.add "genericSub'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (Data_Generic_Rep_Productusd_Ctor(a1, b1), Data_Generic_Rep_Productusd_Ctor(a2, b2)) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Generic_Rep_Productusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_Generic_genericSub_prime))) (box ((box dictGenericRing)))))) (box ((box a1)))))) (box ((box a2))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_Generic_genericSub_prime))) (box ((box dictGenericRing1)))))) (box ((box b1)))))) (box ((box b2))))))))))))))) Map.empty))))))))))

let Data_Ring_Generic_genericRingNoArguments  = (sharpurs_apply (box ((box Data_Ring_Generic_GenericRingusd_Dict))) (box ((box ((Map.add "genericSub'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Generic_Rep_NoArgumentsusd_Ctor))))))) Map.empty))))))

let Data_Ring_Generic_genericRingConstructor  = (box (fun (dictGenericRing: obj) -> (sharpurs_apply (box ((box Data_Ring_Generic_GenericRingusd_Dict))) (box ((box ((Map.add "genericSub'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a1, a2) -> ((sharpurs_apply (box ((box Data_Generic_Rep_Constructor))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_Generic_genericSub_prime))) (box ((box dictGenericRing)))))) (box ((box a1)))))) (box ((box a2))))))))))))))) Map.empty))))))))

let Data_Ring_Generic_genericRingArgument  = (box (fun (dictRing: obj) -> (sharpurs_apply (box ((box Data_Ring_Generic_GenericRingusd_Dict))) (box ((box ((Map.add "genericSub'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((box Data_Generic_Rep_Argument))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box dictRing)))))) (box ((box x)))))) (box ((box y))))))))))))))) Map.empty))))))))
