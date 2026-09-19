[<AutoOpen>]
module PureScript_Data_Number_Format

open System
open System.Collections.Generic

module Data_Number_Format_FFI =
    let toPrecisionNative = 
        fun (dVal: obj) -> fun (numVal: obj) ->
            let d = dVal :?> int
            let num = numVal :?> float
            box (num.ToString("G" + d.ToString(), System.Globalization.CultureInfo.InvariantCulture))
    
    let toFixedNative = 
        fun (dVal: obj) -> fun (numVal: obj) ->
            let d = dVal :?> int
            let num = numVal :?> float
            box (num.ToString("F" + d.ToString(), System.Globalization.CultureInfo.InvariantCulture))
    
    let toExponentialNative = 
        fun (dVal: obj) -> fun (numVal: obj) ->
            let d = dVal :?> int
            let num = numVal :?> float
            box (num.ToString("E" + d.ToString(), System.Globalization.CultureInfo.InvariantCulture))
    
    let toString = 
        fun (numVal: obj) ->
            let num = numVal :?> float
            box (num.ToString(System.Globalization.CultureInfo.InvariantCulture))
    

let Data_Number_Format_toExponentialNative = box (Data_Number_Format_FFI.``toExponentialNative``)
let Data_Number_Format_toFixedNative = box (Data_Number_Format_FFI.``toFixedNative``)
let Data_Number_Format_toPrecisionNative = box (Data_Number_Format_FFI.``toPrecisionNative``)
let Data_Number_Format_toString = box (Data_Number_Format_FFI.``toString``)


type Data_Number_Format_Format =
  | Data_Number_Format_Precisionusd_Ctor of obj
  | Data_Number_Format_Fixedusd_Ctor of obj
  | Data_Number_Format_Exponentialusd_Ctor of obj

let Data_Number_Format_Precision  = (box ((fun (usd__arg1: obj) -> (box (Data_Number_Format_Precisionusd_Ctor(usd__arg1))))))

let Data_Number_Format_Fixed  = (box ((fun (usd__arg1: obj) -> (box (Data_Number_Format_Fixedusd_Ctor(usd__arg1))))))

let Data_Number_Format_Exponential  = (box ((fun (usd__arg1: obj) -> (box (Data_Number_Format_Exponentialusd_Ctor(usd__arg1))))))

let Data_Number_Format_toStringWith  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Number_Format_Precisionusd_Ctor(p) -> ((sharpurs_apply (box ((box Data_Number_Format_toPrecisionNative))) (box ((box p))))) | Data_Number_Format_Fixedusd_Ctor(p) -> ((sharpurs_apply (box ((box Data_Number_Format_toFixedNative))) (box ((box p))))) | Data_Number_Format_Exponentialusd_Ctor(p) -> ((sharpurs_apply (box ((box Data_Number_Format_toExponentialNative))) (box ((box p))))))))

let Data_Number_Format_precision  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Number_Format_Precisionusd_Ctor(usd__arg1))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_clamp))) (box ((box Data_Ord_ordInt)))))) (box ((box 1)))))) (box ((box 21)))))))

let Data_Number_Format_fixed  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Number_Format_Fixedusd_Ctor(usd__arg1))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_clamp))) (box ((box Data_Ord_ordInt)))))) (box ((box 0)))))) (box ((box 20)))))))

let Data_Number_Format_exponential  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Number_Format_Exponentialusd_Ctor(usd__arg1))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_clamp))) (box ((box Data_Ord_ordInt)))))) (box ((box 0)))))) (box ((box 20)))))))
