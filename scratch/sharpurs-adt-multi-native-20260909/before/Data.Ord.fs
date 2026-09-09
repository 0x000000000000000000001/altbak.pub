[<AutoOpen>]
module PureScript_Data_Ord

open System
open System.Collections.Generic

module Data_Ord_FFI =
    let ordIntImpl lt eq gt x y = let x' = unbox<int> x in let y' = unbox<int> y in if x' < y' then lt else if x' = y' then eq else gt
    let ordNumberImpl lt eq gt x y = let x' = unbox<float> x in let y' = unbox<float> y in if x' < y' then lt else if x' = y' then eq else gt
    let ordStringImpl lt eq gt x y = let x' = unbox<string> x in let y' = unbox<string> y in if x' < y' then lt else if x' = y' then eq else gt
    let ordCharImpl lt eq gt x y = let x' = unbox<char> x in let y' = unbox<char> y in if x' < y' then lt else if x' = y' then eq else gt
    let ordBooleanImpl lt eq gt x y = let x' = unbox<bool> x in let y' = unbox<bool> y in if x' < y' then lt else if x' = y' then eq else gt
    let ordArrayImpl = undefined
    

let Data_Ord_ordArrayImpl = box (Data_Ord_FFI.``ordArrayImpl``)
let Data_Ord_ordBooleanImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (fun (arg2: obj) -> box (fun (arg3: obj) -> box (fun (arg4: obj) -> box (Data_Ord_FFI.``ordBooleanImpl`` (unbox arg0) (unbox arg1) (unbox arg2) (unbox arg3) (unbox arg4)))))))
let Data_Ord_ordCharImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (fun (arg2: obj) -> box (fun (arg3: obj) -> box (fun (arg4: obj) -> box (Data_Ord_FFI.``ordCharImpl`` (unbox arg0) (unbox arg1) (unbox arg2) (unbox arg3) (unbox arg4)))))))
let Data_Ord_ordIntImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (fun (arg2: obj) -> box (fun (arg3: obj) -> box (fun (arg4: obj) -> box (Data_Ord_FFI.``ordIntImpl`` (unbox arg0) (unbox arg1) (unbox arg2) (unbox arg3) (unbox arg4)))))))
let Data_Ord_ordNumberImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (fun (arg2: obj) -> box (fun (arg3: obj) -> box (fun (arg4: obj) -> box (Data_Ord_FFI.``ordNumberImpl`` (unbox arg0) (unbox arg1) (unbox arg2) (unbox arg3) (unbox arg4)))))))
let Data_Ord_ordStringImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (fun (arg2: obj) -> box (fun (arg3: obj) -> box (fun (arg4: obj) -> box (Data_Ord_FFI.``ordStringImpl`` (unbox arg0) (unbox arg1) (unbox arg2) (unbox arg3) (unbox arg4)))))))


let Data_Ord_eqRec  = (sharpurs_apply (box ((box Data_Eq_eqRec))) (box ((box Prim_undefined))))

let Data_Ord_OrdRecordusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Ord_Ordusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Ord_Ord1usd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Ord_ordVoid  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Ordering_EQusd_Ctor))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Eq_eqVoid))))) Map.empty)))))))

let Data_Ord_ordUnit  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Ordering_EQusd_Ctor))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Eq_eqUnit))))) Map.empty)))))))

let Data_Ord_ordString  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_ordStringImpl))) (box ((box Data_Ordering_LTusd_Ctor)))))) (box ((box Data_Ordering_EQusd_Ctor)))))) (box ((box Data_Ordering_GTusd_Ctor)))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Eq_eqString))))) Map.empty)))))))

let Data_Ord_ordRecordNil  = (sharpurs_apply (box ((box Data_Ord_OrdRecordusd_Dict))) (box ((box ((Map.add "compareRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box Data_Ordering_EQusd_Ctor))))))))) (Map.add "EqRecord0" (box ((box (fun (usd__unused: obj) -> (box Data_Eq_eqRowNil))))) Map.empty)))))))

let Data_Ord_ordProxy  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Ordering_EQusd_Ctor))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Eq_eqProxy))))) Map.empty)))))))

let Data_Ord_ordOrdering  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (Data_Ordering_LTusd_Ctor, Data_Ordering_LTusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor)) | (Data_Ordering_EQusd_Ctor, Data_Ordering_EQusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor)) | (Data_Ordering_GTusd_Ctor, Data_Ordering_GTusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor)) | (Data_Ordering_LTusd_Ctor, _) -> ((box Data_Ordering_LTusd_Ctor)) | (Data_Ordering_EQusd_Ctor, Data_Ordering_LTusd_Ctor) -> ((box Data_Ordering_GTusd_Ctor)) | (Data_Ordering_EQusd_Ctor, Data_Ordering_GTusd_Ctor) -> ((box Data_Ordering_LTusd_Ctor)) | (Data_Ordering_GTusd_Ctor, _) -> ((box Data_Ordering_GTusd_Ctor))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Ordering_eqOrdering))))) Map.empty)))))))

let Data_Ord_ordNumber  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_ordNumberImpl))) (box ((box Data_Ordering_LTusd_Ctor)))))) (box ((box Data_Ordering_EQusd_Ctor)))))) (box ((box Data_Ordering_GTusd_Ctor)))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Eq_eqNumber))))) Map.empty)))))))

let Data_Ord_ordInt  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_ordIntImpl))) (box ((box Data_Ordering_LTusd_Ctor)))))) (box ((box Data_Ordering_EQusd_Ctor)))))) (box ((box Data_Ordering_GTusd_Ctor)))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Eq_eqInt))))) Map.empty)))))))

let Data_Ord_ordChar  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_ordCharImpl))) (box ((box Data_Ordering_LTusd_Ctor)))))) (box ((box Data_Ordering_EQusd_Ctor)))))) (box ((box Data_Ordering_GTusd_Ctor)))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Eq_eqChar))))) Map.empty)))))))

let Data_Ord_ordBoolean  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_ordBooleanImpl))) (box ((box Data_Ordering_LTusd_Ctor)))))) (box ((box Data_Ordering_EQusd_Ctor)))))) (box ((box Data_Ordering_GTusd_Ctor)))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Eq_eqBoolean))))) Map.empty)))))))

let Data_Ord_compareRecord  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "compareRecord" (unbox<Map<string, obj>> ((box v))))))))

let Data_Ord_ordRecord  = (box (fun (usd__unused: obj) -> (box (fun (dictOrdRecord: obj) -> (let eqRec1 = (sharpurs_apply (box ((box Data_Ord_eqRec))) (box ((sharpurs_apply (box ((Map.find "EqRecord0" (unbox<Map<string, obj>> ((box dictOrdRecord)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compareRecord))) (box ((box dictOrdRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box eqRec1))))) Map.empty))))))))))))

let Data_Ord_compare1  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "compare1" (unbox<Map<string, obj>> ((box v))))))))

let Data_Ord_compare  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "compare" (unbox<Map<string, obj>> ((box v))))))))

let Data_Ord_comparing  = (box (fun (dictOrd: obj) -> (box (fun (f: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((sharpurs_apply (box ((box f))) (box ((box x))))))))) (box ((sharpurs_apply (box ((box f))) (box ((box y)))))))))))))))

let Data_Ord_greaterThan  = (box (fun (dictOrd: obj) -> (box (fun (a1: obj) -> (box (fun (a2: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((box a1)))))) (box ((box a2)))) in (match ((unbox ((box v)))) with | Data_Ordering_GTusd_Ctor -> ((box true)) | _ -> ((box false))))))))))

let Data_Ord_greaterThanOrEq  = (box (fun (dictOrd: obj) -> (box (fun (a1: obj) -> (box (fun (a2: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((box a1)))))) (box ((box a2)))) in (match ((unbox ((box v)))) with | Data_Ordering_LTusd_Ctor -> ((box false)) | _ -> ((box true))))))))))

let Data_Ord_lessThan  = (box (fun (dictOrd: obj) -> (box (fun (a1: obj) -> (box (fun (a2: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((box a1)))))) (box ((box a2)))) in (match ((unbox ((box v)))) with | Data_Ordering_LTusd_Ctor -> ((box true)) | _ -> ((box false))))))))))

let Data_Ord_signum  = (box (fun (dictOrd: obj) -> (box (fun (dictRing: obj) -> (let Semiring0 = (sharpurs_apply (box ((Map.find "Semiring0" (unbox<Map<string, obj>> ((box dictRing)))))) (box ((box Prim_undefined)))) in let zero = (sharpurs_apply (box ((box Data_Semiring_zero))) (box ((box Semiring0)))) in let Semiring01 = (sharpurs_apply (box ((Map.find "Semiring0" (unbox<Map<string, obj>> ((box dictRing)))))) (box ((box Prim_undefined)))) in let one = (sharpurs_apply (box ((box Data_Semiring_one))) (box ((box Semiring0)))) in (box (fun (x: obj) -> (match ((unbox ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box zero))))))) with | LitBool true () -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_negate))) (box ((box dictRing)))))) (box ((sharpurs_apply (box ((box Data_Semiring_one))) (box ((box Semiring01)))))))) | _ -> ((match ((unbox ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_greaterThan))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box zero))))))) with | LitBool true () -> ((box one)) | _ -> ((box x))))))))))))

let Data_Ord_lessThanOrEq  = (box (fun (dictOrd: obj) -> (box (fun (a1: obj) -> (box (fun (a2: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((box a1)))))) (box ((box a2)))) in (match ((unbox ((box v)))) with | Data_Ordering_GTusd_Ctor -> ((box false)) | _ -> ((box true))))))))))

let Data_Ord_max  = (box (fun (dictOrd: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box y)))) in (match ((unbox ((box v)))) with | Data_Ordering_LTusd_Ctor -> ((box y)) | Data_Ordering_EQusd_Ctor -> ((box x)) | Data_Ordering_GTusd_Ctor -> ((box x))))))))))

let Data_Ord_min  = (box (fun (dictOrd: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box y)))) in (match ((unbox ((box v)))) with | Data_Ordering_LTusd_Ctor -> ((box x)) | Data_Ordering_EQusd_Ctor -> ((box x)) | Data_Ordering_GTusd_Ctor -> ((box y))))))))))

let Data_Ord_ordArray  = (box (fun (dictOrd: obj) -> (let eqArray = (sharpurs_apply (box ((box Data_Eq_eqArray))) (box ((sharpurs_apply (box ((Map.find "Eq0" (unbox<Map<string, obj>> ((box dictOrd)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((let toDelta = (box (fun (x: obj) -> (box (fun (y: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box y)))) in (match ((unbox ((box v)))) with | Data_Ordering_EQusd_Ctor -> ((box 0)) | Data_Ordering_LTusd_Ctor -> ((box 1)) | Data_Ordering_GTusd_Ctor -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_negate))) (box ((box Data_Ring_ringInt)))))) (box ((box 1))))))))))) in (box (fun (xs: obj) -> (box (fun (ys: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box Data_Ord_ordInt)))))) (box ((box 0)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_ordArrayImpl))) (box ((box toDelta)))))) (box ((box xs)))))) (box ((box ys)))))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box eqArray))))) Map.empty))))))))))

let Data_Ord_ord1Array  = (sharpurs_apply (box ((box Data_Ord_Ord1usd_Dict))) (box ((box ((Map.add "compare1" (box ((box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_compare))) (box ((sharpurs_apply (box ((box Data_Ord_ordArray))) (box ((box dictOrd))))))))))) (Map.add "Eq10" (box ((box (fun (usd__unused: obj) -> (box Data_Eq_eq1Array))))) Map.empty)))))))

let Data_Ord_ordRecordCons  = (box (fun (dictOrdRecord: obj) -> (let eqRowCons = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eqRowCons))) (box ((sharpurs_apply (box ((Map.find "EqRecord0" (unbox<Map<string, obj>> ((box dictOrdRecord)))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined)))) in (box (fun (usd__unused: obj) -> (box (fun (dictIsSymbol: obj) -> (let eqRowCons1 = (sharpurs_apply (box ((box eqRowCons))) (box ((box dictIsSymbol)))) in (box (fun (dictOrd: obj) -> (let eqRowCons2 = (sharpurs_apply (box ((box eqRowCons1))) (box ((sharpurs_apply (box ((Map.find "Eq0" (unbox<Map<string, obj>> ((box dictOrd)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_OrdRecordusd_Dict))) (box ((box ((Map.add "compareRecord" (box ((box (fun (v: obj) -> (box (fun (ra: obj) -> (box (fun (rb: obj) -> (let unsafeGet_prime = (box Record_Unsafe_unsafeGet) in let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let left = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box unsafeGet_prime))) (box ((box key)))))) (box ((box ra))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box unsafeGet_prime))) (box ((box key)))))) (box ((box rb))))))) in (match ((unbox ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_notEq))) (box ((box Data_Ordering_eqOrdering)))))) (box ((box left)))))) (box ((box Data_Ordering_EQusd_Ctor))))))) with | LitBool true () -> ((box left)) | _ -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compareRecord))) (box ((box dictOrdRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box ra)))))) (box ((box rb))))))))))))))) (Map.add "EqRecord0" (box ((box (fun (usd__unused: obj) -> (box eqRowCons2))))) Map.empty))))))))))))))))))

let Data_Ord_clamp  = (box (fun (dictOrd: obj) -> (box (fun (low: obj) -> (box (fun (hi: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_min))) (box ((box dictOrd)))))) (box ((box hi)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_max))) (box ((box dictOrd)))))) (box ((box low)))))) (box ((box x)))))))))))))))

let Data_Ord_between  = (box (fun (dictOrd: obj) -> (box (fun (low: obj) -> (box (fun (hi: obj) -> (box (fun (x: obj) -> (match (((unbox ((box low))), (unbox ((box hi))), (unbox ((box x))))) with | (low1, hi1, x1) when (unbox (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThan))) (box ((box dictOrd)))))) (box ((box x1)))))) (box ((box low1))))) -> ((box false)) | (low1, hi1, x1) when (unbox (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_greaterThan))) (box ((box dictOrd)))))) (box ((box x1)))))) (box ((box hi1))))) -> ((box false)) | (low1, hi1, x1) when (unbox (box true)) -> ((box true)))))))))))

let Data_Ord_abs  = (box (fun (dictOrd: obj) -> (box (fun (dictRing: obj) -> (let zero = (sharpurs_apply (box ((box Data_Semiring_zero))) (box ((sharpurs_apply (box ((Map.find "Semiring0" (unbox<Map<string, obj>> ((box dictRing)))))) (box ((box Prim_undefined))))))) in (box (fun (x: obj) -> (match ((unbox ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_greaterThanOrEq))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box zero))))))) with | LitBool true () -> ((box x)) | _ -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_negate))) (box ((box dictRing)))))) (box ((box x)))))))))))))
