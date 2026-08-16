module TestLet_FFI

let opaque (a: obj) =
    box (fun (_: obj) -> a)
