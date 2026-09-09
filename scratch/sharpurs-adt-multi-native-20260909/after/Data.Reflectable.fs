[<AutoOpen>]
module PureScript_Data_Reflectable

open System
open System.Collections.Generic

module Data_Reflectable_FFI =
    let unsafeCoerce x = x
    

let Data_Reflectable_unsafeCoerce = box (fun (arg0: obj) -> box (Data_Reflectable_FFI.``unsafeCoerce`` (unbox arg0)))


let Data_Reflectable_Reifiableusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Reflectable_Reflectableusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Reflectable_reifiableString  = (sharpurs_apply (box ((box Data_Reflectable_Reifiableusd_Dict))) (box ((box (Map.empty)))))

let Data_Reflectable_reifiableOrdering  = (sharpurs_apply (box ((box Data_Reflectable_Reifiableusd_Dict))) (box ((box (Map.empty)))))

let Data_Reflectable_reifiableInt  = (sharpurs_apply (box ((box Data_Reflectable_Reifiableusd_Dict))) (box ((box (Map.empty)))))

let Data_Reflectable_reifiableBoolean  = (sharpurs_apply (box ((box Data_Reflectable_Reifiableusd_Dict))) (box ((box (Map.empty)))))

let Data_Reflectable_reifyType  = (box (fun (usd__unused: obj) -> (box (fun (s: obj) -> (box (fun (f: obj) -> (let coerce = (box Data_Reflectable_unsafeCoerce) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box coerce))) (box ((box (fun (dictReflectable: obj) -> (sharpurs_apply (box ((box f))) (box ((box dictReflectable))))))))))) (box ((box ((Map.add "reflectType" (box ((box (fun (v: obj) -> (box s))))) Map.empty)))))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))))))))

let Data_Reflectable_reflectType  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "reflectType" (unbox<Map<string, obj>> ((box v))))))))
