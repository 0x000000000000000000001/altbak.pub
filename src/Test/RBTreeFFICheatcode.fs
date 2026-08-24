module Test.RBTreeFFICheatcode
let runRBTreeFFICheatcode (n: obj) =
    let n' = unbox<int> n
    let set = System.Collections.Generic.SortedSet<int>()
    for i in 1 .. n' do set.Add(i) |> ignore
    set.Count :> obj
