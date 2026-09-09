[<AutoOpen>]
module PureScript_Data_Number

open System
open System.Collections.Generic

module Data_Number_FFI =
    let nan = box System.Double.NaN
    let isNaN = fun (x: obj) -> box (System.Double.IsNaN(x :?> float))
    let infinity = box System.Double.PositiveInfinity
    let isFinite = fun (x: obj) -> box (not (System.Double.IsInfinity(x :?> float)) && not (System.Double.IsNaN(x :?> float)))
    
    let fromStringImpl = 
        fun (strVal: obj) -> fun (isFiniteFn: obj) -> fun (just: obj) -> fun (nothing: obj) ->
            let str = strVal :?> string
            match System.Double.TryParse(str, System.Globalization.NumberStyles.Float, System.Globalization.CultureInfo.InvariantCulture) with
            | true, num -> 
                let isFin = isFiniteFn :?> (obj -> obj)
                let just' = just :?> (obj -> obj)
                if isFin (box num) :?> bool then
                    just' (box num)
                else
                    nothing
            | _ -> nothing
    
    let abs = fun (x: obj) -> box (System.Math.Abs(x :?> float))
    let acos = fun (x: obj) -> box (System.Math.Acos(x :?> float))
    let asin = fun (x: obj) -> box (System.Math.Asin(x :?> float))
    let atan = fun (x: obj) -> box (System.Math.Atan(x :?> float))
    let atan2 = fun (y: obj) -> fun (x: obj) -> box (System.Math.Atan2(y :?> float, x :?> float))
    let ceil = fun (x: obj) -> box (System.Math.Ceiling(x :?> float))
    let cos = fun (x: obj) -> box (System.Math.Cos(x :?> float))
    let exp = fun (x: obj) -> box (System.Math.Exp(x :?> float))
    let floor = fun (x: obj) -> box (System.Math.Floor(x :?> float))
    let log = fun (x: obj) -> box (System.Math.Log(x :?> float))
    let max = fun (n1: obj) -> fun (n2: obj) -> box (System.Math.Max(n1 :?> float, n2 :?> float))
    let min = fun (n1: obj) -> fun (n2: obj) -> box (System.Math.Min(n1 :?> float, n2 :?> float))
    let pow = fun (n: obj) -> fun (p: obj) -> box (System.Math.Pow(n :?> float, p :?> float))
    let remainder = fun (n: obj) -> fun (m: obj) -> box ((n :?> float) % (m :?> float))
    let round = fun (x: obj) -> box (System.Math.Round(x :?> float))
    let sign = fun (x: obj) -> box (float (System.Math.Sign(x :?> float)))
    let sin = fun (x: obj) -> box (System.Math.Sin(x :?> float))
    let sqrt = fun (x: obj) -> box (System.Math.Sqrt(x :?> float))
    let tan = fun (x: obj) -> box (System.Math.Tan(x :?> float))
    let trunc = fun (x: obj) -> box (System.Math.Truncate(x :?> float))
    

let Data_Number_abs = box (Data_Number_FFI.``abs``)
let Data_Number_acos = box (Data_Number_FFI.``acos``)
let Data_Number_asin = box (Data_Number_FFI.``asin``)
let Data_Number_atan = box (Data_Number_FFI.``atan``)
let Data_Number_atan2 = box (Data_Number_FFI.``atan2``)
let Data_Number_ceil = box (Data_Number_FFI.``ceil``)
let Data_Number_cos = box (Data_Number_FFI.``cos``)
let Data_Number_exp = box (Data_Number_FFI.``exp``)
let Data_Number_floor = box (Data_Number_FFI.``floor``)
let Data_Number_fromStringImpl = box (Data_Number_FFI.``fromStringImpl``)
let Data_Number_infinity = box (Data_Number_FFI.``infinity``)
let Data_Number_isFinite = box (Data_Number_FFI.``isFinite``)
let Data_Number_isNaN = box (Data_Number_FFI.``isNaN``)
let Data_Number_log = box (Data_Number_FFI.``log``)
let Data_Number_max = box (Data_Number_FFI.``max``)
let Data_Number_min = box (Data_Number_FFI.``min``)
let Data_Number_nan = box (Data_Number_FFI.``nan``)
let Data_Number_pow = box (Data_Number_FFI.``pow``)
let Data_Number_remainder = box (Data_Number_FFI.``remainder``)
let Data_Number_round = box (Data_Number_FFI.``round``)
let Data_Number_sign = box (Data_Number_FFI.``sign``)
let Data_Number_sin = box (Data_Number_FFI.``sin``)
let Data_Number_sqrt = box (Data_Number_FFI.``sqrt``)
let Data_Number_tan = box (Data_Number_FFI.``tan``)
let Data_Number_trunc = box (Data_Number_FFI.``trunc``)


let Data_Number_tau  = (box 6.283185307179586)

let Data_Number_sqrt2  = (box 1.4142135623730951)

let Data_Number_sqrt1_2  = (box 0.7071067811865476)

let Data_Number_pi  = (box 3.141592653589793)

let Data_Number_log2e  = (box 1.4426950408889634)

let Data_Number_log10e  = (box 0.4342944819032518)

let Data_Number_ln2  = (box 0.6931471805599453)

let Data_Number_ln10  = (box 2.302585092994046)

let Data_Number_fromString  = (box (fun (str: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_Uncurried_runFn4))) (box ((box Data_Number_fromStringImpl)))))) (box ((box str)))))) (box ((box Data_Number_isFinite)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1))))))))))) (box ((box Data_Maybe_Nothingusd_Ctor))))))

let Data_Number_e  = (box 2.718281828459045)
