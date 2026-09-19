[<AutoOpen>]
module PureScript_Data_Array_ST_Partial

open System
open System.Collections.Generic

module Data_Array_ST_Partial_FFI =
    let peekImpl = 
        fun (iVal: obj) -> fun (xs: obj) ->
            let i = iVal :?> int
            let arr = xs :?> System.Collections.Generic.List<obj>
            arr.[i]
    
    let pokeImpl = 
        fun (iVal: obj) -> fun (a: obj) -> fun (xs: obj) ->
            let i = iVal :?> int
            let arr = xs :?> System.Collections.Generic.List<obj>
            arr.[i] <- a
            null :> obj
    

let Data_Array_ST_Partial_peekImpl = box (Data_Array_ST_Partial_FFI.``peekImpl``)
let Data_Array_ST_Partial_pokeImpl = box (Data_Array_ST_Partial_FFI.``pokeImpl``)


let Data_Array_ST_Partial_poke  = (box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn3))) (box ((box Data_Array_ST_Partial_pokeImpl))))))

let Data_Array_ST_Partial_peek  = (box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box Control_Monad_ST_Uncurried_runSTFn2))) (box ((box Data_Array_ST_Partial_peekImpl))))))
