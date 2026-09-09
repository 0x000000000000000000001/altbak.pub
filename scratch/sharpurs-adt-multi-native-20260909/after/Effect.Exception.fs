[<AutoOpen>]
module PureScript_Effect_Exception

open System
open System.Collections.Generic

module Effect_Exception_FFI =
    let showErrorImpl = 
        fun (errVal: obj) ->
            let err = errVal :?> System.Exception
            box (err.ToString())
    
    let error = 
        fun (msgVal: obj) ->
            let msg = msgVal :?> string
            box (new System.Exception(msg))
    
    let errorWithCause = 
        fun (msgVal: obj) -> fun (causeVal: obj) ->
            let msg = msgVal :?> string
            let cause = causeVal :?> System.Exception
            box (new System.Exception(msg, cause))
    
    let errorWithName = 
        fun (nameVal: obj) -> fun (msgVal: obj) ->
            let msg = msgVal :?> string
            let name = nameVal :?> string
            let ex = new System.Exception(msg)
            ex.Data.Add("Name", name)
            box ex
    
    let message = 
        fun (eVal: obj) ->
            let e = eVal :?> System.Exception
            box (e.Message)
    
    let name = 
        fun (eVal: obj) ->
            let e = eVal :?> System.Exception
            if e.Data.Contains("Name") then
                box (e.Data.["Name"])
            else
                box "Error"
    
    let stackImpl = 
        fun (just: obj) -> fun (nothing: obj) -> fun (eVal: obj) ->
            let e = eVal :?> System.Exception
            if not (System.String.IsNullOrEmpty(e.StackTrace)) then
                let justFn = just :?> (obj -> obj)
                justFn (box (e.StackTrace))
            else
                nothing
    
    let throwException = 
        fun (eVal: obj) ->
            fun (dummy: obj) ->
                let e = eVal :?> System.Exception
                raise e
                null :> obj
    
    let catchException = 
        fun (cVal: obj) -> fun (tVal: obj) ->
            fun (dummy: obj) ->
                let t = tVal :?> (obj -> obj)
                let c = cVal :?> (obj -> obj)
                try
                    t null
                with
                | :? System.Exception as ex ->
                    let c1 = c (box ex) :?> (obj -> obj)
                    c1 null
                | ex ->
                    let c1 = c (box (new System.Exception(ex.ToString()))) :?> (obj -> obj)
                    c1 null
    

let Effect_Exception_catchException = box (Effect_Exception_FFI.``catchException``)
let Effect_Exception_error = box (Effect_Exception_FFI.``error``)
let Effect_Exception_errorWithCause = box (Effect_Exception_FFI.``errorWithCause``)
let Effect_Exception_errorWithName = box (Effect_Exception_FFI.``errorWithName``)
let Effect_Exception_message = box (Effect_Exception_FFI.``message``)
let Effect_Exception_name = box (Effect_Exception_FFI.``name``)
let Effect_Exception_showErrorImpl = box (Effect_Exception_FFI.``showErrorImpl``)
let Effect_Exception_stackImpl = box (Effect_Exception_FFI.``stackImpl``)
let Effect_Exception_throwException = box (Effect_Exception_FFI.``throwException``)


let Effect_Exception_pure  = (sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect))))

let Effect_Exception_try  = (box (fun (action: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Exception_catchException))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Effect_Exception_pure)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Leftusd_Ctor(usd__arg1)))))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Effect_functorEffect)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Rightusd_Ctor(usd__arg1))))))))))) (box ((box action)))))))))

let Effect_Exception_throw  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Effect_Exception_throwException)))))) (box ((box Effect_Exception_error))))

let Effect_Exception_stack  = (sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Exception_stackImpl))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1))))))))))) (box ((box Data_Maybe_Nothingusd_Ctor))))

let Effect_Exception_showError  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box Effect_Exception_showErrorImpl))) Map.empty))))))
