[<AutoOpen>]
module PureScript_Data_Int

open System
open System.Collections.Generic

module Data_Int_FFI =
    let fromNumberImpl = 
        fun (just: obj) -> fun (nothing: obj) -> fun (nVal: obj) ->
            let n = nVal :?> float
            if n = float (int n) then
                let just' = just :?> (obj -> obj)
                just' (box (int n))
            else
                nothing
    
    let toNumber = 
        fun (n: obj) -> box (float (n :?> int))
    
    let fromStringAsImpl = 
        fun (just: obj) -> fun (nothing: obj) -> fun (radixVal: obj) ->
            let radix = radixVal :?> int
            let just' = just :?> (obj -> obj)
            fun (sVal: obj) ->
                let s = sVal :?> string
                try
                    let i = System.Convert.ToInt32(s, radix)
                    just' (box i)
                with
                | _ -> nothing
    
    let toStringAs = 
        fun (radixVal: obj) -> fun (iVal: obj) ->
            let radix = radixVal :?> int
            let i = iVal :?> int
            box (System.Convert.ToString(i, radix))
    
    let quot = 
        fun (xVal: obj) -> fun (yVal: obj) ->
            let x = xVal :?> int
            let y = yVal :?> int
            box (x / y)
    
    let rem = 
        fun (xVal: obj) -> fun (yVal: obj) ->
            let x = xVal :?> int
            let y = yVal :?> int
            box (x % y)
    
    let pow = 
        fun (xVal: obj) -> fun (yVal: obj) ->
            let x = xVal :?> int
            let y = yVal :?> int
            box (int (System.Math.Pow(float x, float y)))
    

let Data_Int_fromNumberImpl = box (Data_Int_FFI.``fromNumberImpl``)
let Data_Int_fromStringAsImpl = box (Data_Int_FFI.``fromStringAsImpl``)
let Data_Int_pow = box (Data_Int_FFI.``pow``)
let Data_Int_quot = box (Data_Int_FFI.``quot``)
let Data_Int_rem = box (Data_Int_FFI.``rem``)
let Data_Int_toNumber = box (Data_Int_FFI.``toNumber``)
let Data_Int_toStringAs = box (Data_Int_FFI.``toStringAs``)


type Data_Int_Parity =
  | Data_Int_Evenusd_Ctor
  | Data_Int_Oddusd_Ctor

let Data_Int_top  = (sharpurs_apply (box ((box Data_Bounded_top))) (box ((box Data_Bounded_boundedInt))))

let Data_Int_bottom  = (sharpurs_apply (box ((box Data_Bounded_bottom))) (box ((box Data_Bounded_boundedInt))))

let Data_Int_Radix  = (box (fun (x: obj) -> (box x)))

let Data_Int_Even  = (box Data_Int_Evenusd_Ctor)

let Data_Int_Odd  = (box Data_Int_Oddusd_Ctor)

let Data_Int_showParity  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Int_Evenusd_Ctor -> ((box "Even")) | Data_Int_Oddusd_Ctor -> ((box "Odd"))))))) Map.empty))))))

let Data_Int_radix  = (box (fun (n: obj) -> (match ((unbox ((box n)))) with | n1 when (unbox (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_conj))) (box ((box Data_HeytingAlgebra_heytingAlgebraBoolean)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_greaterThanOrEq))) (box ((box Data_Ord_ordInt)))))) (box ((box n1)))))) (box ((box 2))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThanOrEq))) (box ((box Data_Ord_ordInt)))))) (box ((box n1)))))) (box ((box 36)))))))) -> ((box (Data_Maybe_Justusd_Ctor((sharpurs_apply (box ((box Data_Int_Radix))) (box ((box n1)))))))) | n1 when (unbox (box Data_Boolean_otherwise)) -> ((box Data_Maybe_Nothingusd_Ctor)))))

let Data_Int_odd  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_notEq))) (box ((box Data_Eq_eqInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Int_Bits_and))) (box ((box x)))))) (box ((box 1))))))))) (box ((box 0))))))

let Data_Int_octal  = (sharpurs_apply (box ((box Data_Int_Radix))) (box ((box 8))))

let Data_Int_hexadecimal  = (sharpurs_apply (box ((box Data_Int_Radix))) (box ((box 16))))

let Data_Int_fromStringAs  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Int_fromStringAsImpl))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1))))))))))) (box ((box Data_Maybe_Nothingusd_Ctor))))

let Data_Int_fromString  = (sharpurs_apply (box ((box Data_Int_fromStringAs))) (box ((sharpurs_apply (box ((box Data_Int_Radix))) (box ((box 10)))))))

let Data_Int_fromNumber  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Int_fromNumberImpl))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1))))))))))) (box ((box Data_Maybe_Nothingusd_Ctor))))

let Data_Int_unsafeClamp  = (box (fun (x: obj) -> (match ((unbox ((box x)))) with | x1 when (unbox (sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_not))) (box ((box Data_HeytingAlgebra_heytingAlgebraBoolean)))))) (box ((sharpurs_apply (box ((box Data_Number_isFinite))) (box ((box x1)))))))) -> ((box 0)) | x1 when (unbox (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_greaterThanOrEq))) (box ((box Data_Ord_ordNumber)))))) (box ((box x1)))))) (box ((sharpurs_apply (box ((box Data_Int_toNumber))) (box ((box Data_Int_top)))))))) -> ((box Data_Int_top)) | x1 when (unbox (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThanOrEq))) (box ((box Data_Ord_ordNumber)))))) (box ((box x1)))))) (box ((sharpurs_apply (box ((box Data_Int_toNumber))) (box ((box Data_Int_bottom)))))))) -> ((box Data_Int_bottom)) | x1 when (unbox (box Data_Boolean_otherwise)) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Maybe_fromMaybe))) (box ((box 0)))))) (box ((sharpurs_apply (box ((box Data_Int_fromNumber))) (box ((box x1)))))))))))

let Data_Int_round  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Int_unsafeClamp)))))) (box ((box Data_Number_round))))

let Data_Int_trunc  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Int_unsafeClamp)))))) (box ((box Data_Number_trunc))))

let Data_Int_floor  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Int_unsafeClamp)))))) (box ((box Data_Number_floor))))

let Data_Int_even  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Eq_eqInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Int_Bits_and))) (box ((box x)))))) (box ((box 1))))))))) (box ((box 0))))))

let Data_Int_parity  = (box (fun (n: obj) -> (match ((unbox ((sharpurs_apply (box ((box Data_Int_even))) (box ((box n))))))) with | LitBool true () -> ((box Data_Int_Evenusd_Ctor)) | _ -> ((box Data_Int_Oddusd_Ctor)))))

let Data_Int_eqParity  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (Data_Int_Evenusd_Ctor, Data_Int_Evenusd_Ctor) -> ((box true)) | (Data_Int_Oddusd_Ctor, Data_Int_Oddusd_Ctor) -> ((box true)) | (_, _) -> ((box false))))))))) Map.empty))))))

let Data_Int_ordParity  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (Data_Int_Evenusd_Ctor, Data_Int_Evenusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor)) | (Data_Int_Evenusd_Ctor, _) -> ((box Data_Ordering_LTusd_Ctor)) | (_, Data_Int_Evenusd_Ctor) -> ((box Data_Ordering_GTusd_Ctor)) | (Data_Int_Oddusd_Ctor, Data_Int_Oddusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Int_eqParity))))) Map.empty)))))))

let Data_Int_semiringParity  = (sharpurs_apply (box ((box Data_Semiring_Semiringusd_Dict))) (box ((box ((Map.add "zero" (box ((box Data_Int_Evenusd_Ctor))) (Map.add "add" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match ((unbox ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Int_eqParity)))))) (box ((box x)))))) (box ((box y))))))) with | LitBool true () -> ((box Data_Int_Evenusd_Ctor)) | _ -> ((box Data_Int_Oddusd_Ctor))))))))) (Map.add "one" (box ((box Data_Int_Oddusd_Ctor))) (Map.add "mul" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (Data_Int_Oddusd_Ctor, Data_Int_Oddusd_Ctor) -> ((box Data_Int_Oddusd_Ctor)) | (_, _) -> ((box Data_Int_Evenusd_Ctor))))))))) Map.empty)))))))))

let Data_Int_ringParity  = (sharpurs_apply (box ((box Data_Ring_Ringusd_Dict))) (box ((box ((Map.add "sub" (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Int_semiringParity)))))) (Map.add "Semiring0" (box ((box (fun (usd__unused: obj) -> (box Data_Int_semiringParity))))) Map.empty)))))))

let Data_Int_divisionRingParity  = (sharpurs_apply (box ((box Data_DivisionRing_DivisionRingusd_Dict))) (box ((box ((Map.add "recip" (box ((sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn)))))) (Map.add "Ring0" (box ((box (fun (usd__unused: obj) -> (box Data_Int_ringParity))))) Map.empty)))))))

let Data_Int_decimal  = (sharpurs_apply (box ((box Data_Int_Radix))) (box ((box 10))))

let Data_Int_commutativeRingParity  = (sharpurs_apply (box ((box Data_CommutativeRing_CommutativeRingusd_Dict))) (box ((box ((Map.add "Ring0" (box ((box (fun (usd__unused: obj) -> (box Data_Int_ringParity))))) Map.empty))))))

let Data_Int_euclideanRingParity  = (sharpurs_apply (box ((box Data_EuclideanRing_EuclideanRingusd_Dict))) (box ((box ((Map.add "degree" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Int_Evenusd_Ctor -> ((box 0)) | Data_Int_Oddusd_Ctor -> ((box 1))))))) (Map.add "div" (box ((box (fun (x: obj) -> (box (fun (v: obj) -> (box x))))))) (Map.add "mod" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Int_Evenusd_Ctor))))))) (Map.add "CommutativeRing0" (box ((box (fun (usd__unused: obj) -> (box Data_Int_commutativeRingParity))))) Map.empty)))))))))

let Data_Int_ceil  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Int_unsafeClamp)))))) (box ((box Data_Number_ceil))))

let Data_Int_boundedParity  = (sharpurs_apply (box ((box Data_Bounded_Boundedusd_Dict))) (box ((box ((Map.add "bottom" (box ((box Data_Int_Evenusd_Ctor))) (Map.add "top" (box ((box Data_Int_Oddusd_Ctor))) (Map.add "Ord0" (box ((box (fun (usd__unused: obj) -> (box Data_Int_ordParity))))) Map.empty))))))))

let Data_Int_binary  = (sharpurs_apply (box ((box Data_Int_Radix))) (box ((box 2))))

let Data_Int_base36  = (sharpurs_apply (box ((box Data_Int_Radix))) (box ((box 36))))
