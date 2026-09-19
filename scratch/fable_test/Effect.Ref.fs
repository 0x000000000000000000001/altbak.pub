[<AutoOpen>]
module PureScript_Effect_Ref

open System
open System.Collections.Generic

module Effect_Ref_FFI =
    let _new s = fun _ -> box (ref s)
    
    let newWithSelf f =
        fun _ ->
            let r = ref (box null)
            let f' = unbox<obj -> obj> f
            let s = unbox<obj> (f' (box r))
            r := s
            box r
    
    let read r =
        fun _ ->
            let rRef = unbox<obj ref> r
            lock rRef (fun () -> !rRef)
    
    let write s r =
        fun _ ->
            let rRef = unbox<obj ref> r
            lock rRef (fun () -> rRef := s)
            box null
    
    let modifyImpl f r =
        fun _ ->
            let rRef = unbox<obj ref> r
            let f' = unbox<obj -> Map<string, obj>> f
            lock rRef (fun () ->
                let result = f' (!rRef)
                rRef := Map.find "state" result
                Map.find "value" result
            )
    

let Effect_Ref__new = box (fun (arg0: obj) -> box (Effect_Ref_FFI.``_new`` (unbox arg0)))
let Effect_Ref_modifyImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Effect_Ref_FFI.``modifyImpl`` (unbox arg0) (unbox arg1))))
let Effect_Ref_newWithSelf = box (fun (arg0: obj) -> box (Effect_Ref_FFI.``newWithSelf`` (unbox arg0)))
let Effect_Ref_read = box (fun (arg0: obj) -> box (Effect_Ref_FFI.``read`` (unbox arg0)))
let Effect_Ref_write = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Effect_Ref_FFI.``write`` (unbox arg0) (unbox arg1))))


let Effect_Ref_void  = (sharpurs_apply (box ((box Data_Functor_void))) (box ((box Effect_functorEffect))))

let Effect_Ref_new  = (box Effect_Ref__new)

let Effect_Ref_modify_prime  = (box Effect_Ref_modifyImpl)

let Effect_Ref_modify  = (box (fun (f: obj) -> (sharpurs_apply (box ((box Effect_Ref_modify_prime))) (box ((box (fun (s: obj) -> (let s_prime = (sharpurs_apply (box ((box f))) (box ((box s)))) in (box ((Map.add "state" (box ((box s_prime))) (Map.add "value" (box ((box s_prime))) Map.empty))))))))))))

let Effect_Ref_modify_  = (box (fun (f: obj) -> (box (fun (s: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Effect_Ref_void)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Ref_modify))) (box ((box f)))))) (box ((box s)))))))))))
