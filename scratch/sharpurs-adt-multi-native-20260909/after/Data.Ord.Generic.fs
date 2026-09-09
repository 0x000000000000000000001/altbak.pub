[<AutoOpen>]
module PureScript_Data_Ord_Generic

open System
open System.Collections.Generic

let Data_Ord_Generic_GenericOrdusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Ord_Generic_genericOrdNoConstructors  = (sharpurs_apply (box ((box Data_Ord_Generic_GenericOrdusd_Dict))) (box ((box ((Map.add "genericCompare'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Ordering_EQusd_Ctor))))))) Map.empty))))))

let Data_Ord_Generic_genericOrdNoArguments  = (sharpurs_apply (box ((box Data_Ord_Generic_GenericOrdusd_Dict))) (box ((box ((Map.add "genericCompare'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Ordering_EQusd_Ctor))))))) Map.empty))))))

let Data_Ord_Generic_genericOrdArgument  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_Generic_GenericOrdusd_Dict))) (box ((box ((Map.add "genericCompare'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a1, a2) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((box a1)))))) (box ((box a2)))))))))))) Map.empty))))))))

let Data_Ord_Generic_genericCompare_prime  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "genericCompare'" (unbox<Map<string, obj>> ((box v))))))))

let Data_Ord_Generic_genericOrdConstructor  = (box (fun (dictGenericOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_Generic_GenericOrdusd_Dict))) (box ((box ((Map.add "genericCompare'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a1, a2) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_Generic_genericCompare_prime))) (box ((box dictGenericOrd)))))) (box ((box a1)))))) (box ((box a2)))))))))))) Map.empty))))))))

let Data_Ord_Generic_genericOrdProduct  = (box (fun (dictGenericOrd: obj) -> (box (fun (dictGenericOrd1: obj) -> (sharpurs_apply (box ((box Data_Ord_Generic_GenericOrdusd_Dict))) (box ((box ((Map.add "genericCompare'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (Data_Generic_Rep_Productusd_Ctor(a1, b1), Data_Generic_Rep_Productusd_Ctor(a2, b2)) -> ((let v2 = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_Generic_genericCompare_prime))) (box ((box dictGenericOrd)))))) (box ((box a1)))))) (box ((box a2)))) in (match ((unbox ((box v2)))) with | Data_Ordering_EQusd_Ctor -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_Generic_genericCompare_prime))) (box ((box dictGenericOrd1)))))) (box ((box b1)))))) (box ((box b2))))) | other -> ((box other)))))))))))) Map.empty))))))))))

let Data_Ord_Generic_genericOrdSum  = (box (fun (dictGenericOrd: obj) -> (box (fun (dictGenericOrd1: obj) -> (sharpurs_apply (box ((box Data_Ord_Generic_GenericOrdusd_Dict))) (box ((box ((Map.add "genericCompare'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (Data_Generic_Rep_Inlusd_Ctor(a1), Data_Generic_Rep_Inlusd_Ctor(a2)) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_Generic_genericCompare_prime))) (box ((box dictGenericOrd)))))) (box ((box a1)))))) (box ((box a2))))) | (Data_Generic_Rep_Inrusd_Ctor(b1), Data_Generic_Rep_Inrusd_Ctor(b2)) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_Generic_genericCompare_prime))) (box ((box dictGenericOrd1)))))) (box ((box b1)))))) (box ((box b2))))) | (Data_Generic_Rep_Inlusd_Ctor(_), Data_Generic_Rep_Inrusd_Ctor(_)) -> ((box Data_Ordering_LTusd_Ctor)) | (Data_Generic_Rep_Inrusd_Ctor(_), Data_Generic_Rep_Inlusd_Ctor(_)) -> ((box Data_Ordering_GTusd_Ctor))))))))) Map.empty))))))))))

let Data_Ord_Generic_genericCompare  = (box (fun (dictGeneric: obj) -> (box (fun (dictGenericOrd: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_Generic_genericCompare_prime))) (box ((box dictGenericOrd)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_from))) (box ((box dictGeneric)))))) (box ((box x))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_from))) (box ((box dictGeneric)))))) (box ((box y)))))))))))))))
