[<AutoOpen>]
module PureScript_Data_Eq_Generic

open System
open System.Collections.Generic

let Data_Eq_Generic_GenericEqusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Eq_Generic_genericEqNoConstructors  = (sharpurs_apply (box ((box Data_Eq_Generic_GenericEqusd_Dict))) (box ((box ((Map.add "genericEq'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box true))))))) Map.empty))))))

let Data_Eq_Generic_genericEqNoArguments  = (sharpurs_apply (box ((box Data_Eq_Generic_GenericEqusd_Dict))) (box ((box ((Map.add "genericEq'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box true))))))) Map.empty))))))

let Data_Eq_Generic_genericEqArgument  = (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_Generic_GenericEqusd_Dict))) (box ((box ((Map.add "genericEq'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a1, a2) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box dictEq)))))) (box ((box a1)))))) (box ((box a2)))))))))))) Map.empty))))))))

let Data_Eq_Generic_genericEq_prime  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "genericEq'" (unbox<Map<string, obj>> ((box v))))))))

let Data_Eq_Generic_genericEqConstructor  = (box (fun (dictGenericEq: obj) -> (sharpurs_apply (box ((box Data_Eq_Generic_GenericEqusd_Dict))) (box ((box ((Map.add "genericEq'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a1, a2) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_Generic_genericEq_prime))) (box ((box dictGenericEq)))))) (box ((box a1)))))) (box ((box a2)))))))))))) Map.empty))))))))

let Data_Eq_Generic_genericEqProduct  = (box (fun (dictGenericEq: obj) -> (box (fun (dictGenericEq1: obj) -> (sharpurs_apply (box ((box Data_Eq_Generic_GenericEqusd_Dict))) (box ((box ((Map.add "genericEq'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (Data_Generic_Rep_Productusd_Ctor(a1, b1), Data_Generic_Rep_Productusd_Ctor(a2, b2)) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_conj))) (box ((box Data_HeytingAlgebra_heytingAlgebraBoolean)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_Generic_genericEq_prime))) (box ((box dictGenericEq)))))) (box ((box a1)))))) (box ((box a2))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_Generic_genericEq_prime))) (box ((box dictGenericEq1)))))) (box ((box b1)))))) (box ((box b2))))))))))))))) Map.empty))))))))))

let Data_Eq_Generic_genericEqSum  = (box (fun (dictGenericEq: obj) -> (box (fun (dictGenericEq1: obj) -> (sharpurs_apply (box ((box Data_Eq_Generic_GenericEqusd_Dict))) (box ((box ((Map.add "genericEq'" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (Data_Generic_Rep_Inlusd_Ctor(a1), Data_Generic_Rep_Inlusd_Ctor(a2)) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_Generic_genericEq_prime))) (box ((box dictGenericEq)))))) (box ((box a1)))))) (box ((box a2))))) | (Data_Generic_Rep_Inrusd_Ctor(b1), Data_Generic_Rep_Inrusd_Ctor(b2)) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_Generic_genericEq_prime))) (box ((box dictGenericEq1)))))) (box ((box b1)))))) (box ((box b2))))) | (_, _) -> ((box false))))))))) Map.empty))))))))))

let Data_Eq_Generic_genericEq  = (box (fun (dictGeneric: obj) -> (box (fun (dictGenericEq: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_Generic_genericEq_prime))) (box ((box dictGenericEq)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_from))) (box ((box dictGeneric)))))) (box ((box x))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_from))) (box ((box dictGeneric)))))) (box ((box y)))))))))))))))
