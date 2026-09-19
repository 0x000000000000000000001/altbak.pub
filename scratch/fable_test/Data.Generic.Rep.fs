[<AutoOpen>]
module PureScript_Data_Generic_Rep

open System
open System.Collections.Generic

type Data_Generic_Rep_Sum =
  | Data_Generic_Rep_Inlusd_Ctor of obj
  | Data_Generic_Rep_Inrusd_Ctor of obj

type Data_Generic_Rep_Product =
  | Data_Generic_Rep_Productusd_Ctor of obj * obj

type Data_Generic_Rep_NoArguments =
  | Data_Generic_Rep_NoArgumentsusd_Ctor

let Data_Generic_Rep_Inl  = (box ((fun (usd__arg1: obj) -> (box (Data_Generic_Rep_Inlusd_Ctor(usd__arg1))))))

let Data_Generic_Rep_Inr  = (box ((fun (usd__arg1: obj) -> (box (Data_Generic_Rep_Inrusd_Ctor(usd__arg1))))))

let Data_Generic_Rep_Product  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Generic_Rep_Productusd_Ctor(usd__arg1, usd__arg2)))))))

let Data_Generic_Rep_NoConstructors  = (box (fun (x: obj) -> (box x)))

let Data_Generic_Rep_NoArguments  = (box Data_Generic_Rep_NoArgumentsusd_Ctor)

let Data_Generic_Rep_Genericusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Generic_Rep_Constructor  = (box (fun (x: obj) -> (box x)))

let Data_Generic_Rep_Argument  = (box (fun (x: obj) -> (box x)))

let Data_Generic_Rep_to  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "to" (unbox<Map<string, obj>> ((box v))))))))

let Data_Generic_Rep_showSum  = (box (fun (dictShow: obj) -> (box (fun (dictShow1: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Generic_Rep_Inlusd_Ctor(a) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Inl ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((box ")")))))))) | Data_Generic_Rep_Inrusd_Ctor(b) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Inr ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow1)))))) (box ((box b))))))))) (box ((box ")"))))))))))))) Map.empty))))))))))

let Data_Generic_Rep_showProduct  = (box (fun (dictShow: obj) -> (box (fun (dictShow1: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Generic_Rep_Productusd_Ctor(a, b) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Product ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box " ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow1)))))) (box ((box b))))))))) (box ((box ")"))))))))))))))))))) Map.empty))))))))))

let Data_Generic_Rep_showNoArguments  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (box "NoArguments"))))) Map.empty))))))

let Data_Generic_Rep_showConstructor  = (box (fun (dictIsSymbol: obj) -> (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Constructor @")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box " ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((box ")"))))))))))))))))))) Map.empty))))))))))

let Data_Generic_Rep_showArgument  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Argument ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Generic_Rep_repOf  = (box (fun (dictGeneric: obj) -> (box (fun (v: obj) -> (box Type_Proxy_Proxyusd_Ctor)))))

let Data_Generic_Rep_from  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "from" (unbox<Map<string, obj>> ((box v))))))))
