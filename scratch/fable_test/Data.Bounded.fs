[<AutoOpen>]
module PureScript_Data_Bounded

open System
open System.Collections.Generic

module Data_Bounded_FFI =
    let topChar = '\uffff'
    let bottomChar = '\u0000'
    let topNumber = System.Double.MaxValue
    let bottomNumber = System.Double.MinValue
    let topInt = System.Int32.MaxValue
    let bottomInt = System.Int32.MinValue
    

let Data_Bounded_bottomChar = box (Data_Bounded_FFI.``bottomChar``)
let Data_Bounded_bottomInt = box (Data_Bounded_FFI.``bottomInt``)
let Data_Bounded_bottomNumber = box (Data_Bounded_FFI.``bottomNumber``)
let Data_Bounded_topChar = box (Data_Bounded_FFI.``topChar``)
let Data_Bounded_topInt = box (Data_Bounded_FFI.``topInt``)
let Data_Bounded_topNumber = box (Data_Bounded_FFI.``topNumber``)


let Data_Bounded_ordRecord  = (sharpurs_apply (box ((box Data_Ord_ordRecord))) (box ((box Prim_undefined))))

let Data_Bounded_BoundedRecordusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Bounded_Boundedusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Bounded_topRecord  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "topRecord" (unbox<Map<string, obj>> ((box v))))))))

let Data_Bounded_top  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "top" (unbox<Map<string, obj>> ((box v))))))))

let Data_Bounded_boundedUnit  = (sharpurs_apply (box ((box Data_Bounded_Boundedusd_Dict))) (box ((box ((Map.add "top" (box ((box Data_Unit_unit))) (Map.add "bottom" (box ((box Data_Unit_unit))) (Map.add "Ord0" (box ((box (fun (usd__unused: obj) -> (box Data_Ord_ordUnit))))) Map.empty))))))))

let Data_Bounded_boundedRecordNil  = (sharpurs_apply (box ((box Data_Bounded_BoundedRecordusd_Dict))) (box ((box ((Map.add "topRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (Map.empty)))))))) (Map.add "bottomRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (Map.empty)))))))) (Map.add "OrdRecord0" (box ((box (fun (usd__unused: obj) -> (box Data_Ord_ordRecordNil))))) Map.empty))))))))

let Data_Bounded_boundedProxy  = (sharpurs_apply (box ((box Data_Bounded_Boundedusd_Dict))) (box ((box ((Map.add "bottom" (box ((box Type_Proxy_Proxyusd_Ctor))) (Map.add "top" (box ((box Type_Proxy_Proxyusd_Ctor))) (Map.add "Ord0" (box ((box (fun (usd__unused: obj) -> (box Data_Ord_ordProxy))))) Map.empty))))))))

let Data_Bounded_boundedOrdering  = (sharpurs_apply (box ((box Data_Bounded_Boundedusd_Dict))) (box ((box ((Map.add "top" (box ((box Data_Ordering_GTusd_Ctor))) (Map.add "bottom" (box ((box Data_Ordering_LTusd_Ctor))) (Map.add "Ord0" (box ((box (fun (usd__unused: obj) -> (box Data_Ord_ordOrdering))))) Map.empty))))))))

let Data_Bounded_boundedNumber  = (sharpurs_apply (box ((box Data_Bounded_Boundedusd_Dict))) (box ((box ((Map.add "top" (box ((box Data_Bounded_topNumber))) (Map.add "bottom" (box ((box Data_Bounded_bottomNumber))) (Map.add "Ord0" (box ((box (fun (usd__unused: obj) -> (box Data_Ord_ordNumber))))) Map.empty))))))))

let Data_Bounded_boundedInt  = (sharpurs_apply (box ((box Data_Bounded_Boundedusd_Dict))) (box ((box ((Map.add "top" (box ((box Data_Bounded_topInt))) (Map.add "bottom" (box ((box Data_Bounded_bottomInt))) (Map.add "Ord0" (box ((box (fun (usd__unused: obj) -> (box Data_Ord_ordInt))))) Map.empty))))))))

let Data_Bounded_boundedChar  = (sharpurs_apply (box ((box Data_Bounded_Boundedusd_Dict))) (box ((box ((Map.add "top" (box ((box Data_Bounded_topChar))) (Map.add "bottom" (box ((box Data_Bounded_bottomChar))) (Map.add "Ord0" (box ((box (fun (usd__unused: obj) -> (box Data_Ord_ordChar))))) Map.empty))))))))

let Data_Bounded_boundedBoolean  = (sharpurs_apply (box ((box Data_Bounded_Boundedusd_Dict))) (box ((box ((Map.add "top" (box ((box true))) (Map.add "bottom" (box ((box false))) (Map.add "Ord0" (box ((box (fun (usd__unused: obj) -> (box Data_Ord_ordBoolean))))) Map.empty))))))))

let Data_Bounded_bottomRecord  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "bottomRecord" (unbox<Map<string, obj>> ((box v))))))))

let Data_Bounded_boundedRecord  = (box (fun (usd__unused: obj) -> (box (fun (dictBoundedRecord: obj) -> (let ordRecord1 = (sharpurs_apply (box ((box Data_Bounded_ordRecord))) (box ((sharpurs_apply (box ((Map.find "OrdRecord0" (unbox<Map<string, obj>> ((box dictBoundedRecord)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Bounded_Boundedusd_Dict))) (box ((box ((Map.add "top" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Bounded_topRecord))) (box ((box dictBoundedRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (Map.add "bottom" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Bounded_bottomRecord))) (box ((box dictBoundedRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (Map.add "Ord0" (box ((box (fun (usd__unused: obj) -> (box ordRecord1))))) Map.empty)))))))))))))

let Data_Bounded_bottom  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "bottom" (unbox<Map<string, obj>> ((box v))))))))

let Data_Bounded_boundedRecordCons  = (box (fun (dictIsSymbol: obj) -> (box (fun (dictBounded: obj) -> (let top1 = (sharpurs_apply (box ((box Data_Bounded_top))) (box ((box dictBounded)))) in let bottom1 = (sharpurs_apply (box ((box Data_Bounded_bottom))) (box ((box dictBounded)))) in let Ord0 = (sharpurs_apply (box ((Map.find "Ord0" (unbox<Map<string, obj>> ((box dictBounded)))))) (box ((box Prim_undefined)))) in (box (fun (usd__unused: obj) -> (box (fun (usd__unused: obj) -> (box (fun (dictBoundedRecord: obj) -> (let ordRecordCons = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_ordRecordCons))) (box ((sharpurs_apply (box ((Map.find "OrdRecord0" (unbox<Map<string, obj>> ((box dictBoundedRecord)))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined)))))) (box ((box dictIsSymbol)))))) (box ((box Ord0)))) in (sharpurs_apply (box ((box Data_Bounded_BoundedRecordusd_Dict))) (box ((box ((Map.add "topRecord" (box ((box (fun (v: obj) -> (box (fun (rowProxy: obj) -> (let tail = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Bounded_topRecord))) (box ((box dictBoundedRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box rowProxy)))) in let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let insert = (sharpurs_apply (box ((box Record_Unsafe_unsafeSet))) (box ((box key)))) in (sharpurs_apply (box ((sharpurs_apply (box ((box insert))) (box ((box top1)))))) (box ((box tail))))))))))) (Map.add "bottomRecord" (box ((box (fun (v: obj) -> (box (fun (rowProxy: obj) -> (let tail = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Bounded_bottomRecord))) (box ((box dictBoundedRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box rowProxy)))) in let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let insert = (sharpurs_apply (box ((box Record_Unsafe_unsafeSet))) (box ((box key)))) in (sharpurs_apply (box ((sharpurs_apply (box ((box insert))) (box ((box bottom1)))))) (box ((box tail))))))))))) (Map.add "OrdRecord0" (box ((box (fun (usd__unused: obj) -> (box ordRecordCons))))) Map.empty))))))))))))))))))))
