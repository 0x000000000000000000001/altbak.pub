[<AutoOpen>]
module PureScript_Data_Semigroup

open System
open System.Collections.Generic

module Data_Semigroup_FFI =
    let concatString a b = (unbox<string> a) + (unbox<string> b)
    let concatArray (xs: obj) (ys: obj) =
        let arrX = unbox<obj[]> xs
        let arrY = unbox<obj[]> ys
        if arrX.Length = 0 then ys
        elif arrY.Length = 0 then xs
        else
            let res = Array.zeroCreate (arrX.Length + arrY.Length)
            Array.Copy(arrX, 0, res, 0, arrX.Length)
            Array.Copy(arrY, 0, res, arrX.Length, arrY.Length)
            box res
    

let Data_Semigroup_concatArray = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Semigroup_FFI.``concatArray`` (unbox arg0) (unbox arg1))))
let Data_Semigroup_concatString = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data_Semigroup_FFI.``concatString`` (unbox arg0) (unbox arg1))))


let Data_Semigroup_SemigroupRecordusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Semigroup_Semigroupusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Semigroup_semigroupVoid  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box Data_Void_absurd))))) Map.empty))))))

let Data_Semigroup_semigroupUnit  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Unit_unit))))))) Map.empty))))))

let Data_Semigroup_semigroupString  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box Data_Semigroup_concatString))) Map.empty))))))

let Data_Semigroup_semigroupRecordNil  = (sharpurs_apply (box ((box Data_Semigroup_SemigroupRecordusd_Dict))) (box ((box ((Map.add "appendRecord" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (Map.empty)))))))))) Map.empty))))))

let Data_Semigroup_semigroupProxy  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Type_Proxy_Proxyusd_Ctor))))))) Map.empty))))))

let Data_Semigroup_semigroupArray  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box Data_Semigroup_concatArray))) Map.empty))))))

let Data_Semigroup_appendRecord  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "appendRecord" (unbox<Map<string, obj>> ((box v))))))))

let Data_Semigroup_semigroupRecord  = (box (fun (usd__unused: obj) -> (box (fun (dictSemigroupRecord: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_appendRecord))) (box ((box dictSemigroupRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) Map.empty))))))))))

let Data_Semigroup_append  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "append" (unbox<Map<string, obj>> ((box v))))))))

let Data_Semigroup_semigroupFn  = (box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup)))))) (box ((sharpurs_apply (box ((box f))) (box ((box x))))))))) (box ((sharpurs_apply (box ((box g))) (box ((box x))))))))))))))) Map.empty))))))))

let Data_Semigroup_semigroupRecordCons  = (box (fun (dictIsSymbol: obj) -> (box (fun (usd__unused: obj) -> (box (fun (dictSemigroupRecord: obj) -> (box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((box Data_Semigroup_SemigroupRecordusd_Dict))) (box ((box ((Map.add "appendRecord" (box ((box (fun (v: obj) -> (box (fun (ra: obj) -> (box (fun (rb: obj) -> (let tail = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_appendRecord))) (box ((box dictSemigroupRecord)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))))) (box ((box ra)))))) (box ((box rb)))) in let key = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Symbol_reflectSymbol))) (box ((box dictIsSymbol)))))) (box ((box Type_Proxy_Proxyusd_Ctor)))) in let insert = (sharpurs_apply (box ((box Record_Unsafe_unsafeSet))) (box ((box key)))) in let get = (sharpurs_apply (box ((box Record_Unsafe_unsafeGet))) (box ((box key)))) in (sharpurs_apply (box ((sharpurs_apply (box ((box insert))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup)))))) (box ((sharpurs_apply (box ((box get))) (box ((box ra))))))))) (box ((sharpurs_apply (box ((box get))) (box ((box rb)))))))))))) (box ((box tail))))))))))))) Map.empty))))))))))))))
