[<AutoOpen>]
module PureScript_Data_Array_ST

open System
open System.Collections.Generic

module Data_Array_ST_FFI =
    open System
    open System.Collections.Generic
    open System.Linq
    
    let ``new`` =
        (fun () -> System.Collections.Generic.List<obj>() :> obj) :> obj
    
    let peekImpl =
        fun (just: obj) -> fun (nothing: obj) -> fun (iVal: obj) -> fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                let i = iVal :?> int
                if i >= 0 && i < arr.Count then
                    sharpurs_apply just arr.[i]
                else
                    nothing) :> obj
    
    let pokeImpl =
        fun (iVal: obj) -> fun (a: obj) -> fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                let i = iVal :?> int
                let ret = i >= 0 && i < arr.Count
                if ret then arr.[i] <- a
                box ret) :> obj
    
    let lengthImpl =
        fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                box arr.Count) :> obj
    
    let popImpl =
        fun (just: obj) -> fun (nothing: obj) -> fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                if arr.Count > 0 then
                    let el = arr.[arr.Count - 1]
                    arr.RemoveAt(arr.Count - 1)
                    sharpurs_apply just el
                else
                    nothing) :> obj
    
    let pushAllImpl =
        fun (asVal: obj) -> fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                let as' = asVal :?> obj[]
                arr.AddRange(as')
                box arr.Count) :> obj
    
    let shiftImpl =
        fun (just: obj) -> fun (nothing: obj) -> fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                if arr.Count > 0 then
                    let el = arr.[0]
                    arr.RemoveAt(0)
                    sharpurs_apply just el
                else
                    nothing) :> obj
    
    let unshiftAllImpl =
        fun (asVal: obj) -> fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                let as' = asVal :?> obj[]
                arr.InsertRange(0, as')
                box arr.Count) :> obj
    
    let spliceImpl =
        fun (iVal: obj) -> fun (howManyVal: obj) -> fun (bsVal: obj) -> fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                let i = iVal :?> int
                let howMany = howManyVal :?> int
                let bs = bsVal :?> obj[]
                let removed = arr.GetRange(i, howMany)
                arr.RemoveRange(i, howMany)
                arr.InsertRange(i, bs)
                removed.ToArray() :> obj) :> obj
    
    let unsafeFreezeImpl =
        fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                arr.ToArray() :> obj) :> obj
    
    let unsafeThawImpl =
        fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> obj[]
                System.Collections.Generic.List<obj>(arr) :> obj) :> obj
    
    let freezeImpl =
        fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                arr.ToArray() :> obj) :> obj
    
    let thawImpl =
        fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> obj[]
                System.Collections.Generic.List<obj>(arr) :> obj) :> obj
    
    let cloneImpl =
        fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                System.Collections.Generic.List<obj>(arr) :> obj) :> obj
    
    let sortByImpl =
        fun (compare: obj) -> fun (fromOrdering: obj) -> fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                if arr.Count < 2 then arr :> obj
                else
                    let outArr = arr.ToArray()
                    let comparer =
                        { new IComparer<obj> with
                            member _.Compare(a: obj, b: obj) =
                                let step1 = sharpurs_apply compare a
                                let ord = sharpurs_apply step1 b
                                unbox<int> (sharpurs_apply fromOrdering ord) }
                    let sorted = outArr.OrderBy((fun x -> x), comparer).ToArray()
                    arr.Clear()
                    arr.AddRange(sorted)
                    arr :> obj) :> obj
    
    let toAssocArrayImpl =
        fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                let result = Array.zeroCreate arr.Count
                for i = 0 to arr.Count - 1 do
                    let map = Map.empty |> Map.add "value" arr.[i] |> Map.add "index" (box i)
                    result.[i] <- box map
                result :> obj) :> obj
    
    let pushImpl =
        fun (a: obj) -> fun (xs: obj) ->
            (fun () ->
                let arr = xs :?> System.Collections.Generic.List<obj>
                arr.Add(a)
                box arr.Count) :> obj
    

let Data_Array_ST_cloneImpl = box (Data_Array_ST_FFI.``cloneImpl``)
let Data_Array_ST_freezeImpl = box (Data_Array_ST_FFI.``freezeImpl``)
let Data_Array_ST_lengthImpl = box (Data_Array_ST_FFI.``lengthImpl``)
let Data_Array_ST_new = box (Data_Array_ST_FFI.``new``)
let Data_Array_ST_peekImpl = box (Data_Array_ST_FFI.``peekImpl``)
let Data_Array_ST_pokeImpl = box (Data_Array_ST_FFI.``pokeImpl``)
let Data_Array_ST_popImpl = box (Data_Array_ST_FFI.``popImpl``)
let Data_Array_ST_pushAllImpl = box (Data_Array_ST_FFI.``pushAllImpl``)
let Data_Array_ST_pushImpl = box (Data_Array_ST_FFI.``pushImpl``)
let Data_Array_ST_shiftImpl = box (Data_Array_ST_FFI.``shiftImpl``)
let Data_Array_ST_sortByImpl = box (Data_Array_ST_FFI.``sortByImpl``)
let Data_Array_ST_spliceImpl = box (Data_Array_ST_FFI.``spliceImpl``)
let Data_Array_ST_thawImpl = box (Data_Array_ST_FFI.``thawImpl``)
let Data_Array_ST_toAssocArrayImpl = box (Data_Array_ST_FFI.``toAssocArrayImpl``)
let Data_Array_ST_unsafeFreezeImpl = box (Data_Array_ST_FFI.``unsafeFreezeImpl``)
let Data_Array_ST_unsafeThawImpl = box (Data_Array_ST_FFI.``unsafeThawImpl``)
let Data_Array_ST_unshiftAllImpl = box (Data_Array_ST_FFI.``unshiftAllImpl``)


let Data_Array_ST_unshiftAll  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn2))) (box ((box Data_Array_ST_unshiftAllImpl))))

let Data_Array_ST_unshift  = (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn2))) (box ((box Data_Array_ST_unshiftAllImpl)))))) (box ((box [|(box a)|]))))))

let Data_Array_ST_unsafeThaw  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn1))) (box ((box Data_Array_ST_unsafeThawImpl))))

let Data_Array_ST_unsafeFreeze  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn1))) (box ((box Data_Array_ST_unsafeFreezeImpl))))

let Data_Array_ST_toAssocArray  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn1))) (box ((box Data_Array_ST_toAssocArrayImpl))))

let Data_Array_ST_thaw  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn1))) (box ((box Data_Array_ST_thawImpl))))

let Data_Array_ST_withArray  = (box (fun (f: obj) -> (box (fun (xs: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Control_Monad_ST_Internal_bindST)))))) (box ((sharpurs_apply (box ((box Data_Array_ST_thaw))) (box ((box xs))))))))) (box ((box (fun (result: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Control_Monad_ST_Internal_bindST)))))) (box ((sharpurs_apply (box ((box f))) (box ((box result))))))))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box Data_Array_ST_unsafeFreeze))) (box ((box result))))))))))))))))))

let Data_Array_ST_splice  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn4))) (box ((box Data_Array_ST_spliceImpl))))

let Data_Array_ST_sortBy  = (box (fun (comp: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn3))) (box ((box Data_Array_ST_sortByImpl)))))) (box ((box comp)))))) (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Ordering_GTusd_Ctor -> ((box 1)) | Data_Ordering_EQusd_Ctor -> ((box 0)) | Data_Ordering_LTusd_Ctor -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_negate))) (box ((box Data_Ring_ringInt)))))) (box ((box 1)))))))))))))

let Data_Array_ST_sortWith  = (box (fun (dictOrd: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((box Data_Array_ST_sortBy))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_comparing))) (box ((box dictOrd)))))) (box ((box f)))))))))))

let Data_Array_ST_sort  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Array_ST_sortBy))) (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))))))

let Data_Array_ST_shift  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn3))) (box ((box Data_Array_ST_shiftImpl)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1))))))))))) (box ((box Data_Maybe_Nothingusd_Ctor))))

let Data_Array_ST_run  = (box (fun (st: obj) -> (sharpurs_apply (box ((box Control_Monad_ST_Internal_run))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Control_Monad_ST_Internal_bindST)))))) (box ((box st)))))) (box ((box Data_Array_ST_unsafeFreeze)))))))))

let Data_Array_ST_pushAll  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn2))) (box ((box Data_Array_ST_pushAllImpl))))

let Data_Array_ST_push  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn2))) (box ((box Data_Array_ST_pushImpl))))

let Data_Array_ST_pop  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn3))) (box ((box Data_Array_ST_popImpl)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1))))))))))) (box ((box Data_Maybe_Nothingusd_Ctor))))

let Data_Array_ST_poke  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn3))) (box ((box Data_Array_ST_pokeImpl))))

let Data_Array_ST_peek  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn4))) (box ((box Data_Array_ST_peekImpl)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1))))))))))) (box ((box Data_Maybe_Nothingusd_Ctor))))

let Data_Array_ST_modify  = (box (fun (i: obj) -> (box (fun (f: obj) -> (box (fun (xs: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Control_Monad_ST_Internal_bindST)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Array_ST_peek))) (box ((box i)))))) (box ((box xs))))))))) (box ((box (fun (entry: obj) -> (match ((unbox ((box entry)))) with | Data_Maybe_Justusd_Ctor(x) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Array_ST_poke))) (box ((box i)))))) (box ((sharpurs_apply (box ((box f))) (box ((box x))))))))) (box ((box xs))))) | Data_Maybe_Nothingusd_Ctor -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Control_Monad_ST_Internal_applicativeST)))))) (box ((box false)))))))))))))))))

let Data_Array_ST_length  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn1))) (box ((box Data_Array_ST_lengthImpl))))

let Data_Array_ST_freeze  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn1))) (box ((box Data_Array_ST_freezeImpl))))

let Data_Array_ST_clone  = (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn1))) (box ((box Data_Array_ST_cloneImpl))))
