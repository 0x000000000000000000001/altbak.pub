[<AutoOpen>]
module PureScript_Data_Function_Uncurried

open System
open System.Collections.Generic

module Data_Function_Uncurried_FFI =
    let mkFn0 = fun (fn: obj) ->
        let f = fn :?> (obj -> obj)
        (fun () -> f null) :> obj
        
    let mkFn1 = fun (fn: obj) -> fn
    let mkFn2 = fun (fn: obj) -> fn
    let mkFn3 = fun (fn: obj) -> fn
    let mkFn4 = fun (fn: obj) -> fn
    let mkFn5 = fun (fn: obj) -> fn
    let mkFn6 = fun (fn: obj) -> fn
    let mkFn7 = fun (fn: obj) -> fn
    let mkFn8 = fun (fn: obj) -> fn
    let mkFn9 = fun (fn: obj) -> fn
    let mkFn10 = fun (fn: obj) -> fn
    
    let runFn0 = fun (fn: obj) -> 
        let f = fn :?> (unit -> obj)
        f ()
    
    let runFn1 = fun (fn: obj) -> fn
    let runFn2 = fun (fn: obj) -> fn
    let runFn3 = fun (fn: obj) -> fn
    let runFn4 = fun (fn: obj) -> fn
    let runFn5 = fun (fn: obj) -> fn
    let runFn6 = fun (fn: obj) -> fn
    let runFn7 = fun (fn: obj) -> fn
    let runFn8 = fun (fn: obj) -> fn
    let runFn9 = fun (fn: obj) -> fn
    let runFn10 = fun (fn: obj) -> fn
    

let Data_Function_Uncurried_mkFn0 = box (Data_Function_Uncurried_FFI.``mkFn0``)
let Data_Function_Uncurried_mkFn10 = box (Data_Function_Uncurried_FFI.``mkFn10``)
let Data_Function_Uncurried_mkFn2 = box (Data_Function_Uncurried_FFI.``mkFn2``)
let Data_Function_Uncurried_mkFn3 = box (Data_Function_Uncurried_FFI.``mkFn3``)
let Data_Function_Uncurried_mkFn4 = box (Data_Function_Uncurried_FFI.``mkFn4``)
let Data_Function_Uncurried_mkFn5 = box (Data_Function_Uncurried_FFI.``mkFn5``)
let Data_Function_Uncurried_mkFn6 = box (Data_Function_Uncurried_FFI.``mkFn6``)
let Data_Function_Uncurried_mkFn7 = box (Data_Function_Uncurried_FFI.``mkFn7``)
let Data_Function_Uncurried_mkFn8 = box (Data_Function_Uncurried_FFI.``mkFn8``)
let Data_Function_Uncurried_mkFn9 = box (Data_Function_Uncurried_FFI.``mkFn9``)
let Data_Function_Uncurried_runFn0 = box (Data_Function_Uncurried_FFI.``runFn0``)
let Data_Function_Uncurried_runFn10 = box (Data_Function_Uncurried_FFI.``runFn10``)
let Data_Function_Uncurried_runFn2 = box (Data_Function_Uncurried_FFI.``runFn2``)
let Data_Function_Uncurried_runFn3 = box (Data_Function_Uncurried_FFI.``runFn3``)
let Data_Function_Uncurried_runFn4 = box (Data_Function_Uncurried_FFI.``runFn4``)
let Data_Function_Uncurried_runFn5 = box (Data_Function_Uncurried_FFI.``runFn5``)
let Data_Function_Uncurried_runFn6 = box (Data_Function_Uncurried_FFI.``runFn6``)
let Data_Function_Uncurried_runFn7 = box (Data_Function_Uncurried_FFI.``runFn7``)
let Data_Function_Uncurried_runFn8 = box (Data_Function_Uncurried_FFI.``runFn8``)
let Data_Function_Uncurried_runFn9 = box (Data_Function_Uncurried_FFI.``runFn9``)


let Data_Function_Uncurried_runFn1  = (box (fun (f: obj) -> (box f)))

let Data_Function_Uncurried_mkFn1  = (box (fun (f: obj) -> (box f)))
