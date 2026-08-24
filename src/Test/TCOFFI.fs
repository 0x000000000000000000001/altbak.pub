module Test.TCOFFI
let rec loop n acc = if n = 0 then acc else loop (n - 1) (acc + 1)
let runTCOFFI (n: obj) = loop (unbox<int> n) 0 :> obj
