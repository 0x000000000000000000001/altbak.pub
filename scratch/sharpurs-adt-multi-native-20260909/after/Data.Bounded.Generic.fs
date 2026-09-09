[<AutoOpen>]
module PureScript_Data_Bounded_Generic

open System
open System.Collections.Generic

let Data_Bounded_Generic_GenericTopusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Bounded_Generic_GenericBottomusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Bounded_Generic_genericTopNoArguments  = (sharpurs_apply (box ((box Data_Bounded_Generic_GenericTopusd_Dict))) (box ((box ((Map.add "genericTop'" (box ((box Data_Generic_Rep_NoArgumentsusd_Ctor))) Map.empty))))))

let Data_Bounded_Generic_genericTopArgument  = (box (fun (dictBounded: obj) -> (sharpurs_apply (box ((box Data_Bounded_Generic_GenericTopusd_Dict))) (box ((box ((Map.add "genericTop'" (box ((sharpurs_apply (box ((box Data_Generic_Rep_Argument))) (box ((sharpurs_apply (box ((box Data_Bounded_top))) (box ((box dictBounded))))))))) Map.empty))))))))

let Data_Bounded_Generic_genericTop_prime  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "genericTop'" (unbox<Map<string, obj>> ((box v))))))))

let Data_Bounded_Generic_genericTopConstructor  = (box (fun (dictGenericTop: obj) -> (sharpurs_apply (box ((box Data_Bounded_Generic_GenericTopusd_Dict))) (box ((box ((Map.add "genericTop'" (box ((sharpurs_apply (box ((box Data_Generic_Rep_Constructor))) (box ((sharpurs_apply (box ((box Data_Bounded_Generic_genericTop_prime))) (box ((box dictGenericTop))))))))) Map.empty))))))))

let Data_Bounded_Generic_genericTopProduct  = (box (fun (dictGenericTop: obj) -> (box (fun (dictGenericTop1: obj) -> (sharpurs_apply (box ((box Data_Bounded_Generic_GenericTopusd_Dict))) (box ((box ((Map.add "genericTop'" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Generic_Rep_Productusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((box Data_Bounded_Generic_genericTop_prime))) (box ((box dictGenericTop))))))))) (box ((sharpurs_apply (box ((box Data_Bounded_Generic_genericTop_prime))) (box ((box dictGenericTop1))))))))) Map.empty))))))))))

let Data_Bounded_Generic_genericTopSum  = (box (fun (dictGenericTop: obj) -> (sharpurs_apply (box ((box Data_Bounded_Generic_GenericTopusd_Dict))) (box ((box ((Map.add "genericTop'" (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Generic_Rep_Inrusd_Ctor(usd__arg1)))))))) (box ((sharpurs_apply (box ((box Data_Bounded_Generic_genericTop_prime))) (box ((box dictGenericTop))))))))) Map.empty))))))))

let Data_Bounded_Generic_genericTop  = (box (fun (dictGeneric: obj) -> (box (fun (dictGenericTop: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_to))) (box ((box dictGeneric)))))) (box ((sharpurs_apply (box ((box Data_Bounded_Generic_genericTop_prime))) (box ((box dictGenericTop)))))))))))

let Data_Bounded_Generic_genericBottomNoArguments  = (sharpurs_apply (box ((box Data_Bounded_Generic_GenericBottomusd_Dict))) (box ((box ((Map.add "genericBottom'" (box ((box Data_Generic_Rep_NoArgumentsusd_Ctor))) Map.empty))))))

let Data_Bounded_Generic_genericBottomArgument  = (box (fun (dictBounded: obj) -> (sharpurs_apply (box ((box Data_Bounded_Generic_GenericBottomusd_Dict))) (box ((box ((Map.add "genericBottom'" (box ((sharpurs_apply (box ((box Data_Generic_Rep_Argument))) (box ((sharpurs_apply (box ((box Data_Bounded_bottom))) (box ((box dictBounded))))))))) Map.empty))))))))

let Data_Bounded_Generic_genericBottom_prime  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "genericBottom'" (unbox<Map<string, obj>> ((box v))))))))

let Data_Bounded_Generic_genericBottomConstructor  = (box (fun (dictGenericBottom: obj) -> (sharpurs_apply (box ((box Data_Bounded_Generic_GenericBottomusd_Dict))) (box ((box ((Map.add "genericBottom'" (box ((sharpurs_apply (box ((box Data_Generic_Rep_Constructor))) (box ((sharpurs_apply (box ((box Data_Bounded_Generic_genericBottom_prime))) (box ((box dictGenericBottom))))))))) Map.empty))))))))

let Data_Bounded_Generic_genericBottomProduct  = (box (fun (dictGenericBottom: obj) -> (box (fun (dictGenericBottom1: obj) -> (sharpurs_apply (box ((box Data_Bounded_Generic_GenericBottomusd_Dict))) (box ((box ((Map.add "genericBottom'" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Generic_Rep_Productusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((box Data_Bounded_Generic_genericBottom_prime))) (box ((box dictGenericBottom))))))))) (box ((sharpurs_apply (box ((box Data_Bounded_Generic_genericBottom_prime))) (box ((box dictGenericBottom1))))))))) Map.empty))))))))))

let Data_Bounded_Generic_genericBottomSum  = (box (fun (dictGenericBottom: obj) -> (sharpurs_apply (box ((box Data_Bounded_Generic_GenericBottomusd_Dict))) (box ((box ((Map.add "genericBottom'" (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Generic_Rep_Inlusd_Ctor(usd__arg1)))))))) (box ((sharpurs_apply (box ((box Data_Bounded_Generic_genericBottom_prime))) (box ((box dictGenericBottom))))))))) Map.empty))))))))

let Data_Bounded_Generic_genericBottom  = (box (fun (dictGeneric: obj) -> (box (fun (dictGenericBottom: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_to))) (box ((box dictGeneric)))))) (box ((sharpurs_apply (box ((box Data_Bounded_Generic_genericBottom_prime))) (box ((box dictGenericBottom)))))))))))
