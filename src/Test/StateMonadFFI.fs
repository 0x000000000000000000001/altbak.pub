module Test.StateMonadFFI
type State<'s, 'a> = State of ('s -> 'a * 's)
let runState (State f) s = f s
let bind (State f) g = State (fun s -> let a, s' = f s in runState (g a) s')
let pure' a = State (fun s -> a, s)
let get = State (fun s -> s, s)
let put s = State (fun _ -> (), s)
let modify f = bind get (fun s -> put (f s))
let rec chain n = if n = 0 then pure' () else bind (modify (fun x -> x + 1)) (fun _ -> chain (n - 1))
let runStateMonadFFI (n: obj) =
    let rec loop i acc =
        if i = 0 then acc
        else
            let _, s' = runState (chain 60) 0
            loop (i - 1) (acc + s')
    loop (unbox<int> n) 0 :> obj
