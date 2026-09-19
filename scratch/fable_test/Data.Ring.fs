[<AutoOpen>]
module PureScript_Data_Ring

open System
open System.Collections.Generic

module Data_Ring_FFI =
    let intSub a b = (unbox<int> a) - (unbox<int> b)
    let numSub a b = (unbox<float> a) - (unbox<float> b)
    

let Data_Ring_intSub = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Ring_FFI.``intSub`` (unbox arg0) (unbox arg1))))
let Data_Ring_numSub = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Ring_FFI.``numSub`` (unbox arg0) (unbox arg1))))


let Data_Ring_semiringRecord  = (sharpurs_apply (box ((box Data_Semiring_semiringRecord))) (box ((box Prim_undefined))))

let Data_Ring_RingRecordusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Ring_Ringusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Ring_subRecord  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "subRecord" (unbox<Map<string, obj>> ((box v))))))))

let Data_Ring_sub  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "sub" (unbox<Map<string, obj>> ((box v))))))))

let Data_Ring_ringUnit  = (sharpurs_apply (box ((box Data_Ring_Ringusd_Dict))) (box ((box ((Map.add "sub" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Unit_unit))))))) (Map.add "Semiring0" (box ((box (fun (usd__unused: obj) -> (box Data_Semiring_semiringUnit))))) Map.empty)))))))

let Data_Ring_ringRecordNil  = (sharpurs_apply (box ((box Data_Ring_RingRecordusd_Dict))) (box ((box ((Map.add "subRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (Map.empty)))))))))) (Map.add "SemiringRecord0" (box ((box (fun (usd__unused: obj) -> (box Data_Semiring_semiringRecordNil))))) Map.empty)))))))

let Data_Ring_ringRecordCons  = (box (fun (dictIsSymbol: obj) -> (let semiringRecordCons = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_semiringRecordCons))) (box ((box dictIsSymbol)))))) (box ((box Prim_undefined)))) in (box (fun (usd__unused: obj) -> (box (fun (dictRingRecord: obj) -> (let semiringRecordCons1 = (sharpurs_apply (box ((box semiringRecordCons))) (box ((sharpurs_apply (box ((Map.find "SemiringRecord0" (unbox<Map<string, obj>> ((box dictRingRecord)))))) (box ((box Prim_undefined))))))) in (box (fun (dictRing: obj) -> (let semiringRecordCons2 = (sharpurs_apply (box ((box semiringRecordCons1))) (box ((sharpurs_apply (box ((Map.find "Semiring0" (unbox<Map<string, obj>> ((box dictRing)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ring_RingRecordusd_Dict))) (box ((box ((Map.add "subRecord" (box ((box (fun (v: obj) -> (box (fun (ra: obj) -> (box (fun (rb: obj) -> (let tail = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_subRecord))) (box ((box dictRingRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box ra)))))) (box ((box rb)))) in let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let insert = (sharpurs_apply (box ((box Record_Unsafe_unsafeSet))) (box ((box key)))) in let get = (sharpurs_apply (box ((box Record_Unsafe_unsafeGet))) (box ((box key)))) in (sharpurs_apply (box ((sharpurs_apply (box ((box insert))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box dictRing)))))) (box ((sharpurs_apply (box ((box get))) (box ((box ra))))))))) (box ((sharpurs_apply (box ((box get))) (box ((box rb)))))))))))) (box ((box tail))))))))))))) (Map.add "SemiringRecord0" (box ((box (fun (usd__unused: obj) -> (box semiringRecordCons2))))) Map.empty))))))))))))))))))

let Data_Ring_ringRecord  = (box (fun (usd__unused: obj) -> (box (fun (dictRingRecord: obj) -> (let semiringRecord1 = (sharpurs_apply (box ((box Data_Ring_semiringRecord))) (box ((sharpurs_apply (box ((Map.find "SemiringRecord0" (unbox<Map<string, obj>> ((box dictRingRecord)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ring_Ringusd_Dict))) (box ((box ((Map.add "sub" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_subRecord))) (box ((box dictRingRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (Map.add "Semiring0" (box ((box (fun (usd__unused: obj) -> (box semiringRecord1))))) Map.empty))))))))))))

let Data_Ring_ringProxy  = (sharpurs_apply (box ((box Data_Ring_Ringusd_Dict))) (box ((box ((Map.add "sub" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Type_Proxy_Proxyusd_Ctor))))))) (Map.add "Semiring0" (box ((box (fun (usd__unused: obj) -> (box Data_Semiring_semiringProxy))))) Map.empty)))))))

let Data_Ring_ringNumber  = (sharpurs_apply (box ((box Data_Ring_Ringusd_Dict))) (box ((box ((Map.add "sub" (box ((box Data_Ring_numSub))) (Map.add "Semiring0" (box ((box (fun (usd__unused: obj) -> (box Data_Semiring_semiringNumber))))) Map.empty)))))))

let Data_Ring_ringInt  = (sharpurs_apply (box ((box Data_Ring_Ringusd_Dict))) (box ((box ((Map.add "sub" (box ((box Data_Ring_intSub))) (Map.add "Semiring0" (box ((box (fun (usd__unused: obj) -> (box Data_Semiring_semiringInt))))) Map.empty)))))))

let Data_Ring_ringFn  = (box (fun (dictRing: obj) -> (let semiringFn = (sharpurs_apply (box ((box Data_Semiring_semiringFn))) (box ((sharpurs_apply (box ((Map.find "Semiring0" (unbox<Map<string, obj>> ((box dictRing)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ring_Ringusd_Dict))) (box ((box ((Map.add "sub" (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box dictRing)))))) (box ((sharpurs_apply (box ((box f))) (box ((box x))))))))) (box ((sharpurs_apply (box ((box g))) (box ((box x))))))))))))))) (Map.add "Semiring0" (box ((box (fun (usd__unused: obj) -> (box semiringFn))))) Map.empty))))))))))

let Data_Ring_negate  = (box (fun (dictRing: obj) -> (let Semiring0 = (sharpurs_apply (box ((Map.find "Semiring0" (unbox<Map<string, obj>> ((box dictRing)))))) (box ((box Prim_undefined)))) in (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box dictRing)))))) (box ((sharpurs_apply (box ((box Data_Semiring_zero))) (box ((box Semiring0))))))))) (box ((box a)))))))))
