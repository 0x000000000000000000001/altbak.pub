[<AutoOpen>]
module PureScript_Data_Symbol

open System
open System.Collections.Generic

module Data_Symbol_FFI =
    let unsafeCoerce x = x
    

let Data_Symbol_unsafeCoerce = box (fun (arg0: obj) -> box (Data_Symbol_FFI.``unsafeCoerce`` (unbox arg0)))


let Data_Symbol_IsSymbolusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Symbol_reifySymbol  = (box (fun (s: obj) -> (box (fun (f: obj) -> (let coerce = (box Data_Symbol_unsafeCoerce) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box coerce))) (box ((box (fun (dictIsSymbol: obj) -> (sharpurs_apply (box ((box f))) (box ((box dictIsSymbol))))))))))) (box ((box ((Map.add "reflectSymbol" (box ((box (fun (v: obj) -> (box s))))) Map.empty)))))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))))))

let Data_Symbol_reflectSymbol  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "reflectSymbol" (unbox<Map<string, obj>> ((box v))))))))
