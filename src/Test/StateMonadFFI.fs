module Test.StateMonadFFI
type StateResult<'s, 'a> = { value: 'a; state: 's }
type State<'s, 'a> = State of ('s -> StateResult<'s, 'a>)
let runState (State f) s = f s
let bind (State f) g = State (fun s -> let result = f s in runState (g result.value) result.state)
let pure' a = State (fun s -> { value = a; state = s })
let get = State (fun s -> { value = s; state = s })
let put s = State (fun _ -> { value = (); state = s })
let modify f = bind get (fun s -> put (f s))
let rec chain n = if n = 0 then pure' () else bind (modify (fun x -> x + 1)) (fun _ -> chain (n - 1))
let runStateMonadFFI (n: obj) =
    let rec loop i acc =
        if i = 0 then acc
        else
            let result = runState (chain (unbox<int> n)) 0
            loop (i - 1) (acc + result.state)
    loop 20 0 :> obj
