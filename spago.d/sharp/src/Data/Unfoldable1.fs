module Data.Unfoldable1_FFI

open System

let unfoldr1ArrayImpl (isNothing: obj) (fromJust: obj) (fst: obj) (snd: obj) (f: obj) (b: obj) =
    let isNothing' = isNothing :?> (obj -> obj)
    let fromJust' = fromJust :?> (obj -> obj)
    let fst' = fst :?> (obj -> obj)
    let snd' = snd :?> (obj -> obj)
    let f' = f :?> (obj -> obj)
    
    let res = ResizeArray<obj>()
    let mutable currentB = b
    
    let mutable loop = true
    while loop do
        let tuple = f' currentB
        let a = fst' tuple
        res.Add(a)
        let maybeB = snd' tuple
        let isN = unbox<bool> (isNothing' maybeB)
        if isN then
            loop <- false
        else
            currentB <- fromJust' maybeB
            
    box (res.ToArray())
