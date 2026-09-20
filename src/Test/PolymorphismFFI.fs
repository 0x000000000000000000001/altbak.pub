module Test.PolymorphismFFI

type Monoidish<'a> = { mempty: 'a; mappend: 'a -> 'a -> 'a }
let rec polyLoop dictionary count acc =
    if count = 0 then acc
    else polyLoop dictionary (count - 1) (dictionary.mappend acc dictionary.mempty)
let runPolymorphismFFI (n: obj) =
    polyLoop { mempty = 1; mappend = (+) } (unbox<int> n) 0 :> obj
