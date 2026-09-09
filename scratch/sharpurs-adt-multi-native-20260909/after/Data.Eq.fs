[<AutoOpen>]
module PureScript_Data_Eq

open System
open System.Collections.Generic

module Data_Eq_FFI =
    let eqIntImpl (a: obj) (b: obj) = unbox<int> a = unbox<int> b
    let eqNumberImpl (a: obj) (b: obj) = unbox<float> a = unbox<float> b
    let eqStringImpl (a: obj) (b: obj) = unbox<string> a = unbox<string> b
    let eqCharImpl (a: obj) (b: obj) = unbox<char> a = unbox<char> b
    let eqBooleanImpl (a: obj) (b: obj) = unbox<bool> a = unbox<bool> b
    let eqArrayImpl (f: obj) (xs: obj) (ys: obj) =
        let xs' = unbox<obj[]> xs
        let ys' = unbox<obj[]> ys
        if xs'.Length <> ys'.Length then box false
        else
            let mutable eq = true
            for i = 0 to xs'.Length - 1 do
                let isEq = unbox<bool> (sharpurs_apply (sharpurs_apply f xs'.[i]) ys'.[i])
                if not isEq then
                    eq <- false
            box eq
    

let Data_Eq_eqArrayImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (fun (arg2: obj) -> box (Data_Eq_FFI.``eqArrayImpl`` (unbox arg0) (unbox arg1) (unbox arg2)))))
let Data_Eq_eqBooleanImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Eq_FFI.``eqBooleanImpl`` (unbox arg0) (unbox arg1))))
let Data_Eq_eqCharImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Eq_FFI.``eqCharImpl`` (unbox arg0) (unbox arg1))))
let Data_Eq_eqIntImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Eq_FFI.``eqIntImpl`` (unbox arg0) (unbox arg1))))
let Data_Eq_eqNumberImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Eq_FFI.``eqNumberImpl`` (unbox arg0) (unbox arg1))))
let Data_Eq_eqStringImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Eq_FFI.``eqStringImpl`` (unbox arg0) (unbox arg1))))


let Data_Eq_EqRecordusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Eq_Equsd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Eq_Eq1usd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Eq_eqVoid  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box true))))))) Map.empty))))))

let Data_Eq_eqUnit  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box true))))))) Map.empty))))))

let Data_Eq_eqString  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box Data_Eq_eqStringImpl))) Map.empty))))))

let Data_Eq_eqRowNil  = (sharpurs_apply (box ((box Data_Eq_EqRecordusd_Dict))) (box ((box ((Map.add "eqRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box true))))))))) Map.empty))))))

let Data_Eq_eqRecord  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "eqRecord" (unbox<Map<string, obj>> ((box v))))))))

let Data_Eq_eqRec  = (box (fun (usd__unused: obj) -> (box (fun (dictEqRecord: obj) -> (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eqRecord))) (box ((box dictEqRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) Map.empty))))))))))

let Data_Eq_eqProxy  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box true))))))) Map.empty))))))

let Data_Eq_eqNumber  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box Data_Eq_eqNumberImpl))) Map.empty))))))

let Data_Eq_eqInt  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box Data_Eq_eqIntImpl))) Map.empty))))))

let Data_Eq_eqChar  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box Data_Eq_eqCharImpl))) Map.empty))))))

let Data_Eq_eqBoolean  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box Data_Eq_eqBooleanImpl))) Map.empty))))))

let Data_Eq_eq1  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "eq1" (unbox<Map<string, obj>> ((box v))))))))

let Data_Eq_eq  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "eq" (unbox<Map<string, obj>> ((box v))))))))

let Data_Eq_eqArray  = (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((sharpurs_apply (box ((box Data_Eq_eqArrayImpl))) (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box dictEq))))))))) Map.empty))))))))

let Data_Eq_eq1Array  = (sharpurs_apply (box ((box Data_Eq_Eq1usd_Dict))) (box ((box ((Map.add "eq1" (box ((box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_eq))) (box ((sharpurs_apply (box ((box Data_Eq_eqArray))) (box ((box dictEq))))))))))) Map.empty))))))

let Data_Eq_eqRowCons  = (box (fun (dictEqRecord: obj) -> (box (fun (usd__unused: obj) -> (box (fun (dictIsSymbol: obj) -> (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_EqRecordusd_Dict))) (box ((box ((Map.add "eqRecord" (box ((box (fun (v: obj) -> (box (fun (ra: obj) -> (box (fun (rb: obj) -> (let tail = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eqRecord))) (box ((box dictEqRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box ra)))))) (box ((box rb)))) in let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let get = (sharpurs_apply (box ((box Record_Unsafe_unsafeGet))) (box ((box key)))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_conj))) (box ((box Data_HeytingAlgebra_heytingAlgebraBoolean)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box dictEq)))))) (box ((sharpurs_apply (box ((box get))) (box ((box ra))))))))) (box ((sharpurs_apply (box ((box get))) (box ((box rb)))))))))))) (box ((box tail))))))))))))) Map.empty))))))))))))))

let Data_Eq_notEq  = (box (fun (dictEq: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Eq_eqBoolean)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box dictEq)))))) (box ((box x)))))) (box ((box y))))))))) (box ((box false))))))))))

let Data_Eq_notEq1  = (box (fun (dictEq1: obj) -> (box (fun (dictEq: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Eq_eqBoolean)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq1))) (box ((box dictEq1)))))) (box ((box dictEq)))))) (box ((box x)))))) (box ((box y))))))))) (box ((box false))))))))))))
