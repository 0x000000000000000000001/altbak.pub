[<AutoOpen>]
module PureScript_Data_Show

open System
open System.Collections.Generic

module Data_Show_FFI =
    let showIntImpl (x: obj) : obj = box (string (unbox<int> x))
    let showNumberImpl (x: obj) : obj = box (string (unbox<float> x))
    let showStringImpl (x: obj) : obj = box ("\"" + unbox<string> x + "\"")
    let showCharImpl c = string (unbox<char> c)
    let showArrayImpl (f: obj) (xs: obj) =
        let xs' = unbox<obj[]> xs
        let mutable res = "["
        for i = 0 to xs'.Length - 1 do
            res <- res + unbox<string> (Sharpurs_Prelude.sharpurs_apply f xs'.[i])
            if i < xs'.Length - 1 then
                res <- res + ","
        res <- res + "]"
        box res
    

let Data_Show_showArrayImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Show_FFI.``showArrayImpl`` (unbox arg0) (unbox arg1))))
let Data_Show_showCharImpl = box (fun (arg0: obj) -> box (Data_Show_FFI.``showCharImpl`` (unbox arg0)))
let Data_Show_showIntImpl = box (fun (arg0: obj) -> box (Data_Show_FFI.``showIntImpl`` (unbox arg0)))
let Data_Show_showNumberImpl = box (fun (arg0: obj) -> box (Data_Show_FFI.``showNumberImpl`` (unbox arg0)))
let Data_Show_showStringImpl = box (fun (arg0: obj) -> box (Data_Show_FFI.``showStringImpl`` (unbox arg0)))


let Data_Show_ShowRecordFieldsusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Show_Showusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Show_showVoid  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box Data_Void_absurd))) Map.empty))))))

let Data_Show_showUnit  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (box "unit"))))) Map.empty))))))

let Data_Show_showString  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box Data_Show_showStringImpl))) Map.empty))))))

let Data_Show_showRecordFieldsNil  = (sharpurs_apply (box ((box Data_Show_ShowRecordFieldsusd_Dict))) (box ((box ((Map.add "showRecordFields" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box ""))))))) Map.empty))))))

let Data_Show_showRecordFields  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "showRecordFields" (unbox<Map<string, obj>> ((box v))))))))

let Data_Show_showRecord  = (box (fun (usd__unused: obj) -> (box (fun (usd__unused: obj) -> (box (fun (dictShowRecordFields: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (record: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "{")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_showRecordFields))) (box ((box dictShowRecordFields)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box record))))))))) (box ((box "}"))))))))))) Map.empty))))))))))))

let Data_Show_showProxy  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (box "Proxy"))))) Map.empty))))))

let Data_Show_showNumber  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box Data_Show_showNumberImpl))) Map.empty))))))

let Data_Show_showInt  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box Data_Show_showIntImpl))) Map.empty))))))

let Data_Show_showChar  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box Data_Show_showCharImpl))) Map.empty))))))

let Data_Show_showBoolean  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | LitBool true () -> ((box "true")) | LitBool false () -> ((box "false"))))))) Map.empty))))))

let Data_Show_show  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "show" (unbox<Map<string, obj>> ((box v))))))))

let Data_Show_showArray  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((sharpurs_apply (box ((box Data_Show_showArrayImpl))) (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow))))))))) Map.empty))))))))

let Data_Show_showRecordFieldsCons  = (box (fun (dictIsSymbol: obj) -> (box (fun (dictShowRecordFields: obj) -> (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_ShowRecordFieldsusd_Dict))) (box ((box ((Map.add "showRecordFields" (box ((box (fun (v: obj) -> (box (fun (record: obj) -> (let tail = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_showRecordFields))) (box ((box dictShowRecordFields)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box record)))) in let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let focus = (sharpurs_apply (box ((sharpurs_apply (box ((box Record_Unsafe_unsafeGet))) (box ((box key)))))) (box ((box record)))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box " ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box key)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box ": ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box focus))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box ",")))))) (box ((box tail))))))))))))))))))))))) Map.empty))))))))))))

let Data_Show_showRecordFieldsConsNil  = (box (fun (dictIsSymbol: obj) -> (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_ShowRecordFieldsusd_Dict))) (box ((box ((Map.add "showRecordFields" (box ((box (fun (v: obj) -> (box (fun (record: obj) -> (let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let focus = (sharpurs_apply (box ((sharpurs_apply (box ((box Record_Unsafe_unsafeGet))) (box ((box key)))))) (box ((box record)))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box " ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box key)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box ": ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box focus))))))))) (box ((box " ")))))))))))))))))))) Map.empty))))))))))
