[<AutoOpen>]
module PureScript_Control_Monad_ST_Internal

open System
open System.Collections.Generic

module Control_Monad_ST_Internal_FFI =
    let map_ = 
        fun (f: obj) -> fun (a: obj) ->
            let a' = a :?> (unit -> obj)
            (fun () -> Sharpurs_Prelude.sharpurs_apply f (a' ())) :> obj
    
    let bind_ = 
        fun (a: obj) -> fun (f: obj) ->
            let a' = a :?> (unit -> obj)
            (fun () -> 
                let res = Sharpurs_Prelude.sharpurs_apply f (a' ()) :?> (unit -> obj)
                res ()) :> obj
    
    let pure_ = 
        fun (a: obj) -> 
            (fun () -> a) :> obj
    
    let run = 
        fun (f: obj) -> 
            let f' = f :?> (unit -> obj)
            f' ()
    
    let ``while`` = 
        fun (cond: obj) -> fun (a: obj) ->
            let cond' = cond :?> (unit -> obj)
            let a' = a :?> (unit -> obj)
            (fun () ->
                while (unbox<bool> (cond' ())) do
                    a' () |> ignore
                null :> obj) :> obj
    
    let ``new`` = 
        fun (val': obj) -> 
            (fun () -> ref val' :> obj) :> obj
    
    let read = 
        fun (r: obj) -> 
            let r' = r :?> obj ref
            (fun () -> r'.Value) :> obj
    
    let modifyImpl = 
        fun (f: obj) -> fun (r: obj) ->
            let r' = r :?> obj ref
            (fun () ->
                let res = Sharpurs_Prelude.sharpurs_apply f (r'.Value) :?> Map<string, obj>
                r'.Value <- Map.find "state" res
                Map.find "value" res) :> obj
    
    let write = 
        fun (a: obj) -> fun (r: obj) ->
            let r' = r :?> obj ref
            (fun () ->
                r'.Value <- a
                a) :> obj
    
    let ``for`` =
        fun (lo: obj) -> fun (hi: obj) -> fun (f: obj) ->
            let lo' = lo :?> int
            let hi' = hi :?> int
            (fun () ->
                for i = lo' to hi' - 1 do
                    let step = Sharpurs_Prelude.sharpurs_apply f (box i) :?> (unit -> obj)
                    step () |> ignore
                null :> obj) :> obj
    
    let foreach =
        fun (xs: obj) -> fun (f: obj) ->
            let arr = xs :?> obj[]
            (fun () ->
                for x in arr do
                    let step = Sharpurs_Prelude.sharpurs_apply f x :?> (unit -> obj)
                    step () |> ignore
                null :> obj) :> obj
    

let Control_Monad_ST_Internal_bind_ = box (Control_Monad_ST_Internal_FFI.``bind_``)
let Control_Monad_ST_Internal_for = box (Control_Monad_ST_Internal_FFI.``for``)
let Control_Monad_ST_Internal_foreach = box (Control_Monad_ST_Internal_FFI.``foreach``)
let Control_Monad_ST_Internal_map_ = box (Control_Monad_ST_Internal_FFI.``map_``)
let Control_Monad_ST_Internal_modifyImpl = box (Control_Monad_ST_Internal_FFI.``modifyImpl``)
let Control_Monad_ST_Internal_new = box (Control_Monad_ST_Internal_FFI.``new``)
let Control_Monad_ST_Internal_pure_ = box (Control_Monad_ST_Internal_FFI.``pure_``)
let Control_Monad_ST_Internal_read = box (Control_Monad_ST_Internal_FFI.``read``)
let Control_Monad_ST_Internal_run = box (Control_Monad_ST_Internal_FFI.``run``)
let Control_Monad_ST_Internal_while = box (Control_Monad_ST_Internal_FFI.``while``)
let Control_Monad_ST_Internal_write = box (Control_Monad_ST_Internal_FFI.``write``)


let Control_Monad_ST_Internal_modify_prime  = (box Control_Monad_ST_Internal_modifyImpl)

let Control_Monad_ST_Internal_modify  = (box (fun (f: obj) -> (sharpurs_apply (box ((box Control_Monad_ST_Internal_modify_prime))) (box ((box (fun (s: obj) -> (let s_prime = (sharpurs_apply (box ((box f))) (box ((box s)))) in (box ((Map.add "state" (box ((box s_prime))) (Map.add "value" (box ((box s_prime))) Map.empty))))))))))))

let Control_Monad_ST_Internal_functorST  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box Control_Monad_ST_Internal_map_))) Map.empty))))))

let rec Control_Monad_ST_Internal_monadST : obj = ((sharpurs_apply (box ((box Control_Monad_Monadusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Control_Monad_ST_Internal_applicativeST))))) (Map.add "Bind1" (box ((box (fun (usd__unused: obj) -> (box Control_Monad_ST_Internal_bindST))))) Map.empty))))))))
 and Control_Monad_ST_Internal_bindST : obj = ((sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((box Control_Monad_ST_Internal_bind_))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Control_Monad_ST_Internal_applyST))))) Map.empty))))))))
 and Control_Monad_ST_Internal_applyST : obj = ((sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((sharpurs_apply (box ((box Control_Monad_ap))) (box ((box Control_Monad_ST_Internal_monadST)))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Control_Monad_ST_Internal_functorST))))) Map.empty))))))))
 and Control_Monad_ST_Internal_applicativeST : obj = ((sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box Control_Monad_ST_Internal_pure_))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Control_Monad_ST_Internal_applyST))))) Map.empty))))))))


let Control_Monad_ST_Internal_semigroupST  = (box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_lift2))) (box ((box Control_Monad_ST_Internal_applyST)))))) (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup))))))))) Map.empty))))))))

let Control_Monad_ST_Internal_monadRecST  = (sharpurs_apply (box ((box Control_Monad_Rec_Class_MonadRecusd_Dict))) (box ((box ((Map.add "tailRecM" (box ((box (fun (f: obj) -> (box (fun (a: obj) -> (let isLooping = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | Control_Monad_Rec_Class_Loopusd_Ctor(_) -> ((box true)) | _ -> ((box false))))) in let fromDone = (sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (box (fun (v: obj) -> (sharpurs_apply (box ((box (fun (usd__unused: obj) -> (match ((unbox ((box v)))) with | Control_Monad_Rec_Class_Doneusd_Ctor(b) -> ((box b))))))) (box ((box Prim_undefined))))))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Control_Monad_ST_Internal_bindST)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bindFlipped))) (box ((box Control_Monad_ST_Internal_bindST)))))) (box ((box Control_Monad_ST_Internal_new)))))) (box ((sharpurs_apply (box ((box f))) (box ((box a)))))))))))) (box ((box (fun (r: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_discard))) (box ((box Control_Bind_discardUnit)))))) (box ((box Control_Monad_ST_Internal_bindST)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_ST_Internal_while))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Control_Monad_ST_Internal_functorST)))))) (box ((box isLooping)))))) (box ((sharpurs_apply (box ((box Control_Monad_ST_Internal_read))) (box ((box r)))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Control_Monad_ST_Internal_bindST)))))) (box ((sharpurs_apply (box ((box Control_Monad_ST_Internal_read))) (box ((box r))))))))) (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | Control_Monad_Rec_Class_Loopusd_Ctor(a_prime) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Control_Monad_ST_Internal_bindST)))))) (box ((sharpurs_apply (box ((box f))) (box ((box a_prime))))))))) (box ((box (fun (e: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_void))) (box ((box Control_Monad_ST_Internal_functorST)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_ST_Internal_write))) (box ((box e)))))) (box ((box r))))))))))))) | Control_Monad_Rec_Class_Doneusd_Ctor(_) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Control_Monad_ST_Internal_applicativeST)))))) (box ((box Data_Unit_unit))))))))))))))))))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Control_Monad_ST_Internal_functorST)))))) (box ((box fromDone)))))) (box ((sharpurs_apply (box ((box Control_Monad_ST_Internal_read))) (box ((box r)))))))))))))))))))))))) (Map.add "Monad0" (box ((box (fun (usd__unused: obj) -> (box Control_Monad_ST_Internal_monadST))))) Map.empty)))))))

let Control_Monad_ST_Internal_monoidST  = (box (fun (dictMonoid: obj) -> (let semigroupST1 = (sharpurs_apply (box ((box Control_Monad_ST_Internal_semigroupST))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Control_Monad_ST_Internal_applicativeST)))))) (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box dictMonoid))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupST1))))) Map.empty))))))))))
