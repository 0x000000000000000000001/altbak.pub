[<AutoOpen>]
module PureScript_Data_Unfoldable1

open System
open System.Collections.Generic

module Data_Unfoldable1_FFI =
    let unfoldr1ArrayImpl = 
        fun (isNothing: obj) -> fun (fromJust: obj) -> fun (fst: obj) -> fun (snd: obj) -> fun (f: obj) -> fun (b: obj) ->
            let isNothing' = isNothing :?> (obj -> obj)
            let fromJust' = fromJust :?> (obj -> obj)
            let fst' = fst :?> (obj -> obj)
            let snd' = snd :?> (obj -> obj)
            let f' = f :?> (obj -> obj)
            
            let result = System.Collections.Generic.List<obj>()
            let mutable value = b
            let mutable looping = true
            
            while looping do
                let tuple = f' value
                result.Add(fst' tuple)
                let maybe = snd' tuple
                if (isNothing' maybe :?> bool) then
                    looping <- false
                else
                    value <- fromJust' maybe
                    
            result.ToArray() :> obj
    

let Data_Unfoldable1_unfoldr1ArrayImpl = box (Data_Unfoldable1_FFI.``unfoldr1ArrayImpl``)


let Data_Unfoldable1_Unfoldable1usd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Unfoldable1_unfoldr1  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "unfoldr1" (unbox<Map<string, obj>> ((box v))))))))

let Data_Unfoldable1_unfoldable1Maybe  = (sharpurs_apply (box ((box Data_Unfoldable1_Unfoldable1usd_Dict))) (box ((box ((Map.add "unfoldr1" (box ((box (fun (f: obj) -> (box (fun (b: obj) -> (sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1)))))))) (box ((sharpurs_apply (box ((box Data_Tuple_fst))) (box ((sharpurs_apply (box ((box f))) (box ((box b)))))))))))))))) Map.empty))))))

let Data_Unfoldable1_unfoldable1Array  = (sharpurs_apply (box ((box Data_Unfoldable1_Unfoldable1usd_Dict))) (box ((box ((Map.add "unfoldr1" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Unfoldable1_unfoldr1ArrayImpl))) (box ((box Data_Maybe_isNothing)))))) (box ((sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box Data_Maybe_fromJust))) (box ((box Prim_undefined)))))))))))))) (box ((box Data_Tuple_fst)))))) (box ((box Data_Tuple_snd)))))) Map.empty))))))

let Data_Unfoldable1_replicate1  = (box (fun (dictUnfoldable1: obj) -> (box (fun (n: obj) -> (box (fun (v: obj) -> (let step = (box (fun (i: obj) -> (match ((unbox ((box i)))) with | i1 when (unbox (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThanOrEq))) (box ((box Data_Ord_ordInt)))))) (box ((box i1)))))) (box ((box 0))))) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box v)))))) (box ((box Data_Maybe_Nothingusd_Ctor))))) | i1 when (unbox (box Data_Boolean_otherwise)) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box v)))))) (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1)))))))) (box ((box ((unbox<int> (box ((box i1)))) - (unbox<int> (box ((box 1)))))))))))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Unfoldable1_unfoldr1))) (box ((box dictUnfoldable1)))))) (box ((box step)))))) (box ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))))))))))))

let Data_Unfoldable1_replicate1A  = (box (fun (dictApply: obj) -> (box (fun (dictUnfoldable1: obj) -> (box (fun (dictTraversable1: obj) -> (box (fun (n: obj) -> (box (fun (m: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_Traversable_sequence1))) (box ((box dictTraversable1)))))) (box ((box dictApply)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Unfoldable1_replicate1))) (box ((box dictUnfoldable1)))))) (box ((box n)))))) (box ((box m)))))))))))))))))

let Data_Unfoldable1_singleton  = (box (fun (dictUnfoldable1: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Unfoldable1_replicate1))) (box ((box dictUnfoldable1)))))) (box ((box 1))))))

let Data_Unfoldable1_range  = (box (fun (dictUnfoldable1: obj) -> (box (fun (start: obj) -> (box (fun (end_var: obj) -> (let go = (box (fun (delta: obj) -> (box (fun (i: obj) -> (let i_prime = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((box i)))))) (box ((box delta)))) in (sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box i)))))) (box ((match ((unbox ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Eq_eqInt)))))) (box ((box i)))))) (box ((box end_var))))))) with | LitBool true () -> ((box Data_Maybe_Nothingusd_Ctor)) | _ -> ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1)))))))) (box ((box i_prime)))))))))))))) in (let delta = (match ((unbox ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_greaterThanOrEq))) (box ((box Data_Ord_ordInt)))))) (box ((box end_var)))))) (box ((box start))))))) with | LitBool true () -> ((box 1)) | _ -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_negate))) (box ((box Data_Ring_ringInt)))))) (box ((box 1)))))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Unfoldable1_unfoldr1))) (box ((box dictUnfoldable1)))))) (box ((sharpurs_apply (box ((box go))) (box ((box delta))))))))) (box ((box start))))))))))))

let Data_Unfoldable1_iterateN  = (box (fun (dictUnfoldable1: obj) -> (box (fun (n: obj) -> (box (fun (f: obj) -> (box (fun (s: obj) -> (let go = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Tuple_Tupleusd_Ctor(x, n_prime) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box x)))))) (box ((match ((unbox ((box ((unbox<int> (box ((box n_prime)))) > (unbox<int> (box ((box 0))))))))) with | LitBool true () -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((box f))) (box ((box x)))))))))))) (box ((box ((unbox<int> (box ((box n_prime)))) - (unbox<int> (box ((box 1))))))))))))) | _ -> ((box Data_Maybe_Nothingusd_Ctor)))))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Unfoldable1_unfoldr1))) (box ((box dictUnfoldable1)))))) (box ((box go))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box s)))))) (box ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1)))))))))))))))))))))
