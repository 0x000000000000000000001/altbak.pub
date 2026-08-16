module Data.Unfoldable_FFI

open System

let unfoldrArrayImpl (isNothing: obj) (fromJust: obj) (fst: obj) (snd: obj) (f: obj) (b: obj) =
    let isNothing' = isNothing :?> (obj -> obj)
    let fromJust' = fromJust :?> (obj -> obj)
    let fst' = fst :?> (obj -> obj)
    let snd' = snd :?> (obj -> obj)
    let f' = f :?> (obj -> obj)
    
    let res = ResizeArray<obj>()
    let mutable currentB = b
    
    let mutable loop = true
    while loop do
        let maybeTuple = f' currentB
        let isN = unbox<bool> (isNothing' maybeTuple)
        if isN then
            loop <- false
        else
            let tuple = fromJust' maybeTuple
            let a = fst' tuple
            currentB <- snd' tuple
            res.Add(a)
    
    box (res.ToArray())
