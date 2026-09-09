[<AutoOpen>]
module PureScript_Data_Array_NonEmpty_Internal

open System
open System.Collections.Generic

module Data_Array_NonEmpty_Internal_FFI =
    let foldr1Impl =
        fun (f: obj) -> fun (xs: obj) ->
            let arr = xs :?> obj[]
            let mutable acc = arr.[arr.Length - 1]
            for i = arr.Length - 2 downto 0 do
                let step1 = sharpurs_apply f arr.[i]
                acc <- sharpurs_apply step1 acc
            acc
    
    let foldl1Impl =
        fun (f: obj) -> fun (xs: obj) ->
            let arr = xs :?> obj[]
            let mutable acc = arr.[0]
            for i = 1 to arr.Length - 1 do
                let step1 = sharpurs_apply f acc
                acc <- sharpurs_apply step1 arr.[i]
            acc
    
    let traverse1Impl =
        fun (apply: obj) -> fun (mapFn: obj) -> fun (f: obj) -> fun (arrayVal: obj) ->
            let arr = arrayVal :?> obj[]
    
            let array1 (a: obj) = [| a |] :> obj
            let array2 (a: obj) = 
                (fun (b: obj) -> [| a; b |] :> obj) :> obj
            let array3 (a: obj) = 
                (fun (b: obj) -> 
                    (fun (c: obj) -> [| a; b; c |] :> obj) :> obj) :> obj
            
            let concat2 (xsVal: obj) = 
                (fun (ysVal: obj) ->
                    let xs = xsVal :?> obj[]
                    let ys = ysVal :?> obj[]
                    Array.append xs ys :> obj) :> obj
    
            let rec go bot top =
                let diff = top - bot
                if diff = 1 then
                    let mapped = sharpurs_apply mapFn (array1 :> obj)
                    sharpurs_apply mapped (sharpurs_apply f arr.[bot])
                elif diff = 2 then
                    let mapped = sharpurs_apply mapFn (array2 :> obj)
                    let applied = sharpurs_apply apply (sharpurs_apply mapped (sharpurs_apply f arr.[bot]))
                    sharpurs_apply applied (sharpurs_apply f arr.[bot + 1])
                elif diff = 3 then
                    let mapped = sharpurs_apply mapFn (array3 :> obj)
                    let applied1 = sharpurs_apply apply (sharpurs_apply mapped (sharpurs_apply f arr.[bot]))
                    let applied2 = sharpurs_apply apply (sharpurs_apply applied1 (sharpurs_apply f arr.[bot + 1]))
                    sharpurs_apply applied2 (sharpurs_apply f arr.[bot + 2])
                else
                    let pivot = bot + (diff / 4) * 2
                    let mapped = sharpurs_apply mapFn (concat2 :> obj)
                    let applied = sharpurs_apply apply (sharpurs_apply mapped (go bot pivot))
                    sharpurs_apply applied (go pivot top)
            
            go 0 arr.Length
    

let Data_Array_NonEmpty_Internal_foldl1Impl = box (Data_Array_NonEmpty_Internal_FFI.``foldl1Impl``)
let Data_Array_NonEmpty_Internal_foldr1Impl = box (Data_Array_NonEmpty_Internal_FFI.``foldr1Impl``)
let Data_Array_NonEmpty_Internal_traverse1Impl = box (Data_Array_NonEmpty_Internal_FFI.``traverse1Impl``)


let Data_Array_NonEmpty_Internal_NonEmptyArray  = (box (fun (x: obj) -> (box x)))

let Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray  = (box Data_Unfoldable1_unfoldable1Array)

let Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray  = (box Data_TraversableWithIndex_traversableWithIndexArray)

let Data_Array_NonEmpty_Internal_traversableNonEmptyArray  = (box Data_Traversable_traversableArray)

let Data_Array_NonEmpty_Internal_showNonEmptyArray  = (box (fun (dictShow: obj) -> (let showArray = (sharpurs_apply (box ((box Data_Show_showArray))) (box ((box dictShow)))) in (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | xs -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(NonEmptyArray ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box showArray)))))) (box ((box xs))))))))) (box ((box ")"))))))))))))) Map.empty)))))))))

let Data_Array_NonEmpty_Internal_semigroupNonEmptyArray  = (box Data_Semigroup_semigroupArray)

let Data_Array_NonEmpty_Internal_ordNonEmptyArray  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_ordArray))) (box ((box dictOrd))))))

let Data_Array_NonEmpty_Internal_ord1NonEmptyArray  = (box Data_Ord_ord1Array)

let Data_Array_NonEmpty_Internal_monadNonEmptyArray  = (box Control_Monad_monadArray)

let Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray  = (box Data_FunctorWithIndex_functorWithIndexArray)

let Data_Array_NonEmpty_Internal_functorNonEmptyArray  = (box Data_Functor_functorArray)

let Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray  = (box Data_FoldableWithIndex_foldableWithIndexArray)

let Data_Array_NonEmpty_Internal_foldableNonEmptyArray  = (box Data_Foldable_foldableArray)

let rec Data_Array_NonEmpty_Internal_foldable1NonEmptyArray : obj = ((sharpurs_apply (box ((box Data_Semigroup_Foldable_Foldable1usd_Dict))) (box ((box ((Map.add "foldMap1" (box ((box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_Foldable_foldMap1DefaultL))) (box ((box Data_Array_NonEmpty_Internal_foldable1NonEmptyArray)))))) (box ((box Data_Array_NonEmpty_Internal_functorNonEmptyArray)))))) (box ((box dictSemigroup)))))))) (Map.add "foldr1" (box ((sharpurs_apply (box ((box Data_Function_Uncurried_runFn2))) (box ((box Data_Array_NonEmpty_Internal_foldr1Impl)))))) (Map.add "foldl1" (box ((sharpurs_apply (box ((box Data_Function_Uncurried_runFn2))) (box ((box Data_Array_NonEmpty_Internal_foldl1Impl)))))) (Map.add "Foldable0" (box ((box (fun (usd__unused: obj) -> (box Data_Array_NonEmpty_Internal_foldableNonEmptyArray))))) Map.empty))))))))))


let rec Data_Array_NonEmpty_Internal_traversable1NonEmptyArray : obj = ((sharpurs_apply (box ((box Data_Semigroup_Traversable_Traversable1usd_Dict))) (box ((box ((Map.add "traverse1" (box ((box (fun (dictApply: obj) -> (let apply = (sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))) in let map = (sharpurs_apply (box ((box Data_Functor_map))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictApply)))))) (box ((box Prim_undefined))))))) in (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_Uncurried_runFn3))) (box ((box Data_Array_NonEmpty_Internal_traverse1Impl)))))) (box ((box apply)))))) (box ((box map)))))) (box ((box f))))))))))) (Map.add "sequence1" (box ((box (fun (dictApply: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_Traversable_sequence1Default))) (box ((box Data_Array_NonEmpty_Internal_traversable1NonEmptyArray)))))) (box ((box dictApply)))))))) (Map.add "Foldable10" (box ((box (fun (usd__unused: obj) -> (box Data_Array_NonEmpty_Internal_foldable1NonEmptyArray))))) (Map.add "Traversable1" (box ((box (fun (usd__unused: obj) -> (box Data_Array_NonEmpty_Internal_traversableNonEmptyArray))))) Map.empty))))))))))


let Data_Array_NonEmpty_Internal_eqNonEmptyArray  = (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_eqArray))) (box ((box dictEq))))))

let Data_Array_NonEmpty_Internal_eq1NonEmptyArray  = (box Data_Eq_eq1Array)

let Data_Array_NonEmpty_Internal_bindNonEmptyArray  = (box Control_Bind_bindArray)

let Data_Array_NonEmpty_Internal_applyNonEmptyArray  = (box Control_Apply_applyArray)

let Data_Array_NonEmpty_Internal_applicativeNonEmptyArray  = (box Control_Applicative_applicativeArray)

let Data_Array_NonEmpty_Internal_altNonEmptyArray  = (box Control_Alt_altArray)
