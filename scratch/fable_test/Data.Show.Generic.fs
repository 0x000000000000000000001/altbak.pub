[<AutoOpen>]
module PureScript_Data_Show_Generic

open System
open System.Collections.Generic

module Data_Show_Generic_FFI =
    let intercalate = undefined
    

let Data_Show_Generic_intercalate = box (Data_Show_Generic_FFI.``intercalate``)


let Data_Show_Generic_GenericShowArgsusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Show_Generic_GenericShowusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Show_Generic_genericShowArgsNoArguments  = (sharpurs_apply (box ((box Data_Show_Generic_GenericShowArgsusd_Dict))) (box ((box ((Map.add "genericShowArgs" (box ((box (fun (v: obj) -> (box [||]))))) Map.empty))))))

let Data_Show_Generic_genericShowArgsArgument  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Generic_GenericShowArgsusd_Dict))) (box ((box ((Map.add "genericShowArgs" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((box [|(sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))|]))))))) Map.empty))))))))

let Data_Show_Generic_genericShowArgs  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "genericShowArgs" (unbox<Map<string, obj>> ((box v))))))))

let Data_Show_Generic_genericShowArgsProduct  = (box (fun (dictGenericShowArgs: obj) -> (box (fun (dictGenericShowArgs1: obj) -> (sharpurs_apply (box ((box Data_Show_Generic_GenericShowArgsusd_Dict))) (box ((box ((Map.add "genericShowArgs" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Generic_Rep_Productusd_Ctor(a, b) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupArray)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_Generic_genericShowArgs))) (box ((box dictGenericShowArgs)))))) (box ((box a))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_Generic_genericShowArgs))) (box ((box dictGenericShowArgs1)))))) (box ((box b))))))))))))) Map.empty))))))))))

let Data_Show_Generic_genericShowConstructor  = (box (fun (dictGenericShowArgs: obj) -> (box (fun (dictIsSymbol: obj) -> (sharpurs_apply (box ((box Data_Show_Generic_GenericShowusd_Dict))) (box ((box ((Map.add "genericShow'" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((let ctor = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in (let v1 = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_Generic_genericShowArgs))) (box ((box dictGenericShowArgs)))))) (box ((box a)))) in (match ((unbox ((box v1)))) with | [|  |] -> ((box ctor)) | args -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_Generic_intercalate))) (box ((box " ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupArray)))))) (box ((box [|(box ctor)|])))))) (box ((box args)))))))))))) (box ((box ")"))))))))))))))))) Map.empty))))))))))

let Data_Show_Generic_genericShow_prime  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "genericShow'" (unbox<Map<string, obj>> ((box v))))))))

let rec Data_Show_Generic_genericShowNoConstructors : obj = ((sharpurs_apply (box ((box Data_Show_Generic_GenericShowusd_Dict))) (box ((box ((Map.add "genericShow'" (box ((box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_Generic_genericShow_prime))) (box ((box Data_Show_Generic_genericShowNoConstructors)))))) (box ((box a)))))))) Map.empty)))))))


let Data_Show_Generic_genericShowSum  = (box (fun (dictGenericShow: obj) -> (box (fun (dictGenericShow1: obj) -> (sharpurs_apply (box ((box Data_Show_Generic_GenericShowusd_Dict))) (box ((box ((Map.add "genericShow'" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Generic_Rep_Inlusd_Ctor(a) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_Generic_genericShow_prime))) (box ((box dictGenericShow)))))) (box ((box a))))) | Data_Generic_Rep_Inrusd_Ctor(b) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_Generic_genericShow_prime))) (box ((box dictGenericShow1)))))) (box ((box b)))))))))) Map.empty))))))))))

let Data_Show_Generic_genericShow  = (box (fun (dictGeneric: obj) -> (box (fun (dictGenericShow: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_Generic_genericShow_prime))) (box ((box dictGenericShow)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Generic_Rep_from))) (box ((box dictGeneric)))))) (box ((box x)))))))))))))
