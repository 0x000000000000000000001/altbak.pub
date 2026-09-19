[<AutoOpen>]
module PureScript_Data_EuclideanRing

open System
open System.Collections.Generic

module Data_EuclideanRing_FFI =
    let absInt (n: int) =
        if n = System.Int32.MinValue then System.Int32.MaxValue
        else System.Math.Abs(n)
    
    let intDiv a b = 
        let x = unbox<int> a
        let y = unbox<int> b
        if y = 0 then 0
        else
            let yy = absInt y
            let m = ((x % yy) + yy) % yy
            if y = -1 && x = System.Int32.MinValue then System.Int32.MinValue
            else ((x - m) / y)
            
    let intDegree a = absInt (unbox<int> a)
    let numDiv a b = (unbox<float> a) / (unbox<float> b)
    let intMod a b =
        let x = unbox<int> a
        let y = unbox<int> b
        if y = 0 then 0
        else
            let yy = absInt y
            ((x % yy) + yy) % yy
    

let Data_EuclideanRing_intDegree = box (fun (arg0: obj) -> box (Data_EuclideanRing_FFI.``intDegree`` (unbox arg0)))
let Data_EuclideanRing_intDiv = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_EuclideanRing_FFI.``intDiv`` (unbox arg0) (unbox arg1))))
let Data_EuclideanRing_intMod = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_EuclideanRing_FFI.``intMod`` (unbox arg0) (unbox arg1))))
let Data_EuclideanRing_numDiv = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_EuclideanRing_FFI.``numDiv`` (unbox arg0) (unbox arg1))))


let Data_EuclideanRing_EuclideanRingusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_EuclideanRing_mod  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "mod" (unbox<Map<string, obj>> ((box v))))))))

let rec Data_EuclideanRing_gcd_tco (dictEq: obj) (dictEuclideanRing: obj) : obj = ((let zero = (sharpurs_apply (box ((box Data_Semiring_zero))) (box ((sharpurs_apply (box ((Map.find "Semiring0" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "Ring0" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "CommutativeRing0" (unbox<Map<string, obj>> ((box dictEuclideanRing)))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined))))))) in (box (fun (a: obj) -> (box (fun (b: obj) -> (match ((unbox ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box dictEq)))))) (box ((box b)))))) (box ((box zero))))))) with | LitBool true () -> ((box a)) | _ -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_gcd))) (box ((box dictEq)))))) (box ((box dictEuclideanRing)))))) (box ((box b)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_mod))) (box ((box dictEuclideanRing)))))) (box ((box a)))))) (box ((box b)))))))))))))))
and Data_EuclideanRing_gcd = box (fun (dictEq: obj) ->  (fun (dictEuclideanRing: obj) -> Data_EuclideanRing_gcd_tco dictEq dictEuclideanRing))


let Data_EuclideanRing_euclideanRingNumber  = (sharpurs_apply (box ((box Data_EuclideanRing_EuclideanRingusd_Dict))) (box ((box ((Map.add "degree" (box ((box (fun (v: obj) -> (box 1))))) (Map.add "div" (box ((box Data_EuclideanRing_numDiv))) (Map.add "mod" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box 0.0))))))) (Map.add "CommutativeRing0" (box ((box (fun (usd__unused: obj) -> (box Data_CommutativeRing_commutativeRingNumber))))) Map.empty)))))))))

let Data_EuclideanRing_euclideanRingInt  = (sharpurs_apply (box ((box Data_EuclideanRing_EuclideanRingusd_Dict))) (box ((box ((Map.add "degree" (box ((box Data_EuclideanRing_intDegree))) (Map.add "div" (box ((box Data_EuclideanRing_intDiv))) (Map.add "mod" (box ((box Data_EuclideanRing_intMod))) (Map.add "CommutativeRing0" (box ((box (fun (usd__unused: obj) -> (box Data_CommutativeRing_commutativeRingInt))))) Map.empty)))))))))

let Data_EuclideanRing_div  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "div" (unbox<Map<string, obj>> ((box v))))))))

let Data_EuclideanRing_lcm  = (box (fun (dictEq: obj) -> (box (fun (dictEuclideanRing: obj) -> (let Ring0 = (sharpurs_apply (box ((Map.find "Ring0" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "CommutativeRing0" (unbox<Map<string, obj>> ((box dictEuclideanRing)))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined)))) in let zero = (sharpurs_apply (box ((box Data_Semiring_zero))) (box ((sharpurs_apply (box ((Map.find "Semiring0" (unbox<Map<string, obj>> ((box Ring0)))))) (box ((box Prim_undefined))))))) in let Semiring0 = (sharpurs_apply (box ((Map.find "Semiring0" (unbox<Map<string, obj>> ((box Ring0)))))) (box ((box Prim_undefined)))) in (box (fun (a: obj) -> (box (fun (b: obj) -> (match ((unbox ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_disj))) (box ((box Data_HeytingAlgebra_heytingAlgebraBoolean)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box dictEq)))))) (box ((box a)))))) (box ((box zero))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box dictEq)))))) (box ((box b)))))) (box ((box zero)))))))))) with | LitBool true () -> ((box zero)) | _ -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_div))) (box ((box dictEuclideanRing)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mul))) (box ((box Semiring0)))))) (box ((box a)))))) (box ((box b))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_gcd))) (box ((box dictEq)))))) (box ((box dictEuclideanRing)))))) (box ((box a)))))) (box ((box b))))))))))))))))))

let Data_EuclideanRing_degree  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "degree" (unbox<Map<string, obj>> ((box v))))))))
