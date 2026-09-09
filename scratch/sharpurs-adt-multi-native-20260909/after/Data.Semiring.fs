[<AutoOpen>]
module PureScript_Data_Semiring

open System
open System.Collections.Generic

module Data_Semiring_FFI =
    let intAdd a b = (unbox<int> a) + (unbox<int> b)
    let intMul a b = (unbox<int> a) * (unbox<int> b)
    let numAdd a b = (unbox<float> a) + (unbox<float> b)
    let numMul a b = (unbox<float> a) * (unbox<float> b)
    

let Data_Semiring_intAdd = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Semiring_FFI.``intAdd`` (unbox arg0) (unbox arg1))))
let Data_Semiring_intMul = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Semiring_FFI.``intMul`` (unbox arg0) (unbox arg1))))
let Data_Semiring_numAdd = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Semiring_FFI.``numAdd`` (unbox arg0) (unbox arg1))))
let Data_Semiring_numMul = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Semiring_FFI.``numMul`` (unbox arg0) (unbox arg1))))


let Data_Semiring_SemiringRecordusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Semiring_Semiringusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Semiring_zeroRecord  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "zeroRecord" (unbox<Map<string, obj>> ((box v))))))))

let Data_Semiring_zero  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "zero" (unbox<Map<string, obj>> ((box v))))))))

let Data_Semiring_semiringUnit  = (sharpurs_apply (box ((box Data_Semiring_Semiringusd_Dict))) (box ((box ((Map.add "add" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Unit_unit))))))) (Map.add "zero" (box ((box Data_Unit_unit))) (Map.add "mul" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Unit_unit))))))) (Map.add "one" (box ((box Data_Unit_unit))) Map.empty)))))))))

let Data_Semiring_semiringRecordNil  = (sharpurs_apply (box ((box Data_Semiring_SemiringRecordusd_Dict))) (box ((box ((Map.add "addRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (Map.empty)))))))))) (Map.add "mulRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (Map.empty)))))))))) (Map.add "oneRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (Map.empty)))))))) (Map.add "zeroRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (Map.empty)))))))) Map.empty)))))))))

let Data_Semiring_semiringProxy  = (sharpurs_apply (box ((box Data_Semiring_Semiringusd_Dict))) (box ((box ((Map.add "add" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Type_Proxy_Proxyusd_Ctor))))))) (Map.add "mul" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Type_Proxy_Proxyusd_Ctor))))))) (Map.add "one" (box ((box Type_Proxy_Proxyusd_Ctor))) (Map.add "zero" (box ((box Type_Proxy_Proxyusd_Ctor))) Map.empty)))))))))

let Data_Semiring_semiringNumber  = (sharpurs_apply (box ((box Data_Semiring_Semiringusd_Dict))) (box ((box ((Map.add "add" (box ((box Data_Semiring_numAdd))) (Map.add "zero" (box ((box 0.0))) (Map.add "mul" (box ((box Data_Semiring_numMul))) (Map.add "one" (box ((box 1.0))) Map.empty)))))))))

let Data_Semiring_semiringInt  = (sharpurs_apply (box ((box Data_Semiring_Semiringusd_Dict))) (box ((box ((Map.add "add" (box ((box Data_Semiring_intAdd))) (Map.add "zero" (box ((box 0))) (Map.add "mul" (box ((box Data_Semiring_intMul))) (Map.add "one" (box ((box 1))) Map.empty)))))))))

let Data_Semiring_oneRecord  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "oneRecord" (unbox<Map<string, obj>> ((box v))))))))

let Data_Semiring_one  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "one" (unbox<Map<string, obj>> ((box v))))))))

let Data_Semiring_mulRecord  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "mulRecord" (unbox<Map<string, obj>> ((box v))))))))

let Data_Semiring_mul  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "mul" (unbox<Map<string, obj>> ((box v))))))))

let Data_Semiring_addRecord  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "addRecord" (unbox<Map<string, obj>> ((box v))))))))

let Data_Semiring_semiringRecord  = (box (fun (usd__unused: obj) -> (box (fun (dictSemiringRecord: obj) -> (sharpurs_apply (box ((box Data_Semiring_Semiringusd_Dict))) (box ((box ((Map.add "add" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_addRecord))) (box ((box dictSemiringRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (Map.add "mul" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mulRecord))) (box ((box dictSemiringRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (Map.add "one" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_oneRecord))) (box ((box dictSemiringRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (Map.add "zero" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_zeroRecord))) (box ((box dictSemiringRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) Map.empty)))))))))))))

let Data_Semiring_add  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "add" (unbox<Map<string, obj>> ((box v))))))))

let Data_Semiring_semiringFn  = (box (fun (dictSemiring: obj) -> (sharpurs_apply (box ((box Data_Semiring_Semiringusd_Dict))) (box ((box ((Map.add "add" (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box dictSemiring)))))) (box ((sharpurs_apply (box ((box f))) (box ((box x))))))))) (box ((sharpurs_apply (box ((box g))) (box ((box x))))))))))))))) (Map.add "zero" (box ((box (fun (v: obj) -> (sharpurs_apply (box ((box Data_Semiring_zero))) (box ((box dictSemiring)))))))) (Map.add "mul" (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mul))) (box ((box dictSemiring)))))) (box ((sharpurs_apply (box ((box f))) (box ((box x))))))))) (box ((sharpurs_apply (box ((box g))) (box ((box x))))))))))))))) (Map.add "one" (box ((box (fun (v: obj) -> (sharpurs_apply (box ((box Data_Semiring_one))) (box ((box dictSemiring)))))))) Map.empty)))))))))))

let Data_Semiring_semiringRecordCons  = (box (fun (dictIsSymbol: obj) -> (box (fun (usd__unused: obj) -> (box (fun (dictSemiringRecord: obj) -> (box (fun (dictSemiring: obj) -> (let one1 = (sharpurs_apply (box ((box Data_Semiring_one))) (box ((box dictSemiring)))) in let zero1 = (sharpurs_apply (box ((box Data_Semiring_zero))) (box ((box dictSemiring)))) in (sharpurs_apply (box ((box Data_Semiring_SemiringRecordusd_Dict))) (box ((box ((Map.add "addRecord" (box ((box (fun (v: obj) -> (box (fun (ra: obj) -> (box (fun (rb: obj) -> (let tail = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_addRecord))) (box ((box dictSemiringRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box ra)))))) (box ((box rb)))) in let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let insert = (sharpurs_apply (box ((box Record_Unsafe_unsafeSet))) (box ((box key)))) in let get = (sharpurs_apply (box ((box Record_Unsafe_unsafeGet))) (box ((box key)))) in (sharpurs_apply (box ((sharpurs_apply (box ((box insert))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box dictSemiring)))))) (box ((sharpurs_apply (box ((box get))) (box ((box ra))))))))) (box ((sharpurs_apply (box ((box get))) (box ((box rb)))))))))))) (box ((box tail))))))))))))) (Map.add "mulRecord" (box ((box (fun (v: obj) -> (box (fun (ra: obj) -> (box (fun (rb: obj) -> (let tail = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mulRecord))) (box ((box dictSemiringRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box ra)))))) (box ((box rb)))) in let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let insert = (sharpurs_apply (box ((box Record_Unsafe_unsafeSet))) (box ((box key)))) in let get = (sharpurs_apply (box ((box Record_Unsafe_unsafeGet))) (box ((box key)))) in (sharpurs_apply (box ((sharpurs_apply (box ((box insert))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mul))) (box ((box dictSemiring)))))) (box ((sharpurs_apply (box ((box get))) (box ((box ra))))))))) (box ((sharpurs_apply (box ((box get))) (box ((box rb)))))))))))) (box ((box tail))))))))))))) (Map.add "oneRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (let tail = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_oneRecord))) (box ((box dictSemiringRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let insert = (sharpurs_apply (box ((box Record_Unsafe_unsafeSet))) (box ((box key)))) in (sharpurs_apply (box ((sharpurs_apply (box ((box insert))) (box ((box one1)))))) (box ((box tail))))))))))) (Map.add "zeroRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (let tail = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_zeroRecord))) (box ((box dictSemiringRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let insert = (sharpurs_apply (box ((box Record_Unsafe_unsafeSet))) (box ((box key)))) in (sharpurs_apply (box ((sharpurs_apply (box ((box insert))) (box ((box zero1)))))) (box ((box tail))))))))))) Map.empty))))))))))))))))))
